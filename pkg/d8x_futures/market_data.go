package d8x_futures

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/go-multicall"
)

const MARGIN_ACCOUNT_ABI = `[{
	"inputs": [
		{
			"internalType": "uint24",
			"name": "_perpetualId",
			"type": "uint24"
		},
		{
			"internalType": "address",
			"name": "_traderAddress",
			"type": "address"
		}
	],
	"name": "getMarginAccount",
	"outputs": [
				{
					"internalType": "int128",
					"name": "fLockedInValueQC",
					"type": "int128"
				},
				{
					"internalType": "int128",
					"name": "fCashCC",
					"type": "int128"
				},
				{
					"internalType": "int128",
					"name": "fPositionBC",
					"type": "int128"
				},
				{
					"internalType": "int128",
					"name": "fUnitAccumulatedFundingStart",
					"type": "int128"
				}],
		"stateMutability": "view",
		"type": "function"}]`

const QUERY_PERP_PX_ABI = `[{
        "inputs": [
            {
                "internalType": "uint24",
                "name": "_iPerpetualId",
                "type": "uint24"
            },
            {
                "internalType": "int128",
                "name": "_fTradeAmountBC",
                "type": "int128"
            },
            {
                "internalType": "int128[2]",
                "name": "_fIndexPrice",
                "type": "int128[2]"
            },
            {
                "internalType": "uint16",
                "name": "_confTbps",
                "type": "uint16"
            },
            {
                "internalType": "uint64",
                "name": "_params",
                "type": "uint64"
            }
        ],
        "name": "queryPerpetualPrice",
        "outputs": [
            {
                "internalType": "int128",
                "name": "",
                "type": "int128"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }]`

const GET_LIQUIDATABLE_ACC_ABI = `[
	{
        "inputs": [
            {
                "internalType": "uint24",
                "name": "_perpetualId",
                "type": "uint24"
            },
            {
                "internalType": "int128[2]",
                "name": "_fIndexPrice",
                "type": "int128[2]"
            }
        ],
        "name": "getLiquidatableAccounts",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "unsafeAccounts",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]`

const POOL_SHTKN_PX_ABI = `[  {
	"inputs": [
		{
			"internalType": "uint8",
			"name": "_poolId",
			"type": "uint8"
		}
	],
	"name": "getShareTokenPriceD18",
	"outputs": [
		{
			"internalType": "uint256",
			"name": "price",
			"type": "uint256"
		}
	],
	"stateMutability": "view",
	"type": "function"
}]`

const GET_TRADER_STATE_ABI = `[{
  "name": "getTraderState",
  "type": "function",
  "stateMutability": "view",
  "inputs": [
    { "name": "_perpetualId", "type": "uint24" },
    { "name": "_traderAddress", "type": "address" },
    { "name": "_fIndexPrice", "type": "int128[2]" }
  ],
  "outputs": [
    { "name": "", "type": "int128[11]" }
  ]
}]`

type OptEndPoints struct {
	Rpc            *ethclient.Client
	PriceFeedEndPt string
}

func (sdkRo *SdkRO) GetPositionRisk(symbol string, traderAddr common.Address, optEndPt *OptEndPoints) (PositionRisk, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawGetPositionRisk(sdkRo.Info, optRpc, &traderAddr, symbol, optPyth, sdkRo.ChainConfig.PrdMktFeedEndpoint, sdkRo.ChainConfig.LowLiqFeedEndpoint)
}

// GetPositionRiskAll gets Position Risks for all perpetuals.
// It uses multicall to query traderState and constructs an array of position risks
func (sdkRo *SdkRO) GetPositionRiskAll(traderAddr common.Address, optEndPt *OptEndPoints) ([]PositionRisk, error) {
	syms := make([]string, 0, len(sdkRo.Info.PerpetualSymbolToId))
	for _, perp := range sdkRo.Info.Perpetuals {
		if perp.State == NORMAL {
			sym := sdkRo.Info.PerpetualIdToSymbol[perp.Id]
			syms = append(syms, sym)
		}
	}
	return sdkRo.GetPositionRisks(syms, traderAddr, optEndPt)
}

// GetPositionRisks uses multicall to query traderState and constructs an array of position risks
func (sdkRo *SdkRO) GetPositionRisks(symbols []string, traderAddr common.Address, optEndPt *OptEndPoints) ([]PositionRisk, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	chunks := 10
	pos := make([]PositionRisk, 0, len(symbols))
	start := 0
	end := min(len(symbols), chunks)
	for {
		p, err := RawQueryPositionRisks(
			optRpc,
			&sdkRo.Info,
			optPyth,
			sdkRo.ChainConfig.PrdMktFeedEndpoint,
			sdkRo.ChainConfig.LowLiqFeedEndpoint,
			symbols[start:end],
			traderAddr,
		)
		if err != nil {
			return nil, err
		}
		pos = append(pos, p...)
		start = end
		end = min(len(symbols), end+chunks)
		if start == len(symbols) {
			break
		}
	}
	return pos, nil
}

func (sdkRo *SdkRO) QueryPerpetualState(perpetualIds []int32, optEndPt *OptEndPoints) ([]PerpetualState, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryPerpetualState(optRpc, sdkRo.Info, perpetualIds, optPyth, sdkRo.ChainConfig.PrdMktFeedEndpoint, sdkRo.ChainConfig.LowLiqFeedEndpoint)
}

func (sdkRo *SdkRO) ApproximateOrderBook(symbol string, optEndPt *OptEndPoints) (*OrderBook, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	// example: https://api.binance.com/api/v1/depth?symbol=BTCUSDC&limit=5000
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdkRo.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	if isLowLiqPerp(&sdkRo.Info.Perpetuals[j]) {
		// we calculate off-chain
		return sdkRo.approximateOrderBookOffChain(&sdkRo.Info.Perpetuals[j])
	} else {
		return sdkRo.approximateOrderBookOnChain(&sdkRo.Info.Perpetuals[j], symbol, optRpc, optPyth)
	}
}

func (sdkRo *SdkRO) approximateOrderBookOffChain(pinfo *PerpetualStaticInfo) (*OrderBook, error) {
	m := pinfo.LotSizeBC * 10
	N := 50
	mx := m * 5000
	inc := float64(mx-m) / float64(N)
	var ob OrderBook
	ob.Asks = make([]OrderBookLevel, N)
	ob.Bids = make([]OrderBookLevel, N)
	p, err := fetchPythPrices(pinfo.PriceIds, "", "", sdkRo.ChainConfig.LowLiqFeedEndpoint)
	if err != nil {
		return nil, err
	}
	mid, halfBa, encBk, err := extractLowLiqParams(p[0])
	if err != nil {
		return nil, err
	}
	a0, m0, a1, m1 := decodeOrderBook(encBk)
	var lastQ float64
	for j := 0; j < N; j++ {
		// ask side
		q := m + float64(j)*inc
		ob.Asks[N-1-j].Price = mid + halfBa + math.Exp(a1+m1*q)
		ob.Asks[N-1-j].Quantity = q - lastQ
		lastQ = q
		// bid side
		ob.Bids[j].Price = mid - halfBa - math.Exp(a0+m0*q)
		ob.Bids[j].Quantity = ob.Asks[N-1-j].Quantity
	}
	ob.TimestampMs = time.Now().UnixMilli()
	return &ob, nil
}

func (sdkRo *SdkRO) approximateOrderBookOnChain(pinfo *PerpetualStaticInfo, symbol string, optRpc *ethclient.Client, optPyth string) (*OrderBook, error) {
	m := pinfo.LotSizeBC * 10
	samples := []float64{1, 100, 250, 2500, 5000}
	tradeAmt := make([]float64, len(samples)*2)
	for j := range samples {
		tradeAmt[len(samples)+j] = m * samples[j]
		tradeAmt[len(samples)-1-j] = -m * samples[j]
	}
	prices, err := RawQueryPerpetualPriceTuple(optRpc,
		&sdkRo.Info,
		optPyth,
		sdkRo.ChainConfig.PrdMktFeedEndpoint,
		sdkRo.ChainConfig.LowLiqFeedEndpoint,
		symbol,
		tradeAmt,
	)
	if err != nil {
		return nil, err
	}
	// we want 50 samples per side
	N := 50
	inc := (samples[len(samples)-1] - samples[0]) * m / float64(N)
	var ob OrderBook
	ob.Asks = make([]OrderBookLevel, N)
	ob.Bids = make([]OrderBookLevel, N)
	var lastQ float64
	for j := 0; j < N; j++ {
		// ask side
		q := (m + float64(j)*inc)
		ob.Asks[N-1-j].Price = interpolate(q, tradeAmt, prices)
		ob.Asks[N-1-j].Quantity = q - lastQ
		lastQ = q
		// bid side
		q = -q
		ob.Bids[j].Price = interpolate(q, tradeAmt, prices)
		ob.Bids[j].Quantity = ob.Asks[N-1-j].Quantity

	}
	ob.TimestampMs = time.Now().UnixMilli()

	return &ob, nil
}

// interpolate interpolates linearly at s for the function
// represented by (x, y)
func interpolate(s float64, x, y []float64) float64 {
	if s < x[0] {
		return y[0]
	}
	if s > x[len(x)-1] {
		return y[len(y)-1]
	}
	var j int
	for j = 1; j < len(x); j++ {
		if s < x[j] {
			break
		}
	}
	dx := s - x[j-1]
	return y[j-1] + dx*(y[j]-y[j-1])/(x[j]-x[j-1])
}

func (sdkRo *SdkRO) QueryPerpetualPrices(symbol string, tradeAmt []float64, optEndPt *OptEndPoints) ([]float64, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryPerpetualPriceTuple(optRpc,
		&sdkRo.Info,
		optPyth,
		sdkRo.ChainConfig.PrdMktFeedEndpoint,
		sdkRo.ChainConfig.LowLiqFeedEndpoint,
		symbol,
		tradeAmt,
	)
}

func extractEndpoints(sdkRo *SdkRO, optEndPt *OptEndPoints) (*ethclient.Client, string) {
	optRpc := sdkRo.Conn.Rpc
	optPyth := sdkRo.ChainConfig.PriceFeedEndpoint
	if optEndPt == nil {
		return optRpc, optPyth
	}
	if optEndPt.Rpc != nil {
		optRpc = optEndPt.Rpc
	}
	if optEndPt.PriceFeedEndPt != "" {
		optPyth = optEndPt.PriceFeedEndPt
	}
	return optRpc, optPyth
}

func (sdkRo *SdkRO) QueryPoolStates(optRpc *ethclient.Client) ([]PoolState, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryPoolStates(optRpc, sdkRo.Info)
}

func (sdkRo *SdkRO) QueryOpenOrders(symbol string, traderAddr common.Address, optRpc *ethclient.Client) ([]Order, []string, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryOpenOrders(optRpc, sdkRo.Info, symbol, traderAddr)
}

func (sdkRo *SdkRO) QueryAllOpenOrders(symbol string, optRpc *ethclient.Client) (*OpenOrders, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryAllOpenOrders(optRpc, sdkRo.Info, symbol)
}

func (sdkRo *SdkRO) QueryNumOrders(symbol string, optRpc *ethclient.Client) (int64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryNumOrders(optRpc, sdkRo.Info, symbol)
}

func (sdkRo *SdkRO) QueryOpenOrderRange(symbol string, from, to int, optRpc *ethclient.Client) (*OpenOrders, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryOpenOrderRange(optRpc, sdkRo.Info, symbol, from, to)
}

func (sdkRo *SdkRO) QueryOrderStatus(symbol string, traderAddr common.Address, orderDigest string, optRpc *ethclient.Client) (OrderStatus, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryOrderStatus(optRpc, sdkRo.Info, traderAddr, orderDigest, symbol)
}

func (sdkRo *SdkRO) QueryMaxTradeAmount(symbol string, currentPositionNotional float64, isBuy bool, optRpc *ethclient.Client) (float64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryMaxTradeAmount(optRpc, sdkRo.Info, currentPositionNotional, symbol, isBuy)
}

func (sdkRo *SdkRO) QueryMarginAccounts(symbol string, traderAddrs []common.Address, optRpc *ethclient.Client) ([]MarginAccount, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryMarginAccounts(optRpc, &sdkRo.Info, symbol, traderAddrs)
}

func (sdkRo *SdkRO) QueryTraderVolume(poolId int32, traderAddr common.Address, optRpc *ethclient.Client) (float64, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryTraderVolume(optRpc, sdkRo.Info, traderAddr, poolId)
}

func (sdkRo *SdkRO) QueryExchangeFeeTbpsForTrader(poolId int32, traderAddr common.Address, brokerAddr common.Address, optRpc *ethclient.Client) (uint16, error) {
	if optRpc == nil {
		optRpc = sdkRo.Conn.Rpc
	}
	return RawQueryExchangeFeeTbpsForTrader(optRpc, sdkRo.Info, poolId, traderAddr, brokerAddr)
}

func (sdkRo *SdkRO) GetMinimalPositionSize(symbol string) (float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdkRo.Info, symbol)
	if j == -1 {
		return 0, errors.New("GetMinimalPositionSize: no perpetual " + symbol)
	}
	v := sdkRo.Info.Perpetuals[j]
	return v.LotSizeBC * 10.0, nil
}

// QueryLiquidatableAccounts identifies all liquidatable accounts in the given perpetuals
func (sdkRo *SdkRO) QueryLiquidatableAccounts(perpId int32, optEndPt *OptEndPoints) ([]common.Address, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryLiquidatableAccounts(optRpc, &sdkRo.Info, perpId, optPyth, sdkRo.ChainConfig.PrdMktFeedEndpoint, sdkRo.ChainConfig.LowLiqFeedEndpoint)
}

// QueryLiquidatableAccountsInPool identifies all traders that can be liquidated in all perpetuals of
// the given pool.
func (sdkRo *SdkRO) QueryLiquidatableAccountsInPool(poolId int32, optEndPt *OptEndPoints) ([]LiquidatableAccounts, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryLiquidatableAccountsInPool(optRpc,
		&sdkRo.Info,
		poolId,
		optPyth,
		sdkRo.ChainConfig.PrdMktFeedEndpoint,
		sdkRo.ChainConfig.LowLiqFeedEndpoint,
	)
}

func (sdkRo *SdkRO) FetchPricesForPerpetualId(id int32, optPythEndPt string) (PerpetualPriceInfo, error) {
	if optPythEndPt == "" {
		optPythEndPt = sdkRo.ChainConfig.PriceFeedEndpoint
	}
	return RawFetchPricesForPerpetualId(
		sdkRo.Info,
		id,
		optPythEndPt,
		sdkRo.ChainConfig.PrdMktFeedEndpoint,
		sdkRo.ChainConfig.LowLiqFeedEndpoint,
	)
}

func (sdkRo *SdkRO) FetchPricesForPerpetual(symbol string, optPythEndPt string) (PerpetualPriceInfo, error) {
	if optPythEndPt == "" {
		optPythEndPt = sdkRo.ChainConfig.PriceFeedEndpoint
	}
	return RawFetchPricesForPerpetual(sdkRo.Info,
		symbol,
		optPythEndPt,
		sdkRo.ChainConfig.PrdMktFeedEndpoint,
		sdkRo.ChainConfig.LowLiqFeedEndpoint,
	)
}

// FetchCollToSettlePx gets the conversion from collateral price to settlement price
func (sdkRo *SdkRO) FetchCollToSettlePx(symbol string, optPythEndPt string) (float64, error) {
	info := &sdkRo.Info
	j := GetPoolStaticInfoIdxFromSymbol(info, symbol)
	if j == -1 {
		return 0, errors.New("FetchCollToSettlePx: no perpetual " + symbol)
	}
	settleCCY := info.Pools[j].PoolSettleSymbol
	mgnSymbol := info.Pools[j].PoolMarginSymbol
	if settleCCY == mgnSymbol {
		return float64(1), nil
	}
	if optPythEndPt == "" {
		optPythEndPt = sdkRo.ChainConfig.PriceFeedEndpoint
	}
	pxMap, _, err := fetchIndexPricesForPerpetual(info, j, optPythEndPt, "", sdkRo.ChainConfig.LowLiqFeedEndpoint)
	if err != nil {
		return 0, err
	}
	sym := mgnSymbol + "-" + settleCCY
	px, exists := pxMap[sym]
	if !exists {
		return 0, fmt.Errorf("symbol %s does not exist in index prices", sym)
	}
	return px.Px, nil
}

// GetSettleTokenBalance retreives the balance of the settlement token for traderAddr and the given perpetual or pool symbol
func (sdkRo *SdkRO) GetSettleTokenBalance(symbol string, traderAddr common.Address, optRpc *ethclient.Client) (float64, error) {
	tknAddr, err := RawGetSettleTknAddr(&sdkRo.Info, symbol)
	if err != nil {
		return 0, err
	}
	rpc := sdkRo.Conn.Rpc
	if optRpc != nil {
		rpc = optRpc
	}
	return RawGetTknBalance(tknAddr, traderAddr, rpc)
}

// GetPoolShareTknBalance returns the amount of pool share tokens that the lpAddr holds in decimals
func (sdkRo *SdkRO) GetPoolShareTknBalance(poolId int, lpAddr common.Address, optRpc *ethclient.Client) (float64, error) {
	rpc := sdkRo.Conn.Rpc
	if optRpc != nil {
		rpc = optRpc
	}
	return RawQueryPoolShTknBalance(lpAddr, poolId, sdkRo.Info, rpc)
}

func (sdkRo *SdkRO) GetPoolShareTknPrice(poolIds []int, optRpc *ethclient.Client) ([]float64, error) {
	rpc := sdkRo.Conn.Rpc
	if optRpc != nil {
		rpc = optRpc
	}
	return RawGetPoolShTknPrice(rpc, poolIds, sdkRo.Info)
}

// GetPriceId returns the price id information for the given index,
// e.g BTC-USD
func (sdkRo *SdkRO) GetPriceId(idxSymbol string) (utils.PriceFeedId, error) {
	for _, info := range sdkRo.PxConfig.PriceFeedIds {
		if info.Symbol == idxSymbol {
			return info, nil
		}
	}
	return utils.PriceFeedId{}, fmt.Errorf("no price info found for index symbol %s", idxSymbol)
}

// Allowance checks the allowance of the given address to spend settlement tokens for the given
// pool (via symbol) on the manager. Returns the value in decimals and the decimal-N value (big-int).
// Symbol is a pool symbol like "USDC" (or perpetual symbol like ETH-USDC-USDC works too)
// Use ApproveTknSpending to approve if allowance not enough.
func (sdkRo *SdkRO) Allowance(symbol string, user common.Address, optRpc *ethclient.Client) (float64, *big.Int, error) {
	tknAddr, err := RawGetSettleTknAddr(&sdkRo.Info, symbol)
	if err != nil {
		return 0, nil, err
	}
	rpc := sdkRo.Conn.Rpc
	if optRpc != nil {
		rpc = optRpc
	}
	erc20Instance, err := contracts.NewErc20(tknAddr, rpc)
	if err != nil {
		return 0, nil, fmt.Errorf("error creating instance of token %s: %s", tknAddr.String(), err.Error())
	}
	n, err := erc20Instance.Decimals(nil)
	if err != nil {
		return 0, nil, err
	}
	a, err := erc20Instance.Allowance(nil, user, sdkRo.Info.ProxyAddr)
	if err != nil {
		return 0, nil, fmt.Errorf("error getting allowance for token %s:%s", tknAddr.String(), err.Error())
	}
	amt := utils.DecNToFloat(a, n)
	return amt, a, nil
}

// IsPrdMktPerp returns true if perpetual with symbol of the form
// BTC-USD-USDC is a prediction market perpetual
func (sdkRo *SdkRO) IsPrdMktPerp(symbol string) (bool, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdkRo.Info, symbol)
	if j == -1 {
		return false, errors.New("IsPrdMktPerp: no perpetual " + symbol)
	}
	v := sdkRo.Info.Perpetuals[j]
	return isPrdMktPerp(&v), nil
}

// IsLowLiqPerp returns true if perpetual with symbol of the form
// BTC-USD-USDC is a lowliq market perpetual
func (sdkRo *SdkRO) IsLowLiqPerp(symbol string) (bool, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdkRo.Info, symbol)
	if j == -1 {
		return false, errors.New("IsLowLiqPerp: no perpetual " + symbol)
	}
	v := sdkRo.Info.Perpetuals[j]
	return isLowLiqPerp(&v), nil
}

func RawGetMarginTknAddr(xInfo *StaticExchangeInfo, symbol string) (common.Address, error) {
	j := GetPoolStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return common.Address{}, errors.New("RawGetMarginTknAddr: no perpetual " + symbol)
	}
	pool := xInfo.Pools[j]
	poolId := pool.PoolId

	return (xInfo.Pools[poolId-1].PoolMarginTokenAddr), nil
}

func RawGetSettleTknAddr(xInfo *StaticExchangeInfo, symbol string) (common.Address, error) {
	j := GetPoolStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return common.Address{}, errors.New("RawGetMarginTknAddr: no perpetual " + symbol)
	}
	pool := xInfo.Pools[j]
	poolId := pool.PoolId

	return (xInfo.Pools[poolId-1].PoolSettleTokenAddr), nil
}

func RawGetPositionRisk(
	xInfo StaticExchangeInfo,
	rpc *ethclient.Client,
	traderAddr *common.Address,
	symbol string,
	pythEndpoint string,
	prdMktEndpoint string,
	lowLiqFeedEndpoint string,
) (PositionRisk, error) {
	priceData, err := RawFetchPricesForPerpetual(xInfo, symbol, pythEndpoint, prdMktEndpoint, lowLiqFeedEndpoint)
	if err != nil {
		return PositionRisk{}, err
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		panic("symbol does not exist in static perpetual info")
	}
	indexPrices := [2]*big.Int{utils.Float64ToABDK(priceData.S2Price), utils.Float64ToABDK(priceData.S3Price)}

	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return PositionRisk{}, err
	}
	traderState, err := proxy.GetTraderState(
		nil,
		new(big.Int).SetInt64(int64(xInfo.Perpetuals[j].Id)),
		*traderAddr,
		indexPrices)
	if err != nil {
		return PositionRisk{}, fmt.Errorf("error fetching margin account: %v", err.Error())
	}
	return traderStateToPositionRisk(
		symbol,
		&xInfo.Perpetuals[j],
		priceData.S2Price,
		priceData.S3Price,
		traderState[:],
	), nil
}

func traderStateToPositionRisk(sym string, pInfo *PerpetualStaticInfo, s2, s3 float64, traderState []*big.Int) PositionRisk {
	const idxAvailableCashCC = 2
	const idxCash = 3
	const idxNotional = 4
	const idxLockedIn = 5
	const idxMarkPrice = 8
	const idxLvg = 7
	const idxS3 = 9
	var S3Liq float64
	lockedInValue := utils.ABDKToFloat64(traderState[idxLockedIn])
	cashCC := utils.ABDKToFloat64(traderState[idxCash])
	posBC := utils.ABDKToFloat64(traderState[idxNotional])
	Sm := utils.ABDKToFloat64(traderState[idxMarkPrice])
	S3 := utils.ABDKToFloat64(traderState[idxS3])
	unpaidFundingCC := utils.ABDKToFloat64(traderState[idxAvailableCashCC]) - cashCC
	unpaidFundingQC := unpaidFundingCC
	S2Liq := RawCalculateLiquidationPrice(pInfo.CollateralCurrencyType, lockedInValue, posBC, cashCC, pInfo.MaintenanceMarginRate, S3, Sm)
	switch pInfo.CollateralCurrencyType {
	case BASE:
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / s2
	case QUANTO:
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / s3
		S3Liq = utils.ABDKToFloat64(traderState[idxS3])
	}
	// floor at zero
	S2Liq = math.Max(0, S2Liq)
	S3Liq = math.Max(0, S3Liq)

	var side Side
	if posBC == 0 {
		side = SIDE_CLOSED
	} else if posBC < 0 {
		side = SIDE_SELL
	} else {
		side = SIDE_BUY
	}
	var entryPrice, leverage, pnl, liqLeverage float64
	if posBC != 0 {
		entryPrice = math.Abs(lockedInValue / posBC)
		leverage = utils.ABDKToFloat64(traderState[idxLvg])
		pnl = posBC*Sm - lockedInValue + unpaidFundingQC
		liqLeverage = 1 / pInfo.MaintenanceMarginRate
	} else {
		S3Liq = 0
	}
	m := PositionRisk{
		Symbol:                         sym,
		PositionNotionalBaseCCY:        math.Abs(posBC),
		Side:                           side,
		EntryPrice:                     entryPrice,
		Leverage:                       leverage,
		MarkPrice:                      Sm,
		UnrealizedPnlQuoteCCY:          pnl,
		UnrealizedFundingCollateralCCY: unpaidFundingCC,
		CollateralCC:                   cashCC,
		LiquidationPrice:               [2]float64{S2Liq, S3Liq},
		LiquidationLvg:                 liqLeverage,
		CollToQuoteConversion:          S3,
	}
	return m
}

// QueryPerpetualState collects PerpetualState by calling the off-chain prices and
// blockchain queries. endpoint is the address to get prices from
func RawQueryPerpetualState(
	rpc *ethclient.Client,
	xInfo StaticExchangeInfo,
	perpetualIds []int32, pythEndpoint, prdMktEndpoint, lowLiqEndpoint string,
) ([]PerpetualState, error) {
	bigIntSlice := make([]*big.Int, len(perpetualIds))
	for i, id := range perpetualIds {
		bigIntSlice[i] = big.NewInt(int64(id))
	}
	// perpetual data via blockchain

	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, err
	}
	perps, err := proxy.GetPerpetuals(nil, bigIntSlice)
	if err != nil {
		return nil, err
	}
	// gather perpetual index prices (offchain REST)
	pxInfo := make([]*big.Int, len(perpetualIds)*2)
	pxInfoFloat := make([]float64, len(perpetualIds)*2)
	for i := range perpetualIds {
		p, err := RawFetchPricesForPerpetualId(xInfo, perpetualIds[i], pythEndpoint, prdMktEndpoint, lowLiqEndpoint)
		if err != nil {
			return nil, err
		}
		pxInfoFloat[i*2] = p.S2Price
		pxInfoFloat[i*2+1] = p.S3Price
		pxInfo[i*2] = utils.Float64ToABDK(p.S2Price)
		pxInfo[i*2+1] = utils.Float64ToABDK(p.S3Price)
	}
	// midprices via blockchain query
	pxMid, err := proxy.QueryMidPrices(nil, bigIntSlice, pxInfo)
	if err != nil {
		return []PerpetualState{}, err
	}

	perpStates := make([]PerpetualState, len(perps))
	for i, perpData := range perps {
		perpStates[i].Id = int32(perpData.Id.Int64())
		perpStates[i].State = PerpetualStateEnum(perpData.State)
		perpStates[i].IndexPrice = pxInfoFloat[i*2]
		perpStates[i].CollToQuoteIndexPrice = pxInfoFloat[i*2+1]
		perpStates[i].MarkPrice = pxInfoFloat[i*2] * (1 + utils.ABDKToFloat64(perpData.CurrentMarkPremiumRate.FPrice))
		perpStates[i].CurrentFundingRateBps = utils.ABDKToFloat64(perpData.FCurrentFundingRate)
		perpStates[i].OpenInterestBC = utils.ABDKToFloat64(perpData.FOpenInterest)
		perpStates[i].MidPrice = utils.ABDKToFloat64(pxMid[i])
	}
	return perpStates, nil
}

// RawQueryPoolStates gathers the pool states of all pools by querying the blockchain in
// chunks of 10 pools
func RawQueryPoolStates(rpc *ethclient.Client, xInfo StaticExchangeInfo) ([]PoolState, error) {
	numPools := len(xInfo.Pools)
	poolStates := make([]PoolState, numPools)
	// we query a maximum of 10 pools at once
	const MAXPOOLS = 10
	iterations := (numPools + MAXPOOLS - 1) / MAXPOOLS
	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, err
	}
	for i := 0; i < iterations; i++ {
		from := i * MAXPOOLS
		to := (i+1)*MAXPOOLS - 1
		if to > numPools {
			to = numPools
		}

		pools, err := proxy.GetLiquidityPools(nil, uint8(from+1), uint8(to+1))
		if err != nil {
			return nil, err
		}
		pIdx := 0
		for j := from; j < to; j++ {
			poolStates[j].Id = int32(pools[pIdx].Id)
			poolStates[j].DefaultFundCashCC = utils.ABDKToFloat64(pools[pIdx].FDefaultFundCashCC)
			poolStates[j].IsRunning = pools[pIdx].IsRunning
			poolStates[j].PnlParticipantCashCC = utils.ABDKToFloat64(pools[pIdx].FPnLparticipantsCashCC)
			poolStates[j].TotalSupplyShareToken = utils.ABDKToFloat64(pools[pIdx].TotalSupplyShareToken)
			poolStates[j].TotalTargetAMMFundSizeCC = utils.ABDKToFloat64(pools[pIdx].FTargetAMMFundSize)
			pIdx++
		}
	}
	return poolStates, nil
}

func RawQueryAllOpenOrders(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string) (*OpenOrders, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)

	from := 0
	count := 500
	var orders OpenOrders
	orders.HashIndex = make(map[string]int)
	zeroAddr := common.Address{}
outerLoop:
	for {
		currOrders, err := lob.PollRange(nil, big.NewInt(int64(from)), big.NewInt(int64(count)))
		if err != nil {
			return nil, err
		}
		from = from + count
		for k, corder := range currOrders.Orders {
			if corder.TraderAddr == zeroAddr {
				break outerLoop
			} else {
				order := FromChainType(&corder, &xInfo)
				orders.Orders = append(orders.Orders, order)
				strDigests := "0x" + common.Bytes2Hex(currOrders.OrderHashes[k][:])
				orders.HashIndex[strDigests] = k
				orders.OrderHashes = append(orders.OrderHashes, strDigests)
			}
		}
		orders.SubmittedTs = currOrders.SubmittedTs
		from = from + count
	}

	return &orders, nil
}

func RawQueryNumOrders(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string) (int64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return 0, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	count, err := lob.OrderCount(nil)
	if err != nil {
		return 0, err
	}
	return count.Int64(), nil
}

func RawQueryOpenOrderRange(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string, from, to int) (*OpenOrders, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	ooSc, err := lob.PollRange(nil, big.NewInt(int64(from)), big.NewInt(int64(to-from)))
	var orders OpenOrders
	orders.HashIndex = make(map[string]int)
	for k, scOrder := range ooSc.Orders {
		orders.RawOrders = append(orders.RawOrders, scOrder)
		orders.Orders = append(orders.Orders, FromChainType(&scOrder, &xInfo))
		strDigests := "0x" + common.Bytes2Hex(ooSc.OrderHashes[k][:])
		orders.HashIndex[strDigests] = k
		orders.OrderHashes = append(orders.OrderHashes, strDigests)
	}
	orders.SubmittedTs = ooSc.SubmittedTs
	return &orders, err
}

func RawQueryOpenOrders(rpc *ethclient.Client, xInfo StaticExchangeInfo, symbol string, traderAddr common.Address) ([]Order, []string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return []Order{}, []string{}, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)

	from := 0
	count := 15
	clientOrders := []contracts.IClientOrderClientOrder{}
	zeroAddr := common.Address{}
outerLoop:
	for {
		currOrders, err := lob.GetOrders(nil, traderAddr, big.NewInt(int64(from)), big.NewInt(int64(count)))
		if err != nil {
			return []Order{}, []string{}, err
		}
		from = from + count
		for _, order := range currOrders {
			if order.TraderAddr == zeroAddr {
				break outerLoop
			} else {
				clientOrders = append(clientOrders, order)
			}
		}
	}
	digests, err := lob.LimitDigestsOfTrader(
		nil,
		traderAddr,
		big.NewInt(int64(0)),
		big.NewInt(int64(len(clientOrders))),
	)

	// format digests into strings
	strDigests := make([]string, len(digests))
	// format ClientOrders into orders
	orders := make([]Order, len(clientOrders))
	for i, d := range digests {
		strDigests[i] = "0x" + common.Bytes2Hex(d[:])
		orders[i] = FromChainType(&clientOrders[i], &xInfo)
	}
	if err != nil {
		return []Order{}, []string{}, err
	}
	return orders, strDigests, nil
}

func RawQueryOrderStatus(rpc *ethclient.Client, xInfo StaticExchangeInfo, traderAddr common.Address, orderDigest string, symbol string) (OrderStatus, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return ORDER_STATUS_UNKNOWN, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	// convert digest string to bytes 32
	bytesDigest := common.Hex2Bytes(strings.TrimPrefix(orderDigest, "0x"))
	var orderDigest32 [32]byte
	copy(orderDigest32[:], bytesDigest)
	status, err := lob.GetOrderStatus(nil, orderDigest32)
	if err != nil {
		return ORDER_STATUS_UNKNOWN, err
	}
	return OrderStatus(status), nil
}

// RawQueryPerpetualPriceTuple queries prices for different trade amounts for the given perpetual. The perpetual
// is given by symbol (e.g., "ETH-USD-WEETH")
func RawQueryPerpetualPriceTuple(
	client *ethclient.Client,
	xInfo *StaticExchangeInfo,
	pythEndpoint string,
	prdMktEndpoint string,
	lowLiqEndpoint string,
	symbol string,
	tradeAmt []float64,
) ([]float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
	pxInfo, err := fetchPerpetualPriceInfo(xInfo, j, pythEndpoint, prdMktEndpoint, lowLiqEndpoint)
	if err != nil {
		return nil, errors.New("RawAddCollateral: failed fetching oracle prices " + err.Error())
	}
	pricesAbdk := [2]*big.Int{utils.Float64ToABDK(pxInfo.S2Price), utils.Float64ToABDK(pxInfo.S3Price)}
	caller, err := multicall.New(client)
	if err != nil {
		return nil, err
	}
	type priceOutput struct {
		PriceAbdk *big.Int
	}

	contract, err := multicall.NewContract(QUERY_PERP_PX_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	calls := make([]*multicall.Call, 0, len(tradeAmt))
	for _, amt := range tradeAmt {
		taAbdk := utils.Float64ToABDK(amt)
		c := contract.NewCall(new(priceOutput), "queryPerpetualPrice", perpId, taAbdk, pricesAbdk, pxInfo.Conf, pxInfo.CLOBParams)
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	prices := make([]float64, 0, len(tradeAmt))
	for _, call := range res {
		px := utils.ABDKToFloat64(call.Outputs.(*priceOutput).PriceAbdk)
		prices = append(prices, px)
	}
	return prices, nil
}

func RawQueryPositionRisks(
	client *ethclient.Client,
	xInfo *StaticExchangeInfo,
	pythEndpoint string,
	prdMktEndpoint string,
	lowLiqEndpoint string,
	symbols []string,
	trader common.Address,
) ([]PositionRisk, error) {
	perpIds := make([]*big.Int, len(symbols))
	prices := make([]PerpetualPriceInfo, len(symbols))
	var err error
	for k, s := range symbols {
		j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, s)
		if j == -1 {
			return nil, fmt.Errorf("Symbol " + s + " does not exist in static perpetual info")
		}
		perpIds[k] = big.NewInt(int64(xInfo.Perpetuals[j].Id))
		prices[k], err = fetchPerpetualPriceInfo(xInfo, j, pythEndpoint, prdMktEndpoint, lowLiqEndpoint)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch price info: %v", err)
		}
	}
	caller, err := multicall.New(client)
	if err != nil {
		return nil, err
	}

	contract, err := multicall.NewContract(GET_TRADER_STATE_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	calls := make([]*multicall.Call, 0, len(symbols))
	type TraderStateOutput struct {
		Output [11]*big.Int
	}
	for j := range symbols {
		outputs := new(TraderStateOutput)
		idxPx := [2]*big.Int{
			utils.Float64ToABDK(prices[j].S2Price),
			utils.Float64ToABDK(prices[j].S3Price),
		}
		c := contract.NewCall(outputs, "getTraderState", perpIds[j], trader, idxPx)
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	accounts := make([]PositionRisk, len(symbols))
	for k, call := range res {
		outputs := call.Outputs.(*TraderStateOutput)
		j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbols[k])
		ts := traderStateToPositionRisk(
			symbols[k],
			&xInfo.Perpetuals[j],
			prices[k].S2Price,
			prices[k].S3Price,
			outputs.Output[:],
		)
		accounts[k] = ts
	}
	return accounts, nil
}

func RawQueryMarginAccounts(client *ethclient.Client, xInfo *StaticExchangeInfo, symbol string, traders []common.Address) ([]MarginAccount, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
	caller, err := multicall.New(client)
	if err != nil {
		return nil, err
	}

	type marginOutput struct {
		FLockedInValueQC             *big.Int
		FCashCC                      *big.Int
		FPositionBC                  *big.Int
		FUnitAccumulatedFundingStart *big.Int
	}
	contract, err := multicall.NewContract(MARGIN_ACCOUNT_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	calls := make([]*multicall.Call, 0, len(traders))

	for _, t := range traders {
		c := contract.NewCall(new(marginOutput), "getMarginAccount", perpId, t)
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	accounts := make([]MarginAccount, 0, len(traders))
	for _, call := range res {
		m := MarginAccount{
			FLockedInValueQC:             utils.ABDKToFloat64(call.Outputs.(*marginOutput).FLockedInValueQC),
			FCashCC:                      utils.ABDKToFloat64(call.Outputs.(*marginOutput).FCashCC),
			FPositionBC:                  utils.ABDKToFloat64(call.Outputs.(*marginOutput).FPositionBC),
			FUnitAccumulatedFundingStart: utils.ABDKToFloat64(call.Outputs.(*marginOutput).FUnitAccumulatedFundingStart),
		}
		accounts = append(accounts, m)
	}
	return accounts, nil
}

// Not compatible with zkevm deployment
func RawGetPerpetualData(rpc *ethclient.Client, xInfo *StaticExchangeInfo, symbol string) (*contracts.PerpStoragePerpetualData, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, err
	}
	perpData, err := proxy.GetPerpetual(nil, big.NewInt(int64(xInfo.Perpetuals[j].Id)))
	if err != nil {
		return nil, err
	}
	return &perpData, nil
}

func RawQueryMaxTradeAmount(rpc *ethclient.Client, xInfo StaticExchangeInfo, currentPositionNotional float64, symbol string, isBuy bool) (float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return 0, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	p := utils.Float64ToABDK(currentPositionNotional)
	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return 0, err
	}
	t, err := proxy.GetMaxSignedOpenTradeSizeForPos(nil, big.NewInt(int64(xInfo.Perpetuals[j].Id)), p, isBuy)
	if err != nil {
		return 0, err
	}
	return utils.ABDKToFloat64(t), nil
}

func RawQueryTraderVolume(rpc *ethclient.Client, xInfo StaticExchangeInfo, traderAddr common.Address, poolId int32) (float64, error) {
	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return 0, err
	}
	vol, err := proxy.GetCurrentTraderVolume(nil, uint8(poolId), traderAddr)
	if err != nil {
		return 0, err
	}
	return utils.ABDKToFloat64(vol), nil
}

func RawQueryExchangeFeeTbpsForTrader(rpc *ethclient.Client, xInfo StaticExchangeInfo, poolId int32, traderAddr common.Address, brokerAddr common.Address) (uint16, error) {
	proxy, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return 0, err
	}
	feeTbps, err := proxy.QueryExchangeFee(nil, uint8(poolId), traderAddr, brokerAddr)
	if err != nil {
		return 0, err
	}
	return feeTbps, nil
}

func RawGetPoolShTknPrice(rpc *ethclient.Client, poolIds []int, xInfo StaticExchangeInfo) ([]float64, error) {
	caller, err := multicall.New(rpc)
	if err != nil {
		return nil, err
	}
	type priceOutput struct {
		PriceD18 *big.Int
	}
	contract, err := multicall.NewContract(POOL_SHTKN_PX_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	calls := make([]*multicall.Call, 0, len(poolIds))
	for _, id := range poolIds {
		c := contract.NewCall(new(priceOutput), "getShareTokenPriceD18", uint8(id))
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	prices := make([]float64, 0, len(poolIds))
	for _, call := range res {
		px := utils.DecNToFloat(call.Outputs.(*priceOutput).PriceD18, 18)
		prices = append(prices, px)
	}
	return prices, nil
}

func RawCalculateLiquidationPrice(ccy CollateralCCY, lockedInValue float64, positionBC float64, cashCC float64, tau float64, S3 float64, Sm float64) float64 {
	if positionBC == 0 {
		return float64(0)
	}
	switch ccy {
	case QUANTO:
		numerator := -lockedInValue + cashCC*S3*(1-utils.Sign(positionBC))
		denominator := tau*math.Abs(positionBC) - positionBC - (cashCC*S3*utils.Sign(positionBC))/Sm
		return numerator / denominator
	case BASE:
		numerator := -lockedInValue + cashCC
		denominator := tau*math.Abs(positionBC) - positionBC
		return numerator / denominator
	default:
		return lockedInValue / (positionBC - tau*math.Abs(positionBC) + cashCC)
	}
}

func RawQueryLiquidatableAccounts(
	client *ethclient.Client,
	xInfo *StaticExchangeInfo,
	perpId int32,
	pythEndpoint string,
	prdMktEndpoint string,
	lowLiqEndpoint string,
) ([]common.Address, error) {
	j := GetPerpetualStaticInfoIdxFromId(xInfo, perpId)
	if j == -1 {
		return nil, fmt.Errorf("RawQueryLiquidatableAccounts: perp id %d not found", perpId)
	}
	pxFeed, err := fetchPerpetualPriceInfo(xInfo, j, pythEndpoint, prdMktEndpoint, lowLiqEndpoint)
	if err != nil {
		return nil, errors.New("RawQueryLiquidatableAccounts: failed fetching oracle prices " + err.Error())
	}
	pricesAbdk := [2]*big.Int{utils.Float64ToABDK(pxFeed.S2Price), utils.Float64ToABDK(pxFeed.S3Price)}
	proxy, err := CreatePerpetualManagerInstance(client, xInfo.ProxyAddr)
	if err != nil {
		return nil, err
	}
	id := big.NewInt(int64(perpId))
	return proxy.GetLiquidatableAccounts(nil, id, pricesAbdk)
}

func RawQueryLiquidatableAccountsInPool(client *ethclient.Client,
	xInfo *StaticExchangeInfo,
	poolId int32,
	pythEndpoint, prdMktEndpoint, lowLiqEndpoint string,
) ([]LiquidatableAccounts, error) {
	caller, err := multicall.New(client)
	if err != nil {
		return nil, err
	}
	contract, err := multicall.NewContract(GET_LIQUIDATABLE_ACC_ABI, xInfo.ProxyAddr.Hex())
	if err != nil {
		return nil, err
	}
	// get perpetuals in the pool and gather all symbols in
	// a first loop to query all prices at once
	symbols := make([]string, 0, len(xInfo.Perpetuals))
	priceIdSet := make(map[string]bool)
	symSet := make(map[string]bool)
	priceIds := make([]PriceId, 0, len(xInfo.Perpetuals))
	for k, perp := range xInfo.Perpetuals {
		if perp.PoolId != poolId || perp.State != NORMAL {
			continue
		}

		for _, pxId := range xInfo.Perpetuals[k].PriceIds {
			if priceIdSet[pxId.Id] {
				continue
			}
			priceIdSet[pxId.Id] = true
			priceIds = append(priceIds, pxId)
		}
		for _, sym := range xInfo.Perpetuals[k].OnChainSymbols {
			if symSet[sym] {
				continue
			}
			symSet[sym] = true
			symbols = append(symbols, sym)
		}
	}
	feedData, err := fetchPricesFromAPI(priceIds, pythEndpoint, prdMktEndpoint, lowLiqEndpoint, false)
	if err != nil {
		return nil, err
	}
	pxMap, err := fetchPricesFromChain(symbols, xInfo.ChainOracles)
	if err != nil {
		return nil, err
	}
	for k, id := range feedData.PriceIds {
		syms := xInfo.PriceFeedInfo.PxIdToSymbols[id]
		for _, sym := range syms {
			pxMap[sym] = Price{
				// use EMA (=S2 if regular market)
				Px:         feedData.Prices[k].Ema,
				Ts:         int64(feedData.Prices[k].Ts),
				IsOffChain: true,
			}
		}
	}
	// now that we have all prices, we build the multi-call
	calls := make([]*multicall.Call, 0, len(xInfo.Perpetuals))
	type liqOutput struct {
		LiqAccount []common.Address
	}
	// build multi-call
	perpIds := make([]int32, 0, len(xInfo.Perpetuals))
	for j := range xInfo.Perpetuals {
		if xInfo.Perpetuals[j].PoolId != poolId || xInfo.Perpetuals[j].State != NORMAL {
			continue
		}
		S2Sym := xInfo.Perpetuals[j].S2Symbol
		S3Sym := xInfo.Perpetuals[j].S3Symbol
		s2, s3, isMarketClosedS2, isMarketClosedS3, err := calculatePerpIdxPx(xInfo, pxMap, S2Sym, S3Sym)
		if err != nil {
			return nil, err
		}
		if isMarketClosedS2 || isMarketClosedS3 {
			// skip
			continue
		}
		perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
		pricesAbdk := [2]*big.Int{utils.Float64ToABDK(s2), utils.Float64ToABDK(s3)}
		c := contract.NewCall(new(liqOutput), "getLiquidatableAccounts", perpId, pricesAbdk)

		perpIds = append(perpIds, xInfo.Perpetuals[j].Id)
		calls = append(calls, c)
	}
	res, err := caller.Call(nil, calls...)
	if err != nil {
		return nil, err
	}
	liqAccs := make([]LiquidatableAccounts, 0, len(calls))
	for k, call := range res {
		addr := call.Outputs.(*liqOutput)
		if len(addr.LiqAccount) == 0 {
			continue
		}
		liqAccs = append(liqAccs,
			LiquidatableAccounts{PerpId: perpIds[k], LiqAccounts: addr.LiqAccount})
	}
	return liqAccs, nil
}

func RawFetchPricesForPerpetualId(exchangeInfo StaticExchangeInfo, id int32, pxFeedEndpoint, prdMktEndpoint, lowLiqEndpoint string) (PerpetualPriceInfo, error) {
	j := GetPerpetualStaticInfoIdxFromId(&exchangeInfo, id)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPerpetualPriceInfo(&exchangeInfo, j, pxFeedEndpoint, prdMktEndpoint, lowLiqEndpoint)
}

// FetchPricesForPerpetual queries the REST-endpoints/onchain oracles and calculates S2,S3
// index prices, also returns the price-feed-data required for blockchain submission and
// information whether the market is closed or not. endpoint is the endpoint that provides pyth prices.
func RawFetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, symbol string, pxFeedEndpoint, prdMktEndpoint, lowLiqFeedEndpoint string) (PerpetualPriceInfo, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&exchangeInfo, symbol)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPerpetualPriceInfo(&exchangeInfo, j, pxFeedEndpoint, prdMktEndpoint, lowLiqFeedEndpoint)
}

// fetchPerpetualPriceInfo gets prices from the VAA-endpoint and on-chain if needed,
// j is the index of the perpetual in StaticExchangeInfo
func fetchPerpetualPriceInfo(xInfo *StaticExchangeInfo, j int, pxFeedEndpoint, prdMktEndpoint, lowLiqFeedEndpoint string) (PerpetualPriceInfo, error) {
	pxMap, feedData, err := fetchIndexPricesForPerpetual(xInfo, j, pxFeedEndpoint, prdMktEndpoint, lowLiqFeedEndpoint)
	if err != nil {
		return PerpetualPriceInfo{}, err
	}
	s2, s3, isMarketClosedS2, isMarketClosedS3, err := calculatePerpIdxPx(xInfo, pxMap, xInfo.Perpetuals[j].S2Symbol, xInfo.Perpetuals[j].S3Symbol)
	for _, p := range feedData.Prices {
		isMarketClosedS2 = isMarketClosedS2 || p.IsClosedPrdMkt
	}

	if err != nil {
		return PerpetualPriceInfo{}, err
	}
	p := PerpetualPriceInfo{
		S2Price:          s2,
		S3Price:          s3,
		Ema:              s2,
		IsMarketClosedS2: isMarketClosedS2,
		IsMarketClosedS3: isMarketClosedS3,
		PriceFeed:        feedData,
	}
	p.PriceFeed = feedData
	// Search whether prediction market price
	for _, obs := range feedData.Prices {
		if obs.CLOBParams != 0 {
			p.Conf = obs.Conf
			p.CLOBParams = obs.CLOBParams
			p.Ema = obs.Ema
			break
		}
	}
	return p, nil
}

// fetchIndexPricesForPerpetual gets the index prices required to calculate S2 and S3 indices for the given perpetual.
// prices are gathered on-chain and offchain (depending on feeds)
// j is the index of the perpetual in StaticExchangeInfo
func fetchIndexPricesForPerpetual(xInfo *StaticExchangeInfo, j int, pxFeedEndpoint, prdMktEndpoint, lowLiqFeedEndpoint string) (map[string]Price, PriceFeedData, error) {
	// get underlying data from rest-api with vaa
	feedData, err := fetchPricesFromAPI(xInfo.Perpetuals[j].PriceIds, pxFeedEndpoint, prdMktEndpoint, lowLiqFeedEndpoint, true)
	if err != nil {
		return nil, PriceFeedData{}, err
	}
	pxMap, err := fetchPricesFromChain(xInfo.Perpetuals[j].OnChainSymbols, xInfo.ChainOracles)
	if err != nil {
		return nil, PriceFeedData{}, err
	}
	for k, id := range feedData.PriceIds {
		syms := xInfo.PriceFeedInfo.PxIdToSymbols[id]
		for _, sym := range syms {
			pxMap[sym] = Price{Px: feedData.Prices[k].Px, Ts: int64(feedData.Prices[k].Ts), IsOffChain: true}
		}
	}
	return pxMap, feedData, nil
}

// calculatePerpPxFromFeed calculates the prices for the given symbols via pre-calculated triangulations
// based on feedData
func calculatePerpIdxPx(xInfo *StaticExchangeInfo, pxMap map[string]Price, S2Sym, S3Sym string) (float64, float64, bool, bool, error) {
	// triangulate fetched prices to obtain index prices
	triang := xInfo.IdxPriceTriangulations[S2Sym]
	if len(triang.Symbol) == 0 {
		return 0, 0, false, false, fmt.Errorf("no triangulation for symbol %s", S2Sym)
	}
	thOn := int64(xInfo.PriceFeedInfo.ThreshOnChainFeedOutdatedSec)
	thOff := int64(xInfo.PriceFeedInfo.ThreshOffChainFeedOutdatedSec)
	pxS2, isMarketClosedS2, err := calculateTriangulation(triang, thOff, thOn, pxMap)
	if err != nil {
		return 0, 0, false, false, err
	}
	var pxS3 float64 = 0
	var isMarketClosedS3 bool = false
	// triangulate S3 if there is an S3 index
	if S3Sym != "" {
		triang = xInfo.IdxPriceTriangulations[S3Sym]
		if len(triang.Symbol) == 0 {
			return 0, 0, false, false, fmt.Errorf("no triangulation for symbol %s", S3Sym)
		}
		pxS3, isMarketClosedS3, err = calculateTriangulation(triang, thOff, thOn, pxMap)
		if err != nil {
			return 0, 0, false, false, err
		}
	}
	return pxS2, pxS3, isMarketClosedS2, isMarketClosedS3, nil
}

// fetchPricesFromChain retrieves the requested symbol prices from configured on-chain oracles
func fetchPricesFromChain(symbols []string, oracle *ChainOracles) (map[string]Price, error) {
	prices := make(map[string]Price)
	for _, sym := range symbols {
		px, ts, err := oracle.GetPrice(sym, false)
		if err != nil {
			return nil, err
		}
		prices[sym] = Price{
			Px:         px,
			Ts:         ts,
			IsOffChain: false,
		}
	}
	return prices, nil
}

// fetchPricesFromAPI gets the prices for the given priceIds from the
// configured REST-API. The PriceFeedConfig is needed to store what
// symbol (e.g. BTC-USD) the price feed covers.
func fetchPricesFromAPI(priceIds []PriceId, priceFeedEndpoint, prdMktEndpoint, lowLiqEndpoint string, withVaa bool) (PriceFeedData, error) {
	pxData := PriceFeedData{
		PriceIds: make([]string, len(priceIds)),
		Prices:   make([]PriceObs, len(priceIds)),
	}
	if withVaa {
		pxData.Vaas = make([][]byte, len(priceIds))
	}

	// REST query
	// include VAA
	data, err := fetchPythPrices(priceIds, priceFeedEndpoint, prdMktEndpoint, lowLiqEndpoint)
	if err != nil {
		return PriceFeedData{}, err
	}
	// process data
	for _, d := range data {
		// find idx of d.Id
		for i, id := range priceIds {
			if id.Id == d.Id {

				pxData.PriceIds[i] = d.Id
				if id.Type == utils.PXTYPE_PYTH {
					// regular markets: we set the S2 as EMA and set the parameters
					// to zero
					pxData.Prices[i] = PriceObs{
						Px:             utils.PythNToFloat64(d.Price.Price, d.Price.Expo),
						Ema:            utils.PythNToFloat64(d.Price.Price, d.Price.Expo),
						Conf:           uint16(0),
						CLOBParams:     0,
						Ts:             int64(d.Price.PublishTime),
						IsOffChain:     true,
						IsClosedPrdMkt: false,
					}
				} else {
					// lowliq & prediction markets: set ema and parameters
					conf, err := strconv.Atoi(d.Price.Conf)
					if err != nil {
						return PriceFeedData{}, fmt.Errorf("unable to convert prices.conf %s", err.Error())
					}
					params, _ := new(big.Int).SetString(d.EmaPrice.Conf, 10)
					pxData.Prices[i] = PriceObs{
						Px:             utils.PythNToFloat64(d.Price.Price, d.Price.Expo),
						Ema:            utils.PythNToFloat64(d.EmaPrice.Price, d.EmaPrice.Expo),
						Conf:           uint16(conf),
						CLOBParams:     params.Uint64(),
						Ts:             int64(d.Price.PublishTime),
						IsOffChain:     true,
						IsClosedPrdMkt: d.PrdMktClosed,
					}
				}
				if !withVaa {
					break
				}
				decodedVaaBytes, err := base64.StdEncoding.DecodeString(d.Vaa)
				if err != nil {
					err := errors.New("fetchPricesFromAPI decoding base64:" + err.Error())
					return PriceFeedData{}, err
				}
				pxData.Vaas[i] = decodedVaaBytes
				break
			}
		}
	}
	return pxData, nil
}

// fetchPythPrices gets the specified priceIds from the pyth endpoint 'priceFeedEndpoint'
// and includes VAA (signed prices)
func fetchPythPrices(priceIds []PriceId, priceFeedEndpoint, prdMktEndpoint, lowLiqEndpoint string) ([]ResponsePythLatestPriceFeed, error) {
	// see https://hermes.pyth.network/docs/
	priceFeedEndpoint = strings.TrimSuffix(priceFeedEndpoint, "/api") // cut off legacy url suffix
	priceFeedEndpoint = strings.TrimSuffix(priceFeedEndpoint, "/")
	prdMktEndpoint = strings.TrimSuffix(prdMktEndpoint, "/")
	resCh := make(chan *PythLatestPxV2, len(priceIds))
	errCh := make(chan error, len(priceIds))
	query1 := fmt.Sprintf("%s/v2/updates/price/latest?encoding=base64&ids[]=", priceFeedEndpoint)
	query2 := fmt.Sprintf("%s/v2/updates/price/latest?encoding=base64&ids[]=", prdMktEndpoint)
	query3 := fmt.Sprintf("%s/v2/updates/price/latest?encoding=base64&ids[]=", lowLiqEndpoint)
	count := 0
	for _, id := range priceIds {
		switch id.Type {
		case utils.PXTYPE_PYTH:
			fetchPythPrice(query1+strings.TrimPrefix(id.Id, "0x"), resCh, errCh)
			count++
		case utils.PXTYPE_POLYMARKET:
			fetchPythPrice(query2+strings.TrimPrefix(id.Id, "0x"), resCh, errCh)
			count++
		case utils.PXTYPE_V2, utils.PXTYPE_V3:
			fetchPythPrice(query3+strings.TrimPrefix(id.Id, "0x"), resCh, errCh)
			count++
		}
	}
	// collect the results and errors
	res := make([]ResponsePythLatestPriceFeed, count)
	for i := 0; i < count; i++ {
		select {
		case result := <-resCh:
			res[i] = ResponsePythLatestPriceFeed{
				EmaPrice:     result.Parsed[0].EmaPrice,
				Id:           strings.TrimPrefix(result.Parsed[0].ID, "0x"),
				Price:        result.Parsed[0].Price,
				Vaa:          result.Binary.Data[0],
				PrdMktClosed: result.Parsed[0].Metadata.MarketClosed,
			}
		case err := <-errCh:
			return nil, err
		}
	}
	close(resCh)
	close(errCh)
	return res, nil
}

func fetchPythPrice(query string, resCh chan<- *PythLatestPxV2, errCh chan<- error) {
	response, err := http.Get(query)
	if err != nil {
		errCh <- errors.New("error sending fetchPricesFromAPI request:" + err.Error())
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		errCh <- errors.New("error reading response body:" + err.Error())
	}
	if response.StatusCode != 200 {
		errCh <- errors.New("error fetchPricesFromAPI status " + strconv.Itoa(response.StatusCode) + " " + string(body[:]))
		return
	}
	var data PythLatestPxV2
	err = json.Unmarshal(body, &data)
	if err != nil {
		errCh <- errors.New("fetchPricesFromAPI:" + err.Error())
		return
	}
	resCh <- &data
}

func RawGetTknBalance(tknAddr common.Address, userAddr common.Address, rpc *ethclient.Client) (float64, error) {
	erc20Instance, err := contracts.NewErc20(tknAddr, rpc)
	if err != nil {
		return 0, errors.New("RawGetTknBalance: creating instance of token " + tknAddr.String())
	}
	n, err := erc20Instance.Decimals(nil)
	if err != nil {
		return 0, err
	}
	b, err := erc20Instance.BalanceOf(nil, userAddr)
	if err != nil {
		return 0, err
	}
	bal := utils.DecNToFloat(b, n)
	return bal, nil
}

func RawQueryPoolShTknBalance(lpAddr common.Address, poolId int, xInfo StaticExchangeInfo, rpc *ethclient.Client) (float64, error) {
	shTkn := xInfo.Pools[poolId-1].ShareTokenAddr
	return RawGetTknBalance(shTkn, lpAddr, rpc)
}
