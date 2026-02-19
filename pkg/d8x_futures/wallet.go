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

type NonceProvider interface {
	Next() (uint64, error)
}

type Wallet struct {
	PrivateKey     *ecdsa.PrivateKey
	Address        common.Address
	Auth           *bind.TransactOpts
	ChainId        int64
	IsPostLondon   bool
	gasInitialized bool
	NonceProvider  NonceProvider
}

type GasOptions struct {
	BaseFeeMultiplier int
	TipCapMultiplier  int
	GasLimit          uint64
}

type GasOption func(*GasOptions)

func WithBaseFeeMultiplier(m int) GasOption {
	return func(o *GasOptions) { o.BaseFeeMultiplier = m }
}

func WithTipCapMultiplier(m int) GasOption {
	return func(o *GasOptions) { o.TipCapMultiplier = m }
}

func WithGasLimit(m uint64) GasOption {
	return func(o *GasOptions) { o.GasLimit = m }
}

// NewWallet constructs a new wallet. ChainId must be provided and privatekey must be of the form "abcdef012" (no 0x).
// The wallet is usable for signing immediately.
// The Chain specific gas configuration (IsPostLondon, signer type) is detected automatically on the first call to UpdateNonceAndGasPx.
func NewWallet(privateKeyHex string, chainId int64) (*Wallet, error) {
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
	w.ChainId = chainId
	var chainIdBI big.Int
	chainIdBI.SetInt64(chainId)
	w.Auth, err = bind.NewKeyedTransactorWithChainID(privateKey, &chainIdBI)
	if err != nil {
		return nil, err
	}
	w.Auth.Value = big.NewInt(0)
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
// gas fee cap if w.IsPostLondon. On the first call, it detects the chain type and configures the appropriate signer.
func (w *Wallet) UpdateNonceAndGasPx(rpc *ethclient.Client, opts ...GasOption) error {
	if !w.gasInitialized {
		if err := w.initGas(rpc); err != nil {
			return err
		}
	}
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
	return w.updateGasOptions(rpc, options)
}

// initGas detects the chain type and configures the appropriate transaction signer.
func (w *Wallet) initGas(rpc *ethclient.Client) error {
	head, err := rpc.HeaderByNumber(context.Background(), nil)
	w.IsPostLondon = err == nil && head.BaseFee != nil
	chainId := w.ChainId
	privateKey := w.PrivateKey
	if w.IsPostLondon {
		w.Auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewLondonSigner(big.NewInt(chainId)), privateKey)
		}
	} else {
		w.Auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainId)), privateKey)
		}
	}
	w.gasInitialized = true
	return nil
}

func (w *Wallet) updateGasOptions(rpc *ethclient.Client, opts GasOptions) error {
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
	if opts.GasLimit != 0 {
		w.Auth.GasLimit = opts.GasLimit
	}
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
	if w.NonceProvider != nil {
		n, err := w.NonceProvider.Next()
		if err != nil {
			return fmt.Errorf("nonce provider: %w", err)
		}
		w.Auth.Nonce = big.NewInt(int64(n))
		return nil
	}
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
