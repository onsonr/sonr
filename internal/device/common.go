package device

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/keychain"
)

// Error definitions
var (
	logger = golog.Child("device")
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

// DirType is the type of a directory.
type DirType int

// Directory types
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

	// Textile is the type for Textile folder.
	Textile
)

// Path returns the path for the directory.
func (d DirType) Path() (string, error) {
	// Switch on the directory type
	switch d {
	case Support:
		if SupportPath == "" {
			return "", ErrDirectoryUnset
		}
		return SupportPath, nil
	case Temporary:
		if TempPath == "" {
			return "", ErrDirectoryUnset
		}
		return TempPath, nil
	case Documents:
		if DocsPath == "" {
			return "", ErrDirectoryUnset
		}
		return DocsPath, nil
	case Downloads:
		if DownloadsPath == "" {
			return "", ErrDirectoryUnset
		}
		return DownloadsPath, nil
	case Database:
		if DatabasePath == "" {
			return "", ErrDirectoryUnset
		}
		return DatabasePath, nil
	case Textile:
		if TextilePath == "" {
			return "", ErrDirectoryUnset
		}
		return TextilePath, nil
	default:
		return "", ErrDirectoryInvalid
	}
}

// Exists returns true if the directory exists.
func (d DirType) Exists() bool {
	// Get the directory path
	path, err := d.Path()
	if err != nil {
		logger.Error("Failed to get Directory path", err)
		return false
	}

	// Check if the directory exists
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// Join returns the path for the directory joined with the path.
func (d DirType) Join(path string) (string, error) {
	// Get the directory path
	dir, err := d.Path()
	if err != nil {
		return path, ErrDirectoryJoin
	}

	// Join the directory with the path
	return filepath.Join(dir, path), nil
}

// Has returns true if the directory has the file or directory.
func (d DirType) Has(p string) bool {
	// Get the directory path
	path, err := d.Join(p)
	if err != nil {
		logger.Error("Failed to determine joined path", err)
		return false
	}

	// Check if the directory has the file or directory
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
