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
	SportFeedEndpoint  string         `json:"sportFeedEndpoint"`
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
	SymbolToPxId                  map[string]PriceFeedId // map the symbol (BTC-USD) to PriceFeedId data
	PxIdToSymbols                 map[string][]string    // map the symbol price id 0x382738927... to one or more symbols
}

type PriceType int

const (
	PXTYPE_POLYMARKET PriceType = iota
	PXTYPE_PYTH
	PXTYPE_V2
	PXTYPE_V3
	PXTYPE_ONCHAIN
	PXTYPE_SPORT
	PXTYPE_UNKNOWN
)

// Map to convert strings to AssetType. needs to align with
// priceFeedConfig on https://github.com/D8-X/sync-hub
var PriceTypeMap = map[string]PriceType{
	"polymarket": PXTYPE_POLYMARKET,
	"pyth":       PXTYPE_PYTH,
	"univ2":      PXTYPE_V2,
	"univ3":      PXTYPE_V3,
	"sport":      PXTYPE_SPORT,
	"onchain":    PXTYPE_ONCHAIN,
	"unknown":    PXTYPE_UNKNOWN,
}

// PriceType to string
func (p PriceType) String() string {
	for name, t := range PriceTypeMap {
		if t == p {
			return name
		}
	}
	return "unknown"
}

type AssetClass int

const (
	ACLASS_CRYPTO AssetClass = iota // Crypto & Crypto.Index, https://www.pyth.network/price-feeds/crypto-index-gmci30-usd
	ACLASS_POLYMKT
	ACLASS_FX
	ACLASS_EQUITY    // equity and ETF
	ACLASS_METAL     // Metal.XAG/USD
	ACLASS_COMMODITY //	Commodities.USOILSPOT
	ACLASS_RATES     //	Rates.US1M
	ACLASS_SPORT
	ACLASS_UNKNOWN
)

// Map to convert strings to AssetType. needs to align with
// priceFeedConfig on https://github.com/D8-X/sync-hub
var AssetClassMap = map[string]AssetClass{
	"crypto":      ACLASS_CRYPTO,
	"polymarket":  ACLASS_POLYMKT,
	"sport":       ACLASS_SPORT,
	"fx":          ACLASS_FX,
	"equity":      ACLASS_EQUITY,
	"metal":       ACLASS_METAL,
	"commodities": ACLASS_COMMODITY,
	"rates":       ACLASS_RATES,
	"unknown":     ACLASS_UNKNOWN,
}

// UnmarshalJSON implements custom unmarshaling for PriceType
func (p *PriceType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	val, ok := PriceTypeMap[str]
	if !ok {
		val = PXTYPE_UNKNOWN
	}
	*p = val
	return nil
}

func OriginToAssetClass(origin string, feedType PriceType) AssetClass {
	if feedType == PXTYPE_SPORT {
		return ACLASS_SPORT
	}
	if origin == "" {
		return ACLASS_CRYPTO
	}
	if len(origin) < 2 {
		return ACLASS_UNKNOWN
	}
	if origin[0:2] == "0x" {
		return ACLASS_POLYMKT
	}

	origin = strings.ToLower(origin)
	origin = strings.Split(origin, ".")[0]
	val, ok := AssetClassMap[origin]
	if !ok {
		return ACLASS_UNKNOWN
	}
	return val
}

// AssetClass to string
func (a AssetClass) String() string {
	for name, t := range AssetClassMap {
		if t == a {
			return name
		}
	}
	return "unknown"
}

type PriceFeedId struct {
	Symbol     string     `json:"symbol"`
	Id         string     `json:"id"`
	Type       PriceType  `json:"type"`
	Origin     string     `json:"origin"`
	StorkSym   string     `json:"storkSym"`
	AssetClass AssetClass // based on origin
}

type PriceFeedOnChainConfig struct {
	Name           string   `json:"name"`
	RPCs           []string `json:"rpcs"`
	PxFeedAddress  string   `json:"pxFeedAddress"`
	Decimals       int      `json:"decimals"`
	MaxFeedAgeSec  int64    `json:"maxFeedAgeSec"`
	Type           string   `json:"type"`
	RelevantChains []int    `json:"relevantChains"`
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
func LoadPriceFeedConfig(data []byte, chainId int, configNetwork string) (PriceFeedConfig, error) {
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
	// SymbolToId
	config := configuration[j]
	config.SymbolToPxId = make(map[string]PriceFeedId)
	feedIds2 := make([]PriceFeedId, 0, len(config.PriceFeedIds))
	for j, feed := range config.PriceFeedIds {
		assetClass := OriginToAssetClass(feed.Origin, feed.Type)
		sym := feed.Symbol
		if assetClass == ACLASS_SPORT {
			// symbol of the form "NHL0-USD:84532", with the number = chainId
			// we cut off the chain id if it's the right chain, otherwise we skip
			sp := strings.Split(feed.Symbol, ":")
			if sp[1] != strconv.Itoa(chainId) {
				// skip
				continue
			}
			sym = sp[0]
		}
		config.PriceFeedIds[j].AssetClass = assetClass
		config.PriceFeedIds[j].Symbol = sym
		feedIds2 = append(feedIds2, config.PriceFeedIds[j])
		config.SymbolToPxId[sym] = PriceFeedId{
			Symbol:     sym,
			Id:         strings.TrimPrefix(feed.Id, "0x"),
			Type:       feed.Type,
			Origin:     feed.Origin,
			AssetClass: OriginToAssetClass(feed.Origin, feed.Type),
		}
	}
	config.PriceFeedIds = feedIds2
	// PxIdToSymbols
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
