package stork

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/spf13/viper"
)

func loadCredentials() (string, string) {
	viper.SetConfigFile("../../.env")
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("could not load .env file" + err.Error())
	}
	return viper.GetString("STORK_CREDENTIALS"), viper.GetString("STORK_ENDPOINT")
}

func TestRestFetchStorkPrice(t *testing.T) {
	config, err := config.GetDefaultPriceConfig(421614)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	cred, endp := loadCredentials()
	s := NewStork(endp, cred)
	j := -1
	for k, c := range config.PriceFeedIds {
		if c.StorkSym != "" {
			j = k
		}
	}
	if j == -1 {
		fmt.Println("found no storksym")
		t.FailNow()
	}
	sym := config.PriceFeedIds[j].StorkSym
	fmt.Println("symbol = ", sym)
	px, err := s.RestFetchStorkPrice(sym)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	fmt.Println("price = ", px)
}
