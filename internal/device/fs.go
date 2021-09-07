package device

import (
	"crypto/rand"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
)

// keychain is a keychain implementation that stores keys in a directory.
type keychain struct {
	Keychain
}

var (
	Config   *config.Config
	HomePath string
	fsReady  bool
)

// Init initializes the keychain and returns a Keychain.
func Init() (Keychain, error) {
	// Create Device Config
	configDirs := config.New(VendorName(), AppName())

	// optional: local path has the highest priority
	folder := configDirs.QueryFolderContainsFile("setting.json")
	if folder != nil {
		data, err := folder.ReadFile("setting.json")
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, &Config)
		logger.Debug("Loaded Config from Disk")
		Config = folder
		return loadKeychain()
	} else {
		data, err := json.Marshal(&Config)
		if err != nil {
			return nil, err
		}

		// Stores to user folder
		folders := configDirs.QueryFolders(config.Global)
		err = folders[0].WriteFile("setting.json", data)
		if err != nil {
			return nil, err
		}

		// Create Keychain
		logger.Debug("Created new Config")
		Config = folders[0]
		return newKeychain()
	}

}

// loadKeychain loads a keychain from a file.
func loadKeychain() (Keychain, error) {
	kc := &keychain{}
	return kc, nil
}

// newKeychain creates a new keychain.
func newKeychain() (Keychain, error) {
	// Create New Account Key
	accPrivKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Create New Group Key
	groupPrivKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Create New Link Key
	linkPrivKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Add Account Key to Keychain
	err = writeKey(Account, accPrivKey)
	if err != nil {
		return nil, err
	}

	// Add Group Key to Keychain
	err = writeKey(Group, groupPrivKey)
	if err != nil {
		return nil, err
	}

	// Add Link Key to Keychain
	err = writeKey(Link, linkPrivKey)
	if err != nil {
		return nil, err
	}
	return &keychain{}, nil
}

// writeKey writes a key to the keychain.
func writeKey(kp KeyPair, privKey crypto.PrivKey) error {
	// Write Key to Keychain
	buf, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return err
	}

	err = Config.WriteFile(kp.Path(), buf)
	if err != nil {
		return err
	}
	return nil
}
