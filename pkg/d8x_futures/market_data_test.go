//go:build integration

package d8x_futures

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestFetchPricesFromAPI(t *testing.T) {
	priceIds := []PriceId{
		{
			Id:   "0xeaa020c61cc479712813461ce153894a96a6c00b21ed0cfc2798d1f9a9e9c94a",
			Type: utils.PXTYPE_PYTH, Origin: "",
		},
		{
			Id:   "0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43",
			Type: utils.PXTYPE_PYTH, Origin: "",
		},
	}
	pxEp := PriceFeedEndpoints{PriceFeedEndpoint: "https://hermes.pyth.network/api"}
	data, err := fetchPricesFromAPI(priceIds, pxEp, false)
	if err != nil {
		t.Fatalf("fetchPricesFromAPI: %v", err)
	}
	t.Log(data)
	priceIdsWrong := []PriceId{
		{
			Id:   "0x796d24444ff50728b58e94b1f53dc3a406b2f1ba9d0d0b91d4406c37491a6feb",
			Type: utils.PXTYPE_PYTH, Origin: "",
		},
		{
			Id:   "0x01f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722",
			Type: utils.PXTYPE_PYTH, Origin: "",
		},
	}
	_, err = fetchPricesFromAPI(priceIdsWrong, pxEp, false)
	if err == nil {
		t.Fatal("Error: queried wrong price id but did not fail")
	}
	t.Log(err.Error())
}

func CreateProxiedRpc(proxyRawUrl, rpcRawUrl string) (*ethclient.Client, error) {
	proxyURL, err := url.Parse(proxyRawUrl)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	// Create an ethclient.Client instance using the RPC client
	return ethclient.NewClient(rpcClient), nil
}

func loadProxyUrl() string {
	return os.Getenv("PROXY_URL")
}

func TestCustomRpc(t *testing.T) {
	proxyUrl := loadProxyUrl()
	if proxyUrl == "" {
		t.Skip("need to define PROXY_URL in .env")
	}
	rpcClient, err := CreateProxiedRpc(proxyUrl, "https://rpc.ankr.com/arbitrum")
	if err != nil {
		t.Fatalf("CreateProxiedRpc: %v", err)
	}

	sdkRo, err := NewSdkRO("84532", WithRpcClient(rpcClient))
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	addr := "0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84"
	amt, err := sdkRo.GetPoolShareTknBalance(1, common.HexToAddress(addr), nil)
	if err != nil {
		t.Fatalf("GetPoolShareTknBalance: %v", err)
	}
	t.Log("Amount =", amt)
}

func TestFetchPythPrices(t *testing.T) {
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	idx := make([]int, 0, 2)
	for j := range pxConf.PriceFeedIds {
		if pxConf.PriceFeedIds[j].Type == utils.PXTYPE_PYTH {
			idx = append(idx, j)
		}
		if len(idx) == 2 {
			break
		}
	}
	pxEp := PriceFeedEndpoints{PriceFeedEndpoint: "https://hermes.pyth.network/"}
	r, err := fetchOraclePrices([]PriceId{
		{Id: pxConf.PriceFeedIds[idx[0]].Id, Type: utils.PXTYPE_PYTH, Origin: ""},
		{Id: pxConf.PriceFeedIds[idx[1]].Id, Type: utils.PXTYPE_PYTH, Origin: ""},
	}, pxEp)
	if err != nil {
		t.Fatalf("fetchOraclePrices: %v", err)
	}
	t.Log(r)
}

func TestFetchPolymarket(t *testing.T) {
	priceIds := []string{
		"0x6710f215fe01867219fc338d0a68290f84c471d002b14b6decd92c2260d94cce",
		"0x2f06f5a323466c1ad9a9a8afb19a21dd3fe0f39853a16ae58637086e8ff5838d",
		"0x7eea0b94fc1916efdd031978918abba566b9aef7f12b3d5a5b0f7141fd8e0b33",
		"0x6ec8114a0dba99a037c7122c93a68f2ab905c3e42e2b82e1ce659aef96de58f2",
	}
	for _, id := range priceIds {
		idDec, _ := utils.Hex2Dec(id)
		reqUrl := "https://clob.polymarket.com/midpoint?token_id=" + idDec
		response, err := http.Get(reqUrl)
		if err != nil {
			t.Fatalf("http.Get: %v", err)
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		t.Log(string(body))
		response.Body.Close()
	}
	originIds := []string{
		"0x76dbb81a9fd937efa736aa23e6c0eb33aaf0f2ca4aa6c06da1d5529ed236ebfb",
		"0x265366ede72d73e137b2b9095a6cdc9be6149290caa295738a95e3d881ad0865",
		"0x67e68c5eee8ac767dd1177de8c653b20642fee48f0f2a56d784e4856b130749d",
		"0x1ab07117f9f698f28490f57754d6fe5309374230c95867a7eba572892a11d710",
	}
	for _, id := range originIds {
		reqUrl := "https://clob.polymarket.com/markets/" + id
		response, err := http.Get(reqUrl)
		if err != nil {
			t.Fatalf("http.Get: %v", err)
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		t.Log(string(body))
		t.Log("----")
		response.Body.Close()
	}
}

func TestFetchLowLiqPx(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	id, err := sdkRo.GetPriceId(sym)
	if err != nil {
		t.Fatalf("GetPriceId: %v", err)
	}
	t.Log(id.Origin)
	for _, info := range sdkRo.Info.Perpetuals {
		t.Logf("id: %d perpetual %s-'%s'", info.Id, info.S2Symbol, info.S3Symbol)
	}
	b, err := sdkRo.IsLowLiqPerp(sym)
	if err != nil {
		t.Fatalf("IsLowLiqPerp: %v", err)
	}
	t.Log(b)
	px, err := sdkRo.QueryPerpetualPrices(sym, []float64{0}, nil)
	if err != nil {
		t.Fatalf("QueryPerpetualPrices: %v", err)
	}
	t.Log(px)
}

func TestFetchInfo(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	// Find a prediction market perpetual dynamically
	var prdSym string
	for _, p := range sdkRo.Info.Perpetuals {
		if p.State() != NORMAL {
			continue
		}
		sym, ok := sdkRo.Info.PerpetualIdToSymbol[p.Id]
		if !ok {
			continue
		}
		isPrd, _ := sdkRo.IsPrdMktPerp(sym)
		if isPrd {
			prdSym = sym
			break
		}
	}
	if prdSym == "" {
		t.Skip("no prediction market perpetual found")
	}
	// Extract S2-quote symbol (e.g., "TRUMP24-USD" from "TRUMP24-USD-USDC")
	parts := strings.Split(prdSym, "-")
	s2Sym := parts[0] + "-" + parts[1]
	id, err := sdkRo.GetPriceId(s2Sym)
	if err != nil {
		t.Fatalf("GetPriceId: %v", err)
	}
	t.Log(id.Origin)
	reqUrl := "https://clob.polymarket.com/markets/" + id.Origin
	response, err := http.Get(reqUrl)
	if err != nil {
		t.Fatalf("http.Get: %v", err)
	}
	defer response.Body.Close()

	t.Log(response)
	type ApiResponse struct {
		EndDateISO string `json:"end_date_iso"`
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
}

func TestGetPoolShareTknBalance(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	if len(sdkRo.Info.Pools) == 0 {
		t.Skip("no pools found")
	}
	poolId := sdkRo.Info.Pools[0].PoolId
	addr := "0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84"
	amt, err := sdkRo.GetPoolShareTknBalance(int(poolId), common.HexToAddress(addr), nil)
	if err != nil {
		t.Fatalf("GetPoolShareTknBalance: %v", err)
	}
	t.Log("Amount =", amt)
	var poolIds []int
	for _, p := range sdkRo.Info.Pools {
		poolIds = append(poolIds, int(p.PoolId))
	}
	px, err := sdkRo.GetPoolShareTknPrice(poolIds, nil)
	if err != nil {
		t.Fatalf("GetPoolShareTknPrice: %v", err)
	}
	t.Log("prices =", px)
}

func TestSettlementToken(t *testing.T) {
	// STUSD and WEETH are arbitrum-specific pools
	sdkRo, err := NewSdkRO("42161")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	addr, err := RawGetSettleTknAddr(&sdkRo.Info, "STUSD")
	if err != nil {
		t.Skipf("STUSD pool not found (may not be on this chain): %v", err)
	}
	mgnAddr, err := RawGetMarginTknAddr(&sdkRo.Info, "STUSD")
	if err != nil {
		t.Fatalf("RawGetMarginTknAddr STUSD: %v", err)
	}
	expAddr := "0xaf88d065e77c8cC2239327C5EDb3A432268e5831"
	if strings.EqualFold(addr.Hex(), expAddr) {
		t.Log("correct settlement token address for STUSD pool")
	} else {
		t.Fatalf("incorrect settlement token address %s instead of %s for STUSD pool", addr.Hex(), expAddr)
	}
	if strings.EqualFold(addr.Hex(), mgnAddr.Hex()) {
		t.Fatalf("margin token is equal to settle token for STUSD pool")
	}
	addr, err = RawGetSettleTknAddr(&sdkRo.Info, "WEETH")
	if err != nil {
		t.Skipf("WEETH pool not found: %v", err)
	}
	mgnAddr, err = RawGetMarginTknAddr(&sdkRo.Info, "USDC")
	if err != nil {
		t.Fatalf("RawGetMarginTknAddr USDC: %v", err)
	}
	if !strings.EqualFold(addr.Hex(), mgnAddr.Hex()) {
		t.Log("correct settlement token address for WEETH pool")
	} else {
		t.Fatalf("incorrect settlement token address for WEETH pool")
	}
}

func TestQueryLiquidatableAccounts(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	acc, err := sdkRo.QueryLiquidatableAccounts(100000, nil)
	if err != nil {
		t.Fatalf("QueryLiquidatableAccounts: %v", err)
	}
	t.Log("Accounts =", acc)

	/*accs, err := RawQueryLiquidatableAccountsInPool(sdkRo.Conn.Rpc, &sdkRo.Info, 1, PriceFeedEndpoints{PriceFeedEndpoint: "https://hermes.pyth.network/api"})
	if err != nil {
		t.Fatalf("RawQueryLiquidatableAccountsInPool: %v", err)
	}
	t.Log(accs)
	acc2, err := sdkRo.QueryLiquidatableAccountsInPool(2, nil)
	if err != nil {
		t.Fatalf("QueryLiquidatableAccountsInPool: %v", err)
	}
	t.Log(acc2)*/
}

func TestGetPerpetualData(t *testing.T) {
	// err := sdkRo.New("421614") //arbitrum sepolia
	// err := sdkRo.New("195") //x1
	// err := sdkRo.New("196") //xlayer
	// err := sdkRo.New("80084") //bartio
	sdkRo, err := NewSdkRO("84532") // base sepolia
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	for _, p := range sdkRo.Info.Perpetuals {
		t.Logf("%d s2=%s s3=%s state=%s", p.Id, p.S2Symbol, p.S3Symbol, p.State().String())
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	startTime := time.Now()
	d, err := RawGetPerpetualData(sdkRo.Conn.Rpc, &sdkRo.Info, sym)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Fatalf("RawGetPerpetualData: %v", err)
	}
	f := utils.ABDKToFloat64(d.FCurrentFundingRate)
	t.Logf("current funding rate = %f", f)
	t.Logf("Time taken: %s", elapsedTime)
}

func TestPerpetualPrice(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	startTime := time.Now()
	px, err := sdkRo.QueryPerpetualPrices(sym, []float64{200, 2000}, nil)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Fatalf("QueryPerpetualPrices: %v", err)
	}
	t.Logf("price = %f", px)
	t.Logf("Time taken: %s", elapsedTime)
}

func TestAllowance(t *testing.T) {
	sdkRo, err := NewSdkRO("42161") // arbitrum
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	a, _, err := sdkRo.Allowance("STUSD", common.HexToAddress("0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84"), nil)
	if err != nil {
		t.Fatalf("Allowance: %v", err)
	}
	t.Logf("allowance = %f", a)
}

func TestPerpetualPriceTuple(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	startTime := time.Now()
	tradeAmt := []float64{-0.06, -0.05, -0.01, 0, 0.01, 0.05}
	ep := PriceFeedEndpoints{
		PriceFeedEndpoint:  sdkRo.ChainConfig.PriceFeedEndpoint,
		LowLiqFeedEndpoint: sdkRo.ChainConfig.LowLiqFeedEndpoint,
		PrdMktFeedEndpoint: sdkRo.ChainConfig.PrdMktFeedEndpoint,
		SportFeedEndpoint:  sdkRo.ChainConfig.SportFeedEndpoint,
	}
	px, err := RawQueryPerpetualPriceTuple(sdkRo.Conn.Rpc, &sdkRo.Info, ep, sym, tradeAmt)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	if err != nil {
		t.Fatalf("RawQueryPerpetualPriceTuple: %v", err)
	}
	t.Logf("prices = [%f, %f, %f,%f]", px[0], px[1], px[2], px[3])
	t.Logf("Time taken: %s", elapsedTime)
}

func getOrders(sdkRo SdkRO, sym string, nodeURL string, from, to int, resultChan chan<- *OpenOrders) {
	rpcClient, _ := ethclient.Dial(nodeURL)
	orders, err := sdkRo.QueryOpenOrderRange(sym, from, to, rpcClient) //([]Order, []string, error)
	if err != nil {
		resultChan <- nil
	} else {
		resultChan <- orders
	}
}

func TestMarginAccount(t *testing.T) {
	sdkRo, err := NewSdkRO("84532") // base_sepolia
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	addressStrings := []string{"0xdef43CF2Dd024abc5447C1Dcdc2fE3FE58547b84", "0xFF9C956Cd9eB2D27011F79d6A70F62eE6562C4b6", "0xc4C3694DBdCC41475Ebb8d624ddC8acf66d2609d"}
	var addresses []common.Address
	for _, addrStr := range addressStrings {
		address := common.HexToAddress(addrStr)
		addresses = append(addresses, address)
	}
	m, err := RawQueryMarginAccounts(sdkRo.Conn.Rpc, &sdkRo.Info, sym, addresses)
	if err != nil {
		t.Fatalf("RawQueryMarginAccounts: %v", err)
	}
	m2, err := sdkRo.QueryMarginAccounts(sym, addresses, nil)
	if err != nil {
		t.Fatalf("QueryMarginAccounts: %v", err)
	}
	t.Log(m[0].FPositionBC)
	t.Log(m2[0].FPositionBC)
	t.Log(m[1].FPositionBC)
	t.Log(m2[1].FPositionBC)
}

func TestSdkROOrders(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	startTime := time.Now()
	n, err := sdkRo.QueryNumOrders(sym, nil)
	endTime := time.Now()
	t.Logf("Num orders in %s seconds", endTime.Sub(startTime))
	if err != nil {
		t.Fatalf("QueryNumOrders: %v", err)
	}
	t.Logf("There are %d open orders", n)
	startTime = time.Now()
	rpcUrls := []string{
		"https://bartio.drpc.org",
		"https://bartio.rpc.berachain.com",
		"https://berat2.lava.build",
	}
	orderChan := make(chan *OpenOrders, len(rpcUrls))
	if n == 0 {
		t.Log("No open orders, skipping parallel fetch")
		return
	}
	num := int(n) / min(len(rpcUrls), int(n))
	for i := 0; i < len(rpcUrls); i++ {
		from := i * num
		to := from + num
		go getOrders(*sdkRo, sym, rpcUrls[i], from, to, orderChan)
	}
	orders := make([]*OpenOrders, 0, len(rpcUrls))
	totalOrders := 0
	for i := 0; i < len(rpcUrls); i++ {
		res := <-orderChan
		if res == nil {
			continue
		}
		orders = append(orders, res)
		totalOrders += len(res.OrderHashes)
	}
	endTime = time.Now()
	t.Logf("Found %d orders", totalOrders)
	t.Logf("in %s seconds", endTime.Sub(startTime))
	if len(orders) < 2 {
		t.Log("not enough orders for test")
		return
	}
	t.Log(orders[0].Orders[0].OptTraderAddr.Hex())
	t.Log(orders[0].Orders[1].OptTraderAddr.Hex())
}

func TestFetchCollToSettlePx(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	px, err := sdkRo.FetchCollToSettlePx(sym, "")
	if err != nil {
		t.Fatalf("FetchCollToSettlePx: %v", err)
	}
	t.Logf("collateral to settlement conversion = %f", px)
}

func TestSdkRO(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	perp := getActiveSymbol(t, &sdkRo.Info)
	orders, err := sdkRo.QueryAllOpenOrders(perp, nil) //([]Order, []string, error)
	if err != nil {
		t.Fatalf("QueryAllOpenOrders: %v", err)
	}
	oo := orders.Orders
	dgsts := orders.OrderHashes
	t.Log(oo)
	t.Log(dgsts)

	trader := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	broker := common.HexToAddress("0xB0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
	pr, err := sdkRo.GetPositionRisk(perp, trader, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(pr)
	}
	bal, err := sdkRo.GetSettleTokenBalance(perp, trader, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("balance of margin token=", bal)
	}
	perpState, err := sdkRo.QueryPerpetualState([]int32{100000}, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(perpState)
	}
	poolState, err := sdkRo.QueryPoolStates(nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(poolState)
	}
	oo, dgsts, err = sdkRo.QueryOpenOrders(perp, trader, nil) //([]Order, []string, error)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(oo)
		t.Log(dgsts)
	}
	id := "258ae021f8743b903d8bde405dba7cc7a74d977ce956db8cb6c2c308976ceb89"
	status, err := sdkRo.QueryOrderStatus(perp, trader, id, nil) // (string, error)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(status)
	}
	m, err := sdkRo.QueryMaxTradeAmount(perp, 0, true, nil) // (float64, error) {
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(m)
	}
	vol, err := sdkRo.QueryTraderVolume(1, trader, nil) //(float64, error) {
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(vol)
	}
	fee, err := sdkRo.QueryExchangeFeeTbpsForTrader(1, trader, broker, nil) // (uint16, error) {
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(fee)
	}
	minpos, err := sdkRo.GetMinimalPositionSize(perp) //(float64, error)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(minpos)
	}
}

func TestFetchPricesForPerpetual(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	ep := PriceFeedEndpoints{
		PriceFeedEndpoint:  sdkRo.ChainConfig.PriceFeedEndpoint,
		LowLiqFeedEndpoint: sdkRo.ChainConfig.LowLiqFeedEndpoint,
		PrdMktFeedEndpoint: sdkRo.ChainConfig.PrdMktFeedEndpoint,
		SportFeedEndpoint:  sdkRo.ChainConfig.SportFeedEndpoint,
	}
	pxBundle, err := RawFetchPricesForPerpetual(sdkRo.Info, sym, ep)
	if err != nil {
		t.Fatalf("RawFetchPricesForPerpetual: %v", err)
	}
	t.Log(pxBundle)
}

func TestGetPositionRisks(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	traderAddr := common.HexToAddress("0x96d570211cb7638783Af538DC07122763CF59b3a")
	symbols := []string{sym}
	pRisk, err := sdkRo.GetPositionRisks(symbols, traderAddr, nil)
	if err != nil {
		t.Fatalf("GetPositionRisks: %v", err)
	}
	jason, err := json.MarshalIndent(pRisk, "", " ")
	if err != nil {
		t.Fatalf("MarshalIndent: %v", err)
	}
	t.Log(string(jason))

	pRisk, err = sdkRo.GetPositionRiskAll(traderAddr, nil)
	if err != nil {
		t.Fatalf("GetPositionRiskAll: %v", err)
	}
	jason, err = json.MarshalIndent(pRisk, "", " ")
	if err != nil {
		t.Fatalf("MarshalIndent: %v", err)
	}
	t.Log(string(jason))
}

func TestGetPositionRisk(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	sym := getAnySymbol(t, &info)
	ep := PriceFeedEndpoints{
		PriceFeedEndpoint:  chConf.PriceFeedEndpoint,
		LowLiqFeedEndpoint: chConf.LowLiqFeedEndpoint,
		PrdMktFeedEndpoint: chConf.PrdMktFeedEndpoint,
		SportFeedEndpoint:  chConf.SportFeedEndpoint,
	}
	pRisk, err := RawGetPositionRisk(info, conn.Rpc, (*common.Address)(&traderAddr), sym, ep)
	if err != nil {
		t.Fatalf("RawGetPositionRisk: %v", err)
	}
	t.Log(pRisk)
}

func TestQueryPerpetualState(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	// Get perpIds dynamically from loaded exchange info
	var perpIds []int32
	for _, p := range info.Perpetuals {
		perpIds = append(perpIds, int32(p.Id))
		if len(perpIds) >= 2 {
			break
		}
	}
	if len(perpIds) == 0 {
		t.Skip("no perpetuals found in exchange info")
	}
	ep := PriceFeedEndpoints{
		PriceFeedEndpoint:  chConf.PriceFeedEndpoint,
		LowLiqFeedEndpoint: chConf.LowLiqFeedEndpoint,
		PrdMktFeedEndpoint: chConf.PrdMktFeedEndpoint,
		SportFeedEndpoint:  chConf.SportFeedEndpoint,
	}
	perpState, err := RawQueryPerpetualState(conn.Rpc, info, perpIds, ep)
	if err != nil {
		t.Fatalf("RawQueryPerpetualState: %v", err)
	}
	t.Log(perpState)
}

func TestQueryPoolStates(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	poolStates, err := RawQueryPoolStates(conn.Rpc, info)
	if err != nil {
		t.Fatalf("RawQueryPoolStates: %v", err)
	}
	for _, p := range poolStates {
		t.Logf("--- Pool %d ---", p.Id)
		t.Log(p.IsRunning)
		t.Log("DefaultFundCashCC=", p.DefaultFundCashCC)
		t.Log("PnlParticipantCashCC=", p.PnlParticipantCashCC)
		t.Log("TotalSupplyShareToken=", p.TotalSupplyShareToken)
		t.Log("TotalTargetAMMFundSizeCC=", p.TotalTargetAMMFundSizeCC)
	}
}

func TestQueryOpenOrders(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	sym := getAnySymbol(t, &info)
	orders, digests, err := RawQueryOpenOrders(conn.Rpc, info, sym, traderAddr)
	if err != nil {
		t.Fatalf("RawQueryOpenOrders: %v", err)
	}
	t.Log("--- orders ", orders)
	t.Log("--- digests ", digests)

	// order status
	if len(digests) < 1 {
		return
	}
	d := digests[0]
	status, err := RawQueryOrderStatus(conn.Rpc, info, traderAddr, d, sym)
	if err != nil {
		t.Fatalf("RawQueryOrderStatus: %v", err)
	}
	t.Log("order status: ", status)
}

func TestQueryTraderVolume(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	volume, err := RawQueryTraderVolume(conn.Rpc, info, traderAddr, 1)
	if err != nil {
		t.Fatalf("RawQueryTraderVolume: %v", err)
	}
	t.Log("Trader volume = ", volume)
}

func TestQueryExchangeFeeTbpsForTrader(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	brokerAddr := common.Address{}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	fee, err := RawQueryExchangeFeeTbpsForTrader(conn.Rpc, info, 1, traderAddr, brokerAddr)
	if err != nil {
		t.Fatalf("RawQueryExchangeFeeTbpsForTrader: %v", err)
	}
	t.Log("Fee Tbps = ", fee)
}

func TestQueryMaxTradeAmount(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	sym := getAnySymbol(t, &info)
	trade, err := RawQueryMaxTradeAmount(conn.Rpc, info, 0.01, sym, true)
	if err != nil {
		t.Fatalf("RawQueryMaxTradeAmount: %v", err)
	}
	t.Logf("Max Trade buy (0.01 %s) = %v", sym, trade)
}

func TestApproximateOrderBook(t *testing.T) {
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdkRo.Info)
	ob, err := sdkRo.ApproximateOrderBook(sym, nil)
	if err != nil {
		t.Fatalf("ApproximateOrderBook: %v", err)
	}
	jason, err := json.Marshal(ob)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	t.Log(string(jason))

	var nob OrderBook
	err = json.Unmarshal([]byte(jason), &nob)
	if err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
}
