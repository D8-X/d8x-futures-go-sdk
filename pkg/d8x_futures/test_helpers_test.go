package d8x_futures

import (
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

// getActiveSymbol returns the first NORMAL-state perpetual symbol from the exchange info.
func getActiveSymbol(t *testing.T, info *StaticExchangeInfo) string {
	t.Helper()
	for _, p := range info.Perpetuals {
		if p.State() == NORMAL {
			sym, ok := info.PerpetualIdToSymbol[p.Id]
			if ok {
				return sym
			}
		}
	}
	t.Fatal("no active perpetual found")
	return ""
}

// getLowLiqSymbol returns the first NORMAL-state perpetual symbol that uses a V2-type price feed,
// or skips the test if none is found.
func getLowLiqSymbol(t *testing.T, info *StaticExchangeInfo) string {
	t.Helper()
	for _, p := range info.Perpetuals {
		if p.State() != NORMAL {
			continue
		}
		for _, pid := range p.PriceIds {
			if pid.Type == utils.PXTYPE_V2 || pid.Type == utils.PXTYPE_V3 {
				sym, ok := info.PerpetualIdToSymbol[p.Id]
				if ok {
					return sym
				}
			}
		}
	}
	t.Skip("no low-liquidity perpetual found")
	return ""
}
