package account

import (
	"math/big"

	"github.com/ethereum/go-ethereum/signer/core"
)

type Account struct {
	ID             string         `json:"id,omitempty"`
	OwnershipClaim OwnershipClaim `json:"ownership_claim,omitempty"`
}

type OwnershipClaim struct {
	Claim     core.TypedData `json:"claim,omitempty"`
	Signature []byte         `json:"signature,omitempty"`
}

type AccountState struct {
	ServiceID          string   `json:"service_id,omitempty"`
	AccountID          string   `json:"account_id,omitempty"`
	GapHalvingPeriod   int64    `json:"gap_halving_period,omitempty"`
	Power              *big.Int `json:"power,omitempty"`
	LockedPower        *big.Int `json:"locked_power,omitempty"`
	EnergyCap          *big.Int `json:"energy_cap,omitempty"`
	Energy             *big.Int `json:"energy,omitempty"`
	EnergyCalculatedAt int64    `json:"energy_calculated_at,omitempty"`
}
