package d8x_futures

import (
	"fmt"
	"testing"
)

func TestQueryNestedPerpetualInfo(t *testing.T) {
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
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
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	nest := QueryNestedPerpetualInfo(conn)
	info := QueryExchangeStaticInfo(conn, nest)
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
	v := PythNToFloat64("314159265358979", -14)
	fmt.Println("v=", v)
}

func TestTriangulate(t *testing.T) {
	pxConfig, err := LoadPriceFeedConfig("testnet")
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
