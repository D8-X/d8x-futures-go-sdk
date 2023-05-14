package d8x_futures

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestFetchPricesFromAPI(t *testing.T) {
	pxConfig, err := LoadPriceFeedConfig("testnet")
	if err != nil {
		panic(err)
	}
	priceIds := []string{"0x796d24444ff50728b58e94b1f53dc3a406b2f1ba9d0d0b91d4406c37491a6feb",
		"0x41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722"}
	data := fetchPricesFromAPI(priceIds, pxConfig)
	fmt.Println(data)
}

func TestFetchPricesForPerpetual(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	pxBundle := FetchPricesForPerpetual(info, "BTC-USD-MATIC")
	fmt.Println(pxBundle)
}

func TestGetPositionRisk(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("***REMOVED***")
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	pRisk, err := GetPositionRisk(info, conn, (*common.Address)(&traderAddr), "ETH-USD-MATIC")
	if err != nil {
		panic(err)
	}
	fmt.Println(pRisk)
}
