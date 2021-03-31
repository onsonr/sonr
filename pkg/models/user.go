package models

import (
	"hash/fnv"
	"log"
	"math"
	"time"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ^ Create New Peer from Connection Request and Host ID ^ //
func NewPeer(cr *ConnectionRequest, id peer.ID) (*Peer, error) {
	// Initialize
	deviceID := cr.Device.GetId()
	profile := Profile{
		Username:  cr.GetUsername(),
		FirstName: cr.Contact.GetFirstName(),
		LastName:  cr.Contact.GetLastName(),
		Picture:   cr.Contact.GetPicture(),
		Platform:  cr.Device.GetPlatform(),
	}

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return nil, err
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := mid.ProtectedID("Sonr"); err != nil {
			log.Println(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	// Set Peer
	return &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  &profile,
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
func (p *Peer) SignInviteWithContact(c *Contact) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		Payload: Payload_CONTACT,
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),
			Platform: p.Platform,

			// Transfer Properties
			Status: TransferCard_INVITE,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,

			// Data Properties
			Contact: c,
		},
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithFile(tc *TransferCard) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		Payload: tc.Payload,
		Card:    tc,
	}
}

// ^ Generate AuthInvite with URL Payload from Request and User Peer Data ^ //
func (p *Peer) SignInviteWithLink(req *InviteRequest) AuthInvite {
	// Get URL Data
	urlInfo, err := GetPageInfoFromUrl(req.Url)
	if err != nil {
		urlInfo = &URLLink{
			Link: req.Url,
		}
	}

	// Create Invite
	return AuthInvite{
		From:    p,
		Payload: Payload_CONTACT,
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),
			Platform: p.Platform,

			// Transfer Properties
			Status: TransferCard_INVITE,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,

			// Data Properties
			Url: urlInfo,
		},
	}
}

// ^ SignReply Creates AuthReply ^
func (p *Peer) SignReply(d bool) *AuthReply {
	return &AuthReply{
		From:     p,
		Type:     AuthReply_Transfer,
		Decision: d,
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_UNDEFINED,
			Received: int32(time.Now().Unix()),
			Preview:  p.Profile.Picture,
			Platform: p.Platform,

			// Transfer Properties
			Status: TransferCard_REPLY,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,
		},
	}
}

// ^ SignReply Creates AuthReply with Contact  ^
func (p *Peer) SignReplyWithContact(c *Contact) *AuthReply {
	return &AuthReply{
		From: p,
		Type: AuthReply_Contact,
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),
			Preview:  p.Profile.Picture,
			Platform: p.Platform,

			// Transfer Properties
			Status: TransferCard_REPLY,

			// Owner Properties
			Username:  p.Profile.Username,
			FirstName: p.Profile.FirstName,
			LastName:  p.Profile.LastName,

			// Data Properties
			Contact: c,
		},
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
	if u.Type == UpdateRequest_Direction {
		// Extract Data
		facing := u.Facing
		heading := u.Heading

		// Update User Values
		var faceDir float64
		var faceAnpd float64
		var headDir float64
		var headAnpd float64
		faceDir = math.Round(facing*100) / 100
		headDir = math.Round(heading*100) / 100
		desg := int((facing / 11.25) + 0.25)

		// Find Antipodal
		if facing > 180 {
			faceAnpd = math.Round((facing-180)*100) / 100
		} else {
			faceAnpd = math.Round((facing+180)*100) / 100
		}

		// Find Antipodal
		if heading > 180 {
			headAnpd = math.Round((heading-180)*100) / 100
		} else {
			headAnpd = math.Round((heading+180)*100) / 100
		}

		// Set Position
		p.Position = &Position{
			Facing:           faceDir,
			FacingAntipodal:  faceAnpd,
			Heading:          headDir,
			HeadingAntipodal: headAnpd,
			Designation:      Position_Designation(desg % 32),
		}
	}

	// Set Properties
	if u.Type == UpdateRequest_Properties {
		p.Properties = &Peer_Properties{
			IsFlatMode:      u.GetIsFlatMode(),
			HasPointToShare: u.GetHasPointToShare(),
		}
	}

	// Check for New Contact, Update Peer Profile
	if u.Type == UpdateRequest_Contact {
		p.Profile = &Profile{
			FirstName: u.Contact.GetFirstName(),
			LastName:  u.Contact.GetLastName(),
			Picture:   u.Contact.GetPicture(),
		}
	}
}
