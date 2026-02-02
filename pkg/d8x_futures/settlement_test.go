package d8x_futures

import (
	"fmt"
	"os"
	"testing"
	"time"

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
	code := m.Run() // run tests
	os.Exit(code)   // exit with proper code
}

func TestQueryPerpetualStateEnum(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	sym, err := sdk.internalToSymbol("NHL0-USD")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	state, err := sdk.QueryPerpetualStateEnum(sym, nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(state.String())
}

func TestSetter(t *testing.T) {
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
		fmt.Printf("%s -> %s\n", p.S2Symbol, p.state.String())
	}
}

func TestRefreshPerpetualStateEnum(t *testing.T) {
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
	url := os.Getenv("RPC")
	rpc, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	sym, err := sdk.internalToSymbol("NHL0-USD")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	state, err := sdk.RefreshPerpetualStateEnum(sym, rpc)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(state.String())
	for _, p := range sdk.Info.Perpetuals {
		fmt.Println(p.S2Symbol, p.State().String())
	}
}

func TestSettlement(t *testing.T) {
	url := os.Getenv("RPC")
	rpc, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	sdk, err := NewSdk([]string{os.Getenv("PK")}, "84532", WithRpcClient(rpc))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	for _, p := range sdk.Info.Perpetuals {
		if p.State() == NORMAL || p.State() == INVALID {
			continue
		}

		symbol := p.S2Symbol + "-PUSD"
		px, err := sdk.FetchPricesForPerpetualId(p.Id, "")
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}

		fmt.Println(symbol, "s2=", px.S2Price, "s3=", px.S3Price, p.State().String(), p.Id)
		if p.state != CLEARED {
			continue
		}
		err = sdk.RunSettlementProcess(symbol, max(1, px.S2Price), px.S3Price, rpc)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
	}
	symbols := make([]string, 0)
	for _, p := range sdk.Info.Perpetuals {
		fmt.Printf("%s -> %s\n", p.S2Symbol, p.state.String())
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
			fmt.Println(err)
			t.FailNow()
		}
		time.Sleep(1 * time.Second)
	}
}

func TestEnable(t *testing.T) {
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
	url := os.Getenv("RPC")
	rpc, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	symbol := "NHL2-USD-PUSD"
	tx, err := sdk.ActivatePerpetual(symbol, rpc)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(tx.Hash())
}
