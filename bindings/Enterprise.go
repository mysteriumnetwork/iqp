// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// EnterpriseMetaData contains all meta data concerning the Enterprise contract.
var EnterpriseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"BaseUriChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"name\":\"BondingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"converter\",\"type\":\"address\"}],\"name\":\"ConverterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"}],\"name\":\"EnterpriseCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"EnterpriseOnlyCollectionPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnterpriseShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"EnterpriseWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"}],\"name\":\"FixedReserveChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"percent\",\"type\":\"uint16\"}],\"name\":\"GcFeePercentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PaymentTokenChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"poolFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"serviceFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"renterOnlyReturnTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"enterpriseOnlyCollectionTime\",\"type\":\"uint32\"}],\"name\":\"RentalPeriodExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"returner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"rentalAmount\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"gcRewardAmount\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gcRewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUsedReserve\",\"type\":\"uint256\"}],\"name\":\"RentalReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"rentalAmount\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"poolFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"serviceFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"gcFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"renterOnlyReturnTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"enterpriseOnlyCollectionTime\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUsedReserve\",\"type\":\"uint256\"}],\"name\":\"Rented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"RenterOnlyReturnPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumEnterprise.StakeOperation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReserve\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUsedReserve\",\"type\":\"uint256\"}],\"name\":\"StakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"streamingReserve\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"streamingReserveTarget\",\"type\":\"uint112\"}],\"name\":\"StreamingReserveChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"StreamingReserveHalvingPeriodChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"}],\"name\":\"claimStakingReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmountDelta\",\"type\":\"uint256\"}],\"name\":\"decreaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"disablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"enablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"rentalAmount\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"rentalPeriod\",\"type\":\"uint32\"}],\"name\":\"estimateRentalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"rentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxPayment\",\"type\":\"uint256\"}],\"name\":\"extendRentalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondingCurve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConverter\",\"outputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseOnlyCollectionPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"streamingReserveHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"renterOnlyReturnPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enterpriseOnlyCollectionPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint112\",\"name\":\"streamingReserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"streamingReserveTarget\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"streamingReserveUpdated\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPaymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPaymentTokenIndex\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPowerTokens\",\"outputs\":[{\"internalType\":\"contractIPowerToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"}],\"name\":\"getRentalAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"rentalAmount\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"powerTokenIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"startTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"renterOnlyReturnTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enterpriseOnlyCollectionTime\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"gcRewardAmount\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"gcRewardTokenIndex\",\"type\":\"uint16\"}],\"internalType\":\"structIEnterpriseStorage.RentalAgreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRentalToken\",\"outputs\":[{\"internalType\":\"contractIRentalToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRenterOnlyReturnPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"internalType\":\"structEnterpriseStorage.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeToken\",\"outputs\":[{\"internalType\":\"contractIStakeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"}],\"name\":\"getStakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStreamingReserveHalvingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsedReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmountDelta\",\"type\":\"uint256\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"enterpriseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"contractIConverter\",\"name\":\"converter\",\"type\":\"address\"},{\"internalType\":\"contractProxyAdmin\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"enterpriseToken\",\"type\":\"address\"},{\"internalType\":\"contractIStakeToken\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"contractIRentalToken\",\"name\":\"rentalToken\",\"type\":\"address\"}],\"name\":\"initializeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"}],\"name\":\"isRegisteredPowerToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedPaymentToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceSymbol\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"energyGapHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"serviceFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minRentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxRentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"swappingEnabledForever\",\"type\":\"bool\"}],\"name\":\"registerService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"rentalAmount\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"rentalPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxPayment\",\"type\":\"uint256\"}],\"name\":\"rent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"}],\"name\":\"returnRental\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"name\":\"setBondingCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"newConverter\",\"type\":\"address\"}],\"name\":\"setConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollector\",\"type\":\"address\"}],\"name\":\"setEnterpriseCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newPeriod\",\"type\":\"uint32\"}],\"name\":\"setEnterpriseOnlyCollectionPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"setEnterpriseWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newGcFeePercent\",\"type\":\"uint16\"}],\"name\":\"setGcFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newPeriod\",\"type\":\"uint32\"}],\"name\":\"setRenterOnlyReturnPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"streamingReserveHalvingPeriod\",\"type\":\"uint32\"}],\"name\":\"setStreamingReserveHalvingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownEnterpriseForever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rentalTokenId\",\"type\":\"uint256\"}],\"name\":\"transferRental\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeTokenId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"enterpriseFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"enterpriseImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rentalTokenImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeTokenImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"powerTokenImplementation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"powerTokens\",\"type\":\"address[]\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615c9580620000216000396000f3fe608060405234801561001057600080fd5b506004361061027d5760003560e01c80630103f313146102825780630cac36b2146102ad57806318a8a6cf146102c25780632107730c146102e25780632261b07f146102f7578063239cd4a41461031157806324d86f00146103245780632e17de78146103375780632fb2067f1461034a5780633310df9e1461038b5780633513e0dc1461039e578063449497ec146103b15780634741efb3146103c4578063500a1564146103cc57806359bf5d39146103dd5780635a9b0b89146103f357806360f870871461041357806362c1f388146104245780636815f337146104375780636df0bb141461044a5780637c14d9911461045d5780637e6a7cbb146104705780638245ca1e1461048e57806388cc58e4146104a157806389035b61146104b25780638b3240a0146104de5780638da5cb5b146104e65780638fc49ad2146104ee578063960970c7146105015780639ab71193146105145780639c7aa7f8146105275780639f52673c1461053a578063a0bcfc7f1461054f578063a554b26d14610562578063a694fc3a14610573578063abfe35ad14610586578063aebb504114610599578063b19337a4146105a1578063b406bf6d146105b4578063b714ff53146105d5578063b96266fa146105e6578063bec10cde146105fb578063bfd84fb41461060e578063c4d66de814610621578063c852d20014610634578063c9a304cb14610647578063cdc2aebf1461065a578063ce325bf81461066d578063dccdc7d9146106a2578063dd9919e7146106b3578063ef1f9f39146106c8578063f2fde38b146106db578063f87c4261146106ee578063faac38ef14610701578063fcb2884e1461071c575b600080fd5b610295610290366004614e4c565b610724565b60405160019190910b81526020015b60405180910390f35b6102b5610751565b6040516102a491906155f0565b6102d56102d036600461536b565b6107e3565b6040516102a491906156df565b6102f56102f0366004614e4c565b610938565b005b6004546001600160a01b03165b6040516102a491906154c0565b6102f561031f36600461536b565b610a12565b6102f56103323660046153e0565b610bea565b6102f561034536600461536b565b610e6f565b61037b610358366004614e4c565b6001600160a01b0316600090815260076020526040812054600190810b900b1390565b60405190151581526020016102a4565b6102f561039936600461536b565b6110bf565b6102f56103ac3660046153e0565b6114b4565b6102f56103bf366004615333565b6115c1565b6102f561164b565b6001546001600160a01b0316610304565b6103e5611745565b6040519081526020016102a4565b6103fb61176a565b6040516102a49c9b9a99989796959493929190615603565b6000546001600160a01b0316610304565b6102f5610432366004614e4c565b61192b565b6102f5610445366004615159565b6119f1565b6102f5610458366004614e84565b611d5c565b6102f561046b366004615401565b6120c9565b600654600160a01b900463ffffffff165b6040516102a491906157bc565b6103e561049c366004614f51565b6121a0565b6003546001600160a01b0316610304565b61037b6104c0366004614e4c565b6001600160a01b031660009081526013602052604090205460ff1690565b6103046122f6565b610304612317565b6102f56104fc366004614faa565b61232d565b6102f561050f36600461500b565b612a00565b6102f561052236600461520b565b612cea565b6102f5610535366004614e4c565b6131a2565b600654600160c01b900463ffffffff16610481565b6102f561055d3660046150b1565b6132e8565b6005546001600160a01b0316610304565b6102f561058136600461536b565b613360565b6103e561059436600461536b565b61350a565b6103e5613538565b6102f56105af366004614e4c565b61354a565b600654600160e01b900461ffff1660405161ffff90911681526020016102a4565b6002546001600160a01b0316610304565b6105ee613610565b6040516102a4919061555a565b6102f56106093660046153e0565b613671565b6102f561061c366004614e4c565b613862565b6102f561062f366004614e4c565b6138f8565b6102f561064236600461539b565b6139cd565b6102f5610655366004615401565b613db9565b61030461066836600461536b565b613e8f565b61068061067b36600461536b565b613ecd565b60408051825181526020808401519082015291810151908201526060016102a4565b6006546001600160a01b0316610304565b600354600160a01b900463ffffffff16610481565b6102f56106d6366004615067565b613fa7565b6102f56106e9366004614e4c565b614076565b6102f56106fc366004615401565b614167565b600e54600d54604080519283526020830191909152016102a4565b600a546103e5565b6001600160a01b03811660009081526007602052604081205461074b90600190810b6158c5565b92915050565b60606010805461076090615975565b80601f016020809104026020016040519081016040528092919081815260200182805461078c90615975565b80156107d95780601f106107ae576101008083540402835291602001916107d9565b820191906000526020600020905b8154815290600101906020018083116107bc57829003601f168201915b5050505050905090565b6040805161010081018252600080825260208201819052818301819052606082018190526080820181905260a0820181905260c0820181905260e082015260025491516331a9108f60e11b81526004810184905290916001600160a01b031690636352211e9060240160206040518083038186803b15801561086457600080fd5b505afa158015610878573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089c9190614e68565b505060009081526011602090815260409182902082516101008101845281546001600160701b038082168352600160701b80830461ffff90811696850196909652600160801b830463ffffffff90811697850197909752600160a01b830487166060850152600160c01b830487166080850152600160e01b90920490951660a083015260019092015493841660c082015292041660e082015290565b33610941612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061098c5760405162461bcd60e51b815260040161098391906155f0565b60405180910390fd5b506040805180820190915260028152611b1960f11b60208201526001600160a01b0382166109cd5760405162461bcd60e51b815260040161098391906155f0565b50600680546001600160a01b0319166001600160a01b038316179055604051600080516020615c4083398151915290610a079083906154c0565b60405180910390a150565b6001546040516331a9108f60e11b815260048101839052829133916001600160a01b0390911690636352211e9060240160206040518083038186803b158015610a5a57600080fd5b505afa158015610a6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a929190614e68565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090610ad45760405162461bcd60e51b815260040161098391906155f0565b506000828152601260205260408120805460018201549192909190610af7611745565b90506000610b0683858461422d565b9050610b118261425f565b811115604051806040016040528060028152602001611a1b60f11b81525090610b4d5760405162461bcd60e51b815260040161098391906155f0565b50600054610b65906001600160a01b0316338361426f565b6000610b7185846142c5565b90506000610b7f8286615932565b9050610b8b83826142e0565b600187018290556000336001600160a01b03168a600080516020615a80833981519152868a8688600c548b8d610bc19190615932565b600a54604051610bd7979695949392919061578c565b60405180910390a4505050505050505050565b6001546040516331a9108f60e11b815260048101849052839133916001600160a01b0390911690636352211e9060240160206040518083038186803b158015610c3257600080fd5b505afa158015610c46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6a9190614e68565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090610cac5760405162461bcd60e51b815260040161098391906155f0565b50600083815260126020908152604091829020825160608101845281548152600182015481840152600291820154818501819052845180860190955291845261035360f41b928401929092529091904311610d1a5760405162461bcd60e51b815260040161098391906155f0565b5080516040805180820190915260028152611a1b60f11b602082015290841115610d575760405162461bcd60e51b815260040161098391906155f0565b506000610d62611745565b9050610d6d8161425f565b841115604051806040016040528060028152602001611a1b60f11b81525090610da95760405162461bcd60e51b815260040161098391906155f0565b50600054610dc1906001600160a01b0316338661426f565b6000610dcd85836142c5565b90508260200151811115610de2575060208201515b610dec85826142e0565b6020838101805183900380825285518890038652600089815260129093526040909220600181019290925584519182905551600c5460049233928a92600080516020615a80833981519152928b929091889190610e49858c615932565b600a54604051610e5f979695949392919061578c565b60405180910390a4505050505050565b6001546040516331a9108f60e11b815260048101839052829133916001600160a01b0390911690636352211e9060240160206040518083038186803b158015610eb757600080fd5b505afa158015610ecb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eef9190614e68565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090610f315760405162461bcd60e51b815260040161098391906155f0565b5060006012600084815260200190815260200160002090504381600201541060405180604001604052806002815260200161035360f41b81525090610f895760405162461bcd60e51b815260040161098391906155f0565b5060018101546000610f99611745565b90506000610fa78383614367565b9050610fb28261425f565b811115604051806040016040528060028152602001611a1b60f11b81525090610fee5760405162461bcd60e51b815260040161098391906155f0565b50600054611006906001600160a01b0316338361426f565b61101081846142e0565b600154604051630852cd8d60e31b8152600481018890526001600160a01b03909116906342966c6890602401600060405180830381600087803b15801561105657600080fd5b505af115801561106a573d6000803e3d6000fd5b5050506000878152601260205260408120818155600181018290556002908101919091559050336001600160a01b031687600080516020615a80833981519152846000886000600c54898b610e499190615932565b60008181526011602090815260409182902082516101008101845281546001600160701b0380821680845261ffff600160701b80850482168689015263ffffffff600160801b86048116878b0152600160a01b860481166060880152600160c01b860481166080880152600160e01b90950490941660a086015260019095015491821660c085015291900490921660e082015283518085019094526002845261068760f31b9284019290925290919061118b5760405162461bcd60e51b815260040161098391906155f0565b506002546040516331a9108f60e11b8152600481018490526000916001600160a01b031690636352211e9060240160206040518083038186803b1580156111d157600080fd5b505afa1580156111e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112099190614e68565b905060004290508063ffffffff16836080015163ffffffff1610806112365750336001600160a01b038316145b60405180604001604052806002815260200161353360f01b8152509061126f5760405162461bcd60e51b815260040161098391906155f0565b508063ffffffff168360a0015163ffffffff1610806112965750336001600160a01b038316145b806112ab57506005546001600160a01b031633145b604051806040016040528060028152602001610d4d60f21b815250906112e45760405162461bcd60e51b815260040161098391906155f0565b50600354600160c01b900460ff1661131b5782600001516001600160701b0316600a60008282546113159190615932565b90915550505b6014836020015161ffff168154811061134457634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b0316336001600160a01b0316857fe8c3d212180d23288b990f9938fef98c64caec34bc128fccea9f7eea006ba73886600001518760c0015160088960e0015161ffff16815481106113ca57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03166113e7611745565b600a54604080516001600160701b0396871681529490951660208501526001600160a01b0392909216838501526060830152608082015290519081900360a00190a4600254604051633f34d4cf60e21b8152600481018690523360248201526001600160a01b039091169063fcd3533c90604401600060405180830381600087803b15801561147557600080fd5b505af1158015611489573d6000803e3d6000fd5b50505060009485525050601160205250506040812090815560010180546001600160801b0319169055565b336114bd612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906114ff5760405162461bcd60e51b815260040161098391906155f0565b5061150f600a600360401b615886565b82111560405180604001604052806002815260200161373760f01b8152509061154b5760405162461bcd60e51b815260040161098391906155f0565b5060408051808201909152600281526106e760f31b6020820152600160401b82111561158a5760405162461bcd60e51b815260040161098391906155f0565b50600e829055600d8190556040805183815260208101839052600080516020615ae083398151915291015b60405180910390a15050565b336115ca612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061160c5760405162461bcd60e51b815260040161098391906155f0565b506006805461ffff60e01b1916600160e01b61ffff841690810291909117909155604051908152600080516020615b2083398151915290602001610a07565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156116905760405162461bcd60e51b815260040161098391906155f0565b503361169a612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906116dc5760405162461bcd60e51b815260040161098391906155f0565b5060038054600160c01b60ff60c01b199091161790556000600a819055600b8054600160701b81046001600160701b03166001600160701b03199091161790556040517f6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c0729190a1565b600061174f614377565b6001600160701b0316600954611765919061584f565b905090565b600354600654600c54600954600a54600b54600f80546060978897600097889788978897889788978897889788978897949660109663ffffffff600160a01b95869004811697958504811696600160c01b860482169661ffff600160e01b9788900416969094936001600160701b0380821694600160701b8304909116939290910416908c906117f990615975565b80601f016020809104026020016040519081016040528092919081815260200182805461182590615975565b80156118725780601f1061184757610100808354040283529160200191611872565b820191906000526020600020905b81548152906001019060200180831161185557829003601f168201915b50505050509b508a805461188590615975565b80601f01602080910402602001604051908101604052809291908181526020018280546118b190615975565b80156118fe5780601f106118d3576101008083540402835291602001916118fe565b820191906000526020600020905b8154815290600101906020018083116118e157829003601f168201915b50505050509a509b509b509b509b509b509b509b509b509b509b509b509b50909192939495969798999a9b565b33611934612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906119765760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261363160f01b60208201526001600160a01b0382166119b75760405162461bcd60e51b815260040161098391906155f0565b50600580546001600160a01b0319166001600160a01b038316179055604051600080516020615bc083398151915290610a079083906154c0565b600f80546119fe90615975565b6040805180820190915260018152601960f91b6020820152915015611a365760405162461bcd60e51b815260040161098391906155f0565b50600087511160405180604001604052806002815260200161383160f01b81525090611a755760405162461bcd60e51b815260040161098391906155f0565b50611a7f816138f8565b81600080516020615b0083398151915280546001600160a01b03929092166001600160a01b031992831617905560038054909116331790558651611aca90600f9060208a0190614c4c565b50611ad760108787614cd0565b5060068054600480546001600160a01b03199081166001600160a01b03888116919091179092556005805492861692909116821790556003805463ffffffff60a01b191661127560a71b179055600165ffff0000000160c01b0319909116600160e01b61ffff8816026001600160c01b03191617176102a360a61b1763ffffffff60c01b19166102a360c71b179055611b756064600560401b615886565b600e55611b87600a600360401b615886565b600d55604051600080516020615a6083398151915290611baa90889088906155c1565b60405180910390a1600654604051600160e01b90910461ffff168152600080516020615b208339815191529060200160405180910390a1600454604051600080516020615b6083398151915291611c0c916001600160a01b03909116906154c0565b60405180910390a1600654604051600080516020615c4083398151915291611c3f916001600160a01b03909116906154c0565b60405180910390a1600554604051600080516020615bc083398151915291611c72916001600160a01b03909116906154c0565b60405180910390a1600080516020615b40833981519152600360149054906101000a900463ffffffff16604051611ca991906157bc565b60405180910390a1600080516020615ac0833981519152600660149054906101000a900463ffffffff16604051611ce091906157bc565b60405180910390a1600080516020615be0833981519152600660189054906101000a900463ffffffff16604051611d1791906157bc565b60405180910390a1600080516020615ae0833981519152600e54600d54604051611d4b929190918252602082015260400190565b60405180910390a150505050505050565b33611d65612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090611da75760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261383360f01b60208201526001600160a01b038816611de85760405162461bcd60e51b815260040161098391906155f0565b50600380546001600160a01b0319166001600160a01b0389161790556000611e0e6122f6565b90506001600160a01b03871615611e805760405163266a23b160e21b81526001600160a01b038216906399a88ec490611e4d9030908b906004016155a7565b600060405180830381600087803b158015611e6757600080fd5b505af1158015611e7b573d6000803e3d6000fd5b505050505b6001600160a01b03861615611ef65760025460405163266a23b160e21b81526001600160a01b03838116926399a88ec492611ec392909116908a906004016155a7565b600060405180830381600087803b158015611edd57600080fd5b505af1158015611ef1573d6000803e3d6000fd5b505050505b6001600160a01b03851615611f6c5760015460405163266a23b160e21b81526001600160a01b03838116926399a88ec492611f39929091169089906004016155a7565b600060405180830381600087803b158015611f5357600080fd5b505af1158015611f67573d6000803e3d6000fd5b505050505b6001600160a01b038416156120bf5760005b828110156120bd5760136000858584818110611faa57634e487b7160e01b600052603260045260246000fd5b9050602002016020810190611fbf9190614e4c565b6001600160a01b03168152602080820192909252604090810160002054815180830190925260018252601b60f91b928201929092529060ff166120155760405162461bcd60e51b815260040161098391906155f0565b50816001600160a01b03166399a88ec485858481811061204557634e487b7160e01b600052603260045260246000fd5b905060200201602081019061205a9190614e4c565b876040518363ffffffff1660e01b81526004016120789291906155a7565b600060405180830381600087803b15801561209257600080fd5b505af11580156120a6573d6000803e3d6000fd5b5050505080806120b5906159b0565b915050611f7e565b505b5050505050505050565b336120d2612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906121145760405162461bcd60e51b815260040161098391906155f0565b506006546040805180820190915260028152610d8d60f21b60208201529063ffffffff600160c01b909104811690831611156121635760405162461bcd60e51b815260040161098391906155f0565b506006805463ffffffff60a01b1916600160a01b63ffffffff841602179055604051600080516020615ac083398151915290610a079083906157bc565b600354604080518082019091526002815261373560f01b6020820152600091600160c01b900460ff16156121e75760405162461bcd60e51b815260040161098391906155f0565b506001600160a01b03851660009081526013602090815260409182902054825180840190935260018352601b60f91b9183019190915260ff1661223d5760405162461bcd60e51b815260040161098391906155f0565b506000806000876001600160a01b031663ff1fb8868888886040518463ffffffff1660e01b815260040161227393929190615511565b60606040518083038186803b15801561228b57600080fd5b505afa15801561229f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122c391906152f2565b91945092509050806122d58385615824565b6122df9190615824565b6001600160701b031693505050505b949350505050565b6000600080516020615b008339815191525b546001600160a01b0316919050565b6000600080516020615aa0833981519152612308565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156123725760405162461bcd60e51b815260040161098391906155f0565b506040805180820190915260028152611b9b60f11b60208201526001600160701b0384166123b35760405162461bcd60e51b815260040161098391906155f0565b506001600160a01b03851660009081526013602090815260409182902054825180840190935260018352601b60f91b9183019190915260ff166124095760405162461bcd60e51b815260040161098391906155f0565b50612412613538565b836001600160701b03161115604051806040016040528060028152602001611a1b60f11b815250906124575760405162461bcd60e51b815260040161098391906155f0565b506000806000876001600160a01b031663ff1fb8868888886040518463ffffffff1660e01b815260040161248d93929190615511565b60606040518083038186803b1580156124a557600080fd5b505afa1580156124b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124dd91906152f2565b9194509250905083816124f08486615824565b6124fa9190615824565b6001600160701b0316111560405180604001604052806002815260200161343760f01b8152509061253e5760405162461bcd60e51b815260040161098391906155f0565b5061255387836001600160701b0316856143df565b600254612578906001600160a01b03898116913391166001600160701b038516614587565b856001600160701b0316600a6000828254612593919061584f565b90915550600090506125a58642615867565b6006549091506000906125c590600160a01b900463ffffffff1683615867565b6006549091506000906125e590600160c01b900463ffffffff1684615867565b90506000600260009054906101000a90046001600160a01b03166001600160a01b031663caa0f92a6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561263957600080fd5b505af115801561264d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126719190615383565b90506040518061010001604052808b6001600160701b031681526020018d6001600160a01b03166381045ead6040518163ffffffff1660e01b815260040160206040518083038186803b1580156126c757600080fd5b505afa1580156126db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ff919061534f565b61ffff1681526020014263ffffffff1681526020018563ffffffff1681526020018463ffffffff1681526020018363ffffffff168152602001866001600160701b031681526020016127508d610724565b61ffff9081169091526000838152601160209081526040918290208451815492860151868501516060880151608089015160a08a01516001600160701b039586166001600160801b031998891617600160701b958b16860217600160801b600160c01b031916600160801b63ffffffff9586160263ffffffff60a01b191617600160a01b93851693909302929092176001600160c01b0316600160c01b918416919091026001600160e01b031617600160e01b929091169190910217835560c08701516001909301805460e0909801519390921696909316959095179316029190911790915560025490516335313c2160e11b815282916001600160a01b031690636a627842906128659033906004016154c0565b602060405180830381600087803b15801561287f57600080fd5b505af1158015612893573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128b79190615383565b146128d257634e487b7160e01b600052600160045260246000fd5b604051635a85946560e01b8152600481018290526001600160a01b038d1690635a85946590602401600060405180830381600087803b15801561291457600080fd5b505af1158015612928573d6000803e3d6000fd5b505050508b6001600160a01b0316336001600160a01b0316827f59dee24cee42a5fd3c1be706cc8906676ac6950d523c17fa522275c74e7b9c758e8e8c8c8c428d8d8d612973611745565b600a54604080516001600160a01b03909c168c526001600160701b039a8b1660208d0152988a16988b019890985295881660608a015296909316608088015263ffffffff91821660a0880152811660c087015290811660e08601529092166101008401526101208301919091526101408201526101600160405180910390a4505050505050505050505050565b600280546040805180820190915291825261333960f01b60208301526001600160a01b03163314612a445760405162461bcd60e51b815260040161098391906155f0565b5060008181526011602090815260409182902082516101008101845281546001600160701b0380821680845261ffff600160701b80850482168689015263ffffffff600160801b86048116878b0152600160a01b860481166060880152600160c01b860481166080880152600160e01b90950490941660a086015260019095015491821660c085015291900490921660e082015283518085019094526002845261068760f31b92840192909252909190612b115760405162461bcd60e51b815260040161098391906155f0565b50606081015160208201516014805463ffffffff9093164211926001600160a01b0388811615939088161592600092909161ffff16908110612b6357634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031690508115612be557845160405163079cc67960e41b81526001600160a01b038316916379cc679091612bae918c916004016154ef565b600060405180830381600087803b158015612bc857600080fd5b505af1158015612bdc573d6000803e3d6000fd5b505050506120bf565b8215612c195784516040516340c10f1960e01b81526001600160a01b038316916340c10f1991612bae918b916004016154ef565b83612cba5784516040516333bebb7760e01b81526001600160a01b038a8116600483015289811660248301526001600160701b039092166044820152908216906333bebb7790606401602060405180830381600087803b158015612c7c57600080fd5b505af1158015612c90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cb4919061504b565b506120bf565b60408051808201825260028152611a9960f11b6020820152905162461bcd60e51b815261098391906004016155f0565b33612cf3612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090612d355760405162461bcd60e51b815260040161098391906155f0565b50600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff1615612d7b5760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261034360f41b60208201526001600160a01b038716612dbc5760405162461bcd60e51b815260040161098391906155f0565b50601454604080518082019091526002815261343160f01b60208201529061ffff11612dfb5760405162461bcd60e51b815260040161098391906155f0565b506003546000906001600160a01b031663c31011cc612e186122f6565b6040518263ffffffff1660e01b8152600401612e3491906154c0565b602060405180830381600087803b158015612e4e57600080fd5b505af1158015612e62573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e869190614e68565b905060008060009054906101000a90046001600160a01b03166001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015612ed757600080fd5b505afa158015612eeb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612f1391908101906150f0565b90506000818c604051602001612f2a929190615484565b6040516020818303038152906040529050826001600160a01b0316631624f6c68e8360008054906101000a90046001600160a01b03166001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015612f9857600080fd5b505afa158015612fac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fd0919061541b565b6040518463ffffffff1660e01b8152600401612fee939291906156a6565b600060405180830381600087803b15801561300857600080fd5b505af115801561301c573d6000803e3d6000fd5b505050505050806001600160a01b031663b29c98ad30898b878b8f6014805490508d8d8c6040518b63ffffffff1660e01b81526004016130d79a999897969594939291906001600160a01b039a8b1681529890991660208901526001600160701b039690961660408801526001600160601b0394909416606087015261ffff928316608087015263ffffffff91821660a087015290911660c085015290811660e0840152166101008201529015156101208201526101400190565b600060405180830381600087803b1580156130f157600080fd5b505af1158015613105573d6000803e3d6000fd5b50506014805460018082019092557fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec0180546001600160a01b0319166001600160a01b038616908117909155600081815260136020526040808220805460ff191690941790935591519093507f2fa31fbaacf5eaf61d648ea7528ada6efb69bfb06d2c3bd35ce511a820fce53e9250a25050505050505050505050565b336131ab612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906131ed5760405162461bcd60e51b815260040161098391906155f0565b506001600160a01b03811660009081526007602090815260409182902054825180840190935260028352611b9960f11b91830191909152600190810b900b6132485760405162461bcd60e51b815260040161098391906155f0565b506001600160a01b038116600090815260076020526040812054600190810b900b13156132e5576001600160a01b0381166000908152600760205260409020546132949060010b6159cb565b6001600160a01b038216600090815260076020526040808220805460019490940b61ffff1661ffff19909416939093179092559051600080516020615c0083398151915291610a07918491906154d4565b50565b336132f1612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906133335760405162461bcd60e51b815260040161098391906155f0565b5061334060108383614cd0565b50600080516020615a6083398151915282826040516115b59291906155c1565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156133a55760405162461bcd60e51b815260040161098391906155f0565b506000546133be906001600160a01b0316333084614587565b60006133c8611745565b90506000600c546000146133e5576133e083836142c5565b6133e7565b825b90506133f383826145c5565b6001546040516335313c2160e11b81526000916001600160a01b031690636a627842906134249033906004016154c0565b602060405180830381600087803b15801561343e57600080fd5b505af1158015613452573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134769190615383565b60408051606081018252868152602080820186815243838501908152600086815260129093529390912091518255516001808301919091559151600290910155909150336001600160a01b031682600080516020615a8083398151915287888788600c548c8c6134e6919061584f565b600a546040516134fc979695949392919061578c565b60405180910390a450505050565b600081815260126020526040812060018101548154613531919061352c611745565b61422d565b9392505050565b6000611765613545611745565b61425f565b33613553612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906135955760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261363360f01b60208201526001600160a01b0382166135d65760405162461bcd60e51b815260040161098391906155f0565b50600480546001600160a01b0319166001600160a01b038316179055604051600080516020615b6083398151915290610a079083906154c0565b606060148054806020026020016040519081016040528092919081815260200182805480156107d957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161364a575050505050905090565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156136b65760405162461bcd60e51b815260040161098391906155f0565b506001546040516331a9108f60e11b815260048101849052839133916001600160a01b0390911690636352211e9060240160206040518083038186803b1580156136ff57600080fd5b505afa158015613713573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137379190614e68565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906137795760405162461bcd60e51b815260040161098391906155f0565b50600054613792906001600160a01b0316333085614587565b600061379c611745565b90506000600c546000146137b9576137b484836142c5565b6137bb565b835b90506137c784826145c5565b600085815260126020526040812080549091906137e590879061584f565b905060008383600101546137f9919061584f565b8284556001840181905543600285015590506003336001600160a01b031689600080516020615a808339815191528a868987600c548f8e61383a919061584f565b600a54604051613850979695949392919061578c565b60405180910390a45050505050505050565b3361386b612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906138ad5760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261373160f01b60208201526001600160a01b0382166138ee5760405162461bcd60e51b815260040161098391906155f0565b506132e581614613565b6000613902612317565b6001600160a01b031614604051806040016040528060018152602001601960f91b815250906139445760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261373360f01b60208201526001600160a01b0382166139855760405162461bcd60e51b815260040161098391906155f0565b50600080516020615aa083398151915280546001600160a01b0319166001600160a01b038316908117909155604051600090600080516020615ba0833981519152908290a350565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff1615613a125760405162461bcd60e51b815260040161098391906155f0565b50600084815260116020908152604091829020805483518085019094526002845261068760f31b9284019290925291906001600160701b0316613a685760405162461bcd60e51b815260040161098391906155f0565b50805460148054600092600160701b900461ffff16908110613a9a57634e487b7160e01b600052603260045260246000fd5b60009182526020909120015482546001600160a01b0390911691504290613acf908690600160a01b900463ffffffff16615867565b63ffffffff16101560405180604001604052806002815260200161343960f01b81525090613b105760405162461bcd60e51b815260040161098391906155f0565b50600a548254613b29906001600160701b031682615932565b600a5582546040516001627023bd60e11b0319815260009182916001600160a01b0386169163ff1fb88691613b70918c916001600160701b03909116908c90600401615511565b60606040518083038186803b158015613b8857600080fd5b505afa158015613b9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bc091906152f2565b50600a859055909250905085613bd68284615824565b6001600160701b0316111560405180604001604052806002815260200161343760f01b81525090613c1a5760405162461bcd60e51b815260040161098391906155f0565b50613c2f88826001600160701b0316846143df565b8454600090613c4c908990600160a01b900463ffffffff16615867565b600654909150600090613c6c90600160a01b900463ffffffff1683615867565b600654909150600090613c8c90600160c01b900463ffffffff1684615867565b8854600160a01b600160e01b031916600160a01b63ffffffff8681169190910263ffffffff60c01b191691909117600160c01b85831602176001600160e01b0316600160e01b91831691909102178955604051635a85946560e01b8152600481018e90529091506001600160a01b03881690635a85946590602401600060405180830381600087803b158015613d2157600080fd5b505af1158015613d35573d6000803e3d6000fd5b5050604080516001600160a01b038f1681526001600160701b03898116602083015288168183015263ffffffff87811660608301528681166080830152851660a082015290513393508f92507f012ee01e680dee236f33567d88889822f7521f77a65adc92142abf5a2005d3ef9181900360c00190a3505050505050505050505050565b33613dc2612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090613e045760405162461bcd60e51b815260040161098391906155f0565b50600654604080518082019091526002815261363560f01b60208201529063ffffffff808416600160a01b909204161115613e525760405162461bcd60e51b815260040161098391906155f0565b506006805463ffffffff60c01b1916600160c01b63ffffffff841602179055604051600080516020615be083398151915290610a079083906157bc565b600060088281548110613eb257634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031692915050565b613ef160405180606001604052806000815260200160008152602001600081525090565b6001546040516331a9108f60e11b8152600481018490526001600160a01b0390911690636352211e9060240160206040518083038186803b158015613f3557600080fd5b505afa158015613f49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f6d9190614e68565b5050600090815260126020908152604091829020825160608101845281548152600182015492810192909252600201549181019190915290565b6000546040805180820190915260018152601960f91b6020820152906001600160a01b031615613fea5760405162461bcd60e51b815260040161098391906155f0565b506040805180820190915260018152603560f81b60208201526001600160a01b03841661402a5760405162461bcd60e51b815260040161098391906155f0565b50600080546001600160a01b038086166001600160a01b03199283161790925560018054858416908316179055600280549284169290911691909117905561407183614613565b505050565b3361407f612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906140c15760405162461bcd60e51b815260040161098391906155f0565b50604080518082019091526002815261373360f01b60208201526001600160a01b0382166141025760405162461bcd60e51b815260040161098391906155f0565b50806001600160a01b0316614115612317565b6001600160a01b0316600080516020615ba083398151915260405160405180910390a3600080516020615aa083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b33614170612317565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906141b25760405162461bcd60e51b815260040161098391906155f0565b506040805180820190915260028152611b1b60f11b602082015263ffffffff82166141f05760405162461bcd60e51b815260040161098391906155f0565b506003805463ffffffff60a01b1916600160a01b63ffffffff841602179055604051600080516020615b4083398151915290610a079083906157bc565b60008061423a8584614367565b9050838111156142535761424e8482615932565b614256565b60005b95945050505050565b6000600a548261074b9190615932565b6140718363a9059cbb60e01b848460405160240161428e929190615541565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261475e565b60008183600c546142d691906158a6565b6135319190615886565b6000600954905081600c60008282546142f99190615932565b909155505082811061430d5782900361433b565b6000614317614830565b6001600160701b031690508361432d828461584f565b6143379190615932565b9150505b6009819055604051818152600080516020615c20833981519152906020015b60405180910390a1505050565b600c546000906142d684846158a6565b600b546000906143c29063ffffffff600160e01b820416906143ac906001600160701b0380821691600160701b90041661590a565b600354600160a01b900463ffffffff16426148ee565b600b546117659190600160701b90046001600160701b031661590a565b6001600160701b03811682016144006001600160a01b038516333084614587565b6000546001600160a01b0390811690849084908716831461455b576004805460405163095ea7b360e01b81526001600160a01b03808b169363095ea7b39361444d93921691899101615541565b602060405180830381600087803b15801561446757600080fd5b505af115801561447b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061449f919061504b565b506004805460405163029b465d60e01b81526001600160a01b038a8116938201939093526024810187905285831660448201526000929091169063029b465d90606401602060405180830381600087803b1580156144fc57600080fd5b505af1158015614510573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906145349190615383565b90508461454182896158a6565b61454b9190615886565b92506145578382615932565b9150505b600654614575906001600160a01b0385811691168461426f565b61457e81614a8f565b50505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526145bf9085906323b872dd60e01b9060840161428e565b50505050565b80600c60008282546145d7919061584f565b90915550506009546000906145ed90849061584f565b6009819055604051818152909150600080516020615c208339815191529060200161435a565b6001600160a01b038116600090815260076020526040902054600190810b900b6146c95760088054600181810183557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390910180546001600160a01b0385166001600160a01b03199091168117909155915460009283526007602052604092839020805491830b61ffff1661ffff199092169190911790559051600080516020615c0083398151915291610a07918491906154d4565b6001600160a01b038116600090815260076020526040812054600190810b900b12156132e5576001600160a01b0381166000908152600760205260409020546147149060010b6159cb565b6001600160a01b03821660009081526007602052604090819020805461ffff191661ffff600194850b1617905551600080516020615c0083398151915291610a07918491906154d4565b60006147b3826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614b419092919063ffffffff16565b80519091501561407157808060200190518101906147d1919061504b565b6140715760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610983565b600061483a614377565b600b80546001600160701b0319168082559192508291600e9061486e908490600160701b90046001600160701b031661590a565b82546001600160701b039182166101009390930a928302928202191691909117909155600b805463ffffffff4216600160e01b026001600160e01b038216811792839055604051600080516020615b8083398151915295506148e3949182169282169290921792600160701b90041690615772565b60405180910390a190565b60008463ffffffff168263ffffffff161015604051806040016040528060018152602001600760fb1b815250906149385760405162461bcd60e51b815260040161098391906155f0565b5084820391508263ffffffff168263ffffffff168161496757634e487b7160e01b600052601260045260246000fd5b0463ffffffff16846001600160701b0316901c93508263ffffffff168263ffffffff16816149a557634e487b7160e01b600052601260045260246000fd5b06915063ffffffff821615806149c257506001600160701b038416155b156149ce5750826122ee565b60006001600160701b0385168163ffffffff86811690861671b17217f7d1cf79abc9e3b39803f2f6af40f30281614a1557634e487b7160e01b600052601260045260246000fd5b049050600160901b5b8215614a8257928201928082840281614a4757634e487b7160e01b600052601260045260246000fd5b0493849003939250600160901b018082840281614a7457634e487b7160e01b600052601260045260246000fd5b049250600160901b01614a1e565b5091979650505050505050565b614a97614377565b600b80546001600160701b0319166001600160701b03928316178082558392600e91614acc918591600160701b900416615824565b82546001600160701b039182166101009390930a928302928202191691909117909155600b805463ffffffff4216600160e01b026001600160e01b038216811792839055604051600080516020615b808339815191529550610a07949182169282169290921792600160701b90041690615772565b60606122ee848460008585843b614b9a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610983565b600080866001600160a01b03168587604051614bb69190615468565b60006040518083038185875af1925050503d8060008114614bf3576040519150601f19603f3d011682016040523d82523d6000602084013e614bf8565b606091505b5091509150614c08828286614c13565b979650505050505050565b60608315614c22575081613531565b825115614c325782518084602001fd5b8160405162461bcd60e51b815260040161098391906155f0565b828054614c5890615975565b90600052602060002090601f016020900481019282614c7a5760008555614cc0565b82601f10614c9357805160ff1916838001178555614cc0565b82800160010185558215614cc0579182015b82811115614cc0578251825591602001919060010190614ca5565b50614ccc929150614d44565b5090565b828054614cdc90615975565b90600052602060002090601f016020900481019282614cfe5760008555614cc0565b82601f10614d175782800160ff19823516178555614cc0565b82800160010185558215614cc0579182015b82811115614cc0578235825591602001919060010190614d29565b5b80821115614ccc5760008155600101614d45565b8035614d6481615a17565b919050565b8035614d6481615a2c565b60008083601f840112614d85578182fd5b5081356001600160401b03811115614d9b578182fd5b602083019150836020828501011115614db357600080fd5b9250929050565b600082601f830112614dca578081fd5b8135614ddd614dd8826157fd565b6157cd565b818152846020838601011115614df1578283fd5b816020850160208301379081016020019190915292915050565b8035614d6481615a3a565b8035614d6481615a4f565b803563ffffffff81168114614d6457600080fd5b80356001600160601b0381168114614d6457600080fd5b600060208284031215614e5d578081fd5b813561353181615a17565b600060208284031215614e79578081fd5b815161353181615a17565b600080600080600080600060c0888a031215614e9e578283fd5b8735614ea981615a17565b96506020880135614eb981615a17565b95506040880135614ec981615a17565b94506060880135614ed981615a17565b93506080880135614ee981615a17565b925060a08801356001600160401b0380821115614f04578384fd5b818a0191508a601f830112614f17578384fd5b813581811115614f25578485fd5b8b60208260051b8501011115614f39578485fd5b60208301945080935050505092959891949750929550565b60008060008060808587031215614f66578384fd5b8435614f7181615a17565b93506020850135614f8181615a17565b92506040850135614f9181615a3a565b9150614f9f60608601614e21565b905092959194509250565b600080600080600060a08688031215614fc1578283fd5b8535614fcc81615a17565b94506020860135614fdc81615a17565b93506040860135614fec81615a3a565b9250614ffa60608701614e21565b949793965091946080013592915050565b60008060006060848603121561501f578081fd5b833561502a81615a17565b9250602084013561503a81615a17565b929592945050506040919091013590565b60006020828403121561505c578081fd5b815161353181615a2c565b60008060006060848603121561507b578081fd5b833561508681615a17565b9250602084013561509681615a17565b915060408401356150a681615a17565b809150509250925092565b600080602083850312156150c3578182fd5b82356001600160401b038111156150d8578283fd5b6150e485828601614d74565b90969095509350505050565b600060208284031215615101578081fd5b81516001600160401b03811115615116578182fd5b8201601f81018413615126578182fd5b8051615134614dd8826157fd565b818152856020838501011115615148578384fd5b614256826020830160208601615949565b600080600080600080600060c0888a031215615173578081fd5b87356001600160401b0380821115615189578283fd5b6151958b838c01614dba565b985060208a01359150808211156151aa578283fd5b506151b78a828b01614d74565b90975095505060408801356151cb81615a4f565b935060608801356151db81615a17565b925060808801356151eb81615a17565b915060a08801356151fb81615a17565b8091505092959891949750929550565b6000806000806000806000806000806101408b8d03121561522a578384fd5b8a356001600160401b0380821115615240578586fd5b61524c8e838f01614dba565b9b5060208d0135915080821115615261578586fd5b5061526e8d828e01614dba565b99505061527d60408c01614e21565b975061528b60608c01614e0b565b965061529960808c01614d59565b95506152a760a08c01614e16565b94506152b560c08c01614e21565b93506152c360e08c01614e21565b92506152d26101008c01614e35565b91506152e16101208c01614d69565b90509295989b9194979a5092959850565b600080600060608486031215615306578081fd5b835161531181615a3a565b602085015190935061532281615a3a565b60408501519092506150a681615a3a565b600060208284031215615344578081fd5b813561353181615a4f565b600060208284031215615360578081fd5b815161353181615a4f565b60006020828403121561537c578081fd5b5035919050565b600060208284031215615394578081fd5b5051919050565b600080600080608085870312156153b0578182fd5b8435935060208501356153c281615a17565b92506153d060408601614e21565b9396929550929360600135925050565b600080604083850312156153f2578182fd5b50508035926020909101359150565b600060208284031215615412578081fd5b61353182614e21565b60006020828403121561542c578081fd5b815160ff81168114613531578182fd5b60008151808452615454816020860160208601615949565b601f01601f19169290920160200192915050565b6000825161547a818460208701615949565b9190910192915050565b60008351615496818460208801615949565b600160fd1b90830190815283516154b4816001840160208801615949565b01600101949350505050565b6001600160a01b0391909116815260200190565b6001600160a01b039290921682521515602082015260400190565b6001600160a01b039290921682526001600160701b0316602082015260400190565b6001600160a01b039390931683526001600160701b0391909116602083015263ffffffff16604082015260600190565b6001600160a01b03929092168252602082015260400190565b6020808252825182820181905260009190848201906040850190845b8181101561559b5783516001600160a01b031683529284019291840191600101615576565b50909695505050505050565b6001600160a01b0392831681529116602082015260400190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b602081526000613531602083018461543c565b6101808152600061561861018083018f61543c565b828103602084015261562a818f61543c565b63ffffffff8e811660408601528d811660608601528c16608085015261ffff8b1660a085015260c084018a905260e0840189905261010084018890526001600160701b03878116610120860152861661014085015291506156889050565b63ffffffff83166101608301529d9c50505050505050505050505050565b6060815260006156b9606083018661543c565b82810360208401526156cb818661543c565b91505060ff83166040830152949350505050565b81516001600160701b0316815260208083015161ffff169082015260408083015163ffffffff9081169183019190915260608084015182169083015260808084015182169083015260a080840151918216908301526101008201905060c083015161575560c08401826001600160701b03169052565b5060e083015161576b60e084018261ffff169052565b5092915050565b6001600160701b0392831681529116602082015260400190565b968752602087019590955260408601939093526060850191909152608084015260a083015260c082015260e00190565b63ffffffff91909116815260200190565b604051601f8201601f191681016001600160401b03811182821017156157f5576157f5615a01565b604052919050565b60006001600160401b0382111561581657615816615a01565b50601f01601f191660200190565b60006001600160701b03828116848216808303821115615846576158466159eb565b01949350505050565b60008219821115615862576158626159eb565b500190565b600063ffffffff808316818516808303821115615846576158466159eb565b6000826158a157634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156158c0576158c06159eb565b500290565b60008160010b8360010b82811281617fff19018312811516156158ea576158ea6159eb565b81617fff018313811615615900576159006159eb565b5090039392505050565b60006001600160701b038381169083168181101561592a5761592a6159eb565b039392505050565b600082821015615944576159446159eb565b500390565b60005b8381101561596457818101518382015260200161594c565b838111156145bf5750506000910152565b600181811c9082168061598957607f821691505b602082108114156159aa57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156159c4576159c46159eb565b5060010190565b60008160010b617fff198114156159e4576159e46159eb565b9003919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146132e557600080fd5b80151581146132e557600080fd5b6001600160701b03811681146132e557600080fd5b61ffff811681146132e557600080fdfe87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d603429b4bf882467b990034c9cdcf7d1c4c3f189e62ce27af5e9bd563659fa2864f471908b72bb76dae5bd24599026e7bf3ddb256497722888ffa422f83729ede745abdabfd9615abf44cc5ea5223a16cd377282452510a2de5121d44b8c097f1926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39fd1248cccb5fef9131c731321e43e9a924840ffee7dc68c7d1d3e5cb7dedcae0327894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a156469b2baf1f9580f11e84f351d1ac9c543726f7023ba5d6d278fc487898fac055b6befcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656dfc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a118298be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e07aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4d076b5bc77c447b04d82e76a12ec958bfe6d48418cce446aaf9ebeb3136638eb92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907a70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180471e342623158b92281274ef7263e4f9a0e0b748c1c328113afbb58742c1789fa264697066735822122071e9212e43e9b067e465ad6cfe91cd6ab698473f75bc13d411aff98c714cda1264736f6c63430008040033",
}

// EnterpriseABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseMetaData.ABI instead.
var EnterpriseABI = EnterpriseMetaData.ABI

// EnterpriseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnterpriseMetaData.Bin instead.
var EnterpriseBin = EnterpriseMetaData.Bin

// DeployEnterprise deploys a new Ethereum contract, binding an instance of Enterprise to it.
func DeployEnterprise(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Enterprise, error) {
	parsed, err := EnterpriseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnterpriseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Enterprise{EnterpriseCaller: EnterpriseCaller{contract: contract}, EnterpriseTransactor: EnterpriseTransactor{contract: contract}, EnterpriseFilterer: EnterpriseFilterer{contract: contract}}, nil
}

// Enterprise is an auto generated Go binding around an Ethereum contract.
type Enterprise struct {
	EnterpriseCaller     // Read-only binding to the contract
	EnterpriseTransactor // Write-only binding to the contract
	EnterpriseFilterer   // Log filterer for contract events
}

// EnterpriseCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseSession struct {
	Contract     *Enterprise       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnterpriseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseCallerSession struct {
	Contract *EnterpriseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EnterpriseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseTransactorSession struct {
	Contract     *EnterpriseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EnterpriseRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseRaw struct {
	Contract *Enterprise // Generic contract binding to access the raw methods on
}

// EnterpriseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseCallerRaw struct {
	Contract *EnterpriseCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseTransactorRaw struct {
	Contract *EnterpriseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterprise creates a new instance of Enterprise, bound to a specific deployed contract.
func NewEnterprise(address common.Address, backend bind.ContractBackend) (*Enterprise, error) {
	contract, err := bindEnterprise(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Enterprise{EnterpriseCaller: EnterpriseCaller{contract: contract}, EnterpriseTransactor: EnterpriseTransactor{contract: contract}, EnterpriseFilterer: EnterpriseFilterer{contract: contract}}, nil
}

// NewEnterpriseCaller creates a new read-only instance of Enterprise, bound to a specific deployed contract.
func NewEnterpriseCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseCaller, error) {
	contract, err := bindEnterprise(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseCaller{contract: contract}, nil
}

// NewEnterpriseTransactor creates a new write-only instance of Enterprise, bound to a specific deployed contract.
func NewEnterpriseTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseTransactor, error) {
	contract, err := bindEnterprise(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseTransactor{contract: contract}, nil
}

// NewEnterpriseFilterer creates a new log filterer instance of Enterprise, bound to a specific deployed contract.
func NewEnterpriseFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseFilterer, error) {
	contract, err := bindEnterprise(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseFilterer{contract: contract}, nil
}

// bindEnterprise binds a generic wrapper to an already deployed contract.
func bindEnterprise(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnterpriseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enterprise *EnterpriseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Enterprise.Contract.EnterpriseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enterprise *EnterpriseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enterprise.Contract.EnterpriseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enterprise *EnterpriseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enterprise.Contract.EnterpriseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enterprise *EnterpriseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Enterprise.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enterprise *EnterpriseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enterprise.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enterprise *EnterpriseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enterprise.Contract.contract.Transact(opts, method, params...)
}

// EstimateRentalFee is a free data retrieval call binding the contract method 0x8245ca1e.
//
// Solidity: function estimateRentalFee(address powerToken, address paymentToken, uint112 rentalAmount, uint32 rentalPeriod) view returns(uint256)
func (_Enterprise *EnterpriseCaller) EstimateRentalFee(opts *bind.CallOpts, powerToken common.Address, paymentToken common.Address, rentalAmount *big.Int, rentalPeriod uint32) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "estimateRentalFee", powerToken, paymentToken, rentalAmount, rentalPeriod)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateRentalFee is a free data retrieval call binding the contract method 0x8245ca1e.
//
// Solidity: function estimateRentalFee(address powerToken, address paymentToken, uint112 rentalAmount, uint32 rentalPeriod) view returns(uint256)
func (_Enterprise *EnterpriseSession) EstimateRentalFee(powerToken common.Address, paymentToken common.Address, rentalAmount *big.Int, rentalPeriod uint32) (*big.Int, error) {
	return _Enterprise.Contract.EstimateRentalFee(&_Enterprise.CallOpts, powerToken, paymentToken, rentalAmount, rentalPeriod)
}

// EstimateRentalFee is a free data retrieval call binding the contract method 0x8245ca1e.
//
// Solidity: function estimateRentalFee(address powerToken, address paymentToken, uint112 rentalAmount, uint32 rentalPeriod) view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) EstimateRentalFee(powerToken common.Address, paymentToken common.Address, rentalAmount *big.Int, rentalPeriod uint32) (*big.Int, error) {
	return _Enterprise.Contract.EstimateRentalFee(&_Enterprise.CallOpts, powerToken, paymentToken, rentalAmount, rentalPeriod)
}

// GetAvailableReserve is a free data retrieval call binding the contract method 0xaebb5041.
//
// Solidity: function getAvailableReserve() view returns(uint256)
func (_Enterprise *EnterpriseCaller) GetAvailableReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getAvailableReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableReserve is a free data retrieval call binding the contract method 0xaebb5041.
//
// Solidity: function getAvailableReserve() view returns(uint256)
func (_Enterprise *EnterpriseSession) GetAvailableReserve() (*big.Int, error) {
	return _Enterprise.Contract.GetAvailableReserve(&_Enterprise.CallOpts)
}

// GetAvailableReserve is a free data retrieval call binding the contract method 0xaebb5041.
//
// Solidity: function getAvailableReserve() view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) GetAvailableReserve() (*big.Int, error) {
	return _Enterprise.Contract.GetAvailableReserve(&_Enterprise.CallOpts)
}

// GetBaseUri is a free data retrieval call binding the contract method 0x0cac36b2.
//
// Solidity: function getBaseUri() view returns(string)
func (_Enterprise *EnterpriseCaller) GetBaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getBaseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseUri is a free data retrieval call binding the contract method 0x0cac36b2.
//
// Solidity: function getBaseUri() view returns(string)
func (_Enterprise *EnterpriseSession) GetBaseUri() (string, error) {
	return _Enterprise.Contract.GetBaseUri(&_Enterprise.CallOpts)
}

// GetBaseUri is a free data retrieval call binding the contract method 0x0cac36b2.
//
// Solidity: function getBaseUri() view returns(string)
func (_Enterprise *EnterpriseCallerSession) GetBaseUri() (string, error) {
	return _Enterprise.Contract.GetBaseUri(&_Enterprise.CallOpts)
}

// GetBondingCurve is a free data retrieval call binding the contract method 0xfaac38ef.
//
// Solidity: function getBondingCurve() view returns(uint256 pole, uint256 slope)
func (_Enterprise *EnterpriseCaller) GetBondingCurve(opts *bind.CallOpts) (struct {
	Pole  *big.Int
	Slope *big.Int
}, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getBondingCurve")

	outstruct := new(struct {
		Pole  *big.Int
		Slope *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pole = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondingCurve is a free data retrieval call binding the contract method 0xfaac38ef.
//
// Solidity: function getBondingCurve() view returns(uint256 pole, uint256 slope)
func (_Enterprise *EnterpriseSession) GetBondingCurve() (struct {
	Pole  *big.Int
	Slope *big.Int
}, error) {
	return _Enterprise.Contract.GetBondingCurve(&_Enterprise.CallOpts)
}

// GetBondingCurve is a free data retrieval call binding the contract method 0xfaac38ef.
//
// Solidity: function getBondingCurve() view returns(uint256 pole, uint256 slope)
func (_Enterprise *EnterpriseCallerSession) GetBondingCurve() (struct {
	Pole  *big.Int
	Slope *big.Int
}, error) {
	return _Enterprise.Contract.GetBondingCurve(&_Enterprise.CallOpts)
}

// GetConverter is a free data retrieval call binding the contract method 0x2261b07f.
//
// Solidity: function getConverter() view returns(address)
func (_Enterprise *EnterpriseCaller) GetConverter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getConverter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConverter is a free data retrieval call binding the contract method 0x2261b07f.
//
// Solidity: function getConverter() view returns(address)
func (_Enterprise *EnterpriseSession) GetConverter() (common.Address, error) {
	return _Enterprise.Contract.GetConverter(&_Enterprise.CallOpts)
}

// GetConverter is a free data retrieval call binding the contract method 0x2261b07f.
//
// Solidity: function getConverter() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetConverter() (common.Address, error) {
	return _Enterprise.Contract.GetConverter(&_Enterprise.CallOpts)
}

// GetEnterpriseCollector is a free data retrieval call binding the contract method 0xa554b26d.
//
// Solidity: function getEnterpriseCollector() view returns(address)
func (_Enterprise *EnterpriseCaller) GetEnterpriseCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getEnterpriseCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseCollector is a free data retrieval call binding the contract method 0xa554b26d.
//
// Solidity: function getEnterpriseCollector() view returns(address)
func (_Enterprise *EnterpriseSession) GetEnterpriseCollector() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseCollector(&_Enterprise.CallOpts)
}

// GetEnterpriseCollector is a free data retrieval call binding the contract method 0xa554b26d.
//
// Solidity: function getEnterpriseCollector() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetEnterpriseCollector() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseCollector(&_Enterprise.CallOpts)
}

// GetEnterpriseOnlyCollectionPeriod is a free data retrieval call binding the contract method 0x9f52673c.
//
// Solidity: function getEnterpriseOnlyCollectionPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCaller) GetEnterpriseOnlyCollectionPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getEnterpriseOnlyCollectionPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEnterpriseOnlyCollectionPeriod is a free data retrieval call binding the contract method 0x9f52673c.
//
// Solidity: function getEnterpriseOnlyCollectionPeriod() view returns(uint32)
func (_Enterprise *EnterpriseSession) GetEnterpriseOnlyCollectionPeriod() (uint32, error) {
	return _Enterprise.Contract.GetEnterpriseOnlyCollectionPeriod(&_Enterprise.CallOpts)
}

// GetEnterpriseOnlyCollectionPeriod is a free data retrieval call binding the contract method 0x9f52673c.
//
// Solidity: function getEnterpriseOnlyCollectionPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCallerSession) GetEnterpriseOnlyCollectionPeriod() (uint32, error) {
	return _Enterprise.Contract.GetEnterpriseOnlyCollectionPeriod(&_Enterprise.CallOpts)
}

// GetEnterpriseToken is a free data retrieval call binding the contract method 0x60f87087.
//
// Solidity: function getEnterpriseToken() view returns(address)
func (_Enterprise *EnterpriseCaller) GetEnterpriseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getEnterpriseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseToken is a free data retrieval call binding the contract method 0x60f87087.
//
// Solidity: function getEnterpriseToken() view returns(address)
func (_Enterprise *EnterpriseSession) GetEnterpriseToken() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseToken(&_Enterprise.CallOpts)
}

// GetEnterpriseToken is a free data retrieval call binding the contract method 0x60f87087.
//
// Solidity: function getEnterpriseToken() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetEnterpriseToken() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseToken(&_Enterprise.CallOpts)
}

// GetEnterpriseWallet is a free data retrieval call binding the contract method 0xdccdc7d9.
//
// Solidity: function getEnterpriseWallet() view returns(address)
func (_Enterprise *EnterpriseCaller) GetEnterpriseWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getEnterpriseWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseWallet is a free data retrieval call binding the contract method 0xdccdc7d9.
//
// Solidity: function getEnterpriseWallet() view returns(address)
func (_Enterprise *EnterpriseSession) GetEnterpriseWallet() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseWallet(&_Enterprise.CallOpts)
}

// GetEnterpriseWallet is a free data retrieval call binding the contract method 0xdccdc7d9.
//
// Solidity: function getEnterpriseWallet() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetEnterpriseWallet() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseWallet(&_Enterprise.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_Enterprise *EnterpriseCaller) GetFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_Enterprise *EnterpriseSession) GetFactory() (common.Address, error) {
	return _Enterprise.Contract.GetFactory(&_Enterprise.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetFactory() (common.Address, error) {
	return _Enterprise.Contract.GetFactory(&_Enterprise.CallOpts)
}

// GetGCFeePercent is a free data retrieval call binding the contract method 0xb406bf6d.
//
// Solidity: function getGCFeePercent() view returns(uint16)
func (_Enterprise *EnterpriseCaller) GetGCFeePercent(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getGCFeePercent")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetGCFeePercent is a free data retrieval call binding the contract method 0xb406bf6d.
//
// Solidity: function getGCFeePercent() view returns(uint16)
func (_Enterprise *EnterpriseSession) GetGCFeePercent() (uint16, error) {
	return _Enterprise.Contract.GetGCFeePercent(&_Enterprise.CallOpts)
}

// GetGCFeePercent is a free data retrieval call binding the contract method 0xb406bf6d.
//
// Solidity: function getGCFeePercent() view returns(uint16)
func (_Enterprise *EnterpriseCallerSession) GetGCFeePercent() (uint16, error) {
	return _Enterprise.Contract.GetGCFeePercent(&_Enterprise.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint32 streamingReserveHalvingPeriod, uint32 renterOnlyReturnPeriod, uint32 enterpriseOnlyCollectionPeriod, uint16 gcFeePercent, uint256 totalShares, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_Enterprise *EnterpriseCaller) GetInfo(opts *bind.CallOpts) (struct {
	Name                           string
	BaseUri                        string
	StreamingReserveHalvingPeriod  uint32
	RenterOnlyReturnPeriod         uint32
	EnterpriseOnlyCollectionPeriod uint32
	GcFeePercent                   uint16
	TotalShares                    *big.Int
	FixedReserve                   *big.Int
	UsedReserve                    *big.Int
	StreamingReserve               *big.Int
	StreamingReserveTarget         *big.Int
	StreamingReserveUpdated        uint32
}, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getInfo")

	outstruct := new(struct {
		Name                           string
		BaseUri                        string
		StreamingReserveHalvingPeriod  uint32
		RenterOnlyReturnPeriod         uint32
		EnterpriseOnlyCollectionPeriod uint32
		GcFeePercent                   uint16
		TotalShares                    *big.Int
		FixedReserve                   *big.Int
		UsedReserve                    *big.Int
		StreamingReserve               *big.Int
		StreamingReserveTarget         *big.Int
		StreamingReserveUpdated        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.BaseUri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StreamingReserveHalvingPeriod = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.RenterOnlyReturnPeriod = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.EnterpriseOnlyCollectionPeriod = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.GcFeePercent = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.TotalShares = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.FixedReserve = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsedReserve = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserve = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserveTarget = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserveUpdated = *abi.ConvertType(out[11], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint32 streamingReserveHalvingPeriod, uint32 renterOnlyReturnPeriod, uint32 enterpriseOnlyCollectionPeriod, uint16 gcFeePercent, uint256 totalShares, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_Enterprise *EnterpriseSession) GetInfo() (struct {
	Name                           string
	BaseUri                        string
	StreamingReserveHalvingPeriod  uint32
	RenterOnlyReturnPeriod         uint32
	EnterpriseOnlyCollectionPeriod uint32
	GcFeePercent                   uint16
	TotalShares                    *big.Int
	FixedReserve                   *big.Int
	UsedReserve                    *big.Int
	StreamingReserve               *big.Int
	StreamingReserveTarget         *big.Int
	StreamingReserveUpdated        uint32
}, error) {
	return _Enterprise.Contract.GetInfo(&_Enterprise.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint32 streamingReserveHalvingPeriod, uint32 renterOnlyReturnPeriod, uint32 enterpriseOnlyCollectionPeriod, uint16 gcFeePercent, uint256 totalShares, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_Enterprise *EnterpriseCallerSession) GetInfo() (struct {
	Name                           string
	BaseUri                        string
	StreamingReserveHalvingPeriod  uint32
	RenterOnlyReturnPeriod         uint32
	EnterpriseOnlyCollectionPeriod uint32
	GcFeePercent                   uint16
	TotalShares                    *big.Int
	FixedReserve                   *big.Int
	UsedReserve                    *big.Int
	StreamingReserve               *big.Int
	StreamingReserveTarget         *big.Int
	StreamingReserveUpdated        uint32
}, error) {
	return _Enterprise.Contract.GetInfo(&_Enterprise.CallOpts)
}

// GetPaymentToken is a free data retrieval call binding the contract method 0xcdc2aebf.
//
// Solidity: function getPaymentToken(uint256 index) view returns(address)
func (_Enterprise *EnterpriseCaller) GetPaymentToken(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getPaymentToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPaymentToken is a free data retrieval call binding the contract method 0xcdc2aebf.
//
// Solidity: function getPaymentToken(uint256 index) view returns(address)
func (_Enterprise *EnterpriseSession) GetPaymentToken(index *big.Int) (common.Address, error) {
	return _Enterprise.Contract.GetPaymentToken(&_Enterprise.CallOpts, index)
}

// GetPaymentToken is a free data retrieval call binding the contract method 0xcdc2aebf.
//
// Solidity: function getPaymentToken(uint256 index) view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetPaymentToken(index *big.Int) (common.Address, error) {
	return _Enterprise.Contract.GetPaymentToken(&_Enterprise.CallOpts, index)
}

// GetPaymentTokenIndex is a free data retrieval call binding the contract method 0x0103f313.
//
// Solidity: function getPaymentTokenIndex(address token) view returns(int16)
func (_Enterprise *EnterpriseCaller) GetPaymentTokenIndex(opts *bind.CallOpts, token common.Address) (int16, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getPaymentTokenIndex", token)

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// GetPaymentTokenIndex is a free data retrieval call binding the contract method 0x0103f313.
//
// Solidity: function getPaymentTokenIndex(address token) view returns(int16)
func (_Enterprise *EnterpriseSession) GetPaymentTokenIndex(token common.Address) (int16, error) {
	return _Enterprise.Contract.GetPaymentTokenIndex(&_Enterprise.CallOpts, token)
}

// GetPaymentTokenIndex is a free data retrieval call binding the contract method 0x0103f313.
//
// Solidity: function getPaymentTokenIndex(address token) view returns(int16)
func (_Enterprise *EnterpriseCallerSession) GetPaymentTokenIndex(token common.Address) (int16, error) {
	return _Enterprise.Contract.GetPaymentTokenIndex(&_Enterprise.CallOpts, token)
}

// GetPowerTokens is a free data retrieval call binding the contract method 0xb96266fa.
//
// Solidity: function getPowerTokens() view returns(address[])
func (_Enterprise *EnterpriseCaller) GetPowerTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getPowerTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPowerTokens is a free data retrieval call binding the contract method 0xb96266fa.
//
// Solidity: function getPowerTokens() view returns(address[])
func (_Enterprise *EnterpriseSession) GetPowerTokens() ([]common.Address, error) {
	return _Enterprise.Contract.GetPowerTokens(&_Enterprise.CallOpts)
}

// GetPowerTokens is a free data retrieval call binding the contract method 0xb96266fa.
//
// Solidity: function getPowerTokens() view returns(address[])
func (_Enterprise *EnterpriseCallerSession) GetPowerTokens() ([]common.Address, error) {
	return _Enterprise.Contract.GetPowerTokens(&_Enterprise.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() view returns(address)
func (_Enterprise *EnterpriseCaller) GetProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getProxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() view returns(address)
func (_Enterprise *EnterpriseSession) GetProxyAdmin() (common.Address, error) {
	return _Enterprise.Contract.GetProxyAdmin(&_Enterprise.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0x8b3240a0.
//
// Solidity: function getProxyAdmin() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetProxyAdmin() (common.Address, error) {
	return _Enterprise.Contract.GetProxyAdmin(&_Enterprise.CallOpts)
}

// GetRentalAgreement is a free data retrieval call binding the contract method 0x18a8a6cf.
//
// Solidity: function getRentalAgreement(uint256 rentalTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_Enterprise *EnterpriseCaller) GetRentalAgreement(opts *bind.CallOpts, rentalTokenId *big.Int) (IEnterpriseStorageRentalAgreement, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getRentalAgreement", rentalTokenId)

	if err != nil {
		return *new(IEnterpriseStorageRentalAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(IEnterpriseStorageRentalAgreement)).(*IEnterpriseStorageRentalAgreement)

	return out0, err

}

// GetRentalAgreement is a free data retrieval call binding the contract method 0x18a8a6cf.
//
// Solidity: function getRentalAgreement(uint256 rentalTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_Enterprise *EnterpriseSession) GetRentalAgreement(rentalTokenId *big.Int) (IEnterpriseStorageRentalAgreement, error) {
	return _Enterprise.Contract.GetRentalAgreement(&_Enterprise.CallOpts, rentalTokenId)
}

// GetRentalAgreement is a free data retrieval call binding the contract method 0x18a8a6cf.
//
// Solidity: function getRentalAgreement(uint256 rentalTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_Enterprise *EnterpriseCallerSession) GetRentalAgreement(rentalTokenId *big.Int) (IEnterpriseStorageRentalAgreement, error) {
	return _Enterprise.Contract.GetRentalAgreement(&_Enterprise.CallOpts, rentalTokenId)
}

// GetRentalToken is a free data retrieval call binding the contract method 0xb714ff53.
//
// Solidity: function getRentalToken() view returns(address)
func (_Enterprise *EnterpriseCaller) GetRentalToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getRentalToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRentalToken is a free data retrieval call binding the contract method 0xb714ff53.
//
// Solidity: function getRentalToken() view returns(address)
func (_Enterprise *EnterpriseSession) GetRentalToken() (common.Address, error) {
	return _Enterprise.Contract.GetRentalToken(&_Enterprise.CallOpts)
}

// GetRentalToken is a free data retrieval call binding the contract method 0xb714ff53.
//
// Solidity: function getRentalToken() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetRentalToken() (common.Address, error) {
	return _Enterprise.Contract.GetRentalToken(&_Enterprise.CallOpts)
}

// GetRenterOnlyReturnPeriod is a free data retrieval call binding the contract method 0x7e6a7cbb.
//
// Solidity: function getRenterOnlyReturnPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCaller) GetRenterOnlyReturnPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getRenterOnlyReturnPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRenterOnlyReturnPeriod is a free data retrieval call binding the contract method 0x7e6a7cbb.
//
// Solidity: function getRenterOnlyReturnPeriod() view returns(uint32)
func (_Enterprise *EnterpriseSession) GetRenterOnlyReturnPeriod() (uint32, error) {
	return _Enterprise.Contract.GetRenterOnlyReturnPeriod(&_Enterprise.CallOpts)
}

// GetRenterOnlyReturnPeriod is a free data retrieval call binding the contract method 0x7e6a7cbb.
//
// Solidity: function getRenterOnlyReturnPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCallerSession) GetRenterOnlyReturnPeriod() (uint32, error) {
	return _Enterprise.Contract.GetRenterOnlyReturnPeriod(&_Enterprise.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_Enterprise *EnterpriseCaller) GetReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_Enterprise *EnterpriseSession) GetReserve() (*big.Int, error) {
	return _Enterprise.Contract.GetReserve(&_Enterprise.CallOpts)
}

// GetReserve is a free data retrieval call binding the contract method 0x59bf5d39.
//
// Solidity: function getReserve() view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) GetReserve() (*big.Int, error) {
	return _Enterprise.Contract.GetReserve(&_Enterprise.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0xce325bf8.
//
// Solidity: function getStake(uint256 stakeTokenId) view returns((uint256,uint256,uint256))
func (_Enterprise *EnterpriseCaller) GetStake(opts *bind.CallOpts, stakeTokenId *big.Int) (EnterpriseStorageStake, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getStake", stakeTokenId)

	if err != nil {
		return *new(EnterpriseStorageStake), err
	}

	out0 := *abi.ConvertType(out[0], new(EnterpriseStorageStake)).(*EnterpriseStorageStake)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0xce325bf8.
//
// Solidity: function getStake(uint256 stakeTokenId) view returns((uint256,uint256,uint256))
func (_Enterprise *EnterpriseSession) GetStake(stakeTokenId *big.Int) (EnterpriseStorageStake, error) {
	return _Enterprise.Contract.GetStake(&_Enterprise.CallOpts, stakeTokenId)
}

// GetStake is a free data retrieval call binding the contract method 0xce325bf8.
//
// Solidity: function getStake(uint256 stakeTokenId) view returns((uint256,uint256,uint256))
func (_Enterprise *EnterpriseCallerSession) GetStake(stakeTokenId *big.Int) (EnterpriseStorageStake, error) {
	return _Enterprise.Contract.GetStake(&_Enterprise.CallOpts, stakeTokenId)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_Enterprise *EnterpriseCaller) GetStakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getStakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_Enterprise *EnterpriseSession) GetStakeToken() (common.Address, error) {
	return _Enterprise.Contract.GetStakeToken(&_Enterprise.CallOpts)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetStakeToken() (common.Address, error) {
	return _Enterprise.Contract.GetStakeToken(&_Enterprise.CallOpts)
}

// GetStakingReward is a free data retrieval call binding the contract method 0xabfe35ad.
//
// Solidity: function getStakingReward(uint256 stakeTokenId) view returns(uint256)
func (_Enterprise *EnterpriseCaller) GetStakingReward(opts *bind.CallOpts, stakeTokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getStakingReward", stakeTokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingReward is a free data retrieval call binding the contract method 0xabfe35ad.
//
// Solidity: function getStakingReward(uint256 stakeTokenId) view returns(uint256)
func (_Enterprise *EnterpriseSession) GetStakingReward(stakeTokenId *big.Int) (*big.Int, error) {
	return _Enterprise.Contract.GetStakingReward(&_Enterprise.CallOpts, stakeTokenId)
}

// GetStakingReward is a free data retrieval call binding the contract method 0xabfe35ad.
//
// Solidity: function getStakingReward(uint256 stakeTokenId) view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) GetStakingReward(stakeTokenId *big.Int) (*big.Int, error) {
	return _Enterprise.Contract.GetStakingReward(&_Enterprise.CallOpts, stakeTokenId)
}

// GetStreamingReserveHalvingPeriod is a free data retrieval call binding the contract method 0xdd9919e7.
//
// Solidity: function getStreamingReserveHalvingPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCaller) GetStreamingReserveHalvingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getStreamingReserveHalvingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStreamingReserveHalvingPeriod is a free data retrieval call binding the contract method 0xdd9919e7.
//
// Solidity: function getStreamingReserveHalvingPeriod() view returns(uint32)
func (_Enterprise *EnterpriseSession) GetStreamingReserveHalvingPeriod() (uint32, error) {
	return _Enterprise.Contract.GetStreamingReserveHalvingPeriod(&_Enterprise.CallOpts)
}

// GetStreamingReserveHalvingPeriod is a free data retrieval call binding the contract method 0xdd9919e7.
//
// Solidity: function getStreamingReserveHalvingPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCallerSession) GetStreamingReserveHalvingPeriod() (uint32, error) {
	return _Enterprise.Contract.GetStreamingReserveHalvingPeriod(&_Enterprise.CallOpts)
}

// GetUsedReserve is a free data retrieval call binding the contract method 0xfcb2884e.
//
// Solidity: function getUsedReserve() view returns(uint256)
func (_Enterprise *EnterpriseCaller) GetUsedReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getUsedReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsedReserve is a free data retrieval call binding the contract method 0xfcb2884e.
//
// Solidity: function getUsedReserve() view returns(uint256)
func (_Enterprise *EnterpriseSession) GetUsedReserve() (*big.Int, error) {
	return _Enterprise.Contract.GetUsedReserve(&_Enterprise.CallOpts)
}

// GetUsedReserve is a free data retrieval call binding the contract method 0xfcb2884e.
//
// Solidity: function getUsedReserve() view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) GetUsedReserve() (*big.Int, error) {
	return _Enterprise.Contract.GetUsedReserve(&_Enterprise.CallOpts)
}

// IsRegisteredPowerToken is a free data retrieval call binding the contract method 0x89035b61.
//
// Solidity: function isRegisteredPowerToken(address powerToken) view returns(bool)
func (_Enterprise *EnterpriseCaller) IsRegisteredPowerToken(opts *bind.CallOpts, powerToken common.Address) (bool, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "isRegisteredPowerToken", powerToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredPowerToken is a free data retrieval call binding the contract method 0x89035b61.
//
// Solidity: function isRegisteredPowerToken(address powerToken) view returns(bool)
func (_Enterprise *EnterpriseSession) IsRegisteredPowerToken(powerToken common.Address) (bool, error) {
	return _Enterprise.Contract.IsRegisteredPowerToken(&_Enterprise.CallOpts, powerToken)
}

// IsRegisteredPowerToken is a free data retrieval call binding the contract method 0x89035b61.
//
// Solidity: function isRegisteredPowerToken(address powerToken) view returns(bool)
func (_Enterprise *EnterpriseCallerSession) IsRegisteredPowerToken(powerToken common.Address) (bool, error) {
	return _Enterprise.Contract.IsRegisteredPowerToken(&_Enterprise.CallOpts, powerToken)
}

// IsSupportedPaymentToken is a free data retrieval call binding the contract method 0x2fb2067f.
//
// Solidity: function isSupportedPaymentToken(address token) view returns(bool)
func (_Enterprise *EnterpriseCaller) IsSupportedPaymentToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "isSupportedPaymentToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedPaymentToken is a free data retrieval call binding the contract method 0x2fb2067f.
//
// Solidity: function isSupportedPaymentToken(address token) view returns(bool)
func (_Enterprise *EnterpriseSession) IsSupportedPaymentToken(token common.Address) (bool, error) {
	return _Enterprise.Contract.IsSupportedPaymentToken(&_Enterprise.CallOpts, token)
}

// IsSupportedPaymentToken is a free data retrieval call binding the contract method 0x2fb2067f.
//
// Solidity: function isSupportedPaymentToken(address token) view returns(bool)
func (_Enterprise *EnterpriseCallerSession) IsSupportedPaymentToken(token common.Address) (bool, error) {
	return _Enterprise.Contract.IsSupportedPaymentToken(&_Enterprise.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Enterprise *EnterpriseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Enterprise *EnterpriseSession) Owner() (common.Address, error) {
	return _Enterprise.Contract.Owner(&_Enterprise.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Enterprise *EnterpriseCallerSession) Owner() (common.Address, error) {
	return _Enterprise.Contract.Owner(&_Enterprise.CallOpts)
}

// ClaimStakingReward is a paid mutator transaction binding the contract method 0x239cd4a4.
//
// Solidity: function claimStakingReward(uint256 stakeTokenId) returns()
func (_Enterprise *EnterpriseTransactor) ClaimStakingReward(opts *bind.TransactOpts, stakeTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "claimStakingReward", stakeTokenId)
}

// ClaimStakingReward is a paid mutator transaction binding the contract method 0x239cd4a4.
//
// Solidity: function claimStakingReward(uint256 stakeTokenId) returns()
func (_Enterprise *EnterpriseSession) ClaimStakingReward(stakeTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ClaimStakingReward(&_Enterprise.TransactOpts, stakeTokenId)
}

// ClaimStakingReward is a paid mutator transaction binding the contract method 0x239cd4a4.
//
// Solidity: function claimStakingReward(uint256 stakeTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) ClaimStakingReward(stakeTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ClaimStakingReward(&_Enterprise.TransactOpts, stakeTokenId)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0x24d86f00.
//
// Solidity: function decreaseStake(uint256 stakeTokenId, uint256 stakeAmountDelta) returns()
func (_Enterprise *EnterpriseTransactor) DecreaseStake(opts *bind.TransactOpts, stakeTokenId *big.Int, stakeAmountDelta *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "decreaseStake", stakeTokenId, stakeAmountDelta)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0x24d86f00.
//
// Solidity: function decreaseStake(uint256 stakeTokenId, uint256 stakeAmountDelta) returns()
func (_Enterprise *EnterpriseSession) DecreaseStake(stakeTokenId *big.Int, stakeAmountDelta *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.DecreaseStake(&_Enterprise.TransactOpts, stakeTokenId, stakeAmountDelta)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0x24d86f00.
//
// Solidity: function decreaseStake(uint256 stakeTokenId, uint256 stakeAmountDelta) returns()
func (_Enterprise *EnterpriseTransactorSession) DecreaseStake(stakeTokenId *big.Int, stakeAmountDelta *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.DecreaseStake(&_Enterprise.TransactOpts, stakeTokenId, stakeAmountDelta)
}

// DisablePaymentToken is a paid mutator transaction binding the contract method 0x9c7aa7f8.
//
// Solidity: function disablePaymentToken(address token) returns()
func (_Enterprise *EnterpriseTransactor) DisablePaymentToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "disablePaymentToken", token)
}

// DisablePaymentToken is a paid mutator transaction binding the contract method 0x9c7aa7f8.
//
// Solidity: function disablePaymentToken(address token) returns()
func (_Enterprise *EnterpriseSession) DisablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.DisablePaymentToken(&_Enterprise.TransactOpts, token)
}

// DisablePaymentToken is a paid mutator transaction binding the contract method 0x9c7aa7f8.
//
// Solidity: function disablePaymentToken(address token) returns()
func (_Enterprise *EnterpriseTransactorSession) DisablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.DisablePaymentToken(&_Enterprise.TransactOpts, token)
}

// EnablePaymentToken is a paid mutator transaction binding the contract method 0xbfd84fb4.
//
// Solidity: function enablePaymentToken(address token) returns()
func (_Enterprise *EnterpriseTransactor) EnablePaymentToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "enablePaymentToken", token)
}

// EnablePaymentToken is a paid mutator transaction binding the contract method 0xbfd84fb4.
//
// Solidity: function enablePaymentToken(address token) returns()
func (_Enterprise *EnterpriseSession) EnablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.EnablePaymentToken(&_Enterprise.TransactOpts, token)
}

// EnablePaymentToken is a paid mutator transaction binding the contract method 0xbfd84fb4.
//
// Solidity: function enablePaymentToken(address token) returns()
func (_Enterprise *EnterpriseTransactorSession) EnablePaymentToken(token common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.EnablePaymentToken(&_Enterprise.TransactOpts, token)
}

// ExtendRentalPeriod is a paid mutator transaction binding the contract method 0xc852d200.
//
// Solidity: function extendRentalPeriod(uint256 rentalTokenId, address paymentToken, uint32 rentalPeriod, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactor) ExtendRentalPeriod(opts *bind.TransactOpts, rentalTokenId *big.Int, paymentToken common.Address, rentalPeriod uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "extendRentalPeriod", rentalTokenId, paymentToken, rentalPeriod, maxPayment)
}

// ExtendRentalPeriod is a paid mutator transaction binding the contract method 0xc852d200.
//
// Solidity: function extendRentalPeriod(uint256 rentalTokenId, address paymentToken, uint32 rentalPeriod, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseSession) ExtendRentalPeriod(rentalTokenId *big.Int, paymentToken common.Address, rentalPeriod uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ExtendRentalPeriod(&_Enterprise.TransactOpts, rentalTokenId, paymentToken, rentalPeriod, maxPayment)
}

// ExtendRentalPeriod is a paid mutator transaction binding the contract method 0xc852d200.
//
// Solidity: function extendRentalPeriod(uint256 rentalTokenId, address paymentToken, uint32 rentalPeriod, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactorSession) ExtendRentalPeriod(rentalTokenId *big.Int, paymentToken common.Address, rentalPeriod uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ExtendRentalPeriod(&_Enterprise.TransactOpts, rentalTokenId, paymentToken, rentalPeriod, maxPayment)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xbec10cde.
//
// Solidity: function increaseStake(uint256 stakeTokenId, uint256 stakeAmountDelta) returns()
func (_Enterprise *EnterpriseTransactor) IncreaseStake(opts *bind.TransactOpts, stakeTokenId *big.Int, stakeAmountDelta *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "increaseStake", stakeTokenId, stakeAmountDelta)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xbec10cde.
//
// Solidity: function increaseStake(uint256 stakeTokenId, uint256 stakeAmountDelta) returns()
func (_Enterprise *EnterpriseSession) IncreaseStake(stakeTokenId *big.Int, stakeAmountDelta *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.IncreaseStake(&_Enterprise.TransactOpts, stakeTokenId, stakeAmountDelta)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xbec10cde.
//
// Solidity: function increaseStake(uint256 stakeTokenId, uint256 stakeAmountDelta) returns()
func (_Enterprise *EnterpriseTransactorSession) IncreaseStake(stakeTokenId *big.Int, stakeAmountDelta *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.IncreaseStake(&_Enterprise.TransactOpts, stakeTokenId, stakeAmountDelta)
}

// Initialize is a paid mutator transaction binding the contract method 0x6815f337.
//
// Solidity: function initialize(string enterpriseName, string baseUri, uint16 gcFeePercent, address converter, address proxyAdmin, address owner) returns()
func (_Enterprise *EnterpriseTransactor) Initialize(opts *bind.TransactOpts, enterpriseName string, baseUri string, gcFeePercent uint16, converter common.Address, proxyAdmin common.Address, owner common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "initialize", enterpriseName, baseUri, gcFeePercent, converter, proxyAdmin, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x6815f337.
//
// Solidity: function initialize(string enterpriseName, string baseUri, uint16 gcFeePercent, address converter, address proxyAdmin, address owner) returns()
func (_Enterprise *EnterpriseSession) Initialize(enterpriseName string, baseUri string, gcFeePercent uint16, converter common.Address, proxyAdmin common.Address, owner common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.Initialize(&_Enterprise.TransactOpts, enterpriseName, baseUri, gcFeePercent, converter, proxyAdmin, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x6815f337.
//
// Solidity: function initialize(string enterpriseName, string baseUri, uint16 gcFeePercent, address converter, address proxyAdmin, address owner) returns()
func (_Enterprise *EnterpriseTransactorSession) Initialize(enterpriseName string, baseUri string, gcFeePercent uint16, converter common.Address, proxyAdmin common.Address, owner common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.Initialize(&_Enterprise.TransactOpts, enterpriseName, baseUri, gcFeePercent, converter, proxyAdmin, owner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Enterprise *EnterpriseTransactor) Initialize0(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "initialize0", initialOwner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Enterprise *EnterpriseSession) Initialize0(initialOwner common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.Initialize0(&_Enterprise.TransactOpts, initialOwner)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Enterprise *EnterpriseTransactorSession) Initialize0(initialOwner common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.Initialize0(&_Enterprise.TransactOpts, initialOwner)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address enterpriseToken, address stakeToken, address rentalToken) returns()
func (_Enterprise *EnterpriseTransactor) InitializeTokens(opts *bind.TransactOpts, enterpriseToken common.Address, stakeToken common.Address, rentalToken common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "initializeTokens", enterpriseToken, stakeToken, rentalToken)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address enterpriseToken, address stakeToken, address rentalToken) returns()
func (_Enterprise *EnterpriseSession) InitializeTokens(enterpriseToken common.Address, stakeToken common.Address, rentalToken common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.InitializeTokens(&_Enterprise.TransactOpts, enterpriseToken, stakeToken, rentalToken)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address enterpriseToken, address stakeToken, address rentalToken) returns()
func (_Enterprise *EnterpriseTransactorSession) InitializeTokens(enterpriseToken common.Address, stakeToken common.Address, rentalToken common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.InitializeTokens(&_Enterprise.TransactOpts, enterpriseToken, stakeToken, rentalToken)
}

// RegisterService is a paid mutator transaction binding the contract method 0x9ab71193.
//
// Solidity: function registerService(string serviceName, string serviceSymbol, uint32 energyGapHalvingPeriod, uint112 baseRate, address baseToken, uint16 serviceFeePercent, uint32 minRentalPeriod, uint32 maxRentalPeriod, uint96 minGCFee, bool swappingEnabledForever) returns()
func (_Enterprise *EnterpriseTransactor) RegisterService(opts *bind.TransactOpts, serviceName string, serviceSymbol string, energyGapHalvingPeriod uint32, baseRate *big.Int, baseToken common.Address, serviceFeePercent uint16, minRentalPeriod uint32, maxRentalPeriod uint32, minGCFee *big.Int, swappingEnabledForever bool) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "registerService", serviceName, serviceSymbol, energyGapHalvingPeriod, baseRate, baseToken, serviceFeePercent, minRentalPeriod, maxRentalPeriod, minGCFee, swappingEnabledForever)
}

// RegisterService is a paid mutator transaction binding the contract method 0x9ab71193.
//
// Solidity: function registerService(string serviceName, string serviceSymbol, uint32 energyGapHalvingPeriod, uint112 baseRate, address baseToken, uint16 serviceFeePercent, uint32 minRentalPeriod, uint32 maxRentalPeriod, uint96 minGCFee, bool swappingEnabledForever) returns()
func (_Enterprise *EnterpriseSession) RegisterService(serviceName string, serviceSymbol string, energyGapHalvingPeriod uint32, baseRate *big.Int, baseToken common.Address, serviceFeePercent uint16, minRentalPeriod uint32, maxRentalPeriod uint32, minGCFee *big.Int, swappingEnabledForever bool) (*types.Transaction, error) {
	return _Enterprise.Contract.RegisterService(&_Enterprise.TransactOpts, serviceName, serviceSymbol, energyGapHalvingPeriod, baseRate, baseToken, serviceFeePercent, minRentalPeriod, maxRentalPeriod, minGCFee, swappingEnabledForever)
}

// RegisterService is a paid mutator transaction binding the contract method 0x9ab71193.
//
// Solidity: function registerService(string serviceName, string serviceSymbol, uint32 energyGapHalvingPeriod, uint112 baseRate, address baseToken, uint16 serviceFeePercent, uint32 minRentalPeriod, uint32 maxRentalPeriod, uint96 minGCFee, bool swappingEnabledForever) returns()
func (_Enterprise *EnterpriseTransactorSession) RegisterService(serviceName string, serviceSymbol string, energyGapHalvingPeriod uint32, baseRate *big.Int, baseToken common.Address, serviceFeePercent uint16, minRentalPeriod uint32, maxRentalPeriod uint32, minGCFee *big.Int, swappingEnabledForever bool) (*types.Transaction, error) {
	return _Enterprise.Contract.RegisterService(&_Enterprise.TransactOpts, serviceName, serviceSymbol, energyGapHalvingPeriod, baseRate, baseToken, serviceFeePercent, minRentalPeriod, maxRentalPeriod, minGCFee, swappingEnabledForever)
}

// Rent is a paid mutator transaction binding the contract method 0x8fc49ad2.
//
// Solidity: function rent(address powerToken, address paymentToken, uint112 rentalAmount, uint32 rentalPeriod, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactor) Rent(opts *bind.TransactOpts, powerToken common.Address, paymentToken common.Address, rentalAmount *big.Int, rentalPeriod uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "rent", powerToken, paymentToken, rentalAmount, rentalPeriod, maxPayment)
}

// Rent is a paid mutator transaction binding the contract method 0x8fc49ad2.
//
// Solidity: function rent(address powerToken, address paymentToken, uint112 rentalAmount, uint32 rentalPeriod, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseSession) Rent(powerToken common.Address, paymentToken common.Address, rentalAmount *big.Int, rentalPeriod uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Rent(&_Enterprise.TransactOpts, powerToken, paymentToken, rentalAmount, rentalPeriod, maxPayment)
}

// Rent is a paid mutator transaction binding the contract method 0x8fc49ad2.
//
// Solidity: function rent(address powerToken, address paymentToken, uint112 rentalAmount, uint32 rentalPeriod, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactorSession) Rent(powerToken common.Address, paymentToken common.Address, rentalAmount *big.Int, rentalPeriod uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Rent(&_Enterprise.TransactOpts, powerToken, paymentToken, rentalAmount, rentalPeriod, maxPayment)
}

// ReturnRental is a paid mutator transaction binding the contract method 0x3310df9e.
//
// Solidity: function returnRental(uint256 rentalTokenId) returns()
func (_Enterprise *EnterpriseTransactor) ReturnRental(opts *bind.TransactOpts, rentalTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "returnRental", rentalTokenId)
}

// ReturnRental is a paid mutator transaction binding the contract method 0x3310df9e.
//
// Solidity: function returnRental(uint256 rentalTokenId) returns()
func (_Enterprise *EnterpriseSession) ReturnRental(rentalTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ReturnRental(&_Enterprise.TransactOpts, rentalTokenId)
}

// ReturnRental is a paid mutator transaction binding the contract method 0x3310df9e.
//
// Solidity: function returnRental(uint256 rentalTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) ReturnRental(rentalTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ReturnRental(&_Enterprise.TransactOpts, rentalTokenId)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri) returns()
func (_Enterprise *EnterpriseTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri string) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setBaseUri", baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri) returns()
func (_Enterprise *EnterpriseSession) SetBaseUri(baseUri string) (*types.Transaction, error) {
	return _Enterprise.Contract.SetBaseUri(&_Enterprise.TransactOpts, baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri) returns()
func (_Enterprise *EnterpriseTransactorSession) SetBaseUri(baseUri string) (*types.Transaction, error) {
	return _Enterprise.Contract.SetBaseUri(&_Enterprise.TransactOpts, baseUri)
}

// SetBondingCurve is a paid mutator transaction binding the contract method 0x3513e0dc.
//
// Solidity: function setBondingCurve(uint256 pole, uint256 slope) returns()
func (_Enterprise *EnterpriseTransactor) SetBondingCurve(opts *bind.TransactOpts, pole *big.Int, slope *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setBondingCurve", pole, slope)
}

// SetBondingCurve is a paid mutator transaction binding the contract method 0x3513e0dc.
//
// Solidity: function setBondingCurve(uint256 pole, uint256 slope) returns()
func (_Enterprise *EnterpriseSession) SetBondingCurve(pole *big.Int, slope *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.SetBondingCurve(&_Enterprise.TransactOpts, pole, slope)
}

// SetBondingCurve is a paid mutator transaction binding the contract method 0x3513e0dc.
//
// Solidity: function setBondingCurve(uint256 pole, uint256 slope) returns()
func (_Enterprise *EnterpriseTransactorSession) SetBondingCurve(pole *big.Int, slope *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.SetBondingCurve(&_Enterprise.TransactOpts, pole, slope)
}

// SetConverter is a paid mutator transaction binding the contract method 0xb19337a4.
//
// Solidity: function setConverter(address newConverter) returns()
func (_Enterprise *EnterpriseTransactor) SetConverter(opts *bind.TransactOpts, newConverter common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setConverter", newConverter)
}

// SetConverter is a paid mutator transaction binding the contract method 0xb19337a4.
//
// Solidity: function setConverter(address newConverter) returns()
func (_Enterprise *EnterpriseSession) SetConverter(newConverter common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetConverter(&_Enterprise.TransactOpts, newConverter)
}

// SetConverter is a paid mutator transaction binding the contract method 0xb19337a4.
//
// Solidity: function setConverter(address newConverter) returns()
func (_Enterprise *EnterpriseTransactorSession) SetConverter(newConverter common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetConverter(&_Enterprise.TransactOpts, newConverter)
}

// SetEnterpriseCollector is a paid mutator transaction binding the contract method 0x62c1f388.
//
// Solidity: function setEnterpriseCollector(address newCollector) returns()
func (_Enterprise *EnterpriseTransactor) SetEnterpriseCollector(opts *bind.TransactOpts, newCollector common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setEnterpriseCollector", newCollector)
}

// SetEnterpriseCollector is a paid mutator transaction binding the contract method 0x62c1f388.
//
// Solidity: function setEnterpriseCollector(address newCollector) returns()
func (_Enterprise *EnterpriseSession) SetEnterpriseCollector(newCollector common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseCollector(&_Enterprise.TransactOpts, newCollector)
}

// SetEnterpriseCollector is a paid mutator transaction binding the contract method 0x62c1f388.
//
// Solidity: function setEnterpriseCollector(address newCollector) returns()
func (_Enterprise *EnterpriseTransactorSession) SetEnterpriseCollector(newCollector common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseCollector(&_Enterprise.TransactOpts, newCollector)
}

// SetEnterpriseOnlyCollectionPeriod is a paid mutator transaction binding the contract method 0xc9a304cb.
//
// Solidity: function setEnterpriseOnlyCollectionPeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactor) SetEnterpriseOnlyCollectionPeriod(opts *bind.TransactOpts, newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setEnterpriseOnlyCollectionPeriod", newPeriod)
}

// SetEnterpriseOnlyCollectionPeriod is a paid mutator transaction binding the contract method 0xc9a304cb.
//
// Solidity: function setEnterpriseOnlyCollectionPeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseSession) SetEnterpriseOnlyCollectionPeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseOnlyCollectionPeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetEnterpriseOnlyCollectionPeriod is a paid mutator transaction binding the contract method 0xc9a304cb.
//
// Solidity: function setEnterpriseOnlyCollectionPeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactorSession) SetEnterpriseOnlyCollectionPeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseOnlyCollectionPeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetEnterpriseWallet is a paid mutator transaction binding the contract method 0x2107730c.
//
// Solidity: function setEnterpriseWallet(address newWallet) returns()
func (_Enterprise *EnterpriseTransactor) SetEnterpriseWallet(opts *bind.TransactOpts, newWallet common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setEnterpriseWallet", newWallet)
}

// SetEnterpriseWallet is a paid mutator transaction binding the contract method 0x2107730c.
//
// Solidity: function setEnterpriseWallet(address newWallet) returns()
func (_Enterprise *EnterpriseSession) SetEnterpriseWallet(newWallet common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseWallet(&_Enterprise.TransactOpts, newWallet)
}

// SetEnterpriseWallet is a paid mutator transaction binding the contract method 0x2107730c.
//
// Solidity: function setEnterpriseWallet(address newWallet) returns()
func (_Enterprise *EnterpriseTransactorSession) SetEnterpriseWallet(newWallet common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseWallet(&_Enterprise.TransactOpts, newWallet)
}

// SetGcFeePercent is a paid mutator transaction binding the contract method 0x449497ec.
//
// Solidity: function setGcFeePercent(uint16 newGcFeePercent) returns()
func (_Enterprise *EnterpriseTransactor) SetGcFeePercent(opts *bind.TransactOpts, newGcFeePercent uint16) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setGcFeePercent", newGcFeePercent)
}

// SetGcFeePercent is a paid mutator transaction binding the contract method 0x449497ec.
//
// Solidity: function setGcFeePercent(uint16 newGcFeePercent) returns()
func (_Enterprise *EnterpriseSession) SetGcFeePercent(newGcFeePercent uint16) (*types.Transaction, error) {
	return _Enterprise.Contract.SetGcFeePercent(&_Enterprise.TransactOpts, newGcFeePercent)
}

// SetGcFeePercent is a paid mutator transaction binding the contract method 0x449497ec.
//
// Solidity: function setGcFeePercent(uint16 newGcFeePercent) returns()
func (_Enterprise *EnterpriseTransactorSession) SetGcFeePercent(newGcFeePercent uint16) (*types.Transaction, error) {
	return _Enterprise.Contract.SetGcFeePercent(&_Enterprise.TransactOpts, newGcFeePercent)
}

// SetRenterOnlyReturnPeriod is a paid mutator transaction binding the contract method 0x7c14d991.
//
// Solidity: function setRenterOnlyReturnPeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactor) SetRenterOnlyReturnPeriod(opts *bind.TransactOpts, newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setRenterOnlyReturnPeriod", newPeriod)
}

// SetRenterOnlyReturnPeriod is a paid mutator transaction binding the contract method 0x7c14d991.
//
// Solidity: function setRenterOnlyReturnPeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseSession) SetRenterOnlyReturnPeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetRenterOnlyReturnPeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetRenterOnlyReturnPeriod is a paid mutator transaction binding the contract method 0x7c14d991.
//
// Solidity: function setRenterOnlyReturnPeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactorSession) SetRenterOnlyReturnPeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetRenterOnlyReturnPeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetStreamingReserveHalvingPeriod is a paid mutator transaction binding the contract method 0xf87c4261.
//
// Solidity: function setStreamingReserveHalvingPeriod(uint32 streamingReserveHalvingPeriod) returns()
func (_Enterprise *EnterpriseTransactor) SetStreamingReserveHalvingPeriod(opts *bind.TransactOpts, streamingReserveHalvingPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setStreamingReserveHalvingPeriod", streamingReserveHalvingPeriod)
}

// SetStreamingReserveHalvingPeriod is a paid mutator transaction binding the contract method 0xf87c4261.
//
// Solidity: function setStreamingReserveHalvingPeriod(uint32 streamingReserveHalvingPeriod) returns()
func (_Enterprise *EnterpriseSession) SetStreamingReserveHalvingPeriod(streamingReserveHalvingPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetStreamingReserveHalvingPeriod(&_Enterprise.TransactOpts, streamingReserveHalvingPeriod)
}

// SetStreamingReserveHalvingPeriod is a paid mutator transaction binding the contract method 0xf87c4261.
//
// Solidity: function setStreamingReserveHalvingPeriod(uint32 streamingReserveHalvingPeriod) returns()
func (_Enterprise *EnterpriseTransactorSession) SetStreamingReserveHalvingPeriod(streamingReserveHalvingPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetStreamingReserveHalvingPeriod(&_Enterprise.TransactOpts, streamingReserveHalvingPeriod)
}

// ShutdownEnterpriseForever is a paid mutator transaction binding the contract method 0x4741efb3.
//
// Solidity: function shutdownEnterpriseForever() returns()
func (_Enterprise *EnterpriseTransactor) ShutdownEnterpriseForever(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "shutdownEnterpriseForever")
}

// ShutdownEnterpriseForever is a paid mutator transaction binding the contract method 0x4741efb3.
//
// Solidity: function shutdownEnterpriseForever() returns()
func (_Enterprise *EnterpriseSession) ShutdownEnterpriseForever() (*types.Transaction, error) {
	return _Enterprise.Contract.ShutdownEnterpriseForever(&_Enterprise.TransactOpts)
}

// ShutdownEnterpriseForever is a paid mutator transaction binding the contract method 0x4741efb3.
//
// Solidity: function shutdownEnterpriseForever() returns()
func (_Enterprise *EnterpriseTransactorSession) ShutdownEnterpriseForever() (*types.Transaction, error) {
	return _Enterprise.Contract.ShutdownEnterpriseForever(&_Enterprise.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 stakeAmount) returns()
func (_Enterprise *EnterpriseTransactor) Stake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "stake", stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 stakeAmount) returns()
func (_Enterprise *EnterpriseSession) Stake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Stake(&_Enterprise.TransactOpts, stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 stakeAmount) returns()
func (_Enterprise *EnterpriseTransactorSession) Stake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Stake(&_Enterprise.TransactOpts, stakeAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Enterprise *EnterpriseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Enterprise *EnterpriseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.TransferOwnership(&_Enterprise.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Enterprise *EnterpriseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.TransferOwnership(&_Enterprise.TransactOpts, newOwner)
}

// TransferRental is a paid mutator transaction binding the contract method 0x960970c7.
//
// Solidity: function transferRental(address from, address to, uint256 rentalTokenId) returns()
func (_Enterprise *EnterpriseTransactor) TransferRental(opts *bind.TransactOpts, from common.Address, to common.Address, rentalTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "transferRental", from, to, rentalTokenId)
}

// TransferRental is a paid mutator transaction binding the contract method 0x960970c7.
//
// Solidity: function transferRental(address from, address to, uint256 rentalTokenId) returns()
func (_Enterprise *EnterpriseSession) TransferRental(from common.Address, to common.Address, rentalTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.TransferRental(&_Enterprise.TransactOpts, from, to, rentalTokenId)
}

// TransferRental is a paid mutator transaction binding the contract method 0x960970c7.
//
// Solidity: function transferRental(address from, address to, uint256 rentalTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) TransferRental(from common.Address, to common.Address, rentalTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.TransferRental(&_Enterprise.TransactOpts, from, to, rentalTokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeTokenId) returns()
func (_Enterprise *EnterpriseTransactor) Unstake(opts *bind.TransactOpts, stakeTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "unstake", stakeTokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeTokenId) returns()
func (_Enterprise *EnterpriseSession) Unstake(stakeTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Unstake(&_Enterprise.TransactOpts, stakeTokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) Unstake(stakeTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Unstake(&_Enterprise.TransactOpts, stakeTokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x6df0bb14.
//
// Solidity: function upgrade(address enterpriseFactory, address enterpriseImplementation, address rentalTokenImplementation, address stakeTokenImplementation, address powerTokenImplementation, address[] powerTokens) returns()
func (_Enterprise *EnterpriseTransactor) Upgrade(opts *bind.TransactOpts, enterpriseFactory common.Address, enterpriseImplementation common.Address, rentalTokenImplementation common.Address, stakeTokenImplementation common.Address, powerTokenImplementation common.Address, powerTokens []common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "upgrade", enterpriseFactory, enterpriseImplementation, rentalTokenImplementation, stakeTokenImplementation, powerTokenImplementation, powerTokens)
}

// Upgrade is a paid mutator transaction binding the contract method 0x6df0bb14.
//
// Solidity: function upgrade(address enterpriseFactory, address enterpriseImplementation, address rentalTokenImplementation, address stakeTokenImplementation, address powerTokenImplementation, address[] powerTokens) returns()
func (_Enterprise *EnterpriseSession) Upgrade(enterpriseFactory common.Address, enterpriseImplementation common.Address, rentalTokenImplementation common.Address, stakeTokenImplementation common.Address, powerTokenImplementation common.Address, powerTokens []common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.Upgrade(&_Enterprise.TransactOpts, enterpriseFactory, enterpriseImplementation, rentalTokenImplementation, stakeTokenImplementation, powerTokenImplementation, powerTokens)
}

// Upgrade is a paid mutator transaction binding the contract method 0x6df0bb14.
//
// Solidity: function upgrade(address enterpriseFactory, address enterpriseImplementation, address rentalTokenImplementation, address stakeTokenImplementation, address powerTokenImplementation, address[] powerTokens) returns()
func (_Enterprise *EnterpriseTransactorSession) Upgrade(enterpriseFactory common.Address, enterpriseImplementation common.Address, rentalTokenImplementation common.Address, stakeTokenImplementation common.Address, powerTokenImplementation common.Address, powerTokens []common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.Upgrade(&_Enterprise.TransactOpts, enterpriseFactory, enterpriseImplementation, rentalTokenImplementation, stakeTokenImplementation, powerTokenImplementation, powerTokens)
}

// EnterpriseBaseUriChangedIterator is returned from FilterBaseUriChanged and is used to iterate over the raw logs and unpacked data for BaseUriChanged events raised by the Enterprise contract.
type EnterpriseBaseUriChangedIterator struct {
	Event *EnterpriseBaseUriChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseBaseUriChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseBaseUriChanged)
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
		it.Event = new(EnterpriseBaseUriChanged)
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
func (it *EnterpriseBaseUriChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseBaseUriChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseBaseUriChanged represents a BaseUriChanged event raised by the Enterprise contract.
type EnterpriseBaseUriChanged struct {
	BaseUri string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseUriChanged is a free log retrieval operation binding the contract event 0x87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6.
//
// Solidity: event BaseUriChanged(string baseUri)
func (_Enterprise *EnterpriseFilterer) FilterBaseUriChanged(opts *bind.FilterOpts) (*EnterpriseBaseUriChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "BaseUriChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseBaseUriChangedIterator{contract: _Enterprise.contract, event: "BaseUriChanged", logs: logs, sub: sub}, nil
}

// WatchBaseUriChanged is a free log subscription operation binding the contract event 0x87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6.
//
// Solidity: event BaseUriChanged(string baseUri)
func (_Enterprise *EnterpriseFilterer) WatchBaseUriChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseBaseUriChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "BaseUriChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseBaseUriChanged)
				if err := _Enterprise.contract.UnpackLog(event, "BaseUriChanged", log); err != nil {
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

// ParseBaseUriChanged is a log parse operation binding the contract event 0x87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6.
//
// Solidity: event BaseUriChanged(string baseUri)
func (_Enterprise *EnterpriseFilterer) ParseBaseUriChanged(log types.Log) (*EnterpriseBaseUriChanged, error) {
	event := new(EnterpriseBaseUriChanged)
	if err := _Enterprise.contract.UnpackLog(event, "BaseUriChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseBondingChangedIterator is returned from FilterBondingChanged and is used to iterate over the raw logs and unpacked data for BondingChanged events raised by the Enterprise contract.
type EnterpriseBondingChangedIterator struct {
	Event *EnterpriseBondingChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseBondingChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseBondingChanged)
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
		it.Event = new(EnterpriseBondingChanged)
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
func (it *EnterpriseBondingChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseBondingChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseBondingChanged represents a BondingChanged event raised by the Enterprise contract.
type EnterpriseBondingChanged struct {
	Pole  *big.Int
	Slope *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBondingChanged is a free log retrieval operation binding the contract event 0x926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39f.
//
// Solidity: event BondingChanged(uint256 pole, uint256 slope)
func (_Enterprise *EnterpriseFilterer) FilterBondingChanged(opts *bind.FilterOpts) (*EnterpriseBondingChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "BondingChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseBondingChangedIterator{contract: _Enterprise.contract, event: "BondingChanged", logs: logs, sub: sub}, nil
}

// WatchBondingChanged is a free log subscription operation binding the contract event 0x926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39f.
//
// Solidity: event BondingChanged(uint256 pole, uint256 slope)
func (_Enterprise *EnterpriseFilterer) WatchBondingChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseBondingChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "BondingChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseBondingChanged)
				if err := _Enterprise.contract.UnpackLog(event, "BondingChanged", log); err != nil {
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

// ParseBondingChanged is a log parse operation binding the contract event 0x926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39f.
//
// Solidity: event BondingChanged(uint256 pole, uint256 slope)
func (_Enterprise *EnterpriseFilterer) ParseBondingChanged(log types.Log) (*EnterpriseBondingChanged, error) {
	event := new(EnterpriseBondingChanged)
	if err := _Enterprise.contract.UnpackLog(event, "BondingChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseConverterChangedIterator is returned from FilterConverterChanged and is used to iterate over the raw logs and unpacked data for ConverterChanged events raised by the Enterprise contract.
type EnterpriseConverterChangedIterator struct {
	Event *EnterpriseConverterChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseConverterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseConverterChanged)
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
		it.Event = new(EnterpriseConverterChanged)
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
func (it *EnterpriseConverterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseConverterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseConverterChanged represents a ConverterChanged event raised by the Enterprise contract.
type EnterpriseConverterChanged struct {
	Converter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConverterChanged is a free log retrieval operation binding the contract event 0xbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656df.
//
// Solidity: event ConverterChanged(address converter)
func (_Enterprise *EnterpriseFilterer) FilterConverterChanged(opts *bind.FilterOpts) (*EnterpriseConverterChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "ConverterChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseConverterChangedIterator{contract: _Enterprise.contract, event: "ConverterChanged", logs: logs, sub: sub}, nil
}

// WatchConverterChanged is a free log subscription operation binding the contract event 0xbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656df.
//
// Solidity: event ConverterChanged(address converter)
func (_Enterprise *EnterpriseFilterer) WatchConverterChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseConverterChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "ConverterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseConverterChanged)
				if err := _Enterprise.contract.UnpackLog(event, "ConverterChanged", log); err != nil {
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

// ParseConverterChanged is a log parse operation binding the contract event 0xbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656df.
//
// Solidity: event ConverterChanged(address converter)
func (_Enterprise *EnterpriseFilterer) ParseConverterChanged(log types.Log) (*EnterpriseConverterChanged, error) {
	event := new(EnterpriseConverterChanged)
	if err := _Enterprise.contract.UnpackLog(event, "ConverterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseEnterpriseCollectorChangedIterator is returned from FilterEnterpriseCollectorChanged and is used to iterate over the raw logs and unpacked data for EnterpriseCollectorChanged events raised by the Enterprise contract.
type EnterpriseEnterpriseCollectorChangedIterator struct {
	Event *EnterpriseEnterpriseCollectorChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseEnterpriseCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseEnterpriseCollectorChanged)
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
		it.Event = new(EnterpriseEnterpriseCollectorChanged)
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
func (it *EnterpriseEnterpriseCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseEnterpriseCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseEnterpriseCollectorChanged represents a EnterpriseCollectorChanged event raised by the Enterprise contract.
type EnterpriseEnterpriseCollectorChanged struct {
	Collector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseCollectorChanged is a free log retrieval operation binding the contract event 0x7aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4.
//
// Solidity: event EnterpriseCollectorChanged(address collector)
func (_Enterprise *EnterpriseFilterer) FilterEnterpriseCollectorChanged(opts *bind.FilterOpts) (*EnterpriseEnterpriseCollectorChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "EnterpriseCollectorChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseEnterpriseCollectorChangedIterator{contract: _Enterprise.contract, event: "EnterpriseCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseCollectorChanged is a free log subscription operation binding the contract event 0x7aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4.
//
// Solidity: event EnterpriseCollectorChanged(address collector)
func (_Enterprise *EnterpriseFilterer) WatchEnterpriseCollectorChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseEnterpriseCollectorChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "EnterpriseCollectorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseEnterpriseCollectorChanged)
				if err := _Enterprise.contract.UnpackLog(event, "EnterpriseCollectorChanged", log); err != nil {
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

// ParseEnterpriseCollectorChanged is a log parse operation binding the contract event 0x7aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4.
//
// Solidity: event EnterpriseCollectorChanged(address collector)
func (_Enterprise *EnterpriseFilterer) ParseEnterpriseCollectorChanged(log types.Log) (*EnterpriseEnterpriseCollectorChanged, error) {
	event := new(EnterpriseEnterpriseCollectorChanged)
	if err := _Enterprise.contract.UnpackLog(event, "EnterpriseCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator is returned from FilterEnterpriseOnlyCollectionPeriodChanged and is used to iterate over the raw logs and unpacked data for EnterpriseOnlyCollectionPeriodChanged events raised by the Enterprise contract.
type EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator struct {
	Event *EnterpriseEnterpriseOnlyCollectionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseEnterpriseOnlyCollectionPeriodChanged)
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
		it.Event = new(EnterpriseEnterpriseOnlyCollectionPeriodChanged)
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
func (it *EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseEnterpriseOnlyCollectionPeriodChanged represents a EnterpriseOnlyCollectionPeriodChanged event raised by the Enterprise contract.
type EnterpriseEnterpriseOnlyCollectionPeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseOnlyCollectionPeriodChanged is a free log retrieval operation binding the contract event 0xd076b5bc77c447b04d82e76a12ec958bfe6d48418cce446aaf9ebeb3136638eb.
//
// Solidity: event EnterpriseOnlyCollectionPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) FilterEnterpriseOnlyCollectionPeriodChanged(opts *bind.FilterOpts) (*EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "EnterpriseOnlyCollectionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseEnterpriseOnlyCollectionPeriodChangedIterator{contract: _Enterprise.contract, event: "EnterpriseOnlyCollectionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseOnlyCollectionPeriodChanged is a free log subscription operation binding the contract event 0xd076b5bc77c447b04d82e76a12ec958bfe6d48418cce446aaf9ebeb3136638eb.
//
// Solidity: event EnterpriseOnlyCollectionPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) WatchEnterpriseOnlyCollectionPeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseEnterpriseOnlyCollectionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "EnterpriseOnlyCollectionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseEnterpriseOnlyCollectionPeriodChanged)
				if err := _Enterprise.contract.UnpackLog(event, "EnterpriseOnlyCollectionPeriodChanged", log); err != nil {
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

// ParseEnterpriseOnlyCollectionPeriodChanged is a log parse operation binding the contract event 0xd076b5bc77c447b04d82e76a12ec958bfe6d48418cce446aaf9ebeb3136638eb.
//
// Solidity: event EnterpriseOnlyCollectionPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) ParseEnterpriseOnlyCollectionPeriodChanged(log types.Log) (*EnterpriseEnterpriseOnlyCollectionPeriodChanged, error) {
	event := new(EnterpriseEnterpriseOnlyCollectionPeriodChanged)
	if err := _Enterprise.contract.UnpackLog(event, "EnterpriseOnlyCollectionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseEnterpriseShutdownIterator is returned from FilterEnterpriseShutdown and is used to iterate over the raw logs and unpacked data for EnterpriseShutdown events raised by the Enterprise contract.
type EnterpriseEnterpriseShutdownIterator struct {
	Event *EnterpriseEnterpriseShutdown // Event containing the contract specifics and raw log

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
func (it *EnterpriseEnterpriseShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseEnterpriseShutdown)
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
		it.Event = new(EnterpriseEnterpriseShutdown)
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
func (it *EnterpriseEnterpriseShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseEnterpriseShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseEnterpriseShutdown represents a EnterpriseShutdown event raised by the Enterprise contract.
type EnterpriseEnterpriseShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseShutdown is a free log retrieval operation binding the contract event 0x6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c072.
//
// Solidity: event EnterpriseShutdown()
func (_Enterprise *EnterpriseFilterer) FilterEnterpriseShutdown(opts *bind.FilterOpts) (*EnterpriseEnterpriseShutdownIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "EnterpriseShutdown")
	if err != nil {
		return nil, err
	}
	return &EnterpriseEnterpriseShutdownIterator{contract: _Enterprise.contract, event: "EnterpriseShutdown", logs: logs, sub: sub}, nil
}

// WatchEnterpriseShutdown is a free log subscription operation binding the contract event 0x6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c072.
//
// Solidity: event EnterpriseShutdown()
func (_Enterprise *EnterpriseFilterer) WatchEnterpriseShutdown(opts *bind.WatchOpts, sink chan<- *EnterpriseEnterpriseShutdown) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "EnterpriseShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseEnterpriseShutdown)
				if err := _Enterprise.contract.UnpackLog(event, "EnterpriseShutdown", log); err != nil {
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

// ParseEnterpriseShutdown is a log parse operation binding the contract event 0x6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c072.
//
// Solidity: event EnterpriseShutdown()
func (_Enterprise *EnterpriseFilterer) ParseEnterpriseShutdown(log types.Log) (*EnterpriseEnterpriseShutdown, error) {
	event := new(EnterpriseEnterpriseShutdown)
	if err := _Enterprise.contract.UnpackLog(event, "EnterpriseShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseEnterpriseWalletChangedIterator is returned from FilterEnterpriseWalletChanged and is used to iterate over the raw logs and unpacked data for EnterpriseWalletChanged events raised by the Enterprise contract.
type EnterpriseEnterpriseWalletChangedIterator struct {
	Event *EnterpriseEnterpriseWalletChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseEnterpriseWalletChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseEnterpriseWalletChanged)
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
		it.Event = new(EnterpriseEnterpriseWalletChanged)
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
func (it *EnterpriseEnterpriseWalletChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseEnterpriseWalletChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseEnterpriseWalletChanged represents a EnterpriseWalletChanged event raised by the Enterprise contract.
type EnterpriseEnterpriseWalletChanged struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseWalletChanged is a free log retrieval operation binding the contract event 0x471e342623158b92281274ef7263e4f9a0e0b748c1c328113afbb58742c1789f.
//
// Solidity: event EnterpriseWalletChanged(address wallet)
func (_Enterprise *EnterpriseFilterer) FilterEnterpriseWalletChanged(opts *bind.FilterOpts) (*EnterpriseEnterpriseWalletChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "EnterpriseWalletChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseEnterpriseWalletChangedIterator{contract: _Enterprise.contract, event: "EnterpriseWalletChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseWalletChanged is a free log subscription operation binding the contract event 0x471e342623158b92281274ef7263e4f9a0e0b748c1c328113afbb58742c1789f.
//
// Solidity: event EnterpriseWalletChanged(address wallet)
func (_Enterprise *EnterpriseFilterer) WatchEnterpriseWalletChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseEnterpriseWalletChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "EnterpriseWalletChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseEnterpriseWalletChanged)
				if err := _Enterprise.contract.UnpackLog(event, "EnterpriseWalletChanged", log); err != nil {
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

// ParseEnterpriseWalletChanged is a log parse operation binding the contract event 0x471e342623158b92281274ef7263e4f9a0e0b748c1c328113afbb58742c1789f.
//
// Solidity: event EnterpriseWalletChanged(address wallet)
func (_Enterprise *EnterpriseFilterer) ParseEnterpriseWalletChanged(log types.Log) (*EnterpriseEnterpriseWalletChanged, error) {
	event := new(EnterpriseEnterpriseWalletChanged)
	if err := _Enterprise.contract.UnpackLog(event, "EnterpriseWalletChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseFixedReserveChangedIterator is returned from FilterFixedReserveChanged and is used to iterate over the raw logs and unpacked data for FixedReserveChanged events raised by the Enterprise contract.
type EnterpriseFixedReserveChangedIterator struct {
	Event *EnterpriseFixedReserveChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseFixedReserveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseFixedReserveChanged)
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
		it.Event = new(EnterpriseFixedReserveChanged)
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
func (it *EnterpriseFixedReserveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseFixedReserveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseFixedReserveChanged represents a FixedReserveChanged event raised by the Enterprise contract.
type EnterpriseFixedReserveChanged struct {
	FixedReserve *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFixedReserveChanged is a free log retrieval operation binding the contract event 0xa70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180.
//
// Solidity: event FixedReserveChanged(uint256 fixedReserve)
func (_Enterprise *EnterpriseFilterer) FilterFixedReserveChanged(opts *bind.FilterOpts) (*EnterpriseFixedReserveChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "FixedReserveChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseFixedReserveChangedIterator{contract: _Enterprise.contract, event: "FixedReserveChanged", logs: logs, sub: sub}, nil
}

// WatchFixedReserveChanged is a free log subscription operation binding the contract event 0xa70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180.
//
// Solidity: event FixedReserveChanged(uint256 fixedReserve)
func (_Enterprise *EnterpriseFilterer) WatchFixedReserveChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseFixedReserveChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "FixedReserveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseFixedReserveChanged)
				if err := _Enterprise.contract.UnpackLog(event, "FixedReserveChanged", log); err != nil {
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

// ParseFixedReserveChanged is a log parse operation binding the contract event 0xa70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180.
//
// Solidity: event FixedReserveChanged(uint256 fixedReserve)
func (_Enterprise *EnterpriseFilterer) ParseFixedReserveChanged(log types.Log) (*EnterpriseFixedReserveChanged, error) {
	event := new(EnterpriseFixedReserveChanged)
	if err := _Enterprise.contract.UnpackLog(event, "FixedReserveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseGcFeePercentChangedIterator is returned from FilterGcFeePercentChanged and is used to iterate over the raw logs and unpacked data for GcFeePercentChanged events raised by the Enterprise contract.
type EnterpriseGcFeePercentChangedIterator struct {
	Event *EnterpriseGcFeePercentChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseGcFeePercentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseGcFeePercentChanged)
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
		it.Event = new(EnterpriseGcFeePercentChanged)
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
func (it *EnterpriseGcFeePercentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseGcFeePercentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseGcFeePercentChanged represents a GcFeePercentChanged event raised by the Enterprise contract.
type EnterpriseGcFeePercentChanged struct {
	Percent uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGcFeePercentChanged is a free log retrieval operation binding the contract event 0x27894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a15646.
//
// Solidity: event GcFeePercentChanged(uint16 percent)
func (_Enterprise *EnterpriseFilterer) FilterGcFeePercentChanged(opts *bind.FilterOpts) (*EnterpriseGcFeePercentChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "GcFeePercentChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseGcFeePercentChangedIterator{contract: _Enterprise.contract, event: "GcFeePercentChanged", logs: logs, sub: sub}, nil
}

// WatchGcFeePercentChanged is a free log subscription operation binding the contract event 0x27894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a15646.
//
// Solidity: event GcFeePercentChanged(uint16 percent)
func (_Enterprise *EnterpriseFilterer) WatchGcFeePercentChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseGcFeePercentChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "GcFeePercentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseGcFeePercentChanged)
				if err := _Enterprise.contract.UnpackLog(event, "GcFeePercentChanged", log); err != nil {
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

// ParseGcFeePercentChanged is a log parse operation binding the contract event 0x27894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a15646.
//
// Solidity: event GcFeePercentChanged(uint16 percent)
func (_Enterprise *EnterpriseFilterer) ParseGcFeePercentChanged(log types.Log) (*EnterpriseGcFeePercentChanged, error) {
	event := new(EnterpriseGcFeePercentChanged)
	if err := _Enterprise.contract.UnpackLog(event, "GcFeePercentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Enterprise contract.
type EnterpriseOwnershipTransferredIterator struct {
	Event *EnterpriseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnterpriseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseOwnershipTransferred)
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
		it.Event = new(EnterpriseOwnershipTransferred)
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
func (it *EnterpriseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseOwnershipTransferred represents a OwnershipTransferred event raised by the Enterprise contract.
type EnterpriseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Enterprise *EnterpriseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnterpriseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseOwnershipTransferredIterator{contract: _Enterprise.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Enterprise *EnterpriseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnterpriseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseOwnershipTransferred)
				if err := _Enterprise.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Enterprise *EnterpriseFilterer) ParseOwnershipTransferred(log types.Log) (*EnterpriseOwnershipTransferred, error) {
	event := new(EnterpriseOwnershipTransferred)
	if err := _Enterprise.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterprisePaymentTokenChangeIterator is returned from FilterPaymentTokenChange and is used to iterate over the raw logs and unpacked data for PaymentTokenChange events raised by the Enterprise contract.
type EnterprisePaymentTokenChangeIterator struct {
	Event *EnterprisePaymentTokenChange // Event containing the contract specifics and raw log

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
func (it *EnterprisePaymentTokenChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterprisePaymentTokenChange)
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
		it.Event = new(EnterprisePaymentTokenChange)
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
func (it *EnterprisePaymentTokenChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterprisePaymentTokenChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterprisePaymentTokenChange represents a PaymentTokenChange event raised by the Enterprise contract.
type EnterprisePaymentTokenChange struct {
	PaymentToken common.Address
	Enabled      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenChange is a free log retrieval operation binding the contract event 0x92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907.
//
// Solidity: event PaymentTokenChange(address paymentToken, bool enabled)
func (_Enterprise *EnterpriseFilterer) FilterPaymentTokenChange(opts *bind.FilterOpts) (*EnterprisePaymentTokenChangeIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "PaymentTokenChange")
	if err != nil {
		return nil, err
	}
	return &EnterprisePaymentTokenChangeIterator{contract: _Enterprise.contract, event: "PaymentTokenChange", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenChange is a free log subscription operation binding the contract event 0x92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907.
//
// Solidity: event PaymentTokenChange(address paymentToken, bool enabled)
func (_Enterprise *EnterpriseFilterer) WatchPaymentTokenChange(opts *bind.WatchOpts, sink chan<- *EnterprisePaymentTokenChange) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "PaymentTokenChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterprisePaymentTokenChange)
				if err := _Enterprise.contract.UnpackLog(event, "PaymentTokenChange", log); err != nil {
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

// ParsePaymentTokenChange is a log parse operation binding the contract event 0x92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907.
//
// Solidity: event PaymentTokenChange(address paymentToken, bool enabled)
func (_Enterprise *EnterpriseFilterer) ParsePaymentTokenChange(log types.Log) (*EnterprisePaymentTokenChange, error) {
	event := new(EnterprisePaymentTokenChange)
	if err := _Enterprise.contract.UnpackLog(event, "PaymentTokenChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseRentalPeriodExtendedIterator is returned from FilterRentalPeriodExtended and is used to iterate over the raw logs and unpacked data for RentalPeriodExtended events raised by the Enterprise contract.
type EnterpriseRentalPeriodExtendedIterator struct {
	Event *EnterpriseRentalPeriodExtended // Event containing the contract specifics and raw log

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
func (it *EnterpriseRentalPeriodExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseRentalPeriodExtended)
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
		it.Event = new(EnterpriseRentalPeriodExtended)
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
func (it *EnterpriseRentalPeriodExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseRentalPeriodExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseRentalPeriodExtended represents a RentalPeriodExtended event raised by the Enterprise contract.
type EnterpriseRentalPeriodExtended struct {
	RentalTokenId                *big.Int
	Renter                       common.Address
	PaymentToken                 common.Address
	PoolFee                      *big.Int
	ServiceFee                   *big.Int
	EndTime                      uint32
	RenterOnlyReturnTime         uint32
	EnterpriseOnlyCollectionTime uint32
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterRentalPeriodExtended is a free log retrieval operation binding the contract event 0x012ee01e680dee236f33567d88889822f7521f77a65adc92142abf5a2005d3ef.
//
// Solidity: event RentalPeriodExtended(uint256 indexed rentalTokenId, address indexed renter, address paymentToken, uint112 poolFee, uint112 serviceFee, uint32 endTime, uint32 renterOnlyReturnTime, uint32 enterpriseOnlyCollectionTime)
func (_Enterprise *EnterpriseFilterer) FilterRentalPeriodExtended(opts *bind.FilterOpts, rentalTokenId []*big.Int, renter []common.Address) (*EnterpriseRentalPeriodExtendedIterator, error) {

	var rentalTokenIdRule []interface{}
	for _, rentalTokenIdItem := range rentalTokenId {
		rentalTokenIdRule = append(rentalTokenIdRule, rentalTokenIdItem)
	}
	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "RentalPeriodExtended", rentalTokenIdRule, renterRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseRentalPeriodExtendedIterator{contract: _Enterprise.contract, event: "RentalPeriodExtended", logs: logs, sub: sub}, nil
}

// WatchRentalPeriodExtended is a free log subscription operation binding the contract event 0x012ee01e680dee236f33567d88889822f7521f77a65adc92142abf5a2005d3ef.
//
// Solidity: event RentalPeriodExtended(uint256 indexed rentalTokenId, address indexed renter, address paymentToken, uint112 poolFee, uint112 serviceFee, uint32 endTime, uint32 renterOnlyReturnTime, uint32 enterpriseOnlyCollectionTime)
func (_Enterprise *EnterpriseFilterer) WatchRentalPeriodExtended(opts *bind.WatchOpts, sink chan<- *EnterpriseRentalPeriodExtended, rentalTokenId []*big.Int, renter []common.Address) (event.Subscription, error) {

	var rentalTokenIdRule []interface{}
	for _, rentalTokenIdItem := range rentalTokenId {
		rentalTokenIdRule = append(rentalTokenIdRule, rentalTokenIdItem)
	}
	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "RentalPeriodExtended", rentalTokenIdRule, renterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseRentalPeriodExtended)
				if err := _Enterprise.contract.UnpackLog(event, "RentalPeriodExtended", log); err != nil {
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

// ParseRentalPeriodExtended is a log parse operation binding the contract event 0x012ee01e680dee236f33567d88889822f7521f77a65adc92142abf5a2005d3ef.
//
// Solidity: event RentalPeriodExtended(uint256 indexed rentalTokenId, address indexed renter, address paymentToken, uint112 poolFee, uint112 serviceFee, uint32 endTime, uint32 renterOnlyReturnTime, uint32 enterpriseOnlyCollectionTime)
func (_Enterprise *EnterpriseFilterer) ParseRentalPeriodExtended(log types.Log) (*EnterpriseRentalPeriodExtended, error) {
	event := new(EnterpriseRentalPeriodExtended)
	if err := _Enterprise.contract.UnpackLog(event, "RentalPeriodExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseRentalReturnedIterator is returned from FilterRentalReturned and is used to iterate over the raw logs and unpacked data for RentalReturned events raised by the Enterprise contract.
type EnterpriseRentalReturnedIterator struct {
	Event *EnterpriseRentalReturned // Event containing the contract specifics and raw log

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
func (it *EnterpriseRentalReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseRentalReturned)
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
		it.Event = new(EnterpriseRentalReturned)
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
func (it *EnterpriseRentalReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseRentalReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseRentalReturned represents a RentalReturned event raised by the Enterprise contract.
type EnterpriseRentalReturned struct {
	RentalTokenId    *big.Int
	Returner         common.Address
	PowerToken       common.Address
	RentalAmount     *big.Int
	GcRewardAmount   *big.Int
	GcRewardToken    common.Address
	TotalReserve     *big.Int
	TotalUsedReserve *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRentalReturned is a free log retrieval operation binding the contract event 0xe8c3d212180d23288b990f9938fef98c64caec34bc128fccea9f7eea006ba738.
//
// Solidity: event RentalReturned(uint256 indexed rentalTokenId, address indexed returner, address indexed powerToken, uint112 rentalAmount, uint112 gcRewardAmount, address gcRewardToken, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) FilterRentalReturned(opts *bind.FilterOpts, rentalTokenId []*big.Int, returner []common.Address, powerToken []common.Address) (*EnterpriseRentalReturnedIterator, error) {

	var rentalTokenIdRule []interface{}
	for _, rentalTokenIdItem := range rentalTokenId {
		rentalTokenIdRule = append(rentalTokenIdRule, rentalTokenIdItem)
	}
	var returnerRule []interface{}
	for _, returnerItem := range returner {
		returnerRule = append(returnerRule, returnerItem)
	}
	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "RentalReturned", rentalTokenIdRule, returnerRule, powerTokenRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseRentalReturnedIterator{contract: _Enterprise.contract, event: "RentalReturned", logs: logs, sub: sub}, nil
}

// WatchRentalReturned is a free log subscription operation binding the contract event 0xe8c3d212180d23288b990f9938fef98c64caec34bc128fccea9f7eea006ba738.
//
// Solidity: event RentalReturned(uint256 indexed rentalTokenId, address indexed returner, address indexed powerToken, uint112 rentalAmount, uint112 gcRewardAmount, address gcRewardToken, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) WatchRentalReturned(opts *bind.WatchOpts, sink chan<- *EnterpriseRentalReturned, rentalTokenId []*big.Int, returner []common.Address, powerToken []common.Address) (event.Subscription, error) {

	var rentalTokenIdRule []interface{}
	for _, rentalTokenIdItem := range rentalTokenId {
		rentalTokenIdRule = append(rentalTokenIdRule, rentalTokenIdItem)
	}
	var returnerRule []interface{}
	for _, returnerItem := range returner {
		returnerRule = append(returnerRule, returnerItem)
	}
	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "RentalReturned", rentalTokenIdRule, returnerRule, powerTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseRentalReturned)
				if err := _Enterprise.contract.UnpackLog(event, "RentalReturned", log); err != nil {
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

// ParseRentalReturned is a log parse operation binding the contract event 0xe8c3d212180d23288b990f9938fef98c64caec34bc128fccea9f7eea006ba738.
//
// Solidity: event RentalReturned(uint256 indexed rentalTokenId, address indexed returner, address indexed powerToken, uint112 rentalAmount, uint112 gcRewardAmount, address gcRewardToken, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) ParseRentalReturned(log types.Log) (*EnterpriseRentalReturned, error) {
	event := new(EnterpriseRentalReturned)
	if err := _Enterprise.contract.UnpackLog(event, "RentalReturned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseRentedIterator is returned from FilterRented and is used to iterate over the raw logs and unpacked data for Rented events raised by the Enterprise contract.
type EnterpriseRentedIterator struct {
	Event *EnterpriseRented // Event containing the contract specifics and raw log

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
func (it *EnterpriseRentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseRented)
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
		it.Event = new(EnterpriseRented)
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
func (it *EnterpriseRentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseRentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseRented represents a Rented event raised by the Enterprise contract.
type EnterpriseRented struct {
	RentalTokenId                *big.Int
	Renter                       common.Address
	PowerToken                   common.Address
	PaymentToken                 common.Address
	RentalAmount                 *big.Int
	PoolFee                      *big.Int
	ServiceFee                   *big.Int
	GcFee                        *big.Int
	StartTime                    uint32
	EndTime                      uint32
	RenterOnlyReturnTime         uint32
	EnterpriseOnlyCollectionTime uint32
	TotalReserve                 *big.Int
	TotalUsedReserve             *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterRented is a free log retrieval operation binding the contract event 0x59dee24cee42a5fd3c1be706cc8906676ac6950d523c17fa522275c74e7b9c75.
//
// Solidity: event Rented(uint256 indexed rentalTokenId, address indexed renter, address indexed powerToken, address paymentToken, uint112 rentalAmount, uint112 poolFee, uint112 serviceFee, uint112 gcFee, uint32 startTime, uint32 endTime, uint32 renterOnlyReturnTime, uint32 enterpriseOnlyCollectionTime, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) FilterRented(opts *bind.FilterOpts, rentalTokenId []*big.Int, renter []common.Address, powerToken []common.Address) (*EnterpriseRentedIterator, error) {

	var rentalTokenIdRule []interface{}
	for _, rentalTokenIdItem := range rentalTokenId {
		rentalTokenIdRule = append(rentalTokenIdRule, rentalTokenIdItem)
	}
	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}
	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "Rented", rentalTokenIdRule, renterRule, powerTokenRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseRentedIterator{contract: _Enterprise.contract, event: "Rented", logs: logs, sub: sub}, nil
}

// WatchRented is a free log subscription operation binding the contract event 0x59dee24cee42a5fd3c1be706cc8906676ac6950d523c17fa522275c74e7b9c75.
//
// Solidity: event Rented(uint256 indexed rentalTokenId, address indexed renter, address indexed powerToken, address paymentToken, uint112 rentalAmount, uint112 poolFee, uint112 serviceFee, uint112 gcFee, uint32 startTime, uint32 endTime, uint32 renterOnlyReturnTime, uint32 enterpriseOnlyCollectionTime, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) WatchRented(opts *bind.WatchOpts, sink chan<- *EnterpriseRented, rentalTokenId []*big.Int, renter []common.Address, powerToken []common.Address) (event.Subscription, error) {

	var rentalTokenIdRule []interface{}
	for _, rentalTokenIdItem := range rentalTokenId {
		rentalTokenIdRule = append(rentalTokenIdRule, rentalTokenIdItem)
	}
	var renterRule []interface{}
	for _, renterItem := range renter {
		renterRule = append(renterRule, renterItem)
	}
	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "Rented", rentalTokenIdRule, renterRule, powerTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseRented)
				if err := _Enterprise.contract.UnpackLog(event, "Rented", log); err != nil {
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

// ParseRented is a log parse operation binding the contract event 0x59dee24cee42a5fd3c1be706cc8906676ac6950d523c17fa522275c74e7b9c75.
//
// Solidity: event Rented(uint256 indexed rentalTokenId, address indexed renter, address indexed powerToken, address paymentToken, uint112 rentalAmount, uint112 poolFee, uint112 serviceFee, uint112 gcFee, uint32 startTime, uint32 endTime, uint32 renterOnlyReturnTime, uint32 enterpriseOnlyCollectionTime, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) ParseRented(log types.Log) (*EnterpriseRented, error) {
	event := new(EnterpriseRented)
	if err := _Enterprise.contract.UnpackLog(event, "Rented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseRenterOnlyReturnPeriodChangedIterator is returned from FilterRenterOnlyReturnPeriodChanged and is used to iterate over the raw logs and unpacked data for RenterOnlyReturnPeriodChanged events raised by the Enterprise contract.
type EnterpriseRenterOnlyReturnPeriodChangedIterator struct {
	Event *EnterpriseRenterOnlyReturnPeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseRenterOnlyReturnPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseRenterOnlyReturnPeriodChanged)
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
		it.Event = new(EnterpriseRenterOnlyReturnPeriodChanged)
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
func (it *EnterpriseRenterOnlyReturnPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseRenterOnlyReturnPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseRenterOnlyReturnPeriodChanged represents a RenterOnlyReturnPeriodChanged event raised by the Enterprise contract.
type EnterpriseRenterOnlyReturnPeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRenterOnlyReturnPeriodChanged is a free log retrieval operation binding the contract event 0x745abdabfd9615abf44cc5ea5223a16cd377282452510a2de5121d44b8c097f1.
//
// Solidity: event RenterOnlyReturnPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) FilterRenterOnlyReturnPeriodChanged(opts *bind.FilterOpts) (*EnterpriseRenterOnlyReturnPeriodChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "RenterOnlyReturnPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseRenterOnlyReturnPeriodChangedIterator{contract: _Enterprise.contract, event: "RenterOnlyReturnPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchRenterOnlyReturnPeriodChanged is a free log subscription operation binding the contract event 0x745abdabfd9615abf44cc5ea5223a16cd377282452510a2de5121d44b8c097f1.
//
// Solidity: event RenterOnlyReturnPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) WatchRenterOnlyReturnPeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseRenterOnlyReturnPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "RenterOnlyReturnPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseRenterOnlyReturnPeriodChanged)
				if err := _Enterprise.contract.UnpackLog(event, "RenterOnlyReturnPeriodChanged", log); err != nil {
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

// ParseRenterOnlyReturnPeriodChanged is a log parse operation binding the contract event 0x745abdabfd9615abf44cc5ea5223a16cd377282452510a2de5121d44b8c097f1.
//
// Solidity: event RenterOnlyReturnPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) ParseRenterOnlyReturnPeriodChanged(log types.Log) (*EnterpriseRenterOnlyReturnPeriodChanged, error) {
	event := new(EnterpriseRenterOnlyReturnPeriodChanged)
	if err := _Enterprise.contract.UnpackLog(event, "RenterOnlyReturnPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseServiceRegisteredIterator is returned from FilterServiceRegistered and is used to iterate over the raw logs and unpacked data for ServiceRegistered events raised by the Enterprise contract.
type EnterpriseServiceRegisteredIterator struct {
	Event *EnterpriseServiceRegistered // Event containing the contract specifics and raw log

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
func (it *EnterpriseServiceRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseServiceRegistered)
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
		it.Event = new(EnterpriseServiceRegistered)
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
func (it *EnterpriseServiceRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseServiceRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseServiceRegistered represents a ServiceRegistered event raised by the Enterprise contract.
type EnterpriseServiceRegistered struct {
	PowerToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterServiceRegistered is a free log retrieval operation binding the contract event 0x2fa31fbaacf5eaf61d648ea7528ada6efb69bfb06d2c3bd35ce511a820fce53e.
//
// Solidity: event ServiceRegistered(address indexed powerToken)
func (_Enterprise *EnterpriseFilterer) FilterServiceRegistered(opts *bind.FilterOpts, powerToken []common.Address) (*EnterpriseServiceRegisteredIterator, error) {

	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "ServiceRegistered", powerTokenRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseServiceRegisteredIterator{contract: _Enterprise.contract, event: "ServiceRegistered", logs: logs, sub: sub}, nil
}

// WatchServiceRegistered is a free log subscription operation binding the contract event 0x2fa31fbaacf5eaf61d648ea7528ada6efb69bfb06d2c3bd35ce511a820fce53e.
//
// Solidity: event ServiceRegistered(address indexed powerToken)
func (_Enterprise *EnterpriseFilterer) WatchServiceRegistered(opts *bind.WatchOpts, sink chan<- *EnterpriseServiceRegistered, powerToken []common.Address) (event.Subscription, error) {

	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "ServiceRegistered", powerTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseServiceRegistered)
				if err := _Enterprise.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
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

// ParseServiceRegistered is a log parse operation binding the contract event 0x2fa31fbaacf5eaf61d648ea7528ada6efb69bfb06d2c3bd35ce511a820fce53e.
//
// Solidity: event ServiceRegistered(address indexed powerToken)
func (_Enterprise *EnterpriseFilterer) ParseServiceRegistered(log types.Log) (*EnterpriseServiceRegistered, error) {
	event := new(EnterpriseServiceRegistered)
	if err := _Enterprise.contract.UnpackLog(event, "ServiceRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStakeChangedIterator is returned from FilterStakeChanged and is used to iterate over the raw logs and unpacked data for StakeChanged events raised by the Enterprise contract.
type EnterpriseStakeChangedIterator struct {
	Event *EnterpriseStakeChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStakeChanged)
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
		it.Event = new(EnterpriseStakeChanged)
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
func (it *EnterpriseStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStakeChanged represents a StakeChanged event raised by the Enterprise contract.
type EnterpriseStakeChanged struct {
	StakeTokenId     *big.Int
	Staker           common.Address
	Operation        uint8
	AmountDelta      *big.Int
	Amount           *big.Int
	SharesDelta      *big.Int
	Shares           *big.Int
	TotalShares      *big.Int
	TotalReserve     *big.Int
	TotalUsedReserve *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakeChanged is a free log retrieval operation binding the contract event 0x03429b4bf882467b990034c9cdcf7d1c4c3f189e62ce27af5e9bd563659fa286.
//
// Solidity: event StakeChanged(uint256 indexed stakeTokenId, address indexed staker, uint8 indexed operation, uint256 amountDelta, uint256 amount, uint256 sharesDelta, uint256 shares, uint256 totalShares, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) FilterStakeChanged(opts *bind.FilterOpts, stakeTokenId []*big.Int, staker []common.Address, operation []uint8) (*EnterpriseStakeChangedIterator, error) {

	var stakeTokenIdRule []interface{}
	for _, stakeTokenIdItem := range stakeTokenId {
		stakeTokenIdRule = append(stakeTokenIdRule, stakeTokenIdItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "StakeChanged", stakeTokenIdRule, stakerRule, operationRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseStakeChangedIterator{contract: _Enterprise.contract, event: "StakeChanged", logs: logs, sub: sub}, nil
}

// WatchStakeChanged is a free log subscription operation binding the contract event 0x03429b4bf882467b990034c9cdcf7d1c4c3f189e62ce27af5e9bd563659fa286.
//
// Solidity: event StakeChanged(uint256 indexed stakeTokenId, address indexed staker, uint8 indexed operation, uint256 amountDelta, uint256 amount, uint256 sharesDelta, uint256 shares, uint256 totalShares, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) WatchStakeChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStakeChanged, stakeTokenId []*big.Int, staker []common.Address, operation []uint8) (event.Subscription, error) {

	var stakeTokenIdRule []interface{}
	for _, stakeTokenIdItem := range stakeTokenId {
		stakeTokenIdRule = append(stakeTokenIdRule, stakeTokenIdItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "StakeChanged", stakeTokenIdRule, stakerRule, operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStakeChanged)
				if err := _Enterprise.contract.UnpackLog(event, "StakeChanged", log); err != nil {
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

// ParseStakeChanged is a log parse operation binding the contract event 0x03429b4bf882467b990034c9cdcf7d1c4c3f189e62ce27af5e9bd563659fa286.
//
// Solidity: event StakeChanged(uint256 indexed stakeTokenId, address indexed staker, uint8 indexed operation, uint256 amountDelta, uint256 amount, uint256 sharesDelta, uint256 shares, uint256 totalShares, uint256 totalReserve, uint256 totalUsedReserve)
func (_Enterprise *EnterpriseFilterer) ParseStakeChanged(log types.Log) (*EnterpriseStakeChanged, error) {
	event := new(EnterpriseStakeChanged)
	if err := _Enterprise.contract.UnpackLog(event, "StakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStreamingReserveChangedIterator is returned from FilterStreamingReserveChanged and is used to iterate over the raw logs and unpacked data for StreamingReserveChanged events raised by the Enterprise contract.
type EnterpriseStreamingReserveChangedIterator struct {
	Event *EnterpriseStreamingReserveChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStreamingReserveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStreamingReserveChanged)
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
		it.Event = new(EnterpriseStreamingReserveChanged)
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
func (it *EnterpriseStreamingReserveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStreamingReserveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStreamingReserveChanged represents a StreamingReserveChanged event raised by the Enterprise contract.
type EnterpriseStreamingReserveChanged struct {
	StreamingReserve       *big.Int
	StreamingReserveTarget *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterStreamingReserveChanged is a free log retrieval operation binding the contract event 0xc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a11829.
//
// Solidity: event StreamingReserveChanged(uint112 streamingReserve, uint112 streamingReserveTarget)
func (_Enterprise *EnterpriseFilterer) FilterStreamingReserveChanged(opts *bind.FilterOpts) (*EnterpriseStreamingReserveChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "StreamingReserveChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStreamingReserveChangedIterator{contract: _Enterprise.contract, event: "StreamingReserveChanged", logs: logs, sub: sub}, nil
}

// WatchStreamingReserveChanged is a free log subscription operation binding the contract event 0xc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a11829.
//
// Solidity: event StreamingReserveChanged(uint112 streamingReserve, uint112 streamingReserveTarget)
func (_Enterprise *EnterpriseFilterer) WatchStreamingReserveChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStreamingReserveChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "StreamingReserveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStreamingReserveChanged)
				if err := _Enterprise.contract.UnpackLog(event, "StreamingReserveChanged", log); err != nil {
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

// ParseStreamingReserveChanged is a log parse operation binding the contract event 0xc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a11829.
//
// Solidity: event StreamingReserveChanged(uint112 streamingReserve, uint112 streamingReserveTarget)
func (_Enterprise *EnterpriseFilterer) ParseStreamingReserveChanged(log types.Log) (*EnterpriseStreamingReserveChanged, error) {
	event := new(EnterpriseStreamingReserveChanged)
	if err := _Enterprise.contract.UnpackLog(event, "StreamingReserveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseStreamingReserveHalvingPeriodChangedIterator is returned from FilterStreamingReserveHalvingPeriodChanged and is used to iterate over the raw logs and unpacked data for StreamingReserveHalvingPeriodChanged events raised by the Enterprise contract.
type EnterpriseStreamingReserveHalvingPeriodChangedIterator struct {
	Event *EnterpriseStreamingReserveHalvingPeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseStreamingReserveHalvingPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseStreamingReserveHalvingPeriodChanged)
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
		it.Event = new(EnterpriseStreamingReserveHalvingPeriodChanged)
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
func (it *EnterpriseStreamingReserveHalvingPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseStreamingReserveHalvingPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseStreamingReserveHalvingPeriodChanged represents a StreamingReserveHalvingPeriodChanged event raised by the Enterprise contract.
type EnterpriseStreamingReserveHalvingPeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStreamingReserveHalvingPeriodChanged is a free log retrieval operation binding the contract event 0x9b2baf1f9580f11e84f351d1ac9c543726f7023ba5d6d278fc487898fac055b6.
//
// Solidity: event StreamingReserveHalvingPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) FilterStreamingReserveHalvingPeriodChanged(opts *bind.FilterOpts) (*EnterpriseStreamingReserveHalvingPeriodChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "StreamingReserveHalvingPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseStreamingReserveHalvingPeriodChangedIterator{contract: _Enterprise.contract, event: "StreamingReserveHalvingPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchStreamingReserveHalvingPeriodChanged is a free log subscription operation binding the contract event 0x9b2baf1f9580f11e84f351d1ac9c543726f7023ba5d6d278fc487898fac055b6.
//
// Solidity: event StreamingReserveHalvingPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) WatchStreamingReserveHalvingPeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseStreamingReserveHalvingPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "StreamingReserveHalvingPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseStreamingReserveHalvingPeriodChanged)
				if err := _Enterprise.contract.UnpackLog(event, "StreamingReserveHalvingPeriodChanged", log); err != nil {
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

// ParseStreamingReserveHalvingPeriodChanged is a log parse operation binding the contract event 0x9b2baf1f9580f11e84f351d1ac9c543726f7023ba5d6d278fc487898fac055b6.
//
// Solidity: event StreamingReserveHalvingPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) ParseStreamingReserveHalvingPeriodChanged(log types.Log) (*EnterpriseStreamingReserveHalvingPeriodChanged, error) {
	event := new(EnterpriseStreamingReserveHalvingPeriodChanged)
	if err := _Enterprise.contract.UnpackLog(event, "StreamingReserveHalvingPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
