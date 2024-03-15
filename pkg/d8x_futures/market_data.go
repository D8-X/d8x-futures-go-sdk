package d8x_futures

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/go-multicall"
)

const MARGIN_ACCOUNT_ABI = `[{
	"inputs": [
		{
			"internalType": "uint24",
			"name": "_perpetualId",
			"type": "uint24"
		},
		{
			"internalType": "address",
			"name": "_traderAddress",
			"type": "address"
		}
	],
	"name": "getMarginAccount",
	"outputs": [
				{
					"internalType": "int128",
					"name": "fLockedInValueQC",
					"type": "int128"
				},
				{
					"internalType": "int128",
					"name": "fCashCC",
					"type": "int128"
				},
				{
					"internalType": "int128",
					"name": "fPositionBC",
					"type": "int128"
				},
				{
					"internalType": "int128",
					"name": "fUnitAccumulatedFundingStart",
					"type": "int128"
				},
				{
					"internalType": "uint64",
					"name": "iLastOpenTimestamp",
					"type": "uint64"
				},
				{
					"internalType": "uint16",
					"name": "feeTbps",
					"type": "uint16"
				},
				{
					"internalType": "uint16",
					"name": "brokerFeeTbps",
					"type": "uint16"
				},
				{
					"internalType": "bytes16",
					"name": "positionId",
					"type": "bytes16"
				}],
		"stateMutability": "view",
		"type": "function"}]`

const QUERY_PERP_PX_ABI = `[{
			"inputs": [
				{
					"internalType": "uint24",
					"name": "_iPerpetualId",
					"type": "uint24"
				},
				{
					"internalType": "int128",
					"name": "_fTradeAmountBC",
					"type": "int128"
				},
				{
					"internalType": "int128[2]",
					"name": "_fIndexPrice",
					"type": "int128[2]"
				}
			],
			"name": "queryPerpetualPrice",
			"outputs": [
				{
					"internalType": "int128",
					"name": "",
					"type": "int128"
				}
			],
			"stateMutability": "view",
			"type": "function"
		}]`

func (sdkRo *SdkRO) GetPositionRisk(symbol string, traderAddr common.Address, optRpc *ethclient.Client) (PositionRisk, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawGetPositionRisk(sdkRo.Info, optRpc, &traderAddr, symbol, sdkRo.ChainConfig.PriceFeedEndpoints[0])
}

func (sdkRo *SdkRO) QueryPerpetualState(perpetualIds []int32, optRpc *ethclient.Client) ([]PerpetualState, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryPerpetualState(optRpc, sdkRo.Info, perpetualIds, sdkRo.ChainConfig.PriceFeedEndpoints[0])
}

func (sdkRo *SdkRO) QueryPoolStates(optRpc *ethclient.Client) ([]PoolState, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryPoolStates(optRpc, sdkRo.Info)
}

func (sdkRo *SdkRO) QueryPerpetualPrices(symbol string, tradeAmt []float64, optRpc *ethclient.Client) ([]float64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryPerpetualPriceTuple(optRpc, &sdkRo.Info, sdkRo.ChainConfig.PriceFeedEndpoints[0], symbol, tradeAmt)
}

func (sdkRo *SdkRO) QueryOpenOrders(symbol string, traderAddr common.Address, optRpc *ethclient.Client) ([]Order, []string, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryOpenOrders(optRpc, sdkRo.Info, symbol, traderAddr)
}

func (sdkRo *SdkRO) QueryAllOpenOrders(symbol string, optRpc *ethclient.Client) (*OpenOrders, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryAllOpenOrders(optRpc, sdkRo.Info, symbol)
}

func (sdkRo *SdkRO) QueryNumOrders(symbol string, optRpc *ethclient.Client) (int64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryNumOrders(optRpc, sdkRo.Info, symbol)
}

func (sdkRo *SdkRO) QueryOpenOrderRange(symbol string, from, to int, optRpc *ethclient.Client) (*OpenOrders, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryOpenOrderRange(optRpc, sdkRo.Info, symbol, from, to)
}

func (sdkRo *SdkRO) QueryOrderStatus(symbol string, traderAddr common.Address, orderDigest string, optRpc *ethclient.Client) (string, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryOrderStatus(optRpc, sdkRo.Info, traderAddr, orderDigest, symbol)
}

func (sdkRo *SdkRO) QueryMaxTradeAmount(symbol string, currentPositionNotional float64, isBuy bool, optRpc *ethclient.Client) (float64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryMaxTradeAmount(optRpc, sdkRo.Info, currentPositionNotional, symbol, isBuy)
}

func (sdkRo *SdkRO) QueryMarginAccounts(symbol string, traderAddrs []common.Address, optRpc *ethclient.Client) ([]MarginAccount, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryMarginAccounts(optRpc, &sdkRo.Info, symbol, traderAddrs)
}

func (sdkRo *SdkRO) QueryTraderVolume(poolId int32, traderAddr common.Address, optRpc *ethclient.Client) (float64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryTraderVolume(optRpc, sdkRo.Info, traderAddr, poolId)
}

func (sdkRo *SdkRO) QueryExchangeFeeTbpsForTrader(poolId int32, traderAddr common.Address, brokerAddr common.Address, optRpc *ethclient.Client) (uint16, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryExchangeFeeTbpsForTrader(optRpc, sdkRo.Info, poolId, traderAddr, brokerAddr)
}

func (sdkRo *SdkRO) GetMinimalPositionSize(symbol string) (float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdkRo.Info, symbol)
	if j == -1 {
		return 0, errors.New("GetMinimalPositionSize: no perpetual " + symbol)
	}
	v := sdkRo.Info.Perpetuals[j]
	return v.LotSizeBC * 10.0, nil
}

func (sdkRo *SdkRO) FetchPricesForPerpetualId(id int32) (PerpetualPriceInfo, error) {
	return RawFetchPricesForPerpetualId(sdkRo.Info, id, sdkRo.ChainConfig.PriceFeedEndpoints[0])
}

func (sdkRo *SdkRO) FetchPricesForPerpetual(symbol string, endpoint string) (PerpetualPriceInfo, error) {
	return RawFetchPricesForPerpetual(sdkRo.Info, symbol, sdkRo.ChainConfig.PriceFeedEndpoints[0])
}

func (sdkRo *SdkRO) GetMarginTokenBalance(symbol string, traderAddr common.Address, optRpc *ethclient.Client) (float64, error) {
	tknAddr, err := RawGetMarginTknAddr(&sdkRo.Info, symbol)
	if err != nil {
		return 0, err
	}
	rpc := sdkRo.Conn.Rpc
	if optRpc != nil {
		rpc = optRpc
	}
	erc20Instance, err := contracts.NewErc20(tknAddr, rpc)
	if err != nil {
		return 0, errors.New("GetMarginTokenBalance: creating instance of token " + tknAddr.String())
	}
	n, _ := erc20Instance.Decimals(nil)
	b, _ := erc20Instance.BalanceOf(nil, traderAddr)
	bal := utils.DecNToFloat(b, n)
	return bal, nil
}

func RawGetMarginTknAddr(xInfo *StaticExchangeInfo, symbol string) (common.Address, error) {
	j := GetPoolStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return common.Address{}, errors.New("RawGetMarginTknAddr: no perpetual " + symbol)
	}
	pool := xInfo.Pools[j]
	poolId := pool.PoolId

	return (xInfo.Pools[poolId-1].PoolMarginTokenAddr), nil
}

func RawGetPositionRisk(xInfo StaticExchangeInfo, rpc *ethclient.Client, traderAddr *common.Address, symbol string, endpoint string) (PositionRisk, error) {
	priceData, err := RawFetchPricesForPerpetual(xInfo, symbol, endpoint)
	if err != nil {
		return PositionRisk{}, err
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		panic("symbol does not exist in static perpetual info")
	}
	indexPrices := [2]*big.Int{utils.Float64ToABDK(priceData.S2Price), utils.Float64ToABDK(priceData.S3Price)}

	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)

	traderState, err := proxy.GetTraderState(
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
	S2Liq := RawCalculateLiquidationPrice(xInfo.Perpetuals[j].CollateralCurrencyType, lockedInValue, posBC, cashCC, xInfo.Perpetuals[j].MaintenanceMarginRate, S3, Sm)
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
	} else if posBC < 0 {
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
// blockchain queries. endpoint is the address to get prices from
func RawQueryPerpetualState(rpc *ethclient.Client, xInfo StaticExchangeInfo, perpetualIds []int32, endpoint string) ([]PerpetualState, error) {
	bigIntSlice := make([]*big.Int, len(perpetualIds))
	for i, id := range perpetualIds {
		bigIntSlice[i] = big.NewInt(int64(id))
	}
	// perpetual data via blockchain

	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	perps, err := proxy.GetPerpetuals(nil, bigIntSlice)
	if err != nil {
		return []PerpetualState{}, err
	}
	// gather perpetual index prices (offchain REST)
	pxInfo := make([]*big.Int, len(perpetualIds)*2)
	pxInfoFloat := make([]float64, len(perpetualIds)*2)
	for i := range perpetualIds {
		p, err := RawFetchPricesForPerpetualId(xInfo, perpetualIds[i], endpoint)
		if err != nil {
			return nil, err
		}
		pxInfoFloat[i*2] = p.S2Price
		pxInfoFloat[i*2+1] = p.S3Price
		pxInfo[i*2] = utils.Float64ToABDK(p.S2Price)
		pxInfo[i*2+1] = utils.Float64ToABDK(p.S3Price)
	}
	// midprices via blockchain query
	pxMid, err := proxy.QueryMidPrices(nil, bigIntSlice, pxInfo)
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

// RawQueryPoolStates gathers the pool states of all pools by querying the blockchain in
// chunks of 10 pools
func RawQueryPoolStates(rpc *ethclient.Client, xInfo StaticExchangeInfo) ([]PoolState, error) {
	numPools := len(xInfo.Pools)
	poolStates := make([]PoolState, numPools)
	// we query a maximum of 10 pools at once
	const MAXPOOLS = 10
	iterations := (numPools + MAXPOOLS - 1) / MAXPOOLS
	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	for i := 0; i < iterations; i++ {
		from := i * MAXPOOLS
		to := (i+1)*MAXPOOLS - 1
		if to > numPools {
			to = numPools
		}

		pools, err := proxy.GetLiquidityPools(nil, uint8(from+1), uint8(to+1))
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

func RawQueryAllOpenOrders(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string) (*OpenOrders, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)

	from := 0
	count := 500
	var orders OpenOrders
	orders.HashIndex = make(map[string]int)
	zeroAddr := common.Address{}
outerLoop:
	for {
		currOrders, err := lob.PollRange(nil, big.NewInt(int64(from)), big.NewInt(int64(count)))
		if err != nil {
			return nil, err
		}
		from = from + count
		for k, corder := range currOrders.Orders {
			if corder.TraderAddr == zeroAddr {
				break outerLoop
			} else {
				order := FromChainType(&corder, &xInfo)
				orders.Orders = append(orders.Orders, order)
				strDigests := "0x" + common.Bytes2Hex(currOrders.OrderHashes[k][:])
				orders.HashIndex[strDigests] = k
				orders.OrderHashes = append(orders.OrderHashes, strDigests)
			}
		}
		orders.SubmittedTs = currOrders.SubmittedTs
		from = from + count
	}

	return &orders, nil
}

func RawQueryNumOrders(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string) (int64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return 0, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	count, err := lob.OrderCount(nil)
	if err != nil {
		return 0, err
	}
	return count.Int64(), nil
}

func RawQueryOpenOrderRange(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string, from, to int) (*OpenOrders, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	ooSc, err := lob.PollRange(nil, big.NewInt(int64(from)), big.NewInt(int64(to-from)))
	var orders OpenOrders
	orders.HashIndex = make(map[string]int)
	for k, scOrder := range ooSc.Orders {
		orders.Orders = append(orders.Orders, FromChainType(&scOrder, &xInfo))
		strDigests := "0x" + common.Bytes2Hex(ooSc.OrderHashes[k][:])
		orders.HashIndex[strDigests] = k
		orders.OrderHashes = append(orders.OrderHashes, strDigests)
	}
	orders.SubmittedTs = ooSc.SubmittedTs
	return &orders, err
}

func RawQueryOpenOrders(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string, traderAddr common.Address) ([]Order, []string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return []Order{}, []string{}, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)

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
		orders[i] = FromChainType(&clientOrders[i], &xInfo)
	}
	if err != nil {
		return []Order{}, []string{}, err
	}
	return orders, strDigests, nil
}

func RawQueryOrderStatus(rpc *ethclient.Client, xInfo StaticExchangeInfo, traderAddr common.Address, orderDigest string, symbol string) (string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return "", fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
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

func RawQueryPerpetualPriceTuple(client *ethclient.Client, xInfo *StaticExchangeInfo, pythEndpoint, symbol string, tradeAmt []float64) ([]float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
	pxFeed, err := fetchPricesForPerpetual(*xInfo, j, pythEndpoint)
	if err != nil {
		return nil, errors.New("RawAddCollateral: failed fetching oracle prices " + err.Error())
	}
	pricesAbdk := [2]*big.Int{utils.Float64ToABDK(pxFeed.S2Price), utils.Float64ToABDK(pxFeed.S3Price)}

	caller, err := multicall.New(client)
	if err != nil {
		return nil, err
	}
	type priceOutput struct {
		PriceAbdk *big.Int
	}

	contract, err := multicall.NewContract(QUERY_PERP_PX_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	calls := make([]*multicall.Call, 0, len(tradeAmt))
	for _, amt := range tradeAmt {
		taAbdk := utils.Float64ToABDK(amt)
		c := contract.NewCall(new(priceOutput), "queryPerpetualPrice", perpId, taAbdk, pricesAbdk)
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	prices := make([]float64, 0, len(tradeAmt))
	for _, call := range res {
		px := utils.ABDKToFloat64(call.Outputs.(*priceOutput).PriceAbdk)
		prices = append(prices, px)
	}
	return prices, nil
}

func RawQueryMarginAccounts(client *ethclient.Client, xInfo *StaticExchangeInfo, symbol string, traders []common.Address) ([]MarginAccount, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
	caller, err := multicall.New(client)
	if err != nil {
		return nil, err
	}

	type marginOutput struct {
		FLockedInValueQC             *big.Int
		FCashCC                      *big.Int
		FPositionBC                  *big.Int
		FUnitAccumulatedFundingStart *big.Int
		ILastOpenTimestamp           uint64
		FeeTbps                      uint16
		BrokerFeeTbps                uint16
		PositionId                   [16]byte
	}
	contract, err := multicall.NewContract(MARGIN_ACCOUNT_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	calls := make([]*multicall.Call, 0, len(traders))

	for _, t := range traders {
		c := contract.NewCall(new(marginOutput), "getMarginAccount", perpId, t)
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	accounts := make([]MarginAccount, 0, len(traders))
	for _, call := range res {
		m := MarginAccount{
			FLockedInValueQC:             utils.ABDKToFloat64(call.Outputs.(*marginOutput).FLockedInValueQC),
			FCashCC:                      utils.ABDKToFloat64(call.Outputs.(*marginOutput).FCashCC),
			FPositionBC:                  utils.ABDKToFloat64(call.Outputs.(*marginOutput).FPositionBC),
			FUnitAccumulatedFundingStart: utils.ABDKToFloat64(call.Outputs.(*marginOutput).FUnitAccumulatedFundingStart),
		}
		accounts = append(accounts, m)
	}
	return accounts, nil
}

// Not compatible with zkevm deployment
func RawGetPerpetualData(rpc *ethclient.Client, xInfo *StaticExchangeInfo, symbol string) (*contracts.PerpStoragePerpetualData, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	perpData, err := proxy.GetPerpetual(nil, big.NewInt(int64(xInfo.Perpetuals[j].Id)))
	if err != nil {
		return nil, err
	}
	return &perpData, nil
}

func RawQueryMaxTradeAmount(rpc *ethclient.Client, xInfo StaticExchangeInfo, currentPositionNotional float64, symbol string, isBuy bool) (float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return 0, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	p := utils.Float64ToABDK(currentPositionNotional)
	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	t, err := proxy.GetMaxSignedOpenTradeSizeForPos(nil, big.NewInt(int64(xInfo.Perpetuals[j].Id)), p, isBuy)
	if err != nil {
		return 0, err
	}
	return utils.ABDKToFloat64(t), nil
}

func RawQueryTraderVolume(rpc *ethclient.Client, xInfo StaticExchangeInfo, traderAddr common.Address, poolId int32) (float64, error) {
	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	vol, err := proxy.GetCurrentTraderVolume(nil, uint8(poolId), traderAddr)
	if err != nil {
		return 0, err
	}
	return utils.ABDKToFloat64(vol), nil
}

func RawQueryExchangeFeeTbpsForTrader(rpc *ethclient.Client, xInfo StaticExchangeInfo, poolId int32, traderAddr common.Address, brokerAddr common.Address) (uint16, error) {
	proxy := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	feeTbps, err := proxy.QueryExchangeFee(nil, uint8(poolId), traderAddr, brokerAddr)
	if err != nil {
		return 0, err
	}
	return feeTbps, nil
}

func RawCalculateLiquidationPrice(ccy CollateralCCY, lockedInValue float64, positionBC float64, cashCC float64, tau float64, S3 float64, Sm float64) float64 {
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

func RawFetchPricesForPerpetualId(exchangeInfo StaticExchangeInfo, id int32, endpoint string) (PerpetualPriceInfo, error) {
	j := GetPerpetualStaticInfoIdxFromId(&exchangeInfo, id)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPricesForPerpetual(exchangeInfo, j, endpoint)
}

// FetchPricesForPerpetual queries the REST-endpoints of the oracles and calculates S2,S3
// index prices, also returns the price-feed-data required for blockchain submission and
// information whether the market is closed or not. endpoint is the endpoint that provides pyth prices.
func RawFetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, symbol string, endpoint string) (PerpetualPriceInfo, error) {

	j := GetPerpetualStaticInfoIdxFromSymbol(&exchangeInfo, symbol)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPricesForPerpetual(exchangeInfo, j, endpoint)
}

// fetchPricesForPerpetual gets prices from the VAA-endpoint, perpetual number j
// and endpoint number epNo. j is the index of the perpetual in StaticExchangeInfo
func fetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, j int, endpoint string) (PerpetualPriceInfo, error) {
	// get underlying data from rest-api
	feedData, err := fetchPricesFromAPI(exchangeInfo.Perpetuals[j].PriceIds, exchangeInfo.PriceFeedInfo, endpoint)
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
func fetchPricesFromAPI(priceIds []string, config utils.PriceFeedConfig, priceFeedEndpoint string) (PriceFeedData, error) {
	pxData := PriceFeedData{
		Symbols:        make([]string, len(priceIds)),
		PriceIds:       priceIds,
		Prices:         make([]float64, len(priceIds)),
		IsMarketClosed: make([]bool, len(priceIds)),
		Vaas:           make([][]byte, len(priceIds)),
		PublishTimes:   make([]uint64, len(priceIds)),
	}
	// binary=true <- include VAA
	query := priceFeedEndpoint + "/latest_price_feeds?binary=true&"
	// loop through price id's, find its endpoints and prepare the query
	for i, id := range priceIds {
		priceIds[i] = strings.TrimPrefix(id, "0x")
		pxData.Symbols[i] = "not in config"
		query += "ids[]=" + priceIds[i] + "&"
		for _, c := range config.PriceFeedIds {
			if c.Id == "0x"+priceIds[i] {
				pxData.Symbols[i] = c.Symbol
				break
			}
		}
	}
	query = strings.TrimSuffix(query, "&")
	timestampNow := time.Now().Unix()
	// REST query (#queries == number of endpoints for feeds)
	response, err := http.Get(query)
	if err != nil {
		err := errors.New("Error sending fetchPricesFromAPI request:" + err.Error())
		return PriceFeedData{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		err := errors.New("Error reading response body:" + err.Error())
		return PriceFeedData{}, err
	}
	if response.StatusCode != 200 {
		err := errors.New("Error fetchPricesFromAPI status " + strconv.Itoa(response.StatusCode) + " " + string(body[:]))
		return PriceFeedData{}, err
	}
	var data []ResponsePythLatestPriceFeed
	err = json.Unmarshal(body, &data)
	if err != nil {
		err := errors.New("fetchPricesFromAPI:" + err.Error())
		return PriceFeedData{}, err
	}
	// process data
	for _, d := range data {
		//find idx of d.Id
		for i, id := range priceIds {
			if id == d.Id {
				pxData.Prices[i] = utils.PythNToFloat64(d.Price.Price, d.Price.Expo)
				pxData.IsMarketClosed[i] = timestampNow-int64(d.Price.PublishTime) > int64(config.ThresholdMarketClosedSec)
				decodedVaaBytes, err := base64.StdEncoding.DecodeString(d.Vaa)
				if err != nil {
					err := errors.New("fetchPricesFromAPI decoding base64:" + err.Error())
					return PriceFeedData{}, err
				}
				pxData.Vaas[i] = decodedVaaBytes
				pxData.PublishTimes[i] = uint64(d.Price.PublishTime)
				break
			}
		}
	}
	return pxData, nil
}
