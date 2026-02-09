//go:build integration

package d8x_futures

import (
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/config"
)

func TestPerpetualStaticInfo(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getActiveSymbol(t, &sdk.Info)
	info, err := sdk.GetPerpetualStaticInfo(sym)
	if err != nil {
		t.Fatalf("GetPerpetualStaticInfo: %v", err)
	}
	t.Log(info)
}

func TestQueryNestedPerpetualInfo(t *testing.T) {
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
	p, err := QueryNestedPerpetualInfo(conn)
	if err != nil {
		t.Fatalf("QueryNestedPerpetualInfo: %v", err)
	}
	t.Log(p.PerpetualIds)
}

func TestReadSymbolList(t *testing.T) {
	symMap, err := config.GetSymbolList()
	if err != nil {
		t.Fatalf("GetSymbolList: %v", err)
	}
	t.Log((symMap)["MATC"])
}

func TestQueryPoolStaticInfo(t *testing.T) {
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
	nest, err := QueryNestedPerpetualInfo(conn)
	if err != nil {
		t.Fatalf("QueryNestedPerpetualInfo: %v", err)
	}
	info, err := QueryExchangeStaticInfo(&conn, &chConf, &pxConf, &nest)
	if err != nil {
		t.Fatalf("QueryExchangeStaticInfo: %v", err)
	}
	t.Log(info)
	info.Store("./tmpXchInfo.json")
}

func TestTriangulate(t *testing.T) {
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}

	triangs := Triangulate("CHF-USDC", pxConf.PriceFeedIds)
	t.Log("Triangulate")
	t.Log(triangs)

	// test an impossible path
	triangs2 := Triangulate("CHF-DOGE", pxConf.PriceFeedIds)
	t.Log(triangs2)
}
