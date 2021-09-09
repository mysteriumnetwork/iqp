package account

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/iqp/blockchain"
	"github.com/mysteriumnetwork/iqp/energy"
)

type accountStateStorage interface {
	GetAccountState(serviceID string, accountID string) (AccountState, error)
	DeleteAccountState(serviceID string, accountID string) error
	InitAccountState(accountState AccountState) (AccountState, error)
	ChangeAccountState(previous, new AccountState) (AccountState, error)
}

type bc interface {
	GetServiceInfo(serviceAddress common.Address) (blockchain.ServiceInfo, error)
	GetAccountState(serviceAddress, accountAddress common.Address) (blockchain.AccountState, error)
}

type StateValidator interface {
	Validate(as AccountState) error
}

type StateManager struct {
	storage accountStateStorage
	bc      bc
	sv      StateValidator
}

func NewStateManager(storage accountStateStorage, bc bc, sv StateValidator) *StateManager {
	return &StateManager{
		storage: storage,
		bc:      bc,
		sv:      sv,
	}
}

var ErrNotMatchinChains = errors.New("chains are not matching")

func chainsMatch(a, b AccountID) bool {
	return a.ChainID() == b.ChainID()
}

func checkChains(a, b AccountID) error {
	if chainsMatch(a, b) {
		return nil
	}

	return fmt.Errorf("chain1: %v, chain2: %v, err: %w", a.chainID, b.chainID, ErrNotMatchinChains)
}

func (sm *StateManager) InitAccountState(
	serviceID, accountID AccountID,
	gapHalvingPeriod int64,
	power, lockedPower, energyCap, energy *big.Int,
	energyCalculatedAt int64) (AccountState, error) {
	if err := checkChains(serviceID, accountID); err != nil {
		return AccountState{}, err
	}

	as := AccountState{
		ServiceID:          serviceID.String(),
		AccountID:          accountID.String(),
		GapHalvingPeriod:   gapHalvingPeriod,
		Power:              power,
		LockedPower:        lockedPower,
		EnergyCap:          energyCap,
		Energy:             energy,
		EnergyCalculatedAt: energyCalculatedAt,
	}

	err := sm.sv.Validate(as)
	if err != nil {
		return AccountState{}, err
	}

	return sm.storage.InitAccountState(as)
}

func (sm *StateManager) InitAccountStateFromBlockchain(serviceID, accountID AccountID) (AccountState, error) {
	if err := checkChains(serviceID, accountID); err != nil {
		return AccountState{}, err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	var accState blockchain.AccountState
	var accStateErr error
	go func() {
		defer wg.Done()
		accState, accStateErr = sm.bc.GetAccountState(serviceID.address, accountID.address)
	}()

	// TODO: this will likely only change very rarely, so there's probably little need to look up all the info all the time.
	var serviceInfo blockchain.ServiceInfo
	var serviceInfoErr error
	go func() {
		defer wg.Done()
		serviceInfo, serviceInfoErr = sm.bc.GetServiceInfo(serviceID.address)
	}()

	wg.Wait()

	if accStateErr != nil {
		return AccountState{}, fmt.Errorf("could not get account state from bc: %w", accStateErr)
	}

	if serviceInfoErr != nil {
		return AccountState{}, fmt.Errorf("could not get service info from bc: %w", serviceInfoErr)
	}

	return sm.InitAccountState(
		serviceID, accountID,
		serviceInfo.GapHalvingPeriod,
		accState.Balance,
		new(big.Int),
		accState.Energy,
		accState.Energy,
		accState.Timestamp,
	)
}

func (sm *StateManager) GetBlockchainAccountState(serviceID, accountID AccountID) (blockchain.AccountState, error) {
	return sm.bc.GetAccountState(serviceID.address, accountID.address)
}

func (sm *StateManager) GetAccountState(serviceID, accountID AccountID) (AccountState, error) {
	return sm.storage.GetAccountState(serviceID.String(), accountID.String())
}

func (sm *StateManager) DeleteAccountState(serviceID, accountID AccountID) error {
	return sm.storage.DeleteAccountState(serviceID.String(), accountID.String())
}

func (sm *StateManager) GetOrInitializeAccountState(serviceID, accountID AccountID) (AccountState, error) {
	if err := checkChains(serviceID, accountID); err != nil {
		return AccountState{}, err
	}

	state, err := sm.GetAccountState(serviceID, accountID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return sm.InitAccountStateFromBlockchain(serviceID, accountID)
		}
		return AccountState{}, err
	}

	return state, nil
}

type stateChangeType string

const (
	stateChangeEnergy      stateChangeType = "energy"
	stateChangePower       stateChangeType = "power"
	stateChangeLockedPower stateChangeType = "locked_power"
)

func (sm *StateManager) changeState(serviceID, accountID AccountID, key stateChangeType, delta *big.Int, timestamp int64) (AccountState, error) {
	currentState, err := sm.GetOrInitializeAccountState(serviceID, accountID)
	if err != nil {
		return AccountState{}, err
	}
	power := currentState.Power
	lockedPower := currentState.LockedPower
	effectivePower := new(big.Int).Sub(power, lockedPower)

	switch key {
	case stateChangePower:
		power = new(big.Int).Add(power, delta)
		effectivePower = new(big.Int).Add(effectivePower, delta)
	case stateChangeLockedPower:
		lockedPower = new(big.Int).Add(lockedPower, delta)
		effectivePower = new(big.Int).Sub(effectivePower, delta)
	}

	effective, err := energy.CalculateEffectiveEnergy(
		energy.EffectiveEnergyCalculationParams{
			Energy:    currentState.Energy,
			EnergyCap: currentState.EnergyCap,
			Power:     effectivePower,
			GapHalvingParams: energy.GapHalvingParams{
				GapHalvingPeriod: currentState.GapHalvingPeriod,
				T0:               currentState.EnergyCalculatedAt,
				T1:               timestamp,
			},
		},
	)
	if err != nil {
		return AccountState{}, err
	}

	if key == stateChangeEnergy {
		// Energy cannot be increased directly
		if delta.Sign() == 1 {
			return AccountState{}, ErrPositiveEnergyDelta
		}
		effective.Energy = new(big.Int).Add(effective.Energy, delta)
	}

	newState := AccountState{
		Power:              power,
		LockedPower:        lockedPower,
		EnergyCap:          effective.EnergyCap,
		Energy:             effective.Energy,
		EnergyCalculatedAt: timestamp,
		ServiceID:          currentState.ServiceID,
		AccountID:          currentState.AccountID,
		GapHalvingPeriod:   currentState.GapHalvingPeriod,
	}

	fmt.Println(newState.LockedPower)
	err = sm.sv.Validate(newState)
	if err != nil {
		return AccountState{}, err
	}

	return sm.storage.ChangeAccountState(currentState, newState)

}

var ErrPositiveEnergyDelta = errors.New("positive energy delta is not allowed")

func (sm *StateManager) IncreasePower(serviceID, accountID AccountID, delta *big.Int, timestamp int64) (AccountState, error) {
	if err := validateDelta(delta); err != nil {
		return AccountState{}, err
	}

	return sm.changeState(serviceID, accountID, stateChangePower, delta, timestamp)
}

func (sm *StateManager) DecreasePower(serviceID, accountID AccountID, delta *big.Int, timestamp int64) (AccountState, error) {
	if err := validateDelta(delta); err != nil {
		return AccountState{}, err
	}

	return sm.changeState(serviceID, accountID, stateChangePower, new(big.Int).Neg(delta), timestamp)
}

func (sm *StateManager) LockPower(serviceID, accountID AccountID, delta *big.Int, timestamp int64) (AccountState, error) {
	if err := validateDelta(delta); err != nil {
		return AccountState{}, err
	}

	return sm.changeState(serviceID, accountID, stateChangeLockedPower, delta, timestamp)
}

func (sm *StateManager) UnlockPower(serviceID, accountID AccountID, delta *big.Int, timestamp int64) (AccountState, error) {
	if err := validateDelta(delta); err != nil {
		return AccountState{}, err
	}

	return sm.changeState(serviceID, accountID, stateChangeLockedPower, new(big.Int).Neg(delta), timestamp)
}

func (sm *StateManager) SpendEnergy(serviceID, accountID AccountID, delta *big.Int, timestamp int64) (AccountState, error) {
	if err := validateDelta(delta); err != nil {
		return AccountState{}, err
	}

	return sm.changeState(serviceID, accountID, stateChangeEnergy, new(big.Int).Neg(delta), timestamp)
}

var ErrNegativeDelta = errors.New("delta should be positive")

func validateDelta(delta *big.Int) error {
	if delta.Sign() >= 0 {
		return nil
	}

	return ErrNegativeDelta
}
