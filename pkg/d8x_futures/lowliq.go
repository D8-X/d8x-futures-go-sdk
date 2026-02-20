package d8x_futures

import (
	"fmt"
	"math"
	"strconv"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

const SIGNIFICANT_DIGITS = 3 // accuracy, e.g., 3.14159 -> 3.14 for value 3

// decode16 converts the float encoded in the 16bit number
// into its float64 representation
// |number_sign_bit|exponent_sign_bit|integer_number|exponent
// | 1 bit         | 1 bit           | 10 bits      | 4 bits
func decode16(num uint16) float64 {
	// extract sign bits
	neg := num&(1<<15) != 0
	negExp := num&(1<<14) != 0
	// extract number using mask 0011111111110000
	extrNum := int((num & 0x3FF0) >> 4)
	if neg {
		extrNum = -extrNum
	}
	// extract exponent
	extrExp := int(num & 0xF)
	if negExp {
		extrExp = -extrExp
	}
	val := float64(extrNum) * math.Pow10(-SIGNIFICANT_DIGITS+1)
	val = val * math.Pow10(extrExp)
	return val
}

// decodeOrderBook into a0, m0, a1, m1
func decodeOrderBook(enc uint64) (float64, float64, float64, float64) {
	// Extract the four 16-bit numbers
	m1 := uint16(enc & 0xFFFF)         // First 16 bits from right (0 to 15)
	a1 := uint16((enc >> 16) & 0xFFFF) // Second 16 bits (16 to 31)
	m0 := uint16((enc >> 32) & 0xFFFF) // Third 16 bits (32 to 47)
	a0 := uint16((enc >> 48) & 0xFFFF) // Fourth 16 bits (48 to 63)
	return decode16(a0), decode16(m0), decode16(a1), decode16(m1)
}

// GetLowLiqPrices uses the encoded price information to extract prices for the
// given sample trade-quantities (negative for short trades)
// This applies only to low-liq perpetuals
func GetLowLiqPrices(p ResponsePythLatestPriceFeed, tradeSizes []float64) ([]float64, error) {
	mid, halfBa, encBk, err := extractLowLiqParams(p)
	if err != nil {
		return nil, fmt.Errorf("unable to parse order-book encoding: %s", err.Error())
	}
	return calcPrices(halfBa, mid, encBk, tradeSizes), nil
}

func extractLowLiqParams(p ResponsePythLatestPriceFeed) (float64, float64, uint64, error) {
	mid, err := utils.PythNToFloat64(p.Price.Price, p.Price.Expo)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid mid price: %w", err)
	}
	hbaTbps, err := strconv.Atoi(p.Price.Conf)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("unable to extract half-ba: %s", err.Error())
	}
	halfBa := float64(hbaTbps) / 1e5
	encBk, err := strconv.ParseUint(p.EmaPrice.Conf, 10, 64)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("unable to parse order-book encoding: %s", err.Error())
	}
	return mid, halfBa, encBk, nil
}

// calcPrice uses the encoded order book and estimates the
// prices that the given trade sizes imply
func calcPrices(halfBA, mid float64, enc uint64, tradeSizes []float64) []float64 {
	a0, m0, a1, m1 := decodeOrderBook(enc)
	prices := make([]float64, 0, len(tradeSizes))
	for _, v := range tradeSizes {
		var px float64
		if v < 0 {
			// bid side
			px = mid - halfBA - math.Exp(a0+m0*math.Abs(v))
		} else {
			// ask side
			px = mid + halfBA + math.Exp(a1+m1*v)
		}
		prices = append(prices, px)
	}
	return prices
}
