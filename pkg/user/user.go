package user

import (
	"google.golang.org/protobuf/proto"

	"github.com/libp2p/go-libp2p-core/crypto"

	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"
const K_SONR_PRIV_KEY = "snr-peer.privkey"

// @ Sonr User Struct
type User struct {
	// Properties
	Call md.NodeCallback
	FS   *FileSystem

	// User Data
	contact  *md.Contact
	device   *md.Device
	devices  []*md.Device
	protoRef *md.User
}

// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *md.ConnectionRequest, callback md.NodeCallback) (*User, error) {
	// @ Init FileSystem
	fs, err := SetFS(cr, callback)
	if err != nil {
		return nil, err
	}

	// @ Create Devices
	devices := make([]*md.Device, 32)
	devices = append(devices, cr.Device)

	// @ Return
	return &User{
		Call:    callback,
		contact: cr.GetContact(),
		device:  cr.Device,
		devices: devices,
		FS:      fs,
	}, nil
}

// ^ Method Loads User Data from Disk ^ //
func (u *User) LoadUser() (*md.User, error) {
	// Read File
	dat, err := u.FS.ReadFile(K_SONR_USER_PATH)
	if err != nil {
		return nil, err
	}

	// Get User Data
	user := &md.User{}
	err = proto.Unmarshal(dat, user)
	if err != nil {
		return nil, err
	}

	// Set and Return
	u.protoRef = user
	return user, nil
}

// ^ Get Peer returns Users Contact ^ //
func (u *User) Contact() *md.Contact {
	return u.contact
}

// ^ Get Peer returns Users Current device ^ //
func (u *User) Device() *md.Device {
	return u.device
}

// ^ Get Key: Returns Private key from disk if found ^ //
func (u *User) PrivateKey() (crypto.PrivKey, error) {
	// @ Get Private Key
	if ok := u.FS.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, err := u.FS.ReadFile(K_SONR_PRIV_KEY)
		if err != nil {
			return nil, err
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshalling identity private key")
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, errors.Wrap(err, "generating identity private key")
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling identity private key")
		}

		// Write Key to File
		_, err = u.FS.WriteFile(K_SONR_PRIV_KEY, buf)
		if err != nil {
			return nil, err
		}

		// Set Key Ref
		return privKey, nil
	}

}
