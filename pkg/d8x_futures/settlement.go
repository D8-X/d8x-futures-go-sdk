package d8x_futures

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SetEmergencyState sets the perpetual to emergency state
func (sdk *Sdk) SetEmergencyState(symbol string, optRpc *ethclient.Client, optGas ...GasOption) (*types.Transaction, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return nil, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	return RawSetSetEmergencyState(&sdk.Conn, optRpc, int32(id), &sdk.Info, sdk.Wallets[0], optGas...)
}

// RunSettlementProcess runs the entire settlement process in an idempotent way.
func (sdk *Sdk) RunSettlementProcess(symbol string, pxS2, pxS3 float64, optRpc *ethclient.Client, optGas ...GasOption) error {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	state, err := sdk.QueryPerpetualStateEnum(symbol, optRpc)
	if err != nil {
		return err
	}
	for {
		switch state {
		case NORMAL:
			if err := tryOnchainCall(optRpc,
				func() (*types.Transaction, error) { return sdk.SetEmergencyState(symbol, optRpc, optGas...) },
			); err != nil {
				return err
			}
			fmt.Printf("%s: emergency state set\n", symbol)
		case EMERGENCY:
			var tx *types.Transaction
			if err := tryOnchainCall(optRpc,
				func() (*types.Transaction, error) {
					var err error
					tx, err = sdk.SetSettlementPrice(symbol, pxS2, pxS3, optRpc, optGas...)
					return tx, err
				},
			); err != nil {
				return err
			}
			if _, err := bind.WaitMined(context.Background(), optRpc, tx); err != nil {
				return fmt.Errorf("%s: failed to mine settlement price transaction: %w", symbol, err)
			}

			fmt.Printf("%s: settlement price set\n", symbol)
			state, err := sdk.QueryPerpetualStateEnum(symbol, optRpc)
			if err != nil {
				return err
			}
			if state == EMERGENCY {
				if err := tryOnchainCall(optRpc,
					func() (*types.Transaction, error) {
						var err error
						tx, err = sdk.ToggleEmergencyState(symbol, optRpc, optGas...)
						return tx, err
					},
				); err != nil {
					return err
				}
				if _, err := bind.WaitMined(context.Background(), optRpc, tx); err != nil {
					return fmt.Errorf("%s: failed to mine toggle emergency state transaction: %w", symbol, err)
				}
			}
			fmt.Printf("%s: enter settle state\n", symbol)
		case SETTLE:
			praw, err := sdk.QueryPerpetualRawData(symbol, optRpc)
			if err != nil {
				return fmt.Errorf("%s unable to query state: %w", symbol, err)
			}
			pxS2approx := utils.ABDKToFloat64(praw.FSettlementS2PriceData)
			if math.Abs(pxS2approx-pxS2) > 0.01 {
				fmt.Printf("%s: settlement price is out of range: S2 onchain=%f, S2=%f\n", symbol, pxS2approx, pxS2)
				time.Sleep(1 * time.Second)
				state = EMERGENCY
				continue
			}
			if err := sdk.ClearTradersInPoolOfPerp(symbol, optRpc, optGas...); err != nil {
				return err
			}
			fmt.Printf("%s: enter cleared state\n", symbol)
			time.Sleep(3 * time.Second)
		case CLEARED:
			state, err = sdk.QueryPerpetualStateEnum(symbol, optRpc)
			if err != nil {
				return err
			}
			if state != CLEARED {
				fmt.Printf("%s: state not CLEARED by current rpc but %s\n", symbol, state.String())
				continue
			}
			if err := sdk.SettleTraders(symbol, optRpc, optGas...); err != nil {
				return err
			}
			fmt.Printf("%s: finished cleared state\n", symbol)
		}
		time.Sleep(3 * time.Second) // rpc cooloff
		state, err = sdk.QueryPerpetualStateEnum(symbol, optRpc)
		if err != nil {
			return err
		}
		if state == INVALID {
			fmt.Printf("%s settlement completed\n", symbol)
			break
		}
	}
	return nil
}

func tryOnchainCall(rpc *ethclient.Client, fu func() (*types.Transaction, error)) error {
	tx, err := fu()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err = bind.WaitMined(ctx, rpc, tx)
	if err != nil {
		return err
	}
	return nil
}

func (sdk *Sdk) SetSettlementPrice(symbol string, pxS2, pxS3 float64, optRpc *ethclient.Client, optGas ...GasOption) (*types.Transaction, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return nil, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	s2 := utils.Float64ToABDK(pxS2)
	s3 := utils.Float64ToABDK(pxS3)
	return RawSetSettlementPrice(&sdk.Conn, optRpc, int32(id), s2, s3, &sdk.Info, sdk.Wallets[0], optGas...)
}

// ToggleEmergencyState prepares the perpetual for settlement. Toggles back to emergency when called again
func (sdk *Sdk) ToggleEmergencyState(symbol string, optRpc *ethclient.Client, optGas ...GasOption) (*types.Transaction, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return nil, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	return RawToggleEmergencyState(&sdk.Conn, optRpc, int32(id), &sdk.Info, sdk.Wallets[0], optGas...)
}

func (sdk *Sdk) ClearTradersInPoolOfPerp(symbol string, optRpc *ethclient.Client, optGas ...GasOption) error {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	const bulkSize = 15
	id := uint8(sdk.Info.Perpetuals[j].PoolId)
	state, err := sdk.QueryPerpetualStateEnum(symbol, optRpc)
	if err != nil {
		return err
	}
	if state != SETTLE {
		return fmt.Errorf("perpetual state must be 'SETTLE'")
	}
	for j := 0; state == SETTLE; j++ {
		tx, err := RawClearTradersInPool(&sdk.Conn, optRpc, id, bulkSize, &sdk.Info, sdk.Wallets[0], optGas...)
		if err != nil {
			return fmt.Errorf("clear traders in pool failed in iteration %d: %w", j, err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		_, err = bind.WaitMined(ctx, optRpc, tx)
		if err != nil {
			return fmt.Errorf("clear failed in iteration %d: %w", j, err)
		}
		time.Sleep(1 * time.Second) // rpc cool-off
		state, err = sdk.QueryPerpetualStateEnum(symbol, optRpc)
		if err != nil {
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

func (sdk *Sdk) SettleTraders(symbol string, optRpc *ethclient.Client, optGas ...GasOption) error {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	const bulkSize = 15
	perpId := int64(sdk.Info.Perpetuals[j].Id)
	state, err := sdk.QueryPerpetualStateEnum(symbol, optRpc)
	if err != nil {
		return err
	}
	if state != CLEARED {
		return fmt.Errorf("perpetual state must be 'CLEARED'")
	}
	proxy, err := CreatePerpetualManagerInstance(optRpc, sdk.Info.ProxyAddr)
	if err != nil {
		return err
	}
	for {
		tx, err := RawSettleTradersInPerp(&sdk.Conn, optRpc, perpId, bulkSize, &sdk.Info, sdk.Wallets[0], optGas...)
		if err != nil {
			return fmt.Errorf("clear traders in pool failed in iteration %d: %w", j, err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		_, err = bind.WaitMined(ctx, optRpc, tx)
		if err != nil {
			return fmt.Errorf("clear failed in iteration %d: %w", j, err)
		}
		time.Sleep(1 * time.Second) // rpc cool-off
		acc, err := proxy.GetSettleableAccounts(nil, big.NewInt(perpId), big.NewInt(0), big.NewInt(10))
		if err != nil {
			return err
		}
		if len(acc) == 0 {
			break
		}
	}
	time.Sleep(2 * time.Second) // rpc cool-off
	tx, err := sdk.DeactivatePerp(symbol, optRpc, optGas...)
	if err != nil {
		return fmt.Errorf("failed to deactivate perp: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err = bind.WaitMined(ctx, optRpc, tx)
	if err != nil {
		return fmt.Errorf("failed to mine (deactivate perp): %w", err)
	}
	return nil
}

func (sdk *Sdk) DeactivatePerp(symbol string, optRpc *ethclient.Client, optGas ...GasOption) (*types.Transaction, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return nil, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("no perpetual index found for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	state, err := sdk.QueryPerpetualStateEnum(symbol, optRpc)
	if err != nil {
		return nil, err
	}
	if state != CLEARED {
		return nil, fmt.Errorf("perpetual state must be 'CLEARED'")
	}
	return RawDeactivatePerp(&sdk.Conn, optRpc, id, &sdk.Info, sdk.Wallets[0])
}

// QueryPerpetualStateEnum queries the onchain state of the perpetual
// Internal state no longer correct after calling this function
func (sdk *SdkRO) QueryPerpetualRawData(symbol string, optRpc *ethclient.Client) (*contracts.PerpStoragePerpetualData, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return nil, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("no perpetual id for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	proxy, err := CreatePerpetualManagerInstance(optRpc, sdk.Info.ProxyAddr)
	if err != nil {
		return nil, err
	}
	perp, err := proxy.GetPerpetual(nil, big.NewInt(id))
	if err != nil {
		return nil, err
	}
	return &perp, nil
}

// QueryPerpetualStateEnum queries the onchain state of the perpetual
// Internal state no longer correct after calling this function
func (sdk *SdkRO) QueryPerpetualStateEnum(symbol string, optRpc *ethclient.Client) (PerpetualStateEnum, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return 10, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return 10, fmt.Errorf("no perpetual id for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	proxy, err := CreatePerpetualManagerInstance(optRpc, sdk.Info.ProxyAddr)
	if err != nil {
		return 10, err
	}
	perp, err := proxy.GetPerpetual(nil, big.NewInt(id))
	if err != nil {
		return 10, err
	}
	return PerpetualStateEnum(perp.State), nil
}

// RefreshPerpetualStateEnum queries the onchain state of the perpetual
// and updates the internal state in a thread-safe way, returns the new/current state
func (sdk *SdkRO) RefreshPerpetualStateEnum(symbol string, optRpc *ethclient.Client) (PerpetualStateEnum, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return 10, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return 10, fmt.Errorf("no perpetual id for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	proxy, err := CreatePerpetualManagerInstance(optRpc, sdk.Info.ProxyAddr)
	if err != nil {
		return 10, err
	}
	perp, err := proxy.GetPerpetual(nil, big.NewInt(id))
	if err != nil {
		return 10, err
	}
	sdk.Info.Perpetuals[j].setState(PerpetualStateEnum(perp.State))
	return PerpetualStateEnum(perp.State), nil
}

func (sdk *Sdk) ActivatePerpetual(symbol string, optRpc *ethclient.Client, gasOpts ...GasOption) (*types.Transaction, error) {
	var err error
	symbol, err = sdk.symbolToInternal(symbol)
	if err != nil {
		return nil, err
	}
	if optRpc == nil {
		optRpc = sdk.Conn.Rpc
	}
	state, err := sdk.QueryPerpetualStateEnum(symbol, optRpc)
	if err != nil {
		return nil, err
	}
	if state != INVALID {
		return nil, fmt.Errorf("perp state must be invalid, is %s", state.String())
	}
	j := GetPerpetualStaticInfoIdxFromSymbol(&sdk.Info, symbol)
	if j == -1 {
		return nil, fmt.Errorf("no perpetual id for symbol %s", symbol)
	}
	id := int64(sdk.Info.Perpetuals[j].Id)
	tx, err := RawActivatePerpetual(&sdk.Conn, optRpc, int32(id), &sdk.Info, sdk.Wallets[0], gasOpts...)
	if err != nil {
		return nil, err
	}
	_, err = bind.WaitMined(context.Background(), optRpc, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func RawDeactivatePerp(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	perpId int64,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawLiquidatePosition: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.DeactivatePerp(postingWallet.Auth, big.NewInt(perpId))
}

func RawSettleTradersInPerp(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	perpId int64,
	bulkSize int,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawSettleTradersInPool: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.SettleTraders(postingWallet.Auth, big.NewInt(perpId), big.NewInt(int64(bulkSize)))
}

func RawClearTradersInPool(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	poolId uint8,
	bulkSize int,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawClearTradersInPool: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.ClearTradersInPool(postingWallet.Auth, poolId, big.NewInt(int64(bulkSize)))
}

func RawToggleEmergencyState(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	perpId int32,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawLiquidatePosition: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.TogglePerpEmergencyState(postingWallet.Auth, big.NewInt(int64(perpId)))
}

func RawSetSettlementPrice(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	perpId int32,
	S2, S3 *big.Int,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawLiquidatePosition: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.AdjustSettlementPrice(postingWallet.Auth, big.NewInt(int64(perpId)), S2, S3)
}

func RawActivatePerpetual(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	perpId int32,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawLiquidatePosition: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.ActivatePerpetual(postingWallet.Auth, big.NewInt(int64(perpId)))
}

func RawSetSetEmergencyState(
	conn *BlockChainConnector,
	rpc *ethclient.Client,
	perpId int32,
	xInfo *StaticExchangeInfo,
	postingWallet *Wallet,
	gasOpts ...GasOption,
) (*types.Transaction, error) {
	perpCtrct, err := CreatePerpetualManagerInstance(rpc, xInfo.ProxyAddr)
	if err != nil {
		return nil, fmt.Errorf("failed creating manager instance %w", err)
	}
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	postingWallet.Auth.Value = big.NewInt(0)

	err = postingWallet.UpdateNonceAndGasPx(rpc, gasOpts...)
	if err != nil {
		return nil, fmt.Errorf("RawLiquidatePosition: uptdateNonceAndGasPx %v", err)
	}
	return perpCtrct.SetEmergencyState(postingWallet.Auth, big.NewInt(int64(perpId)))
}
