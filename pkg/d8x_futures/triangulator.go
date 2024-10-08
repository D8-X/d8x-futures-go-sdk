package d8x_futures

import (
	"fmt"
	"strings"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

// calculateTriangulation calculates the triangulated price and reports whether any
// of the price-feeds is outdated, given a triangulation path and price data
func calculateTriangulation(triang Triangulation, outDateOffChain, outDateOnChain int64, symToPrice map[string]Price) (float64, bool, error) {
	var price float64 = 1.0
	timestampNow := time.Now().Unix()
	var isFeedOutdated = false
	// loop over triangulation
	for i, symCurr := range triang.Symbol {
		// find the current symbol symCurr of the triangulation in the
		// price data
		pxObs, exists := symToPrice[symCurr]
		if !exists {
			return 0, true, fmt.Errorf("symbol %s does not exist in config", symCurr)
		}

		// multiply or divide depending on isInverse
		if triang.IsInverse[i] {
			price = price / pxObs.Px
		} else {
			price = price * pxObs.Px
		}
		isFeedOutdated = isFeedOutdated || ((pxObs.IsOffChain && timestampNow-pxObs.Ts > outDateOffChain) ||
			(!pxObs.IsOffChain && timestampNow-pxObs.Ts > outDateOnChain))
	}
	return price, isFeedOutdated, nil
}

// Triangulate finds the shortest triangulation path for symbol (e.g. BTC-USDC) using
// all price sources in pxConfig. Returns an empty array if no triangulation found.
func Triangulate(symbol string, priceFeeds []utils.PriceFeedId) Triangulation {
	var feedSymBase []string
	var feedSymQuote []string
	symbol = strings.ToUpper(symbol)
	// extract all base and quote currencies
	for _, feed := range priceFeeds {
		syms := strings.Split(feed.Symbol, "-")
		feedSymBase = append(feedSymBase, strings.ToUpper(syms[0]))
		feedSymQuote = append(feedSymQuote, strings.ToUpper(syms[1]))
	}
	paths := findPath(append(feedSymBase, feedSymQuote...), append(feedSymQuote, feedSymBase...), symbol)
	if len(paths) == 0 {
		return Triangulation{}
	}
	// get shortest path
	var minPath int = 1000
	var minIdx int = -1
	for j := 0; j < len(paths); j++ {
		if len(paths[j]) < int(minPath) {
			minPath = len(paths[j])
			minIdx = j
		}
	}
	// make it a triangulation
	triang := Triangulation{
		IsInverse: make([]bool, len(paths[minIdx])),
		Symbol:    make([]string, len(paths[minIdx])),
	}

	for j := 0; j < len(paths[minIdx]); j++ {
		currSym := strings.Split(paths[minIdx][j], "-")
		// is this symbol in the original config?
		invert := true
		for k := 0; k < len(feedSymBase); k++ {
			if feedSymBase[k] == currSym[0] && feedSymQuote[k] == currSym[1] {
				// found
				invert = false
				break
			}
		}
		if invert {
			triang.IsInverse[j] = true
			triang.Symbol[j] = currSym[1] + "-" + currSym[0]
		} else {
			triang.IsInverse[j] = false
			triang.Symbol[j] = currSym[0] + "-" + currSym[1]
		}
	}
	return triang
}

// findPath finds all possible non-circular paths to triangulate 'pair' (e.g. ETH-USDC).
// The input requires an array of all base currencies and an array of all quote correncies
// in which each base-quote pair occurs twice (e.g., ccyBase=[BTC, USD], ccyQuote=[USD, BTC])
func findPath(ccyBase []string, ccyQuote []string, pair string) [][]string {
	syms := strings.Split(pair, "-")
	base, quote := syms[0], syms[1]
	var paths [][]string
	for k := 0; k < len(ccyBase); k++ {
		if ccyBase[k] == base {
			currentPath := []string{ccyBase[k] + "-" + ccyQuote[k]}
			if ccyQuote[k] == quote {
				// done
				paths = append(paths, currentPath)
			} else {
				// find path for ccyQuote[k]-quote without the currency
				var newBase []string
				var newQuote []string
				for j := 0; j < len(ccyBase); j++ {
					if j != k && ccyQuote[j]+"-"+ccyBase[j] != ccyBase[k]+"-"+ccyQuote[k] {
						newBase = append(newBase, ccyBase[j])
						newQuote = append(newQuote, ccyQuote[j])
					}
				}
				recursivePath := findPath(newBase, newQuote, ccyQuote[k]+"-"+quote)
				for j := 0; j < len(recursivePath); j++ {
					paths = append(paths, append(currentPath, (recursivePath)[j]...))
				}
			}
		}
	}
	return paths
}
