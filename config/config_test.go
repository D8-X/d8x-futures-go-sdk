package config

import (
	"fmt"
	"testing"
)

func TestGetSettlementConfig(t *testing.T) {
	config, err := GetSettlementConfig()
	if err != nil {
		t.Skip("config not found, skipping in CI")
	}
	fmt.Print(config[0].PerpFlags.String())
}

func TestGetPriceFeedConfig(t *testing.T) {
	config, err := GetDefaultPriceConfig(421614)
	if err != nil {
		t.Skip("config not found, skipping in CI")
	}
	fmt.Print(config.PriceFeedIds)
}
