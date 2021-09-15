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

// ExpMathMockMetaData contains all meta data concerning the ExpMathMock contract.
var ExpMathMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"gas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"t0\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"c0\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"t12\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"t\",\"type\":\"uint32\"}],\"name\":\"halfLife\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"t0\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"c0\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"t12\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"t\",\"type\":\"uint32\"}],\"name\":\"measure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"result\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103f2806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806365372147146100515780636ca7c2161461008157806380756ef2146100985780639d4483a8146100ab575b600080fd5b600054610064906001600160701b031681565b6040516001600160701b0390911681526020015b60405180910390f35b61008a60015481565b604051908152602001610078565b6100646100a63660046102e6565b6100c0565b6100be6100b93660046102e6565b6100d9565b005b60006100ce85858585610123565b90505b949350505050565b60005a905060006100ec868686866100c0565b90505a6100f99083610399565b600155600080546001600160701b0319166001600160701b03929092169190911790555050505050565b60008463ffffffff168263ffffffff161015604051806040016040528060018152602001600760fb1b815250906101765760405162461bcd60e51b815260040161016d9190610346565b60405180910390fd5b5084820391508263ffffffff168263ffffffff16816101a557634e487b7160e01b600052601260045260246000fd5b0463ffffffff16846001600160701b0316901c93508263ffffffff168263ffffffff16816101e357634e487b7160e01b600052601260045260246000fd5b06915063ffffffff8216158061020057506001600160701b038416155b1561020c5750826100d1565b60006001600160701b0385168163ffffffff86811690861671b17217f7d1cf79abc9e3b39803f2f6af40f3028161025357634e487b7160e01b600052601260045260246000fd5b049050600160901b5b82156102c05792820192808284028161028557634e487b7160e01b600052601260045260246000fd5b0493849003939250600160901b0180828402816102b257634e487b7160e01b600052601260045260246000fd5b049250600160901b0161025c565b5091979650505050505050565b803563ffffffff811681146102e157600080fd5b919050565b600080600080608085870312156102fb578384fd5b610304856102cd565b935060208501356001600160701b038116811461031f578384fd5b925061032d604086016102cd565b915061033b606086016102cd565b905092959194509250565b6000602080835283518082850152825b8181101561037257858101830151858201604001528201610356565b818111156103835783604083870101525b50601f01601f1916929092016040019392505050565b6000828210156103b757634e487b7160e01b81526011600452602481fd5b50039056fea264697066735822122038627719d2646cf7d9b16521426fbaf296f38c19fd3dd1db1790c2e53d68ceae64736f6c63430008040033",
}

// ExpMathMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ExpMathMockMetaData.ABI instead.
var ExpMathMockABI = ExpMathMockMetaData.ABI

// ExpMathMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExpMathMockMetaData.Bin instead.
var ExpMathMockBin = ExpMathMockMetaData.Bin

// DeployExpMathMock deploys a new Ethereum contract, binding an instance of ExpMathMock to it.
func DeployExpMathMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExpMathMock, error) {
	parsed, err := ExpMathMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExpMathMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExpMathMock{ExpMathMockCaller: ExpMathMockCaller{contract: contract}, ExpMathMockTransactor: ExpMathMockTransactor{contract: contract}, ExpMathMockFilterer: ExpMathMockFilterer{contract: contract}}, nil
}

// ExpMathMock is an auto generated Go binding around an Ethereum contract.
type ExpMathMock struct {
	ExpMathMockCaller     // Read-only binding to the contract
	ExpMathMockTransactor // Write-only binding to the contract
	ExpMathMockFilterer   // Log filterer for contract events
}

// ExpMathMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExpMathMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpMathMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExpMathMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpMathMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExpMathMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpMathMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExpMathMockSession struct {
	Contract     *ExpMathMock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExpMathMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExpMathMockCallerSession struct {
	Contract *ExpMathMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExpMathMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExpMathMockTransactorSession struct {
	Contract     *ExpMathMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExpMathMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExpMathMockRaw struct {
	Contract *ExpMathMock // Generic contract binding to access the raw methods on
}

// ExpMathMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExpMathMockCallerRaw struct {
	Contract *ExpMathMockCaller // Generic read-only contract binding to access the raw methods on
}

// ExpMathMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExpMathMockTransactorRaw struct {
	Contract *ExpMathMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExpMathMock creates a new instance of ExpMathMock, bound to a specific deployed contract.
func NewExpMathMock(address common.Address, backend bind.ContractBackend) (*ExpMathMock, error) {
	contract, err := bindExpMathMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExpMathMock{ExpMathMockCaller: ExpMathMockCaller{contract: contract}, ExpMathMockTransactor: ExpMathMockTransactor{contract: contract}, ExpMathMockFilterer: ExpMathMockFilterer{contract: contract}}, nil
}

// NewExpMathMockCaller creates a new read-only instance of ExpMathMock, bound to a specific deployed contract.
func NewExpMathMockCaller(address common.Address, caller bind.ContractCaller) (*ExpMathMockCaller, error) {
	contract, err := bindExpMathMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExpMathMockCaller{contract: contract}, nil
}

// NewExpMathMockTransactor creates a new write-only instance of ExpMathMock, bound to a specific deployed contract.
func NewExpMathMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ExpMathMockTransactor, error) {
	contract, err := bindExpMathMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExpMathMockTransactor{contract: contract}, nil
}

// NewExpMathMockFilterer creates a new log filterer instance of ExpMathMock, bound to a specific deployed contract.
func NewExpMathMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ExpMathMockFilterer, error) {
	contract, err := bindExpMathMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExpMathMockFilterer{contract: contract}, nil
}

// bindExpMathMock binds a generic wrapper to an already deployed contract.
func bindExpMathMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExpMathMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpMathMock *ExpMathMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExpMathMock.Contract.ExpMathMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpMathMock *ExpMathMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpMathMock.Contract.ExpMathMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpMathMock *ExpMathMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpMathMock.Contract.ExpMathMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpMathMock *ExpMathMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExpMathMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpMathMock *ExpMathMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpMathMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpMathMock *ExpMathMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpMathMock.Contract.contract.Transact(opts, method, params...)
}

// Gas is a free data retrieval call binding the contract method 0x6ca7c216.
//
// Solidity: function gas() view returns(uint256)
func (_ExpMathMock *ExpMathMockCaller) Gas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpMathMock.contract.Call(opts, &out, "gas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gas is a free data retrieval call binding the contract method 0x6ca7c216.
//
// Solidity: function gas() view returns(uint256)
func (_ExpMathMock *ExpMathMockSession) Gas() (*big.Int, error) {
	return _ExpMathMock.Contract.Gas(&_ExpMathMock.CallOpts)
}

// Gas is a free data retrieval call binding the contract method 0x6ca7c216.
//
// Solidity: function gas() view returns(uint256)
func (_ExpMathMock *ExpMathMockCallerSession) Gas() (*big.Int, error) {
	return _ExpMathMock.Contract.Gas(&_ExpMathMock.CallOpts)
}

// HalfLife is a free data retrieval call binding the contract method 0x80756ef2.
//
// Solidity: function halfLife(uint32 t0, uint112 c0, uint32 t12, uint32 t) pure returns(uint112)
func (_ExpMathMock *ExpMathMockCaller) HalfLife(opts *bind.CallOpts, t0 uint32, c0 *big.Int, t12 uint32, t uint32) (*big.Int, error) {
	var out []interface{}
	err := _ExpMathMock.contract.Call(opts, &out, "halfLife", t0, c0, t12, t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalfLife is a free data retrieval call binding the contract method 0x80756ef2.
//
// Solidity: function halfLife(uint32 t0, uint112 c0, uint32 t12, uint32 t) pure returns(uint112)
func (_ExpMathMock *ExpMathMockSession) HalfLife(t0 uint32, c0 *big.Int, t12 uint32, t uint32) (*big.Int, error) {
	return _ExpMathMock.Contract.HalfLife(&_ExpMathMock.CallOpts, t0, c0, t12, t)
}

// HalfLife is a free data retrieval call binding the contract method 0x80756ef2.
//
// Solidity: function halfLife(uint32 t0, uint112 c0, uint32 t12, uint32 t) pure returns(uint112)
func (_ExpMathMock *ExpMathMockCallerSession) HalfLife(t0 uint32, c0 *big.Int, t12 uint32, t uint32) (*big.Int, error) {
	return _ExpMathMock.Contract.HalfLife(&_ExpMathMock.CallOpts, t0, c0, t12, t)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() view returns(uint112)
func (_ExpMathMock *ExpMathMockCaller) Result(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpMathMock.contract.Call(opts, &out, "result")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() view returns(uint112)
func (_ExpMathMock *ExpMathMockSession) Result() (*big.Int, error) {
	return _ExpMathMock.Contract.Result(&_ExpMathMock.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() view returns(uint112)
func (_ExpMathMock *ExpMathMockCallerSession) Result() (*big.Int, error) {
	return _ExpMathMock.Contract.Result(&_ExpMathMock.CallOpts)
}

// Measure is a paid mutator transaction binding the contract method 0x9d4483a8.
//
// Solidity: function measure(uint32 t0, uint112 c0, uint32 t12, uint32 t) returns()
func (_ExpMathMock *ExpMathMockTransactor) Measure(opts *bind.TransactOpts, t0 uint32, c0 *big.Int, t12 uint32, t uint32) (*types.Transaction, error) {
	return _ExpMathMock.contract.Transact(opts, "measure", t0, c0, t12, t)
}

// Measure is a paid mutator transaction binding the contract method 0x9d4483a8.
//
// Solidity: function measure(uint32 t0, uint112 c0, uint32 t12, uint32 t) returns()
func (_ExpMathMock *ExpMathMockSession) Measure(t0 uint32, c0 *big.Int, t12 uint32, t uint32) (*types.Transaction, error) {
	return _ExpMathMock.Contract.Measure(&_ExpMathMock.TransactOpts, t0, c0, t12, t)
}

// Measure is a paid mutator transaction binding the contract method 0x9d4483a8.
//
// Solidity: function measure(uint32 t0, uint112 c0, uint32 t12, uint32 t) returns()
func (_ExpMathMock *ExpMathMockTransactorSession) Measure(t0 uint32, c0 *big.Int, t12 uint32, t uint32) (*types.Transaction, error) {
	return _ExpMathMock.Contract.Measure(&_ExpMathMock.TransactOpts, t0, c0, t12, t)
}
