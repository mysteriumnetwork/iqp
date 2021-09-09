package account

import (
	"errors"
	"fmt"
	"math/big"
)

type DefaultStateValidator struct{}

var ErrAccountStateValidation = errors.New("invalid account state")

func (d DefaultStateValidator) Validate(as AccountState) error {
	if as.Power == nil {
		return fmt.Errorf("power cannot be nil %w", ErrAccountStateValidation)
	}

	if as.LockedPower == nil {
		return fmt.Errorf("locked power can not be nil %w", ErrAccountStateValidation)
	}

	if as.EnergyCap == nil {
		return fmt.Errorf("energy cap can not be nil %w", ErrAccountStateValidation)
	}

	if as.Energy == nil {
		return fmt.Errorf("energy can not be nil %w", ErrAccountStateValidation)
	}

	if as.Power.Cmp(new(big.Int)) == -1 {
		return fmt.Errorf("power can not be negative %w", ErrAccountStateValidation)
	}

	if as.LockedPower.Cmp(new(big.Int)) == -1 {
		return fmt.Errorf("power can not be negative %w", ErrAccountStateValidation)
	}

	if as.GapHalvingPeriod <= 0 {
		return fmt.Errorf("gap halving period must be positive %w", ErrAccountStateValidation)
	}

	if as.Energy.Cmp(as.EnergyCap) == 1 {
		return fmt.Errorf("energy above cap %w", ErrAccountStateValidation)
	}

	if as.EnergyCalculatedAt < 0 {
		return fmt.Errorf("negative energy calculation cap %w", ErrAccountStateValidation)
	}
	return nil
}
