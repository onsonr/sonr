package models

import (
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-threads/core/thread"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

var (
	NoPubKey = errors.New("Public Key not found from Peer Protobuf.")
)

// ** ─── Topic MANAGEMENT ────────────────────────────────────────────────────────
func (t *Topic) IsLocal() bool {
	return t.Type == Topic_LOCAL
}

// Local Lobby Topic Protocol ID
func (r *User) NewLocalTopic(opts *ConnectionRequest_ServiceOptions) *Topic {
	// Initialize Set OLC Range
	scope := 6
	if opts.GetOlcRange() > 0 {
		scope = int(opts.GetOlcRange())
	}

	// Return Topic
	return &Topic{
		Name: fmt.Sprintf("/sonr/topic/%s", r.Location.OLC(scope)),
		Type: Topic_LOCAL,
	}
}

// Local Lobby Topic Protocol ID
func (r *User) NewDeviceTopic() *Topic {

	// Return Topic
	return &Topic{
		Name: fmt.Sprintf("/sonr/%s/%s", r.SName, r.DeviceID()),
		Type: Topic_LOCAL,
	}
}

// ** ─── Member MANAGEMENT ────────────────────────────────────────────────────────
// Update Position for Primary Peer in Member
func (m *Member) UpdatePosition(p *Position) {
	m.GetPrimary().Position = p
}

// Update Peer Profiles for Member
func (m *Member) UpdateProfile(c *Contact) {
	// Update Primary
	m.GetPrimary().Profile = &Profile{
		SName:     c.GetProfile().GetSName(),
		FirstName: c.GetProfile().GetFirstName(),
		LastName:  c.GetProfile().GetLastName(),
		Picture:   c.GetProfile().GetPicture(),
		Platform:  c.GetProfile().GetPlatform(),
	}

	// Update Associated
	for _, a := range m.GetAssociated() {
		a.Profile = &Profile{
			SName:     c.GetProfile().GetSName(),
			FirstName: c.GetProfile().GetFirstName(),
			LastName:  c.GetProfile().GetLastName(),
			Picture:   c.GetProfile().GetPicture(),
			Platform:  c.GetProfile().GetPlatform(),
		}
	}
}

// Return Users Primary Peer
func (u *User) GetPrimary() *Peer {
	return u.GetMember().GetPrimary()
}

// Set Primary Peer for Member
func (u *User) SetPrimary(id peer.ID, maddr multiaddr.Multiaddr, isLinker bool) {
	// Set Status
	if isLinker {
		u.Member = &Member{
			SName: u.SName,
			Primary: &Peer{
				Id: &Peer_ID{
					Peer:      id.String(),
					Device:    u.DeviceID(),
					MultiAddr: maddr.String(),
					PublicKey: u.KeyPair().PubKeyBase64(),
				},
				Platform: u.Device.Platform,
				Model:    u.GetDevice().GetModel(),
				HostName: u.GetDevice().GetHostName(),
				Status:   Peer_PAIRING,
			},
			Status: Member_GHOST,
		}
	} else {
		u.Member = &Member{
			SName: u.SName,
			Primary: &Peer{
				SName: u.SName,
				Id: &Peer_ID{
					Peer:      id.String(),
					Device:    u.DeviceID(),
					MultiAddr: maddr.String(),
					PublicKey: u.KeyPair().PubKeyBase64(),
					PushToken: u.GetPushToken(),
				},
				Profile:  u.Profile(),
				Platform: u.Device.Platform,
				Model:    u.GetDevice().GetModel(),
				HostName: u.GetDevice().GetHostName(),
				Status:   Peer_ONLINE,
			},
			Status: Member_ONLINE,
		}
	}
}

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────
// Set Peer from Connection Request and Host ID ^ //
func (u *User) SetPeer(id peer.ID, maddr multiaddr.Multiaddr, isLinker bool) *SonrError {
	// Set Status
	if isLinker {
		u.Member = &Member{
			SName: u.SName,
			Primary: &Peer{
				Id: &Peer_ID{
					Peer:      id.String(),
					Device:    u.DeviceID(),
					MultiAddr: maddr.String(),
					PublicKey: u.KeyPair().PubKeyBase64(),
				},
				Platform: u.Device.Platform,
				Model:    u.GetDevice().GetModel(),
				HostName: u.GetDevice().GetHostName(),
				Status:   Peer_PAIRING,
			},
			Status: Member_GHOST,
		}
	} else {
		u.Member = &Member{
			SName: u.SName,
			Primary: &Peer{
				SName: u.SName,
				Id: &Peer_ID{
					Peer:      id.String(),
					Device:    u.DeviceID(),
					MultiAddr: maddr.String(),
					PublicKey: u.KeyPair().PubKeyBase64(),
					PushToken: u.GetPushToken(),
				},
				Profile:  u.Profile(),
				Platform: u.Device.Platform,
				Model:    u.GetDevice().GetModel(),
				HostName: u.GetDevice().GetHostName(),
				Status:   Peer_ONLINE,
			},
			Status: Member_ONLINE,
		}
	}

	// Log Peer
	LogInfo(u.GetPrimary().String())
	return nil
}

// Checks if User Peer is a Linker
func (u *User) IsLinker() bool {
	return u.GetPrimary().Status == Peer_PAIRING
}

// Verify if Passed ShortID is Correct
func (u *User) VerifyLink(req *LinkRequest) (bool, *LinkResponse) {
	// Check if Peer is Linker
	if u.IsLinker() {
		// Verify Strings
		success := req.GetShortID() == u.GetDevice().ShortID()
		if success {
			return true, &LinkResponse{
				Type:    LinkResponse_Type(req.Type),
				To:      req.GetTo(),
				From:    req.GetFrom(),
				Device:  u.GetDevice(),
				Contact: req.GetContact(),
			}
		}
	}

	// Return Response
	return false, &LinkResponse{
		Success: false,
	}
}

// ** ─── Position MANAGEMENT ────────────────────────────────────────────────────────
func DefaultPosition() *Position {
	return &Position{
		Heading: &Position_Compass{
			Direction: 0,
			Antipodal: 180,
			Cardinal:  Cardinal_N,
		},
		Facing: &Position_Compass{
			Direction: 0,
			Antipodal: 180,
			Cardinal:  Cardinal_N,
		},
		Orientation: &Position_Orientation{
			Pitch: 0.0,
			Roll:  0.0,
			Yaw:   0.0,
		},
	}
}

// Returns Facing Direction
func (p *Position) FaceDirection() float64 {
	return p.GetFacing().GetDirection()
}

// Returns Heading Direction
func (p *Position) HeadDirection() float64 {
	return p.GetHeading().GetDirection()
}

// Returns Values Needed to Update Peer Instance
func (p *Position) Parameters() (float64, float64, *Position_Orientation) {
	return p.HeadDirection(), p.FaceDirection(), p.GetOrientation()
}

// ** ─── Local Event MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Join Topic Event
func NewJoinEvent(peer *Peer) *TopicEvent {
	return &TopicEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: TopicEvent_JOIN,
	}
}

// Creates New Update Topic Event
func NewUpdateEvent(peer *Peer, topic *Topic) *TopicEvent {
	return &TopicEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: TopicEvent_UPDATE,
		Topic:   topic,
	}
}

// Creates New Update Topic Event
func NewLinkerEvent(peer *Peer, topic *Topic) *TopicEvent {
	return &TopicEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: TopicEvent_LINKER,
		Topic:   topic,
	}
}

// Creates New Exit Topic Event
func NewExitEvent(id string, topic *Topic) *TopicEvent {
	return &TopicEvent{
		Id:      id,
		Subject: TopicEvent_EXIT,
		Topic:   topic,
	}
}

// Returns Peer as Buffer ^ //
func (p *Peer) Buffer() ([]byte, error) {
	buf, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// Returns Peer User ID ^ //
func (p *Peer) DeviceID() string {
	return string(p.Id.GetDevice())
}

// Returns Peer ID String Value
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// Returns Peer Public Key ^ //
func (p *Peer) PublicKey() (crypto.PubKey, *SonrError) {
	// Get ID from Public Key
	buf, err := crypto.ConfigDecodeKey(p.GetId().GetPublicKey())
	if err != nil {
		return nil, NewError(err, ErrorEvent_PEER_PUBKEY_DECODE)
	}

	// Unmarshal Public Key
	pubKey, err := crypto.UnmarshalPublicKey(buf)
	if err != nil {
		return nil, NewError(err, ErrorEvent_PEER_PUBKEY_UNMARSHAL)
	}
	return pubKey, nil
}

// Checks if Two Peers are the Same by Device ID and Peer ID
func (p *Peer) IsSame(other *Peer) bool {
	return p.PeerID() == other.PeerID() && p.DeviceID() == other.DeviceID() && p.GetSName() == other.GetSName()
}

// Checks if PeerDeviceIDID is the Same
func (p *Peer) IsSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// Checks if PeerID is the Same
func (p *Peer) IsSamePeerID(pid peer.ID) bool {
	return p.PeerID() == pid.String()
}

// Checks if Two Peers are NOT the Same by Device ID and Peer ID
func (p *Peer) IsNotSame(other *Peer) bool {
	return p.PeerID() != other.PeerID() && p.DeviceID() != other.DeviceID() && p.GetSName() != other.GetSName()
}

// Checks if DeviceID is NOT the Same
func (p *Peer) IsNotSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// Checks if PeerID is NOT the Same
func (p *Peer) IsNotSamePeerID(pid peer.ID) bool {
	return p.PeerID() != pid.String()
}

// Converts Peer Public Key into Thread Key
func (p *Peer) ThreadKey() (thread.PubKey, *SonrError) {
	// Get Pub Key
	pubKey, err := p.PublicKey()
	if err != nil {
		return nil, err
	}

	// Create Thread Pub Key
	threadKey := thread.NewLibp2pPubKey(pubKey)
	return threadKey, nil
}

// Returns Peer Push Token
func (p *Peer) PushToken() (string, *SonrError) {
	if p.Id.GetPushToken() == "" {
		return "", NewError(nil, ErrorEvent_PEER_PUSH_TOKEN_EMPTY)
	}
	return p.Id.GetPushToken(), nil
}
