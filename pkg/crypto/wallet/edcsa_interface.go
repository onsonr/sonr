package wallet

import (
	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// MPCWallet is a wallet that can be used to sign messages using ECDSA based on the MPC protocol.
type MPCWallet interface {
	// Returns the Bech32 representation of the given party.
	Address() string

	// Marshal serializes the cmp.Config into a byte slice for local storage
	Marshal() ([]byte, error)

	// PublicKey returns the public key of this wallet.
	PublicKey() (*secp256k1.PubKey, error)

	// Sign a message with the given wallet.
	Sign(m []byte) (*ecdsa.Signature, error)

	// Unmarshal deserializes the given byte slice into a cmp.Config
	Unmarshal([]byte) error

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}
