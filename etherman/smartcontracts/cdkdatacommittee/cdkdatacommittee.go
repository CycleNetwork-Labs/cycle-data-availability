// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cdkdatacommittee

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CdkdatacommitteeMetaData contains all meta data concerning the Cdkdatacommittee contract.
var CdkdatacommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitteeAddressDoesntExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyURLNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequiredSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsAndSignaturesSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedCommitteeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAddrOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"committeeHash\",\"type\":\"bytes32\"}],\"name\":\"CommitteeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"committeeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmountOfMembers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredAmountOfSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requiredAmountOfSignatures\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"addrsBytes\",\"type\":\"bytes\"}],\"name\":\"setupCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signaturesAndAddrs\",\"type\":\"bytes\"}],\"name\":\"verifySignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061121e8061001d5f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c80638129fc1c11610072578063c7a823e011610058578063c7a823e014610147578063dce1e2b61461015a578063f2fde38b14610162575f80fd5b80638129fc1c146101245780638da5cb5b1461012c575f80fd5b8063609d4544116100a2578063609d4544146100fc5780636beedd3914610113578063715018a61461011c575f80fd5b8063078fba2a146100bd5780635daf08ca146100d2575b5f80fd5b6100d06100cb366004610d90565b610175565b005b6100e56100e0366004610e32565b610434565b6040516100f3929190610e49565b60405180910390f35b61010560665481565b6040519081526020016100f3565b61010560655481565b6100d06104f2565b6100d0610505565b6033546040516001600160a01b0390911681526020016100f3565b6100d0610155366004610ea4565b610624565b606754610105565b6100d0610170366004610eec565b610862565b61017d6108ef565b82858110156101b8576040517f2e7dcd6e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101c3601482610f2d565b82146101fb576040517f2ab6a12900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61020660675f610cb3565b5f805b828110156103d8575f61021d601483610f2d565b90505f86828761022e601483610f44565b9261023b93929190610f57565b61024491610f7e565b60601c905088888481811061025b5761025b610fb3565b905060200281019061026d9190610fc7565b90505f036102a7576040517fb54b70e400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b0316846001600160a01b0316106102f2576040517fd53cfbe000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b809350606760405180604001604052808b8b8781811061031457610314610fb3565b90506020028101906103269190610fc7565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160a01b038516602092830152835460018101855593815220815191926002020190819061038c90826110a4565b50602091909101516001909101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055508190506103d081611160565b915050610209565b5083836040516103e9929190611178565b6040519081900381206066819055606589905581527f831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db6579060200160405180910390a150505050505050565b60678181548110610443575f80fd5b905f5260205f2090600202015f91509050805f0180546104629061101e565b80601f016020809104026020016040519081016040528092919081815260200182805461048e9061101e565b80156104d95780601f106104b0576101008083540402835291602001916104d9565b820191905f5260205f20905b8154815290600101906020018083116104bc57829003601f168201915b505050600190930154919250506001600160a01b031682565b6104fa6108ef565b6105035f610949565b565b5f54610100900460ff161580801561052357505f54600160ff909116105b8061053c5750303b15801561053c57505f5460ff166001145b6105b35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f805460ff1916600117905580156105d4575f805461ff0019166101001790555b6105dc6109a7565b8015610621575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b5f60655460416106349190610f2d565b9050808210806106585750601461064b8284611187565b61065591906111ae565b15155b1561068f576040517f6b8eec4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60665461069e83838187610f57565b6040516106ac929190611178565b6040518091039020146106eb576040517f6b156b2800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8060146106f98486611187565b61070391906111c1565b90505f5b606554811015610859575f610780888888610723604187610f2d565b9060416107308189610f2d565b61073a9190610f44565b9261074793929190610f57565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610a2c92505050565b90505f845b8481101561080c575f610799601483610f2d565b6107a39089610f44565b90505f8a828b6107b4601483610f44565b926107c193929190610f57565b6107ca91610f7e565b60601c90506001600160a01b03851681036107f7576107ea836001610f44565b975060019350505061080c565b5050808061080490611160565b915050610785565b5080610844576040517f8431721300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050808061085190611160565b915050610707565b50505050505050565b61086a6108ef565b6001600160a01b0381166108e65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105aa565b61062181610949565b6033546001600160a01b031633146105035760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105aa565b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610a235760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105aa565b61050333610949565b5f805f610a398585610a50565b91509150610a4681610a92565b5090505b92915050565b5f808251604103610a84576020830151604084015160608501515f1a610a7887828585610bf6565b94509450505050610a8b565b505f905060025b9250929050565b5f816004811115610aa557610aa56111d4565b03610aad5750565b6001816004811115610ac157610ac16111d4565b03610b0e5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105aa565b6002816004811115610b2257610b226111d4565b03610b6f5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105aa565b6003816004811115610b8357610b836111d4565b036106215760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016105aa565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610c2b57505f90506003610caa565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610c7c573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116610ca4575f60019250925050610caa565b91505f90505b94509492505050565b5080545f8255600202905f5260205f209081019061062191905b80821115610d08575f610ce08282610d0c565b5060018101805473ffffffffffffffffffffffffffffffffffffffff19169055600201610ccd565b5090565b508054610d189061101e565b5f825580601f10610d27575050565b601f0160209004905f5260205f209081019061062191905b80821115610d08575f8155600101610d3f565b5f8083601f840112610d62575f80fd5b50813567ffffffffffffffff811115610d79575f80fd5b602083019150836020828501011115610a8b575f80fd5b5f805f805f60608688031215610da4575f80fd5b85359450602086013567ffffffffffffffff80821115610dc2575f80fd5b818801915088601f830112610dd5575f80fd5b813581811115610de3575f80fd5b8960208260051b8501011115610df7575f80fd5b602083019650809550506040880135915080821115610e14575f80fd5b50610e2188828901610d52565b969995985093965092949392505050565b5f60208284031215610e42575f80fd5b5035919050565b604081525f83518060408401525f5b81811015610e755760208187018101516060868401015201610e58565b505f606082850101526060601f19601f8301168401019150506001600160a01b03831660208301529392505050565b5f805f60408486031215610eb6575f80fd5b83359250602084013567ffffffffffffffff811115610ed3575f80fd5b610edf86828701610d52565b9497909650939450505050565b5f60208284031215610efc575f80fd5b81356001600160a01b0381168114610f12575f80fd5b9392505050565b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417610a4a57610a4a610f19565b80820180821115610a4a57610a4a610f19565b5f8085851115610f65575f80fd5b83861115610f71575f80fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015610fab5780818660140360031b1b83161692505b505092915050565b634e487b7160e01b5f52603260045260245ffd5b5f808335601e19843603018112610fdc575f80fd5b83018035915067ffffffffffffffff821115610ff6575f80fd5b602001915036819003821315610a8b575f80fd5b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061103257607f821691505b60208210810361105057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561109f575f81815260208120601f850160051c8101602086101561107c5750805b601f850160051c820191505b8181101561109b57828155600101611088565b5050505b505050565b815167ffffffffffffffff8111156110be576110be61100a565b6110d2816110cc845461101e565b84611056565b602080601f831160018114611105575f84156110ee5750858301515b5f19600386901b1c1916600185901b17855561109b565b5f85815260208120601f198616915b8281101561113357888601518255948401946001909101908401611114565b508582101561115057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f6001820161117157611171610f19565b5060010190565b818382375f9101908152919050565b81810381811115610a4a57610a4a610f19565b634e487b7160e01b5f52601260045260245ffd5b5f826111bc576111bc61119a565b500690565b5f826111cf576111cf61119a565b500490565b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220932d8925302346f93d86251cfe14c6b4f8d2958007e06e727fc0eb6ed0b65f9264736f6c63430008140033",
}

// CdkdatacommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use CdkdatacommitteeMetaData.ABI instead.
var CdkdatacommitteeABI = CdkdatacommitteeMetaData.ABI

// CdkdatacommitteeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CdkdatacommitteeMetaData.Bin instead.
var CdkdatacommitteeBin = CdkdatacommitteeMetaData.Bin

// DeployCdkdatacommittee deploys a new Ethereum contract, binding an instance of Cdkdatacommittee to it.
func DeployCdkdatacommittee(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cdkdatacommittee, error) {
	parsed, err := CdkdatacommitteeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CdkdatacommitteeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cdkdatacommittee{CdkdatacommitteeCaller: CdkdatacommitteeCaller{contract: contract}, CdkdatacommitteeTransactor: CdkdatacommitteeTransactor{contract: contract}, CdkdatacommitteeFilterer: CdkdatacommitteeFilterer{contract: contract}}, nil
}

// Cdkdatacommittee is an auto generated Go binding around an Ethereum contract.
type Cdkdatacommittee struct {
	CdkdatacommitteeCaller     // Read-only binding to the contract
	CdkdatacommitteeTransactor // Write-only binding to the contract
	CdkdatacommitteeFilterer   // Log filterer for contract events
}

// CdkdatacommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CdkdatacommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkdatacommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CdkdatacommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkdatacommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CdkdatacommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdkdatacommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CdkdatacommitteeSession struct {
	Contract     *Cdkdatacommittee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdkdatacommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CdkdatacommitteeCallerSession struct {
	Contract *CdkdatacommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CdkdatacommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CdkdatacommitteeTransactorSession struct {
	Contract     *CdkdatacommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CdkdatacommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CdkdatacommitteeRaw struct {
	Contract *Cdkdatacommittee // Generic contract binding to access the raw methods on
}

// CdkdatacommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CdkdatacommitteeCallerRaw struct {
	Contract *CdkdatacommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// CdkdatacommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CdkdatacommitteeTransactorRaw struct {
	Contract *CdkdatacommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCdkdatacommittee creates a new instance of Cdkdatacommittee, bound to a specific deployed contract.
func NewCdkdatacommittee(address common.Address, backend bind.ContractBackend) (*Cdkdatacommittee, error) {
	contract, err := bindCdkdatacommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cdkdatacommittee{CdkdatacommitteeCaller: CdkdatacommitteeCaller{contract: contract}, CdkdatacommitteeTransactor: CdkdatacommitteeTransactor{contract: contract}, CdkdatacommitteeFilterer: CdkdatacommitteeFilterer{contract: contract}}, nil
}

// NewCdkdatacommitteeCaller creates a new read-only instance of Cdkdatacommittee, bound to a specific deployed contract.
func NewCdkdatacommitteeCaller(address common.Address, caller bind.ContractCaller) (*CdkdatacommitteeCaller, error) {
	contract, err := bindCdkdatacommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CdkdatacommitteeCaller{contract: contract}, nil
}

// NewCdkdatacommitteeTransactor creates a new write-only instance of Cdkdatacommittee, bound to a specific deployed contract.
func NewCdkdatacommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*CdkdatacommitteeTransactor, error) {
	contract, err := bindCdkdatacommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CdkdatacommitteeTransactor{contract: contract}, nil
}

// NewCdkdatacommitteeFilterer creates a new log filterer instance of Cdkdatacommittee, bound to a specific deployed contract.
func NewCdkdatacommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*CdkdatacommitteeFilterer, error) {
	contract, err := bindCdkdatacommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CdkdatacommitteeFilterer{contract: contract}, nil
}

// bindCdkdatacommittee binds a generic wrapper to an already deployed contract.
func bindCdkdatacommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CdkdatacommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdkdatacommittee *CdkdatacommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdkdatacommittee.Contract.CdkdatacommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdkdatacommittee *CdkdatacommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.CdkdatacommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdkdatacommittee *CdkdatacommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.CdkdatacommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdkdatacommittee *CdkdatacommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdkdatacommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdkdatacommittee *CdkdatacommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdkdatacommittee *CdkdatacommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.contract.Transact(opts, method, params...)
}

// CommitteeHash is a free data retrieval call binding the contract method 0x609d4544.
//
// Solidity: function committeeHash() view returns(bytes32)
func (_Cdkdatacommittee *CdkdatacommitteeCaller) CommitteeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cdkdatacommittee.contract.Call(opts, &out, "committeeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitteeHash is a free data retrieval call binding the contract method 0x609d4544.
//
// Solidity: function committeeHash() view returns(bytes32)
func (_Cdkdatacommittee *CdkdatacommitteeSession) CommitteeHash() ([32]byte, error) {
	return _Cdkdatacommittee.Contract.CommitteeHash(&_Cdkdatacommittee.CallOpts)
}

// CommitteeHash is a free data retrieval call binding the contract method 0x609d4544.
//
// Solidity: function committeeHash() view returns(bytes32)
func (_Cdkdatacommittee *CdkdatacommitteeCallerSession) CommitteeHash() ([32]byte, error) {
	return _Cdkdatacommittee.Contract.CommitteeHash(&_Cdkdatacommittee.CallOpts)
}

// GetAmountOfMembers is a free data retrieval call binding the contract method 0xdce1e2b6.
//
// Solidity: function getAmountOfMembers() view returns(uint256)
func (_Cdkdatacommittee *CdkdatacommitteeCaller) GetAmountOfMembers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdkdatacommittee.contract.Call(opts, &out, "getAmountOfMembers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOfMembers is a free data retrieval call binding the contract method 0xdce1e2b6.
//
// Solidity: function getAmountOfMembers() view returns(uint256)
func (_Cdkdatacommittee *CdkdatacommitteeSession) GetAmountOfMembers() (*big.Int, error) {
	return _Cdkdatacommittee.Contract.GetAmountOfMembers(&_Cdkdatacommittee.CallOpts)
}

// GetAmountOfMembers is a free data retrieval call binding the contract method 0xdce1e2b6.
//
// Solidity: function getAmountOfMembers() view returns(uint256)
func (_Cdkdatacommittee *CdkdatacommitteeCallerSession) GetAmountOfMembers() (*big.Int, error) {
	return _Cdkdatacommittee.Contract.GetAmountOfMembers(&_Cdkdatacommittee.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(string url, address addr)
func (_Cdkdatacommittee *CdkdatacommitteeCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Url  string
	Addr common.Address
}, error) {
	var out []interface{}
	err := _Cdkdatacommittee.contract.Call(opts, &out, "members", arg0)

	outstruct := new(struct {
		Url  string
		Addr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Url = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(string url, address addr)
func (_Cdkdatacommittee *CdkdatacommitteeSession) Members(arg0 *big.Int) (struct {
	Url  string
	Addr common.Address
}, error) {
	return _Cdkdatacommittee.Contract.Members(&_Cdkdatacommittee.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(string url, address addr)
func (_Cdkdatacommittee *CdkdatacommitteeCallerSession) Members(arg0 *big.Int) (struct {
	Url  string
	Addr common.Address
}, error) {
	return _Cdkdatacommittee.Contract.Members(&_Cdkdatacommittee.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cdkdatacommittee *CdkdatacommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdkdatacommittee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cdkdatacommittee *CdkdatacommitteeSession) Owner() (common.Address, error) {
	return _Cdkdatacommittee.Contract.Owner(&_Cdkdatacommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cdkdatacommittee *CdkdatacommitteeCallerSession) Owner() (common.Address, error) {
	return _Cdkdatacommittee.Contract.Owner(&_Cdkdatacommittee.CallOpts)
}

// RequiredAmountOfSignatures is a free data retrieval call binding the contract method 0x6beedd39.
//
// Solidity: function requiredAmountOfSignatures() view returns(uint256)
func (_Cdkdatacommittee *CdkdatacommitteeCaller) RequiredAmountOfSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cdkdatacommittee.contract.Call(opts, &out, "requiredAmountOfSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredAmountOfSignatures is a free data retrieval call binding the contract method 0x6beedd39.
//
// Solidity: function requiredAmountOfSignatures() view returns(uint256)
func (_Cdkdatacommittee *CdkdatacommitteeSession) RequiredAmountOfSignatures() (*big.Int, error) {
	return _Cdkdatacommittee.Contract.RequiredAmountOfSignatures(&_Cdkdatacommittee.CallOpts)
}

// RequiredAmountOfSignatures is a free data retrieval call binding the contract method 0x6beedd39.
//
// Solidity: function requiredAmountOfSignatures() view returns(uint256)
func (_Cdkdatacommittee *CdkdatacommitteeCallerSession) RequiredAmountOfSignatures() (*big.Int, error) {
	return _Cdkdatacommittee.Contract.RequiredAmountOfSignatures(&_Cdkdatacommittee.CallOpts)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Cdkdatacommittee *CdkdatacommitteeCaller) VerifySignatures(opts *bind.CallOpts, signedHash [32]byte, signaturesAndAddrs []byte) error {
	var out []interface{}
	err := _Cdkdatacommittee.contract.Call(opts, &out, "verifySignatures", signedHash, signaturesAndAddrs)

	if err != nil {
		return err
	}

	return err

}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Cdkdatacommittee *CdkdatacommitteeSession) VerifySignatures(signedHash [32]byte, signaturesAndAddrs []byte) error {
	return _Cdkdatacommittee.Contract.VerifySignatures(&_Cdkdatacommittee.CallOpts, signedHash, signaturesAndAddrs)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Cdkdatacommittee *CdkdatacommitteeCallerSession) VerifySignatures(signedHash [32]byte, signaturesAndAddrs []byte) error {
	return _Cdkdatacommittee.Contract.VerifySignatures(&_Cdkdatacommittee.CallOpts, signedHash, signaturesAndAddrs)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkdatacommittee.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Cdkdatacommittee *CdkdatacommitteeSession) Initialize() (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.Initialize(&_Cdkdatacommittee.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.Initialize(&_Cdkdatacommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdkdatacommittee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cdkdatacommittee *CdkdatacommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.RenounceOwnership(&_Cdkdatacommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.RenounceOwnership(&_Cdkdatacommittee.TransactOpts)
}

// SetupCommittee is a paid mutator transaction binding the contract method 0x078fba2a.
//
// Solidity: function setupCommittee(uint256 _requiredAmountOfSignatures, string[] urls, bytes addrsBytes) returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactor) SetupCommittee(opts *bind.TransactOpts, _requiredAmountOfSignatures *big.Int, urls []string, addrsBytes []byte) (*types.Transaction, error) {
	return _Cdkdatacommittee.contract.Transact(opts, "setupCommittee", _requiredAmountOfSignatures, urls, addrsBytes)
}

// SetupCommittee is a paid mutator transaction binding the contract method 0x078fba2a.
//
// Solidity: function setupCommittee(uint256 _requiredAmountOfSignatures, string[] urls, bytes addrsBytes) returns()
func (_Cdkdatacommittee *CdkdatacommitteeSession) SetupCommittee(_requiredAmountOfSignatures *big.Int, urls []string, addrsBytes []byte) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.SetupCommittee(&_Cdkdatacommittee.TransactOpts, _requiredAmountOfSignatures, urls, addrsBytes)
}

// SetupCommittee is a paid mutator transaction binding the contract method 0x078fba2a.
//
// Solidity: function setupCommittee(uint256 _requiredAmountOfSignatures, string[] urls, bytes addrsBytes) returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactorSession) SetupCommittee(_requiredAmountOfSignatures *big.Int, urls []string, addrsBytes []byte) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.SetupCommittee(&_Cdkdatacommittee.TransactOpts, _requiredAmountOfSignatures, urls, addrsBytes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cdkdatacommittee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cdkdatacommittee *CdkdatacommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.TransferOwnership(&_Cdkdatacommittee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cdkdatacommittee *CdkdatacommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cdkdatacommittee.Contract.TransferOwnership(&_Cdkdatacommittee.TransactOpts, newOwner)
}

// CdkdatacommitteeCommitteeUpdatedIterator is returned from FilterCommitteeUpdated and is used to iterate over the raw logs and unpacked data for CommitteeUpdated events raised by the Cdkdatacommittee contract.
type CdkdatacommitteeCommitteeUpdatedIterator struct {
	Event *CdkdatacommitteeCommitteeUpdated // Event containing the contract specifics and raw log

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
func (it *CdkdatacommitteeCommitteeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkdatacommitteeCommitteeUpdated)
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
		it.Event = new(CdkdatacommitteeCommitteeUpdated)
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
func (it *CdkdatacommitteeCommitteeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkdatacommitteeCommitteeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkdatacommitteeCommitteeUpdated represents a CommitteeUpdated event raised by the Cdkdatacommittee contract.
type CdkdatacommitteeCommitteeUpdated struct {
	CommitteeHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitteeUpdated is a free log retrieval operation binding the contract event 0x831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db657.
//
// Solidity: event CommitteeUpdated(bytes32 committeeHash)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) FilterCommitteeUpdated(opts *bind.FilterOpts) (*CdkdatacommitteeCommitteeUpdatedIterator, error) {

	logs, sub, err := _Cdkdatacommittee.contract.FilterLogs(opts, "CommitteeUpdated")
	if err != nil {
		return nil, err
	}
	return &CdkdatacommitteeCommitteeUpdatedIterator{contract: _Cdkdatacommittee.contract, event: "CommitteeUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitteeUpdated is a free log subscription operation binding the contract event 0x831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db657.
//
// Solidity: event CommitteeUpdated(bytes32 committeeHash)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) WatchCommitteeUpdated(opts *bind.WatchOpts, sink chan<- *CdkdatacommitteeCommitteeUpdated) (event.Subscription, error) {

	logs, sub, err := _Cdkdatacommittee.contract.WatchLogs(opts, "CommitteeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkdatacommitteeCommitteeUpdated)
				if err := _Cdkdatacommittee.contract.UnpackLog(event, "CommitteeUpdated", log); err != nil {
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

// ParseCommitteeUpdated is a log parse operation binding the contract event 0x831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db657.
//
// Solidity: event CommitteeUpdated(bytes32 committeeHash)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) ParseCommitteeUpdated(log types.Log) (*CdkdatacommitteeCommitteeUpdated, error) {
	event := new(CdkdatacommitteeCommitteeUpdated)
	if err := _Cdkdatacommittee.contract.UnpackLog(event, "CommitteeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkdatacommitteeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cdkdatacommittee contract.
type CdkdatacommitteeInitializedIterator struct {
	Event *CdkdatacommitteeInitialized // Event containing the contract specifics and raw log

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
func (it *CdkdatacommitteeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkdatacommitteeInitialized)
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
		it.Event = new(CdkdatacommitteeInitialized)
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
func (it *CdkdatacommitteeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkdatacommitteeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkdatacommitteeInitialized represents a Initialized event raised by the Cdkdatacommittee contract.
type CdkdatacommitteeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) FilterInitialized(opts *bind.FilterOpts) (*CdkdatacommitteeInitializedIterator, error) {

	logs, sub, err := _Cdkdatacommittee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CdkdatacommitteeInitializedIterator{contract: _Cdkdatacommittee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CdkdatacommitteeInitialized) (event.Subscription, error) {

	logs, sub, err := _Cdkdatacommittee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkdatacommitteeInitialized)
				if err := _Cdkdatacommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) ParseInitialized(log types.Log) (*CdkdatacommitteeInitialized, error) {
	event := new(CdkdatacommitteeInitialized)
	if err := _Cdkdatacommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdkdatacommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cdkdatacommittee contract.
type CdkdatacommitteeOwnershipTransferredIterator struct {
	Event *CdkdatacommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CdkdatacommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdkdatacommitteeOwnershipTransferred)
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
		it.Event = new(CdkdatacommitteeOwnershipTransferred)
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
func (it *CdkdatacommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdkdatacommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdkdatacommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the Cdkdatacommittee contract.
type CdkdatacommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CdkdatacommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cdkdatacommittee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CdkdatacommitteeOwnershipTransferredIterator{contract: _Cdkdatacommittee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CdkdatacommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cdkdatacommittee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdkdatacommitteeOwnershipTransferred)
				if err := _Cdkdatacommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cdkdatacommittee *CdkdatacommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*CdkdatacommitteeOwnershipTransferred, error) {
	event := new(CdkdatacommitteeOwnershipTransferred)
	if err := _Cdkdatacommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
