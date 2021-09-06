// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// EnterpriseOwnableMetaData contains all meta data concerning the EnterpriseOwnable contract.
var EnterpriseOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getEnterprise\",\"outputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EnterpriseOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseOwnableMetaData.ABI instead.
var EnterpriseOwnableABI = EnterpriseOwnableMetaData.ABI

// EnterpriseOwnable is an auto generated Go binding around an Ethereum contract.
type EnterpriseOwnable struct {
	EnterpriseOwnableCaller     // Read-only binding to the contract
	EnterpriseOwnableTransactor // Write-only binding to the contract
	EnterpriseOwnableFilterer   // Log filterer for contract events
}

// EnterpriseOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseOwnableSession struct {
	Contract     *EnterpriseOwnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EnterpriseOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseOwnableCallerSession struct {
	Contract *EnterpriseOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EnterpriseOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseOwnableTransactorSession struct {
	Contract     *EnterpriseOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EnterpriseOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseOwnableRaw struct {
	Contract *EnterpriseOwnable // Generic contract binding to access the raw methods on
}

// EnterpriseOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseOwnableCallerRaw struct {
	Contract *EnterpriseOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseOwnableTransactorRaw struct {
	Contract *EnterpriseOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterpriseOwnable creates a new instance of EnterpriseOwnable, bound to a specific deployed contract.
func NewEnterpriseOwnable(address common.Address, backend bind.ContractBackend) (*EnterpriseOwnable, error) {
	contract, err := bindEnterpriseOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterpriseOwnable{EnterpriseOwnableCaller: EnterpriseOwnableCaller{contract: contract}, EnterpriseOwnableTransactor: EnterpriseOwnableTransactor{contract: contract}, EnterpriseOwnableFilterer: EnterpriseOwnableFilterer{contract: contract}}, nil
}

// NewEnterpriseOwnableCaller creates a new read-only instance of EnterpriseOwnable, bound to a specific deployed contract.
func NewEnterpriseOwnableCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseOwnableCaller, error) {
	contract, err := bindEnterpriseOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseOwnableCaller{contract: contract}, nil
}

// NewEnterpriseOwnableTransactor creates a new write-only instance of EnterpriseOwnable, bound to a specific deployed contract.
func NewEnterpriseOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseOwnableTransactor, error) {
	contract, err := bindEnterpriseOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseOwnableTransactor{contract: contract}, nil
}

// NewEnterpriseOwnableFilterer creates a new log filterer instance of EnterpriseOwnable, bound to a specific deployed contract.
func NewEnterpriseOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseOwnableFilterer, error) {
	contract, err := bindEnterpriseOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseOwnableFilterer{contract: contract}, nil
}

// bindEnterpriseOwnable binds a generic wrapper to an already deployed contract.
func bindEnterpriseOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnterpriseOwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseOwnable *EnterpriseOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseOwnable.Contract.EnterpriseOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseOwnable *EnterpriseOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseOwnable.Contract.EnterpriseOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseOwnable *EnterpriseOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseOwnable.Contract.EnterpriseOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseOwnable *EnterpriseOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseOwnable *EnterpriseOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseOwnable *EnterpriseOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseOwnable.Contract.contract.Transact(opts, method, params...)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_EnterpriseOwnable *EnterpriseOwnableCaller) GetEnterprise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseOwnable.contract.Call(opts, &out, "getEnterprise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_EnterpriseOwnable *EnterpriseOwnableSession) GetEnterprise() (common.Address, error) {
	return _EnterpriseOwnable.Contract.GetEnterprise(&_EnterpriseOwnable.CallOpts)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_EnterpriseOwnable *EnterpriseOwnableCallerSession) GetEnterprise() (common.Address, error) {
	return _EnterpriseOwnable.Contract.GetEnterprise(&_EnterpriseOwnable.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_EnterpriseOwnable *EnterpriseOwnableTransactor) Initialize(opts *bind.TransactOpts, enterprise common.Address) (*types.Transaction, error) {
	return _EnterpriseOwnable.contract.Transact(opts, "initialize", enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_EnterpriseOwnable *EnterpriseOwnableSession) Initialize(enterprise common.Address) (*types.Transaction, error) {
	return _EnterpriseOwnable.Contract.Initialize(&_EnterpriseOwnable.TransactOpts, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_EnterpriseOwnable *EnterpriseOwnableTransactorSession) Initialize(enterprise common.Address) (*types.Transaction, error) {
	return _EnterpriseOwnable.Contract.Initialize(&_EnterpriseOwnable.TransactOpts, enterprise)
}
