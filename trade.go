package d8x_futures

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func PostOrder(conn BlockChainConnector, xInfo StaticExchangeInfo, w Wallet, order Order, trader common.Address) (string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, order.Symbol)
	scOrder := order.ToChainType(xInfo, trader)
	//sig := CreateSignature(xInfo, conn.ChainId, scOrder, w)
	w.SetGasLimit(uint64(conn.PostOrderGasLimit))
	ob := CreateLimitOrderBookInstance(conn.Rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	sig := []byte{}
	tx, err := ob.PostOrder(w.Auth, scOrder, sig)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}

/*
	func CreateSignature(xInfo StaticExchangeInfo, chainId int64, o IClientOrderClientOrder, w Wallet) []byte {
		const isNewOrder = true
		var proxyAddr = xInfo.ProxyAddr

}
*/
func CreateOrderDigest(order IClientOrderClientOrder, chainId int, isNewOrder bool, proxyAddress string) (string, error) {
	nameHash := Keccak256FromString("Perpetual Trade Manager")
	fmt.Println("nameHash=", common.Bytes2Hex(nameHash[:]))
	domainHash := Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	types := []string{"bytes32", "bytes32", "uint256", "address"}
	values := []interface{}{
		domainHash,
		nameHash,
		big.NewInt(int64(chainId)),
		common.HexToAddress(proxyAddress),
	}
	domainSeparator2, _ := AbiEncodeHexString(types, values...)
	myres := Keccak256FromHexString(domainSeparator2)
	fmt.Println("myresKeccakDomainSeparator=", common.Bytes2Hex(myres[:]))
	data, _ := hex.DecodeString(strings.TrimPrefix(domainSeparator2, "0x"))
	dH := solsha3.SoliditySHA3(data)
	//dH := Keccak256FromString(domainSeparator2)
	fmt.Println(dH)
	fmt.Println("domainSeparator=", common.Bytes2Hex(dH[:]))
	fmt.Println(domainSeparator2)

	tradeOrderTypeHash := Keccak256FromString("Order(uint24 iPerpetualId,uint16 brokerFeeTbps,address traderAddr,address brokerAddr,int128 fAmount,int128 fLimitPrice,int128 fTriggerPrice,uint32 iDeadline,uint32 flags,uint16 leverageTDR,uint32 executionTimestamp)")
	fmt.Println("tradeOrderTypeHash=", common.Bytes2Hex(tradeOrderTypeHash[:]))
	types = []string{"bytes32",
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
	values = []interface{}{
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
	fmt.Println("orderamt=", order.FAmount)
	structHash, _ := AbiEncodeBytes32(types, values...)
	structHash2, _ := AbiEncodeHexString(types, values...)
	fmt.Println("structHash=", structHash2)
	// identical
	fmt.Println("structHash=", common.Bytes2Hex(structHash[:]))
	types = []string{"bytes32", "bytes32", "bool"}
	var DHbyteArray [32]byte
	copy(DHbyteArray[:], dH)

	data2, _ := hex.DecodeString(strings.TrimPrefix(structHash2, "0x"))
	dH2 := solsha3.SoliditySHA3(data2)
	var DHbyteArray2 [32]byte
	copy(DHbyteArray2[:], dH2)

	values = []interface{}{DHbyteArray, DHbyteArray2, isNewOrder}
	digest0, _ := AbiEncodeHexString(types, values...)
	digest := Keccak256FromHexString(digest0)
	dgstStr := common.Bytes2Hex(digest[:])
	// all identical to typescript -- but a little cumbersome
	return dgstStr, nil
}

func AbiEncodeBytes32(types []string, values ...interface{}) ([32]byte, error) {
	if len(types) != len(values) {
		return [32]byte{}, fmt.Errorf("number of types and values do not match")
	}
	byteSlice, err := AbiEncode(types, values...)
	if err != nil {
		return [32]byte{}, err
	}
	var byteArray [32]byte
	copy(byteArray[:], byteSlice)
	return byteArray, nil
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
