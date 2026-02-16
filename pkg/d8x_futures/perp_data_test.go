package d8x_futures

import (
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
