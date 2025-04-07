package d8x_futures

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	PrivateKey   *ecdsa.PrivateKey
	Address      common.Address
	Auth         *bind.TransactOpts
	IsPostLondon bool
}

type GasOptions struct {
	BaseFeeMultiplier int
	TipCapMultiplier  int
}

type GasOption func(*GasOptions)

func WithBaseFeeMultiplier(m int) GasOption {
	return func(o *GasOptions) { o.BaseFeeMultiplier = m }
}

func WithTipCapMultiplier(m int) GasOption {
	return func(o *GasOptions) { o.TipCapMultiplier = m }
}

// NewWallet constructs a new wallet. ChainId must be provided and privatekey must be of the form "abcdef012" (no 0x)
// rpc can be nil
func NewWallet(privateKeyHex string, chainId int64, rpc *ethclient.Client) (*Wallet, error) {
	var w Wallet
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	w.Address = crypto.PubkeyToAddress(*publicKeyECDSA)
	w.PrivateKey = privateKey
	var chainIdBI big.Int
	chainIdBI.SetInt64(chainId)
	w.Auth, err = bind.NewKeyedTransactorWithChainID(privateKey, &chainIdBI)
	if err != nil {
		return nil, err
	}
	// determine the type of RPC (post London EIP-1559 or pre)
	head, err := rpc.HeaderByNumber(context.Background(), nil)
	w.IsPostLondon = err == nil && head.BaseFee != nil
	if w.IsPostLondon {
		w.Auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			chainID := big.NewInt(chainId)
			return types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		}
		// we only have layer 2 chains with no tip amount, hence set here and
		// no need to touch again
		tip, err := rpc.SuggestGasTipCap(context.Background())
		if err != nil {
			return nil, err
		}
		w.Auth.GasTipCap = tip
		w.updateGasFeeCap(rpc, GasOptions{BaseFeeMultiplier: 5, TipCapMultiplier: 2})
	} else {
		w.Auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			chainID := big.NewInt(chainId)
			return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		}
		w.UpdateGasPrice(rpc)
	}
	w.Auth.Value = big.NewInt(0)
	if rpc == nil {
		return &w, nil
	}
	return &w, nil
}

func (w *Wallet) SetGasPrice(p *big.Int) error {
	if w.Auth.GasFeeCap != nil || w.Auth.GasTipCap != nil {
		return errors.New("both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) specified")
	}
	w.Auth.GasPrice = p
	return nil
}

func (w *Wallet) SetGasLimit(lim uint64) {
	w.Auth.GasLimit = lim
}

func (w *Wallet) SetValue(val int64) {
	w.Auth.Value = big.NewInt(val)
}

// GetLastGasPrix returns the currently set gas price
// or (post London) the gasfeecap+gastip
func (w *Wallet) GetLastGasPrice() *big.Int {
	if w.IsPostLondon {
		if w.Auth.GasFeeCap == nil {
			return nil
		}
		return new(big.Int).Add(w.Auth.GasFeeCap, w.Auth.GasTipCap)
	}
	return w.Auth.GasPrice
}

// UpdateNonceAndGasPx updates nonce and gas price, or nonce and
// gas fee cap if w.IsPostLondon
func (w *Wallet) UpdateNonceAndGasPx(rpc *ethclient.Client, opts ...GasOption) error {
	err := w.UpdateNonce(rpc)
	if err != nil {
		return err
	}
	if !w.IsPostLondon {
		return w.UpdateGasPrice(rpc)
	}

	// defaults for post london
	options := GasOptions{BaseFeeMultiplier: 5, TipCapMultiplier: 2}
	for _, opt := range opts {
		opt(&options)
	}
	return w.updateGasFeeCap(rpc, options)
}

func (w *Wallet) updateGasFeeCap(rpc *ethclient.Client, opts GasOptions) error {
	ctx := context.Background()

	tipCap, err := rpc.SuggestGasTipCap(ctx)
	if err != nil {
		return fmt.Errorf("updateGasFeeCap: failed to get tip cap: %v", err)
	}
	w.Auth.GasTipCap = new(big.Int).Mul(tipCap, big.NewInt(int64(opts.TipCapMultiplier)))
	head, err := rpc.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("updateGasFeeCap: rpc could not get BaseFee: %v", err)
	}
	w.Auth.GasFeeCap = new(big.Int).Add(
		w.Auth.GasTipCap,
		new(big.Int).Mul(head.BaseFee, big.NewInt(int64(opts.BaseFeeMultiplier))),
	)
	return nil
}

func (w *Wallet) UpdateGasPrice(rpc *ethclient.Client) error {
	g, err := GetGasPrice(rpc)
	if err != nil {
		return errors.New("RPC could not determine gas price:" + err.Error())
	}
	w.Auth.GasPrice = g
	return nil
}

func (w *Wallet) UpdateNonce(rpc *ethclient.Client) error {
	n, err := GetNonce(rpc, w.Address)
	if err != nil {
		return errors.New("RPC could not determine nonce:" + err.Error())
	}
	w.Auth.Nonce = big.NewInt(int64(n))
	return nil
}

func GetNonce(rpc *ethclient.Client, a common.Address) (uint64, error) {
	nonce, err := rpc.PendingNonceAt(context.Background(), a)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func GetGasPrice(rpc *ethclient.Client) (*big.Int, error) {
	gasPrice, err := rpc.SuggestGasPrice(context.Background())
	if err != nil {
		return (new(big.Int)).SetInt64(0), err
	}
	return gasPrice, nil
}
