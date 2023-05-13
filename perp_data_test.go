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
	info := QueryPoolStaticInfo(conn, nest)
	fmt.Println(info)
}

func TestFindPath(t *testing.T) {
	ccyBase := []string{"USD", "USDC", "EUR", "BTC", "BTC"}
	ccyQuote := []string{"CHF", "USD", "USD", "CHF", "EUR"}
	pair := "CHF-USDC"

	paths := findPath(append(ccyBase, ccyQuote...), append(ccyQuote, ccyBase...), pair)
	fmt.Println(paths)

}

func TestTriangulate(t *testing.T) {
	pxConfig, err := LoadPriceFeedConfig("testnet")
	if err != nil {
		panic(err)
	}
	triangs := Triangulate("CHF-USDC", pxConfig)
	fmt.Println(triangs)
}
