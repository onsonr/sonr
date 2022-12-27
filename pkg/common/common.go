package common

import (
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	tm_crypto "github.com/tendermint/tendermint/crypto"
"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	tm_json "github.com/tendermint/tendermint/libs/json"
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

// > Loads a private key from a JSON file and returns a `crypto.PrivKey` interface
func LoadPrivKeyFromJsonPath(path string) (crypto.PrivKey, error) {
	// Load the key from the given path.
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Create new private key interface
	var vnPk tm_crypto.PrivKey

	// Unmarshal the key into the interface.
	err = tm_json.Unmarshal(key, &vnPk)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.UnmarshalPrivateKey(vnPk.Bytes())
	if err != nil {
		return nil, err
	}
	return priv, nil
}
