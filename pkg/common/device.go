package common

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
	"github.com/sonr-io/core/internal/wallet"
)

var (
	// General errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")
)

var (
	// deviceID is the device ID. Either provided or found
	deviceID string
)

// Arch returns the current architecture.
func Arch() string {
	return runtime.GOARCH
}

// HostName returns the hostname of the current machine.
func HostName() (string, error) {
	return os.Hostname()
}

// ID returns the device ID.
func ID() (string, error) {
	// Check if Mobile
	if IsMobile() {
		if deviceID != "" {
			return deviceID, nil
		}
		return "", errors.New("Device ID not set for Mobile.")
	}
	return machineid.ID()
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
	return wallet.Sonr.SignHmacWith(wallet.Account, val)
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

// Platform returns formatted GOOS for Text format.
// Returns: ["MacOS", "Windows", "Linux", "Android", "iOS"]
func Platform() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "MacOS"
	case "linux":
		return "Linux"
	case "android":
		return "Android"
	case "ios":
		return "iOS"
	default:
		return "Unknown"
	}
}

// Stat returns the device stat.
func Stat() (map[string]string, error) {
	// Get Device Id
	id, err := ID()
	if err != nil {
		logger.Error("Failed to get Device ID", err)
		return nil, err
	}

	// Get HostName
	hn, err := HostName()
	if err != nil {
		logger.Error("Failed to get HostName", err)
		return nil, err
	}

	// Return the device info for Peer
	return map[string]string{
		"id":       id,
		"hostName": hn,
		"os":       runtime.GOOS,
		"arch":     runtime.GOARCH,
	}, nil
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
	ok, err := wallet.Sonr.VerifyHmacWith(wallet.Account, prefix, val)
	if err != nil {
		logger.Error("Failed to verify prefix", err)
		return false
	}
	return ok
}
