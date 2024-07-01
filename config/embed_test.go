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
