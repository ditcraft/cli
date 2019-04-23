// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KNWToken

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

// KNWTokenABI is the input ABI used to generate the binding from.
const KNWTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"freeBalanceOfLabel\",\"outputs\":[{\"name\":\"freeBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"addVotingContract\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_labelID\",\"type\":\"uint256\"}],\"name\":\"labelOfAddress\",\"outputs\":[{\"name\":\"label\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"balanceOfLabel\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"totalLabelSupply\",\"outputs\":[{\"name\":\"totalSupplyOfLabel\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"labelCountOfAddress\",\"outputs\":[{\"name\":\"labelCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votingContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// KNWTokenBin is the compiled bytecode used for deploying new contracts.
const KNWTokenBin = `0x608060405234801561001057600080fd5b506112f7806100206000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306970f1c81146100df57806306fdde031461011e57806318160ddd146101a857806320c7ad57146101bd578063313ce567146101f257806358e546cf1461021d578063735997981461024157806395d89b4114610271578063b88c0f9814610286578063ba7aef43146102b3578063c1a63f3c146102e3578063c45b71de14610303578063d950df3414610333578063e7017fc414610363578063fa06792b14610384575b600080fd5b3480156100eb57600080fd5b5061010c60048035600160a060020a031690602480359081019101356103a5565b60408051918252519081900360200190f35b34801561012a57600080fd5b50610133610445565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016d578181015183820152602001610155565b50505050905090810190601f16801561019a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101b457600080fd5b5061010c61047c565b3480156101c957600080fd5b506101de600160a060020a0360043516610483565b604080519115158252519081900360200190f35b3480156101fe57600080fd5b50610207610534565b6040805160ff9092168252519081900360200190f35b34801561022957600080fd5b50610133600160a060020a0360043516602435610539565b34801561024d57600080fd5b506101de60048035600160a060020a031690602480359081019101356044356105eb565b34801561027d57600080fd5b506101336107d7565b34801561029257600080fd5b5061010c60048035600160a060020a0316906024803590810191013561080e565b3480156102bf57600080fd5b506101de60048035600160a060020a03169060248035908101910135604435610854565b3480156102ef57600080fd5b5061010c6004803560248101910135610bbf565b34801561030f57600080fd5b506101de60048035600160a060020a03169060248035908101910135604435610bee565b34801561033f57600080fd5b506101de60048035600160a060020a03169060248035908101910135604435610f19565b34801561036f57600080fd5b5061010c600160a060020a03600435166110f1565b34801561039057600080fd5b506101de600160a060020a036004351661110c565b600160a060020a038316600090815260016020526040808220905161043d9190859085908083838082843782019150509250505090815260200160405180910390205460008087600160a060020a0316600160a060020a03168152602001908152602001600020858560405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61112116565b949350505050565b60408051808201909152600f81527f4b6e6f776c6564676520546f6b656e0000000000000000000000000000000000602082015281565b6005545b90565b6000600160a060020a038216151561050b576040805160e560020a62461bcd02815260206004820152602860248201527f566f74696e6720636f6e7472616374732720616464726573732063616e27742060448201527f626520656d707479000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b50600160a060020a03166000908152600760205260409020805460ff1916600190811790915590565b601281565b600160a060020a0382166000908152600260208181526040808420858552825292839020805484516001821615610100026000190190911693909304601f810183900483028401830190945283835260609390918301828280156105de5780601f106105b3576101008083540402835291602001916105de565b820191906000526020600020905b8154815290600101906020018083116105c157829003601f168201915b5050505050905092915050565b336000818152600760205260408120549091829160ff16151561060d57600080fd5b6106ae6001600089600160a060020a0316600160a060020a031681526020019081526020016000208787604051808383808284378201915050925050509081526020016040518091039020546000808a600160a060020a0316600160a060020a03168152602001908152602001600020888860405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61112116565b91508382101561072e576040805160e560020a62461bcd02815260206004820152602560248201527f43616e2774206c6f636b206d6f726520746f6b656e73207468616e206176616960448201527f6c61626c65000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61078884600160008a600160a060020a0316600160a060020a03168152602001908152602001600020888860405180838380828437909101948552505060405192839003602001909220549291505063ffffffff6111a916565b600160a060020a03881660009081526001602052604090819020905188908890808383808284379091019485525050604051928390036020019092209290925550600198975050505050505050565b60408051808201909152600381527f4b4e570000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a03831660009081526020819052604080822090518490849080838380828437909101948552505060405192839003602001909220549695505050505050565b3360008181526007602052604081205490919060ff16151561087557600080fd5b600160a060020a03861615156108d5576040805160e560020a62461bcd02815260206004820152601660248201527f416464726573732063616e277420626520656d70747900000000000000000000604482015290519081900360640190fd5b6000841161092d576040805160e560020a62461bcd02815260206004820152601e60248201527f4b6e6f776c656467652d4c6162656c2063616e277420626520656d7074790000604482015290519081900360640190fd5b600554610940908463ffffffff6111a916565b60058190555061097d836006878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff6111a916565b60068686604051808383808284378201915050925050509081526020016040518091039020819055506109ff8360008089600160a060020a0316600160a060020a03168152602001908152602001600020878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff6111a916565b60008088600160a060020a0316600160a060020a03168152602001908152602001600020868660405180838380828437909101948552505060408051938490036020908101852095909555600160a060020a038b166000908152600490955290932092508791879150808383808284379091019485525050604051928390036020019092205460ff1615159150610b51905057600160a060020a038616600090815260036020526040902054610abc90600163ffffffff6111a916565b600160a060020a03871660009081526003602090815260408083208490556002825280832093835292905220610af3908686611233565b5060016004600088600160a060020a0316600160a060020a03168152602001908152602001600020868660405180838380828437909101948552505060405192839003602001909220805493151560ff199094169390931790925550505b85600160a060020a03167fec4de1eef14af3ae5d77facf1ed7a9d3d50f6285573ee0ec155fc11217fc3442868686604051808060200183815260200182810382528585828181526020019250808284376040519201829003965090945050505050a250600195945050505050565b600060068383604051808383808284379091019485525050604051928390036020019092205495945050505050565b3360008181526007602052604081205490919060ff161515610c0f57600080fd5b600160a060020a0386161515610c6f576040805160e560020a62461bcd02815260206004820152601660248201527f416464726573732063616e277420626520656d70747900000000000000000000604482015290519081900360640190fd5b60008411610cc7576040805160e560020a62461bcd02815260206004820152601e60248201527f4b6e6f776c656467652d4c6162656c2063616e277420626520656d7074790000604482015290519081900360640190fd5b8260008088600160a060020a0316600160a060020a0316815260200190815260200160002086866040518083838082843782019150509250505090815260200160405180910390205410151515610d8e576040805160e560020a62461bcd02815260206004820152602a60248201527f43616e2774206275726e206d6f7265204b4e57207468616e207468652061646460448201527f7265737320686f6c647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b600554610da1908463ffffffff61112116565b600581905550610dde836006878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61112116565b6006868660405180838380828437820191505092505050908152602001604051809103902081905550610e608360008089600160a060020a0316600160a060020a03168152602001908152602001600020878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61112116565b60008088600160a060020a0316600160a060020a0316815260200190815260200160002086866040518083838082843782019150509250505090815260200160405180910390208190555085600160a060020a03167ffdf096248d2b7b0aef506231c043107c21faacc26193881b3f0cdc8b5479692a868686604051808060200183815260200182810382528585828181526020019250808284376040519201829003965090945050505050a250600195945050505050565b3360008181526007602052604081205490919060ff161515610f3a57600080fd5b60008087600160a060020a0316600160a060020a031681526020019081526020016000208585604051808383808284378201915050925050509081526020016040518091039020546001600088600160a060020a0316600160a060020a0316815260200190815260200160002086866040518083838082843782019150509250505090815260200160405180910390205411151515611049576040805160e560020a62461bcd02815260206004820152602660248201527f43616e74206c6f636b206d6f7265204b4e57207468616e20616e20616464726560448201527f7373206861730000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6110a3836001600089600160a060020a0316600160a060020a03168152602001908152602001600020878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61112116565b600160a060020a038716600090815260016020526040908190209051879087908083838082843790910194855250506040519283900360200190922092909255506001979650505050505050565b600160a060020a031660009081526003602052604090205490565b60076020526000908152604090205460ff1681565b600080838311156111a2576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050900390565b60008282018381101561122c576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112745782800160ff198235161785556112a1565b828001600101855582156112a1579182015b828111156112a1578235825591602001919060010190611286565b506112ad9291506112b1565b5090565b61048091905b808211156112ad57600081556001016112b75600a165627a7a723058202606207f26af739b5c06a6330737b278e218ae421e2bf77fb1be33979ad3daba0029`

// DeployKNWToken deploys a new Ethereum contract, binding an instance of KNWToken to it.
func DeployKNWToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KNWToken, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWToken{KNWTokenCaller: KNWTokenCaller{contract: contract}, KNWTokenTransactor: KNWTokenTransactor{contract: contract}, KNWTokenFilterer: KNWTokenFilterer{contract: contract}}, nil
}

// KNWToken is an auto generated Go binding around an Ethereum contract.
type KNWToken struct {
	KNWTokenCaller     // Read-only binding to the contract
	KNWTokenTransactor // Write-only binding to the contract
	KNWTokenFilterer   // Log filterer for contract events
}

// KNWTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWTokenSession struct {
	Contract     *KNWToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNWTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWTokenCallerSession struct {
	Contract *KNWTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KNWTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWTokenTransactorSession struct {
	Contract     *KNWTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KNWTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWTokenRaw struct {
	Contract *KNWToken // Generic contract binding to access the raw methods on
}

// KNWTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWTokenCallerRaw struct {
	Contract *KNWTokenCaller // Generic read-only contract binding to access the raw methods on
}

// KNWTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWTokenTransactorRaw struct {
	Contract *KNWTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWToken creates a new instance of KNWToken, bound to a specific deployed contract.
func NewKNWToken(address common.Address, backend bind.ContractBackend) (*KNWToken, error) {
	contract, err := bindKNWToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWToken{KNWTokenCaller: KNWTokenCaller{contract: contract}, KNWTokenTransactor: KNWTokenTransactor{contract: contract}, KNWTokenFilterer: KNWTokenFilterer{contract: contract}}, nil
}

// NewKNWTokenCaller creates a new read-only instance of KNWToken, bound to a specific deployed contract.
func NewKNWTokenCaller(address common.Address, caller bind.ContractCaller) (*KNWTokenCaller, error) {
	contract, err := bindKNWToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenCaller{contract: contract}, nil
}

// NewKNWTokenTransactor creates a new write-only instance of KNWToken, bound to a specific deployed contract.
func NewKNWTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWTokenTransactor, error) {
	contract, err := bindKNWToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenTransactor{contract: contract}, nil
}

// NewKNWTokenFilterer creates a new log filterer instance of KNWToken, bound to a specific deployed contract.
func NewKNWTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWTokenFilterer, error) {
	contract, err := bindKNWToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWTokenFilterer{contract: contract}, nil
}

// bindKNWToken binds a generic wrapper to an already deployed contract.
func bindKNWToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWToken *KNWTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWToken.Contract.KNWTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWToken *KNWTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWToken.Contract.KNWTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWToken *KNWTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWToken.Contract.KNWTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWToken *KNWTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWToken *KNWTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWToken *KNWTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(address _address, string _label) constant returns(uint256 balance)
func (_KNWToken *KNWTokenCaller) BalanceOfLabel(opts *bind.CallOpts, _address common.Address, _label string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "balanceOfLabel", _address, _label)
	return *ret0, err
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(address _address, string _label) constant returns(uint256 balance)
func (_KNWToken *KNWTokenSession) BalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.BalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(address _address, string _label) constant returns(uint256 balance)
func (_KNWToken *KNWTokenCallerSession) BalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.BalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_KNWToken *KNWTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_KNWToken *KNWTokenSession) Decimals() (uint8, error) {
	return _KNWToken.Contract.Decimals(&_KNWToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_KNWToken *KNWTokenCallerSession) Decimals() (uint8, error) {
	return _KNWToken.Contract.Decimals(&_KNWToken.CallOpts)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(address _address, string _label) constant returns(uint256 freeBalance)
func (_KNWToken *KNWTokenCaller) FreeBalanceOfLabel(opts *bind.CallOpts, _address common.Address, _label string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "freeBalanceOfLabel", _address, _label)
	return *ret0, err
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(address _address, string _label) constant returns(uint256 freeBalance)
func (_KNWToken *KNWTokenSession) FreeBalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.FreeBalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(address _address, string _label) constant returns(uint256 freeBalance)
func (_KNWToken *KNWTokenCallerSession) FreeBalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.FreeBalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// LabelCountOfAddress is a free data retrieval call binding the contract method 0xe7017fc4.
//
// Solidity: function labelCountOfAddress(address _address) constant returns(uint256 labelCount)
func (_KNWToken *KNWTokenCaller) LabelCountOfAddress(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "labelCountOfAddress", _address)
	return *ret0, err
}

// LabelCountOfAddress is a free data retrieval call binding the contract method 0xe7017fc4.
//
// Solidity: function labelCountOfAddress(address _address) constant returns(uint256 labelCount)
func (_KNWToken *KNWTokenSession) LabelCountOfAddress(_address common.Address) (*big.Int, error) {
	return _KNWToken.Contract.LabelCountOfAddress(&_KNWToken.CallOpts, _address)
}

// LabelCountOfAddress is a free data retrieval call binding the contract method 0xe7017fc4.
//
// Solidity: function labelCountOfAddress(address _address) constant returns(uint256 labelCount)
func (_KNWToken *KNWTokenCallerSession) LabelCountOfAddress(_address common.Address) (*big.Int, error) {
	return _KNWToken.Contract.LabelCountOfAddress(&_KNWToken.CallOpts, _address)
}

// LabelOfAddress is a free data retrieval call binding the contract method 0x58e546cf.
//
// Solidity: function labelOfAddress(address _address, uint256 _labelID) constant returns(string label)
func (_KNWToken *KNWTokenCaller) LabelOfAddress(opts *bind.CallOpts, _address common.Address, _labelID *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "labelOfAddress", _address, _labelID)
	return *ret0, err
}

// LabelOfAddress is a free data retrieval call binding the contract method 0x58e546cf.
//
// Solidity: function labelOfAddress(address _address, uint256 _labelID) constant returns(string label)
func (_KNWToken *KNWTokenSession) LabelOfAddress(_address common.Address, _labelID *big.Int) (string, error) {
	return _KNWToken.Contract.LabelOfAddress(&_KNWToken.CallOpts, _address, _labelID)
}

// LabelOfAddress is a free data retrieval call binding the contract method 0x58e546cf.
//
// Solidity: function labelOfAddress(address _address, uint256 _labelID) constant returns(string label)
func (_KNWToken *KNWTokenCallerSession) LabelOfAddress(_address common.Address, _labelID *big.Int) (string, error) {
	return _KNWToken.Contract.LabelOfAddress(&_KNWToken.CallOpts, _address, _labelID)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KNWToken *KNWTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KNWToken *KNWTokenSession) Name() (string, error) {
	return _KNWToken.Contract.Name(&_KNWToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_KNWToken *KNWTokenCallerSession) Name() (string, error) {
	return _KNWToken.Contract.Name(&_KNWToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KNWToken *KNWTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KNWToken *KNWTokenSession) Symbol() (string, error) {
	return _KNWToken.Contract.Symbol(&_KNWToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_KNWToken *KNWTokenCallerSession) Symbol() (string, error) {
	return _KNWToken.Contract.Symbol(&_KNWToken.CallOpts)
}

// TotalLabelSupply is a free data retrieval call binding the contract method 0xc1a63f3c.
//
// Solidity: function totalLabelSupply(string _label) constant returns(uint256 totalSupplyOfLabel)
func (_KNWToken *KNWTokenCaller) TotalLabelSupply(opts *bind.CallOpts, _label string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "totalLabelSupply", _label)
	return *ret0, err
}

// TotalLabelSupply is a free data retrieval call binding the contract method 0xc1a63f3c.
//
// Solidity: function totalLabelSupply(string _label) constant returns(uint256 totalSupplyOfLabel)
func (_KNWToken *KNWTokenSession) TotalLabelSupply(_label string) (*big.Int, error) {
	return _KNWToken.Contract.TotalLabelSupply(&_KNWToken.CallOpts, _label)
}

// TotalLabelSupply is a free data retrieval call binding the contract method 0xc1a63f3c.
//
// Solidity: function totalLabelSupply(string _label) constant returns(uint256 totalSupplyOfLabel)
func (_KNWToken *KNWTokenCallerSession) TotalLabelSupply(_label string) (*big.Int, error) {
	return _KNWToken.Contract.TotalLabelSupply(&_KNWToken.CallOpts, _label)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KNWToken *KNWTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KNWToken *KNWTokenSession) TotalSupply() (*big.Int, error) {
	return _KNWToken.Contract.TotalSupply(&_KNWToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_KNWToken *KNWTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _KNWToken.Contract.TotalSupply(&_KNWToken.CallOpts)
}

// VotingContracts is a free data retrieval call binding the contract method 0xfa06792b.
//
// Solidity: function votingContracts(address ) constant returns(bool)
func (_KNWToken *KNWTokenCaller) VotingContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "votingContracts", arg0)
	return *ret0, err
}

// VotingContracts is a free data retrieval call binding the contract method 0xfa06792b.
//
// Solidity: function votingContracts(address ) constant returns(bool)
func (_KNWToken *KNWTokenSession) VotingContracts(arg0 common.Address) (bool, error) {
	return _KNWToken.Contract.VotingContracts(&_KNWToken.CallOpts, arg0)
}

// VotingContracts is a free data retrieval call binding the contract method 0xfa06792b.
//
// Solidity: function votingContracts(address ) constant returns(bool)
func (_KNWToken *KNWTokenCallerSession) VotingContracts(arg0 common.Address) (bool, error) {
	return _KNWToken.Contract.VotingContracts(&_KNWToken.CallOpts, arg0)
}

// AddVotingContract is a paid mutator transaction binding the contract method 0x20c7ad57.
//
// Solidity: function addVotingContract(address _contractAddress) returns(bool success)
func (_KNWToken *KNWTokenTransactor) AddVotingContract(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "addVotingContract", _contractAddress)
}

// AddVotingContract is a paid mutator transaction binding the contract method 0x20c7ad57.
//
// Solidity: function addVotingContract(address _contractAddress) returns(bool success)
func (_KNWToken *KNWTokenSession) AddVotingContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.AddVotingContract(&_KNWToken.TransactOpts, _contractAddress)
}

// AddVotingContract is a paid mutator transaction binding the contract method 0x20c7ad57.
//
// Solidity: function addVotingContract(address _contractAddress) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) AddVotingContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.AddVotingContract(&_KNWToken.TransactOpts, _contractAddress)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) Burn(opts *bind.TransactOpts, _address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "burn", _address, _label, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) Burn(_address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Burn(&_KNWToken.TransactOpts, _address, _label, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xc45b71de.
//
// Solidity: function burn(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) Burn(_address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Burn(&_KNWToken.TransactOpts, _address, _label, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x73599798.
//
// Solidity: function lockTokens(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) LockTokens(opts *bind.TransactOpts, _address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "lockTokens", _address, _label, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x73599798.
//
// Solidity: function lockTokens(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) LockTokens(_address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.LockTokens(&_KNWToken.TransactOpts, _address, _label, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x73599798.
//
// Solidity: function lockTokens(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) LockTokens(_address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.LockTokens(&_KNWToken.TransactOpts, _address, _label, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) Mint(opts *bind.TransactOpts, _address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "mint", _address, _label, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) Mint(_address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Mint(&_KNWToken.TransactOpts, _address, _label, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xba7aef43.
//
// Solidity: function mint(address _address, string _label, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) Mint(_address common.Address, _label string, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Mint(&_KNWToken.TransactOpts, _address, _label, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(address _address, string _label, uint256 _numberOfTokens) returns(bool success)
func (_KNWToken *KNWTokenTransactor) UnlockTokens(opts *bind.TransactOpts, _address common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "unlockTokens", _address, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(address _address, string _label, uint256 _numberOfTokens) returns(bool success)
func (_KNWToken *KNWTokenSession) UnlockTokens(_address common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.UnlockTokens(&_KNWToken.TransactOpts, _address, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(address _address, string _label, uint256 _numberOfTokens) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) UnlockTokens(_address common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.UnlockTokens(&_KNWToken.TransactOpts, _address, _label, _numberOfTokens)
}

// KNWTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the KNWToken contract.
type KNWTokenBurnIterator struct {
	Event *KNWTokenBurn // Event containing the contract specifics and raw log

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
func (it *KNWTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNWTokenBurn)
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
		it.Event = new(KNWTokenBurn)
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
func (it *KNWTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNWTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNWTokenBurn represents a Burn event raised by the KNWToken contract.
type KNWTokenBurn struct {
	Who   common.Address
	Label string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xfdf096248d2b7b0aef506231c043107c21faacc26193881b3f0cdc8b5479692a.
//
// Solidity: event Burn(address indexed who, string label, uint256 value)
func (_KNWToken *KNWTokenFilterer) FilterBurn(opts *bind.FilterOpts, who []common.Address) (*KNWTokenBurnIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _KNWToken.contract.FilterLogs(opts, "Burn", whoRule)
	if err != nil {
		return nil, err
	}
	return &KNWTokenBurnIterator{contract: _KNWToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xfdf096248d2b7b0aef506231c043107c21faacc26193881b3f0cdc8b5479692a.
//
// Solidity: event Burn(address indexed who, string label, uint256 value)
func (_KNWToken *KNWTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *KNWTokenBurn, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _KNWToken.contract.WatchLogs(opts, "Burn", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNWTokenBurn)
				if err := _KNWToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// KNWTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the KNWToken contract.
type KNWTokenMintIterator struct {
	Event *KNWTokenMint // Event containing the contract specifics and raw log

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
func (it *KNWTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNWTokenMint)
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
		it.Event = new(KNWTokenMint)
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
func (it *KNWTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNWTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNWTokenMint represents a Mint event raised by the KNWToken contract.
type KNWTokenMint struct {
	Who   common.Address
	Label string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xec4de1eef14af3ae5d77facf1ed7a9d3d50f6285573ee0ec155fc11217fc3442.
//
// Solidity: event Mint(address indexed who, string label, uint256 value)
func (_KNWToken *KNWTokenFilterer) FilterMint(opts *bind.FilterOpts, who []common.Address) (*KNWTokenMintIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _KNWToken.contract.FilterLogs(opts, "Mint", whoRule)
	if err != nil {
		return nil, err
	}
	return &KNWTokenMintIterator{contract: _KNWToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xec4de1eef14af3ae5d77facf1ed7a9d3d50f6285573ee0ec155fc11217fc3442.
//
// Solidity: event Mint(address indexed who, string label, uint256 value)
func (_KNWToken *KNWTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *KNWTokenMint, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _KNWToken.contract.WatchLogs(opts, "Mint", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNWTokenMint)
				if err := _KNWToken.contract.UnpackLog(event, "Mint", log); err != nil {
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
