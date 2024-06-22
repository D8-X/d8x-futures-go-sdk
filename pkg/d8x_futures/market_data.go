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

type OptEndPoints struct {
	Rpc            *ethclient.Client
	PriceFeedEndPt string
}

func (sdkRo *SdkRO) GetPositionRisk(symbol string, traderAddr common.Address, optEndPt *OptEndPoints) (PositionRisk, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawGetPositionRisk(sdkRo.Info, optRpc, &traderAddr, symbol, optPyth)
}

func (sdkRo *SdkRO) QueryPerpetualState(perpetualIds []int32, optEndPt *OptEndPoints) ([]PerpetualState, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryPerpetualState(optRpc, sdkRo.Info, perpetualIds, optPyth)
}

func (sdkRo *SdkRO) QueryPerpetualPrices(symbol string, tradeAmt []float64, optEndPt *OptEndPoints) ([]float64, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryPerpetualPriceTuple(optRpc, &sdkRo.Info, optPyth, symbol, tradeAmt)
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

func (sdkRo *SdkRO) QueryOrderStatus(symbol string, traderAddr common.Address, orderDigest string, optRpc *ethclient.Client) (string, error) {
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
	return RawQueryLiquidatableAccounts(optRpc, &sdkRo.Info, perpId, optPyth)
}

// QueryLiquidatableAccountsInPool identifies all traders that can be liquidated in all perpetuals of
// the given pool.
func (sdkRo *SdkRO) QueryLiquidatableAccountsInPool(poolId int32, optEndPt *OptEndPoints) ([]LiquidatableAccounts, error) {
	optRpc, optPyth := extractEndpoints(sdkRo, optEndPt)
	return RawQueryLiquidatableAccountsInPool(optRpc, &sdkRo.Info, poolId, optPyth)
}

func (sdkRo *SdkRO) FetchPricesForPerpetualId(id int32, optPythEndPt string) (PerpetualPriceInfo, error) {
	if optPythEndPt == "" {
		optPythEndPt = sdkRo.ChainConfig.PriceFeedEndpoint
	}
	return RawFetchPricesForPerpetualId(sdkRo.Info, id, optPythEndPt)
}

func (sdkRo *SdkRO) FetchPricesForPerpetual(symbol string, optPythEndPt string) (PerpetualPriceInfo, error) {
	if optPythEndPt == "" {
		optPythEndPt = sdkRo.ChainConfig.PriceFeedEndpoint
	}
	return RawFetchPricesForPerpetual(sdkRo.Info, symbol, optPythEndPt)
}

func (sdkRo *SdkRO) GetMarginTokenBalance(symbol string, traderAddr common.Address, optRpc *ethclient.Client) (float64, error) {
	tknAddr, err := RawGetMarginTknAddr(&sdkRo.Info, symbol)
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

// Allowance checks the allowance of the given address to spend margin tokens for the given
// pool (via symbol) on the manager. Returns the value in decimals and the decimal-N value (big-int).
// Symbol is a pool symbol like "USDC" (or perpetual symbol like MATIC-USDC-USDC works too)
func (sdkRo *SdkRO) Allowance(symbol string, user common.Address, optRpc *ethclient.Client) (float64, *big.Int, error) {
	tknAddr, err := RawGetMarginTknAddr(&sdkRo.Info, symbol)
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

func RawGetMarginTknAddr(xInfo *StaticExchangeInfo, symbol string) (common.Address, error) {
	j := GetPoolStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return common.Address{}, errors.New("RawGetMarginTknAddr: no perpetual " + symbol)
	}
	pool := xInfo.Pools[j]
	poolId := pool.PoolId

	return (xInfo.Pools[poolId-1].PoolMarginTokenAddr), nil
}

func RawGetPositionRisk(xInfo StaticExchangeInfo, rpc *ethclient.Client, traderAddr *common.Address, symbol string, endpoint string) (PositionRisk, error) {
	priceData, err := RawFetchPricesForPerpetual(xInfo, symbol, endpoint)
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
	S2Liq := RawCalculateLiquidationPrice(xInfo.Perpetuals[j].CollateralCurrencyType, lockedInValue, posBC, cashCC, xInfo.Perpetuals[j].MaintenanceMarginRate, S3, Sm)
	if xInfo.Perpetuals[j].CollateralCurrencyType == BASE {
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / priceData.S2Price
	} else if xInfo.Perpetuals[j].CollateralCurrencyType == QUANTO {
		// convert CC to quote
		unpaidFundingQC = unpaidFundingQC / priceData.S3Price
		S3Liq = utils.ABDKToFloat64(traderState[idxS3])
	}
	// floor at zero
	S2Liq = math.Max(0, S2Liq)
	S3Liq = math.Max(0, S3Liq)

	var side string
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
		liqLeverage = 1 / xInfo.Perpetuals[j].MaintenanceMarginRate
	} else {
		S3Liq = 0
	}
	m := PositionRisk{
		Symbol:                         symbol,
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
	return m, nil
}

// QueryPerpetualState collects PerpetualState by calling the off-chain prices and
// blockchain queries. endpoint is the address to get prices from
func RawQueryPerpetualState(rpc *ethclient.Client, xInfo StaticExchangeInfo, perpetualIds []int32, endpoint string) ([]PerpetualState, error) {
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
		p, err := RawFetchPricesForPerpetualId(xInfo, perpetualIds[i], endpoint)
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

func RawQueryOrderStatus(rpc *ethclient.Client, xInfo StaticExchangeInfo, traderAddr common.Address, orderDigest string, symbol string) (string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(&xInfo, symbol)
	if j == -1 {
		return "", fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	lob := CreateLimitOrderBookInstance(rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	// convert digest string to bytes 32
	bytesDigest := common.Hex2Bytes(strings.TrimPrefix(orderDigest, "0x"))
	var orderDigest32 [32]byte
	copy(orderDigest32[:], bytesDigest)
	status, err := lob.GetOrderStatus(nil, orderDigest32)
	if err != nil {
		return "", err
	}
	var statusStr string
	switch status {
	case ENUM_ORDER_STATUS_CANCELED:
		statusStr = ORDER_STATUS_CANCELED
	case ENUM_ORDER_STATUS_EXECUTED:
		statusStr = ORDER_STATUS_EXECUTED
	case ENUM_ORDER_STATUS_OPEN:
		statusStr = ORDER_STATUS_OPEN
	case ENUM_ORDER_STATUS_UNKNOWN:
		statusStr = ORDER_STATUS_UNKNOWN
	}
	return statusStr, nil
}

// RawQueryPerpetualPriceTuple queries prices for different trade amounts for the given perpetual. The perpetual
// is given by symbol (e.g., "ETH-USD-WEETH")
func RawQueryPerpetualPriceTuple(client *ethclient.Client, xInfo *StaticExchangeInfo, pythEndpoint, symbol string, tradeAmt []float64) ([]float64, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, fmt.Errorf("Symbol " + symbol + " does not exist in static perpetual info")
	}
	perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
	pxFeed, err := fetchPerpetualPriceInfo(xInfo, j, pythEndpoint)
	if err != nil {
		return nil, errors.New("RawAddCollateral: failed fetching oracle prices " + err.Error())
	}
	pricesAbdk := [2]*big.Int{utils.Float64ToABDK(pxFeed.S2Price), utils.Float64ToABDK(pxFeed.S3Price)}

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
		c := contract.NewCall(new(priceOutput), "queryPerpetualPrice", perpId, taAbdk, pricesAbdk)
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
	if ccy == QUANTO {
		numerator := -lockedInValue + cashCC*S3*(1-utils.Sign(positionBC))
		denominator := tau*math.Abs(positionBC) - positionBC - (cashCC*S3*utils.Sign(positionBC))/Sm
		return numerator / denominator
	} else if ccy == BASE {
		numerator := -lockedInValue + cashCC
		denominator := tau*math.Abs(positionBC) - positionBC
		return numerator / denominator
	} else {
		return lockedInValue / (positionBC - tau*math.Abs(positionBC) + cashCC)
	}
}

func RawQueryLiquidatableAccounts(client *ethclient.Client, xInfo *StaticExchangeInfo, perpId int32, pythEndpoint string) ([]common.Address, error) {

	j := GetPerpetualStaticInfoIdxFromId(xInfo, perpId)
	if j == -1 {
		return nil, fmt.Errorf("RawQueryLiquidatableAccounts: perp id %d not found", perpId)
	}
	pxFeed, err := fetchPerpetualPriceInfo(xInfo, j, pythEndpoint)
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

func RawQueryLiquidatableAccountsInPool(client *ethclient.Client, xInfo *StaticExchangeInfo, poolId int32, pythEndpoint string) ([]LiquidatableAccounts, error) {
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
	priceIds := make([]string, 0, len(xInfo.Perpetuals))
	for k, perp := range xInfo.Perpetuals {
		if perp.PoolId != poolId {
			continue
		}

		for _, pxId := range xInfo.Perpetuals[k].PriceIds {
			if priceIdSet[pxId] {
				continue
			}
			priceIdSet[pxId] = true
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
	feedData, err := fetchPricesFromAPI(priceIds, pythEndpoint, false)
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
			pxMap[sym] = PriceObs{Px: feedData.Prices[k], Ts: int64(feedData.PublishTimes[k]), IsOffChain: true}
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
		if xInfo.Perpetuals[j].PoolId != poolId {
			continue
		}
		S2Sym := xInfo.Perpetuals[j].S2Symbol
		S3Sym := xInfo.Perpetuals[j].S3Symbol
		pxInfo, err := calculatePerpIdxPx(xInfo, pxMap, S2Sym, S3Sym)
		if err != nil {
			return nil, err
		}
		if pxInfo.IsMarketClosedS2 || pxInfo.IsMarketClosedS3 {
			// skip
			continue
		}
		perpId := big.NewInt(int64(xInfo.Perpetuals[j].Id))
		pricesAbdk := [2]*big.Int{utils.Float64ToABDK(pxInfo.S2Price), utils.Float64ToABDK(pxInfo.S3Price)}
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

func RawFetchPricesForPerpetualId(exchangeInfo StaticExchangeInfo, id int32, endpoint string) (PerpetualPriceInfo, error) {
	j := GetPerpetualStaticInfoIdxFromId(&exchangeInfo, id)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPerpetualPriceInfo(&exchangeInfo, j, endpoint)
}

// FetchPricesForPerpetual queries the REST-endpoints/onchain oracles and calculates S2,S3
// index prices, also returns the price-feed-data required for blockchain submission and
// information whether the market is closed or not. endpoint is the endpoint that provides pyth prices.
func RawFetchPricesForPerpetual(exchangeInfo StaticExchangeInfo, symbol string, endpoint string) (PerpetualPriceInfo, error) {

	j := GetPerpetualStaticInfoIdxFromSymbol(&exchangeInfo, symbol)
	if j == -1 {
		return PerpetualPriceInfo{}, errors.New("symbol does not exist in static perpetual info")
	}
	return fetchPerpetualPriceInfo(&exchangeInfo, j, endpoint)
}

// fetchPerpetualPriceInfo gets prices from the VAA-endpoint and on-chain if needed,
// j is the index of the perpetual in StaticExchangeInfo
func fetchPerpetualPriceInfo(xInfo *StaticExchangeInfo, j int, endpoint string) (PerpetualPriceInfo, error) {
	pxMap, feedData, err := fetchIndexPricesForPerpetual(xInfo, j, endpoint)
	if err != nil {
		return PerpetualPriceInfo{}, err
	}
	p, err := calculatePerpIdxPx(xInfo, pxMap, xInfo.Perpetuals[j].S2Symbol, xInfo.Perpetuals[j].S3Symbol)
	if err != nil {
		return PerpetualPriceInfo{}, err
	}
	p.PriceFeed = feedData
	return p, nil
}

// fetchIndexPricesForPerpetual gets the index prices required to calculate S2 and S3 indices for the given perpetual.
// prices are gathered on-chain and offchain (depending on feeds)
// j is the index of the perpetual in StaticExchangeInfo
func fetchIndexPricesForPerpetual(xInfo *StaticExchangeInfo, j int, endpoint string) (map[string]PriceObs, PriceFeedData, error) {
	// get underlying data from rest-api with vaa
	feedData, err := fetchPricesFromAPI(xInfo.Perpetuals[j].PriceIds, endpoint, true)
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
			pxMap[sym] = PriceObs{Px: feedData.Prices[k], Ts: int64(feedData.PublishTimes[k]), IsOffChain: true}
		}
	}
	return pxMap, feedData, nil
}

// calculatePerpPxFromFeed calculates the prices for the given symbols via pre-calculated triangulations
// based on feedData
func calculatePerpIdxPx(xInfo *StaticExchangeInfo, pxMap map[string]PriceObs, S2Sym, S3Sym string) (PerpetualPriceInfo, error) {
	// triangulate fetched prices to obtain index prices
	triang := xInfo.IdxPriceTriangulations[S2Sym]
	if len(triang.Symbol) == 0 {
		return PerpetualPriceInfo{}, fmt.Errorf("no triangulation for symbol %s", S2Sym)
	}
	th1 := int64(xInfo.PriceFeedInfo.ThreshOnChainFeedOutdatedSec)
	th2 := int64(xInfo.PriceFeedInfo.ThreshOffChainFeedOutdatedSec)
	pxS2, isMarketClosedS2, err := calculateTriangulation(triang, th1, th2, pxMap)
	if err != nil {
		return PerpetualPriceInfo{}, err
	}
	var pxS3 float64 = 0
	var isMarketClosedS3 bool = false
	// triangulate S3 if there is an S3 index
	if S3Sym != "" {
		triang = xInfo.IdxPriceTriangulations[S3Sym]
		if len(triang.Symbol) == 0 {
			return PerpetualPriceInfo{}, fmt.Errorf("no triangulation for symbol %s", S3Sym)
		}
		pxS3, isMarketClosedS3, err = calculateTriangulation(triang, th1, th2, pxMap)
		if err != nil {
			return PerpetualPriceInfo{}, err
		}
	}

	var priceInfo = PerpetualPriceInfo{
		S2Price:          pxS2,
		S3Price:          pxS3,
		IsMarketClosedS2: isMarketClosedS2,
		IsMarketClosedS3: isMarketClosedS3,
	}
	return priceInfo, nil
}

// fetchPricesFromChain retrieves the requested symbol prices from configured on-chain oracles
func fetchPricesFromChain(symbols []string, oracle *ChainOracles) (map[string]PriceObs, error) {
	prices := make(map[string]PriceObs)
	for _, sym := range symbols {
		px, ts, err := oracle.GetPrice(sym, false)
		if err != nil {
			return nil, err
		}
		prices[sym] = PriceObs{Px: px, Ts: ts, IsOffChain: false}
	}
	return prices, nil
}

// fetchPricesFromAPI gets the prices for the given priceIds from the
// configured REST-API. The PriceFeedConfig is needed to store what
// symbol (e.g. BTC-USD) the price feed covers.
func fetchPricesFromAPI(priceIds []string, priceFeedEndpoint string, withVaa bool) (PriceFeedData, error) {
	pxData := PriceFeedData{
		PriceIds:     priceIds,
		Prices:       make([]float64, len(priceIds)),
		PublishTimes: make([]uint64, len(priceIds)),
	}
	if withVaa {
		pxData.Vaas = make([][]byte, len(priceIds))
	}

	// REST query (#queries == number of endpoints for feeds)
	// include VAA
	data, err := fetchPythPrices(priceIds, priceFeedEndpoint)
	if err != nil {
		return PriceFeedData{}, err
	}
	// process data
	for _, d := range data {
		//find idx of d.Id
		for i, id := range priceIds {
			if id == d.Id {
				pxData.Prices[i] = utils.PythNToFloat64(d.Price.Price, d.Price.Expo)
				pxData.PublishTimes[i] = uint64(d.Price.PublishTime)
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
func fetchPythPrices(priceIds []string, priceFeedEndpoint string) ([]ResponsePythLatestPriceFeed, error) {
	// see https://hermes.pyth.network/docs/
	priceFeedEndpoint = strings.TrimSuffix(priceFeedEndpoint, "/api") //cut off legacy url suffix
	priceFeedEndpoint = strings.TrimSuffix(priceFeedEndpoint, "/")
	resCh := make(chan *PythLatestPxV2, len(priceIds))
	errCh := make(chan error, len(priceIds))
	query := fmt.Sprintf("%s/v2/updates/price/latest?encoding=base64&ids[]=", priceFeedEndpoint)
	for _, id := range priceIds {
		fetchPythPrice(query+strings.TrimPrefix(id, "0x"), resCh, errCh)
	}
	// collect the results and errors
	res := make([]ResponsePythLatestPriceFeed, len(priceIds))
	for i := 0; i < len(priceIds); i++ {
		select {
		case result := <-resCh:
			res[i] = ResponsePythLatestPriceFeed{
				EmaPrice: result.Parsed[0].EmaPrice,
				Id:       result.Parsed[0].ID,
				Price:    result.Parsed[0].Price,
				Vaa:      result.Binary.Data[0],
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
		return 0, errors.New("GetMarginTokenBalance: creating instance of token " + tknAddr.String())
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
	var shTkn = xInfo.Pools[poolId-1].ShareTokenAddr
	return RawGetTknBalance(shTkn, lpAddr, rpc)
}
