package fs

import (
	"errors"
	"os"
	"runtime"

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
)

var (
	logger = golog.Child("internal/fs")
	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")
)

// Start creates new FileSystem
func Start(options ...Option) error {
	opts := defaultFsOptions()
	for _, opt := range options {
		opt(opts)
	}
	return opts.Apply()
}

type Option func(o *fsOptions)

// WithHomePath sets the Home Directory
func WithHomePath(p string) Option {
	return func(o *fsOptions) {
		// Set Home Directory
		if p != "" {
			o.HomeDir = p
		}
	}
}

// WithTempPath sets the Temporary Directory
func WithTempPath(p string) Option {
	return func(o *fsOptions) {
		// Set Home Directory
		if p != "" {
			o.TempDir = p
		}
	}
}

// WithSupportPath sets the Support Directory
func WithSupportPath(p string) Option {
	return func(o *fsOptions) {
		// Set Home Directory
		if p != "" {
			o.SupportDir = p
		}
	}
}

// fsOptions holds directory list
type fsOptions struct {
	HomeDir    string
	TempDir    string
	SupportDir string

	walletDir    string
	databaseDir  string
	downloadsDir string
	textileDir   string
}

// defaultFsOptions returns fsOptions
func defaultFsOptions() *fsOptions {
	opts := &fsOptions{}
	if checkIsDesktop() {
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
	}
	return opts
}

// Apply sets device directories for Path
func (fo *fsOptions) Apply() error {
	// Check for Valid
	var err error
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
	if checkIsDesktop() {
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

// IsFile returns true if the given path is a file
func IsFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// checkIsDesktop returns true if the current platform is desktop
func checkIsDesktop() bool {
	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		return false
	}
	return true
}

// checkIsMobile returns true if the current platform is mobile
func checkIsMobile() bool {
	return !checkIsDesktop()
}
