package device

import (
	"crypto/rand"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

type DeviceOptions struct {
	DocumentsPath string
	CachePath     string
	SupportPath   string
}

var (
	KCConfig    *config.Config
	DocsPath    string
	TempPath    string
	SupportPath string
)

type FSDirType int

const (
	Support FSDirType = iota
	Temporary
	Documents
)

type FSOption struct {
	Path string
	Type FSDirType
}

// Init initializes the keychain and returns a Keychain.
func Init(opts ...FSOption) (Keychain, error) {
	// Set environment variables
	err := initEnv()
	if err != nil {
		logger.Error("Failed to initialize environment variables: %s", zap.Error(err))
		return nil, err
	}

	// Check if Opts are set
	if len(opts) == 0 {
		// Create Device Config
		configDirs := config.New(VendorName(), AppName())

		// optional: local path has the highest priority
		folder := configDirs.QueryFolderContainsFile("setting.json")
		if folder != nil {
			data, err := folder.ReadFile("setting.json")
			if err != nil {
				return nil, err
			}
			json.Unmarshal(data, &KCConfig)
			logger.Debug("Loaded Config from Disk")
			KCConfig = folder
			return loadKeychain()
		} else {
			data, err := json.Marshal(&KCConfig)
			if err != nil {
				return nil, err
			}

			// Stores to user folder
			folders := configDirs.QueryFolders(config.Support)
			err = folders[0].WriteFile("setting.json", data)
			if err != nil {
				return nil, err
			}

			// Create Keychain
			logger.Debug("Created new Config")
			KCConfig = folders[0]
			return newKeychain()
		}
	} else {
		// Set Paths from Opts
		for _, opt := range opts {
			switch opt.Type {
			case Support:
				SupportPath = opt.Path
			case Temporary:
				TempPath = opt.Path
			case Documents:
				DocsPath = opt.Path
			}
		}

		// Create Device Config
		KCConfig = &config.Config{
			Path: SupportPath,
			Type: config.Support,
		}

		// Create Keychain
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

	err = KCConfig.WriteFile(kp.Path(), buf)
	if err != nil {
		return err
	}
	return nil
}
