package d8x_futures

import (
	"fmt"
	"testing"
	"time"
)

func TestGetWEETHPrice(t *testing.T) {
	oracles, err := NewChainOracles()
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	for k := 0; k < 4; k++ {
		px, ts, err := oracles.GetPrice("WEETH-ETH", true)
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Printf("px=%.4f ts=%d\n", px, ts)
		// wait
		time.Sleep(5 * time.Second)
	}
}

func TestGetSTUSDPrice(t *testing.T) {
	oracles, err := NewChainOracles()
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	for k := 0; k < 4; k++ {
		px, ts, err := oracles.GetPrice("STUSD-USDC", true)
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Printf("px=%.4f ts=%d\n", px, ts)
		// wait
		time.Sleep(5 * time.Second)
	}
}
