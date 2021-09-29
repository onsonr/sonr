package device

import (
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
)

// DeviceOptions are options for the device.
type DeviceOptions struct {
	// DatabaseDir is provided Database Path.
	DatabaseDir string

	// DocumentsDir is provided Docs Path.
	DocumentsDir string

	// CacheDir is provided Cache Path.
	CacheDir string

	// SupportDir is provided Support Path.
	SupportDir string

	// DownloadsDir is provided Downloads Path.
	DownloadsDir string
}

var (
	// Keychain is the device keychain
	KeyChain Keychain

	// DatabasePath is the path to the database folder
	DatabasePath string

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

	// Database is the type for Database folder.
	Database
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

		// Check if Keychain exists
		if keychainExists(kcConfig) {
			// Load Existing Keychain
			kc, err := loadKeychain(kcConfig)
			if err != nil {
				return err
			}

			// Set Keychain
			KeyChain = kc
		} else {
			// Create Keychain
			kc, err := newKeychain(kcConfig)
			if err != nil {
				return err
			}

			// Set Keychain
			KeyChain = kc
		}
	}
	return nil
}

// SetDeviceID sets the device ID.
func SetDeviceID(id string) error {
	if id != "" {
		deviceID = id
		return nil
	}
	return logger.Error("Failed to Set Device ID", ErrEmptyDeviceID)
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
	// Get Public Key from Private Key
	return privKey, privKey.GetPublic(), nil
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
	case Database:
		DatabasePath = o.Path
	}
}
