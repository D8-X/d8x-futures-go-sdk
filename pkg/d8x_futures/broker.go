package d8x_futures

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (sdk *Sdk) PurchaseBrokerLots(numLots int, symbol string, opts *OptsOverrides, gasOpts ...GasOption) (*types.Transaction, error) {
	var w *Wallet
	var rpc *ethclient.Client
	if opts == nil {
		w = sdk.Wallets[0]
		rpc = sdk.Conn.Rpc
	} else {
		w = sdk.Wallets[opts.WalletIdx]
	}
	if rpc == nil {
		rpc = sdk.Conn.Rpc
	}
	return RawPurchaseBrokerLots(rpc, &sdk.Info, w, symbol, numLots, gasOpts...)
}

// GetBrokerFeeTbps gets the fee for the trader and chain at hand from the url
func GetBrokerFeeTbps(trader common.Address, chainId int, url string) (uint16, error) {
	req := strings.TrimSuffix(url, "/") + "/broker-fee?addr=" + trader.Hex() + "&chain=" + strconv.Itoa(chainId)
	resp, err := http.Get(req)
	if err != nil {
		return 0, err
	}

	type brokerFeeResp struct {
		BrokerFee int `json:"BrokerFeeTbps"`
	}
	defer resp.Body.Close()

	var r brokerFeeResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}

	return uint16(r.BrokerFee), nil
}

// GetBrokerAddress gets the broker address for the broker at
// the given url
func GetBrokerAddress(url string) (common.Address, error) {
	resp, err := http.Get(strings.TrimSuffix(url, "/") + "/broker-address")
	if err != nil {
		return common.Address{}, err
	}
	type brokerAddrResp struct {
		BrokerAddr string `json:"brokerAddr"`
	}
	defer resp.Body.Close()

	var r brokerAddrResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return common.Address{}, err
	}

	return common.HexToAddress(r.BrokerAddr), nil
}

// SignOrder signs an order via remote broker api (url).
// The order needs a trader address, a broker address, and fee.
// Fill broker address and fee via GetBrokerAddress and
// GetBrokerFeeTbps
func (sdk *SdkRO) SignOrder(order *Order, url string) error {
	endpoint := strings.TrimSuffix(url, "/") + "/sign-order"
	if order.OptTraderAddr == (common.Address{}) {
		return fmt.Errorf("order must contain trader address")
	}
	sco, err := order.ToChainType(&sdk.Info, order.OptTraderAddr)
	if err != nil {
		return err
	}
	type reqData struct {
		Flags              uint32         `json:"flags"`
		IPerpetualId       int            `json:"iPerpetualId"`
		TraderAddr         common.Address `json:"traderAddr"`
		BrokerAddr         common.Address `json:"brokerAddr"`
		FAmount            *big.Int       `json:"fAmount"`
		FLimitPrice        *big.Int       `json:"fLimitPrice"`
		FTriggerPrice      *big.Int       `json:"fTriggerPrice"`
		LeverageTDR        uint16         `json:"leverageTDR"`
		IDeadline          uint32         `json:"iDeadline"`
		ExecutionTimestamp uint32         `json:"executionTimestamp"`
	}
	id, err := strconv.Atoi(sco.IPerpetualId.String())
	if err != nil {
		return err
	}
	payload := reqData{
		Flags:              sco.Flags,
		IPerpetualId:       id,
		TraderAddr:         sco.TraderAddr,
		BrokerAddr:         sco.BrokerAddr,
		FAmount:            sco.FAmount,
		FLimitPrice:        sco.FLimitPrice,
		FTriggerPrice:      sco.FTriggerPrice,
		LeverageTDR:        sco.LeverageTDR,
		IDeadline:          sco.IDeadline,
		ExecutionTimestamp: sco.ExecutionTimestamp,
	}
	body, err := json.Marshal(&payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var res struct {
		BrokerSignature string `json:"brokerSignature"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}
	order.BrokerSignature = []byte(res.BrokerSignature)
	return nil
}

// QueryBrokerLots queries the number of lots that the given address purchased in the pool with symbol
// 'symbol'. Returns -1 on error
func (sdk *SdkRO) QueryBrokerLots(symbol string, addr *common.Address, optRpc *ethclient.Client) (int, error) {
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	return RawQueryBrokerLots(optRpc, &sdk.Info, symbol, addr)
}

func RawQueryBrokerLots(rpc *ethclient.Client, xInfo *StaticExchangeInfo, symbol string, addr *common.Address) (int, error) {
	j := GetPoolStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return -1, errors.New("Could not find pool for symbol " + symbol)
	}
	pool := xInfo.Pools[j]
	poolId := pool.PoolId
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return -1, errors.New("RawQueryBrokerLots:" + err.Error())
	}
	lots, err := perpCtrct.GetBrokerDesignation(nil, uint8(poolId), *addr)
	if err != nil {
		return -1, errors.New("RawQueryBrokerLots:" + err.Error())
	}
	return int(lots), nil
}

func RawPurchaseBrokerLots(rpc *ethclient.Client, xInfo *StaticExchangeInfo, postingWallet *Wallet, symbol string, numLots int, gasOpts ...GasOption) (*types.Transaction, error) {
	j := GetPoolStaticInfoIdxFromSymbol(xInfo, symbol)
	if j == -1 {
		return nil, errors.New("Could not find pool for symbol " + symbol)
	}
	pool := xInfo.Pools[j]
	poolId := pool.PoolId
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, errors.New("RawPurchaseBrokerLots:" + err.Error())
	}
	g := postingWallet.Auth.GasLimit
	defer postingWallet.SetGasLimit(g)
	postingWallet.SetGasLimit(uint64(1_000_000))
	postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	tx, err := perpCtrct.DepositBrokerLots(postingWallet.Auth, uint8(poolId), uint32(numLots))
	if err != nil {
		return nil, errors.New("RawPurchaseBrokerLots:" + err.Error())
	}
	return tx, nil
}
