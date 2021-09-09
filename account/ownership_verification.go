package account

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
)

func GenerateAccountOwnershipClaim(accountID AccountID, version string) (core.TypedData, error) {
	if version != "1" {
		return core.TypedData{}, errors.New("unsupported claim version")
	}
	types := core.Types{
		"EIP712Domain": []core.Type{
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "version",
				Type: "string",
			},
			{
				Name: "chainId",
				Type: "uint256",
			},
		},
		"AccountOwnershipClaim": []core.Type{
			{
				Name: "claim",
				Type: "string",
			},
			{
				Name: "address",
				Type: "string",
			},
		},
	}

	domain := core.TypedDataDomain{
		Name:    "IQ Protocol",
		Version: version,
		ChainId: (*math.HexOrDecimal256)(big.NewInt(accountID.ChainID())),
	}

	message := core.TypedDataMessage{
		"claim":   "Hereby I confirm to be the owner of the following address",
		"address": strings.ToLower(accountID.Address().Hex()),
	}

	data := core.TypedData{
		Types:       types,
		Message:     message,
		Domain:      domain,
		PrimaryType: "AccountOwnershipClaim",
	}
	return data, nil
}

type HashSigner interface {
	SignHash(a accounts.Account, hash []byte) ([]byte, error)
}

func SignClaim(addr common.Address, claim core.TypedData, hashSigner HashSigner) ([]byte, error) {
	sighash, err := TypedDataToHash(claim)
	if err != nil {
		return nil, err
	}
	signature, err := hashSigner.SignHash(accounts.Account{
		Address: addr,
	}, sighash)
	if err != nil {
		return nil, err
	}

	var v = 27 + (uint64(signature[len(signature)-1]) % 2)
	buffer := make([]byte, 1)
	_ = binary.PutUvarint(buffer, v)
	signature[64] = buffer[0]
	return signature, nil
}

func TypedDataToHash(claim core.TypedData) ([]byte, error) {
	domainSeparator, err := claim.HashStruct("EIP712Domain", claim.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("has struct EIP712Domain %w", err)
	}
	typedDataHash, err := claim.HashStruct(claim.PrimaryType, claim.Message)
	if err != nil {
		return nil, fmt.Errorf("has struct PrimaryType %w", err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return crypto.Keccak256(rawData), nil
}

func VerifyClaimSignature(typedData core.TypedData, signature []byte, addr common.Address) error {
	if len(signature) != 65 {
		return fmt.Errorf("invalid signature length: %d", len(signature))
	}
	copiedSignature := make([]byte, len(signature))
	copy(copiedSignature, signature)

	if copiedSignature[64] != 27 && copiedSignature[64] != 28 {
		return fmt.Errorf("invalid recovery id: %d", copiedSignature[64])
	}
	copiedSignature[64] -= 27

	hash, err := TypedDataToHash(typedData)
	if err != nil {
		return err
	}

	pubKeyRaw, err := crypto.Ecrecover(hash, copiedSignature)
	if err != nil {
		return fmt.Errorf("invalid signature: %s", err.Error())
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !bytes.Equal(addr.Bytes(), recoveredAddr.Bytes()) {
		return fmt.Errorf("addresses do not match. Expected %v, got %v", addr.Hex(), recoveredAddr.Hex())
	}

	return nil
}
