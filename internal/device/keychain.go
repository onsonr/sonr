package device

import (
	"errors"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
)

// KeyPair is a type of keypair
type KeyPair int64

const (
	// Account is the keypair for the account
	Account KeyPair = iota

	// Link is the keypair for linking Devices
	Link

	// Group is the keypair for created Groups
	Group

	// Directory Name of Private Key Folder
	PRIVATE_KEY_DIR = ".sonr_private"
)

// Keychain Interface for managing device keypairs.
type Keychain interface {
	// Exists Checks if a key pair exists in the keychain.
	Exists(kp KeyPair) bool

	// GetKeyPair Gets a key pair from the keychain.
	GetKeyPair(kp KeyPair) (crypto.PubKey, crypto.PrivKey, error)

	// GetPubKey Gets a public key from the keychain.
	GetPubKey(kp KeyPair) (crypto.PubKey, error)

	// GetPrivKey Gets a private key from the keychain.
	GetPrivKey(kp KeyPair) (crypto.PrivKey, error)

	// RemoveKeyPair Removes a key from the keychain.
	RemoveKeyPair(kp KeyPair) error

	// SignWith returns a signature for a message with specified pair
	SignWith(kp KeyPair, msg []byte) ([]byte, error)

	// VerifyWith verifies a signature with specified pair
	VerifyWith(kp KeyPair, msg []byte, sig []byte) (bool, error)
}

// keychain is a keychain implementation that stores keys in a directory.
type keychain struct {
	Keychain
	config *config.Config
}

// Exists checks if a key pair exists in the keychain.
func (kc *keychain) Exists(kp KeyPair) bool {
	return kc.config.Exists(kp.Path())
}

// GetKeyPair gets a key pair from the keychain.
func (kc *keychain) GetKeyPair(kp KeyPair) (crypto.PubKey, crypto.PrivKey, error) {
	if kc.Exists(kp) {
		if kp == Account {
			return kc.accountKeyPair()
		} else if kp == Link {
			return kc.linkKeyPair()
		} else if kp == Group {
			return kc.groupKeyPair()
		} else {
			return nil, nil, errors.New("Invalid Key Type")
		}
	}
	return nil, nil, errors.New("Keychain not loaded")
}

// GetPubKey gets a public key from the keychain.
func (kc *keychain) GetPubKey(kp KeyPair) (crypto.PubKey, error) {
	if kc.Exists(kp) {
		if kp == Account {
			pub, _, err := kc.accountKeyPair()
			return pub, err
		} else if kp == Group {
			pub, _, err := kc.groupKeyPair()
			return pub, err
		} else if kp == Link {
			pub, _, err := kc.linkKeyPair()
			return pub, err
		} else {
			return nil, errors.New("Invalid Key Type")
		}
	}
	return nil, errors.New("Keychain not loaded")
}

// GetPrivKey gets a private key from the keychain.
func (kc *keychain) GetPrivKey(kp KeyPair) (crypto.PrivKey, error) {
	if kc.Exists(kp) {
		if kp == Account {
			_, priv, err := kc.accountKeyPair()
			return priv, err
		} else if kp == Group {
			_, priv, err := kc.groupKeyPair()
			return priv, err
		} else if kp == Link {
			_, priv, err := kc.linkKeyPair()
			return priv, err
		} else {
			return nil, errors.New("Invalid Key Type")
		}
	}
	return nil, errors.New("Keychain not loaded")
}

// RemoveKeyPair removes a key from the keychain.
func (kc *keychain) RemoveKeyPair(kp KeyPair) error {
	if kc.Exists(kp) {
		return kc.config.Delete(kp.Path())
	}
	return errors.New("Keychain not loaded")
}

// SignWith signs a message with the specified keypair
func (kc *keychain) SignWith(kp KeyPair, msg []byte) ([]byte, error) {
	if kc.Exists(kp) {
		priv, err := kc.GetPrivKey(kp)
		if err != nil {
			return nil, err
		}
		return priv.Sign(msg)
	}
	return nil, errors.New("Keychain not loaded")
}

// VerifyWith verifies a signature with specified pair
func (kc *keychain) VerifyWith(kp KeyPair, msg []byte, sig []byte) (bool, error) {
	if kc.Exists(kp) {
		pub, err := kc.GetPubKey(kp)
		if err != nil {
			return false, err
		}
		return pub.Verify(msg, sig)
	}
	return false, errors.New("Keychain not loaded")
}

// ---------------- Retreiver Functions ----------------
// Get Pub/Priv Key Pair for Account
func (k *keychain) accountKeyPair() (crypto.PubKey, crypto.PrivKey, error) {
	// Get Buffer
	dat, err := k.config.ReadFile(Account.Path())
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
	dat, err := k.config.ReadFile(Group.Path())
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
	dat, err := k.config.ReadFile(Link.Path())
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
