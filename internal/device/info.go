package device

import (
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
)

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

// ID returns the device ID.
func ID() (string, error) {
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

// VendorName returns the vendor name.
func VendorName() string {
	return "Sonr"
}
