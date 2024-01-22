package d8x_futures

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
	pk := os.Getenv("PK")
	if pk == "" {
		fmt.Println("Provide private key for testnet as environment variable PK")
		t.FailNow()
	}
	err := sdk.New([]string{pk}, "testnet")
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	tx, err := sdk.ApproveTknSpending("USDC", nil, nil)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("approved tkn spending, tx = %s\n", tx.Hash().Hex())
	tx, err = sdk.PurchaseBrokerLots(1, "USDC", nil)
	if err != nil {
		t.Logf(err.Error())
		t.FailNow()
	}
	fmt.Printf("success, tx = %s\n", tx.Hash().Hex())
}
