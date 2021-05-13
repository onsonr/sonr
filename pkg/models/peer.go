package models

import (
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────
// ^ Create New Peer from Connection Request and Host ID ^ //
func NewPeer(cr *ConnectionRequest, id peer.ID, maddr multiaddr.Multiaddr) (*Peer, *SonrError) {
	// Initialize
	deviceID := cr.Device.GetId()
	c := cr.GetContact()
	profile := c.GetProfile()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return nil, NewError(err, ErrorMessage_HOST_KEY)
	}

	// Set Peer
	return &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  profile,
		Platform: cr.Device.Platform,
		Model:    cr.Device.Model,
	}, nil
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
		Event:   LobbyEvent_MESSAGE,
		From:    p,
		Id:      p.Id.Peer,
		Message: m,
		To:      to.Id.Peer,
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithContact(c *Contact, flat bool, req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		IsFlat:  flat,
		Data:    c.GetTransfer(),
		Payload: req.GetPayload(),
		Remote:  req.GetRemote(),
		To:      req.GetTo(),
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithFile(req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		To:      req.GetTo(),
		Payload: req.GetPayload(),
		Data:    req.GetData(),
		Remote:  req.GetRemote(),
	}
}

// ^ Generate AuthInvite with URL Payload from Request and User Peer Data ^ //
func (p *Peer) SignInviteWithLink(req *InviteRequest) AuthInvite {
	// Get URL Data
	link := req.GetData().GetLink()
	urlInfo, err := GetPageInfoFromUrl(link)
	if err != nil {
		urlInfo = &URLLink{
			Url: link,
		}
	}

	// Create Invite
	return AuthInvite{
		From:    p,
		Data:    urlInfo.GetTransfer(),
		Payload: req.GetPayload(),
		Remote:  req.GetRemote(),
		To:      req.GetTo(),
	}
}

// ^ SignReply Creates AuthReply ^
func (p *Peer) SignReply(d bool, req *RespondRequest, to *Peer) *AuthReply {
	return &AuthReply{
		From:     p,
		Type:     AuthReply_Transfer,
		Decision: d,
		Remote:   req.GetRemote(),
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_NONE,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    p.Profile,
			Receiver: to.GetProfile(),
		},
	}
}

// ^ SignReply Creates AuthReply with Contact  ^
func (p *Peer) SignReplyWithContact(c *Contact, flat bool, req *RespondRequest, to *Peer) *AuthReply {
	// Set Reply Type
	var kind AuthReply_Type
	if flat {
		kind = AuthReply_FlatContact
	} else {
		kind = AuthReply_Contact
	}

	// Check if Request Provided
	if req != nil {
		// Build Reply
		return &AuthReply{
			From:   p,
			Type:   kind,
			Remote: req.GetRemote(),
			Card: &TransferCard{
				// SQL Properties
				Payload:  Payload_CONTACT,
				Received: int32(time.Now().Unix()),

				// Owner Properties
				Owner:    p.Profile,
				Receiver: to.GetProfile(),

				// Data Properties
				Contact: c,
			},
		}
	} else {
		// Build Reply
		return &AuthReply{
			From: p,
			Type: kind,
			Card: &TransferCard{
				// SQL Properties
				Payload:  Payload_CONTACT,
				Received: int32(time.Now().Unix()),

				// Owner Properties
				Owner:    p.Profile,
				Receiver: to.GetProfile(),

				// Data Properties
				Contact: c,
			},
		}
	}

}

// ^ SignUpdate Creates Lobby Event with Peer Data ^
func (p *Peer) SignUpdate() *LobbyEvent {
	return &LobbyEvent{
		Event: LobbyEvent_UPDATE,
		From:  p,
		Id:    p.Id.Peer,
	}
}

// ^ Processes Update Request ^ //
func (p *Peer) Update(u *UpdateRequest) {
	if u.Type == UpdateRequest_Position {
		// Extract Data
		facing := u.Position.GetFacing()
		heading := u.Position.GetHeading()

		// Update User Values
		var faceDir float64
		var faceAnpd float64
		var headDir float64
		var headAnpd float64
		faceDir = math.Round(facing.Direction*100) / 100
		headDir = math.Round(heading.Direction*100) / 100
		faceDesg := int((facing.Direction / 11.25) + 0.25)
		headDesg := int((heading.Direction / 11.25) + 0.25)

		// Find Antipodal
		if facing.Direction > 180 {
			faceAnpd = math.Round((facing.Direction-180)*100) / 100
		} else {
			faceAnpd = math.Round((facing.Direction+180)*100) / 100
		}

		// Find Antipodal
		if heading.Direction > 180 {
			headAnpd = math.Round((heading.Direction-180)*100) / 100
		} else {
			headAnpd = math.Round((heading.Direction+180)*100) / 100
		}

		// Set Position
		p.Position = &Position{
			Facing: &Position_Compass{
				Direction: faceDir,
				Antipodal: faceAnpd,
				Cardinal:  Cardinal(faceDesg % 32),
			},
			Heading: &Position_Compass{
				Direction: headDir,
				Antipodal: headAnpd,
				Cardinal:  Cardinal(headDesg % 32),
			},
			Orientation: u.Position.GetOrientation(),
		}
	}

	// Set Properties
	if u.Type == UpdateRequest_Properties {
		p.Properties = u.Properties
	}

	// Check for New Contact, Update Peer Profile
	if u.Type == UpdateRequest_Contact {
		profile := u.Contact.GetProfile()
		p.Profile = &Profile{
			FirstName: profile.GetFirstName(),
			LastName:  profile.GetLastName(),
			Picture:   profile.GetPicture(),
		}
	}
}

// ** ─── Lobby MANAGEMENT ────────────────────────────────────────────────────────

// ^ Get Remote Point Info ^
func GetRemoteInfo(list []string) RemoteInfo {
	return RemoteInfo{
		Display: fmt.Sprintf("%s %s %s", list[0], list[1], list[2]),
		Topic:   fmt.Sprintf("%s-%s-%s", list[0], list[1], list[2]),
		Count:   int32(len(list)),
		IsJoin:  false,
		Words:   list,
	}
}

// ^ Returns as Lobby Buffer ^
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

// ^ Add/Update Peer in Lobby ^
func (l *Lobby) Add(peer *Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Remove Peer from Lobby ^
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Sync Between Remote Peers Lobby ^
func (l *Lobby) Sync(ref *Lobby, remotePeer *Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			if l.User.IsNotPeerIDString(id) {
				l.Add(peer)
			}
		}
	}
	l.Add(remotePeer)
}
