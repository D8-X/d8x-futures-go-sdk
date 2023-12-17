package d8x_futures

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// PostOrder posts an order to the corresponding limit order book.
// Returns orderId, tx hash, error
func (sdk *Sdk) PostOrder(order *Order) (string, string, error) {
	return RawPostOrder(&sdk.Conn, &sdk.Info, sdk.Wallet, []byte{}, order, sdk.Wallet.Address)
}

func (sdk *Sdk) CreateOrderBrokerSignature(iPerpetualId int32, brokerFeeTbps uint32, traderAddr string, iDeadline uint32) (string, string, error) {
	return RawCreateOrderBrokerSignature(sdk.ChainConfig.ProxyAddr,
		sdk.ChainConfig.ChainId, sdk.Wallet, iPerpetualId, brokerFeeTbps, traderAddr, iDeadline)
}

func (sdk *Sdk) CreatePaymentBrokerSignature(ps *PaySummary) (string, string, error) {
	return RawCreatePaymentBrokerSignature(ps, sdk.Wallet)
}

func (sdk *Sdk) AddCollateral(symbol string, amountCC float64) (*types.Transaction, error) {
	return RawAddCollateral(&sdk.Conn, &sdk.Info, sdk.ChainConfig.PriceFeedEndpoints[0], sdk.Wallet, symbol, amountCC)
}

// ApproveTknSpending approves the manager to spend the wallet's margin tokens for the given
// pool (via symbol), if amount = nil, max approval. Symbol is a perpetual, but approval
// is for pool.
func (sdk *Sdk) ApproveTknSpending(symbol string, amount *big.Int) (*types.Transaction, error) {
	tknAddr, err := RawGetMarginTknAddr(&sdk.Info, symbol)
	if err != nil {
		return nil, err
	}
	erc20Instance, err := contracts.NewErc20(tknAddr, sdk.Conn.Rpc)
	if err != nil {
		return nil, errors.New("ApproveTknSpending: creating instance of token " + tknAddr.String())
	}
	var amt *big.Int
	if amount == nil {
		amt = utils.MaxUint256()
	} else {
		amt = amount
	}
	sdk.Wallet.UpdateNonce(sdk.Conn.Rpc)
	approvalTx, err := erc20Instance.Approve(sdk.Wallet.Auth, sdk.Info.ProxyAddr, amt)
	if err != nil {
		return nil, errors.New("Error approving token for chain " + strconv.Itoa(int(sdk.Conn.ChainId)) + ": " + err.Error())
	}
	return approvalTx, nil
}

// RawPostOrder posts an order to the correct limit order book.
// Returns orderId, tx hash, error
// It needs the private key for the wallet
// paying the gas fees. If the trader-address is not the address corresponding to the postingWallet, the func
// also needs signature from the trader
func RawPostOrder(conn *BlockChainConnector, xInfo *StaticExchangeInfo, postingWallet *Wallet, traderSig []byte, order *Order, trader common.Address) (string, string, error) {
	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, order.Symbol)
	scOrder := order.ToChainType(xInfo, trader)
	scOrders := []contracts.IClientOrderClientOrder{scOrder}
	tsigs := [][]byte{traderSig}
	g := postingWallet.Auth.GasLimit
	defer postingWallet.SetGasLimit(g)
	postingWallet.SetGasLimit(uint64(conn.PostOrderGasLimit))
	ob := CreateLimitOrderBookInstance(conn.Rpc, xInfo.Perpetuals[j].LimitOrderBookAddr)
	postingWallet.UpdateNonce(conn.Rpc)
	dgst, err := CreateOrderDigest(scOrder, int(conn.ChainId), true, xInfo.ProxyAddr.Hex())
	if err != nil {
		return "", "", err
	}
	id := CreateOrderId(dgst)
	tx, err := ob.PostOrders(postingWallet.Auth, scOrders, tsigs)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return id, tx.Hash().Hex(), nil
}

// RawAddCollateral adds (amountCC>0) or removes (amountCC<0) collateral to/from the margin account of the given perpetual
func RawAddCollateral(conn *BlockChainConnector, xInfo *StaticExchangeInfo, pythEndpoint string, postingWallet *Wallet, symbol string, amountCC float64) (*types.Transaction, error) {
	if amountCC == 0 {
		return nil, errors.New("RawAddCollateral: amount 0")
	}

	j := GetPerpetualStaticInfoIdxFromSymbol(xInfo, symbol)
	id := int64(xInfo.Perpetuals[j].Id)
	amount := utils.Float64ToABDK(math.Abs(amountCC))
	perpCtrct := CreatePerpetualManagerInstance(conn.Rpc, xInfo.ProxyAddr)
	postingWallet.UpdateNonce(conn.Rpc)
	pxFeed, err := fetchPricesForPerpetual(*xInfo, j, pythEndpoint)
	if err != nil {
		return nil, errors.New("RawAddCollateral: failed fetching oracle prices " + err.Error())
	}
	var tx *types.Transaction
	v := postingWallet.Auth.Value
	defer func() { postingWallet.Auth.Value = v }()
	val := conn.PriceFeedConfig.PriceUpdateFeeGwei * int64(len(pxFeed.PriceFeed.PublishTimes))
	postingWallet.Auth.Value = big.NewInt(val)

	g := postingWallet.Auth.GasLimit
	defer postingWallet.SetGasLimit(g)
	postingWallet.SetGasLimit(uint64(15_000_000))

	if amountCC > 0 {
		tx, err = perpCtrct.Deposit(postingWallet.Auth, big.NewInt(id),
			postingWallet.Address, amount, pxFeed.PriceFeed.Vaas, pxFeed.PriceFeed.PublishTimes)
		if err != nil {
			return nil, errors.New("RawAddCollateral:" + err.Error())
		}
	} else {
		tx, err = perpCtrct.Withdraw(postingWallet.Auth, big.NewInt(id),
			postingWallet.Address, amount, pxFeed.PriceFeed.Vaas, pxFeed.PriceFeed.PublishTimes)
		if err != nil {
			return nil, errors.New("RawAddCollateral:" + err.Error())
		}
	}
	return tx, nil
}

func RawCreateOrderBrokerSignature(proxyAddr common.Address, chainId int64, brokerWallet *Wallet, iPerpetualId int32, brokerFeeTbps uint32, traderAddr string, iDeadline uint32) (string, string, error) {
	digestBytes32, err := createOrderBrokerDigest(proxyAddr, chainId, iPerpetualId, brokerFeeTbps, traderAddr, iDeadline)
	if err != nil {
		return "", "", err
	}
	if brokerWallet.PrivateKey == nil {
		return "", "", fmt.Errorf("Broker key not defined")
	}
	sig, err := CreateEvmSignature(digestBytes32[:], brokerWallet.PrivateKey)
	if err != nil {
		return "", "", err
	}
	sigStr := "0x" + common.Bytes2Hex(sig[:])
	dgstStr := common.Bytes2Hex(digestBytes32[:])
	return dgstStr, sigStr, nil
}

func RawCreatePaymentBrokerSignature(ps *PaySummary, brokerWallet *Wallet) (string, string, error) {
	digestBytes32, err := createPaymentBrokerDigest(ps)
	if err != nil {
		return "", "", err
	}
	if brokerWallet.PrivateKey == nil {
		return "", "", fmt.Errorf("Broker key not defined")
	}
	sig, err := CreateEvmSignature(digestBytes32[:], brokerWallet.PrivateKey)
	if err != nil {
		return "", "", err
	}
	sigStr := "0x" + common.Bytes2Hex(sig[:])
	dgstStr := common.Bytes2Hex(digestBytes32[:])
	return dgstStr, sigStr, nil
}

// RecoverPaymentSignatureAddr recovers the address that created the signature of PaySummary data for
// the given chainId and multiPayContract address.
// The function returns the recovered address or an error
func RecoverPaymentSignatureAddr(sig []byte, ps *PaySummary) (common.Address, error) {
	digestBytes32, err := createPaymentBrokerDigest(ps)
	if err != nil {
		return common.Address{}, err
	}

	addr, err := RecoverEvmAddress(digestBytes32[:], sig[:])
	if err != nil {
		return common.Address{}, err
	}
	return addr, nil
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

func RecoverEvmAddress(data, signature []byte) (common.Address, error) {
	recoveryID := signature[len(signature)-1]
	if recoveryID != 27 && recoveryID != 28 {
		return common.Address{}, fmt.Errorf("invalid recovery id")
	}
	s, _ := accounts.TextAndHash(data)
	adjustedSignature := append(signature[:len(signature)-1], recoveryID-27)

	// Recover public key from encoded digest and signature
	publicKey, err := crypto.SigToPub(s, adjustedSignature)
	if err != nil {
		return common.Address{}, err
	}

	// Get Ethereum address from the public key
	recoveredAddress := crypto.PubkeyToAddress(*publicKey)

	return recoveredAddress, nil
}

func createOrderBrokerDigest(proxyAddr common.Address, chainId int64, iPerpetualId int32, brokerFeeTbps uint32, traderAddr string, iDeadline uint32) ([32]byte, error) {
	domainSeparatorHashBytes32 := getDomainHash("Perpetual Trade Manager", int64(chainId), proxyAddr.String())
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

func createPaymentBrokerDigest(ps *PaySummary) ([32]byte, error) {
	domainSeparatorHashBytes32 := getDomainHash("Multipay", ps.ChainId, ps.MultiPayCtrct.String())
	typeHash := Keccak256FromString("PaySummary(address payer,address executor,address token,uint32 timestamp,uint32 id,uint256 totalAmount)")
	types := []string{"bytes32", "address", "address", "address", "uint32", "uint32", "uint256"}
	values := []interface{}{typeHash, ps.Payer, ps.Executor, ps.Token, ps.Timestamp, ps.Id, ps.TotalAmount}
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

func getDomainHash(name string, chainId int64, contractAddr string) [32]byte {
	nameHash := Keccak256FromString(name)
	domainHash := Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	types := []string{"bytes32", "bytes32", "uint256", "address"}
	values := []interface{}{
		domainHash,
		nameHash,
		big.NewInt(chainId),
		common.HexToAddress(contractAddr),
	}
	domainSeparator, _ := AbiEncodeBytes32(types, values...)
	dH := solsha3.SoliditySHA3(domainSeparator)
	var DomainSeparatorHashBytes32 [32]byte
	copy(DomainSeparatorHashBytes32[:], dH)
	return DomainSeparatorHashBytes32
}

func CreateOrderDigest(order contracts.IClientOrderClientOrder, chainId int, isNewOrder bool, proxyAddress string) (string, error) {
	DomainSeparatorHashBytes32 := getDomainHash("Perpetual Trade Manager", int64(chainId), proxyAddress)
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
	structHash, err := AbiEncodeBytes32(types, values...)
	if err != nil {
		return "", err
	}
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

func CreateOrderId(orderDigest string) string {
	bytesDigest := common.Hex2Bytes(strings.TrimPrefix(orderDigest, "0x"))
	var orderDigest32 [32]byte
	copy(orderDigest32[:], bytesDigest)
	// create ethers-style message with prefix "\x19Ethereum Signed Message:\n"
	s, _ := accounts.TextAndHash(orderDigest32[:])
	idStr := common.Bytes2Hex(s)
	return idStr
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

func BytesFromHexString(hexNumber string) ([]byte, error) {
	data, err := hex.DecodeString(strings.TrimPrefix(hexNumber, "0x"))
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func Bytes32FromHexString(hexNumber string) Keccak256Hash {
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
