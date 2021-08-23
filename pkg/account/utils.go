package account

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	ps "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
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

// Method IsReady checks if Active Device Data is ready
func (al *userLinker) IsReady() bool {
	return al.CurrentDevice().IsReady()
}

// SetAvailable Method Sets Account to be Available
func (al *userLinker) SetAvailable(val bool) *md.StatusEvent {
	return al.currentDevice.SetAvailable(val)
}

// SetConnected Method Sets Account to be Connected
func (al *userLinker) SetConnected(val bool) *md.StatusEvent {
	return al.currentDevice.SetConnected(val)
}

// SetStatus Method Updates Status of Account
func (al *userLinker) SetStatus(newStatus md.Status) *md.StatusEvent {
	return al.currentDevice.SetStatus(newStatus)
}

// SignLinkPacket Method Signs Packet with Keys
func (al *userLinker) SignLinkPacket(resp *md.LinkResponse) *md.LinkPacket {
	u := al.user
	return &md.LinkPacket{
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
func (al *userLinker) VerifyRead() *md.VerifyResponse {
	u := al.user
	kp := u.GetKeyChain().GetAccount()
	return &md.VerifyResponse{
		PublicKey: kp.PubKeyBase64(),
		ShortID:   u.GetCurrent().ShortID(),
	}
}
