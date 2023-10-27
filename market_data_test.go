package d8x_futures

import (
	"fmt"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

func TestFetchPricesFromAPI(t *testing.T) {
	pxConfig, err := utils.LoadPriceFeedConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	priceIds := []string{"0x796d24444ff50728b58e94b1f53dc3a406b2f1ba9d0d0b91d4406c37491a6feb",
		"0x41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722"}
	data, _ := fetchPricesFromAPI(priceIds, pxConfig, 0)
	fmt.Println(data)
}

func TestFetchPricesForPerpetual(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	pxBundle, _ := FetchPricesForPerpetual(info, "BTC-USD-MATIC", 0)
	fmt.Println(pxBundle)
}

func TestGetPositionRisk(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	pRisk, err := GetPositionRisk(info, conn, (*common.Address)(&traderAddr), "ETH-USD-MATIC", 0)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println(pRisk)
}

func TestQueryPerpetualState(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	perpIds := []int32{100001, 100002}
	perpState, err := QueryPerpetualState(conn, info, perpIds, 0)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println(perpState)
}

func TestQueryPoolStates(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	poolStates, err := QueryPoolStates(conn, info)
	if err != nil {
		t.Logf(err.Error())
	}
	for _, p := range poolStates {
		fmt.Println("--- Pool ", p.Id, "---")
		fmt.Println("")
		fmt.Println(p.IsRunning)
		fmt.Println("DefaultFundCashCC=", p.DefaultFundCashCC)
		fmt.Println("PnlParticipantCashCC=", p.PnlParticipantCashCC)
		fmt.Println("TotalSupplyShareToken=", p.TotalSupplyShareToken)
		fmt.Println("TotalTargetAMMFundSizeCC=", p.TotalTargetAMMFundSizeCC)
	}

}
func TestQueryOpenOrders(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	conn := CreateBlockChainConnector(config)
	orders, digests, err := QueryOpenOrders(conn, info, "MATIC-USD-MATIC", traderAddr)
	if err != nil {
		t.Logf(err.Error())
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

func TestQueryTraderVolume(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	conn := CreateBlockChainConnector(config)
	volume, err := QueryTraderVolume(conn, info, traderAddr, 1)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Trader volume = ", volume)
}

func TestQueryExchangeFeeTbpsForTrader(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	brokerAddr := common.Address{}
	conn := CreateBlockChainConnector(config)
	fee, err := QueryExchangeFeeTbpsForTrader(conn, info, 1, traderAddr, brokerAddr)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Fee Tbps = ", fee)
}

func TestQueryMaxTradeAmount(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	trade, err := QueryMaxTradeAmount(conn, info, 0.01, "ETH-USD-MATIC", true)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Max Trade buy (0.01 ETH-USD-MATIC) = ", trade)
}
