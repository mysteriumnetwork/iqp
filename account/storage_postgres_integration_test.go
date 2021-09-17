package account

import (
	"database/sql"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestPostgresStorage_integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration testing in short mode")
	}

	var database *sql.DB
	t.Run("connects to db", func(t *testing.T) {
		db, err := sql.Open("postgres", "user=admin password=admin dbname=account sslmode=disable")
		assert.NoError(t, err)

		assert.Eventually(t, func() bool {
			err := db.Ping()
			return err == nil
		}, time.Second*20, time.Millisecond*100)

		database = db
	})
	defer database.Close()

	storage := NewPostgresStorage(database, "account", "state", time.Second*5)

	addr := common.HexToAddress("0x1")
	account := Account{
		ID:             addr.String(),
		OwnershipClaim: OwnershipClaim{},
	}
	id, err := NewAccountID(ChainEthGoerli, addr)
	assert.NoError(t, err)

	claim, err := GenerateAccountOwnershipClaim(id, "1")
	assert.NoError(t, err)

	account.OwnershipClaim.Claim = claim
	account.OwnershipClaim.Signature = []byte("doesnt matter")

	accountState := AccountState{
		AccountID:          account.ID,
		ServiceID:          "test-service",
		GapHalvingPeriod:   86400,
		Power:              big.NewInt(10),
		LockedPower:        big.NewInt(2),
		EnergyCap:          big.NewInt(8),
		Energy:             big.NewInt(5),
		EnergyCalculatedAt: time.Now().Unix(),
	}
	t.Run("creates tables", func(t *testing.T) {
		err := storage.Init()
		assert.NoError(t, err)
	})

	t.Run("when the store is empty", func(t *testing.T) {
		t.Run("errors on account data request", func(t *testing.T) {
			_, err := storage.GetAccount(account.ID)
			assert.ErrorIs(t, err, ErrNotFound)
		})
		t.Run("errors on account delete request", func(t *testing.T) {
			err := storage.DeleteAccount(account.ID)
			assert.ErrorIs(t, err, ErrNotFound)
		})
		t.Run("errors upon account state deletion", func(t *testing.T) {
			err := storage.DeleteAccountState(accountState.ServiceID, account.ID)
			assert.ErrorIs(t, err, ErrNotFound)
		})
		t.Run("does not allow to initialize non-existent account state", func(t *testing.T) {
			as := accountState
			as.AccountID = common.HexToAddress("0x2").Hex()
			_, err := storage.InitAccountState(as)
			assert.ErrorIs(t, err, ErrNotFound)
		})
		t.Run("saves account", func(t *testing.T) {
			acc, err := storage.SaveAccount(account)
			assert.NoError(t, err)

			assert.EqualValues(t, account, acc)
		})
		t.Run("when account exists", func(t *testing.T) {
			t.Run("returns account data", func(t *testing.T) {
				acc, err := storage.GetAccount(account.ID)
				assert.NoError(t, err)

				assert.EqualValues(t, account, acc)
			})
			t.Run("saves account data with the same values", func(t *testing.T) {
				_, err := storage.SaveAccount(account)
				assert.NoError(t, err)

				acc, err := storage.GetAccount(account.ID)
				assert.NoError(t, err)

				assert.EqualValues(t, account, acc)
			})
			t.Run("updates account data", func(t *testing.T) {
				updatedAccount := account
				updatedAccount.OwnershipClaim.Signature = []byte("string")

				acc, err := storage.SaveAccount(updatedAccount)
				assert.NoError(t, err)

				assert.EqualValues(t, updatedAccount, acc)
			})
			t.Run("initializes account sate", func(t *testing.T) {
				res, err := storage.InitAccountState(accountState)
				assert.NoError(t, err)
				assert.EqualValues(t, accountState, res)

				res, err = storage.GetAccountState(accountState.ServiceID, accountState.AccountID)
				assert.NoError(t, err)
				assert.EqualValues(t, accountState, res)
			})
			t.Run("deletes the account", func(t *testing.T) {
				err := storage.DeleteAccount(account.ID)
				assert.NoError(t, err)

				_, err = storage.GetAccount(account.ID)
				assert.ErrorIs(t, err, ErrNotFound)

				err = storage.DeleteAccount(account.ID)
				assert.ErrorIs(t, err, ErrNotFound)
			})
			t.Run("when account state is initialized", func(t *testing.T) {
				t.Run("errors on multiple initializations", func(t *testing.T) {
					_, err := storage.SaveAccount(account)
					assert.NoError(t, err)
					_, err = storage.InitAccountState(accountState)
					assert.NoError(t, err)
					_, err = storage.InitAccountState(accountState)
					assert.ErrorIs(t, err, ErrAlreadyInitialized)
				})
				t.Run("returns account state", func(t *testing.T) {
					state, err := storage.GetAccountState(accountState.ServiceID, accountState.AccountID)
					assert.NoError(t, err)
					assert.EqualValues(t, accountState, state)
				})
				t.Run("update account state and returns correct response", func(t *testing.T) {
					newState := accountState
					newState.Power = big.NewInt(15)
					newState.Energy = big.NewInt(2)
					newState.EnergyCalculatedAt = time.Now().Unix()

					res, err := storage.ChangeAccountState(accountState, newState)
					assert.NoError(t, err)
					assert.EqualValues(t, newState, res)

					state, err := storage.GetAccountState(newState.ServiceID, newState.AccountID)
					assert.NoError(t, err)
					assert.EqualValues(t, newState, state)
				})
				t.Run("does not update state with incorrect previous state and returns correct response", func(t *testing.T) {
					incorrectState := accountState
					incorrectState.Power = big.NewInt(15)

					newState := accountState
					newState.Power = big.NewInt(20)

					_, err := storage.ChangeAccountState(incorrectState, newState)
					assert.ErrorIs(t, err, ErrNotFound)
				})
				t.Run("deletes account state", func(t *testing.T) {
					err := storage.DeleteAccountState(accountState.ServiceID, accountState.AccountID)
					assert.NoError(t, err)

					_, err = storage.GetAccountState(accountState.ServiceID, accountState.AccountID)
					assert.ErrorIs(t, err, ErrNotFound)

					err = storage.DeleteAccountState(accountState.ServiceID, accountState.AccountID)
					assert.ErrorIs(t, err, ErrNotFound)
				})

				t.Run("deletes account state when account is deleted", func(t *testing.T) {
					_, err := storage.InitAccountState(accountState)
					assert.NoError(t, err)

					_, err = storage.GetAccountState(accountState.ServiceID, accountState.AccountID)
					assert.NoError(t, err)

					err = storage.DeleteAccount(account.ID)
					assert.NoError(t, err)

					_, err = storage.GetAccountState(accountState.ServiceID, accountState.AccountID)
					assert.ErrorIs(t, err, ErrNotFound)

					err = storage.DeleteAccountState(accountState.ServiceID, accountState.AccountID)
					assert.ErrorIs(t, err, ErrNotFound)
				})
			})
		})
	})
}
