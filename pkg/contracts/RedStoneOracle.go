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

// RedStoneOracleMetaData contains all meta data concerning the RedStoneOracle contract.
var RedStoneOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"BlockTimestampIsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalldataMustHaveValidPayload\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalldataOverOrUnderFlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanNotPickMedianOfEmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateMoreThanOneDataFeed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"}],\"name\":\"DataFeedIdNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"}],\"name\":\"DataFeedValueCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedDataTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataPackageTimestamp\",\"type\":\"uint256\"}],\"name\":\"DataPackageTimestampMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataPackageTimestampMustNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataPackageTimestampsMustBeEqual\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataTimestamp\",\"type\":\"uint256\"}],\"name\":\"DataTimestampIsTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedDataTimestampMilliseconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDataTimestampMilliseconds\",\"type\":\"uint256\"}],\"name\":\"DataTimestampShouldBeNewerThanBefore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EachSignerMustProvideTheSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyCalldataPointersArr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetDataServiceIdNotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectUnsignedMetadataSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedSignersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignersCount\",\"type\":\"uint256\"}],\"name\":\"InsufficientNumberOfUniqueSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldataPointer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentBlockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIntervalBetweenUpdates\",\"type\":\"uint256\"}],\"name\":\"MinIntervalBetweenUpdatesHasNotPassedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedstonePayloadMustHaveAtLeastOneDataPackage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"RoundNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receivedSigner\",\"type\":\"address\"}],\"name\":\"SignerNotAuthorised\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedTimestampSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"TimestampFromTooLongFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedTimestampSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"TimestampIsTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UnsafeUint256ToUint80Conversion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UnsafeUintToIntConversion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"UpdaterNotAuthorised\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"aggregateValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extractTimestampsAndAssertAllAreEqual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"extractedTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowedTimestampDiffsInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDataAheadSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataDelaySeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"getAuthorisedSignerIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTimestampFromLatestUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataFeedId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataFeedIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataFeedIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"}],\"name\":\"getDataFeedIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataServiceId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataTimestampFromLatestUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastDataTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"latestRoundId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestRoundParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"latestRoundId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"latestRoundDataTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"latestRoundBlockTimestamp\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinIntervalBetweenUpdates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getPackedTimestampsForRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundTimestamps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPackedTimestampsFromLatestUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"packedTimestamps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeedAdapter\",\"outputs\":[{\"internalType\":\"contractIRedstoneAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeedAdapterWithRounds\",\"outputs\":[{\"internalType\":\"contractPriceFeedsAdapterWithRounds\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"requestedRoundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getRoundDataFromAdapter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dataFeedValue\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"roundDataTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"roundBlockTimestamp\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimestampsFromLatestUpdate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"dataTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"blockTimestamp\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUniqueSignersThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"}],\"name\":\"getValueForDataFeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getValueForDataFeedAndRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dataFeedValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"}],\"name\":\"getValueForDataFeedUnsafe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dataFeedValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"dataFeedIds\",\"type\":\"bytes32[]\"}],\"name\":\"getValuesForDataFeeds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"requestedDataFeedIds\",\"type\":\"bytes32[]\"}],\"name\":\"getValuesForDataFeedsUnsafe\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"}],\"name\":\"requireAuthorisedUpdater\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataPackagesTimestamp\",\"type\":\"uint256\"}],\"name\":\"updateDataFeedsValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueForDataFeed\",\"type\":\"uint256\"}],\"name\":\"validateDataFeedValueOnRead\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueForDataFeed\",\"type\":\"uint256\"}],\"name\":\"validateDataFeedValueOnWrite\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataPackagesTimestamp\",\"type\":\"uint256\"}],\"name\":\"validateDataPackagesTimestampOnce\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataPackagesTimestamp\",\"type\":\"uint256\"}],\"name\":\"validateProposedDataPackagesTimestamp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedTimestampMilliseconds\",\"type\":\"uint256\"}],\"name\":\"validateTimestamp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// RedStoneOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use RedStoneOracleMetaData.ABI instead.
var RedStoneOracleABI = RedStoneOracleMetaData.ABI

// RedStoneOracle is an auto generated Go binding around an Ethereum contract.
type RedStoneOracle struct {
	RedStoneOracleCaller     // Read-only binding to the contract
	RedStoneOracleTransactor // Write-only binding to the contract
	RedStoneOracleFilterer   // Log filterer for contract events
}

// RedStoneOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedStoneOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedStoneOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedStoneOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedStoneOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedStoneOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedStoneOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedStoneOracleSession struct {
	Contract     *RedStoneOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedStoneOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedStoneOracleCallerSession struct {
	Contract *RedStoneOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RedStoneOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedStoneOracleTransactorSession struct {
	Contract     *RedStoneOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RedStoneOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedStoneOracleRaw struct {
	Contract *RedStoneOracle // Generic contract binding to access the raw methods on
}

// RedStoneOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedStoneOracleCallerRaw struct {
	Contract *RedStoneOracleCaller // Generic read-only contract binding to access the raw methods on
}

// RedStoneOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedStoneOracleTransactorRaw struct {
	Contract *RedStoneOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedStoneOracle creates a new instance of RedStoneOracle, bound to a specific deployed contract.
func NewRedStoneOracle(address common.Address, backend bind.ContractBackend) (*RedStoneOracle, error) {
	contract, err := bindRedStoneOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedStoneOracle{RedStoneOracleCaller: RedStoneOracleCaller{contract: contract}, RedStoneOracleTransactor: RedStoneOracleTransactor{contract: contract}, RedStoneOracleFilterer: RedStoneOracleFilterer{contract: contract}}, nil
}

// NewRedStoneOracleCaller creates a new read-only instance of RedStoneOracle, bound to a specific deployed contract.
func NewRedStoneOracleCaller(address common.Address, caller bind.ContractCaller) (*RedStoneOracleCaller, error) {
	contract, err := bindRedStoneOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedStoneOracleCaller{contract: contract}, nil
}

// NewRedStoneOracleTransactor creates a new write-only instance of RedStoneOracle, bound to a specific deployed contract.
func NewRedStoneOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*RedStoneOracleTransactor, error) {
	contract, err := bindRedStoneOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedStoneOracleTransactor{contract: contract}, nil
}

// NewRedStoneOracleFilterer creates a new log filterer instance of RedStoneOracle, bound to a specific deployed contract.
func NewRedStoneOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*RedStoneOracleFilterer, error) {
	contract, err := bindRedStoneOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedStoneOracleFilterer{contract: contract}, nil
}

// bindRedStoneOracle binds a generic wrapper to an already deployed contract.
func bindRedStoneOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RedStoneOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedStoneOracle *RedStoneOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedStoneOracle.Contract.RedStoneOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedStoneOracle *RedStoneOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedStoneOracle.Contract.RedStoneOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedStoneOracle *RedStoneOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedStoneOracle.Contract.RedStoneOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedStoneOracle *RedStoneOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedStoneOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedStoneOracle *RedStoneOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedStoneOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedStoneOracle *RedStoneOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedStoneOracle.Contract.contract.Transact(opts, method, params...)
}

// AggregateValues is a free data retrieval call binding the contract method 0xb24ebfcc.
//
// Solidity: function aggregateValues(uint256[] values) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCaller) AggregateValues(opts *bind.CallOpts, values []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "aggregateValues", values)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregateValues is a free data retrieval call binding the contract method 0xb24ebfcc.
//
// Solidity: function aggregateValues(uint256[] values) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleSession) AggregateValues(values []*big.Int) (*big.Int, error) {
	return _RedStoneOracle.Contract.AggregateValues(&_RedStoneOracle.CallOpts, values)
}

// AggregateValues is a free data retrieval call binding the contract method 0xb24ebfcc.
//
// Solidity: function aggregateValues(uint256[] values) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCallerSession) AggregateValues(values []*big.Int) (*big.Int, error) {
	return _RedStoneOracle.Contract.AggregateValues(&_RedStoneOracle.CallOpts, values)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_RedStoneOracle *RedStoneOracleCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_RedStoneOracle *RedStoneOracleSession) Aggregator() (common.Address, error) {
	return _RedStoneOracle.Contract.Aggregator(&_RedStoneOracle.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_RedStoneOracle *RedStoneOracleCallerSession) Aggregator() (common.Address, error) {
	return _RedStoneOracle.Contract.Aggregator(&_RedStoneOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_RedStoneOracle *RedStoneOracleCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_RedStoneOracle *RedStoneOracleSession) Decimals() (uint8, error) {
	return _RedStoneOracle.Contract.Decimals(&_RedStoneOracle.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_RedStoneOracle *RedStoneOracleCallerSession) Decimals() (uint8, error) {
	return _RedStoneOracle.Contract.Decimals(&_RedStoneOracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_RedStoneOracle *RedStoneOracleCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_RedStoneOracle *RedStoneOracleSession) Description() (string, error) {
	return _RedStoneOracle.Contract.Description(&_RedStoneOracle.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_RedStoneOracle *RedStoneOracleCallerSession) Description() (string, error) {
	return _RedStoneOracle.Contract.Description(&_RedStoneOracle.CallOpts)
}

// ExtractTimestampsAndAssertAllAreEqual is a free data retrieval call binding the contract method 0x55a547d5.
//
// Solidity: function extractTimestampsAndAssertAllAreEqual() pure returns(uint256 extractedTimestamp)
func (_RedStoneOracle *RedStoneOracleCaller) ExtractTimestampsAndAssertAllAreEqual(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "extractTimestampsAndAssertAllAreEqual")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtractTimestampsAndAssertAllAreEqual is a free data retrieval call binding the contract method 0x55a547d5.
//
// Solidity: function extractTimestampsAndAssertAllAreEqual() pure returns(uint256 extractedTimestamp)
func (_RedStoneOracle *RedStoneOracleSession) ExtractTimestampsAndAssertAllAreEqual() (*big.Int, error) {
	return _RedStoneOracle.Contract.ExtractTimestampsAndAssertAllAreEqual(&_RedStoneOracle.CallOpts)
}

// ExtractTimestampsAndAssertAllAreEqual is a free data retrieval call binding the contract method 0x55a547d5.
//
// Solidity: function extractTimestampsAndAssertAllAreEqual() pure returns(uint256 extractedTimestamp)
func (_RedStoneOracle *RedStoneOracleCallerSession) ExtractTimestampsAndAssertAllAreEqual() (*big.Int, error) {
	return _RedStoneOracle.Contract.ExtractTimestampsAndAssertAllAreEqual(&_RedStoneOracle.CallOpts)
}

// GetAllowedTimestampDiffsInSeconds is a free data retrieval call binding the contract method 0xaef2f165.
//
// Solidity: function getAllowedTimestampDiffsInSeconds() view returns(uint256 maxDataAheadSeconds, uint256 maxDataDelaySeconds)
func (_RedStoneOracle *RedStoneOracleCaller) GetAllowedTimestampDiffsInSeconds(opts *bind.CallOpts) (struct {
	MaxDataAheadSeconds *big.Int
	MaxDataDelaySeconds *big.Int
}, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getAllowedTimestampDiffsInSeconds")

	outstruct := new(struct {
		MaxDataAheadSeconds *big.Int
		MaxDataDelaySeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxDataAheadSeconds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxDataDelaySeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAllowedTimestampDiffsInSeconds is a free data retrieval call binding the contract method 0xaef2f165.
//
// Solidity: function getAllowedTimestampDiffsInSeconds() view returns(uint256 maxDataAheadSeconds, uint256 maxDataDelaySeconds)
func (_RedStoneOracle *RedStoneOracleSession) GetAllowedTimestampDiffsInSeconds() (struct {
	MaxDataAheadSeconds *big.Int
	MaxDataDelaySeconds *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetAllowedTimestampDiffsInSeconds(&_RedStoneOracle.CallOpts)
}

// GetAllowedTimestampDiffsInSeconds is a free data retrieval call binding the contract method 0xaef2f165.
//
// Solidity: function getAllowedTimestampDiffsInSeconds() view returns(uint256 maxDataAheadSeconds, uint256 maxDataDelaySeconds)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetAllowedTimestampDiffsInSeconds() (struct {
	MaxDataAheadSeconds *big.Int
	MaxDataDelaySeconds *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetAllowedTimestampDiffsInSeconds(&_RedStoneOracle.CallOpts)
}

// GetAuthorisedSignerIndex is a free data retrieval call binding the contract method 0x3ce142f5.
//
// Solidity: function getAuthorisedSignerIndex(address signerAddress) view returns(uint8)
func (_RedStoneOracle *RedStoneOracleCaller) GetAuthorisedSignerIndex(opts *bind.CallOpts, signerAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getAuthorisedSignerIndex", signerAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAuthorisedSignerIndex is a free data retrieval call binding the contract method 0x3ce142f5.
//
// Solidity: function getAuthorisedSignerIndex(address signerAddress) view returns(uint8)
func (_RedStoneOracle *RedStoneOracleSession) GetAuthorisedSignerIndex(signerAddress common.Address) (uint8, error) {
	return _RedStoneOracle.Contract.GetAuthorisedSignerIndex(&_RedStoneOracle.CallOpts, signerAddress)
}

// GetAuthorisedSignerIndex is a free data retrieval call binding the contract method 0x3ce142f5.
//
// Solidity: function getAuthorisedSignerIndex(address signerAddress) view returns(uint8)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetAuthorisedSignerIndex(signerAddress common.Address) (uint8, error) {
	return _RedStoneOracle.Contract.GetAuthorisedSignerIndex(&_RedStoneOracle.CallOpts, signerAddress)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_RedStoneOracle *RedStoneOracleSession) GetBlockTimestamp() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetBlockTimestamp(&_RedStoneOracle.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetBlockTimestamp(&_RedStoneOracle.CallOpts)
}

// GetBlockTimestampFromLatestUpdate is a free data retrieval call binding the contract method 0x1b2758ee.
//
// Solidity: function getBlockTimestampFromLatestUpdate() view returns(uint256 blockTimestamp)
func (_RedStoneOracle *RedStoneOracleCaller) GetBlockTimestampFromLatestUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getBlockTimestampFromLatestUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestampFromLatestUpdate is a free data retrieval call binding the contract method 0x1b2758ee.
//
// Solidity: function getBlockTimestampFromLatestUpdate() view returns(uint256 blockTimestamp)
func (_RedStoneOracle *RedStoneOracleSession) GetBlockTimestampFromLatestUpdate() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetBlockTimestampFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetBlockTimestampFromLatestUpdate is a free data retrieval call binding the contract method 0x1b2758ee.
//
// Solidity: function getBlockTimestampFromLatestUpdate() view returns(uint256 blockTimestamp)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetBlockTimestampFromLatestUpdate() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetBlockTimestampFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetDataFeedId is a free data retrieval call binding the contract method 0xc8337760.
//
// Solidity: function getDataFeedId() pure returns(bytes32)
func (_RedStoneOracle *RedStoneOracleCaller) GetDataFeedId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getDataFeedId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDataFeedId is a free data retrieval call binding the contract method 0xc8337760.
//
// Solidity: function getDataFeedId() pure returns(bytes32)
func (_RedStoneOracle *RedStoneOracleSession) GetDataFeedId() ([32]byte, error) {
	return _RedStoneOracle.Contract.GetDataFeedId(&_RedStoneOracle.CallOpts)
}

// GetDataFeedId is a free data retrieval call binding the contract method 0xc8337760.
//
// Solidity: function getDataFeedId() pure returns(bytes32)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetDataFeedId() ([32]byte, error) {
	return _RedStoneOracle.Contract.GetDataFeedId(&_RedStoneOracle.CallOpts)
}

// GetDataFeedIds is a free data retrieval call binding the contract method 0xfba03158.
//
// Solidity: function getDataFeedIds() view returns(bytes32[] dataFeedIds)
func (_RedStoneOracle *RedStoneOracleCaller) GetDataFeedIds(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getDataFeedIds")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDataFeedIds is a free data retrieval call binding the contract method 0xfba03158.
//
// Solidity: function getDataFeedIds() view returns(bytes32[] dataFeedIds)
func (_RedStoneOracle *RedStoneOracleSession) GetDataFeedIds() ([][32]byte, error) {
	return _RedStoneOracle.Contract.GetDataFeedIds(&_RedStoneOracle.CallOpts)
}

// GetDataFeedIds is a free data retrieval call binding the contract method 0xfba03158.
//
// Solidity: function getDataFeedIds() view returns(bytes32[] dataFeedIds)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetDataFeedIds() ([][32]byte, error) {
	return _RedStoneOracle.Contract.GetDataFeedIds(&_RedStoneOracle.CallOpts)
}

// GetDataFeedIndex is a free data retrieval call binding the contract method 0x6dafaf6a.
//
// Solidity: function getDataFeedIndex(bytes32 dataFeedId) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCaller) GetDataFeedIndex(opts *bind.CallOpts, dataFeedId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getDataFeedIndex", dataFeedId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDataFeedIndex is a free data retrieval call binding the contract method 0x6dafaf6a.
//
// Solidity: function getDataFeedIndex(bytes32 dataFeedId) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleSession) GetDataFeedIndex(dataFeedId [32]byte) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetDataFeedIndex(&_RedStoneOracle.CallOpts, dataFeedId)
}

// GetDataFeedIndex is a free data retrieval call binding the contract method 0x6dafaf6a.
//
// Solidity: function getDataFeedIndex(bytes32 dataFeedId) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetDataFeedIndex(dataFeedId [32]byte) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetDataFeedIndex(&_RedStoneOracle.CallOpts, dataFeedId)
}

// GetDataServiceId is a free data retrieval call binding the contract method 0xc274583a.
//
// Solidity: function getDataServiceId() view returns(string)
func (_RedStoneOracle *RedStoneOracleCaller) GetDataServiceId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getDataServiceId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDataServiceId is a free data retrieval call binding the contract method 0xc274583a.
//
// Solidity: function getDataServiceId() view returns(string)
func (_RedStoneOracle *RedStoneOracleSession) GetDataServiceId() (string, error) {
	return _RedStoneOracle.Contract.GetDataServiceId(&_RedStoneOracle.CallOpts)
}

// GetDataServiceId is a free data retrieval call binding the contract method 0xc274583a.
//
// Solidity: function getDataServiceId() view returns(string)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetDataServiceId() (string, error) {
	return _RedStoneOracle.Contract.GetDataServiceId(&_RedStoneOracle.CallOpts)
}

// GetDataTimestampFromLatestUpdate is a free data retrieval call binding the contract method 0x7a02bdf1.
//
// Solidity: function getDataTimestampFromLatestUpdate() view returns(uint256 lastDataTimestamp)
func (_RedStoneOracle *RedStoneOracleCaller) GetDataTimestampFromLatestUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getDataTimestampFromLatestUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDataTimestampFromLatestUpdate is a free data retrieval call binding the contract method 0x7a02bdf1.
//
// Solidity: function getDataTimestampFromLatestUpdate() view returns(uint256 lastDataTimestamp)
func (_RedStoneOracle *RedStoneOracleSession) GetDataTimestampFromLatestUpdate() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetDataTimestampFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetDataTimestampFromLatestUpdate is a free data retrieval call binding the contract method 0x7a02bdf1.
//
// Solidity: function getDataTimestampFromLatestUpdate() view returns(uint256 lastDataTimestamp)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetDataTimestampFromLatestUpdate() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetDataTimestampFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetLatestRoundId is a free data retrieval call binding the contract method 0x8c3b990b.
//
// Solidity: function getLatestRoundId() view returns(uint256 latestRoundId)
func (_RedStoneOracle *RedStoneOracleCaller) GetLatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getLatestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestRoundId is a free data retrieval call binding the contract method 0x8c3b990b.
//
// Solidity: function getLatestRoundId() view returns(uint256 latestRoundId)
func (_RedStoneOracle *RedStoneOracleSession) GetLatestRoundId() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetLatestRoundId(&_RedStoneOracle.CallOpts)
}

// GetLatestRoundId is a free data retrieval call binding the contract method 0x8c3b990b.
//
// Solidity: function getLatestRoundId() view returns(uint256 latestRoundId)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetLatestRoundId() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetLatestRoundId(&_RedStoneOracle.CallOpts)
}

// GetLatestRoundParams is a free data retrieval call binding the contract method 0xd1375817.
//
// Solidity: function getLatestRoundParams() view returns(uint256 latestRoundId, uint128 latestRoundDataTimestamp, uint128 latestRoundBlockTimestamp)
func (_RedStoneOracle *RedStoneOracleCaller) GetLatestRoundParams(opts *bind.CallOpts) (struct {
	LatestRoundId             *big.Int
	LatestRoundDataTimestamp  *big.Int
	LatestRoundBlockTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getLatestRoundParams")

	outstruct := new(struct {
		LatestRoundId             *big.Int
		LatestRoundDataTimestamp  *big.Int
		LatestRoundBlockTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LatestRoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestRoundDataTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LatestRoundBlockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLatestRoundParams is a free data retrieval call binding the contract method 0xd1375817.
//
// Solidity: function getLatestRoundParams() view returns(uint256 latestRoundId, uint128 latestRoundDataTimestamp, uint128 latestRoundBlockTimestamp)
func (_RedStoneOracle *RedStoneOracleSession) GetLatestRoundParams() (struct {
	LatestRoundId             *big.Int
	LatestRoundDataTimestamp  *big.Int
	LatestRoundBlockTimestamp *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetLatestRoundParams(&_RedStoneOracle.CallOpts)
}

// GetLatestRoundParams is a free data retrieval call binding the contract method 0xd1375817.
//
// Solidity: function getLatestRoundParams() view returns(uint256 latestRoundId, uint128 latestRoundDataTimestamp, uint128 latestRoundBlockTimestamp)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetLatestRoundParams() (struct {
	LatestRoundId             *big.Int
	LatestRoundDataTimestamp  *big.Int
	LatestRoundBlockTimestamp *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetLatestRoundParams(&_RedStoneOracle.CallOpts)
}

// GetMinIntervalBetweenUpdates is a free data retrieval call binding the contract method 0xd149c0d7.
//
// Solidity: function getMinIntervalBetweenUpdates() view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCaller) GetMinIntervalBetweenUpdates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getMinIntervalBetweenUpdates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinIntervalBetweenUpdates is a free data retrieval call binding the contract method 0xd149c0d7.
//
// Solidity: function getMinIntervalBetweenUpdates() view returns(uint256)
func (_RedStoneOracle *RedStoneOracleSession) GetMinIntervalBetweenUpdates() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetMinIntervalBetweenUpdates(&_RedStoneOracle.CallOpts)
}

// GetMinIntervalBetweenUpdates is a free data retrieval call binding the contract method 0xd149c0d7.
//
// Solidity: function getMinIntervalBetweenUpdates() view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetMinIntervalBetweenUpdates() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetMinIntervalBetweenUpdates(&_RedStoneOracle.CallOpts)
}

// GetPackedTimestampsForRound is a free data retrieval call binding the contract method 0x8ec7c821.
//
// Solidity: function getPackedTimestampsForRound(uint256 roundId) view returns(uint256 roundTimestamps)
func (_RedStoneOracle *RedStoneOracleCaller) GetPackedTimestampsForRound(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getPackedTimestampsForRound", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPackedTimestampsForRound is a free data retrieval call binding the contract method 0x8ec7c821.
//
// Solidity: function getPackedTimestampsForRound(uint256 roundId) view returns(uint256 roundTimestamps)
func (_RedStoneOracle *RedStoneOracleSession) GetPackedTimestampsForRound(roundId *big.Int) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetPackedTimestampsForRound(&_RedStoneOracle.CallOpts, roundId)
}

// GetPackedTimestampsForRound is a free data retrieval call binding the contract method 0x8ec7c821.
//
// Solidity: function getPackedTimestampsForRound(uint256 roundId) view returns(uint256 roundTimestamps)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetPackedTimestampsForRound(roundId *big.Int) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetPackedTimestampsForRound(&_RedStoneOracle.CallOpts, roundId)
}

// GetPackedTimestampsFromLatestUpdate is a free data retrieval call binding the contract method 0xfd1f4bef.
//
// Solidity: function getPackedTimestampsFromLatestUpdate() view returns(uint256 packedTimestamps)
func (_RedStoneOracle *RedStoneOracleCaller) GetPackedTimestampsFromLatestUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getPackedTimestampsFromLatestUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPackedTimestampsFromLatestUpdate is a free data retrieval call binding the contract method 0xfd1f4bef.
//
// Solidity: function getPackedTimestampsFromLatestUpdate() view returns(uint256 packedTimestamps)
func (_RedStoneOracle *RedStoneOracleSession) GetPackedTimestampsFromLatestUpdate() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetPackedTimestampsFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetPackedTimestampsFromLatestUpdate is a free data retrieval call binding the contract method 0xfd1f4bef.
//
// Solidity: function getPackedTimestampsFromLatestUpdate() view returns(uint256 packedTimestamps)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetPackedTimestampsFromLatestUpdate() (*big.Int, error) {
	return _RedStoneOracle.Contract.GetPackedTimestampsFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetPriceFeedAdapter is a free data retrieval call binding the contract method 0x47043b00.
//
// Solidity: function getPriceFeedAdapter() view returns(address)
func (_RedStoneOracle *RedStoneOracleCaller) GetPriceFeedAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getPriceFeedAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceFeedAdapter is a free data retrieval call binding the contract method 0x47043b00.
//
// Solidity: function getPriceFeedAdapter() view returns(address)
func (_RedStoneOracle *RedStoneOracleSession) GetPriceFeedAdapter() (common.Address, error) {
	return _RedStoneOracle.Contract.GetPriceFeedAdapter(&_RedStoneOracle.CallOpts)
}

// GetPriceFeedAdapter is a free data retrieval call binding the contract method 0x47043b00.
//
// Solidity: function getPriceFeedAdapter() view returns(address)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetPriceFeedAdapter() (common.Address, error) {
	return _RedStoneOracle.Contract.GetPriceFeedAdapter(&_RedStoneOracle.CallOpts)
}

// GetPriceFeedAdapterWithRounds is a free data retrieval call binding the contract method 0x4b6a9d81.
//
// Solidity: function getPriceFeedAdapterWithRounds() view returns(address)
func (_RedStoneOracle *RedStoneOracleCaller) GetPriceFeedAdapterWithRounds(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getPriceFeedAdapterWithRounds")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceFeedAdapterWithRounds is a free data retrieval call binding the contract method 0x4b6a9d81.
//
// Solidity: function getPriceFeedAdapterWithRounds() view returns(address)
func (_RedStoneOracle *RedStoneOracleSession) GetPriceFeedAdapterWithRounds() (common.Address, error) {
	return _RedStoneOracle.Contract.GetPriceFeedAdapterWithRounds(&_RedStoneOracle.CallOpts)
}

// GetPriceFeedAdapterWithRounds is a free data retrieval call binding the contract method 0x4b6a9d81.
//
// Solidity: function getPriceFeedAdapterWithRounds() view returns(address)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetPriceFeedAdapterWithRounds() (common.Address, error) {
	return _RedStoneOracle.Contract.GetPriceFeedAdapterWithRounds(&_RedStoneOracle.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 requestedRoundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_RedStoneOracle *RedStoneOracleCaller) GetRoundData(opts *bind.CallOpts, requestedRoundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getRoundData", requestedRoundId)

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

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 requestedRoundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_RedStoneOracle *RedStoneOracleSession) GetRoundData(requestedRoundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetRoundData(&_RedStoneOracle.CallOpts, requestedRoundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 requestedRoundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetRoundData(requestedRoundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetRoundData(&_RedStoneOracle.CallOpts, requestedRoundId)
}

// GetRoundDataFromAdapter is a free data retrieval call binding the contract method 0x26bf15ff.
//
// Solidity: function getRoundDataFromAdapter(bytes32 dataFeedId, uint256 roundId) view returns(uint256 dataFeedValue, uint128 roundDataTimestamp, uint128 roundBlockTimestamp)
func (_RedStoneOracle *RedStoneOracleCaller) GetRoundDataFromAdapter(opts *bind.CallOpts, dataFeedId [32]byte, roundId *big.Int) (struct {
	DataFeedValue       *big.Int
	RoundDataTimestamp  *big.Int
	RoundBlockTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getRoundDataFromAdapter", dataFeedId, roundId)

	outstruct := new(struct {
		DataFeedValue       *big.Int
		RoundDataTimestamp  *big.Int
		RoundBlockTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataFeedValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RoundDataTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RoundBlockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundDataFromAdapter is a free data retrieval call binding the contract method 0x26bf15ff.
//
// Solidity: function getRoundDataFromAdapter(bytes32 dataFeedId, uint256 roundId) view returns(uint256 dataFeedValue, uint128 roundDataTimestamp, uint128 roundBlockTimestamp)
func (_RedStoneOracle *RedStoneOracleSession) GetRoundDataFromAdapter(dataFeedId [32]byte, roundId *big.Int) (struct {
	DataFeedValue       *big.Int
	RoundDataTimestamp  *big.Int
	RoundBlockTimestamp *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetRoundDataFromAdapter(&_RedStoneOracle.CallOpts, dataFeedId, roundId)
}

// GetRoundDataFromAdapter is a free data retrieval call binding the contract method 0x26bf15ff.
//
// Solidity: function getRoundDataFromAdapter(bytes32 dataFeedId, uint256 roundId) view returns(uint256 dataFeedValue, uint128 roundDataTimestamp, uint128 roundBlockTimestamp)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetRoundDataFromAdapter(dataFeedId [32]byte, roundId *big.Int) (struct {
	DataFeedValue       *big.Int
	RoundDataTimestamp  *big.Int
	RoundBlockTimestamp *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetRoundDataFromAdapter(&_RedStoneOracle.CallOpts, dataFeedId, roundId)
}

// GetTimestampsFromLatestUpdate is a free data retrieval call binding the contract method 0xb0f106b0.
//
// Solidity: function getTimestampsFromLatestUpdate() view returns(uint128 dataTimestamp, uint128 blockTimestamp)
func (_RedStoneOracle *RedStoneOracleCaller) GetTimestampsFromLatestUpdate(opts *bind.CallOpts) (struct {
	DataTimestamp  *big.Int
	BlockTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getTimestampsFromLatestUpdate")

	outstruct := new(struct {
		DataTimestamp  *big.Int
		BlockTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTimestampsFromLatestUpdate is a free data retrieval call binding the contract method 0xb0f106b0.
//
// Solidity: function getTimestampsFromLatestUpdate() view returns(uint128 dataTimestamp, uint128 blockTimestamp)
func (_RedStoneOracle *RedStoneOracleSession) GetTimestampsFromLatestUpdate() (struct {
	DataTimestamp  *big.Int
	BlockTimestamp *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetTimestampsFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetTimestampsFromLatestUpdate is a free data retrieval call binding the contract method 0xb0f106b0.
//
// Solidity: function getTimestampsFromLatestUpdate() view returns(uint128 dataTimestamp, uint128 blockTimestamp)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetTimestampsFromLatestUpdate() (struct {
	DataTimestamp  *big.Int
	BlockTimestamp *big.Int
}, error) {
	return _RedStoneOracle.Contract.GetTimestampsFromLatestUpdate(&_RedStoneOracle.CallOpts)
}

// GetUniqueSignersThreshold is a free data retrieval call binding the contract method 0xf90c4924.
//
// Solidity: function getUniqueSignersThreshold() view returns(uint8)
func (_RedStoneOracle *RedStoneOracleCaller) GetUniqueSignersThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getUniqueSignersThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUniqueSignersThreshold is a free data retrieval call binding the contract method 0xf90c4924.
//
// Solidity: function getUniqueSignersThreshold() view returns(uint8)
func (_RedStoneOracle *RedStoneOracleSession) GetUniqueSignersThreshold() (uint8, error) {
	return _RedStoneOracle.Contract.GetUniqueSignersThreshold(&_RedStoneOracle.CallOpts)
}

// GetUniqueSignersThreshold is a free data retrieval call binding the contract method 0xf90c4924.
//
// Solidity: function getUniqueSignersThreshold() view returns(uint8)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetUniqueSignersThreshold() (uint8, error) {
	return _RedStoneOracle.Contract.GetUniqueSignersThreshold(&_RedStoneOracle.CallOpts)
}

// GetValueForDataFeed is a free data retrieval call binding the contract method 0x44e02982.
//
// Solidity: function getValueForDataFeed(bytes32 dataFeedId) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCaller) GetValueForDataFeed(opts *bind.CallOpts, dataFeedId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getValueForDataFeed", dataFeedId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueForDataFeed is a free data retrieval call binding the contract method 0x44e02982.
//
// Solidity: function getValueForDataFeed(bytes32 dataFeedId) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleSession) GetValueForDataFeed(dataFeedId [32]byte) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetValueForDataFeed(&_RedStoneOracle.CallOpts, dataFeedId)
}

// GetValueForDataFeed is a free data retrieval call binding the contract method 0x44e02982.
//
// Solidity: function getValueForDataFeed(bytes32 dataFeedId) view returns(uint256)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetValueForDataFeed(dataFeedId [32]byte) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetValueForDataFeed(&_RedStoneOracle.CallOpts, dataFeedId)
}

// GetValueForDataFeedAndRound is a free data retrieval call binding the contract method 0xf34f73d8.
//
// Solidity: function getValueForDataFeedAndRound(bytes32 dataFeedId, uint256 roundId) view returns(uint256 dataFeedValue)
func (_RedStoneOracle *RedStoneOracleCaller) GetValueForDataFeedAndRound(opts *bind.CallOpts, dataFeedId [32]byte, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getValueForDataFeedAndRound", dataFeedId, roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueForDataFeedAndRound is a free data retrieval call binding the contract method 0xf34f73d8.
//
// Solidity: function getValueForDataFeedAndRound(bytes32 dataFeedId, uint256 roundId) view returns(uint256 dataFeedValue)
func (_RedStoneOracle *RedStoneOracleSession) GetValueForDataFeedAndRound(dataFeedId [32]byte, roundId *big.Int) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetValueForDataFeedAndRound(&_RedStoneOracle.CallOpts, dataFeedId, roundId)
}

// GetValueForDataFeedAndRound is a free data retrieval call binding the contract method 0xf34f73d8.
//
// Solidity: function getValueForDataFeedAndRound(bytes32 dataFeedId, uint256 roundId) view returns(uint256 dataFeedValue)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetValueForDataFeedAndRound(dataFeedId [32]byte, roundId *big.Int) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetValueForDataFeedAndRound(&_RedStoneOracle.CallOpts, dataFeedId, roundId)
}

// GetValueForDataFeedUnsafe is a free data retrieval call binding the contract method 0x6668316a.
//
// Solidity: function getValueForDataFeedUnsafe(bytes32 dataFeedId) view returns(uint256 dataFeedValue)
func (_RedStoneOracle *RedStoneOracleCaller) GetValueForDataFeedUnsafe(opts *bind.CallOpts, dataFeedId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getValueForDataFeedUnsafe", dataFeedId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueForDataFeedUnsafe is a free data retrieval call binding the contract method 0x6668316a.
//
// Solidity: function getValueForDataFeedUnsafe(bytes32 dataFeedId) view returns(uint256 dataFeedValue)
func (_RedStoneOracle *RedStoneOracleSession) GetValueForDataFeedUnsafe(dataFeedId [32]byte) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetValueForDataFeedUnsafe(&_RedStoneOracle.CallOpts, dataFeedId)
}

// GetValueForDataFeedUnsafe is a free data retrieval call binding the contract method 0x6668316a.
//
// Solidity: function getValueForDataFeedUnsafe(bytes32 dataFeedId) view returns(uint256 dataFeedValue)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetValueForDataFeedUnsafe(dataFeedId [32]byte) (*big.Int, error) {
	return _RedStoneOracle.Contract.GetValueForDataFeedUnsafe(&_RedStoneOracle.CallOpts, dataFeedId)
}

// GetValuesForDataFeeds is a free data retrieval call binding the contract method 0x971b9c03.
//
// Solidity: function getValuesForDataFeeds(bytes32[] dataFeedIds) view returns(uint256[])
func (_RedStoneOracle *RedStoneOracleCaller) GetValuesForDataFeeds(opts *bind.CallOpts, dataFeedIds [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getValuesForDataFeeds", dataFeedIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetValuesForDataFeeds is a free data retrieval call binding the contract method 0x971b9c03.
//
// Solidity: function getValuesForDataFeeds(bytes32[] dataFeedIds) view returns(uint256[])
func (_RedStoneOracle *RedStoneOracleSession) GetValuesForDataFeeds(dataFeedIds [][32]byte) ([]*big.Int, error) {
	return _RedStoneOracle.Contract.GetValuesForDataFeeds(&_RedStoneOracle.CallOpts, dataFeedIds)
}

// GetValuesForDataFeeds is a free data retrieval call binding the contract method 0x971b9c03.
//
// Solidity: function getValuesForDataFeeds(bytes32[] dataFeedIds) view returns(uint256[])
func (_RedStoneOracle *RedStoneOracleCallerSession) GetValuesForDataFeeds(dataFeedIds [][32]byte) ([]*big.Int, error) {
	return _RedStoneOracle.Contract.GetValuesForDataFeeds(&_RedStoneOracle.CallOpts, dataFeedIds)
}

// GetValuesForDataFeedsUnsafe is a free data retrieval call binding the contract method 0x55d12458.
//
// Solidity: function getValuesForDataFeedsUnsafe(bytes32[] requestedDataFeedIds) view returns(uint256[] values)
func (_RedStoneOracle *RedStoneOracleCaller) GetValuesForDataFeedsUnsafe(opts *bind.CallOpts, requestedDataFeedIds [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "getValuesForDataFeedsUnsafe", requestedDataFeedIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetValuesForDataFeedsUnsafe is a free data retrieval call binding the contract method 0x55d12458.
//
// Solidity: function getValuesForDataFeedsUnsafe(bytes32[] requestedDataFeedIds) view returns(uint256[] values)
func (_RedStoneOracle *RedStoneOracleSession) GetValuesForDataFeedsUnsafe(requestedDataFeedIds [][32]byte) ([]*big.Int, error) {
	return _RedStoneOracle.Contract.GetValuesForDataFeedsUnsafe(&_RedStoneOracle.CallOpts, requestedDataFeedIds)
}

// GetValuesForDataFeedsUnsafe is a free data retrieval call binding the contract method 0x55d12458.
//
// Solidity: function getValuesForDataFeedsUnsafe(bytes32[] requestedDataFeedIds) view returns(uint256[] values)
func (_RedStoneOracle *RedStoneOracleCallerSession) GetValuesForDataFeedsUnsafe(requestedDataFeedIds [][32]byte) ([]*big.Int, error) {
	return _RedStoneOracle.Contract.GetValuesForDataFeedsUnsafe(&_RedStoneOracle.CallOpts, requestedDataFeedIds)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_RedStoneOracle *RedStoneOracleCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_RedStoneOracle *RedStoneOracleSession) LatestAnswer() (*big.Int, error) {
	return _RedStoneOracle.Contract.LatestAnswer(&_RedStoneOracle.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_RedStoneOracle *RedStoneOracleCallerSession) LatestAnswer() (*big.Int, error) {
	return _RedStoneOracle.Contract.LatestAnswer(&_RedStoneOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_RedStoneOracle *RedStoneOracleCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_RedStoneOracle *RedStoneOracleSession) LatestRound() (*big.Int, error) {
	return _RedStoneOracle.Contract.LatestRound(&_RedStoneOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_RedStoneOracle *RedStoneOracleCallerSession) LatestRound() (*big.Int, error) {
	return _RedStoneOracle.Contract.LatestRound(&_RedStoneOracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_RedStoneOracle *RedStoneOracleCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "latestRoundData")

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
func (_RedStoneOracle *RedStoneOracleSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _RedStoneOracle.Contract.LatestRoundData(&_RedStoneOracle.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_RedStoneOracle *RedStoneOracleCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _RedStoneOracle.Contract.LatestRoundData(&_RedStoneOracle.CallOpts)
}

// RequireAuthorisedUpdater is a free data retrieval call binding the contract method 0xa8b940e6.
//
// Solidity: function requireAuthorisedUpdater(address updater) view returns()
func (_RedStoneOracle *RedStoneOracleCaller) RequireAuthorisedUpdater(opts *bind.CallOpts, updater common.Address) error {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "requireAuthorisedUpdater", updater)

	if err != nil {
		return err
	}

	return err

}

// RequireAuthorisedUpdater is a free data retrieval call binding the contract method 0xa8b940e6.
//
// Solidity: function requireAuthorisedUpdater(address updater) view returns()
func (_RedStoneOracle *RedStoneOracleSession) RequireAuthorisedUpdater(updater common.Address) error {
	return _RedStoneOracle.Contract.RequireAuthorisedUpdater(&_RedStoneOracle.CallOpts, updater)
}

// RequireAuthorisedUpdater is a free data retrieval call binding the contract method 0xa8b940e6.
//
// Solidity: function requireAuthorisedUpdater(address updater) view returns()
func (_RedStoneOracle *RedStoneOracleCallerSession) RequireAuthorisedUpdater(updater common.Address) error {
	return _RedStoneOracle.Contract.RequireAuthorisedUpdater(&_RedStoneOracle.CallOpts, updater)
}

// ValidateDataFeedValueOnRead is a free data retrieval call binding the contract method 0xcbc33eb2.
//
// Solidity: function validateDataFeedValueOnRead(bytes32 dataFeedId, uint256 valueForDataFeed) view returns()
func (_RedStoneOracle *RedStoneOracleCaller) ValidateDataFeedValueOnRead(opts *bind.CallOpts, dataFeedId [32]byte, valueForDataFeed *big.Int) error {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "validateDataFeedValueOnRead", dataFeedId, valueForDataFeed)

	if err != nil {
		return err
	}

	return err

}

// ValidateDataFeedValueOnRead is a free data retrieval call binding the contract method 0xcbc33eb2.
//
// Solidity: function validateDataFeedValueOnRead(bytes32 dataFeedId, uint256 valueForDataFeed) view returns()
func (_RedStoneOracle *RedStoneOracleSession) ValidateDataFeedValueOnRead(dataFeedId [32]byte, valueForDataFeed *big.Int) error {
	return _RedStoneOracle.Contract.ValidateDataFeedValueOnRead(&_RedStoneOracle.CallOpts, dataFeedId, valueForDataFeed)
}

// ValidateDataFeedValueOnRead is a free data retrieval call binding the contract method 0xcbc33eb2.
//
// Solidity: function validateDataFeedValueOnRead(bytes32 dataFeedId, uint256 valueForDataFeed) view returns()
func (_RedStoneOracle *RedStoneOracleCallerSession) ValidateDataFeedValueOnRead(dataFeedId [32]byte, valueForDataFeed *big.Int) error {
	return _RedStoneOracle.Contract.ValidateDataFeedValueOnRead(&_RedStoneOracle.CallOpts, dataFeedId, valueForDataFeed)
}

// ValidateDataFeedValueOnWrite is a free data retrieval call binding the contract method 0x6e3e0370.
//
// Solidity: function validateDataFeedValueOnWrite(bytes32 dataFeedId, uint256 valueForDataFeed) view returns()
func (_RedStoneOracle *RedStoneOracleCaller) ValidateDataFeedValueOnWrite(opts *bind.CallOpts, dataFeedId [32]byte, valueForDataFeed *big.Int) error {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "validateDataFeedValueOnWrite", dataFeedId, valueForDataFeed)

	if err != nil {
		return err
	}

	return err

}

// ValidateDataFeedValueOnWrite is a free data retrieval call binding the contract method 0x6e3e0370.
//
// Solidity: function validateDataFeedValueOnWrite(bytes32 dataFeedId, uint256 valueForDataFeed) view returns()
func (_RedStoneOracle *RedStoneOracleSession) ValidateDataFeedValueOnWrite(dataFeedId [32]byte, valueForDataFeed *big.Int) error {
	return _RedStoneOracle.Contract.ValidateDataFeedValueOnWrite(&_RedStoneOracle.CallOpts, dataFeedId, valueForDataFeed)
}

// ValidateDataFeedValueOnWrite is a free data retrieval call binding the contract method 0x6e3e0370.
//
// Solidity: function validateDataFeedValueOnWrite(bytes32 dataFeedId, uint256 valueForDataFeed) view returns()
func (_RedStoneOracle *RedStoneOracleCallerSession) ValidateDataFeedValueOnWrite(dataFeedId [32]byte, valueForDataFeed *big.Int) error {
	return _RedStoneOracle.Contract.ValidateDataFeedValueOnWrite(&_RedStoneOracle.CallOpts, dataFeedId, valueForDataFeed)
}

// ValidateDataPackagesTimestampOnce is a free data retrieval call binding the contract method 0xbb1f29b7.
//
// Solidity: function validateDataPackagesTimestampOnce(uint256 dataPackagesTimestamp) view returns()
func (_RedStoneOracle *RedStoneOracleCaller) ValidateDataPackagesTimestampOnce(opts *bind.CallOpts, dataPackagesTimestamp *big.Int) error {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "validateDataPackagesTimestampOnce", dataPackagesTimestamp)

	if err != nil {
		return err
	}

	return err

}

// ValidateDataPackagesTimestampOnce is a free data retrieval call binding the contract method 0xbb1f29b7.
//
// Solidity: function validateDataPackagesTimestampOnce(uint256 dataPackagesTimestamp) view returns()
func (_RedStoneOracle *RedStoneOracleSession) ValidateDataPackagesTimestampOnce(dataPackagesTimestamp *big.Int) error {
	return _RedStoneOracle.Contract.ValidateDataPackagesTimestampOnce(&_RedStoneOracle.CallOpts, dataPackagesTimestamp)
}

// ValidateDataPackagesTimestampOnce is a free data retrieval call binding the contract method 0xbb1f29b7.
//
// Solidity: function validateDataPackagesTimestampOnce(uint256 dataPackagesTimestamp) view returns()
func (_RedStoneOracle *RedStoneOracleCallerSession) ValidateDataPackagesTimestampOnce(dataPackagesTimestamp *big.Int) error {
	return _RedStoneOracle.Contract.ValidateDataPackagesTimestampOnce(&_RedStoneOracle.CallOpts, dataPackagesTimestamp)
}

// ValidateProposedDataPackagesTimestamp is a free data retrieval call binding the contract method 0xada11457.
//
// Solidity: function validateProposedDataPackagesTimestamp(uint256 dataPackagesTimestamp) view returns()
func (_RedStoneOracle *RedStoneOracleCaller) ValidateProposedDataPackagesTimestamp(opts *bind.CallOpts, dataPackagesTimestamp *big.Int) error {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "validateProposedDataPackagesTimestamp", dataPackagesTimestamp)

	if err != nil {
		return err
	}

	return err

}

// ValidateProposedDataPackagesTimestamp is a free data retrieval call binding the contract method 0xada11457.
//
// Solidity: function validateProposedDataPackagesTimestamp(uint256 dataPackagesTimestamp) view returns()
func (_RedStoneOracle *RedStoneOracleSession) ValidateProposedDataPackagesTimestamp(dataPackagesTimestamp *big.Int) error {
	return _RedStoneOracle.Contract.ValidateProposedDataPackagesTimestamp(&_RedStoneOracle.CallOpts, dataPackagesTimestamp)
}

// ValidateProposedDataPackagesTimestamp is a free data retrieval call binding the contract method 0xada11457.
//
// Solidity: function validateProposedDataPackagesTimestamp(uint256 dataPackagesTimestamp) view returns()
func (_RedStoneOracle *RedStoneOracleCallerSession) ValidateProposedDataPackagesTimestamp(dataPackagesTimestamp *big.Int) error {
	return _RedStoneOracle.Contract.ValidateProposedDataPackagesTimestamp(&_RedStoneOracle.CallOpts, dataPackagesTimestamp)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 receivedTimestampMilliseconds) view returns()
func (_RedStoneOracle *RedStoneOracleCaller) ValidateTimestamp(opts *bind.CallOpts, receivedTimestampMilliseconds *big.Int) error {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "validateTimestamp", receivedTimestampMilliseconds)

	if err != nil {
		return err
	}

	return err

}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 receivedTimestampMilliseconds) view returns()
func (_RedStoneOracle *RedStoneOracleSession) ValidateTimestamp(receivedTimestampMilliseconds *big.Int) error {
	return _RedStoneOracle.Contract.ValidateTimestamp(&_RedStoneOracle.CallOpts, receivedTimestampMilliseconds)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 receivedTimestampMilliseconds) view returns()
func (_RedStoneOracle *RedStoneOracleCallerSession) ValidateTimestamp(receivedTimestampMilliseconds *big.Int) error {
	return _RedStoneOracle.Contract.ValidateTimestamp(&_RedStoneOracle.CallOpts, receivedTimestampMilliseconds)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_RedStoneOracle *RedStoneOracleCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedStoneOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_RedStoneOracle *RedStoneOracleSession) Version() (*big.Int, error) {
	return _RedStoneOracle.Contract.Version(&_RedStoneOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_RedStoneOracle *RedStoneOracleCallerSession) Version() (*big.Int, error) {
	return _RedStoneOracle.Contract.Version(&_RedStoneOracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RedStoneOracle *RedStoneOracleTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedStoneOracle.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RedStoneOracle *RedStoneOracleSession) Initialize() (*types.Transaction, error) {
	return _RedStoneOracle.Contract.Initialize(&_RedStoneOracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_RedStoneOracle *RedStoneOracleTransactorSession) Initialize() (*types.Transaction, error) {
	return _RedStoneOracle.Contract.Initialize(&_RedStoneOracle.TransactOpts)
}

// UpdateDataFeedsValues is a paid mutator transaction binding the contract method 0xc14c9204.
//
// Solidity: function updateDataFeedsValues(uint256 dataPackagesTimestamp) returns()
func (_RedStoneOracle *RedStoneOracleTransactor) UpdateDataFeedsValues(opts *bind.TransactOpts, dataPackagesTimestamp *big.Int) (*types.Transaction, error) {
	return _RedStoneOracle.contract.Transact(opts, "updateDataFeedsValues", dataPackagesTimestamp)
}

// UpdateDataFeedsValues is a paid mutator transaction binding the contract method 0xc14c9204.
//
// Solidity: function updateDataFeedsValues(uint256 dataPackagesTimestamp) returns()
func (_RedStoneOracle *RedStoneOracleSession) UpdateDataFeedsValues(dataPackagesTimestamp *big.Int) (*types.Transaction, error) {
	return _RedStoneOracle.Contract.UpdateDataFeedsValues(&_RedStoneOracle.TransactOpts, dataPackagesTimestamp)
}

// UpdateDataFeedsValues is a paid mutator transaction binding the contract method 0xc14c9204.
//
// Solidity: function updateDataFeedsValues(uint256 dataPackagesTimestamp) returns()
func (_RedStoneOracle *RedStoneOracleTransactorSession) UpdateDataFeedsValues(dataPackagesTimestamp *big.Int) (*types.Transaction, error) {
	return _RedStoneOracle.Contract.UpdateDataFeedsValues(&_RedStoneOracle.TransactOpts, dataPackagesTimestamp)
}

// RedStoneOracleAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the RedStoneOracle contract.
type RedStoneOracleAnswerUpdatedIterator struct {
	Event *RedStoneOracleAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *RedStoneOracleAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedStoneOracleAnswerUpdated)
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
		it.Event = new(RedStoneOracleAnswerUpdated)
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
func (it *RedStoneOracleAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedStoneOracleAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedStoneOracleAnswerUpdated represents a AnswerUpdated event raised by the RedStoneOracle contract.
type RedStoneOracleAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_RedStoneOracle *RedStoneOracleFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*RedStoneOracleAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _RedStoneOracle.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &RedStoneOracleAnswerUpdatedIterator{contract: _RedStoneOracle.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_RedStoneOracle *RedStoneOracleFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *RedStoneOracleAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _RedStoneOracle.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedStoneOracleAnswerUpdated)
				if err := _RedStoneOracle.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_RedStoneOracle *RedStoneOracleFilterer) ParseAnswerUpdated(log types.Log) (*RedStoneOracleAnswerUpdated, error) {
	event := new(RedStoneOracleAnswerUpdated)
	if err := _RedStoneOracle.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedStoneOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RedStoneOracle contract.
type RedStoneOracleInitializedIterator struct {
	Event *RedStoneOracleInitialized // Event containing the contract specifics and raw log

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
func (it *RedStoneOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedStoneOracleInitialized)
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
		it.Event = new(RedStoneOracleInitialized)
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
func (it *RedStoneOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedStoneOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedStoneOracleInitialized represents a Initialized event raised by the RedStoneOracle contract.
type RedStoneOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedStoneOracle *RedStoneOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*RedStoneOracleInitializedIterator, error) {

	logs, sub, err := _RedStoneOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RedStoneOracleInitializedIterator{contract: _RedStoneOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedStoneOracle *RedStoneOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RedStoneOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _RedStoneOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedStoneOracleInitialized)
				if err := _RedStoneOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RedStoneOracle *RedStoneOracleFilterer) ParseInitialized(log types.Log) (*RedStoneOracleInitialized, error) {
	event := new(RedStoneOracleInitialized)
	if err := _RedStoneOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
