package d8x_futures

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CreatePerpetualManagerInstance creates an instance of the Perpetual Contract.
//
// It takes an rpc and proxyAddr as argument
// It returns an instance IPerpetualManager of the proxy contract.
func CreatePerpetualManagerInstance(rpc *ethclient.Client, proxyAddr common.Address) *IPerpetualManager {
	proxy, err := NewIPerpetualManager(proxyAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate Proxy contract: %v", err)
	}
	return proxy
}

func CreateRPC(nodeURL string) *ethclient.Client {
	rpc, err := ethclient.Dial(nodeURL)
	if err != nil {
		panic(err)
	}
	return rpc
}

func CreateBlockChainConnector(config Config) BlockChainConnector {
	rpc := CreateRPC(config.NodeURL)
	proxy := CreatePerpetualManagerInstance(rpc, config.ProxyAddr)
	symbolMap, err := readSymbolList()
	if err != nil {
		panic(err)
	}
	priceFeedNetwork := config.PriceFeedNetwork
	var b = BlockChainConnector{Rpc: rpc, PerpetualManager: proxy, SymbolMapping: symbolMap, PriceFeedNetwork: priceFeedNetwork}
	return b
}

func QueryNestedPerpetualInfo(conn BlockChainConnector) NestedPerpetualIds {

	const idxFrom = 1
	const idxTo = 20
	//([][]*big.Int, []common.Address, []common.Address, common.Address, error)
	nestedPerpetualIds, poolShareTokenAddr, poolMarginTokenAddr, oracleFactory, err := conn.PerpetualManager.GetPoolStaticInfo(nil, idxFrom, idxTo)
	if err != nil {
		panic(err)
	}
	var p = NestedPerpetualIds{
		PerpetualIds:        nestedPerpetualIds,
		PoolShareTokenAddr:  poolShareTokenAddr,
		PoolMarginTokenAddr: poolMarginTokenAddr,
		OracleFactoryAddr:   oracleFactory,
	}
	return p
}

func QueryPoolStaticInfo(conn BlockChainConnector, nest NestedPerpetualIds) StaticExchangeInfo {
	symbolsSet := make(Set)

	perpIds := nest.PerpetualIds
	var pools []PoolStaticInfo
	var perpetuals []PerpetualStaticInfo
	for i, poolPerpIds := range perpIds {

		pool := PoolStaticInfo{
			PoolId:              int32(i + 1),
			PoolMarginSymbol:    "", // fill later
			PoolMarginTokenAddr: nest.PoolMarginTokenAddr[i],
			ShareTokenAddr:      nest.PoolShareTokenAddr[i],
		}
		pools = append(pools, pool)
		// we query all perpetuals within the current pool
		perpGetterStaticInfos, err := conn.PerpetualManager.GetPerpetualStaticInfo(nil, poolPerpIds)
		if err != nil {
			panic(err)
		}

		for _, perpStatic := range perpGetterStaticInfos {
			info := getterDataToPerpetualStaticInfo(perpStatic, conn.SymbolMapping)
			perpetuals = append(perpetuals, info)
			symbolsSet.Add(info.S2Symbol)
			if info.S3Symbol != "" {
				symbolsSet.Add(info.S3Symbol)
			}
		}
		// pool currency
		switch perpetuals[0].CollateralCurrencyType {
		case QUOTE:
			pool.PoolMarginSymbol = strings.Split(perpetuals[0].S2Symbol, "-")[0]
		case BASE:
			pool.PoolMarginSymbol = strings.Split(perpetuals[0].S2Symbol, "-")[1]
		default:
			pool.PoolMarginSymbol = strings.Split(perpetuals[0].S3Symbol, "-")[0]
		}
	}
	pxConfig, err := LoadPriceFeedConfig(conn.PriceFeedNetwork)
	if err != nil {
		panic(err)
	}
	triangulations := initPriceFeeds(pxConfig, symbolsSet)
	xInfo := StaticExchangeInfo{
		Pools:                  pools,
		Perpetuals:             perpetuals,
		OracleFactoryAddr:      nest.OracleFactoryAddr,
		PriceFeedInfo:          pxConfig,
		IdxPriceTriangulations: triangulations,
	}
	return xInfo
}

// initPriceFeeds determines the triangulation for each symbol in symbolSet
func initPriceFeeds(pxConfig PriceFeedConfig, symbolSet Set) Triangulations {
	triangulations := make(Triangulations)
	for sym := range symbolSet {
		triangulations[sym] = Triangulate(sym, pxConfig)
	}
	return triangulations
}

func getterDataToPerpetualStaticInfo(pIn IPerpetualGetterPerpetualStaticInfo, symMap *map[string]string) PerpetualStaticInfo {
	var poolId int32 = int32(pIn.Id.Int64()) - (int32(pIn.Id.Int64())/1000)*1000
	base := ContractSymbolToSymbol(pIn.S2BaseCCY, symMap)
	quote := ContractSymbolToSymbol(pIn.S2QuoteCCY, symMap)
	base3 := ContractSymbolToSymbol(pIn.S3BaseCCY, symMap)
	quote3 := ContractSymbolToSymbol(pIn.S3QuoteCCY, symMap)
	S2Symbol := base + "-" + quote
	S3Symbol := ""
	if base3 != "" {
		S3Symbol = base3 + "-" + quote3
	}
	priceIds := make([]string, len(pIn.PriceIds))
	for i, uint8Array := range pIn.PriceIds {
		byteArray := make([]byte, len(uint8Array))
		for i, v := range uint8Array {
			byteArray[i] = byte(v)
		}
		priceIds[i] = hex.EncodeToString(byteArray)
	}
	var pOut = PerpetualStaticInfo{
		Id:                     int32(pIn.Id.Int64()),
		PoolId:                 poolId,
		LimitOrderBookAddr:     pIn.LimitOrderBookAddr,
		InitialMarginRate:      I32ToFloat64(pIn.FInitialMarginRate),
		MaintenanceMarginRate:  I32ToFloat64(pIn.FMaintenanceMarginRate),
		CollateralCurrencyType: CollateralCCY(pIn.CollCurrencyType),
		S2Symbol:               S2Symbol,
		S3Symbol:               S3Symbol,
		LotSizeBC:              ABDKToFloat64(pIn.FLotSizeBC),
		ReferralRebate:         ABDKToFloat64(pIn.FReferralRebateCC),
		PriceIds:               priceIds,
	}
	return pOut
}

func ContractSymbolToSymbol(cSym [4]byte, symMap *map[string]string) string {
	sym := string(bytes.TrimRight(cSym[:], "\x00"))
	if sym == "" {
		return ""
	}
	return (*symMap)[sym]
}

// readSymbolList reads the mapping from contract symbol to symbol
func readSymbolList() (*map[string]string, error) {
	jsonData, err := ioutil.ReadFile("config/symbolList.json")
	if err != nil {
		return nil, err
	}
	// Define a map to store the data
	data := make(map[string]string)

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
