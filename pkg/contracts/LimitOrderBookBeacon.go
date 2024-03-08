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

// LimitOrderBookBeaconMetaData contains all meta data concerning the LimitOrderBookBeacon contract.
var LimitOrderBookBeaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blueprint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBlueprint\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LimitOrderBookBeaconABI is the input ABI used to generate the binding from.
// Deprecated: Use LimitOrderBookBeaconMetaData.ABI instead.
var LimitOrderBookBeaconABI = LimitOrderBookBeaconMetaData.ABI

// LimitOrderBookBeacon is an auto generated Go binding around an Ethereum contract.
type LimitOrderBookBeacon struct {
	LimitOrderBookBeaconCaller     // Read-only binding to the contract
	LimitOrderBookBeaconTransactor // Write-only binding to the contract
	LimitOrderBookBeaconFilterer   // Log filterer for contract events
}

// LimitOrderBookBeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitOrderBookBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookBeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitOrderBookBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookBeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimitOrderBookBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookBeaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitOrderBookBeaconSession struct {
	Contract     *LimitOrderBookBeacon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LimitOrderBookBeaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitOrderBookBeaconCallerSession struct {
	Contract *LimitOrderBookBeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// LimitOrderBookBeaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitOrderBookBeaconTransactorSession struct {
	Contract     *LimitOrderBookBeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// LimitOrderBookBeaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitOrderBookBeaconRaw struct {
	Contract *LimitOrderBookBeacon // Generic contract binding to access the raw methods on
}

// LimitOrderBookBeaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitOrderBookBeaconCallerRaw struct {
	Contract *LimitOrderBookBeaconCaller // Generic read-only contract binding to access the raw methods on
}

// LimitOrderBookBeaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitOrderBookBeaconTransactorRaw struct {
	Contract *LimitOrderBookBeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitOrderBookBeacon creates a new instance of LimitOrderBookBeacon, bound to a specific deployed contract.
func NewLimitOrderBookBeacon(address common.Address, backend bind.ContractBackend) (*LimitOrderBookBeacon, error) {
	contract, err := bindLimitOrderBookBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookBeacon{LimitOrderBookBeaconCaller: LimitOrderBookBeaconCaller{contract: contract}, LimitOrderBookBeaconTransactor: LimitOrderBookBeaconTransactor{contract: contract}, LimitOrderBookBeaconFilterer: LimitOrderBookBeaconFilterer{contract: contract}}, nil
}

// NewLimitOrderBookBeaconCaller creates a new read-only instance of LimitOrderBookBeacon, bound to a specific deployed contract.
func NewLimitOrderBookBeaconCaller(address common.Address, caller bind.ContractCaller) (*LimitOrderBookBeaconCaller, error) {
	contract, err := bindLimitOrderBookBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookBeaconCaller{contract: contract}, nil
}

// NewLimitOrderBookBeaconTransactor creates a new write-only instance of LimitOrderBookBeacon, bound to a specific deployed contract.
func NewLimitOrderBookBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitOrderBookBeaconTransactor, error) {
	contract, err := bindLimitOrderBookBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookBeaconTransactor{contract: contract}, nil
}

// NewLimitOrderBookBeaconFilterer creates a new log filterer instance of LimitOrderBookBeacon, bound to a specific deployed contract.
func NewLimitOrderBookBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*LimitOrderBookBeaconFilterer, error) {
	contract, err := bindLimitOrderBookBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookBeaconFilterer{contract: contract}, nil
}

// bindLimitOrderBookBeacon binds a generic wrapper to an already deployed contract.
func bindLimitOrderBookBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LimitOrderBookBeaconMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderBookBeacon *LimitOrderBookBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderBookBeacon.Contract.LimitOrderBookBeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderBookBeacon *LimitOrderBookBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.LimitOrderBookBeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderBookBeacon *LimitOrderBookBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.LimitOrderBookBeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderBookBeacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.contract.Transact(opts, method, params...)
}

// Blueprint is a free data retrieval call binding the contract method 0xc05efa15.
//
// Solidity: function blueprint() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCaller) Blueprint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookBeacon.contract.Call(opts, &out, "blueprint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blueprint is a free data retrieval call binding the contract method 0xc05efa15.
//
// Solidity: function blueprint() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconSession) Blueprint() (common.Address, error) {
	return _LimitOrderBookBeacon.Contract.Blueprint(&_LimitOrderBookBeacon.CallOpts)
}

// Blueprint is a free data retrieval call binding the contract method 0xc05efa15.
//
// Solidity: function blueprint() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCallerSession) Blueprint() (common.Address, error) {
	return _LimitOrderBookBeacon.Contract.Blueprint(&_LimitOrderBookBeacon.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookBeacon.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconSession) Implementation() (common.Address, error) {
	return _LimitOrderBookBeacon.Contract.Implementation(&_LimitOrderBookBeacon.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCallerSession) Implementation() (common.Address, error) {
	return _LimitOrderBookBeacon.Contract.Implementation(&_LimitOrderBookBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBookBeacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconSession) Owner() (common.Address, error) {
	return _LimitOrderBookBeacon.Contract.Owner(&_LimitOrderBookBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconCallerSession) Owner() (common.Address, error) {
	return _LimitOrderBookBeacon.Contract.Owner(&_LimitOrderBookBeacon.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconSession) RenounceOwnership() (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.RenounceOwnership(&_LimitOrderBookBeacon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.RenounceOwnership(&_LimitOrderBookBeacon.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.TransferOwnership(&_LimitOrderBookBeacon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.TransferOwnership(&_LimitOrderBookBeacon.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(address _newBlueprint) returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactor) Update(opts *bind.TransactOpts, _newBlueprint common.Address) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.contract.Transact(opts, "update", _newBlueprint)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(address _newBlueprint) returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconSession) Update(_newBlueprint common.Address) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.Update(&_LimitOrderBookBeacon.TransactOpts, _newBlueprint)
}

// Update is a paid mutator transaction binding the contract method 0x1c1b8772.
//
// Solidity: function update(address _newBlueprint) returns()
func (_LimitOrderBookBeacon *LimitOrderBookBeaconTransactorSession) Update(_newBlueprint common.Address) (*types.Transaction, error) {
	return _LimitOrderBookBeacon.Contract.Update(&_LimitOrderBookBeacon.TransactOpts, _newBlueprint)
}

// LimitOrderBookBeaconOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LimitOrderBookBeacon contract.
type LimitOrderBookBeaconOwnershipTransferredIterator struct {
	Event *LimitOrderBookBeaconOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookBeaconOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookBeaconOwnershipTransferred)
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
		it.Event = new(LimitOrderBookBeaconOwnershipTransferred)
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
func (it *LimitOrderBookBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookBeaconOwnershipTransferred represents a OwnershipTransferred event raised by the LimitOrderBookBeacon contract.
type LimitOrderBookBeaconOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LimitOrderBookBeaconOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimitOrderBookBeacon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookBeaconOwnershipTransferredIterator{contract: _LimitOrderBookBeacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimitOrderBookBeacon *LimitOrderBookBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LimitOrderBookBeaconOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimitOrderBookBeacon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookBeaconOwnershipTransferred)
				if err := _LimitOrderBookBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LimitOrderBookBeacon *LimitOrderBookBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*LimitOrderBookBeaconOwnershipTransferred, error) {
	event := new(LimitOrderBookBeaconOwnershipTransferred)
	if err := _LimitOrderBookBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
