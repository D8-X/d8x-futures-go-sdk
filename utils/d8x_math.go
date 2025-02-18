package utils

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// var ONE_64x64 = new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)
// var DECIMALS_18 = big.NewInt(1000000000000000000)
const (
	ONE_64x64   = "18446744073709551616" // 1 << 64
	DECIMALS_18 = 1e18
)

func Max64x64() *big.Int {
	MAX_64x64 := new(big.Int)
	MAX_64x64.Exp(big.NewInt(2), big.NewInt(127), nil)
	MAX_64x64.Sub(MAX_64x64, big.NewInt(1))
	return MAX_64x64
}

func MaxUint256() *big.Int {
	maxUint256 := new(big.Int)
	maxUint256.Exp(big.NewInt(2), big.NewInt(256), nil)
	maxUint256.Sub(maxUint256, big.NewInt(1))
	return maxUint256
}

func I32ToFloat64(x int32) float64 {
	return float64(x) / (math.Pow(2, 29))
}

func Sign(x float64) float64 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func DecNToFloat(decN *big.Int, n uint8) float64 {
	y := new(big.Float).SetInt(decN)
	N := new(big.Int).SetUint64(uint64(n))
	pow := new(big.Int).Exp(big.NewInt(10), N, nil)
	q := new(big.Float).SetInt(pow)
	z := new(big.Float).Quo(y, q)
	f, _ := z.Float64()
	return f
}

func ABDKToDecN(x *big.Int, n uint8) *big.Int {
	one64x64 := new(big.Int)
	one64x64.SetString(ONE_64x64, 10)
	N := new(big.Int).SetUint64(uint64(n))
	decN := new(big.Int).Exp(big.NewInt(10), N, nil)
	y := new(big.Int).Mul(x, decN)
	y = new(big.Int).Div(y, one64x64)
	return y
}

func Float64ToABDK(x float64) *big.Int {
	if x == 0 {
		return big.NewInt(0)
	}
	y := new(big.Float).SetFloat64(x)
	//fmt.Println(y.Float64())
	one64x64 := new(big.Int)
	one64x64.SetString(ONE_64x64, 10)
	y64 := new(big.Float).Mul(y, new(big.Float).SetInt(one64x64))
	// convert to Int
	intVal := new(big.Int)
	//fmt.Println(y64)
	y64.Int(intVal)
	//fmt.Println(intVal)
	return intVal
}

func ABDKToFloat64(xIn *big.Int) float64 {
	if xIn.Cmp(big.NewInt(0)) == 0 {
		return float64(0)
	}
	x := new(big.Int).Set(xIn)
	x.Abs(x)
	var sgn float64 = float64(x.Sign())
	x.Abs(x)
	var xInt, xRemainder, xDec big.Int
	one64x64 := new(big.Int)
	one64x64.SetString(ONE_64x64, 10)
	xInt.Div(x, one64x64)
	xRemainder.Mod(x, one64x64)
	xDec.Mul(&xRemainder, new(big.Int).SetInt64(DECIMALS_18))
	xDec.Div(&xDec, one64x64)
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

func Dec2Hex(num string) (string, error) {
	number := new(big.Int)
	number, ok := number.SetString(num, 10)
	if !ok {
		return "", fmt.Errorf("converting number to BIG int")
	}

	return "0x" + number.Text(16), nil
}

func Hex2Dec(num string) (string, error) {
	number := new(big.Int)
	num = strings.TrimPrefix(num, "0x")
	number, ok := number.SetString(num, 16)
	if !ok {
		return "", fmt.Errorf("converting number to BIG int")
	}

	return number.Text(10), nil
}
