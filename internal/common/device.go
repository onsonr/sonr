package common

import (
	"errors"
	"os"
	"runtime"
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

// SetDeviceID sets the device ID.
func SetDeviceID(id string) error {
	if id != "" {
		deviceID = id
		return nil
	}
	logger.Error("Failed to Set Device ID", ErrEmptyDeviceID)
	return ErrEmptyDeviceID
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

// Arch returns the current architecture.
func Arch() string {
	return runtime.GOARCH
}

// HostName returns the hostname of the current machine.
func HostName() (string, error) {
	return os.Hostname()
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

// VendorName returns the vendor name.
func VendorName() string {
	return "Sonr"
}

// Stat returns the device stat.
func Stat() (map[string]string, error) {
	// Get HostName
	hn, err := HostName()
	if err != nil {
		logger.Error("Failed to get HostName", err)
		return nil, err
	}

	// Return the device info for Peer
	return map[string]string{
		"HostName": hn,
		"Os":       runtime.GOOS,
		"Arch":     runtime.GOARCH,
	}, nil
}
