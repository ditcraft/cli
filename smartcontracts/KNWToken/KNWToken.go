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
const KNWTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"replaceManager\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"totalIDSupply\",\"outputs\":[{\"name\":\"totalSupplyOfID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOfID\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"authorizeAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newLabel\",\"type\":\"string\"}],\"name\":\"addNewLabel\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAddress\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"freeBalanceOfID\",\"outputs\":[{\"name\":\"freeBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"labelOfID\",\"outputs\":[{\"name\":\"label\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountOfIDs\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"}],\"name\":\"NewLabel\",\"type\":\"event\"}]"

// KNWTokenFuncSigs maps the 4-byte function signature to its string representation.
var KNWTokenFuncSigs = map[string]string{
	"4ac99804": "addNewLabel(string)",
	"ffe03eba": "amountOfIDs()",
	"4a5db3b5": "authorizeAddress(address)",
	"f19e207e": "authorizedAddresses(address)",
	"3938400b": "balanceOfID(address,uint256)",
	"f5298aca": "burn(address,uint256,uint256)",
	"313ce567": "decimals()",
	"b5c2cdba": "freeBalanceOfID(address,uint256)",
	"cce38484": "labelOfID(uint256)",
	"a25983e5": "lockTokens(address,uint256,uint256)",
	"481c6a75": "manager()",
	"156e29f6": "mint(address,uint256,uint256)",
	"06fdde03": "name()",
	"4ba79dfe": "removeAddress(address)",
	"23447982": "replaceManager(address)",
	"95d89b41": "symbol()",
	"2d243d1e": "totalIDSupply(uint256)",
	"18160ddd": "totalSupply()",
	"cd3877df": "unlockTokens(address,uint256,uint256)",
}

// KNWTokenBin is the compiled bytecode used for deploying new contracts.
var KNWTokenBin = "0x608060405234801561001057600080fd5b50600580546001600160a01b03191633179055610fec806100326000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80634ac99804116100ad578063cce3848411610071578063cce38484146103d6578063cd3877df146103f3578063f19e207e14610425578063f5298aca1461044b578063ffe03eba1461047d57610121565b80634ac99804146102da5780634ba79dfe1461034a57806395d89b4114610370578063a25983e514610378578063b5c2cdba146103aa57610121565b80632d243d1e116100f45780632d243d1e14610229578063313ce567146102465780633938400b14610264578063481c6a75146102905780634a5db3b5146102b457610121565b806306fdde0314610126578063156e29f6146101a357806318160ddd146101e95780632344798214610203575b600080fd5b61012e610485565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610168578181015183820152602001610150565b50505050905090810190601f1680156101955780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101d5600480360360608110156101b957600080fd5b506001600160a01b0381351690602081013590604001356104b0565b604080519115158252519081900360200190f35b6101f1610641565b60408051918252519081900360200190f35b6101d56004803603602081101561021957600080fd5b50356001600160a01b0316610648565b6101f16004803603602081101561023f57600080fd5b503561069e565b61024e6106b0565b6040805160ff9092168252519081900360200190f35b6101f16004803603604081101561027a57600080fd5b506001600160a01b0381351690602001356106b5565b6102986106db565b604080516001600160a01b039092168252519081900360200190f35b6101d5600480360360208110156102ca57600080fd5b50356001600160a01b03166106ea565b6101d5600480360360208110156102f057600080fd5b81019060208101813564010000000081111561030b57600080fd5b82018360208201111561031d57600080fd5b8035906020019184600183028401116401000000008311171561033f57600080fd5b509092509050610775565b6101d56004803603602081101561036057600080fd5b50356001600160a01b0316610844565b61012e6108cb565b6101d56004803603606081101561038e57600080fd5b506001600160a01b0381351690602081013590604001356108ea565b6101f1600480360360408110156103c057600080fd5b506001600160a01b0381351690602001356109f9565b61012e600480360360208110156103ec57600080fd5b5035610a48565b6101d56004803603606081101561040957600080fd5b506001600160a01b038135169060208101359060400135610af1565b6101d56004803603602081101561043b57600080fd5b50356001600160a01b0316610be8565b6101d56004803603606081101561046157600080fd5b506001600160a01b038135169060208101359060400135610bfd565b6101f1610d8e565b6040518060400160405280600f81526020016e25b737bbb632b233b2902a37b5b2b760891b81525081565b3360008181526006602052604081205490919060ff166104cf57600080fd5b6001600160a01b038516610523576040805162461bcd60e51b8152602060048201526016602482015275416464726573732063616e277420626520656d70747960501b604482015290519081900360640190fd5b60045484106105635760405162461bcd60e51b815260040180806020018281038252602a815260200180610f69602a913960400191505060405180910390fd5b600254610576908463ffffffff610d9416565b600255600084815260036020526040902054610598908463ffffffff610d9416565b6000858152600360209081526040808320939093556001600160a01b0388168252818152828220878352905220546105d6908463ffffffff610d9416565b6001600160a01b038616600081815260208181526040808320898452825291829020939093558051878152928301869052805191927f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f929081900390910190a2506001949350505050565b6002545b90565b60055460009033906001600160a01b0316811461066457600080fd5b6001600160a01b03831661067757600080fd5b600580546001600160a01b0385166001600160a01b03199091161790556001915050919050565b60009081526003602052604090205490565b601281565b6001600160a01b0391909116600090815260208181526040808320938352929052205490565b6005546001600160a01b031681565b60055460009033906001600160a01b0316811461070657600080fd5b6001600160a01b03831661074b5760405162461bcd60e51b815260040180806020018281038252602c815260200180610f13602c913960400191505060405180910390fd5b50506001600160a01b03166000908152600660205260409020805460ff1916600190811790915590565b60055460009033906001600160a01b0316811461079157600080fd5b60048054600181018083556000929092526107cf907f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018686610e1f565b50506001600480549050037ff7e3164ba69dd943954519f6bce38af5edcad970e8160eace0ff12489b95ea3b858560405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a25060019392505050565b60055460009033906001600160a01b0316811461086057600080fd5b6001600160a01b0383166108a55760405162461bcd60e51b815260040180806020018281038252602c815260200180610f13602c913960400191505060405180910390fd5b50506001600160a01b03166000908152600660205260409020805460ff19169055600190565b604051806040016040528060038152602001624b4e5760e81b81525081565b3360008181526006602052604081205490919060ff1661090957600080fd5b6001600160a01b03851660008181526001602090815260408083208884528252808320549383528282528083208884529091528120549091610951919063ffffffff610dd816565b9050838110156109925760405162461bcd60e51b8152600401808060200182810382526025815260200180610f936025913960400191505060405180910390fd5b6001600160a01b03861660009081526001602090815260408083208884529091529020546109c6908563ffffffff610d9416565b6001600160a01b03871660009081526001602081815260408084208a855290915290912091909155925050509392505050565b6001600160a01b03821660008181526001602090815260408083208584528252808320549383528282528083208584529091528120549091610a41919063ffffffff610dd816565b9392505050565b606060048281548110610a5757fe5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610ae55780601f10610aba57610100808354040283529160200191610ae5565b820191906000526020600020905b815481529060010190602001808311610ac857829003601f168201915b50505050509050919050565b3360008181526006602052604081205490919060ff16610b1057600080fd5b6001600160a01b038516600081815260208181526040808320888452825280832054938352600182528083208884529091529020541115610b825760405162461bcd60e51b8152600401808060200182810382526026815260200180610eed6026913960400191505060405180910390fd5b6001600160a01b0385166000908152600160209081526040808320878452909152902054610bb6908463ffffffff610dd816565b6001600160a01b0386166000908152600160208181526040808420898552909152909120919091559150509392505050565b60066020526000908152604090205460ff1681565b3360008181526006602052604081205490919060ff16610c1c57600080fd5b6001600160a01b038516610c70576040805162461bcd60e51b8152602060048201526016602482015275416464726573732063616e277420626520656d70747960501b604482015290519081900360640190fd5b6004548410610cb05760405162461bcd60e51b815260040180806020018281038252602a815260200180610f69602a913960400191505060405180910390fd5b600254610cc3908463ffffffff610dd816565b600255600084815260036020526040902054610ce5908463ffffffff610dd816565b6000858152600360209081526040808320939093556001600160a01b038816825281815282822087835290522054610d23908463ffffffff610dd816565b6001600160a01b038616600081815260208181526040808320898452825291829020939093558051878152928301869052805191927f49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a929081900390910190a2506001949350505050565b60045490565b600082820183811015610a415760405162461bcd60e51b815260040180806020018281038252602a815260200180610f3f602a913960400191505060405180910390fd5b600082821115610e195760405162461bcd60e51b8152600401808060200182810382526035815260200180610eb86035913960400191505060405180910390fd5b50900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e605782800160ff19823516178555610e8d565b82800160010185558215610e8d579182015b82811115610e8d578235825591602001919060010190610e72565b50610e99929150610e9d565b5090565b61064591905b80821115610e995760008155600101610ea356fe43616e27742073756274726163742061206e756d6265722066726f6d206120736d616c6c6572206f6e6520776974682075696e747343616e74206c6f636b206d6f7265204b4e57207468616e20616e206164647265737320686173417574686f72697a656420636f6e7472616374732720616464726573732063616e277420626520656d707479526573756c742068617320746f20626520626967676572207468616e20626f74682073756d6d616e64734944206e6565647320746f2062652077697468696e2072616e6765206f6620616c6c6f7765642049447343616e2774206c6f636b206d6f726520746f6b656e73207468616e20617661696c61626c65a265627a7a72305820f0d82f23a89471622e930bbcb6fd02a2571c01d8fe213ac6411019eca596377664736f6c634300050a0032"

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

// AmountOfIDs is a free data retrieval call binding the contract method 0xffe03eba.
//
// Solidity: function amountOfIDs() constant returns(uint256 amount)
func (_KNWToken *KNWTokenCaller) AmountOfIDs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "amountOfIDs")
	return *ret0, err
}

// AmountOfIDs is a free data retrieval call binding the contract method 0xffe03eba.
//
// Solidity: function amountOfIDs() constant returns(uint256 amount)
func (_KNWToken *KNWTokenSession) AmountOfIDs() (*big.Int, error) {
	return _KNWToken.Contract.AmountOfIDs(&_KNWToken.CallOpts)
}

// AmountOfIDs is a free data retrieval call binding the contract method 0xffe03eba.
//
// Solidity: function amountOfIDs() constant returns(uint256 amount)
func (_KNWToken *KNWTokenCallerSession) AmountOfIDs() (*big.Int, error) {
	return _KNWToken.Contract.AmountOfIDs(&_KNWToken.CallOpts)
}

// AuthorizedAddresses is a free data retrieval call binding the contract method 0xf19e207e.
//
// Solidity: function authorizedAddresses(address ) constant returns(bool)
func (_KNWToken *KNWTokenCaller) AuthorizedAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "authorizedAddresses", arg0)
	return *ret0, err
}

// AuthorizedAddresses is a free data retrieval call binding the contract method 0xf19e207e.
//
// Solidity: function authorizedAddresses(address ) constant returns(bool)
func (_KNWToken *KNWTokenSession) AuthorizedAddresses(arg0 common.Address) (bool, error) {
	return _KNWToken.Contract.AuthorizedAddresses(&_KNWToken.CallOpts, arg0)
}

// AuthorizedAddresses is a free data retrieval call binding the contract method 0xf19e207e.
//
// Solidity: function authorizedAddresses(address ) constant returns(bool)
func (_KNWToken *KNWTokenCallerSession) AuthorizedAddresses(arg0 common.Address) (bool, error) {
	return _KNWToken.Contract.AuthorizedAddresses(&_KNWToken.CallOpts, arg0)
}

// BalanceOfID is a free data retrieval call binding the contract method 0x3938400b.
//
// Solidity: function balanceOfID(address _address, uint256 _id) constant returns(uint256 balance)
func (_KNWToken *KNWTokenCaller) BalanceOfID(opts *bind.CallOpts, _address common.Address, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "balanceOfID", _address, _id)
	return *ret0, err
}

// BalanceOfID is a free data retrieval call binding the contract method 0x3938400b.
//
// Solidity: function balanceOfID(address _address, uint256 _id) constant returns(uint256 balance)
func (_KNWToken *KNWTokenSession) BalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWToken.Contract.BalanceOfID(&_KNWToken.CallOpts, _address, _id)
}

// BalanceOfID is a free data retrieval call binding the contract method 0x3938400b.
//
// Solidity: function balanceOfID(address _address, uint256 _id) constant returns(uint256 balance)
func (_KNWToken *KNWTokenCallerSession) BalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWToken.Contract.BalanceOfID(&_KNWToken.CallOpts, _address, _id)
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

// FreeBalanceOfID is a free data retrieval call binding the contract method 0xb5c2cdba.
//
// Solidity: function freeBalanceOfID(address _address, uint256 _id) constant returns(uint256 freeBalance)
func (_KNWToken *KNWTokenCaller) FreeBalanceOfID(opts *bind.CallOpts, _address common.Address, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "freeBalanceOfID", _address, _id)
	return *ret0, err
}

// FreeBalanceOfID is a free data retrieval call binding the contract method 0xb5c2cdba.
//
// Solidity: function freeBalanceOfID(address _address, uint256 _id) constant returns(uint256 freeBalance)
func (_KNWToken *KNWTokenSession) FreeBalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWToken.Contract.FreeBalanceOfID(&_KNWToken.CallOpts, _address, _id)
}

// FreeBalanceOfID is a free data retrieval call binding the contract method 0xb5c2cdba.
//
// Solidity: function freeBalanceOfID(address _address, uint256 _id) constant returns(uint256 freeBalance)
func (_KNWToken *KNWTokenCallerSession) FreeBalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWToken.Contract.FreeBalanceOfID(&_KNWToken.CallOpts, _address, _id)
}

// LabelOfID is a free data retrieval call binding the contract method 0xcce38484.
//
// Solidity: function labelOfID(uint256 _id) constant returns(string label)
func (_KNWToken *KNWTokenCaller) LabelOfID(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "labelOfID", _id)
	return *ret0, err
}

// LabelOfID is a free data retrieval call binding the contract method 0xcce38484.
//
// Solidity: function labelOfID(uint256 _id) constant returns(string label)
func (_KNWToken *KNWTokenSession) LabelOfID(_id *big.Int) (string, error) {
	return _KNWToken.Contract.LabelOfID(&_KNWToken.CallOpts, _id)
}

// LabelOfID is a free data retrieval call binding the contract method 0xcce38484.
//
// Solidity: function labelOfID(uint256 _id) constant returns(string label)
func (_KNWToken *KNWTokenCallerSession) LabelOfID(_id *big.Int) (string, error) {
	return _KNWToken.Contract.LabelOfID(&_KNWToken.CallOpts, _id)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_KNWToken *KNWTokenCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_KNWToken *KNWTokenSession) Manager() (common.Address, error) {
	return _KNWToken.Contract.Manager(&_KNWToken.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_KNWToken *KNWTokenCallerSession) Manager() (common.Address, error) {
	return _KNWToken.Contract.Manager(&_KNWToken.CallOpts)
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

// TotalIDSupply is a free data retrieval call binding the contract method 0x2d243d1e.
//
// Solidity: function totalIDSupply(uint256 _id) constant returns(uint256 totalSupplyOfID)
func (_KNWToken *KNWTokenCaller) TotalIDSupply(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "totalIDSupply", _id)
	return *ret0, err
}

// TotalIDSupply is a free data retrieval call binding the contract method 0x2d243d1e.
//
// Solidity: function totalIDSupply(uint256 _id) constant returns(uint256 totalSupplyOfID)
func (_KNWToken *KNWTokenSession) TotalIDSupply(_id *big.Int) (*big.Int, error) {
	return _KNWToken.Contract.TotalIDSupply(&_KNWToken.CallOpts, _id)
}

// TotalIDSupply is a free data retrieval call binding the contract method 0x2d243d1e.
//
// Solidity: function totalIDSupply(uint256 _id) constant returns(uint256 totalSupplyOfID)
func (_KNWToken *KNWTokenCallerSession) TotalIDSupply(_id *big.Int) (*big.Int, error) {
	return _KNWToken.Contract.TotalIDSupply(&_KNWToken.CallOpts, _id)
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

// AddNewLabel is a paid mutator transaction binding the contract method 0x4ac99804.
//
// Solidity: function addNewLabel(string _newLabel) returns(bool success)
func (_KNWToken *KNWTokenTransactor) AddNewLabel(opts *bind.TransactOpts, _newLabel string) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "addNewLabel", _newLabel)
}

// AddNewLabel is a paid mutator transaction binding the contract method 0x4ac99804.
//
// Solidity: function addNewLabel(string _newLabel) returns(bool success)
func (_KNWToken *KNWTokenSession) AddNewLabel(_newLabel string) (*types.Transaction, error) {
	return _KNWToken.Contract.AddNewLabel(&_KNWToken.TransactOpts, _newLabel)
}

// AddNewLabel is a paid mutator transaction binding the contract method 0x4ac99804.
//
// Solidity: function addNewLabel(string _newLabel) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) AddNewLabel(_newLabel string) (*types.Transaction, error) {
	return _KNWToken.Contract.AddNewLabel(&_KNWToken.TransactOpts, _newLabel)
}

// AuthorizeAddress is a paid mutator transaction binding the contract method 0x4a5db3b5.
//
// Solidity: function authorizeAddress(address _address) returns(bool success)
func (_KNWToken *KNWTokenTransactor) AuthorizeAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "authorizeAddress", _address)
}

// AuthorizeAddress is a paid mutator transaction binding the contract method 0x4a5db3b5.
//
// Solidity: function authorizeAddress(address _address) returns(bool success)
func (_KNWToken *KNWTokenSession) AuthorizeAddress(_address common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.AuthorizeAddress(&_KNWToken.TransactOpts, _address)
}

// AuthorizeAddress is a paid mutator transaction binding the contract method 0x4a5db3b5.
//
// Solidity: function authorizeAddress(address _address) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) AuthorizeAddress(_address common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.AuthorizeAddress(&_KNWToken.TransactOpts, _address)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) Burn(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "burn", _address, _id, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) Burn(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Burn(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) Burn(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Burn(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0xa25983e5.
//
// Solidity: function lockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) LockTokens(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "lockTokens", _address, _id, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0xa25983e5.
//
// Solidity: function lockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) LockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.LockTokens(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0xa25983e5.
//
// Solidity: function lockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) LockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.LockTokens(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) Mint(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "mint", _address, _id, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) Mint(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Mint(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) Mint(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Mint(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x4ba79dfe.
//
// Solidity: function removeAddress(address _address) returns(bool success)
func (_KNWToken *KNWTokenTransactor) RemoveAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "removeAddress", _address)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x4ba79dfe.
//
// Solidity: function removeAddress(address _address) returns(bool success)
func (_KNWToken *KNWTokenSession) RemoveAddress(_address common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.RemoveAddress(&_KNWToken.TransactOpts, _address)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x4ba79dfe.
//
// Solidity: function removeAddress(address _address) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) RemoveAddress(_address common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.RemoveAddress(&_KNWToken.TransactOpts, _address)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _address) returns(bool success)
func (_KNWToken *KNWTokenTransactor) ReplaceManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "replaceManager", _address)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _address) returns(bool success)
func (_KNWToken *KNWTokenSession) ReplaceManager(_address common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.ReplaceManager(&_KNWToken.TransactOpts, _address)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _address) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) ReplaceManager(_address common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.ReplaceManager(&_KNWToken.TransactOpts, _address)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xcd3877df.
//
// Solidity: function unlockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactor) UnlockTokens(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "unlockTokens", _address, _id, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xcd3877df.
//
// Solidity: function unlockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenSession) UnlockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.UnlockTokens(&_KNWToken.TransactOpts, _address, _id, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xcd3877df.
//
// Solidity: function unlockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWToken *KNWTokenTransactorSession) UnlockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.UnlockTokens(&_KNWToken.TransactOpts, _address, _id, _amount)
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
	Id    *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed who, uint256 id, uint256 value)
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

// WatchBurn is a free log subscription operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed who, uint256 id, uint256 value)
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

// ParseBurn is a log parse operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed who, uint256 id, uint256 value)
func (_KNWToken *KNWTokenFilterer) ParseBurn(log types.Log) (*KNWTokenBurn, error) {
	event := new(KNWTokenBurn)
	if err := _KNWToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Id    *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed who, uint256 id, uint256 value)
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

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed who, uint256 id, uint256 value)
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed who, uint256 id, uint256 value)
func (_KNWToken *KNWTokenFilterer) ParseMint(log types.Log) (*KNWTokenMint, error) {
	event := new(KNWTokenMint)
	if err := _KNWToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KNWTokenNewLabelIterator is returned from FilterNewLabel and is used to iterate over the raw logs and unpacked data for NewLabel events raised by the KNWToken contract.
type KNWTokenNewLabelIterator struct {
	Event *KNWTokenNewLabel // Event containing the contract specifics and raw log

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
func (it *KNWTokenNewLabelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KNWTokenNewLabel)
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
		it.Event = new(KNWTokenNewLabel)
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
func (it *KNWTokenNewLabelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KNWTokenNewLabelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KNWTokenNewLabel represents a NewLabel event raised by the KNWToken contract.
type KNWTokenNewLabel struct {
	Id    *big.Int
	Label string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewLabel is a free log retrieval operation binding the contract event 0xf7e3164ba69dd943954519f6bce38af5edcad970e8160eace0ff12489b95ea3b.
//
// Solidity: event NewLabel(uint256 indexed id, string label)
func (_KNWToken *KNWTokenFilterer) FilterNewLabel(opts *bind.FilterOpts, id []*big.Int) (*KNWTokenNewLabelIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KNWToken.contract.FilterLogs(opts, "NewLabel", idRule)
	if err != nil {
		return nil, err
	}
	return &KNWTokenNewLabelIterator{contract: _KNWToken.contract, event: "NewLabel", logs: logs, sub: sub}, nil
}

// WatchNewLabel is a free log subscription operation binding the contract event 0xf7e3164ba69dd943954519f6bce38af5edcad970e8160eace0ff12489b95ea3b.
//
// Solidity: event NewLabel(uint256 indexed id, string label)
func (_KNWToken *KNWTokenFilterer) WatchNewLabel(opts *bind.WatchOpts, sink chan<- *KNWTokenNewLabel, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KNWToken.contract.WatchLogs(opts, "NewLabel", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KNWTokenNewLabel)
				if err := _KNWToken.contract.UnpackLog(event, "NewLabel", log); err != nil {
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

// ParseNewLabel is a log parse operation binding the contract event 0xf7e3164ba69dd943954519f6bce38af5edcad970e8160eace0ff12489b95ea3b.
//
// Solidity: event NewLabel(uint256 indexed id, string label)
func (_KNWToken *KNWTokenFilterer) ParseNewLabel(log types.Log) (*KNWTokenNewLabel, error) {
	event := new(KNWTokenNewLabel)
	if err := _KNWToken.contract.UnpackLog(event, "NewLabel", log); err != nil {
		return nil, err
	}
	return event, nil
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
