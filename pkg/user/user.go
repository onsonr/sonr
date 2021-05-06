package user

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Sonr UserConfig Struct
type UserConfig struct {
	// Properties
	Call md.NodeCallback

	// User Data
	contact       *md.Contact
	device        *md.Device
	connectivity  md.Connectivity
	settings      map[string]*md.User_Settings
	user          *md.User
	hasPrivateKey bool
	privateKey    crypto.PrivKey
}

// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *md.ConnectionRequest, callback md.NodeCallback) *UserConfig {
	// Initialize
	hasPrivateKey := true
	device := cr.GetDevice()

	// Get Private Key
	privKey, err := device.PrivateKey()
	if err != nil {
		hasPrivateKey = false
	}

	// @ Return
	return &UserConfig{
		Call:          callback,
		contact:       cr.GetContact(),
		device:        cr.GetDevice(),
		connectivity:  cr.Connectivity,
		privateKey:    privKey,
		hasPrivateKey: hasPrivateKey,
	}
}

// ^ Get Peer returns Users Contact ^ //
func (u *UserConfig) Contact() *md.Contact {
	return u.contact
}

// ^ Get Peer returns Users Current device ^ //
func (u *UserConfig) Device() *md.Device {
	return u.device
}

// ^ Return User Host Options ^ //
func (u *UserConfig) PrivateKey() crypto.PrivKey {
	return u.privateKey
}
