package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/getsentry/sentry-go"
	dq "github.com/joncrlsn/dque"
	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
)

const K_SONR_CLIENT_DIR = ".sonr"

// @ Sonr File System Struct
type FileSystem struct {
	// Properties
	Initialized bool
	IsDesktop   bool

	// Directories
	Downloads string
	Root      string
	Temporary string

	// Queue
	Files        []*ProcessedFile
	CurrentCount int
	Call         dt.NodeCallback
	Queue        dq.DQue
}

// ^ Method Initializes Root Sonr Directory ^ //
func InitFS(connEvent *md.ConnectionRequest, profile *md.Profile, callback dt.NodeCallback) *FileSystem {
	// Initialize
	var sonrPath string
	var hasInitialized bool

	// Check for Client Type
	if connEvent.IsDesktop() {
		// Init Path, Check for Path
		sonrPath = filepath.Join(connEvent.Directories.Home, K_SONR_CLIENT_DIR)
		if err := EnsureDir(sonrPath, 0755); err != nil {
			sentry.CaptureException(err)
			hasInitialized = false
		} else {
			hasInitialized = true
		}
	} else {
		// Set Path to Documents for Mobile
		hasInitialized = true
		sonrPath = connEvent.Directories.Documents
	}

	// Create SFS
	sfs := &FileSystem{
		IsDesktop:   connEvent.Device.GetIsDesktop(),
		Initialized: hasInitialized,
		Downloads:   connEvent.Directories.Downloads,
		Root:        sonrPath,
		Temporary:   connEvent.Directories.Temporary,

		Files:        make([]*ProcessedFile, maxFileBufferSize),
		CurrentCount: 0,
		Call:         callback,
	}

	// Write User
	// if hasInitialized {
	// 	// if err := sfs.WriteUser(user); err != nil {
	// 	// 	sentry.CaptureException(err)
	// 	// }
	// }
	return sfs
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
	if sfs.Initialized {
		// Create File Path
		path := filepath.Join(sfs.Root, name)

		// @ Check for Path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, errors.New("User File Does Not Exist")
		} else {
			// @ Read User Data File
			dat, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}
			return dat, nil
		}
	} else {
		return nil, errors.New("Sonr FileSystem not Initialized")
	}

}

// ^ WriteIncomingFile writes file to Disk ^
func (sfs *FileSystem) WriteFile(name string, data []byte) (string, error) {
	if sfs.Initialized {
		// Create File Path
		path := filepath.Join(sfs.Root, name)

		// Check for User File at Path
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}

		// Defer Close
		defer file.Close()

		// Write User Data to File
		_, err = file.Write(data)
		if err != nil {
			return "", err
		}
		return path, nil
	} else {
		return "", errors.New("Sonr FileSystem not Initialized")
	}

}

// ^ WriteIncomingFile writes file to Disk ^
func (sfs *FileSystem) WriteIncomingFile(load md.Payload, props *md.TransferCard_Properties, data []byte) (string, string) {
	// Create File Name
	fileName := props.Name + "." + props.Mime.Subtype
	path := sfs.getIncomingFilePath(load, fileName)

	// Check for User File at Path
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Defer Close
	defer file.Close()

	// Write User Data to File
	_, err = file.Write(data)
	if err != nil {
		sentry.CaptureException(err)
	}
	return fileName, path
}

// @ Helper: Finds Write Path for Incoming File
func (sfs *FileSystem) getIncomingFilePath(load md.Payload, fileName string) string {
	// Check for Desktop
	if sfs.IsDesktop {
		return filepath.Join(sfs.Downloads, fileName)
	} else {
		// Check Load
		if load == md.Payload_MEDIA {
			return filepath.Join(sfs.Temporary, fileName)
		} else {
			return filepath.Join(sfs.Root, fileName)
		}
	}
}
