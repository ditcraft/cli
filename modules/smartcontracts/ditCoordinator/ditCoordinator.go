// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ditCoordinator

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KNWTokenContractABI is the input ABI used to generate the binding from.
const KNWTokenContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newVotingAddress\",\"type\":\"address\"}],\"name\":\"setVotingAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KNWTokenContractBin is the compiled bytecode used for deploying new contracts.
const KNWTokenContractBin = `0x`

// DeployKNWTokenContract deploys a new Ethereum contract, binding an instance of KNWTokenContract to it.
func DeployKNWTokenContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KNWTokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWTokenContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWTokenContract{KNWTokenContractCaller: KNWTokenContractCaller{contract: contract}, KNWTokenContractTransactor: KNWTokenContractTransactor{contract: contract}, KNWTokenContractFilterer: KNWTokenContractFilterer{contract: contract}}, nil
}

// KNWTokenContract is an auto generated Go binding around an Ethereum contract.
type KNWTokenContract struct {
	KNWTokenContractCaller     // Read-only binding to the contract
	KNWTokenContractTransactor // Write-only binding to the contract
	KNWTokenContractFilterer   // Log filterer for contract events
}

// KNWTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWTokenContractSession struct {
	Contract     *KNWTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNWTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWTokenContractCallerSession struct {
	Contract *KNWTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KNWTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWTokenContractTransactorSession struct {
	Contract     *KNWTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KNWTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWTokenContractRaw struct {
	Contract *KNWTokenContract // Generic contract binding to access the raw methods on
}

// KNWTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWTokenContractCallerRaw struct {
	Contract *KNWTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// KNWTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWTokenContractTransactorRaw struct {
	Contract *KNWTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWTokenContract creates a new instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContract(address common.Address, backend bind.ContractBackend) (*KNWTokenContract, error) {
	contract, err := bindKNWTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContract{KNWTokenContractCaller: KNWTokenContractCaller{contract: contract}, KNWTokenContractTransactor: KNWTokenContractTransactor{contract: contract}, KNWTokenContractFilterer: KNWTokenContractFilterer{contract: contract}}, nil
}

// NewKNWTokenContractCaller creates a new read-only instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractCaller(address common.Address, caller bind.ContractCaller) (*KNWTokenContractCaller, error) {
	contract, err := bindKNWTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractCaller{contract: contract}, nil
}

// NewKNWTokenContractTransactor creates a new write-only instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWTokenContractTransactor, error) {
	contract, err := bindKNWTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractTransactor{contract: contract}, nil
}

// NewKNWTokenContractFilterer creates a new log filterer instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWTokenContractFilterer, error) {
	contract, err := bindKNWTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractFilterer{contract: contract}, nil
}

// bindKNWTokenContract binds a generic wrapper to an already deployed contract.
func bindKNWTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWTokenContract *KNWTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWTokenContract.Contract.KNWTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWTokenContract *KNWTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.KNWTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWTokenContract *KNWTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.KNWTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWTokenContract *KNWTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWTokenContract *KNWTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWTokenContract *KNWTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.contract.Transact(opts, method, params...)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(_newVotingAddress address) returns()
func (_KNWTokenContract *KNWTokenContractTransactor) SetVotingAddress(opts *bind.TransactOpts, _newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "setVotingAddress", _newVotingAddress)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(_newVotingAddress address) returns()
func (_KNWTokenContract *KNWTokenContractSession) SetVotingAddress(_newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.SetVotingAddress(&_KNWTokenContract.TransactOpts, _newVotingAddress)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(_newVotingAddress address) returns()
func (_KNWTokenContract *KNWTokenContractTransactorSession) SetVotingAddress(_newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.SetVotingAddress(&_KNWTokenContract.TransactOpts, _newVotingAddress)
}

// KNWVotingContractABI is the input ABI used to generate the binding from.
const KNWVotingContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newKNWTokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"numVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeLabel\",\"type\":\"string\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newContract\",\"type\":\"address\"},{\"name\":\"_majority\",\"type\":\"uint256\"},{\"name\":\"_mintingMethod\",\"type\":\"uint256\"},{\"name\":\"_burningMethod\",\"type\":\"uint256\"}],\"name\":\"addDitContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"resolveVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"resolvePoll\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCoordinatorAddress\",\"type\":\"address\"}],\"name\":\"setCoordinatorAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KNWVotingContractBin is the compiled bytecode used for deploying new contracts.
const KNWVotingContractBin = `0x`

// DeployKNWVotingContract deploys a new Ethereum contract, binding an instance of KNWVotingContract to it.
func DeployKNWVotingContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KNWVotingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWVotingContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWVotingContract{KNWVotingContractCaller: KNWVotingContractCaller{contract: contract}, KNWVotingContractTransactor: KNWVotingContractTransactor{contract: contract}, KNWVotingContractFilterer: KNWVotingContractFilterer{contract: contract}}, nil
}

// KNWVotingContract is an auto generated Go binding around an Ethereum contract.
type KNWVotingContract struct {
	KNWVotingContractCaller     // Read-only binding to the contract
	KNWVotingContractTransactor // Write-only binding to the contract
	KNWVotingContractFilterer   // Log filterer for contract events
}

// KNWVotingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWVotingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWVotingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWVotingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWVotingContractSession struct {
	Contract     *KNWVotingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KNWVotingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWVotingContractCallerSession struct {
	Contract *KNWVotingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KNWVotingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWVotingContractTransactorSession struct {
	Contract     *KNWVotingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KNWVotingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWVotingContractRaw struct {
	Contract *KNWVotingContract // Generic contract binding to access the raw methods on
}

// KNWVotingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWVotingContractCallerRaw struct {
	Contract *KNWVotingContractCaller // Generic read-only contract binding to access the raw methods on
}

// KNWVotingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWVotingContractTransactorRaw struct {
	Contract *KNWVotingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWVotingContract creates a new instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContract(address common.Address, backend bind.ContractBackend) (*KNWVotingContract, error) {
	contract, err := bindKNWVotingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContract{KNWVotingContractCaller: KNWVotingContractCaller{contract: contract}, KNWVotingContractTransactor: KNWVotingContractTransactor{contract: contract}, KNWVotingContractFilterer: KNWVotingContractFilterer{contract: contract}}, nil
}

// NewKNWVotingContractCaller creates a new read-only instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContractCaller(address common.Address, caller bind.ContractCaller) (*KNWVotingContractCaller, error) {
	contract, err := bindKNWVotingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContractCaller{contract: contract}, nil
}

// NewKNWVotingContractTransactor creates a new write-only instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWVotingContractTransactor, error) {
	contract, err := bindKNWVotingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContractTransactor{contract: contract}, nil
}

// NewKNWVotingContractFilterer creates a new log filterer instance of KNWVotingContract, bound to a specific deployed contract.
func NewKNWVotingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWVotingContractFilterer, error) {
	contract, err := bindKNWVotingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWVotingContractFilterer{contract: contract}, nil
}

// bindKNWVotingContract binds a generic wrapper to an already deployed contract.
func bindKNWVotingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVotingContract *KNWVotingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVotingContract.Contract.KNWVotingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVotingContract *KNWVotingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.KNWVotingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVotingContract *KNWVotingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.KNWVotingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVotingContract *KNWVotingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVotingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVotingContract *KNWVotingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVotingContract *KNWVotingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.contract.Transact(opts, method, params...)
}

// ResolveVote is a free data retrieval call binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(_pollID uint256, _voteOption uint256, _address address) constant returns(reward uint256)
func (_KNWVotingContract *KNWVotingContractCaller) ResolveVote(opts *bind.CallOpts, _pollID *big.Int, _voteOption *big.Int, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVotingContract.contract.Call(opts, out, "resolveVote", _pollID, _voteOption, _address)
	return *ret0, err
}

// ResolveVote is a free data retrieval call binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(_pollID uint256, _voteOption uint256, _address address) constant returns(reward uint256)
func (_KNWVotingContract *KNWVotingContractSession) ResolveVote(_pollID *big.Int, _voteOption *big.Int, _address common.Address) (*big.Int, error) {
	return _KNWVotingContract.Contract.ResolveVote(&_KNWVotingContract.CallOpts, _pollID, _voteOption, _address)
}

// ResolveVote is a free data retrieval call binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(_pollID uint256, _voteOption uint256, _address address) constant returns(reward uint256)
func (_KNWVotingContract *KNWVotingContractCallerSession) ResolveVote(_pollID *big.Int, _voteOption *big.Int, _address common.Address) (*big.Int, error) {
	return _KNWVotingContract.Contract.ResolveVote(&_KNWVotingContract.CallOpts, _pollID, _voteOption, _address)
}

// AddDitContract is a paid mutator transaction binding the contract method 0xb6818a52.
//
// Solidity: function addDitContract(_newContract address, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) AddDitContract(opts *bind.TransactOpts, _newContract common.Address, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "addDitContract", _newContract, _majority, _mintingMethod, _burningMethod)
}

// AddDitContract is a paid mutator transaction binding the contract method 0xb6818a52.
//
// Solidity: function addDitContract(_newContract address, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVotingContract *KNWVotingContractSession) AddDitContract(_newContract common.Address, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddDitContract(&_KNWVotingContract.TransactOpts, _newContract, _majority, _mintingMethod, _burningMethod)
}

// AddDitContract is a paid mutator transaction binding the contract method 0xb6818a52.
//
// Solidity: function addDitContract(_newContract address, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) AddDitContract(_newContract common.Address, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddDitContract(&_KNWVotingContract.TransactOpts, _newContract, _majority, _mintingMethod, _burningMethod)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(_pollID uint256, _address address, _secretHash bytes32) returns(numVotes uint256)
func (_KNWVotingContract *KNWVotingContractTransactor) CommitVote(opts *bind.TransactOpts, _pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "commitVote", _pollID, _address, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(_pollID uint256, _address address, _secretHash bytes32) returns(numVotes uint256)
func (_KNWVotingContract *KNWVotingContractSession) CommitVote(_pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.CommitVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(_pollID uint256, _address address, _secretHash bytes32) returns(numVotes uint256)
func (_KNWVotingContract *KNWVotingContractTransactorSession) CommitVote(_pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.CommitVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _secretHash)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(_pollID uint256) returns(votePassed bool)
func (_KNWVotingContract *KNWVotingContractTransactor) ResolvePoll(opts *bind.TransactOpts, _pollID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "resolvePoll", _pollID)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(_pollID uint256) returns(votePassed bool)
func (_KNWVotingContract *KNWVotingContractSession) ResolvePoll(_pollID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.ResolvePoll(&_KNWVotingContract.TransactOpts, _pollID)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(_pollID uint256) returns(votePassed bool)
func (_KNWVotingContract *KNWVotingContractTransactorSession) ResolvePoll(_pollID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.ResolvePoll(&_KNWVotingContract.TransactOpts, _pollID)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(_pollID uint256, _address address, _voteOption uint256, _salt uint256) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) RevealVote(opts *bind.TransactOpts, _pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "revealVote", _pollID, _address, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(_pollID uint256, _address address, _voteOption uint256, _salt uint256) returns()
func (_KNWVotingContract *KNWVotingContractSession) RevealVote(_pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.RevealVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(_pollID uint256, _address address, _voteOption uint256, _salt uint256) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) RevealVote(_pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.RevealVote(&_KNWVotingContract.TransactOpts, _pollID, _address, _voteOption, _salt)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(_newCoordinatorAddress address) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) SetCoordinatorAddress(opts *bind.TransactOpts, _newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "setCoordinatorAddress", _newCoordinatorAddress)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(_newCoordinatorAddress address) returns()
func (_KNWVotingContract *KNWVotingContractSession) SetCoordinatorAddress(_newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetCoordinatorAddress(&_KNWVotingContract.TransactOpts, _newCoordinatorAddress)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(_newCoordinatorAddress address) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) SetCoordinatorAddress(_newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetCoordinatorAddress(&_KNWVotingContract.TransactOpts, _newCoordinatorAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_newKNWTokenAddress address) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) SetTokenAddress(opts *bind.TransactOpts, _newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "setTokenAddress", _newKNWTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_newKNWTokenAddress address) returns()
func (_KNWVotingContract *KNWVotingContractSession) SetTokenAddress(_newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetTokenAddress(&_KNWVotingContract.TransactOpts, _newKNWTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_newKNWTokenAddress address) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) SetTokenAddress(_newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.SetTokenAddress(&_KNWVotingContract.TransactOpts, _newKNWTokenAddress)
}

// StartPoll is a paid mutator transaction binding the contract method 0xa663fce8.
//
// Solidity: function startPoll(_address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVotingContract *KNWVotingContractTransactor) StartPoll(opts *bind.TransactOpts, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "startPoll", _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0xa663fce8.
//
// Solidity: function startPoll(_address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVotingContract *KNWVotingContractSession) StartPoll(_address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartPoll(&_KNWVotingContract.TransactOpts, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0xa663fce8.
//
// Solidity: function startPoll(_address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVotingContract *KNWVotingContractTransactorSession) StartPoll(_address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartPoll(&_KNWVotingContract.TransactOpts, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"sqrt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x61016b610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100575763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663677342ce811461005c575b600080fd5b610067600435610079565b60408051918252519081900360200190f35b6000808083151561008d5760009250610138565b6001840184106100fe57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f466c6177656420696e70757420666f7220737172740000000000000000000000604482015290519081900360640190fd5b505060026001830104825b80821015610134575080600281808681151561012157fe5b040181151561012c57fe5b049150610109565b8192505b50509190505600a165627a7a72305820d59bfc8115a4b45bda73e85f69809a96cd5e3191267d436185f8ed2e3af2acf50029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(a uint256) constant returns(uint256)
func (_SafeMath *SafeMathCaller) Sqrt(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMath.contract.Call(opts, out, "sqrt", a)
	return *ret0, err
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(a uint256) constant returns(uint256)
func (_SafeMath *SafeMathSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(a uint256) constant returns(uint256)
func (_SafeMath *SafeMathCallerSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// DitContractABI is the input ABI used to generate the binding from.
const DitContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isResolved\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_knowledgeLabelIndex\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"knowledgeLabels\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"resolveVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentProposalID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"revealVoteOnProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ditCoordinatorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"voteOnProposal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULT_COMMIT_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULT_REVEAL_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getRequiredStake\",\"outputs\":[{\"name\":\"requiredStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repository\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"},{\"name\":\"_ditCoordinatorAddress\",\"type\":\"address\"},{\"name\":\"_repository\",\"type\":\"string\"},{\"name\":\"_label1\",\"type\":\"string\"},{\"name\":\"_label2\",\"type\":\"string\"},{\"name\":\"_label3\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"CommitProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"Reveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"ProposalResolved\",\"type\":\"event\"}]"

// DitContractBin is the compiled bytecode used for deploying new contracts.
const DitContractBin = `0x60806040523480156200001157600080fd5b5060405162001a9838038062001a9883398101604090815281516020830151918301516060840151608085015160a0860151939592830193918301929081019101600160a060020a03861615801590620000735750600160a060020a03851615155b15156200010757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f4b4e57566f74696e6720616e6420646974436f6f7264696e61746f722061646460448201527f726573732063616e277420626520656d70747900000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03808916600160a060020a0319928316179283905560018054898316908416179055600280549092169216919091179055835162000158906003906020870190620002ca565b506000600581905583511115620001a157600480546001810180835560009290925284516200019e9160008051602062001a7883398151915201906020870190620002ca565b50505b600082511115620001e45760048054600181018083556000929092528351620001e19160008051602062001a7883398151915201906020860190620002ca565b50505b600082511115620002275760048054600181018083556000929092528251620002249160008051602062001a7883398151915201906020850190620002ca565b50505b600454600010620002be57604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f50726f76696465206174206c65617374206f6e65204b6e6f776c65646765204c60448201527f6162656c00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5050505050506200036f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030d57805160ff19168380011785556200033d565b828001600101855582156200033d579182015b828111156200033d57825182559160200191906001019062000320565b506200034b9291506200034f565b5090565b6200036c91905b808211156200034b576000815560010162000356565b90565b6116f9806200037f6000396000f3006080604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663013cf08b81146100d4578063067462d31461019f5780631c97d92b146101cb5780631f8ff8e6146101d857806330f4ccac14610265578063639661901461027d578063900f0d39146102a457806393b17b8e146102c2578063bfac6efa146102f3578063c030638714610301578063c4e4b02a14610301578063dce6636c14610316578063e9176c601461032e578063eb49fe9414610343575b600080fd5b3480156100e057600080fd5b506100ec600435610358565b60408051888152600160a060020a038716918101919091528415156060820152831515608082015260a0810183905260c0810182905260e0602080830182815289519284019290925288516101008401918a019080838360005b8381101561015e578181015183820152602001610146565b50505050905090810190601f16801561018b5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156101ab57600080fd5b506101b7600435610445565b604080519115158252519081900360200190f35b6101d66004356104e9565b005b3480156101e457600080fd5b506101f0600435610a44565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022a578181015183820152602001610212565b50505050905090810190601f1680156102575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027157600080fd5b506101d6600435610aeb565b34801561028957600080fd5b50610292610fad565b60408051918252519081900360200190f35b3480156102b057600080fd5b506101d6600435602435604435610fb3565b3480156102ce57600080fd5b506102d761114a565b60408051600160a060020a039092168252519081900360200190f35b6101d6600435602435611159565b34801561030d57600080fd5b5061029261149b565b34801561032257600080fd5b506102926004356114a1565b34801561033a57600080fd5b506101f06114b6565b34801561034f57600080fd5b506102d7611511565b6006602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f8101869004860283018601909652858252919492939092908301828280156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b505050600284015460038501546004909501549394600160a060020a0382169460ff60a060020a840481169550750100000000000000000000000000000000000000000090930490921692509087565b60008181526006602052604081206002015460a060020a900460ff1615156104b7576040805160e560020a62461bcd02815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b506000908152600660205260409020600201547501000000000000000000000000000000000000000000900460ff1690565b60003411610567576040805160e560020a62461bcd02815260206004820152602860248201527f56616c7565206f6620746865207472616e73616374696f6e2063616e206e6f7460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600454600019018111156105ea576040805160e560020a62461bcd028152602060048201526024808201527f4b6e6f776c656467652d4c6162656c20696e646578206973206e6f7420636f7260448201527f7265637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6005546105fe90600163ffffffff61152016565b6005556040805160e08101909152600254600480548392600160a060020a03169163a663fce8913391908790811061063257fe5b9060005260206000200161012c80346040518663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018086600160a060020a0316600160a060020a031681526020018060200185815260200184815260200183815260200182810382528681815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561071f5780601f106106f45761010080835404028352916020019161071f565b820191906000526020600020905b81548152906001019060200180831161070257829003601f168201915b50509650505050505050602060405180830381600087803b15801561074357600080fd5b505af1158015610757573d6000803e3d6000fd5b505050506040513d602081101561076d57600080fd5b50518152600480546020909201918490811061078557fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156108135780601f106107e857610100808354040283529160200191610813565b820191906000526020600020905b8154815290600101906020018083116107f657829003601f168201915b505050918352505033602080830191909152600060408084018290526060840182905234608085015260a09093018190526005548152600682529190912082518155828201518051919261086f92600185019290910190611632565b5060408281015160028301805460608601516080870151151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001991151560a060020a0274ff000000000000000000000000000000000000000019600160a060020a0390961673ffffffffffffffffffffffffffffffffffffffff199094169390931794909416919091171691909117905560a0830151600383015560c0909201516004918201556005805460009081526006602081815285832033845284018152858320349081905593548352529290922001546109619163ffffffff61152016565b60058054600090815260066020526040902060049081019290925554815433927ff2c3b57260094fc6dc300b8881004fb6c60df6103955d42f7f70bf2bbbc7344f91859081106109ad57fe5b6000918252602091829020604080518481529290910180546002600182161561010002600019019091160493830184905292829182019084908015610a335780601f10610a0857610100808354040283529160200191610a33565b820191906000526020600020905b815481529060010190602001808311610a1657829003601f168201915b50509250505060405180910390a350565b6004805482908110610a5257fe5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815293509091830182828015610ae35780601f10610ab857610100808354040283529160200191610ae3565b820191906000526020600020905b815481529060010190602001808311610ac657829003601f168201915b505050505081565b60008181526006602052604081206002015460a060020a900460ff161515610ce55760025460008381526006602090815260408083205481517fe74fef3700000000000000000000000000000000000000000000000000000000815260048101919091529051600160a060020a039094169363e74fef3793602480840194938390030190829087803b158015610b8057600080fd5b505af1158015610b94573d6000803e3d6000fd5b505050506040513d6020811015610baa57600080fd5b50516000838152600660209081526040918290206002808201805460a060020a75ff00000000000000000000000000000000000000000019909116750100000000000000000000000000000000000000000097151588021774ff000000000000000000000000000000000000000019161790819055845160ff969091049590951680151593860193909352838552600191820180546101009381161593909302600019019092160492840183905285937f201bd7ec9db47a5155ba8bd82a15a9b162113c1dbc4494fe017b1c9ad2d0d0ec939192918190606082019085908015610cd55780601f10610caa57610100808354040283529160200191610cd5565b820191906000526020600020905b815481529060010190602001808311610cb857829003601f168201915b5050935050505060405180910390a25b60008281526006602090815260408083203384526005019091528120600101541180610d2a5750600082815260066020526040902060020154600160a060020a031633145b1515610da6576040805160e560020a62461bcd02815260206004820152603a60248201527f4f6e6c79207061727469636970616e7473206f662074686520766f746520617260448201527f652061626c6520746f207265736f6c76652074686520766f7465000000000000606482015290519081900360840190fd5b600082815260066020908152604080832033845260050190915281205411610e64576040805160e560020a62461bcd02815260206004820152604960248201527f4f6e6c79207061727469636970616e74732077686f20686176656e277420616c60448201527f7265616479207265736f6c7665642074686520766f7465206172652061626c6560648201527f20746f20646f20736f0000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600280546000848152600660209081526040808320805433808652600590920184528285209096015482517fce729fd200000000000000000000000000000000000000000000000000000000815260048101979097526024870152604486015251600160a060020a039093169363ce729fd29360648083019491928390030190829087803b158015610ef557600080fd5b505af1158015610f09573d6000803e3d6000fd5b505050506040513d6020811015610f1f57600080fd5b505190506000811115610f5b57604051339082156108fc029083906000818181858888f19350505050158015610f59573d6000803e3d6000fd5b505b6000828152600660208181526040808420338552600581018352908420849055928590525260040154610f94908263ffffffff6115aa16565b6000928352600660205260409092206004019190915550565b60055481565b6002546000848152600660205260408082205481517f34f2f2d2000000000000000000000000000000000000000000000000000000008152600481019190915233602482015260448101869052606481018590529051600160a060020a03909316926334f2f2d29260848084019391929182900301818387803b15801561103957600080fd5b505af115801561104d573d6000803e3d6000fd5b5050506000848152600660209081526040808320338085526005820184529382902060028082018990556001918201548451838b14968101879052948501819052606080865293830180549384161561010002600019019093169190910492840183905294955088947f0f93246f12982b82ea2f37862d166f15a04cd9aa4a60d80ea28299b2a0ff1c04949193919290919081906080820190869080156111355780601f1061110a57610100808354040283529160200191611135565b820191906000526020600020905b81548152906001019060200180831161111857829003601f168201915b505094505050505060405180910390a3505050565b600154600160a060020a031681565b60008281526006602052604081206003015434146111e7576040805160e560020a62461bcd02815260206004820152603960248201527f56616c7565206f6620746865207472616e73616374696f6e20646f65736e277460448201527f206d6174636820746865207265717569726564207374616b6500000000000000606482015290519081900360840190fd5b600083815260066020526040902060020154600160a060020a031633141561127f576040805160e560020a62461bcd02815260206004820152603160248201527f5468652070726f706f736572206973206e6f7420616c6c6f77656420746f207660448201527f6f746520696e20612070726f706f73616c000000000000000000000000000000606482015290519081900360840190fd5b6000838152600660205260409020600401546112a1903463ffffffff61152016565b6000848152600660208181526040808420600480820196909655338086526005820184528286203490556002548a8752948452905482517f7eb2ff520000000000000000000000000000000000000000000000000000000081529687015260248601526044850187905251600160a060020a0390921693637eb2ff5293606480830194928390030190829087803b15801561133b57600080fd5b505af115801561134f573d6000803e3d6000fd5b505050506040513d602081101561136557600080fd5b50519050600081116113e7576040805160e560020a62461bcd02815260206004820152603360248201527f566f74696e6720636f6e74726163742072657475726e656420616e20696e766160448201527f6c696420616d6f756e74206f6620766f74657300000000000000000000000000606482015290519081900360840190fd5b6000838152600660208181526040808420338086526005820184528286206001908101889055958990529383528151349381018490529182018690526060808352908501805460029681161561010002600019011695909504908201819052929387937ffe25e9ca71c8ef184b9997685aac106e44c6cbbf4986bf70085d7a67a5d9a103939192879181906080820190869080156111355780601f1061110a57610100808354040283529160200191611135565b61012c81565b60009081526006602052604090206003015490565b6003805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610ae35780601f10610ab857610100808354040283529160200191610ae3565b600054600160a060020a031681565b6000828201838110156115a3576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b6000808383111561162b576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061167357805160ff19168380011785556116a0565b828001600101855582156116a0579182015b828111156116a0578251825591602001919060010190611685565b506116ac9291506116b0565b5090565b6116ca91905b808211156116ac57600081556001016116b6565b905600a165627a7a72305820c3aacf13b3fd17948640e8e23c9776d002bc09b3ef5905ff72011c48b50e6d5700298a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b`

// DeployDitContract deploys a new Ethereum contract, binding an instance of DitContract to it.
func DeployDitContract(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWVotingAddress common.Address, _ditCoordinatorAddress common.Address, _repository string, _label1 string, _label2 string, _label3 string) (common.Address, *types.Transaction, *DitContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DitContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DitContractBin), backend, _KNWVotingAddress, _ditCoordinatorAddress, _repository, _label1, _label2, _label3)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DitContract{DitContractCaller: DitContractCaller{contract: contract}, DitContractTransactor: DitContractTransactor{contract: contract}, DitContractFilterer: DitContractFilterer{contract: contract}}, nil
}

// DitContract is an auto generated Go binding around an Ethereum contract.
type DitContract struct {
	DitContractCaller     // Read-only binding to the contract
	DitContractTransactor // Write-only binding to the contract
	DitContractFilterer   // Log filterer for contract events
}

// DitContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DitContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DitContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DitContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DitContractSession struct {
	Contract     *DitContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DitContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DitContractCallerSession struct {
	Contract *DitContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DitContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DitContractTransactorSession struct {
	Contract     *DitContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DitContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DitContractRaw struct {
	Contract *DitContract // Generic contract binding to access the raw methods on
}

// DitContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DitContractCallerRaw struct {
	Contract *DitContractCaller // Generic read-only contract binding to access the raw methods on
}

// DitContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DitContractTransactorRaw struct {
	Contract *DitContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDitContract creates a new instance of DitContract, bound to a specific deployed contract.
func NewDitContract(address common.Address, backend bind.ContractBackend) (*DitContract, error) {
	contract, err := bindDitContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DitContract{DitContractCaller: DitContractCaller{contract: contract}, DitContractTransactor: DitContractTransactor{contract: contract}, DitContractFilterer: DitContractFilterer{contract: contract}}, nil
}

// NewDitContractCaller creates a new read-only instance of DitContract, bound to a specific deployed contract.
func NewDitContractCaller(address common.Address, caller bind.ContractCaller) (*DitContractCaller, error) {
	contract, err := bindDitContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DitContractCaller{contract: contract}, nil
}

// NewDitContractTransactor creates a new write-only instance of DitContract, bound to a specific deployed contract.
func NewDitContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DitContractTransactor, error) {
	contract, err := bindDitContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DitContractTransactor{contract: contract}, nil
}

// NewDitContractFilterer creates a new log filterer instance of DitContract, bound to a specific deployed contract.
func NewDitContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DitContractFilterer, error) {
	contract, err := bindDitContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DitContractFilterer{contract: contract}, nil
}

// bindDitContract binds a generic wrapper to an already deployed contract.
func bindDitContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DitContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitContract *DitContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitContract.Contract.DitContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitContract *DitContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitContract.Contract.DitContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitContract *DitContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitContract.Contract.DitContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitContract *DitContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitContract *DitContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitContract *DitContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTCOMMITDURATION is a free data retrieval call binding the contract method 0xc0306387.
//
// Solidity: function DEFAULT_COMMIT_DURATION() constant returns(uint256)
func (_DitContract *DitContractCaller) DEFAULTCOMMITDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "DEFAULT_COMMIT_DURATION")
	return *ret0, err
}

// DEFAULTCOMMITDURATION is a free data retrieval call binding the contract method 0xc0306387.
//
// Solidity: function DEFAULT_COMMIT_DURATION() constant returns(uint256)
func (_DitContract *DitContractSession) DEFAULTCOMMITDURATION() (*big.Int, error) {
	return _DitContract.Contract.DEFAULTCOMMITDURATION(&_DitContract.CallOpts)
}

// DEFAULTCOMMITDURATION is a free data retrieval call binding the contract method 0xc0306387.
//
// Solidity: function DEFAULT_COMMIT_DURATION() constant returns(uint256)
func (_DitContract *DitContractCallerSession) DEFAULTCOMMITDURATION() (*big.Int, error) {
	return _DitContract.Contract.DEFAULTCOMMITDURATION(&_DitContract.CallOpts)
}

// DEFAULTREVEALDURATION is a free data retrieval call binding the contract method 0xc4e4b02a.
//
// Solidity: function DEFAULT_REVEAL_DURATION() constant returns(uint256)
func (_DitContract *DitContractCaller) DEFAULTREVEALDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "DEFAULT_REVEAL_DURATION")
	return *ret0, err
}

// DEFAULTREVEALDURATION is a free data retrieval call binding the contract method 0xc4e4b02a.
//
// Solidity: function DEFAULT_REVEAL_DURATION() constant returns(uint256)
func (_DitContract *DitContractSession) DEFAULTREVEALDURATION() (*big.Int, error) {
	return _DitContract.Contract.DEFAULTREVEALDURATION(&_DitContract.CallOpts)
}

// DEFAULTREVEALDURATION is a free data retrieval call binding the contract method 0xc4e4b02a.
//
// Solidity: function DEFAULT_REVEAL_DURATION() constant returns(uint256)
func (_DitContract *DitContractCallerSession) DEFAULTREVEALDURATION() (*big.Int, error) {
	return _DitContract.Contract.DEFAULTREVEALDURATION(&_DitContract.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitContract *DitContractCaller) KNWVotingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "KNWVotingAddress")
	return *ret0, err
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitContract *DitContractSession) KNWVotingAddress() (common.Address, error) {
	return _DitContract.Contract.KNWVotingAddress(&_DitContract.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitContract *DitContractCallerSession) KNWVotingAddress() (common.Address, error) {
	return _DitContract.Contract.KNWVotingAddress(&_DitContract.CallOpts)
}

// CurrentProposalID is a free data retrieval call binding the contract method 0x63966190.
//
// Solidity: function currentProposalID() constant returns(uint256)
func (_DitContract *DitContractCaller) CurrentProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "currentProposalID")
	return *ret0, err
}

// CurrentProposalID is a free data retrieval call binding the contract method 0x63966190.
//
// Solidity: function currentProposalID() constant returns(uint256)
func (_DitContract *DitContractSession) CurrentProposalID() (*big.Int, error) {
	return _DitContract.Contract.CurrentProposalID(&_DitContract.CallOpts)
}

// CurrentProposalID is a free data retrieval call binding the contract method 0x63966190.
//
// Solidity: function currentProposalID() constant returns(uint256)
func (_DitContract *DitContractCallerSession) CurrentProposalID() (*big.Int, error) {
	return _DitContract.Contract.CurrentProposalID(&_DitContract.CallOpts)
}

// DitCoordinatorAddress is a free data retrieval call binding the contract method 0x93b17b8e.
//
// Solidity: function ditCoordinatorAddress() constant returns(address)
func (_DitContract *DitContractCaller) DitCoordinatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "ditCoordinatorAddress")
	return *ret0, err
}

// DitCoordinatorAddress is a free data retrieval call binding the contract method 0x93b17b8e.
//
// Solidity: function ditCoordinatorAddress() constant returns(address)
func (_DitContract *DitContractSession) DitCoordinatorAddress() (common.Address, error) {
	return _DitContract.Contract.DitCoordinatorAddress(&_DitContract.CallOpts)
}

// DitCoordinatorAddress is a free data retrieval call binding the contract method 0x93b17b8e.
//
// Solidity: function ditCoordinatorAddress() constant returns(address)
func (_DitContract *DitContractCallerSession) DitCoordinatorAddress() (common.Address, error) {
	return _DitContract.Contract.DitCoordinatorAddress(&_DitContract.CallOpts)
}

// GetRequiredStake is a free data retrieval call binding the contract method 0xdce6636c.
//
// Solidity: function getRequiredStake(_proposalID uint256) constant returns(requiredStake uint256)
func (_DitContract *DitContractCaller) GetRequiredStake(opts *bind.CallOpts, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "getRequiredStake", _proposalID)
	return *ret0, err
}

// GetRequiredStake is a free data retrieval call binding the contract method 0xdce6636c.
//
// Solidity: function getRequiredStake(_proposalID uint256) constant returns(requiredStake uint256)
func (_DitContract *DitContractSession) GetRequiredStake(_proposalID *big.Int) (*big.Int, error) {
	return _DitContract.Contract.GetRequiredStake(&_DitContract.CallOpts, _proposalID)
}

// GetRequiredStake is a free data retrieval call binding the contract method 0xdce6636c.
//
// Solidity: function getRequiredStake(_proposalID uint256) constant returns(requiredStake uint256)
func (_DitContract *DitContractCallerSession) GetRequiredStake(_proposalID *big.Int) (*big.Int, error) {
	return _DitContract.Contract.GetRequiredStake(&_DitContract.CallOpts, _proposalID)
}

// KnowledgeLabels is a free data retrieval call binding the contract method 0x1f8ff8e6.
//
// Solidity: function knowledgeLabels( uint256) constant returns(string)
func (_DitContract *DitContractCaller) KnowledgeLabels(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "knowledgeLabels", arg0)
	return *ret0, err
}

// KnowledgeLabels is a free data retrieval call binding the contract method 0x1f8ff8e6.
//
// Solidity: function knowledgeLabels( uint256) constant returns(string)
func (_DitContract *DitContractSession) KnowledgeLabels(arg0 *big.Int) (string, error) {
	return _DitContract.Contract.KnowledgeLabels(&_DitContract.CallOpts, arg0)
}

// KnowledgeLabels is a free data retrieval call binding the contract method 0x1f8ff8e6.
//
// Solidity: function knowledgeLabels( uint256) constant returns(string)
func (_DitContract *DitContractCallerSession) KnowledgeLabels(arg0 *big.Int) (string, error) {
	return _DitContract.Contract.KnowledgeLabels(&_DitContract.CallOpts, arg0)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x067462d3.
//
// Solidity: function proposalHasPassed(_proposalID uint256) constant returns(hasPassed bool)
func (_DitContract *DitContractCaller) ProposalHasPassed(opts *bind.CallOpts, _proposalID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "proposalHasPassed", _proposalID)
	return *ret0, err
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x067462d3.
//
// Solidity: function proposalHasPassed(_proposalID uint256) constant returns(hasPassed bool)
func (_DitContract *DitContractSession) ProposalHasPassed(_proposalID *big.Int) (bool, error) {
	return _DitContract.Contract.ProposalHasPassed(&_DitContract.CallOpts, _proposalID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x067462d3.
//
// Solidity: function proposalHasPassed(_proposalID uint256) constant returns(hasPassed bool)
func (_DitContract *DitContractCallerSession) ProposalHasPassed(_proposalID *big.Int) (bool, error) {
	return _DitContract.Contract.ProposalHasPassed(&_DitContract.CallOpts, _proposalID)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals( uint256) constant returns(KNWVoteID uint256, knowledgeLabel string, proposer address, isResolved bool, proposalAccepted bool, requiredStake uint256, totalStake uint256)
func (_DitContract *DitContractCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsResolved       bool
	ProposalAccepted bool
	RequiredStake    *big.Int
	TotalStake       *big.Int
}, error) {
	ret := new(struct {
		KNWVoteID        *big.Int
		KnowledgeLabel   string
		Proposer         common.Address
		IsResolved       bool
		ProposalAccepted bool
		RequiredStake    *big.Int
		TotalStake       *big.Int
	})
	out := ret
	err := _DitContract.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals( uint256) constant returns(KNWVoteID uint256, knowledgeLabel string, proposer address, isResolved bool, proposalAccepted bool, requiredStake uint256, totalStake uint256)
func (_DitContract *DitContractSession) Proposals(arg0 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsResolved       bool
	ProposalAccepted bool
	RequiredStake    *big.Int
	TotalStake       *big.Int
}, error) {
	return _DitContract.Contract.Proposals(&_DitContract.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals( uint256) constant returns(KNWVoteID uint256, knowledgeLabel string, proposer address, isResolved bool, proposalAccepted bool, requiredStake uint256, totalStake uint256)
func (_DitContract *DitContractCallerSession) Proposals(arg0 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsResolved       bool
	ProposalAccepted bool
	RequiredStake    *big.Int
	TotalStake       *big.Int
}, error) {
	return _DitContract.Contract.Proposals(&_DitContract.CallOpts, arg0)
}

// Repository is a free data retrieval call binding the contract method 0xe9176c60.
//
// Solidity: function repository() constant returns(string)
func (_DitContract *DitContractCaller) Repository(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DitContract.contract.Call(opts, out, "repository")
	return *ret0, err
}

// Repository is a free data retrieval call binding the contract method 0xe9176c60.
//
// Solidity: function repository() constant returns(string)
func (_DitContract *DitContractSession) Repository() (string, error) {
	return _DitContract.Contract.Repository(&_DitContract.CallOpts)
}

// Repository is a free data retrieval call binding the contract method 0xe9176c60.
//
// Solidity: function repository() constant returns(string)
func (_DitContract *DitContractCallerSession) Repository() (string, error) {
	return _DitContract.Contract.Repository(&_DitContract.CallOpts)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x1c97d92b.
//
// Solidity: function proposeCommit(_knowledgeLabelIndex uint256) returns()
func (_DitContract *DitContractTransactor) ProposeCommit(opts *bind.TransactOpts, _knowledgeLabelIndex *big.Int) (*types.Transaction, error) {
	return _DitContract.contract.Transact(opts, "proposeCommit", _knowledgeLabelIndex)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x1c97d92b.
//
// Solidity: function proposeCommit(_knowledgeLabelIndex uint256) returns()
func (_DitContract *DitContractSession) ProposeCommit(_knowledgeLabelIndex *big.Int) (*types.Transaction, error) {
	return _DitContract.Contract.ProposeCommit(&_DitContract.TransactOpts, _knowledgeLabelIndex)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x1c97d92b.
//
// Solidity: function proposeCommit(_knowledgeLabelIndex uint256) returns()
func (_DitContract *DitContractTransactorSession) ProposeCommit(_knowledgeLabelIndex *big.Int) (*types.Transaction, error) {
	return _DitContract.Contract.ProposeCommit(&_DitContract.TransactOpts, _knowledgeLabelIndex)
}

// ResolveVote is a paid mutator transaction binding the contract method 0x30f4ccac.
//
// Solidity: function resolveVote(_proposalID uint256) returns()
func (_DitContract *DitContractTransactor) ResolveVote(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitContract.contract.Transact(opts, "resolveVote", _proposalID)
}

// ResolveVote is a paid mutator transaction binding the contract method 0x30f4ccac.
//
// Solidity: function resolveVote(_proposalID uint256) returns()
func (_DitContract *DitContractSession) ResolveVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _DitContract.Contract.ResolveVote(&_DitContract.TransactOpts, _proposalID)
}

// ResolveVote is a paid mutator transaction binding the contract method 0x30f4ccac.
//
// Solidity: function resolveVote(_proposalID uint256) returns()
func (_DitContract *DitContractTransactorSession) ResolveVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _DitContract.Contract.ResolveVote(&_DitContract.TransactOpts, _proposalID)
}

// RevealVoteOnProposal is a paid mutator transaction binding the contract method 0x900f0d39.
//
// Solidity: function revealVoteOnProposal(_proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitContract *DitContractTransactor) RevealVoteOnProposal(opts *bind.TransactOpts, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitContract.contract.Transact(opts, "revealVoteOnProposal", _proposalID, _voteOption, _voteSalt)
}

// RevealVoteOnProposal is a paid mutator transaction binding the contract method 0x900f0d39.
//
// Solidity: function revealVoteOnProposal(_proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitContract *DitContractSession) RevealVoteOnProposal(_proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitContract.Contract.RevealVoteOnProposal(&_DitContract.TransactOpts, _proposalID, _voteOption, _voteSalt)
}

// RevealVoteOnProposal is a paid mutator transaction binding the contract method 0x900f0d39.
//
// Solidity: function revealVoteOnProposal(_proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitContract *DitContractTransactorSession) RevealVoteOnProposal(_proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitContract.Contract.RevealVoteOnProposal(&_DitContract.TransactOpts, _proposalID, _voteOption, _voteSalt)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xbfac6efa.
//
// Solidity: function voteOnProposal(_proposalID uint256, _voteHash bytes32) returns()
func (_DitContract *DitContractTransactor) VoteOnProposal(opts *bind.TransactOpts, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitContract.contract.Transact(opts, "voteOnProposal", _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xbfac6efa.
//
// Solidity: function voteOnProposal(_proposalID uint256, _voteHash bytes32) returns()
func (_DitContract *DitContractSession) VoteOnProposal(_proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitContract.Contract.VoteOnProposal(&_DitContract.TransactOpts, _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xbfac6efa.
//
// Solidity: function voteOnProposal(_proposalID uint256, _voteHash bytes32) returns()
func (_DitContract *DitContractTransactorSession) VoteOnProposal(_proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitContract.Contract.VoteOnProposal(&_DitContract.TransactOpts, _proposalID, _voteHash)
}

// DitContractCommitProposalIterator is returned from FilterCommitProposal and is used to iterate over the raw logs and unpacked data for CommitProposal events raised by the DitContract contract.
type DitContractCommitProposalIterator struct {
	Event *DitContractCommitProposal // Event containing the contract specifics and raw log

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
func (it *DitContractCommitProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractCommitProposal)
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
		it.Event = new(DitContractCommitProposal)
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
func (it *DitContractCommitProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractCommitProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractCommitProposal represents a CommitProposal event raised by the DitContract contract.
type DitContractCommitProposal struct {
	Proposal *big.Int
	Who      common.Address
	Label    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitProposal is a free log retrieval operation binding the contract event 0xf2c3b57260094fc6dc300b8881004fb6c60df6103955d42f7f70bf2bbbc7344f.
//
// Solidity: e CommitProposal(proposal indexed uint256, who indexed address, label string)
func (_DitContract *DitContractFilterer) FilterCommitProposal(opts *bind.FilterOpts, proposal []*big.Int, who []common.Address) (*DitContractCommitProposalIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContract.contract.FilterLogs(opts, "CommitProposal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitContractCommitProposalIterator{contract: _DitContract.contract, event: "CommitProposal", logs: logs, sub: sub}, nil
}

// WatchCommitProposal is a free log subscription operation binding the contract event 0xf2c3b57260094fc6dc300b8881004fb6c60df6103955d42f7f70bf2bbbc7344f.
//
// Solidity: e CommitProposal(proposal indexed uint256, who indexed address, label string)
func (_DitContract *DitContractFilterer) WatchCommitProposal(opts *bind.WatchOpts, sink chan<- *DitContractCommitProposal, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContract.contract.WatchLogs(opts, "CommitProposal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractCommitProposal)
				if err := _DitContract.contract.UnpackLog(event, "CommitProposal", log); err != nil {
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

// DitContractProposalResolvedIterator is returned from FilterProposalResolved and is used to iterate over the raw logs and unpacked data for ProposalResolved events raised by the DitContract contract.
type DitContractProposalResolvedIterator struct {
	Event *DitContractProposalResolved // Event containing the contract specifics and raw log

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
func (it *DitContractProposalResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractProposalResolved)
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
		it.Event = new(DitContractProposalResolved)
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
func (it *DitContractProposalResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractProposalResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractProposalResolved represents a ProposalResolved event raised by the DitContract contract.
type DitContractProposalResolved struct {
	Proposal *big.Int
	Label    string
	Accepted bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalResolved is a free log retrieval operation binding the contract event 0x201bd7ec9db47a5155ba8bd82a15a9b162113c1dbc4494fe017b1c9ad2d0d0ec.
//
// Solidity: e ProposalResolved(proposal indexed uint256, label string, accepted bool)
func (_DitContract *DitContractFilterer) FilterProposalResolved(opts *bind.FilterOpts, proposal []*big.Int) (*DitContractProposalResolvedIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitContract.contract.FilterLogs(opts, "ProposalResolved", proposalRule)
	if err != nil {
		return nil, err
	}
	return &DitContractProposalResolvedIterator{contract: _DitContract.contract, event: "ProposalResolved", logs: logs, sub: sub}, nil
}

// WatchProposalResolved is a free log subscription operation binding the contract event 0x201bd7ec9db47a5155ba8bd82a15a9b162113c1dbc4494fe017b1c9ad2d0d0ec.
//
// Solidity: e ProposalResolved(proposal indexed uint256, label string, accepted bool)
func (_DitContract *DitContractFilterer) WatchProposalResolved(opts *bind.WatchOpts, sink chan<- *DitContractProposalResolved, proposal []*big.Int) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitContract.contract.WatchLogs(opts, "ProposalResolved", proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractProposalResolved)
				if err := _DitContract.contract.UnpackLog(event, "ProposalResolved", log); err != nil {
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

// DitContractRevealIterator is returned from FilterReveal and is used to iterate over the raw logs and unpacked data for Reveal events raised by the DitContract contract.
type DitContractRevealIterator struct {
	Event *DitContractReveal // Event containing the contract specifics and raw log

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
func (it *DitContractRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractReveal)
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
		it.Event = new(DitContractReveal)
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
func (it *DitContractRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractReveal represents a Reveal event raised by the DitContract contract.
type DitContractReveal struct {
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Accept        bool
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReveal is a free log retrieval operation binding the contract event 0x0f93246f12982b82ea2f37862d166f15a04cd9aa4a60d80ea28299b2a0ff1c04.
//
// Solidity: e Reveal(proposal indexed uint256, who indexed address, label string, accept bool, numberOfVotes uint256)
func (_DitContract *DitContractFilterer) FilterReveal(opts *bind.FilterOpts, proposal []*big.Int, who []common.Address) (*DitContractRevealIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContract.contract.FilterLogs(opts, "Reveal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitContractRevealIterator{contract: _DitContract.contract, event: "Reveal", logs: logs, sub: sub}, nil
}

// WatchReveal is a free log subscription operation binding the contract event 0x0f93246f12982b82ea2f37862d166f15a04cd9aa4a60d80ea28299b2a0ff1c04.
//
// Solidity: e Reveal(proposal indexed uint256, who indexed address, label string, accept bool, numberOfVotes uint256)
func (_DitContract *DitContractFilterer) WatchReveal(opts *bind.WatchOpts, sink chan<- *DitContractReveal, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContract.contract.WatchLogs(opts, "Reveal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractReveal)
				if err := _DitContract.contract.UnpackLog(event, "Reveal", log); err != nil {
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

// DitContractVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the DitContract contract.
type DitContractVoteIterator struct {
	Event *DitContractVote // Event containing the contract specifics and raw log

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
func (it *DitContractVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractVote)
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
		it.Event = new(DitContractVote)
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
func (it *DitContractVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractVote represents a Vote event raised by the DitContract contract.
type DitContractVote struct {
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Stake         *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xfe25e9ca71c8ef184b9997685aac106e44c6cbbf4986bf70085d7a67a5d9a103.
//
// Solidity: e Vote(proposal indexed uint256, who indexed address, label string, stake uint256, numberOfVotes uint256)
func (_DitContract *DitContractFilterer) FilterVote(opts *bind.FilterOpts, proposal []*big.Int, who []common.Address) (*DitContractVoteIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContract.contract.FilterLogs(opts, "Vote", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitContractVoteIterator{contract: _DitContract.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xfe25e9ca71c8ef184b9997685aac106e44c6cbbf4986bf70085d7a67a5d9a103.
//
// Solidity: e Vote(proposal indexed uint256, who indexed address, label string, stake uint256, numberOfVotes uint256)
func (_DitContract *DitContractFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *DitContractVote, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContract.contract.WatchLogs(opts, "Vote", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractVote)
				if err := _DitContract.contract.UnpackLog(event, "Vote", log); err != nil {
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

// DitContractInterfaceABI is the input ABI used to generate the binding from.
const DitContractInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_knowledgeLabelIndex\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"resolveVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"revealVoteOnProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"voteOnProposal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"CommitProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"Reveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"ProposalResolved\",\"type\":\"event\"}]"

// DitContractInterfaceBin is the compiled bytecode used for deploying new contracts.
const DitContractInterfaceBin = `0x`

// DeployDitContractInterface deploys a new Ethereum contract, binding an instance of DitContractInterface to it.
func DeployDitContractInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DitContractInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DitContractInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DitContractInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DitContractInterface{DitContractInterfaceCaller: DitContractInterfaceCaller{contract: contract}, DitContractInterfaceTransactor: DitContractInterfaceTransactor{contract: contract}, DitContractInterfaceFilterer: DitContractInterfaceFilterer{contract: contract}}, nil
}

// DitContractInterface is an auto generated Go binding around an Ethereum contract.
type DitContractInterface struct {
	DitContractInterfaceCaller     // Read-only binding to the contract
	DitContractInterfaceTransactor // Write-only binding to the contract
	DitContractInterfaceFilterer   // Log filterer for contract events
}

// DitContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DitContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DitContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DitContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DitContractInterfaceSession struct {
	Contract     *DitContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DitContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DitContractInterfaceCallerSession struct {
	Contract *DitContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DitContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DitContractInterfaceTransactorSession struct {
	Contract     *DitContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DitContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DitContractInterfaceRaw struct {
	Contract *DitContractInterface // Generic contract binding to access the raw methods on
}

// DitContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DitContractInterfaceCallerRaw struct {
	Contract *DitContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DitContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DitContractInterfaceTransactorRaw struct {
	Contract *DitContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDitContractInterface creates a new instance of DitContractInterface, bound to a specific deployed contract.
func NewDitContractInterface(address common.Address, backend bind.ContractBackend) (*DitContractInterface, error) {
	contract, err := bindDitContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DitContractInterface{DitContractInterfaceCaller: DitContractInterfaceCaller{contract: contract}, DitContractInterfaceTransactor: DitContractInterfaceTransactor{contract: contract}, DitContractInterfaceFilterer: DitContractInterfaceFilterer{contract: contract}}, nil
}

// NewDitContractInterfaceCaller creates a new read-only instance of DitContractInterface, bound to a specific deployed contract.
func NewDitContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DitContractInterfaceCaller, error) {
	contract, err := bindDitContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceCaller{contract: contract}, nil
}

// NewDitContractInterfaceTransactor creates a new write-only instance of DitContractInterface, bound to a specific deployed contract.
func NewDitContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DitContractInterfaceTransactor, error) {
	contract, err := bindDitContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceTransactor{contract: contract}, nil
}

// NewDitContractInterfaceFilterer creates a new log filterer instance of DitContractInterface, bound to a specific deployed contract.
func NewDitContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DitContractInterfaceFilterer, error) {
	contract, err := bindDitContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceFilterer{contract: contract}, nil
}

// bindDitContractInterface binds a generic wrapper to an already deployed contract.
func bindDitContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DitContractInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitContractInterface *DitContractInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitContractInterface.Contract.DitContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitContractInterface *DitContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitContractInterface.Contract.DitContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitContractInterface *DitContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitContractInterface.Contract.DitContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitContractInterface *DitContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitContractInterface *DitContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitContractInterface *DitContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitContractInterface.Contract.contract.Transact(opts, method, params...)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x067462d3.
//
// Solidity: function proposalHasPassed(_proposalID uint256) constant returns(hasPassed bool)
func (_DitContractInterface *DitContractInterfaceCaller) ProposalHasPassed(opts *bind.CallOpts, _proposalID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitContractInterface.contract.Call(opts, out, "proposalHasPassed", _proposalID)
	return *ret0, err
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x067462d3.
//
// Solidity: function proposalHasPassed(_proposalID uint256) constant returns(hasPassed bool)
func (_DitContractInterface *DitContractInterfaceSession) ProposalHasPassed(_proposalID *big.Int) (bool, error) {
	return _DitContractInterface.Contract.ProposalHasPassed(&_DitContractInterface.CallOpts, _proposalID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x067462d3.
//
// Solidity: function proposalHasPassed(_proposalID uint256) constant returns(hasPassed bool)
func (_DitContractInterface *DitContractInterfaceCallerSession) ProposalHasPassed(_proposalID *big.Int) (bool, error) {
	return _DitContractInterface.Contract.ProposalHasPassed(&_DitContractInterface.CallOpts, _proposalID)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x1c97d92b.
//
// Solidity: function proposeCommit(_knowledgeLabelIndex uint256) returns()
func (_DitContractInterface *DitContractInterfaceTransactor) ProposeCommit(opts *bind.TransactOpts, _knowledgeLabelIndex *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.contract.Transact(opts, "proposeCommit", _knowledgeLabelIndex)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x1c97d92b.
//
// Solidity: function proposeCommit(_knowledgeLabelIndex uint256) returns()
func (_DitContractInterface *DitContractInterfaceSession) ProposeCommit(_knowledgeLabelIndex *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.Contract.ProposeCommit(&_DitContractInterface.TransactOpts, _knowledgeLabelIndex)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x1c97d92b.
//
// Solidity: function proposeCommit(_knowledgeLabelIndex uint256) returns()
func (_DitContractInterface *DitContractInterfaceTransactorSession) ProposeCommit(_knowledgeLabelIndex *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.Contract.ProposeCommit(&_DitContractInterface.TransactOpts, _knowledgeLabelIndex)
}

// ResolveVote is a paid mutator transaction binding the contract method 0x30f4ccac.
//
// Solidity: function resolveVote(_proposalID uint256) returns()
func (_DitContractInterface *DitContractInterfaceTransactor) ResolveVote(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.contract.Transact(opts, "resolveVote", _proposalID)
}

// ResolveVote is a paid mutator transaction binding the contract method 0x30f4ccac.
//
// Solidity: function resolveVote(_proposalID uint256) returns()
func (_DitContractInterface *DitContractInterfaceSession) ResolveVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.Contract.ResolveVote(&_DitContractInterface.TransactOpts, _proposalID)
}

// ResolveVote is a paid mutator transaction binding the contract method 0x30f4ccac.
//
// Solidity: function resolveVote(_proposalID uint256) returns()
func (_DitContractInterface *DitContractInterfaceTransactorSession) ResolveVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.Contract.ResolveVote(&_DitContractInterface.TransactOpts, _proposalID)
}

// RevealVoteOnProposal is a paid mutator transaction binding the contract method 0x900f0d39.
//
// Solidity: function revealVoteOnProposal(_proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitContractInterface *DitContractInterfaceTransactor) RevealVoteOnProposal(opts *bind.TransactOpts, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.contract.Transact(opts, "revealVoteOnProposal", _proposalID, _voteOption, _voteSalt)
}

// RevealVoteOnProposal is a paid mutator transaction binding the contract method 0x900f0d39.
//
// Solidity: function revealVoteOnProposal(_proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitContractInterface *DitContractInterfaceSession) RevealVoteOnProposal(_proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.Contract.RevealVoteOnProposal(&_DitContractInterface.TransactOpts, _proposalID, _voteOption, _voteSalt)
}

// RevealVoteOnProposal is a paid mutator transaction binding the contract method 0x900f0d39.
//
// Solidity: function revealVoteOnProposal(_proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitContractInterface *DitContractInterfaceTransactorSession) RevealVoteOnProposal(_proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitContractInterface.Contract.RevealVoteOnProposal(&_DitContractInterface.TransactOpts, _proposalID, _voteOption, _voteSalt)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xbfac6efa.
//
// Solidity: function voteOnProposal(_proposalID uint256, _voteHash bytes32) returns()
func (_DitContractInterface *DitContractInterfaceTransactor) VoteOnProposal(opts *bind.TransactOpts, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitContractInterface.contract.Transact(opts, "voteOnProposal", _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xbfac6efa.
//
// Solidity: function voteOnProposal(_proposalID uint256, _voteHash bytes32) returns()
func (_DitContractInterface *DitContractInterfaceSession) VoteOnProposal(_proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitContractInterface.Contract.VoteOnProposal(&_DitContractInterface.TransactOpts, _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xbfac6efa.
//
// Solidity: function voteOnProposal(_proposalID uint256, _voteHash bytes32) returns()
func (_DitContractInterface *DitContractInterfaceTransactorSession) VoteOnProposal(_proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitContractInterface.Contract.VoteOnProposal(&_DitContractInterface.TransactOpts, _proposalID, _voteHash)
}

// DitContractInterfaceCommitProposalIterator is returned from FilterCommitProposal and is used to iterate over the raw logs and unpacked data for CommitProposal events raised by the DitContractInterface contract.
type DitContractInterfaceCommitProposalIterator struct {
	Event *DitContractInterfaceCommitProposal // Event containing the contract specifics and raw log

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
func (it *DitContractInterfaceCommitProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractInterfaceCommitProposal)
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
		it.Event = new(DitContractInterfaceCommitProposal)
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
func (it *DitContractInterfaceCommitProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractInterfaceCommitProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractInterfaceCommitProposal represents a CommitProposal event raised by the DitContractInterface contract.
type DitContractInterfaceCommitProposal struct {
	Proposal *big.Int
	Who      common.Address
	Label    string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitProposal is a free log retrieval operation binding the contract event 0xf2c3b57260094fc6dc300b8881004fb6c60df6103955d42f7f70bf2bbbc7344f.
//
// Solidity: e CommitProposal(proposal indexed uint256, who indexed address, label string)
func (_DitContractInterface *DitContractInterfaceFilterer) FilterCommitProposal(opts *bind.FilterOpts, proposal []*big.Int, who []common.Address) (*DitContractInterfaceCommitProposalIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContractInterface.contract.FilterLogs(opts, "CommitProposal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceCommitProposalIterator{contract: _DitContractInterface.contract, event: "CommitProposal", logs: logs, sub: sub}, nil
}

// WatchCommitProposal is a free log subscription operation binding the contract event 0xf2c3b57260094fc6dc300b8881004fb6c60df6103955d42f7f70bf2bbbc7344f.
//
// Solidity: e CommitProposal(proposal indexed uint256, who indexed address, label string)
func (_DitContractInterface *DitContractInterfaceFilterer) WatchCommitProposal(opts *bind.WatchOpts, sink chan<- *DitContractInterfaceCommitProposal, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContractInterface.contract.WatchLogs(opts, "CommitProposal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractInterfaceCommitProposal)
				if err := _DitContractInterface.contract.UnpackLog(event, "CommitProposal", log); err != nil {
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

// DitContractInterfaceProposalResolvedIterator is returned from FilterProposalResolved and is used to iterate over the raw logs and unpacked data for ProposalResolved events raised by the DitContractInterface contract.
type DitContractInterfaceProposalResolvedIterator struct {
	Event *DitContractInterfaceProposalResolved // Event containing the contract specifics and raw log

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
func (it *DitContractInterfaceProposalResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractInterfaceProposalResolved)
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
		it.Event = new(DitContractInterfaceProposalResolved)
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
func (it *DitContractInterfaceProposalResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractInterfaceProposalResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractInterfaceProposalResolved represents a ProposalResolved event raised by the DitContractInterface contract.
type DitContractInterfaceProposalResolved struct {
	Proposal *big.Int
	Label    string
	Accepted bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalResolved is a free log retrieval operation binding the contract event 0x201bd7ec9db47a5155ba8bd82a15a9b162113c1dbc4494fe017b1c9ad2d0d0ec.
//
// Solidity: e ProposalResolved(proposal indexed uint256, label string, accepted bool)
func (_DitContractInterface *DitContractInterfaceFilterer) FilterProposalResolved(opts *bind.FilterOpts, proposal []*big.Int) (*DitContractInterfaceProposalResolvedIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitContractInterface.contract.FilterLogs(opts, "ProposalResolved", proposalRule)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceProposalResolvedIterator{contract: _DitContractInterface.contract, event: "ProposalResolved", logs: logs, sub: sub}, nil
}

// WatchProposalResolved is a free log subscription operation binding the contract event 0x201bd7ec9db47a5155ba8bd82a15a9b162113c1dbc4494fe017b1c9ad2d0d0ec.
//
// Solidity: e ProposalResolved(proposal indexed uint256, label string, accepted bool)
func (_DitContractInterface *DitContractInterfaceFilterer) WatchProposalResolved(opts *bind.WatchOpts, sink chan<- *DitContractInterfaceProposalResolved, proposal []*big.Int) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitContractInterface.contract.WatchLogs(opts, "ProposalResolved", proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractInterfaceProposalResolved)
				if err := _DitContractInterface.contract.UnpackLog(event, "ProposalResolved", log); err != nil {
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

// DitContractInterfaceRevealIterator is returned from FilterReveal and is used to iterate over the raw logs and unpacked data for Reveal events raised by the DitContractInterface contract.
type DitContractInterfaceRevealIterator struct {
	Event *DitContractInterfaceReveal // Event containing the contract specifics and raw log

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
func (it *DitContractInterfaceRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractInterfaceReveal)
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
		it.Event = new(DitContractInterfaceReveal)
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
func (it *DitContractInterfaceRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractInterfaceRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractInterfaceReveal represents a Reveal event raised by the DitContractInterface contract.
type DitContractInterfaceReveal struct {
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Accept        bool
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReveal is a free log retrieval operation binding the contract event 0x0f93246f12982b82ea2f37862d166f15a04cd9aa4a60d80ea28299b2a0ff1c04.
//
// Solidity: e Reveal(proposal indexed uint256, who indexed address, label string, accept bool, numberOfVotes uint256)
func (_DitContractInterface *DitContractInterfaceFilterer) FilterReveal(opts *bind.FilterOpts, proposal []*big.Int, who []common.Address) (*DitContractInterfaceRevealIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContractInterface.contract.FilterLogs(opts, "Reveal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceRevealIterator{contract: _DitContractInterface.contract, event: "Reveal", logs: logs, sub: sub}, nil
}

// WatchReveal is a free log subscription operation binding the contract event 0x0f93246f12982b82ea2f37862d166f15a04cd9aa4a60d80ea28299b2a0ff1c04.
//
// Solidity: e Reveal(proposal indexed uint256, who indexed address, label string, accept bool, numberOfVotes uint256)
func (_DitContractInterface *DitContractInterfaceFilterer) WatchReveal(opts *bind.WatchOpts, sink chan<- *DitContractInterfaceReveal, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContractInterface.contract.WatchLogs(opts, "Reveal", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractInterfaceReveal)
				if err := _DitContractInterface.contract.UnpackLog(event, "Reveal", log); err != nil {
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

// DitContractInterfaceVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the DitContractInterface contract.
type DitContractInterfaceVoteIterator struct {
	Event *DitContractInterfaceVote // Event containing the contract specifics and raw log

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
func (it *DitContractInterfaceVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitContractInterfaceVote)
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
		it.Event = new(DitContractInterfaceVote)
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
func (it *DitContractInterfaceVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitContractInterfaceVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitContractInterfaceVote represents a Vote event raised by the DitContractInterface contract.
type DitContractInterfaceVote struct {
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Stake         *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xfe25e9ca71c8ef184b9997685aac106e44c6cbbf4986bf70085d7a67a5d9a103.
//
// Solidity: e Vote(proposal indexed uint256, who indexed address, label string, stake uint256, numberOfVotes uint256)
func (_DitContractInterface *DitContractInterfaceFilterer) FilterVote(opts *bind.FilterOpts, proposal []*big.Int, who []common.Address) (*DitContractInterfaceVoteIterator, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContractInterface.contract.FilterLogs(opts, "Vote", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitContractInterfaceVoteIterator{contract: _DitContractInterface.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xfe25e9ca71c8ef184b9997685aac106e44c6cbbf4986bf70085d7a67a5d9a103.
//
// Solidity: e Vote(proposal indexed uint256, who indexed address, label string, stake uint256, numberOfVotes uint256)
func (_DitContractInterface *DitContractInterfaceFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *DitContractInterfaceVote, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitContractInterface.contract.WatchLogs(opts, "Vote", proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitContractInterfaceVote)
				if err := _DitContractInterface.contract.UnpackLog(event, "Vote", log); err != nil {
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

// DitCoordinatorABI is the input ABI used to generate the binding from.
const DitCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"}],\"name\":\"getRepository\",\"outputs\":[{\"name\":\"repositoryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"},{\"name\":\"_label1\",\"type\":\"string\"},{\"name\":\"_label2\",\"type\":\"string\"},{\"name\":\"_label3\",\"type\":\"string\"},{\"name\":\"_voteSettings\",\"type\":\"uint256[3]\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"newDitContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DitCoordinatorBin is the compiled bytecode used for deploying new contracts.
const DitCoordinatorBin = `0x608060405234801561001057600080fd5b5060405160408061227a833981016040528051602090910151600160a060020a038116158015906100495750600160a060020a03821615155b15156100dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4b4e57566f74696e6720616e64204b4e57546f6b656e2061646472657373206360448201527f616e277420626520656d70747900000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a0319908116600160a060020a03938416178083556002805491851691831691909117905560018054821694841694909417938490556003805494909316931692909217905561213e90819061013c90396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166322dd1a1e8114610066578063922b1ca4146100a2578063985dbfc5146100e8578063eb49fe94146100fd575b600080fd5b34801561007257600080fd5b506100866004803560248101910135610112565b60408051600160a060020a039092168252519081900360200190f35b3480156100ae57600080fd5b50610086602460048035828101929082013591813580830192908201359160443580830192908201359160643591820191013560846101e2565b3480156100f457600080fd5b506100866105b1565b34801561010957600080fd5b506100866105c0565b60008082116101a857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4e616d65206f6620746865207265706f7369746f72792063616e27742062652060448201527f656d707479000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600483836040518083838082843790910194855250506040519283900360200190922060010154600160a060020a03169250505092915050565b600080891161027857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4e616d65206f6620746865207265706f7369746f72792063616e27742062652060448201527f656d707479000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000600160a060020a031660048b8b6040518083838082843790910194855250506040519283900360200190922060010154600160a060020a031692909214915061034c905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f5265706f7369746f72792063616e277420616c7265616479206861766520612060448201527f646974436f6e7472616374000000000000000000000000000000000000000000606482015290519081900360840190fd5b600054600160a060020a0316308b8b8b8b8b8b8b8b6103696105cf565b600160a060020a03808c1682528a16602082015260c060408201818152908201899052606082016080830160a0840160e085018d8d8082843790910186810385528b815260200190508b8b80828437909101868103845289815260200190508989808284379091018681038352878152602001905087878082843782019150509e505050505050505050505050505050604051809103906000f080158015610415573d6000803e3d6000fd5b506040805160c06020601f8e01819004028201810190925260a081018c81529293509182918d908d9081908501838280828437505050928452505050600160a060020a038316602080830191909152843560408084019190915290850135606083015284810135608090920191909152516004908c908c9080838380828437919091019485525050604051602093819003840190208451805191946104bf945085935001906105df565b5060208281015160018301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392831617905560408085015160028086019190915560608601516003860155608090950151600494850155935484517fb6818a520000000000000000000000000000000000000000000000000000000081528683169481019490945286356024850152918601356044840152858401356064840152925192169163b6818a529160848082019260009290919082900301818387803b15801561058c57600080fd5b505af11580156105a0573d6000803e3d6000fd5b505050509998505050505050505050565b600154600160a060020a031681565b600054600160a060020a031681565b604051611a988061067b83390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061062057805160ff191683800117855561064d565b8280016001018555821561064d579182015b8281111561064d578251825591602001919060010190610632565b5061065992915061065d565b5090565b61067791905b808211156106595760008155600101610663565b90560060806040523480156200001157600080fd5b5060405162001a9838038062001a9883398101604090815281516020830151918301516060840151608085015160a0860151939592830193918301929081019101600160a060020a03861615801590620000735750600160a060020a03851615155b15156200010757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f4b4e57566f74696e6720616e6420646974436f6f7264696e61746f722061646460448201527f726573732063616e277420626520656d70747900000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03808916600160a060020a0319928316179283905560018054898316908416179055600280549092169216919091179055835162000158906003906020870190620002ca565b506000600581905583511115620001a157600480546001810180835560009290925284516200019e9160008051602062001a7883398151915201906020870190620002ca565b50505b600082511115620001e45760048054600181018083556000929092528351620001e19160008051602062001a7883398151915201906020860190620002ca565b50505b600082511115620002275760048054600181018083556000929092528251620002249160008051602062001a7883398151915201906020850190620002ca565b50505b600454600010620002be57604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f50726f76696465206174206c65617374206f6e65204b6e6f776c65646765204c60448201527f6162656c00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5050505050506200036f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030d57805160ff19168380011785556200033d565b828001600101855582156200033d579182015b828111156200033d57825182559160200191906001019062000320565b506200034b9291506200034f565b5090565b6200036c91905b808211156200034b576000815560010162000356565b90565b6116f9806200037f6000396000f3006080604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663013cf08b81146100d4578063067462d31461019f5780631c97d92b146101cb5780631f8ff8e6146101d857806330f4ccac14610265578063639661901461027d578063900f0d39146102a457806393b17b8e146102c2578063bfac6efa146102f3578063c030638714610301578063c4e4b02a14610301578063dce6636c14610316578063e9176c601461032e578063eb49fe9414610343575b600080fd5b3480156100e057600080fd5b506100ec600435610358565b60408051888152600160a060020a038716918101919091528415156060820152831515608082015260a0810183905260c0810182905260e0602080830182815289519284019290925288516101008401918a019080838360005b8381101561015e578181015183820152602001610146565b50505050905090810190601f16801561018b5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156101ab57600080fd5b506101b7600435610445565b604080519115158252519081900360200190f35b6101d66004356104e9565b005b3480156101e457600080fd5b506101f0600435610a44565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022a578181015183820152602001610212565b50505050905090810190601f1680156102575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027157600080fd5b506101d6600435610aeb565b34801561028957600080fd5b50610292610fad565b60408051918252519081900360200190f35b3480156102b057600080fd5b506101d6600435602435604435610fb3565b3480156102ce57600080fd5b506102d761114a565b60408051600160a060020a039092168252519081900360200190f35b6101d6600435602435611159565b34801561030d57600080fd5b5061029261149b565b34801561032257600080fd5b506102926004356114a1565b34801561033a57600080fd5b506101f06114b6565b34801561034f57600080fd5b506102d7611511565b6006602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f8101869004860283018601909652858252919492939092908301828280156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b505050600284015460038501546004909501549394600160a060020a0382169460ff60a060020a840481169550750100000000000000000000000000000000000000000090930490921692509087565b60008181526006602052604081206002015460a060020a900460ff1615156104b7576040805160e560020a62461bcd02815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b506000908152600660205260409020600201547501000000000000000000000000000000000000000000900460ff1690565b60003411610567576040805160e560020a62461bcd02815260206004820152602860248201527f56616c7565206f6620746865207472616e73616374696f6e2063616e206e6f7460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600454600019018111156105ea576040805160e560020a62461bcd028152602060048201526024808201527f4b6e6f776c656467652d4c6162656c20696e646578206973206e6f7420636f7260448201527f7265637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6005546105fe90600163ffffffff61152016565b6005556040805160e08101909152600254600480548392600160a060020a03169163a663fce8913391908790811061063257fe5b9060005260206000200161012c80346040518663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018086600160a060020a0316600160a060020a031681526020018060200185815260200184815260200183815260200182810382528681815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561071f5780601f106106f45761010080835404028352916020019161071f565b820191906000526020600020905b81548152906001019060200180831161070257829003601f168201915b50509650505050505050602060405180830381600087803b15801561074357600080fd5b505af1158015610757573d6000803e3d6000fd5b505050506040513d602081101561076d57600080fd5b50518152600480546020909201918490811061078557fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156108135780601f106107e857610100808354040283529160200191610813565b820191906000526020600020905b8154815290600101906020018083116107f657829003601f168201915b505050918352505033602080830191909152600060408084018290526060840182905234608085015260a09093018190526005548152600682529190912082518155828201518051919261086f92600185019290910190611632565b5060408281015160028301805460608601516080870151151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001991151560a060020a0274ff000000000000000000000000000000000000000019600160a060020a0390961673ffffffffffffffffffffffffffffffffffffffff199094169390931794909416919091171691909117905560a0830151600383015560c0909201516004918201556005805460009081526006602081815285832033845284018152858320349081905593548352529290922001546109619163ffffffff61152016565b60058054600090815260066020526040902060049081019290925554815433927ff2c3b57260094fc6dc300b8881004fb6c60df6103955d42f7f70bf2bbbc7344f91859081106109ad57fe5b6000918252602091829020604080518481529290910180546002600182161561010002600019019091160493830184905292829182019084908015610a335780601f10610a0857610100808354040283529160200191610a33565b820191906000526020600020905b815481529060010190602001808311610a1657829003601f168201915b50509250505060405180910390a350565b6004805482908110610a5257fe5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815293509091830182828015610ae35780601f10610ab857610100808354040283529160200191610ae3565b820191906000526020600020905b815481529060010190602001808311610ac657829003601f168201915b505050505081565b60008181526006602052604081206002015460a060020a900460ff161515610ce55760025460008381526006602090815260408083205481517fe74fef3700000000000000000000000000000000000000000000000000000000815260048101919091529051600160a060020a039094169363e74fef3793602480840194938390030190829087803b158015610b8057600080fd5b505af1158015610b94573d6000803e3d6000fd5b505050506040513d6020811015610baa57600080fd5b50516000838152600660209081526040918290206002808201805460a060020a75ff00000000000000000000000000000000000000000019909116750100000000000000000000000000000000000000000097151588021774ff000000000000000000000000000000000000000019161790819055845160ff969091049590951680151593860193909352838552600191820180546101009381161593909302600019019092160492840183905285937f201bd7ec9db47a5155ba8bd82a15a9b162113c1dbc4494fe017b1c9ad2d0d0ec939192918190606082019085908015610cd55780601f10610caa57610100808354040283529160200191610cd5565b820191906000526020600020905b815481529060010190602001808311610cb857829003601f168201915b5050935050505060405180910390a25b60008281526006602090815260408083203384526005019091528120600101541180610d2a5750600082815260066020526040902060020154600160a060020a031633145b1515610da6576040805160e560020a62461bcd02815260206004820152603a60248201527f4f6e6c79207061727469636970616e7473206f662074686520766f746520617260448201527f652061626c6520746f207265736f6c76652074686520766f7465000000000000606482015290519081900360840190fd5b600082815260066020908152604080832033845260050190915281205411610e64576040805160e560020a62461bcd02815260206004820152604960248201527f4f6e6c79207061727469636970616e74732077686f20686176656e277420616c60448201527f7265616479207265736f6c7665642074686520766f7465206172652061626c6560648201527f20746f20646f20736f0000000000000000000000000000000000000000000000608482015290519081900360a40190fd5b600280546000848152600660209081526040808320805433808652600590920184528285209096015482517fce729fd200000000000000000000000000000000000000000000000000000000815260048101979097526024870152604486015251600160a060020a039093169363ce729fd29360648083019491928390030190829087803b158015610ef557600080fd5b505af1158015610f09573d6000803e3d6000fd5b505050506040513d6020811015610f1f57600080fd5b505190506000811115610f5b57604051339082156108fc029083906000818181858888f19350505050158015610f59573d6000803e3d6000fd5b505b6000828152600660208181526040808420338552600581018352908420849055928590525260040154610f94908263ffffffff6115aa16565b6000928352600660205260409092206004019190915550565b60055481565b6002546000848152600660205260408082205481517f34f2f2d2000000000000000000000000000000000000000000000000000000008152600481019190915233602482015260448101869052606481018590529051600160a060020a03909316926334f2f2d29260848084019391929182900301818387803b15801561103957600080fd5b505af115801561104d573d6000803e3d6000fd5b5050506000848152600660209081526040808320338085526005820184529382902060028082018990556001918201548451838b14968101879052948501819052606080865293830180549384161561010002600019019093169190910492840183905294955088947f0f93246f12982b82ea2f37862d166f15a04cd9aa4a60d80ea28299b2a0ff1c04949193919290919081906080820190869080156111355780601f1061110a57610100808354040283529160200191611135565b820191906000526020600020905b81548152906001019060200180831161111857829003601f168201915b505094505050505060405180910390a3505050565b600154600160a060020a031681565b60008281526006602052604081206003015434146111e7576040805160e560020a62461bcd02815260206004820152603960248201527f56616c7565206f6620746865207472616e73616374696f6e20646f65736e277460448201527f206d6174636820746865207265717569726564207374616b6500000000000000606482015290519081900360840190fd5b600083815260066020526040902060020154600160a060020a031633141561127f576040805160e560020a62461bcd02815260206004820152603160248201527f5468652070726f706f736572206973206e6f7420616c6c6f77656420746f207660448201527f6f746520696e20612070726f706f73616c000000000000000000000000000000606482015290519081900360840190fd5b6000838152600660205260409020600401546112a1903463ffffffff61152016565b6000848152600660208181526040808420600480820196909655338086526005820184528286203490556002548a8752948452905482517f7eb2ff520000000000000000000000000000000000000000000000000000000081529687015260248601526044850187905251600160a060020a0390921693637eb2ff5293606480830194928390030190829087803b15801561133b57600080fd5b505af115801561134f573d6000803e3d6000fd5b505050506040513d602081101561136557600080fd5b50519050600081116113e7576040805160e560020a62461bcd02815260206004820152603360248201527f566f74696e6720636f6e74726163742072657475726e656420616e20696e766160448201527f6c696420616d6f756e74206f6620766f74657300000000000000000000000000606482015290519081900360840190fd5b6000838152600660208181526040808420338086526005820184528286206001908101889055958990529383528151349381018490529182018690526060808352908501805460029681161561010002600019011695909504908201819052929387937ffe25e9ca71c8ef184b9997685aac106e44c6cbbf4986bf70085d7a67a5d9a103939192879181906080820190869080156111355780601f1061110a57610100808354040283529160200191611135565b61012c81565b60009081526006602052604090206003015490565b6003805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610ae35780601f10610ab857610100808354040283529160200191610ae3565b600054600160a060020a031681565b6000828201838110156115a3576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b6000808383111561162b576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061167357805160ff19168380011785556116a0565b828001600101855582156116a0579182015b828111156116a0578251825591602001919060010190611685565b506116ac9291506116b0565b5090565b6116ca91905b808211156116ac57600081556001016116b6565b905600a165627a7a72305820c3aacf13b3fd17948640e8e23c9776d002bc09b3ef5905ff72011c48b50e6d5700298a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19ba165627a7a72305820c62586b9b9e9a8df94c0e4ed425963fd39296f710a4f489beb379538938c383f0029`

// DeployDitCoordinator deploys a new Ethereum contract, binding an instance of DitCoordinator to it.
func DeployDitCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWTokenAddress common.Address, _KNWVotingAddress common.Address) (common.Address, *types.Transaction, *DitCoordinator, error) {
	parsed, err := abi.JSON(strings.NewReader(DitCoordinatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DitCoordinatorBin), backend, _KNWTokenAddress, _KNWVotingAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DitCoordinator{DitCoordinatorCaller: DitCoordinatorCaller{contract: contract}, DitCoordinatorTransactor: DitCoordinatorTransactor{contract: contract}, DitCoordinatorFilterer: DitCoordinatorFilterer{contract: contract}}, nil
}

// DitCoordinator is an auto generated Go binding around an Ethereum contract.
type DitCoordinator struct {
	DitCoordinatorCaller     // Read-only binding to the contract
	DitCoordinatorTransactor // Write-only binding to the contract
	DitCoordinatorFilterer   // Log filterer for contract events
}

// DitCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DitCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DitCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DitCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DitCoordinatorSession struct {
	Contract     *DitCoordinator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DitCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DitCoordinatorCallerSession struct {
	Contract *DitCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DitCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DitCoordinatorTransactorSession struct {
	Contract     *DitCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DitCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DitCoordinatorRaw struct {
	Contract *DitCoordinator // Generic contract binding to access the raw methods on
}

// DitCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DitCoordinatorCallerRaw struct {
	Contract *DitCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// DitCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DitCoordinatorTransactorRaw struct {
	Contract *DitCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDitCoordinator creates a new instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinator(address common.Address, backend bind.ContractBackend) (*DitCoordinator, error) {
	contract, err := bindDitCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DitCoordinator{DitCoordinatorCaller: DitCoordinatorCaller{contract: contract}, DitCoordinatorTransactor: DitCoordinatorTransactor{contract: contract}, DitCoordinatorFilterer: DitCoordinatorFilterer{contract: contract}}, nil
}

// NewDitCoordinatorCaller creates a new read-only instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*DitCoordinatorCaller, error) {
	contract, err := bindDitCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorCaller{contract: contract}, nil
}

// NewDitCoordinatorTransactor creates a new write-only instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DitCoordinatorTransactor, error) {
	contract, err := bindDitCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorTransactor{contract: contract}, nil
}

// NewDitCoordinatorFilterer creates a new log filterer instance of DitCoordinator, bound to a specific deployed contract.
func NewDitCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DitCoordinatorFilterer, error) {
	contract, err := bindDitCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorFilterer{contract: contract}, nil
}

// bindDitCoordinator binds a generic wrapper to an already deployed contract.
func bindDitCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DitCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitCoordinator *DitCoordinatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitCoordinator.Contract.DitCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitCoordinator *DitCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitCoordinator.Contract.DitCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitCoordinator *DitCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitCoordinator.Contract.DitCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitCoordinator *DitCoordinatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitCoordinator *DitCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitCoordinator *DitCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitCoordinator.Contract.contract.Transact(opts, method, params...)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCaller) KNWTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "KNWTokenAddress")
	return *ret0, err
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorSession) KNWTokenAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWTokenAddress(&_DitCoordinator.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCallerSession) KNWTokenAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWTokenAddress(&_DitCoordinator.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCaller) KNWVotingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "KNWVotingAddress")
	return *ret0, err
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorSession) KNWVotingAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWVotingAddress(&_DitCoordinator.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitCoordinator *DitCoordinatorCallerSession) KNWVotingAddress() (common.Address, error) {
	return _DitCoordinator.Contract.KNWVotingAddress(&_DitCoordinator.CallOpts)
}

// GetRepository is a free data retrieval call binding the contract method 0x22dd1a1e.
//
// Solidity: function getRepository(_repository string) constant returns(repositoryAddress address)
func (_DitCoordinator *DitCoordinatorCaller) GetRepository(opts *bind.CallOpts, _repository string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getRepository", _repository)
	return *ret0, err
}

// GetRepository is a free data retrieval call binding the contract method 0x22dd1a1e.
//
// Solidity: function getRepository(_repository string) constant returns(repositoryAddress address)
func (_DitCoordinator *DitCoordinatorSession) GetRepository(_repository string) (common.Address, error) {
	return _DitCoordinator.Contract.GetRepository(&_DitCoordinator.CallOpts, _repository)
}

// GetRepository is a free data retrieval call binding the contract method 0x22dd1a1e.
//
// Solidity: function getRepository(_repository string) constant returns(repositoryAddress address)
func (_DitCoordinator *DitCoordinatorCallerSession) GetRepository(_repository string) (common.Address, error) {
	return _DitCoordinator.Contract.GetRepository(&_DitCoordinator.CallOpts, _repository)
}

// InitRepository is a paid mutator transaction binding the contract method 0x922b1ca4.
//
// Solidity: function initRepository(_repository string, _label1 string, _label2 string, _label3 string, _voteSettings uint256[3]) returns(newDitContract address)
func (_DitCoordinator *DitCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository string, _label1 string, _label2 string, _label3 string, _voteSettings [3]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "initRepository", _repository, _label1, _label2, _label3, _voteSettings)
}

// InitRepository is a paid mutator transaction binding the contract method 0x922b1ca4.
//
// Solidity: function initRepository(_repository string, _label1 string, _label2 string, _label3 string, _voteSettings uint256[3]) returns(newDitContract address)
func (_DitCoordinator *DitCoordinatorSession) InitRepository(_repository string, _label1 string, _label2 string, _label3 string, _voteSettings [3]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _voteSettings)
}

// InitRepository is a paid mutator transaction binding the contract method 0x922b1ca4.
//
// Solidity: function initRepository(_repository string, _label1 string, _label2 string, _label3 string, _voteSettings uint256[3]) returns(newDitContract address)
func (_DitCoordinator *DitCoordinatorTransactorSession) InitRepository(_repository string, _label1 string, _label2 string, _label3 string, _voteSettings [3]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _voteSettings)
}
