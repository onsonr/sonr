// It converts a `WebauthnCredential` to a `webauthn.Credential`
package crypto

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)



// A Network is a channel that sends messages to parties and receives messages from parties.
type Network interface {
	// Ls returns a list of peers that are connected to the network.
	Ls() []party.ID

	// A function that takes in a party ID and returns a channel of protocol messages.
	Next(id party.ID) <-chan *protocol.Message

	// Sending a message to the network.
	Send(msg *protocol.Message)

	// A channel that is closed when the party is done with the protocol.
	Done(id party.ID) chan struct{}

	// A function that is called when a party is done with the protocol.
	Quit(id party.ID)

	// IsOnlineNetwork returns true if the network is an online network.
	IsOnlineNetwork() bool
}

// Wallet is a collection of WalletShares that can be used with a Network in order to utilize the
// multi-party signature protocol.
type Wallet interface {
	// Address returns the address of the Wallet.
	Address() string

	// Bip32Derive creates a new WalletShare that is derived from the given path.
	Bip32Derive(i uint32, prefix string) (WalletShare, error)

	// EncryptKey returns the secret key from Storage
	EncryptKey() ([]byte, error)

	// Find returns the WalletShare with the given ID.
	Find(id party.ID) WalletShare

	// GetConfigMap returns a map of party.ID to cmp.Config.
	GetConfigMap() map[party.ID]*cmp.Config

	// Network returns the Network that this Wallet is associated with.
	Network() Network

	// PublicKey returns the public key of this wallet.
	PublicKey() (*secp256k1.PubKey, error)

	// Refresh the WalletShares.
	Refresh(current party.ID) (Wallet, error)

	// Sign a message with the given wallet.
	Sign(current party.ID, m []byte) ([]byte, error)

	// SignTx signs a transaction with the given wallet.
	SignTx(current party.ID, msgs ...sdk.Msg) ([]byte, error)

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}

// WalletShare is a wallet that can be used to sign messages using ECDSA based on the MPC protocol.
type WalletShare interface {
	// Returns the Bech32 representation of the given party.
	Address() string

	// Bip32Derive creates a new WalletShare that is derived from the given path.
	Bip32Derive(i uint32, prefix string) (WalletShare, error)

	// CMPConfig returns the *cmp.Config of this wallet if it exists.
	CMPConfig() *cmp.Config

	// DID returns the DID of this wallet.
	DID() (string, error)

	// Index returns the index of this wallet. If its the master wallet, it returns -1.
	Index() int

	// Marshal serializes the cmp.Config into a byte slice for local storage
	Marshal() ([]byte, error)

	// Prefix returns the prefix of this wallet.
	Prefix() string

	// PublicKey returns the public key of this wallet.
	PublicKey() (*secp256k1.PubKey, error)

	// SelfID returns the ID of this wallet.
	SelfID() party.ID

	// PartyIDs returns the IDs of all parties in the group.
	PartyIDs() []party.ID

	// Share returns the share of this wallet.
	Share() *common.WalletShareConfig

	// Unmarshal deserializes the given byte slice into a cmp.Config
	Unmarshal([]byte) error

	// Verify a signature with the given wallet.
	Verify(msg, sig []byte) bool
}
