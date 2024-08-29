package d8x_futures

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
)

func TestFetchPricesFromAPI(t *testing.T) {
	priceIds := []PriceId{
		{Id: "0xeaa020c61cc479712813461ce153894a96a6c00b21ed0cfc2798d1f9a9e9c94a",
			Type: PX_PYTH},
		{Id: "0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43",
			Type: PX_PYTH}}
	data, _ := fetchPricesFromAPI(priceIds, "https://hermes.pyth.network/api", "", false)
	fmt.Println(data)
	priceIdsWrong := []PriceId{
		{Id: "0x796d24444ff50728b58e94b1f53dc3a406b2f1ba9d0d0b91d4406c37491a6feb",
			Type: PX_PYTH},
		{Id: "0x01f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722",
			Type: PX_PYTH},
	}
	_, err := fetchPricesFromAPI(priceIdsWrong, "https://hermes.pyth.network/api", "", false)
	if err == nil {
		slog.Error("Error: queried wrong price id but did not fail")
		t.FailNow()
	}
	fmt.Println(err.Error())

}

func CreateProxiedRpc(proxyRawUrl, rpcRawUrl string) (*ethclient.Client, error) {
	proxyURL, err := url.Parse(proxyRawUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proxy URL: %v", err)
	}

	// Create a custom HTTP transport with the proxy
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// Create a custom HTTP client with the transport
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second, // Set a timeout for the HTTP client
	}
	resp, err := httpClient.Get("http://www.google.com")
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	// Use the custom HTTP client to create an RPC client
	rpcClient, err := rpc.DialOptions(context.Background(), rpcRawUrl, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}

	// Create an ethclient.Client instance using the RPC client
	return ethclient.NewClient(rpcClient), nil
}

func loadProxyUrl() string {
	viper.SetConfigFile("../../.env")
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("could not load .env file" + err.Error())
	}
	return viper.GetString("PROXY_URL")
}

func TestCustomRpc(t *testing.T) {
	url := loadProxyUrl()
	if url == "" {
		fmt.Println("need to define a proxy url in .env")
		t.FailNow()
	}
	rpc, err := CreateProxiedRpc(url, "https://rpc.ankr.com/arbitrum")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	var sdkRo SdkRO
	err = sdkRo.New("42161", WithRpcClient(rpc))
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	addr := "0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84"
	amt, err := sdkRo.GetPoolShareTknBalance(1, common.HexToAddress(addr), nil)
	if err != nil {
		t.FailNow()
	}
	fmt.Println("Amount =", amt)
}

func TestFetchPythPrices(t *testing.T) {
	pxConf, err := config.GetDefaultPriceConfig(196)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	r, err := fetchPythPrices([]PriceId{
		{Id: pxConf.PriceFeedIds[0].Id, Type: PX_PYTH},
		{Id: pxConf.PriceFeedIds[1].Id, Type: PX_PYTH}}, "https://hermes.pyth.network/", "")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Print(r)
}

func TestGetPoolShareTknBalance(t *testing.T) {
	var sdkRo SdkRO
	//err := sdkRo.New("195") //xlayer testnet
	err := sdkRo.New("421614",
		WithPriceFeedEndpoint("https://hermes.pyth.network"),
		WithRpcUrl("https://rpc.ankr.com/arbitrum/")) //arbitrum sepolia
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	addr := "0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84"
	amt, err := sdkRo.GetPoolShareTknBalance(1, common.HexToAddress(addr), nil)
	if err != nil {
		t.FailNow()
	}
	fmt.Println("Amount =", amt)
	px, err := sdkRo.GetPoolShareTknPrice([]int{1, 2, 3}, nil)
	if err != nil {
		t.FailNow()
	}
	fmt.Println("prices =", px)
}

func TestSettlementToken(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("42161") //arbitrum
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	addr, err := RawGetSettleTknAddr(&sdkRo.Info, "STUSD")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	mgnAddr, err := RawGetMarginTknAddr(&sdkRo.Info, "STUSD")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	expAddr := "0xaf88d065e77c8cC2239327C5EDb3A432268e5831"
	if strings.EqualFold(addr.Hex(), expAddr) {
		fmt.Println("correct settlement token address for STUSD pool")
	} else {
		fmt.Printf("incorrect settlement token address %s instead of %s for STUSD pool", addr.Hex(), expAddr)
		t.FailNow()
	}
	if strings.EqualFold(addr.Hex(), mgnAddr.Hex()) {
		fmt.Printf("margin token is equal to settle token for STUSD pool")
		t.FailNow()
	}
	addr, err = RawGetSettleTknAddr(&sdkRo.Info, "WEETH")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	mgnAddr, err = RawGetMarginTknAddr(&sdkRo.Info, "USDC")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	if !strings.EqualFold(addr.Hex(), mgnAddr.Hex()) {
		fmt.Println("correct settlement token address for WEETH pool")
	} else {
		fmt.Printf("incorrect settlement token address for WEETH pool")
		t.FailNow()
	}

}

func TestQueryLiquidatableAccounts(t *testing.T) {
	var sdkRo SdkRO
	//err := sdkRo.New("195") //xlayer testnet
	err := sdkRo.New("421614") //arbitrum sepolia
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	acc, err := sdkRo.QueryLiquidatableAccounts(200000, nil)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Println("Accounts =", acc)

	accs, err := RawQueryLiquidatableAccountsInPool(sdkRo.Conn.Rpc, &sdkRo.Info, 1, "https://hermes.pyth.network/api", "")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Print(accs)
	acc2, err := sdkRo.QueryLiquidatableAccountsInPool(2, nil)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Print(acc2)
}

func TestGetPerpetualData(t *testing.T) {
	var sdkRo SdkRO
	//err := sdkRo.New("421614") //arbitrum sepolia
	//err := sdkRo.New("195") //x1
	//err := sdkRo.New("196") //xlayer
	err := sdkRo.New("80084") //bartio
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	startTime := time.Now()
	d, err := RawGetPerpetualData(sdkRo.Conn.Rpc, &sdkRo.Info, "BTC-USDC-USDC")
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	f := utils.ABDKToFloat64(d.FCurrentFundingRate)
	fmt.Printf("current funding rate = %f\n", f)
	fmt.Printf("Time taken: %s\n", elapsedTime)
}

func TestPerpetualPrice(t *testing.T) {
	var sdkRo SdkRO
	//err := sdkRo.New("196")
	err := sdkRo.New("421614") //arbitrum sepolia
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	startTime := time.Now()
	px, err := sdkRo.QueryPerpetualPrices("ETH-USD-WEETH", []float64{0.01}, nil)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Printf("price = %f\n", px)
	fmt.Printf("Time taken: %s\n", elapsedTime)
}

func TestAllowance(t *testing.T) {

	var sdkRo SdkRO
	//err := sdkRo.New("196")
	err := sdkRo.New("42161") //arbitrum
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	a, _, err := sdkRo.Allowance("STUSD", common.HexToAddress("0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84"), nil)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	fmt.Printf("allowance = %f\n", a)
}

func TestPerpetualPriceTuple(t *testing.T) {

	var sdkRo SdkRO
	//err := sdkRo.New("196")
	err := sdkRo.New("421614") //arbitrum sepolia
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	startTime := time.Now()
	tradeAmt := []float64{-0.06, -0.05, -0.01, 0, 0.01, 0.05}
	px, err := RawQueryPerpetualPriceTuple(sdkRo.Conn.Rpc, &sdkRo.Info, sdkRo.ChainConfig.PriceFeedEndpoint, "", "ETH-USD-WEETH", tradeAmt)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Printf("prices = [%f, %f, %f,%f]\n", px[0], px[1], px[2], px[3])
	fmt.Printf("Time taken: %s\n", elapsedTime)
}

func getOrders(sdkRo SdkRO, nodeURL string, from, to int, resultChan chan<- *OpenOrders) {
	rpc, _ := ethclient.Dial(nodeURL)
	orders, err := sdkRo.QueryOpenOrderRange("BTC-USDC-USDC", from, to, rpc) //([]Order, []string, error)
	if err != nil {
		resultChan <- nil
	} else {
		resultChan <- orders
	}
}

func TestMarginAccount(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("196")
	if err != nil {
		t.Log(err.Error())
	}
	addressStrings := []string{"0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84", "0xFF9C956Cd9eB2D27011F79d6A70F62eE6562C4b6", "0xc4C3694DBdCC41475Ebb8d624ddC8acf66d2609d"}
	var addresses []common.Address
	for _, addrStr := range addressStrings {
		address := common.HexToAddress(addrStr)
		addresses = append(addresses, address)
	}
	m, err := RawQueryMarginAccounts(sdkRo.Conn.Rpc, &sdkRo.Info, "WOKB-USD-WOKB", addresses)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	m2, err := sdkRo.QueryMarginAccounts("WOKB-USD-WOKB", addresses, nil)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	println(m[0].FPositionBC)
	println(m2[0].FPositionBC)
	println(m[1].FPositionBC)
	println(m2[1].FPositionBC)
}

func TestSdkROOrders(t *testing.T) {
	var sdkRo SdkRO
	//err := sdkRo.New("x1Testnet")
	err := sdkRo.New("421614") //arb sepolia
	if err != nil {
		t.Log(err.Error())
	}
	startTime := time.Now()
	n, err := sdkRo.QueryNumOrders("BTC-USDC-USDC", nil)
	endTime := time.Now()
	fmt.Printf("Num orders in %s seconds\n", endTime.Sub(startTime))
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Printf("There are %d open orders\n", n)
	startTime = time.Now()
	rpc := []string{"https://x1-testnet.blockpi.network/v1/rpc/public",
		"https://testrpc.x1.tech",
		"https://x1testrpc.okx.com"}
	orderChan := make(chan *OpenOrders, len(rpc))
	num := int(n) / len(rpc)
	for i := 0; i < len(rpc); i++ {
		from := i * num
		to := from + num
		go getOrders(sdkRo, rpc[i], from, to, orderChan)
	}
	var orders = make([]*OpenOrders, 0, len(rpc))
	totalOrders := 0
	for i := 0; i < len(rpc); i++ {
		res := <-orderChan
		if res == nil {
			continue
		}
		orders = append(orders, res)
		totalOrders += len(res.OrderHashes)
	}
	close(orderChan)
	endTime = time.Now()
	fmt.Printf("Found %d orders\n", totalOrders)
	fmt.Printf("in %s seconds\n", endTime.Sub(startTime))
	if len(orders) < 2 {
		fmt.Printf("not enough orders for test")
		return
	}
	k := orders[2].HashIndex[orders[2].OrderHashes[3]]
	if k != 3 {
		t.Logf("hash index test failed")
		t.Fail()
	}
	fmt.Println(orders[0].Orders[0].OptTraderAddr.Hex())
	fmt.Println(orders[0].Orders[1].OptTraderAddr.Hex())
	fmt.Println(orders[0].Orders[2].OptTraderAddr.Hex())
}

func TestFetchCollToSettlePx(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("42161")
	if err != nil {
		t.Log(err.Error())
	}
	perp := "BTC-USD-STUSD"
	px, err := sdkRo.FetchCollToSettlePx(perp, "")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Printf("collateral to settlement conversion STUSD-USDC= %f", px)
	perp = "ETH-USD-WEETH"
	px, err = sdkRo.FetchCollToSettlePx(perp, "")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Printf("collateral to settlement conversion WEETH-WEETH= %f", px)
}

func TestSdkRO(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("42161")
	if err != nil {
		t.Log(err.Error())
	}
	perp := "BTC-USD-STUSD"
	orders, err := sdkRo.QueryAllOpenOrders(perp, nil) //([]Order, []string, error)
	oo := orders.Orders
	dgsts := orders.OrderHashes
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println(oo)
		fmt.Println(dgsts)
	}

	trader := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	broker := common.HexToAddress("0xB0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
	pr, err := sdkRo.GetPositionRisk(perp, trader, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println(pr)
	}
	bal, err := sdkRo.GetSettleTokenBalance(perp, trader, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("balance of margin token=", bal)
	}
	perpState, err := sdkRo.QueryPerpetualState([]int32{100000, 100001, 200002}, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println(perpState)
	}
	poolState, err := sdkRo.QueryPoolStates(nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println(poolState)
	}
	oo, dgsts, err = sdkRo.QueryOpenOrders(perp, trader, nil) //([]Order, []string, error)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println(oo)
		fmt.Println(dgsts)
	}
	id := "258ae021f8743b903d8bde405dba7cc7a74d977ce956db8cb6c2c308976ceb89"
	status, err := sdkRo.QueryOrderStatus(perp, trader, id, nil) // (string, error)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		t.Log(status)
	}
	m, err := sdkRo.QueryMaxTradeAmount(perp, 0, true, nil) // (float64, error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		t.Log(m)
	}
	vol, err := sdkRo.QueryTraderVolume(1, trader, nil) //(float64, error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(vol)
	}
	fee, err := sdkRo.QueryExchangeFeeTbpsForTrader(1, trader, broker, nil) // (uint16, error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fee)
	}
	minpos, err := sdkRo.GetMinimalPositionSize(perp) //(float64, error)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(minpos)
	}

}

func TestFetchPricesForPerpetual(t *testing.T) {
	var sdkRo SdkRO
	err := sdkRo.New("196")
	if err != nil {
		t.Log(err.Error())
	}
	pxBundle, err := RawFetchPricesForPerpetual(sdkRo.Info, "ETH-USDC-USDC", "https://hermes.pyth.network/api", "")
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println(pxBundle)
}

func TestGetPositionRisk(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Log(err.Error())
	}
	pRisk, err := RawGetPositionRisk(info, conn.Rpc, (*common.Address)(&traderAddr), "ETH-USD-MATIC", "https://hermes.pyth.network/api", "")
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println(pRisk)
}

func TestQueryPerpetualState(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConf, nil)
	perpIds := []int32{100001, 100002}
	perpState, err := RawQueryPerpetualState(conn.Rpc, info, perpIds, "https://hermes-beta.pyth.network/api", "")
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println(perpState)
}

func TestQueryPoolStates(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConf, nil)
	poolStates, err := RawQueryPoolStates(conn.Rpc, info)
	if err != nil {
		t.Log(err.Error())
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
		t.Log(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConf, nil)
	orders, digests, err := RawQueryOpenOrders(conn.Rpc, info, "MATIC-USD-MATIC", traderAddr)
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println("--- orders ", orders, "\n---")
	fmt.Println("--- digests ", digests, "\n---")

	// order status
	if len(digests) < 1 {
		return
	}
	d := digests[0]
	status, _ := RawQueryOrderStatus(conn.Rpc, info, traderAddr, d, "MATIC-USD-MATIC")
	fmt.Println("order status: ", status)
}

func TestQueryTraderVolume(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConf, nil)
	volume, err := RawQueryTraderVolume(conn.Rpc, info, traderAddr, 1)
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println("Trader volume = ", volume)
}

func TestQueryExchangeFeeTbpsForTrader(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	brokerAddr := common.Address{}
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConf, nil)
	fee, err := RawQueryExchangeFeeTbpsForTrader(conn.Rpc, info, 1, traderAddr, brokerAddr)
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println("Fee Tbps = ", fee)
}

func TestQueryMaxTradeAmount(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(1442)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConf, nil)
	trade, err := RawQueryMaxTradeAmount(conn.Rpc, info, 0.01, "ETH-USD-MATIC", true)
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Println("Max Trade buy (0.01 ETH-USD-MATIC) = ", trade)
}
