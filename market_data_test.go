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

func TestQueryPerpetualState(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	perpIds := []int32{100001, 100002}
	perpState, err := QueryPerpetualState(conn, info, perpIds)
	if err != nil {
		panic(err)
	}
	fmt.Println(perpState)
}

func TestQueryPoolStates(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	poolStates, err := QueryPoolStates(conn, info)
	if err != nil {
		panic(err)
	}
	for _, p := range poolStates {
		fmt.Println("--- Pool ", p.Id, "---")
		fmt.Println("")
		fmt.Println(p.IsRunning)
		fmt.Println("DefaultFundCashCC=", p.DefaultFundCashCC)
		fmt.Println("PnlParticipantCashCC=", p.PnlParticipantCashCC)
		fmt.Println("TotalAMMFundCashCC=", p.TotalAMMFundCashCC)
		fmt.Println("TotalTargetAMMFundSizeCC=", p.TotalTargetAMMFundSizeCC)
	}

}
func TestQueryOpenOrders(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("***REMOVED***")
	conn := CreateBlockChainConnector(config)
	orders, digests, err := QueryOpenOrders(conn, info, "MATIC-USD-MATIC", traderAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("--- orders ", orders, "\n---")
	fmt.Println("--- digests ", digests, "\n---")

	// order status
	if len(digests) < 1 {
		return
	}
	d := digests[0]
	status, err := QueryOrderStatus(conn, info, traderAddr, d, "MATIC-USD-MATIC")
	fmt.Println("order status: ", status)
}
