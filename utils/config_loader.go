package utils

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	Name               string         `json:"name"`
	PriceFeedNetwork   string         `json:"priceFeedNetwork"`
	PriceFeedEndpoints []string       `json:"priceFeedEndpoints"`
	ChainId            int64          `json:"chainId"`
	ProxyAddr          common.Address `json:"proxyAddr"`
	MultiPayAddr       common.Address `json:"multipayAddr"`
	NodeURL            string         `json:"nodeURL"`
	PostOrderGasLimit  int64          `json:"postOrderGasLimit"`
	PriceUpdateFeeGwei int64          `json:"priceUpdateFeeGwei"`
	SplitExecutionTx   bool           `json:"splitExecutionTx"`
}

type PriceFeedConfig struct {
	Network                  string        `json:"network"`
	PriceFeedIds             []PriceFeedId `json:"ids"`
	ThresholdMarketClosedSec int32         `json:"threshold_market_closed_sec"`
	PriceUpdateFeeGwei       int64
}

type PriceFeedId struct {
	Symbol string `json:"symbol"`
	Id     string `json:"id"`
	Type   string `json:"type"`
	Origin string `json:"origin"`
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
// for the given network name or chain id
func LoadChainConfig(data []byte, configNameOrChainId string) (ChainConfig, error) {

	var configuration []ChainConfig
	// Unmarshal the JSON data into the Configuration struct
	err := json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return ChainConfig{}, err
	}
	chainId, err := strconv.Atoi(configNameOrChainId)
	if err != nil {
		// user provided a name not an id
		chainId = -1
	}
	for i := 0; i < len(configuration); i++ {
		if configuration[i].Name == configNameOrChainId {
			return configuration[i], nil
		}
		if configuration[i].ChainId == int64(chainId) {
			return configuration[i], nil
		}
	}
	return ChainConfig{}, errors.New("config not found")
}

// LoadChainConfig loads the chain-config from data into ChainConfig struct
// for network with chain id chainId
func LoadChainConfigFromId(data []byte, chainId int64) (ChainConfig, error) {

	var configuration []ChainConfig
	// Unmarshal the JSON data into the Configuration struct
	err := json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return ChainConfig{}, err
	}
	for i := 0; i < len(configuration); i++ {
		if configuration[i].ChainId == chainId {
			return configuration[i], nil
		}
	}
	return ChainConfig{}, errors.New("config not found")
}

// LoadChainConfig loads the chain-config from data into ChainConfig struct
// for network configName
func LoadChainConfigNames(data []byte) ([]string, error) {

	var configuration []ChainConfig
	// Unmarshal the JSON data into the Configuration struct
	err := json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return nil, err
	}
	var ret []string
	for i := 0; i < len(configuration); i++ {
		ret = append(ret, configuration[i].Name)
	}
	return ret, nil
}
