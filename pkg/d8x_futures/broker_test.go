//go:build integration

package d8x_futures

import (
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestQueryBrokerLots(t *testing.T) {
	addr := common.HexToAddress("B0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
	sdkRo, err := NewSdkRO("base_sepolia")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	// Get pool symbol dynamically from first active perpetual
	sym := getActiveSymbol(t, &sdkRo.Info)
	parts := strings.Split(sym, "-")
	pool := parts[len(parts)-1] // settlement currency = pool name
	lots, err := sdkRo.QueryBrokerLots(pool, &addr, nil)
	if err != nil {
		t.Fatalf("QueryBrokerLots: %v", err)
	}
	t.Logf("lots = %d", lots)
}

func TestPurchaseBrokerLots(t *testing.T) {
	pk := loadPk()
	if pk == "" {
		t.Fatal("provide private key for testnet as environment variable PK")
	}
	sdk, err := NewSdk([]string{pk}, "84532")
	if err != nil {
		t.Fatalf("NewSdk: %v", err)
	}
	sym := getActiveSymbol(t, &sdk.Info)
	amountFlt, _, err := sdk.Allowance(sym, sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Fatalf("Allowance: %v", err)
	}
	var tx *types.Transaction
	if amountFlt < 100 {
		tx, err = sdk.ApproveTknSpending(sym, nil, nil)
		if err != nil {
			t.Fatalf("ApproveTknSpending: %v", err)
		}
		t.Logf("approved tkn spending, tx = %s", tx.Hash().Hex())
		time.Sleep(10 * time.Second)
		amountFlt, _, err = sdk.Allowance(sym, sdk.Wallets[0].Address, nil)
		if err != nil {
			t.Fatalf("Allowance after approve: %v", err)
		}
	}
	t.Logf("allowance %f = ", amountFlt)
	tx, err = sdk.PurchaseBrokerLots(1, sym, nil)
	if err != nil {
		t.Fatalf("PurchaseBrokerLots: %v", err)
	}
	t.Logf("success, tx = %s", tx.Hash().Hex())
	time.Sleep(10 * time.Second)
	num, err := sdk.QueryBrokerLots(sym, &sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Fatalf("QueryBrokerLots: %v", err)
	}
	t.Logf("num lots = %d", num)
}
