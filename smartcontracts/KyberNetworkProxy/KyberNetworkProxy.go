// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KyberNetwork

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

// KyberNetworkABI is the input ABI used to generate the binding from.
const KyberNetworkABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"},{\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"tradeWithHint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserCapInWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminQuickly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAlerters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"srcQty\",\"type\":\"uint256\"}],\"name\":\"getExpectedRate\",\"outputs\":[{\"name\":\"expectedRate\",\"type\":\"uint256\"},{\"name\":\"slippageRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getUserCapInTokenWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kyberNetworkContract\",\"type\":\"address\"}],\"name\":\"setKyberNetworkContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"info\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"name\":\"dest\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"maxDestAmount\",\"type\":\"uint256\"},{\"name\":\"minConversionRate\",\"type\":\"uint256\"},{\"name\":\"walletId\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"actualSrcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"actualDestAmount\",\"type\":\"uint256\"}],\"name\":\"ExecuteTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newNetworkContract\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldNetworkContract\",\"type\":\"address\"}],\"name\":\"KyberNetworkSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"EtherWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"AdminClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AlerterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"}]"

// KyberNetwork is an auto generated Go binding around an Ethereum contract.
type KyberNetwork struct {
	KyberNetworkCaller     // Read-only binding to the contract
	KyberNetworkTransactor // Write-only binding to the contract
	KyberNetworkFilterer   // Log filterer for contract events
}

// KyberNetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberNetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberNetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberNetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberNetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberNetworkSession struct {
	Contract     *KyberNetwork     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberNetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberNetworkCallerSession struct {
	Contract *KyberNetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberNetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberNetworkTransactorSession struct {
	Contract     *KyberNetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberNetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberNetworkRaw struct {
	Contract *KyberNetwork // Generic contract binding to access the raw methods on
}

// KyberNetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberNetworkCallerRaw struct {
	Contract *KyberNetworkCaller // Generic read-only contract binding to access the raw methods on
}

// KyberNetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberNetworkTransactorRaw struct {
	Contract *KyberNetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberNetwork creates a new instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetwork(address common.Address, backend bind.ContractBackend) (*KyberNetwork, error) {
	contract, err := bindKyberNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KyberNetwork{KyberNetworkCaller: KyberNetworkCaller{contract: contract}, KyberNetworkTransactor: KyberNetworkTransactor{contract: contract}, KyberNetworkFilterer: KyberNetworkFilterer{contract: contract}}, nil
}

// NewKyberNetworkCaller creates a new read-only instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetworkCaller(address common.Address, caller bind.ContractCaller) (*KyberNetworkCaller, error) {
	contract, err := bindKyberNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkCaller{contract: contract}, nil
}

// NewKyberNetworkTransactor creates a new write-only instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberNetworkTransactor, error) {
	contract, err := bindKyberNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkTransactor{contract: contract}, nil
}

// NewKyberNetworkFilterer creates a new log filterer instance of KyberNetwork, bound to a specific deployed contract.
func NewKyberNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberNetworkFilterer, error) {
	contract, err := bindKyberNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkFilterer{contract: contract}, nil
}

// bindKyberNetwork binds a generic wrapper to an already deployed contract.
func bindKyberNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberNetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetwork *KyberNetworkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetwork.Contract.KyberNetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetwork *KyberNetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetwork.Contract.KyberNetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetwork *KyberNetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetwork.Contract.KyberNetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KyberNetwork *KyberNetworkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KyberNetwork.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KyberNetwork *KyberNetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetwork.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KyberNetwork *KyberNetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KyberNetwork.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetwork *KyberNetworkCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetwork *KyberNetworkSession) Admin() (common.Address, error) {
	return _KyberNetwork.Contract.Admin(&_KyberNetwork.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_KyberNetwork *KyberNetworkCallerSession) Admin() (common.Address, error) {
	return _KyberNetwork.Contract.Admin(&_KyberNetwork.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetwork *KyberNetworkCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetwork *KyberNetworkSession) Enabled() (bool, error) {
	return _KyberNetwork.Contract.Enabled(&_KyberNetwork.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_KyberNetwork *KyberNetworkCallerSession) Enabled() (bool, error) {
	return _KyberNetwork.Contract.Enabled(&_KyberNetwork.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetwork *KyberNetworkCaller) GetAlerters(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "getAlerters")
	return *ret0, err
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetwork *KyberNetworkSession) GetAlerters() ([]common.Address, error) {
	return _KyberNetwork.Contract.GetAlerters(&_KyberNetwork.CallOpts)
}

// GetAlerters is a free data retrieval call binding the contract method 0x7c423f54.
//
// Solidity: function getAlerters() constant returns(address[])
func (_KyberNetwork *KyberNetworkCallerSession) GetAlerters() ([]common.Address, error) {
	return _KyberNetwork.Contract.GetAlerters(&_KyberNetwork.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCaller) GetBalance(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "getBalance", token, user)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetwork *KyberNetworkSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _KyberNetwork.Contract.GetBalance(&_KyberNetwork.CallOpts, token, user)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address token, address user) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCallerSession) GetBalance(token common.Address, user common.Address) (*big.Int, error) {
	return _KyberNetwork.Contract.GetBalance(&_KyberNetwork.CallOpts, token, user)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkCaller) GetExpectedRate(opts *bind.CallOpts, src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	ret := new(struct {
		ExpectedRate *big.Int
		SlippageRate *big.Int
	})
	out := ret
	err := _KyberNetwork.contract.Call(opts, out, "getExpectedRate", src, dest, srcQty)
	return *ret, err
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetwork.Contract.GetExpectedRate(&_KyberNetwork.CallOpts, src, dest, srcQty)
}

// GetExpectedRate is a free data retrieval call binding the contract method 0x809a9e55.
//
// Solidity: function getExpectedRate(address src, address dest, uint256 srcQty) constant returns(uint256 expectedRate, uint256 slippageRate)
func (_KyberNetwork *KyberNetworkCallerSession) GetExpectedRate(src common.Address, dest common.Address, srcQty *big.Int) (struct {
	ExpectedRate *big.Int
	SlippageRate *big.Int
}, error) {
	return _KyberNetwork.Contract.GetExpectedRate(&_KyberNetwork.CallOpts, src, dest, srcQty)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetwork *KyberNetworkCaller) GetOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "getOperators")
	return *ret0, err
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetwork *KyberNetworkSession) GetOperators() ([]common.Address, error) {
	return _KyberNetwork.Contract.GetOperators(&_KyberNetwork.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() constant returns(address[])
func (_KyberNetwork *KyberNetworkCallerSession) GetOperators() ([]common.Address, error) {
	return _KyberNetwork.Contract.GetOperators(&_KyberNetwork.CallOpts)
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCaller) GetUserCapInTokenWei(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "getUserCapInTokenWei", user, token)
	return *ret0, err
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetwork *KyberNetworkSession) GetUserCapInTokenWei(user common.Address, token common.Address) (*big.Int, error) {
	return _KyberNetwork.Contract.GetUserCapInTokenWei(&_KyberNetwork.CallOpts, user, token)
}

// GetUserCapInTokenWei is a free data retrieval call binding the contract method 0x8eaaeecf.
//
// Solidity: function getUserCapInTokenWei(address user, address token) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCallerSession) GetUserCapInTokenWei(user common.Address, token common.Address) (*big.Int, error) {
	return _KyberNetwork.Contract.GetUserCapInTokenWei(&_KyberNetwork.CallOpts, user, token)
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCaller) GetUserCapInWei(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "getUserCapInWei", user)
	return *ret0, err
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetwork *KyberNetworkSession) GetUserCapInWei(user common.Address) (*big.Int, error) {
	return _KyberNetwork.Contract.GetUserCapInWei(&_KyberNetwork.CallOpts, user)
}

// GetUserCapInWei is a free data retrieval call binding the contract method 0x6432679f.
//
// Solidity: function getUserCapInWei(address user) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCallerSession) GetUserCapInWei(user common.Address) (*big.Int, error) {
	return _KyberNetwork.Contract.GetUserCapInWei(&_KyberNetwork.CallOpts, user)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCaller) Info(opts *bind.CallOpts, field [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "info", field)
	return *ret0, err
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetwork *KyberNetworkSession) Info(field [32]byte) (*big.Int, error) {
	return _KyberNetwork.Contract.Info(&_KyberNetwork.CallOpts, field)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 field) constant returns(uint256)
func (_KyberNetwork *KyberNetworkCallerSession) Info(field [32]byte) (*big.Int, error) {
	return _KyberNetwork.Contract.Info(&_KyberNetwork.CallOpts, field)
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetwork *KyberNetworkCaller) KyberNetworkContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "kyberNetworkContract")
	return *ret0, err
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetwork *KyberNetworkSession) KyberNetworkContract() (common.Address, error) {
	return _KyberNetwork.Contract.KyberNetworkContract(&_KyberNetwork.CallOpts)
}

// KyberNetworkContract is a free data retrieval call binding the contract method 0x4f61ff8b.
//
// Solidity: function kyberNetworkContract() constant returns(address)
func (_KyberNetwork *KyberNetworkCallerSession) KyberNetworkContract() (common.Address, error) {
	return _KyberNetwork.Contract.KyberNetworkContract(&_KyberNetwork.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetwork *KyberNetworkCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "maxGasPrice")
	return *ret0, err
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetwork *KyberNetworkSession) MaxGasPrice() (*big.Int, error) {
	return _KyberNetwork.Contract.MaxGasPrice(&_KyberNetwork.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() constant returns(uint256)
func (_KyberNetwork *KyberNetworkCallerSession) MaxGasPrice() (*big.Int, error) {
	return _KyberNetwork.Contract.MaxGasPrice(&_KyberNetwork.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetwork *KyberNetworkCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KyberNetwork.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetwork *KyberNetworkSession) PendingAdmin() (common.Address, error) {
	return _KyberNetwork.Contract.PendingAdmin(&_KyberNetwork.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_KyberNetwork *KyberNetworkCallerSession) PendingAdmin() (common.Address, error) {
	return _KyberNetwork.Contract.PendingAdmin(&_KyberNetwork.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetwork *KyberNetworkTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetwork *KyberNetworkSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.AddAlerter(&_KyberNetwork.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(address newAlerter) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.AddAlerter(&_KyberNetwork.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetwork *KyberNetworkTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetwork *KyberNetworkSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.AddOperator(&_KyberNetwork.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address newOperator) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.AddOperator(&_KyberNetwork.TransactOpts, newOperator)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetwork *KyberNetworkTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetwork *KyberNetworkSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberNetwork.Contract.ClaimAdmin(&_KyberNetwork.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_KyberNetwork *KyberNetworkTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _KyberNetwork.Contract.ClaimAdmin(&_KyberNetwork.TransactOpts)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetwork *KyberNetworkTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetwork *KyberNetworkSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.RemoveAlerter(&_KyberNetwork.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(address alerter) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.RemoveAlerter(&_KyberNetwork.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetwork *KyberNetworkTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetwork *KyberNetworkSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.RemoveOperator(&_KyberNetwork.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.RemoveOperator(&_KyberNetwork.TransactOpts, operator)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetwork *KyberNetworkTransactor) SetKyberNetworkContract(opts *bind.TransactOpts, _kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "setKyberNetworkContract", _kyberNetworkContract)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetwork *KyberNetworkSession) SetKyberNetworkContract(_kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SetKyberNetworkContract(&_KyberNetwork.TransactOpts, _kyberNetworkContract)
}

// SetKyberNetworkContract is a paid mutator transaction binding the contract method 0xabd188a8.
//
// Solidity: function setKyberNetworkContract(address _kyberNetworkContract) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) SetKyberNetworkContract(_kyberNetworkContract common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SetKyberNetworkContract(&_KyberNetwork.TransactOpts, _kyberNetworkContract)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapEtherToToken(opts *bind.TransactOpts, token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapEtherToToken", token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapEtherToToken(&_KyberNetwork.TransactOpts, token, minConversionRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x7a2a0456.
//
// Solidity: function swapEtherToToken(address token, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) SwapEtherToToken(token common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapEtherToToken(&_KyberNetwork.TransactOpts, token, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapTokenToEther(opts *bind.TransactOpts, token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapTokenToEther", token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToEther(&_KyberNetwork.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 srcAmount, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) SwapTokenToEther(token common.Address, srcAmount *big.Int, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToEther(&_KyberNetwork.TransactOpts, token, srcAmount, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) SwapTokenToToken(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "swapTokenToToken", src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToToken(&_KyberNetwork.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address src, uint256 srcAmount, address dest, uint256 minConversionRate) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) SwapTokenToToken(src common.Address, srcAmount *big.Int, dest common.Address, minConversionRate *big.Int) (*types.Transaction, error) {
	return _KyberNetwork.Contract.SwapTokenToToken(&_KyberNetwork.TransactOpts, src, srcAmount, dest, minConversionRate)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) Trade(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "trade", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.Trade(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// Trade is a paid mutator transaction binding the contract method 0xcb3c28c7.
//
// Solidity: function trade(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) Trade(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.Trade(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactor) TradeWithHint(opts *bind.TransactOpts, src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "tradeWithHint", src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetwork *KyberNetworkSession) TradeWithHint(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetwork.Contract.TradeWithHint(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TradeWithHint is a paid mutator transaction binding the contract method 0x29589f61.
//
// Solidity: function tradeWithHint(address src, uint256 srcAmount, address dest, address destAddress, uint256 maxDestAmount, uint256 minConversionRate, address walletId, bytes hint) returns(uint256)
func (_KyberNetwork *KyberNetworkTransactorSession) TradeWithHint(src common.Address, srcAmount *big.Int, dest common.Address, destAddress common.Address, maxDestAmount *big.Int, minConversionRate *big.Int, walletId common.Address, hint []byte) (*types.Transaction, error) {
	return _KyberNetwork.Contract.TradeWithHint(&_KyberNetwork.TransactOpts, src, srcAmount, dest, destAddress, maxDestAmount, minConversionRate, walletId, hint)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetwork *KyberNetworkTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetwork *KyberNetworkSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.TransferAdmin(&_KyberNetwork.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.TransferAdmin(&_KyberNetwork.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetwork *KyberNetworkTransactor) TransferAdminQuickly(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "transferAdminQuickly", newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetwork *KyberNetworkSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.TransferAdminQuickly(&_KyberNetwork.TransactOpts, newAdmin)
}

// TransferAdminQuickly is a paid mutator transaction binding the contract method 0x7acc8678.
//
// Solidity: function transferAdminQuickly(address newAdmin) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) TransferAdminQuickly(newAdmin common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.TransferAdminQuickly(&_KyberNetwork.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetwork *KyberNetworkTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetwork *KyberNetworkSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.WithdrawEther(&_KyberNetwork.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(uint256 amount, address sendTo) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.WithdrawEther(&_KyberNetwork.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetwork *KyberNetworkTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetwork.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetwork *KyberNetworkSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.WithdrawToken(&_KyberNetwork.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address sendTo) returns()
func (_KyberNetwork *KyberNetworkTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _KyberNetwork.Contract.WithdrawToken(&_KyberNetwork.TransactOpts, token, amount, sendTo)
}

// KyberNetworkAdminClaimedIterator is returned from FilterAdminClaimed and is used to iterate over the raw logs and unpacked data for AdminClaimed events raised by the KyberNetwork contract.
type KyberNetworkAdminClaimedIterator struct {
	Event *KyberNetworkAdminClaimed // Event containing the contract specifics and raw log

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
func (it *KyberNetworkAdminClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkAdminClaimed)
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
		it.Event = new(KyberNetworkAdminClaimed)
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
func (it *KyberNetworkAdminClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkAdminClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkAdminClaimed represents a AdminClaimed event raised by the KyberNetwork contract.
type KyberNetworkAdminClaimed struct {
	NewAdmin      common.Address
	PreviousAdmin common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminClaimed is a free log retrieval operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetwork *KyberNetworkFilterer) FilterAdminClaimed(opts *bind.FilterOpts) (*KyberNetworkAdminClaimedIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkAdminClaimedIterator{contract: _KyberNetwork.contract, event: "AdminClaimed", logs: logs, sub: sub}, nil
}

// WatchAdminClaimed is a free log subscription operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetwork *KyberNetworkFilterer) WatchAdminClaimed(opts *bind.WatchOpts, sink chan<- *KyberNetworkAdminClaimed) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "AdminClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkAdminClaimed)
				if err := _KyberNetwork.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
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

// ParseAdminClaimed is a log parse operation binding the contract event 0x65da1cfc2c2e81576ad96afb24a581f8e109b7a403b35cbd3243a1c99efdb9ed.
//
// Solidity: event AdminClaimed(address newAdmin, address previousAdmin)
func (_KyberNetwork *KyberNetworkFilterer) ParseAdminClaimed(log types.Log) (*KyberNetworkAdminClaimed, error) {
	event := new(KyberNetworkAdminClaimed)
	if err := _KyberNetwork.contract.UnpackLog(event, "AdminClaimed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkAlerterAddedIterator is returned from FilterAlerterAdded and is used to iterate over the raw logs and unpacked data for AlerterAdded events raised by the KyberNetwork contract.
type KyberNetworkAlerterAddedIterator struct {
	Event *KyberNetworkAlerterAdded // Event containing the contract specifics and raw log

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
func (it *KyberNetworkAlerterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkAlerterAdded)
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
		it.Event = new(KyberNetworkAlerterAdded)
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
func (it *KyberNetworkAlerterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkAlerterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkAlerterAdded represents a AlerterAdded event raised by the KyberNetwork contract.
type KyberNetworkAlerterAdded struct {
	NewAlerter common.Address
	IsAdd      bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAlerterAdded is a free log retrieval operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetwork *KyberNetworkFilterer) FilterAlerterAdded(opts *bind.FilterOpts) (*KyberNetworkAlerterAddedIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkAlerterAddedIterator{contract: _KyberNetwork.contract, event: "AlerterAdded", logs: logs, sub: sub}, nil
}

// WatchAlerterAdded is a free log subscription operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetwork *KyberNetworkFilterer) WatchAlerterAdded(opts *bind.WatchOpts, sink chan<- *KyberNetworkAlerterAdded) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "AlerterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkAlerterAdded)
				if err := _KyberNetwork.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
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

// ParseAlerterAdded is a log parse operation binding the contract event 0x5611bf3e417d124f97bf2c788843ea8bb502b66079fbee02158ef30b172cb762.
//
// Solidity: event AlerterAdded(address newAlerter, bool isAdd)
func (_KyberNetwork *KyberNetworkFilterer) ParseAlerterAdded(log types.Log) (*KyberNetworkAlerterAdded, error) {
	event := new(KyberNetworkAlerterAdded)
	if err := _KyberNetwork.contract.UnpackLog(event, "AlerterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkEtherWithdrawIterator is returned from FilterEtherWithdraw and is used to iterate over the raw logs and unpacked data for EtherWithdraw events raised by the KyberNetwork contract.
type KyberNetworkEtherWithdrawIterator struct {
	Event *KyberNetworkEtherWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberNetworkEtherWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkEtherWithdraw)
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
		it.Event = new(KyberNetworkEtherWithdraw)
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
func (it *KyberNetworkEtherWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkEtherWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkEtherWithdraw represents a EtherWithdraw event raised by the KyberNetwork contract.
type KyberNetworkEtherWithdraw struct {
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdraw is a free log retrieval operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetwork *KyberNetworkFilterer) FilterEtherWithdraw(opts *bind.FilterOpts) (*KyberNetworkEtherWithdrawIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkEtherWithdrawIterator{contract: _KyberNetwork.contract, event: "EtherWithdraw", logs: logs, sub: sub}, nil
}

// WatchEtherWithdraw is a free log subscription operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetwork *KyberNetworkFilterer) WatchEtherWithdraw(opts *bind.WatchOpts, sink chan<- *KyberNetworkEtherWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "EtherWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkEtherWithdraw)
				if err := _KyberNetwork.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
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

// ParseEtherWithdraw is a log parse operation binding the contract event 0xec47e7ed86c86774d1a72c19f35c639911393fe7c1a34031fdbd260890da90de.
//
// Solidity: event EtherWithdraw(uint256 amount, address sendTo)
func (_KyberNetwork *KyberNetworkFilterer) ParseEtherWithdraw(log types.Log) (*KyberNetworkEtherWithdraw, error) {
	event := new(KyberNetworkEtherWithdraw)
	if err := _KyberNetwork.contract.UnpackLog(event, "EtherWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkExecuteTradeIterator is returned from FilterExecuteTrade and is used to iterate over the raw logs and unpacked data for ExecuteTrade events raised by the KyberNetwork contract.
type KyberNetworkExecuteTradeIterator struct {
	Event *KyberNetworkExecuteTrade // Event containing the contract specifics and raw log

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
func (it *KyberNetworkExecuteTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkExecuteTrade)
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
		it.Event = new(KyberNetworkExecuteTrade)
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
func (it *KyberNetworkExecuteTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkExecuteTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkExecuteTrade represents a ExecuteTrade event raised by the KyberNetwork contract.
type KyberNetworkExecuteTrade struct {
	Trader           common.Address
	Src              common.Address
	Dest             common.Address
	ActualSrcAmount  *big.Int
	ActualDestAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecuteTrade is a free log retrieval operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetwork *KyberNetworkFilterer) FilterExecuteTrade(opts *bind.FilterOpts, trader []common.Address) (*KyberNetworkExecuteTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return &KyberNetworkExecuteTradeIterator{contract: _KyberNetwork.contract, event: "ExecuteTrade", logs: logs, sub: sub}, nil
}

// WatchExecuteTrade is a free log subscription operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetwork *KyberNetworkFilterer) WatchExecuteTrade(opts *bind.WatchOpts, sink chan<- *KyberNetworkExecuteTrade, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "ExecuteTrade", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkExecuteTrade)
				if err := _KyberNetwork.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
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

// ParseExecuteTrade is a log parse operation binding the contract event 0x1849bd6a030a1bca28b83437fd3de96f3d27a5d172fa7e9c78e7b61468928a39.
//
// Solidity: event ExecuteTrade(address indexed trader, address src, address dest, uint256 actualSrcAmount, uint256 actualDestAmount)
func (_KyberNetwork *KyberNetworkFilterer) ParseExecuteTrade(log types.Log) (*KyberNetworkExecuteTrade, error) {
	event := new(KyberNetworkExecuteTrade)
	if err := _KyberNetwork.contract.UnpackLog(event, "ExecuteTrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkKyberNetworkSetIterator is returned from FilterKyberNetworkSet and is used to iterate over the raw logs and unpacked data for KyberNetworkSet events raised by the KyberNetwork contract.
type KyberNetworkKyberNetworkSetIterator struct {
	Event *KyberNetworkKyberNetworkSet // Event containing the contract specifics and raw log

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
func (it *KyberNetworkKyberNetworkSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkKyberNetworkSet)
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
		it.Event = new(KyberNetworkKyberNetworkSet)
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
func (it *KyberNetworkKyberNetworkSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkKyberNetworkSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkKyberNetworkSet represents a KyberNetworkSet event raised by the KyberNetwork contract.
type KyberNetworkKyberNetworkSet struct {
	NewNetworkContract common.Address
	OldNetworkContract common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterKyberNetworkSet is a free log retrieval operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetwork *KyberNetworkFilterer) FilterKyberNetworkSet(opts *bind.FilterOpts) (*KyberNetworkKyberNetworkSetIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkKyberNetworkSetIterator{contract: _KyberNetwork.contract, event: "KyberNetworkSet", logs: logs, sub: sub}, nil
}

// WatchKyberNetworkSet is a free log subscription operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetwork *KyberNetworkFilterer) WatchKyberNetworkSet(opts *bind.WatchOpts, sink chan<- *KyberNetworkKyberNetworkSet) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "KyberNetworkSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkKyberNetworkSet)
				if err := _KyberNetwork.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
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

// ParseKyberNetworkSet is a log parse operation binding the contract event 0x8936e1f096bf0a8c9df862b3d1d5b82774cad78116200175f00b5b7ba3010b02.
//
// Solidity: event KyberNetworkSet(address newNetworkContract, address oldNetworkContract)
func (_KyberNetwork *KyberNetworkFilterer) ParseKyberNetworkSet(log types.Log) (*KyberNetworkKyberNetworkSet, error) {
	event := new(KyberNetworkKyberNetworkSet)
	if err := _KyberNetwork.contract.UnpackLog(event, "KyberNetworkSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the KyberNetwork contract.
type KyberNetworkOperatorAddedIterator struct {
	Event *KyberNetworkOperatorAdded // Event containing the contract specifics and raw log

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
func (it *KyberNetworkOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkOperatorAdded)
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
		it.Event = new(KyberNetworkOperatorAdded)
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
func (it *KyberNetworkOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkOperatorAdded represents a OperatorAdded event raised by the KyberNetwork contract.
type KyberNetworkOperatorAdded struct {
	NewOperator common.Address
	IsAdd       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetwork *KyberNetworkFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*KyberNetworkOperatorAddedIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkOperatorAddedIterator{contract: _KyberNetwork.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetwork *KyberNetworkFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *KyberNetworkOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkOperatorAdded)
				if err := _KyberNetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x091a7a4b85135fdd7e8dbc18b12fabe5cc191ea867aa3c2e1a24a102af61d58b.
//
// Solidity: event OperatorAdded(address newOperator, bool isAdd)
func (_KyberNetwork *KyberNetworkFilterer) ParseOperatorAdded(log types.Log) (*KyberNetworkOperatorAdded, error) {
	event := new(KyberNetworkOperatorAdded)
	if err := _KyberNetwork.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the KyberNetwork contract.
type KyberNetworkTokenWithdrawIterator struct {
	Event *KyberNetworkTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *KyberNetworkTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkTokenWithdraw)
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
		it.Event = new(KyberNetworkTokenWithdraw)
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
func (it *KyberNetworkTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkTokenWithdraw represents a TokenWithdraw event raised by the KyberNetwork contract.
type KyberNetworkTokenWithdraw struct {
	Token  common.Address
	Amount *big.Int
	SendTo common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetwork *KyberNetworkFilterer) FilterTokenWithdraw(opts *bind.FilterOpts) (*KyberNetworkTokenWithdrawIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkTokenWithdrawIterator{contract: _KyberNetwork.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetwork *KyberNetworkFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *KyberNetworkTokenWithdraw) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "TokenWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkTokenWithdraw)
				if err := _KyberNetwork.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x72cb8a894ddb372ceec3d2a7648d86f17d5a15caae0e986c53109b8a9a9385e6.
//
// Solidity: event TokenWithdraw(address token, uint256 amount, address sendTo)
func (_KyberNetwork *KyberNetworkFilterer) ParseTokenWithdraw(log types.Log) (*KyberNetworkTokenWithdraw, error) {
	event := new(KyberNetworkTokenWithdraw)
	if err := _KyberNetwork.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KyberNetworkTransferAdminPendingIterator is returned from FilterTransferAdminPending and is used to iterate over the raw logs and unpacked data for TransferAdminPending events raised by the KyberNetwork contract.
type KyberNetworkTransferAdminPendingIterator struct {
	Event *KyberNetworkTransferAdminPending // Event containing the contract specifics and raw log

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
func (it *KyberNetworkTransferAdminPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberNetworkTransferAdminPending)
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
		it.Event = new(KyberNetworkTransferAdminPending)
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
func (it *KyberNetworkTransferAdminPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberNetworkTransferAdminPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberNetworkTransferAdminPending represents a TransferAdminPending event raised by the KyberNetwork contract.
type KyberNetworkTransferAdminPending struct {
	PendingAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminPending is a free log retrieval operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetwork *KyberNetworkFilterer) FilterTransferAdminPending(opts *bind.FilterOpts) (*KyberNetworkTransferAdminPendingIterator, error) {

	logs, sub, err := _KyberNetwork.contract.FilterLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return &KyberNetworkTransferAdminPendingIterator{contract: _KyberNetwork.contract, event: "TransferAdminPending", logs: logs, sub: sub}, nil
}

// WatchTransferAdminPending is a free log subscription operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetwork *KyberNetworkFilterer) WatchTransferAdminPending(opts *bind.WatchOpts, sink chan<- *KyberNetworkTransferAdminPending) (event.Subscription, error) {

	logs, sub, err := _KyberNetwork.contract.WatchLogs(opts, "TransferAdminPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberNetworkTransferAdminPending)
				if err := _KyberNetwork.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
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

// ParseTransferAdminPending is a log parse operation binding the contract event 0x3b81caf78fa51ecbc8acb482fd7012a277b428d9b80f9d156e8a54107496cc40.
//
// Solidity: event TransferAdminPending(address pendingAdmin)
func (_KyberNetwork *KyberNetworkFilterer) ParseTransferAdminPending(log types.Log) (*KyberNetworkTransferAdminPending, error) {
	event := new(KyberNetworkTransferAdminPending)
	if err := _KyberNetwork.contract.UnpackLog(event, "TransferAdminPending", log); err != nil {
		return nil, err
	}
	return event, nil
}
