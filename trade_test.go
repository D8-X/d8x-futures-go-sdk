package d8x_futures

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestABI(t *testing.T) {
	types := []string{"uint256", "address", "int128", "bytes32"}
	//domainBuf := []byte("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	domainHash := Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	//v := "0x" + hex.EncodeToString(hash)
	//var hashArray [32]byte
	//copy(hashArray[:], hash)
	values := []interface{}{
		big.NewInt(123),
		common.HexToAddress("***REMOVED***"),
		big.NewInt(-1211),
		domainHash,
	}

	result, err := AbiEncodeHexString(types, values...)
	if err != nil {
		fmt.Println("Encoding error:", err)
		return
	}

	fmt.Println("Encoded data:", result)
	var resExpected = "0x000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000009d5aab428e98678d0e645ea4aebd25f744341a05fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb458cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866"
	if resExpected != result {
		panic("wrong encoding")
	}
}

func TestOrderHash(t *testing.T) {
	var info StaticExchangeInfo
	info.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("***REMOVED***")
	var emptyArray [32]byte
	order := Order{
		Symbol:              "ETH-USD-MATIC",
		Side:                SIDE_BUY,
		Type:                ORDER_TYPE_MARKET,
		Quantity:            15,
		reduceOnly:          false,
		LimitPrice:          0,
		TriggerPrice:        0,
		KeepPositionLvg:     false,
		BrokerFeeTbps:       0,
		BrokerAddr:          common.Address{},
		BrokerSignature:     []byte{},
		Flags:               MASK_MARKET_ORDER,
		StopPrice:           0,
		Leverage:            5,
		Deadline:            1684863656,
		ExecutionTimestamp:  1684263656,
		parentChildOrderId1: emptyArray,
		parentChildOrderId2: emptyArray,
	}
	scOrder := order.ToChainType(info, traderAddr)
	dgst, err := CreateOrderDigest(scOrder, 80001, true, info.ProxyAddr.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(dgst)
	if dgst != "40745ec5cad712682c21487bcae978059f6ab00a6dbfbbcbb95eecb73e331af9" {
		panic("wrong dgst result")
	}

}

func TestPostOrder(t *testing.T) {
	//privateKey := os.Getenv("PK")
	privateKey, isPkDefined := os.LookupEnv("PK")
	if !isPkDefined {
		panic("no private key defined")
	}
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	var wallet Wallet
	wallet.NewWallet(privateKey, conn)
	var xInfo StaticExchangeInfo
	xInfo.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("***REMOVED***")
	var emptyArray [32]byte
	order := Order{
		Symbol:              "ETH-USD-MATIC",
		Side:                SIDE_BUY,
		Type:                ORDER_TYPE_MARKET,
		Quantity:            15,
		reduceOnly:          false,
		LimitPrice:          0,
		TriggerPrice:        0,
		KeepPositionLvg:     false,
		BrokerFeeTbps:       0,
		BrokerAddr:          common.Address{},
		BrokerSignature:     []byte{},
		Flags:               MASK_MARKET_ORDER,
		StopPrice:           0,
		Leverage:            5,
		Deadline:            1684863656,
		ExecutionTimestamp:  1684263656,
		parentChildOrderId1: emptyArray,
		parentChildOrderId2: emptyArray,
	}
	txHash, _ := PostOrder(conn, xInfo, wallet, order, traderAddr)
	fmt.Println("Tx hash = ", txHash)
}
