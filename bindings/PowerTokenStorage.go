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

// PowerTokenStorageState is an auto generated low-level Go binding around an user-defined struct.
type PowerTokenStorageState struct {
	LockedBalance *big.Int
	Energy        *big.Int
	Timestamp     uint32
}

// PowerTokenStorageMetaData contains all meta data concerning the PowerTokenStorage contract.
var PowerTokenStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"}],\"name\":\"BaseRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"minDuration\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxDuration\",\"type\":\"uint32\"}],\"name\":\"LoanDurationLimitsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PerpetualAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"percent\",\"type\":\"uint16\"}],\"name\":\"ServiceFeePercentChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowPerpetualForever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowsPerpetual\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseRate\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterprise\",\"outputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGapHalvingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxLoanDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinGCFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinLoanDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getServiceFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"lockedBalance\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"energy\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structPowerTokenStorage.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"gapHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minLoanDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxLoanDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"serviceFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"allowsPerpetual\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"isAllowedLoanDuration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"}],\"name\":\"setBaseRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"minLoanDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxLoanDuration\",\"type\":\"uint32\"}],\"name\":\"setLoanDurationLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newServiceFeePercent\",\"type\":\"uint16\"}],\"name\":\"setServiceFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PowerTokenStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use PowerTokenStorageMetaData.ABI instead.
var PowerTokenStorageABI = PowerTokenStorageMetaData.ABI

// PowerTokenStorage is an auto generated Go binding around an Ethereum contract.
type PowerTokenStorage struct {
	PowerTokenStorageCaller     // Read-only binding to the contract
	PowerTokenStorageTransactor // Write-only binding to the contract
	PowerTokenStorageFilterer   // Log filterer for contract events
}

// PowerTokenStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PowerTokenStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTokenStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PowerTokenStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTokenStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PowerTokenStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTokenStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowerTokenStorageSession struct {
	Contract     *PowerTokenStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PowerTokenStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowerTokenStorageCallerSession struct {
	Contract *PowerTokenStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PowerTokenStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowerTokenStorageTransactorSession struct {
	Contract     *PowerTokenStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PowerTokenStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PowerTokenStorageRaw struct {
	Contract *PowerTokenStorage // Generic contract binding to access the raw methods on
}

// PowerTokenStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowerTokenStorageCallerRaw struct {
	Contract *PowerTokenStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PowerTokenStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowerTokenStorageTransactorRaw struct {
	Contract *PowerTokenStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPowerTokenStorage creates a new instance of PowerTokenStorage, bound to a specific deployed contract.
func NewPowerTokenStorage(address common.Address, backend bind.ContractBackend) (*PowerTokenStorage, error) {
	contract, err := bindPowerTokenStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorage{PowerTokenStorageCaller: PowerTokenStorageCaller{contract: contract}, PowerTokenStorageTransactor: PowerTokenStorageTransactor{contract: contract}, PowerTokenStorageFilterer: PowerTokenStorageFilterer{contract: contract}}, nil
}

// NewPowerTokenStorageCaller creates a new read-only instance of PowerTokenStorage, bound to a specific deployed contract.
func NewPowerTokenStorageCaller(address common.Address, caller bind.ContractCaller) (*PowerTokenStorageCaller, error) {
	contract, err := bindPowerTokenStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageCaller{contract: contract}, nil
}

// NewPowerTokenStorageTransactor creates a new write-only instance of PowerTokenStorage, bound to a specific deployed contract.
func NewPowerTokenStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PowerTokenStorageTransactor, error) {
	contract, err := bindPowerTokenStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageTransactor{contract: contract}, nil
}

// NewPowerTokenStorageFilterer creates a new log filterer instance of PowerTokenStorage, bound to a specific deployed contract.
func NewPowerTokenStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PowerTokenStorageFilterer, error) {
	contract, err := bindPowerTokenStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageFilterer{contract: contract}, nil
}

// bindPowerTokenStorage binds a generic wrapper to an already deployed contract.
func bindPowerTokenStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerTokenStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PowerTokenStorage *PowerTokenStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PowerTokenStorage.Contract.PowerTokenStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PowerTokenStorage *PowerTokenStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.PowerTokenStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PowerTokenStorage *PowerTokenStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.PowerTokenStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PowerTokenStorage *PowerTokenStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PowerTokenStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PowerTokenStorage *PowerTokenStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PowerTokenStorage *PowerTokenStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAllowsPerpetual is a free data retrieval call binding the contract method 0x927c9d87.
//
// Solidity: function getAllowsPerpetual() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetAllowsPerpetual(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getAllowsPerpetual")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAllowsPerpetual is a free data retrieval call binding the contract method 0x927c9d87.
//
// Solidity: function getAllowsPerpetual() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageSession) GetAllowsPerpetual() (bool, error) {
	return _PowerTokenStorage.Contract.GetAllowsPerpetual(&_PowerTokenStorage.CallOpts)
}

// GetAllowsPerpetual is a free data retrieval call binding the contract method 0x927c9d87.
//
// Solidity: function getAllowsPerpetual() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetAllowsPerpetual() (bool, error) {
	return _PowerTokenStorage.Contract.GetAllowsPerpetual(&_PowerTokenStorage.CallOpts)
}

// GetBaseRate is a free data retrieval call binding the contract method 0xb655d0c4.
//
// Solidity: function getBaseRate() view returns(uint112)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetBaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getBaseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseRate is a free data retrieval call binding the contract method 0xb655d0c4.
//
// Solidity: function getBaseRate() view returns(uint112)
func (_PowerTokenStorage *PowerTokenStorageSession) GetBaseRate() (*big.Int, error) {
	return _PowerTokenStorage.Contract.GetBaseRate(&_PowerTokenStorage.CallOpts)
}

// GetBaseRate is a free data retrieval call binding the contract method 0xb655d0c4.
//
// Solidity: function getBaseRate() view returns(uint112)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetBaseRate() (*big.Int, error) {
	return _PowerTokenStorage.Contract.GetBaseRate(&_PowerTokenStorage.CallOpts)
}

// GetBaseToken is a free data retrieval call binding the contract method 0x98acd7a6.
//
// Solidity: function getBaseToken() view returns(address)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetBaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getBaseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBaseToken is a free data retrieval call binding the contract method 0x98acd7a6.
//
// Solidity: function getBaseToken() view returns(address)
func (_PowerTokenStorage *PowerTokenStorageSession) GetBaseToken() (common.Address, error) {
	return _PowerTokenStorage.Contract.GetBaseToken(&_PowerTokenStorage.CallOpts)
}

// GetBaseToken is a free data retrieval call binding the contract method 0x98acd7a6.
//
// Solidity: function getBaseToken() view returns(address)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetBaseToken() (common.Address, error) {
	return _PowerTokenStorage.Contract.GetBaseToken(&_PowerTokenStorage.CallOpts)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetEnterprise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getEnterprise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_PowerTokenStorage *PowerTokenStorageSession) GetEnterprise() (common.Address, error) {
	return _PowerTokenStorage.Contract.GetEnterprise(&_PowerTokenStorage.CallOpts)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetEnterprise() (common.Address, error) {
	return _PowerTokenStorage.Contract.GetEnterprise(&_PowerTokenStorage.CallOpts)
}

// GetGapHalvingPeriod is a free data retrieval call binding the contract method 0x42bb702e.
//
// Solidity: function getGapHalvingPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetGapHalvingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getGapHalvingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGapHalvingPeriod is a free data retrieval call binding the contract method 0x42bb702e.
//
// Solidity: function getGapHalvingPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageSession) GetGapHalvingPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetGapHalvingPeriod(&_PowerTokenStorage.CallOpts)
}

// GetGapHalvingPeriod is a free data retrieval call binding the contract method 0x42bb702e.
//
// Solidity: function getGapHalvingPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetGapHalvingPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetGapHalvingPeriod(&_PowerTokenStorage.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x81045ead.
//
// Solidity: function getIndex() view returns(uint16)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetIndex(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getIndex")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0x81045ead.
//
// Solidity: function getIndex() view returns(uint16)
func (_PowerTokenStorage *PowerTokenStorageSession) GetIndex() (uint16, error) {
	return _PowerTokenStorage.Contract.GetIndex(&_PowerTokenStorage.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x81045ead.
//
// Solidity: function getIndex() view returns(uint16)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetIndex() (uint16, error) {
	return _PowerTokenStorage.Contract.GetIndex(&_PowerTokenStorage.CallOpts)
}

// GetMaxLoanDuration is a free data retrieval call binding the contract method 0x5932f4ff.
//
// Solidity: function getMaxLoanDuration() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetMaxLoanDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getMaxLoanDuration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMaxLoanDuration is a free data retrieval call binding the contract method 0x5932f4ff.
//
// Solidity: function getMaxLoanDuration() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageSession) GetMaxLoanDuration() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMaxLoanDuration(&_PowerTokenStorage.CallOpts)
}

// GetMaxLoanDuration is a free data retrieval call binding the contract method 0x5932f4ff.
//
// Solidity: function getMaxLoanDuration() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetMaxLoanDuration() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMaxLoanDuration(&_PowerTokenStorage.CallOpts)
}

// GetMinGCFee is a free data retrieval call binding the contract method 0x0aff9d19.
//
// Solidity: function getMinGCFee() view returns(uint96)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetMinGCFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getMinGCFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinGCFee is a free data retrieval call binding the contract method 0x0aff9d19.
//
// Solidity: function getMinGCFee() view returns(uint96)
func (_PowerTokenStorage *PowerTokenStorageSession) GetMinGCFee() (*big.Int, error) {
	return _PowerTokenStorage.Contract.GetMinGCFee(&_PowerTokenStorage.CallOpts)
}

// GetMinGCFee is a free data retrieval call binding the contract method 0x0aff9d19.
//
// Solidity: function getMinGCFee() view returns(uint96)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetMinGCFee() (*big.Int, error) {
	return _PowerTokenStorage.Contract.GetMinGCFee(&_PowerTokenStorage.CallOpts)
}

// GetMinLoanDuration is a free data retrieval call binding the contract method 0x13e7e63b.
//
// Solidity: function getMinLoanDuration() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetMinLoanDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getMinLoanDuration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMinLoanDuration is a free data retrieval call binding the contract method 0x13e7e63b.
//
// Solidity: function getMinLoanDuration() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageSession) GetMinLoanDuration() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMinLoanDuration(&_PowerTokenStorage.CallOpts)
}

// GetMinLoanDuration is a free data retrieval call binding the contract method 0x13e7e63b.
//
// Solidity: function getMinLoanDuration() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetMinLoanDuration() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMinLoanDuration(&_PowerTokenStorage.CallOpts)
}

// GetServiceFeePercent is a free data retrieval call binding the contract method 0x6dfbee1b.
//
// Solidity: function getServiceFeePercent() view returns(uint16)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetServiceFeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getServiceFeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetServiceFeePercent is a free data retrieval call binding the contract method 0x6dfbee1b.
//
// Solidity: function getServiceFeePercent() view returns(uint16)
func (_PowerTokenStorage *PowerTokenStorageSession) GetServiceFeePercent() (uint16, error) {
	return _PowerTokenStorage.Contract.GetServiceFeePercent(&_PowerTokenStorage.CallOpts)
}

// GetServiceFeePercent is a free data retrieval call binding the contract method 0x6dfbee1b.
//
// Solidity: function getServiceFeePercent() view returns(uint16)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetServiceFeePercent() (uint16, error) {
	return _PowerTokenStorage.Contract.GetServiceFeePercent(&_PowerTokenStorage.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1bab58f5.
//
// Solidity: function getState(address account) view returns((uint112,uint112,uint32))
func (_PowerTokenStorage *PowerTokenStorageCaller) GetState(opts *bind.CallOpts, account common.Address) (PowerTokenStorageState, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getState", account)

	if err != nil {
		return *new(PowerTokenStorageState), err
	}

	out0 := *abi.ConvertType(out[0], new(PowerTokenStorageState)).(*PowerTokenStorageState)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1bab58f5.
//
// Solidity: function getState(address account) view returns((uint112,uint112,uint32))
func (_PowerTokenStorage *PowerTokenStorageSession) GetState(account common.Address) (PowerTokenStorageState, error) {
	return _PowerTokenStorage.Contract.GetState(&_PowerTokenStorage.CallOpts, account)
}

// GetState is a free data retrieval call binding the contract method 0x1bab58f5.
//
// Solidity: function getState(address account) view returns((uint112,uint112,uint32))
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetState(account common.Address) (PowerTokenStorageState, error) {
	return _PowerTokenStorage.Contract.GetState(&_PowerTokenStorage.CallOpts, account)
}

// IsAllowedLoanDuration is a free data retrieval call binding the contract method 0x6e4d8983.
//
// Solidity: function isAllowedLoanDuration(uint32 duration) view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCaller) IsAllowedLoanDuration(opts *bind.CallOpts, duration uint32) (bool, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "isAllowedLoanDuration", duration)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedLoanDuration is a free data retrieval call binding the contract method 0x6e4d8983.
//
// Solidity: function isAllowedLoanDuration(uint32 duration) view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageSession) IsAllowedLoanDuration(duration uint32) (bool, error) {
	return _PowerTokenStorage.Contract.IsAllowedLoanDuration(&_PowerTokenStorage.CallOpts, duration)
}

// IsAllowedLoanDuration is a free data retrieval call binding the contract method 0x6e4d8983.
//
// Solidity: function isAllowedLoanDuration(uint32 duration) view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) IsAllowedLoanDuration(duration uint32) (bool, error) {
	return _PowerTokenStorage.Contract.IsAllowedLoanDuration(&_PowerTokenStorage.CallOpts, duration)
}

// AllowPerpetualForever is a paid mutator transaction binding the contract method 0x27256486.
//
// Solidity: function allowPerpetualForever() returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) AllowPerpetualForever(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "allowPerpetualForever")
}

// AllowPerpetualForever is a paid mutator transaction binding the contract method 0x27256486.
//
// Solidity: function allowPerpetualForever() returns()
func (_PowerTokenStorage *PowerTokenStorageSession) AllowPerpetualForever() (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.AllowPerpetualForever(&_PowerTokenStorage.TransactOpts)
}

// AllowPerpetualForever is a paid mutator transaction binding the contract method 0x27256486.
//
// Solidity: function allowPerpetualForever() returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) AllowPerpetualForever() (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.AllowPerpetualForever(&_PowerTokenStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x2ef0cf8b.
//
// Solidity: function initialize(address enterprise, uint112 baseRate, uint96 minGCFee, uint32 gapHalvingPeriod, uint16 index, address baseToken, uint32 minLoanDuration, uint32 maxLoanDuration, uint16 serviceFeePercent, bool allowsPerpetual) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) Initialize(opts *bind.TransactOpts, enterprise common.Address, baseRate *big.Int, minGCFee *big.Int, gapHalvingPeriod uint32, index uint16, baseToken common.Address, minLoanDuration uint32, maxLoanDuration uint32, serviceFeePercent uint16, allowsPerpetual bool) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "initialize", enterprise, baseRate, minGCFee, gapHalvingPeriod, index, baseToken, minLoanDuration, maxLoanDuration, serviceFeePercent, allowsPerpetual)
}

// Initialize is a paid mutator transaction binding the contract method 0x2ef0cf8b.
//
// Solidity: function initialize(address enterprise, uint112 baseRate, uint96 minGCFee, uint32 gapHalvingPeriod, uint16 index, address baseToken, uint32 minLoanDuration, uint32 maxLoanDuration, uint16 serviceFeePercent, bool allowsPerpetual) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) Initialize(enterprise common.Address, baseRate *big.Int, minGCFee *big.Int, gapHalvingPeriod uint32, index uint16, baseToken common.Address, minLoanDuration uint32, maxLoanDuration uint32, serviceFeePercent uint16, allowsPerpetual bool) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.Initialize(&_PowerTokenStorage.TransactOpts, enterprise, baseRate, minGCFee, gapHalvingPeriod, index, baseToken, minLoanDuration, maxLoanDuration, serviceFeePercent, allowsPerpetual)
}

// Initialize is a paid mutator transaction binding the contract method 0x2ef0cf8b.
//
// Solidity: function initialize(address enterprise, uint112 baseRate, uint96 minGCFee, uint32 gapHalvingPeriod, uint16 index, address baseToken, uint32 minLoanDuration, uint32 maxLoanDuration, uint16 serviceFeePercent, bool allowsPerpetual) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) Initialize(enterprise common.Address, baseRate *big.Int, minGCFee *big.Int, gapHalvingPeriod uint32, index uint16, baseToken common.Address, minLoanDuration uint32, maxLoanDuration uint32, serviceFeePercent uint16, allowsPerpetual bool) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.Initialize(&_PowerTokenStorage.TransactOpts, enterprise, baseRate, minGCFee, gapHalvingPeriod, index, baseToken, minLoanDuration, maxLoanDuration, serviceFeePercent, allowsPerpetual)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) Initialize0(opts *bind.TransactOpts, enterprise common.Address) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "initialize0", enterprise)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) Initialize0(enterprise common.Address) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.Initialize0(&_PowerTokenStorage.TransactOpts, enterprise)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) Initialize0(enterprise common.Address) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.Initialize0(&_PowerTokenStorage.TransactOpts, enterprise)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x759c8be3.
//
// Solidity: function setBaseRate(uint112 baseRate, address baseToken, uint96 minGCFee) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) SetBaseRate(opts *bind.TransactOpts, baseRate *big.Int, baseToken common.Address, minGCFee *big.Int) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "setBaseRate", baseRate, baseToken, minGCFee)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x759c8be3.
//
// Solidity: function setBaseRate(uint112 baseRate, address baseToken, uint96 minGCFee) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) SetBaseRate(baseRate *big.Int, baseToken common.Address, minGCFee *big.Int) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetBaseRate(&_PowerTokenStorage.TransactOpts, baseRate, baseToken, minGCFee)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x759c8be3.
//
// Solidity: function setBaseRate(uint112 baseRate, address baseToken, uint96 minGCFee) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) SetBaseRate(baseRate *big.Int, baseToken common.Address, minGCFee *big.Int) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetBaseRate(&_PowerTokenStorage.TransactOpts, baseRate, baseToken, minGCFee)
}

// SetLoanDurationLimits is a paid mutator transaction binding the contract method 0xc0434352.
//
// Solidity: function setLoanDurationLimits(uint32 minLoanDuration, uint32 maxLoanDuration) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) SetLoanDurationLimits(opts *bind.TransactOpts, minLoanDuration uint32, maxLoanDuration uint32) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "setLoanDurationLimits", minLoanDuration, maxLoanDuration)
}

// SetLoanDurationLimits is a paid mutator transaction binding the contract method 0xc0434352.
//
// Solidity: function setLoanDurationLimits(uint32 minLoanDuration, uint32 maxLoanDuration) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) SetLoanDurationLimits(minLoanDuration uint32, maxLoanDuration uint32) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetLoanDurationLimits(&_PowerTokenStorage.TransactOpts, minLoanDuration, maxLoanDuration)
}

// SetLoanDurationLimits is a paid mutator transaction binding the contract method 0xc0434352.
//
// Solidity: function setLoanDurationLimits(uint32 minLoanDuration, uint32 maxLoanDuration) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) SetLoanDurationLimits(minLoanDuration uint32, maxLoanDuration uint32) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetLoanDurationLimits(&_PowerTokenStorage.TransactOpts, minLoanDuration, maxLoanDuration)
}

// SetServiceFeePercent is a paid mutator transaction binding the contract method 0x0d148ab6.
//
// Solidity: function setServiceFeePercent(uint16 newServiceFeePercent) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) SetServiceFeePercent(opts *bind.TransactOpts, newServiceFeePercent uint16) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "setServiceFeePercent", newServiceFeePercent)
}

// SetServiceFeePercent is a paid mutator transaction binding the contract method 0x0d148ab6.
//
// Solidity: function setServiceFeePercent(uint16 newServiceFeePercent) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) SetServiceFeePercent(newServiceFeePercent uint16) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetServiceFeePercent(&_PowerTokenStorage.TransactOpts, newServiceFeePercent)
}

// SetServiceFeePercent is a paid mutator transaction binding the contract method 0x0d148ab6.
//
// Solidity: function setServiceFeePercent(uint16 newServiceFeePercent) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) SetServiceFeePercent(newServiceFeePercent uint16) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetServiceFeePercent(&_PowerTokenStorage.TransactOpts, newServiceFeePercent)
}

// PowerTokenStorageBaseRateChangedIterator is returned from FilterBaseRateChanged and is used to iterate over the raw logs and unpacked data for BaseRateChanged events raised by the PowerTokenStorage contract.
type PowerTokenStorageBaseRateChangedIterator struct {
	Event *PowerTokenStorageBaseRateChanged // Event containing the contract specifics and raw log

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
func (it *PowerTokenStorageBaseRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStorageBaseRateChanged)
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
		it.Event = new(PowerTokenStorageBaseRateChanged)
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
func (it *PowerTokenStorageBaseRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStorageBaseRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStorageBaseRateChanged represents a BaseRateChanged event raised by the PowerTokenStorage contract.
type PowerTokenStorageBaseRateChanged struct {
	BaseRate  *big.Int
	BaseToken common.Address
	MinGCFee  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseRateChanged is a free log retrieval operation binding the contract event 0x2833a6b2910014b082775316936860d808d821fc59634fa08bc19c19427fa99d.
//
// Solidity: event BaseRateChanged(uint112 baseRate, address baseToken, uint96 minGCFee)
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterBaseRateChanged(opts *bind.FilterOpts) (*PowerTokenStorageBaseRateChangedIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "BaseRateChanged")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageBaseRateChangedIterator{contract: _PowerTokenStorage.contract, event: "BaseRateChanged", logs: logs, sub: sub}, nil
}

// WatchBaseRateChanged is a free log subscription operation binding the contract event 0x2833a6b2910014b082775316936860d808d821fc59634fa08bc19c19427fa99d.
//
// Solidity: event BaseRateChanged(uint112 baseRate, address baseToken, uint96 minGCFee)
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchBaseRateChanged(opts *bind.WatchOpts, sink chan<- *PowerTokenStorageBaseRateChanged) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "BaseRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStorageBaseRateChanged)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "BaseRateChanged", log); err != nil {
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

// ParseBaseRateChanged is a log parse operation binding the contract event 0x2833a6b2910014b082775316936860d808d821fc59634fa08bc19c19427fa99d.
//
// Solidity: event BaseRateChanged(uint112 baseRate, address baseToken, uint96 minGCFee)
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParseBaseRateChanged(log types.Log) (*PowerTokenStorageBaseRateChanged, error) {
	event := new(PowerTokenStorageBaseRateChanged)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "BaseRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowerTokenStorageLoanDurationLimitsChangedIterator is returned from FilterLoanDurationLimitsChanged and is used to iterate over the raw logs and unpacked data for LoanDurationLimitsChanged events raised by the PowerTokenStorage contract.
type PowerTokenStorageLoanDurationLimitsChangedIterator struct {
	Event *PowerTokenStorageLoanDurationLimitsChanged // Event containing the contract specifics and raw log

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
func (it *PowerTokenStorageLoanDurationLimitsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStorageLoanDurationLimitsChanged)
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
		it.Event = new(PowerTokenStorageLoanDurationLimitsChanged)
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
func (it *PowerTokenStorageLoanDurationLimitsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStorageLoanDurationLimitsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStorageLoanDurationLimitsChanged represents a LoanDurationLimitsChanged event raised by the PowerTokenStorage contract.
type PowerTokenStorageLoanDurationLimitsChanged struct {
	MinDuration uint32
	MaxDuration uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLoanDurationLimitsChanged is a free log retrieval operation binding the contract event 0x1ff66a89a0dcb0d94c8284280cbf0fb7190e7d935e15b6b48513e08b84fd9c7b.
//
// Solidity: event LoanDurationLimitsChanged(uint32 minDuration, uint32 maxDuration)
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterLoanDurationLimitsChanged(opts *bind.FilterOpts) (*PowerTokenStorageLoanDurationLimitsChangedIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "LoanDurationLimitsChanged")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageLoanDurationLimitsChangedIterator{contract: _PowerTokenStorage.contract, event: "LoanDurationLimitsChanged", logs: logs, sub: sub}, nil
}

// WatchLoanDurationLimitsChanged is a free log subscription operation binding the contract event 0x1ff66a89a0dcb0d94c8284280cbf0fb7190e7d935e15b6b48513e08b84fd9c7b.
//
// Solidity: event LoanDurationLimitsChanged(uint32 minDuration, uint32 maxDuration)
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchLoanDurationLimitsChanged(opts *bind.WatchOpts, sink chan<- *PowerTokenStorageLoanDurationLimitsChanged) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "LoanDurationLimitsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStorageLoanDurationLimitsChanged)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "LoanDurationLimitsChanged", log); err != nil {
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

// ParseLoanDurationLimitsChanged is a log parse operation binding the contract event 0x1ff66a89a0dcb0d94c8284280cbf0fb7190e7d935e15b6b48513e08b84fd9c7b.
//
// Solidity: event LoanDurationLimitsChanged(uint32 minDuration, uint32 maxDuration)
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParseLoanDurationLimitsChanged(log types.Log) (*PowerTokenStorageLoanDurationLimitsChanged, error) {
	event := new(PowerTokenStorageLoanDurationLimitsChanged)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "LoanDurationLimitsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowerTokenStoragePerpetualAllowedIterator is returned from FilterPerpetualAllowed and is used to iterate over the raw logs and unpacked data for PerpetualAllowed events raised by the PowerTokenStorage contract.
type PowerTokenStoragePerpetualAllowedIterator struct {
	Event *PowerTokenStoragePerpetualAllowed // Event containing the contract specifics and raw log

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
func (it *PowerTokenStoragePerpetualAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStoragePerpetualAllowed)
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
		it.Event = new(PowerTokenStoragePerpetualAllowed)
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
func (it *PowerTokenStoragePerpetualAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStoragePerpetualAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStoragePerpetualAllowed represents a PerpetualAllowed event raised by the PowerTokenStorage contract.
type PowerTokenStoragePerpetualAllowed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPerpetualAllowed is a free log retrieval operation binding the contract event 0x772fb13db6d1f0f35034c4b28dbbb01644e39d7e8b38de2a93b3c9d1fbb1ad4a.
//
// Solidity: event PerpetualAllowed()
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterPerpetualAllowed(opts *bind.FilterOpts) (*PowerTokenStoragePerpetualAllowedIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "PerpetualAllowed")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStoragePerpetualAllowedIterator{contract: _PowerTokenStorage.contract, event: "PerpetualAllowed", logs: logs, sub: sub}, nil
}

// WatchPerpetualAllowed is a free log subscription operation binding the contract event 0x772fb13db6d1f0f35034c4b28dbbb01644e39d7e8b38de2a93b3c9d1fbb1ad4a.
//
// Solidity: event PerpetualAllowed()
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchPerpetualAllowed(opts *bind.WatchOpts, sink chan<- *PowerTokenStoragePerpetualAllowed) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "PerpetualAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStoragePerpetualAllowed)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "PerpetualAllowed", log); err != nil {
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

// ParsePerpetualAllowed is a log parse operation binding the contract event 0x772fb13db6d1f0f35034c4b28dbbb01644e39d7e8b38de2a93b3c9d1fbb1ad4a.
//
// Solidity: event PerpetualAllowed()
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParsePerpetualAllowed(log types.Log) (*PowerTokenStoragePerpetualAllowed, error) {
	event := new(PowerTokenStoragePerpetualAllowed)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "PerpetualAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowerTokenStorageServiceFeePercentChangedIterator is returned from FilterServiceFeePercentChanged and is used to iterate over the raw logs and unpacked data for ServiceFeePercentChanged events raised by the PowerTokenStorage contract.
type PowerTokenStorageServiceFeePercentChangedIterator struct {
	Event *PowerTokenStorageServiceFeePercentChanged // Event containing the contract specifics and raw log

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
func (it *PowerTokenStorageServiceFeePercentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStorageServiceFeePercentChanged)
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
		it.Event = new(PowerTokenStorageServiceFeePercentChanged)
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
func (it *PowerTokenStorageServiceFeePercentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStorageServiceFeePercentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStorageServiceFeePercentChanged represents a ServiceFeePercentChanged event raised by the PowerTokenStorage contract.
type PowerTokenStorageServiceFeePercentChanged struct {
	Percent uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterServiceFeePercentChanged is a free log retrieval operation binding the contract event 0x7b5d5b53df75c87d5bea0f0535cd81a355db537ee2ca5dcea9849433d99847b9.
//
// Solidity: event ServiceFeePercentChanged(uint16 percent)
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterServiceFeePercentChanged(opts *bind.FilterOpts) (*PowerTokenStorageServiceFeePercentChangedIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "ServiceFeePercentChanged")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageServiceFeePercentChangedIterator{contract: _PowerTokenStorage.contract, event: "ServiceFeePercentChanged", logs: logs, sub: sub}, nil
}

// WatchServiceFeePercentChanged is a free log subscription operation binding the contract event 0x7b5d5b53df75c87d5bea0f0535cd81a355db537ee2ca5dcea9849433d99847b9.
//
// Solidity: event ServiceFeePercentChanged(uint16 percent)
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchServiceFeePercentChanged(opts *bind.WatchOpts, sink chan<- *PowerTokenStorageServiceFeePercentChanged) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "ServiceFeePercentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStorageServiceFeePercentChanged)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "ServiceFeePercentChanged", log); err != nil {
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

// ParseServiceFeePercentChanged is a log parse operation binding the contract event 0x7b5d5b53df75c87d5bea0f0535cd81a355db537ee2ca5dcea9849433d99847b9.
//
// Solidity: event ServiceFeePercentChanged(uint16 percent)
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParseServiceFeePercentChanged(log types.Log) (*PowerTokenStorageServiceFeePercentChanged, error) {
	event := new(PowerTokenStorageServiceFeePercentChanged)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "ServiceFeePercentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
