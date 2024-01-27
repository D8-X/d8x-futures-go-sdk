package d8x_futures

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SdkRO is the read-only type
type SdkRO struct {
	Info        StaticExchangeInfo
	Conn        BlockChainConnector
	ChainConfig utils.ChainConfig
}

// Sdk is the read-write type
type Sdk struct {
	SdkRO
	Wallets []*Wallet
}

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
	PythAddr               common.Address
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

type OpenOrders struct {
	Orders      []Order
	OrderHashes []string
	SubmittedTs []uint32
	HashIndex   map[string]int // Map to store OrderHash indices
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
	PerpetualManager  *contracts.IPerpetualManager
	SymbolMapping     *map[string]string //chain-symbol (MATC) to long format (MATIC)
	PriceFeedNetwork  string             //PythEVMBeta or PythEVMStable
	PostOrderGasLimit int64              //gas limit for posting orders
	PriceFeedConfig   utils.PriceFeedConfig
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
	Vaas           [][]byte
	PublishTimes   []uint64
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
	ParentChildOrderId1 [32]byte
	ParentChildOrderId2 [32]byte
	OptTraderAddr       common.Address // populated when querying orders from on chain, otherwise optional
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

type MarginAccount struct {
	FLockedInValueQC             float64
	FCashCC                      float64
	FPositionBC                  float64
	FUnitAccumulatedFundingStart float64
	iLastOpenTimestamp           uint32
	FeeTbps                      uint16
	BrokerFeeTbps                uint16
	PositionId                   string
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

// Type provided to query payment signature from broker
type BrokerPaySignatureReq struct {
	Payment           PaySummary `json:"payment"`
	ExecutorSignature string     `json:"signature"`
}

type OrderOptions struct {
	LimitPrice          float64
	TriggerPrice        float64
	ReduceOnly          bool
	KeepPositionLvg     bool
	Deadline            uint32
	ExecutionTs         uint32
	parentChildOrderId1 *[32]byte
	parentChildOrderId2 *[32]byte
}

// NewOrder creates a new order allowing for minimal specification (function accepts nil for pointers)
func NewOrder(symbol, side, orderType string, quantity float64, leverage float64, options *OrderOptions) *Order {
	if side != SIDE_BUY && side != SIDE_SELL {
		slog.Error("side must be either " + SIDE_BUY + " or " + SIDE_SELL + " but was " + side)
		return nil
	}
	if options == nil {
		var op OrderOptions
		options = &op
	}
	if options.LimitPrice == 0 && side == SIDE_BUY {
		// set limit to a large number
		// if buy market order
		options.LimitPrice = math.MaxFloat64
	}

	ts := time.Now().Unix()
	if options.Deadline == 0 {
		// set to 30*6 days deadline
		options.Deadline = uint32(ts + 86_400*30*6)
	}
	if options.ExecutionTs == 0 {
		// set to immediate execution
		options.ExecutionTs = uint32(ts) - 5
	}
	var pcoId = &[32]byte{}
	if options.parentChildOrderId1 == nil {
		options.parentChildOrderId1 = pcoId
	}
	if options.parentChildOrderId2 == nil {
		options.parentChildOrderId2 = pcoId
	}
	order := &Order{
		Symbol:              symbol,
		Side:                side,
		Type:                orderType,
		Quantity:            quantity,
		ReduceOnly:          options.ReduceOnly,
		LimitPrice:          options.LimitPrice,
		TriggerPrice:        options.TriggerPrice,
		KeepPositionLvg:     options.KeepPositionLvg,
		BrokerFeeTbps:       0,
		BrokerAddr:          common.Address{},
		BrokerSignature:     []byte{},
		Leverage:            leverage,
		Deadline:            options.Deadline,
		ExecutionTimestamp:  options.ExecutionTs,
		ParentChildOrderId1: *options.parentChildOrderId1,
		ParentChildOrderId2: *options.parentChildOrderId1,
	}
	return order
}

func (ps *BrokerPaySignatureReq) UnmarshalJSON(data []byte) error {
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
	type BrokerPaySignatureReqAux struct {
		Payment           PaySummaryAux `json:"payment"`
		ExecutorSignature string        `json:"signature"`
	}
	var aux BrokerPaySignatureReqAux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ps.Payment.Payer = common.HexToAddress(aux.Payment.Payer)
	ps.Payment.Executor = common.HexToAddress(aux.Payment.Executor)
	ps.Payment.Token = common.HexToAddress(aux.Payment.Token)
	ps.Payment.Timestamp = aux.Payment.Timestamp
	ps.Payment.Id = aux.Payment.Id
	ps.Payment.ChainId = aux.Payment.ChainId
	ps.Payment.MultiPayCtrct = common.HexToAddress(aux.Payment.MultiPayCtrct)

	// Convert TotalAmount string to *big.Int
	ps.Payment.TotalAmount = new(big.Int)
	_, success := ps.Payment.TotalAmount.SetString(aux.Payment.TotalAmount, 10)
	if !success {
		return fmt.Errorf("failed to convert TotalAmount to big.Int")
	}
	ps.ExecutorSignature = aux.ExecutorSignature
	return nil
}

// New creates a new read-only SDK instance
// endpoints contain an rpc endpoint and a pyth Endpoint in this order
// use New("network", "", "pythendpoint") to provide a pyth endpoint
// but no RPC
func (sdkRo *SdkRO) New(networkNameOrId string, endpoints ...string) error {
	chConf, err := config.GetDefaultChainConfig(networkNameOrId)
	if err != nil {
		return err
	}
	// Use the first endpoint if provided
	if len(endpoints) > 0 && endpoints[0] != "" {
		chConf.NodeURL = endpoints[0]
	}

	// Use the second endpoint if provided
	if len(endpoints) > 1 && endpoints[1] != "" {
		chConf.PriceFeedEndpoints = []string{endpoints[1]}
	}
	if len(endpoints) > 2 {
		slog.Info("SdkRO.New: only 2 endpoints are used")
	}
	pxConf, err := config.GetDefaultPriceConfig(chConf.ChainId)
	if err != nil {
		return err
	}
	conn := CreateBlockChainConnector(pxConf, chConf)
	nest, err := QueryNestedPerpetualInfo(conn)
	if err != nil {
		return err
	}
	sdkRo.Conn = conn
	sdkRo.Info = QueryExchangeStaticInfo(&conn, &chConf, &nest)
	sdkRo.ChainConfig = chConf
	return nil
}

// New creates a new read/write Sdk instance
// networkname according to chainConfig or a chainId; rpcEndpoint and pythEndpoint can be ""
func (sdk *Sdk) New(privateKeys []string, networkName string, endpoints ...string) error {

	sdk.SdkRO.New(networkName, endpoints...)
	if sdk.Conn.Rpc == nil {
		return errors.New("sdk.Conn.Rpc=nil; required")
	}

	sdk.Wallets = make([]*Wallet, 0, len(privateKeys))
	for _, privateKey := range privateKeys {
		privateKey, _ = strings.CutPrefix(privateKey, "0x")
		w, err := NewWallet(privateKey, sdk.ChainConfig.ChainId, sdk.Conn.Rpc)
		if err != nil {
			return err
		}
		sdk.Wallets = append(sdk.Wallets, w)
	}

	return nil
}
