package account

import (
	"context"
	"os"
	"path"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Set the User with ConnectionRequest
func (al *userLinker) SetConnection(cr *md.ConnectionRequest) {

	// Initialize Account Params
	al.user.PushToken = cr.GetPushToken()
	al.user.SName = cr.GetContact().GetProfile().GetSName()
	al.user.Contact = cr.GetContact()
	al.user.Member.PushToken = cr.GetPushToken()
	al.Save()

	// Initialize Linker Params
	al.ctx = context.Background()
	al.room = al.user.NewDeviceRoom()
	al.activeDevices = make(map[peer.ID]*md.Device, 0)
	al.syncEvents = make(chan *md.SyncEvent)
}

// Method Returns Account KeyPair
func (al *userLinker) AccountKeys() *md.KeyPair {

	return al.user.GetKeyChain().GetAccount()
}

// Return Client API Keys
func (al *userLinker) APIKeys() *md.APIKeys {

	return al.user.GetApiKeys()
}

// Method Returns Device KeyPair
func (al *userLinker) CurrentDeviceKeys() *md.KeyPair {
	return al.currentDevice.AccountKeys()
}

// Method Returns DeviceID
func (al *userLinker) DeviceID() string {

	return al.user.GetCurrent().GetId()
}

// Method Returns Device KeyPair
func (al *userLinker) DeviceKeys() *md.KeyPair {

	return al.user.GetKeyChain().GetDevice()
}

// Method Returns Device Link Public Key
func (al *userLinker) DevicePubKey() *md.KeyPair_Public {

	return al.user.GetKeyChain().GetDevice().GetPublic()
}

// Method Returns support directory file for account
func (al *userLinker) FilePath() string {

	return path.Join(al.user.GetCurrent().GetFileSystem().GetSupport().GetPath(), util.ACCOUNT_FILE)
}

// Method Returns Exportable Keychain for Linked Devices
func (al *userLinker) ExportKeychain() *md.KeyChain {
	return &md.KeyChain{
		Account: al.AccountKeys(),
		Device:  al.DeviceKeys(),
		Group:   al.GroupKeys(),
	}
}

// Method Returns Profile First Name
func (al *userLinker) FirstName() string {

	return al.user.GetContact().GetProfile().GetFirstName()
}

// Method Returns Group KeyPair
func (al *userLinker) GroupKeys() *md.KeyPair {

	return al.user.GetKeyChain().GetGroup()
}

// Method Returns Profile Last Name
func (al *userLinker) LastName() string {

	return al.user.GetContact().GetProfile().GetLastName()
}

// Method Returns Member
func (al *userLinker) Member() *md.Member {
	return al.user.GetMember()
}

// Method Returns Profile
func (al *userLinker) Profile() *md.Profile {

	return al.user.GetContact().GetProfile()
}

func (al *userLinker) Save() error {

	// Marshal Account to Protobuf
	data, err := proto.Marshal(al.user)
	if err != nil {
		return err
	}

	// Open File at Path
	f, err := os.OpenFile(al.FilePath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// Write Data to File
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	// Close File
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
