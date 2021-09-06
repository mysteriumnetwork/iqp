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

// EnterpriseStorageLiquidityInfo is an auto generated low-level Go binding around an user-defined struct.
type EnterpriseStorageLiquidityInfo struct {
	Amount *big.Int
	Shares *big.Int
	Block  *big.Int
}

// EnterpriseStorageLoanInfo is an auto generated low-level Go binding around an user-defined struct.
type EnterpriseStorageLoanInfo struct {
	Amount                     *big.Int
	PowerTokenIndex            uint16
	BorrowingTime              uint32
	MaturityTime               uint32
	BorrowerReturnGraceTime    uint32
	EnterpriseCollectGraceTime uint32
	GcFee                      *big.Int
	GcFeeTokenIndex            uint16
}

// EnterpriseMetaData contains all meta data concerning the Enterprise contract.
var EnterpriseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"BaseUriChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"name\":\"BondingChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"}],\"name\":\"Borrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"BorrowerLoanReturnGracePeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"converter\",\"type\":\"address\"}],\"name\":\"ConverterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"}],\"name\":\"EnterpriseCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"EnterpriseLoanCollectGracePeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnterpriseShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"EnterpriseVaultChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"}],\"name\":\"FixedReserveChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"percent\",\"type\":\"uint16\"}],\"name\":\"GcFeePercentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"InterestGapHalvingPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumEnterprise.LiquidityChangeType\",\"name\":\"changeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"}],\"name\":\"LoanReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PaymentTokenChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"powerToken\",\"type\":\"address\"}],\"name\":\"ServiceRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"streamingReserve\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"streamingReserveTarget\",\"type\":\"uint112\"}],\"name\":\"StreamingReserveChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"TotalSharesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"}],\"name\":\"UsedReserveChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidityAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"powerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxPayment\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"disablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"enablePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"powerToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"}],\"name\":\"estimateLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"}],\"name\":\"getAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondingCurve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowToken\",\"outputs\":[{\"internalType\":\"contractIBorrowToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowerLoanReturnGracePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConverter\",\"outputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseLoanCollectGracePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCFeePercent\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"interestGapHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"borrowerLoanReturnGracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enterpriseLoanCollectGracePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"fixedReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint112\",\"name\":\"streamingReserve\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"streamingReserveTarget\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"streamingReserveUpdated\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterestGapHalvingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterestToken\",\"outputs\":[{\"internalType\":\"contractIInterestToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"}],\"name\":\"getLiquidityInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"internalType\":\"structEnterpriseStorage.LiquidityInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidityToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"}],\"name\":\"getLoanInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"powerTokenIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"borrowingTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maturityTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"borrowerReturnGraceTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enterpriseCollectGraceTime\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"gcFee\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"gcFeeTokenIndex\",\"type\":\"uint16\"}],\"internalType\":\"structEnterpriseStorage.LoanInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPowerTokens\",\"outputs\":[{\"internalType\":\"contractPowerToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsedReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"enterpriseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"contractIConverter\",\"name\":\"converter\",\"type\":\"address\"},{\"internalType\":\"contractProxyAdmin\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractIInterestToken\",\"name\":\"interestToken\",\"type\":\"address\"},{\"internalType\":\"contractIBorrowToken\",\"name\":\"borrowToken\",\"type\":\"address\"}],\"name\":\"initializeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"powerToken\",\"type\":\"address\"}],\"name\":\"isRegisteredPowerToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isSupportedPaymentToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"}],\"name\":\"loanTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"paymentTokenIndex\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"maxPayment\",\"type\":\"uint256\"}],\"name\":\"reborrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"gapHalvingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"baseRate\",\"type\":\"uint112\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"serviceFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minLoanDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxLoanDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"minGCFee\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"allowsPerpetualTokensForever\",\"type\":\"bool\"}],\"name\":\"registerService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowTokenId\",\"type\":\"uint256\"}],\"name\":\"returnLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pole\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slope\",\"type\":\"uint256\"}],\"name\":\"setBondingCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newPeriod\",\"type\":\"uint32\"}],\"name\":\"setBorrowerLoanReturnGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConverter\",\"name\":\"newConverter\",\"type\":\"address\"}],\"name\":\"setConverter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollector\",\"type\":\"address\"}],\"name\":\"setEnterpriseCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newPeriod\",\"type\":\"uint32\"}],\"name\":\"setEnterpriseLoanCollectGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"setEnterpriseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newGcFeePercent\",\"type\":\"uint16\"}],\"name\":\"setGcFeePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"interestGapHalvingPeriod\",\"type\":\"uint32\"}],\"name\":\"setInterestGapHalvingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownEnterpriseForever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeBorrowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeEnterprise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradeInterestToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"powerToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgradePowerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestTokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615baf80620000216000396000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c8062ca339b146102975780630714012c146102ac57806308669aab146102d55780630cac36b2146102e85780630e101b60146102fd5780630ea1dc9e1461031d5780632261b07f1461032e578063229304dc1461033f5780632b2f6095146103605780632e365ade146103735780632fb2067f146103865780633513e0dc146103a9578063449497ec146103bc5780634741efb3146103cf57806349bd4e89146103d7578063512a53bc146103ea57806351c6590a146103fd57806359bf5d39146104105780635a4598df146104185780635a9b0b891461043657806362c1f38814610456578063644e44df146104695780636815f3371461047a578063696f9c811461048d5780637791c42d146104a05780637b024363146104b15780638188d50d146104c457806389035b61146104d95780638b3240a0146105055780638c0114b01461050d5780638da5cb5b1461052057806390f55edd146105285780639339ea691461053b5780639ab711931461054e5780639c7aa7f8146105615780639c8f9f2314610574578063a0bcfc7f14610587578063a0efd1181461059a578063a554b26d146105af578063aebb5041146105c0578063b19337a4146105c8578063b406bf6d146105db578063b56de97e146105fc578063b96266fa1461060f578063bfd84fb414610624578063c1fc16d814610637578063c22947841461064a578063c4d66de81461065d578063c93d0b9214610670578063df57f2fa14610681578063e1a5fa49146106b6578063e9126154146106c9578063ef1f9f39146106dc578063f2fde38b146106ef578063faac38ef14610702578063fbac9a551461071d578063fcb2884e14610743575b600080fd5b6102aa6102a5366004614f48565b61074b565b005b6102bf6102ba36600461526c565b611052565b6040516102cc919061560c565b60405180910390f35b6102aa6102e336600461526c565b6111a7565b6102f061134b565b6040516102cc9190615521565b61031061030b36600461526c565b6113dd565b6040516102cc91906153ce565b6001546001600160a01b0316610310565b6004546001600160a01b0316610310565b61035261034d366004614eef565b61141b565b6040519081526020016102cc565b6102aa61036e366004614dd9565b611544565b6102aa610381366004615302565b611600565b610399610394366004614dd9565b6116e2565b60405190151581526020016102cc565b6102aa6103b73660046152e1565b611705565b6102aa6103ca366004615234565b611812565b6102aa61189c565b6102aa6103e53660046152e1565b611996565b6102aa6103f8366004614e11565b611bc0565b6102aa61040b36600461526c565b611e47565b610352611fc5565b600354600160a01b900463ffffffff165b6040516102cc91906156b1565b61043e611fea565b6040516102cc9c9b9a99989796959493929190615534565b6102aa610464366004614dd9565b6121ad565b6000546001600160a01b0316610310565b6102aa61048836600461505a565b612273565b6102aa61049b3660046152e1565b61259f565b6006546001600160a01b0316610310565b6102aa6104bf366004614dd9565b61273f565b600654600160a01b900463ffffffff16610429565b6103996104e7366004614dd9565b6001600160a01b031660009081526013602052604090205460ff1690565b6103106127c6565b6102aa61051b36600461529c565b6127e7565b610310612e13565b6102aa610536366004614dd9565b612e29565b61035261054936600461526c565b612eef565b6102aa61055c36600461510c565b612f34565b6102aa61056f366004614dd9565b6133ef565b6102aa61058236600461526c565b613535565b6102aa610595366004614fa9565b613770565b600654600160c01b900463ffffffff16610429565b6005546001600160a01b0316610310565b6103526137e8565b6102aa6105d6366004614dd9565b6137ff565b600654600160e01b900461ffff1660405161ffff90911681526020016102cc565b6102aa61060a366004615302565b6138c5565b61061761399b565b6040516102cc9190615438565b6102aa610632366004614dd9565b6139fc565b6102aa610645366004614dd9565b613a92565b6102aa610658366004614eb7565b613b13565b6102aa61066b366004614dd9565b613c20565b6002546001600160a01b0316610310565b61069461068f36600461526c565b613cf5565b60408051825181526020808401519082015291810151908201526060016102cc565b6102aa6106c4366004615302565b613dcf565b6102aa6106d736600461526c565b613e95565b6102aa6106ea366004614e6d565b614153565b6102aa6106fd366004614dd9565b6141e2565b600e54600d54604080519283526020830191909152016102cc565b61073061072b366004614dd9565b6142d3565b60405160019190910b81526020016102cc565b600a54610352565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156107995760405162461bcd60e51b81526004016107909190615521565b60405180910390fd5b506001600160a01b03851660009081526013602090815260409182902054825180840190935260018352601b60f91b9183019190915260ff166107ef5760405162461bcd60e51b81526004016107909190615521565b506107f9846116e2565b604051806040016040528060028152602001610d0d60f21b815250906108325760405162461bcd60e51b81526004016107909190615521565b50604051636e4d898360e01b81526001600160a01b03861690636e4d89839061085f9085906004016156b1565b60206040518083038186803b15801561087757600080fd5b505afa15801561088b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108af9190614e51565b60405180604001604052806002815260200161343560f01b815250906108e85760405162461bcd60e51b81526004016107909190615521565b506040805180820190915260028152611b9b60f11b60208201526001600160701b0384166109295760405162461bcd60e51b81526004016107909190615521565b506109326137e8565b836001600160701b03161115604051806040016040528060028152602001611a1b60f11b815250906109775760405162461bcd60e51b81526004016107909190615521565b50600080600080886001600160a01b0316639d5513d78989896040518463ffffffff1660e01b81526004016109ae93929190615485565b60606040518083038186803b1580156109c657600080fd5b505afa1580156109da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fe91906151f3565b955090935091508390506000610a148385615719565b6001600160701b0316905085856001600160701b031682610a359190615744565b111560405180604001604052806002815260200161343760f01b81525090610a705760405162461bcd60e51b81526004016107909190615521565b50610a866001600160a01b038a16333084614300565b60005481906001600160a01b038b8116911614610bcf576004805460405163095ea7b360e01b81526001600160a01b03808e169363095ea7b393610acf9392169187910161541f565b602060405180830381600087803b158015610ae957600080fd5b505af1158015610afd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b219190614e51565b50600460009054906101000a90046001600160a01b03166001600160a01b031663029b465d8b8460008054906101000a90046001600160a01b03166040518463ffffffff1660e01b8152600401610b7a939291906154b5565b602060405180830381600087803b158015610b9457600080fd5b505af1158015610ba8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bcc9190615284565b90505b600082610be5836001600160701b03881661579b565b610bef919061577b565b600654600054919250610c0f916001600160a01b03908116911683614371565b896001600160701b0316600a6000828254610c2a9190615744565b9091555060009050610c3c8284615827565b9050610c4781614390565b5050600254610c7495506001600160a01b038b81169550339450169150506001600160701b038516614300565b426000610c81858361575c565b90506000600260009054906101000a90046001600160a01b03166001600160a01b031663caa0f92a6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610cd557600080fd5b505af1158015610ce9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0d9190615284565b9050604051806101000160405280886001600160701b031681526020018a6001600160a01b03166381045ead6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d6357600080fd5b505afa158015610d77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9b9190615250565b61ffff16815263ffffffff80861660208301528481166040830152600654606090920191610dd291600160a01b909104168561575c565b63ffffffff168152602001600660189054906101000a900463ffffffff1684610dfb919061575c565b63ffffffff168152602001856001600160701b03168152602001610e1e8a6142d3565b61ffff9081169091526000838152601160209081526040918290208451815492860151868501516060880151608089015160a08a01516001600160701b039586166001600160801b031998891617600160701b958b16860217600160801b600160c01b031916600160801b63ffffffff9586160263ffffffff60a01b191617600160a01b93851693909302929092176001600160c01b0316600160c01b918416919091026001600160e01b031617600160e01b929091169190910217835560c08701516001909301805460e0909801519390921696909316959095179316029190911790915560025490516335313c2160e11b815282916001600160a01b031690636a62784290610f339033906004016153ce565b602060405180830381600087803b158015610f4d57600080fd5b505af1158015610f61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f859190615284565b14610fa057634e487b7160e01b600052600160045260246000fd5b60405163f8642f2960e01b8152600481018290526001600160a01b038a169063f8642f2990602401600060405180830381600087803b158015610fe257600080fd5b505af1158015610ff6573d6000803e3d6000fd5b50506040518392506001600160a01b038c16915060008051602061595a83398151915290600090a3600080516020615b3a833981519152600a5460405161103f91815260200190565b60405180910390a1505050505050505050565b6040805161010081018252600080825260208201819052818301819052606082018190526080820181905260a0820181905260c0820181905260e082015260025491516331a9108f60e11b81526004810184905290916001600160a01b031690636352211e9060240160206040518083038186803b1580156110d357600080fd5b505afa1580156110e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110b9190614df5565b505060009081526011602090815260409182902082516101008101845281546001600160701b038082168352600160701b80830461ffff90811696850196909652600160801b830463ffffffff90811697850197909752600160a01b830487166060850152600160c01b830487166080850152600160e01b90920490951660a083015260019092015493841660c082015292041660e082015290565b6001546040516331a9108f60e11b815260048101839052829133916001600160a01b0390911690636352211e9060240160206040518083038186803b1580156111ef57600080fd5b505afa158015611203573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112279190614df5565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906112695760405162461bcd60e51b81526004016107909190615521565b5060008281526012602052604081206001810154909161128885612eef565b90506112926137e8565b811115604051806040016040528060028152602001611a1b60f11b815250906112ce5760405162461bcd60e51b81526004016107909190615521565b506000546112e6906001600160a01b03163383614371565b60006112f58460000154614442565b60018501819055905061131061130b8285615827565b614464565b611319826144ae565b60008660008051602061599a8339815191528460405161133b91815260200190565b60405180910390a3505050505050565b60606010805461135a9061586a565b80601f01602080910402602001604051908101604052809291908181526020018280546113869061586a565b80156113d35780601f106113a8576101008083540402835291602001916113d3565b820191906000526020600020905b8154815290600101906020018083116113b657829003601f168201915b5050505050905090565b60006008828154811061140057634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031692915050565b600354604080518082019091526002815261373560f01b6020820152600091600160c01b900460ff16156114625760405162461bcd60e51b81526004016107909190615521565b506001600160a01b03851660009081526013602090815260409182902054825180840190935260018352601b60f91b9183019190915260ff166114b85760405162461bcd60e51b81526004016107909190615521565b50604051631466370160e01b81526001600160a01b038616906314663701906114e990879087908790600401615485565b60206040518083038186803b15801561150157600080fd5b505afa158015611515573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115399190615284565b90505b949350505050565b3361154d612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061158f5760405162461bcd60e51b81526004016107909190615521565b506115986127c6565b60015460405163266a23b160e21b81526001600160a01b03928316926399a88ec4926115cb9291169085906004016154d8565b600060405180830381600087803b1580156115e557600080fd5b505af11580156115f9573d6000803e3d6000fd5b5050505050565b33611609612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061164b5760405162461bcd60e51b81526004016107909190615521565b506006546040805180820190915260028152610d8d60f21b60208201529063ffffffff600160c01b9091048116908316111561169a5760405162461bcd60e51b81526004016107909190615521565b506006805463ffffffff60a01b1916600160a01b63ffffffff841602179055604051600080516020615ada833981519152906116d79083906156b1565b60405180910390a150565b6001600160a01b0316600090815260076020526040812054600190810b900b1390565b3361170e612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906117505760405162461bcd60e51b81526004016107909190615521565b50611760600a600360401b61577b565b82111560405180604001604052806002815260200161373760f01b8152509061179c5760405162461bcd60e51b81526004016107909190615521565b5060408051808201909152600281526106e760f31b6020820152600160401b8211156117db5760405162461bcd60e51b81526004016107909190615521565b50600e829055600d81905560408051838152602081018390526000805160206159da83398151915291015b60405180910390a15050565b3361181b612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061185d5760405162461bcd60e51b81526004016107909190615521565b506006805461ffff60e01b1916600160e01b61ffff841690810291909117909155604051908152600080516020615a1a833981519152906020016116d7565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156118e15760405162461bcd60e51b81526004016107909190615521565b50336118eb612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061192d5760405162461bcd60e51b81526004016107909190615521565b5060038054600160c01b60ff60c01b199091161790556000600a819055600b8054600160701b81046001600160701b03166001600160701b03199091161790556040517f6f6348718c9a361558c634b516777b2b06bb2bf4140ad3d3bfaa44270fc2c0729190a1565b6001546040516331a9108f60e11b815260048101849052839133916001600160a01b0390911690636352211e9060240160206040518083038186803b1580156119de57600080fd5b505afa1580156119f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a169190614df5565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090611a585760405162461bcd60e51b81526004016107909190615521565b5060006012600085815260200190815260200160002090504381600201541060405180604001604052806002815260200161035360f41b81525090611ab05760405162461bcd60e51b81526004016107909190615521565b5080546040805180820190915260028152611a1b60f11b602082015290841115611aed5760405162461bcd60e51b81526004016107909190615521565b50611af66137e8565b831115604051806040016040528060028152602001611a1b60f11b81525090611b325760405162461bcd60e51b81526004016107909190615521565b50600054611b4a906001600160a01b03163385614371565b6000611b5584614442565b90508160010154811115611b6a575060018101545b600182018054829003905581548490038255611b8581614464565b611b8e846144ae565b60045b8560008051602061599a83398151915286604051611bb191815260200190565b60405180910390a35050505050565b600280546040805180820190915291825261333960f01b60208301526001600160a01b03163314611c045760405162461bcd60e51b81526004016107909190615521565b506000818152601160209081526040918290205482518084019093526002835261068760f31b918301919091526001600160701b03169081611c595760405162461bcd60e51b81526004016107909190615521565b5060008281526011602052604081205460148054600160a01b830463ffffffff164211936001600160a01b03898116159490891615939192600160701b90910461ffff16908110611cba57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b031690508115611d3c5760405163079cc67960e41b81526001600160a01b038216906379cc679090611d05908b9089906004016153fd565b600060405180830381600087803b158015611d1f57600080fd5b505af1158015611d33573d6000803e3d6000fd5b50505050611e3d565b8215611d70576040516340c10f1960e01b81526001600160a01b038216906340c10f1990611d05908a9089906004016153fd565b83611e0d576040516333bebb7760e01b81526001600160a01b03898116600483015288811660248301526001600160701b03871660448301528216906333bebb7790606401602060405180830381600087803b158015611dcf57600080fd5b505af1158015611de3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e079190614e51565b50611e3d565b60408051808201825260028152611a9960f11b6020820152905162461bcd60e51b81526107909190600401615521565b5050505050505050565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff1615611e8c5760405162461bcd60e51b81526004016107909190615521565b50600054611ea5906001600160a01b0316333084614300565b6000600c54600014611ebf57611eba82614442565b611ec1565b815b9050611ecc82614519565b6001546040516335313c2160e11b81526000916001600160a01b031690636a62784290611efd9033906004016153ce565b602060405180830381600087803b158015611f1757600080fd5b505af1158015611f2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4f9190615284565b6040805160608101825285815260208082018681524383850190815260008681526012909352939091209151825551600182015590516002909101559050611f9682614551565b60018160008051602061599a83398151915285604051611fb891815260200190565b60405180910390a3505050565b6000611fcf614563565b6001600160701b0316600954611fe59190615744565b905090565b600c54600354600654600954600a54600b54600f805460609788976000978897889788978897889788978897889788979496601096939563ffffffff600160a01b94859004811696948404811695600160c01b850482169561ffff600160e01b9687900416959394936001600160701b0380841694600160701b8504909116939190910416908c9061207b9061586a565b80601f01602080910402602001604051908101604052809291908181526020018280546120a79061586a565b80156120f45780601f106120c9576101008083540402835291602001916120f4565b820191906000526020600020905b8154815290600101906020018083116120d757829003601f168201915b50505050509b508a80546121079061586a565b80601f01602080910402602001604051908101604052809291908181526020018280546121339061586a565b80156121805780601f1061215557610100808354040283529160200191612180565b820191906000526020600020905b81548152906001019060200180831161216357829003601f168201915b50505050509a509b509b509b509b509b509b509b509b509b509b509b509b50909192939495969798999a9b565b336121b6612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906121f85760405162461bcd60e51b81526004016107909190615521565b50604080518082019091526002815261363160f01b60208201526001600160a01b0382166122395760405162461bcd60e51b81526004016107909190615521565b50600580546001600160a01b0319166001600160a01b038316179055604051600080516020615aba833981519152906116d79083906153ce565b600f80546122809061586a565b6040805180820190915260018152601960f91b60208201529150156122b85760405162461bcd60e51b81526004016107909190615521565b506122c281613c20565b816000805160206159fa83398151915280546001600160a01b03929092166001600160a01b03199283161790556003805490911633179055865161230d90600f9060208a0190614bd9565b5061231a60108787614c5d565b5060068054600480546001600160a01b03199081166001600160a01b03888116919091179092556005805492861692909116821790556003805463ffffffff60a01b191661127560a71b179055600165ffff0000000160c01b0319909116600160e01b61ffff8816026001600160c01b03191617176102a360a61b1763ffffffff60c01b19166102a360c71b1790556123b86064600560401b61577b565b600e556123ca600a600360401b61577b565b600d5560405160008051602061593a833981519152906123ed90889088906154f2565b60405180910390a1600654604051600160e01b90910461ffff168152600080516020615a1a8339815191529060200160405180910390a1600454604051600080516020615a5a8339815191529161244f916001600160a01b03909116906153ce565b60405180910390a1600654604051600080516020615a3a83398151915291612482916001600160a01b03909116906153ce565b60405180910390a1600554604051600080516020615aba833981519152916124b5916001600160a01b03909116906153ce565b60405180910390a1600080516020615b1a833981519152600360149054906101000a900463ffffffff166040516124ec91906156b1565b60405180910390a1600080516020615ada833981519152600660149054906101000a900463ffffffff1660405161252391906156b1565b60405180910390a16000805160206159ba833981519152600660189054906101000a900463ffffffff1660405161255a91906156b1565b60405180910390a16000805160206159da833981519152600e54600d5460405161258e929190918252602082015260400190565b60405180910390a150505050505050565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff16156125e45760405162461bcd60e51b81526004016107909190615521565b506001546040516331a9108f60e11b815260048101849052839133916001600160a01b0390911690636352211e9060240160206040518083038186803b15801561262d57600080fd5b505afa158015612641573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126659190614df5565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906126a75760405162461bcd60e51b81526004016107909190615521565b506000546126c0906001600160a01b0316333085614300565b6000600c546000146126da576126d583614442565b6126dc565b825b90506126e783614519565b6000848152601260205260408120805490918591839190612709908490615744565b92505081905550818160010160008282546127249190615744565b909155505043600282015561273882614551565b6003611b91565b33612748612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061278a5760405162461bcd60e51b81526004016107909190615521565b506127936127c6565b60025460405163266a23b160e21b81526001600160a01b03928316926399a88ec4926115cb9291169085906004016154d8565b60006000805160206159fa8339815191525b546001600160a01b0316919050565b600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff161561282c5760405162461bcd60e51b81526004016107909190615521565b50612836836116e2565b604051806040016040528060028152602001610d0d60f21b8152509061286f5760405162461bcd60e51b81526004016107909190615521565b50600084815260116020908152604091829020805483518085019094526002845261068760f31b9284019290925291906001600160701b03166128c55760405162461bcd60e51b81526004016107909190615521565b50805460148054600092600160701b900461ffff169081106128f757634e487b7160e01b600052603260045260246000fd5b600091825260209091200154604051636e4d898360e01b81526001600160a01b0390911691508190636e4d8983906129339087906004016156b1565b60206040518083038186803b15801561294b57600080fd5b505afa15801561295f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129839190614e51565b60405180604001604052806002815260200161343560f01b815250906129bc5760405162461bcd60e51b81526004016107909190615521565b50815442906129d9908690600160a01b900463ffffffff1661575c565b63ffffffff16101560405180604001604052806002815260200161343960f01b81525090612a1a5760405162461bcd60e51b81526004016107909190615521565b508154600a80546001600160701b0390921691600090612a3b908490615827565b90915550508154604051639d5513d760e01b815260009182916001600160a01b03851691639d5513d791612a81918b916001600160701b03909116908b90600401615485565b60606040518083038186803b158015612a9957600080fd5b505afa158015612aad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ad191906151f3565b508554600a80546001600160701b03909216909101905590925090506000612af98284615719565b6001600160701b031690508581111560405180604001604052806002815260200161343760f01b81525090612b415760405162461bcd60e51b81526004016107909190615521565b50612b576001600160a01b038916333084614300565b60005481906001600160a01b038a8116911614612c81576004805460405163095ea7b360e01b81526001600160a01b03808d169363095ea7b393612ba09392169187910161541f565b602060405180830381600087803b158015612bba57600080fd5b505af1158015612bce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf29190614e51565b506004805460005460405163029b465d60e01b81526001600160a01b039283169363029b465d93612c2c938f9389939190921691016154b5565b602060405180830381600087803b158015612c4657600080fd5b505af1158015612c5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c7e9190615284565b90505b600082612c97836001600160701b03871661579b565b612ca1919061577b565b600654600054919250612cc1916001600160a01b03908116911683614371565b6000612ccd8284615827565b9050612cd881614390565b8754612cf2908b90600160a01b900463ffffffff1661575c565b885463ffffffff60a01b1916600160a01b63ffffffff928316810291909117808b55600654612d2b93908390048116929091041661575c565b885463ffffffff60c01b1916600160c01b63ffffffff928316810291909117808b55600654612d6893929004821691600160a01b9091041661575c565b885463ffffffff91909116600160e01b026001600160e01b0390911617885560405163f8642f2960e01b8152600481018d90526001600160a01b0388169063f8642f2990602401600060405180830381600087803b158015612dc957600080fd5b505af1158015612ddd573d6000803e3d6000fd5b50506040518e92506001600160a01b038a16915060008051602061595a83398151915290600090a3505050505050505050505050565b600060008051602061597a8339815191526127d8565b33612e32612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090612e745760405162461bcd60e51b81526004016107909190615521565b506040805180820190915260028152611b1960f11b60208201526001600160a01b038216612eb55760405162461bcd60e51b81526004016107909190615521565b50600680546001600160a01b0319166001600160a01b038316179055604051600080516020615a3a833981519152906116d79083906153ce565b600081815260126020526040812060018101548290612f0d906145cb565b8254909150811115612f2a578154612f259082615827565b61153c565b6000949350505050565b33612f3d612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090612f7f5760405162461bcd60e51b81526004016107909190615521565b50600354604080518082019091526002815261373560f01b602082015290600160c01b900460ff1615612fc55760405162461bcd60e51b81526004016107909190615521565b50604080518082019091526002815261034360f41b60208201526001600160a01b0387166130065760405162461bcd60e51b81526004016107909190615521565b50601454604080518082019091526002815261343160f01b60208201529061ffff116130455760405162461bcd60e51b81526004016107909190615521565b506003546000906001600160a01b031663c31011cc6130626127c6565b6040518263ffffffff1660e01b815260040161307e91906153ce565b602060405180830381600087803b15801561309857600080fd5b505af11580156130ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130d09190614df5565b905060008060009054906101000a90046001600160a01b03166001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561312157600080fd5b505afa158015613135573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261315d9190810190614fe8565b90506000818c604051602001613174929190615392565b6040516020818303038152906040529050826001600160a01b0316631624f6c68e8360008054906101000a90046001600160a01b03166001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156131e257600080fd5b505afa1580156131f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061321a919061531c565b6040518463ffffffff1660e01b8152600401613238939291906155d3565b600060405180830381600087803b15801561325257600080fd5b505af1158015613266573d6000803e3d6000fd5b505050505050806001600160a01b0316632ef0cf8b308a868d6014805490508d8c8c8f8c6040518b63ffffffff1660e01b81526004016133249a999897969594939291906001600160a01b039a8b1681526001600160701b039990991660208a01526001600160601b0397909716604089015263ffffffff958616606089015261ffff94851660808901529290971660a0870152831660c08601529490911660e0840152929092166101008201529015156101208201526101400190565b600060405180830381600087803b15801561333e57600080fd5b505af1158015613352573d6000803e3d6000fd5b50506014805460018082019092557fce6d7b5282bd9a3661ae061feed1dbda4e52ab073b1f9285be6e155d9c38d4ec0180546001600160a01b0319166001600160a01b038616908117909155600081815260136020526040808220805460ff191690941790935591519093507f2fa31fbaacf5eaf61d648ea7528ada6efb69bfb06d2c3bd35ce511a820fce53e9250a25050505050505050505050565b336133f8612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061343a5760405162461bcd60e51b81526004016107909190615521565b506001600160a01b03811660009081526007602090815260409182902054825180840190935260028352611b9960f11b91830191909152600190810b900b6134955760405162461bcd60e51b81526004016107909190615521565b506001600160a01b038116600090815260076020526040812054600190810b900b1315613532576001600160a01b0381166000908152600760205260409020546134e19060010b6158a5565b6001600160a01b038216600090815260076020526040808220805460019490940b61ffff1661ffff19909416939093179092559051600080516020615afa833981519152916116d7918491906153e2565b50565b6001546040516331a9108f60e11b815260048101839052829133916001600160a01b0390911690636352211e9060240160206040518083038186803b15801561357d57600080fd5b505afa158015613591573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135b59190614df5565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906135f75760405162461bcd60e51b81526004016107909190615521565b5060006012600084815260200190815260200160002090504381600201541060405180604001604052806002815260200161035360f41b8152509061364f5760405162461bcd60e51b81526004016107909190615521565b5060018101546000613660826145cb565b905061366a6137e8565b811115604051806040016040528060028152602001611a1b60f11b815250906136a65760405162461bcd60e51b81526004016107909190615521565b50600154604051630852cd8d60e31b8152600481018790526001600160a01b03909116906342966c6890602401600060405180830381600087803b1580156136ed57600080fd5b505af1158015613701573d6000803e3d6000fd5b505060005461371d92506001600160a01b031690503383614371565b61372682614464565b61372f816144ae565b6000858152601260205260408120818155600181018290556002908101919091558560008051602061599a83398151915283604051611bb191815260200190565b33613779612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906137bb5760405162461bcd60e51b81526004016107909190615521565b506137c860108383614c5d565b5060008051602061593a83398151915282826040516118069291906154f2565b6000600a546137f5611fc5565b611fe59190615827565b33613808612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061384a5760405162461bcd60e51b81526004016107909190615521565b50604080518082019091526002815261363360f01b60208201526001600160a01b03821661388b5760405162461bcd60e51b81526004016107909190615521565b50600480546001600160a01b0319166001600160a01b038316179055604051600080516020615a5a833981519152906116d79083906153ce565b336138ce612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b815250906139105760405162461bcd60e51b81526004016107909190615521565b50600654604080518082019091526002815261363560f01b60208201529063ffffffff808416600160a01b90920416111561395e5760405162461bcd60e51b81526004016107909190615521565b506006805463ffffffff60c01b1916600160c01b63ffffffff8416021790556040516000805160206159ba833981519152906116d79083906156b1565b606060148054806020026020016040519081016040528092919081815260200182805480156113d357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116139d5575050505050905090565b33613a05612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090613a475760405162461bcd60e51b81526004016107909190615521565b50604080518082019091526002815261373160f01b60208201526001600160a01b038216613a885760405162461bcd60e51b81526004016107909190615521565b50613532816145e3565b33613a9b612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090613add5760405162461bcd60e51b81526004016107909190615521565b50613ae66127c6565b6001600160a01b03166399a88ec430836040518363ffffffff1660e01b81526004016115cb9291906154d8565b33613b1c612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090613b5e5760405162461bcd60e51b81526004016107909190615521565b506001600160a01b03821660009081526013602090815260409182902054825180840190935260018352601b60f91b9183019190915260ff16613bb45760405162461bcd60e51b81526004016107909190615521565b50613bbd6127c6565b6001600160a01b03166399a88ec483836040518363ffffffff1660e01b8152600401613bea9291906154d8565b600060405180830381600087803b158015613c0457600080fd5b505af1158015613c18573d6000803e3d6000fd5b505050505050565b6000613c2a612e13565b6001600160a01b031614604051806040016040528060018152602001601960f91b81525090613c6c5760405162461bcd60e51b81526004016107909190615521565b50604080518082019091526002815261373360f01b60208201526001600160a01b038216613cad5760405162461bcd60e51b81526004016107909190615521565b5060008051602061597a83398151915280546001600160a01b0319166001600160a01b038316908117909155604051600090600080516020615a9a833981519152908290a350565b613d1960405180606001604052806000815260200160008152602001600081525090565b6001546040516331a9108f60e11b8152600481018490526001600160a01b0390911690636352211e9060240160206040518083038186803b158015613d5d57600080fd5b505afa158015613d71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d959190614df5565b5050600090815260126020908152604091829020825160608101845281548152600182015492810192909252600201549181019190915290565b33613dd8612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b81525090613e1a5760405162461bcd60e51b81526004016107909190615521565b506040805180820190915260028152611b1b60f11b602082015263ffffffff8216613e585760405162461bcd60e51b81526004016107909190615521565b506003805463ffffffff60a01b1916600160a01b63ffffffff841602179055604051600080516020615b1a833981519152906116d79083906156b1565b600081815260116020908152604091829020805483518085019094526002845261068760f31b9284019290925291906001600160701b0316613eea5760405162461bcd60e51b81526004016107909190615521565b506002546040516331a9108f60e11b8152600481018490526000916001600160a01b031690636352211e9060240160206040518083038186803b158015613f3057600080fd5b505afa158015613f44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f689190614df5565b8254909150429063ffffffff808316600160c01b909204161080613f945750336001600160a01b038316145b60405180604001604052806002815260200161353360f01b81525090613fcd5760405162461bcd60e51b81526004016107909190615521565b50825463ffffffff808316600160e01b909204161080613ff55750336001600160a01b038316145b8061400a57506005546001600160a01b031633145b604051806040016040528060028152602001610d4d60f21b815250906140435760405162461bcd60e51b81526004016107909190615521565b50600354600160c01b900460ff1661409f578254600a80546001600160701b0390921691600090614075908490615827565b9091555050600a54604051908152600080516020615b3a8339815191529060200160405180910390a15b600254604051633f34d4cf60e21b8152600481018690523360248201526001600160a01b039091169063fcd3533c90604401600060405180830381600087803b1580156140eb57600080fd5b505af11580156140ff573d6000803e3d6000fd5b50505060008581526011602052604080822082815560010180546001600160801b0319169055518692507f80fadef8be1cf6508244c1bbb67077973d214a1462d3ade2908699c3750cd94b9190a250505050565b6000546040805180820190915260018152601960f91b6020820152906001600160a01b0316156141965760405162461bcd60e51b81526004016107909190615521565b50600080546001600160a01b038086166001600160a01b0319928316179092556001805485841690831617905560028054928416929091169190911790556141dd836145e3565b505050565b336141eb612e13565b6001600160a01b031614604051806040016040528060018152602001603360f81b8152509061422d5760405162461bcd60e51b81526004016107909190615521565b50604080518082019091526002815261373360f01b60208201526001600160a01b03821661426e5760405162461bcd60e51b81526004016107909190615521565b50806001600160a01b0316614281612e13565b6001600160a01b0316600080516020615a9a83398151915260405160405180910390a360008051602061597a83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381166000908152600760205260408120546142fa90600190810b6157ba565b92915050565b6040516001600160a01b038085166024830152831660448201526064810182905261436b9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261472e565b50505050565b6141dd8363a9059cbb60e01b848460405160240161433492919061541f565b614398614563565b600b80546001600160701b0319166001600160701b03928316178082558392600e916143cd918591600160701b900416615719565b82546001600160701b039182166101009390930a928302928202191691909117909155600b805463ffffffff4216600160e01b026001600160e01b038216811792839055604051600080516020615a7a83398151915295506116d7949182169282169290921792600160701b90041690615697565b600061444c611fc5565b82600c5461445a919061579b565b6142fa919061577b565b80600c60008282546144769190615827565b9091555050600c546040519081527ff75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264906020016116d7565b80600954106144c5576009805482900390556144f8565b60006144cf614800565b6001600160701b0316905081816009546144e99190615744565b6144f39190615827565b600955505b600080516020615b5a8339815191526009546040516116d791815260200190565b806009600082825461452b9190615744565b9091555050600954604051908152600080516020615b5a833981519152906020016116d7565b80600c60008282546144769190615744565b600b546000906145ae9063ffffffff600160e01b82041690614598906001600160701b0380821691600160701b9004166157ff565b600354600160a01b900463ffffffff16426148be565b600b54611fe59190600160701b90046001600160701b03166157ff565b6000600c54826145d9611fc5565b61445a919061579b565b6001600160a01b038116600090815260076020526040902054600190810b900b6146995760088054600181810183557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390910180546001600160a01b0385166001600160a01b03199091168117909155915460009283526007602052604092839020805491830b61ffff1661ffff199092169190911790559051600080516020615afa833981519152916116d7918491906153e2565b6001600160a01b038116600090815260076020526040812054600190810b900b1215613532576001600160a01b0381166000908152600760205260409020546146e49060010b6158a5565b6001600160a01b03821660009081526007602052604090819020805461ffff191661ffff600194850b1617905551600080516020615afa833981519152916116d7918491906153e2565b6000614783826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614a5f9092919063ffffffff16565b8051909150156141dd57808060200190518101906147a19190614e51565b6141dd5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610790565b600061480a614563565b600b80546001600160701b0319168082559192508291600e9061483e908490600160701b90046001600160701b03166157ff565b82546001600160701b039182166101009390930a928302928202191691909117909155600b805463ffffffff4216600160e01b026001600160e01b038216811792839055604051600080516020615a7a83398151915295506148b3949182169282169290921792600160701b90041690615697565b60405180910390a190565b60008463ffffffff168263ffffffff161015604051806040016040528060018152602001600760fb1b815250906149085760405162461bcd60e51b81526004016107909190615521565b5084820391508263ffffffff168263ffffffff168161493757634e487b7160e01b600052601260045260246000fd5b0463ffffffff16846001600160701b0316901c93508263ffffffff168263ffffffff168161497557634e487b7160e01b600052601260045260246000fd5b06915063ffffffff8216158061499257506001600160701b038416155b1561499e57508261153c565b60006001600160701b0385168163ffffffff86811690861671b17217f7d1cf79abc9e3b39803f2f6af40f302816149e557634e487b7160e01b600052601260045260246000fd5b049050600160901b5b8215614a5257928201928082840281614a1757634e487b7160e01b600052601260045260246000fd5b0493849003939250600160901b018082840281614a4457634e487b7160e01b600052601260045260246000fd5b049250600160901b016149ee565b5091979650505050505050565b6060614a6e8484600085614a78565b90505b9392505050565b606082471015614ad95760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610790565b843b614b275760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610790565b600080866001600160a01b03168587604051614b439190615376565b60006040518083038185875af1925050503d8060008114614b80576040519150601f19603f3d011682016040523d82523d6000602084013e614b85565b606091505b5091509150614b95828286614ba0565b979650505050505050565b60608315614baf575081614a71565b825115614bbf5782518084602001fd5b8160405162461bcd60e51b81526004016107909190615521565b828054614be59061586a565b90600052602060002090601f016020900481019282614c075760008555614c4d565b82601f10614c2057805160ff1916838001178555614c4d565b82800160010185558215614c4d579182015b82811115614c4d578251825591602001919060010190614c32565b50614c59929150614cd1565b5090565b828054614c699061586a565b90600052602060002090601f016020900481019282614c8b5760008555614c4d565b82601f10614ca45782800160ff19823516178555614c4d565b82800160010185558215614c4d579182015b82811115614c4d578235825591602001919060010190614cb6565b5b80821115614c595760008155600101614cd2565b8035614cf181615906565b919050565b8035614cf1816158f1565b60008083601f840112614d12578182fd5b5081356001600160401b03811115614d28578182fd5b602083019150836020828501011115614d4057600080fd5b9250929050565b600082601f830112614d57578081fd5b8135614d6a614d65826156f2565b6156c2565b818152846020838601011115614d7e578283fd5b816020850160208301379081016020019190915292915050565b8035614cf181615914565b8035614cf181615929565b803563ffffffff81168114614cf157600080fd5b80356001600160601b0381168114614cf157600080fd5b600060208284031215614dea578081fd5b8135614a71816158f1565b600060208284031215614e06578081fd5b8151614a71816158f1565b600080600060608486031215614e25578182fd5b8335614e30816158f1565b92506020840135614e40816158f1565b929592945050506040919091013590565b600060208284031215614e62578081fd5b8151614a7181615906565b600080600060608486031215614e81578081fd5b8335614e8c816158f1565b92506020840135614e9c816158f1565b91506040840135614eac816158f1565b809150509250925092565b60008060408385031215614ec9578182fd5b8235614ed4816158f1565b91506020830135614ee4816158f1565b809150509250929050565b60008060008060808587031215614f04578182fd5b8435614f0f816158f1565b93506020850135614f1f816158f1565b92506040850135614f2f81615914565b9150614f3d60608601614dae565b905092959194509250565b600080600080600060a08688031215614f5f578283fd5b8535614f6a816158f1565b94506020860135614f7a816158f1565b93506040860135614f8a81615914565b9250614f9860608701614dae565b949793965091946080013592915050565b60008060208385031215614fbb578182fd5b82356001600160401b03811115614fd0578283fd5b614fdc85828601614d01565b90969095509350505050565b600060208284031215614ff9578081fd5b81516001600160401b0381111561500e578182fd5b8201601f8101841361501e578182fd5b805161502c614d65826156f2565b818152856020838501011115615040578384fd5b61505182602083016020860161583e565b95945050505050565b600080600080600080600060c0888a031215615074578485fd5b87356001600160401b038082111561508a578687fd5b6150968b838c01614d47565b985060208a01359150808211156150ab578687fd5b506150b88a828b01614d01565b90975095505060408801356150cc81615929565b935060608801356150dc816158f1565b925060808801356150ec816158f1565b915060a08801356150fc816158f1565b8091505092959891949750929550565b6000806000806000806000806000806101408b8d03121561512b578384fd5b8a356001600160401b0380821115615141578586fd5b61514d8e838f01614d47565b9b5060208d0135915080821115615162578586fd5b5061516f8d828e01614d47565b99505061517e60408c01614dae565b975061518c60608c01614d98565b965061519a60808c01614cf6565b95506151a860a08c01614da3565b94506151b660c08c01614dae565b93506151c460e08c01614dae565b92506151d36101008c01614dc2565b91506151e26101208c01614ce6565b90509295989b9194979a5092959850565b600080600060608486031215615207578081fd5b835161521281615914565b602085015190935061522381615914565b6040850151909250614eac81615914565b600060208284031215615245578081fd5b8135614a7181615929565b600060208284031215615261578081fd5b8151614a7181615929565b60006020828403121561527d578081fd5b5035919050565b600060208284031215615295578081fd5b5051919050565b600080600080608085870312156152b1578182fd5b8435935060208501356152c3816158f1565b92506152d160408601614dae565b9396929550929360600135925050565b600080604083850312156152f3578182fd5b50508035926020909101359150565b600060208284031215615313578081fd5b614a7182614dae565b60006020828403121561532d578081fd5b815160ff81168114614a71578182fd5b6000815180845261535581602086016020860161583e565b601f01601f19169290920160200192915050565b6001600160701b03169052565b6000825161538881846020870161583e565b9190910192915050565b600083516153a481846020880161583e565b600160fd1b90830190815283516153c281600184016020880161583e565b01600101949350505050565b6001600160a01b0391909116815260200190565b6001600160a01b039290921682521515602082015260400190565b6001600160a01b039290921682526001600160701b0316602082015260400190565b6001600160a01b03929092168252602082015260400190565b6020808252825182820181905260009190848201906040850190845b818110156154795783516001600160a01b031683529284019291840191600101615454565b50909695505050505050565b6001600160a01b039390931683526001600160701b0391909116602083015263ffffffff16604082015260600190565b6001600160a01b0393841681526020810192909252909116604082015260600190565b6001600160a01b0392831681529116602082015260400190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b602081526000614a71602083018461533d565b6101808152600061554961018083018f61533d565b828103602084015261555b818f61533d565b9150508b604083015263ffffffff808c166060840152808b166080840152808a1660a08401525061ffff881660c08301528660e0830152856101008301526155a7610120830186615369565b6155b5610140830185615369565b63ffffffff83166101608301529d9c50505050505050505050505050565b6060815260006155e6606083018661533d565b82810360208401526155f8818661533d565b91505060ff83166040830152949350505050565b81516001600160701b0316815260208083015161ffff169082015260408083015163ffffffff9081169183019190915260608084015182169083015260808084015182169083015260a080840151918216908301526101008201905060c083015161567a60c0840182615369565b5060e083015161569060e084018261ffff169052565b5092915050565b6001600160701b0392831681529116602082015260400190565b63ffffffff91909116815260200190565b604051601f8201601f191681016001600160401b03811182821017156156ea576156ea6158db565b604052919050565b60006001600160401b0382111561570b5761570b6158db565b50601f01601f191660200190565b60006001600160701b0382811684821680830382111561573b5761573b6158c5565b01949350505050565b60008219821115615757576157576158c5565b500190565b600063ffffffff80831681851680830382111561573b5761573b6158c5565b60008261579657634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156157b5576157b56158c5565b500290565b60008160010b8360010b82811281617fff19018312811516156157df576157df6158c5565b81617fff0183138116156157f5576157f56158c5565b5090039392505050565b60006001600160701b038381169083168181101561581f5761581f6158c5565b039392505050565b600082821015615839576158396158c5565b500390565b60005b83811015615859578181015183820152602001615841565b8381111561436b5750506000910152565b600181811c9082168061587e57607f821691505b6020821081141561589f57634e487b7160e01b600052602260045260246000fd5b50919050565b60008160010b617fff198114156158be576158be6158c5565b9003919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461353257600080fd5b801515811461353257600080fd5b6001600160701b038116811461353257600080fd5b61ffff8116811461353257600080fdfe87cdeaffd8e70903d6ce7cc983fac3b09ca79e83818124c98e47a1d70f8027d6ac59582e5396aca512fa873a2047e7f4c80f8f55d4a06cb34a78a0187f62719f4f471908b72bb76dae5bd24599026e7bf3ddb256497722888ffa422f83729ede5fbe47a2294bb165359bd422c05af7405a9f2c0e0f397deef39ade4da37b27647c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b926b69b9f7735227079a308fce165b111b96e2983e70a86a381ee047675ff39fd1248cccb5fef9131c731321e43e9a924840ffee7dc68c7d1d3e5cb7dedcae0327894f6b9c469a3eddbe98dd0df54563b6437b20982876e0c2cd60c211a156460dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762ddbefcd50991bb64333037879282220d78a901bb4809351643f56628af8a4656dfc8dfab451067f726e1077626e126c7f77149630cc1b0056cf60307a289a118298be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e07aa9698a01d2094fb8a580e2c72a6fc3f3bfecfd4cbdc457461118c59a6fefe4f6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a92a4dc1fa715a8cafb0cdc060f729eb1dfdddcbc9effbdc8d822cdf5ebcf4907971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585a70de76be90dc9c5d9243b0bc73c3d963ba703a0a396f29134bf5b4e6197c180a2646970667358221220fbe8100e6a7c4692daaa17f9579527f8ce5498f05782b57298a53f0b93aa40ff64736f6c63430008040033",
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

// EstimateLoan is a free data retrieval call binding the contract method 0x229304dc.
//
// Solidity: function estimateLoan(address powerToken, address paymentToken, uint112 amount, uint32 duration) view returns(uint256)
func (_Enterprise *EnterpriseCaller) EstimateLoan(opts *bind.CallOpts, powerToken common.Address, paymentToken common.Address, amount *big.Int, duration uint32) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "estimateLoan", powerToken, paymentToken, amount, duration)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateLoan is a free data retrieval call binding the contract method 0x229304dc.
//
// Solidity: function estimateLoan(address powerToken, address paymentToken, uint112 amount, uint32 duration) view returns(uint256)
func (_Enterprise *EnterpriseSession) EstimateLoan(powerToken common.Address, paymentToken common.Address, amount *big.Int, duration uint32) (*big.Int, error) {
	return _Enterprise.Contract.EstimateLoan(&_Enterprise.CallOpts, powerToken, paymentToken, amount, duration)
}

// EstimateLoan is a free data retrieval call binding the contract method 0x229304dc.
//
// Solidity: function estimateLoan(address powerToken, address paymentToken, uint112 amount, uint32 duration) view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) EstimateLoan(powerToken common.Address, paymentToken common.Address, amount *big.Int, duration uint32) (*big.Int, error) {
	return _Enterprise.Contract.EstimateLoan(&_Enterprise.CallOpts, powerToken, paymentToken, amount, duration)
}

// GetAccruedInterest is a free data retrieval call binding the contract method 0x9339ea69.
//
// Solidity: function getAccruedInterest(uint256 interestTokenId) view returns(uint256)
func (_Enterprise *EnterpriseCaller) GetAccruedInterest(opts *bind.CallOpts, interestTokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getAccruedInterest", interestTokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccruedInterest is a free data retrieval call binding the contract method 0x9339ea69.
//
// Solidity: function getAccruedInterest(uint256 interestTokenId) view returns(uint256)
func (_Enterprise *EnterpriseSession) GetAccruedInterest(interestTokenId *big.Int) (*big.Int, error) {
	return _Enterprise.Contract.GetAccruedInterest(&_Enterprise.CallOpts, interestTokenId)
}

// GetAccruedInterest is a free data retrieval call binding the contract method 0x9339ea69.
//
// Solidity: function getAccruedInterest(uint256 interestTokenId) view returns(uint256)
func (_Enterprise *EnterpriseCallerSession) GetAccruedInterest(interestTokenId *big.Int) (*big.Int, error) {
	return _Enterprise.Contract.GetAccruedInterest(&_Enterprise.CallOpts, interestTokenId)
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

// GetBorrowToken is a free data retrieval call binding the contract method 0xc93d0b92.
//
// Solidity: function getBorrowToken() view returns(address)
func (_Enterprise *EnterpriseCaller) GetBorrowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getBorrowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowToken is a free data retrieval call binding the contract method 0xc93d0b92.
//
// Solidity: function getBorrowToken() view returns(address)
func (_Enterprise *EnterpriseSession) GetBorrowToken() (common.Address, error) {
	return _Enterprise.Contract.GetBorrowToken(&_Enterprise.CallOpts)
}

// GetBorrowToken is a free data retrieval call binding the contract method 0xc93d0b92.
//
// Solidity: function getBorrowToken() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetBorrowToken() (common.Address, error) {
	return _Enterprise.Contract.GetBorrowToken(&_Enterprise.CallOpts)
}

// GetBorrowerLoanReturnGracePeriod is a free data retrieval call binding the contract method 0x8188d50d.
//
// Solidity: function getBorrowerLoanReturnGracePeriod() view returns(uint32)
func (_Enterprise *EnterpriseCaller) GetBorrowerLoanReturnGracePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getBorrowerLoanReturnGracePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetBorrowerLoanReturnGracePeriod is a free data retrieval call binding the contract method 0x8188d50d.
//
// Solidity: function getBorrowerLoanReturnGracePeriod() view returns(uint32)
func (_Enterprise *EnterpriseSession) GetBorrowerLoanReturnGracePeriod() (uint32, error) {
	return _Enterprise.Contract.GetBorrowerLoanReturnGracePeriod(&_Enterprise.CallOpts)
}

// GetBorrowerLoanReturnGracePeriod is a free data retrieval call binding the contract method 0x8188d50d.
//
// Solidity: function getBorrowerLoanReturnGracePeriod() view returns(uint32)
func (_Enterprise *EnterpriseCallerSession) GetBorrowerLoanReturnGracePeriod() (uint32, error) {
	return _Enterprise.Contract.GetBorrowerLoanReturnGracePeriod(&_Enterprise.CallOpts)
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

// GetEnterpriseLoanCollectGracePeriod is a free data retrieval call binding the contract method 0xa0efd118.
//
// Solidity: function getEnterpriseLoanCollectGracePeriod() view returns(uint32)
func (_Enterprise *EnterpriseCaller) GetEnterpriseLoanCollectGracePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getEnterpriseLoanCollectGracePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEnterpriseLoanCollectGracePeriod is a free data retrieval call binding the contract method 0xa0efd118.
//
// Solidity: function getEnterpriseLoanCollectGracePeriod() view returns(uint32)
func (_Enterprise *EnterpriseSession) GetEnterpriseLoanCollectGracePeriod() (uint32, error) {
	return _Enterprise.Contract.GetEnterpriseLoanCollectGracePeriod(&_Enterprise.CallOpts)
}

// GetEnterpriseLoanCollectGracePeriod is a free data retrieval call binding the contract method 0xa0efd118.
//
// Solidity: function getEnterpriseLoanCollectGracePeriod() view returns(uint32)
func (_Enterprise *EnterpriseCallerSession) GetEnterpriseLoanCollectGracePeriod() (uint32, error) {
	return _Enterprise.Contract.GetEnterpriseLoanCollectGracePeriod(&_Enterprise.CallOpts)
}

// GetEnterpriseVault is a free data retrieval call binding the contract method 0x7791c42d.
//
// Solidity: function getEnterpriseVault() view returns(address)
func (_Enterprise *EnterpriseCaller) GetEnterpriseVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getEnterpriseVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseVault is a free data retrieval call binding the contract method 0x7791c42d.
//
// Solidity: function getEnterpriseVault() view returns(address)
func (_Enterprise *EnterpriseSession) GetEnterpriseVault() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseVault(&_Enterprise.CallOpts)
}

// GetEnterpriseVault is a free data retrieval call binding the contract method 0x7791c42d.
//
// Solidity: function getEnterpriseVault() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetEnterpriseVault() (common.Address, error) {
	return _Enterprise.Contract.GetEnterpriseVault(&_Enterprise.CallOpts)
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
// Solidity: function getInfo() view returns(string name, string baseUri, uint256 totalShares, uint32 interestGapHalvingPeriod, uint32 borrowerLoanReturnGracePeriod, uint32 enterpriseLoanCollectGracePeriod, uint16 gcFeePercent, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_Enterprise *EnterpriseCaller) GetInfo(opts *bind.CallOpts) (struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GcFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getInfo")

	outstruct := new(struct {
		Name                             string
		BaseUri                          string
		TotalShares                      *big.Int
		InterestGapHalvingPeriod         uint32
		BorrowerLoanReturnGracePeriod    uint32
		EnterpriseLoanCollectGracePeriod uint32
		GcFeePercent                     uint16
		FixedReserve                     *big.Int
		UsedReserve                      *big.Int
		StreamingReserve                 *big.Int
		StreamingReserveTarget           *big.Int
		StreamingReserveUpdated          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.BaseUri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TotalShares = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InterestGapHalvingPeriod = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.BorrowerLoanReturnGracePeriod = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.EnterpriseLoanCollectGracePeriod = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.GcFeePercent = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.FixedReserve = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsedReserve = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserve = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserveTarget = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.StreamingReserveUpdated = *abi.ConvertType(out[11], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint256 totalShares, uint32 interestGapHalvingPeriod, uint32 borrowerLoanReturnGracePeriod, uint32 enterpriseLoanCollectGracePeriod, uint16 gcFeePercent, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_Enterprise *EnterpriseSession) GetInfo() (struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GcFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}, error) {
	return _Enterprise.Contract.GetInfo(&_Enterprise.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(string name, string baseUri, uint256 totalShares, uint32 interestGapHalvingPeriod, uint32 borrowerLoanReturnGracePeriod, uint32 enterpriseLoanCollectGracePeriod, uint16 gcFeePercent, uint256 fixedReserve, uint256 usedReserve, uint112 streamingReserve, uint112 streamingReserveTarget, uint32 streamingReserveUpdated)
func (_Enterprise *EnterpriseCallerSession) GetInfo() (struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GcFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}, error) {
	return _Enterprise.Contract.GetInfo(&_Enterprise.CallOpts)
}

// GetInterestGapHalvingPeriod is a free data retrieval call binding the contract method 0x5a4598df.
//
// Solidity: function getInterestGapHalvingPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCaller) GetInterestGapHalvingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getInterestGapHalvingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetInterestGapHalvingPeriod is a free data retrieval call binding the contract method 0x5a4598df.
//
// Solidity: function getInterestGapHalvingPeriod() view returns(uint32)
func (_Enterprise *EnterpriseSession) GetInterestGapHalvingPeriod() (uint32, error) {
	return _Enterprise.Contract.GetInterestGapHalvingPeriod(&_Enterprise.CallOpts)
}

// GetInterestGapHalvingPeriod is a free data retrieval call binding the contract method 0x5a4598df.
//
// Solidity: function getInterestGapHalvingPeriod() view returns(uint32)
func (_Enterprise *EnterpriseCallerSession) GetInterestGapHalvingPeriod() (uint32, error) {
	return _Enterprise.Contract.GetInterestGapHalvingPeriod(&_Enterprise.CallOpts)
}

// GetInterestToken is a free data retrieval call binding the contract method 0x0ea1dc9e.
//
// Solidity: function getInterestToken() view returns(address)
func (_Enterprise *EnterpriseCaller) GetInterestToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getInterestToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestToken is a free data retrieval call binding the contract method 0x0ea1dc9e.
//
// Solidity: function getInterestToken() view returns(address)
func (_Enterprise *EnterpriseSession) GetInterestToken() (common.Address, error) {
	return _Enterprise.Contract.GetInterestToken(&_Enterprise.CallOpts)
}

// GetInterestToken is a free data retrieval call binding the contract method 0x0ea1dc9e.
//
// Solidity: function getInterestToken() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetInterestToken() (common.Address, error) {
	return _Enterprise.Contract.GetInterestToken(&_Enterprise.CallOpts)
}

// GetLiquidityInfo is a free data retrieval call binding the contract method 0xdf57f2fa.
//
// Solidity: function getLiquidityInfo(uint256 interestTokenId) view returns((uint256,uint256,uint256))
func (_Enterprise *EnterpriseCaller) GetLiquidityInfo(opts *bind.CallOpts, interestTokenId *big.Int) (EnterpriseStorageLiquidityInfo, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getLiquidityInfo", interestTokenId)

	if err != nil {
		return *new(EnterpriseStorageLiquidityInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EnterpriseStorageLiquidityInfo)).(*EnterpriseStorageLiquidityInfo)

	return out0, err

}

// GetLiquidityInfo is a free data retrieval call binding the contract method 0xdf57f2fa.
//
// Solidity: function getLiquidityInfo(uint256 interestTokenId) view returns((uint256,uint256,uint256))
func (_Enterprise *EnterpriseSession) GetLiquidityInfo(interestTokenId *big.Int) (EnterpriseStorageLiquidityInfo, error) {
	return _Enterprise.Contract.GetLiquidityInfo(&_Enterprise.CallOpts, interestTokenId)
}

// GetLiquidityInfo is a free data retrieval call binding the contract method 0xdf57f2fa.
//
// Solidity: function getLiquidityInfo(uint256 interestTokenId) view returns((uint256,uint256,uint256))
func (_Enterprise *EnterpriseCallerSession) GetLiquidityInfo(interestTokenId *big.Int) (EnterpriseStorageLiquidityInfo, error) {
	return _Enterprise.Contract.GetLiquidityInfo(&_Enterprise.CallOpts, interestTokenId)
}

// GetLiquidityToken is a free data retrieval call binding the contract method 0x644e44df.
//
// Solidity: function getLiquidityToken() view returns(address)
func (_Enterprise *EnterpriseCaller) GetLiquidityToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getLiquidityToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidityToken is a free data retrieval call binding the contract method 0x644e44df.
//
// Solidity: function getLiquidityToken() view returns(address)
func (_Enterprise *EnterpriseSession) GetLiquidityToken() (common.Address, error) {
	return _Enterprise.Contract.GetLiquidityToken(&_Enterprise.CallOpts)
}

// GetLiquidityToken is a free data retrieval call binding the contract method 0x644e44df.
//
// Solidity: function getLiquidityToken() view returns(address)
func (_Enterprise *EnterpriseCallerSession) GetLiquidityToken() (common.Address, error) {
	return _Enterprise.Contract.GetLiquidityToken(&_Enterprise.CallOpts)
}

// GetLoanInfo is a free data retrieval call binding the contract method 0x0714012c.
//
// Solidity: function getLoanInfo(uint256 borrowTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_Enterprise *EnterpriseCaller) GetLoanInfo(opts *bind.CallOpts, borrowTokenId *big.Int) (EnterpriseStorageLoanInfo, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "getLoanInfo", borrowTokenId)

	if err != nil {
		return *new(EnterpriseStorageLoanInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EnterpriseStorageLoanInfo)).(*EnterpriseStorageLoanInfo)

	return out0, err

}

// GetLoanInfo is a free data retrieval call binding the contract method 0x0714012c.
//
// Solidity: function getLoanInfo(uint256 borrowTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_Enterprise *EnterpriseSession) GetLoanInfo(borrowTokenId *big.Int) (EnterpriseStorageLoanInfo, error) {
	return _Enterprise.Contract.GetLoanInfo(&_Enterprise.CallOpts, borrowTokenId)
}

// GetLoanInfo is a free data retrieval call binding the contract method 0x0714012c.
//
// Solidity: function getLoanInfo(uint256 borrowTokenId) view returns((uint112,uint16,uint32,uint32,uint32,uint32,uint112,uint16))
func (_Enterprise *EnterpriseCallerSession) GetLoanInfo(borrowTokenId *big.Int) (EnterpriseStorageLoanInfo, error) {
	return _Enterprise.Contract.GetLoanInfo(&_Enterprise.CallOpts, borrowTokenId)
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

// PaymentToken is a free data retrieval call binding the contract method 0x0e101b60.
//
// Solidity: function paymentToken(uint256 index) view returns(address)
func (_Enterprise *EnterpriseCaller) PaymentToken(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "paymentToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x0e101b60.
//
// Solidity: function paymentToken(uint256 index) view returns(address)
func (_Enterprise *EnterpriseSession) PaymentToken(index *big.Int) (common.Address, error) {
	return _Enterprise.Contract.PaymentToken(&_Enterprise.CallOpts, index)
}

// PaymentToken is a free data retrieval call binding the contract method 0x0e101b60.
//
// Solidity: function paymentToken(uint256 index) view returns(address)
func (_Enterprise *EnterpriseCallerSession) PaymentToken(index *big.Int) (common.Address, error) {
	return _Enterprise.Contract.PaymentToken(&_Enterprise.CallOpts, index)
}

// PaymentTokenIndex is a free data retrieval call binding the contract method 0xfbac9a55.
//
// Solidity: function paymentTokenIndex(address token) view returns(int16)
func (_Enterprise *EnterpriseCaller) PaymentTokenIndex(opts *bind.CallOpts, token common.Address) (int16, error) {
	var out []interface{}
	err := _Enterprise.contract.Call(opts, &out, "paymentTokenIndex", token)

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// PaymentTokenIndex is a free data retrieval call binding the contract method 0xfbac9a55.
//
// Solidity: function paymentTokenIndex(address token) view returns(int16)
func (_Enterprise *EnterpriseSession) PaymentTokenIndex(token common.Address) (int16, error) {
	return _Enterprise.Contract.PaymentTokenIndex(&_Enterprise.CallOpts, token)
}

// PaymentTokenIndex is a free data retrieval call binding the contract method 0xfbac9a55.
//
// Solidity: function paymentTokenIndex(address token) view returns(int16)
func (_Enterprise *EnterpriseCallerSession) PaymentTokenIndex(token common.Address) (int16, error) {
	return _Enterprise.Contract.PaymentTokenIndex(&_Enterprise.CallOpts, token)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 liquidityAmount) returns()
func (_Enterprise *EnterpriseTransactor) AddLiquidity(opts *bind.TransactOpts, liquidityAmount *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "addLiquidity", liquidityAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 liquidityAmount) returns()
func (_Enterprise *EnterpriseSession) AddLiquidity(liquidityAmount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.AddLiquidity(&_Enterprise.TransactOpts, liquidityAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 liquidityAmount) returns()
func (_Enterprise *EnterpriseTransactorSession) AddLiquidity(liquidityAmount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.AddLiquidity(&_Enterprise.TransactOpts, liquidityAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x00ca339b.
//
// Solidity: function borrow(address powerToken, address paymentToken, uint112 amount, uint32 duration, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactor) Borrow(opts *bind.TransactOpts, powerToken common.Address, paymentToken common.Address, amount *big.Int, duration uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "borrow", powerToken, paymentToken, amount, duration, maxPayment)
}

// Borrow is a paid mutator transaction binding the contract method 0x00ca339b.
//
// Solidity: function borrow(address powerToken, address paymentToken, uint112 amount, uint32 duration, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseSession) Borrow(powerToken common.Address, paymentToken common.Address, amount *big.Int, duration uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Borrow(&_Enterprise.TransactOpts, powerToken, paymentToken, amount, duration, maxPayment)
}

// Borrow is a paid mutator transaction binding the contract method 0x00ca339b.
//
// Solidity: function borrow(address powerToken, address paymentToken, uint112 amount, uint32 duration, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactorSession) Borrow(powerToken common.Address, paymentToken common.Address, amount *big.Int, duration uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Borrow(&_Enterprise.TransactOpts, powerToken, paymentToken, amount, duration, maxPayment)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x49bd4e89.
//
// Solidity: function decreaseLiquidity(uint256 interestTokenId, uint256 amount) returns()
func (_Enterprise *EnterpriseTransactor) DecreaseLiquidity(opts *bind.TransactOpts, interestTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "decreaseLiquidity", interestTokenId, amount)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x49bd4e89.
//
// Solidity: function decreaseLiquidity(uint256 interestTokenId, uint256 amount) returns()
func (_Enterprise *EnterpriseSession) DecreaseLiquidity(interestTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.DecreaseLiquidity(&_Enterprise.TransactOpts, interestTokenId, amount)
}

// DecreaseLiquidity is a paid mutator transaction binding the contract method 0x49bd4e89.
//
// Solidity: function decreaseLiquidity(uint256 interestTokenId, uint256 amount) returns()
func (_Enterprise *EnterpriseTransactorSession) DecreaseLiquidity(interestTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.DecreaseLiquidity(&_Enterprise.TransactOpts, interestTokenId, amount)
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

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x696f9c81.
//
// Solidity: function increaseLiquidity(uint256 interestTokenId, uint256 amount) returns()
func (_Enterprise *EnterpriseTransactor) IncreaseLiquidity(opts *bind.TransactOpts, interestTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "increaseLiquidity", interestTokenId, amount)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x696f9c81.
//
// Solidity: function increaseLiquidity(uint256 interestTokenId, uint256 amount) returns()
func (_Enterprise *EnterpriseSession) IncreaseLiquidity(interestTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.IncreaseLiquidity(&_Enterprise.TransactOpts, interestTokenId, amount)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0x696f9c81.
//
// Solidity: function increaseLiquidity(uint256 interestTokenId, uint256 amount) returns()
func (_Enterprise *EnterpriseTransactorSession) IncreaseLiquidity(interestTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.IncreaseLiquidity(&_Enterprise.TransactOpts, interestTokenId, amount)
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
// Solidity: function initializeTokens(address liquidityToken, address interestToken, address borrowToken) returns()
func (_Enterprise *EnterpriseTransactor) InitializeTokens(opts *bind.TransactOpts, liquidityToken common.Address, interestToken common.Address, borrowToken common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "initializeTokens", liquidityToken, interestToken, borrowToken)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address liquidityToken, address interestToken, address borrowToken) returns()
func (_Enterprise *EnterpriseSession) InitializeTokens(liquidityToken common.Address, interestToken common.Address, borrowToken common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.InitializeTokens(&_Enterprise.TransactOpts, liquidityToken, interestToken, borrowToken)
}

// InitializeTokens is a paid mutator transaction binding the contract method 0xef1f9f39.
//
// Solidity: function initializeTokens(address liquidityToken, address interestToken, address borrowToken) returns()
func (_Enterprise *EnterpriseTransactorSession) InitializeTokens(liquidityToken common.Address, interestToken common.Address, borrowToken common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.InitializeTokens(&_Enterprise.TransactOpts, liquidityToken, interestToken, borrowToken)
}

// LoanTransfer is a paid mutator transaction binding the contract method 0x512a53bc.
//
// Solidity: function loanTransfer(address from, address to, uint256 borrowTokenId) returns()
func (_Enterprise *EnterpriseTransactor) LoanTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, borrowTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "loanTransfer", from, to, borrowTokenId)
}

// LoanTransfer is a paid mutator transaction binding the contract method 0x512a53bc.
//
// Solidity: function loanTransfer(address from, address to, uint256 borrowTokenId) returns()
func (_Enterprise *EnterpriseSession) LoanTransfer(from common.Address, to common.Address, borrowTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.LoanTransfer(&_Enterprise.TransactOpts, from, to, borrowTokenId)
}

// LoanTransfer is a paid mutator transaction binding the contract method 0x512a53bc.
//
// Solidity: function loanTransfer(address from, address to, uint256 borrowTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) LoanTransfer(from common.Address, to common.Address, borrowTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.LoanTransfer(&_Enterprise.TransactOpts, from, to, borrowTokenId)
}

// Reborrow is a paid mutator transaction binding the contract method 0x8c0114b0.
//
// Solidity: function reborrow(uint256 borrowTokenId, address paymentToken, uint32 duration, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactor) Reborrow(opts *bind.TransactOpts, borrowTokenId *big.Int, paymentToken common.Address, duration uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "reborrow", borrowTokenId, paymentToken, duration, maxPayment)
}

// Reborrow is a paid mutator transaction binding the contract method 0x8c0114b0.
//
// Solidity: function reborrow(uint256 borrowTokenId, address paymentToken, uint32 duration, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseSession) Reborrow(borrowTokenId *big.Int, paymentToken common.Address, duration uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Reborrow(&_Enterprise.TransactOpts, borrowTokenId, paymentToken, duration, maxPayment)
}

// Reborrow is a paid mutator transaction binding the contract method 0x8c0114b0.
//
// Solidity: function reborrow(uint256 borrowTokenId, address paymentToken, uint32 duration, uint256 maxPayment) returns()
func (_Enterprise *EnterpriseTransactorSession) Reborrow(borrowTokenId *big.Int, paymentToken common.Address, duration uint32, maxPayment *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.Reborrow(&_Enterprise.TransactOpts, borrowTokenId, paymentToken, duration, maxPayment)
}

// RegisterService is a paid mutator transaction binding the contract method 0x9ab71193.
//
// Solidity: function registerService(string serviceName, string symbol, uint32 gapHalvingPeriod, uint112 baseRate, address baseToken, uint16 serviceFeePercent, uint32 minLoanDuration, uint32 maxLoanDuration, uint96 minGCFee, bool allowsPerpetualTokensForever) returns()
func (_Enterprise *EnterpriseTransactor) RegisterService(opts *bind.TransactOpts, serviceName string, symbol string, gapHalvingPeriod uint32, baseRate *big.Int, baseToken common.Address, serviceFeePercent uint16, minLoanDuration uint32, maxLoanDuration uint32, minGCFee *big.Int, allowsPerpetualTokensForever bool) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "registerService", serviceName, symbol, gapHalvingPeriod, baseRate, baseToken, serviceFeePercent, minLoanDuration, maxLoanDuration, minGCFee, allowsPerpetualTokensForever)
}

// RegisterService is a paid mutator transaction binding the contract method 0x9ab71193.
//
// Solidity: function registerService(string serviceName, string symbol, uint32 gapHalvingPeriod, uint112 baseRate, address baseToken, uint16 serviceFeePercent, uint32 minLoanDuration, uint32 maxLoanDuration, uint96 minGCFee, bool allowsPerpetualTokensForever) returns()
func (_Enterprise *EnterpriseSession) RegisterService(serviceName string, symbol string, gapHalvingPeriod uint32, baseRate *big.Int, baseToken common.Address, serviceFeePercent uint16, minLoanDuration uint32, maxLoanDuration uint32, minGCFee *big.Int, allowsPerpetualTokensForever bool) (*types.Transaction, error) {
	return _Enterprise.Contract.RegisterService(&_Enterprise.TransactOpts, serviceName, symbol, gapHalvingPeriod, baseRate, baseToken, serviceFeePercent, minLoanDuration, maxLoanDuration, minGCFee, allowsPerpetualTokensForever)
}

// RegisterService is a paid mutator transaction binding the contract method 0x9ab71193.
//
// Solidity: function registerService(string serviceName, string symbol, uint32 gapHalvingPeriod, uint112 baseRate, address baseToken, uint16 serviceFeePercent, uint32 minLoanDuration, uint32 maxLoanDuration, uint96 minGCFee, bool allowsPerpetualTokensForever) returns()
func (_Enterprise *EnterpriseTransactorSession) RegisterService(serviceName string, symbol string, gapHalvingPeriod uint32, baseRate *big.Int, baseToken common.Address, serviceFeePercent uint16, minLoanDuration uint32, maxLoanDuration uint32, minGCFee *big.Int, allowsPerpetualTokensForever bool) (*types.Transaction, error) {
	return _Enterprise.Contract.RegisterService(&_Enterprise.TransactOpts, serviceName, symbol, gapHalvingPeriod, baseRate, baseToken, serviceFeePercent, minLoanDuration, maxLoanDuration, minGCFee, allowsPerpetualTokensForever)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 interestTokenId) returns()
func (_Enterprise *EnterpriseTransactor) RemoveLiquidity(opts *bind.TransactOpts, interestTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "removeLiquidity", interestTokenId)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 interestTokenId) returns()
func (_Enterprise *EnterpriseSession) RemoveLiquidity(interestTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.RemoveLiquidity(&_Enterprise.TransactOpts, interestTokenId)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 interestTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) RemoveLiquidity(interestTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.RemoveLiquidity(&_Enterprise.TransactOpts, interestTokenId)
}

// ReturnLoan is a paid mutator transaction binding the contract method 0xe9126154.
//
// Solidity: function returnLoan(uint256 borrowTokenId) returns()
func (_Enterprise *EnterpriseTransactor) ReturnLoan(opts *bind.TransactOpts, borrowTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "returnLoan", borrowTokenId)
}

// ReturnLoan is a paid mutator transaction binding the contract method 0xe9126154.
//
// Solidity: function returnLoan(uint256 borrowTokenId) returns()
func (_Enterprise *EnterpriseSession) ReturnLoan(borrowTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ReturnLoan(&_Enterprise.TransactOpts, borrowTokenId)
}

// ReturnLoan is a paid mutator transaction binding the contract method 0xe9126154.
//
// Solidity: function returnLoan(uint256 borrowTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) ReturnLoan(borrowTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.ReturnLoan(&_Enterprise.TransactOpts, borrowTokenId)
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

// SetBorrowerLoanReturnGracePeriod is a paid mutator transaction binding the contract method 0x2e365ade.
//
// Solidity: function setBorrowerLoanReturnGracePeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactor) SetBorrowerLoanReturnGracePeriod(opts *bind.TransactOpts, newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setBorrowerLoanReturnGracePeriod", newPeriod)
}

// SetBorrowerLoanReturnGracePeriod is a paid mutator transaction binding the contract method 0x2e365ade.
//
// Solidity: function setBorrowerLoanReturnGracePeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseSession) SetBorrowerLoanReturnGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetBorrowerLoanReturnGracePeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetBorrowerLoanReturnGracePeriod is a paid mutator transaction binding the contract method 0x2e365ade.
//
// Solidity: function setBorrowerLoanReturnGracePeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactorSession) SetBorrowerLoanReturnGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetBorrowerLoanReturnGracePeriod(&_Enterprise.TransactOpts, newPeriod)
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

// SetEnterpriseLoanCollectGracePeriod is a paid mutator transaction binding the contract method 0xb56de97e.
//
// Solidity: function setEnterpriseLoanCollectGracePeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactor) SetEnterpriseLoanCollectGracePeriod(opts *bind.TransactOpts, newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setEnterpriseLoanCollectGracePeriod", newPeriod)
}

// SetEnterpriseLoanCollectGracePeriod is a paid mutator transaction binding the contract method 0xb56de97e.
//
// Solidity: function setEnterpriseLoanCollectGracePeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseSession) SetEnterpriseLoanCollectGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseLoanCollectGracePeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetEnterpriseLoanCollectGracePeriod is a paid mutator transaction binding the contract method 0xb56de97e.
//
// Solidity: function setEnterpriseLoanCollectGracePeriod(uint32 newPeriod) returns()
func (_Enterprise *EnterpriseTransactorSession) SetEnterpriseLoanCollectGracePeriod(newPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseLoanCollectGracePeriod(&_Enterprise.TransactOpts, newPeriod)
}

// SetEnterpriseVault is a paid mutator transaction binding the contract method 0x90f55edd.
//
// Solidity: function setEnterpriseVault(address newVault) returns()
func (_Enterprise *EnterpriseTransactor) SetEnterpriseVault(opts *bind.TransactOpts, newVault common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setEnterpriseVault", newVault)
}

// SetEnterpriseVault is a paid mutator transaction binding the contract method 0x90f55edd.
//
// Solidity: function setEnterpriseVault(address newVault) returns()
func (_Enterprise *EnterpriseSession) SetEnterpriseVault(newVault common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseVault(&_Enterprise.TransactOpts, newVault)
}

// SetEnterpriseVault is a paid mutator transaction binding the contract method 0x90f55edd.
//
// Solidity: function setEnterpriseVault(address newVault) returns()
func (_Enterprise *EnterpriseTransactorSession) SetEnterpriseVault(newVault common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.SetEnterpriseVault(&_Enterprise.TransactOpts, newVault)
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

// SetInterestGapHalvingPeriod is a paid mutator transaction binding the contract method 0xe1a5fa49.
//
// Solidity: function setInterestGapHalvingPeriod(uint32 interestGapHalvingPeriod) returns()
func (_Enterprise *EnterpriseTransactor) SetInterestGapHalvingPeriod(opts *bind.TransactOpts, interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "setInterestGapHalvingPeriod", interestGapHalvingPeriod)
}

// SetInterestGapHalvingPeriod is a paid mutator transaction binding the contract method 0xe1a5fa49.
//
// Solidity: function setInterestGapHalvingPeriod(uint32 interestGapHalvingPeriod) returns()
func (_Enterprise *EnterpriseSession) SetInterestGapHalvingPeriod(interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetInterestGapHalvingPeriod(&_Enterprise.TransactOpts, interestGapHalvingPeriod)
}

// SetInterestGapHalvingPeriod is a paid mutator transaction binding the contract method 0xe1a5fa49.
//
// Solidity: function setInterestGapHalvingPeriod(uint32 interestGapHalvingPeriod) returns()
func (_Enterprise *EnterpriseTransactorSession) SetInterestGapHalvingPeriod(interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return _Enterprise.Contract.SetInterestGapHalvingPeriod(&_Enterprise.TransactOpts, interestGapHalvingPeriod)
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

// UpgradeBorrowToken is a paid mutator transaction binding the contract method 0x7b024363.
//
// Solidity: function upgradeBorrowToken(address implementation) returns()
func (_Enterprise *EnterpriseTransactor) UpgradeBorrowToken(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "upgradeBorrowToken", implementation)
}

// UpgradeBorrowToken is a paid mutator transaction binding the contract method 0x7b024363.
//
// Solidity: function upgradeBorrowToken(address implementation) returns()
func (_Enterprise *EnterpriseSession) UpgradeBorrowToken(implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradeBorrowToken(&_Enterprise.TransactOpts, implementation)
}

// UpgradeBorrowToken is a paid mutator transaction binding the contract method 0x7b024363.
//
// Solidity: function upgradeBorrowToken(address implementation) returns()
func (_Enterprise *EnterpriseTransactorSession) UpgradeBorrowToken(implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradeBorrowToken(&_Enterprise.TransactOpts, implementation)
}

// UpgradeEnterprise is a paid mutator transaction binding the contract method 0xc1fc16d8.
//
// Solidity: function upgradeEnterprise(address implementation) returns()
func (_Enterprise *EnterpriseTransactor) UpgradeEnterprise(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "upgradeEnterprise", implementation)
}

// UpgradeEnterprise is a paid mutator transaction binding the contract method 0xc1fc16d8.
//
// Solidity: function upgradeEnterprise(address implementation) returns()
func (_Enterprise *EnterpriseSession) UpgradeEnterprise(implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradeEnterprise(&_Enterprise.TransactOpts, implementation)
}

// UpgradeEnterprise is a paid mutator transaction binding the contract method 0xc1fc16d8.
//
// Solidity: function upgradeEnterprise(address implementation) returns()
func (_Enterprise *EnterpriseTransactorSession) UpgradeEnterprise(implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradeEnterprise(&_Enterprise.TransactOpts, implementation)
}

// UpgradeInterestToken is a paid mutator transaction binding the contract method 0x2b2f6095.
//
// Solidity: function upgradeInterestToken(address implementation) returns()
func (_Enterprise *EnterpriseTransactor) UpgradeInterestToken(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "upgradeInterestToken", implementation)
}

// UpgradeInterestToken is a paid mutator transaction binding the contract method 0x2b2f6095.
//
// Solidity: function upgradeInterestToken(address implementation) returns()
func (_Enterprise *EnterpriseSession) UpgradeInterestToken(implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradeInterestToken(&_Enterprise.TransactOpts, implementation)
}

// UpgradeInterestToken is a paid mutator transaction binding the contract method 0x2b2f6095.
//
// Solidity: function upgradeInterestToken(address implementation) returns()
func (_Enterprise *EnterpriseTransactorSession) UpgradeInterestToken(implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradeInterestToken(&_Enterprise.TransactOpts, implementation)
}

// UpgradePowerToken is a paid mutator transaction binding the contract method 0xc2294784.
//
// Solidity: function upgradePowerToken(address powerToken, address implementation) returns()
func (_Enterprise *EnterpriseTransactor) UpgradePowerToken(opts *bind.TransactOpts, powerToken common.Address, implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "upgradePowerToken", powerToken, implementation)
}

// UpgradePowerToken is a paid mutator transaction binding the contract method 0xc2294784.
//
// Solidity: function upgradePowerToken(address powerToken, address implementation) returns()
func (_Enterprise *EnterpriseSession) UpgradePowerToken(powerToken common.Address, implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradePowerToken(&_Enterprise.TransactOpts, powerToken, implementation)
}

// UpgradePowerToken is a paid mutator transaction binding the contract method 0xc2294784.
//
// Solidity: function upgradePowerToken(address powerToken, address implementation) returns()
func (_Enterprise *EnterpriseTransactorSession) UpgradePowerToken(powerToken common.Address, implementation common.Address) (*types.Transaction, error) {
	return _Enterprise.Contract.UpgradePowerToken(&_Enterprise.TransactOpts, powerToken, implementation)
}

// WithdrawInterest is a paid mutator transaction binding the contract method 0x08669aab.
//
// Solidity: function withdrawInterest(uint256 interestTokenId) returns()
func (_Enterprise *EnterpriseTransactor) WithdrawInterest(opts *bind.TransactOpts, interestTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.contract.Transact(opts, "withdrawInterest", interestTokenId)
}

// WithdrawInterest is a paid mutator transaction binding the contract method 0x08669aab.
//
// Solidity: function withdrawInterest(uint256 interestTokenId) returns()
func (_Enterprise *EnterpriseSession) WithdrawInterest(interestTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.WithdrawInterest(&_Enterprise.TransactOpts, interestTokenId)
}

// WithdrawInterest is a paid mutator transaction binding the contract method 0x08669aab.
//
// Solidity: function withdrawInterest(uint256 interestTokenId) returns()
func (_Enterprise *EnterpriseTransactorSession) WithdrawInterest(interestTokenId *big.Int) (*types.Transaction, error) {
	return _Enterprise.Contract.WithdrawInterest(&_Enterprise.TransactOpts, interestTokenId)
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

// EnterpriseBorrowedIterator is returned from FilterBorrowed and is used to iterate over the raw logs and unpacked data for Borrowed events raised by the Enterprise contract.
type EnterpriseBorrowedIterator struct {
	Event *EnterpriseBorrowed // Event containing the contract specifics and raw log

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
func (it *EnterpriseBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseBorrowed)
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
		it.Event = new(EnterpriseBorrowed)
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
func (it *EnterpriseBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseBorrowed represents a Borrowed event raised by the Enterprise contract.
type EnterpriseBorrowed struct {
	PowerToken    common.Address
	BorrowTokenId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrowed is a free log retrieval operation binding the contract event 0xac59582e5396aca512fa873a2047e7f4c80f8f55d4a06cb34a78a0187f62719f.
//
// Solidity: event Borrowed(address indexed powerToken, uint256 indexed borrowTokenId)
func (_Enterprise *EnterpriseFilterer) FilterBorrowed(opts *bind.FilterOpts, powerToken []common.Address, borrowTokenId []*big.Int) (*EnterpriseBorrowedIterator, error) {

	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}
	var borrowTokenIdRule []interface{}
	for _, borrowTokenIdItem := range borrowTokenId {
		borrowTokenIdRule = append(borrowTokenIdRule, borrowTokenIdItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "Borrowed", powerTokenRule, borrowTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseBorrowedIterator{contract: _Enterprise.contract, event: "Borrowed", logs: logs, sub: sub}, nil
}

// WatchBorrowed is a free log subscription operation binding the contract event 0xac59582e5396aca512fa873a2047e7f4c80f8f55d4a06cb34a78a0187f62719f.
//
// Solidity: event Borrowed(address indexed powerToken, uint256 indexed borrowTokenId)
func (_Enterprise *EnterpriseFilterer) WatchBorrowed(opts *bind.WatchOpts, sink chan<- *EnterpriseBorrowed, powerToken []common.Address, borrowTokenId []*big.Int) (event.Subscription, error) {

	var powerTokenRule []interface{}
	for _, powerTokenItem := range powerToken {
		powerTokenRule = append(powerTokenRule, powerTokenItem)
	}
	var borrowTokenIdRule []interface{}
	for _, borrowTokenIdItem := range borrowTokenId {
		borrowTokenIdRule = append(borrowTokenIdRule, borrowTokenIdItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "Borrowed", powerTokenRule, borrowTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseBorrowed)
				if err := _Enterprise.contract.UnpackLog(event, "Borrowed", log); err != nil {
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

// ParseBorrowed is a log parse operation binding the contract event 0xac59582e5396aca512fa873a2047e7f4c80f8f55d4a06cb34a78a0187f62719f.
//
// Solidity: event Borrowed(address indexed powerToken, uint256 indexed borrowTokenId)
func (_Enterprise *EnterpriseFilterer) ParseBorrowed(log types.Log) (*EnterpriseBorrowed, error) {
	event := new(EnterpriseBorrowed)
	if err := _Enterprise.contract.UnpackLog(event, "Borrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseBorrowerLoanReturnGracePeriodChangedIterator is returned from FilterBorrowerLoanReturnGracePeriodChanged and is used to iterate over the raw logs and unpacked data for BorrowerLoanReturnGracePeriodChanged events raised by the Enterprise contract.
type EnterpriseBorrowerLoanReturnGracePeriodChangedIterator struct {
	Event *EnterpriseBorrowerLoanReturnGracePeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseBorrowerLoanReturnGracePeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseBorrowerLoanReturnGracePeriodChanged)
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
		it.Event = new(EnterpriseBorrowerLoanReturnGracePeriodChanged)
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
func (it *EnterpriseBorrowerLoanReturnGracePeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseBorrowerLoanReturnGracePeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseBorrowerLoanReturnGracePeriodChanged represents a BorrowerLoanReturnGracePeriodChanged event raised by the Enterprise contract.
type EnterpriseBorrowerLoanReturnGracePeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrowerLoanReturnGracePeriodChanged is a free log retrieval operation binding the contract event 0xf6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a.
//
// Solidity: event BorrowerLoanReturnGracePeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) FilterBorrowerLoanReturnGracePeriodChanged(opts *bind.FilterOpts) (*EnterpriseBorrowerLoanReturnGracePeriodChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "BorrowerLoanReturnGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseBorrowerLoanReturnGracePeriodChangedIterator{contract: _Enterprise.contract, event: "BorrowerLoanReturnGracePeriodChanged", logs: logs, sub: sub}, nil
}

// WatchBorrowerLoanReturnGracePeriodChanged is a free log subscription operation binding the contract event 0xf6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a.
//
// Solidity: event BorrowerLoanReturnGracePeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) WatchBorrowerLoanReturnGracePeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseBorrowerLoanReturnGracePeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "BorrowerLoanReturnGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseBorrowerLoanReturnGracePeriodChanged)
				if err := _Enterprise.contract.UnpackLog(event, "BorrowerLoanReturnGracePeriodChanged", log); err != nil {
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

// ParseBorrowerLoanReturnGracePeriodChanged is a log parse operation binding the contract event 0xf6ae16e9e9771497fc501d1caac6168e19c36e4639d7273813e3b57066035f2a.
//
// Solidity: event BorrowerLoanReturnGracePeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) ParseBorrowerLoanReturnGracePeriodChanged(log types.Log) (*EnterpriseBorrowerLoanReturnGracePeriodChanged, error) {
	event := new(EnterpriseBorrowerLoanReturnGracePeriodChanged)
	if err := _Enterprise.contract.UnpackLog(event, "BorrowerLoanReturnGracePeriodChanged", log); err != nil {
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

// EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator is returned from FilterEnterpriseLoanCollectGracePeriodChanged and is used to iterate over the raw logs and unpacked data for EnterpriseLoanCollectGracePeriodChanged events raised by the Enterprise contract.
type EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator struct {
	Event *EnterpriseEnterpriseLoanCollectGracePeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseEnterpriseLoanCollectGracePeriodChanged)
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
		it.Event = new(EnterpriseEnterpriseLoanCollectGracePeriodChanged)
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
func (it *EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseEnterpriseLoanCollectGracePeriodChanged represents a EnterpriseLoanCollectGracePeriodChanged event raised by the Enterprise contract.
type EnterpriseEnterpriseLoanCollectGracePeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseLoanCollectGracePeriodChanged is a free log retrieval operation binding the contract event 0x7c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b.
//
// Solidity: event EnterpriseLoanCollectGracePeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) FilterEnterpriseLoanCollectGracePeriodChanged(opts *bind.FilterOpts) (*EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "EnterpriseLoanCollectGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseEnterpriseLoanCollectGracePeriodChangedIterator{contract: _Enterprise.contract, event: "EnterpriseLoanCollectGracePeriodChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseLoanCollectGracePeriodChanged is a free log subscription operation binding the contract event 0x7c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b.
//
// Solidity: event EnterpriseLoanCollectGracePeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) WatchEnterpriseLoanCollectGracePeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseEnterpriseLoanCollectGracePeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "EnterpriseLoanCollectGracePeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseEnterpriseLoanCollectGracePeriodChanged)
				if err := _Enterprise.contract.UnpackLog(event, "EnterpriseLoanCollectGracePeriodChanged", log); err != nil {
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

// ParseEnterpriseLoanCollectGracePeriodChanged is a log parse operation binding the contract event 0x7c5a703a1928f0d18042c0454d90ddfe036b518c756a39cd82aca939b71d801b.
//
// Solidity: event EnterpriseLoanCollectGracePeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) ParseEnterpriseLoanCollectGracePeriodChanged(log types.Log) (*EnterpriseEnterpriseLoanCollectGracePeriodChanged, error) {
	event := new(EnterpriseEnterpriseLoanCollectGracePeriodChanged)
	if err := _Enterprise.contract.UnpackLog(event, "EnterpriseLoanCollectGracePeriodChanged", log); err != nil {
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

// EnterpriseEnterpriseVaultChangedIterator is returned from FilterEnterpriseVaultChanged and is used to iterate over the raw logs and unpacked data for EnterpriseVaultChanged events raised by the Enterprise contract.
type EnterpriseEnterpriseVaultChangedIterator struct {
	Event *EnterpriseEnterpriseVaultChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseEnterpriseVaultChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseEnterpriseVaultChanged)
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
		it.Event = new(EnterpriseEnterpriseVaultChanged)
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
func (it *EnterpriseEnterpriseVaultChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseEnterpriseVaultChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseEnterpriseVaultChanged represents a EnterpriseVaultChanged event raised by the Enterprise contract.
type EnterpriseEnterpriseVaultChanged struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseVaultChanged is a free log retrieval operation binding the contract event 0x0dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762dd.
//
// Solidity: event EnterpriseVaultChanged(address vault)
func (_Enterprise *EnterpriseFilterer) FilterEnterpriseVaultChanged(opts *bind.FilterOpts) (*EnterpriseEnterpriseVaultChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "EnterpriseVaultChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseEnterpriseVaultChangedIterator{contract: _Enterprise.contract, event: "EnterpriseVaultChanged", logs: logs, sub: sub}, nil
}

// WatchEnterpriseVaultChanged is a free log subscription operation binding the contract event 0x0dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762dd.
//
// Solidity: event EnterpriseVaultChanged(address vault)
func (_Enterprise *EnterpriseFilterer) WatchEnterpriseVaultChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseEnterpriseVaultChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "EnterpriseVaultChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseEnterpriseVaultChanged)
				if err := _Enterprise.contract.UnpackLog(event, "EnterpriseVaultChanged", log); err != nil {
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

// ParseEnterpriseVaultChanged is a log parse operation binding the contract event 0x0dc14fdd751203f0c4dac8cc93c303e991c5f8e60af35674d5c82061929762dd.
//
// Solidity: event EnterpriseVaultChanged(address vault)
func (_Enterprise *EnterpriseFilterer) ParseEnterpriseVaultChanged(log types.Log) (*EnterpriseEnterpriseVaultChanged, error) {
	event := new(EnterpriseEnterpriseVaultChanged)
	if err := _Enterprise.contract.UnpackLog(event, "EnterpriseVaultChanged", log); err != nil {
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

// EnterpriseInterestGapHalvingPeriodChangedIterator is returned from FilterInterestGapHalvingPeriodChanged and is used to iterate over the raw logs and unpacked data for InterestGapHalvingPeriodChanged events raised by the Enterprise contract.
type EnterpriseInterestGapHalvingPeriodChangedIterator struct {
	Event *EnterpriseInterestGapHalvingPeriodChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseInterestGapHalvingPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseInterestGapHalvingPeriodChanged)
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
		it.Event = new(EnterpriseInterestGapHalvingPeriodChanged)
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
func (it *EnterpriseInterestGapHalvingPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseInterestGapHalvingPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseInterestGapHalvingPeriodChanged represents a InterestGapHalvingPeriodChanged event raised by the Enterprise contract.
type EnterpriseInterestGapHalvingPeriodChanged struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterestGapHalvingPeriodChanged is a free log retrieval operation binding the contract event 0x971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f.
//
// Solidity: event InterestGapHalvingPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) FilterInterestGapHalvingPeriodChanged(opts *bind.FilterOpts) (*EnterpriseInterestGapHalvingPeriodChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "InterestGapHalvingPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseInterestGapHalvingPeriodChangedIterator{contract: _Enterprise.contract, event: "InterestGapHalvingPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchInterestGapHalvingPeriodChanged is a free log subscription operation binding the contract event 0x971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f.
//
// Solidity: event InterestGapHalvingPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) WatchInterestGapHalvingPeriodChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseInterestGapHalvingPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "InterestGapHalvingPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseInterestGapHalvingPeriodChanged)
				if err := _Enterprise.contract.UnpackLog(event, "InterestGapHalvingPeriodChanged", log); err != nil {
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

// ParseInterestGapHalvingPeriodChanged is a log parse operation binding the contract event 0x971f32e7fcf39ef3d42dd88b6e08b31dd4bd8f53e7299c95b669c82e23807e2f.
//
// Solidity: event InterestGapHalvingPeriodChanged(uint32 period)
func (_Enterprise *EnterpriseFilterer) ParseInterestGapHalvingPeriodChanged(log types.Log) (*EnterpriseInterestGapHalvingPeriodChanged, error) {
	event := new(EnterpriseInterestGapHalvingPeriodChanged)
	if err := _Enterprise.contract.UnpackLog(event, "InterestGapHalvingPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the Enterprise contract.
type EnterpriseLiquidityChangedIterator struct {
	Event *EnterpriseLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseLiquidityChanged)
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
		it.Event = new(EnterpriseLiquidityChanged)
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
func (it *EnterpriseLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseLiquidityChanged represents a LiquidityChanged event raised by the Enterprise contract.
type EnterpriseLiquidityChanged struct {
	InterestTokenId *big.Int
	ChangeType      uint8
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0x5fbe47a2294bb165359bd422c05af7405a9f2c0e0f397deef39ade4da37b2764.
//
// Solidity: event LiquidityChanged(uint256 indexed interestTokenId, uint8 indexed changeType, uint256 amount)
func (_Enterprise *EnterpriseFilterer) FilterLiquidityChanged(opts *bind.FilterOpts, interestTokenId []*big.Int, changeType []uint8) (*EnterpriseLiquidityChangedIterator, error) {

	var interestTokenIdRule []interface{}
	for _, interestTokenIdItem := range interestTokenId {
		interestTokenIdRule = append(interestTokenIdRule, interestTokenIdItem)
	}
	var changeTypeRule []interface{}
	for _, changeTypeItem := range changeType {
		changeTypeRule = append(changeTypeRule, changeTypeItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "LiquidityChanged", interestTokenIdRule, changeTypeRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseLiquidityChangedIterator{contract: _Enterprise.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0x5fbe47a2294bb165359bd422c05af7405a9f2c0e0f397deef39ade4da37b2764.
//
// Solidity: event LiquidityChanged(uint256 indexed interestTokenId, uint8 indexed changeType, uint256 amount)
func (_Enterprise *EnterpriseFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseLiquidityChanged, interestTokenId []*big.Int, changeType []uint8) (event.Subscription, error) {

	var interestTokenIdRule []interface{}
	for _, interestTokenIdItem := range interestTokenId {
		interestTokenIdRule = append(interestTokenIdRule, interestTokenIdItem)
	}
	var changeTypeRule []interface{}
	for _, changeTypeItem := range changeType {
		changeTypeRule = append(changeTypeRule, changeTypeItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "LiquidityChanged", interestTokenIdRule, changeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseLiquidityChanged)
				if err := _Enterprise.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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

// ParseLiquidityChanged is a log parse operation binding the contract event 0x5fbe47a2294bb165359bd422c05af7405a9f2c0e0f397deef39ade4da37b2764.
//
// Solidity: event LiquidityChanged(uint256 indexed interestTokenId, uint8 indexed changeType, uint256 amount)
func (_Enterprise *EnterpriseFilterer) ParseLiquidityChanged(log types.Log) (*EnterpriseLiquidityChanged, error) {
	event := new(EnterpriseLiquidityChanged)
	if err := _Enterprise.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseLoanReturnedIterator is returned from FilterLoanReturned and is used to iterate over the raw logs and unpacked data for LoanReturned events raised by the Enterprise contract.
type EnterpriseLoanReturnedIterator struct {
	Event *EnterpriseLoanReturned // Event containing the contract specifics and raw log

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
func (it *EnterpriseLoanReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseLoanReturned)
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
		it.Event = new(EnterpriseLoanReturned)
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
func (it *EnterpriseLoanReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseLoanReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseLoanReturned represents a LoanReturned event raised by the Enterprise contract.
type EnterpriseLoanReturned struct {
	BorrowTokenId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLoanReturned is a free log retrieval operation binding the contract event 0x80fadef8be1cf6508244c1bbb67077973d214a1462d3ade2908699c3750cd94b.
//
// Solidity: event LoanReturned(uint256 indexed borrowTokenId)
func (_Enterprise *EnterpriseFilterer) FilterLoanReturned(opts *bind.FilterOpts, borrowTokenId []*big.Int) (*EnterpriseLoanReturnedIterator, error) {

	var borrowTokenIdRule []interface{}
	for _, borrowTokenIdItem := range borrowTokenId {
		borrowTokenIdRule = append(borrowTokenIdRule, borrowTokenIdItem)
	}

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "LoanReturned", borrowTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseLoanReturnedIterator{contract: _Enterprise.contract, event: "LoanReturned", logs: logs, sub: sub}, nil
}

// WatchLoanReturned is a free log subscription operation binding the contract event 0x80fadef8be1cf6508244c1bbb67077973d214a1462d3ade2908699c3750cd94b.
//
// Solidity: event LoanReturned(uint256 indexed borrowTokenId)
func (_Enterprise *EnterpriseFilterer) WatchLoanReturned(opts *bind.WatchOpts, sink chan<- *EnterpriseLoanReturned, borrowTokenId []*big.Int) (event.Subscription, error) {

	var borrowTokenIdRule []interface{}
	for _, borrowTokenIdItem := range borrowTokenId {
		borrowTokenIdRule = append(borrowTokenIdRule, borrowTokenIdItem)
	}

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "LoanReturned", borrowTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseLoanReturned)
				if err := _Enterprise.contract.UnpackLog(event, "LoanReturned", log); err != nil {
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

// ParseLoanReturned is a log parse operation binding the contract event 0x80fadef8be1cf6508244c1bbb67077973d214a1462d3ade2908699c3750cd94b.
//
// Solidity: event LoanReturned(uint256 indexed borrowTokenId)
func (_Enterprise *EnterpriseFilterer) ParseLoanReturned(log types.Log) (*EnterpriseLoanReturned, error) {
	event := new(EnterpriseLoanReturned)
	if err := _Enterprise.contract.UnpackLog(event, "LoanReturned", log); err != nil {
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

// EnterpriseTotalSharesChangedIterator is returned from FilterTotalSharesChanged and is used to iterate over the raw logs and unpacked data for TotalSharesChanged events raised by the Enterprise contract.
type EnterpriseTotalSharesChangedIterator struct {
	Event *EnterpriseTotalSharesChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseTotalSharesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseTotalSharesChanged)
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
		it.Event = new(EnterpriseTotalSharesChanged)
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
func (it *EnterpriseTotalSharesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseTotalSharesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseTotalSharesChanged represents a TotalSharesChanged event raised by the Enterprise contract.
type EnterpriseTotalSharesChanged struct {
	TotalShares *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTotalSharesChanged is a free log retrieval operation binding the contract event 0xf75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264.
//
// Solidity: event TotalSharesChanged(uint256 totalShares)
func (_Enterprise *EnterpriseFilterer) FilterTotalSharesChanged(opts *bind.FilterOpts) (*EnterpriseTotalSharesChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "TotalSharesChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseTotalSharesChangedIterator{contract: _Enterprise.contract, event: "TotalSharesChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSharesChanged is a free log subscription operation binding the contract event 0xf75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264.
//
// Solidity: event TotalSharesChanged(uint256 totalShares)
func (_Enterprise *EnterpriseFilterer) WatchTotalSharesChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseTotalSharesChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "TotalSharesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseTotalSharesChanged)
				if err := _Enterprise.contract.UnpackLog(event, "TotalSharesChanged", log); err != nil {
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

// ParseTotalSharesChanged is a log parse operation binding the contract event 0xf75b433a747ade2e089a9382669e86292bab6b6218af81254ef18440992bb264.
//
// Solidity: event TotalSharesChanged(uint256 totalShares)
func (_Enterprise *EnterpriseFilterer) ParseTotalSharesChanged(log types.Log) (*EnterpriseTotalSharesChanged, error) {
	event := new(EnterpriseTotalSharesChanged)
	if err := _Enterprise.contract.UnpackLog(event, "TotalSharesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseUsedReserveChangedIterator is returned from FilterUsedReserveChanged and is used to iterate over the raw logs and unpacked data for UsedReserveChanged events raised by the Enterprise contract.
type EnterpriseUsedReserveChangedIterator struct {
	Event *EnterpriseUsedReserveChanged // Event containing the contract specifics and raw log

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
func (it *EnterpriseUsedReserveChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseUsedReserveChanged)
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
		it.Event = new(EnterpriseUsedReserveChanged)
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
func (it *EnterpriseUsedReserveChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseUsedReserveChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseUsedReserveChanged represents a UsedReserveChanged event raised by the Enterprise contract.
type EnterpriseUsedReserveChanged struct {
	FixedReserve *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUsedReserveChanged is a free log retrieval operation binding the contract event 0x80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585.
//
// Solidity: event UsedReserveChanged(uint256 fixedReserve)
func (_Enterprise *EnterpriseFilterer) FilterUsedReserveChanged(opts *bind.FilterOpts) (*EnterpriseUsedReserveChangedIterator, error) {

	logs, sub, err := _Enterprise.contract.FilterLogs(opts, "UsedReserveChanged")
	if err != nil {
		return nil, err
	}
	return &EnterpriseUsedReserveChangedIterator{contract: _Enterprise.contract, event: "UsedReserveChanged", logs: logs, sub: sub}, nil
}

// WatchUsedReserveChanged is a free log subscription operation binding the contract event 0x80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585.
//
// Solidity: event UsedReserveChanged(uint256 fixedReserve)
func (_Enterprise *EnterpriseFilterer) WatchUsedReserveChanged(opts *bind.WatchOpts, sink chan<- *EnterpriseUsedReserveChanged) (event.Subscription, error) {

	logs, sub, err := _Enterprise.contract.WatchLogs(opts, "UsedReserveChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseUsedReserveChanged)
				if err := _Enterprise.contract.UnpackLog(event, "UsedReserveChanged", log); err != nil {
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

// ParseUsedReserveChanged is a log parse operation binding the contract event 0x80a06b510f93a5f9143990fd9bd6b2020b17681a30b6420dac3d79b20ce73585.
//
// Solidity: event UsedReserveChanged(uint256 fixedReserve)
func (_Enterprise *EnterpriseFilterer) ParseUsedReserveChanged(log types.Log) (*EnterpriseUsedReserveChanged, error) {
	event := new(EnterpriseUsedReserveChanged)
	if err := _Enterprise.contract.UnpackLog(event, "UsedReserveChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
