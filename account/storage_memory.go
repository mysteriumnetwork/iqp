package account

import (
	"errors"
	"fmt"
)

type MemoryStorage struct {
	accounts map[string]Account
	states   map[string]map[string]AccountState
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		accounts: make(map[string]Account),
		states:   make(map[string]map[string]AccountState),
	}
}

var ErrNotFound = errors.New("not found")

func (ms *MemoryStorage) GetAccount(id string) (Account, error) {
	acc, ok := ms.accounts[id]
	if !ok {
		return Account{}, fmt.Errorf("account %w", ErrNotFound)
	}

	return acc, nil
}

func (ms *MemoryStorage) DeleteAccount(id string) error {
	_, ok := ms.accounts[id]
	if !ok {
		return fmt.Errorf("account %w", ErrNotFound)
	}

	delete(ms.accounts, id)
	return nil
}

func (ms *MemoryStorage) GetAccountState(serviceID string, accountID string) (AccountState, error) {
	v, ok := ms.states[accountID]
	if !ok {
		return AccountState{}, fmt.Errorf("account %w", ErrNotFound)
	}

	s, ok := v[serviceID]
	if !ok {
		return AccountState{}, fmt.Errorf("service %w", ErrNotFound)
	}

	return s, nil
}

func (ms *MemoryStorage) DeleteAccountState(serviceID string, accountID string) error {
	v, ok := ms.states[accountID]
	if !ok {
		return fmt.Errorf("account %w", ErrNotFound)
	}

	if _, ok := v[serviceID]; !ok {
		return fmt.Errorf("service %w", ErrNotFound)
	}

	delete(v, serviceID)
	return nil
}

func (ms *MemoryStorage) SaveAccount(account Account) (Account, error) {
	ms.accounts[account.ID] = account
	return account, nil
}

var ErrAlreadyInitialized = errors.New("already initialized")

func (ms *MemoryStorage) InitAccountState(accountState AccountState) (AccountState, error) {
	_, ok := ms.accounts[accountState.AccountID]
	if !ok {
		ms.accounts[accountState.AccountID] = Account{ID: accountState.AccountID}
	}

	if states, ok := ms.states[accountState.AccountID]; ok {
		if _, ok := states[accountState.ServiceID]; ok {
			return AccountState{}, ErrAlreadyInitialized
		}
	}

	ms.states[accountState.AccountID] = map[string]AccountState{
		accountState.ServiceID: accountState,
	}

	return accountState, nil
}

var ErrStateNotInitialized = errors.New("state not initialized")

// TODO: previous seems redundant
func (ms *MemoryStorage) ChangeAccountState(previous, new AccountState) (AccountState, error) {
	_, ok := ms.accounts[previous.AccountID]
	if !ok {
		return AccountState{}, fmt.Errorf("account %w", ErrNotFound)
	}

	if _, ok := ms.states[previous.AccountID]; !ok {
		return AccountState{}, ErrStateNotInitialized
	}

	if _, ok := ms.states[previous.AccountID][previous.ServiceID]; !ok {
		return AccountState{}, ErrStateNotInitialized
	}

	ms.states[previous.AccountID][previous.ServiceID] = new
	return new, nil
}
