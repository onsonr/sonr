package wallet

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-hq/sonr/internal/node"
	"github.com/sonr-io/multi-party-sig/pkg/party"
)

// WalletShare is a wallet that can be used to sign messages using ECDSA based on the MPC protocol.
type WalletShare interface {
	// Returns the Bech32 representation of the given party.
	Address() string

	// Marshal serializes the cmp.Config into a byte slice for local storage
	Marshal() ([]byte, error)

	// PublicKey returns the public key of this wallet.
	PublicKey() (*secp256k1.PubKey, error)

	// SelfID returns the ID of this wallet.
	SelfID() party.ID

	// GroupIDs returns the IDs of all parties in the group.
	GroupIDs() []party.ID

	// Sign begins a round of the MPC protocol to sign the given message.
	Sign(msg []byte, th node.TopicHandler) ([]byte, error)

	// Unmarshal deserializes the given byte slice into a cmp.Config
	Unmarshal([]byte) error

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}
