package device

import (
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/denisbrodbeck/machineid"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/keychain"
)

var (
	// Keychain is the device keychain
	KeyChain keychain.Keychain

	FS fileSystem

	// deviceID is the device ID. Either provided or found
	deviceID string
)

// Init initializes the keychain and returns a Keychain.
func Init(options ...DeviceOption) error {
	// Set Paths from Opts
	opts := defaultDeviceOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Initialize fileSystem
	err := initFileSystem(opts)
	if err != nil {
		return err
	}

	// Create Keychain
	kc, err := keychain.NewKeychain(FS.Support.Path)
	if err != nil {
		return err
	}
	KeyChain = kc
	return nil
}

type fileSystem struct {
	// Home is the home directory of the user
	Home Folder

	// Database is the path to the database folder
	Database Folder

	// Documents is the path to the documents folder
	Documents Folder

	// Downloads is the path to the downloads folder
	Downloads Folder

	// Temporary is the path to the temporary/cache folder
	Temporary Folder

	// Support is the path to the support folder
	Support Folder

	// Textile is the path to the mailbox folder
	Textile Folder
}

func initFileSystem(opts deviceOptions) error {
	for _, f := range opts.Folders {
		err := f.MkdirAll()
		if err != nil {
			return err
		}
	}
	FS = fileSystem{
		Home:      findFolder(opts.Folders, "home"),
		Database:  findFolder(opts.Folders, "database"),
		Documents: findFolder(opts.Folders, "documents"),
		Downloads: findFolder(opts.Folders, "downloads"),
		Temporary: findFolder(opts.Folders, "temporary"),
		Support:   findFolder(opts.Folders, "support"),
		Textile:   findFolder(opts.Folders, "textile"),
	}
	return nil
}

func findFolder(fs []Folder, name string) Folder {
	for _, dir := range fs {
		if strings.ToLower(dir.Name()) == strings.ToLower(name) {
			return dir
		}
	}
	return Folder{
		Type: FolderType_NONE,
	}
}

// AppName returns the application name.
func AppName() string {
	switch runtime.GOOS {
	case "android":
		return "io.sonr.petitfour"
	case "darwin":
		return "io.sonr.macos"
	case "linux":
		return "io.sonr.linux"
	case "windows":
		return "io.sonr.windows"
	case "ios":
		return "io.sonr.alpine"
	default:
		return "io.sonr.app"
	}
}

// HostName returns the hostname of the current machine.
func HostName() string {
	hostname, err := os.Hostname()
	if err == nil {
		return hostname
	}
	name, err := UserDisplayName()
	if err == nil {
		return name
	}
	return PrettyOS()
}

// ID returns the device ID.
func ID() string {
	// Check if Mobile
	if IsMobile() {
		if deviceID != "" {
			return deviceID
		}
		return "unknown"
	}
	id, err := machineid.ID()
	if err == nil {
		return id
	}
	return "unknown"
}

// IsMobile returns true if the current platform is ANY mobile platform.
func IsMobile() bool {
	return runtime.GOOS == "android" || runtime.GOOS == "ios"
}

// IsDesktop returns true if the current platform is ANY desktop platform.
func IsDesktop() bool {
	return runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin"
}

// IsAndroid returns true if the current platform is android.
func IsAndroid() bool {
	return runtime.GOOS == "android"
}

// IsIOS returns true if the current platform is iOS.
func IsIOS() bool {
	return runtime.GOOS == "ios"
}

// IsWindows returns true if the current platform is windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns true if the current platform is linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMacOS returns true if the current platform is macOS.
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// PrettyArch returns the formatted architecture name.
func PrettyArch() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x64"
	case "arm":
		return "ARM"
	case "arm64":
		return "ARM64"
	case "386":
		return "x86"
	}
	return runtime.GOARCH
}

// PrettyOS returns the formatted OS name.
func PrettyOS() string {
	switch runtime.GOOS {
	case "android":
		return "Android"
	case "darwin":
		return "MacOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	case "ios":
		return "iOS"
	default:
		return "Unknown"
	}
}

// SetDeviceID sets the device ID.
func SetDeviceID(id string) error {
	if id != "" {
		deviceID = id
		return nil
	}
	logger.Error("Failed to Set Device ID", ErrEmptyDeviceID)
	return ErrEmptyDeviceID
}

// UserDisplayName returns the user display name on Desktop.
func UserDisplayName() (string, error) {
	if IsDesktop() {
		user, err := user.Current()
		if err != nil {
			return "", errors.WithMessage(err, "Failed to Lookup Current User")
		}
		return user.Name, nil
	}
	return HostName(), nil
}

// UserHomeDir returns the user home directory.
func UserHomeDir() (Folder, error) {
	if IsDesktop() {
		p, err := os.UserHomeDir()
		if err != nil {
			return Folder{Type: FolderType_NONE}, err
		}
		return NewFolder(p, FolderType_HOME), nil
	}
	return NewFolder("", FolderType_HOME), nil
}

// UserTemporaryDir returns the user temporary directory.
func UserTemporaryDir() (Folder, error) {
	if IsDesktop() {
		p, err := os.UserCacheDir()
		if err != nil {
			return Folder{Type: FolderType_NONE}, err
		}
		return NewFolder(p, FolderType_TEMPORARY), nil
	}
	return NewFolder("", FolderType_TEMPORARY), nil
}

// UserSupportDir returns the user support directory.
func UserSupportDir() (Folder, error) {
	if IsDesktop() {
		p, err := os.UserConfigDir()
		if err != nil {
			return Folder{Type: FolderType_NONE}, err
		}
		return NewFolder(p, FolderType_SUPPORT), nil
	}
	return NewFolder("", FolderType_SUPPORT), nil
}

// UserHomeDir returns the user home directory.
func UserHomePath() (string, error) {
	if IsDesktop() {
		p, err := os.UserHomeDir()
		if err == nil {
			return p, nil
		}
		return "", err
	}
	return "", nil
}

// UserTemporaryDir returns the user temporary directory.
func UserTemporaryPath() (string, error) {
	if IsDesktop() {
		p, err := os.UserCacheDir()
		if err == nil {
			return p, nil
		}
		return "", err
	}
	return "", nil
}

// UserSupportDir returns the user support directory.
func UserSupportPath() (string, error) {
	if IsDesktop() {
		p, err := os.UserConfigDir()
		if err == nil {
			return p, nil
		}
		return "", err
	}
	return "", nil
}

// VendorName returns the vendor name.
func VendorName() string {
	return "Sonr"
}
