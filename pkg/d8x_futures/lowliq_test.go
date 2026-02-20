//go:build integration

package d8x_futures

import (
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

func TestGetLowLiqPrices(t *testing.T) {
	priceIds := []PriceId{
		{
			Id:     "0xb61e541b6dbc43bc3bb368ab3d19625a3a9d32bd567ba0f9f11a2f0a91e560fd",
			Origin: "univ2",
			Type:   utils.PXTYPE_V2,
		},
	}
	samples := []float64{1, 100, 250, 2500, 5000}
	m := float64(10)
	tradeAmt := make([]float64, len(samples)*2)
	for j := range samples {
		tradeAmt[len(samples)+j] = m * samples[j]
		tradeAmt[len(samples)-1-j] = -m * samples[j]
	}
	// onchain
	sdkRo, err := NewSdkRO("84532")
	if err != nil {
		t.Fatalf("NewSdkRO: %v", err)
	}
	sym := getLowLiqSymbol(t, &sdkRo.Info)
	px0, err := sdkRo.QueryPerpetualPrices(sym, tradeAmt, nil)
	if err != nil {
		t.Fatalf("QueryPerpetualPrices: %v", err)
	}
	// directly from odin
	pxEp := PriceFeedEndpoints{PriceFeedEndpoint: "https://hermes.pyth.network", LowLiqFeedEndpoint: "https://odin-lowliq.d8x.xyz"}
	p, err := fetchOraclePrices(priceIds, pxEp)
	if err != nil {
		t.Fatalf("fetchOraclePrices: %v", err)
	}
	px, err := GetLowLiqPrices(p[0], tradeAmt)
	if err != nil {
		t.Fatalf("GetLowLiqPrices: %v", err)
	}
	t.Log("quantity, price odin, price via onchain")
	for k := 0; k < len(tradeAmt); k++ {
		t.Logf("%f, %f, %f", tradeAmt[k], px[k], px0[k])
	}
}
