package user

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	crypto "github.com/libp2p/go-libp2p-crypto"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"
const K_SONR_PRIV_KEY = "snr-peer.privkey"
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

// @ Get Key: Returns Private key from disk if found ^ //
func (sfs *FileSystem) getPrivateKey() (crypto.PrivKey, error) {
	// @ Get Private Key
	if ok := sfs.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, err := sfs.ReadFile(K_SONR_PRIV_KEY)
		if err != nil {
			return nil, err
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, err
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, err
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, err
		}

		// Write Key to File
		_, err = sfs.WriteFile(K_SONR_PRIV_KEY, buf)
		if err != nil {
			return nil, err
		}

		// Set Key Ref
		return privKey, nil
	}
}
