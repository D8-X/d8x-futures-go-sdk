//go:build integration

package d8x_futures

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestAngleFeed(t *testing.T) {

	provider, err := ethclient.Dial("https://arbitrum-one-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	rate, err := STUSDToUSDC(provider)
	if err != nil {
		t.Fatalf("Failed to get conversion rate: %v", err)
	}

	t.Log("Conversion rate:", rate)
}
