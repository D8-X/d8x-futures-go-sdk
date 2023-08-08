package d8x_futures

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NestedPerpetualIds struct {
	PerpetualIds        [][]*big.Int
	PoolShareTokenAddr  []common.Address
	PoolMarginTokenAddr []common.Address
	OracleFactoryAddr   common.Address
}

// StaticExchangeInfo is the main information with which
// many functions can operate on the exchange
// It can be stored into a JSON file (info.Store) and loaded
// from a JSON file (info.Load)
type StaticExchangeInfo struct {
	Pools                  []PoolStaticInfo
	Perpetuals             []PerpetualStaticInfo
	PerpetualSymbolToId    map[string]int32
	PerpetualIdToSymbol    map[int32]string
	PriceFeedInfo          utils.PriceFeedConfig
	IdxPriceTriangulations Triangulations
	OracleFactoryAddr      common.Address
	ProxyAddr              common.Address
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

type PoolState struct {
	Id                       int32
	IsRunning                bool
	DefaultFundCashCC        float64
	PnlParticipantCashCC     float64
	TotalSupplyShareToken    float64
	TotalTargetAMMFundSizeCC float64
}

type PerpetualState struct {
	Id                    int32
	State                 PerpetualStateEnum
	IndexPrice            float64
	CollToQuoteIndexPrice float64
	MarkPrice             float64
	MidPrice              float64
	CurrentFundingRateBps float64
	OpenInterestBC        float64
	IsMarketClosed        bool
}

type CollateralCCY int8

const (
	QUOTE CollateralCCY = iota
	BASE
	QUANTO
)

type PerpetualStateEnum int8

const (
	INVALID PerpetualStateEnum = iota
	INITIALIZING
	NORMAL
	EMERGENCY
	SETTLE
	CLEARED
)

var (
	SIDE_CLOSED = "CLOSED"
	SIDE_BUY    = "BUY"
	SIDE_SELL   = "SELL"
)
var (
	ORDER_TYPE_LIMIT       = "LIMIT"
	ORDER_TYPE_MARKET      = "MARKET"
	ORDER_TYPE_STOP_MARKET = "STOP_MARKET"
	ORDER_TYPE_STOP_LIMIT  = "STOP_LIMIT"
)

const (
	ENUM_ORDER_STATUS_CANCELED = iota
	ENUM_ORDER_STATUS_EXECUTED
	ENUM_ORDER_STATUS_OPEN
	ENUM_ORDER_STATUS_UNKNOWN
)

var (
	ORDER_STATUS_CANCELED = "CANCELED"
	ORDER_STATUS_EXECUTED = "EXECUTED"
	ORDER_STATUS_OPEN     = "OPEN"
	ORDER_STATUS_UNKNOWN  = "UNKNOWN"
)

var (
	MASK_CLOSE_ONLY        uint32 = 0x80000000
	MASK_LIMIT_ORDER       uint32 = 0x04000000
	MASK_MARKET_ORDER      uint32 = 0x40000000
	MASK_STOP_ORDER        uint32 = 0x20000000
	MASK_KEEP_POS_LEVERAGE uint32 = 0x08000000
)

type BlockChainConnector struct {
	Rpc               *ethclient.Client
	ChainId           int64
	PerpetualManager  *IPerpetualManager
	SymbolMapping     *map[string]string //chain-symbol (MATC) to long format (MATIC)
	PriceFeedNetwork  string             //testnet or mainnet
	PostOrderGasLimit int64              //gas limit for posting orders
}

type ExchangeInfo struct {
	pools []PoolState
}

type PerpetualPriceInfo struct {
	S2Price          float64
	S3Price          float64
	IsMarketClosedS2 bool
	IsMarketClosedS3 bool
	PriceFeed        PriceFeedData
}

type Triangulation struct {
	IsInverse []bool   //[false, true]
	Symbol    []string // [BTC-USD, USDC-USD]
}

type Triangulations map[string]Triangulation

type PriceFeedData struct {
	Symbols        []string
	PriceIds       []string
	Prices         []float64
	IsMarketClosed []bool
	Vaas           []string
}

type ResponsePythLatestPriceFeed struct {
	EmaPrice ResponsePythPrice `json:"ema_price"`
	Id       string            `json:"id"`
	Price    ResponsePythPrice `json:"price"`
	Vaa      string            `json:"vaa"`
}

type ResponsePythPrice struct {
	Conf        string `json:"conf"`
	Expo        int32  `json:"expo"`
	Price       string `json:"price"`
	PublishTime int32  `json:"publish_time"`
}

type Keccak256Hash [32]byte

type Order struct {
	Symbol              string //symbol of the form ETH-USD-MATIC
	Side                string
	Type                string
	Quantity            float64
	ReduceOnly          bool
	LimitPrice          float64
	TriggerPrice        float64
	KeepPositionLvg     bool
	BrokerFeeTbps       uint16
	BrokerAddr          common.Address
	BrokerSignature     []byte
	Leverage            float64
	Deadline            uint32
	ExecutionTimestamp  uint32
	parentChildOrderId1 [32]byte
	parentChildOrderId2 [32]byte
}

type PositionRisk struct {
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

// struct to be submitted by front-end provider
// to get broker signature
type PaySummary struct {
	Payer         common.Address `json:"payer"`
	Executor      common.Address `json:"executor"`
	Token         common.Address `json:"token"`
	Timestamp     uint32         `json:"timestamp"`
	Id            uint32         `json:"id"`
	TotalAmount   *big.Int       `json:"totalAmount"`
	ChainId       int64          `json:"chainId"`
	MultiPayCtrct common.Address `json:"multiPayCtrct"`
}

func (ps *PaySummary) UnmarshalJSON(data []byte) error {
	// Define an auxiliary struct with TotalAmount as string to handle conversion
	type PaySummaryAux struct {
		Payer         string `json:"payer"`
		Executor      string `json:"executor"`
		Token         string `json:"token"`
		Timestamp     uint32 `json:"timestamp"`
		Id            uint32 `json:"id"`
		TotalAmount   string `json:"totalAmount"`
		ChainId       int64  `json:"chainId"`
		MultiPayCtrct string `json:"multiPayCtrct"`
	}

	var aux PaySummaryAux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ps.Payer = common.HexToAddress(aux.Payer)
	ps.Executor = common.HexToAddress(aux.Executor)
	ps.Token = common.HexToAddress(aux.Token)
	ps.Timestamp = aux.Timestamp
	ps.Id = aux.Id
	ps.ChainId = aux.ChainId
	ps.MultiPayCtrct = common.HexToAddress(aux.MultiPayCtrct)

	// Convert TotalAmount string to *big.Int
	ps.TotalAmount = new(big.Int)
	_, success := ps.TotalAmount.SetString(aux.TotalAmount, 10)
	if !success {
		return fmt.Errorf("failed to convert TotalAmount to big.Int")
	}

	return nil
}
