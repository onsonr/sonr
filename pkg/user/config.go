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
	var hasPrivateKey bool

	// Set File System
	fs := &FileSystem{
		Call:        cb,
		Device:      req.Device,
		Directories: req.Directories,
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
		isDesktop:     req.Device.IsDesktop(),
		fileSystem:    fs,
		privateKey:    privKey,
		hasPrivateKey: hasPrivateKey,
	}
}
