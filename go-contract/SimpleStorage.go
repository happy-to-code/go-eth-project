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
)

// SimpleStorageMetaData contains all meta data concerning the SimpleStorage contract.
var SimpleStorageMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_msg\",\"type\":\"string\"}],\"name\":\"setMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"src\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// SimpleStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleStorageMetaData.ABI instead.
var SimpleStorageABI = SimpleStorageMetaData.ABI

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleStorageCallerSession struct {
	Contract *SimpleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleStorageTransactorSession struct {
	Contract     *SimpleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleStorageRaw struct {
	Contract *SimpleStorage // Generic contract binding to access the raw methods on
}

// SimpleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleStorageCallerRaw struct {
	Contract *SimpleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleStorageTransactorRaw struct {
	Contract *SimpleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// NewSimpleStorageCaller creates a new read-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageCaller(address common.Address, caller bind.ContractCaller) (*SimpleStorageCaller, error) {
	contract, err := bindSimpleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageCaller{contract: contract}, nil
}

// NewSimpleStorageTransactor creates a new write-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleStorageTransactor, error) {
	contract, err := bindSimpleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageTransactor{contract: contract}, nil
}

// NewSimpleStorageFilterer creates a new log filterer instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleStorageFilterer, error) {
	contract, err := bindSimpleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageFilterer{contract: contract}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.SimpleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transact(opts, method, params...)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_SimpleStorage *SimpleStorageCaller) GetMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_SimpleStorage *SimpleStorageSession) GetMessage() (string, error) {
	return _SimpleStorage.Contract.GetMessage(&_SimpleStorage.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_SimpleStorage *SimpleStorageCallerSession) GetMessage() (string, error) {
	return _SimpleStorage.Contract.GetMessage(&_SimpleStorage.CallOpts)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _msg) returns()
func (_SimpleStorage *SimpleStorageTransactor) SetMessage(opts *bind.TransactOpts, _msg string) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "setMessage", _msg)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _msg) returns()
func (_SimpleStorage *SimpleStorageSession) SetMessage(_msg string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetMessage(&_SimpleStorage.TransactOpts, _msg)
}

// SetMessage is a paid mutator transaction binding the contract method 0x368b8772.
//
// Solidity: function setMessage(string _msg) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) SetMessage(_msg string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetMessage(&_SimpleStorage.TransactOpts, _msg)
}
