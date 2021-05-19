package models

import (
	"fmt"
	"hash/fnv"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────
// ^ Create New Peer from Connection Request and Host ID ^ //
func (u *User) NewPeer(id peer.ID, maddr multiaddr.Multiaddr) *SonrError {
	// Initialize
	deviceID := u.Device.GetId()
	c := u.GetContact()
	profile := c.GetProfile()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return NewError(err, ErrorMessage_HOST_KEY)
	}

	// Set Peer
	u.Peer = &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  profile,
		Platform: u.Device.Platform,
		Model:    u.Device.Model,
	}
	// Set Device Topic
	u.Connection.Router.DeviceTopic = fmt.Sprintf("/sonr/topic/%s", u.Peer.UserID())
	return nil
}

// ^ Returns Peer as Buffer ^ //
func (p *Peer) Buffer() ([]byte, error) {
	buf, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// ^ Returns Peer User ID ^ //
func (p *Peer) DeviceID() string {
	return string(p.Id.GetDevice())
}

// ^ Returns Peer User ID ^ //
func (p *Peer) UserID() string {
	return fmt.Sprintf("%d", p.Id.GetUser())
}

// ^ Checks for Host Peer ID is Same ^ //
func (p *Peer) IsPeerID(pid peer.ID) bool {
	return p.Id.Peer == pid.String()
}

// ^ Checks for Host Peer ID String is Same ^ //
func (p *Peer) IsPeerIDString(pid string) bool {
	return p.Id.Peer == pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (p *Peer) IsNotPeerIDString(pid string) bool {
	return p.Id.Peer != pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// ^ SignMessage Creates Lobby Event with Message ^
func (p *Peer) SignMessage(m string, to *Peer) *LobbyEvent {
	return &LobbyEvent{
		Event: &LobbyEvent_Local{
			Local: LobbyEvent_MESSAGE,
		},
		From:    p,
		Id:      p.Id.Peer,
		Message: m,
		To:      to.Id.Peer,
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (u *User) SignInviteWithContact(req *InviteRequest, isFlat bool) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    u.GetPeer(),
		IsFlat:  isFlat,
		Data:    u.Contact.GetTransfer(),
		Payload: req.GetPayload(),
		To:      req.GetTo(),
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (u *User) SignInviteWithFile(req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    u.GetPeer(),
		To:      req.GetTo(),
		Payload: req.GetPayload(),
		Data:    req.GetData(),
	}
}

// ^ Generate AuthInvite with URL Payload from Request and User Peer Data ^ //
func (u *User) SignInviteWithLink(req *InviteRequest) AuthInvite {
	// Get URL Data
	link := req.GetData().GetUrl()
	link.SetData()

	// Create Invite
	return AuthInvite{
		From:    u.GetPeer(),
		Data:    link.GetTransfer(),
		Payload: req.GetPayload(),
		To:      req.GetTo(),
	}
}

// ^ SignReply Creates AuthReply ^
func (u *User) SignReply(req *RespondRequest) *AuthReply {
	return &AuthReply{
		From:     u.GetPeer(),
		Type:     AuthReply_Transfer,
		Decision: req.GetDecision(),
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_NONE,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    u.GetPeer().Profile,
			Receiver: req.To.GetProfile(),
		},
	}
}

// ^ SignReplyFlat Creates AuthReply with Contact for Flat Mode  ^
func (u *User) SignReplyWithFlat(from *Peer) *AuthReply {
	return &AuthReply{
		From: u.GetPeer(),
		Type: AuthReply_FlatContact,
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    u.GetPeer().Profile,
			Receiver: from.GetProfile(),

			// Data Properties
			Contact: u.GetContact(),
		},
	}
}

// ^ SignReply Creates AuthReply with Contact  ^
func (u *User) SignReplyWithContact(req *RespondRequest) *AuthReply {
	return &AuthReply{
		From: u.GetPeer(),
		Type: AuthReply_Contact,
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    u.GetPeer().Profile,
			Receiver: req.To.GetProfile(),

			// Data Properties
			Contact: u.GetContact(),
		},
	}
}

// ^ SignUpdate Creates Lobby Event with Peer Data ^
func (p *Peer) SignUpdate() *LobbyEvent {
	return &LobbyEvent{
		Event: &LobbyEvent_Local{
			Local: LobbyEvent_UPDATE,
		},
		From: p,
		Id:   p.Id.Peer,
	}
}
