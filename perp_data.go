package d8x_futures

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/big"
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

func CreateLimitOrderBookFactoryInstance(rpc *ethclient.Client, factoryAddr common.Address) *LimitOrderBookFactory {
	loFactory, err := NewLimitOrderBookFactory(factoryAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate LimitOrderBookFactory contract: %v", err)
	}
	return loFactory
}

func CreateLimitOrderBookInstance(rpc *ethclient.Client, lobAddr common.Address) *LimitOrderBook {
	lob, err := NewLimitOrderBook(lobAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate LOB contract: %v", err)
	}
	return lob
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

	var b = BlockChainConnector{
		Rpc:               rpc,
		ChainId:           config.ChainId,
		PerpetualManager:  proxy,
		SymbolMapping:     symbolMap,
		PriceFeedNetwork:  priceFeedNetwork,
		PostOrderGasLimit: config.PostOrderGasLimit,
	}

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

// GetPerpetualStaticInfoIdxFromSymbol returns the idx of the perpetual within StaticExchangeInfo,
// given the perpetual symbol (e.g., MATIC-USD-USDC). Returns -1 if not found.
func GetPerpetualStaticInfoIdxFromSymbol(exchangeInfo StaticExchangeInfo, symbol string) int {
	perpId := exchangeInfo.PerpetualSymbolToId[symbol]
	return GetPerpetualStaticInfoIdxFromId(exchangeInfo, perpId)
}

// GetPerpetualStaticInfoIdxFromId returns the idx of the perpetual within StaticExchangeInfo,
// given the perpetual id (e.g., 10001). Returns -1 if not found.
func GetPerpetualStaticInfoIdxFromId(exchangeInfo StaticExchangeInfo, perpId int32) int {
	for i, p := range exchangeInfo.Perpetuals {
		if p.Id == perpId {
			return i
		}
	}
	return -1
}

func QueryExchangeStaticInfo(conn BlockChainConnector, config Config, nest NestedPerpetualIds) StaticExchangeInfo {
	symbolsSet := make(Set)

	perpIds := nest.PerpetualIds
	pools := make([]PoolStaticInfo, len(perpIds))
	var perpetuals []PerpetualStaticInfo
	perpetualSymbolToId := make(map[string]int32)
	perpetualIdToSymbol := make(map[int32]string)
	for i, poolPerpIds := range perpIds {

		pools[i] = PoolStaticInfo{
			PoolId:              int32(i + 1),
			PoolMarginSymbol:    "", // fill later
			PoolMarginTokenAddr: nest.PoolMarginTokenAddr[i],
			ShareTokenAddr:      nest.PoolShareTokenAddr[i],
		}
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
			pools[i].PoolMarginSymbol = strings.Split(perpetuals[0].S2Symbol, "-")[1]
		case BASE:
			pools[i].PoolMarginSymbol = strings.Split(perpetuals[0].S2Symbol, "-")[0]
		default:
			pools[i].PoolMarginSymbol = strings.Split(perpetuals[0].S3Symbol, "-")[0]
		}
		// amend mapping perpetual symbol -> perpetual Id
		for _, perpStatic := range perpetuals {
			perpSymbol := perpStatic.S2Symbol + "-" + pools[i].PoolMarginSymbol
			perpetualSymbolToId[perpSymbol] = perpStatic.Id
			perpetualIdToSymbol[perpStatic.Id] = perpSymbol
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
		PerpetualSymbolToId:    perpetualSymbolToId,
		OracleFactoryAddr:      nest.OracleFactoryAddr,
		ProxyAddr:              config.ProxyAddr,
		PriceFeedInfo:          pxConfig,
		IdxPriceTriangulations: triangulations,
	}
	return xInfo
}

// Store stores the StaticExchangeInfo in a file
func (s *StaticExchangeInfo) Store(filename string) error {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}
	// Saving JSON to a file
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Load loads StaticExchangeInfo from a JSON-file created with
// the Store function.
func (s *StaticExchangeInfo) Load(filename string) error {
	// Reading JSON from file
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// Unmarshaling JSON into the struct
	err = json.Unmarshal(jsonData, s)
	if err != nil {
		return err
	}
	return nil
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

func (order *Order) ToChainType(xInfo StaticExchangeInfo, traderAddr common.Address) IClientOrderClientOrder {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, order.Symbol)
	if order.LimitPrice == 0 && order.Side == SIDE_BUY {
		order.LimitPrice = math.MaxFloat64
	}
	var flags uint32 = 0
	if order.ReduceOnly {
		flags = flags | MASK_CLOSE_ONLY
	}
	if order.KeepPositionLvg {
		flags = flags | MASK_KEEP_POS_LEVERAGE
	}
	if order.Type == ORDER_TYPE_LIMIT {
		flags = flags | MASK_LIMIT_ORDER
	} else if order.Type == ORDER_TYPE_MARKET {
		flags = flags | MASK_MARKET_ORDER
	} else if order.Type == ORDER_TYPE_STOP_LIMIT {
		flags = flags | MASK_STOP_ORDER
		flags = flags | MASK_LIMIT_ORDER
	} else if order.Type == ORDER_TYPE_STOP_MARKET {
		flags = flags | MASK_STOP_ORDER
		flags = flags | MASK_MARKET_ORDER
	}
	cOrder := IClientOrderClientOrder{
		IPerpetualId:       big.NewInt(int64(xInfo.Perpetuals[j].Id)),
		FLimitPrice:        Float64ToABDK(order.LimitPrice),
		LeverageTDR:        uint16(100 * order.Leverage),
		ExecutionTimestamp: uint32(order.ExecutionTimestamp),
		Flags:              flags,
		IDeadline:          order.Deadline,
		BrokerAddr:         order.BrokerAddr,
		FTriggerPrice:      Float64ToABDK(order.TriggerPrice),
		FAmount:            Float64ToABDK(order.Quantity),
		BrokerSignature:    order.BrokerSignature,
		ParentChildDigest1: order.parentChildOrderId1,
		ParentChildDigest2: order.parentChildOrderId2,
		TraderAddr:         traderAddr,
		BrokerFeeTbps:      uint16(order.BrokerFeeTbps),
	}

	return cOrder
}

func (scOrder *IClientOrderClientOrder) FromChainType(xInfo StaticExchangeInfo) Order {
	perpId := int32(scOrder.IPerpetualId.Int64())
	var side string
	if scOrder.FAmount.Sign() > 0 {
		side = SIDE_BUY
	} else {
		side = SIDE_SELL
	}
	var orderType string
	if scOrder.Flags&MASK_LIMIT_ORDER == MASK_LIMIT_ORDER {
		if scOrder.Flags&MASK_STOP_ORDER == MASK_STOP_ORDER {
			orderType = ORDER_TYPE_STOP_LIMIT
		} else {
			orderType = ORDER_TYPE_LIMIT
		}
	} else if scOrder.Flags&MASK_MARKET_ORDER == MASK_MARKET_ORDER {
		if scOrder.Flags&MASK_STOP_ORDER == MASK_STOP_ORDER {
			orderType = ORDER_TYPE_STOP_MARKET
		} else {
			orderType = ORDER_TYPE_MARKET
		}
	}
	var reduceOnly, keepPositionLvg = false, false
	if scOrder.Flags&MASK_CLOSE_ONLY == MASK_CLOSE_ONLY {
		reduceOnly = true
	}
	if scOrder.Flags&MASK_KEEP_POS_LEVERAGE == MASK_KEEP_POS_LEVERAGE {
		keepPositionLvg = true
	}
	order := Order{
		Symbol:              xInfo.PerpetualIdToSymbol[perpId],
		Side:                side,
		Type:                orderType,
		Quantity:            math.Abs(ABDKToFloat64(scOrder.FAmount)),
		ReduceOnly:          reduceOnly,
		LimitPrice:          ABDKToFloat64(scOrder.FLimitPrice),
		TriggerPrice:        ABDKToFloat64(scOrder.FTriggerPrice),
		KeepPositionLvg:     keepPositionLvg,
		BrokerFeeTbps:       scOrder.BrokerFeeTbps,
		BrokerAddr:          scOrder.BrokerAddr,
		BrokerSignature:     scOrder.BrokerSignature,
		Leverage:            float64(scOrder.LeverageTDR) / 100,
		Deadline:            scOrder.IDeadline,
		ExecutionTimestamp:  scOrder.ExecutionTimestamp,
		parentChildOrderId1: scOrder.ParentChildDigest1,
		parentChildOrderId2: scOrder.ParentChildDigest2,
	}
	return order
}

/*
	FLockedInValueQC             *big.Int
	FCashCC                      *big.Int
	FPositionBC                  *big.Int
	FUnitAccumulatedFundingStart *big.Int
	ILastOpenTimestamp           uint64
	FeeTbps                      uint16
	BrokerFeeTbps                uint16
	PositionId                   [16]byte*/
