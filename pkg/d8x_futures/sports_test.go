//go:build integration

package d8x_futures

import (
	"strings"
	"testing"
)

func TestInternalToSymbol(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	c, err := sdk.internalToSymbol("SP00-USD")
	if err != nil {
		t.Fatalf("internalToSymbol SP00-USD: %v", err)
	}
	t.Log(c)
	c, err = sdk.internalToSymbol("BTC-USD")
	if err != nil {
		t.Fatalf("internalToSymbol BTC-USD: %v", err)
	}
	t.Log(c)
}

func TestSportSlotAssignment(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	// https://sports.quantena.org/slots-info/84532
	c, exists := sdk.SportSlotAssignment("XP00")
	t.Log(exists, c)
	c, exists = sdk.SportSlotAssignment("NHL0")
	t.Log(exists, c)
}

func TestSportSlot(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}

	extSym, err := sdk.internalToSymbol("SP00-USD")
	if err != nil {
		t.Skip("no active sport event for SP00: " + err.Error())
	}
	gameName := strings.Split(extSym, "-")[0]
	c, exists := sdk.SportSlot(gameName)
	t.Log(exists, c)
}

func TestSymbolToInternal(t *testing.T) {
	sdk, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}

	extSym, err := sdk.internalToSymbol("SP00-USD")
	if err != nil {
		t.Skip("no active sport event for SP00: " + err.Error())
	}
	c, err := sdk.symbolToInternal(extSym)
	if err != nil {
		t.Fatalf("symbolToInternal: %v", err)
	}
	t.Log(c)
}
