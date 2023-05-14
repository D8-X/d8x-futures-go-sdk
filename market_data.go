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

func fetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, symbol string) {
	/*
		perpId := exchangeInfo.PerpetualSymbolToId[symbol]
		var symS2, symS3 string
		for _, p := range exchangeInfo.Perpetuals {
			if p.Id == perpId {
				symS2 = p.S2Symbol
				symS3 = p.S3Symbol
				break
			}
		}
		/*
			triang1 := exchangeInfo.IdxPriceTriangulations[symS2]
			triang2 := []TriangulationElement{}
			if symS3 != "" {
				triang2 := exchangeInfo.IdxPriceTriangulations[symS2]
			}

			/*pxS2 := fetchPriceFromTriangulation(triang, exchangeInfo.PriceFeedInfo)
			var pxS3 float64 = 0.0

			return pxS2, pxS3
	*/
}

// fetchPricesFromAPI get the prices for the given priceIds from the
// configured REST-API.
func fetchPricesFromAPI(priceIds []string, config PriceFeedConfig) PriceFeedData {
	pxData := PriceFeedData{
		Symbols:        make([]string, len(priceIds)),
		PriceIds:       priceIds,
		Prices:         make([]float64, len(priceIds)),
		IsMarketClosed: make([]bool, len(priceIds)),
	}
	queries := make(map[string]string)
	for _, el := range config.EndPoints {
		queries[el.Type] = el.EndpointUrl + "/latest_price_feeds?target_chain=default&"
	}
	// loop through price id's, find its endpoints and prepare the query
	for i, id := range priceIds {
		for _, c := range config.PriceFeedIds {
			if c.Id == id {
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
				if id == "0x"+d.Id {
					pxData.Prices[i] = PythNToFloat64(d.Price.Price, d.Price.Expo)
					pxData.IsMarketClosed[i] = timestampNow-int64(d.Price.PublishTime) > int64(config.ThresholdMarketClosedSec)
					break
				}
			}
		}
	}
	return pxData
}
