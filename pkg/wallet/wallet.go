package wallet

import (
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-hq/sonr/x/identity/types"

	// "github.com/sonr-hq/sonr/pkg/node"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// WalletShare is a wallet that can be used to sign messages using ECDSA based on the MPC protocol.
type WalletShare interface {
	// Returns the Bech32 representation of the given party.
	Address() string

	// CMPConfig returns the *cmp.Config of this wallet if it exists.
	CMPConfig() *cmp.Config

	// DID returns the DID of this wallet.
	DID() (*types.DID, error)

	// Marshal serializes the cmp.Config into a byte slice for local storage
	Marshal() ([]byte, error)

	// PublicKey returns the public key of this wallet.
	PublicKey() (*secp256k1.PubKey, error)

	// SelfID returns the ID of this wallet.
	SelfID() party.ID

	// PartyIDs returns the IDs of all parties in the group.
	PartyIDs() []party.ID

	// Unmarshal deserializes the given byte slice into a cmp.Config
	Unmarshal([]byte) error

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}

// SaveToPath saves the wallet to the given path.
func SaveToPath(w WalletShare, path string) error {
	bz, err := w.Marshal()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, bz, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadFromPath loads a wallet from the given path.
func LoadFromPath(path string) (WalletShare, error) {
	config := cmp.Config{}
	bz, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalBinary(bz)
	if err != nil {
		return nil, err
	}
	w := EmptyWallet()
	err = w.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return w, nil
}
