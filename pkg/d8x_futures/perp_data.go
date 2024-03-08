package d8x_futures

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CreatePerpetualManagerInstance creates an instance of the Perpetual Contract.
//
// It takes an rpc and proxyAddr as argument
// It returns an instance IPerpetualManager of the proxy contract.
func CreatePerpetualManagerInstance(rpc *ethclient.Client, proxyAddr common.Address) *contracts.IPerpetualManager {
	proxy, err := contracts.NewIPerpetualManager(proxyAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate Proxy contract: %v", err)
	}
	return proxy
}

func CreateLimitOrderBookFactoryInstance(rpc *ethclient.Client, factoryAddr common.Address) *contracts.LimitOrderBookFactory {
	loFactory, err := contracts.NewLimitOrderBookFactory(factoryAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate LimitOrderBookFactory contract: %v", err)
	}
	return loFactory
}

func CreateLimitOrderBookInstance(rpc *ethclient.Client, lobAddr common.Address) *contracts.LimitOrderBook {
	lob, err := contracts.NewLimitOrderBook(lobAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate LOB contract: %v", err)
	}
	return lob
}

func CreateRPC(nodeURL string) (*ethclient.Client, error) {
	rpc, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, err
	}
	return rpc, nil
}

func CreateBlockChainConnector(pxConfig utils.PriceFeedConfig, chConf utils.ChainConfig) BlockChainConnector {
	rpc, err := CreateRPC(chConf.NodeURL)
	if err != nil {
		panic(err)
	}
	proxy := CreatePerpetualManagerInstance(rpc, chConf.ProxyAddr)
	symbolMap, err := config.GetSymbolList()
	if err != nil {
		panic(err)
	}
	priceFeedNetwork := chConf.PriceFeedNetwork
	var b = BlockChainConnector{
		Rpc:               rpc,
		ChainId:           chConf.ChainId,
		PerpetualManager:  proxy,
		SymbolMapping:     &symbolMap,
		PriceFeedNetwork:  priceFeedNetwork,
		PostOrderGasLimit: chConf.PostOrderGasLimit,
		PriceFeedConfig:   pxConfig,
	}
	b.PriceFeedConfig.PriceUpdateFeeGwei = chConf.PriceUpdateFeeGwei
	return b
}

func QueryNestedPerpetualInfo(conn BlockChainConnector) (NestedPerpetualIds, error) {

	var idxFrom uint8 = 1
	const queryLen uint8 = 5
	var lenReceived = queryLen
	var nestedPerpetualIds [][]*big.Int
	var poolShareTokenAddr []common.Address
	var poolMarginTokenAddr []common.Address
	var oracleFactory common.Address
	for {
		//([][]*big.Int, []common.Address, []common.Address, common.Address, error)
		idxTo := idxFrom + queryLen - 1
		p0, p1, p2, orcfac, err := conn.PerpetualManager.GetPoolStaticInfo(nil, idxFrom, idxTo)
		if err != nil {
			return NestedPerpetualIds{}, err
		}
		lenReceived = uint8(len(p1))
		nestedPerpetualIds = append(nestedPerpetualIds, p0...)
		poolShareTokenAddr = append(poolShareTokenAddr, p1...)
		poolMarginTokenAddr = append(poolMarginTokenAddr, p2...)
		oracleFactory = orcfac
		if lenReceived < queryLen {
			break
		}
		idxFrom = idxFrom + lenReceived
	}

	var p = NestedPerpetualIds{
		PerpetualIds:        nestedPerpetualIds,
		PoolShareTokenAddr:  poolShareTokenAddr,
		PoolMarginTokenAddr: poolMarginTokenAddr,
		OracleFactoryAddr:   oracleFactory,
	}
	return p, nil
}

// GetPerpetualStaticInfoIdxFromSymbol returns the idx of the perpetual within StaticExchangeInfo,
// given the perpetual symbol (e.g., MATIC-USD-USDC). Returns -1 if not found.
func GetPerpetualStaticInfoIdxFromSymbol(exchangeInfo *StaticExchangeInfo, symbol string) int {
	perpId := exchangeInfo.PerpetualSymbolToId[symbol]
	return GetPerpetualStaticInfoIdxFromId(exchangeInfo, perpId)
}

// GetPoolStaticInfoIdxFromSymbol returns exchangeInfo.Pools[j] for the given symbol.
// Either provide the perpetual symbol (BTC-USDC-USDC) or the pool symbol (USDC)
func GetPoolStaticInfoIdxFromSymbol(exchangeInfo *StaticExchangeInfo, symbol string) int {
	// in case the user provided a perpetual symbol, we extract the pool symbol
	parts := strings.Split(symbol, "-")
	symbol = parts[len(parts)-1]
	for j := 0; j < len(exchangeInfo.Pools); j++ {
		if symbol == exchangeInfo.Pools[j].PoolMarginSymbol {
			return j
		}
	}
	return -1
}

// GetPerpetualStaticInfoIdxFromId returns the idx of the perpetual within StaticExchangeInfo,
// given the perpetual id (e.g., 10001). Returns -1 if not found.
func GetPerpetualStaticInfoIdxFromId(exchangeInfo *StaticExchangeInfo, perpId int32) int {
	for i, p := range exchangeInfo.Perpetuals {
		if p.Id == perpId {
			return i
		}
	}
	return -1
}

func QueryExchangeStaticInfo(conn *BlockChainConnector, config *utils.ChainConfig, nest *NestedPerpetualIds) StaticExchangeInfo {
	symbolsSet := make(utils.Set)

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
			info := getterDataToPerpetualStaticInfo(&perpStatic, conn.SymbolMapping)
			perpetuals = append(perpetuals, info)
			symbolsSet.Add(info.S2Symbol)
			if info.S3Symbol != "" {
				symbolsSet.Add(info.S3Symbol)
			}
		}
		// pool currency
		j := len(perpetuals) - 1
		switch perpetuals[j].CollateralCurrencyType {
		case QUOTE:
			pools[i].PoolMarginSymbol = strings.Split(perpetuals[j].S2Symbol, "-")[1]
		case BASE:
			pools[i].PoolMarginSymbol = strings.Split(perpetuals[j].S2Symbol, "-")[0]
		default:
			pools[i].PoolMarginSymbol = strings.Split(perpetuals[j].S3Symbol, "-")[0]
		}
		// if we already have a margin symbol with the same name, we rename it using the pool id
		poolId := i + 1
		for k := 0; k < i; k++ {
			if pools[k].PoolMarginSymbol == pools[i].PoolMarginSymbol {
				pools[i].PoolMarginSymbol = pools[i].PoolMarginSymbol + strconv.Itoa(poolId)
				break
			}
		}

	}
	// amend mapping perpetual symbol -> perpetual Id
	for _, perpStatic := range perpetuals {
		poolId := perpStatic.Id / 100000
		perpSymbol := perpStatic.S2Symbol + "-" + pools[poolId-1].PoolMarginSymbol
		perpetualSymbolToId[perpSymbol] = perpStatic.Id
		perpetualIdToSymbol[perpStatic.Id] = perpSymbol
	}
	of, err := contracts.NewOracleFactory(nest.OracleFactoryAddr, conn.Rpc)
	if err != nil {
		panic(err)
	}
	pythAddr, err := of.Pyth(nil)
	if err != nil {
		panic(err)
	}
	triangulations := initPriceFeeds(&conn.PriceFeedConfig, symbolsSet)
	xInfo := StaticExchangeInfo{
		Pools:                  pools,
		Perpetuals:             perpetuals,
		PerpetualSymbolToId:    perpetualSymbolToId,
		PerpetualIdToSymbol:    perpetualIdToSymbol,
		OracleFactoryAddr:      nest.OracleFactoryAddr,
		ProxyAddr:              config.ProxyAddr,
		PriceFeedInfo:          conn.PriceFeedConfig,
		IdxPriceTriangulations: triangulations,
		PythAddr:               pythAddr,
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
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Load loads StaticExchangeInfo from a JSON-file created with
// the Store function.
func (s *StaticExchangeInfo) Load(filename string) error {
	// Reading JSON from file
	jsonData, err := os.ReadFile(filename)
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
func initPriceFeeds(pxConfig *utils.PriceFeedConfig, symbolSet utils.Set) Triangulations {
	triangulations := make(Triangulations)
	for sym := range symbolSet {
		triangulations[sym] = Triangulate(sym, pxConfig.PriceFeedIds)
	}
	return triangulations
}

func getterDataToPerpetualStaticInfo(pIn *contracts.IPerpetualInfoPerpetualStaticInfo, symMap *map[string]string) PerpetualStaticInfo {
	var poolId int32 = int32(pIn.Id.Int64()) / 100000
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
		InitialMarginRate:      utils.I32ToFloat64(pIn.FInitialMarginRate),
		MaintenanceMarginRate:  utils.I32ToFloat64(pIn.FMaintenanceMarginRate),
		CollateralCurrencyType: CollateralCCY(pIn.CollCurrencyType),
		S2Symbol:               S2Symbol,
		S3Symbol:               S3Symbol,
		LotSizeBC:              utils.ABDKToFloat64(pIn.FLotSizeBC),
		ReferralRebate:         utils.ABDKToFloat64(pIn.FReferralRebateCC),
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

func (order *Order) ToChainType(xInfo *StaticExchangeInfo, traderAddr common.Address) contracts.IClientOrderClientOrder {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, order.Symbol)
	var limitPx *big.Int
	if (order.LimitPrice == 0 || order.LimitPrice == math.MaxFloat64) && order.Side == SIDE_BUY {
		limitPx = utils.Max64x64()
	} else {
		limitPx = utils.Float64ToABDK(order.LimitPrice)
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
	var sgn float64 = 1
	if order.Side == SIDE_SELL {
		sgn = -1
	}
	amount := utils.Float64ToABDK(sgn * math.Abs(order.Quantity))
	cOrder := contracts.IClientOrderClientOrder{
		IPerpetualId:       big.NewInt(int64(xInfo.Perpetuals[j].Id)),
		FLimitPrice:        limitPx,
		LeverageTDR:        uint16(100 * order.Leverage),
		ExecutionTimestamp: uint32(order.ExecutionTimestamp),
		Flags:              flags,
		IDeadline:          order.Deadline,
		BrokerAddr:         order.BrokerAddr,
		FTriggerPrice:      utils.Float64ToABDK(order.TriggerPrice),
		FAmount:            amount,
		BrokerSignature:    order.BrokerSignature,
		ParentChildDigest1: order.ParentChildOrderId1,
		ParentChildDigest2: order.ParentChildOrderId2,
		TraderAddr:         traderAddr,
		BrokerFeeTbps:      uint16(order.BrokerFeeTbps),
	}

	return cOrder
}

func FromChainType(scOrder *contracts.IClientOrderClientOrder, xInfo *StaticExchangeInfo) Order {
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
		Quantity:            math.Abs(utils.ABDKToFloat64(scOrder.FAmount)),
		ReduceOnly:          reduceOnly,
		LimitPrice:          utils.ABDKToFloat64(scOrder.FLimitPrice),
		TriggerPrice:        utils.ABDKToFloat64(scOrder.FTriggerPrice),
		KeepPositionLvg:     keepPositionLvg,
		BrokerFeeTbps:       scOrder.BrokerFeeTbps,
		BrokerAddr:          scOrder.BrokerAddr,
		BrokerSignature:     scOrder.BrokerSignature,
		Leverage:            float64(scOrder.LeverageTDR) / 100,
		Deadline:            scOrder.IDeadline,
		ExecutionTimestamp:  scOrder.ExecutionTimestamp,
		ParentChildOrderId1: scOrder.ParentChildDigest1,
		ParentChildOrderId2: scOrder.ParentChildDigest2,
		OptTraderAddr:       scOrder.TraderAddr,
	}
	return order
}
