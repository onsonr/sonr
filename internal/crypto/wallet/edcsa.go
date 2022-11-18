package wallet

import (
	"github.com/sonr-io/multi-party-sig/pkg/ecdsa"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// EDCSAWallet is a wallet that can be used to sign messages using ECDSA based on the MPC protocol.
type EDCSAWallet interface {
	// Returns the Bech32 representation of the given party.
	Address(id ...party.ID) (string, error)

	// Config returns the configuration of this wallet.
	Config() *cmp.Config

	// PublicKey returns the public key of this wallet.
	PublicKey() ([]byte, error)

	// PublicKeyProto returns the public key of this wallet.
	PublicKeyProto() (*secp256k1.PubKey, error)

	// Sign a message with the given wallet.
	Sign(m []byte) (*ecdsa.Signature, error)

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}
