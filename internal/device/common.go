package device

import (
	"errors"
	"fmt"

	"github.com/sonr-io/core/tools/logger"
)

var (
	ErrVariableNotFound = errors.New("EnvVariable not found on OS.getEnv()")
)

// EnvVariable represents an environment variable
type EnvVariable struct {
	Key   string
	Value string
}

// NewEnvVariable creates a new environment variable
func NewEnvVariable(key string) EnvVariable {
	return EnvVariable{Key: key}
}

// Get the environment variable
func (ev EnvVariable) Get() string {
	return ev.Value
}

// Set the environment variable
func (ev EnvVariable) Set(val string) {
	if len(val) > 0 {
		ev.Value = val
		logger.Info(fmt.Sprintf("Enviornment Variable Set: %s", val))
		return
	}
	logger.Error("Failed to set Enviornment variable", ErrVariableNotFound)
}

// EnvVariableMap is a map of environment variables
type EnvVariableMap map[string]string

// ENV Variables
var (
	// HDNS client key for Namebase.io
	HANDSHAKE_KEY EnvVariable = NewEnvVariable("HANDSHAKE_KEY")

	// HDNS secret key for Namebase.io
	HANDSHAKE_SECRET EnvVariable = NewEnvVariable("HANDSHAKE_SECRET")

	// IP Location API key for IPStack.com
	IP_LOCATION_KEY EnvVariable = NewEnvVariable("IP_LOCATION_KEY")

	// RapidAPI key for RapidAPI.com
	RAPID_API_KEY EnvVariable = NewEnvVariable("RAPID_API_KEY")

	// Textile Hub API key
	TEXTILE_HUB_KEY EnvVariable = NewEnvVariable("TEXTILE_HUB_KEY")

	// Textile Hub secret key
	TEXTILE_HUB_SECRET EnvVariable = NewEnvVariable("TEXTILE_HUB_SECRET")
)

// Error definitions
var (
	// General errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

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
