package d8x_futures

import (
	"fmt"
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

func ABDKToFloat64(x *big.Int) float64 {
	if x.Cmp(big.NewInt(0)) == 0 {
		return float64(0)
	}
	var sgn float64 = float64(x.Sign())
	x.Abs(x)
	var xInt, xRemainder, xDec big.Int
	xInt.Div(x, ONE_64x64)
	fmt.Println(x.String())
	fmt.Println(xInt.String())
	fmt.Println(ONE_64x64.String())
	xRemainder.Mod(x, ONE_64x64)
	xDec.Mul(&xRemainder, DECIMALS_18)
	xDec.Div(&xDec, ONE_64x64)
	digits := 18 - (xDec.BitLen()+3)/4 // 18 - number of digits
	pad := strings.Repeat("0", digits)
	numberStr := xInt.String() + "." + pad + xDec.String()
	res, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		panic(err)
	}
	return res * sgn
}
