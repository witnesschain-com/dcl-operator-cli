// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProverRegistry

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

// ProverRegistryMetaData contains all meta data concerning the ProverRegistry contract.
var ProverRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateProverRegistrationDigestHash\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dePINOperatorProversStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIProverRegistry.DePINProverStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dePINOperatorsStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIProverRegistry.DePINOperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dePINProverSaltUsed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dePINProversStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIProverRegistry.DePINProverStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deRegisterProver\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorStatus\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIProverRegistry.DePINOperatorStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProverStatus\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIProverRegistry.DePINProverStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRegisteredProver\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerProver\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proverSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainAdmin\",\"inputs\":[{\"name\":\"_chainAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ChainAdminInitialized\",\"inputs\":[{\"name\":\"_chainAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DePINProverDeregistered\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DePINProverParticipationKeyRegistered\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"participationKey\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DePINProverRegistered\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proverStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIProverRegistry.DePINProverStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// ProverRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProverRegistryMetaData.ABI instead.
var ProverRegistryABI = ProverRegistryMetaData.ABI

// ProverRegistry is an auto generated Go binding around an Ethereum contract.
type ProverRegistry struct {
	ProverRegistryCaller     // Read-only binding to the contract
	ProverRegistryTransactor // Write-only binding to the contract
	ProverRegistryFilterer   // Log filterer for contract events
}

// ProverRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProverRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProverRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProverRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProverRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProverRegistrySession struct {
	Contract     *ProverRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProverRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProverRegistryCallerSession struct {
	Contract *ProverRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ProverRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProverRegistryTransactorSession struct {
	Contract     *ProverRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ProverRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProverRegistryRaw struct {
	Contract *ProverRegistry // Generic contract binding to access the raw methods on
}

// ProverRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProverRegistryCallerRaw struct {
	Contract *ProverRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProverRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProverRegistryTransactorRaw struct {
	Contract *ProverRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProverRegistry creates a new instance of ProverRegistry, bound to a specific deployed contract.
func NewProverRegistry(address common.Address, backend bind.ContractBackend) (*ProverRegistry, error) {
	contract, err := bindProverRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProverRegistry{ProverRegistryCaller: ProverRegistryCaller{contract: contract}, ProverRegistryTransactor: ProverRegistryTransactor{contract: contract}, ProverRegistryFilterer: ProverRegistryFilterer{contract: contract}}, nil
}

// NewProverRegistryCaller creates a new read-only instance of ProverRegistry, bound to a specific deployed contract.
func NewProverRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProverRegistryCaller, error) {
	contract, err := bindProverRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryCaller{contract: contract}, nil
}

// NewProverRegistryTransactor creates a new write-only instance of ProverRegistry, bound to a specific deployed contract.
func NewProverRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProverRegistryTransactor, error) {
	contract, err := bindProverRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryTransactor{contract: contract}, nil
}

// NewProverRegistryFilterer creates a new log filterer instance of ProverRegistry, bound to a specific deployed contract.
func NewProverRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProverRegistryFilterer, error) {
	contract, err := bindProverRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryFilterer{contract: contract}, nil
}

// bindProverRegistry binds a generic wrapper to an already deployed contract.
func bindProverRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProverRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProverRegistry *ProverRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProverRegistry.Contract.ProverRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProverRegistry *ProverRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProverRegistry.Contract.ProverRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProverRegistry *ProverRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProverRegistry.Contract.ProverRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProverRegistry *ProverRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProverRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProverRegistry *ProverRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProverRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProverRegistry *ProverRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProverRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ProverRegistry *ProverRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ProverRegistry *ProverRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ProverRegistry.Contract.UPGRADEINTERFACEVERSION(&_ProverRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ProverRegistry *ProverRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ProverRegistry.Contract.UPGRADEINTERFACEVERSION(&_ProverRegistry.CallOpts)
}

// CalculateProverRegistrationDigestHash is a free data retrieval call binding the contract method 0xa60ca65a.
//
// Solidity: function calculateProverRegistrationDigestHash(address prover, address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ProverRegistry *ProverRegistryCaller) CalculateProverRegistrationDigestHash(opts *bind.CallOpts, prover common.Address, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "calculateProverRegistrationDigestHash", prover, operator, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateProverRegistrationDigestHash is a free data retrieval call binding the contract method 0xa60ca65a.
//
// Solidity: function calculateProverRegistrationDigestHash(address prover, address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ProverRegistry *ProverRegistrySession) CalculateProverRegistrationDigestHash(prover common.Address, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ProverRegistry.Contract.CalculateProverRegistrationDigestHash(&_ProverRegistry.CallOpts, prover, operator, salt, expiry)
}

// CalculateProverRegistrationDigestHash is a free data retrieval call binding the contract method 0xa60ca65a.
//
// Solidity: function calculateProverRegistrationDigestHash(address prover, address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ProverRegistry *ProverRegistryCallerSession) CalculateProverRegistrationDigestHash(prover common.Address, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ProverRegistry.Contract.CalculateProverRegistrationDigestHash(&_ProverRegistry.CallOpts, prover, operator, salt, expiry)
}

// ChainAdmin is a free data retrieval call binding the contract method 0x6478d8ed.
//
// Solidity: function chainAdmin() view returns(address)
func (_ProverRegistry *ProverRegistryCaller) ChainAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "chainAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainAdmin is a free data retrieval call binding the contract method 0x6478d8ed.
//
// Solidity: function chainAdmin() view returns(address)
func (_ProverRegistry *ProverRegistrySession) ChainAdmin() (common.Address, error) {
	return _ProverRegistry.Contract.ChainAdmin(&_ProverRegistry.CallOpts)
}

// ChainAdmin is a free data retrieval call binding the contract method 0x6478d8ed.
//
// Solidity: function chainAdmin() view returns(address)
func (_ProverRegistry *ProverRegistryCallerSession) ChainAdmin() (common.Address, error) {
	return _ProverRegistry.Contract.ChainAdmin(&_ProverRegistry.CallOpts)
}

// DePINOperatorProversStatus is a free data retrieval call binding the contract method 0x060e86cd.
//
// Solidity: function dePINOperatorProversStatus(address , address ) view returns(uint8)
func (_ProverRegistry *ProverRegistryCaller) DePINOperatorProversStatus(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (uint8, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "dePINOperatorProversStatus", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DePINOperatorProversStatus is a free data retrieval call binding the contract method 0x060e86cd.
//
// Solidity: function dePINOperatorProversStatus(address , address ) view returns(uint8)
func (_ProverRegistry *ProverRegistrySession) DePINOperatorProversStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _ProverRegistry.Contract.DePINOperatorProversStatus(&_ProverRegistry.CallOpts, arg0, arg1)
}

// DePINOperatorProversStatus is a free data retrieval call binding the contract method 0x060e86cd.
//
// Solidity: function dePINOperatorProversStatus(address , address ) view returns(uint8)
func (_ProverRegistry *ProverRegistryCallerSession) DePINOperatorProversStatus(arg0 common.Address, arg1 common.Address) (uint8, error) {
	return _ProverRegistry.Contract.DePINOperatorProversStatus(&_ProverRegistry.CallOpts, arg0, arg1)
}

// DePINOperatorsStatus is a free data retrieval call binding the contract method 0xe240309f.
//
// Solidity: function dePINOperatorsStatus(address ) view returns(uint8)
func (_ProverRegistry *ProverRegistryCaller) DePINOperatorsStatus(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "dePINOperatorsStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DePINOperatorsStatus is a free data retrieval call binding the contract method 0xe240309f.
//
// Solidity: function dePINOperatorsStatus(address ) view returns(uint8)
func (_ProverRegistry *ProverRegistrySession) DePINOperatorsStatus(arg0 common.Address) (uint8, error) {
	return _ProverRegistry.Contract.DePINOperatorsStatus(&_ProverRegistry.CallOpts, arg0)
}

// DePINOperatorsStatus is a free data retrieval call binding the contract method 0xe240309f.
//
// Solidity: function dePINOperatorsStatus(address ) view returns(uint8)
func (_ProverRegistry *ProverRegistryCallerSession) DePINOperatorsStatus(arg0 common.Address) (uint8, error) {
	return _ProverRegistry.Contract.DePINOperatorsStatus(&_ProverRegistry.CallOpts, arg0)
}

// DePINProverSaltUsed is a free data retrieval call binding the contract method 0x376ce2ba.
//
// Solidity: function dePINProverSaltUsed(address , bytes32 ) view returns(bool)
func (_ProverRegistry *ProverRegistryCaller) DePINProverSaltUsed(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "dePINProverSaltUsed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DePINProverSaltUsed is a free data retrieval call binding the contract method 0x376ce2ba.
//
// Solidity: function dePINProverSaltUsed(address , bytes32 ) view returns(bool)
func (_ProverRegistry *ProverRegistrySession) DePINProverSaltUsed(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ProverRegistry.Contract.DePINProverSaltUsed(&_ProverRegistry.CallOpts, arg0, arg1)
}

// DePINProverSaltUsed is a free data retrieval call binding the contract method 0x376ce2ba.
//
// Solidity: function dePINProverSaltUsed(address , bytes32 ) view returns(bool)
func (_ProverRegistry *ProverRegistryCallerSession) DePINProverSaltUsed(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _ProverRegistry.Contract.DePINProverSaltUsed(&_ProverRegistry.CallOpts, arg0, arg1)
}

// DePINProversStatus is a free data retrieval call binding the contract method 0xbbaf9bf4.
//
// Solidity: function dePINProversStatus(address ) view returns(uint8)
func (_ProverRegistry *ProverRegistryCaller) DePINProversStatus(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "dePINProversStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DePINProversStatus is a free data retrieval call binding the contract method 0xbbaf9bf4.
//
// Solidity: function dePINProversStatus(address ) view returns(uint8)
func (_ProverRegistry *ProverRegistrySession) DePINProversStatus(arg0 common.Address) (uint8, error) {
	return _ProverRegistry.Contract.DePINProversStatus(&_ProverRegistry.CallOpts, arg0)
}

// DePINProversStatus is a free data retrieval call binding the contract method 0xbbaf9bf4.
//
// Solidity: function dePINProversStatus(address ) view returns(uint8)
func (_ProverRegistry *ProverRegistryCallerSession) DePINProversStatus(arg0 common.Address) (uint8, error) {
	return _ProverRegistry.Contract.DePINProversStatus(&_ProverRegistry.CallOpts, arg0)
}

// DomainSeperator is a free data retrieval call binding the contract method 0xcb73ac56.
//
// Solidity: function domainSeperator() view returns(bytes32)
func (_ProverRegistry *ProverRegistryCaller) DomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "domainSeperator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeperator is a free data retrieval call binding the contract method 0xcb73ac56.
//
// Solidity: function domainSeperator() view returns(bytes32)
func (_ProverRegistry *ProverRegistrySession) DomainSeperator() ([32]byte, error) {
	return _ProverRegistry.Contract.DomainSeperator(&_ProverRegistry.CallOpts)
}

// DomainSeperator is a free data retrieval call binding the contract method 0xcb73ac56.
//
// Solidity: function domainSeperator() view returns(bytes32)
func (_ProverRegistry *ProverRegistryCallerSession) DomainSeperator() ([32]byte, error) {
	return _ProverRegistry.Contract.DomainSeperator(&_ProverRegistry.CallOpts)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ProverRegistry *ProverRegistryCaller) GetOperatorStatus(opts *bind.CallOpts, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "getOperatorStatus", operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ProverRegistry *ProverRegistrySession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ProverRegistry.Contract.GetOperatorStatus(&_ProverRegistry.CallOpts, operator)
}

// GetOperatorStatus is a free data retrieval call binding the contract method 0xfd39105a.
//
// Solidity: function getOperatorStatus(address operator) view returns(uint8)
func (_ProverRegistry *ProverRegistryCallerSession) GetOperatorStatus(operator common.Address) (uint8, error) {
	return _ProverRegistry.Contract.GetOperatorStatus(&_ProverRegistry.CallOpts, operator)
}

// GetProverStatus is a free data retrieval call binding the contract method 0x6ea988e0.
//
// Solidity: function getProverStatus(address prover) view returns(uint8)
func (_ProverRegistry *ProverRegistryCaller) GetProverStatus(opts *bind.CallOpts, prover common.Address) (uint8, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "getProverStatus", prover)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProverStatus is a free data retrieval call binding the contract method 0x6ea988e0.
//
// Solidity: function getProverStatus(address prover) view returns(uint8)
func (_ProverRegistry *ProverRegistrySession) GetProverStatus(prover common.Address) (uint8, error) {
	return _ProverRegistry.Contract.GetProverStatus(&_ProverRegistry.CallOpts, prover)
}

// GetProverStatus is a free data retrieval call binding the contract method 0x6ea988e0.
//
// Solidity: function getProverStatus(address prover) view returns(uint8)
func (_ProverRegistry *ProverRegistryCallerSession) GetProverStatus(prover common.Address) (uint8, error) {
	return _ProverRegistry.Contract.GetProverStatus(&_ProverRegistry.CallOpts, prover)
}

// IsRegisteredProver is a free data retrieval call binding the contract method 0x58603c75.
//
// Solidity: function isRegisteredProver(address prover, address operator) view returns(bool)
func (_ProverRegistry *ProverRegistryCaller) IsRegisteredProver(opts *bind.CallOpts, prover common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "isRegisteredProver", prover, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredProver is a free data retrieval call binding the contract method 0x58603c75.
//
// Solidity: function isRegisteredProver(address prover, address operator) view returns(bool)
func (_ProverRegistry *ProverRegistrySession) IsRegisteredProver(prover common.Address, operator common.Address) (bool, error) {
	return _ProverRegistry.Contract.IsRegisteredProver(&_ProverRegistry.CallOpts, prover, operator)
}

// IsRegisteredProver is a free data retrieval call binding the contract method 0x58603c75.
//
// Solidity: function isRegisteredProver(address prover, address operator) view returns(bool)
func (_ProverRegistry *ProverRegistryCallerSession) IsRegisteredProver(prover common.Address, operator common.Address) (bool, error) {
	return _ProverRegistry.Contract.IsRegisteredProver(&_ProverRegistry.CallOpts, prover, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProverRegistry *ProverRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProverRegistry *ProverRegistrySession) Owner() (common.Address, error) {
	return _ProverRegistry.Contract.Owner(&_ProverRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProverRegistry *ProverRegistryCallerSession) Owner() (common.Address, error) {
	return _ProverRegistry.Contract.Owner(&_ProverRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ProverRegistry *ProverRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ProverRegistry *ProverRegistrySession) Paused() (bool, error) {
	return _ProverRegistry.Contract.Paused(&_ProverRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ProverRegistry *ProverRegistryCallerSession) Paused() (bool, error) {
	return _ProverRegistry.Contract.Paused(&_ProverRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProverRegistry *ProverRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProverRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProverRegistry *ProverRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _ProverRegistry.Contract.ProxiableUUID(&_ProverRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProverRegistry *ProverRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ProverRegistry.Contract.ProxiableUUID(&_ProverRegistry.CallOpts)
}

// DeRegisterProver is a paid mutator transaction binding the contract method 0xd00ce3d6.
//
// Solidity: function deRegisterProver(address prover) returns()
func (_ProverRegistry *ProverRegistryTransactor) DeRegisterProver(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "deRegisterProver", prover)
}

// DeRegisterProver is a paid mutator transaction binding the contract method 0xd00ce3d6.
//
// Solidity: function deRegisterProver(address prover) returns()
func (_ProverRegistry *ProverRegistrySession) DeRegisterProver(prover common.Address) (*types.Transaction, error) {
	return _ProverRegistry.Contract.DeRegisterProver(&_ProverRegistry.TransactOpts, prover)
}

// DeRegisterProver is a paid mutator transaction binding the contract method 0xd00ce3d6.
//
// Solidity: function deRegisterProver(address prover) returns()
func (_ProverRegistry *ProverRegistryTransactorSession) DeRegisterProver(prover common.Address) (*types.Transaction, error) {
	return _ProverRegistry.Contract.DeRegisterProver(&_ProverRegistry.TransactOpts, prover)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProverRegistry *ProverRegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProverRegistry *ProverRegistrySession) Initialize() (*types.Transaction, error) {
	return _ProverRegistry.Contract.Initialize(&_ProverRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProverRegistry *ProverRegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _ProverRegistry.Contract.Initialize(&_ProverRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ProverRegistry *ProverRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ProverRegistry *ProverRegistrySession) Pause() (*types.Transaction, error) {
	return _ProverRegistry.Contract.Pause(&_ProverRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ProverRegistry *ProverRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _ProverRegistry.Contract.Pause(&_ProverRegistry.TransactOpts)
}

// RegisterProver is a paid mutator transaction binding the contract method 0xb969f136.
//
// Solidity: function registerProver(address prover, bytes32 salt, uint256 expiry, bytes proverSignature) returns()
func (_ProverRegistry *ProverRegistryTransactor) RegisterProver(opts *bind.TransactOpts, prover common.Address, salt [32]byte, expiry *big.Int, proverSignature []byte) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "registerProver", prover, salt, expiry, proverSignature)
}

// RegisterProver is a paid mutator transaction binding the contract method 0xb969f136.
//
// Solidity: function registerProver(address prover, bytes32 salt, uint256 expiry, bytes proverSignature) returns()
func (_ProverRegistry *ProverRegistrySession) RegisterProver(prover common.Address, salt [32]byte, expiry *big.Int, proverSignature []byte) (*types.Transaction, error) {
	return _ProverRegistry.Contract.RegisterProver(&_ProverRegistry.TransactOpts, prover, salt, expiry, proverSignature)
}

// RegisterProver is a paid mutator transaction binding the contract method 0xb969f136.
//
// Solidity: function registerProver(address prover, bytes32 salt, uint256 expiry, bytes proverSignature) returns()
func (_ProverRegistry *ProverRegistryTransactorSession) RegisterProver(prover common.Address, salt [32]byte, expiry *big.Int, proverSignature []byte) (*types.Transaction, error) {
	return _ProverRegistry.Contract.RegisterProver(&_ProverRegistry.TransactOpts, prover, salt, expiry, proverSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProverRegistry *ProverRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProverRegistry *ProverRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ProverRegistry.Contract.RenounceOwnership(&_ProverRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProverRegistry *ProverRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProverRegistry.Contract.RenounceOwnership(&_ProverRegistry.TransactOpts)
}

// SetChainAdmin is a paid mutator transaction binding the contract method 0xf9c182d8.
//
// Solidity: function setChainAdmin(address _chainAdmin) returns()
func (_ProverRegistry *ProverRegistryTransactor) SetChainAdmin(opts *bind.TransactOpts, _chainAdmin common.Address) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "setChainAdmin", _chainAdmin)
}

// SetChainAdmin is a paid mutator transaction binding the contract method 0xf9c182d8.
//
// Solidity: function setChainAdmin(address _chainAdmin) returns()
func (_ProverRegistry *ProverRegistrySession) SetChainAdmin(_chainAdmin common.Address) (*types.Transaction, error) {
	return _ProverRegistry.Contract.SetChainAdmin(&_ProverRegistry.TransactOpts, _chainAdmin)
}

// SetChainAdmin is a paid mutator transaction binding the contract method 0xf9c182d8.
//
// Solidity: function setChainAdmin(address _chainAdmin) returns()
func (_ProverRegistry *ProverRegistryTransactorSession) SetChainAdmin(_chainAdmin common.Address) (*types.Transaction, error) {
	return _ProverRegistry.Contract.SetChainAdmin(&_ProverRegistry.TransactOpts, _chainAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProverRegistry *ProverRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProverRegistry *ProverRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProverRegistry.Contract.TransferOwnership(&_ProverRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProverRegistry *ProverRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProverRegistry.Contract.TransferOwnership(&_ProverRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ProverRegistry *ProverRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ProverRegistry *ProverRegistrySession) Unpause() (*types.Transaction, error) {
	return _ProverRegistry.Contract.Unpause(&_ProverRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ProverRegistry *ProverRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _ProverRegistry.Contract.Unpause(&_ProverRegistry.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProverRegistry *ProverRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProverRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProverRegistry *ProverRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProverRegistry.Contract.UpgradeToAndCall(&_ProverRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProverRegistry *ProverRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProverRegistry.Contract.UpgradeToAndCall(&_ProverRegistry.TransactOpts, newImplementation, data)
}

// ProverRegistryChainAdminInitializedIterator is returned from FilterChainAdminInitialized and is used to iterate over the raw logs and unpacked data for ChainAdminInitialized events raised by the ProverRegistry contract.
type ProverRegistryChainAdminInitializedIterator struct {
	Event *ProverRegistryChainAdminInitialized // Event containing the contract specifics and raw log

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
func (it *ProverRegistryChainAdminInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryChainAdminInitialized)
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
		it.Event = new(ProverRegistryChainAdminInitialized)
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
func (it *ProverRegistryChainAdminInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryChainAdminInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryChainAdminInitialized represents a ChainAdminInitialized event raised by the ProverRegistry contract.
type ProverRegistryChainAdminInitialized struct {
	ChainAdmin common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChainAdminInitialized is a free log retrieval operation binding the contract event 0xd0dfe49bc354505424888f6141a5274c54ee2a9d6adac7d5187916e1c0aea940.
//
// Solidity: event ChainAdminInitialized(address _chainAdmin)
func (_ProverRegistry *ProverRegistryFilterer) FilterChainAdminInitialized(opts *bind.FilterOpts) (*ProverRegistryChainAdminInitializedIterator, error) {

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "ChainAdminInitialized")
	if err != nil {
		return nil, err
	}
	return &ProverRegistryChainAdminInitializedIterator{contract: _ProverRegistry.contract, event: "ChainAdminInitialized", logs: logs, sub: sub}, nil
}

// WatchChainAdminInitialized is a free log subscription operation binding the contract event 0xd0dfe49bc354505424888f6141a5274c54ee2a9d6adac7d5187916e1c0aea940.
//
// Solidity: event ChainAdminInitialized(address _chainAdmin)
func (_ProverRegistry *ProverRegistryFilterer) WatchChainAdminInitialized(opts *bind.WatchOpts, sink chan<- *ProverRegistryChainAdminInitialized) (event.Subscription, error) {

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "ChainAdminInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryChainAdminInitialized)
				if err := _ProverRegistry.contract.UnpackLog(event, "ChainAdminInitialized", log); err != nil {
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

// ParseChainAdminInitialized is a log parse operation binding the contract event 0xd0dfe49bc354505424888f6141a5274c54ee2a9d6adac7d5187916e1c0aea940.
//
// Solidity: event ChainAdminInitialized(address _chainAdmin)
func (_ProverRegistry *ProverRegistryFilterer) ParseChainAdminInitialized(log types.Log) (*ProverRegistryChainAdminInitialized, error) {
	event := new(ProverRegistryChainAdminInitialized)
	if err := _ProverRegistry.contract.UnpackLog(event, "ChainAdminInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryDePINProverDeregisteredIterator is returned from FilterDePINProverDeregistered and is used to iterate over the raw logs and unpacked data for DePINProverDeregistered events raised by the ProverRegistry contract.
type ProverRegistryDePINProverDeregisteredIterator struct {
	Event *ProverRegistryDePINProverDeregistered // Event containing the contract specifics and raw log

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
func (it *ProverRegistryDePINProverDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryDePINProverDeregistered)
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
		it.Event = new(ProverRegistryDePINProverDeregistered)
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
func (it *ProverRegistryDePINProverDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryDePINProverDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryDePINProverDeregistered represents a DePINProverDeregistered event raised by the ProverRegistry contract.
type ProverRegistryDePINProverDeregistered struct {
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDePINProverDeregistered is a free log retrieval operation binding the contract event 0x03fa3d96999d2eb4f9246112eff532300dadca37acdb46130162c9231b83ed14.
//
// Solidity: event DePINProverDeregistered(address indexed prover)
func (_ProverRegistry *ProverRegistryFilterer) FilterDePINProverDeregistered(opts *bind.FilterOpts, prover []common.Address) (*ProverRegistryDePINProverDeregisteredIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "DePINProverDeregistered", proverRule)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryDePINProverDeregisteredIterator{contract: _ProverRegistry.contract, event: "DePINProverDeregistered", logs: logs, sub: sub}, nil
}

// WatchDePINProverDeregistered is a free log subscription operation binding the contract event 0x03fa3d96999d2eb4f9246112eff532300dadca37acdb46130162c9231b83ed14.
//
// Solidity: event DePINProverDeregistered(address indexed prover)
func (_ProverRegistry *ProverRegistryFilterer) WatchDePINProverDeregistered(opts *bind.WatchOpts, sink chan<- *ProverRegistryDePINProverDeregistered, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "DePINProverDeregistered", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryDePINProverDeregistered)
				if err := _ProverRegistry.contract.UnpackLog(event, "DePINProverDeregistered", log); err != nil {
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

// ParseDePINProverDeregistered is a log parse operation binding the contract event 0x03fa3d96999d2eb4f9246112eff532300dadca37acdb46130162c9231b83ed14.
//
// Solidity: event DePINProverDeregistered(address indexed prover)
func (_ProverRegistry *ProverRegistryFilterer) ParseDePINProverDeregistered(log types.Log) (*ProverRegistryDePINProverDeregistered, error) {
	event := new(ProverRegistryDePINProverDeregistered)
	if err := _ProverRegistry.contract.UnpackLog(event, "DePINProverDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryDePINProverParticipationKeyRegisteredIterator is returned from FilterDePINProverParticipationKeyRegistered and is used to iterate over the raw logs and unpacked data for DePINProverParticipationKeyRegistered events raised by the ProverRegistry contract.
type ProverRegistryDePINProverParticipationKeyRegisteredIterator struct {
	Event *ProverRegistryDePINProverParticipationKeyRegistered // Event containing the contract specifics and raw log

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
func (it *ProverRegistryDePINProverParticipationKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryDePINProverParticipationKeyRegistered)
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
		it.Event = new(ProverRegistryDePINProverParticipationKeyRegistered)
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
func (it *ProverRegistryDePINProverParticipationKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryDePINProverParticipationKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryDePINProverParticipationKeyRegistered represents a DePINProverParticipationKeyRegistered event raised by the ProverRegistry contract.
type ProverRegistryDePINProverParticipationKeyRegistered struct {
	Prover           common.Address
	ParticipationKey common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDePINProverParticipationKeyRegistered is a free log retrieval operation binding the contract event 0x31341c94d68af95d7ad4a26e20480da3fd0def077e8b9a533b7b9044c4247623.
//
// Solidity: event DePINProverParticipationKeyRegistered(address indexed prover, address indexed participationKey)
func (_ProverRegistry *ProverRegistryFilterer) FilterDePINProverParticipationKeyRegistered(opts *bind.FilterOpts, prover []common.Address, participationKey []common.Address) (*ProverRegistryDePINProverParticipationKeyRegisteredIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var participationKeyRule []interface{}
	for _, participationKeyItem := range participationKey {
		participationKeyRule = append(participationKeyRule, participationKeyItem)
	}

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "DePINProverParticipationKeyRegistered", proverRule, participationKeyRule)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryDePINProverParticipationKeyRegisteredIterator{contract: _ProverRegistry.contract, event: "DePINProverParticipationKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchDePINProverParticipationKeyRegistered is a free log subscription operation binding the contract event 0x31341c94d68af95d7ad4a26e20480da3fd0def077e8b9a533b7b9044c4247623.
//
// Solidity: event DePINProverParticipationKeyRegistered(address indexed prover, address indexed participationKey)
func (_ProverRegistry *ProverRegistryFilterer) WatchDePINProverParticipationKeyRegistered(opts *bind.WatchOpts, sink chan<- *ProverRegistryDePINProverParticipationKeyRegistered, prover []common.Address, participationKey []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var participationKeyRule []interface{}
	for _, participationKeyItem := range participationKey {
		participationKeyRule = append(participationKeyRule, participationKeyItem)
	}

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "DePINProverParticipationKeyRegistered", proverRule, participationKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryDePINProverParticipationKeyRegistered)
				if err := _ProverRegistry.contract.UnpackLog(event, "DePINProverParticipationKeyRegistered", log); err != nil {
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

// ParseDePINProverParticipationKeyRegistered is a log parse operation binding the contract event 0x31341c94d68af95d7ad4a26e20480da3fd0def077e8b9a533b7b9044c4247623.
//
// Solidity: event DePINProverParticipationKeyRegistered(address indexed prover, address indexed participationKey)
func (_ProverRegistry *ProverRegistryFilterer) ParseDePINProverParticipationKeyRegistered(log types.Log) (*ProverRegistryDePINProverParticipationKeyRegistered, error) {
	event := new(ProverRegistryDePINProverParticipationKeyRegistered)
	if err := _ProverRegistry.contract.UnpackLog(event, "DePINProverParticipationKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryDePINProverRegisteredIterator is returned from FilterDePINProverRegistered and is used to iterate over the raw logs and unpacked data for DePINProverRegistered events raised by the ProverRegistry contract.
type ProverRegistryDePINProverRegisteredIterator struct {
	Event *ProverRegistryDePINProverRegistered // Event containing the contract specifics and raw log

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
func (it *ProverRegistryDePINProverRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryDePINProverRegistered)
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
		it.Event = new(ProverRegistryDePINProverRegistered)
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
func (it *ProverRegistryDePINProverRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryDePINProverRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryDePINProverRegistered represents a DePINProverRegistered event raised by the ProverRegistry contract.
type ProverRegistryDePINProverRegistered struct {
	Prover       common.Address
	Operator     common.Address
	ProverStatus uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDePINProverRegistered is a free log retrieval operation binding the contract event 0xf400e45763c65091f6a4b9ea95f7e17ac4e039a4b9498e5c7bfc2a46bdc2503a.
//
// Solidity: event DePINProverRegistered(address indexed prover, address indexed operator, uint8 proverStatus)
func (_ProverRegistry *ProverRegistryFilterer) FilterDePINProverRegistered(opts *bind.FilterOpts, prover []common.Address, operator []common.Address) (*ProverRegistryDePINProverRegisteredIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "DePINProverRegistered", proverRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryDePINProverRegisteredIterator{contract: _ProverRegistry.contract, event: "DePINProverRegistered", logs: logs, sub: sub}, nil
}

// WatchDePINProverRegistered is a free log subscription operation binding the contract event 0xf400e45763c65091f6a4b9ea95f7e17ac4e039a4b9498e5c7bfc2a46bdc2503a.
//
// Solidity: event DePINProverRegistered(address indexed prover, address indexed operator, uint8 proverStatus)
func (_ProverRegistry *ProverRegistryFilterer) WatchDePINProverRegistered(opts *bind.WatchOpts, sink chan<- *ProverRegistryDePINProverRegistered, prover []common.Address, operator []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "DePINProverRegistered", proverRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryDePINProverRegistered)
				if err := _ProverRegistry.contract.UnpackLog(event, "DePINProverRegistered", log); err != nil {
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

// ParseDePINProverRegistered is a log parse operation binding the contract event 0xf400e45763c65091f6a4b9ea95f7e17ac4e039a4b9498e5c7bfc2a46bdc2503a.
//
// Solidity: event DePINProverRegistered(address indexed prover, address indexed operator, uint8 proverStatus)
func (_ProverRegistry *ProverRegistryFilterer) ParseDePINProverRegistered(log types.Log) (*ProverRegistryDePINProverRegistered, error) {
	event := new(ProverRegistryDePINProverRegistered)
	if err := _ProverRegistry.contract.UnpackLog(event, "DePINProverRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProverRegistry contract.
type ProverRegistryInitializedIterator struct {
	Event *ProverRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ProverRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryInitialized)
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
		it.Event = new(ProverRegistryInitialized)
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
func (it *ProverRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryInitialized represents a Initialized event raised by the ProverRegistry contract.
type ProverRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ProverRegistry *ProverRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProverRegistryInitializedIterator, error) {

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProverRegistryInitializedIterator{contract: _ProverRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ProverRegistry *ProverRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProverRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryInitialized)
				if err := _ProverRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ProverRegistry *ProverRegistryFilterer) ParseInitialized(log types.Log) (*ProverRegistryInitialized, error) {
	event := new(ProverRegistryInitialized)
	if err := _ProverRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProverRegistry contract.
type ProverRegistryOwnershipTransferredIterator struct {
	Event *ProverRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProverRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryOwnershipTransferred)
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
		it.Event = new(ProverRegistryOwnershipTransferred)
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
func (it *ProverRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ProverRegistry contract.
type ProverRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProverRegistry *ProverRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProverRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryOwnershipTransferredIterator{contract: _ProverRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProverRegistry *ProverRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProverRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryOwnershipTransferred)
				if err := _ProverRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProverRegistry *ProverRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ProverRegistryOwnershipTransferred, error) {
	event := new(ProverRegistryOwnershipTransferred)
	if err := _ProverRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ProverRegistry contract.
type ProverRegistryPausedIterator struct {
	Event *ProverRegistryPaused // Event containing the contract specifics and raw log

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
func (it *ProverRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryPaused)
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
		it.Event = new(ProverRegistryPaused)
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
func (it *ProverRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryPaused represents a Paused event raised by the ProverRegistry contract.
type ProverRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ProverRegistry *ProverRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*ProverRegistryPausedIterator, error) {

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ProverRegistryPausedIterator{contract: _ProverRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ProverRegistry *ProverRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ProverRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryPaused)
				if err := _ProverRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ProverRegistry *ProverRegistryFilterer) ParsePaused(log types.Log) (*ProverRegistryPaused, error) {
	event := new(ProverRegistryPaused)
	if err := _ProverRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ProverRegistry contract.
type ProverRegistryUnpausedIterator struct {
	Event *ProverRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *ProverRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryUnpaused)
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
		it.Event = new(ProverRegistryUnpaused)
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
func (it *ProverRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryUnpaused represents a Unpaused event raised by the ProverRegistry contract.
type ProverRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ProverRegistry *ProverRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ProverRegistryUnpausedIterator, error) {

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ProverRegistryUnpausedIterator{contract: _ProverRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ProverRegistry *ProverRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ProverRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryUnpaused)
				if err := _ProverRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ProverRegistry *ProverRegistryFilterer) ParseUnpaused(log types.Log) (*ProverRegistryUnpaused, error) {
	event := new(ProverRegistryUnpaused)
	if err := _ProverRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProverRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ProverRegistry contract.
type ProverRegistryUpgradedIterator struct {
	Event *ProverRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ProverRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProverRegistryUpgraded)
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
		it.Event = new(ProverRegistryUpgraded)
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
func (it *ProverRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProverRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProverRegistryUpgraded represents a Upgraded event raised by the ProverRegistry contract.
type ProverRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProverRegistry *ProverRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProverRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProverRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProverRegistryUpgradedIterator{contract: _ProverRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProverRegistry *ProverRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProverRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProverRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProverRegistryUpgraded)
				if err := _ProverRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ProverRegistry *ProverRegistryFilterer) ParseUpgraded(log types.Log) (*ProverRegistryUpgraded, error) {
	event := new(ProverRegistryUpgraded)
	if err := _ProverRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
