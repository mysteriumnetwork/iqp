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
	storage           accountStorage
	ownershipVerifier OwnershipVerifier
}

func NewManager(storage accountStorage, ownershipVerifier OwnershipVerifier) *Manager {
	return &Manager{
		storage:           storage,
		ownershipVerifier: ownershipVerifier,
	}
}

type OwnershipVerifier interface {
	GenerateAndVerify(acc AccountID, signature []byte, version string) error
}

func (m *Manager) RegisterAccount(accountID AccountID, ownershipClaim OwnershipClaim) (Account, error) {
	err := m.ownershipVerifier.GenerateAndVerify(accountID, ownershipClaim.Signature, ownershipClaim.Claim.Domain.Version)
	if err != nil {
		return Account{}, err
	}

	return m.CreateAccount(accountID, ownershipClaim)
}

// Validation of ownership is up to the caller
func (m *Manager) CreateAccount(accountID AccountID, ownershipClaim OwnershipClaim) (Account, error) {
	acc, err := m.GetAccount(accountID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return acc, err
	}
	return m.storage.SaveAccount(Account{
		ID:             accountID.String(),
		OwnershipClaim: ownershipClaim,
	})
}

func (m *Manager) GetAccount(accountID AccountID) (Account, error) {
	return m.storage.GetAccount(accountID.String())
}

func (m *Manager) DeleteAccount(accountID AccountID) error {
	return m.storage.DeleteAccount(accountID.String())
}
