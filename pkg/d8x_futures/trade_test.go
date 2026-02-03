package d8x_futures

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func loadPk() string {
	viper.SetConfigFile("../../.env")
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("could not load .env file" + err.Error())
	}
	return viper.GetString("PK")
}

func TestSdkExec(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	rpc, err := ethclient.Dial("https://sepolia.base.org")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	// https://sports.quantena.org/slots-info/84532
	perp := "NHL0-USD-PUSD"
	fmt.Printf("wallet addr =%s\n", sdk.Wallets[0].Address.Hex())
	orderObj, err := sdk.QueryAllOpenOrders(perp, nil)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	orders := orderObj.Orders
	ids := orderObj.OrderHashes
	var mktOrderIds []string
	for k, order := range orders {
		if order.Type == ORDER_TYPE_MARKET {
			mktOrderIds = append(mktOrderIds, ids[k])
		}
	}
	if len(mktOrderIds) == 0 {
		// mktOrderIds = append(mktOrderIds, "f3e6741c2eefeb3c5e1c8539f15e5590b8b104ad6c8ed99a2f63151c315b0dd0")

		/*tx, err := sdk.ApproveTknSpending(perp, nil, nil)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		fmt.Println(tx.Hash())
		*/
		order, err := sdk.NewOrder(perp, SIDE_SELL, ORDER_TYPE_MARKET, 20, 1, &OrderOptions{LimitPrice: 0})
		if err != nil {
			t.Log(err.Error())
			t.FailNow()
		}
		orderId, tx, err := sdk.PostOrder(order, nil)
		if err != nil {
			t.Log(err.Error())
			t.FailNow()
		}
		fmt.Println("order id =", orderId)
		_, err = bind.WaitMined(context.Background(), rpc, tx)
		if err != nil {
			t.Log(err.Error())
			t.FailNow()
		}
		// 92d891d3ae6d8695c8732d67fff2b59d309496796e1f369f8d5e0b4ab2a17cd4
		mktOrderIds = append(mktOrderIds, orderId)

	}
	now := time.Now().Unix()
	payoutAddr := common.Address{} // common.HexToAddress("0x98DfAFF5126836E339493a6021FD5B92Bf005F0D")
	tx, err := sdk.ExecuteOrders(perp, []string{mktOrderIds[0]}, &OptsOverridesExec{TsMin: uint32(now), PayoutAddr: payoutAddr})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Println("tx = ", tx.Hash())
	fmt.Println("done")
}

func TestSdkExecWithBroker(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	rpc, err := ethclient.Dial("https://sepolia.base.org")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	// https://sports.quantena.org/slots-info/84532
	perp := "NHL0-USD-PUSD"
	fmt.Printf("wallet addr =%s\n", sdk.Wallets[0].Address.Hex())
	order, err := sdk.NewOrder(perp, SIDE_SELL, ORDER_TYPE_MARKET, 20, 1, &OrderOptions{LimitPrice: 0})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	const url = "https://broker-84532.d8x.xyz"
	order.BrokerAddr, err = GetBrokerAddress(url)
	fmt.Println("broker=", order.BrokerAddr.String())
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	order.OptTraderAddr = sdk.Wallets[0].Address
	order.BrokerFeeTbps, err = GetBrokerFeeTbps(sdk.Wallets[0].Address, int(sdk.ChainConfig.ChainId), url)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	err = sdk.SignOrder(order, url)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	orderId, tx, err := sdk.PostOrder(order, nil)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	fmt.Println("order id =", orderId)
	_, err = bind.WaitMined(context.Background(), rpc, tx)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	now := time.Now().Unix()
	payoutAddr := common.Address{} // common.HexToAddress("0x98DfAFF5126836E339493a6021FD5B92Bf005F0D")
	tx, err = sdk.ExecuteOrders(perp, []string{orderId}, &OptsOverridesExec{TsMin: uint32(now), PayoutAddr: payoutAddr})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Println("tx = ", tx.Hash())
	fmt.Println("done")
}

func TestSdkLiquidatePosition(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	// err := sdk.New([]string{pk}, "195") //x-layer testnet
	// err := sdk.New([]string{pk}, "196") //x-layer testnet
	sdk, err := NewSdk([]string{pk}, "421614") // arbitrum sepolia
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	acc2, err := sdk.QueryLiquidatableAccountsInPool(1, nil)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	fmt.Println(acc2)
	if len(acc2) == 0 {
		return
	}
	for _, el := range acc2 {
		for _, addr := range el.LiqAccounts {
			fmt.Printf("liquidating %s\n", addr.Hex())
			tx, err := sdk.LiquidatePosition(el.PerpId, []common.Address{addr}, nil, nil)
			if err != nil {
				fmt.Printf("error liquidating: %s\n", err.Error())
				continue
			}
			fmt.Printf("liquidated %d trader %s tx=%s\n", el.PerpId, addr.Hex(), tx.Hash())
		}
	}
}

func TestAddCollateral(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	// err := sdk.New([]string{pk}, "421614")
	sdk, err := NewSdk([]string{pk}, "195")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	fmt.Println("wallet = " + sdk.Wallets[0].Address.Hex())
	// opts := nil //OptsOverrides{GasLimit: }
	// tx, err := sdk.AddCollateral("ETH-USD-WEETH", 0.0001, nil)
	tx, err := sdk.AddCollateral("BTC-USDC-USDC", -0.0001, nil)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	} else {
		fmt.Println("tx hash adding collateral=", tx.Hash())
	}
}

func TestTradingFunc(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	sdk, err := NewSdk([]string{pk}, "testnet")
	if err != nil {
		t.Log(err.Error())
	}

	orders, ids, err := sdk.QueryOpenOrders("BTC-USDC-USDC", sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("order ids =", ids)
		fmt.Println("orders =", orders)
	}
	if len(ids) > 0 {
		tx, err := sdk.CancelOrder("BTC-USDC-USDC", ids[0], nil)
		if err != nil {
			t.Log(err.Error())
		} else {
			fmt.Println("tx cancel order=", tx.Hash())
		}
	}
	tx, err := sdk.AddCollateral("ETH-USD-MATIC", 100, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("tx hash adding collateral=", tx.Hash())
	}

	tx, err = sdk.ApproveTknSpending("ETH-USD-MATIC", nil, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("tx hash=", tx.Hash())
	}

	order, err := sdk.NewOrder("BTC-USDC-USDC", SIDE_SELL, ORDER_TYPE_LIMIT, 0.1, 10, &OrderOptions{LimitPrice: 2240})
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	orderId, _, err := sdk.PostOrder(order, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("order id =", orderId)
	}
	tx, err = sdk.CancelOrder("BTC-USDC-USDC", orderId, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("tx cancel order=", tx.Hash())
	}

	status, err := sdk.QueryOrderStatus("ETH-USD-MATIC", sdk.Wallets[0].Address, orderId, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("order status =", status)
	}
	pr, err := sdk.GetPositionRisk("ETH-USD-MATIC", sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("position risk =", pr)
	}
}

func TestABI(t *testing.T) {
	types := []string{"uint256", "address", "int128", "bytes32"}
	// domainBuf := []byte("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	domainHash := Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	// v := "0x" + hex.EncodeToString(hash)
	// var hashArray [32]byte
	// copy(hashArray[:], hash)
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
	resExpected := "0x000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000009d5aab428e98678d0e645ea4aebd25f744341a05fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb458cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866"
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
		Symbol:              "BTC-USD-MATIC",
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
		panic(err)
	}
	fmt.Println(dgst)
	if dgst != "625da4b5a8ccc7a093e88b5908b3019b3272c30058bc17d420e1b0e2a60ce0f5" {
		panic("wrong dgst result")
	}
}

func TestPostOrder(t *testing.T) {
	_, execPk, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}
	chConfig, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	pxConf, err := config.GetDefaultPriceConfig(chConfig.ChainId)
	if err != nil {
		t.Log(err.Error())
	}
	conn, _ := CreateBlockChainConnector(pxConf, chConfig, nil)
	wallet, err := NewWallet(fmt.Sprintf("%x", execPk.D), conn.ChainId, conn.Rpc)
	if err != nil {
		t.Log(err.Error())
	}
	var xInfo StaticExchangeInfo
	xInfo.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
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
		ParentChildOrderId1: emptyArray,
		ParentChildOrderId2: emptyArray,
	}
	_, txHash, _ := RawPostOrder(conn.Rpc, int(conn.ChainId), &xInfo, wallet, []byte{}, &order, traderAddr)
	fmt.Println("Tx hash = ", txHash)
}

func TestBrokerSignature(t *testing.T) {
	_, execPk, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}
	chConfig, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	var xInfo StaticExchangeInfo
	xInfo.Load("./tmpXchInfo.json")
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	wallet, err := NewWallet(fmt.Sprintf("%x", execPk.D), chConfig.ChainId, nil)
	if err != nil {
		panic("error creating wallet")
	}
	const brokerFeeTbps = 110
	dgst, sig, _ := RawCreateOrderBrokerSignature(xInfo.ProxyAddr, 80001, wallet, 10001, brokerFeeTbps, traderAddr.String(), 1684863656)
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
	chConfig, err := config.GetDefaultChainConfig("testnet")
	if err != nil {
		t.Log(err.Error())
	}
	wallet, err := NewWallet(fmt.Sprintf("%x", execPk.D), chConfig.ChainId, nil)
	fmt.Printf("\nwallet addr %s\n", wallet.Address.String())
	if err != nil {
		panic("error creating wallet")
	}
	var totalAmount big.Int
	totalAmount.SetString("1000000000000000000000", 10)

	executorKey, err := crypto.GenerateKey()
	if err != nil {
		t.Log(err.Error())
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
		t.Log(err.Error())
	}
	fmt.Print(dgst, sig)
	sigB, err := BytesFromHexString(sig)
	if err != nil {
		t.Log(err.Error())
	}
	addr, err := RecoverPaymentSignatureAddr(sigB, &ps)
	if err != nil {
		t.Log(err.Error())
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
		t.Log(err.Error())
	}
	// Derive the Ethereum address from the private key
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	pk := fmt.Sprintf("%x", privateKey.D)
	wallet, err := NewWallet(pk, int64(chainId), nil)
	proxyAddr := "0xCdd7C9e07689d1B3D558A714fAa5Cc4B6bA654bD"
	fmt.Printf("\nwallet addr %s\n", wallet.Address.String())
	if err != nil {
		panic("error creating wallet")
	}

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
	fmt.Println("digest = ")
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
		log.Fatal(err)
		return common.Address{}, nil, err
	}
	// Derive the Ethereum address from the private key
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	return addr, privateKey, err
}
