package device

import (
	"errors"
	"runtime"

	"github.com/denisbrodbeck/machineid"
)

func init() {
	if IsDesktop() {
		Init()
	}
}

// Init initializes the device package.
func Init(options ...Option) error {
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	return opts.Apply()
}

// Arch returns the current architecture.
func Arch() string {
	return runtime.GOARCH
}

// HostName returns the hostname of the current machine.
func HostName() (string, error) {
	if hostName != "" {
		return hostName, nil
	}
	return "", errors.New("HostName not set.")
}

// ID returns the device ID.
func ID() (string, error) {
	// Check if the device ID is empty
	if deviceID != "" {
		return deviceID, nil
	}

	if IsDesktop() {
		retdeviceID, err := machineid.ID()
		if err != nil {
			logger.Errorf("%s - Failed to get Device ID", err)
			return "", err
		}
		deviceID = retdeviceID
		return deviceID, nil
	}
	return "", errors.New("Device ID not set.")
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
		logger.Errorf("%s - Failed to get Device ID", err)
		return nil, err
	}

	// Get HostName
	hn, err := HostName()
	if err != nil {
		logger.Errorf("%s - Failed to get HostName", err)
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
