package crypto

import (
	"crypto/rand"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"

	"github.com/sonrhq/core/types/crypto"
)

const (
	// ChallengeLength - Length of bytes to generate for a challenge
	ChallengeLength = 32
)

// GenerateChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func GenerateChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}
	return challenge, nil
}

// Algo is the type of algorithm used for key generation.
type Algo string

const (
	// AlgoSecp256k1 is the secp256k1 algorithm.
	AlgoSecp256k1 Algo = "secp256k1"

	// AlgoEd25519 is the ed25519 algorithm.
	AlgoEd25519 Algo = "ed25519"

	// AlgoSr25519 is the sr25519 algorithm.
	AlgoSr25519 Algo = "sr25519"
)

// KeyType returns the KeyType of the algorithm.
func (a Algo) KeyType() KeyType {
	switch a {
	case AlgoSecp256k1:
		return crypto.KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019
	case AlgoEd25519:
		return crypto.KeyType_KeyType_ED25519_VERIFICATION_KEY_2018
	case AlgoSr25519:
		return crypto.KeyType_KeyType_JSON_WEB_KEY_2020
	default:
		return crypto.KeyType_KeyType_JSON_WEB_KEY_2020
	}
}

// AccountData is the data that is returned by the keygen process. It contains the address, the algorithm used to generate
// the key pair and the public key.
type AccountData = crypto.AccountData

// NewDefaultAccountData creates a new AccountData with the default values.
func NewDefaultAccountData(cointype CoinType, publicKey *PubKey) (*AccountData, error) {
	if publicKey == nil {
		return nil, fmt.Errorf("public key is nil")
	}
	return &AccountData{
		Address:   cointype.FormatAddress(publicKey),
		Algo:      "secp256k1",
		PublicKey: publicKey.Bytes(),
	}, nil
}
