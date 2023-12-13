package d8x_futures

import (
	"fmt"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

func TestQueryNestedPerpetualInfo(t *testing.T) {
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.PriceFeedNetwork)
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	p, err := QueryNestedPerpetualInfo(conn)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println(p.PerpetualIds)
}

func TestReadSymbolList(t *testing.T) {
	symMap, err := config.GetSymbolList()
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println((symMap)["MATC"])

}

func TestQueryPoolStaticInfo(t *testing.T) {
	chConf, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.PriceFeedNetwork)
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	nest, err := QueryNestedPerpetualInfo(conn)
	if err != nil {
		t.Logf(err.Error())
	}
	info := QueryExchangeStaticInfo(&conn, &chConf, &nest)
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
	pxConf, err := config.GetDefaultPriceConfig("PythEVMStable")
	if err != nil {
		t.Logf(err.Error())
	}
	triangs := Triangulate("CHF-USDC", &pxConf)
	fmt.Println("Triangulate")
	fmt.Println(triangs)

	// test an impossible path
	triangs2 := Triangulate("CHF-DOGE", &pxConf)
	fmt.Println(triangs2)
}
