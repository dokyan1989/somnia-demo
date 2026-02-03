// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ReactiveCounterMetaData contains all meta data concerning the ReactiveCounter contract.
var ReactiveCounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"Increment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"inc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"incBy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReactiveCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use ReactiveCounterMetaData.ABI instead.
var ReactiveCounterABI = ReactiveCounterMetaData.ABI

// ReactiveCounter is an auto generated Go binding around an Ethereum contract.
type ReactiveCounter struct {
	ReactiveCounterCaller     // Read-only binding to the contract
	ReactiveCounterTransactor // Write-only binding to the contract
	ReactiveCounterFilterer   // Log filterer for contract events
}

// ReactiveCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReactiveCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReactiveCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReactiveCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReactiveCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReactiveCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReactiveCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReactiveCounterSession struct {
	Contract     *ReactiveCounter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReactiveCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReactiveCounterCallerSession struct {
	Contract *ReactiveCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReactiveCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReactiveCounterTransactorSession struct {
	Contract     *ReactiveCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReactiveCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReactiveCounterRaw struct {
	Contract *ReactiveCounter // Generic contract binding to access the raw methods on
}

// ReactiveCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReactiveCounterCallerRaw struct {
	Contract *ReactiveCounterCaller // Generic read-only contract binding to access the raw methods on
}

// ReactiveCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReactiveCounterTransactorRaw struct {
	Contract *ReactiveCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReactiveCounter creates a new instance of ReactiveCounter, bound to a specific deployed contract.
func NewReactiveCounter(address common.Address, backend bind.ContractBackend) (*ReactiveCounter, error) {
	contract, err := bindReactiveCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReactiveCounter{ReactiveCounterCaller: ReactiveCounterCaller{contract: contract}, ReactiveCounterTransactor: ReactiveCounterTransactor{contract: contract}, ReactiveCounterFilterer: ReactiveCounterFilterer{contract: contract}}, nil
}

// NewReactiveCounterCaller creates a new read-only instance of ReactiveCounter, bound to a specific deployed contract.
func NewReactiveCounterCaller(address common.Address, caller bind.ContractCaller) (*ReactiveCounterCaller, error) {
	contract, err := bindReactiveCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReactiveCounterCaller{contract: contract}, nil
}

// NewReactiveCounterTransactor creates a new write-only instance of ReactiveCounter, bound to a specific deployed contract.
func NewReactiveCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*ReactiveCounterTransactor, error) {
	contract, err := bindReactiveCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReactiveCounterTransactor{contract: contract}, nil
}

// NewReactiveCounterFilterer creates a new log filterer instance of ReactiveCounter, bound to a specific deployed contract.
func NewReactiveCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*ReactiveCounterFilterer, error) {
	contract, err := bindReactiveCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReactiveCounterFilterer{contract: contract}, nil
}

// bindReactiveCounter binds a generic wrapper to an already deployed contract.
func bindReactiveCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReactiveCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReactiveCounter *ReactiveCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReactiveCounter.Contract.ReactiveCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReactiveCounter *ReactiveCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReactiveCounter.Contract.ReactiveCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReactiveCounter *ReactiveCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReactiveCounter.Contract.ReactiveCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReactiveCounter *ReactiveCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReactiveCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReactiveCounter *ReactiveCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReactiveCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReactiveCounter *ReactiveCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReactiveCounter.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_ReactiveCounter *ReactiveCounterCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReactiveCounter.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_ReactiveCounter *ReactiveCounterSession) X() (*big.Int, error) {
	return _ReactiveCounter.Contract.X(&_ReactiveCounter.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_ReactiveCounter *ReactiveCounterCallerSession) X() (*big.Int, error) {
	return _ReactiveCounter.Contract.X(&_ReactiveCounter.CallOpts)
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_ReactiveCounter *ReactiveCounterTransactor) Inc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReactiveCounter.contract.Transact(opts, "inc")
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_ReactiveCounter *ReactiveCounterSession) Inc() (*types.Transaction, error) {
	return _ReactiveCounter.Contract.Inc(&_ReactiveCounter.TransactOpts)
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_ReactiveCounter *ReactiveCounterTransactorSession) Inc() (*types.Transaction, error) {
	return _ReactiveCounter.Contract.Inc(&_ReactiveCounter.TransactOpts)
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_ReactiveCounter *ReactiveCounterTransactor) IncBy(opts *bind.TransactOpts, by *big.Int) (*types.Transaction, error) {
	return _ReactiveCounter.contract.Transact(opts, "incBy", by)
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_ReactiveCounter *ReactiveCounterSession) IncBy(by *big.Int) (*types.Transaction, error) {
	return _ReactiveCounter.Contract.IncBy(&_ReactiveCounter.TransactOpts, by)
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_ReactiveCounter *ReactiveCounterTransactorSession) IncBy(by *big.Int) (*types.Transaction, error) {
	return _ReactiveCounter.Contract.IncBy(&_ReactiveCounter.TransactOpts, by)
}

// ReactiveCounterIncrementIterator is returned from FilterIncrement and is used to iterate over the raw logs and unpacked data for Increment events raised by the ReactiveCounter contract.
type ReactiveCounterIncrementIterator struct {
	Event *ReactiveCounterIncrement // Event containing the contract specifics and raw log

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
func (it *ReactiveCounterIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReactiveCounterIncrement)
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
		it.Event = new(ReactiveCounterIncrement)
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
func (it *ReactiveCounterIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReactiveCounterIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReactiveCounterIncrement represents a Increment event raised by the ReactiveCounter contract.
type ReactiveCounterIncrement struct {
	By  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIncrement is a free log retrieval operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_ReactiveCounter *ReactiveCounterFilterer) FilterIncrement(opts *bind.FilterOpts) (*ReactiveCounterIncrementIterator, error) {

	logs, sub, err := _ReactiveCounter.contract.FilterLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return &ReactiveCounterIncrementIterator{contract: _ReactiveCounter.contract, event: "Increment", logs: logs, sub: sub}, nil
}

// WatchIncrement is a free log subscription operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_ReactiveCounter *ReactiveCounterFilterer) WatchIncrement(opts *bind.WatchOpts, sink chan<- *ReactiveCounterIncrement) (event.Subscription, error) {

	logs, sub, err := _ReactiveCounter.contract.WatchLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReactiveCounterIncrement)
				if err := _ReactiveCounter.contract.UnpackLog(event, "Increment", log); err != nil {
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

// ParseIncrement is a log parse operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_ReactiveCounter *ReactiveCounterFilterer) ParseIncrement(log types.Log) (*ReactiveCounterIncrement, error) {
	event := new(ReactiveCounterIncrement)
	if err := _ReactiveCounter.contract.UnpackLog(event, "Increment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
