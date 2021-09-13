package eip155

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mysteriumnetwork/iqp/bindings"
)

type Client struct {
	c         *ethclient.Client
	bcTimeout time.Duration
	signerFn  bind.SignerFn
}

func NewClient(c *ethclient.Client, bcTimeout time.Duration, signerFunc bind.SignerFn) *Client {
	return &Client{
		c:         c,
		bcTimeout: bcTimeout,
		signerFn:  signerFunc,
	}
}

func (c *Client) GetChainID() (int64, error) {
	ctx, cancel := c.newContext()
	defer cancel()

	chid, err := c.c.ChainID(ctx)
	if err != nil {
		return 0, err
	}

	return chid.Int64(), nil
}

type ERC20Metadata struct {
	Address  common.Address
	Name     string
	Symbol   string
	Decimals uint8
}

func (c *Client) GetERC20Metadata(tokenAddress common.Address) (ERC20Metadata, error) {
	caller, err := bindings.NewERC20Caller(tokenAddress, c.c)
	if err != nil {
		return ERC20Metadata{}, err
	}

	opts, cancel := c.callOpts()
	defer cancel()

	res := ERC20Metadata{
		Address: tokenAddress,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	errChan := make(chan error, 3)
	go func() {
		defer wg.Done()
		decimals, err := caller.Decimals(opts)
		if err != nil {
			errChan <- err
			return
		}

		res.Decimals = decimals
	}()

	go func() {
		defer wg.Done()
		name, err := caller.Name(opts)
		if err != nil {
			errChan <- err
			return
		}

		res.Name = name
	}()

	go func() {
		defer wg.Done()
		symbol, err := caller.Symbol(opts)
		if err != nil {
			errChan <- err
			return
		}

		res.Symbol = symbol
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return ERC20Metadata{}, err
		}
	}

	return res, nil
}

type ERC721Metadata struct {
	Address  common.Address
	Name     string
	Symbol   string
	TokenURI string
}

func (c *Client) GetERC721Metadata(tokenAddress common.Address, tokenID *big.Int) (ERC721Metadata, error) {
	caller, err := bindings.NewERC721Caller(tokenAddress, c.c)
	if err != nil {
		return ERC721Metadata{}, err
	}

	opts, cancel := c.callOpts()
	defer cancel()

	res := ERC721Metadata{
		Address: tokenAddress,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	errChan := make(chan error, 3)
	go func() {
		defer wg.Done()
		uri, err := caller.TokenURI(opts, tokenID)
		if err != nil {
			errChan <- err
			return
		}

		res.TokenURI = uri
	}()

	go func() {
		defer wg.Done()
		name, err := caller.Name(opts)
		if err != nil {
			errChan <- err
			return
		}

		res.Name = name
	}()

	go func() {
		defer wg.Done()
		symbol, err := caller.Symbol(opts)
		if err != nil {
			errChan <- err
			return
		}

		res.Symbol = symbol
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return ERC721Metadata{}, err
		}
	}
	return res, nil
}

func (c *Client) GetTokenBalance(tokenAddress common.Address, accountAddress common.Address) (*big.Int, error) {
	caller, err := bindings.NewERC20Caller(tokenAddress, c.c)
	if err != nil {
		return nil, err
	}

	opts, cancel := c.callOpts()
	defer cancel()

	return caller.BalanceOf(opts, accountAddress)
}

func (c *Client) newContext() (context.Context, func()) {
	return context.WithTimeout(context.Background(), c.bcTimeout)
}

func (c *Client) callOpts() (*bind.CallOpts, func()) {
	ctx, cancel := c.newContext()
	return &bind.CallOpts{
		Context: ctx,
	}, cancel
}

func (c *Client) transactOpts() (*bind.TransactOpts, func()) {
	ctx, cancel := c.newContext()
	return &bind.TransactOpts{
		Context: ctx,
		Signer:  c.signerFn,
	}, cancel
}

type EnterpriseParams struct {
	Name                  string
	LiquidityTokenAddress common.Address
	BaseUri               string
	GCFeePercent          uint16
	Converter             common.Address
}

func (c *Client) DeployEnterprise(enterpriseFactoryAddress common.Address, params EnterpriseParams) (*types.Transaction, error) {
	transactor, err := bindings.NewEnterpriseFactoryTransactor(enterpriseFactoryAddress, c.c)
	if err != nil {
		return nil, err
	}

	opts, cancel := c.transactOpts()
	defer cancel()

	return transactor.Deploy(opts, params.Name, params.LiquidityTokenAddress, params.BaseUri, params.GCFeePercent, params.Converter)
}

func (c *Client) GetServices(enterpriseAddress common.Address) ([]common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.GetPowerTokens(opts)
}

type EnterpriseInfo struct {
	Name                             string
	BaseUri                          string
	TotalShares                      *big.Int
	InterestGapHalvingPeriod         uint32
	BorrowerLoanReturnGracePeriod    uint32
	EnterpriseLoanCollectGracePeriod uint32
	GCFeePercent                     uint16
	FixedReserve                     *big.Int
	UsedReserve                      *big.Int
	StreamingReserve                 *big.Int
	StreamingReserveTarget           *big.Int
	StreamingReserveUpdated          uint32
}

func (c *Client) GetEnterpriseInfo(enterpriseAddress common.Address) (EnterpriseInfo, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return EnterpriseInfo{}, err
	}
	defer cancel()

	info, err := caller.GetInfo(opts)
	if err != nil {
		return EnterpriseInfo{}, err
	}

	return EnterpriseInfo{
		Name:                             info.Name,
		BaseUri:                          info.BaseUri,
		TotalShares:                      info.TotalShares,
		InterestGapHalvingPeriod:         info.InterestGapHalvingPeriod,
		BorrowerLoanReturnGracePeriod:    info.BorrowerLoanReturnGracePeriod,
		EnterpriseLoanCollectGracePeriod: info.EnterpriseLoanCollectGracePeriod,
		GCFeePercent:                     info.GcFeePercent,
		FixedReserve:                     info.FixedReserve,
		UsedReserve:                      info.UsedReserve,
		StreamingReserve:                 info.StreamingReserve,
		StreamingReserveTarget:           info.StreamingReserveTarget,
		StreamingReserveUpdated:          info.StreamingReserveUpdated,
	}, nil
}

func (c *Client) newEnterpriseCaller(enterpriseAddress common.Address) (*bindings.EnterpriseCaller, *bind.CallOpts, func(), error) {
	caller, err := bindings.NewEnterpriseCaller(enterpriseAddress, c.c)
	if err != nil {
		return nil, nil, func() {}, err
	}

	opts, cancel := c.callOpts()
	return caller, opts, cancel, nil
}

func (c *Client) newEnterpriseTransactor(enterpriseAddress common.Address) (*bindings.EnterpriseTransactor, *bind.TransactOpts, func(), error) {
	caller, err := bindings.NewEnterpriseTransactor(enterpriseAddress, c.c)
	if err != nil {
		return nil, nil, func() {}, err
	}

	opts, cancel := c.transactOpts()
	return caller, opts, cancel, nil
}

type ServiceParams struct {
	Name                         string
	Symbol                       string
	GapHalvingPeriod             uint32
	BaseRate                     *big.Int
	BaseToken                    common.Address
	ServiceFeePercent            uint16
	MinLoanDuration              uint32
	MaxLoanDuration              uint32
	MinGCFee                     *big.Int
	AllowsPerpetualTokensForever bool
}

func (c *Client) RegisterService(enterpriseAddress common.Address, params ServiceParams) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.RegisterService(
		opts,
		params.Name,
		params.Symbol,
		params.GapHalvingPeriod,
		params.BaseRate,
		params.BaseToken,
		params.ServiceFeePercent,
		params.MinLoanDuration,
		params.MaxLoanDuration,
		params.MinGCFee,
		params.AllowsPerpetualTokensForever,
	)
}

func (c *Client) AddLiquidity(enterpriseAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.AddLiquidity(opts, amount)
}

func (c *Client) RemoveLiquidity(enterpriseAddress common.Address, interestTokenID *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.RemoveLiquidity(opts, interestTokenID)
}

func (c *Client) IncreaseLiquidity(enterpriseAddress common.Address, interestTokenID, amount *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.IncreaseLiquidity(opts, interestTokenID, amount)
}

func (c *Client) DecreaseLiquidity(enterpriseAddress common.Address, interestTokenID, amount *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.DecreaseLiquidity(opts, interestTokenID, amount)
}

func (c *Client) WithdrawInterest(enterpriseAddress common.Address, interestTokenID *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.WithdrawInterest(opts, interestTokenID)
}

func (c *Client) Borrow(enterpriseAddress, serviceAddress, paymentTokenAddress common.Address, amount, maxPayment *big.Int, duration uint32) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.Borrow(opts, serviceAddress, paymentTokenAddress, amount, duration, maxPayment)
}

func (c *Client) ReturnLoan(enterpriseAddress common.Address, borrowTokenId *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.ReturnLoan(opts, borrowTokenId)
}

func (c *Client) ApproveLiquidityTokensToEnterprise(enterpriseAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	caller, err := bindings.NewEnterprise(enterpriseAddress, c.c)
	if err != nil {
		return nil, err
	}

	copts, ccancel := c.callOpts()
	defer ccancel()

	token, err := caller.GetLiquidityToken(copts)
	if err != nil {
		return nil, err
	}

	erc20Transactor, err := bindings.NewERC20Transactor(token, c.c)
	if err != nil {
		return nil, err
	}

	topts, tcancel := c.transactOpts()
	defer tcancel()

	return erc20Transactor.Approve(topts, enterpriseAddress, amount)
}

func (c *Client) SetBaseUri(enterpriseAddress common.Address, address string) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetBaseUri(opts, address)
}

func (c *Client) SetBondingCurve(enterpriseAddress common.Address, pole, slope *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetBondingCurve(opts, pole, slope)
}

func (c *Client) SetBorrowerLoanReturnGracePeriod(enterpriseAddress common.Address, period uint32) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetBorrowerLoanReturnGracePeriod(opts, period)
}

func (c *Client) SetConverterAddress(enterpriseAddress, converter common.Address) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetConverter(opts, converter)
}

func (c *Client) SetEnterpriseCollectorAddress(enterpriseAddress, collector common.Address) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetEnterpriseCollector(opts, collector)
}

func (c *Client) SetEnterpriseLoanCollectGracePeriod(enterpriseAddress common.Address, period uint32) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetEnterpriseLoanCollectGracePeriod(opts, period)
}

func (c *Client) SetEnterpriseVaultAddress(enterpriseAddress, vaultAddress common.Address) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetEnterpriseVault(opts, vaultAddress)
}

func (c *Client) SetGCFeePercent(enterpriseAddress common.Address, percent uint16) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetGcFeePercent(opts, percent)
}

func (c *Client) SetInterestGapHalvingPeriod(enterpriseAddress common.Address, interestGapHalvingPeriod uint32) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newEnterpriseTransactor(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetInterestGapHalvingPeriod(opts, interestGapHalvingPeriod)
}

func (c *Client) GetReserve(enterpriseAddress common.Address) (*big.Int, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.GetReserve(opts)

}

func (c *Client) GetUsedReserve(enterpriseAddress common.Address) (*big.Int, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.GetUsedReserve(opts)
}

func (c *Client) GetAvailableReserve(enterpriseAddress common.Address) (*big.Int, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.GetAvailableReserve(opts)
}

func (c *Client) EstimateLoan(enterpriseAddress, serviceAddress, paymentTokenAddress common.Address, amount *big.Int, duration uint32) (*big.Int, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.EstimateLoan(opts, serviceAddress, paymentTokenAddress, amount, duration)
}

func (c *Client) GetAccruedInterest(enterpriseAddress common.Address, interestTokenID *big.Int) (*big.Int, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.GetAccruedInterest(opts, interestTokenID)
}

func (c *Client) GetLiquidityTokenEnterpriseAllowance(enterpriseAddress, accountAddress common.Address) (*big.Int, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	tkn, err := caller.GetLiquidityToken(opts)
	if err != nil {
		return nil, err
	}

	tknCaller, err := bindings.NewERC20Caller(tkn, c.c)
	if err != nil {
		return nil, err
	}

	tknCallOpts, tknCancel := c.callOpts()
	defer tknCancel()

	return tknCaller.Allowance(tknCallOpts, accountAddress, enterpriseAddress)
}

func (c *Client) GetLiquidityTokenAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetLiquidityToken(opts)
}

func (c *Client) GetLiquidityTokenMetadata(enterpriseAddress common.Address) (ERC20Metadata, error) {
	tkn, err := c.GetLiquidityTokenAddress(enterpriseAddress)
	if err != nil {
		return ERC20Metadata{}, err
	}
	return c.GetERC20Metadata(tkn)
}

func (c *Client) GetBorrowTokenAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetBorrowToken(opts)
}

func (c *Client) GetBorrowTokenMetadata(enterpriseAddress common.Address) (ERC20Metadata, error) {
	tkn, err := c.GetBorrowTokenAddress(enterpriseAddress)
	if err != nil {
		return ERC20Metadata{}, err
	}
	return c.GetERC20Metadata(tkn)
}

func (c *Client) GetInterestTokenAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetInterestToken(opts)
}

func (c *Client) GetInterestTokenMetadata(enterpriseAddress common.Address) (ERC20Metadata, error) {
	tkn, err := c.GetInterestTokenAddress(enterpriseAddress)
	if err != nil {
		return ERC20Metadata{}, err
	}
	return c.GetERC20Metadata(tkn)
}

func (c *Client) GetInterestTokenIDs(enterpriseAddress, accountAddress common.Address) ([]*big.Int, error) {
	tokenAddress, err := c.GetInterestTokenAddress(enterpriseAddress)
	if err != nil {
		return nil, err
	}

	caller, err := bindings.NewInterestTokenCaller(tokenAddress, c.c)
	if err != nil {
		return nil, err
	}

	opts, cancel := c.callOpts()
	defer cancel()

	tokenCount, err := caller.BalanceOf(opts, accountAddress)
	if err != nil {
		return nil, err
	}

	res := make([]*big.Int, 0)
	var i int64
	for i = 0; i < tokenCount.Int64(); i++ {
		r, err := caller.TokenOfOwnerByIndex(opts, accountAddress, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}

	return res, nil
}

type LiquidityInfo struct {
	Amount *big.Int
	Shares *big.Int
	Block  *big.Int
}

func (c *Client) GetLiquidityInfo(enterpriseAddress common.Address, interestTokenID *big.Int) (LiquidityInfo, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return LiquidityInfo{}, err
	}
	defer cancel()

	info, err := caller.GetLiquidityInfo(opts, interestTokenID)
	if err != nil {
		return LiquidityInfo{}, err
	}
	return LiquidityInfo{
		Amount: info.Amount,
		Shares: info.Shares,
		Block:  info.Block,
	}, nil
}

func (c *Client) GetBorrowTokenIds(enterpriseAddress, accountAddress common.Address) ([]*big.Int, error) {
	borrowToken, err := c.GetBorrowTokenAddress(enterpriseAddress)
	if err != nil {
		return nil, err
	}

	caller, err := bindings.NewBorrowTokenCaller(borrowToken, c.c)
	if err != nil {
		return nil, err
	}

	opts, cancel := c.callOpts()
	defer cancel()
	tokenCount, err := caller.BalanceOf(opts, accountAddress)
	if err != nil {
		return nil, err
	}

	tokenIDs := make([]*big.Int, 0)

	var i int64
	for i = 0; i < tokenCount.Int64(); i++ {
		tknID, err := caller.TokenOfOwnerByIndex(opts, accountAddress, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		tokenIDs = append(tokenIDs, tknID)
	}
	return tokenIDs, nil
}

type LoanInfo struct {
	Amount                     *big.Int
	PowerTokenIndex            uint16
	BorrowingTime              uint32
	MaturityTime               uint32
	BorrowerReturnGraceTime    uint32
	EnterpriseCollectGraceTime uint32
	GcFee                      *big.Int
	GcFeeTokenIndex            uint16
	TokenID                    *big.Int
}

func (c *Client) GetLoanInfo(enterpriseAddress common.Address, borrowTokenID *big.Int) (LoanInfo, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return LoanInfo{}, err
	}
	defer cancel()

	info, err := caller.GetLoanInfo(opts, borrowTokenID)
	if err != nil {
		return LoanInfo{}, err
	}

	return LoanInfo{
		TokenID:                    borrowTokenID,
		Amount:                     info.Amount,
		PowerTokenIndex:            info.PowerTokenIndex,
		BorrowingTime:              info.BorrowingTime,
		MaturityTime:               info.MaturityTime,
		BorrowerReturnGraceTime:    info.BorrowerReturnGraceTime,
		EnterpriseCollectGraceTime: info.EnterpriseCollectGraceTime,
		GcFee:                      info.GcFee,
		GcFeeTokenIndex:            info.GcFeeTokenIndex,
	}, nil
}

func (c *Client) IsRegisteredService(enterpriseAddress, serviceAddress common.Address) (bool, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return false, err
	}
	defer cancel()

	return caller.IsRegisteredPowerToken(opts, serviceAddress)
}

func (c *Client) GetBaseUri(enterpriseAddress common.Address) (string, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return "", err
	}
	defer cancel()

	return caller.GetBaseUri(opts)
}

type BondingCurve struct {
	Pole, Slope *big.Int
}

func (c *Client) GetBondingCurve(enterpriseAddress common.Address) (BondingCurve, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return BondingCurve{}, err
	}
	defer cancel()

	info, err := caller.GetBondingCurve(opts)
	if err != nil {
		return BondingCurve{}, err
	}

	return BondingCurve{
		Pole:  info.Pole,
		Slope: info.Slope,
	}, nil
}

func (c *Client) GetBorrowerLoanReturnGracePeriod(enterpriseAddress common.Address) (uint32, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return 0, err
	}
	defer cancel()

	return caller.GetBorrowerLoanReturnGracePeriod(opts)
}

func (c *Client) GetConverterAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetConverter(opts)
}

func (c *Client) GetEnterpriseCollectorAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetEnterpriseCollector(opts)
}

func (c *Client) GetEnterpriseLoanCollectGracePeriod(enterpriseAddress common.Address) (uint32, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return 0, err
	}
	defer cancel()

	return caller.GetEnterpriseLoanCollectGracePeriod(opts)
}

func (c *Client) GetEnterpriseVaultAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetEnterpriseVault(opts)
}

func (c *Client) GetGCFeePercent(enterpriseAddress common.Address) (uint16, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return 0, err
	}
	defer cancel()

	return caller.GetGCFeePercent(opts)
}

func (c *Client) GetInterestGapHalvingPeriod(enterpriseAddress common.Address) (uint32, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return 0, err
	}
	defer cancel()

	return caller.GetInterestGapHalvingPeriod(opts)
}

func (c *Client) GetProxyAdminAddress(enterpriseAddress common.Address) (common.Address, error) {
	caller, opts, cancel, err := c.newEnterpriseCaller(enterpriseAddress)
	if err != nil {
		return common.Address{}, err
	}
	defer cancel()

	return caller.GetProxyAdmin(opts)
}

func (c *Client) newPowerTokenTransactor(powerTokenAddress common.Address) (*bindings.PowerTokenTransactor, *bind.TransactOpts, func(), error) {
	caller, err := bindings.NewPowerTokenTransactor(powerTokenAddress, c.c)
	if err != nil {
		return nil, nil, func() {}, err
	}

	opts, cancel := c.transactOpts()
	return caller, opts, cancel, nil
}

func (c *Client) newPowerTokenCaller(powerTokenAddress common.Address) (*bindings.PowerTokenCaller, *bind.CallOpts, func(), error) {
	caller, err := bindings.NewPowerTokenCaller(powerTokenAddress, c.c)
	if err != nil {
		return nil, nil, func() {}, err
	}

	opts, cancel := c.callOpts()
	return caller, opts, cancel, nil
}

func (c *Client) ApproveLiquidityTokensToService(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	transactor, err := bindings.NewPowerToken(serviceAddress, c.c)
	if err != nil {
		return nil, err
	}

	callOpts, cancel := c.callOpts()
	defer cancel()

	enterprise, err := transactor.GetEnterprise(callOpts)
	if err != nil {
		return nil, err
	}

	liquidityTokenAddr, err := c.GetLiquidityTokenAddress(enterprise)
	if err != nil {
		return nil, err
	}

	erc20Transactor, err := bindings.NewERC20Transactor(liquidityTokenAddr, c.c)
	if err != nil {
		return nil, err
	}

	topts, tcancel := c.transactOpts()
	defer tcancel()
	return erc20Transactor.Approve(topts, serviceAddress, amount)
}

func (c *Client) Wrap(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newPowerTokenTransactor(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.Wrap(opts, amount)
}

func (c *Client) WrapTo(serviceAddress, accountAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newPowerTokenTransactor(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.WrapTo(opts, accountAddress, amount)
}

func (c *Client) Unwrap(serviceAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newPowerTokenTransactor(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.Unwrap(opts, amount)
}

func (c *Client) SetBaseRate(serviceAddress, baseToken common.Address, baseRate, minGcFee *big.Int) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newPowerTokenTransactor(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetBaseRate(opts, baseRate, baseToken, minGcFee)
}

func (c *Client) SetLoanDurationLimits(serviceAddress common.Address, minLoanDuration, maxLoanDuration uint32) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newPowerTokenTransactor(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetLoanDurationLimits(opts, minLoanDuration, maxLoanDuration)
}

func (c *Client) SetServiceFeePercent(serviceAddress common.Address, feePercent uint16) (*types.Transaction, error) {
	transactor, opts, cancel, err := c.newPowerTokenTransactor(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return transactor.SetServiceFeePercent(opts, feePercent)
}

func (c *Client) GetAllowsPerpetual(serviceAddress common.Address) (bool, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return false, err
	}

	return res.AllowsPerpetual, nil
}

func (c *Client) GetBaseRate(serviceAddress common.Address) (*big.Int, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return nil, err
	}

	return res.BaseRate, nil
}

func (c *Client) GetBaseTokenAddress(serviceAddress common.Address) (common.Address, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return common.Address{}, err
	}

	return res.BaseToken, nil
}

func (c *Client) GetGapHalvingPeriod(serviceAddress common.Address) (int64, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return 0, err
	}

	return res.GapHalvingPeriod, nil
}

func (c *Client) GetMaxLoanDuration(serviceAddress common.Address) (int64, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return 0, err
	}

	return res.MaxLoanDuration, nil
}

func (c *Client) GetMinLoanDuration(serviceAddress common.Address) (int64, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return 0, err
	}

	return res.MinLoanDuration, nil
}

func (c *Client) GetMinGCFee(serviceAddress common.Address) (*big.Int, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return nil, err
	}

	return res.MinGCFee, nil
}

func (c *Client) GetServiceFeePercent(serviceAddress common.Address) (int64, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return 0, err
	}

	return res.ServiceFeePercent, nil
}

func (c *Client) GetServiceIndex(serviceAddress common.Address) (int64, error) {
	res, err := c.GetServiceInfo(serviceAddress)
	if err != nil {
		return 0, err
	}

	return res.Index, nil
}

func (c *Client) GetLiquidityTokenServiceAllowance(serviceAddress, accountAddress common.Address) (*big.Int, error) {
	transactor, err := bindings.NewPowerToken(serviceAddress, c.c)
	if err != nil {
		return nil, err
	}

	callOpts, cancel := c.callOpts()
	defer cancel()

	enterprise, err := transactor.GetEnterprise(callOpts)
	if err != nil {
		return nil, err
	}

	tkn, err := c.GetLiquidityTokenAddress(enterprise)
	if err != nil {
		return nil, err
	}

	erc20Caller, err := bindings.NewERC20Caller(tkn, c.c)
	if err != nil {
		return nil, err
	}

	return erc20Caller.Allowance(callOpts, accountAddress, serviceAddress)
}

func (c *Client) GetPowerTokenAvailableBalance(serviceAddress, accountAddress common.Address) (*big.Int, error) {
	caller, opts, cancel, err := c.newPowerTokenCaller(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.AvailableBalanceOf(opts, accountAddress)
}

func (c *Client) GetPowerTokenBalance(serviceAddress, accountAddress common.Address) (*big.Int, error) {
	caller, opts, cancel, err := c.newPowerTokenCaller(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.BalanceOf(opts, accountAddress)
}

func (c *Client) GetEnergyAt(serviceAddress, accountAddress common.Address, timestamp uint32) (*big.Int, error) {
	caller, opts, cancel, err := c.newPowerTokenCaller(serviceAddress)
	if err != nil {
		return nil, err
	}
	defer cancel()

	return caller.EnergyAt(opts, accountAddress, timestamp)
}

type LoanEstimateDetailed struct {
	Interest   *big.Int
	ServiceFee *big.Int
	GcFee      *big.Int
}

func (c *Client) EstimateLoanDetailed(serviceAddress, paymentTokenAddress common.Address, amount *big.Int, duration uint32) (LoanEstimateDetailed, error) {
	caller, opts, cancel, err := c.newPowerTokenCaller(serviceAddress)
	if err != nil {
		return LoanEstimateDetailed{}, err
	}
	defer cancel()

	info, err := caller.EstimateLoanDetailed(opts, paymentTokenAddress, amount, duration)
	if err != nil {
		return LoanEstimateDetailed{}, err
	}
	return LoanEstimateDetailed{
		Interest:   info.Interest,
		ServiceFee: info.ServiceFee,
		GcFee:      info.GcFee,
	}, nil
}

func (c *Client) GetServiceInfo(serviceAddress common.Address) (ServiceInfo, error) {
	caller, err := bindings.NewPowerTokenCaller(serviceAddress, c.c)
	if err != nil {
		return ServiceInfo{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.bcTimeout)
	defer cancel()

	res, err := caller.GetInfo(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return ServiceInfo{}, err
	}

	return ServiceInfo{
		Address:           serviceAddress,
		Name:              res.Name,
		Symbol:            res.Symbol,
		BaseRate:          res.BaseRate,
		MinGCFee:          res.MinGCFee,
		GapHalvingPeriod:  int64(res.GapHalvingPeriod),
		Index:             int64(res.Index),
		BaseToken:         res.BaseToken,
		MinLoanDuration:   int64(res.MinLoanDuration),
		MaxLoanDuration:   int64(res.MaxLoanDuration),
		ServiceFeePercent: int64(res.ServiceFeePercent),
		AllowsPerpetual:   res.AllowsPerpetual,
	}, nil
}

func (c *Client) GetAccountState(serviceAddress, accountAddress common.Address) (AccountState, error) {
	caller, err := bindings.NewPowerTokenCaller(serviceAddress, c.c)
	if err != nil {
		return AccountState{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.bcTimeout)
	defer cancel()
	balance, err := caller.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, accountAddress)
	if err != nil {
		return AccountState{}, err
	}

	state, err := caller.GetState(&bind.CallOpts{
		Context: ctx,
	}, accountAddress)
	if err != nil {
		return AccountState{}, err
	}

	return AccountState{
		ServiceAddress: serviceAddress,
		AccountAddress: accountAddress,
		Balance:        balance,
		Energy:         state.Energy,
		Timestamp:      int64(state.Timestamp),
	}, nil
}

type AccountState struct {
	ServiceAddress common.Address
	AccountAddress common.Address
	Balance        *big.Int
	Energy         *big.Int
	Timestamp      int64
}

type ServiceInfo struct {
	Address           common.Address
	Name              string
	Symbol            string
	BaseRate          *big.Int
	MinGCFee          *big.Int
	GapHalvingPeriod  int64
	Index             int64
	BaseToken         common.Address
	MinLoanDuration   int64
	MaxLoanDuration   int64
	ServiceFeePercent int64
	AllowsPerpetual   bool
}
