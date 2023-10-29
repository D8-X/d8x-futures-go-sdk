package d8x_futures

import (
	"fmt"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

func TestQueryNestedPerpetualInfo(t *testing.T) {
	config, err := utils.LoadChainConfig("config/chainConfig.json", "testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector("config/priceFeedConfig.json", config)
	p := QueryNestedPerpetualInfo(conn)
	fmt.Println(p.PerpetualIds)
}

func TestReadSymbolList(t *testing.T) {
	symMap, err := readSymbolList()
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println((*symMap)["MATC"])

}

func TestQueryPoolStaticInfo(t *testing.T) {
	config, err := utils.LoadChainConfig("config/chainConfig", "testnet")
	if err != nil {
		t.Logf(err.Error())
	}

	conn := CreateBlockChainConnector("config/priceFeedConfig.json", config)
	nest := QueryNestedPerpetualInfo(conn)
	info := QueryExchangeStaticInfo(conn, config, nest)
	fmt.Println(info)
	info.Store("./tmpXchInfo.json")
}

func TestFindPath(t *testing.T) {
	ccyBase := []string{"USD", "USDC", "EUR", "BTC", "BTC"}
	ccyQuote := []string{"CHF", "USD", "USD", "CHF", "EUR"}
	pair := "CHF-USDC"

	paths := findPath(append(ccyBase, ccyQuote...), append(ccyQuote, ccyBase...), pair)
	fmt.Println(paths)

}

func TestPythNToFloat64(t *testing.T) {
	v := utils.PythNToFloat64("314159265358979", -14)
	fmt.Println("v=", v)
}

func TestTriangulate(t *testing.T) {
	pxConfig, err := utils.LoadPriceFeedConfig("config/priceFeedConfig.json", "testnet")
	if err != nil {
		panic(err)
	}
	triangs := Triangulate("CHF-USDC", pxConfig)
	fmt.Println("Triangulate")
	fmt.Println(triangs)

	// test an impossible path
	triangs2 := Triangulate("CHF-DOGE", pxConfig)
	fmt.Println(triangs2)
}
