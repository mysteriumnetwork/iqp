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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"}],\"name\":\"BaseRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"minRentalPeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxRentalPeriod\",\"type\":\"uint32\"}],\"name\":\"RentalPeriodLimitsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"percent\",\"type\":\"uint16\"}],\"name\":\"ServiceFeePercentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwappingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TransferEnabled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"enableSwappingForever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTransferForever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseRate\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnergyGapHalvingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterprise\",\"outputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxRentalPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinGCFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinRentalPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getServiceFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"lockedBalance\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"energy\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structPowerTokenStorage.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"},{\"internalType\":\"uint16\",\"name\":\"serviceFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"energyGapHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minRentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxRentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"swappingEnabled\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"isAllowedRentalPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSwappingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTransferEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"}],\"name\":\"setBaseRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"minRentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxRentalPeriod\",\"type\":\"uint32\"}],\"name\":\"setRentalPeriodLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newServiceFeePercent\",\"type\":\"uint16\"}],\"name\":\"setServiceFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetEnergyGapHalvingPeriod is a free data retrieval call binding the contract method 0xcec2f246.
//
// Solidity: function getEnergyGapHalvingPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetEnergyGapHalvingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getEnergyGapHalvingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEnergyGapHalvingPeriod is a free data retrieval call binding the contract method 0xcec2f246.
//
// Solidity: function getEnergyGapHalvingPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageSession) GetEnergyGapHalvingPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetEnergyGapHalvingPeriod(&_PowerTokenStorage.CallOpts)
}

// GetEnergyGapHalvingPeriod is a free data retrieval call binding the contract method 0xcec2f246.
//
// Solidity: function getEnergyGapHalvingPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetEnergyGapHalvingPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetEnergyGapHalvingPeriod(&_PowerTokenStorage.CallOpts)
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

// GetMaxRentalPeriod is a free data retrieval call binding the contract method 0xd29cc7c0.
//
// Solidity: function getMaxRentalPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetMaxRentalPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getMaxRentalPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMaxRentalPeriod is a free data retrieval call binding the contract method 0xd29cc7c0.
//
// Solidity: function getMaxRentalPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageSession) GetMaxRentalPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMaxRentalPeriod(&_PowerTokenStorage.CallOpts)
}

// GetMaxRentalPeriod is a free data retrieval call binding the contract method 0xd29cc7c0.
//
// Solidity: function getMaxRentalPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetMaxRentalPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMaxRentalPeriod(&_PowerTokenStorage.CallOpts)
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

// GetMinRentalPeriod is a free data retrieval call binding the contract method 0xfcc4b67d.
//
// Solidity: function getMinRentalPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCaller) GetMinRentalPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "getMinRentalPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMinRentalPeriod is a free data retrieval call binding the contract method 0xfcc4b67d.
//
// Solidity: function getMinRentalPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageSession) GetMinRentalPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMinRentalPeriod(&_PowerTokenStorage.CallOpts)
}

// GetMinRentalPeriod is a free data retrieval call binding the contract method 0xfcc4b67d.
//
// Solidity: function getMinRentalPeriod() view returns(uint32)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) GetMinRentalPeriod() (uint32, error) {
	return _PowerTokenStorage.Contract.GetMinRentalPeriod(&_PowerTokenStorage.CallOpts)
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

// IsAllowedRentalPeriod is a free data retrieval call binding the contract method 0x70005557.
//
// Solidity: function isAllowedRentalPeriod(uint32 period) view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCaller) IsAllowedRentalPeriod(opts *bind.CallOpts, period uint32) (bool, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "isAllowedRentalPeriod", period)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedRentalPeriod is a free data retrieval call binding the contract method 0x70005557.
//
// Solidity: function isAllowedRentalPeriod(uint32 period) view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageSession) IsAllowedRentalPeriod(period uint32) (bool, error) {
	return _PowerTokenStorage.Contract.IsAllowedRentalPeriod(&_PowerTokenStorage.CallOpts, period)
}

// IsAllowedRentalPeriod is a free data retrieval call binding the contract method 0x70005557.
//
// Solidity: function isAllowedRentalPeriod(uint32 period) view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) IsAllowedRentalPeriod(period uint32) (bool, error) {
	return _PowerTokenStorage.Contract.IsAllowedRentalPeriod(&_PowerTokenStorage.CallOpts, period)
}

// IsSwappingEnabled is a free data retrieval call binding the contract method 0x4aa5028b.
//
// Solidity: function isSwappingEnabled() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCaller) IsSwappingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "isSwappingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwappingEnabled is a free data retrieval call binding the contract method 0x4aa5028b.
//
// Solidity: function isSwappingEnabled() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageSession) IsSwappingEnabled() (bool, error) {
	return _PowerTokenStorage.Contract.IsSwappingEnabled(&_PowerTokenStorage.CallOpts)
}

// IsSwappingEnabled is a free data retrieval call binding the contract method 0x4aa5028b.
//
// Solidity: function isSwappingEnabled() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) IsSwappingEnabled() (bool, error) {
	return _PowerTokenStorage.Contract.IsSwappingEnabled(&_PowerTokenStorage.CallOpts)
}

// IsTransferEnabled is a free data retrieval call binding the contract method 0xcca5dcb6.
//
// Solidity: function isTransferEnabled() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCaller) IsTransferEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PowerTokenStorage.contract.Call(opts, &out, "isTransferEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferEnabled is a free data retrieval call binding the contract method 0xcca5dcb6.
//
// Solidity: function isTransferEnabled() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageSession) IsTransferEnabled() (bool, error) {
	return _PowerTokenStorage.Contract.IsTransferEnabled(&_PowerTokenStorage.CallOpts)
}

// IsTransferEnabled is a free data retrieval call binding the contract method 0xcca5dcb6.
//
// Solidity: function isTransferEnabled() view returns(bool)
func (_PowerTokenStorage *PowerTokenStorageCallerSession) IsTransferEnabled() (bool, error) {
	return _PowerTokenStorage.Contract.IsTransferEnabled(&_PowerTokenStorage.CallOpts)
}

// EnableSwappingForever is a paid mutator transaction binding the contract method 0x35528ee0.
//
// Solidity: function enableSwappingForever() returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) EnableSwappingForever(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "enableSwappingForever")
}

// EnableSwappingForever is a paid mutator transaction binding the contract method 0x35528ee0.
//
// Solidity: function enableSwappingForever() returns()
func (_PowerTokenStorage *PowerTokenStorageSession) EnableSwappingForever() (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.EnableSwappingForever(&_PowerTokenStorage.TransactOpts)
}

// EnableSwappingForever is a paid mutator transaction binding the contract method 0x35528ee0.
//
// Solidity: function enableSwappingForever() returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) EnableSwappingForever() (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.EnableSwappingForever(&_PowerTokenStorage.TransactOpts)
}

// EnableTransferForever is a paid mutator transaction binding the contract method 0x659262a0.
//
// Solidity: function enableTransferForever() returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) EnableTransferForever(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "enableTransferForever")
}

// EnableTransferForever is a paid mutator transaction binding the contract method 0x659262a0.
//
// Solidity: function enableTransferForever() returns()
func (_PowerTokenStorage *PowerTokenStorageSession) EnableTransferForever() (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.EnableTransferForever(&_PowerTokenStorage.TransactOpts)
}

// EnableTransferForever is a paid mutator transaction binding the contract method 0x659262a0.
//
// Solidity: function enableTransferForever() returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) EnableTransferForever() (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.EnableTransferForever(&_PowerTokenStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xb29c98ad.
//
// Solidity: function initialize(address enterprise, address baseToken, uint112 baseRate, uint96 minGCFee, uint16 serviceFeePercent, uint32 energyGapHalvingPeriod, uint16 index, uint32 minRentalPeriod, uint32 maxRentalPeriod, bool swappingEnabled) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) Initialize(opts *bind.TransactOpts, enterprise common.Address, baseToken common.Address, baseRate *big.Int, minGCFee *big.Int, serviceFeePercent uint16, energyGapHalvingPeriod uint32, index uint16, minRentalPeriod uint32, maxRentalPeriod uint32, swappingEnabled bool) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "initialize", enterprise, baseToken, baseRate, minGCFee, serviceFeePercent, energyGapHalvingPeriod, index, minRentalPeriod, maxRentalPeriod, swappingEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xb29c98ad.
//
// Solidity: function initialize(address enterprise, address baseToken, uint112 baseRate, uint96 minGCFee, uint16 serviceFeePercent, uint32 energyGapHalvingPeriod, uint16 index, uint32 minRentalPeriod, uint32 maxRentalPeriod, bool swappingEnabled) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) Initialize(enterprise common.Address, baseToken common.Address, baseRate *big.Int, minGCFee *big.Int, serviceFeePercent uint16, energyGapHalvingPeriod uint32, index uint16, minRentalPeriod uint32, maxRentalPeriod uint32, swappingEnabled bool) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.Initialize(&_PowerTokenStorage.TransactOpts, enterprise, baseToken, baseRate, minGCFee, serviceFeePercent, energyGapHalvingPeriod, index, minRentalPeriod, maxRentalPeriod, swappingEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xb29c98ad.
//
// Solidity: function initialize(address enterprise, address baseToken, uint112 baseRate, uint96 minGCFee, uint16 serviceFeePercent, uint32 energyGapHalvingPeriod, uint16 index, uint32 minRentalPeriod, uint32 maxRentalPeriod, bool swappingEnabled) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) Initialize(enterprise common.Address, baseToken common.Address, baseRate *big.Int, minGCFee *big.Int, serviceFeePercent uint16, energyGapHalvingPeriod uint32, index uint16, minRentalPeriod uint32, maxRentalPeriod uint32, swappingEnabled bool) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.Initialize(&_PowerTokenStorage.TransactOpts, enterprise, baseToken, baseRate, minGCFee, serviceFeePercent, energyGapHalvingPeriod, index, minRentalPeriod, maxRentalPeriod, swappingEnabled)
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

// SetRentalPeriodLimits is a paid mutator transaction binding the contract method 0x6e4dba76.
//
// Solidity: function setRentalPeriodLimits(uint32 minRentalPeriod, uint32 maxRentalPeriod) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactor) SetRentalPeriodLimits(opts *bind.TransactOpts, minRentalPeriod uint32, maxRentalPeriod uint32) (*types.Transaction, error) {
	return _PowerTokenStorage.contract.Transact(opts, "setRentalPeriodLimits", minRentalPeriod, maxRentalPeriod)
}

// SetRentalPeriodLimits is a paid mutator transaction binding the contract method 0x6e4dba76.
//
// Solidity: function setRentalPeriodLimits(uint32 minRentalPeriod, uint32 maxRentalPeriod) returns()
func (_PowerTokenStorage *PowerTokenStorageSession) SetRentalPeriodLimits(minRentalPeriod uint32, maxRentalPeriod uint32) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetRentalPeriodLimits(&_PowerTokenStorage.TransactOpts, minRentalPeriod, maxRentalPeriod)
}

// SetRentalPeriodLimits is a paid mutator transaction binding the contract method 0x6e4dba76.
//
// Solidity: function setRentalPeriodLimits(uint32 minRentalPeriod, uint32 maxRentalPeriod) returns()
func (_PowerTokenStorage *PowerTokenStorageTransactorSession) SetRentalPeriodLimits(minRentalPeriod uint32, maxRentalPeriod uint32) (*types.Transaction, error) {
	return _PowerTokenStorage.Contract.SetRentalPeriodLimits(&_PowerTokenStorage.TransactOpts, minRentalPeriod, maxRentalPeriod)
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

// PowerTokenStorageRentalPeriodLimitsChangedIterator is returned from FilterRentalPeriodLimitsChanged and is used to iterate over the raw logs and unpacked data for RentalPeriodLimitsChanged events raised by the PowerTokenStorage contract.
type PowerTokenStorageRentalPeriodLimitsChangedIterator struct {
	Event *PowerTokenStorageRentalPeriodLimitsChanged // Event containing the contract specifics and raw log

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
func (it *PowerTokenStorageRentalPeriodLimitsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStorageRentalPeriodLimitsChanged)
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
		it.Event = new(PowerTokenStorageRentalPeriodLimitsChanged)
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
func (it *PowerTokenStorageRentalPeriodLimitsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStorageRentalPeriodLimitsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStorageRentalPeriodLimitsChanged represents a RentalPeriodLimitsChanged event raised by the PowerTokenStorage contract.
type PowerTokenStorageRentalPeriodLimitsChanged struct {
	MinRentalPeriod uint32
	MaxRentalPeriod uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRentalPeriodLimitsChanged is a free log retrieval operation binding the contract event 0x61db23004a2218e8ecbda348735cc62eaaf3ca857e2705615a8d91acf6d7b817.
//
// Solidity: event RentalPeriodLimitsChanged(uint32 minRentalPeriod, uint32 maxRentalPeriod)
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterRentalPeriodLimitsChanged(opts *bind.FilterOpts) (*PowerTokenStorageRentalPeriodLimitsChangedIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "RentalPeriodLimitsChanged")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageRentalPeriodLimitsChangedIterator{contract: _PowerTokenStorage.contract, event: "RentalPeriodLimitsChanged", logs: logs, sub: sub}, nil
}

// WatchRentalPeriodLimitsChanged is a free log subscription operation binding the contract event 0x61db23004a2218e8ecbda348735cc62eaaf3ca857e2705615a8d91acf6d7b817.
//
// Solidity: event RentalPeriodLimitsChanged(uint32 minRentalPeriod, uint32 maxRentalPeriod)
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchRentalPeriodLimitsChanged(opts *bind.WatchOpts, sink chan<- *PowerTokenStorageRentalPeriodLimitsChanged) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "RentalPeriodLimitsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStorageRentalPeriodLimitsChanged)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "RentalPeriodLimitsChanged", log); err != nil {
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

// ParseRentalPeriodLimitsChanged is a log parse operation binding the contract event 0x61db23004a2218e8ecbda348735cc62eaaf3ca857e2705615a8d91acf6d7b817.
//
// Solidity: event RentalPeriodLimitsChanged(uint32 minRentalPeriod, uint32 maxRentalPeriod)
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParseRentalPeriodLimitsChanged(log types.Log) (*PowerTokenStorageRentalPeriodLimitsChanged, error) {
	event := new(PowerTokenStorageRentalPeriodLimitsChanged)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "RentalPeriodLimitsChanged", log); err != nil {
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

// PowerTokenStorageSwappingEnabledIterator is returned from FilterSwappingEnabled and is used to iterate over the raw logs and unpacked data for SwappingEnabled events raised by the PowerTokenStorage contract.
type PowerTokenStorageSwappingEnabledIterator struct {
	Event *PowerTokenStorageSwappingEnabled // Event containing the contract specifics and raw log

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
func (it *PowerTokenStorageSwappingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStorageSwappingEnabled)
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
		it.Event = new(PowerTokenStorageSwappingEnabled)
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
func (it *PowerTokenStorageSwappingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStorageSwappingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStorageSwappingEnabled represents a SwappingEnabled event raised by the PowerTokenStorage contract.
type PowerTokenStorageSwappingEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwappingEnabled is a free log retrieval operation binding the contract event 0x4f8e8d44a98a830b5142e61afa1cc8dacbf21291bcd657e26d2c7780e2c7fac4.
//
// Solidity: event SwappingEnabled()
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterSwappingEnabled(opts *bind.FilterOpts) (*PowerTokenStorageSwappingEnabledIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "SwappingEnabled")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageSwappingEnabledIterator{contract: _PowerTokenStorage.contract, event: "SwappingEnabled", logs: logs, sub: sub}, nil
}

// WatchSwappingEnabled is a free log subscription operation binding the contract event 0x4f8e8d44a98a830b5142e61afa1cc8dacbf21291bcd657e26d2c7780e2c7fac4.
//
// Solidity: event SwappingEnabled()
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchSwappingEnabled(opts *bind.WatchOpts, sink chan<- *PowerTokenStorageSwappingEnabled) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "SwappingEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStorageSwappingEnabled)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "SwappingEnabled", log); err != nil {
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

// ParseSwappingEnabled is a log parse operation binding the contract event 0x4f8e8d44a98a830b5142e61afa1cc8dacbf21291bcd657e26d2c7780e2c7fac4.
//
// Solidity: event SwappingEnabled()
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParseSwappingEnabled(log types.Log) (*PowerTokenStorageSwappingEnabled, error) {
	event := new(PowerTokenStorageSwappingEnabled)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "SwappingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowerTokenStorageTransferEnabledIterator is returned from FilterTransferEnabled and is used to iterate over the raw logs and unpacked data for TransferEnabled events raised by the PowerTokenStorage contract.
type PowerTokenStorageTransferEnabledIterator struct {
	Event *PowerTokenStorageTransferEnabled // Event containing the contract specifics and raw log

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
func (it *PowerTokenStorageTransferEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowerTokenStorageTransferEnabled)
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
		it.Event = new(PowerTokenStorageTransferEnabled)
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
func (it *PowerTokenStorageTransferEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowerTokenStorageTransferEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowerTokenStorageTransferEnabled represents a TransferEnabled event raised by the PowerTokenStorage contract.
type PowerTokenStorageTransferEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransferEnabled is a free log retrieval operation binding the contract event 0x75fce015c314a132947a3e42f6ab79ab8e05397dabf35b4d742dea228bbadc2d.
//
// Solidity: event TransferEnabled()
func (_PowerTokenStorage *PowerTokenStorageFilterer) FilterTransferEnabled(opts *bind.FilterOpts) (*PowerTokenStorageTransferEnabledIterator, error) {

	logs, sub, err := _PowerTokenStorage.contract.FilterLogs(opts, "TransferEnabled")
	if err != nil {
		return nil, err
	}
	return &PowerTokenStorageTransferEnabledIterator{contract: _PowerTokenStorage.contract, event: "TransferEnabled", logs: logs, sub: sub}, nil
}

// WatchTransferEnabled is a free log subscription operation binding the contract event 0x75fce015c314a132947a3e42f6ab79ab8e05397dabf35b4d742dea228bbadc2d.
//
// Solidity: event TransferEnabled()
func (_PowerTokenStorage *PowerTokenStorageFilterer) WatchTransferEnabled(opts *bind.WatchOpts, sink chan<- *PowerTokenStorageTransferEnabled) (event.Subscription, error) {

	logs, sub, err := _PowerTokenStorage.contract.WatchLogs(opts, "TransferEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowerTokenStorageTransferEnabled)
				if err := _PowerTokenStorage.contract.UnpackLog(event, "TransferEnabled", log); err != nil {
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

// ParseTransferEnabled is a log parse operation binding the contract event 0x75fce015c314a132947a3e42f6ab79ab8e05397dabf35b4d742dea228bbadc2d.
//
// Solidity: event TransferEnabled()
func (_PowerTokenStorage *PowerTokenStorageFilterer) ParseTransferEnabled(log types.Log) (*PowerTokenStorageTransferEnabled, error) {
	event := new(PowerTokenStorageTransferEnabled)
	if err := _PowerTokenStorage.contract.UnpackLog(event, "TransferEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
