package d8x_futures

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestABI(t *testing.T) {
	types := []string{"uint256", "address", "int128", "bytes32"}
	domainHash := Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	values := []interface{}{
		big.NewInt(123),
		common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05"),
		big.NewInt(-1211),
		domainHash,
	}

	result, err := AbiEncodeHexString(types, values...)
	if err != nil {
		t.Fatalf("encoding error: %v", err)
	}

	t.Log("Encoded data:", result)
	resExpected := "0x000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000009d5aab428e98678d0e645ea4aebd25f744341a05fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb458cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866"
	if resExpected != result {
		t.Fatal("wrong encoding")
	}
}

func TestOrderHash(t *testing.T) {
	var info StaticExchangeInfo
	if err := info.Load("./tmpXchInfo.json"); err != nil {
		t.Skip("tmpXchInfo.json not found (run integration tests first): ", err)
	}
	// Check if the required symbol exists in the loaded exchange info
	const testSymbol = "BTC-USD-MATIC"
	found := false
	for _, sym := range info.PerpetualIdToSymbol {
		if sym == testSymbol {
			found = true
			break
		}
	}
	if !found {
		t.Skipf("%s not found in exchange info (test requires specific chain data)", testSymbol)
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	var emptyArray [32]byte
	order := Order{
		Symbol:              testSymbol,
		Side:                SIDE_BUY,
		Type:                ORDER_TYPE_MARKET,
		Quantity:            0.16,
		ReduceOnly:          false,
		LimitPrice:          100,
		TriggerPrice:        0,
		KeepPositionLvg:     false,
		BrokerFeeTbps:       0,
		BrokerAddr:          common.Address{},
		BrokerSignature:     []byte{},
		Leverage:            2,
		Deadline:            1702584746,
		ExecutionTimestamp:  1702504746,
		ParentChildOrderId1: emptyArray,
		ParentChildOrderId2: emptyArray,
	}
	scOrder := order.ToChainType(&info, traderAddr)
	// override
	lim, _ := new(big.Int).SetString("640000000000000000", 16)
	scOrder.FLimitPrice = lim
	q, _ := new(big.Int).SetString("28f5c28f5c28f5f9", 16)
	scOrder.FAmount = q

	dgst, err := CreateOrderDigest(scOrder, 1442, true, info.ProxyAddr.String())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dgst)
	if dgst != "625da4b5a8ccc7a093e88b5908b3019b3272c30058bc17d420e1b0e2a60ce0f5" {
		t.Fatal("wrong dgst result")
	}
}

func TestPaymentSignature(t *testing.T) {
	_, execPk, err := generateKey()
	if err != nil {
		t.Fatalf("generateKey: %v", err)
	}
	chConfig, err := config.GetDefaultChainConfig("arbitrum")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	wallet, err := NewWallet(fmt.Sprintf("%064x", execPk.D), chConfig.ChainId)
	if err != nil {
		t.Fatalf("NewWallet: %v", err)
	}
	t.Logf("wallet addr %s", wallet.Address.String())
	var totalAmount big.Int
	totalAmount.SetString("1000000000000000000000", 10)

	executorKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
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
	ps.ChainId = chConfig.ChainId
	dgst, sig, err := RawCreatePaymentBrokerSignature(&ps, wallet)
	if err != nil {
		t.Fatalf("RawCreatePaymentBrokerSignature: %v", err)
	}
	t.Log(dgst, sig)
	sigB, err := BytesFromHexString(sig)
	if err != nil {
		t.Fatalf("BytesFromHexString: %v", err)
	}
	addr, err := RecoverPaymentSignatureAddr(sigB, &ps)
	if err != nil {
		t.Fatalf("RecoverPaymentSignatureAddr: %v", err)
	}
	t.Logf("recovered address %s", addr.String())
	t.Logf("broker address %s", wallet.Address.String())
	t.Logf("executor address %s", ps.Executor.String())
}

func TestSignOrder(t *testing.T) {
	chainId := 80001
	perpId := 10001
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	// Derive the Ethereum address from the private key
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	pk := fmt.Sprintf("%064x", privateKey.D)
	wallet, err := NewWallet(pk, int64(chainId))
	if err != nil {
		t.Fatalf("NewWallet: %v", err)
	}
	t.Logf("wallet addr %s", wallet.Address.String())
	proxyAddr := "0xCdd7C9e07689d1B3D558A714fAa5Cc4B6bA654bD"

	digest, sig, err := RawCreateOrderBrokerSignature(
		common.HexToAddress(proxyAddr), int64(chainId), wallet, int32(perpId), uint32(4000),
		"0x9d5aaB428e98678d0E645ea4AeBd25f744341a05", 1691249493)
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
	t.Log("digest = ")
	addrRecovered, err := RecoverEvmAddress(digBytes, sigBytes)
	if err != nil {
		t.Errorf("recovering address: %v", err)
	} else {
		t.Log("recovered address")
		t.Log(addrRecovered.String())
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
		return common.Address{}, nil, err
	}
	// Derive the Ethereum address from the private key
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	return addr, privateKey, err
}
