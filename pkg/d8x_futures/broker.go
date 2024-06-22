package d8x_futures

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (sdk *Sdk) PurchaseBrokerLots(numLots int, symbol string, opts *OptsOverrides) (*types.Transaction, error) {
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
	return RawPurchaseBrokerLots(rpc, &sdk.Info, w, symbol, numLots)
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

func RawPurchaseBrokerLots(rpc *ethclient.Client, xInfo *StaticExchangeInfo, postingWallet *Wallet, symbol string, numLots int) (*types.Transaction, error) {
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
	postingWallet.UpdateNonceAndGasPx(rpc)
	tx, err := perpCtrct.DepositBrokerLots(postingWallet.Auth, uint8(poolId), uint32(numLots))
	if err != nil {
		return nil, errors.New("RawPurchaseBrokerLots:" + err.Error())
	}
	return tx, nil
}
