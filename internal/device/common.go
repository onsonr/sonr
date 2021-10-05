package device

import (
	"errors"
	"os"

	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/logger"
)

// Error definitions
var (
	// General errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")

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

// NewDeviceIDPrefix returns a new device ID prefix for users HDNS records
func NewDeviceIDPrefix(sName string) (string, error) {
	// Check if the device ID is empty
	if deviceID == "" {
		return "", ErrEmptyDeviceID
	}

	// Check if the SName is empty
	if sName == "" {
		return "", errors.New("SName cannot by Empty or Less than 4 characters.")
	}
	val := deviceID + sName
	return KeyChain.SignHmacWith(keychain.Account, val)
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

	// Mailbox is the type for Mailbox folder.
	Mailbox
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
	case Mailbox:
		if MailboxPath == "" {
			return "", ErrDirectoryUnset
		}
		return MailboxPath, nil
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
