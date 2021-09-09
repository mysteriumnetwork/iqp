package account

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/iqp/blockchain"
	"github.com/stretchr/testify/assert"
)

func TestAccountStateManager(t *testing.T) {
	serviceAccountID, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x1"))
	assert.NoError(t, err)

	accountID, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x2"))
	assert.NoError(t, err)

	mockBcAccountState := blockchain.AccountState{
		ServiceAddress: serviceAccountID.address,
		AccountAddress: accountID.address,
		Balance:        big.NewInt(1000),
		Energy:         big.NewInt(500),
		Timestamp:      time.Now().Unix(),
	}
	mockBCServiceState := blockchain.ServiceInfo{
		Address:           serviceAccountID.address,
		Name:              "Test Service",
		Symbol:            "TST",
		BaseRate:          big.NewInt(1),
		MinGCFee:          big.NewInt(1),
		GapHalvingPeriod:  3600,
		Index:             0,
		BaseToken:         common.HexToAddress("0x6238F38c32fd76E3189D1EAd943B8342Ff33055D"),
		MinLoanDuration:   3600,
		MaxLoanDuration:   864000,
		ServiceFeePercent: 5000,
		AllowsPerpetual:   false,
	}

	defaultAccountState := AccountState{
		AccountID:          accountID.String(),
		ServiceID:          serviceAccountID.String(),
		GapHalvingPeriod:   mockBCServiceState.GapHalvingPeriod,
		Power:              big.NewInt(1000),
		LockedPower:        big.NewInt(0),
		EnergyCap:          big.NewInt(500),
		Energy:             big.NewInt(500),
		EnergyCalculatedAt: mockBcAccountState.Timestamp,
	}

	t.Run("returns errors on invalid deltas", func(t *testing.T) {
		manager := NewStateManager(NewMemoryStorage(), &mockBC{ServiceInfoError: ErrNotFound}, DefaultStateValidator{})
		t.Run("returns err if delta negative", func(t *testing.T) {
			_, err := manager.DecreasePower(serviceAccountID, accountID, big.NewInt(-1), 0)
			assert.ErrorIs(t, err, ErrNegativeDelta)

			_, err = manager.IncreasePower(serviceAccountID, accountID, big.NewInt(-1), 0)
			assert.ErrorIs(t, err, ErrNegativeDelta)

			_, err = manager.LockPower(serviceAccountID, accountID, big.NewInt(-1), 0)
			assert.ErrorIs(t, err, ErrNegativeDelta)

			_, err = manager.UnlockPower(serviceAccountID, accountID, big.NewInt(-1), 0)
			assert.ErrorIs(t, err, ErrNegativeDelta)

			_, err = manager.SpendEnergy(serviceAccountID, accountID, big.NewInt(-1), 0)
			assert.ErrorIs(t, err, ErrNegativeDelta)
		})

		t.Run("returns bc err if does not exist", func(t *testing.T) {
			_, err := manager.DecreasePower(serviceAccountID, accountID, big.NewInt(1), 0)
			assert.ErrorIs(t, err, ErrNotFound)

			_, err = manager.IncreasePower(serviceAccountID, accountID, big.NewInt(1), 0)
			assert.ErrorIs(t, err, ErrNotFound)

			_, err = manager.LockPower(serviceAccountID, accountID, big.NewInt(1), 0)
			assert.ErrorIs(t, err, ErrNotFound)

			_, err = manager.UnlockPower(serviceAccountID, accountID, big.NewInt(1), 0)
			assert.ErrorIs(t, err, ErrNotFound)

			_, err = manager.SpendEnergy(serviceAccountID, accountID, big.NewInt(1), 0)
			assert.ErrorIs(t, err, ErrNotFound)
		})
	})
	t.Run("when account is registered", func(t *testing.T) {
		t.Run("when account state is not initialized", func(t *testing.T) {
			bc := &mockBC{}
			store := NewMemoryStorage()
			manager := NewStateManager(store, bc, DefaultStateValidator{})
			_, err := store.SaveAccount(Account{ID: accountID.String()})
			assert.NoError(t, err)

			bc.AccountStateResponse = mockBcAccountState

			bc.ServiceInfoResponse = mockBCServiceState

			t.Run("returns err not found upon account state request", func(t *testing.T) {
				_, err := manager.GetAccountState(serviceAccountID, accountID)
				assert.ErrorIs(t, err, ErrNotFound)
			})

			t.Run("returns state from blockchain", func(t *testing.T) {
				accountState, err := manager.GetBlockchainAccountState(serviceAccountID, accountID)
				assert.NoError(t, err)
				assert.EqualValues(t, mockBcAccountState.Balance, accountState.Balance)
				assert.EqualValues(t, mockBcAccountState.Energy, accountState.Energy)
				assert.EqualValues(t, mockBcAccountState.Timestamp, accountState.Timestamp)
			})

			t.Run("returns initialized state", func(t *testing.T) {
				expectedState := defaultAccountState
				accountState, err := manager.GetOrInitializeAccountState(serviceAccountID, accountID)
				assert.NoError(t, err)
				assert.EqualValues(t, expectedState, accountState)
			})

			t.Run("increases power based on blockchain data", func(t *testing.T) {
				delta := big.NewInt(500)
				state, err := manager.IncreasePower(serviceAccountID, accountID, delta, time.Now().Unix())
				assert.NoError(t, err)
				assert.Equal(t, big.NewInt(0).Add(delta, mockBcAccountState.Balance), state.Power)
			})

			t.Run("decreases power based on blockchain data", func(t *testing.T) {
				delta := big.NewInt(500)
				state, err := manager.DecreasePower(serviceAccountID, accountID, delta, time.Now().Unix())
				assert.NoError(t, err)
				assert.Equal(t, mockBcAccountState.Balance, state.Power)
			})

			t.Run("locks power", func(t *testing.T) {
				delta := big.NewInt(300)
				state, err := manager.LockPower(serviceAccountID, accountID, delta, time.Now().Unix())
				assert.NoError(t, err)
				assert.Equal(t, delta, state.LockedPower)
			})

			t.Run("spends energy based on blockchain data", func(t *testing.T) {
				delta := big.NewInt(255)
				state, err := manager.SpendEnergy(serviceAccountID, accountID, delta, time.Now().Unix())
				assert.NoError(t, err)
				assert.Equal(t, big.NewInt(245), state.Energy)
				assert.Equal(t, big.NewInt(500), state.EnergyCap)
			})

			t.Run("unlocks power", func(t *testing.T) {
				delta := big.NewInt(300)
				state, err := manager.UnlockPower(serviceAccountID, accountID, delta, time.Now().Unix())
				assert.NoError(t, err)
				assert.Equal(t, big.NewInt(0).Uint64(), state.LockedPower.Uint64())
			})

			t.Run("errors if unlocking power when none is locked", func(t *testing.T) {
				delta := big.NewInt(1)
				_, err := manager.UnlockPower(serviceAccountID, accountID, delta, time.Now().Unix())
				assert.ErrorIs(t, err, ErrAccountStateValidation)
			})
		})
		t.Run("when account is initialized", func(t *testing.T) {
			bc := &mockBC{}
			store := NewMemoryStorage()
			manager := NewStateManager(store, bc, DefaultStateValidator{})
			_, err := store.SaveAccount(Account{ID: accountID.String()})
			assert.NoError(t, err)

			changeTimestamp := time.Now().Unix()
			changeTimestamp = (changeTimestamp + mockBCServiceState.GapHalvingPeriod)

			initialState := defaultAccountState
			initialState.Power = big.NewInt(1500)
			initialState.LockedPower = big.NewInt(700)

			_, err = store.InitAccountState(initialState)
			assert.NoError(t, err)

			reinitializeState := func(t *testing.T) {
				err := store.DeleteAccountState(initialState.ServiceID, initialState.AccountID)
				assert.NoError(t, err)
				_, err = store.InitAccountState(initialState)
				assert.NoError(t, err)
			}

			t.Run("returns correct state", func(t *testing.T) {
				state, err := manager.GetAccountState(serviceAccountID, accountID)
				assert.NoError(t, err)
				assert.EqualValues(t, initialState, state)
			})

			t.Run("returns initialized state", func(t *testing.T) {
				state, err := manager.GetOrInitializeAccountState(serviceAccountID, accountID)
				assert.NoError(t, err)
				assert.EqualValues(t, initialState, state)
			})

			t.Run("increases power based on stored data", func(t *testing.T) {
				delta := big.NewInt(300)
				state, err := manager.IncreasePower(serviceAccountID, accountID, delta, changeTimestamp)
				assert.NoError(t, err)

				assert.Equal(t, new(big.Int).Add(initialState.Power, delta), state.Power)
				assert.Equal(t, big.NewInt(775), state.Energy)
				assert.Equal(t, big.NewInt(800), state.EnergyCap)
				assert.Equal(t, changeTimestamp, state.EnergyCalculatedAt)
				assert.Equal(t, initialState.AccountID, state.AccountID)
				assert.Equal(t, initialState.ServiceID, state.ServiceID)
				assert.Equal(t, initialState.GapHalvingPeriod, state.GapHalvingPeriod)
				assert.Equal(t, initialState.LockedPower, state.LockedPower)
			})

			t.Run("decreases power based on stored data", func(t *testing.T) {
				reinitializeState(t)
				delta := big.NewInt(300)
				state, err := manager.DecreasePower(serviceAccountID, accountID, delta, changeTimestamp)
				assert.NoError(t, err)
				assert.Equal(t, new(big.Int).Sub(initialState.Power, delta), state.Power)
				assert.Equal(t, big.NewInt(500), state.Energy)
				assert.Equal(t, big.NewInt(500), state.EnergyCap)
				assert.Equal(t, changeTimestamp, state.EnergyCalculatedAt)
				assert.Equal(t, initialState.AccountID, state.AccountID)
				assert.Equal(t, initialState.ServiceID, state.ServiceID)
				assert.Equal(t, initialState.GapHalvingPeriod, state.GapHalvingPeriod)
				assert.Equal(t, initialState.LockedPower, state.LockedPower)
			})

			t.Run("locks power based on stored data", func(t *testing.T) {
				reinitializeState(t)
				delta := big.NewInt(500)
				state, err := manager.LockPower(serviceAccountID, accountID, delta, changeTimestamp)
				assert.NoError(t, err)
				assert.Equal(t, new(big.Int).Add(initialState.LockedPower, delta), state.LockedPower)
				assert.Equal(t, big.NewInt(400), state.Energy)
				assert.Equal(t, big.NewInt(400), state.EnergyCap)
				assert.Equal(t, changeTimestamp, state.EnergyCalculatedAt)
				assert.Equal(t, initialState.AccountID, state.AccountID)
				assert.Equal(t, initialState.ServiceID, state.ServiceID)
				assert.Equal(t, initialState.GapHalvingPeriod, state.GapHalvingPeriod)
				assert.Equal(t, initialState.Power, state.Power)
			})
			t.Run("unlocks power based on stored data", func(t *testing.T) {
				reinitializeState(t)
				delta := big.NewInt(650)
				state, err := manager.UnlockPower(serviceAccountID, accountID, delta, changeTimestamp)
				assert.NoError(t, err)
				assert.Equal(t, new(big.Int).Sub(initialState.LockedPower, delta), state.LockedPower)
				assert.Equal(t, big.NewInt(862), state.Energy)
				assert.Equal(t, big.NewInt(975), state.EnergyCap)
				assert.Equal(t, changeTimestamp, state.EnergyCalculatedAt)
				assert.Equal(t, initialState.AccountID, state.AccountID)
				assert.Equal(t, initialState.ServiceID, state.ServiceID)
				assert.Equal(t, initialState.GapHalvingPeriod, state.GapHalvingPeriod)
				assert.Equal(t, initialState.Power, state.Power)
			})
			t.Run("spends energy based on stored data", func(t *testing.T) {
				reinitializeState(t)
				delta := big.NewInt(255)
				state, err := manager.SpendEnergy(serviceAccountID, accountID, delta, changeTimestamp)
				assert.NoError(t, err)
				assert.Equal(t, big.NewInt(395), state.Energy)
				assert.Equal(t, big.NewInt(650), state.EnergyCap)
				assert.Equal(t, changeTimestamp, state.EnergyCalculatedAt)
				assert.Equal(t, initialState.AccountID, state.AccountID)
				assert.Equal(t, initialState.ServiceID, state.ServiceID)
				assert.Equal(t, initialState.GapHalvingPeriod, state.GapHalvingPeriod)
				assert.Equal(t, initialState.Power, state.Power)
				assert.Equal(t, initialState.LockedPower, state.LockedPower)
			})
			t.Run("allows to delete account state for specific service", func(t *testing.T) {
				t.Run("deletes existing account", func(t *testing.T) {
					err := manager.DeleteAccountState(serviceAccountID, accountID)
					assert.NoError(t, err)
					_, err = manager.GetAccountState(serviceAccountID, accountID)
					assert.ErrorIs(t, err, ErrNotFound)
				})
				t.Run("errors on trying to delete a non existing account", func(t *testing.T) {
					err := manager.DeleteAccountState(serviceAccountID, accountID)
					assert.ErrorIs(t, err, ErrNotFound)
				})
			})
		})
	})
}

type mockBC struct {
	ServiceInfoResponse  blockchain.ServiceInfo
	ServiceInfoError     error
	AccountStateResponse blockchain.AccountState
	AccountStateError    error
}

func (mbc *mockBC) GetAccountState(serviceAddress, accountAddress common.Address) (blockchain.AccountState, error) {
	return mbc.AccountStateResponse, mbc.AccountStateError
}

func (mbc *mockBC) GetServiceInfo(serviceAddress common.Address) (blockchain.ServiceInfo, error) {
	return mbc.ServiceInfoResponse, mbc.ServiceInfoError
}
