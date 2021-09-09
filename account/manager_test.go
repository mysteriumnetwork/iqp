package account

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestAccountManager(t *testing.T) {
	t.Run("when account does not exist", func(t *testing.T) {
		t.Run("registers new account performing ownership verification", func(t *testing.T) {
			tempDir := t.TempDir()
			ks := keystore.NewKeyStore(tempDir, keystore.LightScryptN, keystore.LightScryptP)
			acc, err := ks.NewAccount("")
			assert.NoError(t, err)
			err = ks.Unlock(acc, "")
			assert.NoError(t, err)

			accountID, err := NewAccountID(ChainMaticMumbai, acc.Address)
			assert.NoError(t, err)

			claim, err := GenerateAccountOwnershipClaim(accountID, "1")
			assert.NoError(t, err)

			signature, err := SignClaim(acc.Address, claim, ks)
			assert.NoError(t, err)

			proof := OwnershipClaim{
				Signature: signature,
				Claim:     claim,
			}

			store := NewMemoryStorage()
			manager := NewManager(store, &OwnershipVerifyer{})

			t.Run("can register account with the generated claim", func(t *testing.T) {
				createdAccount, err := manager.RegisterAccount(accountID, proof)
				assert.NoError(t, err)

				assert.EqualValues(t, proof, createdAccount.OwnershipClaim)
				assert.Equal(t, accountID.String(), createdAccount.ID)
			})

			t.Run("can re-check the fetched claim", func(t *testing.T) {
				createdAccount, err := manager.GetAccount(accountID)
				assert.NoError(t, err)

				err = VerifyClaimSignature(createdAccount.OwnershipClaim.Claim, createdAccount.OwnershipClaim.Signature, accountID.Address())
				assert.NoError(t, err)
			})

			t.Run("will reject signature signed by another account", func(t *testing.T) {
				newAccount, err := ks.NewAccount("")
				assert.NoError(t, err)
				err = ks.Unlock(newAccount, "")
				assert.NoError(t, err)

				createdAccount, err := manager.GetAccount(accountID)
				assert.NoError(t, err)

				signature, err := SignClaim(newAccount.Address, createdAccount.OwnershipClaim.Claim, ks)
				assert.NoError(t, err)

				err = VerifyClaimSignature(createdAccount.OwnershipClaim.Claim, signature, accountID.Address())
				assert.Error(t, err)
			})
		})

		t.Run("returns ErrNotFound upon account data request", func(t *testing.T) {
			store := NewMemoryStorage()
			manager := NewManager(store, &OwnershipVerifyer{})

			addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
			assert.NoError(t, err)

			_, err = manager.GetAccount(addr)
			assert.ErrorIs(t, err, ErrNotFound)
		})

		t.Run("returns ErrNotFound upon deletion", func(t *testing.T) {
			store := NewMemoryStorage()
			manager := NewManager(store, &OwnershipVerifyer{})

			addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
			assert.NoError(t, err)

			err = manager.DeleteAccount(addr)
			assert.ErrorIs(t, err, ErrNotFound)
		})

		t.Run("allows to create new account", func(t *testing.T) {
			store := NewMemoryStorage()
			manager := NewManager(store, &OwnershipVerifyer{})

			addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
			assert.NoError(t, err)

			_, err = manager.CreateAccount(addr, OwnershipClaim{})
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
			assert.ErrorIs(t, err, ErrNotFound)
		})
	})
}

func initSeededManager(t *testing.T) (*Manager, AccountID) {
	store := NewMemoryStorage()
	manager := NewManager(store, &OwnershipVerifyer{})

	addr, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x01"))
	assert.NoError(t, err)

	_, err = manager.CreateAccount(addr, OwnershipClaim{})
	assert.NoError(t, err)

	return manager, addr
}
