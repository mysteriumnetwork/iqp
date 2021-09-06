package energy

import (
	"errors"
	"math/big"
)

var one *big.Int
var logHalf *big.Int
var maxSafeUint112 *big.Int

func init() {
	one, _ = new(big.Int).SetString("22300745198530623141535718272648361505980416", 10)
	logHalf, _ = new(big.Int).SetString("15457698658747239244624307340191628289589491", 10)
	maxSafeUint112, _ = new(big.Int).SetString("5192296858534827628530496329220095", 10)
}

type GapHalvingParams struct {
	GapHalvingPeriod int64
	T0               int64
	T1               int64
}

type HalfLifeParams struct {
	InitialValue *big.Int
	GapHalvingParams
}

type EnergyCalculationParams struct {
	HalfLifeParams
	Power *big.Int
}

type EffectiveEnergyCalculationParams struct {
	GapHalvingParams
	Power     *big.Int
	Energy    *big.Int
	EnergyCap *big.Int
}

func HalfLife(in HalfLifeParams) (*big.Int, error) {
	initialValue := in.InitialValue
	if initialValue.Sign() == -1 {
		return nil, errors.New("initial value underflow")
	}

	if initialValue.Cmp(maxSafeUint112) == 1 {
		return nil, errors.New("initial value overflow")
	}

	period := big.NewInt(0).SetInt64(in.T1 - in.T0)
	if period.Sign() == -1 {
		return nil, errors.New("negative period")
	}

	if period.Cmp(new(big.Int)) == 0 {
		return initialValue, nil
	}

	initialValue = new(big.Int).Rsh(initialValue, uint(new(big.Int).Div(period, big.NewInt(in.GapHalvingPeriod)).Uint64()))
	if initialValue.Cmp(new(big.Int)) == 0 {
		return initialValue, nil
	}

	period = new(big.Int).Mod(period, big.NewInt(in.GapHalvingPeriod))
	x := new(big.Int).Div(new(big.Int).Mul(logHalf, period), big.NewInt(in.GapHalvingPeriod))
	z := new(big.Int).Set(initialValue)
	i := new(big.Int).Set(one)
	sum := new(big.Int)

	for z.Cmp(new(big.Int)) != 0 {
		sum = new(big.Int).Add(sum, z)
		z = new(big.Int).Div(new(big.Int).Mul(z, x), i)
		i = new(big.Int).Add(i, one)
		sum = new(big.Int).Sub(sum, z)
		z = new(big.Int).Div(new(big.Int).Mul(z, x), i)
		i = new(big.Int).Add(i, one)
	}

	if sum.Cmp(maxSafeUint112) > 0 {
		return nil, errors.New("result overflow")
	}

	return sum, nil
}

func CalculateLinearEnergy(in EnergyCalculationParams) (*big.Int, error) {
	period := big.NewInt(0).SetInt64(in.T1 - in.T0)
	if period.Sign() == -1 {
		return nil, errors.New("negative period")
	}

	return new(big.Int).Add(in.InitialValue, new(big.Int).Div(new(big.Int).Mul(in.Power, period), new(big.Int).Mul(big.NewInt(in.GapHalvingPeriod), big.NewInt(4)))), nil
}

func CalculateEnergyCap(in EnergyCalculationParams) (*big.Int, error) {
	if in.Power.Cmp(in.InitialValue) > 0 {
		hl, err := HalfLife(HalfLifeParams{
			InitialValue:     new(big.Int).Sub(in.Power, in.InitialValue),
			GapHalvingParams: in.GapHalvingParams,
		})
		if err != nil {
			return nil, err
		}
		return new(big.Int).Sub(in.Power, hl), nil
	}

	hl, err := HalfLife(HalfLifeParams{
		InitialValue:     new(big.Int).Sub(in.InitialValue, in.Power),
		GapHalvingParams: in.GapHalvingParams,
	})
	if err != nil {
		return nil, err
	}
	return new(big.Int).Add(in.Power, hl), nil
}

type EffectiveEnergyResult struct {
	EnergyCap    *big.Int
	LinearEnergy *big.Int
	Energy       *big.Int
}

func CalculateEffectiveEnergy(in EffectiveEnergyCalculationParams) (EffectiveEnergyResult, error) {
	energyCap, err := CalculateEnergyCap(EnergyCalculationParams{
		Power: in.Power,
		HalfLifeParams: HalfLifeParams{
			InitialValue:     in.EnergyCap,
			GapHalvingParams: in.GapHalvingParams,
		},
	})
	if err != nil {
		return EffectiveEnergyResult{}, err
	}

	linearEnergy, err := CalculateLinearEnergy(EnergyCalculationParams{
		Power: in.Power,
		HalfLifeParams: HalfLifeParams{
			InitialValue:     in.Energy,
			GapHalvingParams: in.GapHalvingParams,
		},
	})
	if err != nil {
		return EffectiveEnergyResult{}, err
	}

	energy := energyCap
	if linearEnergy.Cmp(energyCap) < 0 {
		energy = linearEnergy
	}

	return EffectiveEnergyResult{
		EnergyCap:    energyCap,
		LinearEnergy: linearEnergy,
		Energy:       energy,
	}, nil
}
