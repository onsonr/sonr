package device

import (
	"crypto/rand"
	"errors"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
)

// KeyPairType is a type of keypair
type KeyPairType int64

const (
	// Account is the keypair for the account
	Account KeyPairType = iota

	// Link is the keypair for linking Devices
	Link

	// Group is the keypair for created Groups
	Group

	// Directory Name of Private Key Folder
	PRIVATE_KEY_DIR = ".sonr_private"
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

// Keychain Interface for managing device keypairs.
type Keychain interface {
	// Exists Checks if a key pair exists in the keychain.
	Exists(kp KeyPairType) bool

	// GetKeyPair Gets a key pair from the keychain.
	GetKeyPair(kp KeyPairType) (crypto.PubKey, crypto.PrivKey, error)

	// GetPubKey Gets a public key from the keychain.
	GetPubKey(kp KeyPairType) (crypto.PubKey, error)

	// GetPrivKey Gets a private key from the keychain.
	GetPrivKey(kp KeyPairType) (crypto.PrivKey, error)

	// RemoveKeyPair Removes a key from the keychain.
	RemoveKeyPair(kp KeyPairType) error

	// SignWith returns a signature for a message with specified pair
	SignWith(kp KeyPairType, msg []byte) ([]byte, error)

	// VerifyWith verifies a signature with specified pair
	VerifyWith(kp KeyPairType, msg []byte, sig []byte) (bool, error)
}

// keychain is a keychain implementation that stores keys in a directory.
type keychain struct {
	Keychain
	config *config.Config

	// Key Pair References
	accountKeyPair keyPair
	groupKeyPair   keyPair
	linkKeyPair    keyPair
}

// loadKeychain loads a keychain from a file.
func loadKeychain(kcconfig *config.Config) (Keychain, error) {
	// Create Keychain
	kc := &keychain{
		config: kcconfig,
	}

	// Read Account Key
	accPrivKey, accPubKey, err := readKey(kcconfig, Account)
	if err != nil {
		return nil, err
	}

	// Load Account Key to Keychain
	kc.LoadKeyPair(accPubKey, accPrivKey, Account)

	// Read Link Key
	linkPrivKey, linkPubKey, err := readKey(kcconfig, Link)
	if err != nil {
		return nil, err
	}

	// Load Link Key to Keychain
	kc.LoadKeyPair(linkPubKey, linkPrivKey, Link)

	// Read Group Key
	groupPrivKey, groupPubKey, err := readKey(kcconfig, Group)
	if err != nil {
		return nil, err
	}

	// Load Group Key to Keychain
	kc.LoadKeyPair(groupPubKey, groupPrivKey, Group)

	return kc, nil
}

// newKeychain creates a new keychain.
func newKeychain(kcconfig *config.Config) (Keychain, error) {
	// Create Keychain
	kc := &keychain{
		config: kcconfig,
	}

	// Create New Account Key
	accPrivKey, accPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Write Account Key to Disk
	err = writeKey(kcconfig, accPrivKey, Account)
	if err != nil {
		return nil, err
	}

	// Load Account Key to Keychain
	kc.LoadKeyPair(accPubKey, accPrivKey, Account)

	// Create New Link Key
	linkPrivKey, linkPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Write Link Key to Disk
	err = writeKey(kcconfig, linkPrivKey, Link)
	if err != nil {
		return nil, err
	}

	// Load Link Key to Keychain
	kc.LoadKeyPair(linkPubKey, linkPrivKey, Link)

	// Create New Group Key
	groupPrivKey, groupPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Write Group Key to Disk
	err = writeKey(kcconfig, groupPrivKey, Group)
	if err != nil {
		return nil, err
	}

	// Load Group Key to Keychain
	kc.LoadKeyPair(groupPubKey, groupPrivKey, Group)
	return kc, nil
}

// Exists checks if a key pair exists in the keychain.
func (kc *keychain) Exists(kp KeyPairType) bool {
	return kc.config.Exists(kp.Path())
}

// GetKeyPair gets a key pair from the keychain.
func (kc *keychain) GetKeyPair(kp KeyPairType) (crypto.PubKey, crypto.PrivKey, error) {
	if kc.Exists(kp) {
		if kp == Account {
			return kc.accountKeyPair.PrivPubKeys()
		} else if kp == Link {
			return kc.linkKeyPair.PrivPubKeys()
		} else if kp == Group {
			return kc.groupKeyPair.PrivPubKeys()
		} else {
			return nil, nil, errors.New("Invalid Key Type")
		}
	}
	return nil, nil, errors.New("Keychain not loaded")
}

// GetPubKey gets a public key from the keychain.
func (kc *keychain) GetPubKey(kp KeyPairType) (crypto.PubKey, error) {
	if kc.Exists(kp) {
		if kp == Account {
			pub, _, err := kc.accountKeyPair.PrivPubKeys()
			if err != nil {
				return nil, err
			}
			return pub, nil
		} else if kp == Group {
			pub, _, err := kc.groupKeyPair.PrivPubKeys()
			if err != nil {
				return nil, err
			}
			return pub, nil
		} else if kp == Link {
			pub, _, err := kc.linkKeyPair.PrivPubKeys()
			if err != nil {
				return nil, err
			}
			return pub, nil
		}
		return nil, errors.New("Invalid Key Type")
	}
	return nil, errors.New("Keychain not loaded")
}

// GetPrivKey gets a private key from the keychain.
func (kc *keychain) GetPrivKey(kp KeyPairType) (crypto.PrivKey, error) {
	if kc.Exists(kp) {
		if kp == Account {
			_, priv, err := kc.accountKeyPair.PrivPubKeys()
			if err != nil {
				return nil, err
			}
			return priv, nil
		} else if kp == Group {
			_, priv, err := kc.groupKeyPair.PrivPubKeys()
			if err != nil {
				return nil, err
			}
			return priv, nil
		} else if kp == Link {
			_, priv, err := kc.linkKeyPair.PrivPubKeys()
			if err != nil {
				return nil, err
			}
			return priv, nil
		}
		return nil, errors.New("Invalid Key Type")
	}
	return nil, errors.New("Keychain not loaded")
}

// LoadKeyPair loads a keypair set into the keychain.
func (kc *keychain) LoadKeyPair(pub crypto.PubKey, priv crypto.PrivKey, kp KeyPairType) {
	if kp == Account {
		kc.accountKeyPair = keyPair{pub, priv, kp}
	} else if kp == Link {
		kc.linkKeyPair = keyPair{pub, priv, kp}
	} else if kp == Group {
		kc.groupKeyPair = keyPair{pub, priv, kp}
	} else {
		logger.Error("Invalid KeyPair Type provided")
	}

}

// RemoveKeyPair removes a key from the keychain.
func (kc *keychain) RemoveKeyPair(kp KeyPairType) error {
	if kc.Exists(kp) {
		return kc.config.Delete(kp.Path())
	}
	return errors.New("Keychain not loaded")
}

// SignWith signs a message with the specified keypair
func (kc *keychain) SignWith(kp KeyPairType, msg []byte) ([]byte, error) {
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
func (kc *keychain) VerifyWith(kp KeyPairType, msg []byte, sig []byte) (bool, error) {
	if kc.Exists(kp) {
		pub, err := kc.GetPubKey(kp)
		if err != nil {
			return false, err
		}
		return pub.Verify(msg, sig)
	}
	return false, errors.New("Keychain not loaded")
}

// ---------------- FilePath Functions ----------------
type keyPair struct {
	pub    crypto.PubKey
	priv   crypto.PrivKey
	kpType KeyPairType
}

// PrivPubKeys returns the private and public keys for the keypair given keychain
func (kp keyPair) PrivPubKeys() (crypto.PubKey, crypto.PrivKey, error) {
	if kp.priv == nil {
		return nil, nil, errors.New("No Private Key")
	}

	if kp.pub == nil {
		return nil, nil, errors.New("No Public Key")
	}
	return kp.pub, kp.priv, nil
}
