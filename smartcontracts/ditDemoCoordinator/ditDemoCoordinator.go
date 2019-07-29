// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ditDemoCoordinator

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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
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
var SafeMathBin = "0x610129610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063677342ce146038575b600080fd5b605260048036036020811015604c57600080fd5b50356064565b60408051918252519081900360200190f35b60008160715750600060ef565b81826001011160bf576040805162461bcd60e51b8152602060048201526015602482015274119b185dd959081a5b9c1d5d08199bdc881cdc5c9d605a1b604482015290519081900360640190fd5b60026001830104825b8082101560eb57508060028180868160dc57fe5b04018160e457fe5b04915060c8565b5090505b91905056fea265627a7a723058201ead97bc14077cfbde12cbe4476a74a713b639869859a199f04661d43c80b69264736f6c634300050a0032"

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

// DitDemoCoordinatorABI is the input ABI used to generate the binding from.
const DitDemoCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getKNWVoteIDFromProposalID\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getCurrentProposalID\",\"outputs\":[{\"name\":\"currentProposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"openVoteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repositories\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"currentProposalID\",\"type\":\"uint256\"},{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"replaceManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsOfRepository\",\"outputs\":[{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"identifier\",\"type\":\"string\"},{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"individualStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getVotingMajority\",\"outputs\":[{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getIndividualStake\",\"outputs\":[{\"name\":\"individualStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"},{\"name\":\"_knowledgeIDs\",\"type\":\"uint256[]\"},{\"name\":\"_votingMajority\",\"type\":\"uint256\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"voteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedKnowledgeIDs\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURNING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"passedKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xDitTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getKnowledgeIDs\",\"outputs\":[{\"name\":\"knowledgeIDs\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"repositoryIsInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"passKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"string\"}],\"name\":\"migrateRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_identifier\",\"type\":\"string\"},{\"name\":\"_knowledgeID\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"},{\"name\":\"_voteDuration\",\"type\":\"uint256\"},{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[{\"name\":\"proposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"},{\"name\":\"_lastDitCoordinator\",\"type\":\"address\"},{\"name\":\"_xDitTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"}],\"name\":\"InitializeRepository\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"ProposeCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"CommitVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"OpenVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"votedRight\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"FinalizeVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"FinalizeProposal\",\"type\":\"event\"}]"

// DitDemoCoordinatorFuncSigs maps the 4-byte function signature to its string representation.
var DitDemoCoordinatorFuncSigs = map[string]string{
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
	"481c6a75": "manager()",
	"f60055cb": "migrateRepository(string)",
	"99821d9b": "nextDitCoordinator()",
	"0ee62ec0": "openVoteOnProposal(bytes32,uint256,uint256,uint256)",
	"eb931024": "passKYC(address)",
	"ccd9aa68": "passedKYC(address)",
	"87c9203d": "proposalHasPassed(bytes32,uint256)",
	"3e029f63": "proposalsOfRepository(bytes32,uint256)",
	"fb4ede85": "proposeCommit(bytes32,string,string,uint256,uint256,uint256,uint256)",
	"73b0dddd": "removeKYCValidator(address)",
	"23447982": "replaceManager(address)",
	"1f51fd71": "repositories(bytes32)",
	"ea6c6d0f": "repositoryIsInitialized(bytes32)",
	"39ba645b": "revokeKYC(address)",
	"eb2c0223": "upgradeContract(address)",
	"ab4b593e": "voteOnProposal(bytes32,uint256,bytes32,uint256)",
	"ce83732e": "xDitTokenAddress()",
}

// DitDemoCoordinatorBin is the compiled bytecode used for deploying new contracts.
var DitDemoCoordinatorBin = "0x608060405234801561001057600080fd5b5060405162002a7f38038062002a7f8339818101604052608081101561003557600080fd5b50805160208201516040830151606090930151919290916001600160a01b0383161580159061006c57506001600160a01b03841615155b6100c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d81526020018062002a52602d913960400191505060405180910390fd5b600080546001600160a01b039485166001600160a01b031991821617808355600680548316918716919091179055600180549686169682169690961786556002805493861693821693909317928390556007805482169386169390931790925560038054939094169282169290921790925533808252600c6020526040909120805460ff19169093179092556005805490911690911790556128e8806200016a6000396000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c8063985dbfc51161011a578063d7ad0eae116100ad578063eb49fe941161007c578063eb49fe941461080c578063eb93102414610814578063effb21e1146106f8578063f60055cb1461083a578063fb4ede85146108a857610206565b8063d7ad0eae14610754578063df1bd2cd1461075c578063ea6c6d0f146107c9578063eb2c0223146107e657610206565b8063c814af1f116100e9578063c814af1f146106f8578063ccd9aa6814610700578063ce83732e14610726578063d0c397ef1461072e57610206565b8063985dbfc51461069657806399821d9b1461069e578063ab4b593e146106a6578063b3070331146106d557610206565b806339ba645b1161019d578063481c6a751161016c578063481c6a7514610548578063552edc9d1461056c5780636c63b91e1461058f57806373b0dddd1461064d57806387c9203d1461067357610206565b806339ba645b146103ba5780633e029f63146103e05780633eedfc10146105235780633fcc148d1461052b57610206565b80631ecbb9de116101d95780631ecbb9de146102c65780631f51fd71146102ce57806323447982146103715780632e71d0fb1461039757610206565b806306ee45961461020b5780630bdc90e8146102405780630ee62ec01461025d5780631341f25c146102a0575b600080fd5b61022e6004803603604081101561022157600080fd5b508035906020013561097f565b60408051918252519081900360200190f35b61022e6004803603602081101561025657600080fd5b503561099f565b61028c6004803603608081101561027357600080fd5b50803590602081013590604081013590606001356109b4565b604080519115158252519081900360200190f35b61028c600480360360208110156102b657600080fd5b50356001600160a01b0316610b00565b61022e610b15565b6102eb600480360360208110156102e457600080fd5b5035610b1c565b6040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561033457818101518382015260200161031c565b50505050905090810190601f1680156103615780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61028c6004803603602081101561038757600080fd5b50356001600160a01b0316610bc9565b61028c600480360360408110156103ad57600080fd5b5080359060200135610c1b565b61028c600480360360208110156103d057600080fd5b50356001600160a01b03166110a4565b610403600480360360408110156103f657600080fd5b50803590602001356110e9565b60408051908101889052606081018790526001600160a01b038616608082015284151560a082015283151560c082015260e0810183905261010081018290526101208082528a5190820152895181906020808301916101408401918e019080838360005b8381101561047f578181015183820152602001610467565b50505050905090810190601f1680156104ac5780820380516001836020036101000a031916815260200191505b5083810382528b5181528b516020918201918d019080838360005b838110156104df5781810151838201526020016104c7565b50505050905090810190601f16801561050c5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b61022e61126e565b61022e6004803603602081101561054157600080fd5b5035611273565b610550611288565b604080516001600160a01b039092168252519081900360200190f35b61022e6004803603604081101561058257600080fd5b5080359060200135611297565b61028c600480360360608110156105a557600080fd5b810190602081018135600160201b8111156105bf57600080fd5b8201836020820111156105d157600080fd5b803590602001918460018302840111600160201b831117156105f257600080fd5b919390929091602081019035600160201b81111561060f57600080fd5b82018360208201111561062157600080fd5b803590602001918460208302840111600160201b8311171561064257600080fd5b9193509150356112b7565b61028c6004803603602081101561066357600080fd5b50356001600160a01b0316611662565b61028c6004803603604081101561068957600080fd5b50803590602001356116a7565b61055061174a565b610550611759565b61028c600480360360808110156106bc57600080fd5b5080359060208101359060408101359060600135611768565b61028c600480360360408110156106eb57600080fd5b5080359060200135611a3e565b61022e611a5e565b61028c6004803603602081101561071657600080fd5b50356001600160a01b0316611a63565b610550611a78565b61028c6004803603602081101561074457600080fd5b50356001600160a01b0316611a87565b610550611ad0565b6107796004803603602081101561077257600080fd5b5035611adf565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156107b557818101518382015260200161079d565b505050509050019250505060405180910390f35b61028c600480360360208110156107df57600080fd5b5035611b44565b61028c600480360360208110156107fc57600080fd5b50356001600160a01b0316611b5b565b610550611bad565b61028c6004803603602081101561082a57600080fd5b50356001600160a01b0316611bbc565b61028c6004803603602081101561085057600080fd5b810190602081018135600160201b81111561086a57600080fd5b82018360208201111561087c57600080fd5b803590602001918460018302840111600160201b8311171561089d57600080fd5b509092509050611c05565b61022e600480360360e08110156108be57600080fd5b81359190810190604081016020820135600160201b8111156108df57600080fd5b8201836020820111156108f157600080fd5b803590602001918460018302840111600160201b8311171561091257600080fd5b919390929091602081019035600160201b81111561092f57600080fd5b82018360208201111561094157600080fd5b803590602001918460018302840111600160201b8311171561096257600080fd5b91935091508035906020810135906040810135906060013561200f565b6000918252600a6020908152604080842092845291905290206002015490565b60009081526008602052604090206002015490565b336000818152600b602052604081205490919060ff166109d357600080fd5b6006546000878152600a60209081526040808320898452825280832060020154815163cdd6ceb960e01b81526004810191909152336024820152604481018990526064810188905290516001600160a01b039094169363cdd6ceb993608480840194938390030190829087803b158015610a4c57600080fd5b505af1158015610a60573d6000803e3d6000fd5b505050506040513d6020811015610a7657600080fd5b50506000868152600a6020908152604080832088845282528083203380855260078201845293829020600281018990556003909101549054825191825260018914938201939093528082019290925251879189917f1a245c311cb3eef22175d9c7c5458569e91d9a5956430216eaf8de7c1b7387609181900360600190a450600195945050505050565b600c6020526000908152604090205460ff1681565b62093a8081565b60086020908152600091825260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909291839190830182828015610bb35780601f10610b8857610100808354040283529160200191610bb3565b820191906000526020600020905b815481529060010190602001808311610b9657829003601f168201915b5050505050908060020154908060030154905083565b6005546000906001600160a01b03163314610be357600080fd5b6001600160a01b038216610bf657600080fd5b50600580546001600160a01b0383166001600160a01b03199091161790556001919050565b336000818152600b602052604081205490919060ff16610c3a57600080fd5b6000848152600a60209081526040808320868452825280832033845260070190915290206003015460ff1615610ca15760405162461bcd60e51b81526004018080602001828103825260278152602001806126da6027913960400191505060405180910390fd5b6000848152600a602090815260408083208684528252808320338452600701909152902054151580610cf757506000848152600a602090815260408083208684529091529020600401546001600160a01b031633145b610d325760405162461bcd60e51b815260040180806020018281038252603a815260200180612756603a913960400191505060405180910390fd5b6000848152600a60209081526040808320868452909152902060040154600160a01b900460ff16610e7e576006546000858152600a60209081526040808320878452825280832060020154815163865df0ad60e01b8152600481019190915290516001600160a01b039094169363865df0ad93602480840194938390030190829087803b158015610dc257600080fd5b505af1158015610dd6573d6000803e3d6000fd5b505050506040513d6020811015610dec57600080fd5b50516000858152600a60209081526040808320878452825291829020600481018054600160a01b60ff60a81b19909116600160a81b96151587021760ff60a01b19161790819055600390910154835190815260ff94909104939093161515908301528051859287927f7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc192918290030190a35b6006546000858152600a602090815260408083208784528252808320600280820154338087526007909301909452828520015482516336bf4c9160e01b8152600481019490945260248401526044830152519192839283926001600160a01b03909216916336bf4c9191606480830192606092919082900301818787803b158015610f0857600080fd5b505af1158015610f1c573d6000803e3d6000fd5b505050506040513d6060811015610f3257600080fd5b508051602082015160409092015190945090925090508215610fd7576007546040805163a9059cbb60e01b81523360048201526024810186905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b158015610fa257600080fd5b505af1158015610fb6573d6000803e3d6000fd5b505050506040513d6020811015610fcc57600080fd5b5051610fd757600080fd5b6000878152600a60209081526040808320898452909152902060060154611004908463ffffffff61254316565b6000888152600a602090815260408083208a84528083528184206006810195909555338085526007860184528285206003908101805460ff19166001179055948c905290835292909301548351908152851515918101919091528083018490529151909188918a917f45b8e8084f290362eff12aa8825548cacdd5ce5bbd00dbcaeed1bec7cc8dbcd5919081900360600190a45060019695505050505050565b336000818152600c602052604081205490919060ff166110c357600080fd5b50506001600160a01b03166000908152600b60205260409020805460ff19169055600190565b600a60209081526000928352604080842082529183529181902080548251601f6002600019610100600186161502019093169290920491820185900485028101850190935280835290928391908301828280156111875780601f1061115c57610100808354040283529160200191611187565b820191906000526020600020905b81548152906001019060200180831161116a57829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112255780601f106111fa57610100808354040283529160200191611225565b820191906000526020600020905b81548152906001019060200180831161120857829003601f168201915b5050506002840154600385015460048601546005870154600690970154959692959194506001600160a01b038116935060ff600160a01b8204811693600160a81b909204169189565b603c81565b60009081526008602052604090206003015490565b6005546001600160a01b031681565b6000918252600a6020908152604080842092845291905290206005015490565b336000818152600b602052604081205490919060ff166112d657600080fd5b856113125760405162461bcd60e51b81526004018080602001828103825260248152602001806127f66024913960400191505060405180910390fd5b60008787604051602001808383808284378083019250505092505050604051602081830303815290604052805190602001209050600860008281526020019081526020016000206003015460001461139b5760405162461bcd60e51b815260040180806020018281038252602781526020018061283b6027913960400191505060405180910390fd5b60328410156113f1576040805162461bcd60e51b815260206004820152601f60248201527f566f74696e67206d616a6f726974792068617320746f206265203e3d20353000604482015290519081900360640190fd5b8461142d5760405162461bcd60e51b815260040180806020018281038252602181526020018061281a6021913960400191505060405180910390fd5b6004546001600160a01b0316156114755760405162461bcd60e51b81526004018080602001828103825260228152602001806127346022913960400191505060405180910390fd5b60005b858110156114cd57600082815260096020526040812060019189898581811061149d57fe5b60209081029290920135835250810191909152604001600020805460ff1916911515919091179055600101611478565b50604051806080016040528089898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060408051602089810282810182019093528982529283019290918a918a91829185019084908082843760009201829052509385525050506020808301829052604092830188905284825260088152919020825180519192611578928492909101906125d5565b5060208281015180516115919260018501920190612653565b506040828101516002830155606090920151600390910155600654815163828c165360e01b8152600481018490526024810187905291516001600160a01b039091169163828c16539160448083019260209291908290030181600087803b1580156115fb57600080fd5b505af115801561160f573d6000803e3d6000fd5b505050506040513d602081101561162557600080fd5b5050604051339082907fe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be90600090a3506001979650505050505050565b336000818152600c602052604081205490919060ff1661168157600080fd5b50506001600160a01b03166000908152600c60205260409020805460ff19169055600190565b6000828152600a60209081526040808320848452909152812060040154600160a01b900460ff1661171f576040805162461bcd60e51b815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b506000918252600a60209081526040808420928452919052902060040154600160a81b900460ff1690565b6001546001600160a01b031681565b6004546001600160a01b031681565b336000818152600b602052604081205490919060ff1661178757600080fd5b6000868152600a602090815260408083208884529091529020600401546001600160a01b03163314156117eb5760405162461bcd60e51b81526004018080602001828103825260318152602001806127c56031913960400191505060405180910390fd5b6007546000878152600a6020908152604080832089845282528083206005015481516323b872dd60e01b8152336004820152306024820152604481019190915290516001600160a01b03909416936323b872dd93606480840194938390030190829087803b15801561185c57600080fd5b505af1158015611870573d6000803e3d6000fd5b505050506040513d602081101561188657600080fd5b505161189157600080fd5b6000868152600a60209081526040808320888452909152902060058101546006909101546118be9161258a565b6000878152600a6020908152604080832089845282528083206006808201959095559354600290940154815163d4e0ac9560e01b815260048101919091523360248201526044810189905260648101889052905192936001600160a01b03169263d4e0ac959260848084019391929182900301818787803b15801561194257600080fd5b505af1158015611956573d6000803e3d6000fd5b505050506040513d602081101561196c57600080fd5b50519050806119ac5760405162461bcd60e51b81526004018080602001828103825260338152602001806127016033913960400191505060405180910390fd5b6000878152600a60209081526040808320898452808352818420338086526007820185528386208790556003820154958c9052918452600501548251948552928401929092528281018790526060830184905251909188918a917f2edd4beb5f5626bb825280acc138e2ffc3190dff1e0d8e4563febc24ed342549919081900360800190a45060019695505050505050565b600960209081526000928352604080842090915290825290205460ff1681565b600081565b600b6020526000908152604090205460ff1681565b6002546001600160a01b031681565b336000818152600c602052604081205490919060ff16611aa657600080fd5b50506001600160a01b03166000908152600c60205260409020805460ff1916600190811790915590565b6003546001600160a01b031681565b600081815260086020908152604091829020600101805483518184028101840190945280845260609392830182828015611b3857602002820191906000526020600020905b815481526020019060010190808311611b24575b50505050509050919050565b600090815260086020526040902060030154151590565b6005546000906001600160a01b03163314611b7557600080fd5b6001600160a01b038216611b8857600080fd5b50600480546001600160a01b0383166001600160a01b03199091161790556001919050565b6000546001600160a01b031681565b336000818152600c602052604081205490919060ff16611bdb57600080fd5b50506001600160a01b03166000908152600b60205260409020805460ff1916600190811790915590565b336000818152600b602052604081205490919060ff16611c2457600080fd5b6003546001600160a01b0316611c3957600080fd5b6003546001600160a01b031683611c815760405162461bcd60e51b81526004018080602001828103825260248152602001806127f66024913960400191505060405180910390fd5b600085856040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012090506000826001600160a01b0316630bdc90e8836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611cfb57600080fd5b505afa158015611d0f573d6000803e3d6000fd5b505050506040513d6020811015611d2557600080fd5b505160408051633fcc148d60e01b81526004810185905290519192506000916001600160a01b03861691633fcc148d916024808301926020929190829003018186803b158015611d7457600080fd5b505afa158015611d88573d6000803e3d6000fd5b505050506040513d6020811015611d9e57600080fd5b50516040805163df1bd2cd60e01b81526004810186905290519192506060916001600160a01b0387169163df1bd2cd916024808301926000929190829003018186803b158015611ded57600080fd5b505afa158015611e01573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611e2a57600080fd5b810190808051600160201b811115611e4157600080fd5b82016020810184811115611e5457600080fd5b81518560208202830111600160201b82111715611e7057600080fd5b50909450600093505050505b8151811015611ed6576000858152600960205260408120835160019290859085908110611ea557fe5b6020908102919091018101518252810191909152604001600020805460ff1916911515919091179055600101611e7c565b5060405180608001604052808a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250505060208083018590526040808401889052606090930186905287825260088152919020825180519192611f51928492909101906125d5565b506020828101518051611f6a9260018501920190612653565b506040828101516002830155606090920151600390910155600654815163828c165360e01b8152600481018790526024810185905291516001600160a01b039091169163828c16539160448083019260209291908290030181600087803b158015611fd457600080fd5b505af1158015611fe8573d6000803e3d6000fd5b505050506040513d6020811015611ffe57600080fd5b5060019a9950505050505050505050565b336000818152600b602052604081205490919060ff1661202e57600080fd5b6000831161206d5760405162461bcd60e51b815260040180806020018281038252602881526020018061288c6028913960400191505060405180910390fd5b603c8410158015612081575062093a808411155b6120ca576040805162461bcd60e51b8152602060048201526015602482015274159bdd1948191d5c985d1a5bdb881a5b9d985b1a59605a1b604482015290519081900360640190fd5b88158015906120d857508615155b6121135760405162461bcd60e51b815260040180806020018281038252602f8152602001806126ab602f913960400191505060405180910390fd5b6004546001600160a01b03161561215b5760405162461bcd60e51b81526004018080602001828103825260228152602001806127346022913960400191505060405180910390fd5b60008b815260096020908152604080832089845290915290205460ff166121c9576040805162461bcd60e51b815260206004820152601b60248201527f4b6e6f776c65646765204944206973206e6f7420636f72726563740000000000604482015290519081900360640190fd5b60008b8152600860205260408120600201546121ec90600163ffffffff61258a16565b60008d8152600860209081526040918290206002018390558151610140601f8f018390049092028101820190925261012082018d8152929350909182918e908e9081908501838280828437600092019190915250505090825250604080516020601f8d018190048102820181019092528b815291810191908c908c9081908401838280828437600081840152601f19601f820116905080830192505050505050508152602001600660009054906101000a90046001600160a01b03166001600160a01b031663e5023ff28f338c8b8b8e6040518763ffffffff1660e01b815260040180878152602001866001600160a01b03166001600160a01b031681526020018581526020018481526020018381526020018281526020019650505050505050602060405180830381600087803b15801561232757600080fd5b505af115801561233b573d6000803e3d6000fd5b505050506040513d602081101561235157600080fd5b5051815260208181018a9052336040808401919091526000606084018190526080840181905260a0840189905260c09093018890528f8352600a825280832085845282529091208251805191926123ad928492909101906125d5565b5060208281015180516123c692600185019201906125d5565b5060408281015160028301556060830151600383015560808301516004808401805460a087015160c08801516001600160a01b03199092166001600160a01b039586161760ff60a01b1916600160a01b911515919091021760ff60a81b1916600160a81b9115159190910217905560e085015160058501556101009094015160069093019290925560075481516323b872dd60e01b815233948101949094523060248501526044840188905290519116916323b872dd9160648083019260209291908290030181600087803b15801561249e57600080fd5b505af11580156124b2573d6000803e3d6000fd5b505050506040513d60208110156124c857600080fd5b50516124d357600080fd5b336001600160a01b0316600860008e8152602001908152602001600020600201548d7f603f65981d2b692fdcecdd4b4c97ebe242a11b1feb833e2c556a6a3bc27e43148a8a604051808381526020018281526020019250505060405180910390a49b9a5050505050505050505050565b6000828211156125845760405162461bcd60e51b81526004018080602001828103825260358152602001806127906035913960400191505060405180910390fd5b50900390565b6000828201838110156125ce5760405162461bcd60e51b815260040180806020018281038252602a815260200180612862602a913960400191505060405180910390fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061261657805160ff1916838001178555612643565b82800160010185558215612643579182015b82811115612643578251825591602001919060010190612628565b5061264f92915061268d565b5090565b8280548282559060005260206000209081019282156126435791602002820182811115612643578251825591602001919060010190612628565b6126a791905b8082111561264f5760008155600101612693565b9056fe546f70696320616e64206964656e746966696572206f662070726f706f73616c2063616e277420626520656d70747945616368207061727469636970616e742063616e206f6e6c792066696e616c697a65206f6e6365566f74696e6720636f6e74726163742072657475726e656420616e20696e76616c696420616d6f756e74206f6620766f74657354686572652069732061206e6577657220636f6e7472616374206465706c6f7965644f6e6c79207061727469636970616e7473206f662074686520766f7465206172652061626c6520746f207265736f6c76652074686520766f746543616e27742073756274726163742061206e756d6265722066726f6d206120736d616c6c6572206f6e6520776974682075696e74735468652070726f706f736572206973206e6f7420616c6c6f77656420746f20766f746520696e20612070726f706f73616c5265706f7369746f72792064657363726970746f722063616e277420626520656d70747950726f76696465206174206c65617374206f6e65206b6e6f776c656467652049445265706f7369746f72792063616e206f6e6c7920626520696e697469616c697a6564206f6e6365526573756c742068617320746f20626520626967676572207468616e20626f74682073756d6d616e647356616c7565206f6620746865207472616e73616374696f6e2063616e206e6f74206265207a65726fa265627a7a723058208edca67a1ff3ecdc1e249612255fc4fbd715d29c94e031046d29d6241aed310664736f6c634300050a00324b4e57566f74696e6720616e64204b4e57546f6b656e20616464726573732063616e277420626520656d707479"

// DeployDitDemoCoordinator deploys a new Ethereum contract, binding an instance of DitDemoCoordinator to it.
func DeployDitDemoCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWTokenAddress common.Address, _KNWVotingAddress common.Address, _lastDitCoordinator common.Address, _xDitTokenAddress common.Address) (common.Address, *types.Transaction, *DitDemoCoordinator, error) {
	parsed, err := abi.JSON(strings.NewReader(DitDemoCoordinatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DitDemoCoordinatorBin), backend, _KNWTokenAddress, _KNWVotingAddress, _lastDitCoordinator, _xDitTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DitDemoCoordinator{DitDemoCoordinatorCaller: DitDemoCoordinatorCaller{contract: contract}, DitDemoCoordinatorTransactor: DitDemoCoordinatorTransactor{contract: contract}, DitDemoCoordinatorFilterer: DitDemoCoordinatorFilterer{contract: contract}}, nil
}

// DitDemoCoordinator is an auto generated Go binding around an Ethereum contract.
type DitDemoCoordinator struct {
	DitDemoCoordinatorCaller     // Read-only binding to the contract
	DitDemoCoordinatorTransactor // Write-only binding to the contract
	DitDemoCoordinatorFilterer   // Log filterer for contract events
}

// DitDemoCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DitDemoCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitDemoCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DitDemoCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitDemoCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DitDemoCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DitDemoCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DitDemoCoordinatorSession struct {
	Contract     *DitDemoCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DitDemoCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DitDemoCoordinatorCallerSession struct {
	Contract *DitDemoCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DitDemoCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DitDemoCoordinatorTransactorSession struct {
	Contract     *DitDemoCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DitDemoCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DitDemoCoordinatorRaw struct {
	Contract *DitDemoCoordinator // Generic contract binding to access the raw methods on
}

// DitDemoCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DitDemoCoordinatorCallerRaw struct {
	Contract *DitDemoCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// DitDemoCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DitDemoCoordinatorTransactorRaw struct {
	Contract *DitDemoCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDitDemoCoordinator creates a new instance of DitDemoCoordinator, bound to a specific deployed contract.
func NewDitDemoCoordinator(address common.Address, backend bind.ContractBackend) (*DitDemoCoordinator, error) {
	contract, err := bindDitDemoCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinator{DitDemoCoordinatorCaller: DitDemoCoordinatorCaller{contract: contract}, DitDemoCoordinatorTransactor: DitDemoCoordinatorTransactor{contract: contract}, DitDemoCoordinatorFilterer: DitDemoCoordinatorFilterer{contract: contract}}, nil
}

// NewDitDemoCoordinatorCaller creates a new read-only instance of DitDemoCoordinator, bound to a specific deployed contract.
func NewDitDemoCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*DitDemoCoordinatorCaller, error) {
	contract, err := bindDitDemoCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorCaller{contract: contract}, nil
}

// NewDitDemoCoordinatorTransactor creates a new write-only instance of DitDemoCoordinator, bound to a specific deployed contract.
func NewDitDemoCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DitDemoCoordinatorTransactor, error) {
	contract, err := bindDitDemoCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorTransactor{contract: contract}, nil
}

// NewDitDemoCoordinatorFilterer creates a new log filterer instance of DitDemoCoordinator, bound to a specific deployed contract.
func NewDitDemoCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DitDemoCoordinatorFilterer, error) {
	contract, err := bindDitDemoCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorFilterer{contract: contract}, nil
}

// bindDitDemoCoordinator binds a generic wrapper to an already deployed contract.
func bindDitDemoCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DitDemoCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitDemoCoordinator *DitDemoCoordinatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitDemoCoordinator.Contract.DitDemoCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitDemoCoordinator *DitDemoCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.DitDemoCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitDemoCoordinator *DitDemoCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.DitDemoCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DitDemoCoordinator *DitDemoCoordinatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DitDemoCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.contract.Transact(opts, method, params...)
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) BURNINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "BURNING_METHOD")
	return *ret0, err
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) BURNINGMETHOD() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.BURNINGMETHOD(&_DitDemoCoordinator.CallOpts)
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) BURNINGMETHOD() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.BURNINGMETHOD(&_DitDemoCoordinator.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) KNWTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "KNWTokenAddress")
	return *ret0, err
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) KNWTokenAddress() (common.Address, error) {
	return _DitDemoCoordinator.Contract.KNWTokenAddress(&_DitDemoCoordinator.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) KNWTokenAddress() (common.Address, error) {
	return _DitDemoCoordinator.Contract.KNWTokenAddress(&_DitDemoCoordinator.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) KNWVotingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "KNWVotingAddress")
	return *ret0, err
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) KNWVotingAddress() (common.Address, error) {
	return _DitDemoCoordinator.Contract.KNWVotingAddress(&_DitDemoCoordinator.CallOpts)
}

// KNWVotingAddress is a free data retrieval call binding the contract method 0xeb49fe94.
//
// Solidity: function KNWVotingAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) KNWVotingAddress() (common.Address, error) {
	return _DitDemoCoordinator.Contract.KNWVotingAddress(&_DitDemoCoordinator.CallOpts)
}

// MAXVOTEDURATION is a free data retrieval call binding the contract method 0x1ecbb9de.
//
// Solidity: function MAX_VOTE_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) MAXVOTEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "MAX_VOTE_DURATION")
	return *ret0, err
}

// MAXVOTEDURATION is a free data retrieval call binding the contract method 0x1ecbb9de.
//
// Solidity: function MAX_VOTE_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MAXVOTEDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MAXVOTEDURATION(&_DitDemoCoordinator.CallOpts)
}

// MAXVOTEDURATION is a free data retrieval call binding the contract method 0x1ecbb9de.
//
// Solidity: function MAX_VOTE_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) MAXVOTEDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MAXVOTEDURATION(&_DitDemoCoordinator.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) MINTINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "MINTING_METHOD")
	return *ret0, err
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MINTINGMETHOD() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MINTINGMETHOD(&_DitDemoCoordinator.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) MINTINGMETHOD() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MINTINGMETHOD(&_DitDemoCoordinator.CallOpts)
}

// MINVOTEDURATION is a free data retrieval call binding the contract method 0x3eedfc10.
//
// Solidity: function MIN_VOTE_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) MINVOTEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "MIN_VOTE_DURATION")
	return *ret0, err
}

// MINVOTEDURATION is a free data retrieval call binding the contract method 0x3eedfc10.
//
// Solidity: function MIN_VOTE_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MINVOTEDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MINVOTEDURATION(&_DitDemoCoordinator.CallOpts)
}

// MINVOTEDURATION is a free data retrieval call binding the contract method 0x3eedfc10.
//
// Solidity: function MIN_VOTE_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) MINVOTEDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MINVOTEDURATION(&_DitDemoCoordinator.CallOpts)
}

// AllowedKnowledgeIDs is a free data retrieval call binding the contract method 0xb3070331.
//
// Solidity: function allowedKnowledgeIDs(bytes32 , uint256 ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) AllowedKnowledgeIDs(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "allowedKnowledgeIDs", arg0, arg1)
	return *ret0, err
}

// AllowedKnowledgeIDs is a free data retrieval call binding the contract method 0xb3070331.
//
// Solidity: function allowedKnowledgeIDs(bytes32 , uint256 ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) AllowedKnowledgeIDs(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _DitDemoCoordinator.Contract.AllowedKnowledgeIDs(&_DitDemoCoordinator.CallOpts, arg0, arg1)
}

// AllowedKnowledgeIDs is a free data retrieval call binding the contract method 0xb3070331.
//
// Solidity: function allowedKnowledgeIDs(bytes32 , uint256 ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) AllowedKnowledgeIDs(arg0 [32]byte, arg1 *big.Int) (bool, error) {
	return _DitDemoCoordinator.Contract.AllowedKnowledgeIDs(&_DitDemoCoordinator.CallOpts, arg0, arg1)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256 currentProposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) GetCurrentProposalID(opts *bind.CallOpts, _repository [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "getCurrentProposalID", _repository)
	return *ret0, err
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256 currentProposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetCurrentProposalID(&_DitDemoCoordinator.CallOpts, _repository)
}

// GetCurrentProposalID is a free data retrieval call binding the contract method 0x0bdc90e8.
//
// Solidity: function getCurrentProposalID(bytes32 _repository) constant returns(uint256 currentProposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) GetCurrentProposalID(_repository [32]byte) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetCurrentProposalID(&_DitDemoCoordinator.CallOpts, _repository)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) GetIndividualStake(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "getIndividualStake", _repository, _proposalID)
	return *ret0, err
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetIndividualStake(&_DitDemoCoordinator.CallOpts, _repository, _proposalID)
}

// GetIndividualStake is a free data retrieval call binding the contract method 0x552edc9d.
//
// Solidity: function getIndividualStake(bytes32 _repository, uint256 _proposalID) constant returns(uint256 individualStake)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) GetIndividualStake(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetIndividualStake(&_DitDemoCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256 KNWVoteID)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) GetKNWVoteIDFromProposalID(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "getKNWVoteIDFromProposalID", _repository, _proposalID)
	return *ret0, err
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256 KNWVoteID)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitDemoCoordinator.CallOpts, _repository, _proposalID)
}

// GetKNWVoteIDFromProposalID is a free data retrieval call binding the contract method 0x06ee4596.
//
// Solidity: function getKNWVoteIDFromProposalID(bytes32 _repository, uint256 _proposalID) constant returns(uint256 KNWVoteID)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) GetKNWVoteIDFromProposalID(_repository [32]byte, _proposalID *big.Int) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetKNWVoteIDFromProposalID(&_DitDemoCoordinator.CallOpts, _repository, _proposalID)
}

// GetKnowledgeIDs is a free data retrieval call binding the contract method 0xdf1bd2cd.
//
// Solidity: function getKnowledgeIDs(bytes32 _repository) constant returns(uint256[] knowledgeIDs)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) GetKnowledgeIDs(opts *bind.CallOpts, _repository [32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "getKnowledgeIDs", _repository)
	return *ret0, err
}

// GetKnowledgeIDs is a free data retrieval call binding the contract method 0xdf1bd2cd.
//
// Solidity: function getKnowledgeIDs(bytes32 _repository) constant returns(uint256[] knowledgeIDs)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) GetKnowledgeIDs(_repository [32]byte) ([]*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetKnowledgeIDs(&_DitDemoCoordinator.CallOpts, _repository)
}

// GetKnowledgeIDs is a free data retrieval call binding the contract method 0xdf1bd2cd.
//
// Solidity: function getKnowledgeIDs(bytes32 _repository) constant returns(uint256[] knowledgeIDs)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) GetKnowledgeIDs(_repository [32]byte) ([]*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetKnowledgeIDs(&_DitDemoCoordinator.CallOpts, _repository)
}

// GetVotingMajority is a free data retrieval call binding the contract method 0x3fcc148d.
//
// Solidity: function getVotingMajority(bytes32 _repository) constant returns(uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) GetVotingMajority(opts *bind.CallOpts, _repository [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "getVotingMajority", _repository)
	return *ret0, err
}

// GetVotingMajority is a free data retrieval call binding the contract method 0x3fcc148d.
//
// Solidity: function getVotingMajority(bytes32 _repository) constant returns(uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) GetVotingMajority(_repository [32]byte) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetVotingMajority(&_DitDemoCoordinator.CallOpts, _repository)
}

// GetVotingMajority is a free data retrieval call binding the contract method 0x3fcc148d.
//
// Solidity: function getVotingMajority(bytes32 _repository) constant returns(uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) GetVotingMajority(_repository [32]byte) (*big.Int, error) {
	return _DitDemoCoordinator.Contract.GetVotingMajority(&_DitDemoCoordinator.CallOpts, _repository)
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) IsKYCValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "isKYCValidator", arg0)
	return *ret0, err
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) IsKYCValidator(arg0 common.Address) (bool, error) {
	return _DitDemoCoordinator.Contract.IsKYCValidator(&_DitDemoCoordinator.CallOpts, arg0)
}

// IsKYCValidator is a free data retrieval call binding the contract method 0x1341f25c.
//
// Solidity: function isKYCValidator(address ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) IsKYCValidator(arg0 common.Address) (bool, error) {
	return _DitDemoCoordinator.Contract.IsKYCValidator(&_DitDemoCoordinator.CallOpts, arg0)
}

// LastDitCoordinator is a free data retrieval call binding the contract method 0xd7ad0eae.
//
// Solidity: function lastDitCoordinator() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) LastDitCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "lastDitCoordinator")
	return *ret0, err
}

// LastDitCoordinator is a free data retrieval call binding the contract method 0xd7ad0eae.
//
// Solidity: function lastDitCoordinator() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) LastDitCoordinator() (common.Address, error) {
	return _DitDemoCoordinator.Contract.LastDitCoordinator(&_DitDemoCoordinator.CallOpts)
}

// LastDitCoordinator is a free data retrieval call binding the contract method 0xd7ad0eae.
//
// Solidity: function lastDitCoordinator() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) LastDitCoordinator() (common.Address, error) {
	return _DitDemoCoordinator.Contract.LastDitCoordinator(&_DitDemoCoordinator.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) Manager() (common.Address, error) {
	return _DitDemoCoordinator.Contract.Manager(&_DitDemoCoordinator.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) Manager() (common.Address, error) {
	return _DitDemoCoordinator.Contract.Manager(&_DitDemoCoordinator.CallOpts)
}

// NextDitCoordinator is a free data retrieval call binding the contract method 0x99821d9b.
//
// Solidity: function nextDitCoordinator() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) NextDitCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "nextDitCoordinator")
	return *ret0, err
}

// NextDitCoordinator is a free data retrieval call binding the contract method 0x99821d9b.
//
// Solidity: function nextDitCoordinator() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) NextDitCoordinator() (common.Address, error) {
	return _DitDemoCoordinator.Contract.NextDitCoordinator(&_DitDemoCoordinator.CallOpts)
}

// NextDitCoordinator is a free data retrieval call binding the contract method 0x99821d9b.
//
// Solidity: function nextDitCoordinator() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) NextDitCoordinator() (common.Address, error) {
	return _DitDemoCoordinator.Contract.NextDitCoordinator(&_DitDemoCoordinator.CallOpts)
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) PassedKYC(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "passedKYC", arg0)
	return *ret0, err
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) PassedKYC(arg0 common.Address) (bool, error) {
	return _DitDemoCoordinator.Contract.PassedKYC(&_DitDemoCoordinator.CallOpts, arg0)
}

// PassedKYC is a free data retrieval call binding the contract method 0xccd9aa68.
//
// Solidity: function passedKYC(address ) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) PassedKYC(arg0 common.Address) (bool, error) {
	return _DitDemoCoordinator.Contract.PassedKYC(&_DitDemoCoordinator.CallOpts, arg0)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) ProposalHasPassed(opts *bind.CallOpts, _repository [32]byte, _proposalID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "proposalHasPassed", _repository, _proposalID)
	return *ret0, err
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitDemoCoordinator.Contract.ProposalHasPassed(&_DitDemoCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalHasPassed is a free data retrieval call binding the contract method 0x87c9203d.
//
// Solidity: function proposalHasPassed(bytes32 _repository, uint256 _proposalID) constant returns(bool hasPassed)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) ProposalHasPassed(_repository [32]byte, _proposalID *big.Int) (bool, error) {
	return _DitDemoCoordinator.Contract.ProposalHasPassed(&_DitDemoCoordinator.CallOpts, _repository, _proposalID)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(string description, string identifier, uint256 KNWVoteID, uint256 knowledgeID, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) ProposalsOfRepository(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
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
	err := _DitDemoCoordinator.contract.Call(opts, out, "proposalsOfRepository", arg0, arg1)
	return *ret, err
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(string description, string identifier, uint256 KNWVoteID, uint256 knowledgeID, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
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
	return _DitDemoCoordinator.Contract.ProposalsOfRepository(&_DitDemoCoordinator.CallOpts, arg0, arg1)
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(string description, string identifier, uint256 KNWVoteID, uint256 knowledgeID, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
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
	return _DitDemoCoordinator.Contract.ProposalsOfRepository(&_DitDemoCoordinator.CallOpts, arg0, arg1)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(string name, uint256 currentProposalID, uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) Repositories(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _DitDemoCoordinator.contract.Call(opts, out, "repositories", arg0)
	return *ret, err
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(string name, uint256 currentProposalID, uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) Repositories(arg0 [32]byte) (struct {
	Name              string
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	return _DitDemoCoordinator.Contract.Repositories(&_DitDemoCoordinator.CallOpts, arg0)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(string name, uint256 currentProposalID, uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) Repositories(arg0 [32]byte) (struct {
	Name              string
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	return _DitDemoCoordinator.Contract.Repositories(&_DitDemoCoordinator.CallOpts, arg0)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) RepositoryIsInitialized(opts *bind.CallOpts, _repository [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "repositoryIsInitialized", _repository)
	return *ret0, err
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitDemoCoordinator.Contract.RepositoryIsInitialized(&_DitDemoCoordinator.CallOpts, _repository)
}

// RepositoryIsInitialized is a free data retrieval call binding the contract method 0xea6c6d0f.
//
// Solidity: function repositoryIsInitialized(bytes32 _repository) constant returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) RepositoryIsInitialized(_repository [32]byte) (bool, error) {
	return _DitDemoCoordinator.Contract.RepositoryIsInitialized(&_DitDemoCoordinator.CallOpts, _repository)
}

// XDitTokenAddress is a free data retrieval call binding the contract method 0xce83732e.
//
// Solidity: function xDitTokenAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) XDitTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "xDitTokenAddress")
	return *ret0, err
}

// XDitTokenAddress is a free data retrieval call binding the contract method 0xce83732e.
//
// Solidity: function xDitTokenAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) XDitTokenAddress() (common.Address, error) {
	return _DitDemoCoordinator.Contract.XDitTokenAddress(&_DitDemoCoordinator.CallOpts)
}

// XDitTokenAddress is a free data retrieval call binding the contract method 0xce83732e.
//
// Solidity: function xDitTokenAddress() constant returns(address)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) XDitTokenAddress() (common.Address, error) {
	return _DitDemoCoordinator.Contract.XDitTokenAddress(&_DitDemoCoordinator.CallOpts)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) AddKYCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "addKYCValidator", _address)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) AddKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.AddKYCValidator(&_DitDemoCoordinator.TransactOpts, _address)
}

// AddKYCValidator is a paid mutator transaction binding the contract method 0xd0c397ef.
//
// Solidity: function addKYCValidator(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) AddKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.AddKYCValidator(&_DitDemoCoordinator.TransactOpts, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) FinalizeVote(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "finalizeVote", _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.FinalizeVote(&_DitDemoCoordinator.TransactOpts, _repository, _proposalID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x2e71d0fb.
//
// Solidity: function finalizeVote(bytes32 _repository, uint256 _proposalID) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) FinalizeVote(_repository [32]byte, _proposalID *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.FinalizeVote(&_DitDemoCoordinator.TransactOpts, _repository, _proposalID)
}

// InitRepository is a paid mutator transaction binding the contract method 0x6c63b91e.
//
// Solidity: function initRepository(string _repository, uint256[] _knowledgeIDs, uint256 _votingMajority) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository string, _knowledgeIDs []*big.Int, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "initRepository", _repository, _knowledgeIDs, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x6c63b91e.
//
// Solidity: function initRepository(string _repository, uint256[] _knowledgeIDs, uint256 _votingMajority) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) InitRepository(_repository string, _knowledgeIDs []*big.Int, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.InitRepository(&_DitDemoCoordinator.TransactOpts, _repository, _knowledgeIDs, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x6c63b91e.
//
// Solidity: function initRepository(string _repository, uint256[] _knowledgeIDs, uint256 _votingMajority) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) InitRepository(_repository string, _knowledgeIDs []*big.Int, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.InitRepository(&_DitDemoCoordinator.TransactOpts, _repository, _knowledgeIDs, _votingMajority)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xf60055cb.
//
// Solidity: function migrateRepository(string _repository) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) MigrateRepository(opts *bind.TransactOpts, _repository string) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "migrateRepository", _repository)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xf60055cb.
//
// Solidity: function migrateRepository(string _repository) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MigrateRepository(_repository string) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.MigrateRepository(&_DitDemoCoordinator.TransactOpts, _repository)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xf60055cb.
//
// Solidity: function migrateRepository(string _repository) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) MigrateRepository(_repository string) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.MigrateRepository(&_DitDemoCoordinator.TransactOpts, _repository)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) OpenVoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "openVoteOnProposal", _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.OpenVoteOnProposal(&_DitDemoCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// OpenVoteOnProposal is a paid mutator transaction binding the contract method 0x0ee62ec0.
//
// Solidity: function openVoteOnProposal(bytes32 _repository, uint256 _proposalID, uint256 _voteOption, uint256 _voteSalt) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) OpenVoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteOption *big.Int, _voteSalt *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.OpenVoteOnProposal(&_DitDemoCoordinator.TransactOpts, _repository, _proposalID, _voteOption, _voteSalt)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) PassKYC(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "passKYC", _address)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) PassKYC(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.PassKYC(&_DitDemoCoordinator.TransactOpts, _address)
}

// PassKYC is a paid mutator transaction binding the contract method 0xeb931024.
//
// Solidity: function passKYC(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) PassKYC(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.PassKYC(&_DitDemoCoordinator.TransactOpts, _address)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xfb4ede85.
//
// Solidity: function proposeCommit(bytes32 _repository, string _description, string _identifier, uint256 _knowledgeID, uint256 _numberOfKNW, uint256 _voteDuration, uint256 _amountOfTokens) returns(uint256 proposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) ProposeCommit(opts *bind.TransactOpts, _repository [32]byte, _description string, _identifier string, _knowledgeID *big.Int, _numberOfKNW *big.Int, _voteDuration *big.Int, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "proposeCommit", _repository, _description, _identifier, _knowledgeID, _numberOfKNW, _voteDuration, _amountOfTokens)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xfb4ede85.
//
// Solidity: function proposeCommit(bytes32 _repository, string _description, string _identifier, uint256 _knowledgeID, uint256 _numberOfKNW, uint256 _voteDuration, uint256 _amountOfTokens) returns(uint256 proposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ProposeCommit(_repository [32]byte, _description string, _identifier string, _knowledgeID *big.Int, _numberOfKNW *big.Int, _voteDuration *big.Int, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ProposeCommit(&_DitDemoCoordinator.TransactOpts, _repository, _description, _identifier, _knowledgeID, _numberOfKNW, _voteDuration, _amountOfTokens)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xfb4ede85.
//
// Solidity: function proposeCommit(bytes32 _repository, string _description, string _identifier, uint256 _knowledgeID, uint256 _numberOfKNW, uint256 _voteDuration, uint256 _amountOfTokens) returns(uint256 proposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) ProposeCommit(_repository [32]byte, _description string, _identifier string, _knowledgeID *big.Int, _numberOfKNW *big.Int, _voteDuration *big.Int, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ProposeCommit(&_DitDemoCoordinator.TransactOpts, _repository, _description, _identifier, _knowledgeID, _numberOfKNW, _voteDuration, _amountOfTokens)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) RemoveKYCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "removeKYCValidator", _address)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) RemoveKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.RemoveKYCValidator(&_DitDemoCoordinator.TransactOpts, _address)
}

// RemoveKYCValidator is a paid mutator transaction binding the contract method 0x73b0dddd.
//
// Solidity: function removeKYCValidator(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) RemoveKYCValidator(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.RemoveKYCValidator(&_DitDemoCoordinator.TransactOpts, _address)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) ReplaceManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "replaceManager", _newManager)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ReplaceManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ReplaceManager(&_DitDemoCoordinator.TransactOpts, _newManager)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) ReplaceManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ReplaceManager(&_DitDemoCoordinator.TransactOpts, _newManager)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) RevokeKYC(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "revokeKYC", _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) RevokeKYC(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.RevokeKYC(&_DitDemoCoordinator.TransactOpts, _address)
}

// RevokeKYC is a paid mutator transaction binding the contract method 0x39ba645b.
//
// Solidity: function revokeKYC(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) RevokeKYC(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.RevokeKYC(&_DitDemoCoordinator.TransactOpts, _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) UpgradeContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "upgradeContract", _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.UpgradeContract(&_DitDemoCoordinator.TransactOpts, _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.UpgradeContract(&_DitDemoCoordinator.TransactOpts, _address)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xab4b593e.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash, uint256 _numberOfKNW) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) VoteOnProposal(opts *bind.TransactOpts, _repository [32]byte, _proposalID *big.Int, _voteHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "voteOnProposal", _repository, _proposalID, _voteHash, _numberOfKNW)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xab4b593e.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash, uint256 _numberOfKNW) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.VoteOnProposal(&_DitDemoCoordinator.TransactOpts, _repository, _proposalID, _voteHash, _numberOfKNW)
}

// VoteOnProposal is a paid mutator transaction binding the contract method 0xab4b593e.
//
// Solidity: function voteOnProposal(bytes32 _repository, uint256 _proposalID, bytes32 _voteHash, uint256 _numberOfKNW) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) VoteOnProposal(_repository [32]byte, _proposalID *big.Int, _voteHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.VoteOnProposal(&_DitDemoCoordinator.TransactOpts, _repository, _proposalID, _voteHash, _numberOfKNW)
}

// DitDemoCoordinatorCommitVoteIterator is returned from FilterCommitVote and is used to iterate over the raw logs and unpacked data for CommitVote events raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorCommitVoteIterator struct {
	Event *DitDemoCoordinatorCommitVote // Event containing the contract specifics and raw log

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
func (it *DitDemoCoordinatorCommitVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitDemoCoordinatorCommitVote)
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
		it.Event = new(DitDemoCoordinatorCommitVote)
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
func (it *DitDemoCoordinatorCommitVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitDemoCoordinatorCommitVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitDemoCoordinatorCommitVote represents a CommitVote event raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorCommitVote struct {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) FilterCommitVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitDemoCoordinatorCommitVoteIterator, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.FilterLogs(opts, "CommitVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorCommitVoteIterator{contract: _DitDemoCoordinator.contract, event: "CommitVote", logs: logs, sub: sub}, nil
}

// WatchCommitVote is a free log subscription operation binding the contract event 0x2edd4beb5f5626bb825280acc138e2ffc3190dff1e0d8e4563febc24ed342549.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) WatchCommitVote(opts *bind.WatchOpts, sink chan<- *DitDemoCoordinatorCommitVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.WatchLogs(opts, "CommitVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitDemoCoordinatorCommitVote)
				if err := _DitDemoCoordinator.contract.UnpackLog(event, "CommitVote", log); err != nil {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) ParseCommitVote(log types.Log) (*DitDemoCoordinatorCommitVote, error) {
	event := new(DitDemoCoordinatorCommitVote)
	if err := _DitDemoCoordinator.contract.UnpackLog(event, "CommitVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DitDemoCoordinatorFinalizeProposalIterator is returned from FilterFinalizeProposal and is used to iterate over the raw logs and unpacked data for FinalizeProposal events raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorFinalizeProposalIterator struct {
	Event *DitDemoCoordinatorFinalizeProposal // Event containing the contract specifics and raw log

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
func (it *DitDemoCoordinatorFinalizeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitDemoCoordinatorFinalizeProposal)
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
		it.Event = new(DitDemoCoordinatorFinalizeProposal)
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
func (it *DitDemoCoordinatorFinalizeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitDemoCoordinatorFinalizeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitDemoCoordinatorFinalizeProposal represents a FinalizeProposal event raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorFinalizeProposal struct {
	Repository  [32]byte
	Proposal    *big.Int
	KnowledgeID *big.Int
	Accepted    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizeProposal is a free log retrieval operation binding the contract event 0x7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc1.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, uint256 knowledgeID, bool accepted)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) FilterFinalizeProposal(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int) (*DitDemoCoordinatorFinalizeProposalIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitDemoCoordinator.contract.FilterLogs(opts, "FinalizeProposal", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorFinalizeProposalIterator{contract: _DitDemoCoordinator.contract, event: "FinalizeProposal", logs: logs, sub: sub}, nil
}

// WatchFinalizeProposal is a free log subscription operation binding the contract event 0x7bb26c044d8d0bb569982fbbecdc73c4d75f05aa82ae7575f55d136d6d9c4bc1.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, uint256 knowledgeID, bool accepted)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) WatchFinalizeProposal(opts *bind.WatchOpts, sink chan<- *DitDemoCoordinatorFinalizeProposal, repository [][32]byte, proposal []*big.Int) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var proposalRule []interface{}
	for _, proposalItem := range proposal {
		proposalRule = append(proposalRule, proposalItem)
	}

	logs, sub, err := _DitDemoCoordinator.contract.WatchLogs(opts, "FinalizeProposal", repositoryRule, proposalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitDemoCoordinatorFinalizeProposal)
				if err := _DitDemoCoordinator.contract.UnpackLog(event, "FinalizeProposal", log); err != nil {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) ParseFinalizeProposal(log types.Log) (*DitDemoCoordinatorFinalizeProposal, error) {
	event := new(DitDemoCoordinatorFinalizeProposal)
	if err := _DitDemoCoordinator.contract.UnpackLog(event, "FinalizeProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DitDemoCoordinatorFinalizeVoteIterator is returned from FilterFinalizeVote and is used to iterate over the raw logs and unpacked data for FinalizeVote events raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorFinalizeVoteIterator struct {
	Event *DitDemoCoordinatorFinalizeVote // Event containing the contract specifics and raw log

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
func (it *DitDemoCoordinatorFinalizeVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitDemoCoordinatorFinalizeVote)
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
		it.Event = new(DitDemoCoordinatorFinalizeVote)
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
func (it *DitDemoCoordinatorFinalizeVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitDemoCoordinatorFinalizeVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitDemoCoordinatorFinalizeVote represents a FinalizeVote event raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorFinalizeVote struct {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) FilterFinalizeVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitDemoCoordinatorFinalizeVoteIterator, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.FilterLogs(opts, "FinalizeVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorFinalizeVoteIterator{contract: _DitDemoCoordinator.contract, event: "FinalizeVote", logs: logs, sub: sub}, nil
}

// WatchFinalizeVote is a free log subscription operation binding the contract event 0x45b8e8084f290362eff12aa8825548cacdd5ce5bbd00dbcaeed1bec7cc8dbcd5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool votedRight, uint256 numberOfKNW)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) WatchFinalizeVote(opts *bind.WatchOpts, sink chan<- *DitDemoCoordinatorFinalizeVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.WatchLogs(opts, "FinalizeVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitDemoCoordinatorFinalizeVote)
				if err := _DitDemoCoordinator.contract.UnpackLog(event, "FinalizeVote", log); err != nil {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) ParseFinalizeVote(log types.Log) (*DitDemoCoordinatorFinalizeVote, error) {
	event := new(DitDemoCoordinatorFinalizeVote)
	if err := _DitDemoCoordinator.contract.UnpackLog(event, "FinalizeVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DitDemoCoordinatorInitializeRepositoryIterator is returned from FilterInitializeRepository and is used to iterate over the raw logs and unpacked data for InitializeRepository events raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorInitializeRepositoryIterator struct {
	Event *DitDemoCoordinatorInitializeRepository // Event containing the contract specifics and raw log

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
func (it *DitDemoCoordinatorInitializeRepositoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitDemoCoordinatorInitializeRepository)
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
		it.Event = new(DitDemoCoordinatorInitializeRepository)
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
func (it *DitDemoCoordinatorInitializeRepositoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitDemoCoordinatorInitializeRepositoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitDemoCoordinatorInitializeRepository represents a InitializeRepository event raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorInitializeRepository struct {
	Repository [32]byte
	Who        common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitializeRepository is a free log retrieval operation binding the contract event 0xe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be.
//
// Solidity: event InitializeRepository(bytes32 indexed repository, address indexed who)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) FilterInitializeRepository(opts *bind.FilterOpts, repository [][32]byte, who []common.Address) (*DitDemoCoordinatorInitializeRepositoryIterator, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitDemoCoordinator.contract.FilterLogs(opts, "InitializeRepository", repositoryRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorInitializeRepositoryIterator{contract: _DitDemoCoordinator.contract, event: "InitializeRepository", logs: logs, sub: sub}, nil
}

// WatchInitializeRepository is a free log subscription operation binding the contract event 0xe81437d5f3c837aab5b46f923704dc104204a127033e0dd0e24e7ad8d2b3c7be.
//
// Solidity: event InitializeRepository(bytes32 indexed repository, address indexed who)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) WatchInitializeRepository(opts *bind.WatchOpts, sink chan<- *DitDemoCoordinatorInitializeRepository, repository [][32]byte, who []common.Address) (event.Subscription, error) {

	var repositoryRule []interface{}
	for _, repositoryItem := range repository {
		repositoryRule = append(repositoryRule, repositoryItem)
	}
	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _DitDemoCoordinator.contract.WatchLogs(opts, "InitializeRepository", repositoryRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitDemoCoordinatorInitializeRepository)
				if err := _DitDemoCoordinator.contract.UnpackLog(event, "InitializeRepository", log); err != nil {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) ParseInitializeRepository(log types.Log) (*DitDemoCoordinatorInitializeRepository, error) {
	event := new(DitDemoCoordinatorInitializeRepository)
	if err := _DitDemoCoordinator.contract.UnpackLog(event, "InitializeRepository", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DitDemoCoordinatorOpenVoteIterator is returned from FilterOpenVote and is used to iterate over the raw logs and unpacked data for OpenVote events raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorOpenVoteIterator struct {
	Event *DitDemoCoordinatorOpenVote // Event containing the contract specifics and raw log

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
func (it *DitDemoCoordinatorOpenVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitDemoCoordinatorOpenVote)
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
		it.Event = new(DitDemoCoordinatorOpenVote)
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
func (it *DitDemoCoordinatorOpenVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitDemoCoordinatorOpenVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitDemoCoordinatorOpenVote represents a OpenVote event raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorOpenVote struct {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) FilterOpenVote(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitDemoCoordinatorOpenVoteIterator, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.FilterLogs(opts, "OpenVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorOpenVoteIterator{contract: _DitDemoCoordinator.contract, event: "OpenVote", logs: logs, sub: sub}, nil
}

// WatchOpenVote is a free log subscription operation binding the contract event 0x1a245c311cb3eef22175d9c7c5458569e91d9a5956430216eaf8de7c1b738760.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, bool accept, uint256 numberOfVotes)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) WatchOpenVote(opts *bind.WatchOpts, sink chan<- *DitDemoCoordinatorOpenVote, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.WatchLogs(opts, "OpenVote", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitDemoCoordinatorOpenVote)
				if err := _DitDemoCoordinator.contract.UnpackLog(event, "OpenVote", log); err != nil {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) ParseOpenVote(log types.Log) (*DitDemoCoordinatorOpenVote, error) {
	event := new(DitDemoCoordinatorOpenVote)
	if err := _DitDemoCoordinator.contract.UnpackLog(event, "OpenVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DitDemoCoordinatorProposeCommitIterator is returned from FilterProposeCommit and is used to iterate over the raw logs and unpacked data for ProposeCommit events raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorProposeCommitIterator struct {
	Event *DitDemoCoordinatorProposeCommit // Event containing the contract specifics and raw log

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
func (it *DitDemoCoordinatorProposeCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DitDemoCoordinatorProposeCommit)
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
		it.Event = new(DitDemoCoordinatorProposeCommit)
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
func (it *DitDemoCoordinatorProposeCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DitDemoCoordinatorProposeCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DitDemoCoordinatorProposeCommit represents a ProposeCommit event raised by the DitDemoCoordinator contract.
type DitDemoCoordinatorProposeCommit struct {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) FilterProposeCommit(opts *bind.FilterOpts, repository [][32]byte, proposal []*big.Int, who []common.Address) (*DitDemoCoordinatorProposeCommitIterator, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.FilterLogs(opts, "ProposeCommit", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return &DitDemoCoordinatorProposeCommitIterator{contract: _DitDemoCoordinator.contract, event: "ProposeCommit", logs: logs, sub: sub}, nil
}

// WatchProposeCommit is a free log subscription operation binding the contract event 0x603f65981d2b692fdcecdd4b4c97ebe242a11b1feb833e2c556a6a3bc27e4314.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, uint256 knowledgeID, uint256 numberOfKNW)
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) WatchProposeCommit(opts *bind.WatchOpts, sink chan<- *DitDemoCoordinatorProposeCommit, repository [][32]byte, proposal []*big.Int, who []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DitDemoCoordinator.contract.WatchLogs(opts, "ProposeCommit", repositoryRule, proposalRule, whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DitDemoCoordinatorProposeCommit)
				if err := _DitDemoCoordinator.contract.UnpackLog(event, "ProposeCommit", log); err != nil {
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
func (_DitDemoCoordinator *DitDemoCoordinatorFilterer) ParseProposeCommit(log types.Log) (*DitDemoCoordinatorProposeCommit, error) {
	event := new(DitDemoCoordinatorProposeCommit)
	if err := _DitDemoCoordinator.contract.UnpackLog(event, "ProposeCommit", log); err != nil {
		return nil, err
	}
	return event, nil
}
