package keychain

import (
	"crypto/rand"
	"time"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
)

// Keychain Interface for managing device keypairs.
type Keychain interface {
	// CreateUUID makes new UUID value signed by the local node's private key
	CreateUUID() (*SignedUUID, error)

	// CreateMetadata makes message data shared between all node's p2p protocols
	CreateMetadata(peerID peer.ID) (*SignedMetadata, error)

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

// NewKeychain creates a new keychain with Config.
func NewKeychain(config *config.Config) (Keychain, error) {
	// Check if Keychain exists
	if keychainExists(config) {
		// Load Existing Keychain
		kc, err := loadKeychain(config)
		if err != nil {
			return nil, err
		}

		// Return Keychain
		return kc, nil
	} else {
		// Create Keychain
		kc, err := newKeychain(config)
		if err != nil {
			return nil, err
		}

		// Return Keychain
		return kc, nil
	}
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
		return nil, logger.Error("Failed to Read Account Key", err)
	}

	// Load Account Key to Keychain
	kc.LoadKeyPair(accPubKey, accPrivKey, Account)

	// Read Link Key
	linkPrivKey, linkPubKey, err := readKey(kcconfig, Link)
	if err != nil {
		return nil, logger.Error("Failed to Read Link Key", err)
	}

	// Load Link Key to Keychain
	kc.LoadKeyPair(linkPubKey, linkPrivKey, Link)

	// Read Group Key
	groupPrivKey, groupPubKey, err := readKey(kcconfig, Group)
	if err != nil {
		return nil, logger.Error("Failed to Read Group Key", err)
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
		return nil, logger.Error("Failed to generate Account KeyPair", err)
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
		return nil, logger.Error("Failed to generate Link KeyPair", err)
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
		return nil, logger.Error("Failed to generate Group KeyPair", err)
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

// CreateUUID makes a new UUID value signed by the local node's private key
func (kc *keychain) CreateUUID() (*SignedUUID, error) {
	// generate new UUID
	id := uuid.New().String()

	// sign UUID using local node's private key
	sig, err := kc.SignWith(Account, []byte(id))
	if err != nil {
		logger.Error("Failed to sign UUID", err)
		return nil, err
	}

	// Return UUID with signature
	return &SignedUUID{
		Value:     id,
		Signature: sig,
		Timestamp: time.Now().Unix(),
	}, nil
}

// CreateMetadata makes message data shared between all node's p2p protocols
func (kc *keychain) CreateMetadata(peerID peer.ID) (*SignedMetadata, error) {
	// Get local node's public key
	pubKey, err := kc.GetPubKey(Account)
	if err != nil {
		logger.Error("Failed to get local host's public key", err)
		return nil, err
	}

	// Marshal Public key into public key data
	nodePubKey, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Error("Failed to Extract Public Key", err)
		return nil, err
	}

	// Generate new Metadata
	return &SignedMetadata{
		Timestamp: time.Now().Unix(),
		PublicKey: nodePubKey,
		NodeId:    peer.Encode(peerID),
	}, nil
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
			return nil, nil, ErrInvalidKeyType
		}
	}
	return nil, nil, ErrKeychainUnready
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
		return nil, ErrInvalidKeyType
	}
	return nil, ErrKeychainUnready
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
		return nil, ErrInvalidKeyType
	}
	return nil, ErrKeychainUnready
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
		logger.Error("Failed to load Key Pair", ErrInvalidKeyType)
	}
}

// RemoveKeyPair removes a key from the keychain.
func (kc *keychain) RemoveKeyPair(kp KeyPairType) error {
	if kc.Exists(kp) {
		return kc.config.Delete(kp.Path())
	}
	return logger.Error("Failed to Remove Key Pair", ErrKeychainUnready)
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
	return nil, logger.Error("Failed to Sign Data", ErrKeychainUnready)
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
	return false, logger.Error("Failed to Verify Data", ErrKeychainUnready)
}
