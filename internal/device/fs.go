package device

import (
	"crypto/rand"
	"encoding/json"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
)

// DeviceOptions are options for the device.
type DeviceOptions struct {
	// DocumentsPath is provided Docs Path.
	DocumentsPath string

	// CachePath is provided Cache Path.
	CachePath string

	// SupportPath is provided Support Path.
	SupportPath string

	// DownloadsPath is provided Downloads Path.
	DownloadsPath string
}

var (
	// Keychain is the device keychain
	KeyChain Keychain

	// DocsPath is the path to the documents folder
	DocsPath string

	// DownloadsPath is the path to the downloads folder
	DownloadsPath string

	// TempPath is the path to the temporary/cache folder
	TempPath string

	// SupportPath is the path to the support folder
	SupportPath string

	// deviceID is the device ID. Either provided or found
	deviceID string
)

// DirType is the type of a directory.
type DirType int

const (
	// Support is the type for a support directory.
	Support DirType = iota

	// Temporary is the type for a temporary directory.
	Temporary

	// Documents is the type for Documents folder.
	Documents

	// Downloads is the type for Downloads folder.
	Downloads
)

// Init initializes the keychain and returns a Keychain.
func Init(isDev bool, opts ...FSOption) error {
	// Initialize logger
	logger.Init(isDev)

	// Check if Opts are set
	if len(opts) == 0 {
		// Create Device Config
		configDirs := config.New(VendorName(), AppName())

		// optional: local path has the highest priority
		folder := configDirs.QueryFolderContainsFile("setting.json")
		if folder != nil {
			data, err := folder.ReadFile("setting.json")
			if err != nil {
				return err
			}
			kcConfig := &config.Config{}
			json.Unmarshal(data, &kcConfig)
			logger.Debug("Loaded Config from Disk")
			kcConfig = folder
			kc, err := loadKeychain(kcConfig)
			if err != nil {
				return err
			}
			KeyChain = kc
		} else {
			kcConfig := &config.Config{}
			data, err := json.Marshal(&kcConfig)
			if err != nil {
				return err
			}

			// Stores to user folder
			folders := configDirs.QueryFolders(config.Support)
			err = folders[0].WriteFile("setting.json", data)
			if err != nil {
				return err
			}

			// Create Keychain
			logger.Debug("Created new Config")
			kcConfig = folders[0]
			kc, err := newKeychain(kcConfig)
			if err != nil {
				return err
			}
			KeyChain = kc
		}
	} else {
		// Set Paths from Opts
		for _, opt := range opts {
			opt.apply()
		}

		// Create Device Config
		kcConfig := &config.Config{
			Path: SupportPath,
			Type: config.Support,
		}

		// Create Keychain
		kc, err := newKeychain(kcConfig)
		if err != nil {
			return err
		}
		KeyChain = kc
	}
	return nil
}

// SetDeviceID sets the device ID.
func SetDeviceID(id string) error {
	if id != "" {
		deviceID = id
		return nil
	}
	return errors.New("Empty DeviceID provided.")
}

// loadKeychain loads a keychain from a file.
func loadKeychain(kcconfig *config.Config) (Keychain, error) {
	kc := &keychain{
		config: kcconfig,
	}
	return kc, nil
}

// newKeychain creates a new keychain.
func newKeychain(kcconfig *config.Config) (Keychain, error) {
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
	err = writeKey(kcconfig, Account, accPrivKey)
	if err != nil {
		return nil, err
	}

	// Add Group Key to Keychain
	err = writeKey(kcconfig, Group, groupPrivKey)
	if err != nil {
		return nil, err
	}

	// Add Link Key to Keychain
	err = writeKey(kcconfig, Link, linkPrivKey)
	if err != nil {
		return nil, err
	}
	return &keychain{
		config: kcconfig,
	}, nil
}

// writeKey writes a key to the keychain.
func writeKey(kcconfig *config.Config, kp KeyPair, privKey crypto.PrivKey) error {
	// Write Key to Keychain
	buf, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return err
	}

	err = kcconfig.WriteFile(kp.Path(), buf)
	if err != nil {
		return err
	}
	return nil
}

// FSOption is a functional option for configuring the filesystem.
type FSOption struct {
	Path string  // Path to the directory
	Type DirType // Type of the directory
}

// apply applies the given options to the filesystem.
func (o FSOption) apply() {
	switch o.Type {
	case Support:
		SupportPath = o.Path
	case Temporary:
		TempPath = o.Path
	case Documents:
		DocsPath = o.Path
	case Downloads:
		DownloadsPath = o.Path
	}
}
