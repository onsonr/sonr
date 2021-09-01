package account

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// isEventJoin Checks if PeerEvent is Join and NOT User
func (tm *userLinker) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// isEventExit Checks if PeerEvent is Exit and NOT User
func (tm *userLinker) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// isValidMessage Checks if Message is NOT from User
func (tm *userLinker) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}

// loadAccount Method Loads Account from Disk
func loadLinker(ir *data.InitializeRequest, d *data.Device, kc *data.KeyChain) (*userLinker, *data.SonrError) {
	// Load Account
	buf, err := d.GetFileSystem().GetSupport().ReadFile(util.ACCOUNT_FILE)
	if err != nil {
		return nil, err
	}

	// Unmarshal Account
	loadedAccount := &data.User{}
	serr := proto.Unmarshal(buf, loadedAccount)
	if serr != nil {
		return nil, data.NewError(serr, data.ErrorEvent_ACCOUNT_LOAD)
	}

	data.LogInfo(fmt.Sprintf("LoadedAccount: %s", loadedAccount.String()))

	// Set Account
	loadedAccount.KeyChain = kc
	loadedAccount.Current = d
	loadedAccount.ApiKeys = ir.GetApiKeys()

	// Create Account Linker
	linker := &userLinker{
		user: loadedAccount,
		room: loadedAccount.NewDeviceRoom(),
	}
	return linker, nil
}

// loadAccount Method Loads Account from Disk
func newLinker(ir *data.InitializeRequest, d *data.Device, kc *data.KeyChain) (*userLinker, *data.SonrError) {
	// Return User
	u := &data.User{
		KeyChain: kc,
		Current:  d,
		ApiKeys:  ir.GetApiKeys(),
		Devices:  make([]*data.Device, 0),
		Member: &data.Member{
			Reach:      data.Member_ONLINE,
			Associated: make([]*data.Peer, 0),
		},
	}

	// Create Account Linker
	linker := &userLinker{
		user: u,
		room: u.NewDeviceRoom(),
	}
	err := linker.Save()
	if err != nil {
		return nil, data.NewError(err, data.ErrorEvent_ACCOUNT_SAVE)
	}
	return linker, nil
}

// SetAvailable Method Sets Account to be Available
func (al *userLinker) SetAvailable(val bool) *data.StatusEvent {
	return al.user.GetCurrent().SetAvailable(val)
}

// SetConnected Method Sets Account to be Connected
func (al *userLinker) SetConnected(val bool) *data.StatusEvent {
	return al.user.GetCurrent().SetConnected(val)
}

// SetStatus Method Updates Status of Account
func (al *userLinker) SetStatus(newStatus data.Status) *data.StatusEvent {
	return al.user.GetCurrent().SetStatus(newStatus)
}

// SignLinkPacket Method Signs Packet with Keys
func (al *userLinker) SignLinkPacket(resp *data.LinkResponse) *data.LinkPacket {
	u := al.user
	return &data.LinkPacket{
		Primary:   u.GetPrimary(),
		Secondary: resp.GetDevice(),
		KeyChain:  al.ExportKeychain(),
	}
}

// VerifyDevicePubKey Method Verifies the Device Link Public Key
func (al *userLinker) VerifyDevicePubKey(pub crypto.PubKey) bool {
	u := al.user
	return u.GetKeyChain().Device.VerifyPubKey(pub)
}

// VerifyRead Method Returns Keychain Info to Client
func (al *userLinker) VerifyRead() *data.VerifyResponse {
	u := al.user
	kp := u.GetKeyChain().GetAccount()
	return &data.VerifyResponse{
		PublicKey: kp.PubKeyBase64(),
		ShortID:   u.GetCurrent().ShortID(),
	}
}
