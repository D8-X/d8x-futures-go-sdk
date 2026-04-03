package d8x_futures

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

func TestFindPath(t *testing.T) {
	ccyBase := []string{"USD", "USDC", "EUR", "BTC", "BTC"}
	ccyQuote := []string{"CHF", "USD", "USD", "CHF", "EUR"}
	pair := "CHF-USDC"

	paths := findPath(append(ccyBase, ccyQuote...), append(ccyQuote, ccyBase...), pair)
	t.Log(paths)
}

func TestPythNToFloat64(t *testing.T) {
	v, err := utils.PythNToFloat64("314159265358979", -14)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("v=", v)
}

func TestABDKToFloat64(t *testing.T) {
	num := -1.370863664871871
	v := utils.Float64ToABDK(num)
	// 25265520445871489912
	t.Log("v=", v)
}

func TestConfig(t *testing.T) {
	addr, err := config.GetMultiPayAddr(42161)
	if err != nil {
		t.Fatalf("GetMultiPayAddr: %v", err)
	}
	t.Logf("addr = %s", addr)
}

func TestSetPerpetualMarginRates(t *testing.T) {
	skipIfCI(t)
	chainID := "84532" // base sepolia
	pk := os.Getenv("PK")
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, chainID)
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	num := len(sdk.Info.Perpetuals)

	v := sdk.Info.Perpetuals[num-1]
	sym := sdk.Info.PerpetualIdToSymbol[v.Id]
	im := v.InitialMarginRate
	mm := v.MaintenanceMarginRate
	_, err = sdk.SetPerpetualMarginRates(sym, im, mm, nil)
	if err != nil {
		fmt.Println("failed: %w", err)
		t.FailNow()
	}
	sdkRO, err := NewSdkRO(chainID)
	if err != nil {
		fmt.Println("failed: %w", err)
		t.FailNow()
	}
	v2 := sdkRO.Info.Perpetuals[num-1]
	if !approxEq(v2.MaintenanceMarginRate, v.MaintenanceMarginRate, 0.0001) {
		t.Fatalf("maintenance margin rate mismatch")
	}
	if !approxEq(v2.InitialMarginRate, v.InitialMarginRate, 0.0001) {
		t.Fatalf("initial margin rate mismatch")
	}
}

func approxEq(a, b, diff float64) bool {
	return math.Abs(a-b) < diff
}
