package d8x_futures

import (
	"os"
	"testing"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

func skipIfCI(t *testing.T) {
	t.Helper()
	if os.Getenv("CI") != "" {
		t.Skip("skipping write test in CI")
	}
}

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

func getAnySymbol(t *testing.T, info *StaticExchangeInfo) string {
	t.Helper()
	for _, p := range info.Perpetuals {
		sym, ok := info.PerpetualIdToSymbol[p.Id]
		if ok {
			return sym
		}
	}
	t.Skip("no perpetual found in exchange info")
	return ""
}


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
