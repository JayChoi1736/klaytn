// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system_contracts

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IKIP113BlsPublicKeyInfo is an auto generated low-level Go binding around an user-defined struct.
type IKIP113BlsPublicKeyInfo struct {
	PublicKey []byte
	Pop       []byte
}

// IRegistryRecord is an auto generated low-level Go binding around an user-defined struct.
type IRegistryRecord struct {
	Addr       common.Address
	Activation *big.Int
}

// IKIP113ABI is the input ABI used to generate the binding from.
const IKIP113ABI = "[{\"inputs\":[],\"name\":\"getAllBlsInfo\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"nodeIdList\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pop\",\"type\":\"bytes\"}],\"internalType\":\"structIKIP113.BlsPublicKeyInfo[]\",\"name\":\"pubkeyList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IKIP113BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IKIP113BinRuntime = ``

// IKIP113FuncSigs maps the 4-byte function signature to its string representation.
var IKIP113FuncSigs = map[string]string{
	"6968b53f": "getAllBlsInfo()",
}

// IKIP113 is an auto generated Go binding around a Klaytn contract.
type IKIP113 struct {
	IKIP113Caller     // Read-only binding to the contract
	IKIP113Transactor // Write-only binding to the contract
	IKIP113Filterer   // Log filterer for contract events
}

// IKIP113Caller is an auto generated read-only Go binding around a Klaytn contract.
type IKIP113Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP113Transactor is an auto generated write-only Go binding around a Klaytn contract.
type IKIP113Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP113Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IKIP113Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKIP113Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IKIP113Session struct {
	Contract     *IKIP113          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKIP113CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IKIP113CallerSession struct {
	Contract *IKIP113Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IKIP113TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IKIP113TransactorSession struct {
	Contract     *IKIP113Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IKIP113Raw is an auto generated low-level Go binding around a Klaytn contract.
type IKIP113Raw struct {
	Contract *IKIP113 // Generic contract binding to access the raw methods on
}

// IKIP113CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IKIP113CallerRaw struct {
	Contract *IKIP113Caller // Generic read-only contract binding to access the raw methods on
}

// IKIP113TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IKIP113TransactorRaw struct {
	Contract *IKIP113Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIKIP113 creates a new instance of IKIP113, bound to a specific deployed contract.
func NewIKIP113(address common.Address, backend bind.ContractBackend) (*IKIP113, error) {
	contract, err := bindIKIP113(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKIP113{IKIP113Caller: IKIP113Caller{contract: contract}, IKIP113Transactor: IKIP113Transactor{contract: contract}, IKIP113Filterer: IKIP113Filterer{contract: contract}}, nil
}

// NewIKIP113Caller creates a new read-only instance of IKIP113, bound to a specific deployed contract.
func NewIKIP113Caller(address common.Address, caller bind.ContractCaller) (*IKIP113Caller, error) {
	contract, err := bindIKIP113(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP113Caller{contract: contract}, nil
}

// NewIKIP113Transactor creates a new write-only instance of IKIP113, bound to a specific deployed contract.
func NewIKIP113Transactor(address common.Address, transactor bind.ContractTransactor) (*IKIP113Transactor, error) {
	contract, err := bindIKIP113(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKIP113Transactor{contract: contract}, nil
}

// NewIKIP113Filterer creates a new log filterer instance of IKIP113, bound to a specific deployed contract.
func NewIKIP113Filterer(address common.Address, filterer bind.ContractFilterer) (*IKIP113Filterer, error) {
	contract, err := bindIKIP113(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKIP113Filterer{contract: contract}, nil
}

// bindIKIP113 binds a generic wrapper to an already deployed contract.
func bindIKIP113(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKIP113ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP113 *IKIP113Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP113.Contract.IKIP113Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP113 *IKIP113Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP113.Contract.IKIP113Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP113 *IKIP113Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP113.Contract.IKIP113Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKIP113 *IKIP113CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IKIP113.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKIP113 *IKIP113TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKIP113.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKIP113 *IKIP113TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKIP113.Contract.contract.Transact(opts, method, params...)
}

// GetAllBlsInfo is a free data retrieval call binding the contract method 0x6968b53f.
//
// Solidity: function getAllBlsInfo() view returns(address[] nodeIdList, (bytes,bytes)[] pubkeyList)
func (_IKIP113 *IKIP113Caller) GetAllBlsInfo(opts *bind.CallOpts) (struct {
	NodeIdList []common.Address
	PubkeyList []IKIP113BlsPublicKeyInfo
}, error) {
	ret := new(struct {
		NodeIdList []common.Address
		PubkeyList []IKIP113BlsPublicKeyInfo
	})
	out := ret
	err := _IKIP113.contract.Call(opts, out, "getAllBlsInfo")
	return *ret, err
}

// GetAllBlsInfo is a free data retrieval call binding the contract method 0x6968b53f.
//
// Solidity: function getAllBlsInfo() view returns(address[] nodeIdList, (bytes,bytes)[] pubkeyList)
func (_IKIP113 *IKIP113Session) GetAllBlsInfo() (struct {
	NodeIdList []common.Address
	PubkeyList []IKIP113BlsPublicKeyInfo
}, error) {
	return _IKIP113.Contract.GetAllBlsInfo(&_IKIP113.CallOpts)
}

// GetAllBlsInfo is a free data retrieval call binding the contract method 0x6968b53f.
//
// Solidity: function getAllBlsInfo() view returns(address[] nodeIdList, (bytes,bytes)[] pubkeyList)
func (_IKIP113 *IKIP113CallerSession) GetAllBlsInfo() (struct {
	NodeIdList []common.Address
	PubkeyList []IKIP113BlsPublicKeyInfo
}, error) {
	return _IKIP113.Contract.GetAllBlsInfo(&_IKIP113.CallOpts)
}

// IRegistryABI is the input ABI used to generate the binding from.
const IRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getActiveAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IRegistryBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IRegistryBinRuntime = ``

// IRegistryFuncSigs maps the 4-byte function signature to its string representation.
var IRegistryFuncSigs = map[string]string{
	"e2693e3f": "getActiveAddr(string)",
	"fb825e5f": "getAllNames()",
	"78d573a2": "getAllRecords(string)",
	"4622ab03": "names(uint256)",
	"8da5cb5b": "owner()",
	"3b51650d": "records(string,uint256)",
	"d393c871": "register(string,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// IRegistry is an auto generated Go binding around a Klaytn contract.
type IRegistry struct {
	IRegistryCaller     // Read-only binding to the contract
	IRegistryTransactor // Write-only binding to the contract
	IRegistryFilterer   // Log filterer for contract events
}

// IRegistryCaller is an auto generated read-only Go binding around a Klaytn contract.
type IRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrySession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IRegistrySession struct {
	Contract     *IRegistry        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRegistryCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IRegistryCallerSession struct {
	Contract *IRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IRegistryTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IRegistryTransactorSession struct {
	Contract     *IRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRegistryRaw is an auto generated low-level Go binding around a Klaytn contract.
type IRegistryRaw struct {
	Contract *IRegistry // Generic contract binding to access the raw methods on
}

// IRegistryCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IRegistryCallerRaw struct {
	Contract *IRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IRegistryTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IRegistryTransactorRaw struct {
	Contract *IRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRegistry creates a new instance of IRegistry, bound to a specific deployed contract.
func NewIRegistry(address common.Address, backend bind.ContractBackend) (*IRegistry, error) {
	contract, err := bindIRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRegistry{IRegistryCaller: IRegistryCaller{contract: contract}, IRegistryTransactor: IRegistryTransactor{contract: contract}, IRegistryFilterer: IRegistryFilterer{contract: contract}}, nil
}

// NewIRegistryCaller creates a new read-only instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryCaller(address common.Address, caller bind.ContractCaller) (*IRegistryCaller, error) {
	contract, err := bindIRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryCaller{contract: contract}, nil
}

// NewIRegistryTransactor creates a new write-only instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IRegistryTransactor, error) {
	contract, err := bindIRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryTransactor{contract: contract}, nil
}

// NewIRegistryFilterer creates a new log filterer instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IRegistryFilterer, error) {
	contract, err := bindIRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRegistryFilterer{contract: contract}, nil
}

// bindIRegistry binds a generic wrapper to an already deployed contract.
func bindIRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistry *IRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistry.Contract.IRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistry *IRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistry.Contract.IRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistry *IRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistry.Contract.IRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistry *IRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistry *IRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistry *IRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_IRegistry *IRegistryCaller) GetAllNames(opts *bind.CallOpts) ([]string, error) {

	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "getAllNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_IRegistry *IRegistrySession) GetAllNames() ([]string, error) {
	return _IRegistry.Contract.GetAllNames(&_IRegistry.CallOpts)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_IRegistry *IRegistryCallerSession) GetAllNames() ([]string, error) {
	return _IRegistry.Contract.GetAllNames(&_IRegistry.CallOpts)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_IRegistry *IRegistryCaller) GetAllRecords(opts *bind.CallOpts, name string) ([]IRegistryRecord, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "getAllRecords", name)

	if err != nil {
		return *new([]IRegistryRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRegistryRecord)).(*[]IRegistryRecord)

	return out0, err

}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_IRegistry *IRegistrySession) GetAllRecords(name string) ([]IRegistryRecord, error) {
	return _IRegistry.Contract.GetAllRecords(&_IRegistry.CallOpts, name)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_IRegistry *IRegistryCallerSession) GetAllRecords(name string) ([]IRegistryRecord, error) {
	return _IRegistry.Contract.GetAllRecords(&_IRegistry.CallOpts, name)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_IRegistry *IRegistryCaller) Names(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_IRegistry *IRegistrySession) Names(arg0 *big.Int) (string, error) {
	return _IRegistry.Contract.Names(&_IRegistry.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_IRegistry *IRegistryCallerSession) Names(arg0 *big.Int) (string, error) {
	return _IRegistry.Contract.Names(&_IRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IRegistry *IRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IRegistry *IRegistrySession) Owner() (common.Address, error) {
	return _IRegistry.Contract.Owner(&_IRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IRegistry *IRegistryCallerSession) Owner() (common.Address, error) {
	return _IRegistry.Contract.Owner(&_IRegistry.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_IRegistry *IRegistryCaller) Records(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {

	var out []interface{}
	err := _IRegistry.contract.Call(opts, &out, "records", arg0, arg1)

	outstruct := new(struct {
		Addr       common.Address
		Activation *big.Int
	})

	outstruct.Addr = out[0].(common.Address)
	outstruct.Activation = out[1].(*big.Int)
	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_IRegistry *IRegistrySession) Records(arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	return _IRegistry.Contract.Records(&_IRegistry.CallOpts, arg0, arg1)
}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_IRegistry *IRegistryCallerSession) Records(arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	return _IRegistry.Contract.Records(&_IRegistry.CallOpts, arg0, arg1)
}

// GetActiveAddr is a paid mutator transaction binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) returns(address)
func (_IRegistry *IRegistryTransactor) GetActiveAddr(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _IRegistry.contract.Transact(opts, "getActiveAddr", name)
}

// GetActiveAddr is a paid mutator transaction binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) returns(address)
func (_IRegistry *IRegistrySession) GetActiveAddr(name string) (*types.Transaction, error) {
	return _IRegistry.Contract.GetActiveAddr(&_IRegistry.TransactOpts, name)
}

// GetActiveAddr is a paid mutator transaction binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) returns(address)
func (_IRegistry *IRegistryTransactorSession) GetActiveAddr(name string) (*types.Transaction, error) {
	return _IRegistry.Contract.GetActiveAddr(&_IRegistry.TransactOpts, name)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_IRegistry *IRegistryTransactor) Register(opts *bind.TransactOpts, name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _IRegistry.contract.Transact(opts, "register", name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_IRegistry *IRegistrySession) Register(name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _IRegistry.Contract.Register(&_IRegistry.TransactOpts, name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_IRegistry *IRegistryTransactorSession) Register(name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _IRegistry.Contract.Register(&_IRegistry.TransactOpts, name, addr, activation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IRegistry *IRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IRegistry *IRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IRegistry.Contract.TransferOwnership(&_IRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IRegistry *IRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IRegistry.Contract.TransferOwnership(&_IRegistry.TransactOpts, newOwner)
}

// IRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IRegistry contract.
type IRegistryOwnershipTransferredIterator struct {
	Event *IRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistryOwnershipTransferred)
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
		it.Event = new(IRegistryOwnershipTransferred)
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
func (it *IRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the IRegistry contract.
type IRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IRegistry *IRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IRegistryOwnershipTransferredIterator{contract: _IRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IRegistry *IRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistryOwnershipTransferred)
				if err := _IRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IRegistry *IRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*IRegistryOwnershipTransferred, error) {
	event := new(IRegistryOwnershipTransferred)
	if err := _IRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IRegistry contract.
type IRegistryRegisteredIterator struct {
	Event *IRegistryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistryRegistered)
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
		it.Event = new(IRegistryRegistered)
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
func (it *IRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistryRegistered represents a Registered event raised by the IRegistry contract.
type IRegistryRegistered struct {
	Name       string
	Addr       common.Address
	Activation *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_IRegistry *IRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, addr []common.Address, activation []*big.Int) (*IRegistryRegisteredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var activationRule []interface{}
	for _, activationItem := range activation {
		activationRule = append(activationRule, activationItem)
	}

	logs, sub, err := _IRegistry.contract.FilterLogs(opts, "Registered", addrRule, activationRule)
	if err != nil {
		return nil, err
	}
	return &IRegistryRegisteredIterator{contract: _IRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_IRegistry *IRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IRegistryRegistered, addr []common.Address, activation []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var activationRule []interface{}
	for _, activationItem := range activation {
		activationRule = append(activationRule, activationItem)
	}

	logs, sub, err := _IRegistry.contract.WatchLogs(opts, "Registered", addrRule, activationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistryRegistered)
				if err := _IRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_IRegistry *IRegistryFilterer) ParseRegistered(log types.Log) (*IRegistryRegistered, error) {
	event := new(IRegistryRegistered)
	if err := _IRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getActiveAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e2693e3f": "getActiveAddr(string)",
		"fb825e5f": "getAllNames()",
		"78d573a2": "getAllRecords(string)",
		"4622ab03": "names(uint256)",
		"8da5cb5b": "owner()",
		"3b51650d": "records(string,uint256)",
		"d393c871": "register(string,address,uint256)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610d52806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063d393c8711161005b578063d393c87114610129578063e2693e3f1461013e578063f2fde38b14610151578063fb825e5f1461016457600080fd5b80633b51650d1461008d5780634622ab03146100c457806378d573a2146100e45780638da5cb5b14610104575b600080fd5b6100a061009b366004610a12565b610179565b604080516001600160a01b0390931683526020830191909152015b60405180910390f35b6100d76100d2366004610a57565b6101ce565b6040516100bb9190610acc565b6100f76100f2366004610ae6565b61027a565b6040516100bb9190610b23565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100bb565b61013c610137366004610b97565b61030d565b005b61011161014c366004610ae6565b61062f565b61013c61015f366004610bee565b610726565b61016c6107fd565b6040516100bb9190610c09565b815160208184018101805160008252928201918501919091209190528054829081106101a457600080fd5b6000918252602090912060029091020180546001909101546001600160a01b039091169250905082565b600181815481106101de57600080fd5b9060005260206000200160009150905080546101f990610c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461022590610c6b565b80156102725780601f1061024757610100808354040283529160200191610272565b820191906000526020600020905b81548152906001019060200180831161025557829003601f168201915b505050505081565b606060008260405161028c9190610ca6565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610302576000848152602090819020604080518082019091526002850290910180546001600160a01b031682526001908101548284015290835290920191016102ba565b505050509050919050565b6002546001600160a01b031633146103585760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064015b60405180910390fd5b8260008160405160200161036c9190610ca6565b60405160208183030381529060405290508051600014156103be5760405162461bcd60e51b815260206004820152600c60248201526b456d70747920737472696e6760a01b604482015260640161034f565b4383116104175760405162461bcd60e51b815260206004820152602160248201527f43616e277420726567697374657220636f6e74726163742066726f6d207061736044820152601d60fa1b606482015260840161034f565b600080866040516104289190610ca6565b908152604051908190036020019020549050806104f7576001805480820182556000919091528651610481917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019060208901906108d6565b506000866040516104929190610ca6565b90815260408051602092819003830181208183019092526001600160a01b0388811682528382018881528354600180820186556000958652959094209251600290940290920180546001600160a01b03191693909116929092178255519101556105e5565b600080876040516105089190610ca6565b908152604051908190036020019020610522600184610cd8565b8154811061053257610532610cef565b90600052602060002090600202019050438160010154116105c25760008760405161055d9190610ca6565b90815260408051602092819003830181208183019092526001600160a01b0389811682528382018981528354600180820186556000958652959094209251600290940290920180546001600160a01b03191693909116929092178255519101556105e3565b80546001600160a01b0319166001600160a01b038716178155600181018590555b505b83856001600160a01b03167f142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f308860405161061f9190610acc565b60405180910390a3505050505050565b6000806000836040516106429190610ca6565b908152604051908190036020019020549050805b801561071c574360008560405161066d9190610ca6565b908152604051908190036020019020610687600184610cd8565b8154811061069757610697610cef565b9060005260206000209060020201600101541161070a576000846040516106be9190610ca6565b9081526040519081900360200190206106d8600183610cd8565b815481106106e8576106e8610cef565b60009182526020909120600290910201546001600160a01b0316949350505050565b8061071481610d05565b915050610656565b5060009392505050565b6002546001600160a01b0316331461076c5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b604482015260640161034f565b6001600160a01b0381166107b15760405162461bcd60e51b815260206004820152600c60248201526b5a65726f206164647265737360a01b604482015260640161034f565b600280546001600160a01b0319166001600160a01b03831690811790915560405133907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156108cd57838290600052602060002001805461084090610c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461086c90610c6b565b80156108b95780601f1061088e576101008083540402835291602001916108b9565b820191906000526020600020905b81548152906001019060200180831161089c57829003601f168201915b505050505081526020019060010190610821565b50505050905090565b8280546108e290610c6b565b90600052602060002090601f016020900481019282610904576000855561094a565b82601f1061091d57805160ff191683800117855561094a565b8280016001018555821561094a579182015b8281111561094a57825182559160200191906001019061092f565b5061095692915061095a565b5090565b5b80821115610956576000815560010161095b565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261099657600080fd5b813567ffffffffffffffff808211156109b1576109b161096f565b604051601f8301601f19908116603f011681019082821181831017156109d9576109d961096f565b816040528381528660208588010111156109f257600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610a2557600080fd5b823567ffffffffffffffff811115610a3c57600080fd5b610a4885828601610985565b95602094909401359450505050565b600060208284031215610a6957600080fd5b5035919050565b60005b83811015610a8b578181015183820152602001610a73565b83811115610a9a576000848401525b50505050565b60008151808452610ab8816020860160208601610a70565b601f01601f19169290920160200192915050565b602081526000610adf6020830184610aa0565b9392505050565b600060208284031215610af857600080fd5b813567ffffffffffffffff811115610b0f57600080fd5b610b1b84828501610985565b949350505050565b602080825282518282018190526000919060409081850190868401855b82811015610b6e57815180516001600160a01b03168552860151868501529284019290850190600101610b40565b5091979650505050505050565b80356001600160a01b0381168114610b9257600080fd5b919050565b600080600060608486031215610bac57600080fd5b833567ffffffffffffffff811115610bc357600080fd5b610bcf86828701610985565b935050610bde60208501610b7b565b9150604084013590509250925092565b600060208284031215610c0057600080fd5b610adf82610b7b565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610c5e57603f19888603018452610c4c858351610aa0565b94509285019290850190600101610c30565b5092979650505050505050565b600181811c90821680610c7f57607f821691505b60208210811415610ca057634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610cb8818460208701610a70565b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b600082821015610cea57610cea610cc2565b500390565b634e487b7160e01b600052603260045260246000fd5b600081610d1457610d14610cc2565b50600019019056fea2646970667358221220e198b2f0df398df674d3d7591b7f25729f37b0f94ebb38c50269232f27c2d72964736f6c634300080a0033",
}

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getActiveAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const RegistryBinRuntime = `608060405234801561001057600080fd5b50600436106100885760003560e01c8063d393c8711161005b578063d393c87114610129578063e2693e3f1461013e578063f2fde38b14610151578063fb825e5f1461016457600080fd5b80633b51650d1461008d5780634622ab03146100c457806378d573a2146100e45780638da5cb5b14610104575b600080fd5b6100a061009b366004610a12565b610179565b604080516001600160a01b0390931683526020830191909152015b60405180910390f35b6100d76100d2366004610a57565b6101ce565b6040516100bb9190610acc565b6100f76100f2366004610ae6565b61027a565b6040516100bb9190610b23565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100bb565b61013c610137366004610b97565b61030d565b005b61011161014c366004610ae6565b61062f565b61013c61015f366004610bee565b610726565b61016c6107fd565b6040516100bb9190610c09565b815160208184018101805160008252928201918501919091209190528054829081106101a457600080fd5b6000918252602090912060029091020180546001909101546001600160a01b039091169250905082565b600181815481106101de57600080fd5b9060005260206000200160009150905080546101f990610c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461022590610c6b565b80156102725780601f1061024757610100808354040283529160200191610272565b820191906000526020600020905b81548152906001019060200180831161025557829003601f168201915b505050505081565b606060008260405161028c9190610ca6565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610302576000848152602090819020604080518082019091526002850290910180546001600160a01b031682526001908101548284015290835290920191016102ba565b505050509050919050565b6002546001600160a01b031633146103585760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064015b60405180910390fd5b8260008160405160200161036c9190610ca6565b60405160208183030381529060405290508051600014156103be5760405162461bcd60e51b815260206004820152600c60248201526b456d70747920737472696e6760a01b604482015260640161034f565b4383116104175760405162461bcd60e51b815260206004820152602160248201527f43616e277420726567697374657220636f6e74726163742066726f6d207061736044820152601d60fa1b606482015260840161034f565b600080866040516104289190610ca6565b908152604051908190036020019020549050806104f7576001805480820182556000919091528651610481917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019060208901906108d6565b506000866040516104929190610ca6565b90815260408051602092819003830181208183019092526001600160a01b0388811682528382018881528354600180820186556000958652959094209251600290940290920180546001600160a01b03191693909116929092178255519101556105e5565b600080876040516105089190610ca6565b908152604051908190036020019020610522600184610cd8565b8154811061053257610532610cef565b90600052602060002090600202019050438160010154116105c25760008760405161055d9190610ca6565b90815260408051602092819003830181208183019092526001600160a01b0389811682528382018981528354600180820186556000958652959094209251600290940290920180546001600160a01b03191693909116929092178255519101556105e3565b80546001600160a01b0319166001600160a01b038716178155600181018590555b505b83856001600160a01b03167f142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f308860405161061f9190610acc565b60405180910390a3505050505050565b6000806000836040516106429190610ca6565b908152604051908190036020019020549050805b801561071c574360008560405161066d9190610ca6565b908152604051908190036020019020610687600184610cd8565b8154811061069757610697610cef565b9060005260206000209060020201600101541161070a576000846040516106be9190610ca6565b9081526040519081900360200190206106d8600183610cd8565b815481106106e8576106e8610cef565b60009182526020909120600290910201546001600160a01b0316949350505050565b8061071481610d05565b915050610656565b5060009392505050565b6002546001600160a01b0316331461076c5760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b604482015260640161034f565b6001600160a01b0381166107b15760405162461bcd60e51b815260206004820152600c60248201526b5a65726f206164647265737360a01b604482015260640161034f565b600280546001600160a01b0319166001600160a01b03831690811790915560405133907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156108cd57838290600052602060002001805461084090610c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461086c90610c6b565b80156108b95780601f1061088e576101008083540402835291602001916108b9565b820191906000526020600020905b81548152906001019060200180831161089c57829003601f168201915b505050505081526020019060010190610821565b50505050905090565b8280546108e290610c6b565b90600052602060002090601f016020900481019282610904576000855561094a565b82601f1061091d57805160ff191683800117855561094a565b8280016001018555821561094a579182015b8281111561094a57825182559160200191906001019061092f565b5061095692915061095a565b5090565b5b80821115610956576000815560010161095b565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261099657600080fd5b813567ffffffffffffffff808211156109b1576109b161096f565b604051601f8301601f19908116603f011681019082821181831017156109d9576109d961096f565b816040528381528660208588010111156109f257600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610a2557600080fd5b823567ffffffffffffffff811115610a3c57600080fd5b610a4885828601610985565b95602094909401359450505050565b600060208284031215610a6957600080fd5b5035919050565b60005b83811015610a8b578181015183820152602001610a73565b83811115610a9a576000848401525b50505050565b60008151808452610ab8816020860160208601610a70565b601f01601f19169290920160200192915050565b602081526000610adf6020830184610aa0565b9392505050565b600060208284031215610af857600080fd5b813567ffffffffffffffff811115610b0f57600080fd5b610b1b84828501610985565b949350505050565b602080825282518282018190526000919060409081850190868401855b82811015610b6e57815180516001600160a01b03168552860151868501529284019290850190600101610b40565b5091979650505050505050565b80356001600160a01b0381168114610b9257600080fd5b919050565b600080600060608486031215610bac57600080fd5b833567ffffffffffffffff811115610bc357600080fd5b610bcf86828701610985565b935050610bde60208501610b7b565b9150604084013590509250925092565b600060208284031215610c0057600080fd5b610adf82610b7b565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610c5e57603f19888603018452610c4c858351610aa0565b94509285019290850190600101610c30565b5092979650505050505050565b600181811c90821680610c7f57607f821691505b60208210811415610ca057634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610cb8818460208701610a70565b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b600082821015610cea57610cea610cc2565b500390565b634e487b7160e01b600052603260045260246000fd5b600081610d1457610d14610cc2565b50600019019056fea2646970667358221220e198b2f0df398df674d3d7591b7f25729f37b0f94ebb38c50269232f27c2d72964736f6c634300080a0033`

// RegistryFuncSigs maps the 4-byte function signature to its string representation.
var RegistryFuncSigs = map[string]string{
	"e2693e3f": "getActiveAddr(string)",
	"fb825e5f": "getAllNames()",
	"78d573a2": "getAllRecords(string)",
	"4622ab03": "names(uint256)",
	"8da5cb5b": "owner()",
	"3b51650d": "records(string,uint256)",
	"d393c871": "register(string,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405234801561001057600080fd5b50610db9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063d393c8711161005b578063d393c87114610129578063e2693e3f1461013e578063f2fde38b14610151578063fb825e5f1461016457600080fd5b80633b51650d1461008d5780634622ab03146100c457806378d573a2146100e45780638da5cb5b14610104575b600080fd5b6100a061009b366004610975565b610179565b604080516001600160a01b0390931683526020830191909152015b60405180910390f35b6100d76100d23660046109ba565b6101ce565b6040516100bb9190610a23565b6100f76100f2366004610a3d565b61027a565b6040516100bb9190610a7a565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100bb565b61013c610137366004610aee565b61030d565b005b61011161014c366004610a3d565b61062b565b61013c61015f366004610b45565b610722565b61016c6107f9565b6040516100bb9190610b60565b815160208184018101805160008252928201918501919091209190528054829081106101a457600080fd5b6000918252602090912060029091020180546001909101546001600160a01b039091169250905082565b600181815481106101de57600080fd5b9060005260206000200160009150905080546101f990610bc2565b80601f016020809104026020016040519081016040528092919081815260200182805461022590610bc2565b80156102725780601f1061024757610100808354040283529160200191610272565b820191906000526020600020905b81548152906001019060200180831161025557829003601f168201915b505050505081565b606060008260405161028c9190610bfc565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610302576000848152602090819020604080518082019091526002850290910180546001600160a01b031682526001908101548284015290835290920191016102ba565b505050509050919050565b6002546001600160a01b031633146103585760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b60448201526064015b60405180910390fd5b8260008160405160200161036c9190610bfc565b604051602081830303815290604052905080516000036103bd5760405162461bcd60e51b815260206004820152600c60248201526b456d70747920737472696e6760a01b604482015260640161034f565b4383116104165760405162461bcd60e51b815260206004820152602160248201527f43616e277420726567697374657220636f6e74726163742066726f6d207061736044820152601d60fa1b606482015260840161034f565b600080866040516104279190610bfc565b90815260405190819003602001902054905060008190036104f3576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60161047d8782610c67565b5060008660405161048e9190610bfc565b90815260408051602092819003830181208183019092526001600160a01b0388811682528382018881528354600180820186556000958652959094209251600290940290920180546001600160a01b03191693909116929092178255519101556105e1565b600080876040516105049190610bfc565b90815260405190819003602001902061051e600184610d3d565b8154811061052e5761052e610d56565b90600052602060002090600202019050438160010154116105be576000876040516105599190610bfc565b90815260408051602092819003830181208183019092526001600160a01b0389811682528382018981528354600180820186556000958652959094209251600290940290920180546001600160a01b03191693909116929092178255519101556105df565b80546001600160a01b0319166001600160a01b038716178155600181018590555b505b83856001600160a01b03167f142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f308860405161061b9190610a23565b60405180910390a3505050505050565b60008060008360405161063e9190610bfc565b908152604051908190036020019020549050805b801561071857436000856040516106699190610bfc565b908152604051908190036020019020610683600184610d3d565b8154811061069357610693610d56565b90600052602060002090600202016001015411610706576000846040516106ba9190610bfc565b9081526040519081900360200190206106d4600183610d3d565b815481106106e4576106e4610d56565b60009182526020909120600290910201546001600160a01b0316949350505050565b8061071081610d6c565b915050610652565b5060009392505050565b6002546001600160a01b031633146107685760405162461bcd60e51b81526020600482015260096024820152682737ba1037bbb732b960b91b604482015260640161034f565b6001600160a01b0381166107ad5760405162461bcd60e51b815260206004820152600c60248201526b5a65726f206164647265737360a01b604482015260640161034f565b600280546001600160a01b0319166001600160a01b03831690811790915560405133907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156108c957838290600052602060002001805461083c90610bc2565b80601f016020809104026020016040519081016040528092919081815260200182805461086890610bc2565b80156108b55780601f1061088a576101008083540402835291602001916108b5565b820191906000526020600020905b81548152906001019060200180831161089857829003601f168201915b50505050508152602001906001019061081d565b50505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126108f957600080fd5b813567ffffffffffffffff80821115610914576109146108d2565b604051601f8301601f19908116603f0116810190828211818310171561093c5761093c6108d2565b8160405283815286602085880101111561095557600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561098857600080fd5b823567ffffffffffffffff81111561099f57600080fd5b6109ab858286016108e8565b95602094909401359450505050565b6000602082840312156109cc57600080fd5b5035919050565b60005b838110156109ee5781810151838201526020016109d6565b50506000910152565b60008151808452610a0f8160208601602086016109d3565b601f01601f19169290920160200192915050565b602081526000610a3660208301846109f7565b9392505050565b600060208284031215610a4f57600080fd5b813567ffffffffffffffff811115610a6657600080fd5b610a72848285016108e8565b949350505050565b602080825282518282018190526000919060409081850190868401855b82811015610ac557815180516001600160a01b03168552860151868501529284019290850190600101610a97565b5091979650505050505050565b80356001600160a01b0381168114610ae957600080fd5b919050565b600080600060608486031215610b0357600080fd5b833567ffffffffffffffff811115610b1a57600080fd5b610b26868287016108e8565b935050610b3560208501610ad2565b9150604084013590509250925092565b600060208284031215610b5757600080fd5b610a3682610ad2565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610bb557603f19888603018452610ba38583516109f7565b94509285019290850190600101610b87565b5092979650505050505050565b600181811c90821680610bd657607f821691505b602082108103610bf657634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610c0e8184602087016109d3565b9190910192915050565b601f821115610c6257600081815260208120601f850160051c81016020861015610c3f5750805b601f850160051c820191505b81811015610c5e57828155600101610c4b565b5050505b505050565b815167ffffffffffffffff811115610c8157610c816108d2565b610c9581610c8f8454610bc2565b84610c18565b602080601f831160018114610cca5760008415610cb25750858301515b600019600386901b1c1916600185901b178555610c5e565b600085815260208120601f198616915b82811015610cf957888601518255948401946001909101908401610cda565b5085821015610d175787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b81810381811115610d5057610d50610d27565b92915050565b634e487b7160e01b600052603260045260246000fd5b600081610d7b57610d7b610d27565b50600019019056fea2646970667358221220dec0ab0bd90558df37ff0a87cf78fb70b20ae65ac7a4292a8b2a7682ec1fd45d64736f6c63430008110033"

// DeployRegistry deploys a new Klaytn contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around a Klaytn contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around a Klaytn contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around a Klaytn contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around a Klaytn contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetActiveAddr is a free data retrieval call binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) view returns(address)
func (_Registry *RegistryCaller) GetActiveAddr(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getActiveAddr", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetActiveAddr is a free data retrieval call binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) view returns(address)
func (_Registry *RegistrySession) GetActiveAddr(name string) (common.Address, error) {
	return _Registry.Contract.GetActiveAddr(&_Registry.CallOpts, name)
}

// GetActiveAddr is a free data retrieval call binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) view returns(address)
func (_Registry *RegistryCallerSession) GetActiveAddr(name string) (common.Address, error) {
	return _Registry.Contract.GetActiveAddr(&_Registry.CallOpts, name)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_Registry *RegistryCaller) GetAllNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_Registry *RegistrySession) GetAllNames() ([]string, error) {
	return _Registry.Contract.GetAllNames(&_Registry.CallOpts)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_Registry *RegistryCallerSession) GetAllNames() ([]string, error) {
	return _Registry.Contract.GetAllNames(&_Registry.CallOpts)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_Registry *RegistryCaller) GetAllRecords(opts *bind.CallOpts, name string) ([]IRegistryRecord, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllRecords", name)

	if err != nil {
		return *new([]IRegistryRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRegistryRecord)).(*[]IRegistryRecord)

	return out0, err
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_Registry *RegistrySession) GetAllRecords(name string) ([]IRegistryRecord, error) {
	return _Registry.Contract.GetAllRecords(&_Registry.CallOpts, name)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_Registry *RegistryCallerSession) GetAllRecords(name string) ([]IRegistryRecord, error) {
	return _Registry.Contract.GetAllRecords(&_Registry.CallOpts, name)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_Registry *RegistryCaller) Names(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_Registry *RegistrySession) Names(arg0 *big.Int) (string, error) {
	return _Registry.Contract.Names(&_Registry.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_Registry *RegistryCallerSession) Names(arg0 *big.Int) (string, error) {
	return _Registry.Contract.Names(&_Registry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_Registry *RegistryCaller) Records(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "records", arg0, arg1)

	outstruct := new(struct {
		Addr       common.Address
		Activation *big.Int
	})

	outstruct.Addr = out[0].(common.Address)
	outstruct.Activation = out[1].(*big.Int)
	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_Registry *RegistrySession) Records(arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	return _Registry.Contract.Records(&_Registry.CallOpts, arg0, arg1)
}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_Registry *RegistryCallerSession) Records(arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	return _Registry.Contract.Records(&_Registry.CallOpts, arg0, arg1)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_Registry *RegistrySession) Register(name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_Registry *RegistryTransactorSession) Register(name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, name, addr, activation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Registry contract.
type RegistryRegisteredIterator struct {
	Event *RegistryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegistered)
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
		it.Event = new(RegistryRegistered)
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
func (it *RegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegistered represents a Registered event raised by the Registry contract.
type RegistryRegistered struct {
	Name       string
	Addr       common.Address
	Activation *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_Registry *RegistryFilterer) FilterRegistered(opts *bind.FilterOpts, addr []common.Address, activation []*big.Int) (*RegistryRegisteredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var activationRule []interface{}
	for _, activationItem := range activation {
		activationRule = append(activationRule, activationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Registered", addrRule, activationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIterator{contract: _Registry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_Registry *RegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistryRegistered, addr []common.Address, activation []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var activationRule []interface{}
	for _, activationItem := range activation {
		activationRule = append(activationRule, activationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Registered", addrRule, activationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegistered)
				if err := _Registry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_Registry *RegistryFilterer) ParseRegistered(log types.Log) (*RegistryRegistered, error) {
	event := new(RegistryRegistered)
	if err := _Registry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryMockMetaData contains all meta data concerning the RegistryMock contract.
var RegistryMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getActiveAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e2693e3f": "getActiveAddr(string)",
		"fb825e5f": "getAllNames()",
		"78d573a2": "getAllRecords(string)",
		"4622ab03": "names(uint256)",
		"8da5cb5b": "owner()",
		"3b51650d": "records(string,uint256)",
		"d393c871": "register(string,address,uint256)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506109c5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063d393c8711161005b578063d393c87114610129578063e2693e3f1461013e578063f2fde38b14610151578063fb825e5f1461018157600080fd5b80633b51650d1461008d5780634622ab03146100c457806378d573a2146100e45780638da5cb5b14610104575b600080fd5b6100a061009b3660046106aa565b610196565b604080516001600160a01b0390931683526020830191909152015b60405180910390f35b6100d76100d23660046106ef565b6101eb565b6040516100bb9190610764565b6100f76100f236600461077e565b610297565b6040516100bb91906107bb565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100bb565b61013c61013736600461082f565b61032a565b005b61011161014c36600461077e565b610401565b61013c61015f366004610886565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b610189610495565b6040516100bb91906108a1565b815160208184018101805160008252928201918501919091209190528054829081106101c157600080fd5b6000918252602090912060029091020180546001909101546001600160a01b039091169250905082565b600181815481106101fb57600080fd5b90600052602060002001600091509050805461021690610903565b80601f016020809104026020016040519081016040528092919081815260200182805461024290610903565b801561028f5780601f106102645761010080835404028352916020019161028f565b820191906000526020600020905b81548152906001019060200180831161027257829003601f168201915b505050505081565b60606000826040516102a99190610938565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561031f576000848152602090819020604080518082019091526002850290910180546001600160a01b031682526001908101548284015290835290920191016102d7565b505050509050919050565b60008360405161033a9190610938565b90815260405190819003602001902054610392576001805480820182556000919091528351610390917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60190602086019061056e565b505b6000836040516103a29190610938565b90815260408051602092819003830181208183019092526001600160a01b039485168152828101938452815460018082018455600093845293909220905160029092020180546001600160a01b03191691909416178355905191015550565b6000806000836040516104149190610938565b908152604051908190036020019020549050806104345750600092915050565b6000836040516104449190610938565b90815260405190819003602001902061045e600183610954565b8154811061046e5761046e610979565b60009182526020909120600290910201546001600160a01b03169392505050565b50919050565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156105655783829060005260206000200180546104d890610903565b80601f016020809104026020016040519081016040528092919081815260200182805461050490610903565b80156105515780601f1061052657610100808354040283529160200191610551565b820191906000526020600020905b81548152906001019060200180831161053457829003601f168201915b5050505050815260200190600101906104b9565b50505050905090565b82805461057a90610903565b90600052602060002090601f01602090048101928261059c57600085556105e2565b82601f106105b557805160ff19168380011785556105e2565b828001600101855582156105e2579182015b828111156105e25782518255916020019190600101906105c7565b506105ee9291506105f2565b5090565b5b808211156105ee57600081556001016105f3565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261062e57600080fd5b813567ffffffffffffffff8082111561064957610649610607565b604051601f8301601f19908116603f0116810190828211818310171561067157610671610607565b8160405283815286602085880101111561068a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156106bd57600080fd5b823567ffffffffffffffff8111156106d457600080fd5b6106e08582860161061d565b95602094909401359450505050565b60006020828403121561070157600080fd5b5035919050565b60005b8381101561072357818101518382015260200161070b565b83811115610732576000848401525b50505050565b60008151808452610750816020860160208601610708565b601f01601f19169290920160200192915050565b6020815260006107776020830184610738565b9392505050565b60006020828403121561079057600080fd5b813567ffffffffffffffff8111156107a757600080fd5b6107b38482850161061d565b949350505050565b602080825282518282018190526000919060409081850190868401855b8281101561080657815180516001600160a01b031685528601518685015292840192908501906001016107d8565b5091979650505050505050565b80356001600160a01b038116811461082a57600080fd5b919050565b60008060006060848603121561084457600080fd5b833567ffffffffffffffff81111561085b57600080fd5b6108678682870161061d565b93505061087660208501610813565b9150604084013590509250925092565b60006020828403121561089857600080fd5b61077782610813565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156108f657603f198886030184526108e4858351610738565b945092850192908501906001016108c8565b5092979650505050505050565b600181811c9082168061091757607f821691505b6020821081141561048f57634e487b7160e01b600052602260045260246000fd5b6000825161094a818460208701610708565b9190910192915050565b60008282101561097457634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fdfea26469706673582212202083a4fd53d68d8ed4b25cc538f75f639f2bfd93f7df0ed46990e41d9f87e12d64736f6c634300080a0033",
}

// RegistryMockABI is the input ABI used to generate the binding from.
const RegistryMockABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getActiveAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activation\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RegistryMockBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const RegistryMockBinRuntime = `608060405234801561001057600080fd5b50600436106100885760003560e01c8063d393c8711161005b578063d393c87114610129578063e2693e3f1461013e578063f2fde38b14610151578063fb825e5f1461018157600080fd5b80633b51650d1461008d5780634622ab03146100c457806378d573a2146100e45780638da5cb5b14610104575b600080fd5b6100a061009b3660046106aa565b610196565b604080516001600160a01b0390931683526020830191909152015b60405180910390f35b6100d76100d23660046106ef565b6101eb565b6040516100bb9190610764565b6100f76100f236600461077e565b610297565b6040516100bb91906107bb565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100bb565b61013c61013736600461082f565b61032a565b005b61011161014c36600461077e565b610401565b61013c61015f366004610886565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b610189610495565b6040516100bb91906108a1565b815160208184018101805160008252928201918501919091209190528054829081106101c157600080fd5b6000918252602090912060029091020180546001909101546001600160a01b039091169250905082565b600181815481106101fb57600080fd5b90600052602060002001600091509050805461021690610903565b80601f016020809104026020016040519081016040528092919081815260200182805461024290610903565b801561028f5780601f106102645761010080835404028352916020019161028f565b820191906000526020600020905b81548152906001019060200180831161027257829003601f168201915b505050505081565b60606000826040516102a99190610938565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561031f576000848152602090819020604080518082019091526002850290910180546001600160a01b031682526001908101548284015290835290920191016102d7565b505050509050919050565b60008360405161033a9190610938565b90815260405190819003602001902054610392576001805480820182556000919091528351610390917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60190602086019061056e565b505b6000836040516103a29190610938565b90815260408051602092819003830181208183019092526001600160a01b039485168152828101938452815460018082018455600093845293909220905160029092020180546001600160a01b03191691909416178355905191015550565b6000806000836040516104149190610938565b908152604051908190036020019020549050806104345750600092915050565b6000836040516104449190610938565b90815260405190819003602001902061045e600183610954565b8154811061046e5761046e610979565b60009182526020909120600290910201546001600160a01b03169392505050565b50919050565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156105655783829060005260206000200180546104d890610903565b80601f016020809104026020016040519081016040528092919081815260200182805461050490610903565b80156105515780601f1061052657610100808354040283529160200191610551565b820191906000526020600020905b81548152906001019060200180831161053457829003601f168201915b5050505050815260200190600101906104b9565b50505050905090565b82805461057a90610903565b90600052602060002090601f01602090048101928261059c57600085556105e2565b82601f106105b557805160ff19168380011785556105e2565b828001600101855582156105e2579182015b828111156105e25782518255916020019190600101906105c7565b506105ee9291506105f2565b5090565b5b808211156105ee57600081556001016105f3565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261062e57600080fd5b813567ffffffffffffffff8082111561064957610649610607565b604051601f8301601f19908116603f0116810190828211818310171561067157610671610607565b8160405283815286602085880101111561068a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156106bd57600080fd5b823567ffffffffffffffff8111156106d457600080fd5b6106e08582860161061d565b95602094909401359450505050565b60006020828403121561070157600080fd5b5035919050565b60005b8381101561072357818101518382015260200161070b565b83811115610732576000848401525b50505050565b60008151808452610750816020860160208601610708565b601f01601f19169290920160200192915050565b6020815260006107776020830184610738565b9392505050565b60006020828403121561079057600080fd5b813567ffffffffffffffff8111156107a757600080fd5b6107b38482850161061d565b949350505050565b602080825282518282018190526000919060409081850190868401855b8281101561080657815180516001600160a01b031685528601518685015292840192908501906001016107d8565b5091979650505050505050565b80356001600160a01b038116811461082a57600080fd5b919050565b60008060006060848603121561084457600080fd5b833567ffffffffffffffff81111561085b57600080fd5b6108678682870161061d565b93505061087660208501610813565b9150604084013590509250925092565b60006020828403121561089857600080fd5b61077782610813565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156108f657603f198886030184526108e4858351610738565b945092850192908501906001016108c8565b5092979650505050505050565b600181811c9082168061091757607f821691505b6020821081141561048f57634e487b7160e01b600052602260045260246000fd5b6000825161094a818460208701610708565b9190910192915050565b60008282101561097457634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fdfea26469706673582212202083a4fd53d68d8ed4b25cc538f75f639f2bfd93f7df0ed46990e41d9f87e12d64736f6c634300080a0033`


// RegistryMockFuncSigs maps the 4-byte function signature to its string representation.
var RegistryMockFuncSigs = map[string]string{
	"e2693e3f": "getActiveAddr(string)",
	"fb825e5f": "getAllNames()",
	"78d573a2": "getAllRecords(string)",
	"4622ab03": "names(uint256)",
	"8da5cb5b": "owner()",
	"3b51650d": "records(string,uint256)",
	"d393c871": "register(string,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// RegistryMockBin is the compiled bytecode used for deploying new contracts.
var RegistryMockBin = "0x608060405234801561001057600080fd5b50610a30806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063d393c8711161005b578063d393c87114610129578063e2693e3f1461013e578063f2fde38b14610151578063fb825e5f1461018157600080fd5b80633b51650d1461008d5780634622ab03146100c457806378d573a2146100e45780638da5cb5b14610104575b600080fd5b6100a061009b366004610611565b610196565b604080516001600160a01b0390931683526020830191909152015b60405180910390f35b6100d76100d2366004610656565b6101eb565b6040516100bb91906106bf565b6100f76100f23660046106d9565b610297565b6040516100bb9190610716565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016100bb565b61013c61013736600461078a565b61032a565b005b61011161014c3660046106d9565b6103fd565b61013c61015f3660046107e1565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b610189610495565b6040516100bb91906107fc565b815160208184018101805160008252928201918501919091209190528054829081106101c157600080fd5b6000918252602090912060029091020180546001909101546001600160a01b039091169250905082565b600181815481106101fb57600080fd5b9060005260206000200160009150905080546102169061085e565b80601f01602080910402602001604051908101604052809291908181526020018280546102429061085e565b801561028f5780601f106102645761010080835404028352916020019161028f565b820191906000526020600020905b81548152906001019060200180831161027257829003601f168201915b505050505081565b60606000826040516102a99190610892565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561031f576000848152602090819020604080518082019091526002850290910180546001600160a01b031682526001908101548284015290835290920191016102d7565b505050509050919050565b60008360405161033a9190610892565b9081526040519081900360200190205460000361038e576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60161038c84826108fd565b505b60008360405161039e9190610892565b90815260408051602092819003830181208183019092526001600160a01b039485168152828101938452815460018082018455600093845293909220905160029092020180546001600160a01b03191691909416178355905191015550565b6000806000836040516104109190610892565b90815260405190819003602001902054905060008190036104345750600092915050565b6000836040516104449190610892565b90815260405190819003602001902061045e6001836109bd565b8154811061046e5761046e6109e4565b60009182526020909120600290910201546001600160a01b03169392505050565b50919050565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156105655783829060005260206000200180546104d89061085e565b80601f01602080910402602001604051908101604052809291908181526020018280546105049061085e565b80156105515780601f1061052657610100808354040283529160200191610551565b820191906000526020600020905b81548152906001019060200180831161053457829003601f168201915b5050505050815260200190600101906104b9565b50505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261059557600080fd5b813567ffffffffffffffff808211156105b0576105b061056e565b604051601f8301601f19908116603f011681019082821181831017156105d8576105d861056e565b816040528381528660208588010111156105f157600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561062457600080fd5b823567ffffffffffffffff81111561063b57600080fd5b61064785828601610584565b95602094909401359450505050565b60006020828403121561066857600080fd5b5035919050565b60005b8381101561068a578181015183820152602001610672565b50506000910152565b600081518084526106ab81602086016020860161066f565b601f01601f19169290920160200192915050565b6020815260006106d26020830184610693565b9392505050565b6000602082840312156106eb57600080fd5b813567ffffffffffffffff81111561070257600080fd5b61070e84828501610584565b949350505050565b602080825282518282018190526000919060409081850190868401855b8281101561076157815180516001600160a01b03168552860151868501529284019290850190600101610733565b5091979650505050505050565b80356001600160a01b038116811461078557600080fd5b919050565b60008060006060848603121561079f57600080fd5b833567ffffffffffffffff8111156107b657600080fd5b6107c286828701610584565b9350506107d16020850161076e565b9150604084013590509250925092565b6000602082840312156107f357600080fd5b6106d28261076e565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561085157603f1988860301845261083f858351610693565b94509285019290850190600101610823565b5092979650505050505050565b600181811c9082168061087257607f821691505b60208210810361048f57634e487b7160e01b600052602260045260246000fd5b600082516108a481846020870161066f565b9190910192915050565b601f8211156108f857600081815260208120601f850160051c810160208610156108d55750805b601f850160051c820191505b818110156108f4578281556001016108e1565b5050505b505050565b815167ffffffffffffffff8111156109175761091761056e565b61092b81610925845461085e565b846108ae565b602080601f83116001811461096057600084156109485750858301515b600019600386901b1c1916600185901b1785556108f4565b600085815260208120601f198616915b8281101561098f57888601518255948401946001909101908401610970565b50858210156109ad5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b818103818111156109de57634e487b7160e01b600052601160045260246000fd5b92915050565b634e487b7160e01b600052603260045260246000fdfea2646970667358221220da61e5fee94fcad7c7fffa1ab47fba5f2e0df57cb3294b1ec8f4ad0467db40d564736f6c63430008110033"

// DeployRegistryMock deploys a new Klaytn contract, binding an instance of RegistryMock to it.
func DeployRegistryMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RegistryMock, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RegistryMock{RegistryMockCaller: RegistryMockCaller{contract: contract}, RegistryMockTransactor: RegistryMockTransactor{contract: contract}, RegistryMockFilterer: RegistryMockFilterer{contract: contract}}, nil
}

// RegistryMock is an auto generated Go binding around a Klaytn contract.
type RegistryMock struct {
	RegistryMockCaller     // Read-only binding to the contract
	RegistryMockTransactor // Write-only binding to the contract
	RegistryMockFilterer   // Log filterer for contract events
}

// RegistryMockCaller is an auto generated read-only Go binding around a Klaytn contract.
type RegistryMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryMockTransactor is an auto generated write-only Go binding around a Klaytn contract.
type RegistryMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryMockFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type RegistryMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryMockSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type RegistryMockSession struct {
	Contract     *RegistryMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryMockCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type RegistryMockCallerSession struct {
	Contract *RegistryMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistryMockTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type RegistryMockTransactorSession struct {
	Contract     *RegistryMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistryMockRaw is an auto generated low-level Go binding around a Klaytn contract.
type RegistryMockRaw struct {
	Contract *RegistryMock // Generic contract binding to access the raw methods on
}

// RegistryMockCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type RegistryMockCallerRaw struct {
	Contract *RegistryMockCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryMockTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type RegistryMockTransactorRaw struct {
	Contract *RegistryMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistryMock creates a new instance of RegistryMock, bound to a specific deployed contract.
func NewRegistryMock(address common.Address, backend bind.ContractBackend) (*RegistryMock, error) {
	contract, err := bindRegistryMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryMock{RegistryMockCaller: RegistryMockCaller{contract: contract}, RegistryMockTransactor: RegistryMockTransactor{contract: contract}, RegistryMockFilterer: RegistryMockFilterer{contract: contract}}, nil
}

// NewRegistryMockCaller creates a new read-only instance of RegistryMock, bound to a specific deployed contract.
func NewRegistryMockCaller(address common.Address, caller bind.ContractCaller) (*RegistryMockCaller, error) {
	contract, err := bindRegistryMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryMockCaller{contract: contract}, nil
}

// NewRegistryMockTransactor creates a new write-only instance of RegistryMock, bound to a specific deployed contract.
func NewRegistryMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryMockTransactor, error) {
	contract, err := bindRegistryMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryMockTransactor{contract: contract}, nil
}

// NewRegistryMockFilterer creates a new log filterer instance of RegistryMock, bound to a specific deployed contract.
func NewRegistryMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryMockFilterer, error) {
	contract, err := bindRegistryMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryMockFilterer{contract: contract}, nil
}

// bindRegistryMock binds a generic wrapper to an already deployed contract.
func bindRegistryMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryMock *RegistryMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryMock.Contract.RegistryMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryMock *RegistryMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryMock.Contract.RegistryMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryMock *RegistryMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryMock.Contract.RegistryMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryMock *RegistryMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryMock *RegistryMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryMock *RegistryMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryMock.Contract.contract.Transact(opts, method, params...)
}

// GetActiveAddr is a free data retrieval call binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) view returns(address)
func (_RegistryMock *RegistryMockCaller) GetActiveAddr(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _RegistryMock.contract.Call(opts, &out, "getActiveAddr", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetActiveAddr is a free data retrieval call binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) view returns(address)
func (_RegistryMock *RegistryMockSession) GetActiveAddr(name string) (common.Address, error) {
	return _RegistryMock.Contract.GetActiveAddr(&_RegistryMock.CallOpts, name)
}

// GetActiveAddr is a free data retrieval call binding the contract method 0xe2693e3f.
//
// Solidity: function getActiveAddr(string name) view returns(address)
func (_RegistryMock *RegistryMockCallerSession) GetActiveAddr(name string) (common.Address, error) {
	return _RegistryMock.Contract.GetActiveAddr(&_RegistryMock.CallOpts, name)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_RegistryMock *RegistryMockCaller) GetAllNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _RegistryMock.contract.Call(opts, &out, "getAllNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_RegistryMock *RegistryMockSession) GetAllNames() ([]string, error) {
	return _RegistryMock.Contract.GetAllNames(&_RegistryMock.CallOpts)
}

// GetAllNames is a free data retrieval call binding the contract method 0xfb825e5f.
//
// Solidity: function getAllNames() view returns(string[])
func (_RegistryMock *RegistryMockCallerSession) GetAllNames() ([]string, error) {
	return _RegistryMock.Contract.GetAllNames(&_RegistryMock.CallOpts)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_RegistryMock *RegistryMockCaller) GetAllRecords(opts *bind.CallOpts, name string) ([]IRegistryRecord, error) {
	var out []interface{}
	err := _RegistryMock.contract.Call(opts, &out, "getAllRecords", name)

	if err != nil {
		return *new([]IRegistryRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRegistryRecord)).(*[]IRegistryRecord)

	return out0, err

}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_RegistryMock *RegistryMockSession) GetAllRecords(name string) ([]IRegistryRecord, error) {
	return _RegistryMock.Contract.GetAllRecords(&_RegistryMock.CallOpts, name)
}

// GetAllRecords is a free data retrieval call binding the contract method 0x78d573a2.
//
// Solidity: function getAllRecords(string name) view returns((address,uint256)[])
func (_RegistryMock *RegistryMockCallerSession) GetAllRecords(name string) ([]IRegistryRecord, error) {
	return _RegistryMock.Contract.GetAllRecords(&_RegistryMock.CallOpts, name)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_RegistryMock *RegistryMockCaller) Names(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RegistryMock.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_RegistryMock *RegistryMockSession) Names(arg0 *big.Int) (string, error) {
	return _RegistryMock.Contract.Names(&_RegistryMock.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x4622ab03.
//
// Solidity: function names(uint256 ) view returns(string)
func (_RegistryMock *RegistryMockCallerSession) Names(arg0 *big.Int) (string, error) {
	return _RegistryMock.Contract.Names(&_RegistryMock.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryMock *RegistryMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryMock *RegistryMockSession) Owner() (common.Address, error) {
	return _RegistryMock.Contract.Owner(&_RegistryMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryMock *RegistryMockCallerSession) Owner() (common.Address, error) {
	return _RegistryMock.Contract.Owner(&_RegistryMock.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_RegistryMock *RegistryMockCaller) Records(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	var out []interface{}
	err := _RegistryMock.contract.Call(opts, &out, "records", arg0, arg1)

	outstruct := new(struct {
		Addr       common.Address
		Activation *big.Int
	})

	outstruct.Addr = out[0].(common.Address)
	outstruct.Activation = out[1].(*big.Int)
	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_RegistryMock *RegistryMockSession) Records(arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	return _RegistryMock.Contract.Records(&_RegistryMock.CallOpts, arg0, arg1)
}

// Records is a free data retrieval call binding the contract method 0x3b51650d.
//
// Solidity: function records(string , uint256 ) view returns(address addr, uint256 activation)
func (_RegistryMock *RegistryMockCallerSession) Records(arg0 string, arg1 *big.Int) (struct {
	Addr       common.Address
	Activation *big.Int
}, error) {
	return _RegistryMock.Contract.Records(&_RegistryMock.CallOpts, arg0, arg1)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_RegistryMock *RegistryMockTransactor) Register(opts *bind.TransactOpts, name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _RegistryMock.contract.Transact(opts, "register", name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_RegistryMock *RegistryMockSession) Register(name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _RegistryMock.Contract.Register(&_RegistryMock.TransactOpts, name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address addr, uint256 activation) returns()
func (_RegistryMock *RegistryMockTransactorSession) Register(name string, addr common.Address, activation *big.Int) (*types.Transaction, error) {
	return _RegistryMock.Contract.Register(&_RegistryMock.TransactOpts, name, addr, activation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistryMock *RegistryMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RegistryMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistryMock *RegistryMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistryMock.Contract.TransferOwnership(&_RegistryMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistryMock *RegistryMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistryMock.Contract.TransferOwnership(&_RegistryMock.TransactOpts, newOwner)
}

// RegistryMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistryMock contract.
type RegistryMockOwnershipTransferredIterator struct {
	Event *RegistryMockOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMockOwnershipTransferred)
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
		it.Event = new(RegistryMockOwnershipTransferred)
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
func (it *RegistryMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMockOwnershipTransferred represents a OwnershipTransferred event raised by the RegistryMock contract.
type RegistryMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistryMock *RegistryMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistryMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryMockOwnershipTransferredIterator{contract: _RegistryMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistryMock *RegistryMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistryMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMockOwnershipTransferred)
				if err := _RegistryMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RegistryMock *RegistryMockFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryMockOwnershipTransferred, error) {
	event := new(RegistryMockOwnershipTransferred)
	if err := _RegistryMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryMockRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the RegistryMock contract.
type RegistryMockRegisteredIterator struct {
	Event *RegistryMockRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryMockRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMockRegistered)
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
		it.Event = new(RegistryMockRegistered)
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
func (it *RegistryMockRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMockRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMockRegistered represents a Registered event raised by the RegistryMock contract.
type RegistryMockRegistered struct {
	Name       string
	Addr       common.Address
	Activation *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_RegistryMock *RegistryMockFilterer) FilterRegistered(opts *bind.FilterOpts, addr []common.Address, activation []*big.Int) (*RegistryMockRegisteredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var activationRule []interface{}
	for _, activationItem := range activation {
		activationRule = append(activationRule, activationItem)
	}

	logs, sub, err := _RegistryMock.contract.FilterLogs(opts, "Registered", addrRule, activationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryMockRegisteredIterator{contract: _RegistryMock.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_RegistryMock *RegistryMockFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistryMockRegistered, addr []common.Address, activation []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var activationRule []interface{}
	for _, activationItem := range activation {
		activationRule = append(activationRule, activationItem)
	}

	logs, sub, err := _RegistryMock.contract.WatchLogs(opts, "Registered", addrRule, activationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMockRegistered)
				if err := _RegistryMock.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x142e1fdac7ecccbc62af925f0b4039db26847b625602e56b1421dfbc8a0e4f30.
//
// Solidity: event Registered(string name, address indexed addr, uint256 indexed activation)
func (_RegistryMock *RegistryMockFilterer) ParseRegistered(log types.Log) (*RegistryMockRegistered, error) {
	event := new(RegistryMockRegistered)
	if err := _RegistryMock.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	return event, nil
}
