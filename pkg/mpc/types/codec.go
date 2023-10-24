package types

import (
	"errors"
	"math/big"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/sonr-io/kryptology/pkg/core/curves"
	pv1 "github.com/sonr-io/kryptology/pkg/core/protocol"
)

type Secp256k1PublicKey = *secp256k1.PubKey
type PubKey = cryptotypes.PubKey
type KeyshareRole string

const (
	// KeyshareRolePublic is the default role for the alice dkg
	KeyshareRolePublic KeyshareRole = "alice"

	// KeyshareRoleUser is the role for an encrypted keyshare for a user
	KeyshareRoleUser KeyshareRole = "bob"
)

func (ksr KeyshareRole) isAlice() bool {
	return ksr == KeyshareRolePublic
}

func (ksr KeyshareRole) isBob() bool {
	return ksr == KeyshareRoleUser
}

func SerializeECDSASecp256k1Signature(sig *curves.EcdsaSignature) ([]byte, error) {
	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()

	sigBytes := make([]byte, 66) // V (1 byte) + R (32 bytes) + S (32 bytes)
	sigBytes[0] = byte(sig.V)
	copy(sigBytes[33-len(rBytes):33], rBytes)
	copy(sigBytes[66-len(sBytes):66], sBytes)
	return sigBytes, nil
}

func DeserializeECDSASecp256k1Signature(sigBytes []byte) (*curves.EcdsaSignature, error) {
	if len(sigBytes) != 66 {
		return nil, errors.New("malformed signature: not the correct size")
	}
	sig := &curves.EcdsaSignature{
		V: int(sigBytes[0]),
		R: new(big.Int).SetBytes(sigBytes[1:33]),
		S: new(big.Int).SetBytes(sigBytes[33:66]),
	}
	return sig, nil
}

// For DKG bob starts first. For refresh and sign, Alice starts first.
func RunIteratedProtocol(firstParty pv1.Iterator, secondParty pv1.Iterator) (error, error) {
	var (
		message *pv1.Message
		aErr    error
		bErr    error
	)
	for aErr != pv1.ErrProtocolFinished || bErr != pv1.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != pv1.ErrProtocolFinished {
			return nil, bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != pv1.ErrProtocolFinished {
			return aErr, nil
		}
	}
	return aErr, bErr
}
