package d8x_futures

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	STUSD              = "0x0022228a2cc5E7eF0274A7Baa600d44da5aB5776"
	USDC               = "0xaf88d065e77c8cC2239327C5EDb3A432268e5831"
	USDA               = "0x0000206329b97DB379d5E1Bf586BbDB969C63274"
	TRANSMUTER_ADDRESS = "0xD253b62108d1831aEd298Fc2434A5A8e4E418053"
	DECIMALS           = 6
)

var ierc4626ABI = `[{"constant":true,"inputs":[{"name":"_shares","type":"uint256"}],"name":"previewMint","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
var iTransmuterABI = `[{"constant":true,"inputs":[{"name":"_amountIn","type":"uint256"},{"name":"_targetToken","type":"address"},{"name":"_baseToken","type":"address"}],"name":"quoteOut","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

// STUSDToUSDC fetches the onchain conversion rate for STUSD-USDC
// on Arbitrum. Provider needs to be an Arbitrum RPC instance
func STUSDToUSDC(provider *ethclient.Client) (float64, error) {
	// Parse ABIs
	ierc4626, err := abi.JSON(strings.NewReader(ierc4626ABI))
	if err != nil {
		return 0, err
	}
	iTransmuter, err := abi.JSON(strings.NewReader(iTransmuterABI))
	if err != nil {
		return 0, err
	}

	// Create contract instances
	ierc4626Address := common.HexToAddress(STUSD)
	ierc4626Contract := bind.NewBoundContract(ierc4626Address, ierc4626, provider, provider, provider)

	iTransmuterAddress := common.HexToAddress(TRANSMUTER_ADDRESS)
	iTransmuterContract := bind.NewBoundContract(iTransmuterAddress, iTransmuter, provider, provider, provider)

	// Call previewMint to get amountUSDA
	oneSTUSD := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	resultUSDA := make([]interface{}, 0)
	err = ierc4626Contract.Call(nil, &resultUSDA, "previewMint", oneSTUSD)
	if err != nil {
		return 0, err
	}
	amountUSDA := resultUSDA[0].(*big.Int)
	// Call quoteOut to get amountIn
	resultUSDC := make([]interface{}, 0)
	err = iTransmuterContract.Call(nil, &resultUSDC, "quoteOut", amountUSDA, common.HexToAddress(USDC), common.HexToAddress(USDA))
	if err != nil {
		return 0, err
	}
	amountUSDC := resultUSDC[0].(*big.Int)
	// Convert amountUSDC to float
	amountUSDCFloat := new(big.Float).SetInt(amountUSDC)
	decimalsFloat := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(DECIMALS), nil))
	amountUSDCFloat.Quo(amountUSDCFloat, decimalsFloat)

	amountUSDCResult, _ := amountUSDCFloat.Float64()
	return amountUSDCResult, nil
}
