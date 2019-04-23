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

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
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

// KNWTokenContractABI is the input ABI used to generate the binding from.
const KNWTokenContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newContractAddress\",\"type\":\"address\"}],\"name\":\"addVotingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// AddVotingContract is a paid mutator transaction binding the contract method 0x20c7ad57.
//
// Solidity: function addVotingContract(address _newContractAddress) returns()
func (_KNWTokenContract *KNWTokenContractTransactor) AddVotingContract(opts *bind.TransactOpts, _newContractAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "addVotingContract", _newContractAddress)
}

// AddVotingContract is a paid mutator transaction binding the contract method 0x20c7ad57.
//
// Solidity: function addVotingContract(address _newContractAddress) returns()
func (_KNWTokenContract *KNWTokenContractSession) AddVotingContract(_newContractAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.AddVotingContract(&_KNWTokenContract.TransactOpts, _newContractAddress)
}

// AddVotingContract is a paid mutator transaction binding the contract method 0x20c7ad57.
//
// Solidity: function addVotingContract(address _newContractAddress) returns()
func (_KNWTokenContract *KNWTokenContractTransactorSession) AddVotingContract(_newContractAddress common.Address) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.AddVotingContract(&_KNWTokenContract.TransactOpts, _newContractAddress)
}

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

// DitDemoCoordinatorABI is the input ABI used to generate the binding from.
const DitDemoCoordinatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getKNWVoteIDFromProposalID\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getCurrentProposalID\",\"outputs\":[{\"name\":\"currentProposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_voteSalt\",\"type\":\"uint256\"}],\"name\":\"openVoteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repositories\",\"outputs\":[{\"name\":\"currentProposalID\",\"type\":\"uint256\"},{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalsOfRepository\",\"outputs\":[{\"name\":\"KNWVoteID\",\"type\":\"uint256\"},{\"name\":\"knowledgeLabel\",\"type\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"proposalAccepted\",\"type\":\"bool\"},{\"name\":\"individualStake\",\"type\":\"uint256\"},{\"name\":\"totalStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_VOTE_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"getVotingMajority\",\"outputs\":[{\"name\":\"votingMajority\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OPEN_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"getIndividualStake\",\"outputs\":[{\"name\":\"individualStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_OPEN_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"proposalHasPassed\",\"outputs\":[{\"name\":\"hasPassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"replaceDitManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelID\",\"type\":\"uint256\"}],\"name\":\"getKnowledgeLabels\",\"outputs\":[{\"name\":\"knowledgeLabel\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_label1\",\"type\":\"string\"},{\"name\":\"_label2\",\"type\":\"string\"},{\"name\":\"_label3\",\"type\":\"string\"},{\"name\":\"_votingMajority\",\"type\":\"uint256\"}],\"name\":\"initRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"voteOnProposal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURNING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"passedKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_knowledgeLabelIndex\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"},{\"name\":\"_voteCommitDuration\",\"type\":\"uint256\"},{\"name\":\"_voteOpenDuration\",\"type\":\"uint256\"},{\"name\":\"_amountOfTokens\",\"type\":\"uint256\"}],\"name\":\"proposeCommit\",\"outputs\":[{\"name\":\"proposalID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xDitTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addKYCValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDitCoordinator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"repositoryIsInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"}],\"name\":\"migrateRepository\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWVotingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"passKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_KNWVotingAddress\",\"type\":\"address\"},{\"name\":\"_lastDitCoordinator\",\"type\":\"address\"},{\"name\":\"_xDitTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"ProposeCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"CommitVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accept\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"name\":\"OpenVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"votedRight\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"FinalizeVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"repository\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"proposal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"FinalizeProposal\",\"type\":\"event\"}]"

// DitDemoCoordinatorBin is the compiled bytecode used for deploying new contracts.
const DitDemoCoordinatorBin = `0x608060405234801561001057600080fd5b50604051608080612c3f8339810160409081528151602083015191830151606090930151909290600160a060020a038316158015906100575750600160a060020a03841615155b15156100ea57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4b4e57566f74696e6720616e64204b4e57546f6b656e2061646472657373206360448201527f616e277420626520656d70747900000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a03948516600160a060020a03199182161780835560068054831691871691909117905560018054968616968216969096178087556007805483169187169190911790556002805493861693821693909317928390556008805482169386169390931790925560038054939094169282169290921790925533808252600c6020526040909120805460ff1916909317909255600580549091169091179055612a9e806101a16000396000f3006080604052600436106101ab5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306ee459681146101b05780630bdc90e8146101dd5780630ee62ec0146101f55780631341f25c1461022a5780631ecbb9de1461024b5780631f51fd71146102605780632e71d0fb1461029157806339ba645b146102ac5780633e029f63146102cd5780633eedfc101461039b5780633fcc148d146103b0578063466577cd1461024b578063552edc9d146103c85780636fcfeb3b1461039b57806373b0dddd146103e357806387c9203d14610404578063910161571461041f5780639533222914610440578063985dbfc5146104d057806399821d9b146105015780639a401f4814610516578063ab4b593e14610555578063c814af1f14610576578063ccd9aa681461058b578063cda8d421146105ac578063ce83732e146105d3578063d0c397ef146105e8578063d7ad0eae14610609578063ea6c6d0f1461061e578063ea976d2c14610636578063eb2c02231461064e578063eb49fe941461066f578063eb93102414610684578063effb21e114610576575b600080fd5b3480156101bc57600080fd5b506101cb6004356024356106a5565b60408051918252519081900360200190f35b3480156101e957600080fd5b506101cb6004356106c2565b34801561020157600080fd5b506102166004356024356044356064356106d7565b604080519115158252519081900360200190f35b34801561023657600080fd5b50610216600160a060020a03600435166108bb565b34801561025757600080fd5b506101cb6108d0565b34801561026c57600080fd5b506102786004356108d7565b6040805192835260208301919091528051918290030190f35b34801561029d57600080fd5b506102166004356024356108f3565b3480156102b857600080fd5b50610216600160a060020a0360043516610f78565b3480156102d957600080fd5b506102e8600435602435610fbf565b60408051888152600160a060020a038716918101919091528415156060820152831515608082015260a0810183905260c0810182905260e0602080830182815289519284019290925288516101008401918a019080838360005b8381101561035a578181015183820152602001610342565b50505050905090810190601f1680156103875780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156103a757600080fd5b506101cb6110b3565b3480156103bc57600080fd5b506101cb6004356110b8565b3480156103d457600080fd5b506101cb6004356024356110cd565b3480156103ef57600080fd5b50610216600160a060020a03600435166110ed565b34801561041057600080fd5b50610216600435602435611134565b34801561042b57600080fd5b50610216600160a060020a03600435166111ee565b34801561044c57600080fd5b5061045b60043560243561124f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561049557818101518382015260200161047d565b50505050905090810190601f1680156104c25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104dc57600080fd5b506104e56112fd565b60408051600160a060020a039092168252519081900360200190f35b34801561050d57600080fd5b506104e561130c565b34801561052257600080fd5b5061021660048035906024803580820192908101359160443580820192908101359160643590810191013560843561131b565b34801561056157600080fd5b50610216600435602435604435606435611753565b34801561058257600080fd5b506101cb611b5d565b34801561059757600080fd5b50610216600160a060020a0360043516611b62565b3480156105b857600080fd5b506101cb60043560243560443560643560843560a435611b77565b3480156105df57600080fd5b506104e5612328565b3480156105f457600080fd5b50610216600160a060020a0360043516612337565b34801561061557600080fd5b506104e5612382565b34801561062a57600080fd5b50610216600435612391565b34801561064257600080fd5b506102166004356123a7565b34801561065a57600080fd5b50610216600160a060020a036004351661272c565b34801561067b57600080fd5b506104e561278d565b34801561069057600080fd5b50610216600160a060020a036004351661279c565b6000918252600a6020908152604080842092845291905290205490565b60009081526009602052604090206003015490565b336000818152600b602052604081205490919060ff1615156106f857600080fd5b6006546000878152600a6020908152604080832089845282528083205481517fcdd6ceb9000000000000000000000000000000000000000000000000000000008152600481019190915233602482015260448101899052606481018890529051600160a060020a039094169363cdd6ceb993608480840194938390030190829087803b15801561078757600080fd5b505af115801561079b573d6000803e3d6000fd5b505050506040513d60208110156107b157600080fd5b50506000868152600a602090815260408083208884528252808320338085526005820184529382902060028082018a9055905483516001808c1496820187905294810182905260608082529385018054600019968116156101000296909601909516929092049282018390528a948c947f864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab94909391929190819060808201908690801561089f5780601f106108745761010080835404028352916020019161089f565b820191906000526020600020905b81548152906001019060200180831161088257829003601f168201915b505094505050505060405180910390a450600195945050505050565b600c6020526000908152604090205460ff1681565b62093a8081565b6009602052600090815260409020600381015460049091015482565b336000818152600b6020526040812054909182918291829160ff16151561091957600080fd5b6000878152600a60209081526040808320898452825280832033845260050190915290206003015460ff16156109bf576040805160e560020a62461bcd02815260206004820152602760248201527f45616368207061727469636970616e742063616e206f6e6c792066696e616c6960448201527f7a65206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000878152600a6020908152604080832089845282528083203384526005019091528120541180610a1457506000878152600a60209081526040808320898452909152902060020154600160a060020a031633145b1515610a90576040805160e560020a62461bcd02815260206004820152603a60248201527f4f6e6c79207061727469636970616e7473206f662074686520766f746520617260448201527f652061626c6520746f207265736f6c76652074686520766f7465000000000000606482015290519081900360840190fd5b6000878152600a6020908152604080832089845290915290206002015460a060020a900460ff161515610ca6576006546000888152600a602090815260408083208a845282528083205481517f865df0ad00000000000000000000000000000000000000000000000000000000815260048101919091529051600160a060020a039094169363865df0ad93602480840194938390030190829087803b158015610b3857600080fd5b505af1158015610b4c573d6000803e3d6000fd5b505050506040513d6020811015610b6257600080fd5b50516000888152600a602090815260408083208a84528252918290206002808201805474ff000000000000000000000000000000000000000019961515750100000000000000000000000000000000000000000090810275ff00000000000000000000000000000000000000000019909216919091179690961660a060020a1790819055845195900460ff1680151593860193909352838552600191820180549283161561010002600019019092160492840183905289938b937f56c2d720d2b0f46900eca91b5412e4bb9ef934c72b86308c576d07975fac83539391908190606082019085908015610c965780601f10610c6b57610100808354040283529160200191610c96565b820191906000526020600020905b815481529060010190602001808311610c7957829003601f168201915b5050935050505060405180910390a35b6006546000888152600a602090815260408083208a8452825280832080543380865260059092019093528184206002015482517f36bf4c9100000000000000000000000000000000000000000000000000000000815260048101949094526024840152604483015251600160a060020a03909316926336bf4c9192606480840193606093929083900390910190829087803b158015610d4457600080fd5b505af1158015610d58573d6000803e3d6000fd5b505050506040513d6060811015610d6e57600080fd5b508051602082015160409092015190955090935091506000841115610e3157600854604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018790529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b158015610dfa57600080fd5b505af1158015610e0e573d6000803e3d6000fd5b505050506040513d6020811015610e2457600080fd5b50511515610e3157600080fd5b6000878152600a60209081526040808320898452909152902060040154610e5e908563ffffffff6127e716565b6000888152600a602090815260408083208a8452808352818420600481019590955533808552600586018452828520600301805460ff19166001908117909155948c905290835281518815159381019390935290820186905260608083529383018054600294811615610100026000190116939093049382018490529289928b927fa4a57ebc87f48fa9e8fc4d0812bf408bff87f2326e00c7d209a0d42185b79ec592899189918190608082019086908015610f5b5780601f10610f3057610100808354040283529160200191610f5b565b820191906000526020600020905b815481529060010190602001808311610f3e57829003601f168201915b505094505050505060405180910390a45060019695505050505050565b336000818152600c602052604081205490919060ff161515610f9957600080fd5b5050600160a060020a03166000908152600b60205260409020805460ff19169055600190565b600a6020908152600092835260408084208252918352918190208054600180830180548551600293821615610100026000190190911692909204601f81018790048702830187019095528482529194929390928301828280156110635780601f1061103857610100808354040283529160200191611063565b820191906000526020600020905b81548152906001019060200180831161104657829003601f168201915b505050600284015460038501546004909501549394600160a060020a0382169460ff60a060020a840481169550750100000000000000000000000000000000000000000090930490921692509087565b603c81565b60009081526009602052604090206004015490565b6000918252600a6020908152604080842092845291905290206003015490565b336000818152600c602052604081205490919060ff16151561110e57600080fd5b5050600160a060020a03166000908152600c60205260409020805460ff19169055600190565b6000828152600a6020908152604080832084845290915281206002015460a060020a900460ff1615156111b1576040805160e560020a62461bcd02815260206004820152601d60248201527f50726f706f73616c206861736e2774206265656e207265736f6c766564000000604482015290519081900360640190fd5b506000918252600a602090815260408084209284529190529020600201547501000000000000000000000000000000000000000000900460ff1690565b600554600090600160a060020a0316331461120857600080fd5b600160a060020a038216151561121d57600080fd5b5060058054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b6000828152600960205260409020606090826003811061126b57fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156112f05780601f106112c5576101008083540402835291602001916112f0565b820191906000526020600020905b8154815290600101906020018083116112d357829003601f168201915b5050505050905092915050565b600154600160a060020a031681565b600454600160a060020a031681565b336000818152600b602052604081205490919060ff16151561133c57600080fd5b8915156113b9576040805160e560020a62461bcd02815260206004820152602360248201527f5265706f7369746f72792064657363726970746f722063616e2774206265207a60448201527f65726f0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008a81526009602052604090206004015415611446576040805160e560020a62461bcd02815260206004820152602760248201527f5265706f7369746f72792063616e206f6e6c7920626520696e697469616c697a60448201527f6564206f6e636500000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b603283101561149f576040805160e560020a62461bcd02815260206004820152601f60248201527f566f74696e67206d616a6f726974792068617320746f206265203e3d20353000604482015290519081900360640190fd5b60008811806114ae5750600086115b806114b95750600084115b1515611534576040805160e560020a62461bcd028152602060048201526024808201527f50726f76696465206174206c65617374206f6e65204b6e6f776c65646765204c60448201527f6162656c00000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600454600160a060020a0316156115bb576040805160e560020a62461bcd02815260206004820152602260248201527f54686572652069732061206e6577657220636f6e7472616374206465706c6f7960448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6040805160e06020601f8c01819004028201810190925260c081018a81529091829160608301918291908e908e908190870183828082843782019150505050505081526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375050509290935250505081526000602080830182905260409283018790528d82526009905220815161169890829060036128f9565b50602082810151600383015560409283015160049283015560065483517f828c16530000000000000000000000000000000000000000000000000000000081529283018e9052602483018790529251600160a060020a039093169263828c16539260448082019392918290030181600087803b15801561171757600080fd5b505af115801561172b573d6000803e3d6000fd5b505050506040513d602081101561174157600080fd5b5060019b9a5050505050505050505050565b336000818152600b60205260408120549091829160ff16151561177557600080fd5b6000878152600a60209081526040808320898452909152902060020154600160a060020a0316331415611818576040805160e560020a62461bcd02815260206004820152603160248201527f5468652070726f706f736572206973206e6f7420616c6c6f77656420746f207660448201527f6f746520696e20612070726f706f73616c000000000000000000000000000000606482015290519081900360840190fd5b6008546000888152600a602090815260408083208a845282528083206003015481517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101919091529051600160a060020a03909416936323b872dd93606480840194938390030190829087803b1580156118a257600080fd5b505af11580156118b6573d6000803e3d6000fd5b505050506040513d60208110156118cc57600080fd5b505115156118d957600080fd5b6000878152600a60209081526040808320898452909152902060038101546004909101546119069161286f565b6000888152600a602090815260408083208a84528252808320600480820195909555600654905482517fd4e0ac9500000000000000000000000000000000000000000000000000000000815295860152336024860152604485018a9052606485018990529051600160a060020a039091169363d4e0ac959360848083019493928390030190829087803b15801561199c57600080fd5b505af11580156119b0573d6000803e3d6000fd5b505050506040513d60208110156119c657600080fd5b5051915060008211611a48576040805160e560020a62461bcd02815260206004820152603360248201527f566f74696e6720636f6e74726163742072657475726e656420616e20696e766160448201527f6c696420616d6f756e74206f6620766f74657300000000000000000000000000606482015290519081900360840190fd5b6000878152600a6020908152604080832089845280835281842033808652600582018552838620889055948b90529083526003810154825193840181905291830188905260608301869052608080845260019182018054600261010094821615949094026000190116929092049084018190528a938c937f7ee8ecf4d9d20fb6454a55c418451d97bec229903525d8517fcd32db68a479e4939290918b918a9190819060a082019087908015611b3f5780601f10611b1457610100808354040283529160200191611b3f565b820191906000526020600020905b815481529060010190602001808311611b2257829003601f168201915b50509550505050505060405180910390a45060019695505050505050565b600081565b600b6020526000908152604090205460ff1681565b336000818152600b602052604081205490919060ff161515611b9857600080fd5b60008311611c16576040805160e560020a62461bcd02815260206004820152602860248201527f56616c7565206f6620746865207472616e73616374696f6e2063616e206e6f7460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008881526009602052604081208860038110611c2f57fe5b01546002600019610100600184161502019091160411611cbe576040805160e560020a62461bcd028152602060048201526024808201527f4b6e6f776c656467652d4c6162656c20696e646578206973206e6f7420636f7260448201527f7265637400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b603c8510158015611cd2575062093a808511155b1515611d28576040805160e560020a62461bcd02815260206004820152601c60248201527f566f746520636f6d6d6974206475726174696f6e20696e76616c696400000000604482015290519081900360640190fd5b603c8410158015611d3c575062093a808411155b1515611d92576040805160e560020a62461bcd02815260206004820152601a60248201527f566f7465206f70656e206475726174696f6e20696e76616c6964000000000000604482015290519081900360640190fd5b600454600160a060020a031615611e19576040805160e560020a62461bcd02815260206004820152602260248201527f54686572652069732061206e6577657220636f6e7472616374206465706c6f7960448201527f6564000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600088815260096020526040902060030154611e3c90600163ffffffff61286f16565b60008981526009602052604090819020600380820193909355815160e0810190925260065491928392600160a060020a031691636766b4a4918d9133918e908110611e8357fe5b018b8b8b8f6040518863ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180886000191660001916815260200187600160a060020a0316600160a060020a0316815260200180602001868152602001858152602001848152602001838152602001828103825287818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611f7a5780601f10611f4f57610100808354040283529160200191611f7a565b820191906000526020600020905b815481529060010190602001808311611f5d57829003601f168201915b505098505050505050505050602060405180830381600087803b158015611fa057600080fd5b505af1158015611fb4573d6000803e3d6000fd5b505050506040513d6020811015611fca57600080fd5b5051815260008a81526009602090815260409091209101908960038110611fed57fe5b01805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156120725780601f1061204757610100808354040283529160200191612072565b820191906000526020600020905b81548152906001019060200180831161205557829003601f168201915b50505091835250503360208083019190915260006040808401829052606084018290526080840188905260a09093018790528b8152600a8252828120600983528382206003015482528252919091208251815582820151805191926120df92600185019290910190612949565b506040828101516002830180546060860151608087015173ffffffffffffffffffffffffffffffffffffffff19909216600160a060020a039485161774ff0000000000000000000000000000000000000000191660a060020a911515919091021775ff000000000000000000000000000000000000000000191675010000000000000000000000000000000000000000009115159190910217905560a0840151600384015560c09093015160049283015560085481517f23b872dd00000000000000000000000000000000000000000000000000000000815233938101939093523060248401526044830187905290519216916323b872dd916064808201926020929091908290030181600087803b1580156121fa57600080fd5b505af115801561220e573d6000803e3d6000fd5b505050506040513d602081101561222457600080fd5b5051151561223157600080fd5b6000888152600960205260409020600380820154339290918b917f2ba422bdea9179f02f8c9dd0d6285478f7b4c3fa11a812eeb4d3b1b04cc57c35918c90811061227757fe5b60408051602081018e905281815292909101805460026000196101006001841615020190911604918301829052918c9181906060820190859080156122fd5780601f106122d2576101008083540402835291602001916122fd565b820191906000526020600020905b8154815290600101906020018083116122e057829003601f168201915b5050935050505060405180910390a45050506000948552505060096020525050604090206003015490565b600254600160a060020a031681565b336000818152600c602052604081205490919060ff16151561235857600080fd5b5050600160a060020a03166000908152600c60205260409020805460ff1916600190811790915590565b600354600160a060020a031681565b6000908152600960205260408120600401541190565b6000806000806123b56129c3565b336000818152600b602052604081205490919060ff1615156123d657600080fd5b600354600160a060020a031615156123ed57600080fd5b600354604080517f0bdc90e8000000000000000000000000000000000000000000000000000000008152600481018b90529051600160a060020a0390921697508791630bdc90e8916024808201926020929091908290030181600087803b15801561245757600080fd5b505af115801561246b573d6000803e3d6000fd5b505050506040513d602081101561248157600080fd5b5051604080517f3fcc148d000000000000000000000000000000000000000000000000000000008152600481018b90529051919650600160a060020a03881691633fcc148d916024808201926020929091908290030181600087803b1580156124e957600080fd5b505af11580156124fd573d6000803e3d6000fd5b505050506040513d602081101561251357600080fd5b50519350600091505b60038260ff16101561263b57604080517f95332229000000000000000000000000000000000000000000000000000000008152600481018a905260ff841660248201529051600160a060020a03881691639533222991604480830192600092919082900301818387803b15801561259257600080fd5b505af11580156125a6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156125cf57600080fd5b8101908080516401000000008111156125e757600080fd5b820160208101848111156125fa57600080fd5b815164010000000081118282018710171561261457600080fd5b50909350869250505060ff84166003811061262b57fe5b602002015260019091019061251c565b60408051606081018252848152602080820188905281830187905260008b8152600990915291909120815161267390829060036128f9565b50602082810151600383015560409283015160049283015560065483517f828c16530000000000000000000000000000000000000000000000000000000081529283018c9052602483018890529251600160a060020a039093169263828c16539260448082019392918290030181600087803b1580156126f257600080fd5b505af1158015612706573d6000803e3d6000fd5b505050506040513d602081101561271c57600080fd5b5060019998505050505050505050565b600554600090600160a060020a0316331461274657600080fd5b600160a060020a038216151561275b57600080fd5b5060048054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff199091161790556001919050565b600054600160a060020a031681565b336000818152600c602052604081205490919060ff1615156127bd57600080fd5b5050600160a060020a03166000908152600b60205260409020805460ff1916600190811790915590565b60008083831115612868576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b6000828201838110156128f2576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b8260038101928215612939579160200282015b828111156129395782518051612929918491602090910190612949565b509160200191906001019061290c565b506129459291506129eb565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061298a57805160ff19168380011785556129b7565b828001600101855582156129b7579182015b828111156129b757825182559160200191906001019061299c565b50612945929150612a11565b6060604051908101604052806003905b60608152602001906001900390816129d35790505090565b612a0e91905b80821115612945576000612a058282612a2b565b506001016129f1565b90565b612a0e91905b808211156129455760008155600101612a17565b50805460018160011615610100020316600290046000825580601f10612a515750612a6f565b601f016020900490600052602060002090810190612a6f9190612a11565b505600a165627a7a72305820326e30d519b4497b0b1c36a0772a74ca0df7dc86668d15c182981faddd6bf6750029`

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

// MAXOPENDURATION is a free data retrieval call binding the contract method 0x466577cd.
//
// Solidity: function MAX_OPEN_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) MAXOPENDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "MAX_OPEN_DURATION")
	return *ret0, err
}

// MAXOPENDURATION is a free data retrieval call binding the contract method 0x466577cd.
//
// Solidity: function MAX_OPEN_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MAXOPENDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MAXOPENDURATION(&_DitDemoCoordinator.CallOpts)
}

// MAXOPENDURATION is a free data retrieval call binding the contract method 0x466577cd.
//
// Solidity: function MAX_OPEN_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) MAXOPENDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MAXOPENDURATION(&_DitDemoCoordinator.CallOpts)
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

// MINOPENDURATION is a free data retrieval call binding the contract method 0x6fcfeb3b.
//
// Solidity: function MIN_OPEN_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) MINOPENDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "MIN_OPEN_DURATION")
	return *ret0, err
}

// MINOPENDURATION is a free data retrieval call binding the contract method 0x6fcfeb3b.
//
// Solidity: function MIN_OPEN_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MINOPENDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MINOPENDURATION(&_DitDemoCoordinator.CallOpts)
}

// MINOPENDURATION is a free data retrieval call binding the contract method 0x6fcfeb3b.
//
// Solidity: function MIN_OPEN_DURATION() constant returns(uint256)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) MINOPENDURATION() (*big.Int, error) {
	return _DitDemoCoordinator.Contract.MINOPENDURATION(&_DitDemoCoordinator.CallOpts)
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

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) GetKnowledgeLabels(opts *bind.CallOpts, _repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DitDemoCoordinator.contract.Call(opts, out, "getKnowledgeLabels", _repository, _knowledgeLabelID)
	return *ret0, err
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitDemoCoordinator.Contract.GetKnowledgeLabels(&_DitDemoCoordinator.CallOpts, _repository, _knowledgeLabelID)
}

// GetKnowledgeLabels is a free data retrieval call binding the contract method 0x95332229.
//
// Solidity: function getKnowledgeLabels(bytes32 _repository, uint256 _knowledgeLabelID) constant returns(string knowledgeLabel)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) GetKnowledgeLabels(_repository [32]byte, _knowledgeLabelID *big.Int) (string, error) {
	return _DitDemoCoordinator.Contract.GetKnowledgeLabels(&_DitDemoCoordinator.CallOpts, _repository, _knowledgeLabelID)
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) ProposalsOfRepository(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
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
	err := _DitDemoCoordinator.contract.Call(opts, out, "proposalsOfRepository", arg0, arg1)
	return *ret, err
}

// ProposalsOfRepository is a free data retrieval call binding the contract method 0x3e029f63.
//
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
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
// Solidity: function proposalsOfRepository(bytes32 , uint256 ) constant returns(uint256 KNWVoteID, string knowledgeLabel, address proposer, bool isFinalized, bool proposalAccepted, uint256 individualStake, uint256 totalStake)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) ProposalsOfRepository(arg0 [32]byte, arg1 *big.Int) (struct {
	KNWVoteID        *big.Int
	KnowledgeLabel   string
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
// Solidity: function repositories(bytes32 ) constant returns(uint256 currentProposalID, uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorCaller) Repositories(opts *bind.CallOpts, arg0 [32]byte) (struct {
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	ret := new(struct {
		CurrentProposalID *big.Int
		VotingMajority    *big.Int
	})
	out := ret
	err := _DitDemoCoordinator.contract.Call(opts, out, "repositories", arg0)
	return *ret, err
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(uint256 currentProposalID, uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) Repositories(arg0 [32]byte) (struct {
	CurrentProposalID *big.Int
	VotingMajority    *big.Int
}, error) {
	return _DitDemoCoordinator.Contract.Repositories(&_DitDemoCoordinator.CallOpts, arg0)
}

// Repositories is a free data retrieval call binding the contract method 0x1f51fd71.
//
// Solidity: function repositories(bytes32 ) constant returns(uint256 currentProposalID, uint256 votingMajority)
func (_DitDemoCoordinator *DitDemoCoordinatorCallerSession) Repositories(arg0 [32]byte) (struct {
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

// InitRepository is a paid mutator transaction binding the contract method 0x9a401f48.
//
// Solidity: function initRepository(bytes32 _repository, string _label1, string _label2, string _label3, uint256 _votingMajority) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) InitRepository(opts *bind.TransactOpts, _repository [32]byte, _label1 string, _label2 string, _label3 string, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "initRepository", _repository, _label1, _label2, _label3, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x9a401f48.
//
// Solidity: function initRepository(bytes32 _repository, string _label1, string _label2, string _label3, uint256 _votingMajority) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) InitRepository(_repository [32]byte, _label1 string, _label2 string, _label3 string, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.InitRepository(&_DitDemoCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _votingMajority)
}

// InitRepository is a paid mutator transaction binding the contract method 0x9a401f48.
//
// Solidity: function initRepository(bytes32 _repository, string _label1, string _label2, string _label3, uint256 _votingMajority) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) InitRepository(_repository [32]byte, _label1 string, _label2 string, _label3 string, _votingMajority *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.InitRepository(&_DitDemoCoordinator.TransactOpts, _repository, _label1, _label2, _label3, _votingMajority)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xea976d2c.
//
// Solidity: function migrateRepository(bytes32 _repository) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) MigrateRepository(opts *bind.TransactOpts, _repository [32]byte) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "migrateRepository", _repository)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xea976d2c.
//
// Solidity: function migrateRepository(bytes32 _repository) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) MigrateRepository(_repository [32]byte) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.MigrateRepository(&_DitDemoCoordinator.TransactOpts, _repository)
}

// MigrateRepository is a paid mutator transaction binding the contract method 0xea976d2c.
//
// Solidity: function migrateRepository(bytes32 _repository) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) MigrateRepository(_repository [32]byte) (*types.Transaction, error) {
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

// ProposeCommit is a paid mutator transaction binding the contract method 0xcda8d421.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _numberOfKNW, uint256 _voteCommitDuration, uint256 _voteOpenDuration, uint256 _amountOfTokens) returns(uint256 proposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) ProposeCommit(opts *bind.TransactOpts, _repository [32]byte, _knowledgeLabelIndex *big.Int, _numberOfKNW *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "proposeCommit", _repository, _knowledgeLabelIndex, _numberOfKNW, _voteCommitDuration, _voteOpenDuration, _amountOfTokens)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xcda8d421.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _numberOfKNW, uint256 _voteCommitDuration, uint256 _voteOpenDuration, uint256 _amountOfTokens) returns(uint256 proposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _numberOfKNW *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ProposeCommit(&_DitDemoCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _numberOfKNW, _voteCommitDuration, _voteOpenDuration, _amountOfTokens)
}

// ProposeCommit is a paid mutator transaction binding the contract method 0xcda8d421.
//
// Solidity: function proposeCommit(bytes32 _repository, uint256 _knowledgeLabelIndex, uint256 _numberOfKNW, uint256 _voteCommitDuration, uint256 _voteOpenDuration, uint256 _amountOfTokens) returns(uint256 proposalID)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) ProposeCommit(_repository [32]byte, _knowledgeLabelIndex *big.Int, _numberOfKNW *big.Int, _voteCommitDuration *big.Int, _voteOpenDuration *big.Int, _amountOfTokens *big.Int) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ProposeCommit(&_DitDemoCoordinator.TransactOpts, _repository, _knowledgeLabelIndex, _numberOfKNW, _voteCommitDuration, _voteOpenDuration, _amountOfTokens)
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

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactor) ReplaceDitManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.contract.Transact(opts, "replaceDitManager", _newManager)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorSession) ReplaceDitManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ReplaceDitManager(&_DitDemoCoordinator.TransactOpts, _newManager)
}

// ReplaceDitManager is a paid mutator transaction binding the contract method 0x91016157.
//
// Solidity: function replaceDitManager(address _newManager) returns(bool)
func (_DitDemoCoordinator *DitDemoCoordinatorTransactorSession) ReplaceDitManager(_newManager common.Address) (*types.Transaction, error) {
	return _DitDemoCoordinator.Contract.ReplaceDitManager(&_DitDemoCoordinator.TransactOpts, _newManager)
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
	Label         string
	Stake         *big.Int
	NumberOfKNW   *big.Int
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitVote is a free log retrieval operation binding the contract event 0x7ee8ecf4d9d20fb6454a55c418451d97bec229903525d8517fcd32db68a479e4.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
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

// WatchCommitVote is a free log subscription operation binding the contract event 0x7ee8ecf4d9d20fb6454a55c418451d97bec229903525d8517fcd32db68a479e4.
//
// Solidity: event CommitVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 stake, uint256 numberOfKNW, uint256 numberOfVotes)
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
	Repository [32]byte
	Proposal   *big.Int
	Label      string
	Accepted   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeProposal is a free log retrieval operation binding the contract event 0x56c2d720d2b0f46900eca91b5412e4bb9ef934c72b86308c576d07975fac8353.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, string label, bool accepted)
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

// WatchFinalizeProposal is a free log subscription operation binding the contract event 0x56c2d720d2b0f46900eca91b5412e4bb9ef934c72b86308c576d07975fac8353.
//
// Solidity: event FinalizeProposal(bytes32 indexed repository, uint256 indexed proposal, string label, bool accepted)
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
	Label       string
	VotedRight  bool
	NumberOfKNW *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizeVote is a free log retrieval operation binding the contract event 0xa4a57ebc87f48fa9e8fc4d0812bf408bff87f2326e00c7d209a0d42185b79ec5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool votedRight, uint256 numberOfKNW)
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

// WatchFinalizeVote is a free log subscription operation binding the contract event 0xa4a57ebc87f48fa9e8fc4d0812bf408bff87f2326e00c7d209a0d42185b79ec5.
//
// Solidity: event FinalizeVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool votedRight, uint256 numberOfKNW)
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
	Label         string
	Accept        bool
	NumberOfVotes *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenVote is a free log retrieval operation binding the contract event 0x864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool accept, uint256 numberOfVotes)
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

// WatchOpenVote is a free log subscription operation binding the contract event 0x864c0d6987266fd72e7e37f1fbc98b6a3794b7187dae454c67a2a626628a72ab.
//
// Solidity: event OpenVote(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, bool accept, uint256 numberOfVotes)
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
	Label       string
	NumberOfKNW *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposeCommit is a free log retrieval operation binding the contract event 0x2ba422bdea9179f02f8c9dd0d6285478f7b4c3fa11a812eeb4d3b1b04cc57c35.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 numberOfKNW)
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

// WatchProposeCommit is a free log subscription operation binding the contract event 0x2ba422bdea9179f02f8c9dd0d6285478f7b4c3fa11a812eeb4d3b1b04cc57c35.
//
// Solidity: event ProposeCommit(bytes32 indexed repository, uint256 indexed proposal, address indexed who, string label, uint256 numberOfKNW)
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
