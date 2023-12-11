package d8x_futures

import (
	"fmt"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/ethereum/go-ethereum/common"
)

func TestFetchPricesFromAPI(t *testing.T) {
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	priceIds := []string{"0x796d24444ff50728b58e94b1f53dc3a406b2f1ba9d0d0b91d4406c37491a6feb",
		"0x41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722"}
	data, _ := fetchPricesFromAPI(priceIds, pxConf, "https://hermes-beta.pyth.network/api")
	fmt.Println(data)
}

func TestFetchPricesForPerpetual(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	pxBundle, _ := FetchPricesForPerpetual(info, "BTC-USD-MATIC", "https://hermes-beta.pyth.network/api")
	fmt.Println(pxBundle)
}

func TestGetPositionRisk(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	pRisk, err := GetPositionRisk(info, conn, (*common.Address)(&traderAddr), "ETH-USD-MATIC", "https://hermes-beta.pyth.network/api")
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println(pRisk)
}

func TestQueryPerpetualState(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	perpIds := []int32{100001, 100002}
	perpState, err := QueryPerpetualState(conn, info, perpIds, "https://hermes-beta.pyth.network/api")
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println(perpState)
}

func TestQueryPoolStates(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
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
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
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
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	volume, err := QueryTraderVolume(conn, info, traderAddr, 1)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Trader volume = ", volume)
}

func TestQueryExchangeFeeTbpsForTrader(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	brokerAddr := common.Address{}
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	fee, err := QueryExchangeFeeTbpsForTrader(conn, info, 1, traderAddr, brokerAddr)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Fee Tbps = ", fee)
}

func TestQueryMaxTradeAmount(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig("PythEVMBeta")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	trade, err := QueryMaxTradeAmount(conn, info, 0.01, "ETH-USD-MATIC", true)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Max Trade buy (0.01 ETH-USD-MATIC) = ", trade)
}
