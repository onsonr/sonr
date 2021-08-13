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

// ** ─── Room MANAGEMENT ────────────────────────────────────────────────────────
// Checks if Room Type is Local
func (t *Room) IsLocal() bool {
	return t.Type == Room_LOCAL
}

// Checks if Room Type is Devices
func (t *Room) IsDevices() bool {
	return t.Type == Room_DEVICE
}

// Checks if Room Type is Group
func (t *Room) IsGroup() bool {
	return t.Type == Room_GROUP
}

// Local Lobby Room Protocol ID
func (r *Device) NewLocalRoom(opts *ConnectionRequest_ServiceOptions) *Room {
	// Initialize Set OLC Range
	scope := 6
	if opts.GetOlcRange() > 0 {
		scope = int(opts.GetOlcRange())
	}

	// Return Room
	return &Room{
		Name: fmt.Sprintf("/sonr/local/%s", r.Location.OLC(scope)),
		Type: Room_LOCAL,
	}
}

// Local Lobby Room Protocol ID
func (r *Account) NewDeviceRoom() *Room {

	// Return Room
	return &Room{
		Name: fmt.Sprintf("/sonr/device/%s", r.SName),
		Type: Room_DEVICE,
	}
}

// Local Lobby Room Protocol ID
func (r *Account) NewGroupRoom(name string) *Room {

	// Return Room
	return &Room{
		Name: fmt.Sprintf("/sonr/group/%s", name),
		Type: Room_GROUP,
	}
}

// ** ─── Member MANAGEMENT ────────────────────────────────────────────────────────
// Update Peer Profiles for Member
func (m *Member) UpdateProfile(c *Contact) {
	// Update General
	m.SName = c.GetProfile().GetSName()

	// Update Primary
	m.GetActive().Profile = &Profile{
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
func (u *Account) ActivePeer() *Peer {
	return u.GetMember().GetActive()
}

// Set Primary Peer for Member. Returns Peer Ref and if Primary Peer
func (d *Device) SetPeer(id peer.ID, maddr multiaddr.Multiaddr, isLinker bool) (*Peer, bool) {
	// Set Status
	if isLinker {
		peer := &Peer{
			Id: &Peer_ID{
				Peer:      id.String(),
				Device:    d.Id,
				MultiAddr: maddr.String(),
				PublicKey: d.AccountKeys().PubKeyBase64(),
			},
			Platform: d.Platform,
			Model:    d.GetModel(),
			HostName: d.GetHostName(),
			Status:   Peer_PAIRING,
		}

		// Set Primary
		d.Peer = peer
		return peer, d.HasDeviceKeys()

	} else {
		peer := &Peer{
			Id: &Peer_ID{
				Peer:      id.String(),
				Device:    d.Id,
				MultiAddr: maddr.String(),
				PublicKey: d.AccountKeys().PubKeyBase64(),
			},
			Platform: d.Platform,
			Model:    d.GetModel(),
			HostName: d.GetHostName(),
			Status:   Peer_ONLINE,
		}

		// Set Primary
		d.Peer = peer
		return peer, d.HasDeviceKeys()
	}
}

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────

// Checks if User Peer is a Linker
func (u *Device) IsLinker() bool {
	return u.GetPeer().Status == Peer_PAIRING
}

// Verify if Passed ShortID is Correct
func (u *Device) VerifyLink(req *LinkRequest) (bool, *LinkResponse) {
	// Check if Peer is Linker
	if u.IsLinker() {
		// Verify Strings
		success := req.GetShortID() == u.ShortID()
		if success {
			return true, &LinkResponse{
				Type:    LinkResponse_Type(req.Type),
				To:      req.GetTo(),
				From:    req.GetFrom(),
				Device:  u,
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
// Creates New Join Room Event
func (r *Room) NewJoinEvent(peer *Peer) *RoomEvent {
	return &RoomEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: RoomEvent_JOIN,
		Room:    r,
	}
}

// Creates New Update Room Event
func (r *Room) NewUpdateEvent(peer *Peer) *RoomEvent {
	return &RoomEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: RoomEvent_UPDATE,
		Room:    r,
	}
}

// Creates New Update Room Event
func (r *Room) NewLinkerEvent(peer *Peer) *RoomEvent {
	return &RoomEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: RoomEvent_LINKER,
		Room:    r,
	}
}

// Creates New Exit Room Event
func (r *Room) NewExitEvent(id string) *RoomEvent {
	return &RoomEvent{
		Id:      id,
		Subject: RoomEvent_EXIT,
		Room:    r,
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
