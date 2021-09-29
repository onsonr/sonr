package device

import (
	"errors"
	"fmt"
	"os"
)

// ENV Variables
var (
	// HDNS client key for Namebase.io
	HANDSHAKE_KEY = ""

	// HDNS secret key for Namebase.io
	HANDSHAKE_SECRET = ""

	// IP Location API key for IPStack.com
	IP_LOCATION_KEY = ""

	// RapidAPI key for RapidAPI.com
	RAPID_API_KEY = ""

	// Textile Hub API key
	TEXTILE_HUB_KEY = ""

	// Textile Hub secret key
	TEXTILE_HUB_SECRET = ""
)

// Error definitions
var (
	// General errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Missing Env Variable")

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

// InitEnv initializes the environment variables
func InitEnv() error {
	// Set environment variables
	HANDSHAKE_KEY = os.Getenv("HANDSHAKE_KEY")
	HANDSHAKE_SECRET = os.Getenv("HANDSHAKE_SECRET")
	IP_LOCATION_KEY = os.Getenv("IP_LOCATION_KEY")
	RAPID_API_KEY = os.Getenv("RAPID_API_KEY")
	TEXTILE_HUB_KEY = os.Getenv("TEXTILE_HUB_KEY")
	TEXTILE_HUB_SECRET = os.Getenv("TEXTILE_HUB_SECRET")

	// Check for missing environment variables
	if HANDSHAKE_KEY == "" {
		return envVarError("HANDSHAKE_KEY")
	}
	if HANDSHAKE_SECRET == "" {
		return envVarError("HANDSHAKE_SECRET")
	}
	if IP_LOCATION_KEY == "" {
		return envVarError("IP_LOCATION_KEY")
	}
	if RAPID_API_KEY == "" {
		return envVarError("RAPID_API_KEY")
	}
	if TEXTILE_HUB_KEY == "" {
		return envVarError("TEXTILE_HUB_KEY")
	}
	if TEXTILE_HUB_SECRET == "" {
		return envVarError("TEXTILE_HUB_SECRET")
	}
	return nil
}

// envVarError returns an error for missing environment variables
func envVarError(name string) error {
	return fmt.Errorf("Missing Env Variable for: %s", name)
}
