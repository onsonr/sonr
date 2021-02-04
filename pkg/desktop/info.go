package desktop

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	md "github.com/sonr-io/core/internal/models"
)

type SysInfo struct {
	OLC       string
	Device    md.Device
	Directory md.Directories
}

// ^ Returns System Info ^ //
func SystemInfo() SysInfo {
	// Initialize Vars
	var platform md.Platform
	var model string
	var name string
	var docDir string
	var err error

	// Get Operating System
	runOs := runtime.GOOS

	// Check Runtime OS
	switch runOs {
	// @ Windows
	case "windows":
		platform = md.Platform_Windows

		// @ Mac
	case "darwin":
		platform = md.Platform_MacOS

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
	if docDir, err = os.UserHomeDir(); err != nil {
		log.Println(err)
		docDir = "local/temp"
	}

	// Return SysInfo Object
	return SysInfo{
		// Current Hard Code OLC
		OLC: "87C4XFJV+",

		// Retreived Device Info
		Device: md.Device{
			Platform: platform,
			Model:    model,
			Name:     name,
			Desktop:  true,
		},

		// Current Directories
		Directory: md.Directories{
			Documents: docDir,
			Temporary: filepath.Join(docDir, "Downloads"),
		},
	}
}
