package account

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestAccountManager(t *testing.T) {
	t.Run("when account does not exist", func(t *testing.T) {
		t.Run("returns ErrNotFound upon account data request", func(t *testing.T) {
			store := NewMemoryStorage()
			manager := NewManager(store)

			addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
			assert.NoError(t, err)

			_, err = manager.GetAccount(addr)
			assert.True(t, errors.Is(err, ErrNotFound))
		})

		t.Run("returns ErrNotFound upon deletion", func(t *testing.T) {
			store := NewMemoryStorage()
			manager := NewManager(store)

			addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
			assert.NoError(t, err)

			err = manager.DeleteAccount(addr)
			assert.True(t, errors.Is(err, ErrNotFound))
		})

		t.Run("allows to create new account", func(t *testing.T) {
			store := NewMemoryStorage()
			manager := NewManager(store)

			addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
			assert.NoError(t, err)

			_, err = manager.CreateAccount(addr)
			assert.NoError(t, err)

			acc, err := manager.GetAccount(addr)
			assert.NoError(t, err)

			assert.Equal(t, acc.ID, addr.String())
		})
	})
	t.Run("when account exists", func(t *testing.T) {
		t.Run("allows to delete account", func(t *testing.T) {
			manager, address := initSeededManager(t)

			acc, err := manager.GetAccount(address)
			assert.NoError(t, err)
			assert.Equal(t, acc.ID, address.String())

			err = manager.DeleteAccount(address)
			assert.NoError(t, err)

			_, err = manager.GetAccount(address)
			assert.True(t, errors.Is(err, ErrNotFound))
		})
	})
}

func initSeededManager(t *testing.T) (*Manager, AccountID) {
	store := NewMemoryStorage()
	manager := NewManager(store)

	addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
	assert.NoError(t, err)

	_, err = manager.CreateAccount(addr)
	assert.NoError(t, err)

	return manager, addr
}
