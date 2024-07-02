package d8x_futures

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestQueryBrokerLots(t *testing.T) {
	addr := common.HexToAddress("B0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
	var sdkRo SdkRO
	err := sdkRo.New("testnet")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	lots, err := sdkRo.QueryBrokerLots("USDC", &addr, nil)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("lots = %d\n", lots)
}

func TestPurchaseBrokerLots(t *testing.T) {
	var sdk Sdk
	pk := loadPk()
	if pk == "" {
		fmt.Println("provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	err := sdk.New([]string{pk}, "42161")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	amountFlt, _, err := sdk.Allowance("BTC-USD-STUSD", sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	var tx *types.Transaction
	if amountFlt < 100 {
		tx, err = sdk.ApproveTknSpending("BTC-USD-STUSD", nil, nil)
		if err != nil {
			t.Logf(err.Error())
			t.FailNow()
		}
		fmt.Printf("approved tkn spending, tx = %s\n", tx.Hash().Hex())
		time.Sleep(10 * time.Second)
		amountFlt, _, err = sdk.Allowance("BTC-USD-STUSD", sdk.Wallets[0].Address, nil)
		if err != nil {
			t.Logf(err.Error())
			t.FailNow()
		}
	}
	fmt.Printf("allowance %f = \n", amountFlt)
	tx, err = sdk.PurchaseBrokerLots(1, "BTC-USD-STUSD", nil)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("success, tx = %s\n", tx.Hash().Hex())
	time.Sleep(10 * time.Second)
	num, err := sdk.QueryBrokerLots("BTC-USD-STUSD", &sdk.Wallets[0].Address, nil)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("num lots = %d\n", num)
}
