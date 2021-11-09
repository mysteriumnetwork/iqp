package enterprise

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/iqp/account"
	"github.com/mysteriumnetwork/iqp/blockchain/eip155"
)

type EnterpriseBlockchain interface {
	ServiceBlockchain
	GetEnterpriseInfo(enterpriseAddress common.Address) (eip155.EnterpriseInfo, error)
	GetServices(enterpriseAddress common.Address) ([]common.Address, error)
	ClaimStakingReward(enterpriseAddress common.Address, stakeTokenID *big.Int) (*types.Transaction, error)
	GetStakingReward(enterpriseAddress common.Address, interestTokenID *big.Int) (*big.Int, error)
	ReturnRental(enterpriseAddress common.Address, borrowTokenId *big.Int) (*types.Transaction, error)
	GetRentalTokenIDs(enterpriseAddress, accountAddress common.Address) ([]*big.Int, error)
	GetRentalAgreement(enterpriseAddress common.Address, borrowTokenID *big.Int) (eip155.LoanInfo, error)
	GetPaymentTokenIDs(enterpriseAddress, accountAddress, liquidityTokenAddr common.Address) ([]*big.Int, error)
	GetStake(enterpriseAddress common.Address, interestTokenID *big.Int) (eip155.Stake, error)
	Stake(enterpriseAddress common.Address, amount *big.Int) (*types.Transaction, error)
	Unstake(enterpriseAddress common.Address, interestTokenID *big.Int) (*types.Transaction, error)
	IncreaseStake(enterpriseAddress common.Address, interestTokenID, amount *big.Int) (*types.Transaction, error)
	DecreaseStake(enterpriseAddress common.Address, interestTokenID, amount *big.Int) (*types.Transaction, error)
	ApproveLiquidityTokensToEnterprise(enterpriseAddress common.Address, amount *big.Int) (*types.Transaction, error)
	GetLiquidityTokenEnterpriseAllowance(enterpriseAddress, accountAddress common.Address) (*big.Int, error)
	IsRegisteredService(enterpriseAddress, serviceAddress common.Address) (bool, error)
	GetProxyAdminAddress(enterpriseAddress common.Address) (common.Address, error)
	GetEnterpriseCollectorAddress(enterpriseAddress common.Address) (common.Address, error)
	GetEnterpriseWalletAddress(enterpriseAddress common.Address) (common.Address, error)
	GetRenterOnlyReturnPeriod(enterpriseAddress common.Address) (uint32, error)
	GetEnterpriseTokenAddress(enterpriseAddress common.Address) (common.Address, error)
	GetInterestReserveHalvingPeriod(enterpriseAddress common.Address) (uint32, error)
	GetConverterAddress(enterpriseAddress common.Address) (common.Address, error)
	GetBaseUri(enterpriseAddress common.Address) (string, error)
	GetReserve(enterpriseAddress common.Address) (*big.Int, error)
	GetUsedReserve(enterpriseAddress common.Address) (*big.Int, error)
	GetAvailableReserve(enterpriseAddress common.Address) (*big.Int, error)
	GetBondingCurve(enterpriseAddress common.Address) (eip155.BondingCurve, error)
	GetGCFeePercent(enterpriseAddress common.Address) (uint16, error)
	SetEnterpriseWalletAddress(enterpriseAddress, vaultAddress common.Address) (*types.Transaction, error)
	SetEnterpriseCollectorAddress(enterpriseAddress, collector common.Address) (*types.Transaction, error)
	SetConverterAddress(enterpriseAddress, converter common.Address) (*types.Transaction, error)
	SetBondingCurve(enterpriseAddress common.Address, pole, slope *big.Int) (*types.Transaction, error)
	SetRenterOnlyReturnPeriod(enterpriseAddress common.Address, period uint32) (*types.Transaction, error)
	SetEnterpriseLoanCollectionPeriod(enterpriseAddress common.Address, period uint32) (*types.Transaction, error)
	SetBaseUri(enterpriseAddress common.Address, address string) (*types.Transaction, error)
	SetInterestReserveHalvingPeriod(enterpriseAddress common.Address, interestGapHalvingPeriod uint32) (*types.Transaction, error)
	SetGCFeePercent(enterpriseAddress common.Address, percent uint16) (*types.Transaction, error)
}

type Enterprise struct {
	address    common.Address
	blockchain EnterpriseBlockchain
}

func NewEnterprise(address common.Address, blockchain EnterpriseBlockchain) *Enterprise {
	return &Enterprise{
		address:    address,
		blockchain: blockchain,
	}
}

func (e *Enterprise) GetAddress() common.Address {
	return e.address
}

func (e *Enterprise) GetID() (account.AccountID, error) {
	id, err := e.blockchain.GetChainID()
	if err != nil {
		return account.AccountID{}, err
	}

	return account.NewAccountIDWithRawChain(id, e.address)
}

func (e *Enterprise) GetInfo() (eip155.EnterpriseInfo, error) {
	return e.blockchain.GetEnterpriseInfo(e.address)
}

func (e *Enterprise) GetServices() ([]*Service, error) {
	svcs, err := e.blockchain.GetServices(e.address)
	if err != nil {
		return nil, err
	}

	res := make([]*Service, len(svcs))
	for i := range svcs {
		res[i] = NewService(svcs[i], e.blockchain)
	}

	return res, nil
}

func (e *Enterprise) EstimateRentalFee(serviceAddress, paymentTokenAddress common.Address, amount *big.Int, duration uint32) (eip155.LoanEstimateDetailed, error) {
	return e.blockchain.EstimateRentalFee(serviceAddress, paymentTokenAddress, amount, duration)
}

func (e *Enterprise) ClaimStakingReward(stakeTokenID *big.Int) (*types.Transaction, error) {
	return e.blockchain.ClaimStakingReward(e.address, stakeTokenID)
}

func (e *Enterprise) GetStakingReward(interestTokenID *big.Int) (*big.Int, error) {
	return e.blockchain.GetStakingReward(e.address, interestTokenID)
}

func (e *Enterprise) ReturnRental(interestTokenID *big.Int) (*types.Transaction, error) {
	return e.blockchain.ReturnRental(e.address, interestTokenID)
}

func (e *Enterprise) ListLoanInfo(accountAddress common.Address) ([]eip155.LoanInfo, error) {
	borrowTokenIDs, err := e.blockchain.GetRentalTokenIDs(e.address, accountAddress)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(borrowTokenIDs))

	type fetchResult struct {
		loanInfo eip155.LoanInfo
		err      error
	}

	res := make([]fetchResult, len(borrowTokenIDs))
	for i := range borrowTokenIDs {
		go func(idx int) {
			defer wg.Done()
			info, err := e.blockchain.GetRentalAgreement(e.address, borrowTokenIDs[idx])
			res[idx] = fetchResult{
				loanInfo: info,
				err:      err,
			}
		}(i)
	}

	wg.Wait()

	result := make([]eip155.LoanInfo, len(res))
	for i := range res {
		if res[i].err != nil {
			return nil, err
		}
		result[i] = res[i].loanInfo
	}

	return result, nil
}

func (e *Enterprise) GetRentalAgreement(borrowTokenID *big.Int) (eip155.LoanInfo, error) {
	return e.blockchain.GetRentalAgreement(e.address, borrowTokenID)
}

func (e *Enterprise) ListLiquidityInfo(accountAddress, liquidityTokenAddress common.Address) ([]eip155.Stake, error) {
	liquidityTokenIDs, err := e.blockchain.GetPaymentTokenIDs(e.address, accountAddress, liquidityTokenAddress)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(len(liquidityTokenIDs))

	type fetchResult struct {
		liquidityInfo eip155.Stake
		err           error
	}

	res := make([]fetchResult, len(liquidityTokenIDs))
	for i := range liquidityTokenIDs {
		go func(idx int) {
			defer wg.Done()
			info, err := e.blockchain.GetStake(e.address, liquidityTokenIDs[idx])
			res[idx] = fetchResult{
				liquidityInfo: info,
				err:           err,
			}
		}(i)
	}

	wg.Wait()

	result := make([]eip155.Stake, len(res))
	for i := range res {
		if res[i].err != nil {
			return nil, err
		}
		result[i] = res[i].liquidityInfo
	}

	return result, nil
}

func (e *Enterprise) GetStake(interestTokenID *big.Int) (eip155.Stake, error) {
	return e.blockchain.GetStake(e.address, interestTokenID)
}

func (e *Enterprise) Stake(amount *big.Int) (*types.Transaction, error) {
	return e.blockchain.Stake(e.address, amount)
}

func (e *Enterprise) Unstake(amount *big.Int) (*types.Transaction, error) {
	return e.blockchain.Unstake(e.address, amount)
}

func (e *Enterprise) IncreaseStake(interestTokenID, amount *big.Int) (*types.Transaction, error) {
	return e.blockchain.IncreaseStake(e.address, interestTokenID, amount)
}

func (e *Enterprise) DecreaseStake(interestTokenID, amount *big.Int) (*types.Transaction, error) {
	return e.blockchain.DecreaseStake(e.address, interestTokenID, amount)
}

func (e *Enterprise) SetLiquidityAllowance(amount *big.Int) (*types.Transaction, error) {
	return e.blockchain.ApproveLiquidityTokensToEnterprise(e.address, amount)
}

func (e *Enterprise) GetLiquidityAllowance(accountAddress common.Address) (*big.Int, error) {
	return e.blockchain.GetLiquidityTokenEnterpriseAllowance(e.address, accountAddress)
}

func (e *Enterprise) IsRegisteredService(serviceAddress common.Address) (bool, error) {
	return e.blockchain.IsRegisteredService(e.address, serviceAddress)
}

func (e *Enterprise) GetProxyAdminAddress() (common.Address, error) {
	return e.blockchain.GetProxyAdminAddress(e.address)
}

func (e *Enterprise) GetCollectorAddress() (common.Address, error) {
	return e.blockchain.GetEnterpriseCollectorAddress(e.address)
}

func (e *Enterprise) GetWalletAddress() (common.Address, error) {
	return e.blockchain.GetEnterpriseWalletAddress(e.address)
}

func (e *Enterprise) GetRenterOnlyReturnPeriod() (uint32, error) {
	return e.blockchain.GetRenterOnlyReturnPeriod(e.address)
}

func (e *Enterprise) GetEnterpriseTokenAddress() (common.Address, error) {
	return e.blockchain.GetEnterpriseTokenAddress(e.address)
}

func (e *Enterprise) GetStreamingReserveHalvingPeriod() (uint32, error) {
	return e.blockchain.GetInterestReserveHalvingPeriod(e.address)
}
func (e *Enterprise) GetConverterAddress() (common.Address, error) {
	return e.blockchain.GetConverterAddress(e.address)
}

func (e *Enterprise) GetBaseUri() (string, error) {
	return e.blockchain.GetBaseUri(e.address)
}

func (e *Enterprise) GetReserve() (*big.Int, error) {
	return e.blockchain.GetReserve(e.address)
}

func (e *Enterprise) GetUsedReserve() (*big.Int, error) {
	return e.blockchain.GetUsedReserve(e.address)
}

func (e *Enterprise) GetAvailableReserve() (*big.Int, error) {
	return e.blockchain.GetAvailableReserve(e.address)
}

func (e *Enterprise) GetBondingCurve() (eip155.BondingCurve, error) {
	return e.blockchain.GetBondingCurve(e.address)
}

func (e *Enterprise) GetGCFeePercent() (uint16, error) {
	return e.blockchain.GetGCFeePercent(e.address)
}

func (e *Enterprise) SetCollectorAddress(collector common.Address) (*types.Transaction, error) {
	return e.blockchain.SetEnterpriseCollectorAddress(e.address, collector)
}

func (e *Enterprise) SetVaultAddress(vaultAddress common.Address) (*types.Transaction, error) {
	return e.blockchain.SetEnterpriseWalletAddress(e.address, vaultAddress)
}

func (e *Enterprise) SetConverterAddress(converter common.Address) (*types.Transaction, error) {
	return e.blockchain.SetConverterAddress(e.address, converter)
}

func (e *Enterprise) SetBondingCurve(pole, slope *big.Int) (*types.Transaction, error) {
	return e.blockchain.SetBondingCurve(e.address, pole, slope)
}

func (e *Enterprise) SetRenterOnlyReturnPeriod(period uint32) (*types.Transaction, error) {
	return e.blockchain.SetRenterOnlyReturnPeriod(e.address, period)
}

func (e *Enterprise) SetLoanCollectGracePeriod(period uint32) (*types.Transaction, error) {
	return e.blockchain.SetEnterpriseLoanCollectionPeriod(e.address, period)
}

func (e *Enterprise) SetBaseUri(address string) (*types.Transaction, error) {
	return e.blockchain.SetBaseUri(e.address, address)
}

func (e *Enterprise) SetInterestReserveHalvingPeriod(interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	return e.blockchain.SetInterestReserveHalvingPeriod(e.address, interestGapHalvingPeriod)
}

func (e *Enterprise) SetGCFeePercent(percent uint16) (*types.Transaction, error) {
	return e.blockchain.SetGCFeePercent(e.address, percent)
}
