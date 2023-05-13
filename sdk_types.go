package d8x_futures

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NestedPerpetualIds struct {
	PerpetualIds        [][]*big.Int
	PoolShareTokenAddr  []common.Address
	PoolMarginTokenAddr []common.Address
	OracleFactoryAddr   common.Address
}

type StaticExchangeInfo struct {
	Pools                  []PoolStaticInfo
	Perpetuals             []PerpetualStaticInfo
	PerpetualSymbolToId    map[string]int32
	PriceFeedInfo          PriceFeedConfig
	IdxPriceTriangulations Triangulations
	OracleFactoryAddr      common.Address
}

type PoolStaticInfo struct {
	PoolId              int32
	PoolMarginSymbol    string
	PoolMarginTokenAddr common.Address
	ShareTokenAddr      common.Address
}

type PerpetualStaticInfo struct {
	Id                     int32
	PoolId                 int32
	LimitOrderBookAddr     common.Address
	InitialMarginRate      float64
	MaintenanceMarginRate  float64
	CollateralCurrencyType CollateralCCY
	S2Symbol               string
	S3Symbol               string
	LotSizeBC              float64
	ReferralRebate         float64
	PriceIds               []string
}

type CollateralCCY int8

const (
	QUOTE CollateralCCY = iota
	BASE
	QUANTO
)

type BlockChainConnector struct {
	Rpc              *ethclient.Client
	PerpetualManager *IPerpetualManager
	SymbolMapping    *map[string]string
	PriceFeedNetwork string
}

type ExchangeInfo struct {
	pools             []PoolState
	oracleFactoryAddr string
	proxyAddr         string
}

type TriangulationElement struct {
	IsInverse bool
	Symbol    string // BTC-USD
}

type Triangulations map[string][]TriangulationElement

type PoolState struct {
	IsRunning                bool
	PoolSymbol               string
	MarginTokenAddr          string
	PoolShareTokenAddr       string
	DefaultFundCashCC        float64
	PnlParticipantCashCC     float64
	TotalAMMFundCashCC       float64
	TotalTargetAMMFundSizeCC float64
	BrokerCollateralLotSize  int32
	Perpetuals               []PerpetualState
}

type PerpetualState struct {
	Id                    int32
	State                 string
	BaseCurrency          string
	QuoteCurrency         string
	IndexPrice            float64
	CollToQuoteIndexPrice float64
	MarkPrice             float64
	MidPrice              float64
	CurrentFundingRateBps float64
	OpenInterestBC        float64
	IsMarketClosed        bool
}

type MarginAccount struct {
	Symbol                         string
	PositionNotionalBaseCCY        float64
	Side                           string
	EntryPrice                     float64
	Leverage                       float64
	MarkPrice                      float64
	UnrealizedPnlQuoteCCY          float64
	UnrealizedFundingCollateralCCY float64
	CollateralCC                   float64
	LiquidationPrice               [2]float64
	LiquidationLvg                 float64
	CollToQuoteConversion          float64
}
