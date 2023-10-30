package utils

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	Name                   string         `json:"name"`
	PriceFeedNetwork       string         `json:"priceFeedNetwork"`
	ChainId                int64          `json:"chainId"`
	Version                int32          `json:"version"`
	ProxyAddr              common.Address `json:"proxyAddr"`
	NodeURL                string         `json:"nodeURL"`
	PriceFeedConfigNetwork string         `json:"priceFeedConfigNetwork"`
	PostOrderGasLimit      int64          `json:"postOrderGasLimit"`
}

type PriceFeedConfig struct {
	Network                  string              `json:"network"`
	PriceFeedIds             []PriceFeedId       `json:"ids"`
	EndPoints                []PriceFeedEndpoint `json:"endpoints"`
	ThresholdMarketClosedSec int32               `json:"threshold_market_closed_sec"`
}

type PriceFeedId struct {
	Symbol string `json:"symbol"`
	Id     string `json:"id"`
	Type   string `json:"type"`
	Origin string `json:"origin"`
}

type PriceFeedEndpoint struct {
	Type        string   `json:"type"`
	EndpointUrl []string `json:"endpoints"`
}

// LoadPriceFeedConfig loads the price feed config file
// data into struct PriceFeedConfig for the network called configNetwork
// for example LoadPriceFeedConfig("config/priceFeedConfig.json", "testnet")
func LoadPriceFeedConfig(data []byte, configNetwork string) (PriceFeedConfig, error) {
	var configuration []PriceFeedConfig
	// Unmarshal the JSON data into the Configuration struct
	err := json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return PriceFeedConfig{}, err
	}
	for i := 0; i < len(configuration); i++ {
		if configuration[i].Network == configNetwork {
			return configuration[i], nil
		}
	}
	return PriceFeedConfig{}, errors.New("config not found")
}

// LoadChainConfig loads the chain-config from data into ChainConfig struct
// for network configName
func LoadChainConfig(data []byte, configName string) (ChainConfig, error) {

	var configuration []ChainConfig
	// Unmarshal the JSON data into the Configuration struct
	err := json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return ChainConfig{}, err
	}
	for i := 0; i < len(configuration); i++ {
		if configuration[i].Name == configName {
			return configuration[i], nil
		}
	}
	return ChainConfig{}, errors.New("config not found")
}
