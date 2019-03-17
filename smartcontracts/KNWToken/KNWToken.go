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
const KNWTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"freeBalanceOfLabel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"lockTokens\",\"outputs\":[{\"name\":\"numberOfTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_winningPercentage\",\"type\":\"uint256\"},{\"name\":\"_mintingMethod\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_labelID\",\"type\":\"uint256\"}],\"name\":\"labelOfAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newVotingAddress\",\"type\":\"address\"}],\"name\":\"setVotingAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_stakedTokens\",\"type\":\"uint256\"},{\"name\":\"_winningPercentage\",\"type\":\"uint256\"},{\"name\":\"_burningMethod\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"balanceOfLabel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_label\",\"type\":\"string\"}],\"name\":\"totalLabelSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_label\",\"type\":\"string\"},{\"name\":\"_numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"labelCountOfAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"label\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// KNWTokenBin is the compiled bytecode used for deploying new contracts.
const KNWTokenBin = `0x608060405234801561001057600080fd5b506119f7806100206000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306970f1c81146100df57806306fdde031461015857806318160ddd146101e25780631d3316d0146101f7578063313ce5671461025e5780634a5dc3aa1461028957806358e546cf146102be5780637a6cfcab146102e257806395d89b4114610303578063994dd93a14610318578063b88c0f981461034e578063c1a63f3c146103b5578063d2fa71701461040e578063d950df341461043f578063e7017fc4146104a8575b600080fd5b3480156100eb57600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610146958335600160a060020a03169536956044949193909101919081908401838280828437509497506104c99650505050505050565b60408051918252519081900360200190f35b34801561016457600080fd5b5061016d6105ca565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101a757818101518382015260200161018f565b50505050905090810190601f1680156101d45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ee57600080fd5b50610146610601565b34801561020357600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610146958335600160a060020a03169536956044949193909101919081908401838280828437509497506106089650505050505050565b34801561026a57600080fd5b506102736108e0565b6040805160ff9092168252519081900360200190f35b34801561029557600080fd5b506102bc60048035600160a060020a031690602480359081019101356044356064356108e5565b005b3480156102ca57600080fd5b5061016d600160a060020a0360043516602435610c67565b3480156102ee57600080fd5b506102bc600160a060020a0360043516610d19565b34801561030f57600080fd5b5061016d610e0b565b34801561032457600080fd5b506102bc60048035600160a060020a03169060248035908101910135604435606435608435610e42565b34801561035a57600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610146958335600160a060020a03169536956044949193909101919081908401838280828437509497506113339650505050505050565b3480156103c157600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101469436949293602493928401919081908401838280828437509497506113b19650505050505050565b34801561041a57600080fd5b50610423611419565b60408051600160a060020a039092168252519081900360200190f35b34801561044b57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102bc958335600160a060020a031695369560449491939091019190819084018382808284375094975050933594506114289350505050565b3480156104b457600080fd5b50610146600160a060020a03600435166116df565b600160a060020a0382166000908152600160209081526040808320905184516105c39386929182918401908083835b602083106105175780518252601f1990920191602091820191016104f8565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852054600160a060020a038a1660009081528084529190912088519195909450889350918291908401908083835b6020831061058e5780518252601f19909201916020918201910161056f565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054929150506116fa565b9392505050565b60408051808201909152600f81527f4b6e6f776c6564676520546f6b656e0000000000000000000000000000000000602082015281565b6004545b90565b600654600090600160a060020a0316331461066f576040805160e560020a62461bcd02815260206004820152603360248201526000805160206119ac833981519152604482015260008051602061198c833981519152606482015290519081900360840190fd5b50600160a060020a0382166000908152600160209081526040808320905184519192859282918401908083835b602083106106bb5780518252601f19909201916020918201910161069c565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852054600160a060020a03891660009081528084529190912087519195909450879350918291908401908083835b602083106107325780518252601f199092019160209182019101610713565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205411156108da576107c36001600085600160a060020a0316600160a060020a0316815260200190815260200160002083604051808280519060200190808383602083106105175780518252601f1990920191602091820191016104f8565b9050610853816001600086600160a060020a0316600160a060020a03168152602001908152602001600020846040518082805190602001908083835b6020831061081e5780518252601f1990920191602091820191016107ff565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205492915050611789565b6001600085600160a060020a0316600160a060020a03168152602001908152602001600020836040518082805190602001908083835b602083106108a85780518252601f199092019160209182019101610889565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209290925550505b92915050565b601281565b600654600090600160a060020a0316331461094c576040805160e560020a62461bcd02815260206004820152603360248201526000805160206119ac833981519152604482015260008051602061198c833981519152606482015290519081900360840190fd5b600160a060020a03861615156109ac576040805160e560020a62461bcd02815260206004820152601660248201527f416464726573732063616e277420626520656d70747900000000000000000000604482015290519081900360640190fd5b60008411610a04576040805160e560020a62461bcd02815260206004820152601e60248201527f4b6e6f776c656467652d4c6162656c2063616e277420626520656d7074790000604482015290519081900360640190fd5b506000811515610a3957610a3666470de4df820000610a2a85603263ffffffff6116fa16565b9063ffffffff61180c16565b90505b600454610a4c908263ffffffff61178916565b600481905550610a89816005878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61178916565b6005868660405180838380828437909101948552505060408051938490036020908101852095909555600160a060020a038b16600090815294859052909320925087918791508083838082843782019150509250505090815260200160405180910390205460001415610b5957600160a060020a038616600090815260036020526040902054610b2090600163ffffffff61178916565b600160a060020a03871660009081526003602090815260408083208490556002825280832093835292905220610b579086866118f3565b505b610bb28160008089600160a060020a0316600160a060020a03168152602001908152602001600020878760405180838380828437909101948552505060405192839003602001909220549291505063ffffffff61178916565b60008088600160a060020a0316600160a060020a0316815260200190815260200160002086866040518083838082843782019150509250505090815260200160405180910390208190555085600160a060020a03167fec4de1eef14af3ae5d77facf1ed7a9d3d50f6285573ee0ec155fc11217fc3442868684604051808060200183815260200182810382528585828181526020019250808284376040519201829003965090945050505050a2505050505050565b600160a060020a0382166000908152600260208181526040808420858552825292839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383526060939091830182828015610d0c5780601f10610ce157610100808354040283529160200191610d0c565b820191906000526020600020905b815481529060010190602001808311610cef57829003601f168201915b5050505050905092915050565b600160a060020a03811615801590610d3a5750600654600160a060020a0316155b1515610ddc576040805160e560020a62461bcd02815260206004820152604f60248201527f4b4e57566f74696e6720616464726573732063616e206f6e6c7920626520736560448201527f742069662069742773206e6f7420656d70747920616e64206861736e2774206160648201527f6c7265616479206265656e207365740000000000000000000000000000000000608482015290519081900360a40190fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60408051808201909152600381527f4b4e570000000000000000000000000000000000000000000000000000000000602082015281565b60065460009081908190600160a060020a03163314610ead576040805160e560020a62461bcd02815260206004820152603360248201526000805160206119ac833981519152604482015260008051602061198c833981519152606482015290519081900360840190fd5b600160a060020a0389161515610f0d576040805160e560020a62461bcd02815260206004820152601660248201527f416464726573732063616e277420626520656d70747900000000000000000000604482015290519081900360640190fd5b60008711610f65576040805160e560020a62461bcd02815260206004820152601e60248201527f4b6e6f776c656467652d4c6162656c2063616e277420626520656d7074790000604482015290519081900360640190fd5b856000808b600160a060020a0316600160a060020a031681526020019081526020016000208989604051808383808284378201915050925050509081526020016040518091039020541015151561102c576040805160e560020a62461bcd02815260206004820152602a60248201527f43616e2774206275726e206d6f7265204b4e57207468616e207468652061646460448201527f7265737320686f6c647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b85925060008611156113285783151561113d5761110366038d7ea4c6800061105f8864e8d4a5100063ffffffff61188516565b73__libraries/SafeMath.sol:SafeMath_______63677342ce90916040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060206040518083038186803b1580156110cb57600080fd5b505af41580156110df573d6000803e3d6000fd5b505050506040513d60208110156110f557600080fd5b50519063ffffffff61180c16565b9150858210156111245761111d838363ffffffff6116fa16565b9250611138565b61113583600263ffffffff61188516565b92505b6111a8565b83600114156111575761113583600263ffffffff61188516565b83600214156111a857611182606461117687600263ffffffff61180c16565b9063ffffffff6116fa16565b90506111a56064611199858463ffffffff61180c16565b9063ffffffff61188516565b92505b6004546111bb908463ffffffff6116fa16565b6004819055506111f88360058a8a60405180838380828437909101948552505060405192839003602001909220549291505063ffffffff6116fa16565b600589896040518083838082843782019150509250505090815260200160405180910390208190555061127a836000808c600160a060020a0316600160a060020a031681526020019081526020016000208a8a60405180838380828437909101948552505060405192839003602001909220549291505063ffffffff6116fa16565b6000808b600160a060020a0316600160a060020a0316815260200190815260200160002089896040518083838082843782019150509250505090815260200160405180910390208190555088600160a060020a03167ffdf096248d2b7b0aef506231c043107c21faacc26193881b3f0cdc8b5479692a898986604051808060200183815260200182810382528585828181526020019250808284376040519201829003965090945050505050a25b505050505050505050565b600160a060020a038216600090815260208181526040808320905184519192859282918401908083835b6020831061137c5780518252601f19909201916020918201910161135d565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205495945050505050565b60006005826040518082805190602001908083835b602083106113e55780518252601f1990920191602091820191016113c6565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054949350505050565b600654600160a060020a031681565b600654600160a060020a0316331461148c576040805160e560020a62461bcd02815260206004820152603360248201526000805160206119ac833981519152604482015260008051602061198c833981519152606482015290519081900360840190fd5b600160a060020a038316600090815260208181526040918290209151845185928291908401908083835b602083106114d55780518252601f1990920191602091820191016114b6565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852054600160a060020a0389166000908152600184529190912087519195909450879350918291908401908083835b6020831061154d5780518252601f19909201916020918201910161152e565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054929092111591506115fc9050576040805160e560020a62461bcd02815260206004820152602660248201527f43616e74206c6f636b206d6f7265204b4e57207468616e20616e20616464726560448201527f7373206861730000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b611654816001600086600160a060020a0316600160a060020a03168152602001908152602001600020846040518082805190602001908083836020831061058e5780518252601f19909201916020918201910161056f565b6001600085600160a060020a0316600160a060020a03168152602001908152602001600020836040518082805190602001908083835b602083106116a95780518252601f19909201916020918201910161168a565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929092555050505050565b600160a060020a031660009081526003602052604090205490565b6000808383111561177b576040805160e560020a62461bcd02815260206004820152603560248201527f43616e27742073756274726163742061206e756d6265722066726f6d2061207360448201527f6d616c6c6572206f6e6520776974682075696e74730000000000000000000000606482015290519081900360840190fd5b5050808203805b5092915050565b6000828201838110156105c3576040805160e560020a62461bcd02815260206004820152602a60248201527f526573756c742068617320746f20626520626967676572207468616e20626f7460448201527f682073756d6d616e647300000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008083151561181f5760009150611782565b5082820282848281151561182f57fe5b04146105c3576040805160e560020a62461bcd02815260206004820152601f60248201527f466c6177656420696e70757420666f72206d756c7469706c69636174696f6e00604482015290519081900360640190fd5b6000808083116118df576040805160e560020a62461bcd02815260206004820152601460248201527f43616e277420646976696465206279207a65726f000000000000000000000000604482015290519081900360640190fd5b82848115156118ea57fe5b04949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119345782800160ff19823516178555611961565b82800160010185558215611961579182015b82811115611961578235825591602001919060010190611946565b5061196d929150611971565b5090565b61060591905b8082111561196d576000815560010161197756006c6c6f77656420746f2063616c6c2074686973000000000000000000000000004f6e6c7920746865204b4e57566f74696e6720636f6e74726163742069732061a165627a7a723058203a0a305bfaea031add5c672d6d2f405ffc93b4ceb613b4124f09f4420909a3ea0029`

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
// Solidity: function balanceOfLabel(_address address, _label string) constant returns(uint256)
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
// Solidity: function balanceOfLabel(_address address, _label string) constant returns(uint256)
func (_KNWToken *KNWTokenSession) BalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.BalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// BalanceOfLabel is a free data retrieval call binding the contract method 0xb88c0f98.
//
// Solidity: function balanceOfLabel(_address address, _label string) constant returns(uint256)
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
// Solidity: function freeBalanceOfLabel(_address address, _label string) constant returns(uint256)
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
// Solidity: function freeBalanceOfLabel(_address address, _label string) constant returns(uint256)
func (_KNWToken *KNWTokenSession) FreeBalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.FreeBalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// FreeBalanceOfLabel is a free data retrieval call binding the contract method 0x06970f1c.
//
// Solidity: function freeBalanceOfLabel(_address address, _label string) constant returns(uint256)
func (_KNWToken *KNWTokenCallerSession) FreeBalanceOfLabel(_address common.Address, _label string) (*big.Int, error) {
	return _KNWToken.Contract.FreeBalanceOfLabel(&_KNWToken.CallOpts, _address, _label)
}

// LabelCountOfAddress is a free data retrieval call binding the contract method 0xe7017fc4.
//
// Solidity: function labelCountOfAddress(_address address) constant returns(uint256)
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
// Solidity: function labelCountOfAddress(_address address) constant returns(uint256)
func (_KNWToken *KNWTokenSession) LabelCountOfAddress(_address common.Address) (*big.Int, error) {
	return _KNWToken.Contract.LabelCountOfAddress(&_KNWToken.CallOpts, _address)
}

// LabelCountOfAddress is a free data retrieval call binding the contract method 0xe7017fc4.
//
// Solidity: function labelCountOfAddress(_address address) constant returns(uint256)
func (_KNWToken *KNWTokenCallerSession) LabelCountOfAddress(_address common.Address) (*big.Int, error) {
	return _KNWToken.Contract.LabelCountOfAddress(&_KNWToken.CallOpts, _address)
}

// LabelOfAddress is a free data retrieval call binding the contract method 0x58e546cf.
//
// Solidity: function labelOfAddress(_address address, _labelID uint256) constant returns(string)
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
// Solidity: function labelOfAddress(_address address, _labelID uint256) constant returns(string)
func (_KNWToken *KNWTokenSession) LabelOfAddress(_address common.Address, _labelID *big.Int) (string, error) {
	return _KNWToken.Contract.LabelOfAddress(&_KNWToken.CallOpts, _address, _labelID)
}

// LabelOfAddress is a free data retrieval call binding the contract method 0x58e546cf.
//
// Solidity: function labelOfAddress(_address address, _labelID uint256) constant returns(string)
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
// Solidity: function totalLabelSupply(_label string) constant returns(uint256)
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
// Solidity: function totalLabelSupply(_label string) constant returns(uint256)
func (_KNWToken *KNWTokenSession) TotalLabelSupply(_label string) (*big.Int, error) {
	return _KNWToken.Contract.TotalLabelSupply(&_KNWToken.CallOpts, _label)
}

// TotalLabelSupply is a free data retrieval call binding the contract method 0xc1a63f3c.
//
// Solidity: function totalLabelSupply(_label string) constant returns(uint256)
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

// VotingAddress is a free data retrieval call binding the contract method 0xd2fa7170.
//
// Solidity: function votingAddress() constant returns(address)
func (_KNWToken *KNWTokenCaller) VotingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWToken.contract.Call(opts, out, "votingAddress")
	return *ret0, err
}

// VotingAddress is a free data retrieval call binding the contract method 0xd2fa7170.
//
// Solidity: function votingAddress() constant returns(address)
func (_KNWToken *KNWTokenSession) VotingAddress() (common.Address, error) {
	return _KNWToken.Contract.VotingAddress(&_KNWToken.CallOpts)
}

// VotingAddress is a free data retrieval call binding the contract method 0xd2fa7170.
//
// Solidity: function votingAddress() constant returns(address)
func (_KNWToken *KNWTokenCallerSession) VotingAddress() (common.Address, error) {
	return _KNWToken.Contract.VotingAddress(&_KNWToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x994dd93a.
//
// Solidity: function burn(_address address, _label string, _stakedTokens uint256, _winningPercentage uint256, _burningMethod uint256) returns()
func (_KNWToken *KNWTokenTransactor) Burn(opts *bind.TransactOpts, _address common.Address, _label string, _stakedTokens *big.Int, _winningPercentage *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "burn", _address, _label, _stakedTokens, _winningPercentage, _burningMethod)
}

// Burn is a paid mutator transaction binding the contract method 0x994dd93a.
//
// Solidity: function burn(_address address, _label string, _stakedTokens uint256, _winningPercentage uint256, _burningMethod uint256) returns()
func (_KNWToken *KNWTokenSession) Burn(_address common.Address, _label string, _stakedTokens *big.Int, _winningPercentage *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Burn(&_KNWToken.TransactOpts, _address, _label, _stakedTokens, _winningPercentage, _burningMethod)
}

// Burn is a paid mutator transaction binding the contract method 0x994dd93a.
//
// Solidity: function burn(_address address, _label string, _stakedTokens uint256, _winningPercentage uint256, _burningMethod uint256) returns()
func (_KNWToken *KNWTokenTransactorSession) Burn(_address common.Address, _label string, _stakedTokens *big.Int, _winningPercentage *big.Int, _burningMethod *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Burn(&_KNWToken.TransactOpts, _address, _label, _stakedTokens, _winningPercentage, _burningMethod)
}

// LockTokens is a paid mutator transaction binding the contract method 0x1d3316d0.
//
// Solidity: function lockTokens(_address address, _label string) returns(numberOfTokens uint256)
func (_KNWToken *KNWTokenTransactor) LockTokens(opts *bind.TransactOpts, _address common.Address, _label string) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "lockTokens", _address, _label)
}

// LockTokens is a paid mutator transaction binding the contract method 0x1d3316d0.
//
// Solidity: function lockTokens(_address address, _label string) returns(numberOfTokens uint256)
func (_KNWToken *KNWTokenSession) LockTokens(_address common.Address, _label string) (*types.Transaction, error) {
	return _KNWToken.Contract.LockTokens(&_KNWToken.TransactOpts, _address, _label)
}

// LockTokens is a paid mutator transaction binding the contract method 0x1d3316d0.
//
// Solidity: function lockTokens(_address address, _label string) returns(numberOfTokens uint256)
func (_KNWToken *KNWTokenTransactorSession) LockTokens(_address common.Address, _label string) (*types.Transaction, error) {
	return _KNWToken.Contract.LockTokens(&_KNWToken.TransactOpts, _address, _label)
}

// Mint is a paid mutator transaction binding the contract method 0x4a5dc3aa.
//
// Solidity: function mint(_address address, _label string, _winningPercentage uint256, _mintingMethod uint256) returns()
func (_KNWToken *KNWTokenTransactor) Mint(opts *bind.TransactOpts, _address common.Address, _label string, _winningPercentage *big.Int, _mintingMethod *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "mint", _address, _label, _winningPercentage, _mintingMethod)
}

// Mint is a paid mutator transaction binding the contract method 0x4a5dc3aa.
//
// Solidity: function mint(_address address, _label string, _winningPercentage uint256, _mintingMethod uint256) returns()
func (_KNWToken *KNWTokenSession) Mint(_address common.Address, _label string, _winningPercentage *big.Int, _mintingMethod *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Mint(&_KNWToken.TransactOpts, _address, _label, _winningPercentage, _mintingMethod)
}

// Mint is a paid mutator transaction binding the contract method 0x4a5dc3aa.
//
// Solidity: function mint(_address address, _label string, _winningPercentage uint256, _mintingMethod uint256) returns()
func (_KNWToken *KNWTokenTransactorSession) Mint(_address common.Address, _label string, _winningPercentage *big.Int, _mintingMethod *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.Mint(&_KNWToken.TransactOpts, _address, _label, _winningPercentage, _mintingMethod)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(_newVotingAddress address) returns()
func (_KNWToken *KNWTokenTransactor) SetVotingAddress(opts *bind.TransactOpts, _newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "setVotingAddress", _newVotingAddress)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(_newVotingAddress address) returns()
func (_KNWToken *KNWTokenSession) SetVotingAddress(_newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.SetVotingAddress(&_KNWToken.TransactOpts, _newVotingAddress)
}

// SetVotingAddress is a paid mutator transaction binding the contract method 0x7a6cfcab.
//
// Solidity: function setVotingAddress(_newVotingAddress address) returns()
func (_KNWToken *KNWTokenTransactorSession) SetVotingAddress(_newVotingAddress common.Address) (*types.Transaction, error) {
	return _KNWToken.Contract.SetVotingAddress(&_KNWToken.TransactOpts, _newVotingAddress)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(_address address, _label string, _numberOfTokens uint256) returns()
func (_KNWToken *KNWTokenTransactor) UnlockTokens(opts *bind.TransactOpts, _address common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWToken.contract.Transact(opts, "unlockTokens", _address, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(_address address, _label string, _numberOfTokens uint256) returns()
func (_KNWToken *KNWTokenSession) UnlockTokens(_address common.Address, _label string, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _KNWToken.Contract.UnlockTokens(&_KNWToken.TransactOpts, _address, _label, _numberOfTokens)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xd950df34.
//
// Solidity: function unlockTokens(_address address, _label string, _numberOfTokens uint256) returns()
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
// Solidity: e Burn(who indexed address, label string, value uint256)
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
// Solidity: e Burn(who indexed address, label string, value uint256)
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
// Solidity: e Mint(who indexed address, label string, value uint256)
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
// Solidity: e Mint(who indexed address, label string, value uint256)
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
