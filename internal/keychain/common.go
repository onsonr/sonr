package keychain

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
)

// Error definitions
var (
	logger             = golog.Child("internal/keychain")
	ErrInvalidKeyType  = errors.New("Invalid KeyPair Type provided")
	ErrKeychainUnready = errors.New("Keychain has not been loaded")
	ErrNoPrivateKey    = errors.New("No private key in KeyPair")
	ErrNoPublicKey     = errors.New("No public key in KeyPair")
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

// keychainExists checks if EVERY key pair exists in the keychain.
func keychainExists(kcConfig *config.Config) bool {
	accExists := kcConfig.Exists(Account.Path())
	linkExists := kcConfig.Exists(Link.Path())
	groupExists := kcConfig.Exists(Group.Path())
	return accExists && linkExists && groupExists
}

// readKey reads a key from a file and returns privKey and pubKey.
func readKey(kcconfig *config.Config, kp KeyPairType) (crypto.PrivKey, crypto.PubKey, error) {
	// Get Buffer
	dat, err := kcconfig.ReadFile(kp.Path())
	if err != nil {
		return nil, nil, err
	}

	// Get Private Key from Buffer
	privKey, err := crypto.UnmarshalPrivateKey(dat)
	if err != nil {
		return nil, nil, err
	}
	priv, pub := newSnrKeyPair(privKey)

	// Get Public Key from Private Key
	return priv, pub, nil
}

// writeKey writes a key to the keychain.
func writeKey(kcconfig *config.Config, privKey crypto.PrivKey, kp KeyPairType) error {
	// Marshal Private Key
	buf, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return err
	}

	// Write Key to Keychain
	err = kcconfig.WriteFile(kp.Path(), buf)
	if err != nil {
		return err
	}
	return nil
}
