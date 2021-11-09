package wallet

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/pkg/device"
)

// Error Definitions
var (
	logger             = golog.Default.Child("internal/wallet")
	ErrInvalidKeyType  = errors.New("Invalid KeyPair Type provided")
	ErrKeychainUnready = errors.New("Keychain has not been loaded")
	ErrNoPrivateKey    = errors.New("No private key in KeyPair")
	ErrNoPublicKey     = errors.New("No public key in KeyPair")
)

// NodeOption is a function that modifies the node options.
type WalletOption func(*walletOptions)

// walletOptions is a collection of options for the node.
type walletOptions struct {
	directory string
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *walletOptions {
	// path, _ := device.NewSupportPath(".wallet", device.CreateDirIfNotExist())
	return &walletOptions{
		// directory: path,
	}
}

// keychainExists checks if EVERY key pair exists in the keychain.
func keychainExists(folder device.Folder) bool {
	accExists := folder.Exists(Account.Path())
	linkExists := folder.Exists(Link.Path())
	groupExists := folder.Exists(Group.Path())
	return accExists && linkExists && groupExists
}

// readKey reads a key from a file and returns privKey and pubKey.
func readKey(folder device.Folder, kp KeyPairType) (crypto.PrivKey, crypto.PubKey, error) {
	// Get Buffer
	dat, err := folder.ReadFile(kp.Path())
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
func writeKey(folder device.Folder, privKey crypto.PrivKey, kp KeyPairType) error {
	// Marshal Private Key
	buf, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return err
	}

	// Write Key to Keychain
	err = folder.WriteFile(kp.Path(), buf)
	if err != nil {
		return err
	}
	return nil
}
