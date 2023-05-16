package d8x_futures

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func GetPositionRisk(xInfo StaticExchangeInfo, conn BlockChainConnector, traderAddr *common.Address, symbol string) (PositionRisk, error) {
	priceData := FetchPricesForPerpetual(xInfo, symbol)
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		panic("symbol does not exist in static perpetual info")
	}
	indexPrices := [2]*big.Int{Float64ToABDK(priceData.S2Price), Float64ToABDK(priceData.S3Price)}

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
	lockedInValue := ABDKToFloat64(traderState[idxLockedIn])
	cashCC := ABDKToFloat64(traderState[idxCash])
	posBC := ABDKToFloat64(traderState[idxNotional])
	Sm := ABDKToFloat64(traderState[idxMarkPrice])
	S3 := ABDKToFloat64(traderState[idxS3])
	unpaidFundingCC := ABDKToFloat64(traderState[idxAvailableCashCC]) - cashCC
	unpaidFundingQC := unpaidFundingCC
	S2Liq := CalculateLiquidationPrice(xInfo.Perpetuals[j].CollateralCurrencyType, lockedInValue, posBC, cashCC, xInfo.Perpetuals[j].MaintenanceMarginRate, S3, Sm)
	if xInfo.Perpetuals[j].CollateralCurrencyType == BASE {
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / priceData.S2Price
	} else if xInfo.Perpetuals[j].CollateralCurrencyType == QUANTO {
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / priceData.S3Price
		S3Liq = ABDKToFloat64(traderState[idxS3])
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
		leverage = ABDKToFloat64(traderState[idxLvg])
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

func GetMinimalPositionSize(perp PerpetualStaticInfo) float64 {
	return 10 * perp.LotSizeBC
}

func CalculateLiquidationPrice(ccy CollateralCCY, lockedInValue float64, positionBC float64, cashCC float64, tau float64, S3 float64, Sm float64) float64 {
	if positionBC == 0 {
		return float64(0)
	}
	if ccy == QUANTO {
		numerator := -lockedInValue + cashCC*S3*(1-sign(positionBC))
		denominator := tau*math.Abs(positionBC) - positionBC - (cashCC*S3*sign(positionBC))/Sm
		return numerator / denominator
	} else if ccy == BASE {
		numerator := -lockedInValue + cashCC
		denominator := tau*math.Abs(positionBC) - positionBC
		return numerator / denominator
	} else {
		return lockedInValue / (positionBC - tau*math.Abs(positionBC) + cashCC)
	}
}

// FetchPricesForPerpetual queries the REST-endpoints of the oracles and calculates S2,S3
// index prices, also returns the price-feed-data required for blockchain submission and
// information whether the market is closed or not.
func FetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, symbol string) PerpetualPriceInfo {

	j := GetPerpetualStaticInfoIdxFromSymbol(exchangeInfo, symbol)
	if j == -1 {
		panic("symbol does not exist in static perpetual info")
	}
	// get underlying data from rest-api
	feedData := fetchPricesFromAPI(exchangeInfo.Perpetuals[j].PriceIds, exchangeInfo.PriceFeedInfo)
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
	return priceInfo
}

// fetchPricesFromAPI gets the prices for the given priceIds from the
// configured REST-API. The PriceFeedConfig is needed to extract the
// correct endpoint per feed, and store what symbol (e.g. BTC-USD) the
// price feed covers.
func fetchPricesFromAPI(priceIds []string, config PriceFeedConfig) PriceFeedData {
	pxData := PriceFeedData{
		Symbols:        make([]string, len(priceIds)),
		PriceIds:       priceIds,
		Prices:         make([]float64, len(priceIds)),
		IsMarketClosed: make([]bool, len(priceIds)),
		Vaas:           make([]string, len(priceIds)),
	}
	queries := make(map[string]string)
	for _, el := range config.EndPoints {
		queries[el.Type] = el.EndpointUrl + "/latest_price_feeds?target_chain=default&"
	}
	// loop through price id's, find its endpoints and prepare the query
	for i, id := range priceIds {
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
			continue
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		var data []ResponsePythLatestPriceFeed
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}
		// process data
		for _, d := range data {
			//find idx of d.Id
			for i, id := range priceIds {
				if id == d.Id {
					pxData.Prices[i] = PythNToFloat64(d.Price.Price, d.Price.Expo)
					pxData.IsMarketClosed[i] = timestampNow-int64(d.Price.PublishTime) > int64(config.ThresholdMarketClosedSec)
					pxData.Vaas[i] = d.Vaa
					break
				}
			}
		}
	}
	return pxData
}
