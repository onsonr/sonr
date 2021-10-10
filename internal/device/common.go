package device

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/keychain"
)

// Error definitions
var (
	logger = golog.Child("internal/device")
	// General errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")

	// Keychain errors
	ErrInvalidKeyType  = errors.New("Invalid KeyPair Type provided")
	ErrKeychainUnready = errors.New("Keychain has not been loaded")
	ErrNoPrivateKey    = errors.New("No private key in KeyPair")
	ErrNoPublicKey     = errors.New("No public key in KeyPair")

	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")
)

// DeviceOption is a function that modifies the node options.
type DeviceOption func(deviceOptions)

// WithRequest sets the initialize request.
func WithDirectoryPaths(paths map[string]string) DeviceOption {
	return func(o deviceOptions) {
		for k, v := range paths {
			switch strings.ToLower(k) {
			case "support":
				o.Folders = append(o.Folders, Folder{
					Type: FolderType_SUPPORT,
					Path: v,
				})
			case "temporary":
				o.Folders = append(o.Folders, Folder{
					Type: FolderType_TEMPORARY,
					Path: v,
				})
			case "home":
				home := Folder{
					Type: FolderType_HOME,
					Path: v,
				}
				if IsMobile() {
					o.Folders = append(o.Folders, buildMobileFolders(home)...)
				} else {
					o.Folders = append(o.Folders, home)
				}
			}
		}
	}
}

// deviceOptions are options for the device.
type deviceOptions struct {
	Folders []Folder
}

type getDirFunc func() (string, error)

func defaultDeviceOptions() deviceOptions {
	opts := deviceOptions{
		Folders: make([]Folder, 0),
	}
	var AddDir = func(f getDirFunc, t FolderType) {
		p, err := f()
		if err != nil {
			logger.Error("Failed to get directory", err)
			return
		}
		opts.Folders = append(opts.Folders, Folder{
			Type: t,
			Path: p,
		})
	}

	if IsDesktop() {
		AddDir(UserHomePath, FolderType_HOME)
		AddDir(UserTemporaryPath, FolderType_TEMPORARY)
		AddDir(UserSupportPath, FolderType_SUPPORT)
		return opts
	}
	return opts
}

// NewRecordPrefix returns a new device ID prefix for users HDNS records
func NewRecordPrefix(sName string) (string, error) {
	// Check if the device ID is empty
	if deviceID == "" {
		return "", ErrEmptyDeviceID
	}

	// Check if the SName is empty
	if sName == "" {
		return "", errors.New("SName cannot by Empty or Less than 4 characters.")
	}
	val := fmt.Sprintf("%s:%s", deviceID, sName)
	return KeyChain.SignHmacWith(keychain.Account, val)
}

// VerifyRecordPrefix returns true if the prefix is valid for the device ID.
func VerifyRecordPrefix(prefix string, sName string) bool {
	// Check if the prefix is empty
	if prefix == "" {
		logger.Warn("Empty Prefix Provided as Parameter")
		return false
	}

	// Check if the prefix is valid
	val := fmt.Sprintf("%s:%s", deviceID, sName)
	ok, err := KeyChain.VerifyHmacWith(keychain.Account, prefix, val)
	if err != nil {
		logger.Error("Failed to verify prefix", err)
		return false
	}
	return ok
}
