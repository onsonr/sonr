package models

import (
	"os"
	"path/filepath"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

// ** ─── DEVICE MANAGEMENT ────────────────────────────────────────────────────────
func (d *Device) IsDesktop() bool {
	return d.Platform == Platform_MacOS || d.Platform == Platform_Linux || d.Platform == Platform_Windows
}

func (d *Device) IsMobile() bool {
	return d.Platform == Platform_IOS || d.Platform == Platform_Android
}

func (d *Device) IsIOS() bool {
	return d.Platform == Platform_IOS
}

func (d *Device) IsAndroid() bool {
	return d.Platform == Platform_Android
}

func (d *Device) IsMacOS() bool {
	return d.Platform == Platform_MacOS
}

func (d *Device) IsLinux() bool {
	return d.Platform == Platform_Linux
}

func (d *Device) IsWeb() bool {
	return d.Platform == Platform_Web
}

func (d *Device) IsWindows() bool {
	return d.Platform == Platform_Windows
}

// @ Checks if File Exists
func (d *Device) IsFile(name string) bool {
	// Initialize
	var path string

	// Create File Path
	if d.IsDesktop() {
		path = filepath.Join(d.FileSystem.GetLibrary(), name)
	} else {
		path = filepath.Join(d.FileSystem.GetDocuments(), name)
	}

	// Check Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// @ Returns Private key from disk if found
func (d *Device) PrivateKey() (crypto.PrivKey, *SonrError) {
	K_SONR_PRIV_KEY := "snr-peer.privkey"

	// Get Private Key
	if ok := d.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, serr := d.ReadFile(K_SONR_PRIV_KEY)
		if serr != nil {
			return nil, serr
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, NewError(err, ErrorMessage_HOST_KEY)
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, NewError(err, ErrorMessage_HOST_KEY)
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, NewError(err, ErrorMessage_MARSHAL)
		}

		// Write Key to File
		_, werr := d.WriteFile(K_SONR_PRIV_KEY, buf)
		if werr != nil {
			return nil, NewError(err, ErrorMessage_USER_SAVE)
		}

		// Set Key Ref
		return privKey, nil
	}
}

// Loads User File
func (d *Device) ReadFile(name string) ([]byte, *SonrError) {
	// Initialize
	var path string

	// Create File Path
	if d.IsDesktop() {
		path = filepath.Join(d.FileSystem.GetLibrary(), name)
	} else {
		path = filepath.Join(d.FileSystem.GetDocuments(), name)
	}

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, NewError(err, ErrorMessage_USER_LOAD)
	} else {
		// @ Read User Data File
		dat, err := os.ReadFile(path)
		if err != nil {
			return nil, NewError(err, ErrorMessage_USER_LOAD)
		}
		return dat, nil
	}
}

// Saves Transfer File to Disk
func (d *Device) SaveTransfer(f *SonrFile, i int, data []byte) error {
	// Get Save Path
	path := f.Files[i].SetPath(d)

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}

// Writes a File to Disk
func (d *Device) WriteFile(name string, data []byte) (string, *SonrError) {
	// Create File Path
	path := d.FileSystem.DataSavePath(name, d.IsDesktop())

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorMessage_USER_FS)
	}
	return path, nil
}
