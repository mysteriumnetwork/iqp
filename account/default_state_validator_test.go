package account

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestDefaultStateValidator(t *testing.T) {
	serviceAccountID, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x1"))
	assert.NoError(t, err)

	accountID, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x2"))
	assert.NoError(t, err)

	initialStateGetter := func() AccountState {
		return AccountState{
			AccountID:          accountID.String(),
			ServiceID:          serviceAccountID.String(),
			GapHalvingPeriod:   100,
			Power:              big.NewInt(1000),
			LockedPower:        big.NewInt(10),
			EnergyCap:          big.NewInt(500),
			Energy:             big.NewInt(500),
			EnergyCalculatedAt: time.Now().Unix(),
		}
	}

	validator := DefaultStateValidator{}

	t.Run("does not allow nil power", func(t *testing.T) {
		state := initialStateGetter()
		state.Power = nil

		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "power can not be nil")
	})

	t.Run("does not allow nil locked power", func(t *testing.T) {
		state := initialStateGetter()
		state.LockedPower = nil
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "locked power can not be nil")
	})

	t.Run("does not allow nil energy cap", func(t *testing.T) {
		state := initialStateGetter()
		state.EnergyCap = nil
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "energy cap can not be nil")
	})

	t.Run("does not allow nil energy", func(t *testing.T) {
		state := initialStateGetter()
		state.Energy = nil
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "energy can not be nil")
	})

	t.Run("does not allow negative power", func(t *testing.T) {
		state := initialStateGetter()
		state.Power = new(big.Int).Neg(state.Power)
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "power can not be negative")
	})

	t.Run("does not allow negative locked power", func(t *testing.T) {
		state := initialStateGetter()
		state.LockedPower = new(big.Int).Neg(state.LockedPower)
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "locked power can not be negative")
	})

	t.Run("does not allow negative gap halving period", func(t *testing.T) {
		state := initialStateGetter()
		state.GapHalvingPeriod = -state.GapHalvingPeriod
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "gap halving period must be positive")
	})

	t.Run("does not allow energy above cap", func(t *testing.T) {
		state := initialStateGetter()
		state.Energy = new(big.Int).Mul(state.EnergyCap, big.NewInt(5))
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "energy above cap")
	})

	t.Run("does not allow negative calculation date", func(t *testing.T) {
		state := initialStateGetter()
		state.EnergyCalculatedAt = -state.EnergyCalculatedAt
		err := validator.Validate(state)
		assert.ErrorIs(t, err, ErrAccountStateValidation)
		assert.Contains(t, err.Error(), "negative energy calculation")
	})
}
