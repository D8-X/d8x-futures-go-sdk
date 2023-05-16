package d8x_futures

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestABI(t *testing.T) {
	types := []string{"uint256", "address", "int128", "bytes32"}
	//domainBuf := []byte("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	var domainHash Keccak256Hash
	domainHash.Keccak256FromString("EIP712Domain(string name,uint256 chainId,address verifyingContract)")
	//v := "0x" + hex.EncodeToString(hash)
	//var hashArray [32]byte
	//copy(hashArray[:], hash)
	values := []interface{}{
		big.NewInt(123),
		common.HexToAddress("***REMOVED***"),
		big.NewInt(-1211),
		domainHash,
	}

	result, err := AbiEncode(types, values...)
	if err != nil {
		fmt.Println("Encoding error:", err)
		return
	}

	fmt.Println("Encoded data:", result)
	var resExpected = "0x000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000009d5aab428e98678d0e645ea4aebd25f744341a05fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb458cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866"
	if resExpected != result {
		panic("wrong encoding")
	}
}
