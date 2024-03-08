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

// LimitOrderBookFactoryMetaData contains all meta data concerning the LimitOrderBookFactory contract.
var LimitOrderBookFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"perpManagerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"limitOrderBookAddress\",\"type\":\"address\"}],\"name\":\"PerpetualLimitOrderBookDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"SetCallbackGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"SetExecutor\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_perpetualManagerAddr\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"deployLimitOrderBookProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"orderBooks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"removeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LimitOrderBookFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LimitOrderBookFactoryMetaData.ABI instead.
var LimitOrderBookFactoryABI = LimitOrderBookFactoryMetaData.ABI

// LimitOrderBookFactory is an auto generated Go binding around an Ethereum contract.
type LimitOrderBookFactory struct {
	LimitOrderBookFactoryCaller     // Read-only binding to the contract
	LimitOrderBookFactoryTransactor // Write-only binding to the contract
	LimitOrderBookFactoryFilterer   // Log filterer for contract events
}

// LimitOrderBookFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitOrderBookFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitOrderBookFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimitOrderBookFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitOrderBookFactorySession struct {
	Contract     *LimitOrderBookFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LimitOrderBookFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitOrderBookFactoryCallerSession struct {
	Contract *LimitOrderBookFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// LimitOrderBookFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitOrderBookFactoryTransactorSession struct {
	Contract     *LimitOrderBookFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// LimitOrderBookFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitOrderBookFactoryRaw struct {
	Contract *LimitOrderBookFactory // Generic contract binding to access the raw methods on
}

// LimitOrderBookFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitOrderBookFactoryCallerRaw struct {
	Contract *LimitOrderBookFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LimitOrderBookFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitOrderBookFactoryTransactorRaw struct {
	Contract *LimitOrderBookFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitOrderBookFactory creates a new instance of LimitOrderBookFactory, bound to a specific deployed contract.
func NewLimitOrderBookFactory(address common.Address, backend bind.ContractBackend) (*LimitOrderBookFactory, error) {
	contract, err := bindLimitOrderBookFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactory{LimitOrderBookFactoryCaller: LimitOrderBookFactoryCaller{contract: contract}, LimitOrderBookFactoryTransactor: LimitOrderBookFactoryTransactor{contract: contract}, LimitOrderBookFactoryFilterer: LimitOrderBookFactoryFilterer{contract: contract}}, nil
}

// NewLimitOrderBookFactoryCaller creates a new read-only instance of LimitOrderBookFactory, bound to a specific deployed contract.
func NewLimitOrderBookFactoryCaller(address common.Address, caller bind.ContractCaller) (*LimitOrderBookFactoryCaller, error) {
	contract, err := bindLimitOrderBookFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactoryCaller{contract: contract}, nil
}

// NewLimitOrderBookFactoryTransactor creates a new write-only instance of LimitOrderBookFactory, bound to a specific deployed contract.
func NewLimitOrderBookFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitOrderBookFactoryTransactor, error) {
	contract, err := bindLimitOrderBookFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactoryTransactor{contract: contract}, nil
}

// NewLimitOrderBookFactoryFilterer creates a new log filterer instance of LimitOrderBookFactory, bound to a specific deployed contract.
func NewLimitOrderBookFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LimitOrderBookFactoryFilterer, error) {
	contract, err := bindLimitOrderBookFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactoryFilterer{contract: contract}, nil
}

// bindLimitOrderBookFactory binds a generic wrapper to an already deployed contract.
func bindLimitOrderBookFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LimitOrderBookFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderBookFactory *LimitOrderBookFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderBookFactory.Contract.LimitOrderBookFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderBookFactory *LimitOrderBookFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.LimitOrderBookFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderBookFactory *LimitOrderBookFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.LimitOrderBookFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderBookFactory *LimitOrderBookFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderBookFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.contract.Transact(opts, method, params...)
}

// ApprovedExecutor is a free data retrieval call binding the contract method 0xd770f166.
//
// Solidity: function approvedExecutor(address ) view returns(bool)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCaller) ApprovedExecutor(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LimitOrderBookFactory.contract.Call(opts, &out, "approvedExecutor", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedExecutor is a free data retrieval call binding the contract method 0xd770f166.
//
// Solidity: function approvedExecutor(address ) view returns(bool)
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) ApprovedExecutor(arg0 common.Address) (bool, error) {
	return _LimitOrderBookFactory.Contract.ApprovedExecutor(&_LimitOrderBookFactory.CallOpts, arg0)
}

// ApprovedExecutor is a free data retrieval call binding the contract method 0xd770f166.
//
// Solidity: function approvedExecutor(address ) view returns(bool)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCallerSession) ApprovedExecutor(arg0 common.Address) (bool, error) {
	return _LimitOrderBookFactory.Contract.ApprovedExecutor(&_LimitOrderBookFactory.CallOpts, arg0)
}

// GetBeacon is a free data retrieval call binding the contract method 0x2d6b3a6b.
//
// Solidity: function getBeacon() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCaller) GetBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookFactory.contract.Call(opts, &out, "getBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBeacon is a free data retrieval call binding the contract method 0x2d6b3a6b.
//
// Solidity: function getBeacon() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) GetBeacon() (common.Address, error) {
	return _LimitOrderBookFactory.Contract.GetBeacon(&_LimitOrderBookFactory.CallOpts)
}

// GetBeacon is a free data retrieval call binding the contract method 0x2d6b3a6b.
//
// Solidity: function getBeacon() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCallerSession) GetBeacon() (common.Address, error) {
	return _LimitOrderBookFactory.Contract.GetBeacon(&_LimitOrderBookFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookFactory.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) GetImplementation() (common.Address, error) {
	return _LimitOrderBookFactory.Contract.GetImplementation(&_LimitOrderBookFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCallerSession) GetImplementation() (common.Address, error) {
	return _LimitOrderBookFactory.Contract.GetImplementation(&_LimitOrderBookFactory.CallOpts)
}

// OrderBooks is a free data retrieval call binding the contract method 0x2d558047.
//
// Solidity: function orderBooks(uint24 ) view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCaller) OrderBooks(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookFactory.contract.Call(opts, &out, "orderBooks", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderBooks is a free data retrieval call binding the contract method 0x2d558047.
//
// Solidity: function orderBooks(uint24 ) view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) OrderBooks(arg0 *big.Int) (common.Address, error) {
	return _LimitOrderBookFactory.Contract.OrderBooks(&_LimitOrderBookFactory.CallOpts, arg0)
}

// OrderBooks is a free data retrieval call binding the contract method 0x2d558047.
//
// Solidity: function orderBooks(uint24 ) view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCallerSession) OrderBooks(arg0 *big.Int) (common.Address, error) {
	return _LimitOrderBookFactory.Contract.OrderBooks(&_LimitOrderBookFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) Owner() (common.Address, error) {
	return _LimitOrderBookFactory.Contract.Owner(&_LimitOrderBookFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrderBookFactory *LimitOrderBookFactoryCallerSession) Owner() (common.Address, error) {
	return _LimitOrderBookFactory.Contract.Owner(&_LimitOrderBookFactory.CallOpts)
}

// AddExecutor is a paid mutator transaction binding the contract method 0x1f5a0bbe.
//
// Solidity: function addExecutor(address _executor) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactor) AddExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.contract.Transact(opts, "addExecutor", _executor)
}

// AddExecutor is a paid mutator transaction binding the contract method 0x1f5a0bbe.
//
// Solidity: function addExecutor(address _executor) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) AddExecutor(_executor common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.AddExecutor(&_LimitOrderBookFactory.TransactOpts, _executor)
}

// AddExecutor is a paid mutator transaction binding the contract method 0x1f5a0bbe.
//
// Solidity: function addExecutor(address _executor) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorSession) AddExecutor(_executor common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.AddExecutor(&_LimitOrderBookFactory.TransactOpts, _executor)
}

// DeployLimitOrderBookProxy is a paid mutator transaction binding the contract method 0x074fa708.
//
// Solidity: function deployLimitOrderBookProxy(address _perpetualManagerAddr, uint24 _perpetualId) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactor) DeployLimitOrderBookProxy(opts *bind.TransactOpts, _perpetualManagerAddr common.Address, _perpetualId *big.Int) (*types.Transaction, error) {
	return _LimitOrderBookFactory.contract.Transact(opts, "deployLimitOrderBookProxy", _perpetualManagerAddr, _perpetualId)
}

// DeployLimitOrderBookProxy is a paid mutator transaction binding the contract method 0x074fa708.
//
// Solidity: function deployLimitOrderBookProxy(address _perpetualManagerAddr, uint24 _perpetualId) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) DeployLimitOrderBookProxy(_perpetualManagerAddr common.Address, _perpetualId *big.Int) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.DeployLimitOrderBookProxy(&_LimitOrderBookFactory.TransactOpts, _perpetualManagerAddr, _perpetualId)
}

// DeployLimitOrderBookProxy is a paid mutator transaction binding the contract method 0x074fa708.
//
// Solidity: function deployLimitOrderBookProxy(address _perpetualManagerAddr, uint24 _perpetualId) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorSession) DeployLimitOrderBookProxy(_perpetualManagerAddr common.Address, _perpetualId *big.Int) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.DeployLimitOrderBookProxy(&_LimitOrderBookFactory.TransactOpts, _perpetualManagerAddr, _perpetualId)
}

// RemoveExecutor is a paid mutator transaction binding the contract method 0x24788429.
//
// Solidity: function removeExecutor(address _executor) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactor) RemoveExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.contract.Transact(opts, "removeExecutor", _executor)
}

// RemoveExecutor is a paid mutator transaction binding the contract method 0x24788429.
//
// Solidity: function removeExecutor(address _executor) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) RemoveExecutor(_executor common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.RemoveExecutor(&_LimitOrderBookFactory.TransactOpts, _executor)
}

// RemoveExecutor is a paid mutator transaction binding the contract method 0x24788429.
//
// Solidity: function removeExecutor(address _executor) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorSession) RemoveExecutor(_executor common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.RemoveExecutor(&_LimitOrderBookFactory.TransactOpts, _executor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBookFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.RenounceOwnership(&_LimitOrderBookFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.RenounceOwnership(&_LimitOrderBookFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.TransferOwnership(&_LimitOrderBookFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrderBookFactory *LimitOrderBookFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrderBookFactory.Contract.TransferOwnership(&_LimitOrderBookFactory.TransactOpts, newOwner)
}

// LimitOrderBookFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactoryOwnershipTransferredIterator struct {
	Event *LimitOrderBookFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookFactoryOwnershipTransferred)
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
		it.Event = new(LimitOrderBookFactoryOwnershipTransferred)
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
func (it *LimitOrderBookFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LimitOrderBookFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimitOrderBookFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactoryOwnershipTransferredIterator{contract: _LimitOrderBookFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LimitOrderBookFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimitOrderBookFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookFactoryOwnershipTransferred)
				if err := _LimitOrderBookFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*LimitOrderBookFactoryOwnershipTransferred, error) {
	event := new(LimitOrderBookFactoryOwnershipTransferred)
	if err := _LimitOrderBookFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator is returned from FilterPerpetualLimitOrderBookDeployed and is used to iterate over the raw logs and unpacked data for PerpetualLimitOrderBookDeployed events raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator struct {
	Event *LimitOrderBookFactoryPerpetualLimitOrderBookDeployed // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookFactoryPerpetualLimitOrderBookDeployed)
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
		it.Event = new(LimitOrderBookFactoryPerpetualLimitOrderBookDeployed)
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
func (it *LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookFactoryPerpetualLimitOrderBookDeployed represents a PerpetualLimitOrderBookDeployed event raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactoryPerpetualLimitOrderBookDeployed struct {
	PerpetualId           *big.Int
	PerpManagerAddress    common.Address
	LimitOrderBookAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPerpetualLimitOrderBookDeployed is a free log retrieval operation binding the contract event 0x74fd73207824ecfc332f65fbed1e0b27ae41df0b7ce2b3117a661242f3bc415d.
//
// Solidity: event PerpetualLimitOrderBookDeployed(uint24 indexed perpetualId, address perpManagerAddress, address limitOrderBookAddress)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) FilterPerpetualLimitOrderBookDeployed(opts *bind.FilterOpts, perpetualId []*big.Int) (*LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _LimitOrderBookFactory.contract.FilterLogs(opts, "PerpetualLimitOrderBookDeployed", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactoryPerpetualLimitOrderBookDeployedIterator{contract: _LimitOrderBookFactory.contract, event: "PerpetualLimitOrderBookDeployed", logs: logs, sub: sub}, nil
}

// WatchPerpetualLimitOrderBookDeployed is a free log subscription operation binding the contract event 0x74fd73207824ecfc332f65fbed1e0b27ae41df0b7ce2b3117a661242f3bc415d.
//
// Solidity: event PerpetualLimitOrderBookDeployed(uint24 indexed perpetualId, address perpManagerAddress, address limitOrderBookAddress)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) WatchPerpetualLimitOrderBookDeployed(opts *bind.WatchOpts, sink chan<- *LimitOrderBookFactoryPerpetualLimitOrderBookDeployed, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _LimitOrderBookFactory.contract.WatchLogs(opts, "PerpetualLimitOrderBookDeployed", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookFactoryPerpetualLimitOrderBookDeployed)
				if err := _LimitOrderBookFactory.contract.UnpackLog(event, "PerpetualLimitOrderBookDeployed", log); err != nil {
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

// ParsePerpetualLimitOrderBookDeployed is a log parse operation binding the contract event 0x74fd73207824ecfc332f65fbed1e0b27ae41df0b7ce2b3117a661242f3bc415d.
//
// Solidity: event PerpetualLimitOrderBookDeployed(uint24 indexed perpetualId, address perpManagerAddress, address limitOrderBookAddress)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) ParsePerpetualLimitOrderBookDeployed(log types.Log) (*LimitOrderBookFactoryPerpetualLimitOrderBookDeployed, error) {
	event := new(LimitOrderBookFactoryPerpetualLimitOrderBookDeployed)
	if err := _LimitOrderBookFactory.contract.UnpackLog(event, "PerpetualLimitOrderBookDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderBookFactorySetCallbackGasLimitIterator is returned from FilterSetCallbackGasLimit and is used to iterate over the raw logs and unpacked data for SetCallbackGasLimit events raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactorySetCallbackGasLimitIterator struct {
	Event *LimitOrderBookFactorySetCallbackGasLimit // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookFactorySetCallbackGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookFactorySetCallbackGasLimit)
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
		it.Event = new(LimitOrderBookFactorySetCallbackGasLimit)
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
func (it *LimitOrderBookFactorySetCallbackGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookFactorySetCallbackGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookFactorySetCallbackGasLimit represents a SetCallbackGasLimit event raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactorySetCallbackGasLimit struct {
	PerpetualId *big.Int
	GasLimit    uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetCallbackGasLimit is a free log retrieval operation binding the contract event 0x2eee0fffa52baae148e43a0f2391c49c3dcc1ea635c6ce7f91e974973600719b.
//
// Solidity: event SetCallbackGasLimit(uint24 indexed perpetualId, uint32 gasLimit)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) FilterSetCallbackGasLimit(opts *bind.FilterOpts, perpetualId []*big.Int) (*LimitOrderBookFactorySetCallbackGasLimitIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _LimitOrderBookFactory.contract.FilterLogs(opts, "SetCallbackGasLimit", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactorySetCallbackGasLimitIterator{contract: _LimitOrderBookFactory.contract, event: "SetCallbackGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetCallbackGasLimit is a free log subscription operation binding the contract event 0x2eee0fffa52baae148e43a0f2391c49c3dcc1ea635c6ce7f91e974973600719b.
//
// Solidity: event SetCallbackGasLimit(uint24 indexed perpetualId, uint32 gasLimit)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) WatchSetCallbackGasLimit(opts *bind.WatchOpts, sink chan<- *LimitOrderBookFactorySetCallbackGasLimit, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _LimitOrderBookFactory.contract.WatchLogs(opts, "SetCallbackGasLimit", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookFactorySetCallbackGasLimit)
				if err := _LimitOrderBookFactory.contract.UnpackLog(event, "SetCallbackGasLimit", log); err != nil {
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

// ParseSetCallbackGasLimit is a log parse operation binding the contract event 0x2eee0fffa52baae148e43a0f2391c49c3dcc1ea635c6ce7f91e974973600719b.
//
// Solidity: event SetCallbackGasLimit(uint24 indexed perpetualId, uint32 gasLimit)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) ParseSetCallbackGasLimit(log types.Log) (*LimitOrderBookFactorySetCallbackGasLimit, error) {
	event := new(LimitOrderBookFactorySetCallbackGasLimit)
	if err := _LimitOrderBookFactory.contract.UnpackLog(event, "SetCallbackGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderBookFactorySetExecutorIterator is returned from FilterSetExecutor and is used to iterate over the raw logs and unpacked data for SetExecutor events raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactorySetExecutorIterator struct {
	Event *LimitOrderBookFactorySetExecutor // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookFactorySetExecutorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookFactorySetExecutor)
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
		it.Event = new(LimitOrderBookFactorySetExecutor)
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
func (it *LimitOrderBookFactorySetExecutorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookFactorySetExecutorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookFactorySetExecutor represents a SetExecutor event raised by the LimitOrderBookFactory contract.
type LimitOrderBookFactorySetExecutor struct {
	Executor   common.Address
	Authorized bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetExecutor is a free log retrieval operation binding the contract event 0x827c394aebf0fcbf2d4f5c0107a0031a3860ec0d5e06756e03aacccfb8d2836e.
//
// Solidity: event SetExecutor(address executor, bool authorized)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) FilterSetExecutor(opts *bind.FilterOpts) (*LimitOrderBookFactorySetExecutorIterator, error) {

	logs, sub, err := _LimitOrderBookFactory.contract.FilterLogs(opts, "SetExecutor")
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFactorySetExecutorIterator{contract: _LimitOrderBookFactory.contract, event: "SetExecutor", logs: logs, sub: sub}, nil
}

// WatchSetExecutor is a free log subscription operation binding the contract event 0x827c394aebf0fcbf2d4f5c0107a0031a3860ec0d5e06756e03aacccfb8d2836e.
//
// Solidity: event SetExecutor(address executor, bool authorized)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) WatchSetExecutor(opts *bind.WatchOpts, sink chan<- *LimitOrderBookFactorySetExecutor) (event.Subscription, error) {

	logs, sub, err := _LimitOrderBookFactory.contract.WatchLogs(opts, "SetExecutor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookFactorySetExecutor)
				if err := _LimitOrderBookFactory.contract.UnpackLog(event, "SetExecutor", log); err != nil {
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

// ParseSetExecutor is a log parse operation binding the contract event 0x827c394aebf0fcbf2d4f5c0107a0031a3860ec0d5e06756e03aacccfb8d2836e.
//
// Solidity: event SetExecutor(address executor, bool authorized)
func (_LimitOrderBookFactory *LimitOrderBookFactoryFilterer) ParseSetExecutor(log types.Log) (*LimitOrderBookFactorySetExecutor, error) {
	event := new(LimitOrderBookFactorySetExecutor)
	if err := _LimitOrderBookFactory.contract.UnpackLog(event, "SetExecutor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
