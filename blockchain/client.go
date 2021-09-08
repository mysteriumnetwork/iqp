package blockchain

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mysteriumnetwork/iqp/bindings"
)

type Client struct {
	c         *ethclient.Client
	bcTimeout time.Duration
}

func NewClient(c *ethclient.Client, bcTimeout time.Duration) *Client {
	return &Client{
		c:         c,
		bcTimeout: bcTimeout,
	}
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
