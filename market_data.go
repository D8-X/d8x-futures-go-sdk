package d8x_futures

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// func PositionRisk(proxyContract IPerpetualManager, traderAddr string, symbol string) MarginAccount {

//     let obj = await this.priceFeedGetter.fetchPricesForPerpetual(symbol);
//     let mgnAcct = await PerpetualDataHandler.getMarginAccount(
//       traderAddr,
//       symbol,
//       this.symbolToPerpStaticInfo,
//       this.proxyContract,
//       [obj.idxPrices[0], obj.idxPrices[1]]
//     );
//     return mgnAcct;
//   }

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
