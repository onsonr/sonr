package client

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	md "github.com/sonr-io/core/pkg/models"
)

// ^ Returns System Info ^ //
func SystemInfo() SysInfo {
	// Initialize Vars
	var platform md.Platform
	var model string
	var name string
	var homeDir string
	var libDir string
	var last string
	var err error

	// Get Operating System
	runOs := runtime.GOOS

	// Check Runtime OS
	switch runOs {
	// @ Windows
	case "windows":
		platform = md.Platform_Windows
		last = "PC"

		// @ Mac
	case "darwin":
		platform = md.Platform_MacOS
		last = "Mac"

		// @ Linux
	case "linux":
		platform = md.Platform_Linux

		// @ Unknown
	default:
		platform = md.Platform_Unknown
	}

	// Get Hostname
	if name, err = os.Hostname(); err != nil {
		log.Println(err)
		name = "Undefined"
	}

	// Get Directories
	if homeDir, err = os.UserHomeDir(); err != nil {
		log.Println(err)
		homeDir = "local/temp"
	}

	if libDir, err = os.UserConfigDir(); err != nil {
		log.Println(err)
		libDir = "local/temp"
	}

	// Return SysInfo Object
	return SysInfo{
		// Current Hard Code OLC
		OLC:           "87C4XFJV+",
		TempFirstName: "Prad's",
		TempLastName:  last,

		// Retreived Device Info
		Device: md.Device{
			Platform: platform,
			Model:    model,
			Name:     name,
			Desktop:  true,
		},

		// Current Directories
		Directory: md.Directories{
			Temporary: libDir,
			Documents: filepath.Join(homeDir, "Documents"),
			Downloads: filepath.Join(homeDir, "Downloads"),
			Home:      homeDir,
		},
	}
}
