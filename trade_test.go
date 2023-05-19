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
		common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05"),
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
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
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
	if dgst != "901568a48eee260beff786742991585412be01a5fbef7934b679c7980499c9d8" {
		panic("wrong dgst result")
	}

}

func TestPostOrder(t *testing.T) {
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
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	var emptyArray [32]byte
	order := Order{
		Symbol:              "ETH-USD-MATIC",
		Side:                SIDE_BUY,
		Type:                ORDER_TYPE_MARKET,
		Quantity:            333,
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
	txHash, _ := PostOrder(conn, xInfo, wallet, []byte{}, order, traderAddr)
	fmt.Println("Tx hash = ", txHash)
}

func TestBrokerSignature(t *testing.T) {
	privateKey, isPkDefined := os.LookupEnv("PK")
	if !isPkDefined {
		panic("no private key defined")
	}
	config, err := LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	var xInfo StaticExchangeInfo
	xInfo.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	var wallet Wallet
	err = wallet.NewWallet(privateKey, conn)
	if err != nil {
		panic("error creating wallet")
	}
	const brokerFeeTbps = 110
	dgst, sig, _ := CreateBrokerSignature(xInfo, 80001, wallet, 10001, brokerFeeTbps, traderAddr.String(), 1684863656)
	fmt.Print(dgst, sig)
	/* result depend on proxy address
	if dgst != "dead408cb2d42f86476ab484b39e37a354f3cdcbdddb16422af74425324e8755" {
		panic("wrong dgst result")
	}
	if sig != "0x557248de61a8b5b9fb75b51c0b91e42da3be2281d47ed276f5650efaa9b7ada74fac81fb0d34e0887e1a81353f37acdf575d68ae04dce527df4baaa4a41a02f81b" {
		panic("wrong signature result")
	}*/
}
