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

// EnterpriseStorageMetaData contains all meta data concerning the EnterpriseStorage contract.
var EnterpriseStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"BaseUriChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"name\":\"BondingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"BorrowerLoanReturnGracePeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"converter\",\"type\":\"address\"}],\"name\":\"ConverterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"}],\"name\":\"EnterpriseCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"EnterpriseLoanCollectGracePeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnterpriseShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"EnterpriseVaultChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"}],\"name\":\"FixedReserveChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"percent\",\"type\":\"uint16\"}],\"name\":\"GcFeePercentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"InterestGapHalvingPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PaymentTokenChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"streamingReserve\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"streamingReserveTarget\",\"type\":\"uint112\"}],\"name\":\"StreamingReserveChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"TotalSharesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"}],\"name\":\"UsedReserveChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"disablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"enablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondingCurve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowToken\",\"outputs\":[{\"internalType\":\"contractIBorrowToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowerLoanReturnGracePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConverter\",\"outputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseLoanCollectGracePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"interestGapHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"borrowerLoanReturnGracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enterpriseLoanCollectGracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint112\",\"name\":\"streamingReserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"streamingReserveTarget\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"streamingReserveUpdated\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterestGapHalvingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterestToken\",\"outputs\":[{\"internalType\":\"contractIInterestToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"}],\"name\":\"getLiquidityInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"internalType\":\"structEnterpriseStorage.LiquidityInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidityToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"}],\"name\":\"getLoanInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"powerTokenIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"borrowingTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maturityTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"borrowerReturnGraceTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enterpriseCollectGraceTime\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"gcFee\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"gcFeeTokenIndex\",\"type\":\"uint16\"}],\"internalType\":\"structEnterpriseStorage.LoanInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPowerTokens\",\"outputs\":[{\"internalType\":\"contractPowerToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsedReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"enterpriseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"contractIConverter\",\"name\":\"converter\",\"type\":\"address\"},{\"internalType\":\"contractProxyAdmin\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractIInterestToken\",\"name\":\"interestToken\",\"type\":\"address\"},{\"internalType\":\"contractIBorrowToken\",\"name\":\"borrowToken\",\"type\":\"address\"}],\"name\":\"initializeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"powerToken\",\"type\":\"address\"}],\"name\":\"isRegisteredPowerToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedPaymentToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"paymentTokenIndex\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"name\":\"setBondingCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newPeriod\",\"type\":\"uint32\"}],\"name\":\"setBorrowerLoanReturnGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"newConverter\",\"type\":\"address\"}],\"name\":\"setConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollector\",\"type\":\"address\"}],\"name\":\"setEnterpriseCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newPeriod\",\"type\":\"uint32\"}],\"name\":\"setEnterpriseLoanCollectGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"setEnterpriseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newGcFeePercent\",\"type\":\"uint16\"}],\"name\":\"setGcFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"interestGapHalvingPeriod\",\"type\":\"uint32\"}],\"name\":\"setInterestGapHalvingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeBorrowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeEnterprise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeInterestToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"powerToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradePowerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EnterpriseStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseStorageMetaData.ABI instead.
var EnterpriseStorageABI = EnterpriseStorageMetaData.ABI

// EnterpriseStorage is an auto generated Go binding around an Ethereum contract.
type EnterpriseStorage struct {
	EnterpriseStorageCaller     // Read-only binding to the contract
	EnterpriseStorageTransactor // Write-only binding to the contract
	EnterpriseStorageFilterer   // Log filterer for contract events
}

// EnterpriseStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseStorageSession struct {
	Contract     *EnterpriseStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EnterpriseStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseStorageCallerSession struct {
	Contract *EnterpriseStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EnterpriseStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseStorageTransactorSession struct {
	Contract     *EnterpriseStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EnterpriseStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseStorageRaw struct {
	Contract *EnterpriseStorage // Generic contract binding to access the raw methods on
}

// EnterpriseStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseStorageCallerRaw struct {
	Contract *EnterpriseStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseStorageTransactorRaw struct {
	Contract *EnterpriseStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterpriseStorage creates a new instance of EnterpriseStorage, bound to a specific deployed contract.
func NewEnterpriseStorage(address common.Address, backend bind.ContractBackend) (*EnterpriseStorage, error) {
	contract, err := bindEnterpriseStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorage{EnterpriseStorageCaller: EnterpriseStorageCaller{contract: contract}, EnterpriseStorageTransactor: EnterpriseStorageTransactor{contract: contract}, EnterpriseStorageFilterer: EnterpriseStorageFilterer{contract: contract}}, nil
}

// NewEnterpriseStorageCaller creates a new read-only instance of EnterpriseStorage, bound to a specific deployed contract.
func NewEnterpriseStorageCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseStorageCaller, error) {
	contract, err := bindEnterpriseStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageCaller{contract: contract}, nil
}

// NewEnterpriseStorageTransactor creates a new write-only instance of EnterpriseStorage, bound to a specific deployed contract.
func NewEnterpriseStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseStorageTransactor, error) {
	contract, err := bindEnterpriseStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageTransactor{contract: contract}, nil
}

// NewEnterpriseStorageFilterer creates a new log filterer instance of EnterpriseStorage, bound to a specific deployed contract.
func NewEnterpriseStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseStorageFilterer, error) {
	contract, err := bindEnterpriseStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageFilterer{contract: contract}, nil
}

// bindEnterpriseStorage binds a generic wrapper to an already deployed contract.
func bindEnterpriseStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnterpriseStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseStorage *EnterpriseStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseStorage.Contract.EnterpriseStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseStorage *EnterpriseStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.EnterpriseStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseStorage *EnterpriseStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.EnterpriseStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseStorage *EnterpriseStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseStorage *EnterpriseStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseStorage *EnterpriseStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAvailableReserve is a free data retrieval call binding the contract method 0xaebb5041.
//
// Solidity: function getAvailableReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetAvailableReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getAvailableReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableReserve is a free data retrieval call binding the contract method 0xaebb5041.
//
// Solidity: function getAvailableReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageSession) GetAvailableReserve() (*big.Int, error) {
	return _EnterpriseStorage.Contract.GetAvailableReserve(&_EnterpriseStorage.CallOpts)
}

// GetAvailableReserve is a free data retrieval call binding the contract method 0xaebb5041.
//
// Solidity: function getAvailableReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetAvailableReserve() (*big.Int, error) {
	return _EnterpriseStorage.Contract.GetAvailableReserve(&_EnterpriseStorage.CallOpts)
}

// GetBaseUri is a free data retrieval call binding the contract method 0x0cac36b2.
//
// Solidity: function getBaseUri() view returns(string)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetBaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getBaseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseUri is a free data retrieval call binding the contract method 0x0cac36b2.
//
// Solidity: function getBaseUri() view returns(string)
func (_EnterpriseStorage *EnterpriseStorageSession) GetBaseUri() (string, error) {
	return _EnterpriseStorage.Contract.GetBaseUri(&_EnterpriseStorage.CallOpts)
}

// GetBaseUri is a free data retrieval call binding the contract method 0x0cac36b2.
//
// Solidity: function getBaseUri() view returns(string)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetBaseUri() (string, error) {
	return _EnterpriseStorage.Contract.GetBaseUri(&_EnterpriseStorage.CallOpts)
}

// GetBondingCurve is a free data retrieval call binding the contract method 0xfaac38ef.
//
// Solidity: function getBondingCurve() view returns(uint256 pole, uint256 slope)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetBondingCurve(opts *bind.CallOpts) (struct {
	Pole  *big.Int
	Slope *big.Int
}, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getBondingCurve")

	outstruct := new(struct {
		Pole  *big.Int
		Slope *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pole = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondingCurve is a free data retrieval call binding the contract method 0xfaac38ef.
//
// Solidity: function getBondingCurve() view returns(uint256 pole, uint256 slope)
func (_EnterpriseStorage *EnterpriseStorageSession) GetBondingCurve() (struct {
	Pole  *big.Int
	Slope *big.Int
}, error) {
	return _EnterpriseStorage.Contract.GetBondingCurve(&_EnterpriseStorage.CallOpts)
}

// GetBondingCurve is a free data retrieval call binding the contract method 0xfaac38ef.
//
// Solidity: function getBondingCurve() view returns(uint256 pole, uint256 slope)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetBondingCurve() (struct {
	Pole  *big.Int
	Slope *big.Int
}, error) {
	return _EnterpriseStorage.Contract.GetBondingCurve(&_EnterpriseStorage.CallOpts)
}

// GetBorrowToken is a free data retrieval call binding the contract method 0xc93d0b92.
//
// Solidity: function getBorrowToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetBorrowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getBorrowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowToken is a free data retrieval call binding the contract method 0xc93d0b92.
//
// Solidity: function getBorrowToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetBorrowToken() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetBorrowToken(&_EnterpriseStorage.CallOpts)
}

// GetBorrowToken is a free data retrieval call binding the contract method 0xc93d0b92.
//
// Solidity: function getBorrowToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetBorrowToken() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetBorrowToken(&_EnterpriseStorage.CallOpts)
}

// GetBorrowerLoanReturnGracePeriod is a free data retrieval call binding the contract method 0x8188d50d.
//
// Solidity: function getBorrowerLoanReturnGracePeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetBorrowerLoanReturnGracePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getBorrowerLoanReturnGracePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetBorrowerLoanReturnGracePeriod is a free data retrieval call binding the contract method 0x8188d50d.
//
// Solidity: function getBorrowerLoanReturnGracePeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageSession) GetBorrowerLoanReturnGracePeriod() (uint32, error) {
	return _EnterpriseStorage.Contract.GetBorrowerLoanReturnGracePeriod(&_EnterpriseStorage.CallOpts)
}

// GetBorrowerLoanReturnGracePeriod is a free data retrieval call binding the contract method 0x8188d50d.
//
// Solidity: function getBorrowerLoanReturnGracePeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetBorrowerLoanReturnGracePeriod() (uint32, error) {
	return _EnterpriseStorage.Contract.GetBorrowerLoanReturnGracePeriod(&_EnterpriseStorage.CallOpts)
}

// GetConverter is a free data retrieval call binding the contract method 0x2261b07f.
//
// Solidity: function getConverter() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetConverter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getConverter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConverter is a free data retrieval call binding the contract method 0x2261b07f.
//
// Solidity: function getConverter() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetConverter() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetConverter(&_EnterpriseStorage.CallOpts)
}

// GetConverter is a free data retrieval call binding the contract method 0x2261b07f.
//
// Solidity: function getConverter() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetConverter() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetConverter(&_EnterpriseStorage.CallOpts)
}

// GetEnterpriseCollector is a free data retrieval call binding the contract method 0xa554b26d.
//
// Solidity: function getEnterpriseCollector() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetEnterpriseCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getEnterpriseCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseCollector is a free data retrieval call binding the contract method 0xa554b26d.
//
// Solidity: function getEnterpriseCollector() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetEnterpriseCollector() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetEnterpriseCollector(&_EnterpriseStorage.CallOpts)
}

// GetEnterpriseCollector is a free data retrieval call binding the contract method 0xa554b26d.
//
// Solidity: function getEnterpriseCollector() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetEnterpriseCollector() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetEnterpriseCollector(&_EnterpriseStorage.CallOpts)
}

// GetEnterpriseLoanCollectGracePeriod is a free data retrieval call binding the contract method 0xa0efd118.
//
// Solidity: function getEnterpriseLoanCollectGracePeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetEnterpriseLoanCollectGracePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getEnterpriseLoanCollectGracePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEnterpriseLoanCollectGracePeriod is a free data retrieval call binding the contract method 0xa0efd118.
//
// Solidity: function getEnterpriseLoanCollectGracePeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageSession) GetEnterpriseLoanCollectGracePeriod() (uint32, error) {
	return _EnterpriseStorage.Contract.GetEnterpriseLoanCollectGracePeriod(&_EnterpriseStorage.CallOpts)
}

// GetEnterpriseLoanCollectGracePeriod is a free data retrieval call binding the contract method 0xa0efd118.
//
// Solidity: function getEnterpriseLoanCollectGracePeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetEnterpriseLoanCollectGracePeriod() (uint32, error) {
	return _EnterpriseStorage.Contract.GetEnterpriseLoanCollectGracePeriod(&_EnterpriseStorage.CallOpts)
}

// GetEnterpriseVault is a free data retrieval call binding the contract method 0x7791c42d.
//
// Solidity: function getEnterpriseVault() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetEnterpriseVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getEnterpriseVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseVault is a free data retrieval call binding the contract method 0x7791c42d.
//
// Solidity: function getEnterpriseVault() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetEnterpriseVault() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetEnterpriseVault(&_EnterpriseStorage.CallOpts)
}

// GetEnterpriseVault is a free data retrieval call binding the contract method 0x7791c42d.
//
// Solidity: function getEnterpriseVault() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetEnterpriseVault() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetEnterpriseVault(&_EnterpriseStorage.CallOpts)
}

// GetGCFeePercent is a free data retrieval call binding the contract method 0xb406bf6d.
//
// Solidity: function getGCFeePercent() view returns(uint16)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetGCFeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getGCFeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetGCFeePercent is a free data retrieval call binding the contract method 0xb406bf6d.
//
// Solidity: function getGCFeePercent() view returns(uint16)
func (_EnterpriseStorage *EnterpriseStorageSession) GetGCFeePercent() (uint16, error) {
	return _EnterpriseStorage.Contract.GetGCFeePercent(&_EnterpriseStorage.CallOpts)
}

// GetGCFeePercent is a free data retrieval call binding the contract method 0xb406bf6d.
//
// Solidity: function getGCFeePercent() view returns(uint16)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetGCFeePercent() (uint16, error) {
	return _EnterpriseStorage.Contract.GetGCFeePercent(&_EnterpriseStorage.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint256 totalShares, uint32 interestGapHalvingPeriod, uint32 borrowerLoanReturnGracePeriod, uint32 enterpriseLoanCollectGracePeriod, uint16 gcFeePercent, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetInfo(opts *bind.CallOpts) (struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GcFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getInfo")

	outstruct := new(struct {
		Name                             string
		BaseUri                          string
		TotalShares                      *big.Int
		InterestGapHalvingPeriod         uint32
		BorrowerLoanReturnGracePeriod    uint32
		EnterpriseLoanCollectGracePeriod uint32
		GcFeePercent                     uint16
		FixedReserve                     *big.Int
		UsedReserve                      *big.Int
		StreamingReserve                 *big.Int
		StreamingReserveTarget           *big.Int
		StreamingReserveUpdated          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.BaseUri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TotalShares = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InterestGapHalvingPeriod = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.BorrowerLoanReturnGracePeriod = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.EnterpriseLoanCollectGracePeriod = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.GcFeePercent = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.FixedReserve = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsedReserve = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserve = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserveTarget = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserveUpdated = *abi.ConvertType(out[11], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint256 totalShares, uint32 interestGapHalvingPeriod, uint32 borrowerLoanReturnGracePeriod, uint32 enterpriseLoanCollectGracePeriod, uint16 gcFeePercent, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_EnterpriseStorage *EnterpriseStorageSession) GetInfo() (struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GcFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}, error) {
	return _EnterpriseStorage.Contract.GetInfo(&_EnterpriseStorage.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint256 totalShares, uint32 interestGapHalvingPeriod, uint32 borrowerLoanReturnGracePeriod, uint32 enterpriseLoanCollectGracePeriod, uint16 gcFeePercent, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetInfo() (struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GcFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}, error) {
	return _EnterpriseStorage.Contract.GetInfo(&_EnterpriseStorage.CallOpts)
}

// GetInterestGapHalvingPeriod is a free data retrieval call binding the contract method 0x5a4598df.
//
// Solidity: function getInterestGapHalvingPeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetInterestGapHalvingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getInterestGapHalvingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetInterestGapHalvingPeriod is a free data retrieval call binding the contract method 0x5a4598df.
//
// Solidity: function getInterestGapHalvingPeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageSession) GetInterestGapHalvingPeriod() (uint32, error) {
	return _EnterpriseStorage.Contract.GetInterestGapHalvingPeriod(&_EnterpriseStorage.CallOpts)
}

// GetInterestGapHalvingPeriod is a free data retrieval call binding the contract method 0x5a4598df.
//
// Solidity: function getInterestGapHalvingPeriod() view returns(uint32)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetInterestGapHalvingPeriod() (uint32, error) {
	return _EnterpriseStorage.Contract.GetInterestGapHalvingPeriod(&_EnterpriseStorage.CallOpts)
}

// GetInterestToken is a free data retrieval call binding the contract method 0x0ea1dc9e.
//
// Solidity: function getInterestToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetInterestToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getInterestToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestToken is a free data retrieval call binding the contract method 0x0ea1dc9e.
//
// Solidity: function getInterestToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetInterestToken() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetInterestToken(&_EnterpriseStorage.CallOpts)
}

// GetInterestToken is a free data retrieval call binding the contract method 0x0ea1dc9e.
//
// Solidity: function getInterestToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetInterestToken() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetInterestToken(&_EnterpriseStorage.CallOpts)
}

// GetLiquidityInfo is a free data retrieval call binding the contract method 0xdf57f2fa.
//
// Solidity: function getLiquidityInfo(uint256 interestTokenId) view returns((uint256,uint256,uint256))
func (_EnterpriseStorage *EnterpriseStorageCaller) GetLiquidityInfo(opts *bind.CallOpts, interestTokenId *big.Int) (EnterpriseStorageLiquidityInfo, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getLiquidityInfo", interestTokenId)

	if err != nil {
		return *new(EnterpriseStorageLiquidityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EnterpriseStorageLiquidityInfo)).(*EnterpriseStorageLiquidityInfo)

	return out0, err

}

// GetLiquidityInfo is a free data retrieval call binding the contract method 0xdf57f2fa.
//
// Solidity: function getLiquidityInfo(uint256 interestTokenId) view returns((uint256,uint256,uint256))
func (_EnterpriseStorage *EnterpriseStorageSession) GetLiquidityInfo(interestTokenId *big.Int) (EnterpriseStorageLiquidityInfo, error) {
	return _EnterpriseStorage.Contract.GetLiquidityInfo(&_EnterpriseStorage.CallOpts, interestTokenId)
}

// GetLiquidityInfo is a free data retrieval call binding the contract method 0xdf57f2fa.
//
// Solidity: function getLiquidityInfo(uint256 interestTokenId) view returns((uint256,uint256,uint256))
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetLiquidityInfo(interestTokenId *big.Int) (EnterpriseStorageLiquidityInfo, error) {
	return _EnterpriseStorage.Contract.GetLiquidityInfo(&_EnterpriseStorage.CallOpts, interestTokenId)
}

// GetLiquidityToken is a free data retrieval call binding the contract method 0x644e44df.
//
// Solidity: function getLiquidityToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetLiquidityToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getLiquidityToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidityToken is a free data retrieval call binding the contract method 0x644e44df.
//
// Solidity: function getLiquidityToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetLiquidityToken() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetLiquidityToken(&_EnterpriseStorage.CallOpts)
}

// GetLiquidityToken is a free data retrieval call binding the contract method 0x644e44df.
//
// Solidity: function getLiquidityToken() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetLiquidityToken() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetLiquidityToken(&_EnterpriseStorage.CallOpts)
}

// GetLoanInfo is a free data retrieval call binding the contract method 0x0714012c.
//
// Solidity: function getLoanInfo(uint256 borrowTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_EnterpriseStorage *EnterpriseStorageCaller) GetLoanInfo(opts *bind.CallOpts, borrowTokenId *big.Int) (EnterpriseStorageLoanInfo, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getLoanInfo", borrowTokenId)

	if err != nil {
		return *new(EnterpriseStorageLoanInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EnterpriseStorageLoanInfo)).(*EnterpriseStorageLoanInfo)

	return out0, err

}

// GetLoanInfo is a free data retrieval call binding the contract method 0x0714012c.
//
// Solidity: function getLoanInfo(uint256 borrowTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_EnterpriseStorage *EnterpriseStorageSession) GetLoanInfo(borrowTokenId *big.Int) (EnterpriseStorageLoanInfo, error) {
	return _EnterpriseStorage.Contract.GetLoanInfo(&_EnterpriseStorage.CallOpts, borrowTokenId)
}

// GetLoanInfo is a free data retrieval call binding the contract method 0x0714012c.
//
// Solidity: function getLoanInfo(uint256 borrowTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetLoanInfo(borrowTokenId *big.Int) (EnterpriseStorageLoanInfo, error) {
	return _EnterpriseStorage.Contract.GetLoanInfo(&_EnterpriseStorage.CallOpts, borrowTokenId)
}

// GetPowerTokens is a free data retrieval call binding the contract method 0xb96266fa.
//
// Solidity: function getPowerTokens() view returns(address[])
func (_EnterpriseStorage *EnterpriseStorageCaller) GetPowerTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getPowerTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPowerTokens is a free data retrieval call binding the contract method 0xb96266fa.
//
// Solidity: function getPowerTokens() view returns(address[])
func (_EnterpriseStorage *EnterpriseStorageSession) GetPowerTokens() ([]common.Address, error) {
	return _EnterpriseStorage.Contract.GetPowerTokens(&_EnterpriseStorage.CallOpts)
}

// GetPowerTokens is a free data retrieval call binding the contract method 0xb96266fa.
//
// Solidity: function getPowerTokens() view returns(address[])
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetPowerTokens() ([]common.Address, error) {
	return _EnterpriseStorage.Contract.GetPowerTokens(&_EnterpriseStorage.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getProxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) GetProxyAdmin() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetProxyAdmin(&_EnterpriseStorage.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetProxyAdmin() (common.Address, error) {
	return _EnterpriseStorage.Contract.GetProxyAdmin(&_EnterpriseStorage.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageSession) GetReserve() (*big.Int, error) {
	return _EnterpriseStorage.Contract.GetReserve(&_EnterpriseStorage.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetReserve() (*big.Int, error) {
	return _EnterpriseStorage.Contract.GetReserve(&_EnterpriseStorage.CallOpts)
}

// GetUsedReserve is a free data retrieval call binding the contract method 0xfcb2884e.
//
// Solidity: function getUsedReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageCaller) GetUsedReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "getUsedReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedReserve is a free data retrieval call binding the contract method 0xfcb2884e.
//
// Solidity: function getUsedReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageSession) GetUsedReserve() (*big.Int, error) {
	return _EnterpriseStorage.Contract.GetUsedReserve(&_EnterpriseStorage.CallOpts)
}

// GetUsedReserve is a free data retrieval call binding the contract method 0xfcb2884e.
//
// Solidity: function getUsedReserve() view returns(uint256)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) GetUsedReserve() (*big.Int, error) {
	return _EnterpriseStorage.Contract.GetUsedReserve(&_EnterpriseStorage.CallOpts)
}

// IsRegisteredPowerToken is a free data retrieval call binding the contract method 0x89035b61.
//
// Solidity: function isRegisteredPowerToken(address powerToken) view returns(bool)
func (_EnterpriseStorage *EnterpriseStorageCaller) IsRegisteredPowerToken(opts *bind.CallOpts, powerToken common.Address) (bool, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "isRegisteredPowerToken", powerToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredPowerToken is a free data retrieval call binding the contract method 0x89035b61.
//
// Solidity: function isRegisteredPowerToken(address powerToken) view returns(bool)
func (_EnterpriseStorage *EnterpriseStorageSession) IsRegisteredPowerToken(powerToken common.Address) (bool, error) {
	return _EnterpriseStorage.Contract.IsRegisteredPowerToken(&_EnterpriseStorage.CallOpts, powerToken)
}

// IsRegisteredPowerToken is a free data retrieval call binding the contract method 0x89035b61.
//
// Solidity: function isRegisteredPowerToken(address powerToken) view returns(bool)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) IsRegisteredPowerToken(powerToken common.Address) (bool, error) {
	return _EnterpriseStorage.Contract.IsRegisteredPowerToken(&_EnterpriseStorage.CallOpts, powerToken)
}

// IsSupportedPaymentToken is a free data retrieval call binding the contract method 0x2fb2067f.
//
// Solidity: function isSupportedPaymentToken(address token) view returns(bool)
func (_EnterpriseStorage *EnterpriseStorageCaller) IsSupportedPaymentToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "isSupportedPaymentToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedPaymentToken is a free data retrieval call binding the contract method 0x2fb2067f.
//
// Solidity: function isSupportedPaymentToken(address token) view returns(bool)
func (_EnterpriseStorage *EnterpriseStorageSession) IsSupportedPaymentToken(token common.Address) (bool, error) {
	return _EnterpriseStorage.Contract.IsSupportedPaymentToken(&_EnterpriseStorage.CallOpts, token)
}

// IsSupportedPaymentToken is a free data retrieval call binding the contract method 0x2fb2067f.
//
// Solidity: function isSupportedPaymentToken(address token) view returns(bool)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) IsSupportedPaymentToken(token common.Address) (bool, error) {
	return _EnterpriseStorage.Contract.IsSupportedPaymentToken(&_EnterpriseStorage.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) Owner() (common.Address, error) {
	return _EnterpriseStorage.Contract.Owner(&_EnterpriseStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) Owner() (common.Address, error) {
	return _EnterpriseStorage.Contract.Owner(&_EnterpriseStorage.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x0e101b60.
//
// Solidity: function paymentToken(uint256 index) view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCaller) PaymentToken(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "paymentToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x0e101b60.
//
// Solidity: function paymentToken(uint256 index) view returns(address)
func (_EnterpriseStorage *EnterpriseStorageSession) PaymentToken(index *big.Int) (common.Address, error) {
	return _EnterpriseStorage.Contract.PaymentToken(&_EnterpriseStorage.CallOpts, index)
}

// PaymentToken is a free data retrieval call binding the contract method 0x0e101b60.
//
// Solidity: function paymentToken(uint256 index) view returns(address)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) PaymentToken(index *big.Int) (common.Address, error) {
	return _EnterpriseStorage.Contract.PaymentToken(&_EnterpriseStorage.CallOpts, index)
}

// PaymentTokenIndex is a free data retrieval call binding the contract method 0xfbac9a55.
//
// Solidity: function paymentTokenIndex(address token) view returns(int16)
func (_EnterpriseStorage *EnterpriseStorageCaller) PaymentTokenIndex(opts *bind.CallOpts, token common.Address) (int16, error) {
	var out []interface{}
	err := _EnterpriseStorage.contract.Call(opts, &out, "paymentTokenIndex", token)

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// PaymentTokenIndex is a free data retrieval call binding the contract method 0xfbac9a55.
//
// Solidity: function paymentTokenIndex(address token) view returns(int16)
func (_EnterpriseStorage *EnterpriseStorageSession) PaymentTokenIndex(token common.Address) (int16, error) {
	return _EnterpriseStorage.Contract.PaymentTokenIndex(&_EnterpriseStorage.CallOpts, token)
}

// PaymentTokenIndex is a free data retrieval call binding the contract method 0xfbac9a55.
//
// Solidity: function paymentTokenIndex(address token) view returns(int16)
func (_EnterpriseStorage *EnterpriseStorageCallerSession) PaymentTokenIndex(token common.Address) (int16, error) {
	return _EnterpriseStorage.Contract.PaymentTokenIndex(&_EnterpriseStorage.CallOpts, token)
}

// DisablePaymentToken is a paid mutator transaction binding the contract method 0x9c7aa7f8.
//
// Solidity: function disablePaymentToken(address token) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) DisablePaymentToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "disablePaymentToken", token)
}

// DisablePaymentToken is a paid mutator transaction binding the contract method 0x9c7aa7f8.
//
// Solidity: function disablePaymentToken(address token) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) DisablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.DisablePaymentToken(&_EnterpriseStorage.TransactOpts, token)
}

// DisablePaymentToken is a paid mutator transaction binding the contract method 0x9c7aa7f8.
//
// Solidity: function disablePaymentToken(address token) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) DisablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.DisablePaymentToken(&_EnterpriseStorage.TransactOpts, token)
}

// EnablePaymentToken is a paid mutator transaction binding the contract method 0xbfd84fb4.
//
// Solidity: function enablePaymentToken(address token) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) EnablePaymentToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "enablePaymentToken", token)
}

// EnablePaymentToken is a paid mutator transaction binding the contract method 0xbfd84fb4.
//
// Solidity: function enablePaymentToken(address token) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) EnablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.EnablePaymentToken(&_EnterpriseStorage.TransactOpts, token)
}

// EnablePaymentToken is a paid mutator transaction binding the contract method 0xbfd84fb4.
//
// Solidity: function enablePaymentToken(address token) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) EnablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.EnablePaymentToken(&_EnterpriseStorage.TransactOpts, token)
}

// Initialize is a paid mutator transaction binding the contract method 0x6815f337.
//
// Solidity: function initialize(string enterpriseName, string baseUri, uint16 gcFeePercent, address converter, address proxyAdmin, address owner) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) Initialize(opts *bind.TransactOpts, enterpriseName string, baseUri string, gcFeePercent uint16, converter common.Address, proxyAdmin common.Address, owner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "initialize", enterpriseName, baseUri, gcFeePercent, converter, proxyAdmin, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x6815f337.
//
// Solidity: function initialize(string enterpriseName, string baseUri, uint16 gcFeePercent, address converter, address proxyAdmin, address owner) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) Initialize(enterpriseName string, baseUri string, gcFeePercent uint16, converter common.Address, proxyAdmin common.Address, owner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.Initialize(&_EnterpriseStorage.TransactOpts, enterpriseName, baseUri, gcFeePercent, converter, proxyAdmin, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x6815f337.
//
// Solidity: function initialize(string enterpriseName, string baseUri, uint16 gcFeePercent, address converter, address proxyAdmin, address owner) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) Initialize(enterpriseName string, baseUri string, gcFeePercent uint16, converter common.Address, proxyAdmin common.Address, owner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.Initialize(&_EnterpriseStorage.TransactOpts, enterpriseName, baseUri, gcFeePercent, converter, proxyAdmin, owner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) Initialize0(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "initialize0", initialOwner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) Initialize0(initialOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.Initialize0(&_EnterpriseStorage.TransactOpts, initialOwner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) Initialize0(initialOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.Initialize0(&_EnterpriseStorage.TransactOpts, initialOwner)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address liquidityToken, address interestToken, address borrowToken) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) InitializeTokens(opts *bind.TransactOpts, liquidityToken common.Address, interestToken common.Address, borrowToken common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "initializeTokens", liquidityToken, interestToken, borrowToken)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address liquidityToken, address interestToken, address borrowToken) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) InitializeTokens(liquidityToken common.Address, interestToken common.Address, borrowToken common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.InitializeTokens(&_EnterpriseStorage.TransactOpts, liquidityToken, interestToken, borrowToken)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address liquidityToken, address interestToken, address borrowToken) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) InitializeTokens(liquidityToken common.Address, interestToken common.Address, borrowToken common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.InitializeTokens(&_EnterpriseStorage.TransactOpts, liquidityToken, interestToken, borrowToken)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri string) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setBaseUri", baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetBaseUri(baseUri string) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetBaseUri(&_EnterpriseStorage.TransactOpts, baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetBaseUri(baseUri string) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetBaseUri(&_EnterpriseStorage.TransactOpts, baseUri)
}

// SetBondingCurve is a paid mutator transaction binding the contract method 0x3513e0dc.
//
// Solidity: function setBondingCurve(uint256 pole, uint256 slope) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetBondingCurve(opts *bind.TransactOpts, pole *big.Int, slope *big.Int) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setBondingCurve", pole, slope)
}

// SetBondingCurve is a paid mutator transaction binding the contract method 0x3513e0dc.
//
// Solidity: function setBondingCurve(uint256 pole, uint256 slope) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetBondingCurve(pole *big.Int, slope *big.Int) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetBondingCurve(&_EnterpriseStorage.TransactOpts, pole, slope)
}

// SetBondingCurve is a paid mutator transaction binding the contract method 0x3513e0dc.
//
// Solidity: function setBondingCurve(uint256 pole, uint256 slope) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetBondingCurve(pole *big.Int, slope *big.Int) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetBondingCurve(&_EnterpriseStorage.TransactOpts, pole, slope)
}

// SetBorrowerLoanReturnGracePeriod is a paid mutator transaction binding the contract method 0x2e365ade.
//
// Solidity: function setBorrowerLoanReturnGracePeriod(uint32 newPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetBorrowerLoanReturnGracePeriod(opts *bind.TransactOpts, newPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setBorrowerLoanReturnGracePeriod", newPeriod)
}

// SetBorrowerLoanReturnGracePeriod is a paid mutator transaction binding the contract method 0x2e365ade.
//
// Solidity: function setBorrowerLoanReturnGracePeriod(uint32 newPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetBorrowerLoanReturnGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetBorrowerLoanReturnGracePeriod(&_EnterpriseStorage.TransactOpts, newPeriod)
}

// SetBorrowerLoanReturnGracePeriod is a paid mutator transaction binding the contract method 0x2e365ade.
//
// Solidity: function setBorrowerLoanReturnGracePeriod(uint32 newPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetBorrowerLoanReturnGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetBorrowerLoanReturnGracePeriod(&_EnterpriseStorage.TransactOpts, newPeriod)
}

// SetConverter is a paid mutator transaction binding the contract method 0xb19337a4.
//
// Solidity: function setConverter(address newConverter) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetConverter(opts *bind.TransactOpts, newConverter common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setConverter", newConverter)
}

// SetConverter is a paid mutator transaction binding the contract method 0xb19337a4.
//
// Solidity: function setConverter(address newConverter) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetConverter(newConverter common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetConverter(&_EnterpriseStorage.TransactOpts, newConverter)
}

// SetConverter is a paid mutator transaction binding the contract method 0xb19337a4.
//
// Solidity: function setConverter(address newConverter) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetConverter(newConverter common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetConverter(&_EnterpriseStorage.TransactOpts, newConverter)
}

// SetEnterpriseCollector is a paid mutator transaction binding the contract method 0x62c1f388.
//
// Solidity: function setEnterpriseCollector(address newCollector) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetEnterpriseCollector(opts *bind.TransactOpts, newCollector common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setEnterpriseCollector", newCollector)
}

// SetEnterpriseCollector is a paid mutator transaction binding the contract method 0x62c1f388.
//
// Solidity: function setEnterpriseCollector(address newCollector) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetEnterpriseCollector(newCollector common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetEnterpriseCollector(&_EnterpriseStorage.TransactOpts, newCollector)
}

// SetEnterpriseCollector is a paid mutator transaction binding the contract method 0x62c1f388.
//
// Solidity: function setEnterpriseCollector(address newCollector) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetEnterpriseCollector(newCollector common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetEnterpriseCollector(&_EnterpriseStorage.TransactOpts, newCollector)
}

// SetEnterpriseLoanCollectGracePeriod is a paid mutator transaction binding the contract method 0xb56de97e.
//
// Solidity: function setEnterpriseLoanCollectGracePeriod(uint32 newPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetEnterpriseLoanCollectGracePeriod(opts *bind.TransactOpts, newPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setEnterpriseLoanCollectGracePeriod", newPeriod)
}

// SetEnterpriseLoanCollectGracePeriod is a paid mutator transaction binding the contract method 0xb56de97e.
//
// Solidity: function setEnterpriseLoanCollectGracePeriod(uint32 newPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetEnterpriseLoanCollectGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetEnterpriseLoanCollectGracePeriod(&_EnterpriseStorage.TransactOpts, newPeriod)
}

// SetEnterpriseLoanCollectGracePeriod is a paid mutator transaction binding the contract method 0xb56de97e.
//
// Solidity: function setEnterpriseLoanCollectGracePeriod(uint32 newPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetEnterpriseLoanCollectGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetEnterpriseLoanCollectGracePeriod(&_EnterpriseStorage.TransactOpts, newPeriod)
}

// SetEnterpriseVault is a paid mutator transaction binding the contract method 0x90f55edd.
//
// Solidity: function setEnterpriseVault(address newVault) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetEnterpriseVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setEnterpriseVault", newVault)
}

// SetEnterpriseVault is a paid mutator transaction binding the contract method 0x90f55edd.
//
// Solidity: function setEnterpriseVault(address newVault) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetEnterpriseVault(newVault common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetEnterpriseVault(&_EnterpriseStorage.TransactOpts, newVault)
}

// SetEnterpriseVault is a paid mutator transaction binding the contract method 0x90f55edd.
//
// Solidity: function setEnterpriseVault(address newVault) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetEnterpriseVault(newVault common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetEnterpriseVault(&_EnterpriseStorage.TransactOpts, newVault)
}

// SetGcFeePercent is a paid mutator transaction binding the contract method 0x449497ec.
//
// Solidity: function setGcFeePercent(uint16 newGcFeePercent) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetGcFeePercent(opts *bind.TransactOpts, newGcFeePercent uint16) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setGcFeePercent", newGcFeePercent)
}

// SetGcFeePercent is a paid mutator transaction binding the contract method 0x449497ec.
//
// Solidity: function setGcFeePercent(uint16 newGcFeePercent) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetGcFeePercent(newGcFeePercent uint16) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetGcFeePercent(&_EnterpriseStorage.TransactOpts, newGcFeePercent)
}

// SetGcFeePercent is a paid mutator transaction binding the contract method 0x449497ec.
//
// Solidity: function setGcFeePercent(uint16 newGcFeePercent) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetGcFeePercent(newGcFeePercent uint16) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetGcFeePercent(&_EnterpriseStorage.TransactOpts, newGcFeePercent)
}

// SetInterestGapHalvingPeriod is a paid mutator transaction binding the contract method 0xe1a5fa49.
//
// Solidity: function setInterestGapHalvingPeriod(uint32 interestGapHalvingPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) SetInterestGapHalvingPeriod(opts *bind.TransactOpts, interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "setInterestGapHalvingPeriod", interestGapHalvingPeriod)
}

// SetInterestGapHalvingPeriod is a paid mutator transaction binding the contract method 0xe1a5fa49.
//
// Solidity: function setInterestGapHalvingPeriod(uint32 interestGapHalvingPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) SetInterestGapHalvingPeriod(interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetInterestGapHalvingPeriod(&_EnterpriseStorage.TransactOpts, interestGapHalvingPeriod)
}

// SetInterestGapHalvingPeriod is a paid mutator transaction binding the contract method 0xe1a5fa49.
//
// Solidity: function setInterestGapHalvingPeriod(uint32 interestGapHalvingPeriod) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) SetInterestGapHalvingPeriod(interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.SetInterestGapHalvingPeriod(&_EnterpriseStorage.TransactOpts, interestGapHalvingPeriod)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.TransferOwnership(&_EnterpriseStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.TransferOwnership(&_EnterpriseStorage.TransactOpts, newOwner)
}

// UpgradeBorrowToken is a paid mutator transaction binding the contract method 0x7b024363.
//
// Solidity: function upgradeBorrowToken(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) UpgradeBorrowToken(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "upgradeBorrowToken", implementation)
}

// UpgradeBorrowToken is a paid mutator transaction binding the contract method 0x7b024363.
//
// Solidity: function upgradeBorrowToken(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) UpgradeBorrowToken(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradeBorrowToken(&_EnterpriseStorage.TransactOpts, implementation)
}

// UpgradeBorrowToken is a paid mutator transaction binding the contract method 0x7b024363.
//
// Solidity: function upgradeBorrowToken(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) UpgradeBorrowToken(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradeBorrowToken(&_EnterpriseStorage.TransactOpts, implementation)
}

// UpgradeEnterprise is a paid mutator transaction binding the contract method 0xc1fc16d8.
//
// Solidity: function upgradeEnterprise(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) UpgradeEnterprise(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "upgradeEnterprise", implementation)
}

// UpgradeEnterprise is a paid mutator transaction binding the contract method 0xc1fc16d8.
//
// Solidity: function upgradeEnterprise(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) UpgradeEnterprise(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradeEnterprise(&_EnterpriseStorage.TransactOpts, implementation)
}

// UpgradeEnterprise is a paid mutator transaction binding the contract method 0xc1fc16d8.
//
// Solidity: function upgradeEnterprise(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) UpgradeEnterprise(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradeEnterprise(&_EnterpriseStorage.TransactOpts, implementation)
}

// UpgradeInterestToken is a paid mutator transaction binding the contract method 0x2b2f6095.
//
// Solidity: function upgradeInterestToken(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) UpgradeInterestToken(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "upgradeInterestToken", implementation)
}

// UpgradeInterestToken is a paid mutator transaction binding the contract method 0x2b2f6095.
//
// Solidity: function upgradeInterestToken(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) UpgradeInterestToken(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradeInterestToken(&_EnterpriseStorage.TransactOpts, implementation)
}

// UpgradeInterestToken is a paid mutator transaction binding the contract method 0x2b2f6095.
//
// Solidity: function upgradeInterestToken(address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) UpgradeInterestToken(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradeInterestToken(&_EnterpriseStorage.TransactOpts, implementation)
}

// UpgradePowerToken is a paid mutator transaction binding the contract method 0xc2294784.
//
// Solidity: function upgradePowerToken(address powerToken, address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactor) UpgradePowerToken(opts *bind.TransactOpts, powerToken common.Address, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.contract.Transact(opts, "upgradePowerToken", powerToken, implementation)
}

// UpgradePowerToken is a paid mutator transaction binding the contract method 0xc2294784.
//
// Solidity: function upgradePowerToken(address powerToken, address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageSession) UpgradePowerToken(powerToken common.Address, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradePowerToken(&_EnterpriseStorage.TransactOpts, powerToken, implementation)
}

// UpgradePowerToken is a paid mutator transaction binding the contract method 0xc2294784.
//
// Solidity: function upgradePowerToken(address powerToken, address implementation) returns()
func (_EnterpriseStorage *EnterpriseStorageTransactorSession) UpgradePowerToken(powerToken common.Address, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseStorage.Contract.UpgradePowerToken(&_EnterpriseStorage.TransactOpts, powerToken, implementation)
}

// EnterpriseStorageBaseUriChangedIterator is returned from FilterBaseUriChanged and is used to iterate over the raw logs and unpacked data for BaseUriChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageBaseUriChangedIterator struct {
	Event *EnterpriseStorageBaseUriChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageBaseUriChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageBaseUriChanged)
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
		it.Event = new(EnterpriseStorageBaseUriChanged)
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
func (it *EnterpriseStorageBaseUriChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageBaseUriChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageBaseUriChanged represents a BaseUriChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageBaseUriChanged struct {
	BaseUri string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseUriChanged is a free log retrieval operation binding the contract event 0x87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6.
//
// Solidity: event BaseUriChanged(string baseUri)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterBaseUriChanged(opts *bind.FilterOpts) (*EnterpriseStorageBaseUriChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "BaseUriChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageBaseUriChangedIterator{contract: _EnterpriseStorage.contract, event: "BaseUriChanged", logs: logs, sub: sub}, nil
}

// WatchBaseUriChanged is a free log subscription operation binding the contract event 0x87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6.
//
// Solidity: event BaseUriChanged(string baseUri)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchBaseUriChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageBaseUriChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "BaseUriChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageBaseUriChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "BaseUriChanged", log); err != nil {
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

// ParseBaseUriChanged is a log parse operation binding the contract event 0x87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6.
//
// Solidity: event BaseUriChanged(string baseUri)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseBaseUriChanged(log types.Log) (*EnterpriseStorageBaseUriChanged, error) {
	event := new(EnterpriseStorageBaseUriChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "BaseUriChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageBondingChangedIterator is returned from FilterBondingChanged and is used to iterate over the raw logs and unpacked data for BondingChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageBondingChangedIterator struct {
	Event *EnterpriseStorageBondingChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageBondingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageBondingChanged)
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
		it.Event = new(EnterpriseStorageBondingChanged)
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
func (it *EnterpriseStorageBondingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageBondingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageBondingChanged represents a BondingChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageBondingChanged struct {
	Pole  *big.Int
	Slope *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBondingChanged is a free log retrieval operation binding the contract event 0x926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39f.
//
// Solidity: event BondingChanged(uint256 pole, uint256 slope)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterBondingChanged(opts *bind.FilterOpts) (*EnterpriseStorageBondingChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "BondingChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageBondingChangedIterator{contract: _EnterpriseStorage.contract, event: "BondingChanged", logs: logs, sub: sub}, nil
}

// WatchBondingChanged is a free log subscription operation binding the contract event 0x926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39f.
//
// Solidity: event BondingChanged(uint256 pole, uint256 slope)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchBondingChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageBondingChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "BondingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageBondingChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "BondingChanged", log); err != nil {
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

// ParseBondingChanged is a log parse operation binding the contract event 0x926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39f.
//
// Solidity: event BondingChanged(uint256 pole, uint256 slope)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseBondingChanged(log types.Log) (*EnterpriseStorageBondingChanged, error) {
	event := new(EnterpriseStorageBondingChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "BondingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator is returned from FilterBorrowerLoanReturnGracePeriodChanged and is used to iterate over the raw logs and unpacked data for BorrowerLoanReturnGracePeriodChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator struct {
	Event *EnterpriseStorageBorrowerLoanReturnGracePeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageBorrowerLoanReturnGracePeriodChanged)
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
		it.Event = new(EnterpriseStorageBorrowerLoanReturnGracePeriodChanged)
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
func (it *EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageBorrowerLoanReturnGracePeriodChanged represents a BorrowerLoanReturnGracePeriodChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageBorrowerLoanReturnGracePeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrowerLoanReturnGracePeriodChanged is a free log retrieval operation binding the contract event 0xf6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a.
//
// Solidity: event BorrowerLoanReturnGracePeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterBorrowerLoanReturnGracePeriodChanged(opts *bind.FilterOpts) (*EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "BorrowerLoanReturnGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageBorrowerLoanReturnGracePeriodChangedIterator{contract: _EnterpriseStorage.contract, event: "BorrowerLoanReturnGracePeriodChanged", logs: logs, sub: sub}, nil
}

// WatchBorrowerLoanReturnGracePeriodChanged is a free log subscription operation binding the contract event 0xf6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a.
//
// Solidity: event BorrowerLoanReturnGracePeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchBorrowerLoanReturnGracePeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageBorrowerLoanReturnGracePeriodChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "BorrowerLoanReturnGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageBorrowerLoanReturnGracePeriodChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "BorrowerLoanReturnGracePeriodChanged", log); err != nil {
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

// ParseBorrowerLoanReturnGracePeriodChanged is a log parse operation binding the contract event 0xf6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a.
//
// Solidity: event BorrowerLoanReturnGracePeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseBorrowerLoanReturnGracePeriodChanged(log types.Log) (*EnterpriseStorageBorrowerLoanReturnGracePeriodChanged, error) {
	event := new(EnterpriseStorageBorrowerLoanReturnGracePeriodChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "BorrowerLoanReturnGracePeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageConverterChangedIterator is returned from FilterConverterChanged and is used to iterate over the raw logs and unpacked data for ConverterChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageConverterChangedIterator struct {
	Event *EnterpriseStorageConverterChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageConverterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageConverterChanged)
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
		it.Event = new(EnterpriseStorageConverterChanged)
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
func (it *EnterpriseStorageConverterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageConverterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageConverterChanged represents a ConverterChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageConverterChanged struct {
	Converter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConverterChanged is a free log retrieval operation binding the contract event 0xbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656df.
//
// Solidity: event ConverterChanged(address converter)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterConverterChanged(opts *bind.FilterOpts) (*EnterpriseStorageConverterChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "ConverterChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageConverterChangedIterator{contract: _EnterpriseStorage.contract, event: "ConverterChanged", logs: logs, sub: sub}, nil
}

// WatchConverterChanged is a free log subscription operation binding the contract event 0xbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656df.
//
// Solidity: event ConverterChanged(address converter)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchConverterChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageConverterChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "ConverterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageConverterChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "ConverterChanged", log); err != nil {
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

// ParseConverterChanged is a log parse operation binding the contract event 0xbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656df.
//
// Solidity: event ConverterChanged(address converter)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseConverterChanged(log types.Log) (*EnterpriseStorageConverterChanged, error) {
	event := new(EnterpriseStorageConverterChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "ConverterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageEnterpriseCollectorChangedIterator is returned from FilterEnterpriseCollectorChanged and is used to iterate over the raw logs and unpacked data for EnterpriseCollectorChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseCollectorChangedIterator struct {
	Event *EnterpriseStorageEnterpriseCollectorChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageEnterpriseCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageEnterpriseCollectorChanged)
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
		it.Event = new(EnterpriseStorageEnterpriseCollectorChanged)
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
func (it *EnterpriseStorageEnterpriseCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageEnterpriseCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageEnterpriseCollectorChanged represents a EnterpriseCollectorChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseCollectorChanged struct {
	Collector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseCollectorChanged is a free log retrieval operation binding the contract event 0x7aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4.
//
// Solidity: event EnterpriseCollectorChanged(address collector)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterEnterpriseCollectorChanged(opts *bind.FilterOpts) (*EnterpriseStorageEnterpriseCollectorChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "EnterpriseCollectorChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageEnterpriseCollectorChangedIterator{contract: _EnterpriseStorage.contract, event: "EnterpriseCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseCollectorChanged is a free log subscription operation binding the contract event 0x7aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4.
//
// Solidity: event EnterpriseCollectorChanged(address collector)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchEnterpriseCollectorChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageEnterpriseCollectorChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "EnterpriseCollectorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageEnterpriseCollectorChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseCollectorChanged", log); err != nil {
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

// ParseEnterpriseCollectorChanged is a log parse operation binding the contract event 0x7aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4.
//
// Solidity: event EnterpriseCollectorChanged(address collector)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseEnterpriseCollectorChanged(log types.Log) (*EnterpriseStorageEnterpriseCollectorChanged, error) {
	event := new(EnterpriseStorageEnterpriseCollectorChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator is returned from FilterEnterpriseLoanCollectGracePeriodChanged and is used to iterate over the raw logs and unpacked data for EnterpriseLoanCollectGracePeriodChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator struct {
	Event *EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged)
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
		it.Event = new(EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged)
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
func (it *EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged represents a EnterpriseLoanCollectGracePeriodChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseLoanCollectGracePeriodChanged is a free log retrieval operation binding the contract event 0x7c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b.
//
// Solidity: event EnterpriseLoanCollectGracePeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterEnterpriseLoanCollectGracePeriodChanged(opts *bind.FilterOpts) (*EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "EnterpriseLoanCollectGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageEnterpriseLoanCollectGracePeriodChangedIterator{contract: _EnterpriseStorage.contract, event: "EnterpriseLoanCollectGracePeriodChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseLoanCollectGracePeriodChanged is a free log subscription operation binding the contract event 0x7c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b.
//
// Solidity: event EnterpriseLoanCollectGracePeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchEnterpriseLoanCollectGracePeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "EnterpriseLoanCollectGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseLoanCollectGracePeriodChanged", log); err != nil {
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

// ParseEnterpriseLoanCollectGracePeriodChanged is a log parse operation binding the contract event 0x7c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b.
//
// Solidity: event EnterpriseLoanCollectGracePeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseEnterpriseLoanCollectGracePeriodChanged(log types.Log) (*EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged, error) {
	event := new(EnterpriseStorageEnterpriseLoanCollectGracePeriodChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseLoanCollectGracePeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageEnterpriseShutdownIterator is returned from FilterEnterpriseShutdown and is used to iterate over the raw logs and unpacked data for EnterpriseShutdown events raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseShutdownIterator struct {
	Event *EnterpriseStorageEnterpriseShutdown // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageEnterpriseShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageEnterpriseShutdown)
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
		it.Event = new(EnterpriseStorageEnterpriseShutdown)
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
func (it *EnterpriseStorageEnterpriseShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageEnterpriseShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageEnterpriseShutdown represents a EnterpriseShutdown event raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseShutdown is a free log retrieval operation binding the contract event 0x6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c072.
//
// Solidity: event EnterpriseShutdown()
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterEnterpriseShutdown(opts *bind.FilterOpts) (*EnterpriseStorageEnterpriseShutdownIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "EnterpriseShutdown")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageEnterpriseShutdownIterator{contract: _EnterpriseStorage.contract, event: "EnterpriseShutdown", logs: logs, sub: sub}, nil
}

// WatchEnterpriseShutdown is a free log subscription operation binding the contract event 0x6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c072.
//
// Solidity: event EnterpriseShutdown()
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchEnterpriseShutdown(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageEnterpriseShutdown) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "EnterpriseShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageEnterpriseShutdown)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseShutdown", log); err != nil {
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

// ParseEnterpriseShutdown is a log parse operation binding the contract event 0x6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c072.
//
// Solidity: event EnterpriseShutdown()
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseEnterpriseShutdown(log types.Log) (*EnterpriseStorageEnterpriseShutdown, error) {
	event := new(EnterpriseStorageEnterpriseShutdown)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageEnterpriseVaultChangedIterator is returned from FilterEnterpriseVaultChanged and is used to iterate over the raw logs and unpacked data for EnterpriseVaultChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseVaultChangedIterator struct {
	Event *EnterpriseStorageEnterpriseVaultChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageEnterpriseVaultChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageEnterpriseVaultChanged)
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
		it.Event = new(EnterpriseStorageEnterpriseVaultChanged)
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
func (it *EnterpriseStorageEnterpriseVaultChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageEnterpriseVaultChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageEnterpriseVaultChanged represents a EnterpriseVaultChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageEnterpriseVaultChanged struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseVaultChanged is a free log retrieval operation binding the contract event 0x0dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762dd.
//
// Solidity: event EnterpriseVaultChanged(address vault)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterEnterpriseVaultChanged(opts *bind.FilterOpts) (*EnterpriseStorageEnterpriseVaultChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "EnterpriseVaultChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageEnterpriseVaultChangedIterator{contract: _EnterpriseStorage.contract, event: "EnterpriseVaultChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseVaultChanged is a free log subscription operation binding the contract event 0x0dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762dd.
//
// Solidity: event EnterpriseVaultChanged(address vault)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchEnterpriseVaultChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageEnterpriseVaultChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "EnterpriseVaultChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageEnterpriseVaultChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseVaultChanged", log); err != nil {
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

// ParseEnterpriseVaultChanged is a log parse operation binding the contract event 0x0dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762dd.
//
// Solidity: event EnterpriseVaultChanged(address vault)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseEnterpriseVaultChanged(log types.Log) (*EnterpriseStorageEnterpriseVaultChanged, error) {
	event := new(EnterpriseStorageEnterpriseVaultChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "EnterpriseVaultChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageFixedReserveChangedIterator is returned from FilterFixedReserveChanged and is used to iterate over the raw logs and unpacked data for FixedReserveChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageFixedReserveChangedIterator struct {
	Event *EnterpriseStorageFixedReserveChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageFixedReserveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageFixedReserveChanged)
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
		it.Event = new(EnterpriseStorageFixedReserveChanged)
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
func (it *EnterpriseStorageFixedReserveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageFixedReserveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageFixedReserveChanged represents a FixedReserveChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageFixedReserveChanged struct {
	FixedReserve *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFixedReserveChanged is a free log retrieval operation binding the contract event 0xa70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180.
//
// Solidity: event FixedReserveChanged(uint256 fixedReserve)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterFixedReserveChanged(opts *bind.FilterOpts) (*EnterpriseStorageFixedReserveChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "FixedReserveChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageFixedReserveChangedIterator{contract: _EnterpriseStorage.contract, event: "FixedReserveChanged", logs: logs, sub: sub}, nil
}

// WatchFixedReserveChanged is a free log subscription operation binding the contract event 0xa70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180.
//
// Solidity: event FixedReserveChanged(uint256 fixedReserve)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchFixedReserveChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageFixedReserveChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "FixedReserveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageFixedReserveChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "FixedReserveChanged", log); err != nil {
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

// ParseFixedReserveChanged is a log parse operation binding the contract event 0xa70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180.
//
// Solidity: event FixedReserveChanged(uint256 fixedReserve)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseFixedReserveChanged(log types.Log) (*EnterpriseStorageFixedReserveChanged, error) {
	event := new(EnterpriseStorageFixedReserveChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "FixedReserveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageGcFeePercentChangedIterator is returned from FilterGcFeePercentChanged and is used to iterate over the raw logs and unpacked data for GcFeePercentChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageGcFeePercentChangedIterator struct {
	Event *EnterpriseStorageGcFeePercentChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageGcFeePercentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageGcFeePercentChanged)
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
		it.Event = new(EnterpriseStorageGcFeePercentChanged)
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
func (it *EnterpriseStorageGcFeePercentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageGcFeePercentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageGcFeePercentChanged represents a GcFeePercentChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageGcFeePercentChanged struct {
	Percent uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGcFeePercentChanged is a free log retrieval operation binding the contract event 0x27894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a15646.
//
// Solidity: event GcFeePercentChanged(uint16 percent)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterGcFeePercentChanged(opts *bind.FilterOpts) (*EnterpriseStorageGcFeePercentChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "GcFeePercentChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageGcFeePercentChangedIterator{contract: _EnterpriseStorage.contract, event: "GcFeePercentChanged", logs: logs, sub: sub}, nil
}

// WatchGcFeePercentChanged is a free log subscription operation binding the contract event 0x27894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a15646.
//
// Solidity: event GcFeePercentChanged(uint16 percent)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchGcFeePercentChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageGcFeePercentChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "GcFeePercentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageGcFeePercentChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "GcFeePercentChanged", log); err != nil {
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

// ParseGcFeePercentChanged is a log parse operation binding the contract event 0x27894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a15646.
//
// Solidity: event GcFeePercentChanged(uint16 percent)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseGcFeePercentChanged(log types.Log) (*EnterpriseStorageGcFeePercentChanged, error) {
	event := new(EnterpriseStorageGcFeePercentChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "GcFeePercentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageInterestGapHalvingPeriodChangedIterator is returned from FilterInterestGapHalvingPeriodChanged and is used to iterate over the raw logs and unpacked data for InterestGapHalvingPeriodChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageInterestGapHalvingPeriodChangedIterator struct {
	Event *EnterpriseStorageInterestGapHalvingPeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageInterestGapHalvingPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageInterestGapHalvingPeriodChanged)
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
		it.Event = new(EnterpriseStorageInterestGapHalvingPeriodChanged)
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
func (it *EnterpriseStorageInterestGapHalvingPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageInterestGapHalvingPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageInterestGapHalvingPeriodChanged represents a InterestGapHalvingPeriodChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageInterestGapHalvingPeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterestGapHalvingPeriodChanged is a free log retrieval operation binding the contract event 0x971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f.
//
// Solidity: event InterestGapHalvingPeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterInterestGapHalvingPeriodChanged(opts *bind.FilterOpts) (*EnterpriseStorageInterestGapHalvingPeriodChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "InterestGapHalvingPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageInterestGapHalvingPeriodChangedIterator{contract: _EnterpriseStorage.contract, event: "InterestGapHalvingPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchInterestGapHalvingPeriodChanged is a free log subscription operation binding the contract event 0x971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f.
//
// Solidity: event InterestGapHalvingPeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchInterestGapHalvingPeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageInterestGapHalvingPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "InterestGapHalvingPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageInterestGapHalvingPeriodChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "InterestGapHalvingPeriodChanged", log); err != nil {
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

// ParseInterestGapHalvingPeriodChanged is a log parse operation binding the contract event 0x971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f.
//
// Solidity: event InterestGapHalvingPeriodChanged(uint32 period)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseInterestGapHalvingPeriodChanged(log types.Log) (*EnterpriseStorageInterestGapHalvingPeriodChanged, error) {
	event := new(EnterpriseStorageInterestGapHalvingPeriodChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "InterestGapHalvingPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EnterpriseStorage contract.
type EnterpriseStorageOwnershipTransferredIterator struct {
	Event *EnterpriseStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageOwnershipTransferred)
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
		it.Event = new(EnterpriseStorageOwnershipTransferred)
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
func (it *EnterpriseStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageOwnershipTransferred represents a OwnershipTransferred event raised by the EnterpriseStorage contract.
type EnterpriseStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnterpriseStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageOwnershipTransferredIterator{contract: _EnterpriseStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageOwnershipTransferred)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseOwnershipTransferred(log types.Log) (*EnterpriseStorageOwnershipTransferred, error) {
	event := new(EnterpriseStorageOwnershipTransferred)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStoragePaymentTokenChangeIterator is returned from FilterPaymentTokenChange and is used to iterate over the raw logs and unpacked data for PaymentTokenChange events raised by the EnterpriseStorage contract.
type EnterpriseStoragePaymentTokenChangeIterator struct {
	Event *EnterpriseStoragePaymentTokenChange // Event containing the contract specifics and raw log

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
func (it *EnterpriseStoragePaymentTokenChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStoragePaymentTokenChange)
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
		it.Event = new(EnterpriseStoragePaymentTokenChange)
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
func (it *EnterpriseStoragePaymentTokenChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStoragePaymentTokenChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStoragePaymentTokenChange represents a PaymentTokenChange event raised by the EnterpriseStorage contract.
type EnterpriseStoragePaymentTokenChange struct {
	PaymentToken common.Address
	Enabled      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenChange is a free log retrieval operation binding the contract event 0x92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907.
//
// Solidity: event PaymentTokenChange(address paymentToken, bool enabled)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterPaymentTokenChange(opts *bind.FilterOpts) (*EnterpriseStoragePaymentTokenChangeIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "PaymentTokenChange")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStoragePaymentTokenChangeIterator{contract: _EnterpriseStorage.contract, event: "PaymentTokenChange", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenChange is a free log subscription operation binding the contract event 0x92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907.
//
// Solidity: event PaymentTokenChange(address paymentToken, bool enabled)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchPaymentTokenChange(opts *bind.WatchOpts, sink chan<- *EnterpriseStoragePaymentTokenChange) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "PaymentTokenChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStoragePaymentTokenChange)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "PaymentTokenChange", log); err != nil {
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

// ParsePaymentTokenChange is a log parse operation binding the contract event 0x92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907.
//
// Solidity: event PaymentTokenChange(address paymentToken, bool enabled)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParsePaymentTokenChange(log types.Log) (*EnterpriseStoragePaymentTokenChange, error) {
	event := new(EnterpriseStoragePaymentTokenChange)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "PaymentTokenChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageStreamingReserveChangedIterator is returned from FilterStreamingReserveChanged and is used to iterate over the raw logs and unpacked data for StreamingReserveChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageStreamingReserveChangedIterator struct {
	Event *EnterpriseStorageStreamingReserveChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageStreamingReserveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageStreamingReserveChanged)
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
		it.Event = new(EnterpriseStorageStreamingReserveChanged)
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
func (it *EnterpriseStorageStreamingReserveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageStreamingReserveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageStreamingReserveChanged represents a StreamingReserveChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageStreamingReserveChanged struct {
	StreamingReserve       *big.Int
	StreamingReserveTarget *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterStreamingReserveChanged is a free log retrieval operation binding the contract event 0xc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a11829.
//
// Solidity: event StreamingReserveChanged(uint112 streamingReserve, uint112 streamingReserveTarget)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterStreamingReserveChanged(opts *bind.FilterOpts) (*EnterpriseStorageStreamingReserveChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "StreamingReserveChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageStreamingReserveChangedIterator{contract: _EnterpriseStorage.contract, event: "StreamingReserveChanged", logs: logs, sub: sub}, nil
}

// WatchStreamingReserveChanged is a free log subscription operation binding the contract event 0xc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a11829.
//
// Solidity: event StreamingReserveChanged(uint112 streamingReserve, uint112 streamingReserveTarget)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchStreamingReserveChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageStreamingReserveChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "StreamingReserveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageStreamingReserveChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "StreamingReserveChanged", log); err != nil {
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

// ParseStreamingReserveChanged is a log parse operation binding the contract event 0xc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a11829.
//
// Solidity: event StreamingReserveChanged(uint112 streamingReserve, uint112 streamingReserveTarget)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseStreamingReserveChanged(log types.Log) (*EnterpriseStorageStreamingReserveChanged, error) {
	event := new(EnterpriseStorageStreamingReserveChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "StreamingReserveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageTotalSharesChangedIterator is returned from FilterTotalSharesChanged and is used to iterate over the raw logs and unpacked data for TotalSharesChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageTotalSharesChangedIterator struct {
	Event *EnterpriseStorageTotalSharesChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageTotalSharesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageTotalSharesChanged)
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
		it.Event = new(EnterpriseStorageTotalSharesChanged)
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
func (it *EnterpriseStorageTotalSharesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageTotalSharesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageTotalSharesChanged represents a TotalSharesChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageTotalSharesChanged struct {
	TotalShares *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTotalSharesChanged is a free log retrieval operation binding the contract event 0xf75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264.
//
// Solidity: event TotalSharesChanged(uint256 totalShares)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterTotalSharesChanged(opts *bind.FilterOpts) (*EnterpriseStorageTotalSharesChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "TotalSharesChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageTotalSharesChangedIterator{contract: _EnterpriseStorage.contract, event: "TotalSharesChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSharesChanged is a free log subscription operation binding the contract event 0xf75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264.
//
// Solidity: event TotalSharesChanged(uint256 totalShares)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchTotalSharesChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageTotalSharesChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "TotalSharesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageTotalSharesChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "TotalSharesChanged", log); err != nil {
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

// ParseTotalSharesChanged is a log parse operation binding the contract event 0xf75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264.
//
// Solidity: event TotalSharesChanged(uint256 totalShares)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseTotalSharesChanged(log types.Log) (*EnterpriseStorageTotalSharesChanged, error) {
	event := new(EnterpriseStorageTotalSharesChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "TotalSharesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStorageUsedReserveChangedIterator is returned from FilterUsedReserveChanged and is used to iterate over the raw logs and unpacked data for UsedReserveChanged events raised by the EnterpriseStorage contract.
type EnterpriseStorageUsedReserveChangedIterator struct {
	Event *EnterpriseStorageUsedReserveChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStorageUsedReserveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStorageUsedReserveChanged)
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
		it.Event = new(EnterpriseStorageUsedReserveChanged)
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
func (it *EnterpriseStorageUsedReserveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStorageUsedReserveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStorageUsedReserveChanged represents a UsedReserveChanged event raised by the EnterpriseStorage contract.
type EnterpriseStorageUsedReserveChanged struct {
	FixedReserve *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUsedReserveChanged is a free log retrieval operation binding the contract event 0x80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585.
//
// Solidity: event UsedReserveChanged(uint256 fixedReserve)
func (_EnterpriseStorage *EnterpriseStorageFilterer) FilterUsedReserveChanged(opts *bind.FilterOpts) (*EnterpriseStorageUsedReserveChangedIterator, error) {

	logs, sub, err := _EnterpriseStorage.contract.FilterLogs(opts, "UsedReserveChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStorageUsedReserveChangedIterator{contract: _EnterpriseStorage.contract, event: "UsedReserveChanged", logs: logs, sub: sub}, nil
}

// WatchUsedReserveChanged is a free log subscription operation binding the contract event 0x80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585.
//
// Solidity: event UsedReserveChanged(uint256 fixedReserve)
func (_EnterpriseStorage *EnterpriseStorageFilterer) WatchUsedReserveChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStorageUsedReserveChanged) (event.Subscription, error) {

	logs, sub, err := _EnterpriseStorage.contract.WatchLogs(opts, "UsedReserveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStorageUsedReserveChanged)
				if err := _EnterpriseStorage.contract.UnpackLog(event, "UsedReserveChanged", log); err != nil {
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

// ParseUsedReserveChanged is a log parse operation binding the contract event 0x80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585.
//
// Solidity: event UsedReserveChanged(uint256 fixedReserve)
func (_EnterpriseStorage *EnterpriseStorageFilterer) ParseUsedReserveChanged(log types.Log) (*EnterpriseStorageUsedReserveChanged, error) {
	event := new(EnterpriseStorageUsedReserveChanged)
	if err := _EnterpriseStorage.contract.UnpackLog(event, "UsedReserveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
