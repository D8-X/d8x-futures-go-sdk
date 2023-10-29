package d8x_futures

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

func GetPositionRisk(xInfo StaticExchangeInfo, conn BlockChainConnector, traderAddr *common.Address, symbol string, epNo int) (PositionRisk, error) {
	priceData, err := FetchPricesForPerpetual(xInfo, symbol, epNo)
	if err != nil {
		return PositionRisk{}, err
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		panic("symbol does not exist in static perpetual info")
	}
	indexPrices := [2]*big.Int{utils.Float64ToABDK(priceData.S2Price), utils.Float64ToABDK(priceData.S3Price)}

	traderState, err := conn.PerpetualManager.GetTraderState(
		nil,
		new(big.Int).SetInt64(int64(xInfo.Perpetuals[j].Id)),
		*traderAddr,
		indexPrices)
	if err != nil {
		fmt.Println("Error fetching margin account:", err)
		return PositionRisk{}, err
	}
	const idxAvailableCashCC = 2
	const idxCash = 3
	const idxNotional = 4
	const idxLockedIn = 5
	const idxMarkPrice = 8
	const idxLvg = 7
	const idxS3 = 9
	var S3Liq float64
	lockedInValue := utils.ABDKToFloat64(traderState[idxLockedIn])
	cashCC := utils.ABDKToFloat64(traderState[idxCash])
	posBC := utils.ABDKToFloat64(traderState[idxNotional])
	Sm := utils.ABDKToFloat64(traderState[idxMarkPrice])
	S3 := utils.ABDKToFloat64(traderState[idxS3])
	unpaidFundingCC := utils.ABDKToFloat64(traderState[idxAvailableCashCC]) - cashCC
	unpaidFundingQC := unpaidFundingCC
	S2Liq := CalculateLiquidationPrice(xInfo.Perpetuals[j].CollateralCurrencyType, lockedInValue, posBC, cashCC, xInfo.Perpetuals[j].MaintenanceMarginRate, S3, Sm)
	if xInfo.Perpetuals[j].CollateralCurrencyType == BASE {
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / priceData.S2Price
	} else if xInfo.Perpetuals[j].CollateralCurrencyType == QUANTO {
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / priceData.S3Price
		S3Liq = utils.ABDKToFloat64(traderState[idxS3])
	}
	// floor at zero
	S2Liq = math.Max(0, S2Liq)
	S3Liq = math.Max(0, S3Liq)

	var side string
	if posBC == 0 {
		side = SIDE_CLOSED
	} else if lockedInValue < 0 {
		side = SIDE_SELL
	} else {
		side = SIDE_BUY
	}
	var entryPrice, leverage, pnl, liqLeverage float64
	if posBC != 0 {
		entryPrice = math.Abs(lockedInValue / posBC)
		leverage = utils.ABDKToFloat64(traderState[idxLvg])
		pnl = posBC*Sm - lockedInValue + unpaidFundingQC
		liqLeverage = 1 / xInfo.Perpetuals[j].MaintenanceMarginRate
	} else {
		S3Liq = 0
	}
	m := PositionRisk{
		Symbol:                         symbol,
		PositionNotionalBaseCCY:        math.Abs(posBC),
		Side:                           side,
		EntryPrice:                     entryPrice,
		Leverage:                       leverage,
		MarkPrice:                      Sm,
		UnrealizedPnlQuoteCCY:          pnl,
		UnrealizedFundingCollateralCCY: unpaidFundingCC,
		CollateralCC:                   cashCC,
		LiquidationPrice:               [2]float64{S2Liq, S3Liq},
		LiquidationLvg:                 liqLeverage,
		CollToQuoteConversion:          S3,
	}
	return m, nil
}

// QueryPerpetualState collects PerpetualState by calling the off-chain prices and
// blockchain queries. epNo is the endpoint number to choose (see priceFeedConfig)
func QueryPerpetualState(conn BlockChainConnector, xInfo StaticExchangeInfo, perpetualIds []int32, epNo int) ([]PerpetualState, error) {
	bigIntSlice := make([]*big.Int, len(perpetualIds))
	for i, id := range perpetualIds {
		bigIntSlice[i] = big.NewInt(int64(id))
	}
	// perpetual data via blockchain
	perps, err := conn.PerpetualManager.GetPerpetuals(nil, bigIntSlice)
	if err != nil {
		return []PerpetualState{}, err
	}
	// gather perpetual index prices (offchain REST)
	pxInfo := make([]*big.Int, len(perpetualIds)*2)
	pxInfoFloat := make([]float64, len(perpetualIds)*2)
	for i := range perpetualIds {
		p, err := FetchPricesForPerpetualId(xInfo, perpetualIds[i], epNo)
		if err != nil {
			return nil, err
		}
		pxInfoFloat[i*2] = p.S2Price
		pxInfoFloat[i*2+1] = p.S3Price
		pxInfo[i*2] = utils.Float64ToABDK(p.S2Price)
		pxInfo[i*2+1] = utils.Float64ToABDK(p.S3Price)
	}
	// midprices via blockchain query
	pxMid, err := conn.PerpetualManager.QueryMidPrices(nil, bigIntSlice, pxInfo)
	if err != nil {
		return []PerpetualState{}, err
	}

	perpStates := make([]PerpetualState, len(perps))
	for i, perpData := range perps {
		perpStates[i].Id = int32(perpData.Id.Int64())
		perpStates[i].State = PerpetualStateEnum(perpData.State)
		perpStates[i].IndexPrice = pxInfoFloat[i*2]
		perpStates[i].CollToQuoteIndexPrice = pxInfoFloat[i*2+1]
		perpStates[i].MarkPrice = pxInfoFloat[i*2] * (1 + utils.ABDKToFloat64(perpData.CurrentMarkPremiumRate.FPrice))
		perpStates[i].CurrentFundingRateBps = utils.ABDKToFloat64(perpData.FCurrentFundingRate)
		perpStates[i].OpenInterestBC = utils.ABDKToFloat64(perpData.FOpenInterest)
		perpStates[i].MidPrice = utils.ABDKToFloat64(pxMid[i])
	}
	return perpStates, nil
}

// QueryPoolStates gathers the pool states of all pools by querying the blockchain in
// chunks of 10 pools
func QueryPoolStates(conn BlockChainConnector, xInfo StaticExchangeInfo) ([]PoolState, error) {
	numPools := len(xInfo.Pools)
	poolStates := make([]PoolState, numPools)
	// we query a maximum of 10 pools at once
	const MAXPOOLS = 10
	iterations := (numPools + MAXPOOLS - 1) / MAXPOOLS
	for i := 0; i < iterations; i++ {
		from := i * MAXPOOLS
		to := (i+1)*MAXPOOLS - 1
		if to > numPools {
			to = numPools
		}
		pools, err := conn.PerpetualManager.GetLiquidityPools(nil, uint8(from+1), uint8(to+1))
		if err != nil {
			return []PoolState{}, nil
		}
		pIdx := 0
		for j := from; j < to; j++ {
			poolStates[j].Id = int32(pools[pIdx].Id)
			poolStates[j].DefaultFundCashCC = utils.ABDKToFloat64(pools[pIdx].FDefaultFundCashCC)
			poolStates[j].IsRunning = pools[pIdx].IsRunning
			poolStates[j].PnlParticipantCashCC = utils.ABDKToFloat64(pools[pIdx].FPnLparticipantsCashCC)
			poolStates[j].TotalSupplyShareToken = utils.ABDKToFloat64(pools[pIdx].TotalSupplyShareToken)
			poolStates[j].TotalTargetAMMFundSizeCC = utils.ABDKToFloat64(pools[pIdx].FTargetAMMFundSize)
			pIdx++
		}
	}
	return poolStates, nil
}

func QueryOpenOrders(conn BlockChainConnector, xInfo StaticExchangeInfo, symbol string, traderAddr common.Address) ([]Order, []string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return []Order{}, []string{}, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(conn.Rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)

	from := 0
	count := 15
	clientOrders := []contracts.IClientOrderClientOrder{}
	zeroAddr := common.Address{}
outerLoop:
	for {
		currOrders, err := lob.GetOrders(nil, traderAddr, big.NewInt(int64(from)), big.NewInt(int64(count)))
		if err != nil {
			return []Order{}, []string{}, err
		}
		from = from + count
		for _, order := range currOrders {
			if order.TraderAddr == zeroAddr {
				break outerLoop
			} else {
				clientOrders = append(clientOrders, order)
			}
		}
	}
	digests, err := lob.LimitDigestsOfTrader(
		nil,
		traderAddr,
		big.NewInt(int64(0)),
		big.NewInt(int64(len(clientOrders))),
	)
	// format digests into strings
	strDigests := make([]string, len(digests))
	// format ClientOrders into orders
	orders := make([]Order, len(clientOrders))
	for i, d := range digests {
		strDigests[i] = "0x" + common.Bytes2Hex(d[:])
		orders[i] = FromChainType(&clientOrders[i], xInfo)
	}
	if err != nil {
		return []Order{}, []string{}, err
	}
	return orders, strDigests, nil
}

func QueryOrderStatus(conn BlockChainConnector, xInfo StaticExchangeInfo, traderAddr common.Address, orderDigest string, symbol string) (string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return "", fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(conn.Rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	// convert digest string to bytes 32
	bytesDigest := common.Hex2Bytes(strings.TrimPrefix(orderDigest, "0x"))
	var orderDigest32 [32]byte
	copy(orderDigest32[:], bytesDigest)
	status, err := lob.GetOrderStatus(nil, orderDigest32)
	if err != nil {
		return "", err
	}
	var statusStr string
	switch status {
	case ENUM_ORDER_STATUS_CANCELED:
		statusStr = ORDER_STATUS_CANCELED
	case ENUM_ORDER_STATUS_EXECUTED:
		statusStr = ORDER_STATUS_EXECUTED
	case ENUM_ORDER_STATUS_OPEN:
		statusStr = ORDER_STATUS_OPEN
	case ENUM_ORDER_STATUS_UNKNOWN:
		statusStr = ORDER_STATUS_UNKNOWN
	}
	return statusStr, nil
}

func QueryMaxTradeAmount(conn BlockChainConnector, xInfo StaticExchangeInfo, currentPositionNotional float64, symbol string, isBuy bool) (float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return 0, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	p := utils.Float64ToABDK(currentPositionNotional)
	t, err := conn.PerpetualManager.GetMaxSignedOpenTradeSizeForPos(nil, big.NewInt(int64(xInfo.Perpetuals[j].Id)), p, isBuy)
	if err != nil {
		return 0, err
	}
	return utils.ABDKToFloat64(t), nil
}

func QueryTraderVolume(conn BlockChainConnector, xInfo StaticExchangeInfo, traderAddr common.Address, poolId int32) (float64, error) {
	vol, err := conn.PerpetualManager.GetCurrentTraderVolume(nil, uint8(poolId), traderAddr)
	if err != nil {
		return 0, err
	}
	return utils.ABDKToFloat64(vol), nil
}

func QueryExchangeFeeTbpsForTrader(conn BlockChainConnector, xInfo StaticExchangeInfo, poolId int32, traderAddr common.Address, brokerAddr common.Address) (uint16, error) {
	feeTbps, err := conn.PerpetualManager.QueryExchangeFee(nil, uint8(poolId), traderAddr, brokerAddr)
	if err != nil {
		return 0, err
	}
	return feeTbps, nil
}

func GetMinimalPositionSize(perp PerpetualStaticInfo) float64 {
	return 10 * perp.LotSizeBC
}

func CalculateLiquidationPrice(ccy CollateralCCY, lockedInValue float64, positionBC float64, cashCC float64, tau float64, S3 float64, Sm float64) float64 {
	if positionBC == 0 {
		return float64(0)
	}
	if ccy == QUANTO {
		numerator := -lockedInValue + cashCC*S3*(1-utils.Sign(positionBC))
		denominator := tau*math.Abs(positionBC) - positionBC - (cashCC*S3*utils.Sign(positionBC))/Sm
		return numerator / denominator
	} else if ccy == BASE {
		numerator := -lockedInValue + cashCC
		denominator := tau*math.Abs(positionBC) - positionBC
		return numerator / denominator
	} else {
		return lockedInValue / (positionBC - tau*math.Abs(positionBC) + cashCC)
	}
}

func FetchPricesForPerpetualId(exchangeInfo StaticExchangeInfo, id int32, epNo int) (PerpetualPriceInfo, error) {
	j := GetPerpetualStaticInfoIdxFromId(exchangeInfo, id)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPricesForPerpetual(exchangeInfo, j, epNo)
}

// FetchPricesForPerpetual queries the REST-endpoints of the oracles and calculates S2,S3
// index prices, also returns the price-feed-data required for blockchain submission and
// information whether the market is closed or not. epNo is the endpoint number from the config.
func FetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, symbol string, epNo int) (PerpetualPriceInfo, error) {

	j := GetPerpetualStaticInfoIdxFromSymbol(exchangeInfo, symbol)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPricesForPerpetual(exchangeInfo, j, epNo)
}

// fetchPricesForPerpetual gets prices from the VAA-endpoint, perpetual number j
// and endpoint number epNo
func fetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, j int, epNo int) (PerpetualPriceInfo, error) {
	// get underlying data from rest-api
	feedData, err := fetchPricesFromAPI(exchangeInfo.Perpetuals[j].PriceIds, exchangeInfo.PriceFeedInfo, epNo)
	if err != nil {
		return PerpetualPriceInfo{}, err
	}
	// triangulate fetched prices to obtain index prices
	triang := exchangeInfo.IdxPriceTriangulations[exchangeInfo.Perpetuals[j].S2Symbol]
	pxS2, isMarketClosedS2 := CalculateTriangulation(triang, feedData)
	var pxS3 float64 = 0
	var isMarketClosedS3 bool = false
	// triangulate S3 if there is an S3 index
	if exchangeInfo.Perpetuals[j].S3Symbol != "" {
		triang = exchangeInfo.IdxPriceTriangulations[exchangeInfo.Perpetuals[j].S3Symbol]
		pxS3, isMarketClosedS3 = CalculateTriangulation(triang, feedData)
	}
	var priceInfo = PerpetualPriceInfo{
		S2Price:          pxS2,
		S3Price:          pxS3,
		IsMarketClosedS2: isMarketClosedS2,
		IsMarketClosedS3: isMarketClosedS3,
		PriceFeed:        feedData,
	}
	return priceInfo, nil
}

// fetchPricesFromAPI gets the prices for the given priceIds from the
// configured REST-API. The PriceFeedConfig is needed to extract the
// correct endpoint per feed, and store what symbol (e.g. BTC-USD) the
// price feed covers. endPtNo is the endpoint number in the config
func fetchPricesFromAPI(priceIds []string, config utils.PriceFeedConfig, endPtNo int) (PriceFeedData, error) {
	pxData := PriceFeedData{
		Symbols:        make([]string, len(priceIds)),
		PriceIds:       priceIds,
		Prices:         make([]float64, len(priceIds)),
		IsMarketClosed: make([]bool, len(priceIds)),
		Vaas:           make([]string, len(priceIds)),
	}
	queries := make(map[string]string)
	for _, el := range config.EndPoints {
		queries[el.Type] = el.EndpointUrl[endPtNo] + "/latest_price_feeds?target_chain=default&"
	}
	// loop through price id's, find its endpoints and prepare the query
	for i, id := range priceIds {
		id = strings.TrimPrefix(id, "0x")
		for _, c := range config.PriceFeedIds {
			if c.Id == "0x"+id {
				queries[c.Type] += "ids[]=" + id + "&"
				pxData.Symbols[i] = c.Symbol
				break
			}
		}
	}
	timestampNow := time.Now().Unix()
	// REST query (#queries == number of endpoints for feeds)
	for _, q := range queries {
		response, err := http.Get(q)
		if err != nil {
			fmt.Println("Error sending fetchPricesFromAPI request:", err)
			return PriceFeedData{}, err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return PriceFeedData{}, err
		}

		var data []ResponsePythLatestPriceFeed
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return PriceFeedData{}, err
		}
		// process data
		for _, d := range data {
			//find idx of d.Id
			for i, id := range priceIds {
				if id == d.Id {
					pxData.Prices[i] = utils.PythNToFloat64(d.Price.Price, d.Price.Expo)
					pxData.IsMarketClosed[i] = timestampNow-int64(d.Price.PublishTime) > int64(config.ThresholdMarketClosedSec)
					pxData.Vaas[i] = d.Vaa
					break
				}
			}
		}
	}
	return pxData, nil
}
