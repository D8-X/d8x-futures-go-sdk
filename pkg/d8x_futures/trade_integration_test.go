//go:build integration

package d8x_futures

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	rpc, err := ethclient.Dial("https://sepolia.base.org")
	if err != nil {
		t.Fatalf("ethclient.Dial: %v", err)
	}
	// https://sports.quantena.org/slots-info/84532
	perp := getActiveSymbol(t, &sdk.Info)
	t.Logf("wallet addr =%s", sdk.Wallets[0].Address.Hex())
	orderObj, err := sdk.QueryAllOpenOrders(perp, nil)
	if err != nil {
		t.Fatalf("QueryAllOpenOrders: %v", err)
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
		order, err := sdk.NewOrder(perp, SIDE_SELL, ORDER_TYPE_MARKET, 2, 5, &OrderOptions{LimitPrice: 0})
		if err != nil {
			t.Fatalf("NewOrder: %v", err)
		}
		orderId, tx, err := sdk.PostOrder(order, nil)
		if err != nil {
			t.Fatalf("PostOrder: %v", err)
		}
		t.Log("order id =", orderId)
		_, err = bind.WaitMined(context.Background(), rpc, tx)
		if err != nil {
			t.Fatalf("WaitMined: %v", err)
		}
		// 92d891d3ae6d8695c8732d67fff2b59d309496796e1f369f8d5e0b4ab2a17cd4
		mktOrderIds = append(mktOrderIds, orderId)

	}
	now := time.Now().Unix()
	payoutAddr := common.Address{} // common.HexToAddress("0x98DfAFF5126836E339493a6021FD5B92Bf005F0D")
	tx, err := sdk.ExecuteOrders(perp, []string{mktOrderIds[0]}, &OptsOverridesExec{TsMin: uint32(now), PayoutAddr: payoutAddr})
	if err != nil {
		t.Fatalf("ExecuteOrders: %v", err)
	}
	t.Log("tx = ", tx.Hash())
	t.Log("done")
}

func TestSdkLiquidatePosition(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "421614") // arbitrum sepolia
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	acc2, err := sdk.QueryLiquidatableAccountsInPool(1, nil)
	if err != nil {
		t.Fatalf("QueryLiquidatableAccountsInPool: %v", err)
	}
	t.Log(acc2)
	if len(acc2) == 0 {
		return
	}
	for _, el := range acc2 {
		for _, addr := range el.LiqAccounts {
			t.Logf("liquidating %s", addr.Hex())
			tx, err := sdk.LiquidatePosition(el.PerpId, []common.Address{addr}, nil, nil)
			if err != nil {
				t.Logf("error liquidating: %s", err.Error())
				continue
			}
			t.Logf("liquidated %d trader %s tx=%s", el.PerpId, addr.Hex(), tx.Hash())
		}
	}
}

func TestAddCollateral(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "195")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	t.Log("wallet = " + sdk.Wallets[0].Address.Hex())
	sym := getActiveSymbol(t, &sdk.Info)
	tx, err := sdk.AddCollateral(sym, -0.0001, nil)
	if err != nil {
		t.Fatalf("AddCollateral: %v", err)
	}
	t.Log("tx hash adding collateral=", tx.Hash())
}

func TestTradingFunc(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		t.Fatal("Provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "base_sepolia")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}

	orders, ids, err := sdk.QueryOpenOrders("BTC-USDC-USDC", sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("order ids =", ids)
		t.Log("orders =", orders)
	}
	if len(ids) > 0 {
		tx, err := sdk.CancelOrder("BTC-USDC-USDC", ids[0], nil)
		if err != nil {
			t.Log(err.Error())
		} else {
			t.Log("tx cancel order=", tx.Hash())
		}
	}
	tx, err := sdk.AddCollateral("ETH-USD-MATIC", 100, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("tx hash adding collateral=", tx.Hash())
	}

	tx, err = sdk.ApproveTknSpending("ETH-USD-MATIC", nil, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("tx hash=", tx.Hash())
	}

	order, err := sdk.NewOrder("BTC-USDC-USDC", SIDE_SELL, ORDER_TYPE_LIMIT, 0.1, 10, &OrderOptions{LimitPrice: 2240})
	if err != nil {
		t.Fatalf("NewOrder: %v", err)
	}
	orderId, _, err := sdk.PostOrder(order, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("order id =", orderId)
	}
	tx, err = sdk.CancelOrder("BTC-USDC-USDC", orderId, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("tx cancel order=", tx.Hash())
	}

	status, err := sdk.QueryOrderStatus("ETH-USD-MATIC", sdk.Wallets[0].Address, orderId, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("order status =", status)
	}
	pr, err := sdk.GetPositionRisk("ETH-USD-MATIC", sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log("position risk =", pr)
	}
}

func TestPostOrder(t *testing.T) {
	_, execPk, err := generateKey()
	if err != nil {
		t.Fatalf("generateKey: %v", err)
	}
	chConfig, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	pxConf, err := config.GetDefaultPriceConfig(chConfig.ChainId)
	if err != nil {
		t.Fatalf("GetDefaultPriceConfig: %v", err)
	}
	conn, err := CreateBlockChainConnector(pxConf, chConfig, nil)
	if err != nil {
		t.Fatalf("CreateBlockChainConnector: %v", err)
	}
	wallet, err := NewWallet(fmt.Sprintf("%x", execPk.D), conn.ChainId)
	if err != nil {
		t.Fatalf("NewWallet: %v", err)
	}
	var xInfo StaticExchangeInfo
	if err := xInfo.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
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
	t.Log("Tx hash = ", txHash)
}

func TestBrokerSignature(t *testing.T) {
	_, execPk, err := generateKey()
	if err != nil {
		t.Fatalf("generateKey: %v", err)
	}
	chConfig, err := config.GetDefaultChainConfig("base_sepolia")
	if err != nil {
		t.Fatalf("GetDefaultChainConfig: %v", err)
	}
	var xInfo StaticExchangeInfo
	if err := xInfo.Load("./tmpXchInfo.json"); err != nil {
		t.Fatalf("Load tmpXchInfo.json: %v", err)
	}
	traderAddr := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
	wallet, err := NewWallet(fmt.Sprintf("%x", execPk.D), chConfig.ChainId)
	if err != nil {
		t.Fatalf("NewWallet: %v", err)
	}
	const brokerFeeTbps = 110
	dgst, sig, _ := RawCreateOrderBrokerSignature(xInfo.ProxyAddr, 80001, wallet, 10001, brokerFeeTbps, traderAddr.String(), 1684863656)
	t.Log(dgst, sig)
}
