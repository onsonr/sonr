package wallet

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/cosmos/go-bip39"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/pkg/device"
)

var (
	// Sonr is the KeyChain for the Sonr Protocol
	Sonr KeyChain
)

// SignedMetadata is a struct to be used for signing metadata.
type SignedMetadata struct {
	Timestamp int64
	PublicKey []byte
	NodeId    string
}

// SignedUUID is a struct to be converted into a UUID.
type SignedUUID struct {
	Timestamp int64
	Signature []byte
	Value     string
}

// Keychain Interface for managing device keypairs.
type KeyChain interface {
	// CreateUUID makes new UUID value signed by the local node's private key
	CreateUUID() (*SignedUUID, error)

	// CreateFingerprint makes a new fingerprint value signed by the local node's private key
	CreateFingerprint(sName string, deviceID string) (*SignedFingerprint, error)

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

	// GetSnrPubKey Gets a public key from the keychain with Snr wrapper.
	GetSnrPubKey(kp KeyPairType) (*SnrPubKey, error)

	// GetSnrPrivKey Gets a private key from the keychain with Snr wrapper.
	GetSnrPrivKey(kp KeyPairType) (*SnrPrivKey, error)

	// RemoveKeyPair Removes a key from the keychain.
	RemoveKeyPair(kp KeyPairType) error

	// SignWith returns a signature for a message with specified pair
	SignWith(kp KeyPairType, msg []byte) ([]byte, error)

	// SignHmacWith returns a signature for a message with specified pair using HMAC(256)
	// returning: signature, error
	SignHmacWith(kp KeyPairType, msg string) (string, error)

	// ValidateFingerprint validates a fingerprint value
	ValidateFingerprint(fp *SignedFingerprint) (bool, error)

	// VerifyWith verifies a signature with specified pair
	VerifyWith(kp KeyPairType, msg []byte, sig []byte) (bool, error)

	// VerifyHmacWith verifies a signature with specified pair using HMAC(256)
	// returning: true/false, error
	VerifyHmacWith(kp KeyPairType, msg string, sig string) (bool, error)
}

// Open creates a new keychain with Wallet Folder.
func Open() error {
	config := device.Wallet

	// Check if Keychain exists
	if keychainExists(config) {
		// Load Existing Keychain
		kc, err := loadKeychain(config)
		if err != nil {
			return err
		}
		Sonr = kc
		return nil
	} else {
		logger.Debug("Creating new Keychain")
		// Create Keychain
		kc, err := newKeychain(config)
		if err != nil {
			return err
		}

		// Return Keychain
		Sonr = kc
		return nil
	}
}

// keychain is a keychain implementation that stores keys in a directory.
type keychain struct {
	KeyChain
	config device.Folder

	// Key Pair References
	accountKeyPair keyPair
	groupKeyPair   keyPair
	linkKeyPair    keyPair
}

// loadKeychain loads a keychain from a file.
func loadKeychain(kcconfig device.Folder) (KeyChain, error) {
	// Create Keychain
	kc := &keychain{
		config: kcconfig,
	}

	// Read Account Key
	accPrivKey, accPubKey, err := readKey(kcconfig, Account)
	if err != nil {
		logger.Errorf("%s - Failed to Read Account Key", err)
		return nil, err
	}

	// Load Account Key to Keychain
	kc.LoadKeyPair(accPubKey, accPrivKey, Account)

	// Read Link Key
	linkPrivKey, linkPubKey, err := readKey(kcconfig, Link)
	if err != nil {
		logger.Errorf("%s - Failed to Read Link Key", err)
		return nil, err
	}

	// Load Link Key to Keychain
	kc.LoadKeyPair(linkPubKey, linkPrivKey, Link)

	// Read Group Key
	groupPrivKey, groupPubKey, err := readKey(kcconfig, Group)
	if err != nil {
		logger.Errorf("%s - Failed to Read Group Key", err)
		return nil, err
	}

	// Load Group Key to Keychain
	kc.LoadKeyPair(groupPubKey, groupPrivKey, Group)
	return kc, nil
}

// newKeychain creates a new keychain.
func newKeychain(folder device.Folder) (KeyChain, error) {
	// Create Keychain
	kc := &keychain{
		config: folder,
	}

	// Create New Account Key
	accPrivKey, accPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		logger.Errorf("%s - Failed to generate Account KeyPair", err)
		return nil, err
	}

	// Write Account Key to Disk
	err = writeKey(folder, accPrivKey, Account)
	if err != nil {
		return nil, err
	}

	// Load Account Key to Keychain
	kc.LoadKeyPair(accPubKey, accPrivKey, Account)

	// Create New Link Key
	linkPrivKey, linkPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		logger.Errorf("%s - Failed to generate Link KeyPair", err)
		return nil, err
	}

	// Write Link Key to Disk
	err = writeKey(folder, linkPrivKey, Link)
	if err != nil {
		return nil, err
	}

	// Load Link Key to Keychain
	kc.LoadKeyPair(linkPubKey, linkPrivKey, Link)

	// Create New Group Key
	groupPrivKey, groupPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		logger.Errorf("%s - Failed to generate Group KeyPair", err)
		return nil, err
	}

	// Write Group Key to Disk
	err = writeKey(folder, groupPrivKey, Group)
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
		logger.Errorf("%s - Failed to sign UUID", err)
		return nil, err
	}

	// Return UUID with signature
	return &SignedUUID{
		Value:     id,
		Signature: sig,
		Timestamp: time.Now().Unix(),
	}, nil
}

// CreateUUID makes a new UUID value signed by the local node's private key
func (kc *keychain) CreateFingerprint(sname string, deviceID string) (*SignedFingerprint, error) {
	// Initialize entropy
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		logger.Errorf("%s - Failed to generate entropy", err)
		return nil, err
	}

	// Generate mnemonic
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		logger.Errorf("%s - Failed to generate mnemonic", err)
		return nil, err
	}

	// Convert mnemonic to Buffer
	bufMnemonic, err := bip39.MnemonicToByteArray(mnemonic)
	if err != nil {
		logger.Errorf("%s - Failed to convert mnemonic to byte array", err)
		return nil, err
	}

	// sign UUID using local node's private key
	fingerprint, err := kc.SignWith(Account, bufMnemonic)
	if err != nil {
		logger.Errorf("%s - Failed to sign UUID", err)
		return nil, err
	}

	// Get Full Prefix without Substring
	prefixRaw, err := kc.SignHmacWith(Account, fmt.Sprintf("%s:%s", deviceID, sname))
	if err != nil {
		logger.Errorf("%s - Failed to sign UUID", err)
		return nil, err
	}

	// Find SNR PubKey Instance
	pubKey, err := kc.GetSnrPubKey(Account)
	if err != nil {
		logger.Errorf("%s - Failed to get public key", err)
		return nil, err
	}

	// Get Public KEy as string
	pubStr, err := pubKey.String()
	if err != nil {
		logger.Errorf("%s - Failed to convert public key to string", err)
		return nil, err
	}

	// Create Signed Fingerprint
	signed := &SignedFingerprint{
		Mnemonic:    mnemonic,
		SName:       sname,
		Prefix:      prefixRaw[:16],
		Fingerprint: fingerprint,
		DeviceID:    deviceID,
		PublicKey:   pubStr,
	}
	return signed, nil
}

// CreateMetadata makes message data shared between all node's p2p protocols
func (kc *keychain) CreateMetadata(peerID peer.ID) (*SignedMetadata, error) {
	// Get local node's public key
	pubKey, err := kc.GetPubKey(Account)
	if err != nil {
		logger.Errorf("%s - Failed to get local host's public key", err)
		return nil, err
	}

	// Marshal Public key into public key data
	nodePubKey, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Errorf("%s - Failed to Extract Public Key", err)
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

// GetSnrPubKey gets a public key from the keychain with a given keypair type
// as a SnrPubKey
func (kc *keychain) GetSnrPubKey(kp KeyPairType) (*SnrPubKey, error) {
	pub, err := kc.GetPubKey(kp)
	if err != nil {
		logger.Errorf("%s - Failed to get SnrPubKey", err)
		return nil, err
	}
	return NewSnrPubKey(pub), nil
}

// GetSnrPrivKey gets a private key from the keychain with a given keypair type
// as a SnrPrivKey
func (kc *keychain) GetSnrPrivKey(kp KeyPairType) (*SnrPrivKey, error) {
	priv, err := kc.GetPrivKey(kp)
	if err != nil {
		logger.Errorf("%s - Failed to get SnrPrivKey", err)
		return nil, err
	}
	return NewSnrPrivKey(priv), nil
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
		logger.Errorf("%s - Failed to load Key Pair", ErrInvalidKeyType)
	}
}

// RemoveKeyPair removes a key from the keychain.
func (kc *keychain) RemoveKeyPair(kp KeyPairType) error {
	if kc.Exists(kp) {
		return kc.config.Delete(kp.Path())
	}
	logger.Errorf("%s - Failed to Remove Key Pair", ErrKeychainUnready)
	return ErrKeychainUnready
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
	logger.Errorf("%s - Failed to Sign Data", ErrKeychainUnready)
	return nil, ErrKeychainUnready
}

// SignWith signs a message with the specified keypair with Hmac - Used to Sign Fingerprint of RecoveryCode
// on HDNS subdomain
func (kc *keychain) SignHmacWith(kp KeyPairType, msg string) (string, error) {
	if kc.Exists(kp) {
		// Find the private key
		priv, err := kc.GetPrivKey(kp)
		if err != nil {
			return "", err
		}

		// Get the private key as a byte array
		privBuf, err := priv.Raw()
		if err != nil {
			logger.Errorf("%s - Failed to Get PrivKey Raw Buffer", err)
			return "", err
		}

		// Create a new HMAC object
		h := hmac.New(sha256.New, privBuf)
		h.Write([]byte(msg))
		return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
	}
	logger.Errorf("%s - Failed to Sign Data", ErrKeychainUnready)
	return "", ErrKeychainUnready
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
	logger.Errorf("%s - Failed to Verify Data", ErrKeychainUnready)
	return false, ErrKeychainUnready
}

// VerifyHmacWith verifies a signature with specified pair - Used to Verify Fingerprint of RecoveryCode
// on HDNS subdomain
func (kc *keychain) VerifyHmacWith(kp KeyPairType, msg string, sig string) (bool, error) {
	if kc.Exists(kp) {
		// Find the public key
		pub, err := kc.GetPubKey(kp)
		if err != nil {
			return false, err
		}

		// Get the public key as a byte array
		pubBuf, err := pub.Raw()
		if err != nil {
			logger.Errorf("%s - Failed to Get PubKey Raw Buffer", err)
			return false, err
		}

		// Decode the signature
		sigBuf, err := base64.StdEncoding.DecodeString(sig)
		if err != nil {
			logger.Errorf("%s - Failed to Decode Signature", err)
			return false, err
		}

		// Create a new HMAC object
		h := hmac.New(sha256.New, pubBuf)
		h.Write([]byte(msg))
		return hmac.Equal(h.Sum(nil), sigBuf), nil
	}
	logger.Errorf("%s - Failed to Verify Data", ErrKeychainUnready)
	return false, ErrKeychainUnready
}
