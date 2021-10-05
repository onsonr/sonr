package device

import (
	"encoding/json"

	"github.com/sonr-io/core/internal/keychain"
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

	// TextileDir is provided Mailbox Path.
	TextileDir string
}

var (
	// Keychain is the device keychain
	KeyChain keychain.Keychain

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

	// MailboxPath is the path to the mailbox folder
	TextilePath string

	// deviceID is the device ID. Either provided or found
	deviceID string
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

			// Create Keychain
			kc, err := keychain.NewKeychain(kcConfig)
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

			// Create Keychain
			kc, err := keychain.NewKeychain(kcConfig)
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
		kc, err := keychain.NewKeychain(kcConfig)
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
	return logger.Error("Failed to Set Device ID", ErrEmptyDeviceID)
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
	case Textile:
		TextilePath = o.Path
	}
}
