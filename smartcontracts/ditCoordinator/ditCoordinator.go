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

// KNWVotingContractABI is the input ABI used to generate the binding from.
const KNWVotingContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"},{\"name\":\"winningSide\",\"type\":\"bool\"},{\"name\":\"amountOfKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeLabel\",\"type\":\"string\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"startVote\",\"outputs\":[{\"name\":\"voteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRepository\",\"type\":\"bytes32\"},{\"name\":\"_majority\",\"type\":\"uint256\"}],\"name\":\"addNewRepository\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"endVote\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"openVote\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"amountOfVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVotingContract *KNWVotingContractTransactor) AddNewRepository(opts *bind.TransactOpts, _newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "addNewRepository", _newRepository, _majority)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVotingContract *KNWVotingContractSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddNewRepository(&_KNWVotingContract.TransactOpts, _newRepository, _majority)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVotingContract *KNWVotingContractTransactorSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.AddNewRepository(&_KNWVotingContract.TransactOpts, _newRepository, _majority)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 amountOfVotes)
func (_KNWVotingContract *KNWVotingContractTransactor) CommitVote(opts *bind.TransactOpts, _voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "commitVote", _voteID, _address, _secretHash, _numberOfKNW)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 amountOfVotes)
func (_KNWVotingContract *KNWVotingContractSession) CommitVote(_voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.CommitVote(&_KNWVotingContract.TransactOpts, _voteID, _address, _secretHash, _numberOfKNW)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 amountOfVotes)
func (_KNWVotingContract *KNWVotingContractTransactorSession) CommitVote(_voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.CommitVote(&_KNWVotingContract.TransactOpts, _voteID, _address, _secretHash, _numberOfKNW)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVotingContract *KNWVotingContractTransactor) EndVote(opts *bind.TransactOpts, _voteID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "endVote", _voteID)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVotingContract *KNWVotingContractSession) EndVote(_voteID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.EndVote(&_KNWVotingContract.TransactOpts, _voteID)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVotingContract *KNWVotingContractTransactorSession) EndVote(_voteID *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.EndVote(&_KNWVotingContract.TransactOpts, _voteID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 amountOfKNW)
func (_KNWVotingContract *KNWVotingContractTransactor) FinalizeVote(opts *bind.TransactOpts, _voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "finalizeVote", _voteID, _voteOption, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 amountOfKNW)
func (_KNWVotingContract *KNWVotingContractSession) FinalizeVote(_voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.FinalizeVote(&_KNWVotingContract.TransactOpts, _voteID, _voteOption, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 amountOfKNW)
func (_KNWVotingContract *KNWVotingContractTransactorSession) FinalizeVote(_voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.FinalizeVote(&_KNWVotingContract.TransactOpts, _voteID, _voteOption, _address)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVotingContract *KNWVotingContractTransactor) OpenVote(opts *bind.TransactOpts, _voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "openVote", _voteID, _address, _voteOption, _salt)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVotingContract *KNWVotingContractSession) OpenVote(_voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.OpenVote(&_KNWVotingContract.TransactOpts, _voteID, _address, _voteOption, _salt)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVotingContract *KNWVotingContractTransactorSession) OpenVote(_voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.OpenVote(&_KNWVotingContract.TransactOpts, _voteID, _address, _voteOption, _salt)
}

// StartVote is a paid mutator transaction binding the contract method 0x6766b4a4.
//
// Solidity: function startVote(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVotingContract *KNWVotingContractTransactor) StartVote(opts *bind.TransactOpts, _repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "startVote", _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0x6766b4a4.
//
// Solidity: function startVote(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVotingContract *KNWVotingContractSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartVote(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0x6766b4a4.
//
// Solidity: function startVote(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVotingContract *KNWVotingContractTransactorSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartVote(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake, _numberOfKNW)
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
// Solidity: function sqrt(uint256 a) constant returns(uint256)
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
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathCallerSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// DitCoordinatorABI is the input ABI used to generate the binding from.
const DitCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getKNWVoteIDFromProposalID\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getCurrentProposalID\",\"outputs\":[{\"name\":\"currentProposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"openVoteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repositories\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"currentProposalID\",\"type\":\"uint256\"},{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsOfRepository\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"individualStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getVotingMajority\",\"outputs\":[{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OPEN_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getIndividualStake\",\"outputs\":[{\"name\":\"individualStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"},{\"name\":\"_label1\",\"type\":\"string\"},{\"name\":\"_label2\",\"type\":\"string\"},{\"name\":\"_label3\",\"type\":\"string\"},{\"name\":\"_votingMajority\",\"type\":\"uint256\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_OPEN_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"replaceDitManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelID\",\"type\":\"uint256\"}],\"name\":\"getKnowledgeLabels\",\"outputs\":[{\"name\":\"knowledgeLabel\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"voteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelIndex\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"},{\"name\":\"_voteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"_voteOpenDuration\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[{\"name\":\"proposalID\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURNING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"passedKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"repositoryIsInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"passKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"}],\"name\":\"migrateRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"},{\"name\":\"_lastDitCoordinator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"}],\"name\":\"InitializeRepository\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"ProposeCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"CommitVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"OpenVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"votedRight\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"FinalizeVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"FinalizeProposal\",\"type\":\"event\"}]"

// DitCoordinatorBin is the compiled bytecode used for deploying new contracts.
const DitCoordinatorBin = `0x608060405234801561001057600080fd5b50604051606080612e0a833981016040908152815160208301519190920151600160a060020a0382161580159061004f5750600160a060020a03831615155b15156100e257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4b4e57566f74696e6720616e64204b4e57546f6b656e2061646472657373206360448201527f616e277420626520656d70747900000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03938416600160a060020a031991821617808355600580548316918616919091179055600180549585169582169590951785556002805493909416928116929092179092553380835260096020526040909220805460ff191690931790925560048054909216179055612ca5806101656000396000f3006080604052600436106101a05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306ee459681146101a55780630bdc90e8146101d25780630ee62ec0146101ea5780631341f25c1461021f5780631ecbb9de146102405780631f51fd71146102555780632e71d0fb146102f357806339ba645b1461030e5780633e029f631461032f5780633eedfc10146103fd5780633fcc148d14610412578063466577cd14610240578063552edc9d1461042a5780635d5809bb146104455780636fcfeb3b146103fd57806373b0dddd1461048c57806387c9203d146104ad57806391016157146104c857806395332229146104e9578063985dbfc51461057957806399821d9b146105aa578063ab4b593e146105bf578063bef7a4da146105d3578063c814af1f146105ea578063ccd9aa68146105ff578063d0c397ef14610620578063d7ad0eae14610641578063ea6c6d0f14610656578063eb2c02231461066e578063eb49fe941461068f578063eb931024146106a4578063effb21e1146105ea578063f60055cb146106c5575b600080fd5b3480156101b157600080fd5b506101c06004356024356106e5565b60408051918252519081900360200190f35b3480156101de57600080fd5b506101c0600435610702565b3480156101f657600080fd5b5061020b600435602435604435606435610717565b604080519115158252519081900360200190f35b34801561022b57600080fd5b5061020b600160a060020a03600435166108fb565b34801561024c57600080fd5b506101c0610910565b34801561026157600080fd5b5061026d600435610917565b6040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156102b657818101518382015260200161029e565b50505050905090810190601f1680156102e35780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156102ff57600080fd5b5061020b6004356024356109c4565b34801561031a57600080fd5b5061020b600160a060020a0360043516610fd5565b34801561033b57600080fd5b5061034a60043560243561101c565b60408051888152600160a060020a038716918101919091528415156060820152831515608082015260a0810183905260c0810182905260e0602080830182815289519284019290925288516101008401918a019080838360005b838110156103bc5781810151838201526020016103a4565b50505050905090810190601f1680156103e95780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561040957600080fd5b506101c0611110565b34801561041e57600080fd5b506101c0600435611116565b34801561043657600080fd5b506101c060043560243561112b565b34801561045157600080fd5b5061020b602460048035828101929082013591813580830192908201359160443580830192908201359160643591820191013560843561114b565b34801561049857600080fd5b5061020b600160a060020a0360043516611698565b3480156104b957600080fd5b5061020b6004356024356116df565b3480156104d457600080fd5b5061020b600160a060020a0360043516611799565b3480156104f557600080fd5b506105046004356024356117fa565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561053e578181015183820152602001610526565b50505050905090810190601f16801561056b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058557600080fd5b5061058e6118ab565b60408051600160a060020a039092168252519081900360200190f35b3480156105b657600080fd5b5061058e6118ba565b61020b6004356024356044356064356118c9565b6101c0600435602435604435606435608435611ca4565b3480156105f657600080fd5b506101c06123c8565b34801561060b57600080fd5b5061020b600160a060020a03600435166123cd565b34801561062c57600080fd5b5061020b600160a060020a03600435166123e2565b34801561064d57600080fd5b5061058e61242d565b34801561066257600080fd5b5061020b60043561243c565b34801561067a57600080fd5b5061020b600160a060020a0360043516612452565b34801561069b57600080fd5b5061058e6124b3565b3480156106b057600080fd5b5061020b600160a060020a03600435166124c2565b3480156106d157600080fd5b5061020b600480356024810191013561250d565b600091825260076020908152604080842092845291905290205490565b60009081526006602052604090206004015490565b3360008181526008602052604081205490919060ff16151561073857600080fd5b600554600087815260076020908152604080832089845282528083205481517fcdd6ceb9000000000000000000000000000000000000000000000000000000008152600481019190915233602482015260448101899052606481018890529051600160a060020a039094169363cdd6ceb993608480840194938390030190829087803b1580156107c757600080fd5b505af11580156107db573d6000803e3d6000fd5b505050506040513d60208110156107f157600080fd5b505060008681526007602090815260408083208884528252808320338085526005820184529382902060028082018a9055905483516001808c1496820187905294810182905260608082529385018054600019968116156101000296909601909516929092049282018390528a948c947f864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab9490939192919081906080820190869080156108df5780601f106108b4576101008083540402835291602001916108df565b820191906000526020600020905b8154815290600101906020018083116108c257829003601f168201915b505094505050505060405180910390a450600195945050505050565b60096020526000908152604090205460ff1681565b6212750081565b60066020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529092918391908301828280156109ae5780601f10610983576101008083540402835291602001916109ae565b820191906000526020600020905b81548152906001019060200180831161099157829003601f168201915b5050505050908060040154908060050154905083565b33600081815260086020526040812054909182918291829160ff1615156109ea57600080fd5b6000878152600760209081526040808320898452825280832033845260050190915290206003015460ff1615610a90576040805160e560020a62461bcd02815260206004820152602760248201527f45616368207061727469636970616e742063616e206f6e6c792066696e616c6960448201527f7a65206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600087815260076020908152604080832089845282528083203384526005019091528120541180610ae557506000878152600760209081526040808320898452909152902060020154600160a060020a031633145b1515610b61576040805160e560020a62461bcd02815260206004820152603a60248201527f4f6e6c79207061727469636970616e7473206f662074686520766f746520617260448201527f652061626c6520746f207265736f6c76652074686520766f7465000000000000606482015290519081900360840190fd5b600087815260076020908152604080832089845290915290206002015460a060020a900460ff161515610d775760055460008881526007602090815260408083208a845282528083205481517f865df0ad00000000000000000000000000000000000000000000000000000000815260048101919091529051600160a060020a039094169363865df0ad93602480840194938390030190829087803b158015610c0957600080fd5b505af1158015610c1d573d6000803e3d6000fd5b505050506040513d6020811015610c3357600080fd5b505160008881526007602090815260408083208a84528252918290206002808201805474ff000000000000000000000000000000000000000019961515750100000000000000000000000000000000000000000090810275ff00000000000000000000000000000000000000000019909216919091179690961660a060020a1790819055845195900460ff1680151593860193909352838552600191820180549283161561010002600019019092160492840183905289938b937f56c2d720d2b0f46900eca91b5412e4bb9ef934c72b86308c576d07975fac83539391908190606082019085908015610d675780601f10610d3c57610100808354040283529160200191610d67565b820191906000526020600020905b815481529060010190602001808311610d4a57829003601f168201915b5050935050505060405180910390a35b6005805460008981526007602090815260408083208b84528252808320805433808652919096019092528083206002015481517f36bf4c9100000000000000000000000000000000000000000000000000000000815260048101969096526024860152604485019190915251600160a060020a03909216926336bf4c919260648083019360609383900390910190829087803b158015610e1657600080fd5b505af1158015610e2a573d6000803e3d6000fd5b505050506040513d6060811015610e4057600080fd5b508051602082015160409092015190955090935091506000841115610e8e57604051339085156108fc029086906000818181858888f19350505050158015610e8c573d6000803e3d6000fd5b505b6000878152600760209081526040808320898452909152902060040154610ebb908563ffffffff6129ee16565b60008881526007602090815260408083208a8452808352818420600481019590955533808552600586018452828520600301805460ff19166001908117909155948c905290835281518815159381019390935290820186905260608083529383018054600294811615610100026000190116939093049382018490529289928b927fa4a57ebc87f48fa9e8fc4d0812bf408bff87f2326e00c7d209a0d42185b79ec592899189918190608082019086908015610fb85780601f10610f8d57610100808354040283529160200191610fb8565b820191906000526020600020905b815481529060010190602001808311610f9b57829003601f168201915b505094505050505060405180910390a45060019695505050505050565b3360008181526009602052604081205490919060ff161515610ff657600080fd5b5050600160a060020a03166000908152600860205260409020805460ff19169055600190565b60076020908152600092835260408084208252918352918190208054600180830180548551600293821615610100026000190190911692909204601f81018790048702830187019095528482529194929390928301828280156110c05780601f10611095576101008083540402835291602001916110c0565b820191906000526020600020905b8154815290600101906020018083116110a357829003601f168201915b505050600284015460038501546004909501549394600160a060020a0382169460ff60a060020a840481169550750100000000000000000000000000000000000000000090930490921692509087565b61025881565b60009081526006602052604090206005015490565b600091825260076020908152604080842092845291905290206003015490565b336000818152600860205260408120549091829160ff16151561116d57600080fd5b8a15156111e9576040805160e560020a62461bcd028152602060048201526024808201527f5265706f7369746f72792064657363726970746f722063616e2774206265206560448201527f6d70747900000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b8b8b604051602001808383808284378201915050925050506040516020818303038152906040526040518082805190602001908083835b6020831061123f5780518252601f199092019160209182019101611220565b51815160209384036101000a600019018019909216911617905260408051929094018290039091206000818152600690925292902060050154919550501591506112fb9050576040805160e560020a62461bcd02815260206004820152602760248201527f5265706f7369746f72792063616e206f6e6c7920626520696e697469616c697a60448201527f6564206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6032841015611354576040805160e560020a62461bcd02815260206004820152601f60248201527f566f74696e67206d616a6f726974792068617320746f206265203e3d20353000604482015290519081900360640190fd5b60008911806113635750600087115b8061136e5750600085115b15156113e9576040805160e560020a62461bcd028152602060048201526024808201527f50726f76696465206174206c65617374206f6e65204b6e6f776c65646765204c60448201527f6162656c00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354600160a060020a031615611470576040805160e560020a62461bcd02815260206004820152602260248201527f54686572652069732061206e6577657220636f6e7472616374206465706c6f7960448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6040805160a06020601f8f018190040282018101909252608081018d815290918291908f908f908190850183828082843782019150505050505081526020016060604051908101604052808d8d8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200189898080601f0160208091040260200160405190810160405280939291908181526020018383808284375050509290935250505081526000602080830182905260409283018890528582526006815291902082518051919261159292849290910190612b00565b5060208201516115a89060018301906003612b7e565b50604082810151600480840191909155606090930151600592830155905481517f828c1653000000000000000000000000000000000000000000000000000000008152928301859052602483018790529051600160a060020a039091169163828c16539160448083019260209291908290030181600087803b15801561162d57600080fd5b505af1158015611641573d6000803e3d6000fd5b505050506040513d602081101561165757600080fd5b5050604051339083907fe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be90600090a35060019b9a5050505050505050505050565b3360008181526009602052604081205490919060ff1615156116b957600080fd5b5050600160a060020a03166000908152600960205260409020805460ff19169055600190565b600082815260076020908152604080832084845290915281206002015460a060020a900460ff16151561175c576040805160e560020a62461bcd02815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b5060009182526007602090815260408084209284529190529020600201547501000000000000000000000000000000000000000000900460ff1690565b600454600090600160a060020a031633146117b357600080fd5b600160a060020a03821615156117c857600080fd5b5060048054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b6000828152600660205260409020606090600101826003811061181957fe5b01805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561189e5780601f106118735761010080835404028352916020019161189e565b820191906000526020600020905b81548152906001019060200180831161188157829003601f168201915b5050505050905092915050565b600154600160a060020a031681565b600354600160a060020a031681565b336000818152600860205260408120549091829160ff1615156118eb57600080fd5b60008781526007602090815260408083208984529091529020600301543414611984576040805160e560020a62461bcd02815260206004820152603960248201527f56616c7565206f6620746865207472616e73616374696f6e20646f65736e277460448201527f206d6174636820746865207265717569726564207374616b6500000000000000606482015290519081900360840190fd5b6000878152600760209081526040808320898452909152902060020154600160a060020a0316331415611a27576040805160e560020a62461bcd02815260206004820152603160248201527f5468652070726f706f736572206973206e6f7420616c6c6f77656420746f207660448201527f6f746520696e20612070726f706f73616c000000000000000000000000000000606482015290519081900360840190fd5b6000878152600760209081526040808320898452909152902060040154611a54903463ffffffff612a7616565b60008881526007602090815260408083208a84528252808320600480820195909555600554905482517fd4e0ac9500000000000000000000000000000000000000000000000000000000815295860152336024860152604485018a9052606485018990529051600160a060020a039091169363d4e0ac959360848083019493928390030190829087803b158015611aea57600080fd5b505af1158015611afe573d6000803e3d6000fd5b505050506040513d6020811015611b1457600080fd5b5051915060008211611b96576040805160e560020a62461bcd02815260206004820152603360248201527f566f74696e6720636f6e74726163742072657475726e656420616e20696e766160448201527f6c696420616d6f756e74206f6620766f74657300000000000000000000000000606482015290519081900360840190fd5b600087815260076020908152604080832089845280835281842033808652600582018552838620889055948b905290835281513493810184905291820188905260608201869052608080835260019182018054600261010094821615949094026000190116929092049083018190528a938c937f7ee8ecf4d9d20fb6454a55c418451d97bec229903525d8517fcd32db68a479e493928b918a91819060a082019087908015611c865780601f10611c5b57610100808354040283529160200191611c86565b820191906000526020600020905b815481529060010190602001808311611c6957829003601f168201915b50509550505050505060405180910390a45060019695505050505050565b3360008181526008602052604081205490919060ff161515611cc557600080fd5b60003411611d43576040805160e560020a62461bcd02815260206004820152602860248201527f56616c7565206f6620746865207472616e73616374696f6e2063616e206e6f7460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008781526006602052604081206001018760038110611d5f57fe5b01546002600019610100600184161502019091160411611dee576040805160e560020a62461bcd028152602060048201526024808201527f4b6e6f776c656467652d4c6162656c20696e646578206973206e6f7420636f7260448201527f7265637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6102588410158015611e035750621275008411155b1515611e59576040805160e560020a62461bcd02815260206004820152601c60248201527f566f746520636f6d6d6974206475726174696f6e20696e76616c696400000000604482015290519081900360640190fd5b6102588310158015611e6e5750621275008311155b1515611ec4576040805160e560020a62461bcd02815260206004820152601a60248201527f566f7465206f70656e206475726174696f6e20696e76616c6964000000000000604482015290519081900360640190fd5b600354600160a060020a031615611f4b576040805160e560020a62461bcd02815260206004820152602260248201527f54686572652069732061206e6577657220636f6e7472616374206465706c6f7960448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600087815260066020526040902060040154611f6e90600163ffffffff612a7616565b600088815260066020526040908190206004810192909255805160e0810190915260055490918291600160a060020a031690636766b4a4908b9033906001018c60038110611fb857fe5b018a8a348e6040518863ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180886000191660001916815260200187600160a060020a0316600160a060020a03168152602001806020018681526020018581526020018481526020018381526020018281038252878181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156120af5780601f10612084576101008083540402835291602001916120af565b820191906000526020600020905b81548152906001019060200180831161209257829003601f168201915b505098505050505050505050602060405180830381600087803b1580156120d557600080fd5b505af11580156120e9573d6000803e3d6000fd5b505050506040513d60208110156120ff57600080fd5b505181526000898152600660209081526040909120910190600101886003811061212557fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156121aa5780601f1061217f576101008083540402835291602001916121aa565b820191906000526020600020905b81548152906001019060200180831161218d57829003601f168201915b5050509183525050336020808301919091526000604080840182905260608401829052346080850181905260a0909401939093528a8152600782528281206006835283822060040154825282529190912082518155828201518051919261221992600185019290910190612b00565b5060408281015160028301805460608601516080870151151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001991151560a060020a0274ff000000000000000000000000000000000000000019600160a060020a0390961673ffffffffffffffffffffffffffffffffffffffff199094169390931794909416919091171691909117905560a083015160038084019190915560c09093015160049283015560008a815260066020522090810154339290918a917f2ba422bdea9179f02f8c9dd0d6285478f7b4c3fa11a812eeb4d3b1b04cc57c3591600101908b90811061231857fe5b60408051602081018d905281815292909101805460026000196101006001841615020190911604918301829052918b91819060608201908590801561239e5780601f106123735761010080835404028352916020019161239e565b820191906000526020600020905b81548152906001019060200180831161238157829003601f168201915b5050935050505060405180910390a450505060009384525050600660205250604090206004015490565b600081565b60086020526000908152604090205460ff1681565b3360008181526009602052604081205490919060ff16151561240357600080fd5b5050600160a060020a03166000908152600960205260409020805460ff1916600190811790915590565b600254600160a060020a031681565b6000908152600660205260408120600501541190565b600454600090600160a060020a0316331461246c57600080fd5b600160a060020a038216151561248157600080fd5b5060038054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600054600160a060020a031681565b3360008181526009602052604081205490919060ff1615156124e357600080fd5b5050600160a060020a03166000908152600860205260409020805460ff1916600190811790915590565b600080600080600061251d612bca565b3360008181526008602052604081205490919060ff16151561253e57600080fd5b600254600160a060020a0316151561255557600080fd5b600254600160a060020a031696508815156125df576040805160e560020a62461bcd028152602060048201526024808201527f5265706f7369746f72792064657363726970746f722063616e2774206265206560448201527f6d70747900000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b8989604051602001808383808284378201915050925050506040516020818303038152906040526040518082805190602001908083835b602083106126355780518252601f199092019160209182019101612616565b51815160209384036101000a6000190180199092169116179052604080519290940182900382207f0bdc90e8000000000000000000000000000000000000000000000000000000008352600483018190529351939b50600160a060020a038d169550630bdc90e8945060248083019491935090918290030181600087803b1580156126bf57600080fd5b505af11580156126d3573d6000803e3d6000fd5b505050506040513d60208110156126e957600080fd5b5051604080517f3fcc148d000000000000000000000000000000000000000000000000000000008152600481018990529051919650600160a060020a03891691633fcc148d916024808201926020929091908290030181600087803b15801561275157600080fd5b505af1158015612765573d6000803e3d6000fd5b505050506040513d602081101561277b57600080fd5b50519350600091505b60038260ff1610156128a357604080517f953322290000000000000000000000000000000000000000000000000000000081526004810188905260ff841660248201529051600160a060020a03891691639533222991604480830192600092919082900301818387803b1580156127fa57600080fd5b505af115801561280e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561283757600080fd5b81019080805164010000000081111561284f57600080fd5b8201602081018481111561286257600080fd5b815164010000000081118282018710171561287c57600080fd5b50909350869250505060ff84166003811061289357fe5b6020020152600190910190612784565b6040805160a06020601f8d018190040282018101909252608081018b815290918291908d908d90819085018382808284375050509284525050506020808201869052604080830189905260609092018790526000898152600682529190912082518051919261291792849290910190612b00565b50602082015161292d9060018301906003612b7e565b50604082810151600480840191909155606090930151600592830155905481517f828c1653000000000000000000000000000000000000000000000000000000008152928301899052602483018790529051600160a060020a039091169163828c16539160448083019260209291908290030181600087803b1580156129b257600080fd5b505af11580156129c6573d6000803e3d6000fd5b505050506040513d60208110156129dc57600080fd5b5060019b9a5050505050505050505050565b60008083831115612a6f576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b600082820183811015612af9576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612b4157805160ff1916838001178555612b6e565b82800160010185558215612b6e579182015b82811115612b6e578251825591602001919060010190612b53565b50612b7a929150612bf2565b5090565b8260038101928215612bbe579160200282015b82811115612bbe5782518051612bae918491602090910190612b00565b5091602001919060010190612b91565b50612b7a929150612c0f565b6060604051908101604052806003905b6060815260200190600190039081612bda5790505090565b612c0c91905b80821115612b7a5760008155600101612bf8565b90565b612c0c91905b80821115612b7a576000612c298282612c32565b50600101612c15565b50805460018160011615610100020316600290046000825580601f10612c585750612c76565b601f016020900490600052602060002090810190612c769190612bf2565b505600a165627a7a7230582083a202fdb6344db89f10fc09e66ec45922a3181e552f3c9a0fb8fc6ad58def0b0029`

// DeployDitCoordinator deploys a new Ethereum contract, binding an instance of DitCoordinator to it.
func DeployDitCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWTokenAddress common.Address, _KNWVotingAddress common.Address, _lastDitCoordinator common.Address) (common.Address, *types.Transaction, *DitCoordinator, error) {
	parsed, err := abi.JSON(strings.NewReader(DitCoordinatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DitCoordinatorBin), backend, _KNWTokenAddress, _KNWVotingAddress, _lastDitCoordinator)
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

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) BURNINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "BURNING_METHOD")
	return *ret0, err
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) BURNINGMETHOD() (*big.Int, error) {
	return _DitCoordinator.Contract.BURNINGMETHOD(&_DitCoordinator.CallOpts)
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) BURNINGMETHOD() (*big.Int, error) {
	return _DitCoordinator.Contract.BURNINGMETHOD(&_DitCoordinator.CallOpts)
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

// MAXOPENDURATION is a free data retrieval call binding the contract method 0x466577cd.
//
// Solidity: function MAX_OPEN_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) MAXOPENDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "MAX_OPEN_DURATION")
	return *ret0, err
}

// MAXOPENDURATION is a free data retrieval call binding the contract method 0x466577cd.
//
// Solidity: function MAX_OPEN_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) MAXOPENDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MAXOPENDURATION(&_DitCoordinator.CallOpts)
}

// MAXOPENDURATION is a free data retrieval call binding the contract method 0x466577cd.
//
// Solidity: function MAX_OPEN_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) MAXOPENDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MAXOPENDURATION(&_DitCoordinator.CallOpts)
}

// MAXVOTEDURATION is a free data retrieval call binding the contract method 0x1ecbb9de.
//
// Solidity: function MAX_VOTE_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) MAXVOTEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "MAX_VOTE_DURATION")
	return *ret0, err
}

// MAXVOTEDURATION is a free data retrieval call binding the contract method 0x1ecbb9de.
//
// Solidity: function MAX_VOTE_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) MAXVOTEDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MAXVOTEDURATION(&_DitCoordinator.CallOpts)
}

// MAXVOTEDURATION is a free data retrieval call binding the contract method 0x1ecbb9de.
//
// Solidity: function MAX_VOTE_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) MAXVOTEDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MAXVOTEDURATION(&_DitCoordinator.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) MINTINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "MINTING_METHOD")
	return *ret0, err
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) MINTINGMETHOD() (*big.Int, error) {
	return _DitCoordinator.Contract.MINTINGMETHOD(&_DitCoordinator.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) MINTINGMETHOD() (*big.Int, error) {
	return _DitCoordinator.Contract.MINTINGMETHOD(&_DitCoordinator.CallOpts)
}

// MINOPENDURATION is a free data retrieval call binding the contract method 0x6fcfeb3b.
//
// Solidity: function MIN_OPEN_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) MINOPENDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "MIN_OPEN_DURATION")
	return *ret0, err
}

// MINOPENDURATION is a free data retrieval call binding the contract method 0x6fcfeb3b.
//
// Solidity: function MIN_OPEN_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) MINOPENDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MINOPENDURATION(&_DitCoordinator.CallOpts)
}

// MINOPENDURATION is a free data retrieval call binding the contract method 0x6fcfeb3b.
//
// Solidity: function MIN_OPEN_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) MINOPENDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MINOPENDURATION(&_DitCoordinator.CallOpts)
}

// MINVOTEDURATION is a free data retrieval call binding the contract method 0x3eedfc10.
//
// Solidity: function MIN_VOTE_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCaller) MINVOTEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "MIN_VOTE_DURATION")
	return *ret0, err
}

// MINVOTEDURATION is a free data retrieval call binding the contract method 0x3eedfc10.
//
// Solidity: function MIN_VOTE_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorSession) MINVOTEDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MINVOTEDURATION(&_DitCoordinator.CallOpts)
}

// MINVOTEDURATION is a free data retrieval call binding the contract method 0x3eedfc10.
//
// Solidity: function MIN_VOTE_DURATION() constant returns(uint256)
func (_DitCoordinator *DitCoordinatorCallerSession) MINVOTEDURATION() (*big.Int, error) {
	return _DitCoordinator.Contract.MINVOTEDURATION(&_DitCoordinator.CallOpts)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256 currentProposalID)
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
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256 currentProposalID)
func (_DitCoordinator *DitCoordinatorSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetCurrentProposalID(&_DitCoordinator.CallOpts, _repository)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256 currentProposalID)
func (_DitCoordinator *DitCoordinatorCallerSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetCurrentProposalID(&_DitCoordinator.CallOpts, _repository)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
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
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitCoordinator *DitCoordinatorSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetIndividualStake(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitCoordinator *DitCoordinatorCallerSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetIndividualStake(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256 KNWVoteID)
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
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256 KNWVoteID)
func (_DitCoordinator *DitCoordinatorSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256 KNWVoteID)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
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
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitCoordinator *DitCoordinatorSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitCoordinator.Contract.GetKnowledgeLabels(&_DitCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitCoordinator.Contract.GetKnowledgeLabels(&_DitCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// GetVotingMajority is a free data retrieval call binding the contract method 0x3fcc148d.
//
// Solidity: function getVotingMajority(bytes32 _repository) constant returns(uint256 votingMajority)
func (_DitCoordinator *DitCoordinatorCaller) GetVotingMajority(opts *bind.CallOpts, _repository [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getVotingMajority", _repository)
	return *ret0, err
}

// GetVotingMajority is a free data retrieval call binding the contract method 0x3fcc148d.
//
// Solidity: function getVotingMajority(bytes32 _repository) constant returns(uint256 votingMajority)
func (_DitCoordinator *DitCoordinatorSession) GetVotingMajority(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetVotingMajority(&_DitCoordinator.CallOpts, _repository)
}

// GetVotingMajority is a free data retrieval call binding the contract method 0x3fcc148d.
//
// Solidity: function getVotingMajority(bytes32 _repository) constant returns(uint256 votingMajority)
func (_DitCoordinator *DitCoordinatorCallerSession) GetVotingMajority(_repository [32]byte) (*big.Int, error) {
	return _DitCoordinator.Contract.GetVotingMajority(&_DitCoordinator.CallOpts, _repository)
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) IsKYCValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "isKYCValidator", arg0)
	return *ret0, err
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) IsKYCValidator(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.IsKYCValidator(&_DitCoordinator.CallOpts, arg0)
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) IsKYCValidator(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.IsKYCValidator(&_DitCoordinator.CallOpts, arg0)
}

// LastDitCoordinator is a free data retrieval call binding the contract method 0xd7ad0eae.
//
// Solidity: function lastDitCoordinator() constant returns(address)
func (_DitCoordinator *DitCoordinatorCaller) LastDitCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "lastDitCoordinator")
	return *ret0, err
}

// LastDitCoordinator is a free data retrieval call binding the contract method 0xd7ad0eae.
//
// Solidity: function lastDitCoordinator() constant returns(address)
func (_DitCoordinator *DitCoordinatorSession) LastDitCoordinator() (common.Address, error) {
	return _DitCoordinator.Contract.LastDitCoordinator(&_DitCoordinator.CallOpts)
}

// LastDitCoordinator is a free data retrieval call binding the contract method 0xd7ad0eae.
//
// Solidity: function lastDitCoordinator() constant returns(address)
func (_DitCoordinator *DitCoordinatorCallerSession) LastDitCoordinator() (common.Address, error) {
	return _DitCoordinator.Contract.LastDitCoordinator(&_DitCoordinator.CallOpts)
}

// NextDitCoordinator is a free data retrieval call binding the contract method 0x99821d9b.
//
// Solidity: function nextDitCoordinator() constant returns(address)
func (_DitCoordinator *DitCoordinatorCaller) NextDitCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "nextDitCoordinator")
	return *ret0, err
}

// NextDitCoordinator is a free data retrieval call binding the contract method 0x99821d9b.
//
// Solidity: function nextDitCoordinator() constant returns(address)
func (_DitCoordinator *DitCoordinatorSession) NextDitCoordinator() (common.Address, error) {
	return _DitCoordinator.Contract.NextDitCoordinator(&_DitCoordinator.CallOpts)
}

// NextDitCoordinator is a free data retrieval call binding the contract method 0x99821d9b.
//
// Solidity: function nextDitCoordinator() constant returns(address)
func (_DitCoordinator *DitCoordinatorCallerSession) NextDitCoordinator() (common.Address, error) {
	return _DitCoordinator.Contract.NextDitCoordinator(&_DitCoordinator.CallOpts)
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) PassedKYC(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "passedKYC", arg0)
	return *ret0, err
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) PassedKYC(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.PassedKYC(&_DitCoordinator.CallOpts, arg0)
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) PassedKYC(arg0 common.Address) (bool, error) {
	return _DitCoordinator.Contract.PassedKYC(&_DitCoordinator.CallOpts, arg0)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
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
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitCoordinator *DitCoordinatorSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitCoordinator.Contract.ProposalHasPassed(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitCoordinator *DitCoordinatorCallerSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitCoordinator.Contract.ProposalHasPassed(&_DitCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
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
// Solidity: function repositories(bytes32 ) constant returns(string name, uint256 currentProposalID, uint256 votingMajority)
func (_DitCoordinator *DitCoordinatorCaller) Repositories(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Name              string
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	ret := new(struct {
		Name              string
		CurrentProposalID *big.Int
		VotingMajority    *big.Int
	})
	out := ret
	err := _DitCoordinator.contract.Call(opts, out, "repositories", arg0)
	return *ret, err
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(string name, uint256 currentProposalID, uint256 votingMajority)
func (_DitCoordinator *DitCoordinatorSession) Repositories(arg0 [32]byte) (struct {
	Name              string
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	return _DitCoordinator.Contract.Repositories(&_DitCoordinator.CallOpts, arg0)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(string name, uint256 currentProposalID, uint256 votingMajority)
func (_DitCoordinator *DitCoordinatorCallerSession) Repositories(arg0 [32]byte) (struct {
	Name              string
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	return _DitCoordinator.Contract.Repositories(&_DitCoordinator.CallOpts, arg0)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
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
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitCoordinator.Contract.RepositoryIsInitialized(&_DitCoordinator.CallOpts, _repository)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitCoordinator.Contract.RepositoryIsInitialized(&_DitCoordinator.CallOpts, _repository)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) AddKYCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "addKYCValidator", _address)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) AddKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.AddKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) AddKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.AddKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) FinalizeVote(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "finalizeVote", _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.FinalizeVote(&_DitCoordinator.TransactOpts, _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.FinalizeVote(&_DitCoordinator.TransactOpts, _repository, _proposalID)
}

// InitRepository is a paid mutator transaction binding the contract method 0x5d5809bb.
//
// Solidity: function initRepository(string _repository, string _label1, string _label2, string _label3, uint256 _votingMajority) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository string, _label1 string, _label2 string, _label3 string, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "initRepository", _repository, _label1, _label2, _label3, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x5d5809bb.
//
// Solidity: function initRepository(string _repository, string _label1, string _label2, string _label3, uint256 _votingMajority) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) InitRepository(_repository string, _label1 string, _label2 string, _label3 string, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x5d5809bb.
//
// Solidity: function initRepository(string _repository, string _label1, string _label2, string _label3, uint256 _votingMajority) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) InitRepository(_repository string, _label1 string, _label2 string, _label3 string, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _votingMajority)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xf60055cb.
//
// Solidity: function migrateRepository(string _repository) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) MigrateRepository(opts *bind.TransactOpts, _repository string) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "migrateRepository", _repository)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xf60055cb.
//
// Solidity: function migrateRepository(string _repository) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) MigrateRepository(_repository string) (*types.Transaction, error) {
	return _DitCoordinator.Contract.MigrateRepository(&_DitCoordinator.TransactOpts, _repository)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xf60055cb.
//
// Solidity: function migrateRepository(string _repository) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) MigrateRepository(_repository string) (*types.Transaction, error) {
	return _DitCoordinator.Contract.MigrateRepository(&_DitCoordinator.TransactOpts, _repository)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) OpenVoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "openVoteOnProposal", _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.OpenVoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.OpenVoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) PassKYC(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "passKYC", _address)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) PassKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.PassKYC(&_DitCoordinator.TransactOpts, _address)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) PassKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.PassKYC(&_DitCoordinator.TransactOpts, _address)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xbef7a4da.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _numberOfKNW, uint256 _voteCommitDuration, uint256 _voteOpenDuration) returns(uint256 proposalID)
func (_DitCoordinator *DitCoordinatorTransactor) ProposeCommit(opts *bind.TransactOpts, _repository [32]byte, _knowledgeLabelIndex *big.Int, _numberOfKNW *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "proposeCommit", _repository, _knowledgeLabelIndex, _numberOfKNW, _voteCommitDuration, _voteOpenDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xbef7a4da.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _numberOfKNW, uint256 _voteCommitDuration, uint256 _voteOpenDuration) returns(uint256 proposalID)
func (_DitCoordinator *DitCoordinatorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _numberOfKNW *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _numberOfKNW, _voteCommitDuration, _voteOpenDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xbef7a4da.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _numberOfKNW, uint256 _voteCommitDuration, uint256 _voteOpenDuration) returns(uint256 proposalID)
func (_DitCoordinator *DitCoordinatorTransactorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _numberOfKNW *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _numberOfKNW, _voteCommitDuration, _voteOpenDuration)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) RemoveKYCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "removeKYCValidator", _address)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) RemoveKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RemoveKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) RemoveKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RemoveKYCValidator(&_DitCoordinator.TransactOpts, _address)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) ReplaceDitManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "replaceDitManager", _newManager)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) ReplaceDitManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ReplaceDitManager(&_DitCoordinator.TransactOpts, _newManager)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) ReplaceDitManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ReplaceDitManager(&_DitCoordinator.TransactOpts, _newManager)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) RevokeKYC(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "revokeKYC", _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) RevokeKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RevokeKYC(&_DitCoordinator.TransactOpts, _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) RevokeKYC(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.RevokeKYC(&_DitCoordinator.TransactOpts, _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) UpgradeContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "upgradeContract", _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.UpgradeContract(&_DitCoordinator.TransactOpts, _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.UpgradeContract(&_DitCoordinator.TransactOpts, _address)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xab4b593e.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash, uint256 _numberOfKNW) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) VoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "voteOnProposal", _repository, _proposalID, _voteHash, _numberOfKNW)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xab4b593e.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash, uint256 _numberOfKNW) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.VoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteHash, _numberOfKNW)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xab4b593e.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash, uint256 _numberOfKNW) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.VoteOnProposal(&_DitCoordinator.TransactOpts, _repository, _proposalID, _voteHash, _numberOfKNW)
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
	NumberOfKNW   *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitVote is a free log retrieval operation binding the contract event 0x7ee8ecf4d9d20fb6454a55c418451d97bec229903525d8517fcd32db68a479e4.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
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

// WatchCommitVote is a free log subscription operation binding the contract event 0x7ee8ecf4d9d20fb6454a55c418451d97bec229903525d8517fcd32db68a479e4.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
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

// DitCoordinatorFinalizeProposalIterator is returned from FilterFinalizeProposal and is used to iterate over the raw logs and unpacked data for FinalizeProposal events raised by the DitCoordinator contract.
type DitCoordinatorFinalizeProposalIterator struct {
	Event *DitCoordinatorFinalizeProposal // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorFinalizeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorFinalizeProposal)
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
		it.Event = new(DitCoordinatorFinalizeProposal)
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
func (it *DitCoordinatorFinalizeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorFinalizeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorFinalizeProposal represents a FinalizeProposal event raised by the DitCoordinator contract.
type DitCoordinatorFinalizeProposal struct {
	Repository [32]byte
	Proposal   *big.Int
	Label      string
	Accepted   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeProposal is a free log retrieval operation binding the contract event 0x56c2d720d2b0f46900eca91b5412e4bb9ef934c72b86308c576d07975fac8353.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, string label, bool accepted)
func (_DitCoordinator *DitCoordinatorFilterer) FilterFinalizeProposal(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int) (*DitCoordinatorFinalizeProposalIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "FinalizeProposal", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorFinalizeProposalIterator{contract: _DitCoordinator.contract, event: "FinalizeProposal", logs: logs, sub: sub}, nil
}

// WatchFinalizeProposal is a free log subscription operation binding the contract event 0x56c2d720d2b0f46900eca91b5412e4bb9ef934c72b86308c576d07975fac8353.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, string label, bool accepted)
func (_DitCoordinator *DitCoordinatorFilterer) WatchFinalizeProposal(opts *bind.WatchOpts, sink chan<- *DitCoordinatorFinalizeProposal, repository [][32]byte, proposal []*big.Int) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "FinalizeProposal", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorFinalizeProposal)
				if err := _DitCoordinator.contract.UnpackLog(event, "FinalizeProposal", log); err != nil {
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
	Repository  [32]byte
	Proposal    *big.Int
	Who         common.Address
	Label       string
	VotedRight  bool
	NumberOfKNW *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizeVote is a free log retrieval operation binding the contract event 0xa4a57ebc87f48fa9e8fc4d0812bf408bff87f2326e00c7d209a0d42185b79ec5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool votedRight, uint256 numberOfKNW)
func (_DitCoordinator *DitCoordinatorFilterer) FilterFinalizeVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitCoordinatorFinalizeVoteIterator, error) {

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

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "FinalizeVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorFinalizeVoteIterator{contract: _DitCoordinator.contract, event: "FinalizeVote", logs: logs, sub: sub}, nil
}

// WatchFinalizeVote is a free log subscription operation binding the contract event 0xa4a57ebc87f48fa9e8fc4d0812bf408bff87f2326e00c7d209a0d42185b79ec5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool votedRight, uint256 numberOfKNW)
func (_DitCoordinator *DitCoordinatorFilterer) WatchFinalizeVote(opts *bind.WatchOpts, sink chan<- *DitCoordinatorFinalizeVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "FinalizeVote", repositoryRule, proposalRule, whoRule)
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

// DitCoordinatorInitializeRepositoryIterator is returned from FilterInitializeRepository and is used to iterate over the raw logs and unpacked data for InitializeRepository events raised by the DitCoordinator contract.
type DitCoordinatorInitializeRepositoryIterator struct {
	Event *DitCoordinatorInitializeRepository // Event containing the contract specifics and raw log

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
func (it *DitCoordinatorInitializeRepositoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitCoordinatorInitializeRepository)
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
		it.Event = new(DitCoordinatorInitializeRepository)
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
func (it *DitCoordinatorInitializeRepositoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitCoordinatorInitializeRepositoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitCoordinatorInitializeRepository represents a InitializeRepository event raised by the DitCoordinator contract.
type DitCoordinatorInitializeRepository struct {
	Repository [32]byte
	Who        common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitializeRepository is a free log retrieval operation binding the contract event 0xe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be.
//
// Solidity: event InitializeRepository(bytes32 indexed repository, address indexed who)
func (_DitCoordinator *DitCoordinatorFilterer) FilterInitializeRepository(opts *bind.FilterOpts, repository [][32]byte, who []common.Address) (*DitCoordinatorInitializeRepositoryIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.FilterLogs(opts, "InitializeRepository", repositoryRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitCoordinatorInitializeRepositoryIterator{contract: _DitCoordinator.contract, event: "InitializeRepository", logs: logs, sub: sub}, nil
}

// WatchInitializeRepository is a free log subscription operation binding the contract event 0xe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be.
//
// Solidity: event InitializeRepository(bytes32 indexed repository, address indexed who)
func (_DitCoordinator *DitCoordinatorFilterer) WatchInitializeRepository(opts *bind.WatchOpts, sink chan<- *DitCoordinatorInitializeRepository, repository [][32]byte, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitCoordinator.contract.WatchLogs(opts, "InitializeRepository", repositoryRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitCoordinatorInitializeRepository)
				if err := _DitCoordinator.contract.UnpackLog(event, "InitializeRepository", log); err != nil {
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
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool accept, uint256 numberOfVotes)
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
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool accept, uint256 numberOfVotes)
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
	Repository  [32]byte
	Proposal    *big.Int
	Who         common.Address
	Label       string
	NumberOfKNW *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposeCommit is a free log retrieval operation binding the contract event 0x2ba422bdea9179f02f8c9dd0d6285478f7b4c3fa11a812eeb4d3b1b04cc57c35.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 numberOfKNW)
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

// WatchProposeCommit is a free log subscription operation binding the contract event 0x2ba422bdea9179f02f8c9dd0d6285478f7b4c3fa11a812eeb4d3b1b04cc57c35.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 numberOfKNW)
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
