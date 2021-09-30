package device

import (
	"errors"
	"fmt"
	"os"

	"github.com/sonr-io/core/tools/logger"
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

// Exists checks if the environment variable exists
func (ev EnvVariable) Exists() bool {
	return len(ev.Value) > 0 && len(ev.Key) > 0
}

// Get the environment variable
func (ev EnvVariable) Get() string {
	return ev.Value
}

// Set the environment variable
func (ev EnvVariable) Set(val string) {
	if !ev.Exists() {
		ev.Value = val
	}
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

// InitEnv initializes the environment variables
func InitEnv(isDev bool, envVars EnvVariableMap) {
	// Initialize logger
	logger.Init(isDev)

	// Check Map length
	if len(envVars) > 0 {
		// Set Variables from Map
		for k, v := range envVars {
			if len(v) > 0 {
				os.Setenv(k, v)
			} else {
				logger.Error(fmt.Sprintf("Failed to set Env Var: %s", k), ErrMissingEnvVar)
			}
		}
	}

	// Set Variables from OS
	HANDSHAKE_KEY.Set(os.Getenv("HANDSHAKE_KEY"))
	HANDSHAKE_SECRET.Set(os.Getenv("HANDSHAKE_SECRET"))
	IP_LOCATION_KEY.Set(os.Getenv("IP_LOCATION_KEY"))
	RAPID_API_KEY.Set(os.Getenv("RAPID_API_KEY"))
	TEXTILE_HUB_KEY.Set(os.Getenv("TEXTILE_HUB_KEY"))
	TEXTILE_HUB_SECRET.Set(os.Getenv("TEXTILE_HUB_SECRET"))
}
