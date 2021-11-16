package device

import (
	"errors"
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/kataras/golog"
)

var (
	// Determined/Provided Paths
	Home      Folder // ApplicationDocumentsDir on Mobile, HOME_DIR on Desktop
	Support   Folder // AppSupport Directory
	Temporary Folder // AppCache Directory

	// Calculated Paths
	Database   Folder // Device DB Folder
	Downloads  Folder // Temporary Directory on Mobile for Export, Downloads on Desktop
	Wallet     Folder // Encrypted Storage Directory
	ThirdParty Folder // Sub-Directory of Support, used for Textile

	// deviceID is the device ID. Either provided or found
	deviceID string

	// hostName is the host name. Either provided or found
	hostName string
)

var (
	logger = golog.Default.Child("internal/device")

	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")

	// Device ID Errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")
)

type Option func(o *options)

// SetDeviceID sets the device ID
func SetDeviceID(id string) Option {
	return func(o *options) {
		// Set Home Directory
		if id != "" {
			o.deviceID = id
		}
	}
}

// WithHomePath sets the Home Directory
func WithHomePath(p string) Option {
	return func(o *options) {
		// Set Home Directory
		if p != "" {
			o.HomeDir = p
		}
	}
}

// WithTempPath sets the Temporary Directory
func WithTempPath(p string) Option {
	return func(o *options) {
		// Set Home Directory
		if p != "" {
			o.TempDir = p
		}
	}
}

// WithSupportPath sets the Support Directory
func WithSupportPath(p string) Option {
	return func(o *options) {
		// Set Home Directory
		if p != "" {
			o.SupportDir = p
		}
	}
}

// options holds directory list
type options struct {
	HomeDir    string
	TempDir    string
	SupportDir string

	walletDir    string
	databaseDir  string
	downloadsDir string
	textileDir   string
	deviceID     string
}

// defaultOptions returns fsOptions
func defaultOptions() *options {
	opts := &options{}
	if IsDesktop() {
		hp, err := os.UserHomeDir()
		if err != nil {
			logger.Errorf("%s - Failed to get HomeDir, ", err)
		} else {
			opts.HomeDir = hp
		}

		tp, err := os.UserCacheDir()
		if err != nil {
			logger.Errorf("%s - Failed to get TempDir, ", err)
		} else {
			opts.TempDir = tp
		}

		sp, err := os.UserConfigDir()
		if err != nil {
			logger.Errorf("%s - Failed to get SupportDir, ", err)
		} else {
			opts.SupportDir = sp
		}

		id, err := machineid.ID()
		if err != nil {
			logger.Errorf("%s - Failed to get Device ID", err)
		} else {
			opts.deviceID = id
		}
	}
	return opts
}

// Apply sets device directories for Path
func (fo *options) Apply() error {
	// Get the hostname
	hn, err := os.Hostname()
	if err != nil {
		logger.Errorf("%s - Failed to get HostName", err)
		return err
	}
	hostName = hn

	// Check if deviceID is set
	if fo.deviceID == "" {
		logger.Errorf("%s - Device ID is empty", ErrEmptyDeviceID)
		return ErrEmptyDeviceID
	}
	deviceID = fo.deviceID

	// Check for Valid
	if fo.HomeDir == "" {
		return errors.New("Home Directory was not set.")
	}
	if fo.SupportDir == "" {
		return errors.New("Support Directory was not set.")
	}
	if fo.TempDir == "" {
		return errors.New("Temporary Directory was not set.")
	}

	// Set Home Folder
	Home = Folder(fo.HomeDir)
	Support = Folder(fo.SupportDir)
	Temporary = Folder(fo.TempDir)

	// Create Downloads Folder
	if IsDesktop() {
		Downloads, err = Home.CreateFolder("Downloads")
		if err != nil {
			return err
		}
	} else {
		Downloads, err = Temporary.CreateFolder("Downloads")
		if err != nil {
			return err
		}
	}

	// Create Database Folder
	Database, err = Support.CreateFolder(".db")
	if err != nil {
		return err
	}

	// Create Wallet Folder
	Wallet, err = Support.CreateFolder(".wallet")
	if err != nil {
		return err
	}

	// Create Third Party Folder
	ThirdParty, err = Support.CreateFolder("third_party")
	if err != nil {
		return err
	}
	return nil
}
