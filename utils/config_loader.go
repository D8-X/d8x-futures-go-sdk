package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	Name               string         `json:"name"`
	PriceFeedNetwork   string         `json:"priceFeedNetwork"`
	PriceFeedEndpoint  string         `json:"priceFeedEndpoint"`
	PrdMktFeedEndpoint string         `json:"prdMktFeedEndpoint"`
	LowLiqFeedEndpoint string         `json:"lowLiqFeedEndpoint"`
	ChainId            int64          `json:"chainId"`
	ProxyAddr          common.Address `json:"proxyAddr"`
	MultiPayAddr       common.Address `json:"multipayAddr"`
	NodeURL            string         `json:"nodeURL"`
	PostOrderGasLimit  int64          `json:"postOrderGasLimit"`
	PriceUpdateFeeGwei int64          `json:"priceUpdateFeeGwei"`
}

type PriceFeedConfig struct {
	Network                       string        `json:"network"`
	CandleIrrelevant              []string      `json:"candleIrrelevant"`
	PriceFeedIds                  []PriceFeedId `json:"ids"`
	ThreshOnChainFeedOutdatedSec  int32         `json:"threshOnChainFeedOutdatedSec"`
	ThreshOffChainFeedOutdatedSec int32         `json:"threshOffChainFeedOutdatedSec"`
	PriceUpdateFeeGwei            int64
	SymbolToPxId                  map[string]PriceFeedId //map the symbol (BTC-USD) to PriceFeedId data
	PxIdToSymbols                 map[string][]string    //map the symbol price id 0x382738927... to one or more symbols
}

type PriceType int

const (
	PXTYPE_POLYMARKET PriceType = iota
	PXTYPE_PYTH
	PXTYPE_V2
	PXTYPE_V3
	PXTYPE_ONCHAIN
	PXTYPE_UNKNOWN
)

// Map to convert strings to AssetType. needs to align with
// priceFeedConfig on https://github.com/D8-X/sync-hub
var PriceTypeMap = map[string]PriceType{
	"polymarket": PXTYPE_POLYMARKET,
	"pyth":       PXTYPE_PYTH,
	"low-liq":    PXTYPE_V3, //legacy
	"univ2":      PXTYPE_V2,
	"univ3":      PXTYPE_V3,
	"onchain":    PXTYPE_ONCHAIN,
	"unknown":    PXTYPE_UNKNOWN,
}

// PriceType to string
func (p PriceType) ToString() string {
	for name, t := range PriceTypeMap {
		if t == p {
			return name
		}
	}
	return "unknown"
}

// UnmarshalJSON implements custom unmarshaling for PriceType
func (p *PriceType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	val, ok := PriceTypeMap[str]
	if !ok {
		return errors.New("invalid asset type: " + str)
	}
	*p = val
	return nil
}

type PriceFeedId struct {
	Symbol   string    `json:"symbol"`
	Id       string    `json:"id"`
	Type     PriceType `json:"type"`
	Origin   string    `json:"origin"`
	StorkSym string    `json:"storkSym"`
}

type PriceFeedOnChainConfig struct {
	Name          string   `json:"name"`
	RPCs          []string `json:"rpcs"`
	PxFeedAddress string   `json:"pxFeedAddress"`
	Decimals      int      `json:"decimals"`
	MaxFeedAgeSec int64    `json:"maxFeedAgeSec"`
	Type          string   `json:"type"`
}

type SettlementConfig struct {
	Description         string   `json:"description"`
	PerpFlags           *big.Int `json:"perpFlags"`
	SettleTokenDecimals int      `json:"settleTokenDecimals"`
	SettleCCY           string   `json:"settleCCY"`
	SettleCCYAddr       string   `json:"settleCCYAddr"`
	CollateralCCY       string   `json:"collateralCCY"`
	CollateralCCYAddr   string   `json:"collateralCCYAddr"`
	Triangulation       []string `json:"triangulation"`
}

func (s *SettlementConfig) UnmarshalJSON(data []byte) error {
	type Alias SettlementConfig
	aux := &struct {
		PerpFlags string `json:"perpFlags"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	perpFlagsInt := new(big.Int)
	perpFlagsInt, ok := perpFlagsInt.SetString(aux.PerpFlags, 0)
	if !ok {
		return fmt.Errorf("invalid value for PerpFlags: %s", aux.PerpFlags)
	}
	s.PerpFlags = perpFlagsInt

	return nil
}

// LoadPriceFeedConfig loads the price feed config file
// data into struct PriceFeedConfig for the network called configNetwork
// for example LoadPriceFeedConfig("config/priceFeedConfig.json", "PythEVMStable")
func LoadPriceFeedConfig(data []byte, configNetwork string) (PriceFeedConfig, error) {
	var configuration []PriceFeedConfig
	// Unmarshal the JSON data into the Configuration struct
	err := json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
		return PriceFeedConfig{}, err
	}
	j := -1
	for i := 0; i < len(configuration); i++ {
		if configuration[i].Network == configNetwork {
			j = i
			break
		}
	}
	if j == -1 {
		return PriceFeedConfig{}, errors.New("config not found")
	}
	//SymbolToId
	config := configuration[j]
	config.SymbolToPxId = make(map[string]PriceFeedId)
	for _, feed := range config.PriceFeedIds {
		config.SymbolToPxId[feed.Symbol] = PriceFeedId{
			Symbol: feed.Symbol,
			Id:     strings.TrimPrefix(feed.Id, "0x"),
			Type:   feed.Type,
			Origin: feed.Origin,
		}
	}
	//PxIdToSymbols
	config.PxIdToSymbols = make(map[string][]string)
	for _, feed := range config.PriceFeedIds {
		id := strings.TrimPrefix(feed.Id, "0x")
		if _, exists := config.PxIdToSymbols[id]; !exists {
			config.PxIdToSymbols[id] = make([]string, 0)
		}
		config.PxIdToSymbols[id] = append(config.PxIdToSymbols[id], feed.Symbol)
	}
	return config, nil
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
