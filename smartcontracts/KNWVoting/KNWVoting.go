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
const KNWTokenContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"freeBalanceOfLabel\",\"outputs\":[{\"name\":\"freeBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"balanceOfLabel\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
// Solidity: function balanceOfLabel(address _account, string _label) constant returns(uint256 balance)
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
// Solidity: function balanceOfLabel(address _account, string _label) constant returns(uint256 balance)
func (_KNWTokenContract *KNWTokenContractSession) BalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.BalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(address _account, string _label) constant returns(uint256 balance)
func (_KNWTokenContract *KNWTokenContractCallerSession) BalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.BalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(address _account, string _label) constant returns(uint256 freeBalance)
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
// Solidity: function freeBalanceOfLabel(address _account, string _label) constant returns(uint256 freeBalance)
func (_KNWTokenContract *KNWTokenContractSession) FreeBalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.FreeBalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(address _account, string _label) constant returns(uint256 freeBalance)
func (_KNWTokenContract *KNWTokenContractCallerSession) FreeBalanceOfLabel(_account common.Address, _label string) (*big.Int, error) {
	return _KNWTokenContract.Contract.FreeBalanceOfLabel(&_KNWTokenContract.CallOpts, _account, _label)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "burn", _account, _label, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) Burn(_account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Burn(&_KNWTokenContract.TransactOpts, _account, _label, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) Burn(_account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Burn(&_KNWTokenContract.TransactOpts, _account, _label, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x73599798.
//
// Solidity: function lockTokens(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) LockTokens(opts *bind.TransactOpts, _account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "lockTokens", _account, _label, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x73599798.
//
// Solidity: function lockTokens(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) LockTokens(_account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.LockTokens(&_KNWTokenContract.TransactOpts, _account, _label, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x73599798.
//
// Solidity: function lockTokens(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) LockTokens(_account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.LockTokens(&_KNWTokenContract.TransactOpts, _account, _label, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "mint", _account, _label, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) Mint(_account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Mint(&_KNWTokenContract.TransactOpts, _account, _label, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address _account, string _label, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) Mint(_account common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Mint(&_KNWTokenContract.TransactOpts, _account, _label, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(address _account, string _label, uint256 _numberOfTokens) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) UnlockTokens(opts *bind.TransactOpts, _account common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "unlockTokens", _account, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(address _account, string _label, uint256 _numberOfTokens) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) UnlockTokens(_account common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.UnlockTokens(&_KNWTokenContract.TransactOpts, _account, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(address _account, string _label, uint256 _numberOfTokens) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) UnlockTokens(_account common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.UnlockTokens(&_KNWTokenContract.TransactOpts, _account, _label, _numberOfTokens)
}

// KNWVotingABI is the input ABI used to generate the binding from.
const KNWVotingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"nextKNWVoting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakesOfVote\",\"outputs\":[{\"name\":\"proposersStake\",\"type\":\"uint256\"},{\"name\":\"proposersReward\",\"type\":\"uint256\"},{\"name\":\"returnPool\",\"type\":\"uint256\"},{\"name\":\"rewardPool\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"},{\"name\":\"winningSide\",\"type\":\"bool\"},{\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ditCoordinatorContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentVoteID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"repository\",\"type\":\"bytes32\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"name\":\"openEndDate\",\"type\":\"uint256\"},{\"name\":\"neededMajority\",\"type\":\"uint256\"},{\"name\":\"winningPercentage\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"name\":\"votesUnrevealed\",\"type\":\"uint256\"},{\"name\":\"participantsFor\",\"type\":\"uint256\"},{\"name\":\"participantsAgainst\",\"type\":\"uint256\"},{\"name\":\"participantsUnrevealed\",\"type\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeLabel\",\"type\":\"string\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"startVote\",\"outputs\":[{\"name\":\"voteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"didOpen\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"voteExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRepository\",\"type\":\"bytes32\"},{\"name\":\"_majority\",\"type\":\"uint256\"}],\"name\":\"addNewRepository\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startingVoteID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"endVote\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getUsedKNW\",\"outputs\":[{\"name\":\"usedKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"replaceDitManager\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDitCoordinatorAddress\",\"type\":\"address\"}],\"name\":\"addDitCoordinator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"voteEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getGrossStake\",\"outputs\":[{\"name\":\"grossStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURNING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"isResolved\",\"outputs\":[{\"name\":\"resolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getAmountOfVotes\",\"outputs\":[{\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"openVote\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getNetStake\",\"outputs\":[{\"name\":\"netStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"openPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastKNWVoting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_lastKNWVoting\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// KNWVotingBin is the compiled bytecode used for deploying new contracts.
const KNWVotingBin = `0x60806040523480156200001157600080fd5b506040516040806200337b8339810160405280516020909101516000600160a060020a0383161515620000a557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4b4e57546f6b656e20616464726573732063616e277420626520656d70747900604482015290519081900360640190fd5b60058054600160a060020a03808616600160a060020a03199283161792839055600680549092169281169290921790558216156200018e575060038054600160a060020a031916600160a060020a038381169190911791829055604080517f5a981129000000000000000000000000000000000000000000000000000000008152905192909116918291635a9811299160048083019260209291908290030181600087803b1580156200015757600080fd5b505af11580156200016c573d6000803e3d6000fd5b505050506040513d60208110156200018357600080fd5b505160075562000194565b60006007555b50506007546008555060028054600160a060020a031916331790556131bc80620001bf6000396000f30060806040526004361061017c5763ffffffff60e060020a6000350416630bf682e681146101815780631f84950f146101b257806336bf4c91146101f05780633fdf3dab14610235578063494031831461026a5780635a981129146102825780635df81330146102a95780636766b4a41461039957806370d18152146103d65780637e4173a8146103fa5780637f97e83614610412578063828c16531461043657806382ec45a314610451578063865df0ad146104665780638e1b660f1461047e57806391016157146104a2578063985dbfc5146104c3578063a326690a146104d8578063a4439dc5146104f9578063b43a401d14610511578063c3eb970814610529578063c814af1f14610541578063cc4e195414610556578063cd6ef46b1461056e578063cdd6ceb914610592578063d4e0ac95146105bc578063d9548e53146105e6578063dcfde092146105fe578063ea521a5214610616578063eb2c02231461062e578063effb21e114610541578063ff77f4b81461064f575b600080fd5b34801561018d57600080fd5b50610196610664565b60408051600160a060020a039092168252519081900360200190f35b3480156101be57600080fd5b506101ca600435610673565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156101fc57600080fd5b50610217600435602435600160a060020a036044351661069a565b60408051938452911515602084015282820152519081900360600190f35b34801561024157600080fd5b50610256600160a060020a0360043516610d28565b604080519115158252519081900360200190f35b34801561027657600080fd5b50610256600435610d3d565b34801561028e57600080fd5b50610297610f03565b60408051918252519081900360200190f35b3480156102b557600080fd5b506102c1600435610f09565b604051808e60001916600019168152602001806020018d81526020018c81526020018b81526020018a81526020018981526020018881526020018781526020018681526020018581526020018481526020018315151515815260200182810382528e818151815260200191508051906020019080838360005b8381101561035257818101518382015260200161033a565b50505050905090810190601f16801561037f5780820380516001836020036101000a031916815260200191505b509e50505050505050505050505050505060405180910390f35b3480156103a557600080fd5b50610297600480359060248035600160a060020a03169160443591820191013560643560843560a43560c435610fff565b3480156103e257600080fd5b50610256600160a060020a0360043516602435611479565b34801561040657600080fd5b506102566004356114fc565b34801561041e57600080fd5b50610256600160a060020a0360043516602435611515565b34801561044257600080fd5b50610256600435602435611593565b34801561045d57600080fd5b5061029761161f565b34801561047257600080fd5b50610256600435611625565b34801561048a57600080fd5b50610297600160a060020a0360043516602435611aca565b3480156104ae57600080fd5b50610256600160a060020a0360043516611af8565b3480156104cf57600080fd5b50610196611b59565b3480156104e457600080fd5b50610256600160a060020a0360043516611b68565b34801561050557600080fd5b50610256600435611ca2565b34801561051d57600080fd5b50610256600435611d13565b34801561053557600080fd5b50610297600435611d7d565b34801561054d57600080fd5b50610297611d8f565b34801561056257600080fd5b50610256600435611d94565b34801561057a57600080fd5b50610297600160a060020a0360043516602435611dac565b34801561059e57600080fd5b50610256600435600160a060020a0360243516604435606435611dda565b3480156105c857600080fd5b50610297600435600160a060020a03602435166044356064356122b2565b3480156105f257600080fd5b50610256600435612799565b34801561060a57600080fd5b5061029760043561279e565b34801561062257600080fd5b50610256600435612819565b34801561063a57600080fd5b50610256600160a060020a0360043516612894565b34801561065b57600080fd5b506101966128f5565b600454600160a060020a031681565b600a6020526000908152604090208054600182015460028301546003909301549192909184565b3360008181526020819052604081205490918291829182918291829160ff161515610711576040805160e560020a62461bcd02815260206004820152602b60248201526000805160206131518339815191526044820152600080516020613171833981519152606482015290519081900360840190fd5b60008a8152600960205260409020600c81015490945060ff161515610780576040805160e560020a62461bcd02815260206004820152601760248201527f506f6c6c2068617320746f206265207265736f6c766564000000000000000000604482015290519081900360640190fd5b600160a060020a0388166000908152600d8501602052604081206001015411156108e557600654600160a060020a038981166000818152600d880160205260409081902060019081015491517fd950df3400000000000000000000000000000000000000000000000000000000815260048101938452604481018390526060602482019081528a830180546002600019958216156101000295909501169390930460648301819052959096169563d950df34958f9593949392608401908590801561088c5780601f106108615761010080835404028352916020019161088c565b820191906000526020600020905b81548152906001019060200180831161086f57829003601f168201915b5050945050505050602060405180830381600087803b1580156108ae57600080fd5b505af11580156108c2573d6000803e3d6000fd5b505050506040513d60208110156108d857600080fd5b505115156108e557600080fd5b6108ee8a610d3d565b9250826108fc5760006108ff565b60015b600160a060020a0389166000908152600d8601602052604090205460ff9182168b1493506201000090041615610ae65782156109e7576109e088856001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109d15780601f106109a6576101008083540402835291602001916109d1565b820191906000526020600020905b8154815290600101906020018083116109b457829003601f168201915b50505050508660050154612904565b9450610acd565b60008a8152600a60205260409020600101541515610acd57610aca88856001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a9b5780601f10610a7057610100808354040283529160200191610a9b565b820191906000526020600020905b815481529060010190602001808311610a7e57829003601f168201915b50505050600160a060020a038c166000908152600d890160205260409020600101546005890154909150612a59565b94505b60008a8152600a60205260409020600101549650610d17565b610af0888b611479565b15610c275782158015610b0a575083600701548460060154145b15610b2257610b1b8a836001612c72565b9650610c22565b610b2e8a836000612c72565b96508115610bae57610ba788856001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109d15780601f106109a6576101008083540402835291602001916109d1565b9450610c22565b610c1f88856001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a9b5780601f10610a7057610100808354040283529160200191610a9b565b94505b610d17565b610c31888b611479565b158015610c435750610c43888b611515565b15610cc757610c548a600080612c72565b9650610c1f88856001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a9b5780601f10610a7057610100808354040283529160200191610a9b565b6040805160e560020a62461bcd02815260206004820152601d60248201527f4e6f742061207061727469636970616e74206f662074686520766f7465000000604482015290519081900360640190fd5b509498949750919550929350505050565b60006020819052908152604090205460ff1681565b6000610d47613027565b610d5083611d13565b1515610da6576040805160e560020a62461bcd02815260206004820152601660248201527f506f6c6c2068617320746f206861766520656e64656400000000000000000000604482015290519081900360640190fd5b60008381526009602090815260409182902082516101a08101845281548152600180830180548651600260001994831615610100029490940190911692909204601f8101869004860283018601909652858252919492938581019391929190830182828015610e565780601f10610e2b57610100808354040283529160200191610e56565b820191906000526020600020905b815481529060010190602001808311610e3957829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c820160009054906101000a900460ff16151515158152505090508060e001518160c00151018160800151028160c001516064021191505b50919050565b60075481565b6009602090815260009182526040918290208054600180830180548651600293821615610100026000190190911692909204601f810186900486028301860190965285825291949293909290830182828015610fa65780601f10610f7b57610100808354040283529160200191610fa6565b820191906000526020600020905b815481529060010190602001808311610f8957829003601f168201915b50505050509080600201549080600301549080600401549080600501549080600601549080600701549080600801549080600901549080600a01549080600b01549080600c0160009054906101000a900460ff1690508d565b336000818152602081905260408120549091829182919060ff161515611071576040805160e560020a62461bcd02815260206004820152602b60248201526000805160206131518339815191526044820152600080516020613171833981519152606482015290519081900360840190fd5b60075461108590600163ffffffff612e2e16565b600755611098428963ffffffff612e2e16565b92506110aa838863ffffffff612e2e16565b604080516101a0810182528e815281516020601f8e018190048102820181019093528c815292945091818301918d908d90819084018382808284378201915050505050508152602001848152602001838152602001600160008f60001916600019168152602001908152602001600020600001548152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581525060096000600754815260200190815260200160002060008201518160000190600019169055602082015181600101908051906020019061119c929190613095565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155610120820151816009015561014082015181600a015561016082015181600b015561018082015181600c0160006101000a81548160ff02191690831515021790555090505060806040519081016040528087815260200160008152602001600081526020016000815250600a60006007548152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301559050506001600960006007548152602001908152602001600020600d0160008d600160a060020a0316600160a060020a0316815260200190815260200160002060000160026101000a81548160ff021916908315150217905550600660009054906101000a9004600160a060020a0316600160a060020a031663735997988c600960006007548152602001908152602001600020600101886040518463ffffffff1660e060020a0281526004018084600160a060020a0316600160a060020a03168152602001806020018381526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156113dc5780601f106113b1576101008083540402835291602001916113dc565b820191906000526020600020905b8154815290600101906020018083116113bf57829003601f168201915b5050945050505050602060405180830381600087803b1580156113fe57600080fd5b505af1158015611412573d6000803e3d6000fd5b505050506040513d602081101561142857600080fd5b5051151561143557600080fd5b5050600780546000908152600960209081526040808320600160a060020a03909d168352600d909c0190529990992060010192909255505094549695505050505050565b6000611484826114fc565b15156114c8576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613131833981519152604482015290519081900360640190fd5b506000908152600960209081526040808320600160a060020a03949094168352600d90930190522054610100900460ff1690565b6000811580159061150f57506007548211155b92915050565b6000611520826114fc565b1515611564576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613131833981519152604482015290519081900360640190fd5b506000908152600960209081526040808320600160a060020a03949094168352600d9093019052205460ff1690565b3360008181526020819052604081205490919060ff161515611601576040805160e560020a62461bcd02815260206004820152602b60248201526000805160206131518339815191526044820152600080516020613171833981519152606482015290519081900360840190fd5b600084815260016020819052604090912084905591505b5092915050565b60085481565b33600081815260208190526040812054909182918291829182919060ff16151561169b576040805160e560020a62461bcd02815260206004820152602b60248201526000805160206131518339815191526044820152600080516020613171833981519152606482015290519081900360840190fd5b6116a487611d13565b15156116fa576040805160e560020a62461bcd02815260206004820152601660248201527f506f6c6c2068617320746f206861766520656e64656400000000000000000000604482015290519081900360640190fd5b600087815260096020526040902060078101546006909101546117229163ffffffff612e2e16565b6000888152600960208190526040909120600b81015491810154600a90910154929750611764926117589163ffffffff612e2e16565b9063ffffffff612e2e16565b93508315156117a9576000878152600a6020908152604080832080546001918201556009909252822060058101839055600c01805460ff191690911790559550611ac0565b6000878152600a60205260409020546117f8906117eb906117d0908763ffffffff612eb816565b60008a8152600a60205260409020549063ffffffff612f2616565b859063ffffffff612fae16565b6000888152600a6020526040812060020191909155925061181887610d3d565b955085156118a25760008781526009602052604090206006015461185590869061184990606463ffffffff612fae16565b9063ffffffff612eb816565b60008881526009602052604090206005810191909155600b810154600a909101546118859163ffffffff612e2e16565b6000888152600a6020526040902080546001909101559250611965565b6000878152600960205260409020600781015460069091015414611936576000878152600960205260409020600701546118e990869061184990606463ffffffff612fae16565b60008881526009602081905260409091206005810192909255600b8201549101546119199163ffffffff612e2e16565b6000888152600a6020526040902080546003909101559250611965565b600087815260096020908152604080832060326005820155600b0154600a909252909120805460019091015592505b6000878152600a602052604081206002015411156119f7576000878152600a60205260409020600201546119e4906119c6906117eb906119ab908863ffffffff612eb816565b60008b8152600a60205260409020549063ffffffff612f2616565b6000898152600a60205260409020600301549063ffffffff612e2e16565b6000888152600a60205260409020600301555b6000878152600960209081526040808320600c01805460ff19166001908117909155600a909252822001541115611ac057611a5e611a406001611758878763ffffffff612f2616565b6000898152600a60205260409020600301549063ffffffff612eb816565b6000888152600a6020526040902060010154909250611a83908363ffffffff612e2e16565b6000888152600a60205260409020600181019190915560030154611aad908363ffffffff612f2616565b6000888152600a60205260409020600301555b5050505050919050565b6000908152600960209081526040808320600160a060020a03949094168352600d9093019052206001015490565b600254600090600160a060020a03163314611b1257600080fd5b600160a060020a0382161515611b2757600080fd5b5060028054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600554600160a060020a031681565b600254600090600160a060020a03163314611bf3576040805160e560020a62461bcd02815260206004820152602160248201527f4f6e6c7920746865206469744d616e616765722063616e2063616c6c2074686960448201527f7300000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0382161515611c79576040805160e560020a62461bcd02815260206004820152603a60248201527f646974436f6f7264696e61746f7220616464726573732063616e206f6e6c792060448201527f62652061646465642069662069742773206e6f7420656d707479000000000000606482015290519081900360840190fd5b50600160a060020a03166000908152602081905260409020805460ff1916600190811790915590565b6000611cad826114fc565b1515611cf1576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613131833981519152604482015290519081900360640190fd5b600082815260096020526040902060020154611d0c90612799565b1592915050565b6000611d1e826114fc565b1515611d62576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613131833981519152604482015290519081900360640190fd5b60008281526009602052604090206003015461150f90612799565b6000908152600a602052604090205490565b600081565b6000908152600960205260409020600c015460ff1690565b6000908152600960209081526040808320600160a060020a03949094168352600d9093019052206002015490565b336000818152602081905260408120549091829160ff161515611e49576040805160e560020a62461bcd02815260206004820152602b60248201526000805160206131518339815191526044820152600080516020613171833981519152606482015290519081900360840190fd5b611e5287612819565b1515611ea8576040805160e560020a62461bcd02815260206004820152601e60248201527f52657665616c20706572696f642068617320746f206265206163746976650000604482015290519081900360640190fd5b6000878152600960209081526040808320600160a060020a038a168452600d0190915290205460ff161515611f4d576040805160e560020a62461bcd02815260206004820152602760248201527f5061727469636970616e742068617320746f2068617665206120766f7465206360448201527f6f6d6d6974656400000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000878152600960209081526040808320600160a060020a038a168452600d01909152902054610100900460ff1615611ff6576040805160e560020a62461bcd02815260206004820152602260248201527f43616e27742072657665616c206120766f7465206d6f7265207468616e206f6e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000878152600960209081526040808320600160a060020a038a168452600d01825291829020600301548251808301899052808401889052835180820385018152606090910193849052805191939092909182918401908083835b602083106120705780518252601f199092019160209182019101612051565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614151561211e576040805160e560020a62461bcd02815260206004820152603660248201527f43686f69636520616e642053616c74206861766520746f20626520746865207360448201527f616d6520617320696e2074686520766f74656861736800000000000000000000606482015290519081900360840190fd5b6000878152600960208181526040808420600160a060020a038b168552600d8101835290842060020154938b905291905260080154909250612166908363ffffffff612f2616565b60008881526009602052604090206008810191909155600b015461219190600163ffffffff612f2616565b6000888152600960205260409020600b01556001851415612212576000878152600960205260409020600601546121ce908363ffffffff612e2e16565b6000888152600960208190526040909120600681019290925501546121fa90600163ffffffff612e2e16565b60008881526009602081905260409091200155612272565b600087815260096020526040902060070154612234908363ffffffff612e2e16565b60008881526009602052604090206007810191909155600a015461225f90600163ffffffff612e2e16565b6000888152600960205260409020600a01555b6000878152600960209081526040808320600160a060020a038a168452600d019091529020805461ff001916610100179055600192505050949350505050565b336000818152602081905260408120549091829160ff161515612321576040805160e560020a62461bcd02815260206004820152602b60248201526000805160206131518339815191526044820152600080516020613171833981519152606482015290519081900360840190fd5b861515612378576040805160e560020a62461bcd02815260206004820152601460248201527f766f746549442063616e2774206265207a65726f000000000000000000000000604482015290519081900360640190fd5b61238187611ca2565b15156123d7576040805160e560020a62461bcd02815260206004820152601e60248201527f436f6d6d697420706572696f642068617320746f206265206163746976650000604482015290519081900360640190fd5b6123e18688611515565b15612436576040805160e560020a62461bcd02815260206004820152601f60248201527f43616e277420636f6d6d6974206d6f7265207468616e206f6e6520766f746500604482015290519081900360640190fd5b84151561248d576040805160e560020a62461bcd02815260206004820152601b60248201527f43616e277420766f746520776974682061207a65726f20686173680000000000604482015290519081900360640190fd5b6000878152600a602090815260408083205460065460099093529281902090517f73599798000000000000000000000000000000000000000000000000000000008152600160a060020a038a811660048301908152604483018a905260606024840190815260019485018054600260001997821615610100029790970116959095046064850181905296995091909416946373599798948c94938b939192608490910190859080156125805780601f1061255557610100808354040283529160200191612580565b820191906000526020600020905b81548152906001019060200180831161256357829003601f168201915b5050945050505050602060405180830381600087803b1580156125a257600080fd5b505af11580156125b6573d6000803e3d6000fd5b505050506040513d60208110156125cc57600080fd5b505115156125d957600080fd5b6000878152600960209081526040808320600160a060020a038a168452600d0190915290206001018490556126198464e8d4a5100063ffffffff612eb816565b73__libraries/SafeMath.sol:SafeMath_______63677342ce90916040518263ffffffff1660e060020a0281526004018082815260200191505060206040518083038186803b15801561266c57600080fd5b505af4158015612680573d6000803e3d6000fd5b505050506040513d602081101561269657600080fd5b505191506126b18466038d7ea4c6800063ffffffff612eb816565b82106126d1576126ce8466038d7ea4c6800063ffffffff612eb816565b91505b6126f76126ea6103e8611849858763ffffffff612fae16565b849063ffffffff612e2e16565b6000888152600960208181526040808420600160a060020a038c168552600d8101835290842060028101869055600381018b9055805460ff19166001179055928b90525260080154909350612752908463ffffffff612e2e16565b60008881526009602052604090206008810191909155600b015461277d90600163ffffffff612e2e16565b6000888152600960205260409020600b01555050949350505050565b421190565b60008181526009602081905260408220600b81015491810154600a9091015483926127d49290916117589163ffffffff612e2e16565b90506000811115612805576000838152600a60205260409020546127fe908263ffffffff612eb816565b9150610efd565b50506000908152600a602052604090205490565b6000612824826114fc565b1515612868576040805160e560020a62461bcd0281526020600482015260116024820152600080516020613131833981519152604482015290519081900360640190fd5b60008281526009602052604090206003015461288390612799565b15801561150f5750611d0c82611ca2565b600254600090600160a060020a031633146128ae57600080fd5b600160a060020a03821615156128c357600080fd5b5060048054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600354600160a060020a031681565b60008061292f66470de4df82000061292385603263ffffffff612f2616565b9063ffffffff612fae16565b90506000811115612a51576006546040517fba7aef43000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830190815260448301859052606060248401908152885160648501528851929094169363ba7aef43938a938a93889390929091608490910190602086019080838360005b838110156129cc5781810151838201526020016129b4565b50505050905090810190601f1680156129f95780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015612a1a57600080fd5b505af1158015612a2e573d6000803e3d6000fd5b505050506040513d6020811015612a4457600080fd5b50511515612a5157600080fd5b949350505050565b600080808080861115612b4657612b1166038d7ea4c68000612a868864e8d4a5100063ffffffff612eb816565b73__libraries/SafeMath.sol:SafeMath_______63677342ce90916040518263ffffffff1660e060020a0281526004018082815260200191505060206040518083038186803b158015612ad957600080fd5b505af4158015612aed573d6000803e3d6000fd5b505050506040513d6020811015612b0357600080fd5b50519063ffffffff612fae16565b915085821015612b3257612b2b868363ffffffff612f2616565b9250612b46565b612b4386600263ffffffff612eb816565b92505b6000831115612c66576006546040517fc45b71de000000000000000000000000000000000000000000000000000000008152600160a060020a038a811660048301908152604483018790526060602484019081528b5160648501528b51929094169363c45b71de938d938d938a9390929091608490910190602086019080838360005b83811015612be1578181015183820152602001612bc9565b50505050905090810190601f168015612c0e5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015612c2f57600080fd5b505af1158015612c43573d6000803e3d6000fd5b505050506040513d6020811015612c5957600080fd5b50511515612c6657600080fd5b50909695505050505050565b60008381526009602081905260408220600b81015491810154600a909101548392839283928392612cae9290916117589163ffffffff612e2e16565b935060009250851515612d6d576000888152600a6020526040902060020154612cdd908563ffffffff612eb816565b91508192508615612d6857506000612cf488610d3d565b15612d12575060008781526009602081905260409091200154612d26565b506000878152600960205260409020600a01545b6000888152600a6020526040902060030154612d6590612d4c908363ffffffff612eb816565b611758612d588b61279e565b869063ffffffff612e2e16565b92505b612c66565b6000888152600a602090815260408083205460099092528220600b0154909450118015612daa57506000888152600960205260409020600b015484115b15612c66576000888152600960205260409020600b0154612e2290612e0790612dda90879063ffffffff612f2616565b60008b8152600a60209081526040808320546009909252909120600b01546118499163ffffffff612fae16565b60008a8152600a60205260409020549063ffffffff612e2e16565b98975050505050505050565b600082820183811015612eb1576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b600080808311612f12576040805160e560020a62461bcd02815260206004820152601460248201527f43616e277420646976696465206279207a65726f000000000000000000000000604482015290519081900360640190fd5b8284811515612f1d57fe5b04949350505050565b60008083831115612fa7576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b600080831515612fc15760009150611618565b50828202828482811515612fd157fe5b0414612eb1576040805160e560020a62461bcd02815260206004820152601f60248201527f466c6177656420696e70757420666f72206d756c7469706c69636174696f6e00604482015290519081900360640190fd5b6101a0604051908101604052806000801916815260200160608152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106130d657805160ff1916838001178555613103565b82800160010185558215613103579182015b828111156131035782518255916020019190600101906130e8565b5061310f929150613113565b5090565b61312d91905b8082111561310f5760008155600101613119565b905600506f6c6c2068617320746f2065786973740000000000000000000000000000004f6e6c79206120646974436f6f7264696e61746f7220697320616c6c6f7720746f2063616c6c2074686973000000000000000000000000000000000000000000a165627a7a723058207923a47ac34515274810a2b1d9cf681562e35bf72d7558100efe4b78e39573d80029`

// DeployKNWVoting deploys a new Ethereum contract, binding an instance of KNWVoting to it.
func DeployKNWVoting(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWTokenAddress common.Address, _lastKNWVoting common.Address) (common.Address, *types.Transaction, *KNWVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWVotingBin), backend, _KNWTokenAddress, _lastKNWVoting)
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

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) BURNINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "BURNING_METHOD")
	return *ret0, err
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) BURNINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.BURNINGMETHOD(&_KNWVoting.CallOpts)
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) BURNINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.BURNINGMETHOD(&_KNWVoting.CallOpts)
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

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) MINTINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "MINTING_METHOD")
	return *ret0, err
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) MINTINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.MINTINGMETHOD(&_KNWVoting.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) MINTINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.MINTINGMETHOD(&_KNWVoting.CallOpts)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCaller) CommitPeriodActive(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "commitPeriodActive", _voteID)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingSession) CommitPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.CommitPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCallerSession) CommitPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.CommitPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// CurrentVoteID is a free data retrieval call binding the contract method 0x5a981129.
//
// Solidity: function currentVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) CurrentVoteID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "currentVoteID")
	return *ret0, err
}

// CurrentVoteID is a free data retrieval call binding the contract method 0x5a981129.
//
// Solidity: function currentVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) CurrentVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.CurrentVoteID(&_KNWVoting.CallOpts)
}

// CurrentVoteID is a free data retrieval call binding the contract method 0x5a981129.
//
// Solidity: function currentVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) CurrentVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.CurrentVoteID(&_KNWVoting.CallOpts)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(address _address, uint256 _voteID) constant returns(bool committed)
func (_KNWVoting *KNWVotingCaller) DidCommit(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "didCommit", _address, _voteID)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(address _address, uint256 _voteID) constant returns(bool committed)
func (_KNWVoting *KNWVotingSession) DidCommit(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidCommit(&_KNWVoting.CallOpts, _address, _voteID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(address _address, uint256 _voteID) constant returns(bool committed)
func (_KNWVoting *KNWVotingCallerSession) DidCommit(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidCommit(&_KNWVoting.CallOpts, _address, _voteID)
}

// DidOpen is a free data retrieval call binding the contract method 0x70d18152.
//
// Solidity: function didOpen(address _address, uint256 _voteID) constant returns(bool revealed)
func (_KNWVoting *KNWVotingCaller) DidOpen(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "didOpen", _address, _voteID)
	return *ret0, err
}

// DidOpen is a free data retrieval call binding the contract method 0x70d18152.
//
// Solidity: function didOpen(address _address, uint256 _voteID) constant returns(bool revealed)
func (_KNWVoting *KNWVotingSession) DidOpen(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidOpen(&_KNWVoting.CallOpts, _address, _voteID)
}

// DidOpen is a free data retrieval call binding the contract method 0x70d18152.
//
// Solidity: function didOpen(address _address, uint256 _voteID) constant returns(bool revealed)
func (_KNWVoting *KNWVotingCallerSession) DidOpen(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidOpen(&_KNWVoting.CallOpts, _address, _voteID)
}

// DitCoordinatorContracts is a free data retrieval call binding the contract method 0x3fdf3dab.
//
// Solidity: function ditCoordinatorContracts(address ) constant returns(bool)
func (_KNWVoting *KNWVotingCaller) DitCoordinatorContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "ditCoordinatorContracts", arg0)
	return *ret0, err
}

// DitCoordinatorContracts is a free data retrieval call binding the contract method 0x3fdf3dab.
//
// Solidity: function ditCoordinatorContracts(address ) constant returns(bool)
func (_KNWVoting *KNWVotingSession) DitCoordinatorContracts(arg0 common.Address) (bool, error) {
	return _KNWVoting.Contract.DitCoordinatorContracts(&_KNWVoting.CallOpts, arg0)
}

// DitCoordinatorContracts is a free data retrieval call binding the contract method 0x3fdf3dab.
//
// Solidity: function ditCoordinatorContracts(address ) constant returns(bool)
func (_KNWVoting *KNWVotingCallerSession) DitCoordinatorContracts(arg0 common.Address) (bool, error) {
	return _KNWVoting.Contract.DitCoordinatorContracts(&_KNWVoting.CallOpts, arg0)
}

// GetAmountOfVotes is a free data retrieval call binding the contract method 0xcd6ef46b.
//
// Solidity: function getAmountOfVotes(address _address, uint256 _voteID) constant returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingCaller) GetAmountOfVotes(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getAmountOfVotes", _address, _voteID)
	return *ret0, err
}

// GetAmountOfVotes is a free data retrieval call binding the contract method 0xcd6ef46b.
//
// Solidity: function getAmountOfVotes(address _address, uint256 _voteID) constant returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingSession) GetAmountOfVotes(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetAmountOfVotes(&_KNWVoting.CallOpts, _address, _voteID)
}

// GetAmountOfVotes is a free data retrieval call binding the contract method 0xcd6ef46b.
//
// Solidity: function getAmountOfVotes(address _address, uint256 _voteID) constant returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingCallerSession) GetAmountOfVotes(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetAmountOfVotes(&_KNWVoting.CallOpts, _address, _voteID)
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(uint256 _voteID) constant returns(uint256 grossStake)
func (_KNWVoting *KNWVotingCaller) GetGrossStake(opts *bind.CallOpts, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getGrossStake", _voteID)
	return *ret0, err
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(uint256 _voteID) constant returns(uint256 grossStake)
func (_KNWVoting *KNWVotingSession) GetGrossStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetGrossStake(&_KNWVoting.CallOpts, _voteID)
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(uint256 _voteID) constant returns(uint256 grossStake)
func (_KNWVoting *KNWVotingCallerSession) GetGrossStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetGrossStake(&_KNWVoting.CallOpts, _voteID)
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(uint256 _voteID) constant returns(uint256 netStake)
func (_KNWVoting *KNWVotingCaller) GetNetStake(opts *bind.CallOpts, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getNetStake", _voteID)
	return *ret0, err
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(uint256 _voteID) constant returns(uint256 netStake)
func (_KNWVoting *KNWVotingSession) GetNetStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNetStake(&_KNWVoting.CallOpts, _voteID)
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(uint256 _voteID) constant returns(uint256 netStake)
func (_KNWVoting *KNWVotingCallerSession) GetNetStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNetStake(&_KNWVoting.CallOpts, _voteID)
}

// GetUsedKNW is a free data retrieval call binding the contract method 0x8e1b660f.
//
// Solidity: function getUsedKNW(address _address, uint256 _voteID) constant returns(uint256 usedKNW)
func (_KNWVoting *KNWVotingCaller) GetUsedKNW(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getUsedKNW", _address, _voteID)
	return *ret0, err
}

// GetUsedKNW is a free data retrieval call binding the contract method 0x8e1b660f.
//
// Solidity: function getUsedKNW(address _address, uint256 _voteID) constant returns(uint256 usedKNW)
func (_KNWVoting *KNWVotingSession) GetUsedKNW(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetUsedKNW(&_KNWVoting.CallOpts, _address, _voteID)
}

// GetUsedKNW is a free data retrieval call binding the contract method 0x8e1b660f.
//
// Solidity: function getUsedKNW(address _address, uint256 _voteID) constant returns(uint256 usedKNW)
func (_KNWVoting *KNWVotingCallerSession) GetUsedKNW(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetUsedKNW(&_KNWVoting.CallOpts, _address, _voteID)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 _terminationDate) constant returns(bool expired)
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
// Solidity: function isExpired(uint256 _terminationDate) constant returns(bool expired)
func (_KNWVoting *KNWVotingSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsExpired(&_KNWVoting.CallOpts, _terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 _terminationDate) constant returns(bool expired)
func (_KNWVoting *KNWVotingCallerSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsExpired(&_KNWVoting.CallOpts, _terminationDate)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(uint256 _voteID) constant returns(bool passed)
func (_KNWVoting *KNWVotingCaller) IsPassed(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isPassed", _voteID)
	return *ret0, err
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(uint256 _voteID) constant returns(bool passed)
func (_KNWVoting *KNWVotingSession) IsPassed(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsPassed(&_KNWVoting.CallOpts, _voteID)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(uint256 _voteID) constant returns(bool passed)
func (_KNWVoting *KNWVotingCallerSession) IsPassed(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsPassed(&_KNWVoting.CallOpts, _voteID)
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(uint256 _voteID) constant returns(bool resolved)
func (_KNWVoting *KNWVotingCaller) IsResolved(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isResolved", _voteID)
	return *ret0, err
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(uint256 _voteID) constant returns(bool resolved)
func (_KNWVoting *KNWVotingSession) IsResolved(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsResolved(&_KNWVoting.CallOpts, _voteID)
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(uint256 _voteID) constant returns(bool resolved)
func (_KNWVoting *KNWVotingCallerSession) IsResolved(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsResolved(&_KNWVoting.CallOpts, _voteID)
}

// LastKNWVoting is a free data retrieval call binding the contract method 0xff77f4b8.
//
// Solidity: function lastKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCaller) LastKNWVoting(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "lastKNWVoting")
	return *ret0, err
}

// LastKNWVoting is a free data retrieval call binding the contract method 0xff77f4b8.
//
// Solidity: function lastKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingSession) LastKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.LastKNWVoting(&_KNWVoting.CallOpts)
}

// LastKNWVoting is a free data retrieval call binding the contract method 0xff77f4b8.
//
// Solidity: function lastKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) LastKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.LastKNWVoting(&_KNWVoting.CallOpts)
}

// NextKNWVoting is a free data retrieval call binding the contract method 0x0bf682e6.
//
// Solidity: function nextKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCaller) NextKNWVoting(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "nextKNWVoting")
	return *ret0, err
}

// NextKNWVoting is a free data retrieval call binding the contract method 0x0bf682e6.
//
// Solidity: function nextKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingSession) NextKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.NextKNWVoting(&_KNWVoting.CallOpts)
}

// NextKNWVoting is a free data retrieval call binding the contract method 0x0bf682e6.
//
// Solidity: function nextKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) NextKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.NextKNWVoting(&_KNWVoting.CallOpts)
}

// OpenPeriodActive is a free data retrieval call binding the contract method 0xea521a52.
//
// Solidity: function openPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCaller) OpenPeriodActive(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "openPeriodActive", _voteID)
	return *ret0, err
}

// OpenPeriodActive is a free data retrieval call binding the contract method 0xea521a52.
//
// Solidity: function openPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingSession) OpenPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.OpenPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// OpenPeriodActive is a free data retrieval call binding the contract method 0xea521a52.
//
// Solidity: function openPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCallerSession) OpenPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.OpenPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// StakesOfVote is a free data retrieval call binding the contract method 0x1f84950f.
//
// Solidity: function stakesOfVote(uint256 ) constant returns(uint256 proposersStake, uint256 proposersReward, uint256 returnPool, uint256 rewardPool)
func (_KNWVoting *KNWVotingCaller) StakesOfVote(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _KNWVoting.contract.Call(opts, out, "stakesOfVote", arg0)
	return *ret, err
}

// StakesOfVote is a free data retrieval call binding the contract method 0x1f84950f.
//
// Solidity: function stakesOfVote(uint256 ) constant returns(uint256 proposersStake, uint256 proposersReward, uint256 returnPool, uint256 rewardPool)
func (_KNWVoting *KNWVotingSession) StakesOfVote(arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	return _KNWVoting.Contract.StakesOfVote(&_KNWVoting.CallOpts, arg0)
}

// StakesOfVote is a free data retrieval call binding the contract method 0x1f84950f.
//
// Solidity: function stakesOfVote(uint256 ) constant returns(uint256 proposersStake, uint256 proposersReward, uint256 returnPool, uint256 rewardPool)
func (_KNWVoting *KNWVotingCallerSession) StakesOfVote(arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	return _KNWVoting.Contract.StakesOfVote(&_KNWVoting.CallOpts, arg0)
}

// StartingVoteID is a free data retrieval call binding the contract method 0x82ec45a3.
//
// Solidity: function startingVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) StartingVoteID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "startingVoteID")
	return *ret0, err
}

// StartingVoteID is a free data retrieval call binding the contract method 0x82ec45a3.
//
// Solidity: function startingVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) StartingVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.StartingVoteID(&_KNWVoting.CallOpts)
}

// StartingVoteID is a free data retrieval call binding the contract method 0x82ec45a3.
//
// Solidity: function startingVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) StartingVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.StartingVoteID(&_KNWVoting.CallOpts)
}

// VoteEnded is a free data retrieval call binding the contract method 0xb43a401d.
//
// Solidity: function voteEnded(uint256 _voteID) constant returns(bool ended)
func (_KNWVoting *KNWVotingCaller) VoteEnded(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "voteEnded", _voteID)
	return *ret0, err
}

// VoteEnded is a free data retrieval call binding the contract method 0xb43a401d.
//
// Solidity: function voteEnded(uint256 _voteID) constant returns(bool ended)
func (_KNWVoting *KNWVotingSession) VoteEnded(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteEnded(&_KNWVoting.CallOpts, _voteID)
}

// VoteEnded is a free data retrieval call binding the contract method 0xb43a401d.
//
// Solidity: function voteEnded(uint256 _voteID) constant returns(bool ended)
func (_KNWVoting *KNWVotingCallerSession) VoteEnded(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteEnded(&_KNWVoting.CallOpts, _voteID)
}

// VoteExists is a free data retrieval call binding the contract method 0x7e4173a8.
//
// Solidity: function voteExists(uint256 _voteID) constant returns(bool exists)
func (_KNWVoting *KNWVotingCaller) VoteExists(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "voteExists", _voteID)
	return *ret0, err
}

// VoteExists is a free data retrieval call binding the contract method 0x7e4173a8.
//
// Solidity: function voteExists(uint256 _voteID) constant returns(bool exists)
func (_KNWVoting *KNWVotingSession) VoteExists(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteExists(&_KNWVoting.CallOpts, _voteID)
}

// VoteExists is a free data retrieval call binding the contract method 0x7e4173a8.
//
// Solidity: function voteExists(uint256 _voteID) constant returns(bool exists)
func (_KNWVoting *KNWVotingCallerSession) VoteExists(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteExists(&_KNWVoting.CallOpts, _voteID)
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) constant returns(bytes32 repository, string knowledgeLabel, uint256 commitEndDate, uint256 openEndDate, uint256 neededMajority, uint256 winningPercentage, uint256 votesFor, uint256 votesAgainst, uint256 votesUnrevealed, uint256 participantsFor, uint256 participantsAgainst, uint256 participantsUnrevealed, bool isResolved)
func (_KNWVoting *KNWVotingCaller) Votes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Repository             [32]byte
	KnowledgeLabel         string
	CommitEndDate          *big.Int
	OpenEndDate            *big.Int
	NeededMajority         *big.Int
	WinningPercentage      *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	ret := new(struct {
		Repository             [32]byte
		KnowledgeLabel         string
		CommitEndDate          *big.Int
		OpenEndDate            *big.Int
		NeededMajority         *big.Int
		WinningPercentage      *big.Int
		VotesFor               *big.Int
		VotesAgainst           *big.Int
		VotesUnrevealed        *big.Int
		ParticipantsFor        *big.Int
		ParticipantsAgainst    *big.Int
		ParticipantsUnrevealed *big.Int
		IsResolved             bool
	})
	out := ret
	err := _KNWVoting.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) constant returns(bytes32 repository, string knowledgeLabel, uint256 commitEndDate, uint256 openEndDate, uint256 neededMajority, uint256 winningPercentage, uint256 votesFor, uint256 votesAgainst, uint256 votesUnrevealed, uint256 participantsFor, uint256 participantsAgainst, uint256 participantsUnrevealed, bool isResolved)
func (_KNWVoting *KNWVotingSession) Votes(arg0 *big.Int) (struct {
	Repository             [32]byte
	KnowledgeLabel         string
	CommitEndDate          *big.Int
	OpenEndDate            *big.Int
	NeededMajority         *big.Int
	WinningPercentage      *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	return _KNWVoting.Contract.Votes(&_KNWVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) constant returns(bytes32 repository, string knowledgeLabel, uint256 commitEndDate, uint256 openEndDate, uint256 neededMajority, uint256 winningPercentage, uint256 votesFor, uint256 votesAgainst, uint256 votesUnrevealed, uint256 participantsFor, uint256 participantsAgainst, uint256 participantsUnrevealed, bool isResolved)
func (_KNWVoting *KNWVotingCallerSession) Votes(arg0 *big.Int) (struct {
	Repository             [32]byte
	KnowledgeLabel         string
	CommitEndDate          *big.Int
	OpenEndDate            *big.Int
	NeededMajority         *big.Int
	WinningPercentage      *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	return _KNWVoting.Contract.Votes(&_KNWVoting.CallOpts, arg0)
}

// AddDitCoordinator is a paid mutator transaction binding the contract method 0xa326690a.
//
// Solidity: function addDitCoordinator(address _newDitCoordinatorAddress) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) AddDitCoordinator(opts *bind.TransactOpts, _newDitCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "addDitCoordinator", _newDitCoordinatorAddress)
}

// AddDitCoordinator is a paid mutator transaction binding the contract method 0xa326690a.
//
// Solidity: function addDitCoordinator(address _newDitCoordinatorAddress) returns(bool success)
func (_KNWVoting *KNWVotingSession) AddDitCoordinator(_newDitCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddDitCoordinator(&_KNWVoting.TransactOpts, _newDitCoordinatorAddress)
}

// AddDitCoordinator is a paid mutator transaction binding the contract method 0xa326690a.
//
// Solidity: function addDitCoordinator(address _newDitCoordinatorAddress) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) AddDitCoordinator(_newDitCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddDitCoordinator(&_KNWVoting.TransactOpts, _newDitCoordinatorAddress)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) AddNewRepository(opts *bind.TransactOpts, _newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "addNewRepository", _newRepository, _majority)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVoting *KNWVotingSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddNewRepository(&_KNWVoting.TransactOpts, _newRepository, _majority)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddNewRepository(&_KNWVoting.TransactOpts, _newRepository, _majority)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingTransactor) CommitVote(opts *bind.TransactOpts, _voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "commitVote", _voteID, _address, _secretHash, _numberOfKNW)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingSession) CommitVote(_voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.CommitVote(&_KNWVoting.TransactOpts, _voteID, _address, _secretHash, _numberOfKNW)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingTransactorSession) CommitVote(_voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.CommitVote(&_KNWVoting.TransactOpts, _voteID, _address, _secretHash, _numberOfKNW)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVoting *KNWVotingTransactor) EndVote(opts *bind.TransactOpts, _voteID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "endVote", _voteID)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVoting *KNWVotingSession) EndVote(_voteID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.EndVote(&_KNWVoting.TransactOpts, _voteID)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVoting *KNWVotingTransactorSession) EndVote(_voteID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.EndVote(&_KNWVoting.TransactOpts, _voteID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 numberOfKNW)
func (_KNWVoting *KNWVotingTransactor) FinalizeVote(opts *bind.TransactOpts, _voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "finalizeVote", _voteID, _voteOption, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 numberOfKNW)
func (_KNWVoting *KNWVotingSession) FinalizeVote(_voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.FinalizeVote(&_KNWVoting.TransactOpts, _voteID, _voteOption, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 numberOfKNW)
func (_KNWVoting *KNWVotingTransactorSession) FinalizeVote(_voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.FinalizeVote(&_KNWVoting.TransactOpts, _voteID, _voteOption, _address)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) OpenVote(opts *bind.TransactOpts, _voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "openVote", _voteID, _address, _voteOption, _salt)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVoting *KNWVotingSession) OpenVote(_voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.OpenVote(&_KNWVoting.TransactOpts, _voteID, _address, _voteOption, _salt)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) OpenVote(_voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.OpenVote(&_KNWVoting.TransactOpts, _voteID, _address, _voteOption, _salt)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) ReplaceDitManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "replaceDitManager", _newManager)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool success)
func (_KNWVoting *KNWVotingSession) ReplaceDitManager(_newManager common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.ReplaceDitManager(&_KNWVoting.TransactOpts, _newManager)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) ReplaceDitManager(_newManager common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.ReplaceDitManager(&_KNWVoting.TransactOpts, _newManager)
}

// StartVote is a paid mutator transaction binding the contract method 0x6766b4a4.
//
// Solidity: function startVote(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVoting *KNWVotingTransactor) StartVote(opts *bind.TransactOpts, _repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "startVote", _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0x6766b4a4.
//
// Solidity: function startVote(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVoting *KNWVotingSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.StartVote(&_KNWVoting.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0x6766b4a4.
//
// Solidity: function startVote(bytes32 _repository, address _address, string _knowledgeLabel, uint256 _commitDuration, uint256 _revealDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVoting *KNWVotingTransactorSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeLabel string, _commitDuration *big.Int, _revealDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.StartVote(&_KNWVoting.TransactOpts, _repository, _address, _knowledgeLabel, _commitDuration, _revealDuration, _proposersStake, _numberOfKNW)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) UpgradeContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "upgradeContract", _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool success)
func (_KNWVoting *KNWVotingSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.UpgradeContract(&_KNWVoting.TransactOpts, _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.UpgradeContract(&_KNWVoting.TransactOpts, _address)
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
