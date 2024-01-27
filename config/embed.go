package config

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

//go:embed embedded/*
var EmbededConfigs embed.FS

func GetSymbolList() (map[string]string, error) {
	fs, err := EmbededConfigs.Open("embedded/symbolList.json")
	if err != nil {
		return nil, err
	}
	// Define a map to store the data
	data := make(map[string]string)
	jsonData, err := io.ReadAll(fs)
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
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
	fs, err := EmbededConfigs.Open("embedded/priceFeedConfig.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return utils.PriceFeedConfig{}, err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return utils.PriceFeedConfig{}, err
	}
	ch, err := GetDefaultChainConfigFromId(chainId)
	if err != nil {
		return utils.PriceFeedConfig{}, err
	}
	return utils.LoadPriceFeedConfig(data, ch.PriceFeedNetwork)
}

func GetDefaultPriceConfigByName(configNetwork string) (utils.PriceFeedConfig, error) {
	fs, err := EmbededConfigs.Open("embedded/priceFeedConfig.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return utils.PriceFeedConfig{}, err
	}
	data, err := io.ReadAll(fs)
	if err != nil {
		return utils.PriceFeedConfig{}, err
	}
	return utils.LoadPriceFeedConfig(data, configNetwork)
}
