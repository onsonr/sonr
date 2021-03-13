package data

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Constant Variables
const K_SONR_ROOT_DIR = ".sonr"


// @ Sonr File System Struct
type SonrFS struct {
	Devices   []*md.Device
	User      *md.User
	Cache     string
	Documents string
	Downloads string
	Home      string
	Root      string
	Temporary string
}

// ^ Method Initializes Root Sonr Directory ^ //
func InitFS(connEvent *md.ConnectionRequest, profile *md.Profile) *SonrFS {
	// Initialize
	var sonrPath string
	devices := make([]*md.Device, 32)
	devices = append(devices, connEvent.Device)
	user := &md.User{
		Contact: connEvent.Contact,
		Profile: profile,
		Devices: devices,
	}

	// Check for Client Type
	if connEvent.Device.Desktop {
		// Init Path, Check for Path
		sonrPath = filepath.Join(connEvent.Directories.Home, K_SONR_ROOT_DIR)
		if err := EnsureDir(sonrPath, os.ModeAppend); err != nil {
			sentry.CaptureException(err)
		}
	} else {
		// Init Path, Check for Path
		sonrPath = filepath.Join(connEvent.Directories.Documents, K_SONR_ROOT_DIR)
		if err := EnsureDir(sonrPath, os.ModeAppend); err != nil {
			sentry.CaptureException(err)
		}
	}

	// Create SFS
	sfs := &SonrFS{
		Cache:     connEvent.Directories.Cache,
		Documents: connEvent.Directories.Documents,
		Downloads: connEvent.Directories.Downloads,
		Home:      connEvent.Directories.Home,
		Root:      sonrPath,
		Temporary: connEvent.Directories.Temporary,
	}

	// Write User

	sfs.WriteUser(user)
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

// ^ Write User Data at Path ^
func (sfs *SonrFS) WriteUser(user *md.User) error {
	userPath := filepath.Join(sfs.Root, K_SONR_USER_PATH)

	// Convert User to Bytes
	userData, err := proto.Marshal(user)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Check for User File at Path
	file, err := os.OpenFile(userPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Defer Close
	defer file.Close()

	// Write User Data to File
	_, err = file.Write(userData)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}
