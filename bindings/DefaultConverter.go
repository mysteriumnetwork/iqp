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

// DefaultConverterMetaData contains all meta data concerning the DefaultConverter contract.
var DefaultConverterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"estimateConvert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101a2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063029b465d1461003b57806329fb92ce1461003b575b600080fd5b61004e6100493660046100de565b610060565b60405190815260200160405180910390f35b6000816001600160a01b0316846001600160a01b03161460405180604001604052806002815260200161199b60f11b815250906100b95760405162461bcd60e51b81526004016100b09190610119565b60405180910390fd5b50919392505050565b80356001600160a01b03811681146100d957600080fd5b919050565b6000806000606084860312156100f2578283fd5b6100fb846100c2565b925060208401359150610110604085016100c2565b90509250925092565b6000602080835283518082850152825b8181101561014557858101830151858201604001528201610129565b818111156101565783604083870101525b50601f01601f191692909201604001939250505056fea26469706673582212201b26b3b55f4d405a1083fb953a687509e0cc862aab50fad9f0fdb19abb81d38f64736f6c63430008040033",
}

// DefaultConverterABI is the input ABI used to generate the binding from.
// Deprecated: Use DefaultConverterMetaData.ABI instead.
var DefaultConverterABI = DefaultConverterMetaData.ABI

// DefaultConverterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DefaultConverterMetaData.Bin instead.
var DefaultConverterBin = DefaultConverterMetaData.Bin

// DeployDefaultConverter deploys a new Ethereum contract, binding an instance of DefaultConverter to it.
func DeployDefaultConverter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DefaultConverter, error) {
	parsed, err := DefaultConverterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefaultConverterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefaultConverter{DefaultConverterCaller: DefaultConverterCaller{contract: contract}, DefaultConverterTransactor: DefaultConverterTransactor{contract: contract}, DefaultConverterFilterer: DefaultConverterFilterer{contract: contract}}, nil
}

// DefaultConverter is an auto generated Go binding around an Ethereum contract.
type DefaultConverter struct {
	DefaultConverterCaller     // Read-only binding to the contract
	DefaultConverterTransactor // Write-only binding to the contract
	DefaultConverterFilterer   // Log filterer for contract events
}

// DefaultConverterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefaultConverterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultConverterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefaultConverterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultConverterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefaultConverterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultConverterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefaultConverterSession struct {
	Contract     *DefaultConverter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefaultConverterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefaultConverterCallerSession struct {
	Contract *DefaultConverterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DefaultConverterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefaultConverterTransactorSession struct {
	Contract     *DefaultConverterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DefaultConverterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefaultConverterRaw struct {
	Contract *DefaultConverter // Generic contract binding to access the raw methods on
}

// DefaultConverterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefaultConverterCallerRaw struct {
	Contract *DefaultConverterCaller // Generic read-only contract binding to access the raw methods on
}

// DefaultConverterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefaultConverterTransactorRaw struct {
	Contract *DefaultConverterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefaultConverter creates a new instance of DefaultConverter, bound to a specific deployed contract.
func NewDefaultConverter(address common.Address, backend bind.ContractBackend) (*DefaultConverter, error) {
	contract, err := bindDefaultConverter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefaultConverter{DefaultConverterCaller: DefaultConverterCaller{contract: contract}, DefaultConverterTransactor: DefaultConverterTransactor{contract: contract}, DefaultConverterFilterer: DefaultConverterFilterer{contract: contract}}, nil
}

// NewDefaultConverterCaller creates a new read-only instance of DefaultConverter, bound to a specific deployed contract.
func NewDefaultConverterCaller(address common.Address, caller bind.ContractCaller) (*DefaultConverterCaller, error) {
	contract, err := bindDefaultConverter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultConverterCaller{contract: contract}, nil
}

// NewDefaultConverterTransactor creates a new write-only instance of DefaultConverter, bound to a specific deployed contract.
func NewDefaultConverterTransactor(address common.Address, transactor bind.ContractTransactor) (*DefaultConverterTransactor, error) {
	contract, err := bindDefaultConverter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultConverterTransactor{contract: contract}, nil
}

// NewDefaultConverterFilterer creates a new log filterer instance of DefaultConverter, bound to a specific deployed contract.
func NewDefaultConverterFilterer(address common.Address, filterer bind.ContractFilterer) (*DefaultConverterFilterer, error) {
	contract, err := bindDefaultConverter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefaultConverterFilterer{contract: contract}, nil
}

// bindDefaultConverter binds a generic wrapper to an already deployed contract.
func bindDefaultConverter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefaultConverterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultConverter *DefaultConverterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultConverter.Contract.DefaultConverterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultConverter *DefaultConverterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultConverter.Contract.DefaultConverterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultConverter *DefaultConverterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultConverter.Contract.DefaultConverterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultConverter *DefaultConverterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultConverter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultConverter *DefaultConverterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultConverter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultConverter *DefaultConverterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultConverter.Contract.contract.Transact(opts, method, params...)
}

// Convert is a free data retrieval call binding the contract method 0x029b465d.
//
// Solidity: function convert(address source, uint256 amount, address target) pure returns(uint256)
func (_DefaultConverter *DefaultConverterCaller) Convert(opts *bind.CallOpts, source common.Address, amount *big.Int, target common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefaultConverter.contract.Call(opts, &out, "convert", source, amount, target)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Convert is a free data retrieval call binding the contract method 0x029b465d.
//
// Solidity: function convert(address source, uint256 amount, address target) pure returns(uint256)
func (_DefaultConverter *DefaultConverterSession) Convert(source common.Address, amount *big.Int, target common.Address) (*big.Int, error) {
	return _DefaultConverter.Contract.Convert(&_DefaultConverter.CallOpts, source, amount, target)
}

// Convert is a free data retrieval call binding the contract method 0x029b465d.
//
// Solidity: function convert(address source, uint256 amount, address target) pure returns(uint256)
func (_DefaultConverter *DefaultConverterCallerSession) Convert(source common.Address, amount *big.Int, target common.Address) (*big.Int, error) {
	return _DefaultConverter.Contract.Convert(&_DefaultConverter.CallOpts, source, amount, target)
}

// EstimateConvert is a free data retrieval call binding the contract method 0x29fb92ce.
//
// Solidity: function estimateConvert(address source, uint256 amount, address target) pure returns(uint256)
func (_DefaultConverter *DefaultConverterCaller) EstimateConvert(opts *bind.CallOpts, source common.Address, amount *big.Int, target common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefaultConverter.contract.Call(opts, &out, "estimateConvert", source, amount, target)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateConvert is a free data retrieval call binding the contract method 0x29fb92ce.
//
// Solidity: function estimateConvert(address source, uint256 amount, address target) pure returns(uint256)
func (_DefaultConverter *DefaultConverterSession) EstimateConvert(source common.Address, amount *big.Int, target common.Address) (*big.Int, error) {
	return _DefaultConverter.Contract.EstimateConvert(&_DefaultConverter.CallOpts, source, amount, target)
}

// EstimateConvert is a free data retrieval call binding the contract method 0x29fb92ce.
//
// Solidity: function estimateConvert(address source, uint256 amount, address target) pure returns(uint256)
func (_DefaultConverter *DefaultConverterCallerSession) EstimateConvert(source common.Address, amount *big.Int, target common.Address) (*big.Int, error) {
	return _DefaultConverter.Contract.EstimateConvert(&_DefaultConverter.CallOpts, source, amount, target)
}
