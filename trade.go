package d8x_futures

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

/*
	func PostOrder(conn BlockChainConnector, xInfo StaticExchangeInfo, w Wallet, order Order, trader common.Address) {
		j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, order.Symbol)
		scOrder := order.ToChainType(xInfo, trader)
		sig := CreateSignature(xInfo, conn.ChainId, scOrder, w)
		ob := CreateLimitOrderBookInstance(conn.Rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
		tx, err := ob.PostOrder(w.Auth, scOrder, sig)
	}

	func CreateSignature(xInfo StaticExchangeInfo, chainId int64, o IClientOrderClientOrder, w Wallet) []byte {
		const isNewOrder = true
		var proxyAddr = xInfo.ProxyAddr

}

	func CreateDigest(order IClientOrderClientOrder, chainID int, isNewOrder bool, proxyAddress string) (string, error) {
		var NAME = "Perpetual Trade Manager"
		var DOMAIN_TYPEHASH = "0x" + hashString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")

		domainSeparator := "0x" + hashBytes(abiEncode([]interface{}{
			[]byte(NAME),
			big.NewInt(int64(chainID)),
			proxyAddress,
		}))

		var TRADE_ORDER_TYPEHASH = "0x" + hashString("Order(uint24 iPerpetualId,uint16 brokerFeeTbps,address traderAddr,address brokerAddr,int128 fAmount,int128 fLimitPrice,int128 fTriggerPrice,uint32 iDeadline,uint32 flags,uint16 leverageTDR,uint32 executionTimestamp)")

		structHash := "0x" + hashBytes(abiEncode([]interface{}{
			order.iPerpetualId,
			order.brokerFeeTbps,
			order.traderAddr,
			order.brokerAddr,
			order.fAmount,
			order.fLimitPrice,
			order.fTriggerPrice,
			order.iDeadline,
			order.flags,
			order.leverageTDR,
			order.executionTimestamp,
		}))

		digest := "0x" + hashBytes(abiEncode([]interface{}{
			domainSeparator,
			structHash,
			isNewOrder,
		}))

		return digest, nil
	}

	func hashString(data string) string {
		hash := sha256.New()
		hash.Write([]byte(data))
		return hex.EncodeToString(hash.Sum(nil))
	}

	func hashBytes(data []byte) string {
		hash := sha256.New()
		hash.Write(data)
		return hex.EncodeToString(hash.Sum(nil))
	}

	func abiEncode(params []interface{}) []byte {
		// Encode parameters using the desired encoding format
		// This implementation assumes the use of the Ethereum ABI encoding
		var buf bytes.Buffer
		for _, param := range params {
			// Perform the appropriate encoding for each parameter type
			switch param := param.(type) {
			case int:
				buf.Write(big.NewInt(int64(param)).Bytes())
			case int64:
				buf.Write(big.NewInt(param).Bytes())
			case uint32:
				buf.Write(big.NewInt(int64(param)).Bytes())
			case uint64:
				buf.Write(big.NewInt(int64(param)).Bytes())
			case *big.Int:
				buf.Write(param.Bytes())
			case bool:
				if param {
					buf.WriteByte(0x01)
				} else {
					buf.WriteByte(0x00)
				}
			case string:
				buf.Write([]byte(param))
			case []byte:
				buf.Write(param)
				// Add cases for other parameter types if needed
			}
		}
		return buf.Bytes()
	}
*/

func BufferFrom(v string) []byte {
	decoded, err := hex.DecodeString(v)
	if err != nil {
		panic(err)
	}
	return decoded
}

func AbiEncode(types []string, values ...interface{}) (string, error) {
	if len(types) != len(values) {
		return "", fmt.Errorf("number of types and values do not match")
	}

	arguments := abi.Arguments{}
	for _, typ := range types {
		t, err := abi.NewType(typ, "", nil)
		if err != nil {
			return "", fmt.Errorf("failed to create ABI type: %v", err)
		}
		arguments = append(arguments, abi.Argument{Type: t})
	}

	bytes, err := arguments.Pack(values...)
	if err != nil {
		return "", fmt.Errorf("failed to encode arguments: %v", err)
	}

	result := "0x" + common.Bytes2Hex(bytes)
	return result, nil
}

func (hArr *Keccak256Hash) Keccak256FromString(v string) {
	domainBuf := []byte(v)
	hash := solsha3.SoliditySHA3([]string{"string"}, []interface{}{domainBuf})
	copy(hArr[:], hash)
}
