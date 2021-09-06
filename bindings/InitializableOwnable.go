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

// InitializableOwnableMetaData contains all meta data concerning the InitializableOwnable contract.
var InitializableOwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InitializableOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableOwnableMetaData.ABI instead.
var InitializableOwnableABI = InitializableOwnableMetaData.ABI

// InitializableOwnable is an auto generated Go binding around an Ethereum contract.
type InitializableOwnable struct {
	InitializableOwnableCaller     // Read-only binding to the contract
	InitializableOwnableTransactor // Write-only binding to the contract
	InitializableOwnableFilterer   // Log filterer for contract events
}

// InitializableOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableOwnableSession struct {
	Contract     *InitializableOwnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InitializableOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableOwnableCallerSession struct {
	Contract *InitializableOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InitializableOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableOwnableTransactorSession struct {
	Contract     *InitializableOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InitializableOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableOwnableRaw struct {
	Contract *InitializableOwnable // Generic contract binding to access the raw methods on
}

// InitializableOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableOwnableCallerRaw struct {
	Contract *InitializableOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableOwnableTransactorRaw struct {
	Contract *InitializableOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializableOwnable creates a new instance of InitializableOwnable, bound to a specific deployed contract.
func NewInitializableOwnable(address common.Address, backend bind.ContractBackend) (*InitializableOwnable, error) {
	contract, err := bindInitializableOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InitializableOwnable{InitializableOwnableCaller: InitializableOwnableCaller{contract: contract}, InitializableOwnableTransactor: InitializableOwnableTransactor{contract: contract}, InitializableOwnableFilterer: InitializableOwnableFilterer{contract: contract}}, nil
}

// NewInitializableOwnableCaller creates a new read-only instance of InitializableOwnable, bound to a specific deployed contract.
func NewInitializableOwnableCaller(address common.Address, caller bind.ContractCaller) (*InitializableOwnableCaller, error) {
	contract, err := bindInitializableOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableOwnableCaller{contract: contract}, nil
}

// NewInitializableOwnableTransactor creates a new write-only instance of InitializableOwnable, bound to a specific deployed contract.
func NewInitializableOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableOwnableTransactor, error) {
	contract, err := bindInitializableOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableOwnableTransactor{contract: contract}, nil
}

// NewInitializableOwnableFilterer creates a new log filterer instance of InitializableOwnable, bound to a specific deployed contract.
func NewInitializableOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableOwnableFilterer, error) {
	contract, err := bindInitializableOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableOwnableFilterer{contract: contract}, nil
}

// bindInitializableOwnable binds a generic wrapper to an already deployed contract.
func bindInitializableOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableOwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitializableOwnable *InitializableOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitializableOwnable.Contract.InitializableOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitializableOwnable *InitializableOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.InitializableOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitializableOwnable *InitializableOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.InitializableOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InitializableOwnable *InitializableOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InitializableOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InitializableOwnable *InitializableOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InitializableOwnable *InitializableOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InitializableOwnable *InitializableOwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InitializableOwnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InitializableOwnable *InitializableOwnableSession) Owner() (common.Address, error) {
	return _InitializableOwnable.Contract.Owner(&_InitializableOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InitializableOwnable *InitializableOwnableCallerSession) Owner() (common.Address, error) {
	return _InitializableOwnable.Contract.Owner(&_InitializableOwnable.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_InitializableOwnable *InitializableOwnableTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _InitializableOwnable.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_InitializableOwnable *InitializableOwnableSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.Initialize(&_InitializableOwnable.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_InitializableOwnable *InitializableOwnableTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.Initialize(&_InitializableOwnable.TransactOpts, initialOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InitializableOwnable *InitializableOwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InitializableOwnable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InitializableOwnable *InitializableOwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.TransferOwnership(&_InitializableOwnable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InitializableOwnable *InitializableOwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InitializableOwnable.Contract.TransferOwnership(&_InitializableOwnable.TransactOpts, newOwner)
}

// InitializableOwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InitializableOwnable contract.
type InitializableOwnableOwnershipTransferredIterator struct {
	Event *InitializableOwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InitializableOwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableOwnableOwnershipTransferred)
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
		it.Event = new(InitializableOwnableOwnershipTransferred)
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
func (it *InitializableOwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableOwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableOwnableOwnershipTransferred represents a OwnershipTransferred event raised by the InitializableOwnable contract.
type InitializableOwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InitializableOwnable *InitializableOwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InitializableOwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InitializableOwnable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InitializableOwnableOwnershipTransferredIterator{contract: _InitializableOwnable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InitializableOwnable *InitializableOwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InitializableOwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InitializableOwnable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableOwnableOwnershipTransferred)
				if err := _InitializableOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InitializableOwnable *InitializableOwnableFilterer) ParseOwnershipTransferred(log types.Log) (*InitializableOwnableOwnershipTransferred, error) {
	event := new(InitializableOwnableOwnershipTransferred)
	if err := _InitializableOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
