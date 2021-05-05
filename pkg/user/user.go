package user

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Sonr User Struct
type User struct {
	config *UserConfig
	// Properties
	Call md.NodeCallback

	// User Data
	contact *md.Contact
	device  *md.Device
}

// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *md.ConnectionRequest, callback md.NodeCallback) *User {
	// @ Return
	return &User{
		Call:    callback,
		contact: cr.GetContact(),
		device:  cr.GetDevice(),
		config:  InitUserConfig(cr, callback),
	}
}

// ^ Get Peer returns Users Contact ^ //
func (u *User) Contact() *md.Contact {
	return u.contact
}

// ^ Get Peer returns Users Current device ^ //
func (u *User) Device() *md.Device {
	return u.device
}

// ^ Return User FileSystem ^ //
func (u *User) FileSystem() *FileSystem {
	return u.config.fileSystem
}

// ^ Return User Host Options ^ //
func (u *User) PrivateKey() crypto.PrivKey {
	return u.config.privateKey
}
