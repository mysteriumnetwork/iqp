package account

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
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
	t.Run("is compatibile with js sdk signature", func(t *testing.T) {
		signature := "a487cdd056e902dc4e1a00e34f746a024196ed5b09306e00301003eb475ac7f57f84d22b7427020667af41598ce44339f152f20d3b31bedcf14ee9a4b3a27da61c"
		address := common.HexToAddress("0xD4E8BC65f04B9f23bA8683776D36C12773D56500")
		accID, err := NewAccountID(ChainEthMainnet, address)
		assert.NoError(t, err)

		claim, err := GenerateAccountOwnershipClaim(accID, "1")
		assert.NoError(t, err)

		res, _ := json.Marshal(claim)
		fmt.Println(string(res))

		decoded, err := hex.DecodeString(signature)
		assert.NoError(t, err)

		err = VerifyClaimSignature(claim, decoded, address)
		assert.NoError(t, err)
	})
}
