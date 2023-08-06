package crypto

import (
	"crypto/rand"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"

	"github.com/sonrhq/core/types/crypto"
)

// PartyID is a type alias for party.ID in pkg/party.
type PartyID = party.ID

// MPCPool is a type alias for pool.Pool in pkg/pool.
type MPCPool = pool.Pool

// MPCCmpConfig is a type alias for pool.CmpConfig in pkg/pool.
type MPCCmpConfig = cmp.Config

// MPCECDSASignature is a type alias for ecdsa.Signature in pkg/ecdsa.
type MPCECDSASignature = ecdsa.Signature

// MPCSecp256k1Curve is a type alias for curve.Secp256k1Point in pkg/math/curve.
type MPCSecp256k1Curve = curve.Secp256k1

// MPCSecp256k1Point is a type alias for curve.Secp256k1Point in pkg/math/curve.
type MPCSecp256k1Point = curve.Secp256k1Point

// NewMPCPool creates a new MPCPool with the given size.
func NewMPCPool(size int) *MPCPool {
	return pool.NewPool(0)
}

// NewEmptyECDSASecp256k1Signature creates a new empty MPCECDSASignature.
func NewEmptyECDSASecp256k1Signature() MPCECDSASignature {
	return ecdsa.EmptySignature(MPCSecp256k1Curve{})
}

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
