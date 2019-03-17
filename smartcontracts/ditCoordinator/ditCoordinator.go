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
const KNWVotingContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newKNWTokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"numVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeLabel\",\"type\":\"string\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRepository\",\"type\":\"bytes32\"},{\"name\":\"_majority\",\"type\":\"uint256\"},{\"name\":\"_mintingMethod\",\"type\":\"uint256\"},{\"name\":\"_burningMethod\",\"type\":\"uint256\"}],\"name\":\"addNewRepository\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"resolveVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"resolvePoll\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCoordinatorAddress\",\"type\":\"address\"}],\"name\":\"setCoordinatorAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AddNewRepository is a paid mutator transaction binding the contract method 0xa3fba060.
//
// Solidity: function addNewRepository(_newRepository bytes32, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVotingContract *KNWVotingContractTransactor) AddNewRepository(opts *bind.TransactOpts, _newRepository [32]byte, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "addNewRepository", _newRepository, _majority, _mintingMethod, _burningMethod)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0xa3fba060.
//
// Solidity: function addNewRepository(_newRepository bytes32, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVotingContract *KNWVotingContractSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddNewRepository(&_KNWVotingContract.TransactOpts, _newRepository, _majority, _mintingMethod, _burningMethod)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0xa3fba060.
//
// Solidity: function addNewRepository(_newRepository bytes32, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVotingContract *KNWVotingContractTransactorSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddNewRepository(&_KNWVotingContract.TransactOpts, _newRepository, _majority, _mintingMethod, _burningMethod)
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

// StartPoll is a paid mutator transaction binding the contract method 0x9156cd07.
//
// Solidity: function startPoll(_repository bytes32, _address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVotingContract *KNWVotingContractTransactor) StartPoll(opts *bind.TransactOpts, _repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "startPoll", _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0x9156cd07.
//
// Solidity: function startPoll(_repository bytes32, _address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVotingContract *KNWVotingContractSession) StartPoll(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartPoll(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0x9156cd07.
//
// Solidity: function startPoll(_repository bytes32, _address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVotingContract *KNWVotingContractTransactorSession) StartPoll(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartPoll(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"sqrt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x61016b610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100575763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663677342ce811461005c575b600080fd5b610067600435610079565b60408051918252519081900360200190f35b6000808083151561008d5760009250610138565b6001840184106100fe57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f466c6177656420696e70757420666f7220737172740000000000000000000000604482015290519081900360640190fd5b505060026001830104825b80821015610134575080600281808681151561012157fe5b040181151561012c57fe5b049150610109565b8192505b50509190505600a165627a7a72305820ce9eba5c2c770e6489e8634ca0a14db23dee696aa56e1247aac6317c7eab01bd0029`

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

// DitCoordinatorABI is the input ABI used to generate the binding from.
const DitCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getKNWVoteIDFromProposalID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelIndex\",\"type\":\"uint256\"},{\"name\":\"_voteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"_voteOpenDuration\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getCurrentProposalID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"openVoteOnProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repositories\",\"outputs\":[{\"name\":\"votingMajority\",\"type\":\"uint256\"},{\"name\":\"mintingMethod\",\"type\":\"uint256\"},{\"name\":\"burningMethod\",\"type\":\"uint256\"},{\"name\":\"currentProposalID\",\"type\":\"uint256\"},{\"name\":\"minVoteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"maxVoteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"minVoteOpenDuration\",\"type\":\"uint256\"},{\"name\":\"maxVoteOpenDuration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"finalizeVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsOfRepository\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"individualStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_label1\",\"type\":\"string\"},{\"name\":\"_label2\",\"type\":\"string\"},{\"name\":\"_label3\",\"type\":\"string\"},{\"name\":\"_voteSettings\",\"type\":\"uint256[7]\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getIndividualStake\",\"outputs\":[{\"name\":\"individualStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelID\",\"type\":\"uint256\"}],\"name\":\"getKnowledgeLabels\",\"outputs\":[{\"name\":\"knowledgeLabel\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"voteOnProposal\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"repositoryIsInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"ProposeCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"CommitVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"OpenVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"FinalizeVote\",\"type\":\"event\"}]"

// DitCoordinatorBin is the compiled bytecode used for deploying new contracts.
const DitCoordinatorBin = `0x608060405234801561001057600080fd5b50604051604080611f05833981016040528051602090910151600160a060020a038116158015906100495750600160a060020a03821615155b15156100dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4b4e57566f74696e6720616e64204b4e57546f6b656e2061646472657373206360448201527f616e277420626520656d70747900000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a0319908116600160a060020a039384161780835560028054918516918316919091179055600180548216948416949094179384905560038054949093169316929092179055611dc990819061013c90396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306ee459681146100df5780630aba86881461010c5780630bdc90e8146101225780630ee62ec01461013a5780631f51fd711461015b5780632e71d0fb146101b45780633e029f63146101cf57806351f43c241461029d578063552edc9d146102ef57806387c9203d1461030a5780639533222914610325578063985dbfc5146103b5578063a34c299a146103e6578063ea6c6d0f146103f7578063eb49fe941461040f575b600080fd5b3480156100eb57600080fd5b506100fa600435602435610424565b60408051918252519081900360200190f35b610120600435602435604435606435610441565b005b34801561012e57600080fd5b506100fa600435610b02565b34801561014657600080fd5b50610120600435602435604435606435610b17565b34801561016757600080fd5b50610173600435610cbf565b604080519889526020890197909752878701959095526060870193909352608086019190915260a085015260c084015260e083015251908190036101000190f35b3480156101c057600080fd5b50610120600435602435610d07565b3480156101db57600080fd5b506101ea6004356024356111f6565b60408051888152600160a060020a038716918101919091528415156060820152831515608082015260a0810183905260c0810182905260e0602080830182815289519284019290925288516101008401918a019080838360005b8381101561025c578181015183820152602001610244565b50505050905090810190601f1680156102895780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156102a957600080fd5b506102db60048035906024803580820192908101359160443580820192908101359160643590810191013560846112ea565b604080519115158252519081900360200190f35b3480156102fb57600080fd5b506100fa60043560243561161c565b34801561031657600080fd5b506102db60043560243561163c565b34801561033157600080fd5b506103406004356024356116f6565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561037a578181015183820152602001610362565b50505050905090810190601f1680156103a75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c157600080fd5b506103ca6117a4565b60408051600160a060020a039092168252519081900360200190f35b6101206004356024356044356117b3565b34801561040357600080fd5b506102db600435611b15565b34801561041b57600080fd5b506103ca611b2b565b600091825260056020908152604080842092845291905290205490565b600034116104bf576040805160e560020a62461bcd02815260206004820152602860248201527f56616c7565206f6620746865207472616e73616374696f6e2063616e206e6f7460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600084815260046020526040812084600381106104d857fe5b01546002600019610100600184161502019091160411610567576040805160e560020a62461bcd028152602060048201526024808201527f4b6e6f776c656467652d4c6162656c20696e646578206973206e6f7420636f7260448201527f7265637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600084815260046020526040902060070154821080159061059957506000848152600460205260409020600801548211155b15156105ef576040805160e560020a62461bcd02815260206004820152601c60248201527f566f746520636f6d6d6974206475726174696f6e20696e76616c696400000000604482015290519081900360640190fd5b600084815260046020526040902060090154811080159061062157506000848152600460205260409020600a01548111155b1515610677576040805160e560020a62461bcd02815260206004820152601a60248201527f566f7465206f70656e206475726174696f6e20696e76616c6964000000000000604482015290519081900360640190fd5b60008481526004602052604090206006015461069a90600163ffffffff611b3a16565b600085815260046020526040908190206006810192909255805160e0810190915260025490918291600160a060020a031690639156cd07908890339089600381106106e157fe5b6040517c010000000000000000000000000000000000000000000000000000000063ffffffff871602815260048101858152600160a060020a0385166024830152606482018c9052608482018b90523460a4830181905260c0604484019081529490930180546002610100600183161502600019019091160460c4840181905290948d948d9490939260e490910190879080156107bf5780601f10610794576101008083540402835291602001916107bf565b820191906000526020600020905b8154815290600101906020018083116107a257829003601f168201915b5050975050505050505050602060405180830381600087803b1580156107e457600080fd5b505af11580156107f8573d6000803e3d6000fd5b505050506040513d602081101561080e57600080fd5b505181526000868152600460209081526040909120910190856003811061083157fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108b65780601f1061088b576101008083540402835291602001916108b6565b820191906000526020600020905b81548152906001019060200180831161089957829003601f168201915b505050918352505033602080830191909152600060408084018290526060840182905234608085015260a0909301819052878152600582528281206004835283822060060154825282529190912082518155828201518051919261092292600185019290910190611c4c565b5060408281015160028301805460608601516080870151151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001991151560a060020a0274ff000000000000000000000000000000000000000019600160a060020a0390961673ffffffffffffffffffffffffffffffffffffffff199094169390931794909416919091171691909117905560a0830151600383015560c0909201516004918201556000868152600560209081528382208382528483206006015483529052919091200154610a0b903463ffffffff611b3a16565b6000858152600560209081526040808320600480845282852060068101805487529285529285208101959095559288905292905254339186907f171fe77c3addce776991159eb3eb73b14d9187ebd06c1c34ea12355a84ddbd83908760038110610a7157fe5b6040805160208082529390920180546002600019610100600184161502019091160493830184905292829182019084908015610aee5780601f10610ac357610100808354040283529160200191610aee565b820191906000526020600020905b815481529060010190602001808311610ad157829003601f168201915b50509250505060405180910390a450505050565b60009081526004602052604090206006015490565b60025460008581526005602090815260408083208784529091528082205481517f34f2f2d2000000000000000000000000000000000000000000000000000000008152600481019190915233602482015260448101869052606481018590529051600160a060020a03909316926334f2f2d29260848084019391929182900301818387803b158015610ba857600080fd5b505af1158015610bbc573d6000803e3d6000fd5b5050506000858152600560208181526040808420888552825280842033808652938101835293819020600180820189905590548251898314948101859052928301819052606080845295820180546002600019948216156101000294909401169290920495830186905293955088948a947f864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab94929390918190608082019086908015610ca95780601f10610c7e57610100808354040283529160200191610ca9565b820191906000526020600020905b815481529060010190602001808311610c8c57829003601f168201915b505094505050505060405180910390a450505050565b600460205280600052604060002060009150905080600301549080600401549080600501549080600601549080600701549080600801549080600901549080600a0154905088565b60008281526005602081815260408084208585528252808420338552909201905281206002015460ff1615610dac576040805160e560020a62461bcd02815260206004820152602760248201527f45616368207061727469636970616e742063616e206f6e6c792066696e616c6960448201527f7a65206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000838152600560208181526040808420868552825280842033855290920190528120541180610e0057506000838152600560209081526040808320858452909152902060020154600160a060020a031633145b1515610e7c576040805160e560020a62461bcd02815260206004820152603a60248201527f4f6e6c79207061727469636970616e7473206f662074686520766f746520617260448201527f652061626c6520746f207265736f6c76652074686520766f7465000000000000606482015290519081900360840190fd5b600083815260056020908152604080832085845290915290206002015460a060020a900460ff16151561109257600254600084815260056020908152604080832086845282528083205481517fe74fef3700000000000000000000000000000000000000000000000000000000815260048101919091529051600160a060020a039094169363e74fef3793602480840194938390030190829087803b158015610f2457600080fd5b505af1158015610f38573d6000803e3d6000fd5b505050506040513d6020811015610f4e57600080fd5b505160008481526005602090815260408083208684528252918290206002808201805474ff000000000000000000000000000000000000000019961515750100000000000000000000000000000000000000000090810275ff00000000000000000000000000000000000000000019909216919091179690961660a060020a1790819055845195900460ff16801515938601939093528385526001918201805492831615610100026000190190921604928401839052859387937f6bd2699645e0f6c5547bdf0d053280e48fef1ab21514bd02c88610b1279b942a93919081906060820190859080156110825780601f1061105757610100808354040283529160200191611082565b820191906000526020600020905b81548152906001019060200180831161106557829003601f168201915b5050935050505060405180910390a35b600254600084815260056020818152604080842087855282528084208054338087529190940183528185206001015482517fce729fd200000000000000000000000000000000000000000000000000000000815260048101959095526024850152604484015251600160a060020a039094169363ce729fd2936064808501948390030190829087803b15801561112757600080fd5b505af115801561113b573d6000803e3d6000fd5b505050506040513d602081101561115157600080fd5b50519050600081111561118d57604051339082156108fc029083906000818181858888f1935050505015801561118b573d6000803e3d6000fd5b505b60008381526005602090815260408083208584529091529020600401546111ba908263ffffffff611bc416565b6000938452600560208181526040808720958752948152848620600481019390935533865291019052509020600201805460ff19166001179055565b60056020908152600092835260408084208252918352918190208054600180830180548551600293821615610100026000190190911692909204601f810187900487028301870190955284825291949293909283018282801561129a5780601f1061126f5761010080835404028352916020019161129a565b820191906000526020600020905b81548152906001019060200180831161127d57829003601f168201915b505050600284015460038501546004909501549394600160a060020a0382169460ff60a060020a840481169550750100000000000000000000000000000000000000000090930490921692509087565b6000881515611369576040805160e560020a62461bcd02815260206004820152602360248201527f5265706f7369746f72792064657363726970746f722063616e2774206265207a60448201527f65726f0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600089815260046020526040902060030154156113f6576040805160e560020a62461bcd02815260206004820152602760248201527f5265706f7369746f72792063616e206f6e6c7920626520696e697469616c697a60448201527f6564206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008711806114055750600085115b806114105750600083115b151561148b576040805160e560020a62461bcd028152602060048201526024808201527f50726f76696465206174206c65617374206f6e65204b6e6f776c65646765204c60448201527f6162656c00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b604080516101a06020601f8b0181900402820181019092526101808101898152909182916101208301918291908d908d9081908701838280828437820191505050505050815260200189898080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437505050929093525050508152833560208083019190915284013560408083019190915284013560608201526000608082015260a00183600360209081029190910135825260808501358282015260a085013560408084019190915260c086013560609093019290925260008c815260049091522081516115b99082906003611cca565b506020820151600382015560408201516004820155606082015160058201556080820151600682015560a0820151600782015560c0820151600882015560e0820151600982015561010090910151600a9091015550600198975050505050505050565b600091825260056020908152604080842092845291905290206003015490565b600082815260056020908152604080832084845290915281206002015460a060020a900460ff1615156116b9576040805160e560020a62461bcd02815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b5060009182526005602090815260408084209284529190529020600201547501000000000000000000000000000000000000000000900460ff1690565b6000828152600460205260409020606090826003811061171257fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156117975780601f1061176c57610100808354040283529160200191611797565b820191906000526020600020905b81548152906001019060200180831161177a57829003601f168201915b5050505050905092915050565b600154600160a060020a031681565b6000838152600560209081526040808320858452909152812060030154341461184c576040805160e560020a62461bcd02815260206004820152603960248201527f56616c7565206f6620746865207472616e73616374696f6e20646f65736e277460448201527f206d6174636820746865207265717569726564207374616b6500000000000000606482015290519081900360840190fd5b6000848152600560209081526040808320868452909152902060020154600160a060020a03163314156118ef576040805160e560020a62461bcd02815260206004820152603160248201527f5468652070726f706f736572206973206e6f7420616c6c6f77656420746f207660448201527f6f746520696e20612070726f706f73616c000000000000000000000000000000606482015290519081900360840190fd5b600084815260056020908152604080832086845290915290206004015461191c903463ffffffff611b3a16565b60008581526005602090815260408083208784528252808320600480820195909555600254905482517f7eb2ff5200000000000000000000000000000000000000000000000000000000815295860152336024860152604485018790529051600160a060020a0390911693637eb2ff529360648083019493928390030190829087803b1580156119ab57600080fd5b505af11580156119bf573d6000803e3d6000fd5b505050506040513d60208110156119d557600080fd5b5051905060008111611a57576040805160e560020a62461bcd02815260206004820152603360248201527f566f74696e6720636f6e74726163742072657475726e656420616e20696e766160448201527f6c696420616d6f756e74206f6620766f74657300000000000000000000000000606482015290519081900360840190fd5b6000848152600560208181526040808420878552808352818520338087529481018452828620879055948890528252805134928101839052908101859052606080825260019485018054600261010097821615979097026000190116959095049082018190529293879389937fa01eea487bb3ec75528c167ccf90452d4164ddda7b13c55b2a89751a8dc5fbc19390918891908190608082019086908015610ca95780601f10610c7e57610100808354040283529160200191610ca9565b6000908152600460205260408120600301541190565b600054600160a060020a031681565b600082820183811015611bbd576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b60008083831115611c45576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611c8d57805160ff1916838001178555611cba565b82800160010185558215611cba579182015b82811115611cba578251825591602001919060010190611c9f565b50611cc6929150611d16565b5090565b8260038101928215611d0a579160200282015b82811115611d0a5782518051611cfa918491602090910190611c4c565b5091602001919060010190611cdd565b50611cc6929150611d33565b611d3091905b80821115611cc65760008155600101611d1c565b90565b611d3091905b80821115611cc6576000611d4d8282611d56565b50600101611d39565b50805460018160011615610100020316600290046000825580601f10611d7c5750611d9a565b601f016020900490600052602060002090810190611d9a9190611d16565b505600a165627a7a72305820e9c1101f47bce68951ee43e14c20769457864427337fc75be0ce242ca649aa0b0029`

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

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(_repository bytes32) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) GetCurrentProposalID(opts *bind.CallOpts, _repository [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getCurrentProposalID", _repository)
	return *ret0, err
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(_repository bytes32) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetCurrentProposalID(&_DitCoordinator.CallOpts, _repository)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(_repository bytes32) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetCurrentProposalID(&_DitCoordinator.CallOpts, _repository)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(_repository bytes32, _proposalID uint256) constant returns(individualStake uint256)
func (_DitCoordinator *DitCoordinatorCaller) GetIndividualStake(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getIndividualStake", _repository, _proposalID)
	return *ret0, err
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(_repository bytes32, _proposalID uint256) constant returns(individualStake uint256)
func (_DitCoordinator *DitCoordinatorSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetIndividualStake(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(_repository bytes32, _proposalID uint256) constant returns(individualStake uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetIndividualStake(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(_repository bytes32, _proposalID uint256) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) GetKNWVoteIDFromProposalID(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getKNWVoteIDFromProposalID", _repository, _proposalID)
	return *ret0, err
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(_repository bytes32, _proposalID uint256) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(_repository bytes32, _proposalID uint256) constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(_repository bytes32, _knowledgeLabelID uint256) constant returns(knowledgeLabel string)
func (_DitCoordinator *DitCoordinatorCaller) GetKnowledgeLabels(opts *bind.CallOpts, _repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getKnowledgeLabels", _repository, _knowledgeLabelID)
	return *ret0, err
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(_repository bytes32, _knowledgeLabelID uint256) constant returns(knowledgeLabel string)
func (_DitCoordinator *DitCoordinatorSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitCoordinator.Contract.GetKnowledgeLabels(&_DitCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(_repository bytes32, _knowledgeLabelID uint256) constant returns(knowledgeLabel string)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitCoordinator.Contract.GetKnowledgeLabels(&_DitCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(_repository bytes32, _proposalID uint256) constant returns(hasPassed bool)
func (_DitCoordinator *DitCoordinatorCaller) ProposalHasPassed(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "proposalHasPassed", _repository, _proposalID)
	return *ret0, err
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(_repository bytes32, _proposalID uint256) constant returns(hasPassed bool)
func (_DitCoordinator *DitCoordinatorSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitCoordinator.Contract.ProposalHasPassed(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(_repository bytes32, _proposalID uint256) constant returns(hasPassed bool)
func (_DitCoordinator *DitCoordinatorCallerSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitCoordinator.Contract.ProposalHasPassed(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository( bytes32,  uint256) constant returns(KNWVoteID uint256, knowledgeLabel string, proposer address, isFinalized bool, proposalAccepted bool, individualStake uint256, totalStake uint256)
func (_DitCoordinator *DitCoordinatorCaller) ProposalsOfRepository(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	ret := new(struct {
		KNWVoteID        *big.Int
		KnowledgeLabel   string
		Proposer         common.Address
		IsFinalized      bool
		ProposalAccepted bool
		IndividualStake  *big.Int
		TotalStake       *big.Int
	})
	out := ret
	err := _DitCoordinator.contract.Call(opts, out, "proposalsOfRepository", arg0, arg1)
	return *ret, err
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository( bytes32,  uint256) constant returns(KNWVoteID uint256, knowledgeLabel string, proposer address, isFinalized bool, proposalAccepted bool, individualStake uint256, totalStake uint256)
func (_DitCoordinator *DitCoordinatorSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	return _DitCoordinator.Contract.ProposalsOfRepository(&_DitCoordinator.CallOpts, arg0, arg1)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository( bytes32,  uint256) constant returns(KNWVoteID uint256, knowledgeLabel string, proposer address, isFinalized bool, proposalAccepted bool, individualStake uint256, totalStake uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	return _DitCoordinator.Contract.ProposalsOfRepository(&_DitCoordinator.CallOpts, arg0, arg1)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories( bytes32) constant returns(votingMajority uint256, mintingMethod uint256, burningMethod uint256, currentProposalID uint256, minVoteCommitDuration uint256, maxVoteCommitDuration uint256, minVoteOpenDuration uint256, maxVoteOpenDuration uint256)
func (_DitCoordinator *DitCoordinatorCaller) Repositories(opts *bind.CallOpts, arg0 [32]byte) (struct {
	VotingMajority        *big.Int
	MintingMethod         *big.Int
	BurningMethod         *big.Int
	CurrentProposalID     *big.Int
	MinVoteCommitDuration *big.Int
	MaxVoteCommitDuration *big.Int
	MinVoteOpenDuration   *big.Int
	MaxVoteOpenDuration   *big.Int
}, error) {
	ret := new(struct {
		VotingMajority        *big.Int
		MintingMethod         *big.Int
		BurningMethod         *big.Int
		CurrentProposalID     *big.Int
		MinVoteCommitDuration *big.Int
		MaxVoteCommitDuration *big.Int
		MinVoteOpenDuration   *big.Int
		MaxVoteOpenDuration   *big.Int
	})
	out := ret
	err := _DitCoordinator.contract.Call(opts, out, "repositories", arg0)
	return *ret, err
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories( bytes32) constant returns(votingMajority uint256, mintingMethod uint256, burningMethod uint256, currentProposalID uint256, minVoteCommitDuration uint256, maxVoteCommitDuration uint256, minVoteOpenDuration uint256, maxVoteOpenDuration uint256)
func (_DitCoordinator *DitCoordinatorSession) Repositories(arg0 [32]byte) (struct {
	VotingMajority        *big.Int
	MintingMethod         *big.Int
	BurningMethod         *big.Int
	CurrentProposalID     *big.Int
	MinVoteCommitDuration *big.Int
	MaxVoteCommitDuration *big.Int
	MinVoteOpenDuration   *big.Int
	MaxVoteOpenDuration   *big.Int
}, error) {
	return _DitCoordinator.Contract.Repositories(&_DitCoordinator.CallOpts, arg0)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories( bytes32) constant returns(votingMajority uint256, mintingMethod uint256, burningMethod uint256, currentProposalID uint256, minVoteCommitDuration uint256, maxVoteCommitDuration uint256, minVoteOpenDuration uint256, maxVoteOpenDuration uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) Repositories(arg0 [32]byte) (struct {
	VotingMajority        *big.Int
	MintingMethod         *big.Int
	BurningMethod         *big.Int
	CurrentProposalID     *big.Int
	MinVoteCommitDuration *big.Int
	MaxVoteCommitDuration *big.Int
	MinVoteOpenDuration   *big.Int
	MaxVoteOpenDuration   *big.Int
}, error) {
	return _DitCoordinator.Contract.Repositories(&_DitCoordinator.CallOpts, arg0)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(_repository bytes32) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) RepositoryIsInitialized(opts *bind.CallOpts, _repository [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "repositoryIsInitialized", _repository)
	return *ret0, err
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(_repository bytes32) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitCoordinator.Contract.RepositoryIsInitialized(&_DitCoordinator.CallOpts, _repository)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(_repository bytes32) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitCoordinator.Contract.RepositoryIsInitialized(&_DitCoordinator.CallOpts, _repository)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(_repository bytes32, _proposalID uint256) returns()
func (_DitCoordinator *DitCoordinatorTransactor) FinalizeVote(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "finalizeVote", _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(_repository bytes32, _proposalID uint256) returns()
func (_DitCoordinator *DitCoordinatorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.FinalizeVote(&_DitCoordinator.TransactOpts, _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(_repository bytes32, _proposalID uint256) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.FinalizeVote(&_DitCoordinator.TransactOpts, _repository, _proposalID)
}

// InitRepository is a paid mutator transaction binding the contract method 0x51f43c24.
//
// Solidity: function initRepository(_repository bytes32, _label1 string, _label2 string, _label3 string, _voteSettings uint256[7]) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository [32]byte, _label1 string, _label2 string, _label3 string, _voteSettings [7]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "initRepository", _repository, _label1, _label2, _label3, _voteSettings)
}

// InitRepository is a paid mutator transaction binding the contract method 0x51f43c24.
//
// Solidity: function initRepository(_repository bytes32, _label1 string, _label2 string, _label3 string, _voteSettings uint256[7]) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) InitRepository(_repository [32]byte, _label1 string, _label2 string, _label3 string, _voteSettings [7]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _voteSettings)
}

// InitRepository is a paid mutator transaction binding the contract method 0x51f43c24.
//
// Solidity: function initRepository(_repository bytes32, _label1 string, _label2 string, _label3 string, _voteSettings uint256[7]) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) InitRepository(_repository [32]byte, _label1 string, _label2 string, _label3 string, _voteSettings [7]*big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _voteSettings)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(_repository bytes32, _proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitCoordinator *DitCoordinatorTransactor) OpenVoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "openVoteOnProposal", _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(_repository bytes32, _proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitCoordinator *DitCoordinatorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.OpenVoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(_repository bytes32, _proposalID uint256, _voteOption uint256, _voteSalt uint256) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.OpenVoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x0aba8688.
//
// Solidity: function proposeCommit(_repository bytes32, _knowledgeLabelIndex uint256, _voteCommitDuration uint256, _voteOpenDuration uint256) returns()
func (_DitCoordinator *DitCoordinatorTransactor) ProposeCommit(opts *bind.TransactOpts, _repository [32]byte, _knowledgeLabelIndex *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "proposeCommit", _repository, _knowledgeLabelIndex, _voteCommitDuration, _voteOpenDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x0aba8688.
//
// Solidity: function proposeCommit(_repository bytes32, _knowledgeLabelIndex uint256, _voteCommitDuration uint256, _voteOpenDuration uint256) returns()
func (_DitCoordinator *DitCoordinatorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _voteCommitDuration, _voteOpenDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x0aba8688.
//
// Solidity: function proposeCommit(_repository bytes32, _knowledgeLabelIndex uint256, _voteCommitDuration uint256, _voteOpenDuration uint256) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _voteCommitDuration, _voteOpenDuration)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xa34c299a.
//
// Solidity: function voteOnProposal(_repository bytes32, _proposalID uint256, _voteHash bytes32) returns()
func (_DitCoordinator *DitCoordinatorTransactor) VoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "voteOnProposal", _repository, _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xa34c299a.
//
// Solidity: function voteOnProposal(_repository bytes32, _proposalID uint256, _voteHash bytes32) returns()
func (_DitCoordinator *DitCoordinatorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitCoordinator.Contract.VoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteHash)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xa34c299a.
//
// Solidity: function voteOnProposal(_repository bytes32, _proposalID uint256, _voteHash bytes32) returns()
func (_DitCoordinator *DitCoordinatorTransactorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte) (*types.Transaction, error) {
	return _DitCoordinator.Contract.VoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteHash)
}

// DitCoordinatorCommitVoteIterator is returned from FilterCommitVote and is used to iterate over the raw logs and unpacked data for CommitVote events raised by the DitCoordinator contract.
type DitCoordinatorCommitVoteIterator struct {
	Event *DitCoordinatorCommitVote // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorCommitVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorCommitVote)
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
		it.Event = new(DitCoordinatorCommitVote)
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
func (it *DitCoordinatorCommitVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorCommitVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorCommitVote represents a CommitVote event raised by the DitCoordinator contract.
type DitCoordinatorCommitVote struct {
	Repository    [32]byte
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Stake         *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitVote is a free log retrieval operation binding the contract event 0xa01eea487bb3ec75528c167ccf90452d4164ddda7b13c55b2a89751a8dc5fbc1.
//
// Solidity: e CommitVote(repository indexed bytes32, proposal indexed uint256, who indexed address, label string, stake uint256, numberOfVotes uint256)
func (_DitCoordinator *DitCoordinatorFilterer) FilterCommitVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorCommitVoteIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "CommitVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorCommitVoteIterator{contract: _DitCoordinator.contract, event: "CommitVote", logs: logs, sub: sub}, nil
}

// WatchCommitVote is a free log subscription operation binding the contract event 0xa01eea487bb3ec75528c167ccf90452d4164ddda7b13c55b2a89751a8dc5fbc1.
//
// Solidity: e CommitVote(repository indexed bytes32, proposal indexed uint256, who indexed address, label string, stake uint256, numberOfVotes uint256)
func (_DitCoordinator *DitCoordinatorFilterer) WatchCommitVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorCommitVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "CommitVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorCommitVote)
				if err := _DitCoordinator.contract.UnpackLog(event, "CommitVote", log); err != nil {
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

// DitCoordinatorFinalizeVoteIterator is returned from FilterFinalizeVote and is used to iterate over the raw logs and unpacked data for FinalizeVote events raised by the DitCoordinator contract.
type DitCoordinatorFinalizeVoteIterator struct {
	Event *DitCoordinatorFinalizeVote // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorFinalizeVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorFinalizeVote)
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
		it.Event = new(DitCoordinatorFinalizeVote)
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
func (it *DitCoordinatorFinalizeVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorFinalizeVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorFinalizeVote represents a FinalizeVote event raised by the DitCoordinator contract.
type DitCoordinatorFinalizeVote struct {
	Repository [32]byte
	Proposal   *big.Int
	Label      string
	Accepted   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeVote is a free log retrieval operation binding the contract event 0x6bd2699645e0f6c5547bdf0d053280e48fef1ab21514bd02c88610b1279b942a.
//
// Solidity: e FinalizeVote(repository indexed bytes32, proposal indexed uint256, label string, accepted bool)
func (_DitCoordinator *DitCoordinatorFilterer) FilterFinalizeVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int) (*DitCoordinatorFinalizeVoteIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "FinalizeVote", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorFinalizeVoteIterator{contract: _DitCoordinator.contract, event: "FinalizeVote", logs: logs, sub: sub}, nil
}

// WatchFinalizeVote is a free log subscription operation binding the contract event 0x6bd2699645e0f6c5547bdf0d053280e48fef1ab21514bd02c88610b1279b942a.
//
// Solidity: e FinalizeVote(repository indexed bytes32, proposal indexed uint256, label string, accepted bool)
func (_DitCoordinator *DitCoordinatorFilterer) WatchFinalizeVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorFinalizeVote, repository [][32]byte, proposal []*big.Int) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "FinalizeVote", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorFinalizeVote)
				if err := _DitCoordinator.contract.UnpackLog(event, "FinalizeVote", log); err != nil {
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

// DitCoordinatorOpenVoteIterator is returned from FilterOpenVote and is used to iterate over the raw logs and unpacked data for OpenVote events raised by the DitCoordinator contract.
type DitCoordinatorOpenVoteIterator struct {
	Event *DitCoordinatorOpenVote // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorOpenVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorOpenVote)
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
		it.Event = new(DitCoordinatorOpenVote)
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
func (it *DitCoordinatorOpenVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorOpenVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorOpenVote represents a OpenVote event raised by the DitCoordinator contract.
type DitCoordinatorOpenVote struct {
	Repository    [32]byte
	Proposal      *big.Int
	Who           common.Address
	Label         string
	Accept        bool
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenVote is a free log retrieval operation binding the contract event 0x864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab.
//
// Solidity: e OpenVote(repository indexed bytes32, proposal indexed uint256, who indexed address, label string, accept bool, numberOfVotes uint256)
func (_DitCoordinator *DitCoordinatorFilterer) FilterOpenVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorOpenVoteIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "OpenVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorOpenVoteIterator{contract: _DitCoordinator.contract, event: "OpenVote", logs: logs, sub: sub}, nil
}

// WatchOpenVote is a free log subscription operation binding the contract event 0x864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab.
//
// Solidity: e OpenVote(repository indexed bytes32, proposal indexed uint256, who indexed address, label string, accept bool, numberOfVotes uint256)
func (_DitCoordinator *DitCoordinatorFilterer) WatchOpenVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorOpenVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "OpenVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorOpenVote)
				if err := _DitCoordinator.contract.UnpackLog(event, "OpenVote", log); err != nil {
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

// DitCoordinatorProposeCommitIterator is returned from FilterProposeCommit and is used to iterate over the raw logs and unpacked data for ProposeCommit events raised by the DitCoordinator contract.
type DitCoordinatorProposeCommitIterator struct {
	Event *DitCoordinatorProposeCommit // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorProposeCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorProposeCommit)
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
		it.Event = new(DitCoordinatorProposeCommit)
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
func (it *DitCoordinatorProposeCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorProposeCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorProposeCommit represents a ProposeCommit event raised by the DitCoordinator contract.
type DitCoordinatorProposeCommit struct {
	Repository [32]byte
	Proposal   *big.Int
	Who        common.Address
	Label      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposeCommit is a free log retrieval operation binding the contract event 0x171fe77c3addce776991159eb3eb73b14d9187ebd06c1c34ea12355a84ddbd83.
//
// Solidity: e ProposeCommit(repository indexed bytes32, proposal indexed uint256, who indexed address, label string)
func (_DitCoordinator *DitCoordinatorFilterer) FilterProposeCommit(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorProposeCommitIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "ProposeCommit", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorProposeCommitIterator{contract: _DitCoordinator.contract, event: "ProposeCommit", logs: logs, sub: sub}, nil
}

// WatchProposeCommit is a free log subscription operation binding the contract event 0x171fe77c3addce776991159eb3eb73b14d9187ebd06c1c34ea12355a84ddbd83.
//
// Solidity: e ProposeCommit(repository indexed bytes32, proposal indexed uint256, who indexed address, label string)
func (_DitCoordinator *DitCoordinatorFilterer) WatchProposeCommit(opts *bind.WatchOpts, sink chan<- *DitCoordinatorProposeCommit, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "ProposeCommit", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorProposeCommit)
				if err := _DitCoordinator.contract.UnpackLog(event, "ProposeCommit", log); err != nil {
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
