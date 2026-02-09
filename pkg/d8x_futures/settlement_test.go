//go:build integration

package d8x_futures

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load .env before any tests run
	var err error
	if err = godotenv.Load(".env"); err != nil {
		// Try parent
		err = godotenv.Load("../../.env")
		if err != nil {
			err = godotenv.Load("../.env")
		}
	}
	if err != nil {
		wd, _ := os.Getwd()
		fmt.Println("working dir:", wd)
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate tmpXchInfo.json so tests that depend on it can load it
	if err := generateTmpXchInfo(); err != nil {
		fmt.Println("warning: could not generate tmpXchInfo.json:", err)
	}

	code := m.Run() // run tests
	os.Exit(code)   // exit with proper code
}

// generateTmpXchInfo queries on-chain static info and writes tmpXchInfo.json.
func generateTmpXchInfo() error {
	chConf, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		return err
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		return err
	}
	conn, err := CreateBlockChainConnector(pxConf, chConf, nil)
	if err != nil {
		return err
	}
	nest, err := QueryNestedPerpetualInfo(conn)
	if err != nil {
		return err
	}
	info, err := QueryExchangeStaticInfo(&conn, &chConf, &pxConf, &nest)
	if err != nil {
		return err
	}
	return info.Store("./tmpXchInfo.json")
}

func TestQueryPerpetualStateEnum(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym, err := sdk.internalToSymbol("NHL0-USD")
	if err != nil {
		t.Fatalf("internalToSymbol: %v", err)
	}
	state, err := sdk.QueryPerpetualStateEnum(sym, nil)
	if err != nil {
		t.Fatalf("QueryPerpetualStateEnum: %v", err)
	}
	t.Log(state.String())
}

func TestSetter(t *testing.T) {
	pk := os.Getenv("PK")
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	for _, p := range sdk.Info.Perpetuals {
		t.Logf("%s -> %s", p.S2Symbol, p.state.String())
	}
}

func TestRefreshPerpetualStateEnum(t *testing.T) {
	pk := os.Getenv("PK")
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	url := os.Getenv("RPC")
	if url == "" {
		t.Fatal("Provide RPC URL as environment variable RPC")
	}
	rpc, err := ethclient.Dial(url)
	if err != nil {
		t.Fatalf("ethclient.Dial: %v", err)
	}
	sym, err := sdk.internalToSymbol("NHL0-USD")
	if err != nil {
		t.Fatalf("internalToSymbol: %v", err)
	}
	state, err := sdk.RefreshPerpetualStateEnum(sym, rpc)
	if err != nil {
		t.Fatalf("RefreshPerpetualStateEnum: %v", err)
	}
	t.Log(state.String())
	for _, p := range sdk.Info.Perpetuals {
		t.Log(p.S2Symbol, p.State().String())
	}
}

func TestInfo(t *testing.T) {
	pk := os.Getenv("PK")
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	for _, p := range sdk.Info.Perpetuals {
		symbol := p.S2Symbol + "-PUSD"
		fmt.Println(symbol, p.Id)
	}
}

func TestSettlement(t *testing.T) {
	pk := os.Getenv("PK")
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	url := os.Getenv("RPC")
	if url == "" {
		t.Fatal("Provide RPC URL as environment variable RPC")
	}
	rpc, err := ethclient.Dial(url)
	if err != nil {
		t.Fatalf("ethclient.Dial: %v", err)
	}
	sdk, err := NewSdk([]string{pk}, "84532", WithRpcClient(rpc))
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	symbol := "SP07-USD-PUSD"
	err = sdk.RunSettlementProcess(symbol, 1.5, 1, rpc)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	symbols := make([]string, 0)
	for _, p := range sdk.Info.Perpetuals {
		t.Logf("%s -> %s", p.S2Symbol, p.state.String())
		if p.state.String() == EMERGENCY.String() {
			symbols = append(symbols, p.S2Symbol)
		}
	}
	for _, s := range symbols {
		symbol := s + "-PUSD"
		s2 := float64(2)
		s3 := float64(1)
		err = sdk.RunSettlementProcess(symbol, s2, s3, rpc)
		if err != nil {
			t.Fatalf("RunSettlementProcess: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}

func TestEnable(t *testing.T) {
	pk := os.Getenv("PK")
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	url := os.Getenv("RPC")
	if url == "" {
		t.Fatal("Provide RPC URL as environment variable RPC")
	}
	rpc, err := ethclient.Dial(url)
	if err != nil {
		t.Fatalf("ethclient.Dial: %v", err)
	}
	
	var symbol string
	for _, p := range sdk.Info.Perpetuals {
		if p.State() == INVALID {
			sym, ok := sdk.Info.PerpetualIdToSymbol[p.Id]
			if ok {
				symbol = sym
				break
			}
		}
	}
	if symbol == "" {
		t.Skip("no INVALID perpetual found to activate")
	}
	t.Logf("activating %s", symbol)
	tx, err := sdk.ActivatePerpetual(symbol, rpc)
	if err != nil {
		if strings.Contains(err.Error(), "only maintainer") {
			t.Skipf("wallet lacks maintainer role: %v", err)
		}
		t.Fatalf("ActivatePerpetual: %v", err)
	}
	t.Log(tx.Hash())
}
