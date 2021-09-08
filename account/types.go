package account

import (
	"math/big"
)

type Account struct {
	ID string
}

type AccountState struct {
	ServiceID          string
	AccountID          string
	GapHalvingPeriod   int64
	Power              *big.Int
	EnergyCap          *big.Int
	Energy             *big.Int
	EnergyCalculatedAt int64
}
