package d8x_futures

import (
	"fmt"
	"testing"
	"time"
)

func TestGetPrice(t *testing.T) {
	oracles, err := NewChainOracles()
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	for k := 0; k < 4; k++ {
		px, err := oracles.GetPrice("WEETH-ETH", true)
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		fmt.Printf("px=%.4f\n", px)
		// wait
		time.Sleep(5 * time.Second)
	}
}
