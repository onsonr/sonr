package device

import (
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

// DeviceStat is the device info struct
type DeviceStat struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Os        string `json:"os"`
	Arch      string `json:"arch"`
	IsDesktop bool
	IsMobile  bool
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

// ID returns the device ID.
func ID() (string, error) {
	return machineid.ID()
}

// Info returns the device info.
func Info() *common.Peer_Device {
	// Get HostName
	hn, err := HostName()
	if err != nil {
		hn = "unknown"
	}

	// Get Devices ID
	id, err := ID()
	if err != nil {
		id = "unknown"
	}

	// Return the device info for Peer
	return &common.Peer_Device{
		HostName: hn,
		Os:       runtime.GOOS,
		Id:       id,
		Arch:     runtime.GOARCH,
	}
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
func Stat() *DeviceStat {
	// Get Device Id
	id, err := ID()
	if err != nil {
		logger.Error("Failed to get Device ID", zap.Error(err))
	}

	// Get HostName
	hn, err := HostName()
	if err != nil {
		logger.Error("Failed to get HostName", zap.Error(err))
	}

	// Return the device info for Peer
	return &DeviceStat{
		Id:        id,
		Name:      hn,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		IsDesktop: IsDesktop(),
		IsMobile:  IsMobile(),
	}
}
