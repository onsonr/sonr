package user

import (
	"google.golang.org/protobuf/proto"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/pkg/errors"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/internal/network"
	dt "github.com/sonr-io/core/pkg/data"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"
const K_SONR_PRIV_KEY = "snr-peer.privkey"

// @ Sonr User Struct
type User struct {
	// Properties
	Call dt.NodeCallback
	FS   *sf.FileSystem

	// User Data
	contact  *md.Contact
	device   *md.Device
	devices  []*md.Device
	peer     *md.Peer
	privKey  crypto.PrivKey
	profile  *md.Profile
	protoRef *md.User
}

// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *md.ConnectionRequest, callback dt.NodeCallback) (*User, error) {
	// @ Init FileSystem
	fs, err := sf.NewFs(cr, callback)
	if err != nil {
		return nil, err
	}

	// @ Get Private Key
	var privKey crypto.PrivKey
	if ok := fs.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, err := fs.ReadFile(K_SONR_PRIV_KEY)
		if err != nil {
			return nil, err
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshalling identity private key")
		}

		// Set Key Ref
		privKey = key
	} else {
		// Create New Key
		key, buf, err := network.Ed25519KeyBuf()
		if err != nil {
			return nil, err
		}

		// Write Key to File
		_, err = fs.WriteFile(K_SONR_PRIV_KEY, buf)
		if err != nil {
			return nil, err
		}

		// Set Key Ref
		privKey = key
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
		//Peer: *md.Peer,
		profile: cr.GetProfile(),
		FS:      fs,
		privKey: privKey,
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
