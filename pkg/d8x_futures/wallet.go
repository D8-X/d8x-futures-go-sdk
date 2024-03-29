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
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Auth       *bind.TransactOpts
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
	signerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		chainID := big.NewInt(chainId)
		return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	}
	w.Auth.Signer = signerFn
	// set default values
	w.Auth.Value = big.NewInt(0)
	w.Auth.GasLimit = uint64(300_000)

	if rpc == nil {
		return &w, nil
	}
	// query current values
	w.Auth.GasPrice, err = GetGasPrice(rpc)
	if err != nil {
		return nil, errors.New("RPC could not determine gas price:" + err.Error())
	}
	n, err := GetNonce(rpc, w.Address)
	if err != nil {
		return nil, errors.New("RPC could not determine nonce:" + err.Error())
	}
	w.Auth.Nonce = big.NewInt(int64(n))

	return &w, nil
}

func (w *Wallet) SetGasPrice(p *big.Int) {
	w.Auth.GasPrice = p
}

func (w *Wallet) SetGasLimit(lim uint64) {
	w.Auth.GasLimit = lim
}

func (w *Wallet) SetValue(val int64) {
	w.Auth.Value = big.NewInt(val)
}

func (w *Wallet) UpdateNonceAndGasPx(rpc *ethclient.Client) error {
	err := w.UpdateNonce(rpc)
	if err != nil {
		return err
	}
	return w.UpdateGasPrice(rpc)
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
