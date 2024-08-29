package config

import (
	"fmt"
	"testing"
)

func TestGetSettlementConfig(t *testing.T) {
	config, err := GetSettlementConfig()
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	fmt.Print(config[0].PerpFlags.String())
}

func TestGetPriceFeedConfig(t *testing.T) {
	config, err := GetDefaultPriceConfig(421614)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	fmt.Print(config.PriceFeedIds)
}
