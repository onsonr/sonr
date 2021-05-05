package user

import (
	"fmt"
	"os"
	"path/filepath"

	crypto "github.com/libp2p/go-libp2p-crypto"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"
const K_SONR_PRIV_KEY = "snr-peer.privkey"
const K_FILE_QUEUE_NAME = "file-queue"

// @ Sonr File System Struct
type FileSystem struct {
	// Properties
	Call        md.NodeCallback
	Device      *md.Device
	Directories *md.Directories
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
	// Initialize
	var path string

	// Create File Path
	if sfs.Device.IsDesktop() {
		path = filepath.Join(sfs.Directories.GetLibrary(), name)
	} else {
		path = filepath.Join(sfs.Directories.GetDocuments(), name)
	}

	// Check Path
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
func (sfs *FileSystem) ReadFile(name string) ([]byte, *md.SonrError) {
	// Initialize
	var path string

	// Create File Path
	if sfs.Device.IsDesktop() {
		path = filepath.Join(sfs.Directories.GetLibrary(), name)
	} else {
		path = filepath.Join(sfs.Directories.GetDocuments(), name)
	}

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, md.NewError(err, md.ErrorMessage_USER_LOAD)
	} else {
		// @ Read User Data File
		dat, err := os.ReadFile(path)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_USER_LOAD)
		}
		return dat, nil
	}
}

// ^ WriteIncomingFile writes file to Disk ^
func (sfs *FileSystem) WriteFile(name string, data []byte) (string, *md.SonrError) {
	// Create File Path
	path := sfs.Directories.DataSavePath(name, sfs.Device.IsDesktop())

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", md.NewError(err, md.ErrorMessage_USER_FS)
	}
	return path, nil
}

// @ Helper: Finds Write Path for Incoming File
func (sfs *FileSystem) GetPathForPayload(file *md.SonrFile) string {
	// Check count
	if file.IsSingle() {
		f := file.SingleFile()
		return sfs.Directories.TransferSavePath(f.Name, f.Mime, sfs.Device.IsDesktop())
	} else {
		return ""
	}
}

// @ Get Key: Returns Private key from disk if found ^ //
func (sfs *FileSystem) getPrivateKey() (crypto.PrivKey, *md.SonrError) {
	// @ Get Private Key
	if ok := sfs.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, serr := sfs.ReadFile(K_SONR_PRIV_KEY)
		if serr != nil {
			return nil, serr
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_KEY)
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_KEY)
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_MARSHAL)
		}

		// Write Key to File
		_, werr := sfs.WriteFile(K_SONR_PRIV_KEY, buf)
		if werr != nil {
			return nil, md.NewError(err, md.ErrorMessage_USER_SAVE)
		}

		// Set Key Ref
		return privKey, nil
	}
}
