package d8x_futures

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestAngleFeed(t *testing.T) {

	provider, err := ethclient.Dial("https://rpc.ankr.com/arbitrum")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return
	}

	rate, err := STUSDToUSDC(provider)
	if err != nil {
		fmt.Println("Failed to get conversion rate:", err)
		return
	}

	fmt.Println("Conversion rate:", rate)
}
