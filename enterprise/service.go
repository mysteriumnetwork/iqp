package enterprise

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mysteriumnetwork/iqp/account"
	"github.com/mysteriumnetwork/iqp/blockchain/eip155"
)

type ServiceBlockchain interface {
	GetChainID() (int64, error)
	GetServiceInfo(serviceAddress common.Address) (eip155.ServiceInfo, error)
	GetAccountState(serviceAddress, accountAddress common.Address) (eip155.AccountState, error)
	GetBaseRate(serviceAddress common.Address) (*big.Int, error)
	GetMinGCFee(serviceAddress common.Address) (*big.Int, error)
	GetEnergyGapHalvingPeriod(serviceAddress common.Address) (uint32, error)
	GetServiceIndex(serviceAddress common.Address) (uint16, error)
	GetBaseTokenAddress(serviceAddress common.Address) (common.Address, error)
	GetMinRentalDuration(serviceAddress common.Address) (uint32, error)
	GetMaxRentalDuration(serviceAddress common.Address) (uint32, error)
	GetServiceFeePercent(serviceAddress common.Address) (uint16, error)
	GetAllowsSwapping(serviceAddress common.Address) (bool, error)
	GetProxyAdminTokenServiceAllowance(serviceAddress, accountAddress common.Address) (*big.Int, error)
	ApproveLiquidityTokensToService(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error)
	SetBaseRate(serviceAddress, baseToken common.Address, baseRate, minGcFee *big.Int) (*types.Transaction, error)
	SetRentalPeriodLimits(serviceAddress common.Address, minLoanDuration, maxLoanDuration uint32) (*types.Transaction, error)
	SetServiceFeePercent(serviceAddress common.Address, feePercent uint16) (*types.Transaction, error)
	SwapIn(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error)
	SwapOut(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error)
	GetPowerTokenAvailableBalance(serviceAddress, accountAddress common.Address) (*big.Int, error)
	GetPowerTokenBalance(serviceAddress, accountAddress common.Address) (*big.Int, error)
	GetEnergyAt(serviceAddress, accountAddress common.Address, timestamp uint32) (*big.Int, error)
	EstimateRentalFee(serviceAddress, paymentTokenAddress common.Address, amount *big.Int, duration uint32) (eip155.LoanEstimateDetailed, error)
}

type Service struct {
	address    common.Address
	blockchain ServiceBlockchain
}

func NewService(address common.Address, blockchain ServiceBlockchain) *Service {
	return &Service{
		address:    address,
		blockchain: blockchain,
	}
}

func (s *Service) GetID() (account.AccountID, error) {
	id, err := s.blockchain.GetChainID()
	if err != nil {
		return account.AccountID{}, err
	}

	return account.NewAccountIDWithRawChain(id, s.address)
}

func (s *Service) GetAddress() common.Address {
	return s.address
}

func (s *Service) GetInfo() (eip155.ServiceInfo, error) {
	return s.blockchain.GetServiceInfo(s.address)
}

func (s *Service) GetAccountState(accountAddress common.Address) (eip155.AccountState, error) {
	return s.blockchain.GetAccountState(s.address, accountAddress)
}

func (s *Service) GetBaseRate() (*big.Int, error) {
	return s.blockchain.GetBaseRate(s.address)
}

func (s *Service) GetMinGCFee() (*big.Int, error) {
	return s.blockchain.GetMinGCFee(s.address)
}

func (s *Service) GetEnergyGapHalvingPeriod() (uint32, error) {
	return s.blockchain.GetEnergyGapHalvingPeriod(s.address)
}

func (s *Service) GetServiceIndex() (uint16, error) {
	return s.blockchain.GetServiceIndex(s.address)
}

func (s *Service) GetBaseTokenAddres() (common.Address, error) {
	return s.blockchain.GetBaseTokenAddress(s.address)
}

func (s *Service) GetMinRentalDuration() (uint32, error) {
	return s.blockchain.GetMinRentalDuration(s.address)
}

func (s *Service) GetMaxRentalDuration() (uint32, error) {
	return s.blockchain.GetMaxRentalDuration(s.address)
}

func (s *Service) GetServiceFeePercent() (uint16, error) {
	return s.blockchain.GetServiceFeePercent(s.address)
}

func (s *Service) IsSwappingAllowed() (bool, error) {
	return s.blockchain.GetAllowsSwapping(s.address)
}

func (s *Service) GetLiquidityAllowance(account common.Address) (*big.Int, error) {
	return s.blockchain.GetProxyAdminTokenServiceAllowance(s.address, account)
}

func (s *Service) SetLiquidityAllowance(amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.ApproveLiquidityTokensToService(s.address, amount)
}

func (s *Service) SetBaseRate(baseRate, minGCFee *big.Int, baseToken common.Address) (*types.Transaction, error) {
	return s.blockchain.SetBaseRate(s.address, baseToken, baseRate, minGCFee)
}

func (s *Service) SetRentalPeriodLimits(minLoanDuration, maxLoanDuration uint32) (*types.Transaction, error) {
	return s.blockchain.SetRentalPeriodLimits(s.address, minLoanDuration, maxLoanDuration)
}

func (s *Service) SetServiceFeePercent(feePercent uint16) (*types.Transaction, error) {
	return s.blockchain.SetServiceFeePercent(s.address, feePercent)
}

func (s *Service) SwapIn(amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.SwapIn(s.address, amount)
}

func (s *Service) SwapOut(amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.SwapOut(s.address, amount)
}

func (s *Service) GetAvailableBalance(account common.Address) (*big.Int, error) {
	return s.blockchain.GetPowerTokenAvailableBalance(s.address, account)
}

func (s *Service) GetBalance(account common.Address) (*big.Int, error) {
	return s.blockchain.GetPowerTokenBalance(s.address, account)
}

func (s *Service) GetEnergyAt(timestamp uint32, account common.Address) (*big.Int, error) {
	return s.blockchain.GetEnergyAt(s.address, account, timestamp)
}

func (s *Service) EstimateLoanDetailed(paymentTokenAddress common.Address, amount *big.Int, duration uint32) (eip155.LoanEstimateDetailed, error) {
	return s.blockchain.EstimateRentalFee(s.address, paymentTokenAddress, amount, duration)
}
