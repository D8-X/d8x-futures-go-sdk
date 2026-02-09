//go:build integration

package d8x_futures

import (
	"testing"
	"time"
)

func TestGetWEETHPrice(t *testing.T) {
	oracles, err := NewChainOracles(42161)
	if err != nil {
		t.Fatalf("NewChainOracles: %v", err)
	}
	for k := 0; k < 4; k++ {
		px, ts, err := oracles.GetPrice("WEETH-ETH", true)
		if err != nil {
			t.Fatalf("GetPrice WEETH-ETH: %v", err)
		}
		t.Logf("px=%.4f ts=%d", px, ts)
		// wait
		time.Sleep(5 * time.Second)
	}
}

func TestGetSTUSDPrice(t *testing.T) {
	oracles, err := NewChainOracles(42161)
	if err != nil {
		t.Fatalf("NewChainOracles: %v", err)
	}
	for k := 0; k < 4; k++ {
		px, ts, err := oracles.GetPrice("STUSD-USDC", true)
		if err != nil {
			if k == 0 {
				t.Skipf("STUSD-USDC oracle not available: %v", err)
			}
			t.Fatalf("GetPrice STUSD-USDC: %v", err)
		}
		t.Logf("px=%.4f ts=%d", px, ts)
		// wait
		time.Sleep(5 * time.Second)
	}
}
