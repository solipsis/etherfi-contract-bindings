// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

// TODO: this is wrong
package events

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

// MembershipNFTMetaData contains all meta data concerning the MembershipNFT contract.
var MembershipNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DisallowZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEAPRollover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintingIsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMembershipManagerContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequireTokenUnlocked\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MerkleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"MintingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"TokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPRECATED_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"accruedLoyaltyPointsOf\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"accruedStakingRewardsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"accruedTierPointsOf\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endID\",\"type\":\"uint256\"}],\"name\":\"alertBatchMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"alertMetadataUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"allTimeHighDepositOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOfUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_amountForPoints\",\"type\":\"uint128\"}],\"name\":\"canTopUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimableTier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eapDepositBlockNumber\",\"type\":\"uint32\"}],\"name\":\"computeTierPointsForEap\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"eapDepositProcessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eapMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_blocks\",\"type\":\"uint32\"}],\"name\":\"incrementLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_membershipManagerInstance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityPoolAddress\",\"type\":\"address\"}],\"name\":\"initializeOnUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityPool\",\"outputs\":[{\"internalType\":\"contractILiquidityPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"loyaltyPointsOf\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_since\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_until\",\"type\":\"uint256\"}],\"name\":\"membershipPointsEarning\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMintTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftData\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"transferLockedUntil\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eapDepositBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_snapshotEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_points\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"processDepositFromEapUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requiredEapPointsPerEapDeposit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"}],\"name\":\"setContractMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maxTokenId\",\"type\":\"uint32\"}],\"name\":\"setMaxTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newURI\",\"type\":\"string\"}],\"name\":\"setMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setMintingPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64[]\",\"name\":\"_requiredEapPointsPerEapDeposit\",\"type\":\"uint64[]\"}],\"name\":\"setUpForEap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tierOf\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tierPointsOf\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferLockedUntil\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"valueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MembershipNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use MembershipNFTMetaData.ABI instead.
var MembershipNFTABI = MembershipNFTMetaData.ABI

// MembershipNFT is an auto generated Go binding around an Ethereum contract.
type MembershipNFT struct {
	MembershipNFTCaller     // Read-only binding to the contract
	MembershipNFTTransactor // Write-only binding to the contract
	MembershipNFTFilterer   // Log filterer for contract events
}

// MembershipNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MembershipNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MembershipNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MembershipNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MembershipNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MembershipNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MembershipNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MembershipNFTSession struct {
	Contract     *MembershipNFT    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MembershipNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MembershipNFTCallerSession struct {
	Contract *MembershipNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MembershipNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MembershipNFTTransactorSession struct {
	Contract     *MembershipNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MembershipNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MembershipNFTRaw struct {
	Contract *MembershipNFT // Generic contract binding to access the raw methods on
}

// MembershipNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MembershipNFTCallerRaw struct {
	Contract *MembershipNFTCaller // Generic read-only contract binding to access the raw methods on
}

// MembershipNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MembershipNFTTransactorRaw struct {
	Contract *MembershipNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMembershipNFT creates a new instance of MembershipNFT, bound to a specific deployed contract.
func NewMembershipNFT(address common.Address, backend bind.ContractBackend) (*MembershipNFT, error) {
	contract, err := bindMembershipNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MembershipNFT{MembershipNFTCaller: MembershipNFTCaller{contract: contract}, MembershipNFTTransactor: MembershipNFTTransactor{contract: contract}, MembershipNFTFilterer: MembershipNFTFilterer{contract: contract}}, nil
}

// NewMembershipNFTCaller creates a new read-only instance of MembershipNFT, bound to a specific deployed contract.
func NewMembershipNFTCaller(address common.Address, caller bind.ContractCaller) (*MembershipNFTCaller, error) {
	contract, err := bindMembershipNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTCaller{contract: contract}, nil
}

// NewMembershipNFTTransactor creates a new write-only instance of MembershipNFT, bound to a specific deployed contract.
func NewMembershipNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*MembershipNFTTransactor, error) {
	contract, err := bindMembershipNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTTransactor{contract: contract}, nil
}

// NewMembershipNFTFilterer creates a new log filterer instance of MembershipNFT, bound to a specific deployed contract.
func NewMembershipNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*MembershipNFTFilterer, error) {
	contract, err := bindMembershipNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTFilterer{contract: contract}, nil
}

// bindMembershipNFT binds a generic wrapper to an already deployed contract.
func bindMembershipNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MembershipNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MembershipNFT *MembershipNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MembershipNFT.Contract.MembershipNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MembershipNFT *MembershipNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MembershipNFT.Contract.MembershipNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MembershipNFT *MembershipNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MembershipNFT.Contract.MembershipNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MembershipNFT *MembershipNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MembershipNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MembershipNFT *MembershipNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MembershipNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MembershipNFT *MembershipNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MembershipNFT.Contract.contract.Transact(opts, method, params...)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_MembershipNFT *MembershipNFTCaller) DEPRECATEDAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "DEPRECATED_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_MembershipNFT *MembershipNFTSession) DEPRECATEDAdmin() (common.Address, error) {
	return _MembershipNFT.Contract.DEPRECATEDAdmin(&_MembershipNFT.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0x50a8a553.
//
// Solidity: function DEPRECATED_admin() view returns(address)
func (_MembershipNFT *MembershipNFTCallerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _MembershipNFT.Contract.DEPRECATEDAdmin(&_MembershipNFT.CallOpts)
}

// AccruedLoyaltyPointsOf is a free data retrieval call binding the contract method 0x29b69ad1.
//
// Solidity: function accruedLoyaltyPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCaller) AccruedLoyaltyPointsOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "accruedLoyaltyPointsOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedLoyaltyPointsOf is a free data retrieval call binding the contract method 0x29b69ad1.
//
// Solidity: function accruedLoyaltyPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTSession) AccruedLoyaltyPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AccruedLoyaltyPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AccruedLoyaltyPointsOf is a free data retrieval call binding the contract method 0x29b69ad1.
//
// Solidity: function accruedLoyaltyPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCallerSession) AccruedLoyaltyPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AccruedLoyaltyPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AccruedStakingRewardsOf is a free data retrieval call binding the contract method 0x335a0e94.
//
// Solidity: function accruedStakingRewardsOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTCaller) AccruedStakingRewardsOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "accruedStakingRewardsOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStakingRewardsOf is a free data retrieval call binding the contract method 0x335a0e94.
//
// Solidity: function accruedStakingRewardsOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTSession) AccruedStakingRewardsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AccruedStakingRewardsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AccruedStakingRewardsOf is a free data retrieval call binding the contract method 0x335a0e94.
//
// Solidity: function accruedStakingRewardsOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTCallerSession) AccruedStakingRewardsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AccruedStakingRewardsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AccruedTierPointsOf is a free data retrieval call binding the contract method 0x5518787c.
//
// Solidity: function accruedTierPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCaller) AccruedTierPointsOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "accruedTierPointsOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedTierPointsOf is a free data retrieval call binding the contract method 0x5518787c.
//
// Solidity: function accruedTierPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTSession) AccruedTierPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AccruedTierPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AccruedTierPointsOf is a free data retrieval call binding the contract method 0x5518787c.
//
// Solidity: function accruedTierPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCallerSession) AccruedTierPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AccruedTierPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_MembershipNFT *MembershipNFTSession) Admins(arg0 common.Address) (bool, error) {
	return _MembershipNFT.Contract.Admins(&_MembershipNFT.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _MembershipNFT.Contract.Admins(&_MembershipNFT.CallOpts, arg0)
}

// AllTimeHighDepositOf is a free data retrieval call binding the contract method 0x1c187780.
//
// Solidity: function allTimeHighDepositOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTCaller) AllTimeHighDepositOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "allTimeHighDepositOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllTimeHighDepositOf is a free data retrieval call binding the contract method 0x1c187780.
//
// Solidity: function allTimeHighDepositOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTSession) AllTimeHighDepositOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AllTimeHighDepositOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AllTimeHighDepositOf is a free data retrieval call binding the contract method 0x1c187780.
//
// Solidity: function allTimeHighDepositOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTCallerSession) AllTimeHighDepositOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.AllTimeHighDepositOf(&_MembershipNFT.CallOpts, _tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MembershipNFT *MembershipNFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MembershipNFT *MembershipNFTSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.BalanceOf(&_MembershipNFT.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_MembershipNFT *MembershipNFTCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.BalanceOf(&_MembershipNFT.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MembershipNFT *MembershipNFTCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MembershipNFT *MembershipNFTSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MembershipNFT.Contract.BalanceOfBatch(&_MembershipNFT.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_MembershipNFT *MembershipNFTCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _MembershipNFT.Contract.BalanceOfBatch(&_MembershipNFT.CallOpts, accounts, ids)
}

// BalanceOfUser is a free data retrieval call binding the contract method 0xf43cc3b3.
//
// Solidity: function balanceOfUser(address _user, uint256 _id) view returns(uint256)
func (_MembershipNFT *MembershipNFTCaller) BalanceOfUser(opts *bind.CallOpts, _user common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "balanceOfUser", _user, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUser is a free data retrieval call binding the contract method 0xf43cc3b3.
//
// Solidity: function balanceOfUser(address _user, uint256 _id) view returns(uint256)
func (_MembershipNFT *MembershipNFTSession) BalanceOfUser(_user common.Address, _id *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.BalanceOfUser(&_MembershipNFT.CallOpts, _user, _id)
}

// BalanceOfUser is a free data retrieval call binding the contract method 0xf43cc3b3.
//
// Solidity: function balanceOfUser(address _user, uint256 _id) view returns(uint256)
func (_MembershipNFT *MembershipNFTCallerSession) BalanceOfUser(_user common.Address, _id *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.BalanceOfUser(&_MembershipNFT.CallOpts, _user, _id)
}

// CanTopUp is a free data retrieval call binding the contract method 0x5707a765.
//
// Solidity: function canTopUp(uint256 _tokenId, uint256 _totalAmount, uint128 _amount, uint128 _amountForPoints) view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) CanTopUp(opts *bind.CallOpts, _tokenId *big.Int, _totalAmount *big.Int, _amount *big.Int, _amountForPoints *big.Int) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "canTopUp", _tokenId, _totalAmount, _amount, _amountForPoints)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTopUp is a free data retrieval call binding the contract method 0x5707a765.
//
// Solidity: function canTopUp(uint256 _tokenId, uint256 _totalAmount, uint128 _amount, uint128 _amountForPoints) view returns(bool)
func (_MembershipNFT *MembershipNFTSession) CanTopUp(_tokenId *big.Int, _totalAmount *big.Int, _amount *big.Int, _amountForPoints *big.Int) (bool, error) {
	return _MembershipNFT.Contract.CanTopUp(&_MembershipNFT.CallOpts, _tokenId, _totalAmount, _amount, _amountForPoints)
}

// CanTopUp is a free data retrieval call binding the contract method 0x5707a765.
//
// Solidity: function canTopUp(uint256 _tokenId, uint256 _totalAmount, uint128 _amount, uint128 _amountForPoints) view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) CanTopUp(_tokenId *big.Int, _totalAmount *big.Int, _amount *big.Int, _amountForPoints *big.Int) (bool, error) {
	return _MembershipNFT.Contract.CanTopUp(&_MembershipNFT.CallOpts, _tokenId, _totalAmount, _amount, _amountForPoints)
}

// ClaimableTier is a free data retrieval call binding the contract method 0xd774babe.
//
// Solidity: function claimableTier(uint256 _tokenId) view returns(uint8)
func (_MembershipNFT *MembershipNFTCaller) ClaimableTier(opts *bind.CallOpts, _tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "claimableTier", _tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ClaimableTier is a free data retrieval call binding the contract method 0xd774babe.
//
// Solidity: function claimableTier(uint256 _tokenId) view returns(uint8)
func (_MembershipNFT *MembershipNFTSession) ClaimableTier(_tokenId *big.Int) (uint8, error) {
	return _MembershipNFT.Contract.ClaimableTier(&_MembershipNFT.CallOpts, _tokenId)
}

// ClaimableTier is a free data retrieval call binding the contract method 0xd774babe.
//
// Solidity: function claimableTier(uint256 _tokenId) view returns(uint8)
func (_MembershipNFT *MembershipNFTCallerSession) ClaimableTier(_tokenId *big.Int) (uint8, error) {
	return _MembershipNFT.Contract.ClaimableTier(&_MembershipNFT.CallOpts, _tokenId)
}

// ComputeTierPointsForEap is a free data retrieval call binding the contract method 0x2d3d21bb.
//
// Solidity: function computeTierPointsForEap(uint32 _eapDepositBlockNumber) view returns(uint40)
func (_MembershipNFT *MembershipNFTCaller) ComputeTierPointsForEap(opts *bind.CallOpts, _eapDepositBlockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "computeTierPointsForEap", _eapDepositBlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeTierPointsForEap is a free data retrieval call binding the contract method 0x2d3d21bb.
//
// Solidity: function computeTierPointsForEap(uint32 _eapDepositBlockNumber) view returns(uint40)
func (_MembershipNFT *MembershipNFTSession) ComputeTierPointsForEap(_eapDepositBlockNumber uint32) (*big.Int, error) {
	return _MembershipNFT.Contract.ComputeTierPointsForEap(&_MembershipNFT.CallOpts, _eapDepositBlockNumber)
}

// ComputeTierPointsForEap is a free data retrieval call binding the contract method 0x2d3d21bb.
//
// Solidity: function computeTierPointsForEap(uint32 _eapDepositBlockNumber) view returns(uint40)
func (_MembershipNFT *MembershipNFTCallerSession) ComputeTierPointsForEap(_eapDepositBlockNumber uint32) (*big.Int, error) {
	return _MembershipNFT.Contract.ComputeTierPointsForEap(&_MembershipNFT.CallOpts, _eapDepositBlockNumber)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_MembershipNFT *MembershipNFTCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_MembershipNFT *MembershipNFTSession) ContractURI() (string, error) {
	return _MembershipNFT.Contract.ContractURI(&_MembershipNFT.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_MembershipNFT *MembershipNFTCallerSession) ContractURI() (string, error) {
	return _MembershipNFT.Contract.ContractURI(&_MembershipNFT.CallOpts)
}

// EapDepositProcessed is a free data retrieval call binding the contract method 0x27570411.
//
// Solidity: function eapDepositProcessed(address ) view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) EapDepositProcessed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "eapDepositProcessed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EapDepositProcessed is a free data retrieval call binding the contract method 0x27570411.
//
// Solidity: function eapDepositProcessed(address ) view returns(bool)
func (_MembershipNFT *MembershipNFTSession) EapDepositProcessed(arg0 common.Address) (bool, error) {
	return _MembershipNFT.Contract.EapDepositProcessed(&_MembershipNFT.CallOpts, arg0)
}

// EapDepositProcessed is a free data retrieval call binding the contract method 0x27570411.
//
// Solidity: function eapDepositProcessed(address ) view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) EapDepositProcessed(arg0 common.Address) (bool, error) {
	return _MembershipNFT.Contract.EapDepositProcessed(&_MembershipNFT.CallOpts, arg0)
}

// EapMerkleRoot is a free data retrieval call binding the contract method 0x7e328ae4.
//
// Solidity: function eapMerkleRoot() view returns(bytes32)
func (_MembershipNFT *MembershipNFTCaller) EapMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "eapMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EapMerkleRoot is a free data retrieval call binding the contract method 0x7e328ae4.
//
// Solidity: function eapMerkleRoot() view returns(bytes32)
func (_MembershipNFT *MembershipNFTSession) EapMerkleRoot() ([32]byte, error) {
	return _MembershipNFT.Contract.EapMerkleRoot(&_MembershipNFT.CallOpts)
}

// EapMerkleRoot is a free data retrieval call binding the contract method 0x7e328ae4.
//
// Solidity: function eapMerkleRoot() view returns(bytes32)
func (_MembershipNFT *MembershipNFTCallerSession) EapMerkleRoot() ([32]byte, error) {
	return _MembershipNFT.Contract.EapMerkleRoot(&_MembershipNFT.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_MembershipNFT *MembershipNFTCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_MembershipNFT *MembershipNFTSession) GetImplementation() (common.Address, error) {
	return _MembershipNFT.Contract.GetImplementation(&_MembershipNFT.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_MembershipNFT *MembershipNFTCallerSession) GetImplementation() (common.Address, error) {
	return _MembershipNFT.Contract.GetImplementation(&_MembershipNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MembershipNFT *MembershipNFTSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MembershipNFT.Contract.IsApprovedForAll(&_MembershipNFT.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _MembershipNFT.Contract.IsApprovedForAll(&_MembershipNFT.CallOpts, account, operator)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x3ca5444b.
//
// Solidity: function isWithdrawable(uint256 _tokenId, uint256 _withdrawalAmount) view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) IsWithdrawable(opts *bind.CallOpts, _tokenId *big.Int, _withdrawalAmount *big.Int) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "isWithdrawable", _tokenId, _withdrawalAmount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawable is a free data retrieval call binding the contract method 0x3ca5444b.
//
// Solidity: function isWithdrawable(uint256 _tokenId, uint256 _withdrawalAmount) view returns(bool)
func (_MembershipNFT *MembershipNFTSession) IsWithdrawable(_tokenId *big.Int, _withdrawalAmount *big.Int) (bool, error) {
	return _MembershipNFT.Contract.IsWithdrawable(&_MembershipNFT.CallOpts, _tokenId, _withdrawalAmount)
}

// IsWithdrawable is a free data retrieval call binding the contract method 0x3ca5444b.
//
// Solidity: function isWithdrawable(uint256 _tokenId, uint256 _withdrawalAmount) view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) IsWithdrawable(_tokenId *big.Int, _withdrawalAmount *big.Int) (bool, error) {
	return _MembershipNFT.Contract.IsWithdrawable(&_MembershipNFT.CallOpts, _tokenId, _withdrawalAmount)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_MembershipNFT *MembershipNFTCaller) LiquidityPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "liquidityPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_MembershipNFT *MembershipNFTSession) LiquidityPool() (common.Address, error) {
	return _MembershipNFT.Contract.LiquidityPool(&_MembershipNFT.CallOpts)
}

// LiquidityPool is a free data retrieval call binding the contract method 0x665a11ca.
//
// Solidity: function liquidityPool() view returns(address)
func (_MembershipNFT *MembershipNFTCallerSession) LiquidityPool() (common.Address, error) {
	return _MembershipNFT.Contract.LiquidityPool(&_MembershipNFT.CallOpts)
}

// LoyaltyPointsOf is a free data retrieval call binding the contract method 0xe0879a33.
//
// Solidity: function loyaltyPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCaller) LoyaltyPointsOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "loyaltyPointsOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LoyaltyPointsOf is a free data retrieval call binding the contract method 0xe0879a33.
//
// Solidity: function loyaltyPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTSession) LoyaltyPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.LoyaltyPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// LoyaltyPointsOf is a free data retrieval call binding the contract method 0xe0879a33.
//
// Solidity: function loyaltyPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCallerSession) LoyaltyPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.LoyaltyPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// MaxTokenId is a free data retrieval call binding the contract method 0x91ba317a.
//
// Solidity: function maxTokenId() view returns(uint32)
func (_MembershipNFT *MembershipNFTCaller) MaxTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "maxTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxTokenId is a free data retrieval call binding the contract method 0x91ba317a.
//
// Solidity: function maxTokenId() view returns(uint32)
func (_MembershipNFT *MembershipNFTSession) MaxTokenId() (uint32, error) {
	return _MembershipNFT.Contract.MaxTokenId(&_MembershipNFT.CallOpts)
}

// MaxTokenId is a free data retrieval call binding the contract method 0x91ba317a.
//
// Solidity: function maxTokenId() view returns(uint32)
func (_MembershipNFT *MembershipNFTCallerSession) MaxTokenId() (uint32, error) {
	return _MembershipNFT.Contract.MaxTokenId(&_MembershipNFT.CallOpts)
}

// MembershipPointsEarning is a free data retrieval call binding the contract method 0xfdb65879.
//
// Solidity: function membershipPointsEarning(uint256 _tokenId, uint256 _since, uint256 _until) view returns(uint40)
func (_MembershipNFT *MembershipNFTCaller) MembershipPointsEarning(opts *bind.CallOpts, _tokenId *big.Int, _since *big.Int, _until *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "membershipPointsEarning", _tokenId, _since, _until)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MembershipPointsEarning is a free data retrieval call binding the contract method 0xfdb65879.
//
// Solidity: function membershipPointsEarning(uint256 _tokenId, uint256 _since, uint256 _until) view returns(uint40)
func (_MembershipNFT *MembershipNFTSession) MembershipPointsEarning(_tokenId *big.Int, _since *big.Int, _until *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.MembershipPointsEarning(&_MembershipNFT.CallOpts, _tokenId, _since, _until)
}

// MembershipPointsEarning is a free data retrieval call binding the contract method 0xfdb65879.
//
// Solidity: function membershipPointsEarning(uint256 _tokenId, uint256 _since, uint256 _until) view returns(uint40)
func (_MembershipNFT *MembershipNFTCallerSession) MembershipPointsEarning(_tokenId *big.Int, _since *big.Int, _until *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.MembershipPointsEarning(&_MembershipNFT.CallOpts, _tokenId, _since, _until)
}

// MintingPaused is a free data retrieval call binding the contract method 0xe1a283d6.
//
// Solidity: function mintingPaused() view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) MintingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "mintingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintingPaused is a free data retrieval call binding the contract method 0xe1a283d6.
//
// Solidity: function mintingPaused() view returns(bool)
func (_MembershipNFT *MembershipNFTSession) MintingPaused() (bool, error) {
	return _MembershipNFT.Contract.MintingPaused(&_MembershipNFT.CallOpts)
}

// MintingPaused is a free data retrieval call binding the contract method 0xe1a283d6.
//
// Solidity: function mintingPaused() view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) MintingPaused() (bool, error) {
	return _MembershipNFT.Contract.MintingPaused(&_MembershipNFT.CallOpts)
}

// NextMintTokenId is a free data retrieval call binding the contract method 0xe88fbd41.
//
// Solidity: function nextMintTokenId() view returns(uint32)
func (_MembershipNFT *MembershipNFTCaller) NextMintTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "nextMintTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextMintTokenId is a free data retrieval call binding the contract method 0xe88fbd41.
//
// Solidity: function nextMintTokenId() view returns(uint32)
func (_MembershipNFT *MembershipNFTSession) NextMintTokenId() (uint32, error) {
	return _MembershipNFT.Contract.NextMintTokenId(&_MembershipNFT.CallOpts)
}

// NextMintTokenId is a free data retrieval call binding the contract method 0xe88fbd41.
//
// Solidity: function nextMintTokenId() view returns(uint32)
func (_MembershipNFT *MembershipNFTCallerSession) NextMintTokenId() (uint32, error) {
	return _MembershipNFT.Contract.NextMintTokenId(&_MembershipNFT.CallOpts)
}

// NftData is a free data retrieval call binding the contract method 0x8c210975.
//
// Solidity: function nftData(uint256 ) view returns(uint32 transferLockedUntil)
func (_MembershipNFT *MembershipNFTCaller) NftData(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "nftData", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NftData is a free data retrieval call binding the contract method 0x8c210975.
//
// Solidity: function nftData(uint256 ) view returns(uint32 transferLockedUntil)
func (_MembershipNFT *MembershipNFTSession) NftData(arg0 *big.Int) (uint32, error) {
	return _MembershipNFT.Contract.NftData(&_MembershipNFT.CallOpts, arg0)
}

// NftData is a free data retrieval call binding the contract method 0x8c210975.
//
// Solidity: function nftData(uint256 ) view returns(uint32 transferLockedUntil)
func (_MembershipNFT *MembershipNFTCallerSession) NftData(arg0 *big.Int) (uint32, error) {
	return _MembershipNFT.Contract.NftData(&_MembershipNFT.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MembershipNFT *MembershipNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MembershipNFT *MembershipNFTSession) Owner() (common.Address, error) {
	return _MembershipNFT.Contract.Owner(&_MembershipNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MembershipNFT *MembershipNFTCallerSession) Owner() (common.Address, error) {
	return _MembershipNFT.Contract.Owner(&_MembershipNFT.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MembershipNFT *MembershipNFTCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MembershipNFT *MembershipNFTSession) ProxiableUUID() ([32]byte, error) {
	return _MembershipNFT.Contract.ProxiableUUID(&_MembershipNFT.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MembershipNFT *MembershipNFTCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MembershipNFT.Contract.ProxiableUUID(&_MembershipNFT.CallOpts)
}

// RequiredEapPointsPerEapDeposit is a free data retrieval call binding the contract method 0xdd2460a8.
//
// Solidity: function requiredEapPointsPerEapDeposit(uint256 ) view returns(uint64)
func (_MembershipNFT *MembershipNFTCaller) RequiredEapPointsPerEapDeposit(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "requiredEapPointsPerEapDeposit", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RequiredEapPointsPerEapDeposit is a free data retrieval call binding the contract method 0xdd2460a8.
//
// Solidity: function requiredEapPointsPerEapDeposit(uint256 ) view returns(uint64)
func (_MembershipNFT *MembershipNFTSession) RequiredEapPointsPerEapDeposit(arg0 *big.Int) (uint64, error) {
	return _MembershipNFT.Contract.RequiredEapPointsPerEapDeposit(&_MembershipNFT.CallOpts, arg0)
}

// RequiredEapPointsPerEapDeposit is a free data retrieval call binding the contract method 0xdd2460a8.
//
// Solidity: function requiredEapPointsPerEapDeposit(uint256 ) view returns(uint64)
func (_MembershipNFT *MembershipNFTCallerSession) RequiredEapPointsPerEapDeposit(arg0 *big.Int) (uint64, error) {
	return _MembershipNFT.Contract.RequiredEapPointsPerEapDeposit(&_MembershipNFT.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MembershipNFT *MembershipNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MembershipNFT *MembershipNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MembershipNFT.Contract.SupportsInterface(&_MembershipNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MembershipNFT *MembershipNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MembershipNFT.Contract.SupportsInterface(&_MembershipNFT.CallOpts, interfaceId)
}

// TierOf is a free data retrieval call binding the contract method 0x53f96df2.
//
// Solidity: function tierOf(uint256 _tokenId) view returns(uint8)
func (_MembershipNFT *MembershipNFTCaller) TierOf(opts *bind.CallOpts, _tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "tierOf", _tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TierOf is a free data retrieval call binding the contract method 0x53f96df2.
//
// Solidity: function tierOf(uint256 _tokenId) view returns(uint8)
func (_MembershipNFT *MembershipNFTSession) TierOf(_tokenId *big.Int) (uint8, error) {
	return _MembershipNFT.Contract.TierOf(&_MembershipNFT.CallOpts, _tokenId)
}

// TierOf is a free data retrieval call binding the contract method 0x53f96df2.
//
// Solidity: function tierOf(uint256 _tokenId) view returns(uint8)
func (_MembershipNFT *MembershipNFTCallerSession) TierOf(_tokenId *big.Int) (uint8, error) {
	return _MembershipNFT.Contract.TierOf(&_MembershipNFT.CallOpts, _tokenId)
}

// TierPointsOf is a free data retrieval call binding the contract method 0xa88bd1e6.
//
// Solidity: function tierPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCaller) TierPointsOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "tierPointsOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierPointsOf is a free data retrieval call binding the contract method 0xa88bd1e6.
//
// Solidity: function tierPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTSession) TierPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.TierPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// TierPointsOf is a free data retrieval call binding the contract method 0xa88bd1e6.
//
// Solidity: function tierPointsOf(uint256 _tokenId) view returns(uint40)
func (_MembershipNFT *MembershipNFTCallerSession) TierPointsOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.TierPointsOf(&_MembershipNFT.CallOpts, _tokenId)
}

// TransferLockedUntil is a free data retrieval call binding the contract method 0xc90681db.
//
// Solidity: function transferLockedUntil(uint256 _tokenId) view returns(uint32)
func (_MembershipNFT *MembershipNFTCaller) TransferLockedUntil(opts *bind.CallOpts, _tokenId *big.Int) (uint32, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "transferLockedUntil", _tokenId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TransferLockedUntil is a free data retrieval call binding the contract method 0xc90681db.
//
// Solidity: function transferLockedUntil(uint256 _tokenId) view returns(uint32)
func (_MembershipNFT *MembershipNFTSession) TransferLockedUntil(_tokenId *big.Int) (uint32, error) {
	return _MembershipNFT.Contract.TransferLockedUntil(&_MembershipNFT.CallOpts, _tokenId)
}

// TransferLockedUntil is a free data retrieval call binding the contract method 0xc90681db.
//
// Solidity: function transferLockedUntil(uint256 _tokenId) view returns(uint32)
func (_MembershipNFT *MembershipNFTCallerSession) TransferLockedUntil(_tokenId *big.Int) (uint32, error) {
	return _MembershipNFT.Contract.TransferLockedUntil(&_MembershipNFT.CallOpts, _tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_MembershipNFT *MembershipNFTCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_MembershipNFT *MembershipNFTSession) Uri(arg0 *big.Int) (string, error) {
	return _MembershipNFT.Contract.Uri(&_MembershipNFT.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_MembershipNFT *MembershipNFTCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _MembershipNFT.Contract.Uri(&_MembershipNFT.CallOpts, arg0)
}

// ValueOf is a free data retrieval call binding the contract method 0xcadf338f.
//
// Solidity: function valueOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTCaller) ValueOf(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MembershipNFT.contract.Call(opts, &out, "valueOf", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValueOf is a free data retrieval call binding the contract method 0xcadf338f.
//
// Solidity: function valueOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTSession) ValueOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.ValueOf(&_MembershipNFT.CallOpts, _tokenId)
}

// ValueOf is a free data retrieval call binding the contract method 0xcadf338f.
//
// Solidity: function valueOf(uint256 _tokenId) view returns(uint256)
func (_MembershipNFT *MembershipNFTCallerSession) ValueOf(_tokenId *big.Int) (*big.Int, error) {
	return _MembershipNFT.Contract.ValueOf(&_MembershipNFT.CallOpts, _tokenId)
}

// AlertBatchMetadataUpdate is a paid mutator transaction binding the contract method 0x5afcbef9.
//
// Solidity: function alertBatchMetadataUpdate(uint256 startID, uint256 endID) returns()
func (_MembershipNFT *MembershipNFTTransactor) AlertBatchMetadataUpdate(opts *bind.TransactOpts, startID *big.Int, endID *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "alertBatchMetadataUpdate", startID, endID)
}

// AlertBatchMetadataUpdate is a paid mutator transaction binding the contract method 0x5afcbef9.
//
// Solidity: function alertBatchMetadataUpdate(uint256 startID, uint256 endID) returns()
func (_MembershipNFT *MembershipNFTSession) AlertBatchMetadataUpdate(startID *big.Int, endID *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.AlertBatchMetadataUpdate(&_MembershipNFT.TransactOpts, startID, endID)
}

// AlertBatchMetadataUpdate is a paid mutator transaction binding the contract method 0x5afcbef9.
//
// Solidity: function alertBatchMetadataUpdate(uint256 startID, uint256 endID) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) AlertBatchMetadataUpdate(startID *big.Int, endID *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.AlertBatchMetadataUpdate(&_MembershipNFT.TransactOpts, startID, endID)
}

// AlertMetadataUpdate is a paid mutator transaction binding the contract method 0x6800a4f4.
//
// Solidity: function alertMetadataUpdate(uint256 id) returns()
func (_MembershipNFT *MembershipNFTTransactor) AlertMetadataUpdate(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "alertMetadataUpdate", id)
}

// AlertMetadataUpdate is a paid mutator transaction binding the contract method 0x6800a4f4.
//
// Solidity: function alertMetadataUpdate(uint256 id) returns()
func (_MembershipNFT *MembershipNFTSession) AlertMetadataUpdate(id *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.AlertMetadataUpdate(&_MembershipNFT.TransactOpts, id)
}

// AlertMetadataUpdate is a paid mutator transaction binding the contract method 0x6800a4f4.
//
// Solidity: function alertMetadataUpdate(uint256 id) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) AlertMetadataUpdate(id *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.AlertMetadataUpdate(&_MembershipNFT.TransactOpts, id)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_MembershipNFT *MembershipNFTTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "burn", _from, _tokenId, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_MembershipNFT *MembershipNFTSession) Burn(_from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.Burn(&_MembershipNFT.TransactOpts, _from, _tokenId, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) Burn(_from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.Burn(&_MembershipNFT.TransactOpts, _from, _tokenId, _amount)
}

// IncrementLock is a paid mutator transaction binding the contract method 0x803b6605.
//
// Solidity: function incrementLock(uint256 _tokenId, uint32 _blocks) returns()
func (_MembershipNFT *MembershipNFTTransactor) IncrementLock(opts *bind.TransactOpts, _tokenId *big.Int, _blocks uint32) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "incrementLock", _tokenId, _blocks)
}

// IncrementLock is a paid mutator transaction binding the contract method 0x803b6605.
//
// Solidity: function incrementLock(uint256 _tokenId, uint32 _blocks) returns()
func (_MembershipNFT *MembershipNFTSession) IncrementLock(_tokenId *big.Int, _blocks uint32) (*types.Transaction, error) {
	return _MembershipNFT.Contract.IncrementLock(&_MembershipNFT.TransactOpts, _tokenId, _blocks)
}

// IncrementLock is a paid mutator transaction binding the contract method 0x803b6605.
//
// Solidity: function incrementLock(uint256 _tokenId, uint32 _blocks) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) IncrementLock(_tokenId *big.Int, _blocks uint32) (*types.Transaction, error) {
	return _MembershipNFT.Contract.IncrementLock(&_MembershipNFT.TransactOpts, _tokenId, _blocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string _metadataURI, address _membershipManagerInstance) returns()
func (_MembershipNFT *MembershipNFTTransactor) Initialize(opts *bind.TransactOpts, _metadataURI string, _membershipManagerInstance common.Address) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "initialize", _metadataURI, _membershipManagerInstance)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string _metadataURI, address _membershipManagerInstance) returns()
func (_MembershipNFT *MembershipNFTSession) Initialize(_metadataURI string, _membershipManagerInstance common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.Initialize(&_MembershipNFT.TransactOpts, _metadataURI, _membershipManagerInstance)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string _metadataURI, address _membershipManagerInstance) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) Initialize(_metadataURI string, _membershipManagerInstance common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.Initialize(&_MembershipNFT.TransactOpts, _metadataURI, _membershipManagerInstance)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0xe587a617.
//
// Solidity: function initializeOnUpgrade(address _liquidityPoolAddress) returns()
func (_MembershipNFT *MembershipNFTTransactor) InitializeOnUpgrade(opts *bind.TransactOpts, _liquidityPoolAddress common.Address) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "initializeOnUpgrade", _liquidityPoolAddress)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0xe587a617.
//
// Solidity: function initializeOnUpgrade(address _liquidityPoolAddress) returns()
func (_MembershipNFT *MembershipNFTSession) InitializeOnUpgrade(_liquidityPoolAddress common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.InitializeOnUpgrade(&_MembershipNFT.TransactOpts, _liquidityPoolAddress)
}

// InitializeOnUpgrade is a paid mutator transaction binding the contract method 0xe587a617.
//
// Solidity: function initializeOnUpgrade(address _liquidityPoolAddress) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) InitializeOnUpgrade(_liquidityPoolAddress common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.InitializeOnUpgrade(&_MembershipNFT.TransactOpts, _liquidityPoolAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(uint256)
func (_MembershipNFT *MembershipNFTTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(uint256)
func (_MembershipNFT *MembershipNFTSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.Mint(&_MembershipNFT.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(uint256)
func (_MembershipNFT *MembershipNFTTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MembershipNFT.Contract.Mint(&_MembershipNFT.TransactOpts, _to, _amount)
}

// ProcessDepositFromEapUser is a paid mutator transaction binding the contract method 0xe7c88ef4.
//
// Solidity: function processDepositFromEapUser(address _user, uint32 _eapDepositBlockNumber, uint256 _snapshotEthAmount, uint256 _points, bytes32[] _merkleProof) returns()
func (_MembershipNFT *MembershipNFTTransactor) ProcessDepositFromEapUser(opts *bind.TransactOpts, _user common.Address, _eapDepositBlockNumber uint32, _snapshotEthAmount *big.Int, _points *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "processDepositFromEapUser", _user, _eapDepositBlockNumber, _snapshotEthAmount, _points, _merkleProof)
}

// ProcessDepositFromEapUser is a paid mutator transaction binding the contract method 0xe7c88ef4.
//
// Solidity: function processDepositFromEapUser(address _user, uint32 _eapDepositBlockNumber, uint256 _snapshotEthAmount, uint256 _points, bytes32[] _merkleProof) returns()
func (_MembershipNFT *MembershipNFTSession) ProcessDepositFromEapUser(_user common.Address, _eapDepositBlockNumber uint32, _snapshotEthAmount *big.Int, _points *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.ProcessDepositFromEapUser(&_MembershipNFT.TransactOpts, _user, _eapDepositBlockNumber, _snapshotEthAmount, _points, _merkleProof)
}

// ProcessDepositFromEapUser is a paid mutator transaction binding the contract method 0xe7c88ef4.
//
// Solidity: function processDepositFromEapUser(address _user, uint32 _eapDepositBlockNumber, uint256 _snapshotEthAmount, uint256 _points, bytes32[] _merkleProof) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) ProcessDepositFromEapUser(_user common.Address, _eapDepositBlockNumber uint32, _snapshotEthAmount *big.Int, _points *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.ProcessDepositFromEapUser(&_MembershipNFT.TransactOpts, _user, _eapDepositBlockNumber, _snapshotEthAmount, _points, _merkleProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MembershipNFT *MembershipNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MembershipNFT *MembershipNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _MembershipNFT.Contract.RenounceOwnership(&_MembershipNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MembershipNFT *MembershipNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MembershipNFT.Contract.RenounceOwnership(&_MembershipNFT.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MembershipNFT *MembershipNFTTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MembershipNFT *MembershipNFTSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SafeBatchTransferFrom(&_MembershipNFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SafeBatchTransferFrom(&_MembershipNFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MembershipNFT *MembershipNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MembershipNFT *MembershipNFTSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SafeTransferFrom(&_MembershipNFT.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SafeTransferFrom(&_MembershipNFT.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MembershipNFT *MembershipNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MembershipNFT *MembershipNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetApprovalForAll(&_MembershipNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetApprovalForAll(&_MembershipNFT.TransactOpts, operator, approved)
}

// SetContractMetadataURI is a paid mutator transaction binding the contract method 0xa201fc50.
//
// Solidity: function setContractMetadataURI(string _newURI) returns()
func (_MembershipNFT *MembershipNFTTransactor) SetContractMetadataURI(opts *bind.TransactOpts, _newURI string) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "setContractMetadataURI", _newURI)
}

// SetContractMetadataURI is a paid mutator transaction binding the contract method 0xa201fc50.
//
// Solidity: function setContractMetadataURI(string _newURI) returns()
func (_MembershipNFT *MembershipNFTSession) SetContractMetadataURI(_newURI string) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetContractMetadataURI(&_MembershipNFT.TransactOpts, _newURI)
}

// SetContractMetadataURI is a paid mutator transaction binding the contract method 0xa201fc50.
//
// Solidity: function setContractMetadataURI(string _newURI) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SetContractMetadataURI(_newURI string) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetContractMetadataURI(&_MembershipNFT.TransactOpts, _newURI)
}

// SetMaxTokenId is a paid mutator transaction binding the contract method 0xe2fab5b2.
//
// Solidity: function setMaxTokenId(uint32 _maxTokenId) returns()
func (_MembershipNFT *MembershipNFTTransactor) SetMaxTokenId(opts *bind.TransactOpts, _maxTokenId uint32) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "setMaxTokenId", _maxTokenId)
}

// SetMaxTokenId is a paid mutator transaction binding the contract method 0xe2fab5b2.
//
// Solidity: function setMaxTokenId(uint32 _maxTokenId) returns()
func (_MembershipNFT *MembershipNFTSession) SetMaxTokenId(_maxTokenId uint32) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetMaxTokenId(&_MembershipNFT.TransactOpts, _maxTokenId)
}

// SetMaxTokenId is a paid mutator transaction binding the contract method 0xe2fab5b2.
//
// Solidity: function setMaxTokenId(uint32 _maxTokenId) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SetMaxTokenId(_maxTokenId uint32) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetMaxTokenId(&_MembershipNFT.TransactOpts, _maxTokenId)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _newURI) returns()
func (_MembershipNFT *MembershipNFTTransactor) SetMetadataURI(opts *bind.TransactOpts, _newURI string) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "setMetadataURI", _newURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _newURI) returns()
func (_MembershipNFT *MembershipNFTSession) SetMetadataURI(_newURI string) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetMetadataURI(&_MembershipNFT.TransactOpts, _newURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _newURI) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SetMetadataURI(_newURI string) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetMetadataURI(&_MembershipNFT.TransactOpts, _newURI)
}

// SetMintingPaused is a paid mutator transaction binding the contract method 0xad13419d.
//
// Solidity: function setMintingPaused(bool _paused) returns()
func (_MembershipNFT *MembershipNFTTransactor) SetMintingPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "setMintingPaused", _paused)
}

// SetMintingPaused is a paid mutator transaction binding the contract method 0xad13419d.
//
// Solidity: function setMintingPaused(bool _paused) returns()
func (_MembershipNFT *MembershipNFTSession) SetMintingPaused(_paused bool) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetMintingPaused(&_MembershipNFT.TransactOpts, _paused)
}

// SetMintingPaused is a paid mutator transaction binding the contract method 0xad13419d.
//
// Solidity: function setMintingPaused(bool _paused) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SetMintingPaused(_paused bool) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetMintingPaused(&_MembershipNFT.TransactOpts, _paused)
}

// SetUpForEap is a paid mutator transaction binding the contract method 0xa0f1f5e7.
//
// Solidity: function setUpForEap(bytes32 _newMerkleRoot, uint64[] _requiredEapPointsPerEapDeposit) returns()
func (_MembershipNFT *MembershipNFTTransactor) SetUpForEap(opts *bind.TransactOpts, _newMerkleRoot [32]byte, _requiredEapPointsPerEapDeposit []uint64) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "setUpForEap", _newMerkleRoot, _requiredEapPointsPerEapDeposit)
}

// SetUpForEap is a paid mutator transaction binding the contract method 0xa0f1f5e7.
//
// Solidity: function setUpForEap(bytes32 _newMerkleRoot, uint64[] _requiredEapPointsPerEapDeposit) returns()
func (_MembershipNFT *MembershipNFTSession) SetUpForEap(_newMerkleRoot [32]byte, _requiredEapPointsPerEapDeposit []uint64) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetUpForEap(&_MembershipNFT.TransactOpts, _newMerkleRoot, _requiredEapPointsPerEapDeposit)
}

// SetUpForEap is a paid mutator transaction binding the contract method 0xa0f1f5e7.
//
// Solidity: function setUpForEap(bytes32 _newMerkleRoot, uint64[] _requiredEapPointsPerEapDeposit) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) SetUpForEap(_newMerkleRoot [32]byte, _requiredEapPointsPerEapDeposit []uint64) (*types.Transaction, error) {
	return _MembershipNFT.Contract.SetUpForEap(&_MembershipNFT.TransactOpts, _newMerkleRoot, _requiredEapPointsPerEapDeposit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MembershipNFT *MembershipNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MembershipNFT *MembershipNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.TransferOwnership(&_MembershipNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.TransferOwnership(&_MembershipNFT.TransactOpts, newOwner)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_MembershipNFT *MembershipNFTTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_MembershipNFT *MembershipNFTSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _MembershipNFT.Contract.UpdateAdmin(&_MembershipNFT.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _MembershipNFT.Contract.UpdateAdmin(&_MembershipNFT.TransactOpts, _address, _isAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MembershipNFT *MembershipNFTTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MembershipNFT *MembershipNFTSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.UpgradeTo(&_MembershipNFT.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MembershipNFT *MembershipNFTTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MembershipNFT.Contract.UpgradeTo(&_MembershipNFT.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MembershipNFT *MembershipNFTTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MembershipNFT *MembershipNFTSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.UpgradeToAndCall(&_MembershipNFT.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MembershipNFT *MembershipNFTTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MembershipNFT.Contract.UpgradeToAndCall(&_MembershipNFT.TransactOpts, newImplementation, data)
}

// MembershipNFTAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the MembershipNFT contract.
type MembershipNFTAdminChangedIterator struct {
	Event *MembershipNFTAdminChanged // Event containing the contract specifics and raw log

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
func (it *MembershipNFTAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTAdminChanged)
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
		it.Event = new(MembershipNFTAdminChanged)
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
func (it *MembershipNFTAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTAdminChanged represents a AdminChanged event raised by the MembershipNFT contract.
type MembershipNFTAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MembershipNFT *MembershipNFTFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MembershipNFTAdminChangedIterator, error) {

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MembershipNFTAdminChangedIterator{contract: _MembershipNFT.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MembershipNFT *MembershipNFTFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MembershipNFTAdminChanged) (event.Subscription, error) {

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTAdminChanged)
				if err := _MembershipNFT.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_MembershipNFT *MembershipNFTFilterer) ParseAdminChanged(log types.Log) (*MembershipNFTAdminChanged, error) {
	event := new(MembershipNFTAdminChanged)
	if err := _MembershipNFT.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MembershipNFT contract.
type MembershipNFTApprovalForAllIterator struct {
	Event *MembershipNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MembershipNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTApprovalForAll)
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
		it.Event = new(MembershipNFTApprovalForAll)
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
func (it *MembershipNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTApprovalForAll represents a ApprovalForAll event raised by the MembershipNFT contract.
type MembershipNFTApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MembershipNFT *MembershipNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*MembershipNFTApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTApprovalForAllIterator{contract: _MembershipNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MembershipNFT *MembershipNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MembershipNFTApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTApprovalForAll)
				if err := _MembershipNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_MembershipNFT *MembershipNFTFilterer) ParseApprovalForAll(log types.Log) (*MembershipNFTApprovalForAll, error) {
	event := new(MembershipNFTApprovalForAll)
	if err := _MembershipNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the MembershipNFT contract.
type MembershipNFTBatchMetadataUpdateIterator struct {
	Event *MembershipNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MembershipNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTBatchMetadataUpdate)
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
		it.Event = new(MembershipNFTBatchMetadataUpdate)
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
func (it *MembershipNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the MembershipNFT contract.
type MembershipNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MembershipNFT *MembershipNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*MembershipNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MembershipNFTBatchMetadataUpdateIterator{contract: _MembershipNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MembershipNFT *MembershipNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MembershipNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTBatchMetadataUpdate)
				if err := _MembershipNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_MembershipNFT *MembershipNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*MembershipNFTBatchMetadataUpdate, error) {
	event := new(MembershipNFTBatchMetadataUpdate)
	if err := _MembershipNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the MembershipNFT contract.
type MembershipNFTBeaconUpgradedIterator struct {
	Event *MembershipNFTBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *MembershipNFTBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTBeaconUpgraded)
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
		it.Event = new(MembershipNFTBeaconUpgraded)
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
func (it *MembershipNFTBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTBeaconUpgraded represents a BeaconUpgraded event raised by the MembershipNFT contract.
type MembershipNFTBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MembershipNFT *MembershipNFTFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MembershipNFTBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTBeaconUpgradedIterator{contract: _MembershipNFT.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MembershipNFT *MembershipNFTFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MembershipNFTBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTBeaconUpgraded)
				if err := _MembershipNFT.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_MembershipNFT *MembershipNFTFilterer) ParseBeaconUpgraded(log types.Log) (*MembershipNFTBeaconUpgraded, error) {
	event := new(MembershipNFTBeaconUpgraded)
	if err := _MembershipNFT.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MembershipNFT contract.
type MembershipNFTInitializedIterator struct {
	Event *MembershipNFTInitialized // Event containing the contract specifics and raw log

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
func (it *MembershipNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTInitialized)
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
		it.Event = new(MembershipNFTInitialized)
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
func (it *MembershipNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTInitialized represents a Initialized event raised by the MembershipNFT contract.
type MembershipNFTInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MembershipNFT *MembershipNFTFilterer) FilterInitialized(opts *bind.FilterOpts) (*MembershipNFTInitializedIterator, error) {

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MembershipNFTInitializedIterator{contract: _MembershipNFT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MembershipNFT *MembershipNFTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MembershipNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTInitialized)
				if err := _MembershipNFT.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MembershipNFT *MembershipNFTFilterer) ParseInitialized(log types.Log) (*MembershipNFTInitialized, error) {
	event := new(MembershipNFTInitialized)
	if err := _MembershipNFT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTMerkleUpdatedIterator is returned from FilterMerkleUpdated and is used to iterate over the raw logs and unpacked data for MerkleUpdated events raised by the MembershipNFT contract.
type MembershipNFTMerkleUpdatedIterator struct {
	Event *MembershipNFTMerkleUpdated // Event containing the contract specifics and raw log

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
func (it *MembershipNFTMerkleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTMerkleUpdated)
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
		it.Event = new(MembershipNFTMerkleUpdated)
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
func (it *MembershipNFTMerkleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTMerkleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTMerkleUpdated represents a MerkleUpdated event raised by the MembershipNFT contract.
type MembershipNFTMerkleUpdated struct {
	Arg0 [32]byte
	Arg1 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMerkleUpdated is a free log retrieval operation binding the contract event 0x00d0856a1dd55617bdcd651751cfd654397ebc2a435d173e3780203cf2f39a17.
//
// Solidity: event MerkleUpdated(bytes32 arg0, bytes32 arg1)
func (_MembershipNFT *MembershipNFTFilterer) FilterMerkleUpdated(opts *bind.FilterOpts) (*MembershipNFTMerkleUpdatedIterator, error) {

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "MerkleUpdated")
	if err != nil {
		return nil, err
	}
	return &MembershipNFTMerkleUpdatedIterator{contract: _MembershipNFT.contract, event: "MerkleUpdated", logs: logs, sub: sub}, nil
}

// WatchMerkleUpdated is a free log subscription operation binding the contract event 0x00d0856a1dd55617bdcd651751cfd654397ebc2a435d173e3780203cf2f39a17.
//
// Solidity: event MerkleUpdated(bytes32 arg0, bytes32 arg1)
func (_MembershipNFT *MembershipNFTFilterer) WatchMerkleUpdated(opts *bind.WatchOpts, sink chan<- *MembershipNFTMerkleUpdated) (event.Subscription, error) {

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "MerkleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTMerkleUpdated)
				if err := _MembershipNFT.contract.UnpackLog(event, "MerkleUpdated", log); err != nil {
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

// ParseMerkleUpdated is a log parse operation binding the contract event 0x00d0856a1dd55617bdcd651751cfd654397ebc2a435d173e3780203cf2f39a17.
//
// Solidity: event MerkleUpdated(bytes32 arg0, bytes32 arg1)
func (_MembershipNFT *MembershipNFTFilterer) ParseMerkleUpdated(log types.Log) (*MembershipNFTMerkleUpdated, error) {
	event := new(MembershipNFTMerkleUpdated)
	if err := _MembershipNFT.contract.UnpackLog(event, "MerkleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the MembershipNFT contract.
type MembershipNFTMetadataUpdateIterator struct {
	Event *MembershipNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MembershipNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTMetadataUpdate)
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
		it.Event = new(MembershipNFTMetadataUpdate)
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
func (it *MembershipNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTMetadataUpdate represents a MetadataUpdate event raised by the MembershipNFT contract.
type MembershipNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MembershipNFT *MembershipNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*MembershipNFTMetadataUpdateIterator, error) {

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MembershipNFTMetadataUpdateIterator{contract: _MembershipNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MembershipNFT *MembershipNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MembershipNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTMetadataUpdate)
				if err := _MembershipNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_MembershipNFT *MembershipNFTFilterer) ParseMetadataUpdate(log types.Log) (*MembershipNFTMetadataUpdate, error) {
	event := new(MembershipNFTMetadataUpdate)
	if err := _MembershipNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTMintingPausedIterator is returned from FilterMintingPaused and is used to iterate over the raw logs and unpacked data for MintingPaused events raised by the MembershipNFT contract.
type MembershipNFTMintingPausedIterator struct {
	Event *MembershipNFTMintingPaused // Event containing the contract specifics and raw log

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
func (it *MembershipNFTMintingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTMintingPaused)
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
		it.Event = new(MembershipNFTMintingPaused)
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
func (it *MembershipNFTMintingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTMintingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTMintingPaused represents a MintingPaused event raised by the MembershipNFT contract.
type MembershipNFTMintingPaused struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintingPaused is a free log retrieval operation binding the contract event 0x5708b32f825aa2f1192da18d8205d00f2cb633e72d1c5bb33a3111ccc61be8f4.
//
// Solidity: event MintingPaused(bool isPaused)
func (_MembershipNFT *MembershipNFTFilterer) FilterMintingPaused(opts *bind.FilterOpts) (*MembershipNFTMintingPausedIterator, error) {

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "MintingPaused")
	if err != nil {
		return nil, err
	}
	return &MembershipNFTMintingPausedIterator{contract: _MembershipNFT.contract, event: "MintingPaused", logs: logs, sub: sub}, nil
}

// WatchMintingPaused is a free log subscription operation binding the contract event 0x5708b32f825aa2f1192da18d8205d00f2cb633e72d1c5bb33a3111ccc61be8f4.
//
// Solidity: event MintingPaused(bool isPaused)
func (_MembershipNFT *MembershipNFTFilterer) WatchMintingPaused(opts *bind.WatchOpts, sink chan<- *MembershipNFTMintingPaused) (event.Subscription, error) {

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "MintingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTMintingPaused)
				if err := _MembershipNFT.contract.UnpackLog(event, "MintingPaused", log); err != nil {
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

// ParseMintingPaused is a log parse operation binding the contract event 0x5708b32f825aa2f1192da18d8205d00f2cb633e72d1c5bb33a3111ccc61be8f4.
//
// Solidity: event MintingPaused(bool isPaused)
func (_MembershipNFT *MembershipNFTFilterer) ParseMintingPaused(log types.Log) (*MembershipNFTMintingPaused, error) {
	event := new(MembershipNFTMintingPaused)
	if err := _MembershipNFT.contract.UnpackLog(event, "MintingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MembershipNFT contract.
type MembershipNFTOwnershipTransferredIterator struct {
	Event *MembershipNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MembershipNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTOwnershipTransferred)
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
		it.Event = new(MembershipNFTOwnershipTransferred)
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
func (it *MembershipNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTOwnershipTransferred represents a OwnershipTransferred event raised by the MembershipNFT contract.
type MembershipNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MembershipNFT *MembershipNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MembershipNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTOwnershipTransferredIterator{contract: _MembershipNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MembershipNFT *MembershipNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MembershipNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTOwnershipTransferred)
				if err := _MembershipNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MembershipNFT *MembershipNFTFilterer) ParseOwnershipTransferred(log types.Log) (*MembershipNFTOwnershipTransferred, error) {
	event := new(MembershipNFTOwnershipTransferred)
	if err := _MembershipNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTTokenLockedIterator is returned from FilterTokenLocked and is used to iterate over the raw logs and unpacked data for TokenLocked events raised by the MembershipNFT contract.
type MembershipNFTTokenLockedIterator struct {
	Event *MembershipNFTTokenLocked // Event containing the contract specifics and raw log

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
func (it *MembershipNFTTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTTokenLocked)
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
		it.Event = new(MembershipNFTTokenLocked)
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
func (it *MembershipNFTTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTTokenLocked represents a TokenLocked event raised by the MembershipNFT contract.
type MembershipNFTTokenLocked struct {
	TokenId *big.Int
	Until   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenLocked is a free log retrieval operation binding the contract event 0xf43c632cabf9b23317efd3d40fdd5be89b56d735e96a36181c46684f9f9ea81c.
//
// Solidity: event TokenLocked(uint256 indexed _tokenId, uint256 until)
func (_MembershipNFT *MembershipNFTFilterer) FilterTokenLocked(opts *bind.FilterOpts, _tokenId []*big.Int) (*MembershipNFTTokenLockedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "TokenLocked", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTTokenLockedIterator{contract: _MembershipNFT.contract, event: "TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTokenLocked is a free log subscription operation binding the contract event 0xf43c632cabf9b23317efd3d40fdd5be89b56d735e96a36181c46684f9f9ea81c.
//
// Solidity: event TokenLocked(uint256 indexed _tokenId, uint256 until)
func (_MembershipNFT *MembershipNFTFilterer) WatchTokenLocked(opts *bind.WatchOpts, sink chan<- *MembershipNFTTokenLocked, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "TokenLocked", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTTokenLocked)
				if err := _MembershipNFT.contract.UnpackLog(event, "TokenLocked", log); err != nil {
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

// ParseTokenLocked is a log parse operation binding the contract event 0xf43c632cabf9b23317efd3d40fdd5be89b56d735e96a36181c46684f9f9ea81c.
//
// Solidity: event TokenLocked(uint256 indexed _tokenId, uint256 until)
func (_MembershipNFT *MembershipNFTFilterer) ParseTokenLocked(log types.Log) (*MembershipNFTTokenLocked, error) {
	event := new(MembershipNFTTokenLocked)
	if err := _MembershipNFT.contract.UnpackLog(event, "TokenLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the MembershipNFT contract.
type MembershipNFTTransferBatchIterator struct {
	Event *MembershipNFTTransferBatch // Event containing the contract specifics and raw log

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
func (it *MembershipNFTTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTTransferBatch)
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
		it.Event = new(MembershipNFTTransferBatch)
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
func (it *MembershipNFTTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTTransferBatch represents a TransferBatch event raised by the MembershipNFT contract.
type MembershipNFTTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MembershipNFT *MembershipNFTFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MembershipNFTTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTTransferBatchIterator{contract: _MembershipNFT.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MembershipNFT *MembershipNFTFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *MembershipNFTTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTTransferBatch)
				if err := _MembershipNFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_MembershipNFT *MembershipNFTFilterer) ParseTransferBatch(log types.Log) (*MembershipNFTTransferBatch, error) {
	event := new(MembershipNFTTransferBatch)
	if err := _MembershipNFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the MembershipNFT contract.
type MembershipNFTTransferSingleIterator struct {
	Event *MembershipNFTTransferSingle // Event containing the contract specifics and raw log

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
func (it *MembershipNFTTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTTransferSingle)
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
		it.Event = new(MembershipNFTTransferSingle)
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
func (it *MembershipNFTTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTTransferSingle represents a TransferSingle event raised by the MembershipNFT contract.
type MembershipNFTTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MembershipNFT *MembershipNFTFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MembershipNFTTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTTransferSingleIterator{contract: _MembershipNFT.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MembershipNFT *MembershipNFTFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *MembershipNFTTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTTransferSingle)
				if err := _MembershipNFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_MembershipNFT *MembershipNFTFilterer) ParseTransferSingle(log types.Log) (*MembershipNFTTransferSingle, error) {
	event := new(MembershipNFTTransferSingle)
	if err := _MembershipNFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the MembershipNFT contract.
type MembershipNFTURIIterator struct {
	Event *MembershipNFTURI // Event containing the contract specifics and raw log

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
func (it *MembershipNFTURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTURI)
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
		it.Event = new(MembershipNFTURI)
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
func (it *MembershipNFTURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTURI represents a URI event raised by the MembershipNFT contract.
type MembershipNFTURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MembershipNFT *MembershipNFTFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MembershipNFTURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTURIIterator{contract: _MembershipNFT.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MembershipNFT *MembershipNFTFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MembershipNFTURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTURI)
				if err := _MembershipNFT.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MembershipNFT *MembershipNFTFilterer) ParseURI(log types.Log) (*MembershipNFTURI, error) {
	event := new(MembershipNFTURI)
	if err := _MembershipNFT.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MembershipNFTUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MembershipNFT contract.
type MembershipNFTUpgradedIterator struct {
	Event *MembershipNFTUpgraded // Event containing the contract specifics and raw log

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
func (it *MembershipNFTUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MembershipNFTUpgraded)
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
		it.Event = new(MembershipNFTUpgraded)
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
func (it *MembershipNFTUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MembershipNFTUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MembershipNFTUpgraded represents a Upgraded event raised by the MembershipNFT contract.
type MembershipNFTUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MembershipNFT *MembershipNFTFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MembershipNFTUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MembershipNFT.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MembershipNFTUpgradedIterator{contract: _MembershipNFT.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MembershipNFT *MembershipNFTFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MembershipNFTUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MembershipNFT.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MembershipNFTUpgraded)
				if err := _MembershipNFT.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_MembershipNFT *MembershipNFTFilterer) ParseUpgraded(log types.Log) (*MembershipNFTUpgraded, error) {
	event := new(MembershipNFTUpgraded)
	if err := _MembershipNFT.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
