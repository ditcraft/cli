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
const KNWTokenContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"amountOfIDs\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KNWTokenContractFuncSigs maps the 4-byte function signature to its string representation.
var KNWTokenContractFuncSigs = map[string]string{
	"ffe03eba": "amountOfIDs()",
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

// AmountOfIDs is a free data retrieval call binding the contract method 0xffe03eba.
//
// Solidity: function amountOfIDs() constant returns(uint256 amount)
func (_KNWTokenContract *KNWTokenContractCaller) AmountOfIDs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWTokenContract.contract.Call(opts, out, "amountOfIDs")
	return *ret0, err
}

// AmountOfIDs is a free data retrieval call binding the contract method 0xffe03eba.
//
// Solidity: function amountOfIDs() constant returns(uint256 amount)
func (_KNWTokenContract *KNWTokenContractSession) AmountOfIDs() (*big.Int, error) {
	return _KNWTokenContract.Contract.AmountOfIDs(&_KNWTokenContract.CallOpts)
}

// AmountOfIDs is a free data retrieval call binding the contract method 0xffe03eba.
//
// Solidity: function amountOfIDs() constant returns(uint256 amount)
func (_KNWTokenContract *KNWTokenContractCallerSession) AmountOfIDs() (*big.Int, error) {
	return _KNWTokenContract.Contract.AmountOfIDs(&_KNWTokenContract.CallOpts)
}

// KNWVotingContractABI is the input ABI used to generate the binding from.
const KNWVotingContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"},{\"name\":\"winningSide\",\"type\":\"bool\"},{\"name\":\"amountOfKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRepository\",\"type\":\"bytes32\"},{\"name\":\"_majority\",\"type\":\"uint256\"}],\"name\":\"addNewRepository\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"endVote\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"openVote\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"amountOfVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeID\",\"type\":\"uint256\"},{\"name\":\"_voteDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"startVote\",\"outputs\":[{\"name\":\"voteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KNWVotingContractFuncSigs maps the 4-byte function signature to its string representation.
var KNWVotingContractFuncSigs = map[string]string{
	"828c1653": "addNewRepository(bytes32,uint256)",
	"d4e0ac95": "commitVote(uint256,address,bytes32,uint256)",
	"865df0ad": "endVote(uint256)",
	"36bf4c91": "finalizeVote(uint256,uint256,address)",
	"cdd6ceb9": "openVote(uint256,address,uint256,uint256)",
	"e5023ff2": "startVote(bytes32,address,uint256,uint256,uint256,uint256)",
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

// StartVote is a paid mutator transaction binding the contract method 0xe5023ff2.
//
// Solidity: function startVote(bytes32 _repository, address _address, uint256 _knowledgeID, uint256 _voteDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVotingContract *KNWVotingContractTransactor) StartVote(opts *bind.TransactOpts, _repository [32]byte, _address common.Address, _knowledgeID *big.Int, _voteDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.contract.Transact(opts, "startVote", _repository, _address, _knowledgeID, _voteDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0xe5023ff2.
//
// Solidity: function startVote(bytes32 _repository, address _address, uint256 _knowledgeID, uint256 _voteDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVotingContract *KNWVotingContractSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeID *big.Int, _voteDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartVote(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeID, _voteDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0xe5023ff2.
//
// Solidity: function startVote(bytes32 _repository, address _address, uint256 _knowledgeID, uint256 _voteDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVotingContract *KNWVotingContractTransactorSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeID *big.Int, _voteDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVotingContract.Contract.StartVote(&_KNWVotingContract.TransactOpts, _repository, _address, _knowledgeID, _voteDuration, _proposersStake, _numberOfKNW)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"sqrt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathFuncSigs maps the 4-byte function signature to its string representation.
var SafeMathFuncSigs = map[string]string{
	"677342ce": "sqrt(uint256)",
}

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x610129610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063677342ce146038575b600080fd5b605260048036036020811015604c57600080fd5b50356064565b60408051918252519081900360200190f35b60008160715750600060ef565b81826001011160bf576040805162461bcd60e51b8152602060048201526015602482015274119b185dd959081a5b9c1d5d08199bdc881cdc5c9d605a1b604482015290519081900360640190fd5b60026001830104825b8082101560eb57508060028180868160dc57fe5b04018160e457fe5b04915060c8565b5090505b91905056fea265627a7a72305820953020e84b9f63c03b694ac9fe34b1131c3ea11c415e206ee9a52f32f26c3bd364736f6c634300050a0032"

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
const DitCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getKNWVoteIDFromProposalID\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getCurrentProposalID\",\"outputs\":[{\"name\":\"currentProposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"openVoteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repositories\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"currentProposalID\",\"type\":\"uint256\"},{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"replaceManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsOfRepository\",\"outputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"},{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"individualStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getVotingMajority\",\"outputs\":[{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getIndividualStake\",\"outputs\":[{\"name\":\"individualStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_identifier\",\"type\":\"string\"},{\"name\":\"_knowledgeID\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"},{\"name\":\"_voteDuration\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[{\"name\":\"proposalID\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"},{\"name\":\"_knowledgeIDs\",\"type\":\"uint256[]\"},{\"name\":\"_votingMajority\",\"type\":\"uint256\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"voteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedKnowledgeIDs\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURNING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"passedKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getKnowledgeIDs\",\"outputs\":[{\"name\":\"knowledgeIDs\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"repositoryIsInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"passKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"}],\"name\":\"migrateRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"},{\"name\":\"_lastDitCoordinator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"}],\"name\":\"InitializeRepository\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"ProposeCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"CommitVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"OpenVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votedRight\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"FinalizeVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"FinalizeProposal\",\"type\":\"event\"}]"

// DitCoordinatorFuncSigs maps the 4-byte function signature to its string representation.
var DitCoordinatorFuncSigs = map[string]string{
	"c814af1f": "BURNING_METHOD()",
	"985dbfc5": "KNWTokenAddress()",
	"eb49fe94": "KNWVotingAddress()",
	"1ecbb9de": "MAX_VOTE_DURATION()",
	"effb21e1": "MINTING_METHOD()",
	"3eedfc10": "MIN_VOTE_DURATION()",
	"d0c397ef": "addKYCValidator(address)",
	"b3070331": "allowedKnowledgeIDs(bytes32,uint256)",
	"2e71d0fb": "finalizeVote(bytes32,uint256)",
	"0bdc90e8": "getCurrentProposalID(bytes32)",
	"552edc9d": "getIndividualStake(bytes32,uint256)",
	"06ee4596": "getKNWVoteIDFromProposalID(bytes32,uint256)",
	"df1bd2cd": "getKnowledgeIDs(bytes32)",
	"3fcc148d": "getVotingMajority(bytes32)",
	"6c63b91e": "initRepository(string,uint256[],uint256)",
	"1341f25c": "isKYCValidator(address)",
	"d7ad0eae": "lastDitCoordinator()",
	"f60055cb": "migrateRepository(string)",
	"99821d9b": "nextDitCoordinator()",
	"0ee62ec0": "openVoteOnProposal(bytes32,uint256,uint256,uint256)",
	"eb931024": "passKYC(address)",
	"ccd9aa68": "passedKYC(address)",
	"87c9203d": "proposalHasPassed(bytes32,uint256)",
	"3e029f63": "proposalsOfRepository(bytes32,uint256)",
	"627b90b3": "proposeCommit(bytes32,string,string,uint256,uint256,uint256)",
	"73b0dddd": "removeKYCValidator(address)",
	"23447982": "replaceManager(address)",
	"1f51fd71": "repositories(bytes32)",
	"ea6c6d0f": "repositoryIsInitialized(bytes32)",
	"39ba645b": "revokeKYC(address)",
	"eb2c0223": "upgradeContract(address)",
	"ab4b593e": "voteOnProposal(bytes32,uint256,bytes32,uint256)",
}

// DitCoordinatorBin is the compiled bytecode used for deploying new contracts.
var DitCoordinatorBin = "0x608060405234801561001057600080fd5b5060405162002c3338038062002c338339818101604052606081101561003557600080fd5b50805160208201516040909201519091906001600160a01b0382161580159061006657506001600160a01b03831615155b6100bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d81526020018062002c06602d913960400191505060405180910390fd5b600080546001600160a01b039384166001600160a01b031991821617808355600580548316918616919091179055600180549585169582169590951780865560068054831691861691909117905560028054939094169281169290921790925533808352600b6020526040909220805460ff191690931790925560048054909216179055612ab680620001506000396000f3fe6080604052600436106101e35760003560e01c806387c9203d11610102578063d7ad0eae11610095578063eb49fe9411610064578063eb49fe94146109fc578063eb93102414610a11578063effb21e114610895578063f60055cb14610a44576101e3565b8063d7ad0eae14610910578063df1bd2cd14610925578063ea6c6d0f1461099f578063eb2c0223146109c9576101e3565b8063b3070331116100d1578063b307033114610865578063c814af1f14610895578063ccd9aa68146108aa578063d0c397ef146108dd576101e3565b806387c9203d146107c0578063985dbfc5146107f057806399821d9b14610821578063ab4b593e14610836576101e3565b806339ba645b1161017a578063552edc9d11610149578063552edc9d146105c1578063627b90b3146105f15780636c63b91e146106c257806373b0dddd1461078d576101e3565b806339ba645b146103ff5780633e029f63146104325780633eedfc10146105825780633fcc148d14610597576101e3565b80631ecbb9de116101b65780631ecbb9de146102d75780631f51fd71146102ec578063234479821461039c5780632e71d0fb146103cf576101e3565b806306ee4596146101e85780630bdc90e81461022a5780630ee62ec0146102545780631341f25c146102a4575b600080fd5b3480156101f457600080fd5b506102186004803603604081101561020b57600080fd5b5080359060200135610abf565b60408051918252519081900360200190f35b34801561023657600080fd5b506102186004803603602081101561024d57600080fd5b5035610adf565b34801561026057600080fd5b506102906004803603608081101561027757600080fd5b5080359060208101359060408101359060600135610af4565b604080519115158252519081900360200190f35b3480156102b057600080fd5b50610290600480360360208110156102c757600080fd5b50356001600160a01b0316610c40565b3480156102e357600080fd5b50610218610c55565b3480156102f857600080fd5b506103166004803603602081101561030f57600080fd5b5035610c5c565b6040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561035f578181015183820152602001610347565b50505050905090810190601f16801561038c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156103a857600080fd5b50610290600480360360208110156103bf57600080fd5b50356001600160a01b0316610d09565b3480156103db57600080fd5b50610290600480360360408110156103f257600080fd5b5080359060200135610d5b565b34801561040b57600080fd5b506102906004803603602081101561042257600080fd5b50356001600160a01b031661118a565b34801561043e57600080fd5b506104626004803603604081101561045557600080fd5b50803590602001356111cf565b60408051908101889052606081018790526001600160a01b038616608082015284151560a082015283151560c082015260e0810183905261010081018290526101208082528a5190820152895181906020808301916101408401918e019080838360005b838110156104de5781810151838201526020016104c6565b50505050905090810190601f16801561050b5780820380516001836020036101000a031916815260200191505b5083810382528b5181528b516020918201918d019080838360005b8381101561053e578181015183820152602001610526565b50505050905090810190601f16801561056b5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b34801561058e57600080fd5b50610218611354565b3480156105a357600080fd5b50610218600480360360208110156105ba57600080fd5b5035611359565b3480156105cd57600080fd5b50610218600480360360408110156105e457600080fd5b508035906020013561136e565b610218600480360360c081101561060757600080fd5b81359190810190604081016020820135600160201b81111561062857600080fd5b82018360208201111561063a57600080fd5b803590602001918460018302840111600160201b8311171561065b57600080fd5b919390929091602081019035600160201b81111561067857600080fd5b82018360208201111561068a57600080fd5b803590602001918460018302840111600160201b831117156106ab57600080fd5b91935091508035906020810135906040013561138e565b3480156106ce57600080fd5b50610290600480360360608110156106e557600080fd5b810190602081018135600160201b8111156106ff57600080fd5b82018360208201111561071157600080fd5b803590602001918460018302840111600160201b8311171561073257600080fd5b919390929091602081019035600160201b81111561074f57600080fd5b82018360208201111561076157600080fd5b803590602001918460208302840111600160201b8311171561078257600080fd5b919350915035611837565b34801561079957600080fd5b50610290600480360360208110156107b057600080fd5b50356001600160a01b0316611cb6565b3480156107cc57600080fd5b50610290600480360360408110156107e357600080fd5b5080359060200135611cfb565b3480156107fc57600080fd5b50610805611d9e565b604080516001600160a01b039092168252519081900360200190f35b34801561082d57600080fd5b50610805611dad565b6102906004803603608081101561084c57600080fd5b5080359060208101359060408101359060600135611dbc565b34801561087157600080fd5b506102906004803603604081101561088857600080fd5b508035906020013561203f565b3480156108a157600080fd5b5061021861205f565b3480156108b657600080fd5b50610290600480360360208110156108cd57600080fd5b50356001600160a01b0316612064565b3480156108e957600080fd5b506102906004803603602081101561090057600080fd5b50356001600160a01b0316612079565b34801561091c57600080fd5b506108056120c2565b34801561093157600080fd5b5061094f6004803603602081101561094857600080fd5b50356120d1565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561098b578181015183820152602001610973565b505050509050019250505060405180910390f35b3480156109ab57600080fd5b50610290600480360360208110156109c257600080fd5b5035612136565b3480156109d557600080fd5b50610290600480360360208110156109ec57600080fd5b50356001600160a01b031661214d565b348015610a0857600080fd5b5061080561219f565b348015610a1d57600080fd5b5061029060048036036020811015610a3457600080fd5b50356001600160a01b03166121ae565b348015610a5057600080fd5b5061029060048036036020811015610a6757600080fd5b810190602081018135600160201b811115610a8157600080fd5b820183602082011115610a9357600080fd5b803590602001918460018302840111600160201b83111715610ab457600080fd5b5090925090506121f7565b600091825260096020908152604080842092845291905290206002015490565b60009081526007602052604090206002015490565b336000818152600a602052604081205490919060ff16610b1357600080fd5b6005546000878152600960209081526040808320898452825280832060020154815163cdd6ceb960e01b81526004810191909152336024820152604481018990526064810188905290516001600160a01b039094169363cdd6ceb993608480840194938390030190829087803b158015610b8c57600080fd5b505af1158015610ba0573d6000803e3d6000fd5b505050506040513d6020811015610bb657600080fd5b5050600086815260096020908152604080832088845282528083203380855260078201845293829020600281018990556003909101549054825191825260018914938201939093528082019290925251879189917f1a245c311cb3eef22175d9c7c5458569e91d9a5956430216eaf8de7c1b7387609181900360600190a450600195945050505050565b600b6020526000908152604090205460ff1681565b6212750081565b60076020908152600091825260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909291839190830182828015610cf35780601f10610cc857610100808354040283529160200191610cf3565b820191906000526020600020905b815481529060010190602001808311610cd657829003601f168201915b5050505050908060020154908060030154905083565b6004546000906001600160a01b03163314610d2357600080fd5b6001600160a01b038216610d3657600080fd5b50600480546001600160a01b0383166001600160a01b03199091161790556001919050565b336000818152600a602052604081205490919060ff16610d7a57600080fd5b6000848152600960209081526040808320868452825280832033845260070190915290206003015460ff1615610de15760405162461bcd60e51b81526004018080602001828103825260278152602001806128a86027913960400191505060405180910390fd5b60008481526009602090815260408083208684528252808320338452600701909152902054151580610e37575060008481526009602090815260408083208684529091529020600401546001600160a01b031633145b610e725760405162461bcd60e51b815260040180806020018281038252603a815260200180612924603a913960400191505060405180910390fd5b6000848152600960209081526040808320868452909152902060040154600160a01b900460ff16610fbe576005546000858152600960209081526040808320878452825280832060020154815163865df0ad60e01b8152600481019190915290516001600160a01b039094169363865df0ad93602480840194938390030190829087803b158015610f0257600080fd5b505af1158015610f16573d6000803e3d6000fd5b505050506040513d6020811015610f2c57600080fd5b50516000858152600960209081526040808320878452825291829020600481018054600160a01b60ff60a81b19909116600160a81b96151587021760ff60a01b19161790819055600390910154835190815260ff94909104939093161515908301528051859287927f7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc192918290030190a35b60055460008581526009602090815260408083208784528252808320600280820154338087526007909301909452828520015482516336bf4c9160e01b8152600481019490945260248401526044830152519192839283926001600160a01b03909216916336bf4c9191606480830192606092919082900301818787803b15801561104857600080fd5b505af115801561105c573d6000803e3d6000fd5b505050506040513d606081101561107257600080fd5b5080516020820151604090920151909450909250905082156110bd57604051339084156108fc029085906000818181858888f193505050501580156110bb573d6000803e3d6000fd5b505b60008781526009602090815260408083208984529091529020600601546110ea908463ffffffff6126d816565b60008881526009602090815260408083208a84528083528184206006810195909555338085526007860184528285206003908101805460ff19166001179055948c905290835292909301548351908152851515918101919091528083018490529151909188918a917f45b8e8084f290362eff12aa8825548cacdd5ce5bbd00dbcaeed1bec7cc8dbcd5919081900360600190a45060019695505050505050565b336000818152600b602052604081205490919060ff166111a957600080fd5b50506001600160a01b03166000908152600a60205260409020805460ff19169055600190565b600960209081526000928352604080842082529183529181902080548251601f60026000196101006001861615020190931692909204918201859004850281018501909352808352909283919083018282801561126d5780601f106112425761010080835404028352916020019161126d565b820191906000526020600020905b81548152906001019060200180831161125057829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561130b5780601f106112e05761010080835404028352916020019161130b565b820191906000526020600020905b8154815290600101906020018083116112ee57829003601f168201915b5050506002840154600385015460048601546005870154600690970154959692959194506001600160a01b038116935060ff600160a01b8204811693600160a81b909204169189565b603c81565b60009081526007602052604090206003015490565b600091825260096020908152604080842092845291905290206005015490565b336000818152600a602052604081205490919060ff166113ad57600080fd5b600034116113ec5760405162461bcd60e51b8152600401808060200182810382526028815260200180612a5a6028913960400191505060405180910390fd5b603c83101580156114005750621275008311155b611449576040805162461bcd60e51b8152602060048201526015602482015274159bdd1948191d5c985d1a5bdb881a5b9d985b1a59605a1b604482015290519081900360640190fd5b871580159061145757508515155b6114925760405162461bcd60e51b815260040180806020018281038252602f815260200180612879602f913960400191505060405180910390fd5b6003546001600160a01b0316156114da5760405162461bcd60e51b81526004018080602001828103825260228152602001806129026022913960400191505060405180910390fd5b60008a815260086020908152604080832088845290915290205460ff16611548576040805162461bcd60e51b815260206004820152601b60248201527f4b6e6f776c65646765204944206973206e6f7420636f72726563740000000000604482015290519081900360640190fd5b60008a81526007602052604081206002015461156b90600163ffffffff61271f16565b60008c8152600760209081526040918290206002018390558151610140601f8e018390049092028101820190925261012082018c8152929350909182918d908d9081908501838280828437600092019190915250505090825250604080516020601f8c018190048102820181019092528a815291810191908b908b9081908401838280828437600081840152601f19601f820116905080830192505050505050508152602001600560009054906101000a90046001600160a01b03166001600160a01b031663e5023ff28e338b8a348d6040518763ffffffff1660e01b815260040180878152602001866001600160a01b03166001600160a01b031681526020018581526020018481526020018381526020018281526020019650505050505050602060405180830381600087803b1580156116a657600080fd5b505af11580156116ba573d6000803e3d6000fd5b505050506040513d60208110156116d057600080fd5b50518152602081810189905233604080840191909152600060608401819052608084018190523460a0850181905260c0909401939093528e835260098252808320858452825290912082518051919261172e9284929091019061276a565b506020828101518051611747926001850192019061276a565b5060408281015160028084019190915560608401516003840155608084015160048401805460a087015160c08801511515600160a81b0260ff60a81b19911515600160a01b0260ff60a01b196001600160a01b039096166001600160a01b03199094169390931794909416919091171691909117905560e084015160058401556101009093015160069092019190915560008d815260076020908152908290209092015481518981529283018890528151339391928f927f603f65981d2b692fdcecdd4b4c97ebe242a11b1feb833e2c556a6a3bc27e4314929081900390910190a49a9950505050505050505050565b336000818152600a602052604081205490919060ff1661185657600080fd5b856118925760405162461bcd60e51b81526004018080602001828103825260248152602001806129c46024913960400191505060405180910390fd5b60008787604051602001808383808284378083019250505092505050604051602081830303815290604052805190602001209050600760008281526020019081526020016000206003015460001461191b5760405162461bcd60e51b8152600401808060200182810382526027815260200180612a096027913960400191505060405180910390fd5b6032841015611971576040805162461bcd60e51b815260206004820152601f60248201527f566f74696e67206d616a6f726974792068617320746f206265203e3d20353000604482015290519081900360640190fd5b846119ad5760405162461bcd60e51b81526004018080602001828103825260218152602001806129e86021913960400191505060405180910390fd5b6003546001600160a01b0316156119f55760405162461bcd60e51b81526004018080602001828103825260228152602001806129026022913960400191505060405180910390fd5b600654604080516001620fe0a360e11b0319815290516000926001600160a01b03169163ffe03eba916004808301926020929190829003018186803b158015611a3d57600080fd5b505afa158015611a51573d6000803e3d6000fd5b505050506040513d6020811015611a6757600080fd5b5051905060005b86811015611b205781888883818110611a8357fe5b9050602002013510611ad3576040805162461bcd60e51b8152602060048201526014602482015273125b9d985b1a590812db9bdddb195919d948125160621b604482015290519081900360640190fd5b60008381526008602052604081206001918a8a85818110611af057fe5b60209081029290920135835250810191909152604001600020805460ff1916911515919091179055600101611a6e565b5060405180608001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506040805160208a810282810182019093528a82529283019290918b918b91829185019084908082843760009201829052509385525050506020808301829052604092830189905285825260078152919020825180519192611bcb9284929091019061276a565b506020828101518051611be492600185019201906127e8565b506040828101516002830155606090920151600390910155600554815163828c165360e01b8152600481018590526024810188905291516001600160a01b039091169163828c16539160448083019260209291908290030181600087803b158015611c4e57600080fd5b505af1158015611c62573d6000803e3d6000fd5b505050506040513d6020811015611c7857600080fd5b5050604051339083907fe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be90600090a350600198975050505050505050565b336000818152600b602052604081205490919060ff16611cd557600080fd5b50506001600160a01b03166000908152600b60205260409020805460ff19169055600190565b6000828152600960209081526040808320848452909152812060040154600160a01b900460ff16611d73576040805162461bcd60e51b815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b506000918252600960209081526040808420928452919052902060040154600160a81b900460ff1690565b6001546001600160a01b031681565b6003546001600160a01b031681565b336000818152600a602052604081205490919060ff16611ddb57600080fd5b60008681526009602090815260408083208884529091529020600501543414611e355760405162461bcd60e51b81526004018080602001828103825260398152602001806128406039913960400191505060405180910390fd5b60008681526009602090815260408083208884529091529020600401546001600160a01b0316331415611e995760405162461bcd60e51b81526004018080602001828103825260318152602001806129936031913960400191505060405180910390fd5b6000868152600960209081526040808320888452909152902060060154611ec6903463ffffffff61271f16565b600087815260096020908152604080832089845282528083206006810194909455600554600290940154815163d4e0ac9560e01b815260048101919091523360248201526044810189905260648101889052905192936001600160a01b03169263d4e0ac959260848084019391929182900301818787803b158015611f4a57600080fd5b505af1158015611f5e573d6000803e3d6000fd5b505050506040513d6020811015611f7457600080fd5b5051905080611fb45760405162461bcd60e51b81526004018080602001828103825260338152602001806128cf6033913960400191505060405180910390fd5b600087815260096020908152604080832089845280835281842033808652600782018552838620879055948b90529083526003015481519081523492810192909252818101879052606082018490525188918a917f2edd4beb5f5626bb825280acc138e2ffc3190dff1e0d8e4563febc24ed3425499181900360800190a45060019695505050505050565b600860209081526000928352604080842090915290825290205460ff1681565b600081565b600a6020526000908152604090205460ff1681565b336000818152600b602052604081205490919060ff1661209857600080fd5b50506001600160a01b03166000908152600b60205260409020805460ff1916600190811790915590565b6002546001600160a01b031681565b60008181526007602090815260409182902060010180548351818402810184019094528084526060939283018282801561212a57602002820191906000526020600020905b815481526020019060010190808311612116575b50505050509050919050565b600090815260076020526040902060030154151590565b6004546000906001600160a01b0316331461216757600080fd5b6001600160a01b03821661217a57600080fd5b50600380546001600160a01b0383166001600160a01b03199091161790556001919050565b6000546001600160a01b031681565b336000818152600b602052604081205490919060ff166121cd57600080fd5b50506001600160a01b03166000908152600a60205260409020805460ff1916600190811790915590565b336000818152600a602052604081205490919060ff1661221657600080fd5b6002546001600160a01b031661222b57600080fd5b6002546001600160a01b0316836122735760405162461bcd60e51b81526004018080602001828103825260248152602001806129c46024913960400191505060405180910390fd5b600085856040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012090506000826001600160a01b0316630bdc90e8836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156122ed57600080fd5b505afa158015612301573d6000803e3d6000fd5b505050506040513d602081101561231757600080fd5b505160408051633fcc148d60e01b81526004810185905290519192506000916001600160a01b03861691633fcc148d916024808301926020929190829003018186803b15801561236657600080fd5b505afa15801561237a573d6000803e3d6000fd5b505050506040513d602081101561239057600080fd5b50516040805163df1bd2cd60e01b81526004810186905290519192506060916001600160a01b0387169163df1bd2cd916024808301926000929190829003018186803b1580156123df57600080fd5b505afa1580156123f3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561241c57600080fd5b810190808051600160201b81111561243357600080fd5b8201602081018481111561244657600080fd5b81518560208202830111600160201b8211171561246257600080fd5b5050600654604080516001620fe0a360e11b031981529051929650600095506001600160a01b03909116935063ffe03eba9250600480820192602092909190829003018186803b1580156124b557600080fd5b505afa1580156124c9573d6000803e3d6000fd5b505050506040513d60208110156124df57600080fd5b5051905060005b825181101561259e57818382815181106124fc57fe5b60200260200101511061254d576040805162461bcd60e51b8152602060048201526014602482015273125b9d985b1a590812db9bdddb195919d948125160621b604482015290519081900360640190fd5b600086815260086020526040812084516001929086908590811061256d57fe5b6020908102919091018101518252810191909152604001600020805460ff19169115159190911790556001016124e6565b5060405180608001604052808b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250938552505050602080830186905260408084018990526060909301879052888252600781529190208251805191926126199284929091019061276a565b50602082810151805161263292600185019201906127e8565b506040828101516002830155606090920151600390910155600554815163828c165360e01b8152600481018890526024810186905291516001600160a01b039091169163828c16539160448083019260209291908290030181600087803b15801561269c57600080fd5b505af11580156126b0573d6000803e3d6000fd5b505050506040513d60208110156126c657600080fd5b5060019b9a5050505050505050505050565b6000828211156127195760405162461bcd60e51b815260040180806020018281038252603581526020018061295e6035913960400191505060405180910390fd5b50900390565b6000828201838110156127635760405162461bcd60e51b815260040180806020018281038252602a815260200180612a30602a913960400191505060405180910390fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106127ab57805160ff19168380011785556127d8565b828001600101855582156127d8579182015b828111156127d85782518255916020019190600101906127bd565b506127e4929150612822565b5090565b8280548282559060005260206000209081019282156127d857916020028201828111156127d85782518255916020019190600101906127bd565b61283c91905b808211156127e45760008155600101612828565b9056fe56616c7565206f6620746865207472616e73616374696f6e20646f65736e2774206d6174636820746865207265717569726564207374616b65546f70696320616e64206964656e746966696572206f662070726f706f73616c2063616e277420626520656d70747945616368207061727469636970616e742063616e206f6e6c792066696e616c697a65206f6e6365566f74696e6720636f6e74726163742072657475726e656420616e20696e76616c696420616d6f756e74206f6620766f74657354686572652069732061206e6577657220636f6e7472616374206465706c6f7965644f6e6c79207061727469636970616e7473206f662074686520766f7465206172652061626c6520746f207265736f6c76652074686520766f746543616e27742073756274726163742061206e756d6265722066726f6d206120736d616c6c6572206f6e6520776974682075696e74735468652070726f706f736572206973206e6f7420616c6c6f77656420746f20766f746520696e20612070726f706f73616c5265706f7369746f72792064657363726970746f722063616e277420626520656d70747950726f76696465206174206c65617374206f6e65206b6e6f776c656467652049445265706f7369746f72792063616e206f6e6c7920626520696e697469616c697a6564206f6e6365526573756c742068617320746f20626520626967676572207468616e20626f74682073756d6d616e647356616c7565206f6620746865207472616e73616374696f6e2063616e206e6f74206265207a65726fa265627a7a723058201f80b32a2d32cd86104425e8d66c375a22be670e06d5c57e2e780a3d12afecf964736f6c634300050a00324b4e57566f74696e6720616e64204b4e57546f6b656e20616464726573732063616e277420626520656d707479"

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

// AllowedKnowledgeIDs is a free data retrieval call binding the contract method 0xb3070331.
//
// Solidity: function allowedKnowledgeIDs(bytes32 , uint256 ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCaller) AllowedKnowledgeIDs(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "allowedKnowledgeIDs", arg0, arg1)
	return *ret0, err
}

// AllowedKnowledgeIDs is a free data retrieval call binding the contract method 0xb3070331.
//
// Solidity: function allowedKnowledgeIDs(bytes32 , uint256 ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorSession) AllowedKnowledgeIDs(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _DitCoordinator.Contract.AllowedKnowledgeIDs(&_DitCoordinator.CallOpts, arg0, arg1)
}

// AllowedKnowledgeIDs is a free data retrieval call binding the contract method 0xb3070331.
//
// Solidity: function allowedKnowledgeIDs(bytes32 , uint256 ) constant returns(bool)
func (_DitCoordinator *DitCoordinatorCallerSession) AllowedKnowledgeIDs(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _DitCoordinator.Contract.AllowedKnowledgeIDs(&_DitCoordinator.CallOpts, arg0, arg1)
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

// GetKnowledgeIDs is a free data retrieval call binding the contract method 0xdf1bd2cd.
//
// Solidity: function getKnowledgeIDs(bytes32 _repository) constant returns(uint256[] knowledgeIDs)
func (_DitCoordinator *DitCoordinatorCaller) GetKnowledgeIDs(opts *bind.CallOpts, _repository [32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DitCoordinator.contract.Call(opts, out, "getKnowledgeIDs", _repository)
	return *ret0, err
}

// GetKnowledgeIDs is a free data retrieval call binding the contract method 0xdf1bd2cd.
//
// Solidity: function getKnowledgeIDs(bytes32 _repository) constant returns(uint256[] knowledgeIDs)
func (_DitCoordinator *DitCoordinatorSession) GetKnowledgeIDs(_repository [32]byte) ([]*big.Int, error) {
	return _DitCoordinator.Contract.GetKnowledgeIDs(&_DitCoordinator.CallOpts, _repository)
}

// GetKnowledgeIDs is a free data retrieval call binding the contract method 0xdf1bd2cd.
//
// Solidity: function getKnowledgeIDs(bytes32 _repository) constant returns(uint256[] knowledgeIDs)
func (_DitCoordinator *DitCoordinatorCallerSession) GetKnowledgeIDs(_repository [32]byte) ([]*big.Int, error) {
	return _DitCoordinator.Contract.GetKnowledgeIDs(&_DitCoordinator.CallOpts, _repository)
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(string description, string identifier, uint256 KNWVoteID, uint256 knowledgeID, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitCoordinator *DitCoordinatorCaller) ProposalsOfRepository(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	Description      string
	Identifier       string
	KNWVoteID        *big.Int
	KnowledgeID      *big.Int
	Proposer         common.Address
	IsFinalized      bool
	ProposalAccepted bool
	IndividualStake  *big.Int
	TotalStake       *big.Int
}, error) {
	ret := new(struct {
		Description      string
		Identifier       string
		KNWVoteID        *big.Int
		KnowledgeID      *big.Int
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(string description, string identifier, uint256 KNWVoteID, uint256 knowledgeID, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitCoordinator *DitCoordinatorSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	Description      string
	Identifier       string
	KNWVoteID        *big.Int
	KnowledgeID      *big.Int
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(string description, string identifier, uint256 KNWVoteID, uint256 knowledgeID, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitCoordinator *DitCoordinatorCallerSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	Description      string
	Identifier       string
	KNWVoteID        *big.Int
	KnowledgeID      *big.Int
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

// InitRepository is a paid mutator transaction binding the contract method 0x6c63b91e.
//
// Solidity: function initRepository(string _repository, uint256[] _knowledgeIDs, uint256 _votingMajority) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository string, _knowledgeIDs []*big.Int, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "initRepository", _repository, _knowledgeIDs, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x6c63b91e.
//
// Solidity: function initRepository(string _repository, uint256[] _knowledgeIDs, uint256 _votingMajority) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) InitRepository(_repository string, _knowledgeIDs []*big.Int, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _knowledgeIDs, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x6c63b91e.
//
// Solidity: function initRepository(string _repository, uint256[] _knowledgeIDs, uint256 _votingMajority) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) InitRepository(_repository string, _knowledgeIDs []*big.Int, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.InitRepository(&_DitCoordinator.TransactOpts, _repository, _knowledgeIDs, _votingMajority)
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

// ProposeCommit is a paid mutator transaction binding the contract method 0x627b90b3.
//
// Solidity: function proposeCommit(bytes32 _repository, string _description, string _identifier, uint256 _knowledgeID, uint256 _numberOfKNW, uint256 _voteDuration) returns(uint256 proposalID)
func (_DitCoordinator *DitCoordinatorTransactor) ProposeCommit(opts *bind.TransactOpts, _repository [32]byte, _description string, _identifier string, _knowledgeID *big.Int, _numberOfKNW *big.Int, _voteDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "proposeCommit", _repository, _description, _identifier, _knowledgeID, _numberOfKNW, _voteDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x627b90b3.
//
// Solidity: function proposeCommit(bytes32 _repository, string _description, string _identifier, uint256 _knowledgeID, uint256 _numberOfKNW, uint256 _voteDuration) returns(uint256 proposalID)
func (_DitCoordinator *DitCoordinatorSession) ProposeCommit(_repository [32]byte, _description string, _identifier string, _knowledgeID *big.Int, _numberOfKNW *big.Int, _voteDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _description, _identifier, _knowledgeID, _numberOfKNW, _voteDuration)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0x627b90b3.
//
// Solidity: function proposeCommit(bytes32 _repository, string _description, string _identifier, uint256 _knowledgeID, uint256 _numberOfKNW, uint256 _voteDuration) returns(uint256 proposalID)
func (_DitCoordinator *DitCoordinatorTransactorSession) ProposeCommit(_repository [32]byte, _description string, _identifier string, _knowledgeID *big.Int, _numberOfKNW *big.Int, _voteDuration *big.Int) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ProposeCommit(&_DitCoordinator.TransactOpts, _repository, _description, _identifier, _knowledgeID, _numberOfKNW, _voteDuration)
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

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactor) ReplaceManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _DitCoordinator.contract.Transact(opts, "replaceManager", _newManager)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool)
func (_DitCoordinator *DitCoordinatorSession) ReplaceManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ReplaceManager(&_DitCoordinator.TransactOpts, _newManager)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool)
func (_DitCoordinator *DitCoordinatorTransactorSession) ReplaceManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitCoordinator.Contract.ReplaceManager(&_DitCoordinator.TransactOpts, _newManager)
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
	KnowledgeID   *big.Int
	Stake         *big.Int
	NumberOfKNW   *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitVote is a free log retrieval operation binding the contract event 0x2edd4beb5f5626bb825280acc138e2ffc3190dff1e0d8e4563febc24ed342549.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
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

// WatchCommitVote is a free log subscription operation binding the contract event 0x2edd4beb5f5626bb825280acc138e2ffc3190dff1e0d8e4563febc24ed342549.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
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

// ParseCommitVote is a log parse operation binding the contract event 0x2edd4beb5f5626bb825280acc138e2ffc3190dff1e0d8e4563febc24ed342549.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
func (_DitCoordinator *DitCoordinatorFilterer) ParseCommitVote(log types.Log) (*DitCoordinatorCommitVote, error) {
	event := new(DitCoordinatorCommitVote)
	if err := _DitCoordinator.contract.UnpackLog(event, "CommitVote", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Repository  [32]byte
	Proposal    *big.Int
	KnowledgeID *big.Int
	Accepted    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizeProposal is a free log retrieval operation binding the contract event 0x7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc1.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, uint256 knowledgeID, bool accepted)
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

// WatchFinalizeProposal is a free log subscription operation binding the contract event 0x7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc1.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, uint256 knowledgeID, bool accepted)
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

// ParseFinalizeProposal is a log parse operation binding the contract event 0x7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc1.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, uint256 knowledgeID, bool accepted)
func (_DitCoordinator *DitCoordinatorFilterer) ParseFinalizeProposal(log types.Log) (*DitCoordinatorFinalizeProposal, error) {
	event := new(DitCoordinatorFinalizeProposal)
	if err := _DitCoordinator.contract.UnpackLog(event, "FinalizeProposal", log); err != nil {
		return nil, err
	}
	return event, nil
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
	KnowledgeID *big.Int
	VotedRight  bool
	NumberOfKNW *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizeVote is a free log retrieval operation binding the contract event 0x45b8e8084f290362eff12aa8825548cacdd5ce5bbd00dbcaeed1bec7cc8dbcd5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool votedRight, uint256 numberOfKNW)
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

// WatchFinalizeVote is a free log subscription operation binding the contract event 0x45b8e8084f290362eff12aa8825548cacdd5ce5bbd00dbcaeed1bec7cc8dbcd5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool votedRight, uint256 numberOfKNW)
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

// ParseFinalizeVote is a log parse operation binding the contract event 0x45b8e8084f290362eff12aa8825548cacdd5ce5bbd00dbcaeed1bec7cc8dbcd5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool votedRight, uint256 numberOfKNW)
func (_DitCoordinator *DitCoordinatorFilterer) ParseFinalizeVote(log types.Log) (*DitCoordinatorFinalizeVote, error) {
	event := new(DitCoordinatorFinalizeVote)
	if err := _DitCoordinator.contract.UnpackLog(event, "FinalizeVote", log); err != nil {
		return nil, err
	}
	return event, nil
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

// ParseInitializeRepository is a log parse operation binding the contract event 0xe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be.
//
// Solidity: event InitializeRepository(bytes32 indexed repository, address indexed who)
func (_DitCoordinator *DitCoordinatorFilterer) ParseInitializeRepository(log types.Log) (*DitCoordinatorInitializeRepository, error) {
	event := new(DitCoordinatorInitializeRepository)
	if err := _DitCoordinator.contract.UnpackLog(event, "InitializeRepository", log); err != nil {
		return nil, err
	}
	return event, nil
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
	KnowledgeID   *big.Int
	Accept        bool
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenVote is a free log retrieval operation binding the contract event 0x1a245c311cb3eef22175d9c7c5458569e91d9a5956430216eaf8de7c1b738760.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool accept, uint256 numberOfVotes)
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

// WatchOpenVote is a free log subscription operation binding the contract event 0x1a245c311cb3eef22175d9c7c5458569e91d9a5956430216eaf8de7c1b738760.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool accept, uint256 numberOfVotes)
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

// ParseOpenVote is a log parse operation binding the contract event 0x1a245c311cb3eef22175d9c7c5458569e91d9a5956430216eaf8de7c1b738760.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool accept, uint256 numberOfVotes)
func (_DitCoordinator *DitCoordinatorFilterer) ParseOpenVote(log types.Log) (*DitCoordinatorOpenVote, error) {
	event := new(DitCoordinatorOpenVote)
	if err := _DitCoordinator.contract.UnpackLog(event, "OpenVote", log); err != nil {
		return nil, err
	}
	return event, nil
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
	KnowledgeID *big.Int
	NumberOfKNW *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposeCommit is a free log retrieval operation binding the contract event 0x603f65981d2b692fdcecdd4b4c97ebe242a11b1feb833e2c556a6a3bc27e4314.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 numberOfKNW)
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

// WatchProposeCommit is a free log subscription operation binding the contract event 0x603f65981d2b692fdcecdd4b4c97ebe242a11b1feb833e2c556a6a3bc27e4314.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 numberOfKNW)
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

// ParseProposeCommit is a log parse operation binding the contract event 0x603f65981d2b692fdcecdd4b4c97ebe242a11b1feb833e2c556a6a3bc27e4314.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 numberOfKNW)
func (_DitCoordinator *DitCoordinatorFilterer) ParseProposeCommit(log types.Log) (*DitCoordinatorProposeCommit, error) {
	event := new(DitCoordinatorProposeCommit)
	if err := _DitCoordinator.contract.UnpackLog(event, "ProposeCommit", log); err != nil {
		return nil, err
	}
	return event, nil
}
