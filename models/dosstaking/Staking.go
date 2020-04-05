// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dosstaking

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

// DosstakingABI is the input ABI used to generate the binding from.
const DosstakingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nodeRunner\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"selfStakedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedDB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardCut\",\"type\":\"uint256\"}],\"name\":\"NewNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nodeRunner\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropburnAmount\",\"type\":\"uint256\"}],\"name\":\"Unbond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCirculatingSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCirculatingSupply\",\"type\":\"uint256\"}],\"name\":\"UpdateCirculatingSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"UpdateDropBurnMaxQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMinStakePerNode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinStakePerNode\",\"type\":\"uint256\"}],\"name\":\"UpdateMinStakePerNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"UpdateUnbondDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nodeRunner\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DBDECIMAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DBTOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOSDECIMAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOSTOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ONEYEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accumulatedRewardIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"delegatorClaimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"delegatorUnbond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"delegatorWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"delegatorWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"delegatedNode\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRewardIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingWithdraw\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dropburnMaxQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorRewardTokensRT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"getNodeRewardTokensRT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"getNodeUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initBlkN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dostoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dbtoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"isValidStakingNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRateUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakePerNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dropburnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardCut\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_desc\",\"type\":\"string\"}],\"name\":\"newNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeClaimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeRunners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeStart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeStop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dropburnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeUnbond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeUnregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardCut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedDB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"selfStakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalOtherDelegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRewardIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingWithdrawToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingWithdrawDB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastStartTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"running\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newSupply\",\"type\":\"uint256\"}],\"name\":\"setCirculatingSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quota\",\"type\":\"uint256\"}],\"name\":\"setDropBurnMaxQuota\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStake\",\"type\":\"uint256\"}],\"name\":\"setMinStakePerNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setUnbondDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingRewardsVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newDropburnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newCut\",\"type\":\"uint256\"}],\"name\":\"updateNodeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Dosstaking is an auto generated Go binding around an Ethereum contract.
type Dosstaking struct {
	DosstakingCaller     // Read-only binding to the contract
	DosstakingTransactor // Write-only binding to the contract
	DosstakingFilterer   // Log filterer for contract events
}

// DosstakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type DosstakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosstakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DosstakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosstakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DosstakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosstakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DosstakingSession struct {
	Contract     *Dosstaking       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DosstakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DosstakingCallerSession struct {
	Contract *DosstakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DosstakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DosstakingTransactorSession struct {
	Contract     *DosstakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DosstakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type DosstakingRaw struct {
	Contract *Dosstaking // Generic contract binding to access the raw methods on
}

// DosstakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DosstakingCallerRaw struct {
	Contract *DosstakingCaller // Generic read-only contract binding to access the raw methods on
}

// DosstakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DosstakingTransactorRaw struct {
	Contract *DosstakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDosstaking creates a new instance of Dosstaking, bound to a specific deployed contract.
func NewDosstaking(address common.Address, backend bind.ContractBackend) (*Dosstaking, error) {
	contract, err := bindDosstaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dosstaking{DosstakingCaller: DosstakingCaller{contract: contract}, DosstakingTransactor: DosstakingTransactor{contract: contract}, DosstakingFilterer: DosstakingFilterer{contract: contract}}, nil
}

// NewDosstakingCaller creates a new read-only instance of Dosstaking, bound to a specific deployed contract.
func NewDosstakingCaller(address common.Address, caller bind.ContractCaller) (*DosstakingCaller, error) {
	contract, err := bindDosstaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DosstakingCaller{contract: contract}, nil
}

// NewDosstakingTransactor creates a new write-only instance of Dosstaking, bound to a specific deployed contract.
func NewDosstakingTransactor(address common.Address, transactor bind.ContractTransactor) (*DosstakingTransactor, error) {
	contract, err := bindDosstaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DosstakingTransactor{contract: contract}, nil
}

// NewDosstakingFilterer creates a new log filterer instance of Dosstaking, bound to a specific deployed contract.
func NewDosstakingFilterer(address common.Address, filterer bind.ContractFilterer) (*DosstakingFilterer, error) {
	contract, err := bindDosstaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DosstakingFilterer{contract: contract}, nil
}

// bindDosstaking binds a generic wrapper to an already deployed contract.
func bindDosstaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DosstakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dosstaking *DosstakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dosstaking.Contract.DosstakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dosstaking *DosstakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosstaking.Contract.DosstakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dosstaking *DosstakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dosstaking.Contract.DosstakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dosstaking *DosstakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dosstaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dosstaking *DosstakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosstaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dosstaking *DosstakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dosstaking.Contract.contract.Transact(opts, method, params...)
}

// DBDECIMAL is a free data retrieval call binding the contract method 0x91bf6960.
//
// Solidity: function DBDECIMAL() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) DBDECIMAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "DBDECIMAL")
	return *ret0, err
}

// DBDECIMAL is a free data retrieval call binding the contract method 0x91bf6960.
//
// Solidity: function DBDECIMAL() constant returns(uint256)
func (_Dosstaking *DosstakingSession) DBDECIMAL() (*big.Int, error) {
	return _Dosstaking.Contract.DBDECIMAL(&_Dosstaking.CallOpts)
}

// DBDECIMAL is a free data retrieval call binding the contract method 0x91bf6960.
//
// Solidity: function DBDECIMAL() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) DBDECIMAL() (*big.Int, error) {
	return _Dosstaking.Contract.DBDECIMAL(&_Dosstaking.CallOpts)
}

// DBTOKEN is a free data retrieval call binding the contract method 0xd2a02541.
//
// Solidity: function DBTOKEN() constant returns(address)
func (_Dosstaking *DosstakingCaller) DBTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "DBTOKEN")
	return *ret0, err
}

// DBTOKEN is a free data retrieval call binding the contract method 0xd2a02541.
//
// Solidity: function DBTOKEN() constant returns(address)
func (_Dosstaking *DosstakingSession) DBTOKEN() (common.Address, error) {
	return _Dosstaking.Contract.DBTOKEN(&_Dosstaking.CallOpts)
}

// DBTOKEN is a free data retrieval call binding the contract method 0xd2a02541.
//
// Solidity: function DBTOKEN() constant returns(address)
func (_Dosstaking *DosstakingCallerSession) DBTOKEN() (common.Address, error) {
	return _Dosstaking.Contract.DBTOKEN(&_Dosstaking.CallOpts)
}

// DOSDECIMAL is a free data retrieval call binding the contract method 0xc48f4d47.
//
// Solidity: function DOSDECIMAL() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) DOSDECIMAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "DOSDECIMAL")
	return *ret0, err
}

// DOSDECIMAL is a free data retrieval call binding the contract method 0xc48f4d47.
//
// Solidity: function DOSDECIMAL() constant returns(uint256)
func (_Dosstaking *DosstakingSession) DOSDECIMAL() (*big.Int, error) {
	return _Dosstaking.Contract.DOSDECIMAL(&_Dosstaking.CallOpts)
}

// DOSDECIMAL is a free data retrieval call binding the contract method 0xc48f4d47.
//
// Solidity: function DOSDECIMAL() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) DOSDECIMAL() (*big.Int, error) {
	return _Dosstaking.Contract.DOSDECIMAL(&_Dosstaking.CallOpts)
}

// DOSTOKEN is a free data retrieval call binding the contract method 0x0fd1d0ea.
//
// Solidity: function DOSTOKEN() constant returns(address)
func (_Dosstaking *DosstakingCaller) DOSTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "DOSTOKEN")
	return *ret0, err
}

// DOSTOKEN is a free data retrieval call binding the contract method 0x0fd1d0ea.
//
// Solidity: function DOSTOKEN() constant returns(address)
func (_Dosstaking *DosstakingSession) DOSTOKEN() (common.Address, error) {
	return _Dosstaking.Contract.DOSTOKEN(&_Dosstaking.CallOpts)
}

// DOSTOKEN is a free data retrieval call binding the contract method 0x0fd1d0ea.
//
// Solidity: function DOSTOKEN() constant returns(address)
func (_Dosstaking *DosstakingCallerSession) DOSTOKEN() (common.Address, error) {
	return _Dosstaking.Contract.DOSTOKEN(&_Dosstaking.CallOpts)
}

// ONEYEAR is a free data retrieval call binding the contract method 0x195cb3ab.
//
// Solidity: function ONEYEAR() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) ONEYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "ONEYEAR")
	return *ret0, err
}

// ONEYEAR is a free data retrieval call binding the contract method 0x195cb3ab.
//
// Solidity: function ONEYEAR() constant returns(uint256)
func (_Dosstaking *DosstakingSession) ONEYEAR() (*big.Int, error) {
	return _Dosstaking.Contract.ONEYEAR(&_Dosstaking.CallOpts)
}

// ONEYEAR is a free data retrieval call binding the contract method 0x195cb3ab.
//
// Solidity: function ONEYEAR() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) ONEYEAR() (*big.Int, error) {
	return _Dosstaking.Contract.ONEYEAR(&_Dosstaking.CallOpts)
}

// AccumulatedRewardIndex is a free data retrieval call binding the contract method 0x37308854.
//
// Solidity: function accumulatedRewardIndex() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) AccumulatedRewardIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "accumulatedRewardIndex")
	return *ret0, err
}

// AccumulatedRewardIndex is a free data retrieval call binding the contract method 0x37308854.
//
// Solidity: function accumulatedRewardIndex() constant returns(uint256)
func (_Dosstaking *DosstakingSession) AccumulatedRewardIndex() (*big.Int, error) {
	return _Dosstaking.Contract.AccumulatedRewardIndex(&_Dosstaking.CallOpts)
}

// AccumulatedRewardIndex is a free data retrieval call binding the contract method 0x37308854.
//
// Solidity: function accumulatedRewardIndex() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) AccumulatedRewardIndex() (*big.Int, error) {
	return _Dosstaking.Contract.AccumulatedRewardIndex(&_Dosstaking.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Dosstaking *DosstakingCaller) BridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "bridgeAddr")
	return *ret0, err
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Dosstaking *DosstakingSession) BridgeAddr() (common.Address, error) {
	return _Dosstaking.Contract.BridgeAddr(&_Dosstaking.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Dosstaking *DosstakingCallerSession) BridgeAddr() (common.Address, error) {
	return _Dosstaking.Contract.BridgeAddr(&_Dosstaking.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) CirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "circulatingSupply")
	return *ret0, err
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_Dosstaking *DosstakingSession) CirculatingSupply() (*big.Int, error) {
	return _Dosstaking.Contract.CirculatingSupply(&_Dosstaking.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) CirculatingSupply() (*big.Int, error) {
	return _Dosstaking.Contract.CirculatingSupply(&_Dosstaking.CallOpts)
}

// DelegatorWithdrawable is a free data retrieval call binding the contract method 0x0bac06cb.
//
// Solidity: function delegatorWithdrawable(_owner address, _nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCaller) DelegatorWithdrawable(opts *bind.CallOpts, _owner common.Address, _nodeAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "delegatorWithdrawable", _owner, _nodeAddr)
	return *ret0, err
}

// DelegatorWithdrawable is a free data retrieval call binding the contract method 0x0bac06cb.
//
// Solidity: function delegatorWithdrawable(_owner address, _nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingSession) DelegatorWithdrawable(_owner common.Address, _nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.DelegatorWithdrawable(&_Dosstaking.CallOpts, _owner, _nodeAddr)
}

// DelegatorWithdrawable is a free data retrieval call binding the contract method 0x0bac06cb.
//
// Solidity: function delegatorWithdrawable(_owner address, _nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) DelegatorWithdrawable(_owner common.Address, _nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.DelegatorWithdrawable(&_Dosstaking.CallOpts, _owner, _nodeAddr)
}

// Delegators is a free data retrieval call binding the contract method 0xa2526bd3.
//
// Solidity: function delegators( address,  address) constant returns(delegatedNode address, delegatedAmount uint256, accumulatedRewards uint256, accumulatedRewardIndex uint256, pendingWithdraw uint256)
func (_Dosstaking *DosstakingCaller) Delegators(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	DelegatedNode          common.Address
	DelegatedAmount        *big.Int
	AccumulatedRewards     *big.Int
	AccumulatedRewardIndex *big.Int
	PendingWithdraw        *big.Int
}, error) {
	ret := new(struct {
		DelegatedNode          common.Address
		DelegatedAmount        *big.Int
		AccumulatedRewards     *big.Int
		AccumulatedRewardIndex *big.Int
		PendingWithdraw        *big.Int
	})
	out := ret
	err := _Dosstaking.contract.Call(opts, out, "delegators", arg0, arg1)
	return *ret, err
}

// Delegators is a free data retrieval call binding the contract method 0xa2526bd3.
//
// Solidity: function delegators( address,  address) constant returns(delegatedNode address, delegatedAmount uint256, accumulatedRewards uint256, accumulatedRewardIndex uint256, pendingWithdraw uint256)
func (_Dosstaking *DosstakingSession) Delegators(arg0 common.Address, arg1 common.Address) (struct {
	DelegatedNode          common.Address
	DelegatedAmount        *big.Int
	AccumulatedRewards     *big.Int
	AccumulatedRewardIndex *big.Int
	PendingWithdraw        *big.Int
}, error) {
	return _Dosstaking.Contract.Delegators(&_Dosstaking.CallOpts, arg0, arg1)
}

// Delegators is a free data retrieval call binding the contract method 0xa2526bd3.
//
// Solidity: function delegators( address,  address) constant returns(delegatedNode address, delegatedAmount uint256, accumulatedRewards uint256, accumulatedRewardIndex uint256, pendingWithdraw uint256)
func (_Dosstaking *DosstakingCallerSession) Delegators(arg0 common.Address, arg1 common.Address) (struct {
	DelegatedNode          common.Address
	DelegatedAmount        *big.Int
	AccumulatedRewards     *big.Int
	AccumulatedRewardIndex *big.Int
	PendingWithdraw        *big.Int
}, error) {
	return _Dosstaking.Contract.Delegators(&_Dosstaking.CallOpts, arg0, arg1)
}

// DropburnMaxQuota is a free data retrieval call binding the contract method 0xe8c3470c.
//
// Solidity: function dropburnMaxQuota() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) DropburnMaxQuota(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "dropburnMaxQuota")
	return *ret0, err
}

// DropburnMaxQuota is a free data retrieval call binding the contract method 0xe8c3470c.
//
// Solidity: function dropburnMaxQuota() constant returns(uint256)
func (_Dosstaking *DosstakingSession) DropburnMaxQuota() (*big.Int, error) {
	return _Dosstaking.Contract.DropburnMaxQuota(&_Dosstaking.CallOpts)
}

// DropburnMaxQuota is a free data retrieval call binding the contract method 0xe8c3470c.
//
// Solidity: function dropburnMaxQuota() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) DropburnMaxQuota() (*big.Int, error) {
	return _Dosstaking.Contract.DropburnMaxQuota(&_Dosstaking.CallOpts)
}

// GetCurrentAPR is a free data retrieval call binding the contract method 0x21ae05c8.
//
// Solidity: function getCurrentAPR() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) GetCurrentAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "getCurrentAPR")
	return *ret0, err
}

// GetCurrentAPR is a free data retrieval call binding the contract method 0x21ae05c8.
//
// Solidity: function getCurrentAPR() constant returns(uint256)
func (_Dosstaking *DosstakingSession) GetCurrentAPR() (*big.Int, error) {
	return _Dosstaking.Contract.GetCurrentAPR(&_Dosstaking.CallOpts)
}

// GetCurrentAPR is a free data retrieval call binding the contract method 0x21ae05c8.
//
// Solidity: function getCurrentAPR() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) GetCurrentAPR() (*big.Int, error) {
	return _Dosstaking.Contract.GetCurrentAPR(&_Dosstaking.CallOpts)
}

// GetDelegatorRewardTokensRT is a free data retrieval call binding the contract method 0xfb7f2488.
//
// Solidity: function getDelegatorRewardTokensRT(_delegator address, _nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCaller) GetDelegatorRewardTokensRT(opts *bind.CallOpts, _delegator common.Address, _nodeAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "getDelegatorRewardTokensRT", _delegator, _nodeAddr)
	return *ret0, err
}

// GetDelegatorRewardTokensRT is a free data retrieval call binding the contract method 0xfb7f2488.
//
// Solidity: function getDelegatorRewardTokensRT(_delegator address, _nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingSession) GetDelegatorRewardTokensRT(_delegator common.Address, _nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.GetDelegatorRewardTokensRT(&_Dosstaking.CallOpts, _delegator, _nodeAddr)
}

// GetDelegatorRewardTokensRT is a free data retrieval call binding the contract method 0xfb7f2488.
//
// Solidity: function getDelegatorRewardTokensRT(_delegator address, _nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) GetDelegatorRewardTokensRT(_delegator common.Address, _nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.GetDelegatorRewardTokensRT(&_Dosstaking.CallOpts, _delegator, _nodeAddr)
}

// GetNodeAddrs is a free data retrieval call binding the contract method 0x1017bf56.
//
// Solidity: function getNodeAddrs() constant returns(address[])
func (_Dosstaking *DosstakingCaller) GetNodeAddrs(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "getNodeAddrs")
	return *ret0, err
}

// GetNodeAddrs is a free data retrieval call binding the contract method 0x1017bf56.
//
// Solidity: function getNodeAddrs() constant returns(address[])
func (_Dosstaking *DosstakingSession) GetNodeAddrs() ([]common.Address, error) {
	return _Dosstaking.Contract.GetNodeAddrs(&_Dosstaking.CallOpts)
}

// GetNodeAddrs is a free data retrieval call binding the contract method 0x1017bf56.
//
// Solidity: function getNodeAddrs() constant returns(address[])
func (_Dosstaking *DosstakingCallerSession) GetNodeAddrs() ([]common.Address, error) {
	return _Dosstaking.Contract.GetNodeAddrs(&_Dosstaking.CallOpts)
}

// GetNodeRewardTokensRT is a free data retrieval call binding the contract method 0x4d0b5ea9.
//
// Solidity: function getNodeRewardTokensRT(_nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCaller) GetNodeRewardTokensRT(opts *bind.CallOpts, _nodeAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "getNodeRewardTokensRT", _nodeAddr)
	return *ret0, err
}

// GetNodeRewardTokensRT is a free data retrieval call binding the contract method 0x4d0b5ea9.
//
// Solidity: function getNodeRewardTokensRT(_nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingSession) GetNodeRewardTokensRT(_nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.GetNodeRewardTokensRT(&_Dosstaking.CallOpts, _nodeAddr)
}

// GetNodeRewardTokensRT is a free data retrieval call binding the contract method 0x4d0b5ea9.
//
// Solidity: function getNodeRewardTokensRT(_nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) GetNodeRewardTokensRT(_nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.GetNodeRewardTokensRT(&_Dosstaking.CallOpts, _nodeAddr)
}

// GetNodeUptime is a free data retrieval call binding the contract method 0x7f92ca19.
//
// Solidity: function getNodeUptime(nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCaller) GetNodeUptime(opts *bind.CallOpts, nodeAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "getNodeUptime", nodeAddr)
	return *ret0, err
}

// GetNodeUptime is a free data retrieval call binding the contract method 0x7f92ca19.
//
// Solidity: function getNodeUptime(nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingSession) GetNodeUptime(nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.GetNodeUptime(&_Dosstaking.CallOpts, nodeAddr)
}

// GetNodeUptime is a free data retrieval call binding the contract method 0x7f92ca19.
//
// Solidity: function getNodeUptime(nodeAddr address) constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) GetNodeUptime(nodeAddr common.Address) (*big.Int, error) {
	return _Dosstaking.Contract.GetNodeUptime(&_Dosstaking.CallOpts, nodeAddr)
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) InitBlkN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "initBlkN")
	return *ret0, err
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() constant returns(uint256)
func (_Dosstaking *DosstakingSession) InitBlkN() (*big.Int, error) {
	return _Dosstaking.Contract.InitBlkN(&_Dosstaking.CallOpts)
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) InitBlkN() (*big.Int, error) {
	return _Dosstaking.Contract.InitBlkN(&_Dosstaking.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosstaking *DosstakingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosstaking *DosstakingSession) IsOwner() (bool, error) {
	return _Dosstaking.Contract.IsOwner(&_Dosstaking.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosstaking *DosstakingCallerSession) IsOwner() (bool, error) {
	return _Dosstaking.Contract.IsOwner(&_Dosstaking.CallOpts)
}

// IsValidStakingNode is a free data retrieval call binding the contract method 0xa8e8ab38.
//
// Solidity: function isValidStakingNode(nodeAddr address) constant returns(bool)
func (_Dosstaking *DosstakingCaller) IsValidStakingNode(opts *bind.CallOpts, nodeAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "isValidStakingNode", nodeAddr)
	return *ret0, err
}

// IsValidStakingNode is a free data retrieval call binding the contract method 0xa8e8ab38.
//
// Solidity: function isValidStakingNode(nodeAddr address) constant returns(bool)
func (_Dosstaking *DosstakingSession) IsValidStakingNode(nodeAddr common.Address) (bool, error) {
	return _Dosstaking.Contract.IsValidStakingNode(&_Dosstaking.CallOpts, nodeAddr)
}

// IsValidStakingNode is a free data retrieval call binding the contract method 0xa8e8ab38.
//
// Solidity: function isValidStakingNode(nodeAddr address) constant returns(bool)
func (_Dosstaking *DosstakingCallerSession) IsValidStakingNode(nodeAddr common.Address) (bool, error) {
	return _Dosstaking.Contract.IsValidStakingNode(&_Dosstaking.CallOpts, nodeAddr)
}

// LastRateUpdatedTime is a free data retrieval call binding the contract method 0xbe568059.
//
// Solidity: function lastRateUpdatedTime() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) LastRateUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "lastRateUpdatedTime")
	return *ret0, err
}

// LastRateUpdatedTime is a free data retrieval call binding the contract method 0xbe568059.
//
// Solidity: function lastRateUpdatedTime() constant returns(uint256)
func (_Dosstaking *DosstakingSession) LastRateUpdatedTime() (*big.Int, error) {
	return _Dosstaking.Contract.LastRateUpdatedTime(&_Dosstaking.CallOpts)
}

// LastRateUpdatedTime is a free data retrieval call binding the contract method 0xbe568059.
//
// Solidity: function lastRateUpdatedTime() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) LastRateUpdatedTime() (*big.Int, error) {
	return _Dosstaking.Contract.LastRateUpdatedTime(&_Dosstaking.CallOpts)
}

// MinStakePerNode is a free data retrieval call binding the contract method 0xe5f95a99.
//
// Solidity: function minStakePerNode() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) MinStakePerNode(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "minStakePerNode")
	return *ret0, err
}

// MinStakePerNode is a free data retrieval call binding the contract method 0xe5f95a99.
//
// Solidity: function minStakePerNode() constant returns(uint256)
func (_Dosstaking *DosstakingSession) MinStakePerNode() (*big.Int, error) {
	return _Dosstaking.Contract.MinStakePerNode(&_Dosstaking.CallOpts)
}

// MinStakePerNode is a free data retrieval call binding the contract method 0xe5f95a99.
//
// Solidity: function minStakePerNode() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) MinStakePerNode() (*big.Int, error) {
	return _Dosstaking.Contract.MinStakePerNode(&_Dosstaking.CallOpts)
}

// NodeAddrs is a free data retrieval call binding the contract method 0xb4a26490.
//
// Solidity: function nodeAddrs( uint256) constant returns(address)
func (_Dosstaking *DosstakingCaller) NodeAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "nodeAddrs", arg0)
	return *ret0, err
}

// NodeAddrs is a free data retrieval call binding the contract method 0xb4a26490.
//
// Solidity: function nodeAddrs( uint256) constant returns(address)
func (_Dosstaking *DosstakingSession) NodeAddrs(arg0 *big.Int) (common.Address, error) {
	return _Dosstaking.Contract.NodeAddrs(&_Dosstaking.CallOpts, arg0)
}

// NodeAddrs is a free data retrieval call binding the contract method 0xb4a26490.
//
// Solidity: function nodeAddrs( uint256) constant returns(address)
func (_Dosstaking *DosstakingCallerSession) NodeAddrs(arg0 *big.Int) (common.Address, error) {
	return _Dosstaking.Contract.NodeAddrs(&_Dosstaking.CallOpts, arg0)
}

// NodeRunners is a free data retrieval call binding the contract method 0xbe4a455f.
//
// Solidity: function nodeRunners( address,  address) constant returns(bool)
func (_Dosstaking *DosstakingCaller) NodeRunners(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "nodeRunners", arg0, arg1)
	return *ret0, err
}

// NodeRunners is a free data retrieval call binding the contract method 0xbe4a455f.
//
// Solidity: function nodeRunners( address,  address) constant returns(bool)
func (_Dosstaking *DosstakingSession) NodeRunners(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Dosstaking.Contract.NodeRunners(&_Dosstaking.CallOpts, arg0, arg1)
}

// NodeRunners is a free data retrieval call binding the contract method 0xbe4a455f.
//
// Solidity: function nodeRunners( address,  address) constant returns(bool)
func (_Dosstaking *DosstakingCallerSession) NodeRunners(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Dosstaking.Contract.NodeRunners(&_Dosstaking.CallOpts, arg0, arg1)
}

// NodeWithdrawable is a free data retrieval call binding the contract method 0x70841a0b.
//
// Solidity: function nodeWithdrawable(_owner address, _nodeAddr address) constant returns(uint256, uint256)
func (_Dosstaking *DosstakingCaller) NodeWithdrawable(opts *bind.CallOpts, _owner common.Address, _nodeAddr common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dosstaking.contract.Call(opts, out, "nodeWithdrawable", _owner, _nodeAddr)
	return *ret0, *ret1, err
}

// NodeWithdrawable is a free data retrieval call binding the contract method 0x70841a0b.
//
// Solidity: function nodeWithdrawable(_owner address, _nodeAddr address) constant returns(uint256, uint256)
func (_Dosstaking *DosstakingSession) NodeWithdrawable(_owner common.Address, _nodeAddr common.Address) (*big.Int, *big.Int, error) {
	return _Dosstaking.Contract.NodeWithdrawable(&_Dosstaking.CallOpts, _owner, _nodeAddr)
}

// NodeWithdrawable is a free data retrieval call binding the contract method 0x70841a0b.
//
// Solidity: function nodeWithdrawable(_owner address, _nodeAddr address) constant returns(uint256, uint256)
func (_Dosstaking *DosstakingCallerSession) NodeWithdrawable(_owner common.Address, _nodeAddr common.Address) (*big.Int, *big.Int, error) {
	return _Dosstaking.Contract.NodeWithdrawable(&_Dosstaking.CallOpts, _owner, _nodeAddr)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes( address) constant returns(ownerAddr address, rewardCut uint256, stakedDB uint256, selfStakedAmount uint256, totalOtherDelegatedAmount uint256, accumulatedRewards uint256, accumulatedRewardIndex uint256, pendingWithdrawToken uint256, pendingWithdrawDB uint256, lastStartTime uint256, running bool, description string)
func (_Dosstaking *DosstakingCaller) Nodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	OwnerAddr                 common.Address
	RewardCut                 *big.Int
	StakedDB                  *big.Int
	SelfStakedAmount          *big.Int
	TotalOtherDelegatedAmount *big.Int
	AccumulatedRewards        *big.Int
	AccumulatedRewardIndex    *big.Int
	PendingWithdrawToken      *big.Int
	PendingWithdrawDB         *big.Int
	LastStartTime             *big.Int
	Running                   bool
	Description               string
}, error) {
	ret := new(struct {
		OwnerAddr                 common.Address
		RewardCut                 *big.Int
		StakedDB                  *big.Int
		SelfStakedAmount          *big.Int
		TotalOtherDelegatedAmount *big.Int
		AccumulatedRewards        *big.Int
		AccumulatedRewardIndex    *big.Int
		PendingWithdrawToken      *big.Int
		PendingWithdrawDB         *big.Int
		LastStartTime             *big.Int
		Running                   bool
		Description               string
	})
	out := ret
	err := _Dosstaking.contract.Call(opts, out, "nodes", arg0)
	return *ret, err
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes( address) constant returns(ownerAddr address, rewardCut uint256, stakedDB uint256, selfStakedAmount uint256, totalOtherDelegatedAmount uint256, accumulatedRewards uint256, accumulatedRewardIndex uint256, pendingWithdrawToken uint256, pendingWithdrawDB uint256, lastStartTime uint256, running bool, description string)
func (_Dosstaking *DosstakingSession) Nodes(arg0 common.Address) (struct {
	OwnerAddr                 common.Address
	RewardCut                 *big.Int
	StakedDB                  *big.Int
	SelfStakedAmount          *big.Int
	TotalOtherDelegatedAmount *big.Int
	AccumulatedRewards        *big.Int
	AccumulatedRewardIndex    *big.Int
	PendingWithdrawToken      *big.Int
	PendingWithdrawDB         *big.Int
	LastStartTime             *big.Int
	Running                   bool
	Description               string
}, error) {
	return _Dosstaking.Contract.Nodes(&_Dosstaking.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x189a5a17.
//
// Solidity: function nodes( address) constant returns(ownerAddr address, rewardCut uint256, stakedDB uint256, selfStakedAmount uint256, totalOtherDelegatedAmount uint256, accumulatedRewards uint256, accumulatedRewardIndex uint256, pendingWithdrawToken uint256, pendingWithdrawDB uint256, lastStartTime uint256, running bool, description string)
func (_Dosstaking *DosstakingCallerSession) Nodes(arg0 common.Address) (struct {
	OwnerAddr                 common.Address
	RewardCut                 *big.Int
	StakedDB                  *big.Int
	SelfStakedAmount          *big.Int
	TotalOtherDelegatedAmount *big.Int
	AccumulatedRewards        *big.Int
	AccumulatedRewardIndex    *big.Int
	PendingWithdrawToken      *big.Int
	PendingWithdrawDB         *big.Int
	LastStartTime             *big.Int
	Running                   bool
	Description               string
}, error) {
	return _Dosstaking.Contract.Nodes(&_Dosstaking.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosstaking *DosstakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosstaking *DosstakingSession) Owner() (common.Address, error) {
	return _Dosstaking.Contract.Owner(&_Dosstaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosstaking *DosstakingCallerSession) Owner() (common.Address, error) {
	return _Dosstaking.Contract.Owner(&_Dosstaking.CallOpts)
}

// StakingRewardsVault is a free data retrieval call binding the contract method 0xcd45dbcb.
//
// Solidity: function stakingRewardsVault() constant returns(address)
func (_Dosstaking *DosstakingCaller) StakingRewardsVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "stakingRewardsVault")
	return *ret0, err
}

// StakingRewardsVault is a free data retrieval call binding the contract method 0xcd45dbcb.
//
// Solidity: function stakingRewardsVault() constant returns(address)
func (_Dosstaking *DosstakingSession) StakingRewardsVault() (common.Address, error) {
	return _Dosstaking.Contract.StakingRewardsVault(&_Dosstaking.CallOpts)
}

// StakingRewardsVault is a free data retrieval call binding the contract method 0xcd45dbcb.
//
// Solidity: function stakingRewardsVault() constant returns(address)
func (_Dosstaking *DosstakingCallerSession) StakingRewardsVault() (common.Address, error) {
	return _Dosstaking.Contract.StakingRewardsVault(&_Dosstaking.CallOpts)
}

// TotalStakedTokens is a free data retrieval call binding the contract method 0x3ae73259.
//
// Solidity: function totalStakedTokens() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) TotalStakedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "totalStakedTokens")
	return *ret0, err
}

// TotalStakedTokens is a free data retrieval call binding the contract method 0x3ae73259.
//
// Solidity: function totalStakedTokens() constant returns(uint256)
func (_Dosstaking *DosstakingSession) TotalStakedTokens() (*big.Int, error) {
	return _Dosstaking.Contract.TotalStakedTokens(&_Dosstaking.CallOpts)
}

// TotalStakedTokens is a free data retrieval call binding the contract method 0x3ae73259.
//
// Solidity: function totalStakedTokens() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) TotalStakedTokens() (*big.Int, error) {
	return _Dosstaking.Contract.TotalStakedTokens(&_Dosstaking.CallOpts)
}

// UnbondDuration is a free data retrieval call binding the contract method 0xc4393444.
//
// Solidity: function unbondDuration() constant returns(uint256)
func (_Dosstaking *DosstakingCaller) UnbondDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosstaking.contract.Call(opts, out, "unbondDuration")
	return *ret0, err
}

// UnbondDuration is a free data retrieval call binding the contract method 0xc4393444.
//
// Solidity: function unbondDuration() constant returns(uint256)
func (_Dosstaking *DosstakingSession) UnbondDuration() (*big.Int, error) {
	return _Dosstaking.Contract.UnbondDuration(&_Dosstaking.CallOpts)
}

// UnbondDuration is a free data retrieval call binding the contract method 0xc4393444.
//
// Solidity: function unbondDuration() constant returns(uint256)
func (_Dosstaking *DosstakingCallerSession) UnbondDuration() (*big.Int, error) {
	return _Dosstaking.Contract.UnbondDuration(&_Dosstaking.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(_tokenAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) Delegate(opts *bind.TransactOpts, _tokenAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "delegate", _tokenAmount, _nodeAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(_tokenAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) Delegate(_tokenAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.Delegate(&_Dosstaking.TransactOpts, _tokenAmount, _nodeAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(_tokenAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) Delegate(_tokenAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.Delegate(&_Dosstaking.TransactOpts, _tokenAmount, _nodeAddr)
}

// DelegatorClaimReward is a paid mutator transaction binding the contract method 0x1cac57ec.
//
// Solidity: function delegatorClaimReward(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) DelegatorClaimReward(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "delegatorClaimReward", _nodeAddr)
}

// DelegatorClaimReward is a paid mutator transaction binding the contract method 0x1cac57ec.
//
// Solidity: function delegatorClaimReward(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) DelegatorClaimReward(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.DelegatorClaimReward(&_Dosstaking.TransactOpts, _nodeAddr)
}

// DelegatorClaimReward is a paid mutator transaction binding the contract method 0x1cac57ec.
//
// Solidity: function delegatorClaimReward(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) DelegatorClaimReward(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.DelegatorClaimReward(&_Dosstaking.TransactOpts, _nodeAddr)
}

// DelegatorUnbond is a paid mutator transaction binding the contract method 0x95d516dd.
//
// Solidity: function delegatorUnbond(_tokenAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) DelegatorUnbond(opts *bind.TransactOpts, _tokenAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "delegatorUnbond", _tokenAmount, _nodeAddr)
}

// DelegatorUnbond is a paid mutator transaction binding the contract method 0x95d516dd.
//
// Solidity: function delegatorUnbond(_tokenAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) DelegatorUnbond(_tokenAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.DelegatorUnbond(&_Dosstaking.TransactOpts, _tokenAmount, _nodeAddr)
}

// DelegatorUnbond is a paid mutator transaction binding the contract method 0x95d516dd.
//
// Solidity: function delegatorUnbond(_tokenAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) DelegatorUnbond(_tokenAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.DelegatorUnbond(&_Dosstaking.TransactOpts, _tokenAmount, _nodeAddr)
}

// DelegatorWithdraw is a paid mutator transaction binding the contract method 0xc314bb99.
//
// Solidity: function delegatorWithdraw(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) DelegatorWithdraw(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "delegatorWithdraw", _nodeAddr)
}

// DelegatorWithdraw is a paid mutator transaction binding the contract method 0xc314bb99.
//
// Solidity: function delegatorWithdraw(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) DelegatorWithdraw(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.DelegatorWithdraw(&_Dosstaking.TransactOpts, _nodeAddr)
}

// DelegatorWithdraw is a paid mutator transaction binding the contract method 0xc314bb99.
//
// Solidity: function delegatorWithdraw(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) DelegatorWithdraw(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.DelegatorWithdraw(&_Dosstaking.TransactOpts, _nodeAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_dostoken address, _dbtoken address, _vault address, _bridgeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) Initialize(opts *bind.TransactOpts, _dostoken common.Address, _dbtoken common.Address, _vault common.Address, _bridgeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "initialize", _dostoken, _dbtoken, _vault, _bridgeAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_dostoken address, _dbtoken address, _vault address, _bridgeAddr address) returns()
func (_Dosstaking *DosstakingSession) Initialize(_dostoken common.Address, _dbtoken common.Address, _vault common.Address, _bridgeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.Initialize(&_Dosstaking.TransactOpts, _dostoken, _dbtoken, _vault, _bridgeAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(_dostoken address, _dbtoken address, _vault address, _bridgeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) Initialize(_dostoken common.Address, _dbtoken common.Address, _vault common.Address, _bridgeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.Initialize(&_Dosstaking.TransactOpts, _dostoken, _dbtoken, _vault, _bridgeAddr)
}

// NewNode is a paid mutator transaction binding the contract method 0x867f121a.
//
// Solidity: function newNode(_nodeAddr address, _tokenAmount uint256, _dropburnAmount uint256, _rewardCut uint256, _desc string) returns()
func (_Dosstaking *DosstakingTransactor) NewNode(opts *bind.TransactOpts, _nodeAddr common.Address, _tokenAmount *big.Int, _dropburnAmount *big.Int, _rewardCut *big.Int, _desc string) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "newNode", _nodeAddr, _tokenAmount, _dropburnAmount, _rewardCut, _desc)
}

// NewNode is a paid mutator transaction binding the contract method 0x867f121a.
//
// Solidity: function newNode(_nodeAddr address, _tokenAmount uint256, _dropburnAmount uint256, _rewardCut uint256, _desc string) returns()
func (_Dosstaking *DosstakingSession) NewNode(_nodeAddr common.Address, _tokenAmount *big.Int, _dropburnAmount *big.Int, _rewardCut *big.Int, _desc string) (*types.Transaction, error) {
	return _Dosstaking.Contract.NewNode(&_Dosstaking.TransactOpts, _nodeAddr, _tokenAmount, _dropburnAmount, _rewardCut, _desc)
}

// NewNode is a paid mutator transaction binding the contract method 0x867f121a.
//
// Solidity: function newNode(_nodeAddr address, _tokenAmount uint256, _dropburnAmount uint256, _rewardCut uint256, _desc string) returns()
func (_Dosstaking *DosstakingTransactorSession) NewNode(_nodeAddr common.Address, _tokenAmount *big.Int, _dropburnAmount *big.Int, _rewardCut *big.Int, _desc string) (*types.Transaction, error) {
	return _Dosstaking.Contract.NewNode(&_Dosstaking.TransactOpts, _nodeAddr, _tokenAmount, _dropburnAmount, _rewardCut, _desc)
}

// NodeClaimReward is a paid mutator transaction binding the contract method 0xb41b0f96.
//
// Solidity: function nodeClaimReward(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) NodeClaimReward(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "nodeClaimReward", _nodeAddr)
}

// NodeClaimReward is a paid mutator transaction binding the contract method 0xb41b0f96.
//
// Solidity: function nodeClaimReward(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) NodeClaimReward(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeClaimReward(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeClaimReward is a paid mutator transaction binding the contract method 0xb41b0f96.
//
// Solidity: function nodeClaimReward(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) NodeClaimReward(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeClaimReward(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeStart is a paid mutator transaction binding the contract method 0x4c542d3d.
//
// Solidity: function nodeStart(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) NodeStart(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "nodeStart", _nodeAddr)
}

// NodeStart is a paid mutator transaction binding the contract method 0x4c542d3d.
//
// Solidity: function nodeStart(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) NodeStart(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeStart(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeStart is a paid mutator transaction binding the contract method 0x4c542d3d.
//
// Solidity: function nodeStart(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) NodeStart(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeStart(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeStop is a paid mutator transaction binding the contract method 0xc5375c29.
//
// Solidity: function nodeStop(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) NodeStop(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "nodeStop", _nodeAddr)
}

// NodeStop is a paid mutator transaction binding the contract method 0xc5375c29.
//
// Solidity: function nodeStop(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) NodeStop(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeStop(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeStop is a paid mutator transaction binding the contract method 0xc5375c29.
//
// Solidity: function nodeStop(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) NodeStop(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeStop(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeUnbond is a paid mutator transaction binding the contract method 0x508b74fe.
//
// Solidity: function nodeUnbond(_tokenAmount uint256, _dropburnAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) NodeUnbond(opts *bind.TransactOpts, _tokenAmount *big.Int, _dropburnAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "nodeUnbond", _tokenAmount, _dropburnAmount, _nodeAddr)
}

// NodeUnbond is a paid mutator transaction binding the contract method 0x508b74fe.
//
// Solidity: function nodeUnbond(_tokenAmount uint256, _dropburnAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) NodeUnbond(_tokenAmount *big.Int, _dropburnAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeUnbond(&_Dosstaking.TransactOpts, _tokenAmount, _dropburnAmount, _nodeAddr)
}

// NodeUnbond is a paid mutator transaction binding the contract method 0x508b74fe.
//
// Solidity: function nodeUnbond(_tokenAmount uint256, _dropburnAmount uint256, _nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) NodeUnbond(_tokenAmount *big.Int, _dropburnAmount *big.Int, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeUnbond(&_Dosstaking.TransactOpts, _tokenAmount, _dropburnAmount, _nodeAddr)
}

// NodeUnregister is a paid mutator transaction binding the contract method 0xa2772193.
//
// Solidity: function nodeUnregister(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) NodeUnregister(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "nodeUnregister", _nodeAddr)
}

// NodeUnregister is a paid mutator transaction binding the contract method 0xa2772193.
//
// Solidity: function nodeUnregister(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) NodeUnregister(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeUnregister(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeUnregister is a paid mutator transaction binding the contract method 0xa2772193.
//
// Solidity: function nodeUnregister(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) NodeUnregister(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeUnregister(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeWithdraw is a paid mutator transaction binding the contract method 0x4ecea80d.
//
// Solidity: function nodeWithdraw(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactor) NodeWithdraw(opts *bind.TransactOpts, _nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "nodeWithdraw", _nodeAddr)
}

// NodeWithdraw is a paid mutator transaction binding the contract method 0x4ecea80d.
//
// Solidity: function nodeWithdraw(_nodeAddr address) returns()
func (_Dosstaking *DosstakingSession) NodeWithdraw(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeWithdraw(&_Dosstaking.TransactOpts, _nodeAddr)
}

// NodeWithdraw is a paid mutator transaction binding the contract method 0x4ecea80d.
//
// Solidity: function nodeWithdraw(_nodeAddr address) returns()
func (_Dosstaking *DosstakingTransactorSession) NodeWithdraw(_nodeAddr common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.NodeWithdraw(&_Dosstaking.TransactOpts, _nodeAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosstaking *DosstakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosstaking *DosstakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dosstaking.Contract.RenounceOwnership(&_Dosstaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosstaking *DosstakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dosstaking.Contract.RenounceOwnership(&_Dosstaking.TransactOpts)
}

// SetCirculatingSupply is a paid mutator transaction binding the contract method 0xb1764071.
//
// Solidity: function setCirculatingSupply(_newSupply uint256) returns()
func (_Dosstaking *DosstakingTransactor) SetCirculatingSupply(opts *bind.TransactOpts, _newSupply *big.Int) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "setCirculatingSupply", _newSupply)
}

// SetCirculatingSupply is a paid mutator transaction binding the contract method 0xb1764071.
//
// Solidity: function setCirculatingSupply(_newSupply uint256) returns()
func (_Dosstaking *DosstakingSession) SetCirculatingSupply(_newSupply *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetCirculatingSupply(&_Dosstaking.TransactOpts, _newSupply)
}

// SetCirculatingSupply is a paid mutator transaction binding the contract method 0xb1764071.
//
// Solidity: function setCirculatingSupply(_newSupply uint256) returns()
func (_Dosstaking *DosstakingTransactorSession) SetCirculatingSupply(_newSupply *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetCirculatingSupply(&_Dosstaking.TransactOpts, _newSupply)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(_quota uint256) returns()
func (_Dosstaking *DosstakingTransactor) SetDropBurnMaxQuota(opts *bind.TransactOpts, _quota *big.Int) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "setDropBurnMaxQuota", _quota)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(_quota uint256) returns()
func (_Dosstaking *DosstakingSession) SetDropBurnMaxQuota(_quota *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetDropBurnMaxQuota(&_Dosstaking.TransactOpts, _quota)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(_quota uint256) returns()
func (_Dosstaking *DosstakingTransactorSession) SetDropBurnMaxQuota(_quota *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetDropBurnMaxQuota(&_Dosstaking.TransactOpts, _quota)
}

// SetMinStakePerNode is a paid mutator transaction binding the contract method 0x8b0bc845.
//
// Solidity: function setMinStakePerNode(_minStake uint256) returns()
func (_Dosstaking *DosstakingTransactor) SetMinStakePerNode(opts *bind.TransactOpts, _minStake *big.Int) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "setMinStakePerNode", _minStake)
}

// SetMinStakePerNode is a paid mutator transaction binding the contract method 0x8b0bc845.
//
// Solidity: function setMinStakePerNode(_minStake uint256) returns()
func (_Dosstaking *DosstakingSession) SetMinStakePerNode(_minStake *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetMinStakePerNode(&_Dosstaking.TransactOpts, _minStake)
}

// SetMinStakePerNode is a paid mutator transaction binding the contract method 0x8b0bc845.
//
// Solidity: function setMinStakePerNode(_minStake uint256) returns()
func (_Dosstaking *DosstakingTransactorSession) SetMinStakePerNode(_minStake *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetMinStakePerNode(&_Dosstaking.TransactOpts, _minStake)
}

// SetUnbondDuration is a paid mutator transaction binding the contract method 0xb16d2d58.
//
// Solidity: function setUnbondDuration(_duration uint256) returns()
func (_Dosstaking *DosstakingTransactor) SetUnbondDuration(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "setUnbondDuration", _duration)
}

// SetUnbondDuration is a paid mutator transaction binding the contract method 0xb16d2d58.
//
// Solidity: function setUnbondDuration(_duration uint256) returns()
func (_Dosstaking *DosstakingSession) SetUnbondDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetUnbondDuration(&_Dosstaking.TransactOpts, _duration)
}

// SetUnbondDuration is a paid mutator transaction binding the contract method 0xb16d2d58.
//
// Solidity: function setUnbondDuration(_duration uint256) returns()
func (_Dosstaking *DosstakingTransactorSession) SetUnbondDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.SetUnbondDuration(&_Dosstaking.TransactOpts, _duration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosstaking *DosstakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosstaking *DosstakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.TransferOwnership(&_Dosstaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosstaking *DosstakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosstaking.Contract.TransferOwnership(&_Dosstaking.TransactOpts, newOwner)
}

// UpdateNodeStaking is a paid mutator transaction binding the contract method 0x1ece5950.
//
// Solidity: function updateNodeStaking(_nodeAddr address, _newTokenAmount uint256, _newDropburnAmount uint256, _newCut uint256) returns()
func (_Dosstaking *DosstakingTransactor) UpdateNodeStaking(opts *bind.TransactOpts, _nodeAddr common.Address, _newTokenAmount *big.Int, _newDropburnAmount *big.Int, _newCut *big.Int) (*types.Transaction, error) {
	return _Dosstaking.contract.Transact(opts, "updateNodeStaking", _nodeAddr, _newTokenAmount, _newDropburnAmount, _newCut)
}

// UpdateNodeStaking is a paid mutator transaction binding the contract method 0x1ece5950.
//
// Solidity: function updateNodeStaking(_nodeAddr address, _newTokenAmount uint256, _newDropburnAmount uint256, _newCut uint256) returns()
func (_Dosstaking *DosstakingSession) UpdateNodeStaking(_nodeAddr common.Address, _newTokenAmount *big.Int, _newDropburnAmount *big.Int, _newCut *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.UpdateNodeStaking(&_Dosstaking.TransactOpts, _nodeAddr, _newTokenAmount, _newDropburnAmount, _newCut)
}

// UpdateNodeStaking is a paid mutator transaction binding the contract method 0x1ece5950.
//
// Solidity: function updateNodeStaking(_nodeAddr address, _newTokenAmount uint256, _newDropburnAmount uint256, _newCut uint256) returns()
func (_Dosstaking *DosstakingTransactorSession) UpdateNodeStaking(_nodeAddr common.Address, _newTokenAmount *big.Int, _newDropburnAmount *big.Int, _newCut *big.Int) (*types.Transaction, error) {
	return _Dosstaking.Contract.UpdateNodeStaking(&_Dosstaking.TransactOpts, _nodeAddr, _newTokenAmount, _newDropburnAmount, _newCut)
}

// DosstakingClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the Dosstaking contract.
type DosstakingClaimRewardIterator struct {
	Event *DosstakingClaimReward // Event containing the contract specifics and raw log

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
func (it *DosstakingClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingClaimReward)
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
		it.Event = new(DosstakingClaimReward)
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
func (it *DosstakingClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingClaimReward represents a ClaimReward event raised by the Dosstaking contract.
type DosstakingClaimReward struct {
	To         common.Address
	NodeRunner bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0x70f2476b8214c2f4e6fc47e4adb3c9adca3cc0e0e6c9a9d7b0e2626ce83cbc40.
//
// Solidity: e ClaimReward(to indexed address, nodeRunner bool, amount uint256)
func (_Dosstaking *DosstakingFilterer) FilterClaimReward(opts *bind.FilterOpts, to []common.Address) (*DosstakingClaimRewardIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "ClaimReward", toRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingClaimRewardIterator{contract: _Dosstaking.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x70f2476b8214c2f4e6fc47e4adb3c9adca3cc0e0e6c9a9d7b0e2626ce83cbc40.
//
// Solidity: e ClaimReward(to indexed address, nodeRunner bool, amount uint256)
func (_Dosstaking *DosstakingFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *DosstakingClaimReward, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "ClaimReward", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingClaimReward)
				if err := _Dosstaking.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// DosstakingDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the Dosstaking contract.
type DosstakingDelegateIterator struct {
	Event *DosstakingDelegate // Event containing the contract specifics and raw log

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
func (it *DosstakingDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingDelegate)
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
		it.Event = new(DosstakingDelegate)
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
func (it *DosstakingDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingDelegate represents a Delegate event raised by the Dosstaking contract.
type DosstakingDelegate struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: e Delegate(from indexed address, to indexed address, amount uint256)
func (_Dosstaking *DosstakingFilterer) FilterDelegate(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DosstakingDelegateIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "Delegate", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingDelegateIterator{contract: _Dosstaking.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: e Delegate(from indexed address, to indexed address, amount uint256)
func (_Dosstaking *DosstakingFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *DosstakingDelegate, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "Delegate", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingDelegate)
				if err := _Dosstaking.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// DosstakingNewNodeIterator is returned from FilterNewNode and is used to iterate over the raw logs and unpacked data for NewNode events raised by the Dosstaking contract.
type DosstakingNewNodeIterator struct {
	Event *DosstakingNewNode // Event containing the contract specifics and raw log

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
func (it *DosstakingNewNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingNewNode)
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
		it.Event = new(DosstakingNewNode)
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
func (it *DosstakingNewNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingNewNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingNewNode represents a NewNode event raised by the Dosstaking contract.
type DosstakingNewNode struct {
	Owner            common.Address
	NodeAddress      common.Address
	SelfStakedAmount *big.Int
	StakedDB         *big.Int
	RewardCut        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewNode is a free log retrieval operation binding the contract event 0x08f38bdd2a12d30a92014fb5523763d3d147ea395e40f81c241eea1a6af0cdfc.
//
// Solidity: e NewNode(owner indexed address, nodeAddress indexed address, selfStakedAmount uint256, stakedDB uint256, rewardCut uint256)
func (_Dosstaking *DosstakingFilterer) FilterNewNode(opts *bind.FilterOpts, owner []common.Address, nodeAddress []common.Address) (*DosstakingNewNodeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "NewNode", ownerRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingNewNodeIterator{contract: _Dosstaking.contract, event: "NewNode", logs: logs, sub: sub}, nil
}

// WatchNewNode is a free log subscription operation binding the contract event 0x08f38bdd2a12d30a92014fb5523763d3d147ea395e40f81c241eea1a6af0cdfc.
//
// Solidity: e NewNode(owner indexed address, nodeAddress indexed address, selfStakedAmount uint256, stakedDB uint256, rewardCut uint256)
func (_Dosstaking *DosstakingFilterer) WatchNewNode(opts *bind.WatchOpts, sink chan<- *DosstakingNewNode, owner []common.Address, nodeAddress []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "NewNode", ownerRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingNewNode)
				if err := _Dosstaking.contract.UnpackLog(event, "NewNode", log); err != nil {
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

// DosstakingOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Dosstaking contract.
type DosstakingOwnershipRenouncedIterator struct {
	Event *DosstakingOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DosstakingOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingOwnershipRenounced)
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
		it.Event = new(DosstakingOwnershipRenounced)
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
func (it *DosstakingOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingOwnershipRenounced represents a OwnershipRenounced event raised by the Dosstaking contract.
type DosstakingOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Dosstaking *DosstakingFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DosstakingOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingOwnershipRenouncedIterator{contract: _Dosstaking.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Dosstaking *DosstakingFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DosstakingOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingOwnershipRenounced)
				if err := _Dosstaking.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DosstakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dosstaking contract.
type DosstakingOwnershipTransferredIterator struct {
	Event *DosstakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DosstakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingOwnershipTransferred)
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
		it.Event = new(DosstakingOwnershipTransferred)
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
func (it *DosstakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingOwnershipTransferred represents a OwnershipTransferred event raised by the Dosstaking contract.
type DosstakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Dosstaking *DosstakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DosstakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingOwnershipTransferredIterator{contract: _Dosstaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Dosstaking *DosstakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DosstakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingOwnershipTransferred)
				if err := _Dosstaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DosstakingUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the Dosstaking contract.
type DosstakingUnbondIterator struct {
	Event *DosstakingUnbond // Event containing the contract specifics and raw log

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
func (it *DosstakingUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingUnbond)
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
		it.Event = new(DosstakingUnbond)
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
func (it *DosstakingUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingUnbond represents a Unbond event raised by the Dosstaking contract.
type DosstakingUnbond struct {
	From           common.Address
	To             common.Address
	NodeRunner     bool
	TokenAmount    *big.Int
	DropburnAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x882043d8baf498c3af74efc34e0973ddc34e0e5a0351d84132ab0c7032b89160.
//
// Solidity: e Unbond(from indexed address, to indexed address, nodeRunner bool, tokenAmount uint256, dropburnAmount uint256)
func (_Dosstaking *DosstakingFilterer) FilterUnbond(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DosstakingUnbondIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "Unbond", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingUnbondIterator{contract: _Dosstaking.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x882043d8baf498c3af74efc34e0973ddc34e0e5a0351d84132ab0c7032b89160.
//
// Solidity: e Unbond(from indexed address, to indexed address, nodeRunner bool, tokenAmount uint256, dropburnAmount uint256)
func (_Dosstaking *DosstakingFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *DosstakingUnbond, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "Unbond", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingUnbond)
				if err := _Dosstaking.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// DosstakingUpdateCirculatingSupplyIterator is returned from FilterUpdateCirculatingSupply and is used to iterate over the raw logs and unpacked data for UpdateCirculatingSupply events raised by the Dosstaking contract.
type DosstakingUpdateCirculatingSupplyIterator struct {
	Event *DosstakingUpdateCirculatingSupply // Event containing the contract specifics and raw log

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
func (it *DosstakingUpdateCirculatingSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingUpdateCirculatingSupply)
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
		it.Event = new(DosstakingUpdateCirculatingSupply)
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
func (it *DosstakingUpdateCirculatingSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingUpdateCirculatingSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingUpdateCirculatingSupply represents a UpdateCirculatingSupply event raised by the Dosstaking contract.
type DosstakingUpdateCirculatingSupply struct {
	OldCirculatingSupply *big.Int
	NewCirculatingSupply *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateCirculatingSupply is a free log retrieval operation binding the contract event 0xef08ea314c41dc78bc73676b156d3b7802c096dd7a552f7f514a80967aba10ac.
//
// Solidity: e UpdateCirculatingSupply(oldCirculatingSupply uint256, newCirculatingSupply uint256)
func (_Dosstaking *DosstakingFilterer) FilterUpdateCirculatingSupply(opts *bind.FilterOpts) (*DosstakingUpdateCirculatingSupplyIterator, error) {

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "UpdateCirculatingSupply")
	if err != nil {
		return nil, err
	}
	return &DosstakingUpdateCirculatingSupplyIterator{contract: _Dosstaking.contract, event: "UpdateCirculatingSupply", logs: logs, sub: sub}, nil
}

// WatchUpdateCirculatingSupply is a free log subscription operation binding the contract event 0xef08ea314c41dc78bc73676b156d3b7802c096dd7a552f7f514a80967aba10ac.
//
// Solidity: e UpdateCirculatingSupply(oldCirculatingSupply uint256, newCirculatingSupply uint256)
func (_Dosstaking *DosstakingFilterer) WatchUpdateCirculatingSupply(opts *bind.WatchOpts, sink chan<- *DosstakingUpdateCirculatingSupply) (event.Subscription, error) {

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "UpdateCirculatingSupply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingUpdateCirculatingSupply)
				if err := _Dosstaking.contract.UnpackLog(event, "UpdateCirculatingSupply", log); err != nil {
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

// DosstakingUpdateDropBurnMaxQuotaIterator is returned from FilterUpdateDropBurnMaxQuota and is used to iterate over the raw logs and unpacked data for UpdateDropBurnMaxQuota events raised by the Dosstaking contract.
type DosstakingUpdateDropBurnMaxQuotaIterator struct {
	Event *DosstakingUpdateDropBurnMaxQuota // Event containing the contract specifics and raw log

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
func (it *DosstakingUpdateDropBurnMaxQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingUpdateDropBurnMaxQuota)
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
		it.Event = new(DosstakingUpdateDropBurnMaxQuota)
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
func (it *DosstakingUpdateDropBurnMaxQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingUpdateDropBurnMaxQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingUpdateDropBurnMaxQuota represents a UpdateDropBurnMaxQuota event raised by the Dosstaking contract.
type DosstakingUpdateDropBurnMaxQuota struct {
	OldQuota *big.Int
	NewQuota *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateDropBurnMaxQuota is a free log retrieval operation binding the contract event 0x0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e99.
//
// Solidity: e UpdateDropBurnMaxQuota(oldQuota uint256, newQuota uint256)
func (_Dosstaking *DosstakingFilterer) FilterUpdateDropBurnMaxQuota(opts *bind.FilterOpts) (*DosstakingUpdateDropBurnMaxQuotaIterator, error) {

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "UpdateDropBurnMaxQuota")
	if err != nil {
		return nil, err
	}
	return &DosstakingUpdateDropBurnMaxQuotaIterator{contract: _Dosstaking.contract, event: "UpdateDropBurnMaxQuota", logs: logs, sub: sub}, nil
}

// WatchUpdateDropBurnMaxQuota is a free log subscription operation binding the contract event 0x0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e99.
//
// Solidity: e UpdateDropBurnMaxQuota(oldQuota uint256, newQuota uint256)
func (_Dosstaking *DosstakingFilterer) WatchUpdateDropBurnMaxQuota(opts *bind.WatchOpts, sink chan<- *DosstakingUpdateDropBurnMaxQuota) (event.Subscription, error) {

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "UpdateDropBurnMaxQuota")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingUpdateDropBurnMaxQuota)
				if err := _Dosstaking.contract.UnpackLog(event, "UpdateDropBurnMaxQuota", log); err != nil {
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

// DosstakingUpdateMinStakePerNodeIterator is returned from FilterUpdateMinStakePerNode and is used to iterate over the raw logs and unpacked data for UpdateMinStakePerNode events raised by the Dosstaking contract.
type DosstakingUpdateMinStakePerNodeIterator struct {
	Event *DosstakingUpdateMinStakePerNode // Event containing the contract specifics and raw log

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
func (it *DosstakingUpdateMinStakePerNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingUpdateMinStakePerNode)
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
		it.Event = new(DosstakingUpdateMinStakePerNode)
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
func (it *DosstakingUpdateMinStakePerNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingUpdateMinStakePerNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingUpdateMinStakePerNode represents a UpdateMinStakePerNode event raised by the Dosstaking contract.
type DosstakingUpdateMinStakePerNode struct {
	OldMinStakePerNode *big.Int
	NewMinStakePerNode *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinStakePerNode is a free log retrieval operation binding the contract event 0x89892efa8f66455fa9bf996b1444e79778d71795cf00089fadf89071a3896ebb.
//
// Solidity: e UpdateMinStakePerNode(oldMinStakePerNode uint256, newMinStakePerNode uint256)
func (_Dosstaking *DosstakingFilterer) FilterUpdateMinStakePerNode(opts *bind.FilterOpts) (*DosstakingUpdateMinStakePerNodeIterator, error) {

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "UpdateMinStakePerNode")
	if err != nil {
		return nil, err
	}
	return &DosstakingUpdateMinStakePerNodeIterator{contract: _Dosstaking.contract, event: "UpdateMinStakePerNode", logs: logs, sub: sub}, nil
}

// WatchUpdateMinStakePerNode is a free log subscription operation binding the contract event 0x89892efa8f66455fa9bf996b1444e79778d71795cf00089fadf89071a3896ebb.
//
// Solidity: e UpdateMinStakePerNode(oldMinStakePerNode uint256, newMinStakePerNode uint256)
func (_Dosstaking *DosstakingFilterer) WatchUpdateMinStakePerNode(opts *bind.WatchOpts, sink chan<- *DosstakingUpdateMinStakePerNode) (event.Subscription, error) {

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "UpdateMinStakePerNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingUpdateMinStakePerNode)
				if err := _Dosstaking.contract.UnpackLog(event, "UpdateMinStakePerNode", log); err != nil {
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

// DosstakingUpdateUnbondDurationIterator is returned from FilterUpdateUnbondDuration and is used to iterate over the raw logs and unpacked data for UpdateUnbondDuration events raised by the Dosstaking contract.
type DosstakingUpdateUnbondDurationIterator struct {
	Event *DosstakingUpdateUnbondDuration // Event containing the contract specifics and raw log

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
func (it *DosstakingUpdateUnbondDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingUpdateUnbondDuration)
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
		it.Event = new(DosstakingUpdateUnbondDuration)
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
func (it *DosstakingUpdateUnbondDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingUpdateUnbondDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingUpdateUnbondDuration represents a UpdateUnbondDuration event raised by the Dosstaking contract.
type DosstakingUpdateUnbondDuration struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateUnbondDuration is a free log retrieval operation binding the contract event 0x8671d68a1c48206a7cde676b141ee222482f2014d09ad75af8f5ee2118af9b99.
//
// Solidity: e UpdateUnbondDuration(oldDuration uint256, newDuration uint256)
func (_Dosstaking *DosstakingFilterer) FilterUpdateUnbondDuration(opts *bind.FilterOpts) (*DosstakingUpdateUnbondDurationIterator, error) {

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "UpdateUnbondDuration")
	if err != nil {
		return nil, err
	}
	return &DosstakingUpdateUnbondDurationIterator{contract: _Dosstaking.contract, event: "UpdateUnbondDuration", logs: logs, sub: sub}, nil
}

// WatchUpdateUnbondDuration is a free log subscription operation binding the contract event 0x8671d68a1c48206a7cde676b141ee222482f2014d09ad75af8f5ee2118af9b99.
//
// Solidity: e UpdateUnbondDuration(oldDuration uint256, newDuration uint256)
func (_Dosstaking *DosstakingFilterer) WatchUpdateUnbondDuration(opts *bind.WatchOpts, sink chan<- *DosstakingUpdateUnbondDuration) (event.Subscription, error) {

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "UpdateUnbondDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingUpdateUnbondDuration)
				if err := _Dosstaking.contract.UnpackLog(event, "UpdateUnbondDuration", log); err != nil {
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

// DosstakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Dosstaking contract.
type DosstakingWithdrawIterator struct {
	Event *DosstakingWithdraw // Event containing the contract specifics and raw log

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
func (it *DosstakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosstakingWithdraw)
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
		it.Event = new(DosstakingWithdraw)
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
func (it *DosstakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosstakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosstakingWithdraw represents a Withdraw event raised by the Dosstaking contract.
type DosstakingWithdraw struct {
	From        common.Address
	To          common.Address
	NodeRunner  bool
	TokenAmount *big.Int
	DbAmount    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: e Withdraw(from indexed address, to indexed address, nodeRunner bool, tokenAmount uint256, dbAmount uint256)
func (_Dosstaking *DosstakingFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DosstakingWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.FilterLogs(opts, "Withdraw", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DosstakingWithdrawIterator{contract: _Dosstaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: e Withdraw(from indexed address, to indexed address, nodeRunner bool, tokenAmount uint256, dbAmount uint256)
func (_Dosstaking *DosstakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DosstakingWithdraw, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dosstaking.contract.WatchLogs(opts, "Withdraw", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosstakingWithdraw)
				if err := _Dosstaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
