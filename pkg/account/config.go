package account

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/logger"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// Method Returns Account KeyPair
func (al *userLinker) AccountKeys() *data.KeyPair {
	return al.user.GetKeyChain().GetAccount()
}

// Return Client API Keys
func (al *userLinker) APIKeys() *data.APIKeys {
	return al.user.GetApiKeys()
}

// Method Returns Current Device
func (al *userLinker) CurrentDevice() *data.Device {
	return al.user.GetCurrent()
}

// Method Returns Device KeyPair
func (al *userLinker) CurrentDeviceKeys() *data.KeyPair {
	return al.user.GetCurrent().AccountKeys()
}

// Method Returns DeviceID
func (al *userLinker) DeviceID() string {

	return al.user.GetCurrent().GetId()
}

// Method Returns Device KeyPair
func (al *userLinker) DeviceKeys() *data.KeyPair {

	return al.user.GetKeyChain().GetDevice()
}

// Method Returns Device Link Public Key
func (al *userLinker) DevicePubKey() *data.KeyPair_Public {

	return al.user.GetKeyChain().GetDevice().GetPublic()
}

// Method Returns support directory file for account
func (al *userLinker) FilePath() string {

	return path.Join(al.user.GetCurrent().GetFileSystem().GetSupport().GetPath(), util.ACCOUNT_FILE)
}

// Method Returns Exportable Keychain for Linked Devices
func (al *userLinker) ExportKeychain() *data.KeyChain {
	return &data.KeyChain{
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
func (al *userLinker) GroupKeys() *data.KeyPair {

	return al.user.GetKeyChain().GetGroup()
}

// Method Returns Profile Last Name
func (al *userLinker) LastName() string {

	return al.user.GetContact().GetProfile().GetLastName()
}

// Method Returns Member
func (al *userLinker) Member() *data.Member {
	return al.user.GetMember()
}

// Method Returns Profile
func (al *userLinker) Profile() *data.Profile {

	return al.user.GetContact().GetProfile()
}

func (al *userLinker) Save() error {

	// Marshal Account to Protobuf
	buf, err := proto.Marshal(al.user)
	if err != nil {
		return err
	}

	// Open File at Path
	f, err := os.OpenFile(al.FilePath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// Write Data to File
	_, err = f.Write(buf)
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
func (al *userLinker) NewUpdateEvent(room *data.Room, id peer.ID) *data.RoomEvent {
	return &data.RoomEvent{
		Subject: data.RoomEvent_UPDATE,
		Member:  al.Member(),
		Id:      id.String(),
		Room:    room,
	}
}

// NewDefaultUpdateEvent Updates Peer with Default Position and Returns Lobby Event with Peer Data ^
func (al *userLinker) NewDefaultUpdateEvent(room *data.Room, id peer.ID) *data.RoomEvent {
	// Update Peer
	al.user.GetCurrent().UpdatePosition(data.DefaultPosition().Parameters())

	// Check if User is Linker
	if al.user.GetCurrent().GetStatus() == data.Status_LINKER {
		// Return Event
		return &data.RoomEvent{
			Subject: data.RoomEvent_LINKER,
			Member:  al.Member(),
			Id:      id.String(),
			Room:    room,
		}
	} else {
		// Return Event
		return &data.RoomEvent{
			Subject: data.RoomEvent_UPDATE,
			Member:  al.Member(),
			Id:      id.String(),
			Room:    room,
		}
	}

}

// NewUpdateEvent Creates Lobby Event with Peer Data ^
func (al *userLinker) NewExitEvent(room *data.Room, id peer.ID) *data.RoomEvent {
	return &data.RoomEvent{
		Subject: data.RoomEvent_EXIT,
		Id:      id.String(),
		Room:    room,
	}
}

// Validates InviteRequest has From Parameter
func (al *userLinker) SignInvite(i *data.InviteRequest) *data.InviteRequest {
	// Set From
	if i.From == nil {
		i.From = al.Member()
	}

	// Convert all Thumbnails to Buffers
	if i.IsPayloadTransfer() {
		// Get File
		f := i.GetFile()
		if f != nil {
			// Convert Thumbnails to Buffers
			for _, t := range f.Items {
				if t.GetProperties().GetIsThumbPath() {
					// Fetch Buffer from Path
					buffer, err := ioutil.ReadFile(t.GetThumbPath())
					if err != nil {
						logger.Error("Failed to get buffer for thumbnail at path.", zap.Error(err))
						continue
					}

					// Set Buffer
					t.Thumbnail = &data.SFile_Item_ThumbBuffer{
						ThumbBuffer: buffer,
					}

					// Update Properties
					oldProps := t.GetProperties()
					t.Properties = &data.SFile_Item_Properties{
						IsThumbPath:  false,
						IsAudio:      oldProps.GetIsAudio(),
						IsVideo:      oldProps.GetIsVideo(),
						IsImage:      oldProps.GetIsImage(),
						HasThumbnail: oldProps.GetHasThumbnail(),
						Width:        oldProps.GetWidth(),
						Height:       oldProps.GetHeight(),
						Duration:     oldProps.GetDuration(),
					}
				}
			}
		}
	}

	// Set Type
	if i.Type == data.InviteRequest_NONE {
		i.Type = data.InviteRequest_LOCAL
	}
	return i
}
