package account

import (
	"errors"
)

type accountStorage interface {
	SaveAccount(account Account) (Account, error)
	GetAccount(id string) (Account, error)
	DeleteAccount(id string) error
}

type Manager struct {
	storage accountStorage
}

func NewManager(storage accountStorage) *Manager {
	return &Manager{
		storage: storage,
	}
}

// Validation of ownership is up to the caller
func (m *Manager) CreateAccount(accountID AccountID) (Account, error) {
	acc, err := m.GetAccount(accountID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return acc, err
	}
	return m.storage.SaveAccount(Account{
		ID: accountID.String(),
	})
}

func (m *Manager) GetAccount(accountID AccountID) (Account, error) {
	return m.storage.GetAccount(accountID.String())
}

func (m *Manager) DeleteAccount(accountID AccountID) error {
	return m.storage.DeleteAccount(accountID.String())
}
