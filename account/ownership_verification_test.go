package account

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/stretchr/testify/assert"
)

func TestOwnershipVerification(t *testing.T) {
	t.Run("generates and verifies signature", func(t *testing.T) {
		dir := t.TempDir()
		ks := keystore.NewKeyStore(dir, keystore.LightScryptN, keystore.LightScryptP)
		acc, err := ks.NewAccount("")
		assert.NoError(t, err)

		err = ks.Unlock(acc, "")
		assert.NoError(t, err)

		accID, err := NewAccountID(ChainMaticMainnet, acc.Address)
		assert.NoError(t, err)

		claim, err := GenerateAccountOwnershipClaim(accID, "1")
		assert.NoError(t, err)

		signature, err := SignClaim(acc.Address, claim, ks)
		assert.NoError(t, err)

		err = VerifyClaimSignature(claim, signature, acc.Address)
		assert.NoError(t, err)
	})
}
