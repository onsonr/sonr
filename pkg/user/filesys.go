package user

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	md "github.com/sonr-io/core/pkg/models"
)

const K_SONR_CLIENT_DIR = ".sonr"
const K_FILE_QUEUE_NAME = "file-queue"

// @ Sonr File System Struct
type FileSystem struct {
	// Properties
	IsDesktop bool
	Call      md.NodeCallback

	// Directories
	Downloads string
	Main      string
	Temporary string
}

// ^ Method Initializes Root Sonr Directory ^ //
func SetFS(connEvent *md.ConnectionRequest, callback md.NodeCallback) (*FileSystem, error) {
	// Initialize
	var sonrPath string

	// Check for Client Type
	if connEvent.Device.GetIsDesktop() {
		// Init Path, Check for Path
		sonrPath = filepath.Join(connEvent.Directories.Home, K_SONR_CLIENT_DIR)
		if err := EnsureDir(sonrPath, 0755); err != nil {
			return nil, err
		}
	} else {
		// Set Path to Documents for Mobile
		sonrPath = connEvent.Directories.Documents
	}

	// Create SFS
	sfs := &FileSystem{
		IsDesktop: connEvent.Device.GetIsDesktop(),
		Downloads: connEvent.Directories.Downloads,
		Main:      sonrPath,
		Temporary: connEvent.Directories.Temporary,
		Call:      callback,
	}

	// sfs.Queue = q
	return sfs, nil
}

// ^ EnsureDir creates directory if it doesnt exist ^
func EnsureDir(path string, perm os.FileMode) error {
	_, err := IsDir(path)

	if os.IsNotExist(err) {
		err = os.Mkdir(path, perm)
		if err != nil {
			return fmt.Errorf("failed to ensure directory at %q: %q", path, err)
		}
	}
	return err
}

// ^ EnsureDir creates directory if it doesnt exist ^
func (sfs *FileSystem) IsFile(name string) bool {
	// Create File Path
	path := filepath.Join(sfs.Main, name)

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// ^ IsDir determines is the path given is a directory or not. ^
func IsDir(name string) (bool, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	if !fi.IsDir() {
		return false, fmt.Errorf("%q is not a directory", name)
	}
	return true, nil
}

// ^ WriteIncomingFile writes file to Disk ^
func (sfs *FileSystem) ReadFile(name string) ([]byte, error) {
	// Create File Path
	path := filepath.Join(sfs.Main, name)

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("User File Does Not Exist")
	} else {
		// @ Read User Data File
		dat, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return dat, nil
	}
}

// ^ WriteIncomingFile writes file to Disk ^
func (sfs *FileSystem) WriteFile(name string, data []byte) (string, error) {
	// Create File Path
	path := filepath.Join(sfs.Main, name)

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", err
	}
	return path, nil
}

// @ Helper: Finds Write Path for Incoming File
func (sfs *FileSystem) GetPathForPayload(load md.Payload, fileName string) string {
	// Check for Desktop
	if sfs.IsDesktop {
		return filepath.Join(sfs.Downloads, fileName)
	} else {
		// Check Load
		if load == md.Payload_MEDIA {
			return filepath.Join(sfs.Temporary, fileName)
		} else {
			return filepath.Join(sfs.Main, fileName)
		}
	}
}
