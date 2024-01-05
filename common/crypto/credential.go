package crypto

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
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
		return KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019
	case AlgoEd25519:
		return KeyType_KeyType_ED25519_VERIFICATION_KEY_2018
	case AlgoSr25519:
		return KeyType_KeyType_JSON_WEB_KEY_2020
	default:
		return KeyType_KeyType_JSON_WEB_KEY_2020
	}
}

// Base64UrlToBytes converts a base64url string to bytes.
func Base64UrlToBytes(base64Url string) ([]byte, error) {
	base64String := strings.ReplaceAll(strings.ReplaceAll(base64Url, "-", "+"), "_", "/")
	missingPadding := len(base64String) % 4
	if missingPadding > 0 {
		base64String += strings.Repeat("=", 4-missingPadding)
	}
	return base64.StdEncoding.DecodeString(base64String)
}

// ParseCredentialPublicKey parses a public key from a base64url string.
func ParseCredentialPublicKey(pubStr string) (interface{}, error) {
	derEncodedPublicKey, err := Base64UrlToBytes(pubStr)
	if err != nil {
		return nil, err
	}

	publicKey, err := x509.ParsePKIXPublicKey(derEncodedPublicKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
