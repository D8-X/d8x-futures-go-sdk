package d8x_futures

import (
	"fmt"
	"testing"
)

func TestInternalToSymbol(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	c, err := sdk.internalToSymbol("NHL0-USD")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(c)
	c, err = sdk.internalToSymbol("BTC-USD")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(c)
}

func TestSymbolToInternal(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	c, err := sdk.symbolToInternal("NHL_FLA_DET_251015-USD")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(c)
}
