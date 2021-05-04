package user

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
)

type UserConfig struct {
	connectivity  md.Connectivity
	isDesktop     bool
	hasPrivateKey bool
	privateKey    crypto.PrivKey
	fileSystem    *FileSystem
	user          *md.User
	settings      map[string]*md.User_Settings
}

// ^ Initialize User ^ //
func InitUserConfig(req *md.ConnectionRequest, cb md.NodeCallback) *UserConfig {
	// Initialize
	var sonrPath string
	var hasPrivateKey bool

	// Check for Client Type
	if req.Device.GetIsDesktop() {
		// Init Path, Check for Path
		sonrPath = req.Directories.Library
		if err := EnsureDir(req.Directories.Library, 0755); err != nil {
			cb.Error(md.NewError(err, md.ErrorMessage_USER_FS))
			return nil
		}
	} else {
		// Set Path to Documents for Mobile
		sonrPath = req.Directories.Support
	}

	// Set File System
	fs := &FileSystem{
		IsDesktop: req.Device.GetIsDesktop(),
		Downloads: req.Directories.Downloads,
		Main:      sonrPath,
		Temporary: req.Directories.Temporary,
		Call:      cb,
	}

	// Get Private Key
	privKey, err := fs.getPrivateKey()
	if err != nil {
		hasPrivateKey = false
	} else {
		hasPrivateKey = true
	}

	// Build Config
	return &UserConfig{
		connectivity:  req.Connectivity,
		isDesktop:     req.Device.IsDesktop,
		fileSystem:    fs,
		privateKey:    privKey,
		hasPrivateKey: hasPrivateKey,
	}
}
