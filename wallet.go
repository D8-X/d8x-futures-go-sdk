package d8x_futures

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Auth       *bind.TransactOpts
}

func (w *Wallet) NewWallet(privateKeyHex string, conn BlockChainConnector) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	w.Address = crypto.PubkeyToAddress(*publicKeyECDSA)
	w.Auth = bind.NewKeyedTransactor(privateKey)
	// set default values
	w.Auth.Value = big.NewInt(0)
	w.Auth.GasLimit = uint64(300000)
	// query current values
	w.Auth.GasPrice, err = GetGasPrice(conn)
	if err != nil {
		log.Fatal("RPC could not determine gas price")
	}
	n, err := GetNonce(conn, w.Address)
	if err != nil {
		log.Fatal("RPC could not determine gas price")
	}
	w.Auth.Nonce = big.NewInt(int64(n))
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

func (w *Wallet) UpdateNonce(conn BlockChainConnector) {
	n, err := GetNonce(conn, w.Address)
	if err != nil {
		log.Fatal("RPC could not determine gas price")
	}
	w.Auth.Nonce = big.NewInt(int64(n))
}

func GetNonce(conn BlockChainConnector, a common.Address) (uint64, error) {
	nonce, err := conn.Rpc.PendingNonceAt(context.Background(), a)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return nonce, nil
}

func GetGasPrice(conn BlockChainConnector) (*big.Int, error) {
	gasPrice, err := conn.Rpc.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return (new(big.Int)).SetInt64(0), err
	}
	return gasPrice, nil
}
