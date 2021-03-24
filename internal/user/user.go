package user

import (
	"google.golang.org/protobuf/proto"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"

// @ Sonr User Struct
type User struct {
	// Inherited
	Call dt.NodeCallback

	// User Properties
	Contact *md.Contact
	Device  *md.Device
	Devices []*md.Device
	Peer    *md.Peer
	Profile *md.Profile
	User    *md.User

	// Data
	FS *sf.FileSystem
}

// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *md.ConnectionRequest, callback dt.NodeCallback) *User {
	// Create Devices
	devices := make([]*md.Device, 32)
	devices = append(devices, cr.Device)

	// Create User
	user := &User{
		Call:    callback,
		Contact: cr.GetContact(),
		Device:  cr.Device,
		Devices: devices,
		//Peer: *md.Peer,
		Profile: cr.GetProfile(),
	}

	// Init FileSystem
	user.FS = sf.InitFS(cr, user.Profile, callback)
	return user
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
	u.User = user
	return user, nil
}
