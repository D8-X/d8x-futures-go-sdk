// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CompositeTokenMetaData contains all meta data concerning the CompositeToken contract.
var CompositeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InvalidAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"InvalidFeed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"ControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"name\":\"TokenSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CMP_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_feeds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"effectiveBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_feeds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"registerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedToken\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"feedDecimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalInvested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// CompositeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CompositeTokenMetaData.ABI instead.
var CompositeTokenABI = CompositeTokenMetaData.ABI

// CompositeToken is an auto generated Go binding around an Ethereum contract.
type CompositeToken struct {
	CompositeTokenCaller     // Read-only binding to the contract
	CompositeTokenTransactor // Write-only binding to the contract
	CompositeTokenFilterer   // Log filterer for contract events
}

// CompositeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompositeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompositeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompositeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompositeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompositeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompositeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompositeTokenSession struct {
	Contract     *CompositeToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompositeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompositeTokenCallerSession struct {
	Contract *CompositeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CompositeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompositeTokenTransactorSession struct {
	Contract     *CompositeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CompositeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompositeTokenRaw struct {
	Contract *CompositeToken // Generic contract binding to access the raw methods on
}

// CompositeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompositeTokenCallerRaw struct {
	Contract *CompositeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CompositeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompositeTokenTransactorRaw struct {
	Contract *CompositeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompositeToken creates a new instance of CompositeToken, bound to a specific deployed contract.
func NewCompositeToken(address common.Address, backend bind.ContractBackend) (*CompositeToken, error) {
	contract, err := bindCompositeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompositeToken{CompositeTokenCaller: CompositeTokenCaller{contract: contract}, CompositeTokenTransactor: CompositeTokenTransactor{contract: contract}, CompositeTokenFilterer: CompositeTokenFilterer{contract: contract}}, nil
}

// NewCompositeTokenCaller creates a new read-only instance of CompositeToken, bound to a specific deployed contract.
func NewCompositeTokenCaller(address common.Address, caller bind.ContractCaller) (*CompositeTokenCaller, error) {
	contract, err := bindCompositeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenCaller{contract: contract}, nil
}

// NewCompositeTokenTransactor creates a new write-only instance of CompositeToken, bound to a specific deployed contract.
func NewCompositeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CompositeTokenTransactor, error) {
	contract, err := bindCompositeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenTransactor{contract: contract}, nil
}

// NewCompositeTokenFilterer creates a new log filterer instance of CompositeToken, bound to a specific deployed contract.
func NewCompositeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CompositeTokenFilterer, error) {
	contract, err := bindCompositeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenFilterer{contract: contract}, nil
}

// bindCompositeToken binds a generic wrapper to an already deployed contract.
func bindCompositeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompositeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompositeToken *CompositeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompositeToken.Contract.CompositeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompositeToken *CompositeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompositeToken.Contract.CompositeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompositeToken *CompositeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompositeToken.Contract.CompositeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompositeToken *CompositeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompositeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompositeToken *CompositeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompositeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompositeToken *CompositeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompositeToken.Contract.contract.Transact(opts, method, params...)
}

// CMPDECIMALS is a free data retrieval call binding the contract method 0x6506504c.
//
// Solidity: function CMP_DECIMALS() view returns(uint8)
func (_CompositeToken *CompositeTokenCaller) CMPDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "CMP_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CMPDECIMALS is a free data retrieval call binding the contract method 0x6506504c.
//
// Solidity: function CMP_DECIMALS() view returns(uint8)
func (_CompositeToken *CompositeTokenSession) CMPDECIMALS() (uint8, error) {
	return _CompositeToken.Contract.CMPDECIMALS(&_CompositeToken.CallOpts)
}

// CMPDECIMALS is a free data retrieval call binding the contract method 0x6506504c.
//
// Solidity: function CMP_DECIMALS() view returns(uint8)
func (_CompositeToken *CompositeTokenCallerSession) CMPDECIMALS() (uint8, error) {
	return _CompositeToken.Contract.CMPDECIMALS(&_CompositeToken.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CompositeToken *CompositeTokenCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CompositeToken *CompositeTokenSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CompositeToken.Contract.UPGRADEINTERFACEVERSION(&_CompositeToken.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CompositeToken *CompositeTokenCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CompositeToken.Contract.UPGRADEINTERFACEVERSION(&_CompositeToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_CompositeToken *CompositeTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_CompositeToken *CompositeTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.Allowance(&_CompositeToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_CompositeToken *CompositeTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.Allowance(&_CompositeToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CompositeToken *CompositeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CompositeToken *CompositeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.BalanceOf(&_CompositeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CompositeToken *CompositeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.BalanceOf(&_CompositeToken.CallOpts, account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CompositeToken *CompositeTokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CompositeToken *CompositeTokenSession) Controller() (common.Address, error) {
	return _CompositeToken.Contract.Controller(&_CompositeToken.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CompositeToken *CompositeTokenCallerSession) Controller() (common.Address, error) {
	return _CompositeToken.Contract.Controller(&_CompositeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CompositeToken *CompositeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CompositeToken *CompositeTokenSession) Decimals() (uint8, error) {
	return _CompositeToken.Contract.Decimals(&_CompositeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_CompositeToken *CompositeTokenCallerSession) Decimals() (uint8, error) {
	return _CompositeToken.Contract.Decimals(&_CompositeToken.CallOpts)
}

// EffectiveBalanceOf is a free data retrieval call binding the contract method 0xc7a64723.
//
// Solidity: function effectiveBalanceOf(address _account) view returns(uint256 balance)
func (_CompositeToken *CompositeTokenCaller) EffectiveBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "effectiveBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EffectiveBalanceOf is a free data retrieval call binding the contract method 0xc7a64723.
//
// Solidity: function effectiveBalanceOf(address _account) view returns(uint256 balance)
func (_CompositeToken *CompositeTokenSession) EffectiveBalanceOf(_account common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.EffectiveBalanceOf(&_CompositeToken.CallOpts, _account)
}

// EffectiveBalanceOf is a free data retrieval call binding the contract method 0xc7a64723.
//
// Solidity: function effectiveBalanceOf(address _account) view returns(uint256 balance)
func (_CompositeToken *CompositeTokenCallerSession) EffectiveBalanceOf(_account common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.EffectiveBalanceOf(&_CompositeToken.CallOpts, _account)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_CompositeToken *CompositeTokenCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_CompositeToken *CompositeTokenSession) GetSupportedTokens() ([]common.Address, error) {
	return _CompositeToken.Contract.GetSupportedTokens(&_CompositeToken.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_CompositeToken *CompositeTokenCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _CompositeToken.Contract.GetSupportedTokens(&_CompositeToken.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_CompositeToken *CompositeTokenCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_CompositeToken *CompositeTokenSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _CompositeToken.Contract.LatestRoundData(&_CompositeToken.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_CompositeToken *CompositeTokenCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _CompositeToken.Contract.LatestRoundData(&_CompositeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CompositeToken *CompositeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CompositeToken *CompositeTokenSession) Name() (string, error) {
	return _CompositeToken.Contract.Name(&_CompositeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CompositeToken *CompositeTokenCallerSession) Name() (string, error) {
	return _CompositeToken.Contract.Name(&_CompositeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompositeToken *CompositeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompositeToken *CompositeTokenSession) Owner() (common.Address, error) {
	return _CompositeToken.Contract.Owner(&_CompositeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CompositeToken *CompositeTokenCallerSession) Owner() (common.Address, error) {
	return _CompositeToken.Contract.Owner(&_CompositeToken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CompositeToken *CompositeTokenCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CompositeToken *CompositeTokenSession) ProxiableUUID() ([32]byte, error) {
	return _CompositeToken.Contract.ProxiableUUID(&_CompositeToken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CompositeToken *CompositeTokenCallerSession) ProxiableUUID() ([32]byte, error) {
	return _CompositeToken.Contract.ProxiableUUID(&_CompositeToken.CallOpts)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x3c89c490.
//
// Solidity: function registeredToken(address ) view returns(address)
func (_CompositeToken *CompositeTokenCaller) RegisteredToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "registeredToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredToken is a free data retrieval call binding the contract method 0x3c89c490.
//
// Solidity: function registeredToken(address ) view returns(address)
func (_CompositeToken *CompositeTokenSession) RegisteredToken(arg0 common.Address) (common.Address, error) {
	return _CompositeToken.Contract.RegisteredToken(&_CompositeToken.CallOpts, arg0)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x3c89c490.
//
// Solidity: function registeredToken(address ) view returns(address)
func (_CompositeToken *CompositeTokenCallerSession) RegisteredToken(arg0 common.Address) (common.Address, error) {
	return _CompositeToken.Contract.RegisteredToken(&_CompositeToken.CallOpts, arg0)
}

// SupportedToken is a free data retrieval call binding the contract method 0xe2c6e438.
//
// Solidity: function supportedToken(address ) view returns(bytes32 id, address feed, uint8 tokenDecimals, uint8 feedDecimals)
func (_CompositeToken *CompositeTokenCaller) SupportedToken(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id            [32]byte
	Feed          common.Address
	TokenDecimals uint8
	FeedDecimals  uint8
}, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "supportedToken", arg0)

	outstruct := new(struct {
		Id            [32]byte
		Feed          common.Address
		TokenDecimals uint8
		FeedDecimals  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Feed = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenDecimals = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.FeedDecimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// SupportedToken is a free data retrieval call binding the contract method 0xe2c6e438.
//
// Solidity: function supportedToken(address ) view returns(bytes32 id, address feed, uint8 tokenDecimals, uint8 feedDecimals)
func (_CompositeToken *CompositeTokenSession) SupportedToken(arg0 common.Address) (struct {
	Id            [32]byte
	Feed          common.Address
	TokenDecimals uint8
	FeedDecimals  uint8
}, error) {
	return _CompositeToken.Contract.SupportedToken(&_CompositeToken.CallOpts, arg0)
}

// SupportedToken is a free data retrieval call binding the contract method 0xe2c6e438.
//
// Solidity: function supportedToken(address ) view returns(bytes32 id, address feed, uint8 tokenDecimals, uint8 feedDecimals)
func (_CompositeToken *CompositeTokenCallerSession) SupportedToken(arg0 common.Address) (struct {
	Id            [32]byte
	Feed          common.Address
	TokenDecimals uint8
	FeedDecimals  uint8
}, error) {
	return _CompositeToken.Contract.SupportedToken(&_CompositeToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CompositeToken *CompositeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CompositeToken *CompositeTokenSession) Symbol() (string, error) {
	return _CompositeToken.Contract.Symbol(&_CompositeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CompositeToken *CompositeTokenCallerSession) Symbol() (string, error) {
	return _CompositeToken.Contract.Symbol(&_CompositeToken.CallOpts)
}

// TotalInvested is a free data retrieval call binding the contract method 0x90838e09.
//
// Solidity: function totalInvested(address ) view returns(uint256)
func (_CompositeToken *CompositeTokenCaller) TotalInvested(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "totalInvested", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInvested is a free data retrieval call binding the contract method 0x90838e09.
//
// Solidity: function totalInvested(address ) view returns(uint256)
func (_CompositeToken *CompositeTokenSession) TotalInvested(arg0 common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.TotalInvested(&_CompositeToken.CallOpts, arg0)
}

// TotalInvested is a free data retrieval call binding the contract method 0x90838e09.
//
// Solidity: function totalInvested(address ) view returns(uint256)
func (_CompositeToken *CompositeTokenCallerSession) TotalInvested(arg0 common.Address) (*big.Int, error) {
	return _CompositeToken.Contract.TotalInvested(&_CompositeToken.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CompositeToken *CompositeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CompositeToken *CompositeTokenSession) TotalSupply() (*big.Int, error) {
	return _CompositeToken.Contract.TotalSupply(&_CompositeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CompositeToken *CompositeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CompositeToken.Contract.TotalSupply(&_CompositeToken.CallOpts)
}

// TotalValue is a free data retrieval call binding the contract method 0xd4c3eea0.
//
// Solidity: function totalValue() view returns(uint256 value)
func (_CompositeToken *CompositeTokenCaller) TotalValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompositeToken.contract.Call(opts, &out, "totalValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValue is a free data retrieval call binding the contract method 0xd4c3eea0.
//
// Solidity: function totalValue() view returns(uint256 value)
func (_CompositeToken *CompositeTokenSession) TotalValue() (*big.Int, error) {
	return _CompositeToken.Contract.TotalValue(&_CompositeToken.CallOpts)
}

// TotalValue is a free data retrieval call binding the contract method 0xd4c3eea0.
//
// Solidity: function totalValue() view returns(uint256 value)
func (_CompositeToken *CompositeTokenCallerSession) TotalValue() (*big.Int, error) {
	return _CompositeToken.Contract.TotalValue(&_CompositeToken.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0xc3bcb4d6.
//
// Solidity: function addTokens(address[] _tokens, address[] _feeds, bytes32[] _ids) returns()
func (_CompositeToken *CompositeTokenTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address, _feeds []common.Address, _ids [][32]byte) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "addTokens", _tokens, _feeds, _ids)
}

// AddTokens is a paid mutator transaction binding the contract method 0xc3bcb4d6.
//
// Solidity: function addTokens(address[] _tokens, address[] _feeds, bytes32[] _ids) returns()
func (_CompositeToken *CompositeTokenSession) AddTokens(_tokens []common.Address, _feeds []common.Address, _ids [][32]byte) (*types.Transaction, error) {
	return _CompositeToken.Contract.AddTokens(&_CompositeToken.TransactOpts, _tokens, _feeds, _ids)
}

// AddTokens is a paid mutator transaction binding the contract method 0xc3bcb4d6.
//
// Solidity: function addTokens(address[] _tokens, address[] _feeds, bytes32[] _ids) returns()
func (_CompositeToken *CompositeTokenTransactorSession) AddTokens(_tokens []common.Address, _feeds []common.Address, _ids [][32]byte) (*types.Transaction, error) {
	return _CompositeToken.Contract.AddTokens(&_CompositeToken.TransactOpts, _tokens, _feeds, _ids)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.Contract.Approve(&_CompositeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.Contract.Approve(&_CompositeToken.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0x9b188519.
//
// Solidity: function initialize(address _owner, address[] _tokens, address[] _feeds, bytes32[] _ids) returns()
func (_CompositeToken *CompositeTokenTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _tokens []common.Address, _feeds []common.Address, _ids [][32]byte) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "initialize", _owner, _tokens, _feeds, _ids)
}

// Initialize is a paid mutator transaction binding the contract method 0x9b188519.
//
// Solidity: function initialize(address _owner, address[] _tokens, address[] _feeds, bytes32[] _ids) returns()
func (_CompositeToken *CompositeTokenSession) Initialize(_owner common.Address, _tokens []common.Address, _feeds []common.Address, _ids [][32]byte) (*types.Transaction, error) {
	return _CompositeToken.Contract.Initialize(&_CompositeToken.TransactOpts, _owner, _tokens, _feeds, _ids)
}

// Initialize is a paid mutator transaction binding the contract method 0x9b188519.
//
// Solidity: function initialize(address _owner, address[] _tokens, address[] _feeds, bytes32[] _ids) returns()
func (_CompositeToken *CompositeTokenTransactorSession) Initialize(_owner common.Address, _tokens []common.Address, _feeds []common.Address, _ids [][32]byte) (*types.Transaction, error) {
	return _CompositeToken.Contract.Initialize(&_CompositeToken.TransactOpts, _owner, _tokens, _feeds, _ids)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x97c414df.
//
// Solidity: function registerAccount(address _token) returns()
func (_CompositeToken *CompositeTokenTransactor) RegisterAccount(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "registerAccount", _token)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x97c414df.
//
// Solidity: function registerAccount(address _token) returns()
func (_CompositeToken *CompositeTokenSession) RegisterAccount(_token common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.RegisterAccount(&_CompositeToken.TransactOpts, _token)
}

// RegisterAccount is a paid mutator transaction binding the contract method 0x97c414df.
//
// Solidity: function registerAccount(address _token) returns()
func (_CompositeToken *CompositeTokenTransactorSession) RegisterAccount(_token common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.RegisterAccount(&_CompositeToken.TransactOpts, _token)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_CompositeToken *CompositeTokenTransactor) RemoveTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "removeTokens", _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_CompositeToken *CompositeTokenSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.RemoveTokens(&_CompositeToken.TransactOpts, _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_CompositeToken *CompositeTokenTransactorSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.RemoveTokens(&_CompositeToken.TransactOpts, _tokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CompositeToken *CompositeTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CompositeToken *CompositeTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _CompositeToken.Contract.RenounceOwnership(&_CompositeToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CompositeToken *CompositeTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CompositeToken.Contract.RenounceOwnership(&_CompositeToken.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _newController) returns()
func (_CompositeToken *CompositeTokenTransactor) SetController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "setController", _newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _newController) returns()
func (_CompositeToken *CompositeTokenSession) SetController(_newController common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.SetController(&_CompositeToken.TransactOpts, _newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _newController) returns()
func (_CompositeToken *CompositeTokenTransactorSession) SetController(_newController common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.SetController(&_CompositeToken.TransactOpts, _newController)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.Contract.Transfer(&_CompositeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.Contract.Transfer(&_CompositeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.Contract.TransferFrom(&_CompositeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CompositeToken *CompositeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CompositeToken.Contract.TransferFrom(&_CompositeToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CompositeToken *CompositeTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CompositeToken *CompositeTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.TransferOwnership(&_CompositeToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CompositeToken *CompositeTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CompositeToken.Contract.TransferOwnership(&_CompositeToken.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CompositeToken *CompositeTokenTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CompositeToken.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CompositeToken *CompositeTokenSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CompositeToken.Contract.UpgradeToAndCall(&_CompositeToken.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CompositeToken *CompositeTokenTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CompositeToken.Contract.UpgradeToAndCall(&_CompositeToken.TransactOpts, newImplementation, data)
}

// CompositeTokenAccountRegisteredIterator is returned from FilterAccountRegistered and is used to iterate over the raw logs and unpacked data for AccountRegistered events raised by the CompositeToken contract.
type CompositeTokenAccountRegisteredIterator struct {
	Event *CompositeTokenAccountRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenAccountRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenAccountRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenAccountRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenAccountRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenAccountRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenAccountRegistered represents a AccountRegistered event raised by the CompositeToken contract.
type CompositeTokenAccountRegistered struct {
	Token   common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountRegistered is a free log retrieval operation binding the contract event 0x415c9c2b518f5b14d4a1dd06e9ea39351fbcf788a8000a8726eb2ad48fca781f.
//
// Solidity: event AccountRegistered(address indexed token, address account)
func (_CompositeToken *CompositeTokenFilterer) FilterAccountRegistered(opts *bind.FilterOpts, token []common.Address) (*CompositeTokenAccountRegisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "AccountRegistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenAccountRegisteredIterator{contract: _CompositeToken.contract, event: "AccountRegistered", logs: logs, sub: sub}, nil
}

// WatchAccountRegistered is a free log subscription operation binding the contract event 0x415c9c2b518f5b14d4a1dd06e9ea39351fbcf788a8000a8726eb2ad48fca781f.
//
// Solidity: event AccountRegistered(address indexed token, address account)
func (_CompositeToken *CompositeTokenFilterer) WatchAccountRegistered(opts *bind.WatchOpts, sink chan<- *CompositeTokenAccountRegistered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "AccountRegistered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenAccountRegistered)
				if err := _CompositeToken.contract.UnpackLog(event, "AccountRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountRegistered is a log parse operation binding the contract event 0x415c9c2b518f5b14d4a1dd06e9ea39351fbcf788a8000a8726eb2ad48fca781f.
//
// Solidity: event AccountRegistered(address indexed token, address account)
func (_CompositeToken *CompositeTokenFilterer) ParseAccountRegistered(log types.Log) (*CompositeTokenAccountRegistered, error) {
	event := new(CompositeTokenAccountRegistered)
	if err := _CompositeToken.contract.UnpackLog(event, "AccountRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CompositeToken contract.
type CompositeTokenApprovalIterator struct {
	Event *CompositeTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenApproval represents a Approval event raised by the CompositeToken contract.
type CompositeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CompositeToken *CompositeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CompositeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenApprovalIterator{contract: _CompositeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CompositeToken *CompositeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CompositeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenApproval)
				if err := _CompositeToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CompositeToken *CompositeTokenFilterer) ParseApproval(log types.Log) (*CompositeTokenApproval, error) {
	event := new(CompositeTokenApproval)
	if err := _CompositeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenControllerSetIterator is returned from FilterControllerSet and is used to iterate over the raw logs and unpacked data for ControllerSet events raised by the CompositeToken contract.
type CompositeTokenControllerSetIterator struct {
	Event *CompositeTokenControllerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenControllerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenControllerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenControllerSet represents a ControllerSet event raised by the CompositeToken contract.
type CompositeTokenControllerSet struct {
	OldController common.Address
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterControllerSet is a free log retrieval operation binding the contract event 0x007582b62407f53d49cfc72e7ddab574c06ef3b8aced104b54b7bed4681ee54a.
//
// Solidity: event ControllerSet(address oldController, address newController)
func (_CompositeToken *CompositeTokenFilterer) FilterControllerSet(opts *bind.FilterOpts) (*CompositeTokenControllerSetIterator, error) {

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "ControllerSet")
	if err != nil {
		return nil, err
	}
	return &CompositeTokenControllerSetIterator{contract: _CompositeToken.contract, event: "ControllerSet", logs: logs, sub: sub}, nil
}

// WatchControllerSet is a free log subscription operation binding the contract event 0x007582b62407f53d49cfc72e7ddab574c06ef3b8aced104b54b7bed4681ee54a.
//
// Solidity: event ControllerSet(address oldController, address newController)
func (_CompositeToken *CompositeTokenFilterer) WatchControllerSet(opts *bind.WatchOpts, sink chan<- *CompositeTokenControllerSet) (event.Subscription, error) {

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "ControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenControllerSet)
				if err := _CompositeToken.contract.UnpackLog(event, "ControllerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseControllerSet is a log parse operation binding the contract event 0x007582b62407f53d49cfc72e7ddab574c06ef3b8aced104b54b7bed4681ee54a.
//
// Solidity: event ControllerSet(address oldController, address newController)
func (_CompositeToken *CompositeTokenFilterer) ParseControllerSet(log types.Log) (*CompositeTokenControllerSet, error) {
	event := new(CompositeTokenControllerSet)
	if err := _CompositeToken.contract.UnpackLog(event, "ControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CompositeToken contract.
type CompositeTokenInitializedIterator struct {
	Event *CompositeTokenInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenInitialized represents a Initialized event raised by the CompositeToken contract.
type CompositeTokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CompositeToken *CompositeTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*CompositeTokenInitializedIterator, error) {

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CompositeTokenInitializedIterator{contract: _CompositeToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CompositeToken *CompositeTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CompositeTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenInitialized)
				if err := _CompositeToken.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CompositeToken *CompositeTokenFilterer) ParseInitialized(log types.Log) (*CompositeTokenInitialized, error) {
	event := new(CompositeTokenInitialized)
	if err := _CompositeToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CompositeToken contract.
type CompositeTokenOwnershipTransferredIterator struct {
	Event *CompositeTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CompositeToken contract.
type CompositeTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CompositeToken *CompositeTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CompositeTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenOwnershipTransferredIterator{contract: _CompositeToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CompositeToken *CompositeTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CompositeTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenOwnershipTransferred)
				if err := _CompositeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CompositeToken *CompositeTokenFilterer) ParseOwnershipTransferred(log types.Log) (*CompositeTokenOwnershipTransferred, error) {
	event := new(CompositeTokenOwnershipTransferred)
	if err := _CompositeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenTokenSupportedIterator is returned from FilterTokenSupported and is used to iterate over the raw logs and unpacked data for TokenSupported events raised by the CompositeToken contract.
type CompositeTokenTokenSupportedIterator struct {
	Event *CompositeTokenTokenSupported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenTokenSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenTokenSupported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenTokenSupported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenTokenSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenTokenSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenTokenSupported represents a TokenSupported event raised by the CompositeToken contract.
type CompositeTokenTokenSupported struct {
	Token     common.Address
	Supported bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenSupported is a free log retrieval operation binding the contract event 0xd86c88b4a367d9b81a6e00c9d3f7d01657bffef358fa8968c91c045ae99f0d52.
//
// Solidity: event TokenSupported(address token, bool supported)
func (_CompositeToken *CompositeTokenFilterer) FilterTokenSupported(opts *bind.FilterOpts) (*CompositeTokenTokenSupportedIterator, error) {

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "TokenSupported")
	if err != nil {
		return nil, err
	}
	return &CompositeTokenTokenSupportedIterator{contract: _CompositeToken.contract, event: "TokenSupported", logs: logs, sub: sub}, nil
}

// WatchTokenSupported is a free log subscription operation binding the contract event 0xd86c88b4a367d9b81a6e00c9d3f7d01657bffef358fa8968c91c045ae99f0d52.
//
// Solidity: event TokenSupported(address token, bool supported)
func (_CompositeToken *CompositeTokenFilterer) WatchTokenSupported(opts *bind.WatchOpts, sink chan<- *CompositeTokenTokenSupported) (event.Subscription, error) {

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "TokenSupported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenTokenSupported)
				if err := _CompositeToken.contract.UnpackLog(event, "TokenSupported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSupported is a log parse operation binding the contract event 0xd86c88b4a367d9b81a6e00c9d3f7d01657bffef358fa8968c91c045ae99f0d52.
//
// Solidity: event TokenSupported(address token, bool supported)
func (_CompositeToken *CompositeTokenFilterer) ParseTokenSupported(log types.Log) (*CompositeTokenTokenSupported, error) {
	event := new(CompositeTokenTokenSupported)
	if err := _CompositeToken.contract.UnpackLog(event, "TokenSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CompositeToken contract.
type CompositeTokenTransferIterator struct {
	Event *CompositeTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenTransfer represents a Transfer event raised by the CompositeToken contract.
type CompositeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CompositeToken *CompositeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CompositeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenTransferIterator{contract: _CompositeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CompositeToken *CompositeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CompositeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenTransfer)
				if err := _CompositeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CompositeToken *CompositeTokenFilterer) ParseTransfer(log types.Log) (*CompositeTokenTransfer, error) {
	event := new(CompositeTokenTransfer)
	if err := _CompositeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompositeTokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CompositeToken contract.
type CompositeTokenUpgradedIterator struct {
	Event *CompositeTokenUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CompositeTokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompositeTokenUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CompositeTokenUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CompositeTokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompositeTokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompositeTokenUpgraded represents a Upgraded event raised by the CompositeToken contract.
type CompositeTokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CompositeToken *CompositeTokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CompositeTokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CompositeToken.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CompositeTokenUpgradedIterator{contract: _CompositeToken.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CompositeToken *CompositeTokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CompositeTokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CompositeToken.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompositeTokenUpgraded)
				if err := _CompositeToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CompositeToken *CompositeTokenFilterer) ParseUpgraded(log types.Log) (*CompositeTokenUpgraded, error) {
	event := new(CompositeTokenUpgraded)
	if err := _CompositeToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
