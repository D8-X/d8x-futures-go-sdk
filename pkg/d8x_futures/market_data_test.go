package d8x_futures

import (
	"fmt"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
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
	priceIdsWrong := []string{"0x796d24444ff50728b58e94b1f53dc3a406b2f1ba9d0d0b91d4406c37491a6feb",
		"0x01f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722"}
	data, err = fetchPricesFromAPI(priceIdsWrong, pxConf, "https://hermes-beta.pyth.network/api")
	fmt.Println(err.Error())

}

func TestGetPerpetualData(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("x1Testnet")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	startTime := time.Now()
	d, err := RawGetPerpetualData(sdkRo.Conn, &sdkRo.Info, "BTC-USDC-USDC")
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	f := utils.ABDKToFloat64(d.FCurrentFundingRate)
	fmt.Printf("current funding rate = %f\n", f)
	fmt.Printf("Time taken: %s\n", elapsedTime)
}

func TestPerpetualPrice(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("x1Testnet")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	startTime := time.Now()
	px, err := sdkRo.QueryPerpetualPrice("BTC-USDC-USDC", 0.01)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("price = %f\n", px)
	fmt.Printf("Time taken: %s\n", elapsedTime)
}

func TestPerpetualPriceTuple(t *testing.T) {

	var sdkRo SdkRO
	err := sdkRo.New("x1Testnet")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	startTime := time.Now()
	tradeAmt := []float64{-0.01, 0, 0.01, 0.05}
	px, err := RawQueryPerpetualPriceTuple(&sdkRo.Conn, &sdkRo.Info, "https://testrpc.x1.tech", sdkRo.ChainConfig.PriceFeedEndpoints[0], "BTC-USDC-USDC", tradeAmt)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("prices = [%f, %f, %f,%f]\n", px[0], px[1], px[2], px[3])
	fmt.Printf("Time taken: %s\n", elapsedTime)
}

func TestSdkROOrders(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("x1Testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	n, err := sdkRo.QueryNumOrders("BTC-USDC-USDC")
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Printf("There are %d open orders", n)
	orders, err := sdkRo.QueryOpenOrderRange("BTC-USDC-USDC", 0, 10) //([]Order, []string, error)
	oo := orders.Orders
	dgsts := orders.OrderHashes
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println(oo)
		fmt.Println(dgsts)
	}
}

func TestSdkRO(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("x1Testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	orders, err := sdkRo.QueryAllOpenOrders("BTC-USDC-USDC") //([]Order, []string, error)
	oo := orders.Orders
	dgsts := orders.OrderHashes
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println(oo)
		fmt.Println(dgsts)
	}

	trader := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	broker := common.HexToAddress("0xB0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
	pr, err := sdkRo.GetPositionRisk("BTC-USDC-USDC", trader)
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println(pr)
	}
	bal, err := sdkRo.GetMarginTokenBalance("BTC-USDC-USDC", trader)
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println("balance of margin token=", bal)
	}
	perpState, err := sdkRo.QueryPerpetualState([]int32{100000, 100001, 200002})
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println(perpState)
	}
	poolState, err := sdkRo.QueryPoolStates()
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println(poolState)
	}
	oo, dgsts, err = sdkRo.QueryOpenOrders("BTC-USD-MATIC", trader) //([]Order, []string, error)
	if err != nil {
		t.Logf(err.Error())
	} else {
		fmt.Println(oo)
		fmt.Println(dgsts)
	}
	id := "258ae021f8743b903d8bde405dba7cc7a74d977ce956db8cb6c2c308976ceb89"
	status, err := sdkRo.QueryOrderStatus("BTC-USDC-USDC", trader, id) // (string, error)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		t.Log(status)
	}
	m, err := sdkRo.QueryMaxTradeAmount("BTC-USD-MATIC", 0, true) // (float64, error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		t.Log(m)
	}
	vol, err := sdkRo.QueryTraderVolume(1, trader) //(float64, error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(vol)
	}
	fee, err := sdkRo.QueryExchangeFeeTbpsForTrader(1, trader, broker) // (uint16, error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fee)
	}
	minpos, err := sdkRo.GetMinimalPositionSize("BTC-USD-MATIC") //(float64, error)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(minpos)
	}

}

func TestFetchPricesForPerpetual(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	pxBundle, _ := RawFetchPricesForPerpetual(info, "BTC-USD-MATIC", "https://hermes-beta.pyth.network/api")
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
	pRisk, err := RawGetPositionRisk(info, conn, (*common.Address)(&traderAddr), "ETH-USD-MATIC", "https://hermes-beta.pyth.network/api")
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
	perpState, err := RawQueryPerpetualState(conn, info, perpIds, "https://hermes-beta.pyth.network/api")
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
	poolStates, err := RawQueryPoolStates(conn, info)
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
	orders, digests, err := RawQueryOpenOrders(conn, info, "MATIC-USD-MATIC", traderAddr)
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
	status, err := RawQueryOrderStatus(conn, info, traderAddr, d, "MATIC-USD-MATIC")
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
	volume, err := RawQueryTraderVolume(conn, info, traderAddr, 1)
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
	fee, err := RawQueryExchangeFeeTbpsForTrader(conn, info, 1, traderAddr, brokerAddr)
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
	trade, err := RawQueryMaxTradeAmount(conn, info, 0.01, "ETH-USD-MATIC", true)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println("Max Trade buy (0.01 ETH-USD-MATIC) = ", trade)
}
