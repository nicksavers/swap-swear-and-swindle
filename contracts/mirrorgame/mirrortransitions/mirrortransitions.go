// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package mirrortransitions

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MirrorTransistionsABI is the input ABI used to generate the binding from.
const MirrorTransistionsABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getInitialStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"transitions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"startGracePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"witnesses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"witnessState\",\"type\":\"uint8\"},{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"gracePeriods\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"expired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"paymentValidatorContract\",\"type\":\"address\"},{\"name\":\"ENSMirrotValidatorContract\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// MirrorTransistionsBin is the compiled bytecode used for deploying new contracts.
const MirrorTransistionsBin = `0x6060604052341561000f57600080fd5b60405160408061063083398101604052808051919060200180519150505b8160008060025b60010160ff1660ff16815260200190815260200160002060006101000a815481600160a060020a030219169083600160a060020a031602179055508060008060028081111561007f57fe5b600290810160ff16825260208201929092526040016000208054600160a060020a031916600160a060020a0393909316929092179091555b60020160016000805b60ff168152602081019190915260400160009081209060025b60010160ff90811682526020820192909252604001600020805460ff19169290911691909117905560025b60016000815b60ff168152602081019190915260400160009081209060025b600190810160ff908116835260208301939093526040909101600020805460ff1916939092169290921790555b60016000805b60ff168152602081019190915260400160009081209060025b600290810160ff908116835260208301939093526040909101600020805460ff1916939092169290921790555b60016000815b60ff168152602081019190915260400160009081209060025b60020160ff1660ff16815260200190815260200160002060006101000a81548160ff021916908360ff1602179055506023600260006002808111156101fc57fe5b60010160ff1660ff1681526020019081526020016000208190555060236002600060028081111561022957fe5b60020160ff1660ff168152602001908152602001600020819055505b50505b6103d9806102576000396000f3006060604052361561008b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307e5ae5f81146100905780639c2a1029146100b9578063a7edcbe4146100ee578063b2400f091461011e578063d39eb30114610160578063db06156314610195578063fa778a0f146101d7578063fcc0a2ae14610202575b600080fd5b341561009b57600080fd5b6100a3610232565b60405160ff909116815260200160405180910390f35b34156100c457600080fd5b6100a360ff60043581169060243516610240565b60405160ff909116815260200160405180910390f35b34156100f957600080fd5b61010a60043560ff60243516610260565b604051901515815260200160405180910390f35b341561012957600080fd5b61013760ff600435166102b4565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561016b57600080fd5b6100a360ff600435811690602435166102dc565b60405160ff909116815260200160405180910390f35b34156101a057600080fd5b61013760ff60043516610305565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101e257600080fd5b6101f060ff60043516610334565b60405190815260200160405180910390f35b341561020d57600080fd5b61010a60043560ff60243516610346565b604051901515815260200160405180910390f35b600060025b60010190505b90565b600160209081526000928352604080842090915290825290205460ff1681565b600082815260036020908152604080832060ff8516845290915281205415156102aa5750600082815260036020908152604080832060ff85168452909152902043905560016102ae565b5060005b92915050565b60006020819052908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60ff80831660009081526001602090815260408083208585168452909152902054165b92915050565b60ff811660009081526020819052604090205473ffffffffffffffffffffffffffffffffffffffff165b919050565b60026020526000908152604090205481565b600082815260036020908152604080832060ff85168452909152812054156102aa5760ff82166000818152600260209081526040808320548784526003835281842094845293909152902054430311156102aa575060016102ae565b5b5060005b929150505600a165627a7a7230582089c6597772c87b1cd36779ee3a054bc07677d644f57c9af5d129c61e3d2fac8c0029`

// DeployMirrorTransistions deploys a new Ethereum contract, binding an instance of MirrorTransistions to it.
func DeployMirrorTransistions(auth *bind.TransactOpts, backend bind.ContractBackend, paymentValidatorContract common.Address, ENSMirrotValidatorContract common.Address) (common.Address, *types.Transaction, *MirrorTransistions, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorTransistionsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MirrorTransistionsBin), backend, paymentValidatorContract, ENSMirrotValidatorContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MirrorTransistions{MirrorTransistionsCaller: MirrorTransistionsCaller{contract: contract}, MirrorTransistionsTransactor: MirrorTransistionsTransactor{contract: contract}}, nil
}

// MirrorTransistions is an auto generated Go binding around an Ethereum contract.
type MirrorTransistions struct {
	MirrorTransistionsCaller     // Read-only binding to the contract
	MirrorTransistionsTransactor // Write-only binding to the contract
}

// MirrorTransistionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MirrorTransistionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorTransistionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MirrorTransistionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorTransistionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MirrorTransistionsSession struct {
	Contract     *MirrorTransistions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MirrorTransistionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MirrorTransistionsCallerSession struct {
	Contract *MirrorTransistionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MirrorTransistionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MirrorTransistionsTransactorSession struct {
	Contract     *MirrorTransistionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MirrorTransistionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MirrorTransistionsRaw struct {
	Contract *MirrorTransistions // Generic contract binding to access the raw methods on
}

// MirrorTransistionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MirrorTransistionsCallerRaw struct {
	Contract *MirrorTransistionsCaller // Generic read-only contract binding to access the raw methods on
}

// MirrorTransistionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MirrorTransistionsTransactorRaw struct {
	Contract *MirrorTransistionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMirrorTransistions creates a new instance of MirrorTransistions, bound to a specific deployed contract.
func NewMirrorTransistions(address common.Address, backend bind.ContractBackend) (*MirrorTransistions, error) {
	contract, err := bindMirrorTransistions(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MirrorTransistions{MirrorTransistionsCaller: MirrorTransistionsCaller{contract: contract}, MirrorTransistionsTransactor: MirrorTransistionsTransactor{contract: contract}}, nil
}

// NewMirrorTransistionsCaller creates a new read-only instance of MirrorTransistions, bound to a specific deployed contract.
func NewMirrorTransistionsCaller(address common.Address, caller bind.ContractCaller) (*MirrorTransistionsCaller, error) {
	contract, err := bindMirrorTransistions(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorTransistionsCaller{contract: contract}, nil
}

// NewMirrorTransistionsTransactor creates a new write-only instance of MirrorTransistions, bound to a specific deployed contract.
func NewMirrorTransistionsTransactor(address common.Address, transactor bind.ContractTransactor) (*MirrorTransistionsTransactor, error) {
	contract, err := bindMirrorTransistions(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MirrorTransistionsTransactor{contract: contract}, nil
}

// bindMirrorTransistions binds a generic wrapper to an already deployed contract.
func bindMirrorTransistions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorTransistionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorTransistions *MirrorTransistionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MirrorTransistions.Contract.MirrorTransistionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorTransistions *MirrorTransistionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.MirrorTransistionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorTransistions *MirrorTransistionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.MirrorTransistionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorTransistions *MirrorTransistionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MirrorTransistions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorTransistions *MirrorTransistionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorTransistions *MirrorTransistionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.contract.Transact(opts, method, params...)
}

// GracePeriods is a free data retrieval call binding the contract method 0xfa778a0f.
//
// Solidity: function gracePeriods( uint8) constant returns(uint256)
func (_MirrorTransistions *MirrorTransistionsCaller) GracePeriods(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MirrorTransistions.contract.Call(opts, out, "gracePeriods", arg0)
	return *ret0, err
}

// GracePeriods is a free data retrieval call binding the contract method 0xfa778a0f.
//
// Solidity: function gracePeriods( uint8) constant returns(uint256)
func (_MirrorTransistions *MirrorTransistionsSession) GracePeriods(arg0 uint8) (*big.Int, error) {
	return _MirrorTransistions.Contract.GracePeriods(&_MirrorTransistions.CallOpts, arg0)
}

// GracePeriods is a free data retrieval call binding the contract method 0xfa778a0f.
//
// Solidity: function gracePeriods( uint8) constant returns(uint256)
func (_MirrorTransistions *MirrorTransistionsCallerSession) GracePeriods(arg0 uint8) (*big.Int, error) {
	return _MirrorTransistions.Contract.GracePeriods(&_MirrorTransistions.CallOpts, arg0)
}

// Transitions is a free data retrieval call binding the contract method 0x9c2a1029.
//
// Solidity: function transitions( uint8,  uint8) constant returns(uint8)
func (_MirrorTransistions *MirrorTransistionsCaller) Transitions(opts *bind.CallOpts, arg0 uint8, arg1 uint8) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MirrorTransistions.contract.Call(opts, out, "transitions", arg0, arg1)
	return *ret0, err
}

// Transitions is a free data retrieval call binding the contract method 0x9c2a1029.
//
// Solidity: function transitions( uint8,  uint8) constant returns(uint8)
func (_MirrorTransistions *MirrorTransistionsSession) Transitions(arg0 uint8, arg1 uint8) (uint8, error) {
	return _MirrorTransistions.Contract.Transitions(&_MirrorTransistions.CallOpts, arg0, arg1)
}

// Transitions is a free data retrieval call binding the contract method 0x9c2a1029.
//
// Solidity: function transitions( uint8,  uint8) constant returns(uint8)
func (_MirrorTransistions *MirrorTransistionsCallerSession) Transitions(arg0 uint8, arg1 uint8) (uint8, error) {
	return _MirrorTransistions.Contract.Transitions(&_MirrorTransistions.CallOpts, arg0, arg1)
}

// Witnesses is a free data retrieval call binding the contract method 0xb2400f09.
//
// Solidity: function witnesses( uint8) constant returns(address)
func (_MirrorTransistions *MirrorTransistionsCaller) Witnesses(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MirrorTransistions.contract.Call(opts, out, "witnesses", arg0)
	return *ret0, err
}

// Witnesses is a free data retrieval call binding the contract method 0xb2400f09.
//
// Solidity: function witnesses( uint8) constant returns(address)
func (_MirrorTransistions *MirrorTransistionsSession) Witnesses(arg0 uint8) (common.Address, error) {
	return _MirrorTransistions.Contract.Witnesses(&_MirrorTransistions.CallOpts, arg0)
}

// Witnesses is a free data retrieval call binding the contract method 0xb2400f09.
//
// Solidity: function witnesses( uint8) constant returns(address)
func (_MirrorTransistions *MirrorTransistionsCallerSession) Witnesses(arg0 uint8) (common.Address, error) {
	return _MirrorTransistions.Contract.Witnesses(&_MirrorTransistions.CallOpts, arg0)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_MirrorTransistions *MirrorTransistionsTransactor) Expired(opts *bind.TransactOpts, caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorTransistions.contract.Transact(opts, "expired", caseId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_MirrorTransistions *MirrorTransistionsSession) Expired(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.Expired(&_MirrorTransistions.TransactOpts, caseId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(caseId bytes32, status uint8) returns(bool)
func (_MirrorTransistions *MirrorTransistionsTransactorSession) Expired(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.Expired(&_MirrorTransistions.TransactOpts, caseId, status)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_MirrorTransistions *MirrorTransistionsTransactor) GetInitialStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorTransistions.contract.Transact(opts, "getInitialStatus")
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_MirrorTransistions *MirrorTransistionsSession) GetInitialStatus() (*types.Transaction, error) {
	return _MirrorTransistions.Contract.GetInitialStatus(&_MirrorTransistions.TransactOpts)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_MirrorTransistions *MirrorTransistionsTransactorSession) GetInitialStatus() (*types.Transaction, error) {
	return _MirrorTransistions.Contract.GetInitialStatus(&_MirrorTransistions.TransactOpts)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_MirrorTransistions *MirrorTransistionsTransactor) GetStatus(opts *bind.TransactOpts, witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorTransistions.contract.Transact(opts, "getStatus", witnessState, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_MirrorTransistions *MirrorTransistionsSession) GetStatus(witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.GetStatus(&_MirrorTransistions.TransactOpts, witnessState, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_MirrorTransistions *MirrorTransistionsTransactorSession) GetStatus(witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.GetStatus(&_MirrorTransistions.TransactOpts, witnessState, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_MirrorTransistions *MirrorTransistionsTransactor) GetWitness(opts *bind.TransactOpts, trialStatus uint8) (*types.Transaction, error) {
	return _MirrorTransistions.contract.Transact(opts, "getWitness", trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_MirrorTransistions *MirrorTransistionsSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.GetWitness(&_MirrorTransistions.TransactOpts, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_MirrorTransistions *MirrorTransistionsTransactorSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.GetWitness(&_MirrorTransistions.TransactOpts, trialStatus)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_MirrorTransistions *MirrorTransistionsTransactor) StartGracePeriod(opts *bind.TransactOpts, caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorTransistions.contract.Transact(opts, "startGracePeriod", caseId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_MirrorTransistions *MirrorTransistionsSession) StartGracePeriod(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.StartGracePeriod(&_MirrorTransistions.TransactOpts, caseId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(caseId bytes32, status uint8) returns(bool)
func (_MirrorTransistions *MirrorTransistionsTransactorSession) StartGracePeriod(caseId [32]byte, status uint8) (*types.Transaction, error) {
	return _MirrorTransistions.Contract.StartGracePeriod(&_MirrorTransistions.TransactOpts, caseId, status)
}

// TrialTransistionsAbstractABI is the input ABI used to generate the binding from.
const TrialTransistionsAbstractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getInitialStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"startGracePeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"witnessState\",\"type\":\"uint8\"},{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"trialStatus\",\"type\":\"uint8\"}],\"name\":\"getWitness\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"claimId\",\"type\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"expired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"}]"

// TrialTransistionsAbstractBin is the compiled bytecode used for deploying new contracts.
const TrialTransistionsAbstractBin = `0x`

// DeployTrialTransistionsAbstract deploys a new Ethereum contract, binding an instance of TrialTransistionsAbstract to it.
func DeployTrialTransistionsAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrialTransistionsAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialTransistionsAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TrialTransistionsAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrialTransistionsAbstract{TrialTransistionsAbstractCaller: TrialTransistionsAbstractCaller{contract: contract}, TrialTransistionsAbstractTransactor: TrialTransistionsAbstractTransactor{contract: contract}}, nil
}

// TrialTransistionsAbstract is an auto generated Go binding around an Ethereum contract.
type TrialTransistionsAbstract struct {
	TrialTransistionsAbstractCaller     // Read-only binding to the contract
	TrialTransistionsAbstractTransactor // Write-only binding to the contract
}

// TrialTransistionsAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialTransistionsAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialTransistionsAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrialTransistionsAbstractSession struct {
	Contract     *TrialTransistionsAbstract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TrialTransistionsAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrialTransistionsAbstractCallerSession struct {
	Contract *TrialTransistionsAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TrialTransistionsAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrialTransistionsAbstractTransactorSession struct {
	Contract     *TrialTransistionsAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TrialTransistionsAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrialTransistionsAbstractRaw struct {
	Contract *TrialTransistionsAbstract // Generic contract binding to access the raw methods on
}

// TrialTransistionsAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractCallerRaw struct {
	Contract *TrialTransistionsAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// TrialTransistionsAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrialTransistionsAbstractTransactorRaw struct {
	Contract *TrialTransistionsAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrialTransistionsAbstract creates a new instance of TrialTransistionsAbstract, bound to a specific deployed contract.
func NewTrialTransistionsAbstract(address common.Address, backend bind.ContractBackend) (*TrialTransistionsAbstract, error) {
	contract, err := bindTrialTransistionsAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrialTransistionsAbstract{TrialTransistionsAbstractCaller: TrialTransistionsAbstractCaller{contract: contract}, TrialTransistionsAbstractTransactor: TrialTransistionsAbstractTransactor{contract: contract}}, nil
}

// NewTrialTransistionsAbstractCaller creates a new read-only instance of TrialTransistionsAbstract, bound to a specific deployed contract.
func NewTrialTransistionsAbstractCaller(address common.Address, caller bind.ContractCaller) (*TrialTransistionsAbstractCaller, error) {
	contract, err := bindTrialTransistionsAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TrialTransistionsAbstractCaller{contract: contract}, nil
}

// NewTrialTransistionsAbstractTransactor creates a new write-only instance of TrialTransistionsAbstract, bound to a specific deployed contract.
func NewTrialTransistionsAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*TrialTransistionsAbstractTransactor, error) {
	contract, err := bindTrialTransistionsAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TrialTransistionsAbstractTransactor{contract: contract}, nil
}

// bindTrialTransistionsAbstract binds a generic wrapper to an already deployed contract.
func bindTrialTransistionsAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialTransistionsAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrialTransistionsAbstract.Contract.TrialTransistionsAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.TrialTransistionsAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.TrialTransistionsAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrialTransistionsAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.contract.Transact(opts, method, params...)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(claimId bytes32, status uint8) returns(bool)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) Expired(opts *bind.TransactOpts, claimId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "expired", claimId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(claimId bytes32, status uint8) returns(bool)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) Expired(claimId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.Expired(&_TrialTransistionsAbstract.TransactOpts, claimId, status)
}

// Expired is a paid mutator transaction binding the contract method 0xfcc0a2ae.
//
// Solidity: function expired(claimId bytes32, status uint8) returns(bool)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) Expired(claimId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.Expired(&_TrialTransistionsAbstract.TransactOpts, claimId, status)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetInitialStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getInitialStatus")
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetInitialStatus() (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetInitialStatus(&_TrialTransistionsAbstract.TransactOpts)
}

// GetInitialStatus is a paid mutator transaction binding the contract method 0x07e5ae5f.
//
// Solidity: function getInitialStatus() returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetInitialStatus() (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetInitialStatus(&_TrialTransistionsAbstract.TransactOpts)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetStatus(opts *bind.TransactOpts, witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getStatus", witnessState, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetStatus(witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetStatus(&_TrialTransistionsAbstract.TransactOpts, witnessState, trialStatus)
}

// GetStatus is a paid mutator transaction binding the contract method 0xd39eb301.
//
// Solidity: function getStatus(witnessState uint8, trialStatus uint8) returns(status uint8)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetStatus(witnessState uint8, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetStatus(&_TrialTransistionsAbstract.TransactOpts, witnessState, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) GetWitness(opts *bind.TransactOpts, trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "getWitness", trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetWitness(&_TrialTransistionsAbstract.TransactOpts, trialStatus)
}

// GetWitness is a paid mutator transaction binding the contract method 0xdb061563.
//
// Solidity: function getWitness(trialStatus uint8) returns(address)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) GetWitness(trialStatus uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.GetWitness(&_TrialTransistionsAbstract.TransactOpts, trialStatus)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(claimId bytes32, status uint8) returns(bool)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactor) StartGracePeriod(opts *bind.TransactOpts, claimId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.contract.Transact(opts, "startGracePeriod", claimId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(claimId bytes32, status uint8) returns(bool)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractSession) StartGracePeriod(claimId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.StartGracePeriod(&_TrialTransistionsAbstract.TransactOpts, claimId, status)
}

// StartGracePeriod is a paid mutator transaction binding the contract method 0xa7edcbe4.
//
// Solidity: function startGracePeriod(claimId bytes32, status uint8) returns(bool)
func (_TrialTransistionsAbstract *TrialTransistionsAbstractTransactorSession) StartGracePeriod(claimId [32]byte, status uint8) (*types.Transaction, error) {
	return _TrialTransistionsAbstract.Contract.StartGracePeriod(&_TrialTransistionsAbstract.TransactOpts, claimId, status)
}

// WitnessAbstractABI is the input ABI used to generate the binding from.
const WitnessAbstractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"isEvidentSubmited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"caseId\",\"type\":\"bytes32\"},{\"name\":\"serviceId\",\"type\":\"bytes32\"},{\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"testimonyFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"}]"

// WitnessAbstractBin is the compiled bytecode used for deploying new contracts.
const WitnessAbstractBin = `0x`

// DeployWitnessAbstract deploys a new Ethereum contract, binding an instance of WitnessAbstract to it.
func DeployWitnessAbstract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WitnessAbstract, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessAbstractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WitnessAbstractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WitnessAbstract{WitnessAbstractCaller: WitnessAbstractCaller{contract: contract}, WitnessAbstractTransactor: WitnessAbstractTransactor{contract: contract}}, nil
}

// WitnessAbstract is an auto generated Go binding around an Ethereum contract.
type WitnessAbstract struct {
	WitnessAbstractCaller     // Read-only binding to the contract
	WitnessAbstractTransactor // Write-only binding to the contract
}

// WitnessAbstractCaller is an auto generated read-only Go binding around an Ethereum contract.
type WitnessAbstractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessAbstractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WitnessAbstractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessAbstractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WitnessAbstractSession struct {
	Contract     *WitnessAbstract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WitnessAbstractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WitnessAbstractCallerSession struct {
	Contract *WitnessAbstractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WitnessAbstractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WitnessAbstractTransactorSession struct {
	Contract     *WitnessAbstractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WitnessAbstractRaw is an auto generated low-level Go binding around an Ethereum contract.
type WitnessAbstractRaw struct {
	Contract *WitnessAbstract // Generic contract binding to access the raw methods on
}

// WitnessAbstractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WitnessAbstractCallerRaw struct {
	Contract *WitnessAbstractCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessAbstractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WitnessAbstractTransactorRaw struct {
	Contract *WitnessAbstractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessAbstract creates a new instance of WitnessAbstract, bound to a specific deployed contract.
func NewWitnessAbstract(address common.Address, backend bind.ContractBackend) (*WitnessAbstract, error) {
	contract, err := bindWitnessAbstract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessAbstract{WitnessAbstractCaller: WitnessAbstractCaller{contract: contract}, WitnessAbstractTransactor: WitnessAbstractTransactor{contract: contract}}, nil
}

// NewWitnessAbstractCaller creates a new read-only instance of WitnessAbstract, bound to a specific deployed contract.
func NewWitnessAbstractCaller(address common.Address, caller bind.ContractCaller) (*WitnessAbstractCaller, error) {
	contract, err := bindWitnessAbstract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessAbstractCaller{contract: contract}, nil
}

// NewWitnessAbstractTransactor creates a new write-only instance of WitnessAbstract, bound to a specific deployed contract.
func NewWitnessAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessAbstractTransactor, error) {
	contract, err := bindWitnessAbstract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WitnessAbstractTransactor{contract: contract}, nil
}

// bindWitnessAbstract binds a generic wrapper to an already deployed contract.
func bindWitnessAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessAbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessAbstract *WitnessAbstractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessAbstract.Contract.WitnessAbstractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessAbstract *WitnessAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.WitnessAbstractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessAbstract *WitnessAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.WitnessAbstractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessAbstract *WitnessAbstractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessAbstract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessAbstract *WitnessAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessAbstract *WitnessAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.contract.Transact(opts, method, params...)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractTransactor) IsEvidentSubmited(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.contract.Transact(opts, "isEvidentSubmited", caseId, serviceId, clientAddress)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractSession) IsEvidentSubmited(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.IsEvidentSubmited(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// IsEvidentSubmited is a paid mutator transaction binding the contract method 0x70e39acd.
//
// Solidity: function isEvidentSubmited(caseId bytes32, serviceId bytes32, clientAddress address) returns(bool)
func (_WitnessAbstract *WitnessAbstractTransactorSession) IsEvidentSubmited(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.IsEvidentSubmited(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractTransactor) TestimonyFor(opts *bind.TransactOpts, caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.contract.Transact(opts, "testimonyFor", caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.TestimonyFor(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}

// TestimonyFor is a paid mutator transaction binding the contract method 0xc4b45283.
//
// Solidity: function testimonyFor(caseId bytes32, serviceId bytes32, clientAddress address) returns(uint8)
func (_WitnessAbstract *WitnessAbstractTransactorSession) TestimonyFor(caseId [32]byte, serviceId [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _WitnessAbstract.Contract.TestimonyFor(&_WitnessAbstract.TransactOpts, caseId, serviceId, clientAddress)
}
