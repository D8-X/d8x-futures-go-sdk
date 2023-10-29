package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

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

// LoadPriceFeedConfig loads the price feed config file cfFilePath into struct PriceFeedConfig
// for the network called configNetwork
// for example LoadPriceFeedConfig("config/priceFeedConfig.json", "testnet")
func LoadPriceFeedConfig(cfFilePath string, configNetwork string) (PriceFeedConfig, error) {
	// Read the JSON file
	data, err := os.ReadFile(cfFilePath)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return PriceFeedConfig{}, err
	}
	var configuration []PriceFeedConfig
	// Unmarshal the JSON data into the Configuration struct
	err = json.Unmarshal(data, &configuration)
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

// LoadChainConfig loads the chain-config from cfFilePath (e.g. "config/config.json") into ChainConfig struct
// for network configName
func LoadChainConfig(cfFilePath string, configName string) (ChainConfig, error) {
	// Read the JSON file
	data, err := os.ReadFile(cfFilePath)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return ChainConfig{}, err
	}
	var configuration []ChainConfig
	// Unmarshal the JSON data into the Configuration struct
	err = json.Unmarshal(data, &configuration)
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
