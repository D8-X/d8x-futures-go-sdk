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

// OracleFactoryOracleData is an auto generated low-level Go binding around an user-defined struct.
type OracleFactoryOracleData struct {
	Oracle    common.Address
	IsInverse bool
}

// OracleFactoryMetaData contains all meta data concerning the OracleFactory contract.
var OracleFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxFeedTimeGapSec\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pythFeedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_onDemandfeedAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"baseCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"quoteCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"OracleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"baseCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"quoteCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"OracleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"baseCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"quoteCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"oracle\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"isInverse\",\"type\":\"bool[]\"}],\"name\":\"RouteAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"baseCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"quoteCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"marketClosed\",\"type\":\"bool\"}],\"name\":\"SetMarketClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"baseCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"quoteCurrency\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ShortRouteAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"address[]\",\"name\":\"_oracles\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isInverse\",\"type\":\"bool[]\"}],\"name\":\"addRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"_tradingBreakMins\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_feedAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_priceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_feedPeriod\",\"type\":\"uint256\"}],\"name\":\"createOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"}],\"name\":\"existsRoute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"}],\"name\":\"getRoute\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isInverse\",\"type\":\"bool\"}],\"internalType\":\"structOracleFactory.OracleData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4[2]\",\"name\":\"_baseQuote\",\"type\":\"bytes4[2]\"}],\"name\":\"getRouteIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onDemandFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pyth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"_marketClosed\",\"type\":\"bool\"}],\"name\":\"setMarketClosed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_priceIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxAcceptableFeedAge\",\"type\":\"uint256\"}],\"name\":\"updatePriceFeeds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OracleFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleFactoryMetaData.ABI instead.
var OracleFactoryABI = OracleFactoryMetaData.ABI

// OracleFactory is an auto generated Go binding around an Ethereum contract.
type OracleFactory struct {
	OracleFactoryCaller     // Read-only binding to the contract
	OracleFactoryTransactor // Write-only binding to the contract
	OracleFactoryFilterer   // Log filterer for contract events
}

// OracleFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleFactorySession struct {
	Contract     *OracleFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleFactoryCallerSession struct {
	Contract *OracleFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OracleFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleFactoryTransactorSession struct {
	Contract     *OracleFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OracleFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleFactoryRaw struct {
	Contract *OracleFactory // Generic contract binding to access the raw methods on
}

// OracleFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleFactoryCallerRaw struct {
	Contract *OracleFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OracleFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleFactoryTransactorRaw struct {
	Contract *OracleFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleFactory creates a new instance of OracleFactory, bound to a specific deployed contract.
func NewOracleFactory(address common.Address, backend bind.ContractBackend) (*OracleFactory, error) {
	contract, err := bindOracleFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleFactory{OracleFactoryCaller: OracleFactoryCaller{contract: contract}, OracleFactoryTransactor: OracleFactoryTransactor{contract: contract}, OracleFactoryFilterer: OracleFactoryFilterer{contract: contract}}, nil
}

// NewOracleFactoryCaller creates a new read-only instance of OracleFactory, bound to a specific deployed contract.
func NewOracleFactoryCaller(address common.Address, caller bind.ContractCaller) (*OracleFactoryCaller, error) {
	contract, err := bindOracleFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleFactoryCaller{contract: contract}, nil
}

// NewOracleFactoryTransactor creates a new write-only instance of OracleFactory, bound to a specific deployed contract.
func NewOracleFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleFactoryTransactor, error) {
	contract, err := bindOracleFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleFactoryTransactor{contract: contract}, nil
}

// NewOracleFactoryFilterer creates a new log filterer instance of OracleFactory, bound to a specific deployed contract.
func NewOracleFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFactoryFilterer, error) {
	contract, err := bindOracleFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFactoryFilterer{contract: contract}, nil
}

// bindOracleFactory binds a generic wrapper to an already deployed contract.
func bindOracleFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleFactory *OracleFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleFactory.Contract.OracleFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleFactory *OracleFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleFactory.Contract.OracleFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleFactory *OracleFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleFactory.Contract.OracleFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleFactory *OracleFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleFactory *OracleFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleFactory *OracleFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleFactory.Contract.contract.Transact(opts, method, params...)
}

// ExistsRoute is a free data retrieval call binding the contract method 0x781690a8.
//
// Solidity: function existsRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(bool)
func (_OracleFactory *OracleFactoryCaller) ExistsRoute(opts *bind.CallOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte) (bool, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "existsRoute", _baseCurrency, _quoteCurrency)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsRoute is a free data retrieval call binding the contract method 0x781690a8.
//
// Solidity: function existsRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(bool)
func (_OracleFactory *OracleFactorySession) ExistsRoute(_baseCurrency [4]byte, _quoteCurrency [4]byte) (bool, error) {
	return _OracleFactory.Contract.ExistsRoute(&_OracleFactory.CallOpts, _baseCurrency, _quoteCurrency)
}

// ExistsRoute is a free data retrieval call binding the contract method 0x781690a8.
//
// Solidity: function existsRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(bool)
func (_OracleFactory *OracleFactoryCallerSession) ExistsRoute(_baseCurrency [4]byte, _quoteCurrency [4]byte) (bool, error) {
	return _OracleFactory.Contract.ExistsRoute(&_OracleFactory.CallOpts, _baseCurrency, _quoteCurrency)
}

// GetRoute is a free data retrieval call binding the contract method 0xc35c334c.
//
// Solidity: function getRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns((address,bool)[])
func (_OracleFactory *OracleFactoryCaller) GetRoute(opts *bind.CallOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte) ([]OracleFactoryOracleData, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "getRoute", _baseCurrency, _quoteCurrency)

	if err != nil {
		return *new([]OracleFactoryOracleData), err
	}

	out0 := *abi.ConvertType(out[0], new([]OracleFactoryOracleData)).(*[]OracleFactoryOracleData)

	return out0, err

}

// GetRoute is a free data retrieval call binding the contract method 0xc35c334c.
//
// Solidity: function getRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns((address,bool)[])
func (_OracleFactory *OracleFactorySession) GetRoute(_baseCurrency [4]byte, _quoteCurrency [4]byte) ([]OracleFactoryOracleData, error) {
	return _OracleFactory.Contract.GetRoute(&_OracleFactory.CallOpts, _baseCurrency, _quoteCurrency)
}

// GetRoute is a free data retrieval call binding the contract method 0xc35c334c.
//
// Solidity: function getRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns((address,bool)[])
func (_OracleFactory *OracleFactoryCallerSession) GetRoute(_baseCurrency [4]byte, _quoteCurrency [4]byte) ([]OracleFactoryOracleData, error) {
	return _OracleFactory.Contract.GetRoute(&_OracleFactory.CallOpts, _baseCurrency, _quoteCurrency)
}

// GetRouteIds is a free data retrieval call binding the contract method 0x9cbbe3f6.
//
// Solidity: function getRouteIds(bytes4[2] _baseQuote) view returns(bytes32[], bool[])
func (_OracleFactory *OracleFactoryCaller) GetRouteIds(opts *bind.CallOpts, _baseQuote [2][4]byte) ([][32]byte, []bool, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "getRouteIds", _baseQuote)

	if err != nil {
		return *new([][32]byte), *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return out0, out1, err

}

// GetRouteIds is a free data retrieval call binding the contract method 0x9cbbe3f6.
//
// Solidity: function getRouteIds(bytes4[2] _baseQuote) view returns(bytes32[], bool[])
func (_OracleFactory *OracleFactorySession) GetRouteIds(_baseQuote [2][4]byte) ([][32]byte, []bool, error) {
	return _OracleFactory.Contract.GetRouteIds(&_OracleFactory.CallOpts, _baseQuote)
}

// GetRouteIds is a free data retrieval call binding the contract method 0x9cbbe3f6.
//
// Solidity: function getRouteIds(bytes4[2] _baseQuote) view returns(bytes32[], bool[])
func (_OracleFactory *OracleFactoryCallerSession) GetRouteIds(_baseQuote [2][4]byte) ([][32]byte, []bool, error) {
	return _OracleFactory.Contract.GetRouteIds(&_OracleFactory.CallOpts, _baseQuote)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0xd99454c3.
//
// Solidity: function getSpotPrice(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(int128, uint256)
func (_OracleFactory *OracleFactoryCaller) GetSpotPrice(opts *bind.CallOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "getSpotPrice", _baseCurrency, _quoteCurrency)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0xd99454c3.
//
// Solidity: function getSpotPrice(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(int128, uint256)
func (_OracleFactory *OracleFactorySession) GetSpotPrice(_baseCurrency [4]byte, _quoteCurrency [4]byte) (*big.Int, *big.Int, error) {
	return _OracleFactory.Contract.GetSpotPrice(&_OracleFactory.CallOpts, _baseCurrency, _quoteCurrency)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0xd99454c3.
//
// Solidity: function getSpotPrice(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(int128, uint256)
func (_OracleFactory *OracleFactoryCallerSession) GetSpotPrice(_baseCurrency [4]byte, _quoteCurrency [4]byte) (*big.Int, *big.Int, error) {
	return _OracleFactory.Contract.GetSpotPrice(&_OracleFactory.CallOpts, _baseCurrency, _quoteCurrency)
}

// OnDemandFeed is a free data retrieval call binding the contract method 0x3b12b600.
//
// Solidity: function onDemandFeed() view returns(address)
func (_OracleFactory *OracleFactoryCaller) OnDemandFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "onDemandFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnDemandFeed is a free data retrieval call binding the contract method 0x3b12b600.
//
// Solidity: function onDemandFeed() view returns(address)
func (_OracleFactory *OracleFactorySession) OnDemandFeed() (common.Address, error) {
	return _OracleFactory.Contract.OnDemandFeed(&_OracleFactory.CallOpts)
}

// OnDemandFeed is a free data retrieval call binding the contract method 0x3b12b600.
//
// Solidity: function onDemandFeed() view returns(address)
func (_OracleFactory *OracleFactoryCallerSession) OnDemandFeed() (common.Address, error) {
	return _OracleFactory.Contract.OnDemandFeed(&_OracleFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleFactory *OracleFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleFactory *OracleFactorySession) Owner() (common.Address, error) {
	return _OracleFactory.Contract.Owner(&_OracleFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleFactory *OracleFactoryCallerSession) Owner() (common.Address, error) {
	return _OracleFactory.Contract.Owner(&_OracleFactory.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_OracleFactory *OracleFactoryCaller) Pyth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleFactory.contract.Call(opts, &out, "pyth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_OracleFactory *OracleFactorySession) Pyth() (common.Address, error) {
	return _OracleFactory.Contract.Pyth(&_OracleFactory.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_OracleFactory *OracleFactoryCallerSession) Pyth() (common.Address, error) {
	return _OracleFactory.Contract.Pyth(&_OracleFactory.CallOpts)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address _oracle) returns()
func (_OracleFactory *OracleFactoryTransactor) AddOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "addOracle", _oracle)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address _oracle) returns()
func (_OracleFactory *OracleFactorySession) AddOracle(_oracle common.Address) (*types.Transaction, error) {
	return _OracleFactory.Contract.AddOracle(&_OracleFactory.TransactOpts, _oracle)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address _oracle) returns()
func (_OracleFactory *OracleFactoryTransactorSession) AddOracle(_oracle common.Address) (*types.Transaction, error) {
	return _OracleFactory.Contract.AddOracle(&_OracleFactory.TransactOpts, _oracle)
}

// AddRoute is a paid mutator transaction binding the contract method 0x38125219.
//
// Solidity: function addRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency, address[] _oracles, bool[] _isInverse) returns()
func (_OracleFactory *OracleFactoryTransactor) AddRoute(opts *bind.TransactOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte, _oracles []common.Address, _isInverse []bool) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "addRoute", _baseCurrency, _quoteCurrency, _oracles, _isInverse)
}

// AddRoute is a paid mutator transaction binding the contract method 0x38125219.
//
// Solidity: function addRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency, address[] _oracles, bool[] _isInverse) returns()
func (_OracleFactory *OracleFactorySession) AddRoute(_baseCurrency [4]byte, _quoteCurrency [4]byte, _oracles []common.Address, _isInverse []bool) (*types.Transaction, error) {
	return _OracleFactory.Contract.AddRoute(&_OracleFactory.TransactOpts, _baseCurrency, _quoteCurrency, _oracles, _isInverse)
}

// AddRoute is a paid mutator transaction binding the contract method 0x38125219.
//
// Solidity: function addRoute(bytes4 _baseCurrency, bytes4 _quoteCurrency, address[] _oracles, bool[] _isInverse) returns()
func (_OracleFactory *OracleFactoryTransactorSession) AddRoute(_baseCurrency [4]byte, _quoteCurrency [4]byte, _oracles []common.Address, _isInverse []bool) (*types.Transaction, error) {
	return _OracleFactory.Contract.AddRoute(&_OracleFactory.TransactOpts, _baseCurrency, _quoteCurrency, _oracles, _isInverse)
}

// CreateOracle is a paid mutator transaction binding the contract method 0x4ce0763f.
//
// Solidity: function createOracle(bytes4 _baseCurrency, bytes4 _quoteCurrency, uint16 _tradingBreakMins, address _feedAddress, bytes32 _priceId, uint256 _feedPeriod) returns(address)
func (_OracleFactory *OracleFactoryTransactor) CreateOracle(opts *bind.TransactOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte, _tradingBreakMins uint16, _feedAddress common.Address, _priceId [32]byte, _feedPeriod *big.Int) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "createOracle", _baseCurrency, _quoteCurrency, _tradingBreakMins, _feedAddress, _priceId, _feedPeriod)
}

// CreateOracle is a paid mutator transaction binding the contract method 0x4ce0763f.
//
// Solidity: function createOracle(bytes4 _baseCurrency, bytes4 _quoteCurrency, uint16 _tradingBreakMins, address _feedAddress, bytes32 _priceId, uint256 _feedPeriod) returns(address)
func (_OracleFactory *OracleFactorySession) CreateOracle(_baseCurrency [4]byte, _quoteCurrency [4]byte, _tradingBreakMins uint16, _feedAddress common.Address, _priceId [32]byte, _feedPeriod *big.Int) (*types.Transaction, error) {
	return _OracleFactory.Contract.CreateOracle(&_OracleFactory.TransactOpts, _baseCurrency, _quoteCurrency, _tradingBreakMins, _feedAddress, _priceId, _feedPeriod)
}

// CreateOracle is a paid mutator transaction binding the contract method 0x4ce0763f.
//
// Solidity: function createOracle(bytes4 _baseCurrency, bytes4 _quoteCurrency, uint16 _tradingBreakMins, address _feedAddress, bytes32 _priceId, uint256 _feedPeriod) returns(address)
func (_OracleFactory *OracleFactoryTransactorSession) CreateOracle(_baseCurrency [4]byte, _quoteCurrency [4]byte, _tradingBreakMins uint16, _feedAddress common.Address, _priceId [32]byte, _feedPeriod *big.Int) (*types.Transaction, error) {
	return _OracleFactory.Contract.CreateOracle(&_OracleFactory.TransactOpts, _baseCurrency, _quoteCurrency, _tradingBreakMins, _feedAddress, _priceId, _feedPeriod)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleFactory *OracleFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleFactory *OracleFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleFactory.Contract.RenounceOwnership(&_OracleFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OracleFactory *OracleFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OracleFactory.Contract.RenounceOwnership(&_OracleFactory.TransactOpts)
}

// SetMarketClosed is a paid mutator transaction binding the contract method 0x0b3bea39.
//
// Solidity: function setMarketClosed(bytes4 _baseCurrency, bytes4 _quoteCurrency, bool _marketClosed) returns()
func (_OracleFactory *OracleFactoryTransactor) SetMarketClosed(opts *bind.TransactOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte, _marketClosed bool) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "setMarketClosed", _baseCurrency, _quoteCurrency, _marketClosed)
}

// SetMarketClosed is a paid mutator transaction binding the contract method 0x0b3bea39.
//
// Solidity: function setMarketClosed(bytes4 _baseCurrency, bytes4 _quoteCurrency, bool _marketClosed) returns()
func (_OracleFactory *OracleFactorySession) SetMarketClosed(_baseCurrency [4]byte, _quoteCurrency [4]byte, _marketClosed bool) (*types.Transaction, error) {
	return _OracleFactory.Contract.SetMarketClosed(&_OracleFactory.TransactOpts, _baseCurrency, _quoteCurrency, _marketClosed)
}

// SetMarketClosed is a paid mutator transaction binding the contract method 0x0b3bea39.
//
// Solidity: function setMarketClosed(bytes4 _baseCurrency, bytes4 _quoteCurrency, bool _marketClosed) returns()
func (_OracleFactory *OracleFactoryTransactorSession) SetMarketClosed(_baseCurrency [4]byte, _quoteCurrency [4]byte, _marketClosed bool) (*types.Transaction, error) {
	return _OracleFactory.Contract.SetMarketClosed(&_OracleFactory.TransactOpts, _baseCurrency, _quoteCurrency, _marketClosed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleFactory *OracleFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleFactory *OracleFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleFactory.Contract.TransferOwnership(&_OracleFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OracleFactory *OracleFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OracleFactory.Contract.TransferOwnership(&_OracleFactory.TransactOpts, newOwner)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xd09a9043.
//
// Solidity: function updatePriceFeeds(bytes[] _updateData, bytes32[] _priceIds, uint64[] _publishTimes, uint256 _maxAcceptableFeedAge) payable returns()
func (_OracleFactory *OracleFactoryTransactor) UpdatePriceFeeds(opts *bind.TransactOpts, _updateData [][]byte, _priceIds [][32]byte, _publishTimes []uint64, _maxAcceptableFeedAge *big.Int) (*types.Transaction, error) {
	return _OracleFactory.contract.Transact(opts, "updatePriceFeeds", _updateData, _priceIds, _publishTimes, _maxAcceptableFeedAge)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xd09a9043.
//
// Solidity: function updatePriceFeeds(bytes[] _updateData, bytes32[] _priceIds, uint64[] _publishTimes, uint256 _maxAcceptableFeedAge) payable returns()
func (_OracleFactory *OracleFactorySession) UpdatePriceFeeds(_updateData [][]byte, _priceIds [][32]byte, _publishTimes []uint64, _maxAcceptableFeedAge *big.Int) (*types.Transaction, error) {
	return _OracleFactory.Contract.UpdatePriceFeeds(&_OracleFactory.TransactOpts, _updateData, _priceIds, _publishTimes, _maxAcceptableFeedAge)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xd09a9043.
//
// Solidity: function updatePriceFeeds(bytes[] _updateData, bytes32[] _priceIds, uint64[] _publishTimes, uint256 _maxAcceptableFeedAge) payable returns()
func (_OracleFactory *OracleFactoryTransactorSession) UpdatePriceFeeds(_updateData [][]byte, _priceIds [][32]byte, _publishTimes []uint64, _maxAcceptableFeedAge *big.Int) (*types.Transaction, error) {
	return _OracleFactory.Contract.UpdatePriceFeeds(&_OracleFactory.TransactOpts, _updateData, _priceIds, _publishTimes, _maxAcceptableFeedAge)
}

// OracleFactoryOracleAddedIterator is returned from FilterOracleAdded and is used to iterate over the raw logs and unpacked data for OracleAdded events raised by the OracleFactory contract.
type OracleFactoryOracleAddedIterator struct {
	Event *OracleFactoryOracleAdded // Event containing the contract specifics and raw log

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
func (it *OracleFactoryOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFactoryOracleAdded)
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
		it.Event = new(OracleFactoryOracleAdded)
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
func (it *OracleFactoryOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFactoryOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFactoryOracleAdded represents a OracleAdded event raised by the OracleFactory contract.
type OracleFactoryOracleAdded struct {
	BaseCurrency  [4]byte
	QuoteCurrency [4]byte
	Oracle        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleAdded is a free log retrieval operation binding the contract event 0xe621fbccb18360e45070d412142ac8af44c72c78a29680523d70980a8c0f9cd0.
//
// Solidity: event OracleAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) FilterOracleAdded(opts *bind.FilterOpts) (*OracleFactoryOracleAddedIterator, error) {

	logs, sub, err := _OracleFactory.contract.FilterLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return &OracleFactoryOracleAddedIterator{contract: _OracleFactory.contract, event: "OracleAdded", logs: logs, sub: sub}, nil
}

// WatchOracleAdded is a free log subscription operation binding the contract event 0xe621fbccb18360e45070d412142ac8af44c72c78a29680523d70980a8c0f9cd0.
//
// Solidity: event OracleAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) WatchOracleAdded(opts *bind.WatchOpts, sink chan<- *OracleFactoryOracleAdded) (event.Subscription, error) {

	logs, sub, err := _OracleFactory.contract.WatchLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFactoryOracleAdded)
				if err := _OracleFactory.contract.UnpackLog(event, "OracleAdded", log); err != nil {
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

// ParseOracleAdded is a log parse operation binding the contract event 0xe621fbccb18360e45070d412142ac8af44c72c78a29680523d70980a8c0f9cd0.
//
// Solidity: event OracleAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) ParseOracleAdded(log types.Log) (*OracleFactoryOracleAdded, error) {
	event := new(OracleFactoryOracleAdded)
	if err := _OracleFactory.contract.UnpackLog(event, "OracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFactoryOracleCreatedIterator is returned from FilterOracleCreated and is used to iterate over the raw logs and unpacked data for OracleCreated events raised by the OracleFactory contract.
type OracleFactoryOracleCreatedIterator struct {
	Event *OracleFactoryOracleCreated // Event containing the contract specifics and raw log

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
func (it *OracleFactoryOracleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFactoryOracleCreated)
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
		it.Event = new(OracleFactoryOracleCreated)
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
func (it *OracleFactoryOracleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFactoryOracleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFactoryOracleCreated represents a OracleCreated event raised by the OracleFactory contract.
type OracleFactoryOracleCreated struct {
	BaseCurrency  [4]byte
	QuoteCurrency [4]byte
	Oracle        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleCreated is a free log retrieval operation binding the contract event 0x5c7565604e115e293a940d15499869f3be962833e1741dcc29035b3a9e4a3697.
//
// Solidity: event OracleCreated(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) FilterOracleCreated(opts *bind.FilterOpts) (*OracleFactoryOracleCreatedIterator, error) {

	logs, sub, err := _OracleFactory.contract.FilterLogs(opts, "OracleCreated")
	if err != nil {
		return nil, err
	}
	return &OracleFactoryOracleCreatedIterator{contract: _OracleFactory.contract, event: "OracleCreated", logs: logs, sub: sub}, nil
}

// WatchOracleCreated is a free log subscription operation binding the contract event 0x5c7565604e115e293a940d15499869f3be962833e1741dcc29035b3a9e4a3697.
//
// Solidity: event OracleCreated(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) WatchOracleCreated(opts *bind.WatchOpts, sink chan<- *OracleFactoryOracleCreated) (event.Subscription, error) {

	logs, sub, err := _OracleFactory.contract.WatchLogs(opts, "OracleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFactoryOracleCreated)
				if err := _OracleFactory.contract.UnpackLog(event, "OracleCreated", log); err != nil {
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

// ParseOracleCreated is a log parse operation binding the contract event 0x5c7565604e115e293a940d15499869f3be962833e1741dcc29035b3a9e4a3697.
//
// Solidity: event OracleCreated(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) ParseOracleCreated(log types.Log) (*OracleFactoryOracleCreated, error) {
	event := new(OracleFactoryOracleCreated)
	if err := _OracleFactory.contract.UnpackLog(event, "OracleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OracleFactory contract.
type OracleFactoryOwnershipTransferredIterator struct {
	Event *OracleFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFactoryOwnershipTransferred)
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
		it.Event = new(OracleFactoryOwnershipTransferred)
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
func (it *OracleFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the OracleFactory contract.
type OracleFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleFactory *OracleFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleFactoryOwnershipTransferredIterator{contract: _OracleFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OracleFactory *OracleFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OracleFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFactoryOwnershipTransferred)
				if err := _OracleFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OracleFactory *OracleFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*OracleFactoryOwnershipTransferred, error) {
	event := new(OracleFactoryOwnershipTransferred)
	if err := _OracleFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFactoryRouteAddedIterator is returned from FilterRouteAdded and is used to iterate over the raw logs and unpacked data for RouteAdded events raised by the OracleFactory contract.
type OracleFactoryRouteAddedIterator struct {
	Event *OracleFactoryRouteAdded // Event containing the contract specifics and raw log

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
func (it *OracleFactoryRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFactoryRouteAdded)
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
		it.Event = new(OracleFactoryRouteAdded)
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
func (it *OracleFactoryRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFactoryRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFactoryRouteAdded represents a RouteAdded event raised by the OracleFactory contract.
type OracleFactoryRouteAdded struct {
	BaseCurrency  [4]byte
	QuoteCurrency [4]byte
	Oracle        []common.Address
	IsInverse     []bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRouteAdded is a free log retrieval operation binding the contract event 0x2207316e7d706f640f70ee13b93bbec0ed047a07a5cf63e6c0bf112c156fb918.
//
// Solidity: event RouteAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address[] oracle, bool[] isInverse)
func (_OracleFactory *OracleFactoryFilterer) FilterRouteAdded(opts *bind.FilterOpts) (*OracleFactoryRouteAddedIterator, error) {

	logs, sub, err := _OracleFactory.contract.FilterLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return &OracleFactoryRouteAddedIterator{contract: _OracleFactory.contract, event: "RouteAdded", logs: logs, sub: sub}, nil
}

// WatchRouteAdded is a free log subscription operation binding the contract event 0x2207316e7d706f640f70ee13b93bbec0ed047a07a5cf63e6c0bf112c156fb918.
//
// Solidity: event RouteAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address[] oracle, bool[] isInverse)
func (_OracleFactory *OracleFactoryFilterer) WatchRouteAdded(opts *bind.WatchOpts, sink chan<- *OracleFactoryRouteAdded) (event.Subscription, error) {

	logs, sub, err := _OracleFactory.contract.WatchLogs(opts, "RouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFactoryRouteAdded)
				if err := _OracleFactory.contract.UnpackLog(event, "RouteAdded", log); err != nil {
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

// ParseRouteAdded is a log parse operation binding the contract event 0x2207316e7d706f640f70ee13b93bbec0ed047a07a5cf63e6c0bf112c156fb918.
//
// Solidity: event RouteAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address[] oracle, bool[] isInverse)
func (_OracleFactory *OracleFactoryFilterer) ParseRouteAdded(log types.Log) (*OracleFactoryRouteAdded, error) {
	event := new(OracleFactoryRouteAdded)
	if err := _OracleFactory.contract.UnpackLog(event, "RouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFactorySetMarketClosedIterator is returned from FilterSetMarketClosed and is used to iterate over the raw logs and unpacked data for SetMarketClosed events raised by the OracleFactory contract.
type OracleFactorySetMarketClosedIterator struct {
	Event *OracleFactorySetMarketClosed // Event containing the contract specifics and raw log

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
func (it *OracleFactorySetMarketClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFactorySetMarketClosed)
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
		it.Event = new(OracleFactorySetMarketClosed)
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
func (it *OracleFactorySetMarketClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFactorySetMarketClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFactorySetMarketClosed represents a SetMarketClosed event raised by the OracleFactory contract.
type OracleFactorySetMarketClosed struct {
	BaseCurrency  [4]byte
	QuoteCurrency [4]byte
	Oracle        common.Address
	MarketClosed  bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetMarketClosed is a free log retrieval operation binding the contract event 0xa820cee6e9183dedbc6cc3fcb50a4aa424da7a07bed41482a24542fcfdd4b7da.
//
// Solidity: event SetMarketClosed(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle, bool marketClosed)
func (_OracleFactory *OracleFactoryFilterer) FilterSetMarketClosed(opts *bind.FilterOpts) (*OracleFactorySetMarketClosedIterator, error) {

	logs, sub, err := _OracleFactory.contract.FilterLogs(opts, "SetMarketClosed")
	if err != nil {
		return nil, err
	}
	return &OracleFactorySetMarketClosedIterator{contract: _OracleFactory.contract, event: "SetMarketClosed", logs: logs, sub: sub}, nil
}

// WatchSetMarketClosed is a free log subscription operation binding the contract event 0xa820cee6e9183dedbc6cc3fcb50a4aa424da7a07bed41482a24542fcfdd4b7da.
//
// Solidity: event SetMarketClosed(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle, bool marketClosed)
func (_OracleFactory *OracleFactoryFilterer) WatchSetMarketClosed(opts *bind.WatchOpts, sink chan<- *OracleFactorySetMarketClosed) (event.Subscription, error) {

	logs, sub, err := _OracleFactory.contract.WatchLogs(opts, "SetMarketClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFactorySetMarketClosed)
				if err := _OracleFactory.contract.UnpackLog(event, "SetMarketClosed", log); err != nil {
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

// ParseSetMarketClosed is a log parse operation binding the contract event 0xa820cee6e9183dedbc6cc3fcb50a4aa424da7a07bed41482a24542fcfdd4b7da.
//
// Solidity: event SetMarketClosed(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle, bool marketClosed)
func (_OracleFactory *OracleFactoryFilterer) ParseSetMarketClosed(log types.Log) (*OracleFactorySetMarketClosed, error) {
	event := new(OracleFactorySetMarketClosed)
	if err := _OracleFactory.contract.UnpackLog(event, "SetMarketClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFactoryShortRouteAddedIterator is returned from FilterShortRouteAdded and is used to iterate over the raw logs and unpacked data for ShortRouteAdded events raised by the OracleFactory contract.
type OracleFactoryShortRouteAddedIterator struct {
	Event *OracleFactoryShortRouteAdded // Event containing the contract specifics and raw log

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
func (it *OracleFactoryShortRouteAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFactoryShortRouteAdded)
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
		it.Event = new(OracleFactoryShortRouteAdded)
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
func (it *OracleFactoryShortRouteAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFactoryShortRouteAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFactoryShortRouteAdded represents a ShortRouteAdded event raised by the OracleFactory contract.
type OracleFactoryShortRouteAdded struct {
	BaseCurrency  [4]byte
	QuoteCurrency [4]byte
	Oracle        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterShortRouteAdded is a free log retrieval operation binding the contract event 0xa85cc73a090fbae723a8a5430875c4e7a3779228ef04f65a9421310f98a90b25.
//
// Solidity: event ShortRouteAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) FilterShortRouteAdded(opts *bind.FilterOpts) (*OracleFactoryShortRouteAddedIterator, error) {

	logs, sub, err := _OracleFactory.contract.FilterLogs(opts, "ShortRouteAdded")
	if err != nil {
		return nil, err
	}
	return &OracleFactoryShortRouteAddedIterator{contract: _OracleFactory.contract, event: "ShortRouteAdded", logs: logs, sub: sub}, nil
}

// WatchShortRouteAdded is a free log subscription operation binding the contract event 0xa85cc73a090fbae723a8a5430875c4e7a3779228ef04f65a9421310f98a90b25.
//
// Solidity: event ShortRouteAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) WatchShortRouteAdded(opts *bind.WatchOpts, sink chan<- *OracleFactoryShortRouteAdded) (event.Subscription, error) {

	logs, sub, err := _OracleFactory.contract.WatchLogs(opts, "ShortRouteAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFactoryShortRouteAdded)
				if err := _OracleFactory.contract.UnpackLog(event, "ShortRouteAdded", log); err != nil {
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

// ParseShortRouteAdded is a log parse operation binding the contract event 0xa85cc73a090fbae723a8a5430875c4e7a3779228ef04f65a9421310f98a90b25.
//
// Solidity: event ShortRouteAdded(bytes4 baseCurrency, bytes4 quoteCurrency, address oracle)
func (_OracleFactory *OracleFactoryFilterer) ParseShortRouteAdded(log types.Log) (*OracleFactoryShortRouteAdded, error) {
	event := new(OracleFactoryShortRouteAdded)
	if err := _OracleFactory.contract.UnpackLog(event, "ShortRouteAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
