package d8x_futures

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Name                             string         `json:"name"`
	ChainId                          int64          `json:"chainId"`
	Version                          int32          `json:"version"`
	ProxyAddr                        common.Address `json:"proxyAddr"`
	NodeURL                          string         `json:"nodeURL"`
	PriceFeedConfigNetwork           string         `json:"priceFeedConfigNetwork"`
	ShareTokenABILocation            string         `json:"shareTokenABILocation"`
	ProxyABILocation                 string         `json:"proxyABILocation"`
	LimitOrderBookFactoryABILocation string         `json:"limitOrderBookFactoryABILocation"`
	LimitOrderBookABILocation        string         `json:"limitOrderBookABILocation"`
}

type PriceFeedConfig struct {
	Network      string              `json:"network"`
	PriceFeedIds []PriceFeedId       `json:"ids"`
	EndPoints    []PriceFeedEndpoint `json:"endpoints"`
}

type PriceFeedId struct {
	Symbol string `json:"symbol"`
	Id     string `json:"id"`
	Type   string `json:"type"`
	Origin string `json:"origin"`
}

type PriceFeedEndpoint struct {
	Type        string `json:"type"`
	EndpointUrl string `json:"endpoint"`
}

func loadPriceFeedConfig(configNetwork string) (PriceFeedConfig, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile("config/priceFeedConfig.json")
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

func loadConfig(configName string) (Config, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return Config{}, err
	}
	var configuration []Config
	// Unmarshal the JSON data into the Configuration struct
	err = json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return Config{}, err
	}
	for i := 0; i < len(configuration); i++ {
		if configuration[i].Name == configName {
			return configuration[i], nil
		}
	}
	return Config{}, errors.New("config not found")
}
