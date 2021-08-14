package account

import (
	"context"
	"os"
	"path"

	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Set the User with ConnectionRequest
func (al *accountLinker) SetConnection(cr *md.ConnectionRequest) {
	u := al.account
	// Initialize Account Params
	u.PushToken = cr.GetPushToken()
	u.SName = cr.GetContact().GetProfile().GetSName()
	u.Contact = cr.GetContact()
	u.Member.PushToken = cr.GetPushToken()
	al.Save()

	// Initialize Linker Params
	al.ctx = context.Background()
	al.room = al.account.NewDeviceRoom()
	al.activeDevices = make([]*md.Device, 0)
	al.syncEvents = make(chan *md.SyncEvent)
}

// Method Returns Account KeyPair
func (al *accountLinker) AccountKeys() *md.KeyPair {
	u := al.account
	return u.GetKeyChain().GetAccount()
}

// Return Client API Keys
func (al *accountLinker) APIKeys() *md.APIKeys {
	u := al.account
	return u.GetApiKeys()
}

// Method Returns DeviceID
func (al *accountLinker) DeviceID() string {
	u := al.account
	return u.GetCurrent().GetId()
}

// Method Returns Device KeyPair
func (al *accountLinker) DeviceKeys() *md.KeyPair {
	u := al.account
	return u.GetKeyChain().GetDevice()
}

// Method Returns Device Link Public Key
func (al *accountLinker) DevicePubKey() *md.KeyPair_Public {
	u := al.account
	return u.GetKeyChain().GetDevice().GetPublic()
}

// Method Returns support directory file for account
func (al *accountLinker) FilePath() string {
	u := al.account
	return path.Join(u.GetCurrent().GetFileSystem().GetSupport().GetPath(), util.ACCOUNT_FILE)
}

// Method Returns Exportable Keychain for Linked Devices
func (al *accountLinker) ExportKeychain() *md.KeyChain {
	return &md.KeyChain{
		Account: al.AccountKeys(),
		Device:  al.DeviceKeys(),
		Group:   al.GroupKeys(),
	}
}

// Method Returns Profile First Name
func (al *accountLinker) FirstName() string {
	u := al.account
	return u.GetContact().GetProfile().GetFirstName()
}

// Method Returns Group KeyPair
func (al *accountLinker) GroupKeys() *md.KeyPair {
	u := al.account
	return u.GetKeyChain().GetGroup()
}

// Method Returns Profile Last Name
func (al *accountLinker) LastName() string {
	u := al.account
	return u.GetContact().GetProfile().GetLastName()
}

// Method Returns Member
func (al *accountLinker) Member() *md.Member {
	u := al.account
	return u.GetMember()
}

// Method Returns Profile
func (al *accountLinker) Profile() *md.Profile {
	u := al.account
	return u.GetContact().GetProfile()
}

func (al *accountLinker) Save() error {
	u := al.account
	// Marshal Account to Protobuf
	data, err := proto.Marshal(u)
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
