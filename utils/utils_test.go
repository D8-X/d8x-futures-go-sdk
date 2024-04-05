package utils

import (
	"fmt"
	"math/big"
	"testing"
)

func TestABDKToDecN(t *testing.T) {
	x := new(big.Int)
	x.SetString(ONE_64x64, 10)
	v := ABDKToDecN(x, 5)
	vi := v.Int64()
	if vi != 100000 {
		t.Fail()
	}
	fmt.Println(vi)
}
