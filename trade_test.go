package d8x_futures

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
		ReduceOnly:          false,
		LimitPrice:          0,
		TriggerPrice:        0,
		KeepPositionLvg:     false,
		BrokerFeeTbps:       0,
		BrokerAddr:          common.Address{},
		BrokerSignature:     []byte{},
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
	_, execPk, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	var wallet Wallet
	wallet.NewWallet(fmt.Sprintf("%x", execPk.D), conn.ChainId, conn.Rpc)
	var xInfo StaticExchangeInfo
	xInfo.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("***REMOVED***")
	var emptyArray [32]byte
	order := Order{
		Symbol:              "ETH-USD-MATIC",
		Side:                SIDE_BUY,
		Type:                ORDER_TYPE_MARKET,
		Quantity:            333,
		ReduceOnly:          false,
		LimitPrice:          0,
		TriggerPrice:        0,
		KeepPositionLvg:     false,
		BrokerFeeTbps:       0,
		BrokerAddr:          common.Address{},
		BrokerSignature:     []byte{},
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
	_, execPk, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	var xInfo StaticExchangeInfo
	xInfo.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("***REMOVED***")
	var wallet Wallet
	err = wallet.NewWallet(fmt.Sprintf("%x", execPk.D), config.ChainId, nil)
	if err != nil {
		panic("error creating wallet")
	}
	const brokerFeeTbps = 110
	dgst, sig, _ := CreateOrderBrokerSignature(xInfo.ProxyAddr, 80001, wallet, 10001, brokerFeeTbps, traderAddr.String(), 1684863656)
	fmt.Print(dgst, sig)
	/* result depend on proxy address
	if dgst != "dead408cb2d42f86476ab484b39e37a354f3cdcbdddb16422af74425324e8755" {
		panic("wrong dgst result")
	}
	if sig != "0x557248de61a8b5b9fb75b51c0b91e42da3be2281d47ed276f5650efaa9b7ada74fac81fb0d34e0887e1a81353f37acdf575d68ae04dce527df4baaa4a41a02f81b" {
		panic("wrong signature result")
	}*/
}

func TestPaymentSignature(t *testing.T) {
	_, execPk, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}
	config, err := utils.LoadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	var wallet Wallet
	err = wallet.NewWallet(fmt.Sprintf("%x", execPk.D), config.ChainId, nil)
	fmt.Printf("\nwallet addr %s\n", wallet.Address.String())
	if err != nil {
		panic("error creating wallet")
	}
	var totalAmount big.Int
	totalAmount.SetString("1000000000000000000000", 10)

	executorKey, err := crypto.GenerateKey()
	if err != nil {
		t.Logf(err.Error())
	}
	tokenAddr := common.HexToAddress("0xeE53d536DFC355017147058a4197cAD3b4ac1964")
	multiPayCtrct := common.HexToAddress("0x30b55550e02B663E15A95B50850ebD20363c2AD5")
	timestamp := time.Now().UTC().Unix()
	var ps PaySummary
	ps.Payer = wallet.Address
	ps.Executor = crypto.PubkeyToAddress(executorKey.PublicKey)
	ps.Timestamp = uint32(timestamp)
	ps.Token = tokenAddr
	ps.TotalAmount = &totalAmount
	ps.MultiPayCtrct = multiPayCtrct
	ps.ChainId = config.ChainId
	dgst, sig, err := CreatePaymentBrokerSignature(ps, wallet)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Print(dgst, sig)
	sigB, err := BytesFromHexString(sig)
	if err != nil {
		t.Logf(err.Error())
	}
	addr, err := RecoverPaymentSignatureAddr(sigB, ps)
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Printf("\nrecovered address %s", addr.String())
	fmt.Printf("\nbroker address %s", wallet.Address.String())
	fmt.Printf("\nexecutor address %s", ps.Executor.String())
}

func TestSignOrder(t *testing.T) {
	chainId := 80001
	perpId := 10001
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf(err.Error())
	}
	// Derive the Ethereum address from the private key
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	var wallet Wallet
	pk := fmt.Sprintf("%x", privateKey.D)
	err = wallet.NewWallet(pk, int64(chainId), nil)
	proxyAddr := "0xCdd7C9e07689d1B3D558A714fAa5Cc4B6bA654bD"
	fmt.Printf("\nwallet addr %s\n", wallet.Address.String())
	if err != nil {
		panic("error creating wallet")
	}

	digest, sig, err := CreateOrderBrokerSignature(
		common.HexToAddress(proxyAddr), int64(chainId), wallet, int32(perpId), uint32(4000),
		"***REMOVED***", 1691249493)
	if err != nil {
		t.Errorf("signing order: %v", err)
	}
	// now encode again
	sigBytes, err := BytesFromHexString(sig)
	if err != nil {
		t.Errorf("decoding signature: %v", err)
	}
	digBytes, err := BytesFromHexString(digest)
	if err != nil {
		t.Errorf("decoding digest: %v", err)
	}
	fmt.Println("digest = ")
	addrRecovered, err := RecoverEvmAddress(digBytes, sigBytes)
	if err != nil {
		t.Errorf("recovering address: %v", err)
	} else {
		t.Logf("recovered address")
		t.Logf(addrRecovered.String())
	}
	t.Log("recovered addr = ", addrRecovered.String())
	t.Log("signer    addr = ", addr.String())
	if addrRecovered.String() == addr.String() {
		t.Logf("recovered address correct")
	} else {
		t.Errorf("recovering address incorrect")
	}
}

func generateKey() (common.Address, *ecdsa.PrivateKey, error) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return common.Address{}, nil, err
	}
	// Derive the Ethereum address from the private key
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	return addr, privateKey, err
}
