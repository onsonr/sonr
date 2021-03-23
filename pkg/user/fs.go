package data

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/getsentry/sentry-go"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/pkg/errors"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
)

// @ Constant Variables
const K_SONR_CLIENT_DIR = ".sonr"

// @ Sonr File System Struct
type SonrFS struct {
	// Properties
	Initialized bool
	Devices     []*md.Device
	User        *md.User
	IsDesktop   bool

	// Directories
	Downloads string
	Root      string
	Temporary string

	// Queue
	Files        []*sf.ProcessedFile
	CurrentCount int
	Call         md.FileCallback
	Profile      *md.Profile
}

// ^ Method Initializes Root Sonr Directory ^ //
func InitFS(connEvent *md.ConnectionRequest, profile *md.Profile, callback md.FileCallback) *SonrFS {
	// Initialize
	var sonrPath string
	var hasInitialized bool
	devices := make([]*md.Device, 32)
	devices = append(devices, connEvent.Device)
	user := &md.User{
		Contact: connEvent.Contact,
		Profile: profile,
		Devices: devices,
	}

	// Check for Client Type
	if connEvent.Device.IsDesktop {
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
	sfs := &SonrFS{
		IsDesktop:   connEvent.Device.GetIsDesktop(),
		Initialized: hasInitialized,
		Downloads:   connEvent.Directories.Downloads,
		Root:        sonrPath,
		Temporary:   connEvent.Directories.Temporary,

		Files:        make([]*sf.ProcessedFile, maxFileBufferSize),
		CurrentCount: 0,
		Call:         callback,
		Profile:      profile,
	}

	// Write User
	if hasInitialized {
		if err := sfs.WriteUser(user); err != nil {
			sentry.CaptureException(err)
		}
	}
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

// ^ Get Key: Returns Private key from disk if found ^ //
func (fs *SonrFS) GetPrivateKey() (crypto.PrivKey, error) {
	if fs.Initialized {
		// @ Set Path
		privKeyFileName := filepath.Join(fs.Root, "snr-peer.privkey")
		var generate bool

		// @ Find Key
		privKeyBytes, err := ioutil.ReadFile(privKeyFileName)
		if os.IsNotExist(err) {
			generate = true
		} else if err != nil {
			sentry.CaptureException(err)
			return nil, err
		}

		// @ Check for Generate
		if generate {
			// Create New Key
			privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
			if err != nil {
				sentry.CaptureException(err)
				return nil, errors.Wrap(err, "generating identity private key")
			}

			// Marshal Data
			privKeyBytes, err := crypto.MarshalPrivateKey(privKey)
			if err != nil {
				sentry.CaptureException(err)
				return nil, errors.Wrap(err, "marshalling identity private key")
			}

			// Create File
			f, err := os.Create(privKeyFileName)
			if err != nil {
				sentry.CaptureException(err)
				return nil, errors.Wrap(err, "creating identity private key file")
			}
			defer f.Close()

			// Write Key to file
			if _, err := f.Write(privKeyBytes); err != nil {
				sentry.CaptureException(err)
				return nil, errors.Wrap(err, "writing identity private key to file")
			}
			return privKey, nil
		}

		privKey, err := crypto.UnmarshalPrivateKey(privKeyBytes)
		if err != nil {
			sentry.CaptureException(err)
			return nil, errors.Wrap(err, "unmarshalling identity private key")
		}
		return privKey, nil
	} else {
		return nil, errors.New("FileSystem not initialized.")
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
