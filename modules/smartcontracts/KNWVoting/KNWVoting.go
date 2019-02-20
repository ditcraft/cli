// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KNWVoting

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
const KNWTokenContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"freeBalanceOfLabel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"lockTokens\",\"outputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_winningPercentage\",\"type\":\"uint256\"},{\"name\":\"_mintingMethod\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_numberOfTokens\",\"type\":\"uint256\"},{\"name\":\"_winningPercentage\",\"type\":\"uint256\"},{\"name\":\"_burningMethod\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"balanceOfLabel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(_account address, _label string) constant returns(uint256)
func (_KNWTokenContract *KNWTokenContractCaller) BalanceOfLabel(opts *bind.CallOpts, _account common.Address, _label string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWTokenContract.contract.Call(opts, out, "balanceOfLabel", _account, _label)
	return *ret0, err
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(_account address, _label string) constant returns(uint256)
func (_KNWTokenContract *KNWTokenContractSession) BalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.BalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(_account address, _label string) constant returns(uint256)
func (_KNWTokenContract *KNWTokenContractCallerSession) BalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.BalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(_account address, _label string) constant returns(uint256)
func (_KNWTokenContract *KNWTokenContractCaller) FreeBalanceOfLabel(opts *bind.CallOpts, _account common.Address, _label string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWTokenContract.contract.Call(opts, out, "freeBalanceOfLabel", _account, _label)
	return *ret0, err
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(_account address, _label string) constant returns(uint256)
func (_KNWTokenContract *KNWTokenContractSession) FreeBalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.FreeBalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(_account address, _label string) constant returns(uint256)
func (_KNWTokenContract *KNWTokenContractCallerSession) FreeBalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.FreeBalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// Burn is a paid mutator transaction binding the contract method 0x994dd93a.
//
// Solidity: function burn(_account address, _label string, _numberOfTokens uint256, _winningPercentage uint256, _burningMethod uint256) returns()
func (_KNWTokenContract *KNWTokenContractTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _label string, _numberOfTokens *big.Int, _winningPercentage *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "burn", _account, _label, _numberOfTokens, _winningPercentage, _burningMethod)
}

// Burn is a paid mutator transaction binding the contract method 0x994dd93a.
//
// Solidity: function burn(_account address, _label string, _numberOfTokens uint256, _winningPercentage uint256, _burningMethod uint256) returns()
func (_KNWTokenContract *KNWTokenContractSession) Burn(_account common.Address, _label string, _numberOfTokens *big.Int, _winningPercentage *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Burn(&_KNWTokenContract.TransactOpts, _account, _label, _numberOfTokens, _winningPercentage, _burningMethod)
}

// Burn is a paid mutator transaction binding the contract method 0x994dd93a.
//
// Solidity: function burn(_account address, _label string, _numberOfTokens uint256, _winningPercentage uint256, _burningMethod uint256) returns()
func (_KNWTokenContract *KNWTokenContractTransactorSession) Burn(_account common.Address, _label string, _numberOfTokens *big.Int, _winningPercentage *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Burn(&_KNWTokenContract.TransactOpts, _account, _label, _numberOfTokens, _winningPercentage, _burningMethod)
}

// LockTokens is a paid mutator transaction binding the contract method 0x1d3316d0.
//
// Solidity: function lockTokens(_account address, _label string) returns(numberOfTokens uint256)
func (_KNWTokenContract *KNWTokenContractTransactor) LockTokens(opts *bind.TransactOpts, _account common.Address, _label string) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "lockTokens", _account, _label)
}

// LockTokens is a paid mutator transaction binding the contract method 0x1d3316d0.
//
// Solidity: function lockTokens(_account address, _label string) returns(numberOfTokens uint256)
func (_KNWTokenContract *KNWTokenContractSession) LockTokens(_account common.Address, _label string) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.LockTokens(&_KNWTokenContract.TransactOpts, _account, _label)
}

// LockTokens is a paid mutator transaction binding the contract method 0x1d3316d0.
//
// Solidity: function lockTokens(_account address, _label string) returns(numberOfTokens uint256)
func (_KNWTokenContract *KNWTokenContractTransactorSession) LockTokens(_account common.Address, _label string) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.LockTokens(&_KNWTokenContract.TransactOpts, _account, _label)
}

// Mint is a paid mutator transaction binding the contract method 0x4a5dc3aa.
//
// Solidity: function mint(_account address, _label string, _winningPercentage uint256, _mintingMethod uint256) returns()
func (_KNWTokenContract *KNWTokenContractTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _label string, _winningPercentage *big.Int, _mintingMethod *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "mint", _account, _label, _winningPercentage, _mintingMethod)
}

// Mint is a paid mutator transaction binding the contract method 0x4a5dc3aa.
//
// Solidity: function mint(_account address, _label string, _winningPercentage uint256, _mintingMethod uint256) returns()
func (_KNWTokenContract *KNWTokenContractSession) Mint(_account common.Address, _label string, _winningPercentage *big.Int, _mintingMethod *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Mint(&_KNWTokenContract.TransactOpts, _account, _label, _winningPercentage, _mintingMethod)
}

// Mint is a paid mutator transaction binding the contract method 0x4a5dc3aa.
//
// Solidity: function mint(_account address, _label string, _winningPercentage uint256, _mintingMethod uint256) returns()
func (_KNWTokenContract *KNWTokenContractTransactorSession) Mint(_account common.Address, _label string, _winningPercentage *big.Int, _mintingMethod *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Mint(&_KNWTokenContract.TransactOpts, _account, _label, _winningPercentage, _mintingMethod)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(_account address, _label string, _numberOfTokens uint256) returns()
func (_KNWTokenContract *KNWTokenContractTransactor) UnlockTokens(opts *bind.TransactOpts, _account common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "unlockTokens", _account, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(_account address, _label string, _numberOfTokens uint256) returns()
func (_KNWTokenContract *KNWTokenContractSession) UnlockTokens(_account common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.UnlockTokens(&_KNWTokenContract.TransactOpts, _account, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(_account address, _label string, _numberOfTokens uint256) returns()
func (_KNWTokenContract *KNWTokenContractTransactorSession) UnlockTokens(_account common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.UnlockTokens(&_KNWTokenContract.TransactOpts, _account, _label, _numberOfTokens)
}

// KNWVotingABI is the input ABI used to generate the binding from.
const KNWVotingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_POLL_NONCE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newKNWTokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_obsoleteContract\",\"type\":\"address\"}],\"name\":\"removeDitContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"revealVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getNumKNW\",\"outputs\":[{\"name\":\"numKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"revealPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pollMap\",\"outputs\":[{\"name\":\"initiatingContract\",\"type\":\"address\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"name\":\"voteQuorum\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"name\":\"votesUnrevealed\",\"type\":\"uint256\"},{\"name\":\"winningPercentage\",\"type\":\"uint256\"},{\"name\":\"participantsFor\",\"type\":\"uint256\"},{\"name\":\"participantsAgainst\",\"type\":\"uint256\"},{\"name\":\"participantsUnrevealed\",\"type\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"numVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeMap\",\"outputs\":[{\"name\":\"proposersStake\",\"type\":\"uint256\"},{\"name\":\"proposersReward\",\"type\":\"uint256\"},{\"name\":\"returnPool\",\"type\":\"uint256\"},{\"name\":\"rewardPool\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ditCoordinatorAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pollNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeLabel\",\"type\":\"string\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"}],\"name\":\"startPoll\",\"outputs\":[{\"name\":\"pollID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"didReveal\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newContract\",\"type\":\"address\"},{\"name\":\"_majority\",\"type\":\"uint256\"},{\"name\":\"_mintingMethod\",\"type\":\"uint256\"},{\"name\":\"_burningMethod\",\"type\":\"uint256\"}],\"name\":\"addDitContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getGrossStake\",\"outputs\":[{\"name\":\"grossStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"isResolved\",\"outputs\":[{\"name\":\"resolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"resolveVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getNumVotes\",\"outputs\":[{\"name\":\"numVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"getNetStake\",\"outputs\":[{\"name\":\"netStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"resolvePoll\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pollID\",\"type\":\"uint256\"}],\"name\":\"pollEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCoordinatorAddress\",\"type\":\"address\"}],\"name\":\"setCoordinatorAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// KNWVotingBin is the compiled bytecode used for deploying new contracts.
const KNWVotingBin = `0x608060405234801561001057600080fd5b5060006004556132cc806100256000396000f3006080604052600436106101505763ffffffff60e060020a6000350416632173a10f811461015557806326a4e8d21461017c5780632a518d2a1461019f57806334f2f2d2146101c057806337071611146101ea578063441c77c01461020e578063494031831461023a5780636148fed5146102525780637eb2ff521461034c5780637f97e8361461037357806388a0c9291461039757806388d21ff3146103d557806393b17b8e146103ed57806397508f361461041e578063985dbfc514610433578063a4439dc514610448578063a663fce814610460578063aa7ca46414610496578063b6818a52146104ba578063c3eb9708146104e4578063cc4e1954146104fc578063ce729fd214610514578063d8476fec1461053b578063d9548e531461055f578063dcfde09214610577578063e74fef371461058f578063ee684830146105a7578063f354b838146105bf575b600080fd5b34801561016157600080fd5b5061016a6105e0565b60408051918252519081900360200190f35b34801561018857600080fd5b5061019d600160a060020a03600435166105e5565b005b3480156101ab57600080fd5b5061019d600160a060020a03600435166106e4565b3480156101cc57600080fd5b5061019d600435600160a060020a036024351660443560643561078d565b3480156101f657600080fd5b5061016a600160a060020a0360043516602435610c66565b34801561021a57600080fd5b50610226600435610c94565b604080519115158252519081900360200190f35b34801561024657600080fd5b50610226600435610d17565b34801561025e57600080fd5b5061026a600435610ee6565b604051808e600160a060020a0316600160a060020a03168152602001806020018d81526020018c81526020018b81526020018a81526020018981526020018881526020018781526020018681526020018581526020018481526020018315151515815260200182810382528e818151815260200191508051906020019080838360005b838110156103055781810151838201526020016102ed565b50505050905090810190601f1680156103325780820380516001836020036101000a031916815260200191505b509e50505050505050505050505050505060405180910390f35b34801561035857600080fd5b5061016a600435600160a060020a0360243516604435610fe9565b34801561037f57600080fd5b50610226600160a060020a03600435166024356114d3565b3480156103a357600080fd5b506103af600435611551565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156103e157600080fd5b50610226600435611578565b3480156103f957600080fd5b5061040261158d565b60408051600160a060020a039092168252519081900360200190f35b34801561042a57600080fd5b5061016a61159c565b34801561043f57600080fd5b506104026115a2565b34801561045457600080fd5b506102266004356115b1565b34801561046c57600080fd5b5061016a60048035600160a060020a0316906024803590810191013560443560643560843561161b565b3480156104a257600080fd5b50610226600160a060020a0360043516602435611add565b3480156104c657600080fd5b5061019d600160a060020a0360043516602435604435606435611b60565b3480156104f057600080fd5b5061016a600435611c24565b34801561050857600080fd5b50610226600435611c36565b34801561052057600080fd5b5061016a600435602435600160a060020a0360443516611c4e565b34801561054757600080fd5b5061016a600160a060020a03600435166024356126c2565b34801561056b57600080fd5b506102266004356126f0565b34801561058357600080fd5b5061016a6004356126f5565b34801561059b57600080fd5b5061022660043561277a565b3480156105b357600080fd5b50610226600435612c10565b3480156105cb57600080fd5b5061019d600160a060020a0360043516612c7a565b600081565b600160a060020a038116158015906106065750600154600160a060020a0316155b15156106a8576040805160e560020a62461bcd02815260206004820152604e60248201527f4b4e57546f6b656e20616464726573732063616e206f6e6c792062652073657460448201527f2069662069742773206e6f7420656d70747920616e64206861736e277420616c60648201527f7265616479206265656e20736574000000000000000000000000000000000000608482015290519081900360a40190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03938416179182905560028054929093169116179055565b600054600160a060020a0316331461076c576040805160e560020a62461bcd02815260206004820152602560248201527f4f6e6c792074686520646974436f6f7264696e61746f722063616e2063616c6c60448201527f2074686973000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a03166000908152600360205260409020805460ff19169055565b60008481526005602052604081205485903390600160a060020a03168114610801576040805160e560020a62461bcd02815260206004820152603260248201526000805160206132818339815191526044820152600080516020613241833981519152606482015290519081900360840190fd5b61080a87610c94565b1515610860576040805160e560020a62461bcd02815260206004820152601e60248201527f52657665616c20706572696f642068617320746f206265206163746976650000604482015290519081900360640190fd5b6000878152600560209081526040808320600160a060020a038a168452600d0190915290205460ff161515610905576040805160e560020a62461bcd02815260206004820152602760248201527f5061727469636970616e742068617320746f2068617665206120766f7465206360448201527f6f6d6d6974656400000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000878152600560209081526040808320600160a060020a038a168452600d01909152902054610100900460ff16156109ae576040805160e560020a62461bcd02815260206004820152602260248201527f43616e27742072657665616c206120766f7465206d6f7265207468616e206f6e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000878152600560209081526040808320600160a060020a038a168452600d01825291829020600301548251808301899052808401889052835180820385018152606090910193849052805191939092909182918401908083835b60208310610a285780518252601f199092019160209182019101610a09565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916141515610ad6576040805160e560020a62461bcd02815260206004820152603660248201527f43686f69636520616e642053616c74206861766520746f20626520746865207360448201527f616d6520617320696e2074686520766f74656861736800000000000000000000606482015290519081900360840190fd5b6000878152600560208181526040808420600160a060020a038b168552600d8101835290842060020154938b905291905260070154909350610b1e908463ffffffff612d8216565b60008881526005602052604090206007810191909155600b0154610b4990600163ffffffff612d8216565b6000888152600560205260409020600b01556001851415610bcb5760008781526005602081905260409091200154610b87908463ffffffff612e1116565b60008881526005602081905260409091209081019190915560090154610bb490600163ffffffff612e1116565b600088815260056020526040902060090155610c2b565b600087815260056020526040902060060154610bed908463ffffffff612e1116565b60008881526005602052604090206006810191909155600a0154610c1890600163ffffffff612e1116565b6000888152600560205260409020600a01555b50505060009384525050600560209081526040808420600160a060020a039093168452600d90920190529020805461ff001916610100179055565b6000908152600560209081526040808320600160a060020a03949094168352600d9093019052206001015490565b6000610c9f82611578565b1515610ce3576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613261833981519152604482015290519081900360640190fd5b600082815260056020526040902060030154610cfe906126f0565b158015610d115750610d0f826115b1565b155b92915050565b6000610d21613131565b610d2a83612c10565b1515610d80576040805160e560020a62461bcd02815260206004820152601660248201527f506f6c6c2068617320746f206861766520656e64656400000000000000000000604482015290519081900360640190fd5b60008381526005602090815260409182902082516101a0810184528154600160a060020a03168152600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252919492938581019391929190830182828015610e395780601f10610e0e57610100808354040283529160200191610e39565b820191906000526020600020905b815481529060010190602001808311610e1c57829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c820160009054906101000a900460ff16151515158152505090508060c001518160a00151018160800151028160a001516064021191505b50919050565b6005602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252600160a060020a03909216949293909290830182828015610f905780601f10610f6557610100808354040283529160200191610f90565b820191906000526020600020905b815481529060010190602001808311610f7357829003601f168201915b50505050509080600201549080600301549080600401549080600501549080600601549080600701549080600801549080600901549080600a01549080600b01549080600c0160009054906101000a900460ff1690508d565b6000838152600560205260408120548190819086903390600160a060020a03168114611061576040805160e560020a62461bcd02815260206004820152603260248201526000805160206132818339815191526044820152600080516020613241833981519152606482015290519081900360840190fd5b8715156110b8576040805160e560020a62461bcd02815260206004820152601460248201527f706f6c6c49442063616e2774206265207a65726f000000000000000000000000604482015290519081900360640190fd5b6110c1886115b1565b1515611117576040805160e560020a62461bcd02815260206004820152601e60248201527f436f6d6d697420706572696f642068617320746f206265206163746976650000604482015290519081900360640190fd5b61112187896114d3565b15611176576040805160e560020a62461bcd02815260206004820152601f60248201527f43616e277420636f6d6d6974206d6f7265207468616e206f6e6520766f746500604482015290519081900360640190fd5b8515156111cd576040805160e560020a62461bcd02815260206004820152601b60248201527f43616e277420766f746520776974682061207a65726f20686173680000000000604482015290519081900360640190fd5b6000888152600660209081526040808320546002805460059094529382902082517f1d3316d0000000000000000000000000000000000000000000000000000000008152600160a060020a038d81166004830190815260248301958652600193840180546000199581161561010002959095019094169790970460448301819052939b5090941694631d3316d0948d949293919291606490910190849080156112b75780601f1061128c576101008083540402835291602001916112b7565b820191906000526020600020905b81548152906001019060200180831161129a57829003601f168201915b50509350505050602060405180830381600087803b1580156112d857600080fd5b505af11580156112ec573d6000803e3d6000fd5b505050506040513d602081101561130257600080fd5b50516000898152600560209081526040808320600160a060020a038c168452600d01909152902060010181905593506113468464e8d4a5100063ffffffff612e9b16565b73__libraries/SafeMath.sol:SafeMath_______63677342ce90916040518263ffffffff1660e060020a0281526004018082815260200191505060206040518083038186803b15801561139957600080fd5b505af41580156113ad573d6000803e3d6000fd5b505050506040513d60208110156113c357600080fd5b505192506113de8466038d7ea4c6800063ffffffff612e9b16565b83106113fe576113fb8466038d7ea4c6800063ffffffff612e9b16565b92505b6114306114236103e8611417868963ffffffff612f0916565b9063ffffffff612e9b16565b869063ffffffff612e1116565b6000898152600560208181526040808420600160a060020a038d168552600d8101835290842060028101869055600381018c9055805460ff19166001179055928c9052526007015490955061148b908663ffffffff612e1116565b60008981526005602052604090206007810191909155600b01546114b690600163ffffffff612e1116565b6000898152600560205260409020600b0155505050509392505050565b60006114de82611578565b1515611522576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613261833981519152604482015290519081900360640190fd5b506000908152600560209081526040808320600160a060020a03949094168352600d9093019052205460ff1690565b60066020526000908152604090208054600182015460028301546003909301549192909184565b60008115801590610d11575050600454101590565b600054600160a060020a031681565b60045481565b600154600160a060020a031681565b60006115bc82611578565b1515611600576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613261833981519152604482015290519081900360640190fd5b600082815260056020526040902060020154610d0f906126f0565b33600081815260036020526040812054909182918291829160ff1615156001146116b5576040805160e560020a62461bcd02815260206004820152602860248201527f4f6e6c79206120646974436f6e747261637420697320616c6c6f7720746f206360448201527f616c6c2074686973000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6004546116c990600163ffffffff612e1116565b6004556116dc428963ffffffff612e1116565b93506116ee848863ffffffff612e1116565b92506101a06040519081016040528033600160a060020a031681526020018b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437505050928452505050602080820187905260408083018790523360009081526003808452828220015460608501526080840181905260a0840181905260c0840181905260e084018190526101008401819052610120840181905261014084018190526101609093018390526004548352600582529091208251815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390911617815582820151805191926117f0926001850192909101906131a5565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155610120820151816009015561014082015181600a015561016082015181600b015561018082015181600c0160006101000a81548160ff02191690831515021790555090505060806040519081016040528087815260200160008152602001600081526020016000815250600660006004548152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301559050506001600560006004548152602001908152602001600020600d0160008d600160a060020a0316600160a060020a0316815260200190815260200160002060000160026101000a81548160ff021916908315150217905550600260009054906101000a9004600160a060020a0316600160a060020a0316631d3316d08c6005600060045481526020019081526020016000206001016040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200180602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611a295780601f106119fe57610100808354040283529160200191611a29565b820191906000526020600020905b815481529060010190602001808311611a0c57829003601f168201915b50509350505050602060405180830381600087803b158015611a4a57600080fd5b505af1158015611a5e573d6000803e3d6000fd5b505050506040513d6020811015611a7457600080fd5b8101908080519060200190929190505050915081600560006004548152602001908152602001600020600d0160008d600160a060020a0316600160a060020a03168152602001908152602001600020600101819055506004549450505050509695505050505050565b6000611ae882611578565b1515611b2c576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613261833981519152604482015290519081900360640190fd5b506000908152600560209081526040808320600160a060020a03949094168352600d90930190522054610100900460ff1690565b600054600160a060020a03163314611be8576040805160e560020a62461bcd02815260206004820152602560248201527f4f6e6c792074686520646974436f6f7264696e61746f722063616e2063616c6c60448201527f2074686973000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a039093166000908152600360208190526040909120805460ff1916600190811782559181019390935560028301919091550155565b60009081526006602052604090205490565b6000908152600560205260409020600c015460ff1690565b60008381526005602052604081205481908190819087903390600160a060020a03168114611cc8576040805160e560020a62461bcd02815260206004820152603260248201526000805160206132818339815191526044820152600080516020613241833981519152606482015290519081900360840190fd5b6000898152600560205260409020600c81015490955060ff161515611d37576040805160e560020a62461bcd02815260206004820152601760248201527f506f6c6c2068617320746f206265207265736f6c766564000000000000000000604482015290519081900360640190fd5b600160a060020a0387166000908152600d860160205260408120600101541115611e805760028054600160a060020a038981166000818152600d8a0160205260409081902060019081015491517fd950df3400000000000000000000000000000000000000000000000000000000815260048101938452604481018390526060602482019081528c830180546000199481161561010002949094019093169790970460648201819052949095169563d950df34958e9592949260849091019085908015611e455780601f10611e1a57610100808354040283529160200191611e45565b820191906000526020600020905b815481529060010190602001808311611e2857829003601f168201915b5050945050505050600060405180830381600087803b158015611e6757600080fd5b505af1158015611e7b573d6000803e3d6000fd5b505050505b611e8989610d17565b935083611e97576000611e9a565b60015b600160a060020a0388166000908152600d8701602052604090205460ff9182168a14945062010000900416156121d857831561200c576002805460088701548754600160a060020a039081166000908152600360205260409081902085015490517f4a5dc3aa0000000000000000000000000000000000000000000000000000000081528c831660048201908152604482018590526064820183905260806024830190815260018d8101805460001992811615610100029290920190911698909804608484018190529490961696634a5dc3aa968f969195919493909160a49091019086908015611fcc5780601f10611fa157610100808354040283529160200191611fcc565b820191906000526020600020905b815481529060010190602001808311611faf57829003601f168201915b505095505050505050600060405180830381600087803b158015611fef57600080fd5b505af1158015612003573d6000803e3d6000fd5b505050506121bf565b60008981526006602052604090206001015415156121bf57600260009054906101000a9004600160a060020a0316600160a060020a031663994dd93a888760010188600d0160008c600160a060020a0316600160a060020a03168152602001908152602001600020600101548960080154600360008c60000160009054906101000a9004600160a060020a0316600160a060020a0316600160a060020a03168152602001908152602001600020600101546040518663ffffffff1660e060020a0281526004018086600160a060020a0316600160a060020a03168152602001806020018581526020018481526020018381526020018281038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156121825780601f1061215757610100808354040283529160200191612182565b820191906000526020600020905b81548152906001019060200180831161216557829003601f168201915b50509650505050505050600060405180830381600087803b1580156121a657600080fd5b505af11580156121ba573d6000803e3d6000fd5b505050505b60008981526006602052604090206001015495506126b6565b6121e2878a611add565b1561250457831580156121fc575084600601548560050154145b156122145761220d89846001612f82565b95506124ff565b61222089846000612f82565b95508215612364576002805460088701548754600160a060020a039081166000908152600360205260409081902085015490517f4a5dc3aa0000000000000000000000000000000000000000000000000000000081528c831660048201908152604482018590526064820183905260806024830190815260018d8101805460001992811615610100029290920190911698909804608484018190529490961696634a5dc3aa968f969195919493909160a490910190869080156123245780601f106122f957610100808354040283529160200191612324565b820191906000526020600020905b81548152906001019060200180831161230757829003601f168201915b505095505050505050600060405180830381600087803b15801561234757600080fd5b505af115801561235b573d6000803e3d6000fd5b505050506124ff565b600260009054906101000a9004600160a060020a0316600160a060020a031663994dd93a888760010188600d0160008c600160a060020a0316600160a060020a03168152602001908152602001600020600101548960080154600360008c60000160009054906101000a9004600160a060020a0316600160a060020a0316600160a060020a03168152602001908152602001600020600101546040518663ffffffff1660e060020a0281526004018086600160a060020a0316600160a060020a03168152602001806020018581526020018481526020018381526020018281038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156124c25780601f10612497576101008083540402835291602001916124c2565b820191906000526020600020905b8154815290600101906020018083116124a557829003601f168201915b50509650505050505050600060405180830381600087803b1580156124e657600080fd5b505af11580156124fa573d6000803e3d6000fd5b505050505b6126b6565b61250e878a611add565b1580156125205750612520878a6114d3565b156126665761253189600080612f82565b9550600260009054906101000a9004600160a060020a0316600160a060020a031663994dd93a888760010188600d0160008c600160a060020a0316600160a060020a03168152602001908152602001600020600101548960080154600360008c60000160009054906101000a9004600160a060020a0316600160a060020a0316600160a060020a03168152602001908152602001600020600101546040518663ffffffff1660e060020a0281526004018086600160a060020a0316600160a060020a03168152602001806020018581526020018481526020018381526020018281038252868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156124c25780601f10612497576101008083540402835291602001916124c2565b6040805160e560020a62461bcd02815260206004820152601d60248201527f4e6f742061207061727469636970616e74206f662074686520766f7465000000604482015290519081900360640190fd5b50505050509392505050565b6000908152600560209081526040808320600160a060020a03949094168352600d9093019052206002015490565b421190565b6000818152600560205260408120600b8101546009820154600a90920154839261273592916127299163ffffffff612e1116565b9063ffffffff612e1116565b905060008111156127665760008381526006602052604090205461275f908263ffffffff612e9b16565b9150610ee0565b505060009081526006602052604090205490565b600081815260056020526040812054819081908190819086903390600160a060020a031681146127f6576040805160e560020a62461bcd02815260206004820152603260248201526000805160206132818339815191526044820152600080516020613241833981519152606482015290519081900360840190fd5b6127ff88612c10565b1515612855576040805160e560020a62461bcd02815260206004820152601660248201527f506f6c6c2068617320746f206861766520656e64656400000000000000000000604482015290519081900360640190fd5b6000888152600560208190526040909120600681015491015461287d9163ffffffff612e1116565b6000898152600560205260409020600b8101546009820154600a909201549298506128b4929091612729919063ffffffff612e1116565b94508415156128f957600088815260066020908152604080832080546001918201556005909252822060088101839055600c01805460ff191690911790559650612c05565b6000888152600660205260409020546129489061293b90612920908863ffffffff612e9b16565b60008b8152600660205260409020549063ffffffff612d8216565b869063ffffffff612f0916565b600089815260066020526040812060020191909155935061296888610d17565b965086156129e7576000888152600560208190526040909120015461299a90879061141790606463ffffffff612f0916565b60008981526005602052604090206008810191909155600b810154600a909101546129ca9163ffffffff612e1116565b600089815260066020526040902080546001909101559350612aaa565b6000888152600560208190526040909120600681015491015414612a7b57600088815260056020526040902060060154612a2e90879061141790606463ffffffff612f0916565b60008981526005602052604090206008810191909155600b810154600990910154612a5e9163ffffffff612e1116565b600089815260066020526040902080546003909101559350612aaa565b600088815260056020908152604080832060326008820155600b01546006909252909120805460019091015593505b6000888152600660205260408120600201541115612b3c57600088815260066020526040902060020154612b2990612b0b9061293b90612af0908963ffffffff612e9b16565b60008c8152600660205260409020549063ffffffff612d8216565b60008a8152600660205260409020600301549063ffffffff612e1116565b6000898152600660205260409020600301555b6000888152600560209081526040808320600c01805460ff191660019081179091556006909252822001541115612c0557612ba3612b856001612729888863ffffffff612d8216565b60008a8152600660205260409020600301549063ffffffff612e9b16565b600089815260066020526040902060010154909350612bc8908463ffffffff612e1116565b6000898152600660205260409020600181019190915560030154612bf2908463ffffffff612d8216565b6000898152600660205260409020600301555b505050505050919050565b6000612c1b82611578565b1515612c5f576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613261833981519152604482015290519081900360640190fd5b600082815260056020526040902060030154610d11906126f0565b600160a060020a03811615801590612c9b5750600054600160a060020a0316155b1515612d3d576040805160e560020a62461bcd02815260206004820152605460248201527f646974436f6f7264696e61746f7220616464726573732063616e206f6e6c792060448201527f6265207365742069662069742773206e6f7420656d70747920616e642068617360648201527f6e277420616c7265616479206265656e20736574000000000000000000000000608482015290519081900360a40190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03928316178082559091168152600360205260409020805460ff19166001179055565b60008083831115612e03576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050808203805b5092915050565b600082820183811015612e94576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b600080808311612ef5576040805160e560020a62461bcd02815260206004820152601460248201527f43616e277420646976696465206279207a65726f000000000000000000000000604482015290519081900360640190fd5b8284811515612f0057fe5b04949350505050565b600080831515612f1c5760009150612e0a565b50828202828482811515612f2c57fe5b0414612e94576040805160e560020a62461bcd02815260206004820152601f60248201527f466c6177656420696e70757420666f72206d756c7469706c69636174696f6e00604482015290519081900360640190fd5b6000838152600560205260408120600b8101546009820154600a909201548392839283928392612fbc92916127299163ffffffff612e1116565b93506000925085151561306d57600088815260066020526040902060020154612feb908563ffffffff612e9b16565b915081925086156130685750600061300288610d17565b1561301f5750600087815260056020526040902060090154613033565b506000878152600560205260409020600a01545b60008881526006602052604090206003015461306590613059908363ffffffff612e9b16565b6127296114238b6126f5565b92505b613125565b60008881526006602090815260408083205460059092528220600b01549094501180156130aa57506000888152600560205260409020600b015484115b15613125576000888152600560205260409020600b015461312290613107906130da90879063ffffffff612d8216565b60008b8152600660209081526040808320546005909252909120600b01546114179163ffffffff612f0916565b60008a8152600660205260409020549063ffffffff612e1116565b92505b50909695505050505050565b6101a0604051908101604052806000600160a060020a0316815260200160608152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106131e657805160ff1916838001178555613213565b82800160010185558215613213579182015b828111156132135782518255916020019190600101906131f8565b5061321f929150613223565b5090565b61323d91905b8082111561321f5760008155600101613229565b905600616c6c6f7720746f2063616c6c20746869730000000000000000000000000000506f6c6c2068617320746f2065786973740000000000000000000000000000004f6e6c792074686520696e6974696174696e6720636f6e747261637420697320a165627a7a72305820fe5d84b0fb3c3aeb2d72da544557e2c71d33ca941c59f1459b2923bf0752c4ff0029`

// DeployKNWVoting deploys a new Ethereum contract, binding an instance of KNWVoting to it.
func DeployKNWVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KNWVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWVoting{KNWVotingCaller: KNWVotingCaller{contract: contract}, KNWVotingTransactor: KNWVotingTransactor{contract: contract}, KNWVotingFilterer: KNWVotingFilterer{contract: contract}}, nil
}

// KNWVoting is an auto generated Go binding around an Ethereum contract.
type KNWVoting struct {
	KNWVotingCaller     // Read-only binding to the contract
	KNWVotingTransactor // Write-only binding to the contract
	KNWVotingFilterer   // Log filterer for contract events
}

// KNWVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWVotingSession struct {
	Contract     *KNWVoting        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNWVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWVotingCallerSession struct {
	Contract *KNWVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KNWVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWVotingTransactorSession struct {
	Contract     *KNWVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KNWVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWVotingRaw struct {
	Contract *KNWVoting // Generic contract binding to access the raw methods on
}

// KNWVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWVotingCallerRaw struct {
	Contract *KNWVotingCaller // Generic read-only contract binding to access the raw methods on
}

// KNWVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWVotingTransactorRaw struct {
	Contract *KNWVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWVoting creates a new instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVoting(address common.Address, backend bind.ContractBackend) (*KNWVoting, error) {
	contract, err := bindKNWVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWVoting{KNWVotingCaller: KNWVotingCaller{contract: contract}, KNWVotingTransactor: KNWVotingTransactor{contract: contract}, KNWVotingFilterer: KNWVotingFilterer{contract: contract}}, nil
}

// NewKNWVotingCaller creates a new read-only instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVotingCaller(address common.Address, caller bind.ContractCaller) (*KNWVotingCaller, error) {
	contract, err := bindKNWVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingCaller{contract: contract}, nil
}

// NewKNWVotingTransactor creates a new write-only instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWVotingTransactor, error) {
	contract, err := bindKNWVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingTransactor{contract: contract}, nil
}

// NewKNWVotingFilterer creates a new log filterer instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWVotingFilterer, error) {
	contract, err := bindKNWVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWVotingFilterer{contract: contract}, nil
}

// bindKNWVoting binds a generic wrapper to an already deployed contract.
func bindKNWVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVoting *KNWVotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVoting.Contract.KNWVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVoting *KNWVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVoting.Contract.KNWVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVoting *KNWVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVoting.Contract.KNWVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVoting *KNWVotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVoting *KNWVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVoting *KNWVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVoting.Contract.contract.Transact(opts, method, params...)
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) INITIALPOLLNONCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "INITIAL_POLL_NONCE")
	return *ret0, err
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) INITIALPOLLNONCE() (*big.Int, error) {
	return _KNWVoting.Contract.INITIALPOLLNONCE(&_KNWVoting.CallOpts)
}

// INITIALPOLLNONCE is a free data retrieval call binding the contract method 0x2173a10f.
//
// Solidity: function INITIAL_POLL_NONCE() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) INITIALPOLLNONCE() (*big.Int, error) {
	return _KNWVoting.Contract.INITIALPOLLNONCE(&_KNWVoting.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_KNWVoting *KNWVotingCaller) KNWTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "KNWTokenAddress")
	return *ret0, err
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_KNWVoting *KNWVotingSession) KNWTokenAddress() (common.Address, error) {
	return _KNWVoting.Contract.KNWTokenAddress(&_KNWVoting.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) KNWTokenAddress() (common.Address, error) {
	return _KNWVoting.Contract.KNWTokenAddress(&_KNWVoting.CallOpts)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_KNWVoting *KNWVotingCaller) CommitPeriodActive(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "commitPeriodActive", _pollID)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_KNWVoting *KNWVotingSession) CommitPeriodActive(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.CommitPeriodActive(&_KNWVoting.CallOpts, _pollID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(_pollID uint256) constant returns(active bool)
func (_KNWVoting *KNWVotingCallerSession) CommitPeriodActive(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.CommitPeriodActive(&_KNWVoting.CallOpts, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_address address, _pollID uint256) constant returns(committed bool)
func (_KNWVoting *KNWVotingCaller) DidCommit(opts *bind.CallOpts, _address common.Address, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "didCommit", _address, _pollID)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_address address, _pollID uint256) constant returns(committed bool)
func (_KNWVoting *KNWVotingSession) DidCommit(_address common.Address, _pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidCommit(&_KNWVoting.CallOpts, _address, _pollID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(_address address, _pollID uint256) constant returns(committed bool)
func (_KNWVoting *KNWVotingCallerSession) DidCommit(_address common.Address, _pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidCommit(&_KNWVoting.CallOpts, _address, _pollID)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_address address, _pollID uint256) constant returns(revealed bool)
func (_KNWVoting *KNWVotingCaller) DidReveal(opts *bind.CallOpts, _address common.Address, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "didReveal", _address, _pollID)
	return *ret0, err
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_address address, _pollID uint256) constant returns(revealed bool)
func (_KNWVoting *KNWVotingSession) DidReveal(_address common.Address, _pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidReveal(&_KNWVoting.CallOpts, _address, _pollID)
}

// DidReveal is a free data retrieval call binding the contract method 0xaa7ca464.
//
// Solidity: function didReveal(_address address, _pollID uint256) constant returns(revealed bool)
func (_KNWVoting *KNWVotingCallerSession) DidReveal(_address common.Address, _pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidReveal(&_KNWVoting.CallOpts, _address, _pollID)
}

// DitCoordinatorAddress is a free data retrieval call binding the contract method 0x93b17b8e.
//
// Solidity: function ditCoordinatorAddress() constant returns(address)
func (_KNWVoting *KNWVotingCaller) DitCoordinatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "ditCoordinatorAddress")
	return *ret0, err
}

// DitCoordinatorAddress is a free data retrieval call binding the contract method 0x93b17b8e.
//
// Solidity: function ditCoordinatorAddress() constant returns(address)
func (_KNWVoting *KNWVotingSession) DitCoordinatorAddress() (common.Address, error) {
	return _KNWVoting.Contract.DitCoordinatorAddress(&_KNWVoting.CallOpts)
}

// DitCoordinatorAddress is a free data retrieval call binding the contract method 0x93b17b8e.
//
// Solidity: function ditCoordinatorAddress() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) DitCoordinatorAddress() (common.Address, error) {
	return _KNWVoting.Contract.DitCoordinatorAddress(&_KNWVoting.CallOpts)
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(_pollID uint256) constant returns(grossStake uint256)
func (_KNWVoting *KNWVotingCaller) GetGrossStake(opts *bind.CallOpts, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getGrossStake", _pollID)
	return *ret0, err
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(_pollID uint256) constant returns(grossStake uint256)
func (_KNWVoting *KNWVotingSession) GetGrossStake(_pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetGrossStake(&_KNWVoting.CallOpts, _pollID)
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(_pollID uint256) constant returns(grossStake uint256)
func (_KNWVoting *KNWVotingCallerSession) GetGrossStake(_pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetGrossStake(&_KNWVoting.CallOpts, _pollID)
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(_pollID uint256) constant returns(netStake uint256)
func (_KNWVoting *KNWVotingCaller) GetNetStake(opts *bind.CallOpts, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getNetStake", _pollID)
	return *ret0, err
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(_pollID uint256) constant returns(netStake uint256)
func (_KNWVoting *KNWVotingSession) GetNetStake(_pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNetStake(&_KNWVoting.CallOpts, _pollID)
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(_pollID uint256) constant returns(netStake uint256)
func (_KNWVoting *KNWVotingCallerSession) GetNetStake(_pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNetStake(&_KNWVoting.CallOpts, _pollID)
}

// GetNumKNW is a free data retrieval call binding the contract method 0x37071611.
//
// Solidity: function getNumKNW(_address address, _pollID uint256) constant returns(numKNW uint256)
func (_KNWVoting *KNWVotingCaller) GetNumKNW(opts *bind.CallOpts, _address common.Address, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getNumKNW", _address, _pollID)
	return *ret0, err
}

// GetNumKNW is a free data retrieval call binding the contract method 0x37071611.
//
// Solidity: function getNumKNW(_address address, _pollID uint256) constant returns(numKNW uint256)
func (_KNWVoting *KNWVotingSession) GetNumKNW(_address common.Address, _pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNumKNW(&_KNWVoting.CallOpts, _address, _pollID)
}

// GetNumKNW is a free data retrieval call binding the contract method 0x37071611.
//
// Solidity: function getNumKNW(_address address, _pollID uint256) constant returns(numKNW uint256)
func (_KNWVoting *KNWVotingCallerSession) GetNumKNW(_address common.Address, _pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNumKNW(&_KNWVoting.CallOpts, _address, _pollID)
}

// GetNumVotes is a free data retrieval call binding the contract method 0xd8476fec.
//
// Solidity: function getNumVotes(_address address, _pollID uint256) constant returns(numVotes uint256)
func (_KNWVoting *KNWVotingCaller) GetNumVotes(opts *bind.CallOpts, _address common.Address, _pollID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getNumVotes", _address, _pollID)
	return *ret0, err
}

// GetNumVotes is a free data retrieval call binding the contract method 0xd8476fec.
//
// Solidity: function getNumVotes(_address address, _pollID uint256) constant returns(numVotes uint256)
func (_KNWVoting *KNWVotingSession) GetNumVotes(_address common.Address, _pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNumVotes(&_KNWVoting.CallOpts, _address, _pollID)
}

// GetNumVotes is a free data retrieval call binding the contract method 0xd8476fec.
//
// Solidity: function getNumVotes(_address address, _pollID uint256) constant returns(numVotes uint256)
func (_KNWVoting *KNWVotingCallerSession) GetNumVotes(_address common.Address, _pollID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNumVotes(&_KNWVoting.CallOpts, _address, _pollID)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_KNWVoting *KNWVotingCaller) IsExpired(opts *bind.CallOpts, _terminationDate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isExpired", _terminationDate)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_KNWVoting *KNWVotingSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsExpired(&_KNWVoting.CallOpts, _terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(_terminationDate uint256) constant returns(expired bool)
func (_KNWVoting *KNWVotingCallerSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsExpired(&_KNWVoting.CallOpts, _terminationDate)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_KNWVoting *KNWVotingCaller) IsPassed(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isPassed", _pollID)
	return *ret0, err
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_KNWVoting *KNWVotingSession) IsPassed(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsPassed(&_KNWVoting.CallOpts, _pollID)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(_pollID uint256) constant returns(passed bool)
func (_KNWVoting *KNWVotingCallerSession) IsPassed(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsPassed(&_KNWVoting.CallOpts, _pollID)
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(_pollID uint256) constant returns(resolved bool)
func (_KNWVoting *KNWVotingCaller) IsResolved(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isResolved", _pollID)
	return *ret0, err
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(_pollID uint256) constant returns(resolved bool)
func (_KNWVoting *KNWVotingSession) IsResolved(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsResolved(&_KNWVoting.CallOpts, _pollID)
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(_pollID uint256) constant returns(resolved bool)
func (_KNWVoting *KNWVotingCallerSession) IsResolved(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsResolved(&_KNWVoting.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_KNWVoting *KNWVotingCaller) PollEnded(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "pollEnded", _pollID)
	return *ret0, err
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_KNWVoting *KNWVotingSession) PollEnded(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.PollEnded(&_KNWVoting.CallOpts, _pollID)
}

// PollEnded is a free data retrieval call binding the contract method 0xee684830.
//
// Solidity: function pollEnded(_pollID uint256) constant returns(ended bool)
func (_KNWVoting *KNWVotingCallerSession) PollEnded(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.PollEnded(&_KNWVoting.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_KNWVoting *KNWVotingCaller) PollExists(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "pollExists", _pollID)
	return *ret0, err
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_KNWVoting *KNWVotingSession) PollExists(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.PollExists(&_KNWVoting.CallOpts, _pollID)
}

// PollExists is a free data retrieval call binding the contract method 0x88d21ff3.
//
// Solidity: function pollExists(_pollID uint256) constant returns(exists bool)
func (_KNWVoting *KNWVotingCallerSession) PollExists(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.PollExists(&_KNWVoting.CallOpts, _pollID)
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(initiatingContract address, knowledgeLabel string, commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256, votesUnrevealed uint256, winningPercentage uint256, participantsFor uint256, participantsAgainst uint256, participantsUnrevealed uint256, isResolved bool)
func (_KNWVoting *KNWVotingCaller) PollMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InitiatingContract     common.Address
	KnowledgeLabel         string
	CommitEndDate          *big.Int
	RevealEndDate          *big.Int
	VoteQuorum             *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	WinningPercentage      *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	ret := new(struct {
		InitiatingContract     common.Address
		KnowledgeLabel         string
		CommitEndDate          *big.Int
		RevealEndDate          *big.Int
		VoteQuorum             *big.Int
		VotesFor               *big.Int
		VotesAgainst           *big.Int
		VotesUnrevealed        *big.Int
		WinningPercentage      *big.Int
		ParticipantsFor        *big.Int
		ParticipantsAgainst    *big.Int
		ParticipantsUnrevealed *big.Int
		IsResolved             bool
	})
	out := ret
	err := _KNWVoting.contract.Call(opts, out, "pollMap", arg0)
	return *ret, err
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(initiatingContract address, knowledgeLabel string, commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256, votesUnrevealed uint256, winningPercentage uint256, participantsFor uint256, participantsAgainst uint256, participantsUnrevealed uint256, isResolved bool)
func (_KNWVoting *KNWVotingSession) PollMap(arg0 *big.Int) (struct {
	InitiatingContract     common.Address
	KnowledgeLabel         string
	CommitEndDate          *big.Int
	RevealEndDate          *big.Int
	VoteQuorum             *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	WinningPercentage      *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	return _KNWVoting.Contract.PollMap(&_KNWVoting.CallOpts, arg0)
}

// PollMap is a free data retrieval call binding the contract method 0x6148fed5.
//
// Solidity: function pollMap( uint256) constant returns(initiatingContract address, knowledgeLabel string, commitEndDate uint256, revealEndDate uint256, voteQuorum uint256, votesFor uint256, votesAgainst uint256, votesUnrevealed uint256, winningPercentage uint256, participantsFor uint256, participantsAgainst uint256, participantsUnrevealed uint256, isResolved bool)
func (_KNWVoting *KNWVotingCallerSession) PollMap(arg0 *big.Int) (struct {
	InitiatingContract     common.Address
	KnowledgeLabel         string
	CommitEndDate          *big.Int
	RevealEndDate          *big.Int
	VoteQuorum             *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	WinningPercentage      *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	return _KNWVoting.Contract.PollMap(&_KNWVoting.CallOpts, arg0)
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) PollNonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "pollNonce")
	return *ret0, err
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) PollNonce() (*big.Int, error) {
	return _KNWVoting.Contract.PollNonce(&_KNWVoting.CallOpts)
}

// PollNonce is a free data retrieval call binding the contract method 0x97508f36.
//
// Solidity: function pollNonce() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) PollNonce() (*big.Int, error) {
	return _KNWVoting.Contract.PollNonce(&_KNWVoting.CallOpts)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_KNWVoting *KNWVotingCaller) RevealPeriodActive(opts *bind.CallOpts, _pollID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "revealPeriodActive", _pollID)
	return *ret0, err
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_KNWVoting *KNWVotingSession) RevealPeriodActive(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.RevealPeriodActive(&_KNWVoting.CallOpts, _pollID)
}

// RevealPeriodActive is a free data retrieval call binding the contract method 0x441c77c0.
//
// Solidity: function revealPeriodActive(_pollID uint256) constant returns(active bool)
func (_KNWVoting *KNWVotingCallerSession) RevealPeriodActive(_pollID *big.Int) (bool, error) {
	return _KNWVoting.Contract.RevealPeriodActive(&_KNWVoting.CallOpts, _pollID)
}

// StakeMap is a free data retrieval call binding the contract method 0x88a0c929.
//
// Solidity: function stakeMap( uint256) constant returns(proposersStake uint256, proposersReward uint256, returnPool uint256, rewardPool uint256)
func (_KNWVoting *KNWVotingCaller) StakeMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	ret := new(struct {
		ProposersStake  *big.Int
		ProposersReward *big.Int
		ReturnPool      *big.Int
		RewardPool      *big.Int
	})
	out := ret
	err := _KNWVoting.contract.Call(opts, out, "stakeMap", arg0)
	return *ret, err
}

// StakeMap is a free data retrieval call binding the contract method 0x88a0c929.
//
// Solidity: function stakeMap( uint256) constant returns(proposersStake uint256, proposersReward uint256, returnPool uint256, rewardPool uint256)
func (_KNWVoting *KNWVotingSession) StakeMap(arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	return _KNWVoting.Contract.StakeMap(&_KNWVoting.CallOpts, arg0)
}

// StakeMap is a free data retrieval call binding the contract method 0x88a0c929.
//
// Solidity: function stakeMap( uint256) constant returns(proposersStake uint256, proposersReward uint256, returnPool uint256, rewardPool uint256)
func (_KNWVoting *KNWVotingCallerSession) StakeMap(arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	return _KNWVoting.Contract.StakeMap(&_KNWVoting.CallOpts, arg0)
}

// AddDitContract is a paid mutator transaction binding the contract method 0xb6818a52.
//
// Solidity: function addDitContract(_newContract address, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVoting *KNWVotingTransactor) AddDitContract(opts *bind.TransactOpts, _newContract common.Address, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "addDitContract", _newContract, _majority, _mintingMethod, _burningMethod)
}

// AddDitContract is a paid mutator transaction binding the contract method 0xb6818a52.
//
// Solidity: function addDitContract(_newContract address, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVoting *KNWVotingSession) AddDitContract(_newContract common.Address, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddDitContract(&_KNWVoting.TransactOpts, _newContract, _majority, _mintingMethod, _burningMethod)
}

// AddDitContract is a paid mutator transaction binding the contract method 0xb6818a52.
//
// Solidity: function addDitContract(_newContract address, _majority uint256, _mintingMethod uint256, _burningMethod uint256) returns()
func (_KNWVoting *KNWVotingTransactorSession) AddDitContract(_newContract common.Address, _majority *big.Int, _mintingMethod *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddDitContract(&_KNWVoting.TransactOpts, _newContract, _majority, _mintingMethod, _burningMethod)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(_pollID uint256, _address address, _secretHash bytes32) returns(numVotes uint256)
func (_KNWVoting *KNWVotingTransactor) CommitVote(opts *bind.TransactOpts, _pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "commitVote", _pollID, _address, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(_pollID uint256, _address address, _secretHash bytes32) returns(numVotes uint256)
func (_KNWVoting *KNWVotingSession) CommitVote(_pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVoting.Contract.CommitVote(&_KNWVoting.TransactOpts, _pollID, _address, _secretHash)
}

// CommitVote is a paid mutator transaction binding the contract method 0x7eb2ff52.
//
// Solidity: function commitVote(_pollID uint256, _address address, _secretHash bytes32) returns(numVotes uint256)
func (_KNWVoting *KNWVotingTransactorSession) CommitVote(_pollID *big.Int, _address common.Address, _secretHash [32]byte) (*types.Transaction, error) {
	return _KNWVoting.Contract.CommitVote(&_KNWVoting.TransactOpts, _pollID, _address, _secretHash)
}

// RemoveDitContract is a paid mutator transaction binding the contract method 0x2a518d2a.
//
// Solidity: function removeDitContract(_obsoleteContract address) returns()
func (_KNWVoting *KNWVotingTransactor) RemoveDitContract(opts *bind.TransactOpts, _obsoleteContract common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "removeDitContract", _obsoleteContract)
}

// RemoveDitContract is a paid mutator transaction binding the contract method 0x2a518d2a.
//
// Solidity: function removeDitContract(_obsoleteContract address) returns()
func (_KNWVoting *KNWVotingSession) RemoveDitContract(_obsoleteContract common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.RemoveDitContract(&_KNWVoting.TransactOpts, _obsoleteContract)
}

// RemoveDitContract is a paid mutator transaction binding the contract method 0x2a518d2a.
//
// Solidity: function removeDitContract(_obsoleteContract address) returns()
func (_KNWVoting *KNWVotingTransactorSession) RemoveDitContract(_obsoleteContract common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.RemoveDitContract(&_KNWVoting.TransactOpts, _obsoleteContract)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(_pollID uint256) returns(votePassed bool)
func (_KNWVoting *KNWVotingTransactor) ResolvePoll(opts *bind.TransactOpts, _pollID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "resolvePoll", _pollID)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(_pollID uint256) returns(votePassed bool)
func (_KNWVoting *KNWVotingSession) ResolvePoll(_pollID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.ResolvePoll(&_KNWVoting.TransactOpts, _pollID)
}

// ResolvePoll is a paid mutator transaction binding the contract method 0xe74fef37.
//
// Solidity: function resolvePoll(_pollID uint256) returns(votePassed bool)
func (_KNWVoting *KNWVotingTransactorSession) ResolvePoll(_pollID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.ResolvePoll(&_KNWVoting.TransactOpts, _pollID)
}

// ResolveVote is a paid mutator transaction binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(_pollID uint256, _voteOption uint256, _address address) returns(reward uint256)
func (_KNWVoting *KNWVotingTransactor) ResolveVote(opts *bind.TransactOpts, _pollID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "resolveVote", _pollID, _voteOption, _address)
}

// ResolveVote is a paid mutator transaction binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(_pollID uint256, _voteOption uint256, _address address) returns(reward uint256)
func (_KNWVoting *KNWVotingSession) ResolveVote(_pollID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.ResolveVote(&_KNWVoting.TransactOpts, _pollID, _voteOption, _address)
}

// ResolveVote is a paid mutator transaction binding the contract method 0xce729fd2.
//
// Solidity: function resolveVote(_pollID uint256, _voteOption uint256, _address address) returns(reward uint256)
func (_KNWVoting *KNWVotingTransactorSession) ResolveVote(_pollID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.ResolveVote(&_KNWVoting.TransactOpts, _pollID, _voteOption, _address)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(_pollID uint256, _address address, _voteOption uint256, _salt uint256) returns()
func (_KNWVoting *KNWVotingTransactor) RevealVote(opts *bind.TransactOpts, _pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "revealVote", _pollID, _address, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(_pollID uint256, _address address, _voteOption uint256, _salt uint256) returns()
func (_KNWVoting *KNWVotingSession) RevealVote(_pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.RevealVote(&_KNWVoting.TransactOpts, _pollID, _address, _voteOption, _salt)
}

// RevealVote is a paid mutator transaction binding the contract method 0x34f2f2d2.
//
// Solidity: function revealVote(_pollID uint256, _address address, _voteOption uint256, _salt uint256) returns()
func (_KNWVoting *KNWVotingTransactorSession) RevealVote(_pollID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.RevealVote(&_KNWVoting.TransactOpts, _pollID, _address, _voteOption, _salt)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(_newCoordinatorAddress address) returns()
func (_KNWVoting *KNWVotingTransactor) SetCoordinatorAddress(opts *bind.TransactOpts, _newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "setCoordinatorAddress", _newCoordinatorAddress)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(_newCoordinatorAddress address) returns()
func (_KNWVoting *KNWVotingSession) SetCoordinatorAddress(_newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.SetCoordinatorAddress(&_KNWVoting.TransactOpts, _newCoordinatorAddress)
}

// SetCoordinatorAddress is a paid mutator transaction binding the contract method 0xf354b838.
//
// Solidity: function setCoordinatorAddress(_newCoordinatorAddress address) returns()
func (_KNWVoting *KNWVotingTransactorSession) SetCoordinatorAddress(_newCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.SetCoordinatorAddress(&_KNWVoting.TransactOpts, _newCoordinatorAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_newKNWTokenAddress address) returns()
func (_KNWVoting *KNWVotingTransactor) SetTokenAddress(opts *bind.TransactOpts, _newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "setTokenAddress", _newKNWTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_newKNWTokenAddress address) returns()
func (_KNWVoting *KNWVotingSession) SetTokenAddress(_newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.SetTokenAddress(&_KNWVoting.TransactOpts, _newKNWTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_newKNWTokenAddress address) returns()
func (_KNWVoting *KNWVotingTransactorSession) SetTokenAddress(_newKNWTokenAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.SetTokenAddress(&_KNWVoting.TransactOpts, _newKNWTokenAddress)
}

// StartPoll is a paid mutator transaction binding the contract method 0xa663fce8.
//
// Solidity: function startPoll(_address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVoting *KNWVotingTransactor) StartPoll(opts *bind.TransactOpts, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "startPoll", _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0xa663fce8.
//
// Solidity: function startPoll(_address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVoting *KNWVotingSession) StartPoll(_address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.StartPoll(&_KNWVoting.TransactOpts, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
}

// StartPoll is a paid mutator transaction binding the contract method 0xa663fce8.
//
// Solidity: function startPoll(_address address, _knowledgeLabel string, _commitDuration uint256, _revealDuration uint256, _proposersStake uint256) returns(pollID uint256)
func (_KNWVoting *KNWVotingTransactorSession) StartPoll(_address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.StartPoll(&_KNWVoting.TransactOpts, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake)
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
