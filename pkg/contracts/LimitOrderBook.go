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

// IClientOrderClientOrder is an auto generated low-level Go binding around an user-defined struct.
type IClientOrderClientOrder struct {
	IPerpetualId       *big.Int
	FLimitPrice        *big.Int
	LeverageTDR        uint16
	ExecutionTimestamp uint32
	Flags              uint32
	IDeadline          uint32
	BrokerAddr         common.Address
	FTriggerPrice      *big.Int
	FAmount            *big.Int
	ParentChildDigest1 [32]byte
	TraderAddr         common.Address
	ParentChildDigest2 [32]byte
	BrokerFeeTbps      uint16
	BrokerSignature    []byte
	CallbackTarget     common.Address
}

// LimitOrderBookMetaData contains all meta data concerning the LimitOrderBook contract.
var LimitOrderBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"Callback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"executorAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"PerpetualLimitOrderCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CALLBACK_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"actvDigestPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"callbackFunctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"digestsOfTrader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_digests\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_executorAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"executeOrders\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"enumLimitOrderBook.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest1\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callbackTarget\",\"type\":\"address\"}],\"internalType\":\"structIClientOrder.ClientOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getTrader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_perpetualManagerAddr\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"_iCancelDelaySec\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"limitDigestsOfTrader\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCloseSwitchTimestamp\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"numberOfDigestsOfTrader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderDependency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest2\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderOfDigest\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"executorAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perpManager\",\"outputs\":[{\"internalType\":\"contractIPerpetualManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perpetualId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numElements\",\"type\":\"uint256\"}],\"name\":\"pollRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest1\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callbackTarget\",\"type\":\"address\"}],\"internalType\":\"structIClientOrder.ClientOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"submittedTs\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest1\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"parentChildDigest2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callbackTarget\",\"type\":\"address\"}],\"internalType\":\"structIClientOrder.ClientOrder[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"postOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LimitOrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use LimitOrderBookMetaData.ABI instead.
var LimitOrderBookABI = LimitOrderBookMetaData.ABI

// LimitOrderBook is an auto generated Go binding around an Ethereum contract.
type LimitOrderBook struct {
	LimitOrderBookCaller     // Read-only binding to the contract
	LimitOrderBookTransactor // Write-only binding to the contract
	LimitOrderBookFilterer   // Log filterer for contract events
}

// LimitOrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitOrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitOrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimitOrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitOrderBookSession struct {
	Contract     *LimitOrderBook   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LimitOrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitOrderBookCallerSession struct {
	Contract *LimitOrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LimitOrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitOrderBookTransactorSession struct {
	Contract     *LimitOrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LimitOrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitOrderBookRaw struct {
	Contract *LimitOrderBook // Generic contract binding to access the raw methods on
}

// LimitOrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitOrderBookCallerRaw struct {
	Contract *LimitOrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// LimitOrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitOrderBookTransactorRaw struct {
	Contract *LimitOrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitOrderBook creates a new instance of LimitOrderBook, bound to a specific deployed contract.
func NewLimitOrderBook(address common.Address, backend bind.ContractBackend) (*LimitOrderBook, error) {
	contract, err := bindLimitOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBook{LimitOrderBookCaller: LimitOrderBookCaller{contract: contract}, LimitOrderBookTransactor: LimitOrderBookTransactor{contract: contract}, LimitOrderBookFilterer: LimitOrderBookFilterer{contract: contract}}, nil
}

// NewLimitOrderBookCaller creates a new read-only instance of LimitOrderBook, bound to a specific deployed contract.
func NewLimitOrderBookCaller(address common.Address, caller bind.ContractCaller) (*LimitOrderBookCaller, error) {
	contract, err := bindLimitOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookCaller{contract: contract}, nil
}

// NewLimitOrderBookTransactor creates a new write-only instance of LimitOrderBook, bound to a specific deployed contract.
func NewLimitOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitOrderBookTransactor, error) {
	contract, err := bindLimitOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookTransactor{contract: contract}, nil
}

// NewLimitOrderBookFilterer creates a new log filterer instance of LimitOrderBook, bound to a specific deployed contract.
func NewLimitOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*LimitOrderBookFilterer, error) {
	contract, err := bindLimitOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookFilterer{contract: contract}, nil
}

// bindLimitOrderBook binds a generic wrapper to an already deployed contract.
func bindLimitOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LimitOrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderBook *LimitOrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderBook.Contract.LimitOrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderBook *LimitOrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.LimitOrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderBook *LimitOrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.LimitOrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrderBook *LimitOrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrderBook *LimitOrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrderBook *LimitOrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.contract.Transact(opts, method, params...)
}

// CALLBACKGASLIMIT is a free data retrieval call binding the contract method 0x33d608f1.
//
// Solidity: function CALLBACK_GAS_LIMIT() view returns(uint32)
func (_LimitOrderBook *LimitOrderBookCaller) CALLBACKGASLIMIT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "CALLBACK_GAS_LIMIT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALLBACKGASLIMIT is a free data retrieval call binding the contract method 0x33d608f1.
//
// Solidity: function CALLBACK_GAS_LIMIT() view returns(uint32)
func (_LimitOrderBook *LimitOrderBookSession) CALLBACKGASLIMIT() (uint32, error) {
	return _LimitOrderBook.Contract.CALLBACKGASLIMIT(&_LimitOrderBook.CallOpts)
}

// CALLBACKGASLIMIT is a free data retrieval call binding the contract method 0x33d608f1.
//
// Solidity: function CALLBACK_GAS_LIMIT() view returns(uint32)
func (_LimitOrderBook *LimitOrderBookCallerSession) CALLBACKGASLIMIT() (uint32, error) {
	return _LimitOrderBook.Contract.CALLBACKGASLIMIT(&_LimitOrderBook.CallOpts)
}

// ActvDigestPos is a free data retrieval call binding the contract method 0x82227d87.
//
// Solidity: function actvDigestPos(bytes32 ) view returns(uint256)
func (_LimitOrderBook *LimitOrderBookCaller) ActvDigestPos(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "actvDigestPos", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActvDigestPos is a free data retrieval call binding the contract method 0x82227d87.
//
// Solidity: function actvDigestPos(bytes32 ) view returns(uint256)
func (_LimitOrderBook *LimitOrderBookSession) ActvDigestPos(arg0 [32]byte) (*big.Int, error) {
	return _LimitOrderBook.Contract.ActvDigestPos(&_LimitOrderBook.CallOpts, arg0)
}

// ActvDigestPos is a free data retrieval call binding the contract method 0x82227d87.
//
// Solidity: function actvDigestPos(bytes32 ) view returns(uint256)
func (_LimitOrderBook *LimitOrderBookCallerSession) ActvDigestPos(arg0 [32]byte) (*big.Int, error) {
	return _LimitOrderBook.Contract.ActvDigestPos(&_LimitOrderBook.CallOpts, arg0)
}

// CallbackFunctions is a free data retrieval call binding the contract method 0x67be6a64.
//
// Solidity: function callbackFunctions(bytes32 ) view returns(address)
func (_LimitOrderBook *LimitOrderBookCaller) CallbackFunctions(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "callbackFunctions", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallbackFunctions is a free data retrieval call binding the contract method 0x67be6a64.
//
// Solidity: function callbackFunctions(bytes32 ) view returns(address)
func (_LimitOrderBook *LimitOrderBookSession) CallbackFunctions(arg0 [32]byte) (common.Address, error) {
	return _LimitOrderBook.Contract.CallbackFunctions(&_LimitOrderBook.CallOpts, arg0)
}

// CallbackFunctions is a free data retrieval call binding the contract method 0x67be6a64.
//
// Solidity: function callbackFunctions(bytes32 ) view returns(address)
func (_LimitOrderBook *LimitOrderBookCallerSession) CallbackFunctions(arg0 [32]byte) (common.Address, error) {
	return _LimitOrderBook.Contract.CallbackFunctions(&_LimitOrderBook.CallOpts, arg0)
}

// DigestsOfTrader is a free data retrieval call binding the contract method 0x79b8783c.
//
// Solidity: function digestsOfTrader(address , uint256 ) view returns(bytes32)
func (_LimitOrderBook *LimitOrderBookCaller) DigestsOfTrader(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "digestsOfTrader", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DigestsOfTrader is a free data retrieval call binding the contract method 0x79b8783c.
//
// Solidity: function digestsOfTrader(address , uint256 ) view returns(bytes32)
func (_LimitOrderBook *LimitOrderBookSession) DigestsOfTrader(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderBook.Contract.DigestsOfTrader(&_LimitOrderBook.CallOpts, arg0, arg1)
}

// DigestsOfTrader is a free data retrieval call binding the contract method 0x79b8783c.
//
// Solidity: function digestsOfTrader(address , uint256 ) view returns(bytes32)
func (_LimitOrderBook *LimitOrderBookCallerSession) DigestsOfTrader(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _LimitOrderBook.Contract.DigestsOfTrader(&_LimitOrderBook.CallOpts, arg0, arg1)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 _digest) view returns(uint8)
func (_LimitOrderBook *LimitOrderBookCaller) GetOrderStatus(opts *bind.CallOpts, _digest [32]byte) (uint8, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "getOrderStatus", _digest)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 _digest) view returns(uint8)
func (_LimitOrderBook *LimitOrderBookSession) GetOrderStatus(_digest [32]byte) (uint8, error) {
	return _LimitOrderBook.Contract.GetOrderStatus(&_LimitOrderBook.CallOpts, _digest)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 _digest) view returns(uint8)
func (_LimitOrderBook *LimitOrderBookCallerSession) GetOrderStatus(_digest [32]byte) (uint8, error) {
	return _LimitOrderBook.Contract.GetOrderStatus(&_LimitOrderBook.CallOpts, _digest)
}

// GetOrders is a free data retrieval call binding the contract method 0x27b62ad9.
//
// Solidity: function getOrders(address trader, uint256 offset, uint256 limit) view returns((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] orders)
func (_LimitOrderBook *LimitOrderBookCaller) GetOrders(opts *bind.CallOpts, trader common.Address, offset *big.Int, limit *big.Int) ([]IClientOrderClientOrder, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "getOrders", trader, offset, limit)

	if err != nil {
		return *new([]IClientOrderClientOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]IClientOrderClientOrder)).(*[]IClientOrderClientOrder)

	return out0, err

}

// GetOrders is a free data retrieval call binding the contract method 0x27b62ad9.
//
// Solidity: function getOrders(address trader, uint256 offset, uint256 limit) view returns((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] orders)
func (_LimitOrderBook *LimitOrderBookSession) GetOrders(trader common.Address, offset *big.Int, limit *big.Int) ([]IClientOrderClientOrder, error) {
	return _LimitOrderBook.Contract.GetOrders(&_LimitOrderBook.CallOpts, trader, offset, limit)
}

// GetOrders is a free data retrieval call binding the contract method 0x27b62ad9.
//
// Solidity: function getOrders(address trader, uint256 offset, uint256 limit) view returns((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] orders)
func (_LimitOrderBook *LimitOrderBookCallerSession) GetOrders(trader common.Address, offset *big.Int, limit *big.Int) ([]IClientOrderClientOrder, error) {
	return _LimitOrderBook.Contract.GetOrders(&_LimitOrderBook.CallOpts, trader, offset, limit)
}

// GetTrader is a free data retrieval call binding the contract method 0xda5c2078.
//
// Solidity: function getTrader(bytes32 digest) view returns(address)
func (_LimitOrderBook *LimitOrderBookCaller) GetTrader(opts *bind.CallOpts, digest [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "getTrader", digest)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTrader is a free data retrieval call binding the contract method 0xda5c2078.
//
// Solidity: function getTrader(bytes32 digest) view returns(address)
func (_LimitOrderBook *LimitOrderBookSession) GetTrader(digest [32]byte) (common.Address, error) {
	return _LimitOrderBook.Contract.GetTrader(&_LimitOrderBook.CallOpts, digest)
}

// GetTrader is a free data retrieval call binding the contract method 0xda5c2078.
//
// Solidity: function getTrader(bytes32 digest) view returns(address)
func (_LimitOrderBook *LimitOrderBookCallerSession) GetTrader(digest [32]byte) (common.Address, error) {
	return _LimitOrderBook.Contract.GetTrader(&_LimitOrderBook.CallOpts, digest)
}

// LastOrderHash is a free data retrieval call binding the contract method 0xcccc2ede.
//
// Solidity: function lastOrderHash() view returns(bytes32)
func (_LimitOrderBook *LimitOrderBookCaller) LastOrderHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "lastOrderHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastOrderHash is a free data retrieval call binding the contract method 0xcccc2ede.
//
// Solidity: function lastOrderHash() view returns(bytes32)
func (_LimitOrderBook *LimitOrderBookSession) LastOrderHash() ([32]byte, error) {
	return _LimitOrderBook.Contract.LastOrderHash(&_LimitOrderBook.CallOpts)
}

// LastOrderHash is a free data retrieval call binding the contract method 0xcccc2ede.
//
// Solidity: function lastOrderHash() view returns(bytes32)
func (_LimitOrderBook *LimitOrderBookCallerSession) LastOrderHash() ([32]byte, error) {
	return _LimitOrderBook.Contract.LastOrderHash(&_LimitOrderBook.CallOpts)
}

// LimitDigestsOfTrader is a free data retrieval call binding the contract method 0x95fafddb.
//
// Solidity: function limitDigestsOfTrader(address trader, uint256 page, uint256 limit) view returns(bytes32[])
func (_LimitOrderBook *LimitOrderBookCaller) LimitDigestsOfTrader(opts *bind.CallOpts, trader common.Address, page *big.Int, limit *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "limitDigestsOfTrader", trader, page, limit)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// LimitDigestsOfTrader is a free data retrieval call binding the contract method 0x95fafddb.
//
// Solidity: function limitDigestsOfTrader(address trader, uint256 page, uint256 limit) view returns(bytes32[])
func (_LimitOrderBook *LimitOrderBookSession) LimitDigestsOfTrader(trader common.Address, page *big.Int, limit *big.Int) ([][32]byte, error) {
	return _LimitOrderBook.Contract.LimitDigestsOfTrader(&_LimitOrderBook.CallOpts, trader, page, limit)
}

// LimitDigestsOfTrader is a free data retrieval call binding the contract method 0x95fafddb.
//
// Solidity: function limitDigestsOfTrader(address trader, uint256 page, uint256 limit) view returns(bytes32[])
func (_LimitOrderBook *LimitOrderBookCallerSession) LimitDigestsOfTrader(trader common.Address, page *big.Int, limit *big.Int) ([][32]byte, error) {
	return _LimitOrderBook.Contract.LimitDigestsOfTrader(&_LimitOrderBook.CallOpts, trader, page, limit)
}

// MarketCloseSwitchTimestamp is a free data retrieval call binding the contract method 0x1a604cba.
//
// Solidity: function marketCloseSwitchTimestamp() view returns(int64)
func (_LimitOrderBook *LimitOrderBookCaller) MarketCloseSwitchTimestamp(opts *bind.CallOpts) (int64, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "marketCloseSwitchTimestamp")

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// MarketCloseSwitchTimestamp is a free data retrieval call binding the contract method 0x1a604cba.
//
// Solidity: function marketCloseSwitchTimestamp() view returns(int64)
func (_LimitOrderBook *LimitOrderBookSession) MarketCloseSwitchTimestamp() (int64, error) {
	return _LimitOrderBook.Contract.MarketCloseSwitchTimestamp(&_LimitOrderBook.CallOpts)
}

// MarketCloseSwitchTimestamp is a free data retrieval call binding the contract method 0x1a604cba.
//
// Solidity: function marketCloseSwitchTimestamp() view returns(int64)
func (_LimitOrderBook *LimitOrderBookCallerSession) MarketCloseSwitchTimestamp() (int64, error) {
	return _LimitOrderBook.Contract.MarketCloseSwitchTimestamp(&_LimitOrderBook.CallOpts)
}

// NumberOfDigestsOfTrader is a free data retrieval call binding the contract method 0x0ee27e34.
//
// Solidity: function numberOfDigestsOfTrader(address trader) view returns(uint256)
func (_LimitOrderBook *LimitOrderBookCaller) NumberOfDigestsOfTrader(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "numberOfDigestsOfTrader", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfDigestsOfTrader is a free data retrieval call binding the contract method 0x0ee27e34.
//
// Solidity: function numberOfDigestsOfTrader(address trader) view returns(uint256)
func (_LimitOrderBook *LimitOrderBookSession) NumberOfDigestsOfTrader(trader common.Address) (*big.Int, error) {
	return _LimitOrderBook.Contract.NumberOfDigestsOfTrader(&_LimitOrderBook.CallOpts, trader)
}

// NumberOfDigestsOfTrader is a free data retrieval call binding the contract method 0x0ee27e34.
//
// Solidity: function numberOfDigestsOfTrader(address trader) view returns(uint256)
func (_LimitOrderBook *LimitOrderBookCallerSession) NumberOfDigestsOfTrader(trader common.Address) (*big.Int, error) {
	return _LimitOrderBook.Contract.NumberOfDigestsOfTrader(&_LimitOrderBook.CallOpts, trader)
}

// OrderCount is a free data retrieval call binding the contract method 0x2453ffa8.
//
// Solidity: function orderCount() view returns(uint256)
func (_LimitOrderBook *LimitOrderBookCaller) OrderCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "orderCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderCount is a free data retrieval call binding the contract method 0x2453ffa8.
//
// Solidity: function orderCount() view returns(uint256)
func (_LimitOrderBook *LimitOrderBookSession) OrderCount() (*big.Int, error) {
	return _LimitOrderBook.Contract.OrderCount(&_LimitOrderBook.CallOpts)
}

// OrderCount is a free data retrieval call binding the contract method 0x2453ffa8.
//
// Solidity: function orderCount() view returns(uint256)
func (_LimitOrderBook *LimitOrderBookCallerSession) OrderCount() (*big.Int, error) {
	return _LimitOrderBook.Contract.OrderCount(&_LimitOrderBook.CallOpts)
}

// OrderDependency is a free data retrieval call binding the contract method 0x1e8dd78a.
//
// Solidity: function orderDependency(bytes32 ) view returns(bytes32 parentChildDigest1, bytes32 parentChildDigest2)
func (_LimitOrderBook *LimitOrderBookCaller) OrderDependency(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ParentChildDigest1 [32]byte
	ParentChildDigest2 [32]byte
}, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "orderDependency", arg0)

	outstruct := new(struct {
		ParentChildDigest1 [32]byte
		ParentChildDigest2 [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ParentChildDigest1 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ParentChildDigest2 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// OrderDependency is a free data retrieval call binding the contract method 0x1e8dd78a.
//
// Solidity: function orderDependency(bytes32 ) view returns(bytes32 parentChildDigest1, bytes32 parentChildDigest2)
func (_LimitOrderBook *LimitOrderBookSession) OrderDependency(arg0 [32]byte) (struct {
	ParentChildDigest1 [32]byte
	ParentChildDigest2 [32]byte
}, error) {
	return _LimitOrderBook.Contract.OrderDependency(&_LimitOrderBook.CallOpts, arg0)
}

// OrderDependency is a free data retrieval call binding the contract method 0x1e8dd78a.
//
// Solidity: function orderDependency(bytes32 ) view returns(bytes32 parentChildDigest1, bytes32 parentChildDigest2)
func (_LimitOrderBook *LimitOrderBookCallerSession) OrderDependency(arg0 [32]byte) (struct {
	ParentChildDigest1 [32]byte
	ParentChildDigest2 [32]byte
}, error) {
	return _LimitOrderBook.Contract.OrderDependency(&_LimitOrderBook.CallOpts, arg0)
}

// OrderOfDigest is a free data retrieval call binding the contract method 0x9023dc5b.
//
// Solidity: function orderOfDigest(bytes32 ) view returns(uint16 leverageTDR, uint16 brokerFeeTbps, uint24 iPerpetualId, address traderAddr, uint32 executionTimestamp, address brokerAddr, uint32 submittedTimestamp, uint32 flags, uint32 iDeadline, address executorAddr, int128 fAmount, int128 fLimitPrice, int128 fTriggerPrice, bytes brokerSignature)
func (_LimitOrderBook *LimitOrderBookCaller) OrderOfDigest(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LeverageTDR        uint16
	BrokerFeeTbps      uint16
	IPerpetualId       *big.Int
	TraderAddr         common.Address
	ExecutionTimestamp uint32
	BrokerAddr         common.Address
	SubmittedTimestamp uint32
	Flags              uint32
	IDeadline          uint32
	ExecutorAddr       common.Address
	FAmount            *big.Int
	FLimitPrice        *big.Int
	FTriggerPrice      *big.Int
	BrokerSignature    []byte
}, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "orderOfDigest", arg0)

	outstruct := new(struct {
		LeverageTDR        uint16
		BrokerFeeTbps      uint16
		IPerpetualId       *big.Int
		TraderAddr         common.Address
		ExecutionTimestamp uint32
		BrokerAddr         common.Address
		SubmittedTimestamp uint32
		Flags              uint32
		IDeadline          uint32
		ExecutorAddr       common.Address
		FAmount            *big.Int
		FLimitPrice        *big.Int
		FTriggerPrice      *big.Int
		BrokerSignature    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LeverageTDR = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.BrokerFeeTbps = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.IPerpetualId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TraderAddr = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ExecutionTimestamp = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.BrokerAddr = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.SubmittedTimestamp = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Flags = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.IDeadline = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.ExecutorAddr = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.FAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.FLimitPrice = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.FTriggerPrice = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.BrokerSignature = *abi.ConvertType(out[13], new([]byte)).(*[]byte)

	return *outstruct, err

}

// OrderOfDigest is a free data retrieval call binding the contract method 0x9023dc5b.
//
// Solidity: function orderOfDigest(bytes32 ) view returns(uint16 leverageTDR, uint16 brokerFeeTbps, uint24 iPerpetualId, address traderAddr, uint32 executionTimestamp, address brokerAddr, uint32 submittedTimestamp, uint32 flags, uint32 iDeadline, address executorAddr, int128 fAmount, int128 fLimitPrice, int128 fTriggerPrice, bytes brokerSignature)
func (_LimitOrderBook *LimitOrderBookSession) OrderOfDigest(arg0 [32]byte) (struct {
	LeverageTDR        uint16
	BrokerFeeTbps      uint16
	IPerpetualId       *big.Int
	TraderAddr         common.Address
	ExecutionTimestamp uint32
	BrokerAddr         common.Address
	SubmittedTimestamp uint32
	Flags              uint32
	IDeadline          uint32
	ExecutorAddr       common.Address
	FAmount            *big.Int
	FLimitPrice        *big.Int
	FTriggerPrice      *big.Int
	BrokerSignature    []byte
}, error) {
	return _LimitOrderBook.Contract.OrderOfDigest(&_LimitOrderBook.CallOpts, arg0)
}

// OrderOfDigest is a free data retrieval call binding the contract method 0x9023dc5b.
//
// Solidity: function orderOfDigest(bytes32 ) view returns(uint16 leverageTDR, uint16 brokerFeeTbps, uint24 iPerpetualId, address traderAddr, uint32 executionTimestamp, address brokerAddr, uint32 submittedTimestamp, uint32 flags, uint32 iDeadline, address executorAddr, int128 fAmount, int128 fLimitPrice, int128 fTriggerPrice, bytes brokerSignature)
func (_LimitOrderBook *LimitOrderBookCallerSession) OrderOfDigest(arg0 [32]byte) (struct {
	LeverageTDR        uint16
	BrokerFeeTbps      uint16
	IPerpetualId       *big.Int
	TraderAddr         common.Address
	ExecutionTimestamp uint32
	BrokerAddr         common.Address
	SubmittedTimestamp uint32
	Flags              uint32
	IDeadline          uint32
	ExecutorAddr       common.Address
	FAmount            *big.Int
	FLimitPrice        *big.Int
	FTriggerPrice      *big.Int
	BrokerSignature    []byte
}, error) {
	return _LimitOrderBook.Contract.OrderOfDigest(&_LimitOrderBook.CallOpts, arg0)
}

// PerpManager is a free data retrieval call binding the contract method 0x99578534.
//
// Solidity: function perpManager() view returns(address)
func (_LimitOrderBook *LimitOrderBookCaller) PerpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "perpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PerpManager is a free data retrieval call binding the contract method 0x99578534.
//
// Solidity: function perpManager() view returns(address)
func (_LimitOrderBook *LimitOrderBookSession) PerpManager() (common.Address, error) {
	return _LimitOrderBook.Contract.PerpManager(&_LimitOrderBook.CallOpts)
}

// PerpManager is a free data retrieval call binding the contract method 0x99578534.
//
// Solidity: function perpManager() view returns(address)
func (_LimitOrderBook *LimitOrderBookCallerSession) PerpManager() (common.Address, error) {
	return _LimitOrderBook.Contract.PerpManager(&_LimitOrderBook.CallOpts)
}

// PerpetualId is a free data retrieval call binding the contract method 0x20ef8157.
//
// Solidity: function perpetualId() view returns(uint24)
func (_LimitOrderBook *LimitOrderBookCaller) PerpetualId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "perpetualId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerpetualId is a free data retrieval call binding the contract method 0x20ef8157.
//
// Solidity: function perpetualId() view returns(uint24)
func (_LimitOrderBook *LimitOrderBookSession) PerpetualId() (*big.Int, error) {
	return _LimitOrderBook.Contract.PerpetualId(&_LimitOrderBook.CallOpts)
}

// PerpetualId is a free data retrieval call binding the contract method 0x20ef8157.
//
// Solidity: function perpetualId() view returns(uint24)
func (_LimitOrderBook *LimitOrderBookCallerSession) PerpetualId() (*big.Int, error) {
	return _LimitOrderBook.Contract.PerpetualId(&_LimitOrderBook.CallOpts)
}

// PollRange is a free data retrieval call binding the contract method 0x61fe71c6.
//
// Solidity: function pollRange(uint256 _from, uint256 _numElements) view returns((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] orders, bytes32[] orderHashes, uint32[] submittedTs)
func (_LimitOrderBook *LimitOrderBookCaller) PollRange(opts *bind.CallOpts, _from *big.Int, _numElements *big.Int) (struct {
	Orders      []IClientOrderClientOrder
	OrderHashes [][32]byte
	SubmittedTs []uint32
}, error) {
	var out []interface{}
	err := _LimitOrderBook.contract.Call(opts, &out, "pollRange", _from, _numElements)

	outstruct := new(struct {
		Orders      []IClientOrderClientOrder
		OrderHashes [][32]byte
		SubmittedTs []uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Orders = *abi.ConvertType(out[0], new([]IClientOrderClientOrder)).(*[]IClientOrderClientOrder)
	outstruct.OrderHashes = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.SubmittedTs = *abi.ConvertType(out[2], new([]uint32)).(*[]uint32)

	return *outstruct, err

}

// PollRange is a free data retrieval call binding the contract method 0x61fe71c6.
//
// Solidity: function pollRange(uint256 _from, uint256 _numElements) view returns((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] orders, bytes32[] orderHashes, uint32[] submittedTs)
func (_LimitOrderBook *LimitOrderBookSession) PollRange(_from *big.Int, _numElements *big.Int) (struct {
	Orders      []IClientOrderClientOrder
	OrderHashes [][32]byte
	SubmittedTs []uint32
}, error) {
	return _LimitOrderBook.Contract.PollRange(&_LimitOrderBook.CallOpts, _from, _numElements)
}

// PollRange is a free data retrieval call binding the contract method 0x61fe71c6.
//
// Solidity: function pollRange(uint256 _from, uint256 _numElements) view returns((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] orders, bytes32[] orderHashes, uint32[] submittedTs)
func (_LimitOrderBook *LimitOrderBookCallerSession) PollRange(_from *big.Int, _numElements *big.Int) (struct {
	Orders      []IClientOrderClientOrder
	OrderHashes [][32]byte
	SubmittedTs []uint32
}, error) {
	return _LimitOrderBook.Contract.PollRange(&_LimitOrderBook.CallOpts, _from, _numElements)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xb67f7613.
//
// Solidity: function cancelOrder(bytes32 _digest, bytes _signature, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_LimitOrderBook *LimitOrderBookTransactor) CancelOrder(opts *bind.TransactOpts, _digest [32]byte, _signature []byte, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _LimitOrderBook.contract.Transact(opts, "cancelOrder", _digest, _signature, _updateData, _publishTimes)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xb67f7613.
//
// Solidity: function cancelOrder(bytes32 _digest, bytes _signature, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_LimitOrderBook *LimitOrderBookSession) CancelOrder(_digest [32]byte, _signature []byte, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.CancelOrder(&_LimitOrderBook.TransactOpts, _digest, _signature, _updateData, _publishTimes)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xb67f7613.
//
// Solidity: function cancelOrder(bytes32 _digest, bytes _signature, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_LimitOrderBook *LimitOrderBookTransactorSession) CancelOrder(_digest [32]byte, _signature []byte, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.CancelOrder(&_LimitOrderBook.TransactOpts, _digest, _signature, _updateData, _publishTimes)
}

// ExecuteOrders is a paid mutator transaction binding the contract method 0xe905f3d1.
//
// Solidity: function executeOrders(bytes32[] _digests, address _executorAddr, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_LimitOrderBook *LimitOrderBookTransactor) ExecuteOrders(opts *bind.TransactOpts, _digests [][32]byte, _executorAddr common.Address, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _LimitOrderBook.contract.Transact(opts, "executeOrders", _digests, _executorAddr, _updateData, _publishTimes)
}

// ExecuteOrders is a paid mutator transaction binding the contract method 0xe905f3d1.
//
// Solidity: function executeOrders(bytes32[] _digests, address _executorAddr, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_LimitOrderBook *LimitOrderBookSession) ExecuteOrders(_digests [][32]byte, _executorAddr common.Address, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.ExecuteOrders(&_LimitOrderBook.TransactOpts, _digests, _executorAddr, _updateData, _publishTimes)
}

// ExecuteOrders is a paid mutator transaction binding the contract method 0xe905f3d1.
//
// Solidity: function executeOrders(bytes32[] _digests, address _executorAddr, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_LimitOrderBook *LimitOrderBookTransactorSession) ExecuteOrders(_digests [][32]byte, _executorAddr common.Address, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.ExecuteOrders(&_LimitOrderBook.TransactOpts, _digests, _executorAddr, _updateData, _publishTimes)
}

// Initialize is a paid mutator transaction binding the contract method 0x749fb279.
//
// Solidity: function initialize(address _perpetualManagerAddr, uint24 _perpetualId, uint8 _iCancelDelaySec) returns()
func (_LimitOrderBook *LimitOrderBookTransactor) Initialize(opts *bind.TransactOpts, _perpetualManagerAddr common.Address, _perpetualId *big.Int, _iCancelDelaySec uint8) (*types.Transaction, error) {
	return _LimitOrderBook.contract.Transact(opts, "initialize", _perpetualManagerAddr, _perpetualId, _iCancelDelaySec)
}

// Initialize is a paid mutator transaction binding the contract method 0x749fb279.
//
// Solidity: function initialize(address _perpetualManagerAddr, uint24 _perpetualId, uint8 _iCancelDelaySec) returns()
func (_LimitOrderBook *LimitOrderBookSession) Initialize(_perpetualManagerAddr common.Address, _perpetualId *big.Int, _iCancelDelaySec uint8) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.Initialize(&_LimitOrderBook.TransactOpts, _perpetualManagerAddr, _perpetualId, _iCancelDelaySec)
}

// Initialize is a paid mutator transaction binding the contract method 0x749fb279.
//
// Solidity: function initialize(address _perpetualManagerAddr, uint24 _perpetualId, uint8 _iCancelDelaySec) returns()
func (_LimitOrderBook *LimitOrderBookTransactorSession) Initialize(_perpetualManagerAddr common.Address, _perpetualId *big.Int, _iCancelDelaySec uint8) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.Initialize(&_LimitOrderBook.TransactOpts, _perpetualManagerAddr, _perpetualId, _iCancelDelaySec)
}

// PostOrders is a paid mutator transaction binding the contract method 0xa9a5e49c.
//
// Solidity: function postOrders((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] _orders, bytes[] _signatures) returns()
func (_LimitOrderBook *LimitOrderBookTransactor) PostOrders(opts *bind.TransactOpts, _orders []IClientOrderClientOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _LimitOrderBook.contract.Transact(opts, "postOrders", _orders, _signatures)
}

// PostOrders is a paid mutator transaction binding the contract method 0xa9a5e49c.
//
// Solidity: function postOrders((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] _orders, bytes[] _signatures) returns()
func (_LimitOrderBook *LimitOrderBookSession) PostOrders(_orders []IClientOrderClientOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.PostOrders(&_LimitOrderBook.TransactOpts, _orders, _signatures)
}

// PostOrders is a paid mutator transaction binding the contract method 0xa9a5e49c.
//
// Solidity: function postOrders((uint24,int128,uint16,uint32,uint32,uint32,address,int128,int128,bytes32,address,bytes32,uint16,bytes,address)[] _orders, bytes[] _signatures) returns()
func (_LimitOrderBook *LimitOrderBookTransactorSession) PostOrders(_orders []IClientOrderClientOrder, _signatures [][]byte) (*types.Transaction, error) {
	return _LimitOrderBook.Contract.PostOrders(&_LimitOrderBook.TransactOpts, _orders, _signatures)
}

// LimitOrderBookCallbackIterator is returned from FilterCallback and is used to iterate over the raw logs and unpacked data for Callback events raised by the LimitOrderBook contract.
type LimitOrderBookCallbackIterator struct {
	Event *LimitOrderBookCallback // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookCallback)
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
		it.Event = new(LimitOrderBookCallback)
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
func (it *LimitOrderBookCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookCallback represents a Callback event raised by the LimitOrderBook contract.
type LimitOrderBookCallback struct {
	CallbackTarget common.Address
	Success        bool
	GasLimit       uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCallback is a free log retrieval operation binding the contract event 0x1d4b3142757d5200ad5437c0348f6047f6f312b4ccadda70ecf72cd93ea1e260.
//
// Solidity: event Callback(address callbackTarget, bool success, uint32 gasLimit)
func (_LimitOrderBook *LimitOrderBookFilterer) FilterCallback(opts *bind.FilterOpts) (*LimitOrderBookCallbackIterator, error) {

	logs, sub, err := _LimitOrderBook.contract.FilterLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookCallbackIterator{contract: _LimitOrderBook.contract, event: "Callback", logs: logs, sub: sub}, nil
}

// WatchCallback is a free log subscription operation binding the contract event 0x1d4b3142757d5200ad5437c0348f6047f6f312b4ccadda70ecf72cd93ea1e260.
//
// Solidity: event Callback(address callbackTarget, bool success, uint32 gasLimit)
func (_LimitOrderBook *LimitOrderBookFilterer) WatchCallback(opts *bind.WatchOpts, sink chan<- *LimitOrderBookCallback) (event.Subscription, error) {

	logs, sub, err := _LimitOrderBook.contract.WatchLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookCallback)
				if err := _LimitOrderBook.contract.UnpackLog(event, "Callback", log); err != nil {
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

// ParseCallback is a log parse operation binding the contract event 0x1d4b3142757d5200ad5437c0348f6047f6f312b4ccadda70ecf72cd93ea1e260.
//
// Solidity: event Callback(address callbackTarget, bool success, uint32 gasLimit)
func (_LimitOrderBook *LimitOrderBookFilterer) ParseCallback(log types.Log) (*LimitOrderBookCallback, error) {
	event := new(LimitOrderBookCallback)
	if err := _LimitOrderBook.contract.UnpackLog(event, "Callback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderBookExecutionFailedIterator is returned from FilterExecutionFailed and is used to iterate over the raw logs and unpacked data for ExecutionFailed events raised by the LimitOrderBook contract.
type LimitOrderBookExecutionFailedIterator struct {
	Event *LimitOrderBookExecutionFailed // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookExecutionFailed)
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
		it.Event = new(LimitOrderBookExecutionFailed)
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
func (it *LimitOrderBookExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookExecutionFailed represents a ExecutionFailed event raised by the LimitOrderBook contract.
type LimitOrderBookExecutionFailed struct {
	PerpetualId *big.Int
	Trader      common.Address
	Digest      [32]byte
	Reason      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailed is a free log retrieval operation binding the contract event 0xaabfe27de8303a4c68aa0a50788ef1207117ec13d2fba01e5b1c0750705988a0.
//
// Solidity: event ExecutionFailed(uint24 indexed perpetualId, address indexed trader, bytes32 digest, string reason)
func (_LimitOrderBook *LimitOrderBookFilterer) FilterExecutionFailed(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address) (*LimitOrderBookExecutionFailedIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LimitOrderBook.contract.FilterLogs(opts, "ExecutionFailed", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookExecutionFailedIterator{contract: _LimitOrderBook.contract, event: "ExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchExecutionFailed is a free log subscription operation binding the contract event 0xaabfe27de8303a4c68aa0a50788ef1207117ec13d2fba01e5b1c0750705988a0.
//
// Solidity: event ExecutionFailed(uint24 indexed perpetualId, address indexed trader, bytes32 digest, string reason)
func (_LimitOrderBook *LimitOrderBookFilterer) WatchExecutionFailed(opts *bind.WatchOpts, sink chan<- *LimitOrderBookExecutionFailed, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LimitOrderBook.contract.WatchLogs(opts, "ExecutionFailed", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookExecutionFailed)
				if err := _LimitOrderBook.contract.UnpackLog(event, "ExecutionFailed", log); err != nil {
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

// ParseExecutionFailed is a log parse operation binding the contract event 0xaabfe27de8303a4c68aa0a50788ef1207117ec13d2fba01e5b1c0750705988a0.
//
// Solidity: event ExecutionFailed(uint24 indexed perpetualId, address indexed trader, bytes32 digest, string reason)
func (_LimitOrderBook *LimitOrderBookFilterer) ParseExecutionFailed(log types.Log) (*LimitOrderBookExecutionFailed, error) {
	event := new(LimitOrderBookExecutionFailed)
	if err := _LimitOrderBook.contract.UnpackLog(event, "ExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderBookInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LimitOrderBook contract.
type LimitOrderBookInitializedIterator struct {
	Event *LimitOrderBookInitialized // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookInitialized)
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
		it.Event = new(LimitOrderBookInitialized)
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
func (it *LimitOrderBookInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookInitialized represents a Initialized event raised by the LimitOrderBook contract.
type LimitOrderBookInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LimitOrderBook *LimitOrderBookFilterer) FilterInitialized(opts *bind.FilterOpts) (*LimitOrderBookInitializedIterator, error) {

	logs, sub, err := _LimitOrderBook.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookInitializedIterator{contract: _LimitOrderBook.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LimitOrderBook *LimitOrderBookFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LimitOrderBookInitialized) (event.Subscription, error) {

	logs, sub, err := _LimitOrderBook.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookInitialized)
				if err := _LimitOrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LimitOrderBook *LimitOrderBookFilterer) ParseInitialized(log types.Log) (*LimitOrderBookInitialized, error) {
	event := new(LimitOrderBookInitialized)
	if err := _LimitOrderBook.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderBookPerpetualLimitOrderCreatedIterator is returned from FilterPerpetualLimitOrderCreated and is used to iterate over the raw logs and unpacked data for PerpetualLimitOrderCreated events raised by the LimitOrderBook contract.
type LimitOrderBookPerpetualLimitOrderCreatedIterator struct {
	Event *LimitOrderBookPerpetualLimitOrderCreated // Event containing the contract specifics and raw log

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
func (it *LimitOrderBookPerpetualLimitOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderBookPerpetualLimitOrderCreated)
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
		it.Event = new(LimitOrderBookPerpetualLimitOrderCreated)
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
func (it *LimitOrderBookPerpetualLimitOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderBookPerpetualLimitOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderBookPerpetualLimitOrderCreated represents a PerpetualLimitOrderCreated event raised by the LimitOrderBook contract.
type LimitOrderBookPerpetualLimitOrderCreated struct {
	PerpetualId *big.Int
	Trader      common.Address
	BrokerAddr  common.Address
	Order       IPerpetualOrderOrder
	Digest      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPerpetualLimitOrderCreated is a free log retrieval operation binding the contract event 0xb8ace0e652054ba9160467efc4bd026f89b793c887035f7d88d7b7b0615c4c5f.
//
// Solidity: event PerpetualLimitOrderCreated(uint24 indexed perpetualId, address indexed trader, address brokerAddr, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) order, bytes32 digest)
func (_LimitOrderBook *LimitOrderBookFilterer) FilterPerpetualLimitOrderCreated(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address) (*LimitOrderBookPerpetualLimitOrderCreatedIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LimitOrderBook.contract.FilterLogs(opts, "PerpetualLimitOrderCreated", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderBookPerpetualLimitOrderCreatedIterator{contract: _LimitOrderBook.contract, event: "PerpetualLimitOrderCreated", logs: logs, sub: sub}, nil
}

// WatchPerpetualLimitOrderCreated is a free log subscription operation binding the contract event 0xb8ace0e652054ba9160467efc4bd026f89b793c887035f7d88d7b7b0615c4c5f.
//
// Solidity: event PerpetualLimitOrderCreated(uint24 indexed perpetualId, address indexed trader, address brokerAddr, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) order, bytes32 digest)
func (_LimitOrderBook *LimitOrderBookFilterer) WatchPerpetualLimitOrderCreated(opts *bind.WatchOpts, sink chan<- *LimitOrderBookPerpetualLimitOrderCreated, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _LimitOrderBook.contract.WatchLogs(opts, "PerpetualLimitOrderCreated", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderBookPerpetualLimitOrderCreated)
				if err := _LimitOrderBook.contract.UnpackLog(event, "PerpetualLimitOrderCreated", log); err != nil {
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

// ParsePerpetualLimitOrderCreated is a log parse operation binding the contract event 0xb8ace0e652054ba9160467efc4bd026f89b793c887035f7d88d7b7b0615c4c5f.
//
// Solidity: event PerpetualLimitOrderCreated(uint24 indexed perpetualId, address indexed trader, address brokerAddr, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) order, bytes32 digest)
func (_LimitOrderBook *LimitOrderBookFilterer) ParsePerpetualLimitOrderCreated(log types.Log) (*LimitOrderBookPerpetualLimitOrderCreated, error) {
	event := new(LimitOrderBookPerpetualLimitOrderCreated)
	if err := _LimitOrderBook.contract.UnpackLog(event, "PerpetualLimitOrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
