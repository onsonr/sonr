package wallet

import (
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// MPCWallet is a wallet that can be used to sign messages using ECDSA based on the MPC protocol.
type MPCWallet interface {
	crypto.PrivKey
	// Returns the Bech32 representation of the given party.
	Address() string

	// Marshal serializes the cmp.Config into a byte slice for local storage
	Marshal() ([]byte, error)

	// PublicKey returns the public key of this wallet.
	PublicKey() (*secp256k1.PubKey, error)

	// GetPubKey returns the public key of this wallet.
	GetPublic() crypto.PubKey

	// Sign a message with the given wallet.
	Sign(m []byte) ([]byte, error)

	// Unmarshal deserializes the given byte slice into a cmp.Config
	Unmarshal([]byte) error

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}
