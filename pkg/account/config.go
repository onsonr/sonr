package account

import (
	"os"
	"path"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Method Returns Account KeyPair
func (al *userLinker) AccountKeys() *md.KeyPair {

	return al.user.GetKeyChain().GetAccount()
}

// Return Client API Keys
func (al *userLinker) APIKeys() *md.APIKeys {
	return al.user.GetApiKeys()
}

// Method Returns Current Device
func (al *userLinker) CurrentDevice() *md.Device {
	return al.currentDevice
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

// Method IsReady checks if Active Device Data is ready
func (al *userLinker) IsReady() bool {
	return al.CurrentDevice().IsReady()
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

// NewUpdateEvent Creates Lobby Event with Peer Data ^
func (al *userLinker) NewUpdateEvent(room *md.Room, id peer.ID) *md.RoomEvent {
	return &md.RoomEvent{
		Subject: md.RoomEvent_UPDATE,
		Member:  al.Member(),
		Id:      id.String(),
		Room:    room,
	}
}

// NewDefaultUpdateEvent Updates Peer with Default Position and Returns Lobby Event with Peer Data ^
func (al *userLinker) NewDefaultUpdateEvent(room *md.Room, id peer.ID) *md.RoomEvent {
	// Update Peer
	al.currentDevice.UpdatePosition(md.DefaultPosition().Parameters())

	// Check if User is Linker
	if al.currentDevice.GetStatus() == md.Status_LINKER {
		// Return Event
		return &md.RoomEvent{
			Subject: md.RoomEvent_LINKER,
			Member:  al.Member(),
			Id:      id.String(),
			Room:    room,
		}
	} else {
		// Return Event
		return &md.RoomEvent{
			Subject: md.RoomEvent_UPDATE,
			Member:  al.Member(),
			Id:      id.String(),
			Room:    room,
		}
	}

}

// NewUpdateEvent Creates Lobby Event with Peer Data ^
func (al *userLinker) NewExitEvent(room *md.Room, id peer.ID) *md.RoomEvent {
	return &md.RoomEvent{
		Subject: md.RoomEvent_EXIT,
		Id:      id.String(),
		Room:    room,
	}
}
