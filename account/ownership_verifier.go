package account

type OwnershipVerifyer struct{}

func (ov *OwnershipVerifyer) GenerateAndVerify(acc AccountID, signature []byte, version string) error {
	data, err := GenerateAccountOwnershipClaim(acc, version)
	if err != nil {
		return err
	}

	return VerifyClaimSignature(data, signature, acc.Address())
}
