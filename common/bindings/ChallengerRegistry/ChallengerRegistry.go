// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ChallengerRegistry

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

// ChallengerRegistryMetaData contains all meta data concerning the ChallengerRegistry contract.
var ChallengerRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowOperators\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateChallengerRegistrationDigestHash\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengerSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengersStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTypes.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deRegisterChallenger\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManagerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengerLastSeen\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengerStatus\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTypes.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperator\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowlisted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidChallenger\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"polOperatorAllowlist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumTypes.Status\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerChallenger\",\"inputs\":[{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengerSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasherAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ChallengerDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengerRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorsWhiteListed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// ChallengerRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengerRegistryMetaData.ABI instead.
var ChallengerRegistryABI = ChallengerRegistryMetaData.ABI

// ChallengerRegistry is an auto generated Go binding around an Ethereum contract.
type ChallengerRegistry struct {
	ChallengerRegistryCaller     // Read-only binding to the contract
	ChallengerRegistryTransactor // Write-only binding to the contract
	ChallengerRegistryFilterer   // Log filterer for contract events
}

// ChallengerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengerRegistrySession struct {
	Contract     *ChallengerRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChallengerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengerRegistryCallerSession struct {
	Contract *ChallengerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ChallengerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengerRegistryTransactorSession struct {
	Contract     *ChallengerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ChallengerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengerRegistryRaw struct {
	Contract *ChallengerRegistry // Generic contract binding to access the raw methods on
}

// ChallengerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengerRegistryCallerRaw struct {
	Contract *ChallengerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengerRegistryTransactorRaw struct {
	Contract *ChallengerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengerRegistry creates a new instance of ChallengerRegistry, bound to a specific deployed contract.
func NewChallengerRegistry(address common.Address, backend bind.ContractBackend) (*ChallengerRegistry, error) {
	contract, err := bindChallengerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistry{ChallengerRegistryCaller: ChallengerRegistryCaller{contract: contract}, ChallengerRegistryTransactor: ChallengerRegistryTransactor{contract: contract}, ChallengerRegistryFilterer: ChallengerRegistryFilterer{contract: contract}}, nil
}

// NewChallengerRegistryCaller creates a new read-only instance of ChallengerRegistry, bound to a specific deployed contract.
func NewChallengerRegistryCaller(address common.Address, caller bind.ContractCaller) (*ChallengerRegistryCaller, error) {
	contract, err := bindChallengerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryCaller{contract: contract}, nil
}

// NewChallengerRegistryTransactor creates a new write-only instance of ChallengerRegistry, bound to a specific deployed contract.
func NewChallengerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengerRegistryTransactor, error) {
	contract, err := bindChallengerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryTransactor{contract: contract}, nil
}

// NewChallengerRegistryFilterer creates a new log filterer instance of ChallengerRegistry, bound to a specific deployed contract.
func NewChallengerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengerRegistryFilterer, error) {
	contract, err := bindChallengerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryFilterer{contract: contract}, nil
}

// bindChallengerRegistry binds a generic wrapper to an already deployed contract.
func bindChallengerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengerRegistry *ChallengerRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengerRegistry.Contract.ChallengerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengerRegistry *ChallengerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.ChallengerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengerRegistry *ChallengerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.ChallengerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengerRegistry *ChallengerRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengerRegistry *ChallengerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengerRegistry *ChallengerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChallengerRegistry *ChallengerRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChallengerRegistry *ChallengerRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChallengerRegistry.Contract.UPGRADEINTERFACEVERSION(&_ChallengerRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChallengerRegistry.Contract.UPGRADEINTERFACEVERSION(&_ChallengerRegistry.CallOpts)
}

// CalculateChallengerRegistrationDigestHash is a free data retrieval call binding the contract method 0xd417a559.
//
// Solidity: function calculateChallengerRegistrationDigestHash(address challenger, address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistryCaller) CalculateChallengerRegistrationDigestHash(opts *bind.CallOpts, challenger common.Address, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "calculateChallengerRegistrationDigestHash", challenger, operator, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateChallengerRegistrationDigestHash is a free data retrieval call binding the contract method 0xd417a559.
//
// Solidity: function calculateChallengerRegistrationDigestHash(address challenger, address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistrySession) CalculateChallengerRegistrationDigestHash(challenger common.Address, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ChallengerRegistry.Contract.CalculateChallengerRegistrationDigestHash(&_ChallengerRegistry.CallOpts, challenger, operator, salt, expiry)
}

// CalculateChallengerRegistrationDigestHash is a free data retrieval call binding the contract method 0xd417a559.
//
// Solidity: function calculateChallengerRegistrationDigestHash(address challenger, address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) CalculateChallengerRegistrationDigestHash(challenger common.Address, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ChallengerRegistry.Contract.CalculateChallengerRegistrationDigestHash(&_ChallengerRegistry.CallOpts, challenger, operator, salt, expiry)
}

// ChallengerSaltUsed is a free data retrieval call binding the contract method 0x382ff9b3.
//
// Solidity: function challengerSaltUsed(address , bytes32 ) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCaller) ChallengerSaltUsed(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "challengerSaltUsed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChallengerSaltUsed is a free data retrieval call binding the contract method 0x382ff9b3.
//
// Solidity: function challengerSaltUsed(address , bytes32 ) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistrySession) ChallengerSaltUsed(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ChallengerRegistry.Contract.ChallengerSaltUsed(&_ChallengerRegistry.CallOpts, arg0, arg1)
}

// ChallengerSaltUsed is a free data retrieval call binding the contract method 0x382ff9b3.
//
// Solidity: function challengerSaltUsed(address , bytes32 ) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) ChallengerSaltUsed(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ChallengerRegistry.Contract.ChallengerSaltUsed(&_ChallengerRegistry.CallOpts, arg0, arg1)
}

// ChallengersStatus is a free data retrieval call binding the contract method 0x882047e9.
//
// Solidity: function challengersStatus(address ) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistryCaller) ChallengersStatus(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "challengersStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChallengersStatus is a free data retrieval call binding the contract method 0x882047e9.
//
// Solidity: function challengersStatus(address ) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistrySession) ChallengersStatus(arg0 common.Address) (uint8, error) {
	return _ChallengerRegistry.Contract.ChallengersStatus(&_ChallengerRegistry.CallOpts, arg0)
}

// ChallengersStatus is a free data retrieval call binding the contract method 0x882047e9.
//
// Solidity: function challengersStatus(address ) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) ChallengersStatus(arg0 common.Address) (uint8, error) {
	return _ChallengerRegistry.Contract.ChallengersStatus(&_ChallengerRegistry.CallOpts, arg0)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_ChallengerRegistry *ChallengerRegistryCaller) DelegationManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "delegationManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_ChallengerRegistry *ChallengerRegistrySession) DelegationManagerAddress() (common.Address, error) {
	return _ChallengerRegistry.Contract.DelegationManagerAddress(&_ChallengerRegistry.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) DelegationManagerAddress() (common.Address, error) {
	return _ChallengerRegistry.Contract.DelegationManagerAddress(&_ChallengerRegistry.CallOpts)
}

// DomainSeperator is a free data retrieval call binding the contract method 0xcb73ac56.
//
// Solidity: function domainSeperator() view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistryCaller) DomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "domainSeperator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeperator is a free data retrieval call binding the contract method 0xcb73ac56.
//
// Solidity: function domainSeperator() view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistrySession) DomainSeperator() ([32]byte, error) {
	return _ChallengerRegistry.Contract.DomainSeperator(&_ChallengerRegistry.CallOpts)
}

// DomainSeperator is a free data retrieval call binding the contract method 0xcb73ac56.
//
// Solidity: function domainSeperator() view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) DomainSeperator() ([32]byte, error) {
	return _ChallengerRegistry.Contract.DomainSeperator(&_ChallengerRegistry.CallOpts)
}

// GetChallengerLastSeen is a free data retrieval call binding the contract method 0xa86013c6.
//
// Solidity: function getChallengerLastSeen(address challenger) view returns(uint256)
func (_ChallengerRegistry *ChallengerRegistryCaller) GetChallengerLastSeen(opts *bind.CallOpts, challenger common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "getChallengerLastSeen", challenger)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChallengerLastSeen is a free data retrieval call binding the contract method 0xa86013c6.
//
// Solidity: function getChallengerLastSeen(address challenger) view returns(uint256)
func (_ChallengerRegistry *ChallengerRegistrySession) GetChallengerLastSeen(challenger common.Address) (*big.Int, error) {
	return _ChallengerRegistry.Contract.GetChallengerLastSeen(&_ChallengerRegistry.CallOpts, challenger)
}

// GetChallengerLastSeen is a free data retrieval call binding the contract method 0xa86013c6.
//
// Solidity: function getChallengerLastSeen(address challenger) view returns(uint256)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) GetChallengerLastSeen(challenger common.Address) (*big.Int, error) {
	return _ChallengerRegistry.Contract.GetChallengerLastSeen(&_ChallengerRegistry.CallOpts, challenger)
}

// GetChallengerStatus is a free data retrieval call binding the contract method 0x639c257b.
//
// Solidity: function getChallengerStatus(address challenger) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistryCaller) GetChallengerStatus(opts *bind.CallOpts, challenger common.Address) (uint8, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "getChallengerStatus", challenger)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetChallengerStatus is a free data retrieval call binding the contract method 0x639c257b.
//
// Solidity: function getChallengerStatus(address challenger) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistrySession) GetChallengerStatus(challenger common.Address) (uint8, error) {
	return _ChallengerRegistry.Contract.GetChallengerStatus(&_ChallengerRegistry.CallOpts, challenger)
}

// GetChallengerStatus is a free data retrieval call binding the contract method 0x639c257b.
//
// Solidity: function getChallengerStatus(address challenger) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) GetChallengerStatus(challenger common.Address) (uint8, error) {
	return _ChallengerRegistry.Contract.GetChallengerStatus(&_ChallengerRegistry.CallOpts, challenger)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address challenger) view returns(address operator)
func (_ChallengerRegistry *ChallengerRegistryCaller) GetOperator(opts *bind.CallOpts, challenger common.Address) (common.Address, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "getOperator", challenger)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address challenger) view returns(address operator)
func (_ChallengerRegistry *ChallengerRegistrySession) GetOperator(challenger common.Address) (common.Address, error) {
	return _ChallengerRegistry.Contract.GetOperator(&_ChallengerRegistry.CallOpts, challenger)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address challenger) view returns(address operator)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) GetOperator(challenger common.Address) (common.Address, error) {
	return _ChallengerRegistry.Contract.GetOperator(&_ChallengerRegistry.CallOpts, challenger)
}

// IsAllowlisted is a free data retrieval call binding the contract method 0x05a3b809.
//
// Solidity: function isAllowlisted(address operator) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCaller) IsAllowlisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "isAllowlisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowlisted is a free data retrieval call binding the contract method 0x05a3b809.
//
// Solidity: function isAllowlisted(address operator) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistrySession) IsAllowlisted(operator common.Address) (bool, error) {
	return _ChallengerRegistry.Contract.IsAllowlisted(&_ChallengerRegistry.CallOpts, operator)
}

// IsAllowlisted is a free data retrieval call binding the contract method 0x05a3b809.
//
// Solidity: function isAllowlisted(address operator) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) IsAllowlisted(operator common.Address) (bool, error) {
	return _ChallengerRegistry.Contract.IsAllowlisted(&_ChallengerRegistry.CallOpts, operator)
}

// IsValidChallenger is a free data retrieval call binding the contract method 0xb6fe6d1b.
//
// Solidity: function isValidChallenger(address challenger) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCaller) IsValidChallenger(opts *bind.CallOpts, challenger common.Address) (bool, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "isValidChallenger", challenger)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidChallenger is a free data retrieval call binding the contract method 0xb6fe6d1b.
//
// Solidity: function isValidChallenger(address challenger) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistrySession) IsValidChallenger(challenger common.Address) (bool, error) {
	return _ChallengerRegistry.Contract.IsValidChallenger(&_ChallengerRegistry.CallOpts, challenger)
}

// IsValidChallenger is a free data retrieval call binding the contract method 0xb6fe6d1b.
//
// Solidity: function isValidChallenger(address challenger) view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) IsValidChallenger(challenger common.Address) (bool, error) {
	return _ChallengerRegistry.Contract.IsValidChallenger(&_ChallengerRegistry.CallOpts, challenger)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChallengerRegistry *ChallengerRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChallengerRegistry *ChallengerRegistrySession) Owner() (common.Address, error) {
	return _ChallengerRegistry.Contract.Owner(&_ChallengerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) Owner() (common.Address, error) {
	return _ChallengerRegistry.Contract.Owner(&_ChallengerRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ChallengerRegistry *ChallengerRegistrySession) Paused() (bool, error) {
	return _ChallengerRegistry.Contract.Paused(&_ChallengerRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) Paused() (bool, error) {
	return _ChallengerRegistry.Contract.Paused(&_ChallengerRegistry.CallOpts)
}

// PolOperatorAllowlist is a free data retrieval call binding the contract method 0x27b00650.
//
// Solidity: function polOperatorAllowlist(address ) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistryCaller) PolOperatorAllowlist(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "polOperatorAllowlist", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PolOperatorAllowlist is a free data retrieval call binding the contract method 0x27b00650.
//
// Solidity: function polOperatorAllowlist(address ) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistrySession) PolOperatorAllowlist(arg0 common.Address) (uint8, error) {
	return _ChallengerRegistry.Contract.PolOperatorAllowlist(&_ChallengerRegistry.CallOpts, arg0)
}

// PolOperatorAllowlist is a free data retrieval call binding the contract method 0x27b00650.
//
// Solidity: function polOperatorAllowlist(address ) view returns(uint8)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) PolOperatorAllowlist(arg0 common.Address) (uint8, error) {
	return _ChallengerRegistry.Contract.PolOperatorAllowlist(&_ChallengerRegistry.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ChallengerRegistry.Contract.ProxiableUUID(&_ChallengerRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ChallengerRegistry.Contract.ProxiableUUID(&_ChallengerRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_ChallengerRegistry *ChallengerRegistryCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengerRegistry.contract.Call(opts, &out, "slasherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_ChallengerRegistry *ChallengerRegistrySession) SlasherAddress() (common.Address, error) {
	return _ChallengerRegistry.Contract.SlasherAddress(&_ChallengerRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_ChallengerRegistry *ChallengerRegistryCallerSession) SlasherAddress() (common.Address, error) {
	return _ChallengerRegistry.Contract.SlasherAddress(&_ChallengerRegistry.CallOpts)
}

// AllowOperators is a paid mutator transaction binding the contract method 0x432de9c8.
//
// Solidity: function allowOperators(address[] operators) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) AllowOperators(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "allowOperators", operators)
}

// AllowOperators is a paid mutator transaction binding the contract method 0x432de9c8.
//
// Solidity: function allowOperators(address[] operators) returns()
func (_ChallengerRegistry *ChallengerRegistrySession) AllowOperators(operators []common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.AllowOperators(&_ChallengerRegistry.TransactOpts, operators)
}

// AllowOperators is a paid mutator transaction binding the contract method 0x432de9c8.
//
// Solidity: function allowOperators(address[] operators) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) AllowOperators(operators []common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.AllowOperators(&_ChallengerRegistry.TransactOpts, operators)
}

// DeRegisterChallenger is a paid mutator transaction binding the contract method 0xf91e67db.
//
// Solidity: function deRegisterChallenger(address challenger) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) DeRegisterChallenger(opts *bind.TransactOpts, challenger common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "deRegisterChallenger", challenger)
}

// DeRegisterChallenger is a paid mutator transaction binding the contract method 0xf91e67db.
//
// Solidity: function deRegisterChallenger(address challenger) returns()
func (_ChallengerRegistry *ChallengerRegistrySession) DeRegisterChallenger(challenger common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.DeRegisterChallenger(&_ChallengerRegistry.TransactOpts, challenger)
}

// DeRegisterChallenger is a paid mutator transaction binding the contract method 0xf91e67db.
//
// Solidity: function deRegisterChallenger(address challenger) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) DeRegisterChallenger(challenger common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.DeRegisterChallenger(&_ChallengerRegistry.TransactOpts, challenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChallengerRegistry *ChallengerRegistrySession) Initialize() (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.Initialize(&_ChallengerRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.Initialize(&_ChallengerRegistry.TransactOpts)
}

// RegisterChallenger is a paid mutator transaction binding the contract method 0xf3dd0e46.
//
// Solidity: function registerChallenger(address challenger, bytes32 salt, uint256 expiry, bytes challengerSignature) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) RegisterChallenger(opts *bind.TransactOpts, challenger common.Address, salt [32]byte, expiry *big.Int, challengerSignature []byte) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "registerChallenger", challenger, salt, expiry, challengerSignature)
}

// RegisterChallenger is a paid mutator transaction binding the contract method 0xf3dd0e46.
//
// Solidity: function registerChallenger(address challenger, bytes32 salt, uint256 expiry, bytes challengerSignature) returns()
func (_ChallengerRegistry *ChallengerRegistrySession) RegisterChallenger(challenger common.Address, salt [32]byte, expiry *big.Int, challengerSignature []byte) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.RegisterChallenger(&_ChallengerRegistry.TransactOpts, challenger, salt, expiry, challengerSignature)
}

// RegisterChallenger is a paid mutator transaction binding the contract method 0xf3dd0e46.
//
// Solidity: function registerChallenger(address challenger, bytes32 salt, uint256 expiry, bytes challengerSignature) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) RegisterChallenger(challenger common.Address, salt [32]byte, expiry *big.Int, challengerSignature []byte) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.RegisterChallenger(&_ChallengerRegistry.TransactOpts, challenger, salt, expiry, challengerSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChallengerRegistry *ChallengerRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.RenounceOwnership(&_ChallengerRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.RenounceOwnership(&_ChallengerRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChallengerRegistry *ChallengerRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.TransferOwnership(&_ChallengerRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.TransferOwnership(&_ChallengerRegistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChallengerRegistry *ChallengerRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChallengerRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChallengerRegistry *ChallengerRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.UpgradeToAndCall(&_ChallengerRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChallengerRegistry *ChallengerRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChallengerRegistry.Contract.UpgradeToAndCall(&_ChallengerRegistry.TransactOpts, newImplementation, data)
}

// ChallengerRegistryChallengerDeregisteredIterator is returned from FilterChallengerDeregistered and is used to iterate over the raw logs and unpacked data for ChallengerDeregistered events raised by the ChallengerRegistry contract.
type ChallengerRegistryChallengerDeregisteredIterator struct {
	Event *ChallengerRegistryChallengerDeregistered // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryChallengerDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryChallengerDeregistered)
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
		it.Event = new(ChallengerRegistryChallengerDeregistered)
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
func (it *ChallengerRegistryChallengerDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryChallengerDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryChallengerDeregistered represents a ChallengerDeregistered event raised by the ChallengerRegistry contract.
type ChallengerRegistryChallengerDeregistered struct {
	Operator   common.Address
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengerDeregistered is a free log retrieval operation binding the contract event 0x6c4abdb05e86bfac0db1853dab79515093329275ce863844577b435b1479e840.
//
// Solidity: event ChallengerDeregistered(address indexed operator, address indexed challenger)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterChallengerDeregistered(opts *bind.FilterOpts, operator []common.Address, challenger []common.Address) (*ChallengerRegistryChallengerDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "ChallengerDeregistered", operatorRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryChallengerDeregisteredIterator{contract: _ChallengerRegistry.contract, event: "ChallengerDeregistered", logs: logs, sub: sub}, nil
}

// WatchChallengerDeregistered is a free log subscription operation binding the contract event 0x6c4abdb05e86bfac0db1853dab79515093329275ce863844577b435b1479e840.
//
// Solidity: event ChallengerDeregistered(address indexed operator, address indexed challenger)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchChallengerDeregistered(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryChallengerDeregistered, operator []common.Address, challenger []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "ChallengerDeregistered", operatorRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryChallengerDeregistered)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "ChallengerDeregistered", log); err != nil {
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

// ParseChallengerDeregistered is a log parse operation binding the contract event 0x6c4abdb05e86bfac0db1853dab79515093329275ce863844577b435b1479e840.
//
// Solidity: event ChallengerDeregistered(address indexed operator, address indexed challenger)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseChallengerDeregistered(log types.Log) (*ChallengerRegistryChallengerDeregistered, error) {
	event := new(ChallengerRegistryChallengerDeregistered)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "ChallengerDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryChallengerRegisteredIterator is returned from FilterChallengerRegistered and is used to iterate over the raw logs and unpacked data for ChallengerRegistered events raised by the ChallengerRegistry contract.
type ChallengerRegistryChallengerRegisteredIterator struct {
	Event *ChallengerRegistryChallengerRegistered // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryChallengerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryChallengerRegistered)
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
		it.Event = new(ChallengerRegistryChallengerRegistered)
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
func (it *ChallengerRegistryChallengerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryChallengerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryChallengerRegistered represents a ChallengerRegistered event raised by the ChallengerRegistry contract.
type ChallengerRegistryChallengerRegistered struct {
	Operator   common.Address
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengerRegistered is a free log retrieval operation binding the contract event 0xd2eb8b3ee9ef8c9c3812b694915ad69a399942ff211baa267c3683a2ee4679f8.
//
// Solidity: event ChallengerRegistered(address indexed operator, address indexed challenger)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterChallengerRegistered(opts *bind.FilterOpts, operator []common.Address, challenger []common.Address) (*ChallengerRegistryChallengerRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "ChallengerRegistered", operatorRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryChallengerRegisteredIterator{contract: _ChallengerRegistry.contract, event: "ChallengerRegistered", logs: logs, sub: sub}, nil
}

// WatchChallengerRegistered is a free log subscription operation binding the contract event 0xd2eb8b3ee9ef8c9c3812b694915ad69a399942ff211baa267c3683a2ee4679f8.
//
// Solidity: event ChallengerRegistered(address indexed operator, address indexed challenger)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchChallengerRegistered(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryChallengerRegistered, operator []common.Address, challenger []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "ChallengerRegistered", operatorRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryChallengerRegistered)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "ChallengerRegistered", log); err != nil {
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

// ParseChallengerRegistered is a log parse operation binding the contract event 0xd2eb8b3ee9ef8c9c3812b694915ad69a399942ff211baa267c3683a2ee4679f8.
//
// Solidity: event ChallengerRegistered(address indexed operator, address indexed challenger)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseChallengerRegistered(log types.Log) (*ChallengerRegistryChallengerRegistered, error) {
	event := new(ChallengerRegistryChallengerRegistered)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "ChallengerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChallengerRegistry contract.
type ChallengerRegistryInitializedIterator struct {
	Event *ChallengerRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryInitialized)
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
		it.Event = new(ChallengerRegistryInitialized)
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
func (it *ChallengerRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryInitialized represents a Initialized event raised by the ChallengerRegistry contract.
type ChallengerRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChallengerRegistryInitializedIterator, error) {

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryInitializedIterator{contract: _ChallengerRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryInitialized)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseInitialized(log types.Log) (*ChallengerRegistryInitialized, error) {
	event := new(ChallengerRegistryInitialized)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ChallengerRegistry contract.
type ChallengerRegistryOperatorDeregisteredIterator struct {
	Event *ChallengerRegistryOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryOperatorDeregistered)
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
		it.Event = new(ChallengerRegistryOperatorDeregistered)
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
func (it *ChallengerRegistryOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryOperatorDeregistered represents a OperatorDeregistered event raised by the ChallengerRegistry contract.
type ChallengerRegistryOperatorDeregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*ChallengerRegistryOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryOperatorDeregisteredIterator{contract: _ChallengerRegistry.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryOperatorDeregistered)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseOperatorDeregistered(log types.Log) (*ChallengerRegistryOperatorDeregistered, error) {
	event := new(ChallengerRegistryOperatorDeregistered)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryOperatorsWhiteListedIterator is returned from FilterOperatorsWhiteListed and is used to iterate over the raw logs and unpacked data for OperatorsWhiteListed events raised by the ChallengerRegistry contract.
type ChallengerRegistryOperatorsWhiteListedIterator struct {
	Event *ChallengerRegistryOperatorsWhiteListed // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryOperatorsWhiteListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryOperatorsWhiteListed)
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
		it.Event = new(ChallengerRegistryOperatorsWhiteListed)
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
func (it *ChallengerRegistryOperatorsWhiteListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryOperatorsWhiteListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryOperatorsWhiteListed represents a OperatorsWhiteListed event raised by the ChallengerRegistry contract.
type ChallengerRegistryOperatorsWhiteListed struct {
	Operator    []common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorsWhiteListed is a free log retrieval operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterOperatorsWhiteListed(opts *bind.FilterOpts) (*ChallengerRegistryOperatorsWhiteListedIterator, error) {

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "OperatorsWhiteListed")
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryOperatorsWhiteListedIterator{contract: _ChallengerRegistry.contract, event: "OperatorsWhiteListed", logs: logs, sub: sub}, nil
}

// WatchOperatorsWhiteListed is a free log subscription operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchOperatorsWhiteListed(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryOperatorsWhiteListed) (event.Subscription, error) {

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "OperatorsWhiteListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryOperatorsWhiteListed)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "OperatorsWhiteListed", log); err != nil {
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

// ParseOperatorsWhiteListed is a log parse operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseOperatorsWhiteListed(log types.Log) (*ChallengerRegistryOperatorsWhiteListed, error) {
	event := new(ChallengerRegistryOperatorsWhiteListed)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "OperatorsWhiteListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChallengerRegistry contract.
type ChallengerRegistryOwnershipTransferredIterator struct {
	Event *ChallengerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryOwnershipTransferred)
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
		it.Event = new(ChallengerRegistryOwnershipTransferred)
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
func (it *ChallengerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ChallengerRegistry contract.
type ChallengerRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChallengerRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryOwnershipTransferredIterator{contract: _ChallengerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryOwnershipTransferred)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ChallengerRegistryOwnershipTransferred, error) {
	event := new(ChallengerRegistryOwnershipTransferred)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ChallengerRegistry contract.
type ChallengerRegistryPausedIterator struct {
	Event *ChallengerRegistryPaused // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryPaused)
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
		it.Event = new(ChallengerRegistryPaused)
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
func (it *ChallengerRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryPaused represents a Paused event raised by the ChallengerRegistry contract.
type ChallengerRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*ChallengerRegistryPausedIterator, error) {

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryPausedIterator{contract: _ChallengerRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryPaused)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParsePaused(log types.Log) (*ChallengerRegistryPaused, error) {
	event := new(ChallengerRegistryPaused)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ChallengerRegistry contract.
type ChallengerRegistryUnpausedIterator struct {
	Event *ChallengerRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryUnpaused)
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
		it.Event = new(ChallengerRegistryUnpaused)
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
func (it *ChallengerRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryUnpaused represents a Unpaused event raised by the ChallengerRegistry contract.
type ChallengerRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ChallengerRegistryUnpausedIterator, error) {

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryUnpausedIterator{contract: _ChallengerRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryUnpaused)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseUnpaused(log types.Log) (*ChallengerRegistryUnpaused, error) {
	event := new(ChallengerRegistryUnpaused)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengerRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ChallengerRegistry contract.
type ChallengerRegistryUpgradedIterator struct {
	Event *ChallengerRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ChallengerRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengerRegistryUpgraded)
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
		it.Event = new(ChallengerRegistryUpgraded)
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
func (it *ChallengerRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengerRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengerRegistryUpgraded represents a Upgraded event raised by the ChallengerRegistry contract.
type ChallengerRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChallengerRegistry *ChallengerRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChallengerRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChallengerRegistryUpgradedIterator{contract: _ChallengerRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChallengerRegistry *ChallengerRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChallengerRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChallengerRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengerRegistryUpgraded)
				if err := _ChallengerRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChallengerRegistry *ChallengerRegistryFilterer) ParseUpgraded(log types.Log) (*ChallengerRegistryUpgraded, error) {
	event := new(ChallengerRegistryUpgraded)
	if err := _ChallengerRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
