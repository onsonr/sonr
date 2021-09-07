package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CacheFolder() string {
	// Check Desktop
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if runtime.GOOS == "darwin" {
			return os.Getenv("HOME") + "/Library/Caches"
		} else if runtime.GOOS == "windows" {
			return os.Getenv("LOCALAPPDATA")
		} else {
			return filepath.Join(os.Getenv("HOME"), ".cache")
		}
	} else {
		return ""
	}
}

func GlobalSettingFolder() string {
	// Check Desktop
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		// Check OS
		if runtime.GOOS == "darwin" {
			// MacOS
			return os.Getenv("HOME") + "/Library/Application Support"
		} else if runtime.GOOS == "windows" {
			// Windows
			return os.Getenv("APPDATA")
		} else {
			// Linux or other
			if os.Getenv("XDG_CONFIG_HOME") != "" {
				return os.Getenv("XDG_CONFIG_HOME")
			} else {
				return filepath.Join(os.Getenv("HOME"), ".config")
			}
		}
	} else {
		return ""
	}
}

func SystemSettingFolders() []string {
	// Check Desktop
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		// Check OS
		if runtime.GOOS == "darwin" {
			// MacOS
			return []string{"/Library/Application Support"}
		} else if runtime.GOOS == "windows" {
			// Windows
			return []string{os.Getenv("PROGRAMDATA")}
		} else {
			// Linux or other
			if os.Getenv("XDG_CONFIG_DIRS") != "" {
				return strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")
			} else {
				return []string{"/etc/xdg"}
			}
		}
	} else {
		return []string{}
	}

}
