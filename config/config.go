package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

const (
	SYNC_HUB_URL = "https://raw.githubusercontent.com/D8-X/sync-hub/dev/d8x-futures-go-sdk/"
)

//go:embed embedded/*
var EmbededConfigs embed.FS

// fetchConfigFromRepo gets the config file from Github
func fetchConfigFromRepo(configName string) ([]byte, error) {
	url := SYNC_HUB_URL + configName
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching config: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch config, status code: %d", resp.StatusCode)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}

	return jsonData, nil
}

func GetSymbolList() (map[string]string, error) {
	jsonData, err := fetchConfigFromRepo("symbolList.json")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch symbolList.json:" + err.Error())
	}
	// Define a map to store the data
	data := make(map[string]string)
	// Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetSettlementConfig() ([]utils.SettlementConfig, error) {
	fs, err := EmbededConfigs.Open("embedded/settlement.json")
	if err != nil {
		return nil, err
	}
	jsonStr, err := io.ReadAll(fs)
	if err != nil {
		return nil, err
	}
	var settlements []utils.SettlementConfig
	err = json.Unmarshal([]byte(jsonStr), &settlements)
	if err != nil {
		return nil, err
	}
	return settlements, nil
}

func GetDefaultChainConfigNames() ([]string, error) {
	// Read the JSON file
	fs, err := EmbededConfigs.Open("embedded/chainConfig.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return nil, err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return nil, err
	}
	return utils.LoadChainConfigNames(data)
}

func GetMultiPayAddr(chainId int64) (string, error) {
	// Read the JSON file
	fs, err := EmbededConfigs.Open("embedded/chainConfig.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return "", err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return "", err
	}
	c, err := utils.LoadChainConfigFromId(data, chainId)
	if err != nil {
		return "", err
	}
	return c.MultiPayAddr.Hex(), nil
}

func GetCustomChainConfig(filePath, configName string) (utils.ChainConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return utils.ChainConfig{}, err
	}
	return utils.LoadChainConfig(data, configName)
}

func GetDefaultChainConfig(configName string) (utils.ChainConfig, error) {
	// Read the JSON file
	fs, err := EmbededConfigs.Open("embedded/chainConfig.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return utils.ChainConfig{}, err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return utils.ChainConfig{}, err
	}
	return utils.LoadChainConfig(data, configName)
}

func GetDefaultChainConfigFromId(chainId int64) (utils.ChainConfig, error) {
	// Read the JSON file
	fs, err := EmbededConfigs.Open("embedded/chainConfig.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return utils.ChainConfig{}, err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return utils.ChainConfig{}, err
	}
	return utils.LoadChainConfigFromId(data, chainId)
}

func GetDefaultPriceConfig(chainId int64) (utils.PriceFeedConfig, error) {

	data, err := fetchConfigFromRepo("priceFeedConfig.json")
	if err != nil {
		return utils.PriceFeedConfig{}, fmt.Errorf("unable to fetch priceFeedConfig.json:" + err.Error())
	}
	ch, err := GetDefaultChainConfigFromId(chainId)
	if err != nil {
		return utils.PriceFeedConfig{}, err
	}
	return utils.LoadPriceFeedConfig(data, ch.PriceFeedNetwork)
}

func GetDefaultPriceConfigByName(configNetwork string) (utils.PriceFeedConfig, error) {
	data, err := fetchConfigFromRepo("priceFeedConfig.json")
	if err != nil {
		return utils.PriceFeedConfig{}, fmt.Errorf("unable to fetch priceFeedConfig.json:" + err.Error())
	}
	return utils.LoadPriceFeedConfig(data, configNetwork)
}

// GetPriceFeedOnChain loads priceFeedOnChain into the corresponding structure
func GetPriceFeedOnChain() ([]utils.PriceFeedOnChainConfig, error) {
	fs, err := EmbededConfigs.Open("embedded/priceFeedOnChain.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return nil, err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return nil, err
	}
	var configuration []utils.PriceFeedOnChainConfig
	// Unmarshal the JSON data into the Configuration struct
	err = json.Unmarshal(data, &configuration)
	if err != nil {
		return nil, fmt.Errorf("decoding JSON: %s", err.Error())
	}
	return configuration, nil
}
