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

// PythStructsPrice is an auto generated low-level Go binding around an user-defined struct.
type PythStructsPrice struct {
	Price       int64
	Conf        uint64
	Expo        int32
	PublishTime *big.Int
}

// PythStructsPriceFeed is an auto generated low-level Go binding around an user-defined struct.
type PythStructsPriceFeed struct {
	Id       [32]byte
	Price    PythStructsPrice
	EmaPrice PythStructsPrice
}

// IPythMetaData contains all meta data concerning the IPyth contract.
var IPythMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"BatchPriceFeedUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"publishTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"}],\"name\":\"PriceFeedUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getEmaPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"name\":\"getEmaPriceNoOlderThan\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getEmaPriceUnsafe\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"name\":\"getPriceNoOlderThan\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPriceUnsafe\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"getUpdateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidTimePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validTimePeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"priceIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"minPublishTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPublishTime\",\"type\":\"uint64\"}],\"name\":\"parsePriceFeedUpdates\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"emaPrice\",\"type\":\"tuple\"}],\"internalType\":\"structPythStructs.PriceFeed[]\",\"name\":\"priceFeeds\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"}],\"name\":\"updatePriceFeeds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"priceIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"updatePriceFeedsIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IPythABI is the input ABI used to generate the binding from.
// Deprecated: Use IPythMetaData.ABI instead.
var IPythABI = IPythMetaData.ABI

// IPyth is an auto generated Go binding around an Ethereum contract.
type IPyth struct {
	IPythCaller     // Read-only binding to the contract
	IPythTransactor // Write-only binding to the contract
	IPythFilterer   // Log filterer for contract events
}

// IPythCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPythCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPythTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPythTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPythFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPythFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPythSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPythSession struct {
	Contract     *IPyth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPythCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPythCallerSession struct {
	Contract *IPythCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPythTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPythTransactorSession struct {
	Contract     *IPythTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPythRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPythRaw struct {
	Contract *IPyth // Generic contract binding to access the raw methods on
}

// IPythCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPythCallerRaw struct {
	Contract *IPythCaller // Generic read-only contract binding to access the raw methods on
}

// IPythTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPythTransactorRaw struct {
	Contract *IPythTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPyth creates a new instance of IPyth, bound to a specific deployed contract.
func NewIPyth(address common.Address, backend bind.ContractBackend) (*IPyth, error) {
	contract, err := bindIPyth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPyth{IPythCaller: IPythCaller{contract: contract}, IPythTransactor: IPythTransactor{contract: contract}, IPythFilterer: IPythFilterer{contract: contract}}, nil
}

// NewIPythCaller creates a new read-only instance of IPyth, bound to a specific deployed contract.
func NewIPythCaller(address common.Address, caller bind.ContractCaller) (*IPythCaller, error) {
	contract, err := bindIPyth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPythCaller{contract: contract}, nil
}

// NewIPythTransactor creates a new write-only instance of IPyth, bound to a specific deployed contract.
func NewIPythTransactor(address common.Address, transactor bind.ContractTransactor) (*IPythTransactor, error) {
	contract, err := bindIPyth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPythTransactor{contract: contract}, nil
}

// NewIPythFilterer creates a new log filterer instance of IPyth, bound to a specific deployed contract.
func NewIPythFilterer(address common.Address, filterer bind.ContractFilterer) (*IPythFilterer, error) {
	contract, err := bindIPyth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPythFilterer{contract: contract}, nil
}

// bindIPyth binds a generic wrapper to an already deployed contract.
func bindIPyth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPythMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPyth *IPythRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPyth.Contract.IPythCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPyth *IPythRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPyth.Contract.IPythTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPyth *IPythRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPyth.Contract.IPythTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPyth *IPythCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPyth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPyth *IPythTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPyth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPyth *IPythTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPyth.Contract.contract.Transact(opts, method, params...)
}

// GetEmaPrice is a free data retrieval call binding the contract method 0xb5dcc911.
//
// Solidity: function getEmaPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCaller) GetEmaPrice(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getEmaPrice", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetEmaPrice is a free data retrieval call binding the contract method 0xb5dcc911.
//
// Solidity: function getEmaPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythSession) GetEmaPrice(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetEmaPrice(&_IPyth.CallOpts, id)
}

// GetEmaPrice is a free data retrieval call binding the contract method 0xb5dcc911.
//
// Solidity: function getEmaPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCallerSession) GetEmaPrice(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetEmaPrice(&_IPyth.CallOpts, id)
}

// GetEmaPriceNoOlderThan is a free data retrieval call binding the contract method 0x711a2e28.
//
// Solidity: function getEmaPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCaller) GetEmaPriceNoOlderThan(opts *bind.CallOpts, id [32]byte, age *big.Int) (PythStructsPrice, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getEmaPriceNoOlderThan", id, age)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetEmaPriceNoOlderThan is a free data retrieval call binding the contract method 0x711a2e28.
//
// Solidity: function getEmaPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythSession) GetEmaPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _IPyth.Contract.GetEmaPriceNoOlderThan(&_IPyth.CallOpts, id, age)
}

// GetEmaPriceNoOlderThan is a free data retrieval call binding the contract method 0x711a2e28.
//
// Solidity: function getEmaPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCallerSession) GetEmaPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _IPyth.Contract.GetEmaPriceNoOlderThan(&_IPyth.CallOpts, id, age)
}

// GetEmaPriceUnsafe is a free data retrieval call binding the contract method 0x9474f45b.
//
// Solidity: function getEmaPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCaller) GetEmaPriceUnsafe(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getEmaPriceUnsafe", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetEmaPriceUnsafe is a free data retrieval call binding the contract method 0x9474f45b.
//
// Solidity: function getEmaPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythSession) GetEmaPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetEmaPriceUnsafe(&_IPyth.CallOpts, id)
}

// GetEmaPriceUnsafe is a free data retrieval call binding the contract method 0x9474f45b.
//
// Solidity: function getEmaPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCallerSession) GetEmaPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetEmaPriceUnsafe(&_IPyth.CallOpts, id)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCaller) GetPrice(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getPrice", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythSession) GetPrice(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetPrice(&_IPyth.CallOpts, id)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCallerSession) GetPrice(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetPrice(&_IPyth.CallOpts, id)
}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCaller) GetPriceNoOlderThan(opts *bind.CallOpts, id [32]byte, age *big.Int) (PythStructsPrice, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getPriceNoOlderThan", id, age)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythSession) GetPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _IPyth.Contract.GetPriceNoOlderThan(&_IPyth.CallOpts, id, age)
}

// GetPriceNoOlderThan is a free data retrieval call binding the contract method 0xa4ae35e0.
//
// Solidity: function getPriceNoOlderThan(bytes32 id, uint256 age) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCallerSession) GetPriceNoOlderThan(id [32]byte, age *big.Int) (PythStructsPrice, error) {
	return _IPyth.Contract.GetPriceNoOlderThan(&_IPyth.CallOpts, id, age)
}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCaller) GetPriceUnsafe(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getPriceUnsafe", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythSession) GetPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetPriceUnsafe(&_IPyth.CallOpts, id)
}

// GetPriceUnsafe is a free data retrieval call binding the contract method 0x96834ad3.
//
// Solidity: function getPriceUnsafe(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_IPyth *IPythCallerSession) GetPriceUnsafe(id [32]byte) (PythStructsPrice, error) {
	return _IPyth.Contract.GetPriceUnsafe(&_IPyth.CallOpts, id)
}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] updateData) view returns(uint256 feeAmount)
func (_IPyth *IPythCaller) GetUpdateFee(opts *bind.CallOpts, updateData [][]byte) (*big.Int, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getUpdateFee", updateData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] updateData) view returns(uint256 feeAmount)
func (_IPyth *IPythSession) GetUpdateFee(updateData [][]byte) (*big.Int, error) {
	return _IPyth.Contract.GetUpdateFee(&_IPyth.CallOpts, updateData)
}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] updateData) view returns(uint256 feeAmount)
func (_IPyth *IPythCallerSession) GetUpdateFee(updateData [][]byte) (*big.Int, error) {
	return _IPyth.Contract.GetUpdateFee(&_IPyth.CallOpts, updateData)
}

// GetValidTimePeriod is a free data retrieval call binding the contract method 0xe18910a3.
//
// Solidity: function getValidTimePeriod() view returns(uint256 validTimePeriod)
func (_IPyth *IPythCaller) GetValidTimePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPyth.contract.Call(opts, &out, "getValidTimePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidTimePeriod is a free data retrieval call binding the contract method 0xe18910a3.
//
// Solidity: function getValidTimePeriod() view returns(uint256 validTimePeriod)
func (_IPyth *IPythSession) GetValidTimePeriod() (*big.Int, error) {
	return _IPyth.Contract.GetValidTimePeriod(&_IPyth.CallOpts)
}

// GetValidTimePeriod is a free data retrieval call binding the contract method 0xe18910a3.
//
// Solidity: function getValidTimePeriod() view returns(uint256 validTimePeriod)
func (_IPyth *IPythCallerSession) GetValidTimePeriod() (*big.Int, error) {
	return _IPyth.Contract.GetValidTimePeriod(&_IPyth.CallOpts)
}

// ParsePriceFeedUpdates is a paid mutator transaction binding the contract method 0x4716e9c5.
//
// Solidity: function parsePriceFeedUpdates(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_IPyth *IPythTransactor) ParsePriceFeedUpdates(opts *bind.TransactOpts, updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _IPyth.contract.Transact(opts, "parsePriceFeedUpdates", updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdates is a paid mutator transaction binding the contract method 0x4716e9c5.
//
// Solidity: function parsePriceFeedUpdates(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_IPyth *IPythSession) ParsePriceFeedUpdates(updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _IPyth.Contract.ParsePriceFeedUpdates(&_IPyth.TransactOpts, updateData, priceIds, minPublishTime, maxPublishTime)
}

// ParsePriceFeedUpdates is a paid mutator transaction binding the contract method 0x4716e9c5.
//
// Solidity: function parsePriceFeedUpdates(bytes[] updateData, bytes32[] priceIds, uint64 minPublishTime, uint64 maxPublishTime) payable returns((bytes32,(int64,uint64,int32,uint256),(int64,uint64,int32,uint256))[] priceFeeds)
func (_IPyth *IPythTransactorSession) ParsePriceFeedUpdates(updateData [][]byte, priceIds [][32]byte, minPublishTime uint64, maxPublishTime uint64) (*types.Transaction, error) {
	return _IPyth.Contract.ParsePriceFeedUpdates(&_IPyth.TransactOpts, updateData, priceIds, minPublishTime, maxPublishTime)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] updateData) payable returns()
func (_IPyth *IPythTransactor) UpdatePriceFeeds(opts *bind.TransactOpts, updateData [][]byte) (*types.Transaction, error) {
	return _IPyth.contract.Transact(opts, "updatePriceFeeds", updateData)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] updateData) payable returns()
func (_IPyth *IPythSession) UpdatePriceFeeds(updateData [][]byte) (*types.Transaction, error) {
	return _IPyth.Contract.UpdatePriceFeeds(&_IPyth.TransactOpts, updateData)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] updateData) payable returns()
func (_IPyth *IPythTransactorSession) UpdatePriceFeeds(updateData [][]byte) (*types.Transaction, error) {
	return _IPyth.Contract.UpdatePriceFeeds(&_IPyth.TransactOpts, updateData)
}

// UpdatePriceFeedsIfNecessary is a paid mutator transaction binding the contract method 0xb9256d28.
//
// Solidity: function updatePriceFeedsIfNecessary(bytes[] updateData, bytes32[] priceIds, uint64[] publishTimes) payable returns()
func (_IPyth *IPythTransactor) UpdatePriceFeedsIfNecessary(opts *bind.TransactOpts, updateData [][]byte, priceIds [][32]byte, publishTimes []uint64) (*types.Transaction, error) {
	return _IPyth.contract.Transact(opts, "updatePriceFeedsIfNecessary", updateData, priceIds, publishTimes)
}

// UpdatePriceFeedsIfNecessary is a paid mutator transaction binding the contract method 0xb9256d28.
//
// Solidity: function updatePriceFeedsIfNecessary(bytes[] updateData, bytes32[] priceIds, uint64[] publishTimes) payable returns()
func (_IPyth *IPythSession) UpdatePriceFeedsIfNecessary(updateData [][]byte, priceIds [][32]byte, publishTimes []uint64) (*types.Transaction, error) {
	return _IPyth.Contract.UpdatePriceFeedsIfNecessary(&_IPyth.TransactOpts, updateData, priceIds, publishTimes)
}

// UpdatePriceFeedsIfNecessary is a paid mutator transaction binding the contract method 0xb9256d28.
//
// Solidity: function updatePriceFeedsIfNecessary(bytes[] updateData, bytes32[] priceIds, uint64[] publishTimes) payable returns()
func (_IPyth *IPythTransactorSession) UpdatePriceFeedsIfNecessary(updateData [][]byte, priceIds [][32]byte, publishTimes []uint64) (*types.Transaction, error) {
	return _IPyth.Contract.UpdatePriceFeedsIfNecessary(&_IPyth.TransactOpts, updateData, priceIds, publishTimes)
}

// IPythBatchPriceFeedUpdateIterator is returned from FilterBatchPriceFeedUpdate and is used to iterate over the raw logs and unpacked data for BatchPriceFeedUpdate events raised by the IPyth contract.
type IPythBatchPriceFeedUpdateIterator struct {
	Event *IPythBatchPriceFeedUpdate // Event containing the contract specifics and raw log

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
func (it *IPythBatchPriceFeedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPythBatchPriceFeedUpdate)
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
		it.Event = new(IPythBatchPriceFeedUpdate)
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
func (it *IPythBatchPriceFeedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPythBatchPriceFeedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPythBatchPriceFeedUpdate represents a BatchPriceFeedUpdate event raised by the IPyth contract.
type IPythBatchPriceFeedUpdate struct {
	ChainId        uint16
	SequenceNumber uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchPriceFeedUpdate is a free log retrieval operation binding the contract event 0x943f0e8a16c19895fb87cbeb1a349ed86d7f31923089dd36c1a1ed5e300f267b.
//
// Solidity: event BatchPriceFeedUpdate(uint16 chainId, uint64 sequenceNumber)
func (_IPyth *IPythFilterer) FilterBatchPriceFeedUpdate(opts *bind.FilterOpts) (*IPythBatchPriceFeedUpdateIterator, error) {

	logs, sub, err := _IPyth.contract.FilterLogs(opts, "BatchPriceFeedUpdate")
	if err != nil {
		return nil, err
	}
	return &IPythBatchPriceFeedUpdateIterator{contract: _IPyth.contract, event: "BatchPriceFeedUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchPriceFeedUpdate is a free log subscription operation binding the contract event 0x943f0e8a16c19895fb87cbeb1a349ed86d7f31923089dd36c1a1ed5e300f267b.
//
// Solidity: event BatchPriceFeedUpdate(uint16 chainId, uint64 sequenceNumber)
func (_IPyth *IPythFilterer) WatchBatchPriceFeedUpdate(opts *bind.WatchOpts, sink chan<- *IPythBatchPriceFeedUpdate) (event.Subscription, error) {

	logs, sub, err := _IPyth.contract.WatchLogs(opts, "BatchPriceFeedUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPythBatchPriceFeedUpdate)
				if err := _IPyth.contract.UnpackLog(event, "BatchPriceFeedUpdate", log); err != nil {
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

// ParseBatchPriceFeedUpdate is a log parse operation binding the contract event 0x943f0e8a16c19895fb87cbeb1a349ed86d7f31923089dd36c1a1ed5e300f267b.
//
// Solidity: event BatchPriceFeedUpdate(uint16 chainId, uint64 sequenceNumber)
func (_IPyth *IPythFilterer) ParseBatchPriceFeedUpdate(log types.Log) (*IPythBatchPriceFeedUpdate, error) {
	event := new(IPythBatchPriceFeedUpdate)
	if err := _IPyth.contract.UnpackLog(event, "BatchPriceFeedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPythPriceFeedUpdateIterator is returned from FilterPriceFeedUpdate and is used to iterate over the raw logs and unpacked data for PriceFeedUpdate events raised by the IPyth contract.
type IPythPriceFeedUpdateIterator struct {
	Event *IPythPriceFeedUpdate // Event containing the contract specifics and raw log

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
func (it *IPythPriceFeedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPythPriceFeedUpdate)
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
		it.Event = new(IPythPriceFeedUpdate)
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
func (it *IPythPriceFeedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPythPriceFeedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPythPriceFeedUpdate represents a PriceFeedUpdate event raised by the IPyth contract.
type IPythPriceFeedUpdate struct {
	Id          [32]byte
	PublishTime uint64
	Price       int64
	Conf        uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdate is a free log retrieval operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_IPyth *IPythFilterer) FilterPriceFeedUpdate(opts *bind.FilterOpts, id [][32]byte) (*IPythPriceFeedUpdateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IPyth.contract.FilterLogs(opts, "PriceFeedUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return &IPythPriceFeedUpdateIterator{contract: _IPyth.contract, event: "PriceFeedUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdate is a free log subscription operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_IPyth *IPythFilterer) WatchPriceFeedUpdate(opts *bind.WatchOpts, sink chan<- *IPythPriceFeedUpdate, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IPyth.contract.WatchLogs(opts, "PriceFeedUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPythPriceFeedUpdate)
				if err := _IPyth.contract.UnpackLog(event, "PriceFeedUpdate", log); err != nil {
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

// ParsePriceFeedUpdate is a log parse operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_IPyth *IPythFilterer) ParsePriceFeedUpdate(log types.Log) (*IPythPriceFeedUpdate, error) {
	event := new(IPythPriceFeedUpdate)
	if err := _IPyth.contract.UnpackLog(event, "PriceFeedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
