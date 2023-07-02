package d8x_futures

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
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
	Rpc        *ethclient.Client
}

// NewWallet constructs a new wallet. RPC can be nil in which case the nonce will not be
// queried. ChainId must be provided and privatekey must be of the form "abcdef012" (no 0x)
func (w *Wallet) NewWallet(privateKeyHex string, chainId int64, rpc *ethclient.Client) error {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	w.Address = crypto.PubkeyToAddress(*publicKeyECDSA)
	w.PrivateKey = privateKey

	w.Auth = bind.NewKeyedTransactor(privateKey)
	signerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		chainID := big.NewInt(chainId)
		return types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	}
	w.Auth.Signer = signerFn
	// set default values
	w.Auth.Value = big.NewInt(0)
	w.Auth.GasLimit = uint64(300000)

	w.Rpc = rpc
	if rpc != nil {
		// query current values
		w.Auth.GasPrice, err = GetGasPrice(rpc)
		if err != nil {
			return fmt.Errorf("RPC could not determine gas price")
		}
		n, err := GetNonce(rpc, w.Address)
		if err != nil {
			return fmt.Errorf("RPC could not determine nonce")
		}
		w.Auth.Nonce = big.NewInt(int64(n))
	}
	return nil
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

func (w *Wallet) UpdateNonce(rpc *ethclient.Client) {
	n, err := GetNonce(w.Rpc, w.Address)
	if err != nil {
		log.Fatal("RPC could not determine gas price")
	}
	w.Auth.Nonce = big.NewInt(int64(n))
}

func GetNonce(rpc *ethclient.Client, a common.Address) (uint64, error) {
	nonce, err := rpc.PendingNonceAt(context.Background(), a)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return nonce, nil
}

func GetGasPrice(rpc *ethclient.Client) (*big.Int, error) {
	gasPrice, err := rpc.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return (new(big.Int)).SetInt64(0), err
	}
	return gasPrice, nil
}
