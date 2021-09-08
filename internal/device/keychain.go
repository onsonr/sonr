package device

import (
	"errors"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/crypto"
)

// Directory Name of Private Key Folder
const PRIVATE_KEY_DIR = ".sonr_private"

type KeyPair int64

const (
	Account KeyPair = iota
	Link
	Group
)

type Keychain interface {
	// Checks if a key pair exists in the keychain.
	Exists(key KeyPair) bool

	// Gets a key pair from the keychain.
	GetKeyPair(key KeyPair) (crypto.PubKey, crypto.PrivKey, error)

	// Gets a public key from the keychain.
	GetPubKey(key KeyPair) (crypto.PubKey, error)

	// Gets a private key from the keychain.
	GetPrivKey(key KeyPair) (crypto.PrivKey, error)

	// Removes a key from the keychain.
	RemoveKeyPair(key KeyPair) error
}

// keychain is a keychain implementation that stores keys in a directory.
type keychain struct {
	Keychain
}

// Exists checks if a key pair exists in the keychain.
func (kc *keychain) Exists(kp KeyPair) bool {
	return Config.Exists(kp.Path())
}

// GetKeyPair gets a key pair from the keychain.
func (kc *keychain) GetKeyPair(key KeyPair) (crypto.PubKey, crypto.PrivKey, error) {
	if kc.Exists(key) {
		if key == Account {
			return kc.accountKeyPair()
		} else if key == Link {
			return kc.linkKeyPair()
		} else if key == Group {
			return kc.groupKeyPair()
		} else {
			return nil, nil, errors.New("Invalid Key Type")
		}
	}
	return nil, nil, errors.New("Keychain not loaded")
}

// GetPubKey gets a public key from the keychain.
func (kc *keychain) GetPubKey(key KeyPair) (crypto.PubKey, error) {
	if kc.Exists(key) {
		if key == Account {
			pub, _, err := kc.accountKeyPair()
			return pub, err
		} else if key == Group {
			pub, _, err := kc.groupKeyPair()
			return pub, err
		} else if key == Link {
			pub, _, err := kc.linkKeyPair()
			return pub, err
		} else {
			return nil, errors.New("Invalid Key Type")
		}
	}
	return nil, errors.New("Keychain not loaded")
}

// GetPrivKey gets a private key from the keychain.
func (kc *keychain) GetPrivKey(key KeyPair) (crypto.PrivKey, error) {
	if kc.Exists(key) {
		if key == Account {
			_, priv, err := kc.accountKeyPair()
			return priv, err
		} else if key == Group {
			_, priv, err := kc.groupKeyPair()
			return priv, err
		} else if key == Link {
			_, priv, err := kc.linkKeyPair()
			return priv, err
		} else {
			return nil, errors.New("Invalid Key Type")
		}
	}
	return nil, errors.New("Keychain not loaded")
}

// RemoveKeyPair removes a key from the keychain.
func (kc *keychain) RemoveKeyPair(key KeyPair) error {
	if kc.Exists(key) {
		return Config.Delete(key.Path())
	} else {
		return errors.New("Keychain not loaded")
	}
}

// ---------------- Retreiver Functions ----------------
// Get Pub/Priv Key Pair for Account
func (k *keychain) accountKeyPair() (crypto.PubKey, crypto.PrivKey, error) {
	// Get Buffer
	dat, err := Config.ReadFile(Account.Path())
	if err != nil {
		return nil, nil, err
	}

	// Get Private Key from Buffer
	privKey, err := crypto.UnmarshalPrivateKey(dat)
	if err != nil {
		return nil, nil, err
	}

	// Get Public Key from Private Key
	pubKey := privKey.GetPublic()
	return pubKey, privKey, nil
}

// Get Pub/Priv Key Pair for Group
func (k *keychain) groupKeyPair() (crypto.PubKey, crypto.PrivKey, error) {
	// Get Buffer
	dat, err := Config.ReadFile(Group.Path())
	if err != nil {
		return nil, nil, err
	}

	// Get Private Key from Buffer
	privKey, err := crypto.UnmarshalPrivateKey(dat)
	if err != nil {
		return nil, nil, err
	}

	// Get Public Key from Private Key
	pubKey := privKey.GetPublic()
	return pubKey, privKey, nil
}

// Get Pub/Priv Key Pair for Link
func (k *keychain) linkKeyPair() (crypto.PubKey, crypto.PrivKey, error) {
	// Get Buffer
	dat, err := Config.ReadFile(Link.Path())
	if err != nil {
		return nil, nil, err
	}

	// Get Private Key from Buffer
	privKey, err := crypto.UnmarshalPrivateKey(dat)
	if err != nil {
		return nil, nil, err
	}

	// Get Public Key from Private Key
	pubKey := privKey.GetPublic()
	return pubKey, privKey, nil
}

// ---------------- FilePath Functions ----------------
// Get Account Private Key File Path
func (kp KeyPair) Path() string {
	switch kp {
	case Account:
		return filepath.Join("keychain", "account_private_key")
	case Group:
		return filepath.Join("keychain", "group_private_key")
	case Link:
		return filepath.Join("keychain", "link_private_key")
	}
	return ""
}
