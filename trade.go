package d8x_futures

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func PostOrder(conn BlockChainConnector, xInfo StaticExchangeInfo, trdrWallet Wallet, order Order, trader common.Address) (string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, order.Symbol)
	scOrder := order.ToChainType(xInfo, trader)
	//sig := CreateSignature(xInfo, conn.ChainId, scOrder, w)
	trdrWallet.SetGasLimit(uint64(conn.PostOrderGasLimit))
	ob := CreateLimitOrderBookInstance(conn.Rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	sig := []byte{}
	tx, err := ob.PostOrder(trdrWallet.Auth, scOrder, sig)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func CreateBrokerSignature(xInfo StaticExchangeInfo, chainId int64, brokerWallet Wallet, iPerpetualId int32, brokerFeeTbps uint32, traderAddr string, iDeadline uint32) (string, string, error) {
	digestBytes32, err := createBrokerDigest(xInfo, chainId, brokerWallet, iPerpetualId, brokerFeeTbps, traderAddr, iDeadline)
	if err != nil {
		return "", "", err
	}

	sig, err := CreateEvmSignature(digestBytes32[:], brokerWallet.PrivateKey)
	if err != nil {
		return "", "", err
	}
	sigStr := "0x" + common.Bytes2Hex(sig[:])
	dgstStr := common.Bytes2Hex(digestBytes32[:])
	return dgstStr, sigStr, nil
}

// CreateEvmSignature signs the byte data with the given private key
// per EVM convention, so that smart contracts can check the signature
func CreateEvmSignature(data []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	// create ethers-style message to be signed
	s, _ := accounts.TextAndHash(data)
	// Sign and fix postfix
	// See https://medium.com/mycrypto/the-magic-of-digital-signatures-on-ethereum-98fe184dc9c7
	// "The recovery identifier `v`"
	sig, err := crypto.Sign(s, prv)
	if err != nil {
		return []byte{}, err
	}
	if sig[len(sig)-1] == 1 {
		sig[len(sig)-1] = 28
	} else {
		sig[len(sig)-1] = 27
	}
	return sig, nil
}

func createBrokerDigest(xInfo StaticExchangeInfo, chainId int64, w Wallet, iPerpetualId int32, brokerFeeTbps uint32, traderAddr string, iDeadline uint32) ([32]byte, error) {
	domainSeparatorHashBytes32 := getDomainHash(int64(chainId), xInfo.ProxyAddr.String())
	brokerTypeHash := Keccak256FromString("Order(uint24 iPerpetualId,uint16 brokerFeeTbps,address traderAddr,uint32 iDeadline)")
	types := []string{"bytes32", "uint32", "uint16", "address", "uint32"}
	values := []interface{}{brokerTypeHash, uint32(iPerpetualId), uint16(brokerFeeTbps), common.HexToAddress(traderAddr), uint32(iDeadline)}
	structHash, err := AbiEncodeBytes32(types, values...)
	if err != nil {
		return [32]byte{}, err
	}
	var StructHashBytes32 [32]byte
	copy(StructHashBytes32[:], solsha3.SoliditySHA3(structHash))
	types = []string{"bytes32", "bytes32"}
	values = []interface{}{domainSeparatorHashBytes32, StructHashBytes32}
	digest0, err := AbiEncodeBytes32(types, values...)
	if err != nil {
		return [32]byte{}, err
	}
	var digestBytes32 [32]byte
	copy(digestBytes32[:], solsha3.SoliditySHA3(digest0))
	return digestBytes32, nil
}

func getDomainHash(chainId int64, proxyAddress string) [32]byte {
	nameHash := Keccak256FromString("Perpetual Trade Manager")
	domainHash := Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	types := []string{"bytes32", "bytes32", "uint256", "address"}
	values := []interface{}{
		domainHash,
		nameHash,
		big.NewInt(chainId),
		common.HexToAddress(proxyAddress),
	}
	domainSeparator, _ := AbiEncodeBytes32(types, values...)
	dH := solsha3.SoliditySHA3(domainSeparator)
	var DomainSeparatorHashBytes32 [32]byte
	copy(DomainSeparatorHashBytes32[:], dH)
	return DomainSeparatorHashBytes32
}

func CreateOrderDigest(order IClientOrderClientOrder, chainId int, isNewOrder bool, proxyAddress string) (string, error) {
	DomainSeparatorHashBytes32 := getDomainHash(int64(chainId), proxyAddress)
	tradeOrderTypeHash := Keccak256FromString("Order(uint24 iPerpetualId,uint16 brokerFeeTbps,address traderAddr,address brokerAddr,int128 fAmount,int128 fLimitPrice,int128 fTriggerPrice,uint32 iDeadline,uint32 flags,uint16 leverageTDR,uint32 executionTimestamp)")
	types := []string{"bytes32",
		"uint24",
		"uint16",
		"address",
		"address",
		"int128",
		"int128",
		"int128",
		"uint32",
		"uint32",
		"uint16",
		"uint32"}
	values := []interface{}{
		tradeOrderTypeHash,
		order.IPerpetualId,
		order.BrokerFeeTbps,
		order.TraderAddr,
		order.BrokerAddr,
		order.FAmount,
		order.FLimitPrice,
		order.FTriggerPrice,
		order.IDeadline,
		order.Flags,
		order.LeverageTDR,
		order.ExecutionTimestamp,
	}
	structHash, _ := AbiEncodeBytes32(types, values...)
	dH2 := solsha3.SoliditySHA3(structHash)
	var StructHashBytes32 [32]byte
	copy(StructHashBytes32[:], dH2)

	types = []string{"bytes32", "bytes32", "bool"}
	values = []interface{}{DomainSeparatorHashBytes32, StructHashBytes32, isNewOrder}

	digest0, _ := AbiEncodeBytes32(types, values...)
	h := solsha3.SoliditySHA3(digest0)
	var digestBytes32 [32]byte
	copy(digestBytes32[:], h)
	dgstStr := common.Bytes2Hex(digestBytes32[:])
	return dgstStr, nil
}

func AbiEncodeBytes32(types []string, values ...interface{}) ([]byte, error) {
	if len(types) != len(values) {
		return []byte{}, fmt.Errorf("number of types and values do not match")
	}
	byteSlice, err := AbiEncode(types, values...)
	if err != nil {
		return []byte{}, err
	}
	return byteSlice, nil
}

func AbiEncodeHexString(types []string, values ...interface{}) (string, error) {
	if len(types) != len(values) {
		return "", fmt.Errorf("number of types and values do not match")
	}
	byteSlice, err := AbiEncode(types, values...)
	if err != nil {
		return "", err
	}
	result := "0x" + common.Bytes2Hex(byteSlice)
	return result, nil
}

// AbiEncode encodes the provided types (e.g., string, uint256, int32) and
// corresponding values into a hex-string for EVM
func AbiEncode(types []string, values ...interface{}) ([]byte, error) {
	if len(types) != len(values) {
		return []byte{}, fmt.Errorf("number of types and values do not match")
	}

	arguments := abi.Arguments{}
	for _, typ := range types {
		t, err := abi.NewType(typ, "", nil)
		if err != nil {
			return []byte{}, fmt.Errorf("failed to create ABI type: %v", err)
		}
		arguments = append(arguments, abi.Argument{Type: t})
	}

	bytes, err := arguments.Pack(values...)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to encode arguments: %v", err)
	}
	return bytes, nil

}

func Keccak256FromHexString(hexNumber string) Keccak256Hash {
	data, err := hex.DecodeString(strings.TrimPrefix(hexNumber, "0x"))
	if err != nil {
		panic(err)
	}
	hash := solsha3.SoliditySHA3(data)
	var hArr Keccak256Hash
	copy(hArr[:], hash)
	return hArr
}

func Keccak256FromString(v string) Keccak256Hash {
	domainBuf := []byte(v)
	hash := solsha3.SoliditySHA3([]string{"string"}, []interface{}{domainBuf})
	var hArr Keccak256Hash
	copy(hArr[:], hash)
	return hArr
}
