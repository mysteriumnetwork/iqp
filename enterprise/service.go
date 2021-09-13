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
	GetGapHalvingPeriod(serviceAddress common.Address) (int64, error)
	GetServiceIndex(serviceAddress common.Address) (int64, error)
	GetBaseTokenAddress(serviceAddress common.Address) (common.Address, error)
	GetMinLoanDuration(serviceAddress common.Address) (int64, error)
	GetMaxLoanDuration(serviceAddress common.Address) (int64, error)
	GetServiceFeePercent(serviceAddress common.Address) (int64, error)
	GetAllowsPerpetual(serviceAddress common.Address) (bool, error)
	GetLiquidityTokenServiceAllowance(serviceAddress, accountAddress common.Address) (*big.Int, error)
	ApproveLiquidityTokensToService(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error)
	SetBaseRate(serviceAddress, baseToken common.Address, baseRate, minGcFee *big.Int) (*types.Transaction, error)
	SetLoanDurationLimits(serviceAddress common.Address, minLoanDuration, maxLoanDuration uint32) (*types.Transaction, error)
	SetServiceFeePercent(serviceAddress common.Address, feePercent uint16) (*types.Transaction, error)
	Wrap(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error)
	WrapTo(serviceAddress, accountAddress common.Address, amount *big.Int) (*types.Transaction, error)
	Unwrap(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error)
	GetPowerTokenAvailableBalance(serviceAddress, accountAddress common.Address) (*big.Int, error)
	GetPowerTokenBalance(serviceAddress, accountAddress common.Address) (*big.Int, error)
	GetEnergyAt(serviceAddress, accountAddress common.Address, timestamp uint32) (*big.Int, error)
	EstimateLoanDetailed(serviceAddress, paymentTokenAddress common.Address, amount *big.Int, duration uint32) (eip155.LoanEstimateDetailed, error)
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

func (s *Service) GetGapHalvingPeriod() (int64, error) {
	return s.blockchain.GetGapHalvingPeriod(s.address)
}

func (s *Service) GetServiceIndex() (int64, error) {
	return s.blockchain.GetServiceIndex(s.address)
}

func (s *Service) GetBaseTokenAddres() (common.Address, error) {
	return s.blockchain.GetBaseTokenAddress(s.address)
}

func (s *Service) GetMinLoanDuration() (int64, error) {
	return s.blockchain.GetMinLoanDuration(s.address)
}

func (s *Service) GetMaxLoanDuration() (int64, error) {
	return s.blockchain.GetMaxLoanDuration(s.address)
}

func (s *Service) GetServiceFeePercent() (int64, error) {
	return s.blockchain.GetServiceFeePercent(s.address)
}

func (s *Service) IsPerpetualTokenAllowed() (bool, error) {
	return s.blockchain.GetAllowsPerpetual(s.address)
}

func (s *Service) GetLiquidityAllowance(account common.Address) (*big.Int, error) {
	return s.blockchain.GetLiquidityTokenServiceAllowance(s.address, account)
}

func (s *Service) SetLiquidityAllowance(amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.ApproveLiquidityTokensToService(s.address, amount)
}

func (s *Service) SetBaseRate(baseRate, minGCFee *big.Int, baseToken common.Address) (*types.Transaction, error) {
	return s.blockchain.SetBaseRate(s.address, baseToken, baseRate, minGCFee)
}

func (s *Service) SetLoanDurationLimits(minLoanDuration, maxLoanDuration uint32) (*types.Transaction, error) {
	return s.blockchain.SetLoanDurationLimits(s.address, minLoanDuration, maxLoanDuration)
}

func (s *Service) SetServiceFeePercent(feePercent uint16) (*types.Transaction, error) {
	return s.blockchain.SetServiceFeePercent(s.address, feePercent)
}

func (s *Service) Wrap(amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.Wrap(s.address, amount)
}

func (s *Service) WrapTo(accountAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.WrapTo(s.address, accountAddress, amount)
}

func (s *Service) Unwrap(amount *big.Int) (*types.Transaction, error) {
	return s.blockchain.Unwrap(s.address, amount)
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
	return s.blockchain.EstimateLoanDetailed(s.address, paymentTokenAddress, amount, duration)
}
