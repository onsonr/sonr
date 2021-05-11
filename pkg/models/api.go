package models

import (
	"errors"
	"fmt"
	"hash/fnv"
	"math"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── ConnectionRequest MANAGEMENT ────────────────────────────────────────────────────────
func (req *ConnectionRequest) NewRouter() *Router {
	return &Router{
		Location:     req.GetLocation(),
		Connectivity: req.GetConnectivity(),
	}
}

// ************************** //
// ** MIME Info Management ** //
// ************************** //
// Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// ** ─── PEER MANAGEMENT ────────────────────────────────────────────────────────

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
		Contact: c,
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
		File:    req.GetFile(),
		Remote:  req.GetRemote(),
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
		Url:     urlInfo,
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

// ** ─── Location MANAGEMENT ────────────────────────────────────────────────────────
func (l *Location) MinorOLC() string {
	lat := l.Latitude()
	lon := l.Longitude()
	return olc.Encode(lat, lon, 6)
}

func (l *Location) MajorOLC() string {
	lat := l.Latitude()
	lon := l.Longitude()
	return olc.Encode(lat, lon, 4)
}

func (l *Location) Latitude() float64 {
	if l.Geo != nil {
		return l.Geo.GetLatitude()
	}
	return l.Ip.GetLatitude()
}

func (l *Location) Longitude() float64 {
	if l.Geo != nil {
		return l.Geo.GetLongitude()
	}
	return l.Ip.GetLongitude()
}

func (l *Location) GeoOLC() (string, error) {
	if l.Geo != nil {
		return "", errors.New("Geo Location doesnt exist")
	}
	return olc.Encode(float64(l.Geo.GetLatitude()), float64(l.Geo.GetLongitude()), 5), nil
}

func (l *Location) IPOLC() string {
	return olc.Encode(float64(l.Ip.GetLatitude()), float64(l.Ip.GetLongitude()), 5)
}
