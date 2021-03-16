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
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Constant Variables
const K_SONR_CLIENT_DIR = ".sonr"

// @ Sonr File System Struct
type SonrFS struct {
	// Properties
	Initialized bool
	Devices     []*md.Device
	User        *md.User

	// Directories
	Cache     string
	Documents string
	Downloads string
	Home      string
	Root      string
	Temporary string

	// Queue
	files        []*sf.ProcessedFile
	currentCount int
	directories  *md.Directories
	call         md.FileCallback
	profile      *md.Profile
}

// ^ Method Initializes Root Sonr Directory ^ //
func InitFS(connEvent *md.ConnectionRequest, profile *md.Profile, qc md.OnQueued, mqc md.OnMultiQueued, ec md.OnError) *SonrFS {
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
	if connEvent.Device.Desktop {
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
		sonrPath = connEvent.Directories.Documents
	}

	// Create SFS
	sfs := &SonrFS{
		Initialized: hasInitialized,
		Cache:       connEvent.Directories.Cache,
		Documents:   connEvent.Directories.Documents,
		Downloads:   connEvent.Directories.Downloads,
		Home:        connEvent.Directories.Home,
		Root:        sonrPath,
		Temporary:   connEvent.Directories.Temporary,
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
func (sfs *SonrFS) WriteFile(load md.Payload, props *md.TransferCard_Properties, data []byte) (string, string) {
	// Create File Name
	fileName := props.Name + "." + props.Mime.Subtype
	var path string

	// Check Load
	if load == md.Payload_MEDIA {
		path = filepath.Join(sfs.Temporary, fileName)
	} else {
		path = filepath.Join(sfs.Root, fileName)
	}

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

// ^ Write User Data at Path ^
func (sfs *SonrFS) WriteUser(user *md.User) error {
	userPath := filepath.Join(sfs.Root, K_SONR_USER_PATH)

	// Convert User to Bytes
	userData, err := proto.Marshal(user)
	if err != nil {
		return err
	}

	// Check for User File at Path
	file, err := os.OpenFile(userPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Defer Close
	defer file.Close()

	// Write User Data to File
	_, err = file.Write(userData)
	if err != nil {
		return err
	}
	return nil
}
