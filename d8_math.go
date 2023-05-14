package d8x_futures

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

var ONE_64x64 = new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)
var DECIMALS_18 = big.NewInt(1000000000000000000)

func I32ToFloat64(x int32) float64 {
	return float64(x) / (math.Pow(2, 29))
}

func sign(x float64) float64 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func Float64ToABDK(x float64) *big.Int {
	if x == 0 {
		return big.NewInt(0)
	}
	y := new(big.Float).SetFloat64(x)
	y.Mul(y, new(big.Float).SetInt(ONE_64x64))
	// convert to Int
	intVal := new(big.Int)
	y.Int(intVal)
	return intVal
}

func ABDKToFloat64(x *big.Int) float64 {
	if x.Cmp(big.NewInt(0)) == 0 {
		return float64(0)
	}
	var sgn float64 = float64(x.Sign())
	x.Abs(x)
	var xInt, xRemainder, xDec big.Int
	xInt.Div(x, ONE_64x64)
	xRemainder.Mod(x, ONE_64x64)
	xDec.Mul(&xRemainder, DECIMALS_18)
	xDec.Div(&xDec, ONE_64x64)
	digits := 18 - len(xDec.String()) // 18 - number of digits
	pad := strings.Repeat("0", digits)
	numberStr := xInt.String() + "." + pad + xDec.String()
	res, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		panic(err)
	}
	return res * sgn
}

func PythNToFloat64(price string, expo int32) float64 {
	var x big.Int
	x.SetString(price, 10)
	var dec = new(big.Int).Exp(big.NewInt(10), big.NewInt(-int64(expo)), nil)
	var sgn float64 = float64(x.Sign())
	x.Abs(&x)
	var xInt, xRemainder big.Int
	xInt.Div(&x, dec)
	xRemainder.Mod(&x, dec)
	digits := -int(expo) - len(xRemainder.String()) // digits - number of digits used
	pad := strings.Repeat("0", digits)
	numberStr := xInt.String() + "." + pad + xRemainder.String()
	res, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		panic(err)
	}
	return res * sgn
}
