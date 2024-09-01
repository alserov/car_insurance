// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_activeTill\",\"type\":\"uint256\"}],\"name\":\"Insure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mult\",\"type\":\"uint256\"}],\"name\":\"Payoff\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50600180546001600160a01b031916331790555f805561048e806100315f395ff3fe608060405260043610610028575f3560e01c806306f864191461002c57806368187a321461004d575b5f80fd5b348015610037575f80fd5b5061004b6100463660046103ab565b610060565b005b61004b61005b3660046103cb565b610199565b5f82116100b45760405162461bcd60e51b815260206004820152601b60248201527f76616c75652073686f756c64206265206d6f7265207468616e2030000000000060448201526064015b60405180910390fd5b335f9081526003602052604090205460ff16156101135760405162461bcd60e51b815260206004820152601860248201527f696e737572616e636520616c726561647920657869737473000000000000000060448201526064016100ab565b604080516060810182523380825260208083018681528385018681525f9384526002808452868520955186546001600160a01b0319166001600160a01b0390911617865591516001808701919091559051949091019390935560039052918220805460ff19169091179055805483919081906101909084906103f6565b90915550505050565b335f9081526003602052604090205460ff16156101f85760405162461bcd60e51b815260206004820152601860248201527f696e737572616e636520646f6573206e6f74206578697374000000000000000060448201526064016100ab565b335f90815260026020818152604092839020835160608101855281546001600160a01b031681526001820154928101929092529091015491810182905290421161027b5760405162461bcd60e51b81526020600482015260146024820152731a5b9cdd5c985b98d9481a5cc8195e1c1a5c995960621b60448201526064016100ab565b81156102be5760405162461bcd60e51b815260206004820152601260248201527134b73b30b634b21036bab63a34b83634b2b960711b60448201526064016100ab565b5f60648383602001516102d1919061040f565b6102db9190610426565b82602001516102ea91906103f6565b82516040519192506001600160a01b03169082156108fc029083905f818181858888f193505050506103555760405162461bcd60e51b81526020600482015260146024820152733330b4b632b2103a379039b2b7321022ba3432b960611b60448201526064016100ab565b335f90815260026020818152604080842080546001600160a01b031916815560018101859055909201839055600390528120805460ff19169055805482919081906103a1908490610445565b9091555050505050565b5f80604083850312156103bc575f80fd5b50508035926020909101359150565b5f602082840312156103db575f80fd5b5035919050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610409576104096103e2565b92915050565b8082028115828204841417610409576104096103e2565b5f8261044057634e487b7160e01b5f52601260045260245ffd5b500490565b81810381811115610409576104096103e256fea2646970667358221220a7bcab167a66eef40f9a50554528511e56b53f59b667976988cde05bb1b3473964736f6c634300081a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Insure is a paid mutator transaction binding the contract method 0x06f86419.
//
// Solidity: function Insure(uint256 _amount, uint256 _activeTill) returns()
func (_Api *ApiTransactor) Insure(opts *bind.TransactOpts, _amount *big.Int, _activeTill *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Insure", _amount, _activeTill)
}

// Insure is a paid mutator transaction binding the contract method 0x06f86419.
//
// Solidity: function Insure(uint256 _amount, uint256 _activeTill) returns()
func (_Api *ApiSession) Insure(_amount *big.Int, _activeTill *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Insure(&_Api.TransactOpts, _amount, _activeTill)
}

// Insure is a paid mutator transaction binding the contract method 0x06f86419.
//
// Solidity: function Insure(uint256 _amount, uint256 _activeTill) returns()
func (_Api *ApiTransactorSession) Insure(_amount *big.Int, _activeTill *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Insure(&_Api.TransactOpts, _amount, _activeTill)
}

// Payoff is a paid mutator transaction binding the contract method 0x68187a32.
//
// Solidity: function Payoff(uint256 _mult) payable returns()
func (_Api *ApiTransactor) Payoff(opts *bind.TransactOpts, _mult *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Payoff", _mult)
}

// Payoff is a paid mutator transaction binding the contract method 0x68187a32.
//
// Solidity: function Payoff(uint256 _mult) payable returns()
func (_Api *ApiSession) Payoff(_mult *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Payoff(&_Api.TransactOpts, _mult)
}

// Payoff is a paid mutator transaction binding the contract method 0x68187a32.
//
// Solidity: function Payoff(uint256 _mult) payable returns()
func (_Api *ApiTransactorSession) Payoff(_mult *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Payoff(&_Api.TransactOpts, _mult)
}
