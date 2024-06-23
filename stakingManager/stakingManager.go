// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingManager

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

// IStakingManagerDepositData is an auto generated low-level Go binding around an user-defined struct.
type IStakingManagerDepositData struct {
	PublicKey                        []byte
	Signature                        []byte
	DepositDataRoot                  [32]byte
	IpfsHashForEncryptedValidatorKey string
}

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ALREADY_SET\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DepositCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawSafe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restaked\",\"type\":\"bool\"}],\"name\":\"StakeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumILiquidityPool.SourceOfFunds\",\"name\":\"source\",\"type\":\"uint8\"}],\"name\":\"StakeSource\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bNftOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tNftOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsHashForEncryptedValidatorKey\",\"type\":\"string\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BNFTInterfaceInstance\",\"outputs\":[{\"internalType\":\"contractIBNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPRECATED_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TNFTInterfaceInstance\",\"outputs\":[{\"internalType\":\"contractITNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionManager\",\"outputs\":[{\"internalType\":\"contractIAuctionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorId\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubKey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signature\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRootApproval\",\"type\":\"bytes32[]\"}],\"name\":\"batchApproveRegistration\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_validatorIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"batchCancelDepositAsBnftHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_candidateBidIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfValidators\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tnftHolder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnftHolder\",\"type\":\"address\"},{\"internalType\":\"enumILiquidityPool.SourceOfFunds\",\"name\":\"_source\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_enableRestaking\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_validatorIdToShareWithdrawalSafe\",\"type\":\"uint256\"}],\"name\":\"batchDepositWithBidIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_candidateBidIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_enableRestaking\",\"type\":\"bool\"}],\"name\":\"batchDepositWithBidIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_validatorId\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_bNftRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tNftRecipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsHashForEncryptedValidatorKey\",\"type\":\"string\"}],\"internalType\":\"structIStakingManager.DepositData[]\",\"name\":\"_depositData\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"batchRegisterValidators\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_validatorId\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsHashForEncryptedValidatorKey\",\"type\":\"string\"}],\"internalType\":\"structIStakingManager.DepositData[]\",\"name\":\"_depositData\",\"type\":\"tuple[]\"}],\"name\":\"batchRegisterValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"bidIdToStaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bidIdToStakerInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumILiquidityPool.SourceOfFunds\",\"name\":\"sourceOfFund\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContractEth2\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEtherFiNodeBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auctionAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContractAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_etherFiAdmin\",\"type\":\"address\"}],\"name\":\"initializeOnUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_createEigenPod\",\"type\":\"bool\"}],\"name\":\"instantiateEtherFiNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFullStakeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPoolContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchDepositSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodesManager\",\"outputs\":[{\"internalType\":\"contractIEtherFiNodesManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bnftAddress\",\"type\":\"address\"}],\"name\":\"registerBNFTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_etherFiNodeImplementationContract\",\"type\":\"address\"}],\"name\":\"registerEtherFiNodeImplementationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tnftAddress\",\"type\":\"address\"}],\"name\":\"registerTNFTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodesManagerAddress\",\"type\":\"address\"}],\"name\":\"setEtherFiNodesManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityPoolAddress\",\"type\":\"address\"}],\"name\":\"setLiquidityPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_newMaxBatchDepositSize\",\"type\":\"uint128\"}],\"name\":\"setMaxBatchDepositSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperateManager\",\"type\":\"address\"}],\"name\":\"setNodeOperatorManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateFullStakingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeEtherFiNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// BNFTInterfaceInstance is a free data retrieval call binding the contract method 0x2273b411.
//
// Solidity: function BNFTInterfaceInstance() view returns(address)
func (_StakingManager *StakingManagerCaller) BNFTInterfaceInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "BNFTInterfaceInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BNFTInterfaceInstance is a free data retrieval call binding the contract method 0x2273b411.
//
// Solidity: function BNFTInterfaceInstance() view returns(address)
func (_StakingManager *StakingManagerSession) BNFTInterfaceInstance() (common.Address, error) {
	return _StakingManager.Contract.BNFTInterfaceInstance(&_StakingManager.CallOpts)
}

// BNFTInterfaceInstance is a free data retrieval call binding the contract method 0x2273b411.
//
// Solidity: function BNFTInterfaceInstance() view returns(address)
func (_StakingManager *StakingManagerCallerSession) BNFTInterfaceInstance() (common.Address, error) {
	return _StakingManager.Contract.BNFTInterfaceInstance(&_StakingManager.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_StakingManager *StakingManagerCaller) DEPRECATEDAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "DEPRECATED_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_StakingManager *StakingManagerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _StakingManager.Contract.DEPRECATEDAdmin(&_StakingManager.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_StakingManager *StakingManagerCallerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _StakingManager.Contract.DEPRECATEDAdmin(&_StakingManager.CallOpts)
}

// TNFTInterfaceInstance is a free data retrieval call binding the contract method 0x6d4e148d.
//
// Solidity: function TNFTInterfaceInstance() view returns(address)
func (_StakingManager *StakingManagerCaller) TNFTInterfaceInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "TNFTInterfaceInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TNFTInterfaceInstance is a free data retrieval call binding the contract method 0x6d4e148d.
//
// Solidity: function TNFTInterfaceInstance() view returns(address)
func (_StakingManager *StakingManagerSession) TNFTInterfaceInstance() (common.Address, error) {
	return _StakingManager.Contract.TNFTInterfaceInstance(&_StakingManager.CallOpts)
}

// TNFTInterfaceInstance is a free data retrieval call binding the contract method 0x6d4e148d.
//
// Solidity: function TNFTInterfaceInstance() view returns(address)
func (_StakingManager *StakingManagerCallerSession) TNFTInterfaceInstance() (common.Address, error) {
	return _StakingManager.Contract.TNFTInterfaceInstance(&_StakingManager.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_StakingManager *StakingManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_StakingManager *StakingManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _StakingManager.Contract.Admins(&_StakingManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _StakingManager.Contract.Admins(&_StakingManager.CallOpts, arg0)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_StakingManager *StakingManagerCaller) AuctionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "auctionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_StakingManager *StakingManagerSession) AuctionManager() (common.Address, error) {
	return _StakingManager.Contract.AuctionManager(&_StakingManager.CallOpts)
}

// AuctionManager is a free data retrieval call binding the contract method 0xb0192f9a.
//
// Solidity: function auctionManager() view returns(address)
func (_StakingManager *StakingManagerCallerSession) AuctionManager() (common.Address, error) {
	return _StakingManager.Contract.AuctionManager(&_StakingManager.CallOpts)
}

// BidIdToStaker is a free data retrieval call binding the contract method 0x1183b6b8.
//
// Solidity: function bidIdToStaker(uint256 id) view returns(address)
func (_StakingManager *StakingManagerCaller) BidIdToStaker(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "bidIdToStaker", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BidIdToStaker is a free data retrieval call binding the contract method 0x1183b6b8.
//
// Solidity: function bidIdToStaker(uint256 id) view returns(address)
func (_StakingManager *StakingManagerSession) BidIdToStaker(id *big.Int) (common.Address, error) {
	return _StakingManager.Contract.BidIdToStaker(&_StakingManager.CallOpts, id)
}

// BidIdToStaker is a free data retrieval call binding the contract method 0x1183b6b8.
//
// Solidity: function bidIdToStaker(uint256 id) view returns(address)
func (_StakingManager *StakingManagerCallerSession) BidIdToStaker(id *big.Int) (common.Address, error) {
	return _StakingManager.Contract.BidIdToStaker(&_StakingManager.CallOpts, id)
}

// BidIdToStakerInfo is a free data retrieval call binding the contract method 0xe4b94b21.
//
// Solidity: function bidIdToStakerInfo(uint256 ) view returns(address staker, uint8 sourceOfFund)
func (_StakingManager *StakingManagerCaller) BidIdToStakerInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Staker       common.Address
	SourceOfFund uint8
}, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "bidIdToStakerInfo", arg0)

	outstruct := new(struct {
		Staker       common.Address
		SourceOfFund uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Staker = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SourceOfFund = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// BidIdToStakerInfo is a free data retrieval call binding the contract method 0xe4b94b21.
//
// Solidity: function bidIdToStakerInfo(uint256 ) view returns(address staker, uint8 sourceOfFund)
func (_StakingManager *StakingManagerSession) BidIdToStakerInfo(arg0 *big.Int) (struct {
	Staker       common.Address
	SourceOfFund uint8
}, error) {
	return _StakingManager.Contract.BidIdToStakerInfo(&_StakingManager.CallOpts, arg0)
}

// BidIdToStakerInfo is a free data retrieval call binding the contract method 0xe4b94b21.
//
// Solidity: function bidIdToStakerInfo(uint256 ) view returns(address staker, uint8 sourceOfFund)
func (_StakingManager *StakingManagerCallerSession) BidIdToStakerInfo(arg0 *big.Int) (struct {
	Staker       common.Address
	SourceOfFund uint8
}, error) {
	return _StakingManager.Contract.BidIdToStakerInfo(&_StakingManager.CallOpts, arg0)
}

// DepositContractEth2 is a free data retrieval call binding the contract method 0x82699d5a.
//
// Solidity: function depositContractEth2() view returns(address)
func (_StakingManager *StakingManagerCaller) DepositContractEth2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "depositContractEth2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContractEth2 is a free data retrieval call binding the contract method 0x82699d5a.
//
// Solidity: function depositContractEth2() view returns(address)
func (_StakingManager *StakingManagerSession) DepositContractEth2() (common.Address, error) {
	return _StakingManager.Contract.DepositContractEth2(&_StakingManager.CallOpts)
}

// DepositContractEth2 is a free data retrieval call binding the contract method 0x82699d5a.
//
// Solidity: function depositContractEth2() view returns(address)
func (_StakingManager *StakingManagerCallerSession) DepositContractEth2() (common.Address, error) {
	return _StakingManager.Contract.DepositContractEth2(&_StakingManager.CallOpts)
}

// GetEtherFiNodeBeacon is a free data retrieval call binding the contract method 0xe9c99b6b.
//
// Solidity: function getEtherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerCaller) GetEtherFiNodeBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getEtherFiNodeBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEtherFiNodeBeacon is a free data retrieval call binding the contract method 0xe9c99b6b.
//
// Solidity: function getEtherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerSession) GetEtherFiNodeBeacon() (common.Address, error) {
	return _StakingManager.Contract.GetEtherFiNodeBeacon(&_StakingManager.CallOpts)
}

// GetEtherFiNodeBeacon is a free data retrieval call binding the contract method 0xe9c99b6b.
//
// Solidity: function getEtherFiNodeBeacon() view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetEtherFiNodeBeacon() (common.Address, error) {
	return _StakingManager.Contract.GetEtherFiNodeBeacon(&_StakingManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StakingManager *StakingManagerCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StakingManager *StakingManagerSession) GetImplementation() (common.Address, error) {
	return _StakingManager.Contract.GetImplementation(&_StakingManager.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetImplementation() (common.Address, error) {
	return _StakingManager.Contract.GetImplementation(&_StakingManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StakingManager *StakingManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StakingManager *StakingManagerSession) Implementation() (common.Address, error) {
	return _StakingManager.Contract.Implementation(&_StakingManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Implementation() (common.Address, error) {
	return _StakingManager.Contract.Implementation(&_StakingManager.CallOpts)
}

// ImplementationContract is a free data retrieval call binding the contract method 0x99e7d056.
//
// Solidity: function implementationContract() view returns(address)
func (_StakingManager *StakingManagerCaller) ImplementationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "implementationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ImplementationContract is a free data retrieval call binding the contract method 0x99e7d056.
//
// Solidity: function implementationContract() view returns(address)
func (_StakingManager *StakingManagerSession) ImplementationContract() (common.Address, error) {
	return _StakingManager.Contract.ImplementationContract(&_StakingManager.CallOpts)
}

// ImplementationContract is a free data retrieval call binding the contract method 0x99e7d056.
//
// Solidity: function implementationContract() view returns(address)
func (_StakingManager *StakingManagerCallerSession) ImplementationContract() (common.Address, error) {
	return _StakingManager.Contract.ImplementationContract(&_StakingManager.CallOpts)
}

// IsFullStakeEnabled is a free data retrieval call binding the contract method 0xcf12557a.
//
// Solidity: function isFullStakeEnabled() view returns(bool)
func (_StakingManager *StakingManagerCaller) IsFullStakeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "isFullStakeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFullStakeEnabled is a free data retrieval call binding the contract method 0xcf12557a.
//
// Solidity: function isFullStakeEnabled() view returns(bool)
func (_StakingManager *StakingManagerSession) IsFullStakeEnabled() (bool, error) {
	return _StakingManager.Contract.IsFullStakeEnabled(&_StakingManager.CallOpts)
}

// IsFullStakeEnabled is a free data retrieval call binding the contract method 0xcf12557a.
//
// Solidity: function isFullStakeEnabled() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsFullStakeEnabled() (bool, error) {
	return _StakingManager.Contract.IsFullStakeEnabled(&_StakingManager.CallOpts)
}

// LiquidityPoolContract is a free data retrieval call binding the contract method 0x1d8ef0fe.
//
// Solidity: function liquidityPoolContract() view returns(address)
func (_StakingManager *StakingManagerCaller) LiquidityPoolContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "liquidityPoolContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPoolContract is a free data retrieval call binding the contract method 0x1d8ef0fe.
//
// Solidity: function liquidityPoolContract() view returns(address)
func (_StakingManager *StakingManagerSession) LiquidityPoolContract() (common.Address, error) {
	return _StakingManager.Contract.LiquidityPoolContract(&_StakingManager.CallOpts)
}

// LiquidityPoolContract is a free data retrieval call binding the contract method 0x1d8ef0fe.
//
// Solidity: function liquidityPoolContract() view returns(address)
func (_StakingManager *StakingManagerCallerSession) LiquidityPoolContract() (common.Address, error) {
	return _StakingManager.Contract.LiquidityPoolContract(&_StakingManager.CallOpts)
}

// MaxBatchDepositSize is a free data retrieval call binding the contract method 0x3c219b49.
//
// Solidity: function maxBatchDepositSize() view returns(uint128)
func (_StakingManager *StakingManagerCaller) MaxBatchDepositSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "maxBatchDepositSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBatchDepositSize is a free data retrieval call binding the contract method 0x3c219b49.
//
// Solidity: function maxBatchDepositSize() view returns(uint128)
func (_StakingManager *StakingManagerSession) MaxBatchDepositSize() (*big.Int, error) {
	return _StakingManager.Contract.MaxBatchDepositSize(&_StakingManager.CallOpts)
}

// MaxBatchDepositSize is a free data retrieval call binding the contract method 0x3c219b49.
//
// Solidity: function maxBatchDepositSize() view returns(uint128)
func (_StakingManager *StakingManagerCallerSession) MaxBatchDepositSize() (*big.Int, error) {
	return _StakingManager.Contract.MaxBatchDepositSize(&_StakingManager.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_StakingManager *StakingManagerSession) MerkleRoot() ([32]byte, error) {
	return _StakingManager.Contract.MerkleRoot(&_StakingManager.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) MerkleRoot() ([32]byte, error) {
	return _StakingManager.Contract.MerkleRoot(&_StakingManager.CallOpts)
}

// NodeOperatorManager is a free data retrieval call binding the contract method 0xe30f7c4f.
//
// Solidity: function nodeOperatorManager() view returns(address)
func (_StakingManager *StakingManagerCaller) NodeOperatorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "nodeOperatorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorManager is a free data retrieval call binding the contract method 0xe30f7c4f.
//
// Solidity: function nodeOperatorManager() view returns(address)
func (_StakingManager *StakingManagerSession) NodeOperatorManager() (common.Address, error) {
	return _StakingManager.Contract.NodeOperatorManager(&_StakingManager.CallOpts)
}

// NodeOperatorManager is a free data retrieval call binding the contract method 0xe30f7c4f.
//
// Solidity: function nodeOperatorManager() view returns(address)
func (_StakingManager *StakingManagerCallerSession) NodeOperatorManager() (common.Address, error) {
	return _StakingManager.Contract.NodeOperatorManager(&_StakingManager.CallOpts)
}

// NodesManager is a free data retrieval call binding the contract method 0x469963aa.
//
// Solidity: function nodesManager() view returns(address)
func (_StakingManager *StakingManagerCaller) NodesManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "nodesManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodesManager is a free data retrieval call binding the contract method 0x469963aa.
//
// Solidity: function nodesManager() view returns(address)
func (_StakingManager *StakingManagerSession) NodesManager() (common.Address, error) {
	return _StakingManager.Contract.NodesManager(&_StakingManager.CallOpts)
}

// NodesManager is a free data retrieval call binding the contract method 0x469963aa.
//
// Solidity: function nodesManager() view returns(address)
func (_StakingManager *StakingManagerCallerSession) NodesManager() (common.Address, error) {
	return _StakingManager.Contract.NodesManager(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerSession) Paused() (bool, error) {
	return _StakingManager.Contract.Paused(&_StakingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) Paused() (bool, error) {
	return _StakingManager.Contract.Paused(&_StakingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingManager.Contract.ProxiableUUID(&_StakingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StakingManager.Contract.ProxiableUUID(&_StakingManager.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint128)
func (_StakingManager *StakingManagerCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint128)
func (_StakingManager *StakingManagerSession) StakeAmount() (*big.Int, error) {
	return _StakingManager.Contract.StakeAmount(&_StakingManager.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint128)
func (_StakingManager *StakingManagerCallerSession) StakeAmount() (*big.Int, error) {
	return _StakingManager.Contract.StakeAmount(&_StakingManager.CallOpts)
}

// BatchApproveRegistration is a paid mutator transaction binding the contract method 0x4cfc6c73.
//
// Solidity: function batchApproveRegistration(uint256[] _validatorId, bytes[] _pubKey, bytes[] _signature, bytes32[] _depositDataRootApproval) payable returns()
func (_StakingManager *StakingManagerTransactor) BatchApproveRegistration(opts *bind.TransactOpts, _validatorId []*big.Int, _pubKey [][]byte, _signature [][]byte, _depositDataRootApproval [][32]byte) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchApproveRegistration", _validatorId, _pubKey, _signature, _depositDataRootApproval)
}

// BatchApproveRegistration is a paid mutator transaction binding the contract method 0x4cfc6c73.
//
// Solidity: function batchApproveRegistration(uint256[] _validatorId, bytes[] _pubKey, bytes[] _signature, bytes32[] _depositDataRootApproval) payable returns()
func (_StakingManager *StakingManagerSession) BatchApproveRegistration(_validatorId []*big.Int, _pubKey [][]byte, _signature [][]byte, _depositDataRootApproval [][32]byte) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchApproveRegistration(&_StakingManager.TransactOpts, _validatorId, _pubKey, _signature, _depositDataRootApproval)
}

// BatchApproveRegistration is a paid mutator transaction binding the contract method 0x4cfc6c73.
//
// Solidity: function batchApproveRegistration(uint256[] _validatorId, bytes[] _pubKey, bytes[] _signature, bytes32[] _depositDataRootApproval) payable returns()
func (_StakingManager *StakingManagerTransactorSession) BatchApproveRegistration(_validatorId []*big.Int, _pubKey [][]byte, _signature [][]byte, _depositDataRootApproval [][32]byte) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchApproveRegistration(&_StakingManager.TransactOpts, _validatorId, _pubKey, _signature, _depositDataRootApproval)
}

// BatchCancelDeposit is a paid mutator transaction binding the contract method 0x89a8d9f5.
//
// Solidity: function batchCancelDeposit(uint256[] _validatorIds) returns()
func (_StakingManager *StakingManagerTransactor) BatchCancelDeposit(opts *bind.TransactOpts, _validatorIds []*big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchCancelDeposit", _validatorIds)
}

// BatchCancelDeposit is a paid mutator transaction binding the contract method 0x89a8d9f5.
//
// Solidity: function batchCancelDeposit(uint256[] _validatorIds) returns()
func (_StakingManager *StakingManagerSession) BatchCancelDeposit(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchCancelDeposit(&_StakingManager.TransactOpts, _validatorIds)
}

// BatchCancelDeposit is a paid mutator transaction binding the contract method 0x89a8d9f5.
//
// Solidity: function batchCancelDeposit(uint256[] _validatorIds) returns()
func (_StakingManager *StakingManagerTransactorSession) BatchCancelDeposit(_validatorIds []*big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchCancelDeposit(&_StakingManager.TransactOpts, _validatorIds)
}

// BatchCancelDepositAsBnftHolder is a paid mutator transaction binding the contract method 0x91515281.
//
// Solidity: function batchCancelDepositAsBnftHolder(uint256[] _validatorIds, address _caller) returns()
func (_StakingManager *StakingManagerTransactor) BatchCancelDepositAsBnftHolder(opts *bind.TransactOpts, _validatorIds []*big.Int, _caller common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchCancelDepositAsBnftHolder", _validatorIds, _caller)
}

// BatchCancelDepositAsBnftHolder is a paid mutator transaction binding the contract method 0x91515281.
//
// Solidity: function batchCancelDepositAsBnftHolder(uint256[] _validatorIds, address _caller) returns()
func (_StakingManager *StakingManagerSession) BatchCancelDepositAsBnftHolder(_validatorIds []*big.Int, _caller common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchCancelDepositAsBnftHolder(&_StakingManager.TransactOpts, _validatorIds, _caller)
}

// BatchCancelDepositAsBnftHolder is a paid mutator transaction binding the contract method 0x91515281.
//
// Solidity: function batchCancelDepositAsBnftHolder(uint256[] _validatorIds, address _caller) returns()
func (_StakingManager *StakingManagerTransactorSession) BatchCancelDepositAsBnftHolder(_validatorIds []*big.Int, _caller common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchCancelDepositAsBnftHolder(&_StakingManager.TransactOpts, _validatorIds, _caller)
}

// BatchDepositWithBidIds is a paid mutator transaction binding the contract method 0x47d67886.
//
// Solidity: function batchDepositWithBidIds(uint256[] _candidateBidIds, uint256 _numberOfValidators, address _staker, address _tnftHolder, address _bnftHolder, uint8 _source, bool _enableRestaking, uint256 _validatorIdToShareWithdrawalSafe) returns(uint256[])
func (_StakingManager *StakingManagerTransactor) BatchDepositWithBidIds(opts *bind.TransactOpts, _candidateBidIds []*big.Int, _numberOfValidators *big.Int, _staker common.Address, _tnftHolder common.Address, _bnftHolder common.Address, _source uint8, _enableRestaking bool, _validatorIdToShareWithdrawalSafe *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchDepositWithBidIds", _candidateBidIds, _numberOfValidators, _staker, _tnftHolder, _bnftHolder, _source, _enableRestaking, _validatorIdToShareWithdrawalSafe)
}

// BatchDepositWithBidIds is a paid mutator transaction binding the contract method 0x47d67886.
//
// Solidity: function batchDepositWithBidIds(uint256[] _candidateBidIds, uint256 _numberOfValidators, address _staker, address _tnftHolder, address _bnftHolder, uint8 _source, bool _enableRestaking, uint256 _validatorIdToShareWithdrawalSafe) returns(uint256[])
func (_StakingManager *StakingManagerSession) BatchDepositWithBidIds(_candidateBidIds []*big.Int, _numberOfValidators *big.Int, _staker common.Address, _tnftHolder common.Address, _bnftHolder common.Address, _source uint8, _enableRestaking bool, _validatorIdToShareWithdrawalSafe *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchDepositWithBidIds(&_StakingManager.TransactOpts, _candidateBidIds, _numberOfValidators, _staker, _tnftHolder, _bnftHolder, _source, _enableRestaking, _validatorIdToShareWithdrawalSafe)
}

// BatchDepositWithBidIds is a paid mutator transaction binding the contract method 0x47d67886.
//
// Solidity: function batchDepositWithBidIds(uint256[] _candidateBidIds, uint256 _numberOfValidators, address _staker, address _tnftHolder, address _bnftHolder, uint8 _source, bool _enableRestaking, uint256 _validatorIdToShareWithdrawalSafe) returns(uint256[])
func (_StakingManager *StakingManagerTransactorSession) BatchDepositWithBidIds(_candidateBidIds []*big.Int, _numberOfValidators *big.Int, _staker common.Address, _tnftHolder common.Address, _bnftHolder common.Address, _source uint8, _enableRestaking bool, _validatorIdToShareWithdrawalSafe *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchDepositWithBidIds(&_StakingManager.TransactOpts, _candidateBidIds, _numberOfValidators, _staker, _tnftHolder, _bnftHolder, _source, _enableRestaking, _validatorIdToShareWithdrawalSafe)
}

// BatchDepositWithBidIds0 is a paid mutator transaction binding the contract method 0xc77477a1.
//
// Solidity: function batchDepositWithBidIds(uint256[] _candidateBidIds, bool _enableRestaking) payable returns(uint256[])
func (_StakingManager *StakingManagerTransactor) BatchDepositWithBidIds0(opts *bind.TransactOpts, _candidateBidIds []*big.Int, _enableRestaking bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchDepositWithBidIds0", _candidateBidIds, _enableRestaking)
}

// BatchDepositWithBidIds0 is a paid mutator transaction binding the contract method 0xc77477a1.
//
// Solidity: function batchDepositWithBidIds(uint256[] _candidateBidIds, bool _enableRestaking) payable returns(uint256[])
func (_StakingManager *StakingManagerSession) BatchDepositWithBidIds0(_candidateBidIds []*big.Int, _enableRestaking bool) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchDepositWithBidIds0(&_StakingManager.TransactOpts, _candidateBidIds, _enableRestaking)
}

// BatchDepositWithBidIds0 is a paid mutator transaction binding the contract method 0xc77477a1.
//
// Solidity: function batchDepositWithBidIds(uint256[] _candidateBidIds, bool _enableRestaking) payable returns(uint256[])
func (_StakingManager *StakingManagerTransactorSession) BatchDepositWithBidIds0(_candidateBidIds []*big.Int, _enableRestaking bool) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchDepositWithBidIds0(&_StakingManager.TransactOpts, _candidateBidIds, _enableRestaking)
}

// BatchRegisterValidators is a paid mutator transaction binding the contract method 0x6413cc08.
//
// Solidity: function batchRegisterValidators(bytes32 _depositRoot, uint256[] _validatorId, address _bNftRecipient, address _tNftRecipient, (bytes,bytes,bytes32,string)[] _depositData, address _staker) payable returns()
func (_StakingManager *StakingManagerTransactor) BatchRegisterValidators(opts *bind.TransactOpts, _depositRoot [32]byte, _validatorId []*big.Int, _bNftRecipient common.Address, _tNftRecipient common.Address, _depositData []IStakingManagerDepositData, _staker common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchRegisterValidators", _depositRoot, _validatorId, _bNftRecipient, _tNftRecipient, _depositData, _staker)
}

// BatchRegisterValidators is a paid mutator transaction binding the contract method 0x6413cc08.
//
// Solidity: function batchRegisterValidators(bytes32 _depositRoot, uint256[] _validatorId, address _bNftRecipient, address _tNftRecipient, (bytes,bytes,bytes32,string)[] _depositData, address _staker) payable returns()
func (_StakingManager *StakingManagerSession) BatchRegisterValidators(_depositRoot [32]byte, _validatorId []*big.Int, _bNftRecipient common.Address, _tNftRecipient common.Address, _depositData []IStakingManagerDepositData, _staker common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchRegisterValidators(&_StakingManager.TransactOpts, _depositRoot, _validatorId, _bNftRecipient, _tNftRecipient, _depositData, _staker)
}

// BatchRegisterValidators is a paid mutator transaction binding the contract method 0x6413cc08.
//
// Solidity: function batchRegisterValidators(bytes32 _depositRoot, uint256[] _validatorId, address _bNftRecipient, address _tNftRecipient, (bytes,bytes,bytes32,string)[] _depositData, address _staker) payable returns()
func (_StakingManager *StakingManagerTransactorSession) BatchRegisterValidators(_depositRoot [32]byte, _validatorId []*big.Int, _bNftRecipient common.Address, _tNftRecipient common.Address, _depositData []IStakingManagerDepositData, _staker common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchRegisterValidators(&_StakingManager.TransactOpts, _depositRoot, _validatorId, _bNftRecipient, _tNftRecipient, _depositData, _staker)
}

// BatchRegisterValidators0 is a paid mutator transaction binding the contract method 0xb75a60d8.
//
// Solidity: function batchRegisterValidators(bytes32 _depositRoot, uint256[] _validatorId, (bytes,bytes,bytes32,string)[] _depositData) returns()
func (_StakingManager *StakingManagerTransactor) BatchRegisterValidators0(opts *bind.TransactOpts, _depositRoot [32]byte, _validatorId []*big.Int, _depositData []IStakingManagerDepositData) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "batchRegisterValidators0", _depositRoot, _validatorId, _depositData)
}

// BatchRegisterValidators0 is a paid mutator transaction binding the contract method 0xb75a60d8.
//
// Solidity: function batchRegisterValidators(bytes32 _depositRoot, uint256[] _validatorId, (bytes,bytes,bytes32,string)[] _depositData) returns()
func (_StakingManager *StakingManagerSession) BatchRegisterValidators0(_depositRoot [32]byte, _validatorId []*big.Int, _depositData []IStakingManagerDepositData) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchRegisterValidators0(&_StakingManager.TransactOpts, _depositRoot, _validatorId, _depositData)
}

// BatchRegisterValidators0 is a paid mutator transaction binding the contract method 0xb75a60d8.
//
// Solidity: function batchRegisterValidators(bytes32 _depositRoot, uint256[] _validatorId, (bytes,bytes,bytes32,string)[] _depositData) returns()
func (_StakingManager *StakingManagerTransactorSession) BatchRegisterValidators0(_depositRoot [32]byte, _validatorId []*big.Int, _depositData []IStakingManagerDepositData) (*types.Transaction, error) {
	return _StakingManager.Contract.BatchRegisterValidators0(&_StakingManager.TransactOpts, _depositRoot, _validatorId, _depositData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _auctionAddress, address _depositContractAddress) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, _auctionAddress common.Address, _depositContractAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", _auctionAddress, _depositContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _auctionAddress, address _depositContractAddress) returns()
func (_StakingManager *StakingManagerSession) Initialize(_auctionAddress common.Address, _depositContractAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, _auctionAddress, _depositContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _auctionAddress, address _depositContractAddress) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(_auctionAddress common.Address, _depositContractAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, _auctionAddress, _depositContractAddress)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0x4c73f498.
//
// Solidity: function initializeOnUpgrade(address _nodeOperatorManager, address _etherFiAdmin) returns()
func (_StakingManager *StakingManagerTransactor) InitializeOnUpgrade(opts *bind.TransactOpts, _nodeOperatorManager common.Address, _etherFiAdmin common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initializeOnUpgrade", _nodeOperatorManager, _etherFiAdmin)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0x4c73f498.
//
// Solidity: function initializeOnUpgrade(address _nodeOperatorManager, address _etherFiAdmin) returns()
func (_StakingManager *StakingManagerSession) InitializeOnUpgrade(_nodeOperatorManager common.Address, _etherFiAdmin common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.InitializeOnUpgrade(&_StakingManager.TransactOpts, _nodeOperatorManager, _etherFiAdmin)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0x4c73f498.
//
// Solidity: function initializeOnUpgrade(address _nodeOperatorManager, address _etherFiAdmin) returns()
func (_StakingManager *StakingManagerTransactorSession) InitializeOnUpgrade(_nodeOperatorManager common.Address, _etherFiAdmin common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.InitializeOnUpgrade(&_StakingManager.TransactOpts, _nodeOperatorManager, _etherFiAdmin)
}

// InstantiateEtherFiNode is a paid mutator transaction binding the contract method 0xaeeb9556.
//
// Solidity: function instantiateEtherFiNode(bool _createEigenPod) returns(address)
func (_StakingManager *StakingManagerTransactor) InstantiateEtherFiNode(opts *bind.TransactOpts, _createEigenPod bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "instantiateEtherFiNode", _createEigenPod)
}

// InstantiateEtherFiNode is a paid mutator transaction binding the contract method 0xaeeb9556.
//
// Solidity: function instantiateEtherFiNode(bool _createEigenPod) returns(address)
func (_StakingManager *StakingManagerSession) InstantiateEtherFiNode(_createEigenPod bool) (*types.Transaction, error) {
	return _StakingManager.Contract.InstantiateEtherFiNode(&_StakingManager.TransactOpts, _createEigenPod)
}

// InstantiateEtherFiNode is a paid mutator transaction binding the contract method 0xaeeb9556.
//
// Solidity: function instantiateEtherFiNode(bool _createEigenPod) returns(address)
func (_StakingManager *StakingManagerTransactorSession) InstantiateEtherFiNode(_createEigenPod bool) (*types.Transaction, error) {
	return _StakingManager.Contract.InstantiateEtherFiNode(&_StakingManager.TransactOpts, _createEigenPod)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_StakingManager *StakingManagerTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_StakingManager *StakingManagerSession) PauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.PauseContract(&_StakingManager.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_StakingManager *StakingManagerTransactorSession) PauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.PauseContract(&_StakingManager.TransactOpts)
}

// RegisterBNFTContract is a paid mutator transaction binding the contract method 0xac347303.
//
// Solidity: function registerBNFTContract(address _bnftAddress) returns()
func (_StakingManager *StakingManagerTransactor) RegisterBNFTContract(opts *bind.TransactOpts, _bnftAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerBNFTContract", _bnftAddress)
}

// RegisterBNFTContract is a paid mutator transaction binding the contract method 0xac347303.
//
// Solidity: function registerBNFTContract(address _bnftAddress) returns()
func (_StakingManager *StakingManagerSession) RegisterBNFTContract(_bnftAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterBNFTContract(&_StakingManager.TransactOpts, _bnftAddress)
}

// RegisterBNFTContract is a paid mutator transaction binding the contract method 0xac347303.
//
// Solidity: function registerBNFTContract(address _bnftAddress) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterBNFTContract(_bnftAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterBNFTContract(&_StakingManager.TransactOpts, _bnftAddress)
}

// RegisterEtherFiNodeImplementationContract is a paid mutator transaction binding the contract method 0xb0a6fbaa.
//
// Solidity: function registerEtherFiNodeImplementationContract(address _etherFiNodeImplementationContract) returns()
func (_StakingManager *StakingManagerTransactor) RegisterEtherFiNodeImplementationContract(opts *bind.TransactOpts, _etherFiNodeImplementationContract common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerEtherFiNodeImplementationContract", _etherFiNodeImplementationContract)
}

// RegisterEtherFiNodeImplementationContract is a paid mutator transaction binding the contract method 0xb0a6fbaa.
//
// Solidity: function registerEtherFiNodeImplementationContract(address _etherFiNodeImplementationContract) returns()
func (_StakingManager *StakingManagerSession) RegisterEtherFiNodeImplementationContract(_etherFiNodeImplementationContract common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterEtherFiNodeImplementationContract(&_StakingManager.TransactOpts, _etherFiNodeImplementationContract)
}

// RegisterEtherFiNodeImplementationContract is a paid mutator transaction binding the contract method 0xb0a6fbaa.
//
// Solidity: function registerEtherFiNodeImplementationContract(address _etherFiNodeImplementationContract) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterEtherFiNodeImplementationContract(_etherFiNodeImplementationContract common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterEtherFiNodeImplementationContract(&_StakingManager.TransactOpts, _etherFiNodeImplementationContract)
}

// RegisterTNFTContract is a paid mutator transaction binding the contract method 0x2ae1fc9c.
//
// Solidity: function registerTNFTContract(address _tnftAddress) returns()
func (_StakingManager *StakingManagerTransactor) RegisterTNFTContract(opts *bind.TransactOpts, _tnftAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerTNFTContract", _tnftAddress)
}

// RegisterTNFTContract is a paid mutator transaction binding the contract method 0x2ae1fc9c.
//
// Solidity: function registerTNFTContract(address _tnftAddress) returns()
func (_StakingManager *StakingManagerSession) RegisterTNFTContract(_tnftAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTNFTContract(&_StakingManager.TransactOpts, _tnftAddress)
}

// RegisterTNFTContract is a paid mutator transaction binding the contract method 0x2ae1fc9c.
//
// Solidity: function registerTNFTContract(address _tnftAddress) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterTNFTContract(_tnftAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTNFTContract(&_StakingManager.TransactOpts, _tnftAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// SetEtherFiNodesManagerAddress is a paid mutator transaction binding the contract method 0x9f6ad1f3.
//
// Solidity: function setEtherFiNodesManagerAddress(address _nodesManagerAddress) returns()
func (_StakingManager *StakingManagerTransactor) SetEtherFiNodesManagerAddress(opts *bind.TransactOpts, _nodesManagerAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setEtherFiNodesManagerAddress", _nodesManagerAddress)
}

// SetEtherFiNodesManagerAddress is a paid mutator transaction binding the contract method 0x9f6ad1f3.
//
// Solidity: function setEtherFiNodesManagerAddress(address _nodesManagerAddress) returns()
func (_StakingManager *StakingManagerSession) SetEtherFiNodesManagerAddress(_nodesManagerAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetEtherFiNodesManagerAddress(&_StakingManager.TransactOpts, _nodesManagerAddress)
}

// SetEtherFiNodesManagerAddress is a paid mutator transaction binding the contract method 0x9f6ad1f3.
//
// Solidity: function setEtherFiNodesManagerAddress(address _nodesManagerAddress) returns()
func (_StakingManager *StakingManagerTransactorSession) SetEtherFiNodesManagerAddress(_nodesManagerAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetEtherFiNodesManagerAddress(&_StakingManager.TransactOpts, _nodesManagerAddress)
}

// SetLiquidityPoolAddress is a paid mutator transaction binding the contract method 0x63986aba.
//
// Solidity: function setLiquidityPoolAddress(address _liquidityPoolAddress) returns()
func (_StakingManager *StakingManagerTransactor) SetLiquidityPoolAddress(opts *bind.TransactOpts, _liquidityPoolAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setLiquidityPoolAddress", _liquidityPoolAddress)
}

// SetLiquidityPoolAddress is a paid mutator transaction binding the contract method 0x63986aba.
//
// Solidity: function setLiquidityPoolAddress(address _liquidityPoolAddress) returns()
func (_StakingManager *StakingManagerSession) SetLiquidityPoolAddress(_liquidityPoolAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetLiquidityPoolAddress(&_StakingManager.TransactOpts, _liquidityPoolAddress)
}

// SetLiquidityPoolAddress is a paid mutator transaction binding the contract method 0x63986aba.
//
// Solidity: function setLiquidityPoolAddress(address _liquidityPoolAddress) returns()
func (_StakingManager *StakingManagerTransactorSession) SetLiquidityPoolAddress(_liquidityPoolAddress common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetLiquidityPoolAddress(&_StakingManager.TransactOpts, _liquidityPoolAddress)
}

// SetMaxBatchDepositSize is a paid mutator transaction binding the contract method 0xeefc4f25.
//
// Solidity: function setMaxBatchDepositSize(uint128 _newMaxBatchDepositSize) returns()
func (_StakingManager *StakingManagerTransactor) SetMaxBatchDepositSize(opts *bind.TransactOpts, _newMaxBatchDepositSize *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMaxBatchDepositSize", _newMaxBatchDepositSize)
}

// SetMaxBatchDepositSize is a paid mutator transaction binding the contract method 0xeefc4f25.
//
// Solidity: function setMaxBatchDepositSize(uint128 _newMaxBatchDepositSize) returns()
func (_StakingManager *StakingManagerSession) SetMaxBatchDepositSize(_newMaxBatchDepositSize *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaxBatchDepositSize(&_StakingManager.TransactOpts, _newMaxBatchDepositSize)
}

// SetMaxBatchDepositSize is a paid mutator transaction binding the contract method 0xeefc4f25.
//
// Solidity: function setMaxBatchDepositSize(uint128 _newMaxBatchDepositSize) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMaxBatchDepositSize(_newMaxBatchDepositSize *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaxBatchDepositSize(&_StakingManager.TransactOpts, _newMaxBatchDepositSize)
}

// SetNodeOperatorManager is a paid mutator transaction binding the contract method 0x6bd76ff0.
//
// Solidity: function setNodeOperatorManager(address _nodeOperateManager) returns()
func (_StakingManager *StakingManagerTransactor) SetNodeOperatorManager(opts *bind.TransactOpts, _nodeOperateManager common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setNodeOperatorManager", _nodeOperateManager)
}

// SetNodeOperatorManager is a paid mutator transaction binding the contract method 0x6bd76ff0.
//
// Solidity: function setNodeOperatorManager(address _nodeOperateManager) returns()
func (_StakingManager *StakingManagerSession) SetNodeOperatorManager(_nodeOperateManager common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetNodeOperatorManager(&_StakingManager.TransactOpts, _nodeOperateManager)
}

// SetNodeOperatorManager is a paid mutator transaction binding the contract method 0x6bd76ff0.
//
// Solidity: function setNodeOperatorManager(address _nodeOperateManager) returns()
func (_StakingManager *StakingManagerTransactorSession) SetNodeOperatorManager(_nodeOperateManager common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetNodeOperatorManager(&_StakingManager.TransactOpts, _nodeOperateManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_StakingManager *StakingManagerTransactor) UnPauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unPauseContract")
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_StakingManager *StakingManagerSession) UnPauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.UnPauseContract(&_StakingManager.TransactOpts)
}

// UnPauseContract is a paid mutator transaction binding the contract method 0xbac15203.
//
// Solidity: function unPauseContract() returns()
func (_StakingManager *StakingManagerTransactorSession) UnPauseContract() (*types.Transaction, error) {
	return _StakingManager.Contract.UnPauseContract(&_StakingManager.TransactOpts)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_StakingManager *StakingManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_StakingManager *StakingManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _StakingManager.Contract.UpdateAdmin(&_StakingManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_StakingManager *StakingManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _StakingManager.Contract.UpdateAdmin(&_StakingManager.TransactOpts, _address, _isAdmin)
}

// UpdateFullStakingStatus is a paid mutator transaction binding the contract method 0xb9e76c14.
//
// Solidity: function updateFullStakingStatus(bool _status) returns()
func (_StakingManager *StakingManagerTransactor) UpdateFullStakingStatus(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "updateFullStakingStatus", _status)
}

// UpdateFullStakingStatus is a paid mutator transaction binding the contract method 0xb9e76c14.
//
// Solidity: function updateFullStakingStatus(bool _status) returns()
func (_StakingManager *StakingManagerSession) UpdateFullStakingStatus(_status bool) (*types.Transaction, error) {
	return _StakingManager.Contract.UpdateFullStakingStatus(&_StakingManager.TransactOpts, _status)
}

// UpdateFullStakingStatus is a paid mutator transaction binding the contract method 0xb9e76c14.
//
// Solidity: function updateFullStakingStatus(bool _status) returns()
func (_StakingManager *StakingManagerTransactorSession) UpdateFullStakingStatus(_status bool) (*types.Transaction, error) {
	return _StakingManager.Contract.UpdateFullStakingStatus(&_StakingManager.TransactOpts, _status)
}

// UpgradeEtherFiNode is a paid mutator transaction binding the contract method 0x49370974.
//
// Solidity: function upgradeEtherFiNode(address _newImplementation) returns()
func (_StakingManager *StakingManagerTransactor) UpgradeEtherFiNode(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeEtherFiNode", _newImplementation)
}

// UpgradeEtherFiNode is a paid mutator transaction binding the contract method 0x49370974.
//
// Solidity: function upgradeEtherFiNode(address _newImplementation) returns()
func (_StakingManager *StakingManagerSession) UpgradeEtherFiNode(_newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeEtherFiNode(&_StakingManager.TransactOpts, _newImplementation)
}

// UpgradeEtherFiNode is a paid mutator transaction binding the contract method 0x49370974.
//
// Solidity: function upgradeEtherFiNode(address _newImplementation) returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeEtherFiNode(_newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeEtherFiNode(&_StakingManager.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingManager *StakingManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingManager *StakingManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeTo(&_StakingManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeTo(&_StakingManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeToAndCall(&_StakingManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StakingManager *StakingManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.UpgradeToAndCall(&_StakingManager.TransactOpts, newImplementation, data)
}

// StakingManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StakingManager contract.
type StakingManagerAdminChangedIterator struct {
	Event *StakingManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAdminChanged)
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
		it.Event = new(StakingManagerAdminChanged)
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
func (it *StakingManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAdminChanged represents a AdminChanged event raised by the StakingManager contract.
type StakingManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingManager *StakingManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StakingManagerAdminChangedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAdminChangedIterator{contract: _StakingManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingManager *StakingManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAdminChanged)
				if err := _StakingManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StakingManager *StakingManagerFilterer) ParseAdminChanged(log types.Log) (*StakingManagerAdminChanged, error) {
	event := new(StakingManagerAdminChanged)
	if err := _StakingManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the StakingManager contract.
type StakingManagerBeaconUpgradedIterator struct {
	Event *StakingManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerBeaconUpgraded)
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
		it.Event = new(StakingManagerBeaconUpgraded)
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
func (it *StakingManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerBeaconUpgraded represents a BeaconUpgraded event raised by the StakingManager contract.
type StakingManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingManager *StakingManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StakingManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerBeaconUpgradedIterator{contract: _StakingManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingManager *StakingManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StakingManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerBeaconUpgraded)
				if err := _StakingManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StakingManager *StakingManagerFilterer) ParseBeaconUpgraded(log types.Log) (*StakingManagerBeaconUpgraded, error) {
	event := new(StakingManagerBeaconUpgraded)
	if err := _StakingManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerDepositCancelledIterator is returned from FilterDepositCancelled and is used to iterate over the raw logs and unpacked data for DepositCancelled events raised by the StakingManager contract.
type StakingManagerDepositCancelledIterator struct {
	Event *StakingManagerDepositCancelled // Event containing the contract specifics and raw log

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
func (it *StakingManagerDepositCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerDepositCancelled)
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
		it.Event = new(StakingManagerDepositCancelled)
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
func (it *StakingManagerDepositCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerDepositCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerDepositCancelled represents a DepositCancelled event raised by the StakingManager contract.
type StakingManagerDepositCancelled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositCancelled is a free log retrieval operation binding the contract event 0xc67e37700e70a47022e92eff0449cbf934acde8733b7dc953d826e156db8e6ce.
//
// Solidity: event DepositCancelled(uint256 id)
func (_StakingManager *StakingManagerFilterer) FilterDepositCancelled(opts *bind.FilterOpts) (*StakingManagerDepositCancelledIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "DepositCancelled")
	if err != nil {
		return nil, err
	}
	return &StakingManagerDepositCancelledIterator{contract: _StakingManager.contract, event: "DepositCancelled", logs: logs, sub: sub}, nil
}

// WatchDepositCancelled is a free log subscription operation binding the contract event 0xc67e37700e70a47022e92eff0449cbf934acde8733b7dc953d826e156db8e6ce.
//
// Solidity: event DepositCancelled(uint256 id)
func (_StakingManager *StakingManagerFilterer) WatchDepositCancelled(opts *bind.WatchOpts, sink chan<- *StakingManagerDepositCancelled) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "DepositCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerDepositCancelled)
				if err := _StakingManager.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
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

// ParseDepositCancelled is a log parse operation binding the contract event 0xc67e37700e70a47022e92eff0449cbf934acde8733b7dc953d826e156db8e6ce.
//
// Solidity: event DepositCancelled(uint256 id)
func (_StakingManager *StakingManagerFilterer) ParseDepositCancelled(log types.Log) (*StakingManagerDepositCancelled, error) {
	event := new(StakingManagerDepositCancelled)
	if err := _StakingManager.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingManager contract.
type StakingManagerInitializedIterator struct {
	Event *StakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialized)
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
		it.Event = new(StakingManagerInitialized)
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
func (it *StakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialized represents a Initialized event raised by the StakingManager contract.
type StakingManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingManager *StakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingManagerInitializedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitializedIterator{contract: _StakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakingManager *StakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialized)
				if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseInitialized(log types.Log) (*StakingManagerInitialized, error) {
	event := new(StakingManagerInitialized)
	if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingManager contract.
type StakingManagerPausedIterator struct {
	Event *StakingManagerPaused // Event containing the contract specifics and raw log

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
func (it *StakingManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerPaused)
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
		it.Event = new(StakingManagerPaused)
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
func (it *StakingManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerPaused represents a Paused event raised by the StakingManager contract.
type StakingManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingManagerPausedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingManagerPausedIterator{contract: _StakingManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingManagerPaused) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerPaused)
				if err := _StakingManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParsePaused(log types.Log) (*StakingManagerPaused, error) {
	event := new(StakingManagerPaused)
	if err := _StakingManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeDepositIterator is returned from FilterStakeDeposit and is used to iterate over the raw logs and unpacked data for StakeDeposit events raised by the StakingManager contract.
type StakingManagerStakeDepositIterator struct {
	Event *StakingManagerStakeDeposit // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeDeposit)
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
		it.Event = new(StakingManagerStakeDeposit)
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
func (it *StakingManagerStakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeDeposit represents a StakeDeposit event raised by the StakingManager contract.
type StakingManagerStakeDeposit struct {
	Staker       common.Address
	BidId        *big.Int
	WithdrawSafe common.Address
	Restaked     bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposit is a free log retrieval operation binding the contract event 0x71055523e548abbe2e8f8cf6ae548d26ec4e36ea7d62caecfe259670393e27a7.
//
// Solidity: event StakeDeposit(address indexed staker, uint256 indexed bidId, address indexed withdrawSafe, bool restaked)
func (_StakingManager *StakingManagerFilterer) FilterStakeDeposit(opts *bind.FilterOpts, staker []common.Address, bidId []*big.Int, withdrawSafe []common.Address) (*StakingManagerStakeDepositIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}
	var withdrawSafeRule []interface{}
	for _, withdrawSafeItem := range withdrawSafe {
		withdrawSafeRule = append(withdrawSafeRule, withdrawSafeItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeDeposit", stakerRule, bidIdRule, withdrawSafeRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeDepositIterator{contract: _StakingManager.contract, event: "StakeDeposit", logs: logs, sub: sub}, nil
}

// WatchStakeDeposit is a free log subscription operation binding the contract event 0x71055523e548abbe2e8f8cf6ae548d26ec4e36ea7d62caecfe259670393e27a7.
//
// Solidity: event StakeDeposit(address indexed staker, uint256 indexed bidId, address indexed withdrawSafe, bool restaked)
func (_StakingManager *StakingManagerFilterer) WatchStakeDeposit(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeDeposit, staker []common.Address, bidId []*big.Int, withdrawSafe []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}
	var withdrawSafeRule []interface{}
	for _, withdrawSafeItem := range withdrawSafe {
		withdrawSafeRule = append(withdrawSafeRule, withdrawSafeItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeDeposit", stakerRule, bidIdRule, withdrawSafeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeDeposit)
				if err := _StakingManager.contract.UnpackLog(event, "StakeDeposit", log); err != nil {
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

// ParseStakeDeposit is a log parse operation binding the contract event 0x71055523e548abbe2e8f8cf6ae548d26ec4e36ea7d62caecfe259670393e27a7.
//
// Solidity: event StakeDeposit(address indexed staker, uint256 indexed bidId, address indexed withdrawSafe, bool restaked)
func (_StakingManager *StakingManagerFilterer) ParseStakeDeposit(log types.Log) (*StakingManagerStakeDeposit, error) {
	event := new(StakingManagerStakeDeposit)
	if err := _StakingManager.contract.UnpackLog(event, "StakeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeSourceIterator is returned from FilterStakeSource and is used to iterate over the raw logs and unpacked data for StakeSource events raised by the StakingManager contract.
type StakingManagerStakeSourceIterator struct {
	Event *StakingManagerStakeSource // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeSourceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeSource)
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
		it.Event = new(StakingManagerStakeSource)
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
func (it *StakingManagerStakeSourceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeSourceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeSource represents a StakeSource event raised by the StakingManager contract.
type StakingManagerStakeSource struct {
	BidId  *big.Int
	Source uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeSource is a free log retrieval operation binding the contract event 0xbea568068b14765232e3c45d3dca5ecb83824011eaf1950e32719fd82aefb108.
//
// Solidity: event StakeSource(uint256 bidId, uint8 source)
func (_StakingManager *StakingManagerFilterer) FilterStakeSource(opts *bind.FilterOpts) (*StakingManagerStakeSourceIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeSource")
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeSourceIterator{contract: _StakingManager.contract, event: "StakeSource", logs: logs, sub: sub}, nil
}

// WatchStakeSource is a free log subscription operation binding the contract event 0xbea568068b14765232e3c45d3dca5ecb83824011eaf1950e32719fd82aefb108.
//
// Solidity: event StakeSource(uint256 bidId, uint8 source)
func (_StakingManager *StakingManagerFilterer) WatchStakeSource(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeSource) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeSource")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeSource)
				if err := _StakingManager.contract.UnpackLog(event, "StakeSource", log); err != nil {
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

// ParseStakeSource is a log parse operation binding the contract event 0xbea568068b14765232e3c45d3dca5ecb83824011eaf1950e32719fd82aefb108.
//
// Solidity: event StakeSource(uint256 bidId, uint8 source)
func (_StakingManager *StakingManagerFilterer) ParseStakeSource(log types.Log) (*StakingManagerStakeSource, error) {
	event := new(StakingManagerStakeSource)
	if err := _StakingManager.contract.UnpackLog(event, "StakeSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingManager contract.
type StakingManagerUnpausedIterator struct {
	Event *StakingManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnpaused)
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
		it.Event = new(StakingManagerUnpaused)
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
func (it *StakingManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnpaused represents a Unpaused event raised by the StakingManager contract.
type StakingManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingManagerUnpausedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnpausedIterator{contract: _StakingManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnpaused)
				if err := _StakingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseUnpaused(log types.Log) (*StakingManagerUnpaused, error) {
	event := new(StakingManagerUnpaused)
	if err := _StakingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StakingManager contract.
type StakingManagerUpgradedIterator struct {
	Event *StakingManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *StakingManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUpgraded)
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
		it.Event = new(StakingManagerUpgraded)
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
func (it *StakingManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUpgraded represents a Upgraded event raised by the StakingManager contract.
type StakingManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StakingManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUpgradedIterator{contract: _StakingManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StakingManager *StakingManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StakingManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUpgraded)
				if err := _StakingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseUpgraded(log types.Log) (*StakingManagerUpgraded, error) {
	event := new(StakingManagerUpgraded)
	if err := _StakingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the StakingManager contract.
type StakingManagerValidatorRegisteredIterator struct {
	Event *StakingManagerValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *StakingManagerValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerValidatorRegistered)
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
		it.Event = new(StakingManagerValidatorRegistered)
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
func (it *StakingManagerValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerValidatorRegistered represents a ValidatorRegistered event raised by the StakingManager contract.
type StakingManagerValidatorRegistered struct {
	Operator                         common.Address
	BNftOwner                        common.Address
	TNftOwner                        common.Address
	ValidatorId                      *big.Int
	ValidatorPubKey                  []byte
	IpfsHashForEncryptedValidatorKey string
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x0b43d988cd5ab75ae318de41d6871d4b26efe57c3f3975331873f4dc073041fc.
//
// Solidity: event ValidatorRegistered(address indexed operator, address indexed bNftOwner, address indexed tNftOwner, uint256 validatorId, bytes validatorPubKey, string ipfsHashForEncryptedValidatorKey)
func (_StakingManager *StakingManagerFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, operator []common.Address, bNftOwner []common.Address, tNftOwner []common.Address) (*StakingManagerValidatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var bNftOwnerRule []interface{}
	for _, bNftOwnerItem := range bNftOwner {
		bNftOwnerRule = append(bNftOwnerRule, bNftOwnerItem)
	}
	var tNftOwnerRule []interface{}
	for _, tNftOwnerItem := range tNftOwner {
		tNftOwnerRule = append(tNftOwnerRule, tNftOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ValidatorRegistered", operatorRule, bNftOwnerRule, tNftOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerValidatorRegisteredIterator{contract: _StakingManager.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x0b43d988cd5ab75ae318de41d6871d4b26efe57c3f3975331873f4dc073041fc.
//
// Solidity: event ValidatorRegistered(address indexed operator, address indexed bNftOwner, address indexed tNftOwner, uint256 validatorId, bytes validatorPubKey, string ipfsHashForEncryptedValidatorKey)
func (_StakingManager *StakingManagerFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *StakingManagerValidatorRegistered, operator []common.Address, bNftOwner []common.Address, tNftOwner []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var bNftOwnerRule []interface{}
	for _, bNftOwnerItem := range bNftOwner {
		bNftOwnerRule = append(bNftOwnerRule, bNftOwnerItem)
	}
	var tNftOwnerRule []interface{}
	for _, tNftOwnerItem := range tNftOwner {
		tNftOwnerRule = append(tNftOwnerRule, tNftOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ValidatorRegistered", operatorRule, bNftOwnerRule, tNftOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerValidatorRegistered)
				if err := _StakingManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0x0b43d988cd5ab75ae318de41d6871d4b26efe57c3f3975331873f4dc073041fc.
//
// Solidity: event ValidatorRegistered(address indexed operator, address indexed bNftOwner, address indexed tNftOwner, uint256 validatorId, bytes validatorPubKey, string ipfsHashForEncryptedValidatorKey)
func (_StakingManager *StakingManagerFilterer) ParseValidatorRegistered(log types.Log) (*StakingManagerValidatorRegistered, error) {
	event := new(StakingManagerValidatorRegistered)
	if err := _StakingManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
