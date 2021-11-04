package wallet

import (
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type PrivKey interface {
	crypto.PrivKey
	Marshal() ([]byte, error)
	SignHmac(msg string) (string, error)
	VerifyHmac(msg string, sig string) (bool, error)
}

type PubKey interface {
	crypto.PubKey
	Marshal() ([]byte, error)
	PeerID() peer.ID
}

// KeyPairType is a type of keypair
type KeyPairType int64

const (
	// Account is the keypair for the account
	Account KeyPairType = iota

	// Link is the keypair for linking Devices
	Link

	// Group is the keypair for created Groups
	Group
)

// Path returns the path to the keypair
func (kpt KeyPairType) Path() string {
	switch kpt {
	case Account:
		return filepath.Join("keychain", "account_private_key")
	case Group:
		return filepath.Join("keychain", "group_private_key")
	case Link:
		return filepath.Join("keychain", "link_private_key")
	}
	return ""
}

type SignedFingerprint struct {
	Mnemonic    string
	SName       string
	Prefix      string
	Fingerprint []byte
	PublicKey   string
	DeviceID    string
}
