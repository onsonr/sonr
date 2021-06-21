package models

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── Lobby MANAGEMENT ────────────────────────────────────────────────────────
// Creates Local Lobby from User Data
func NewLocalLobby(u *User) *Lobby {
	// Get Info
	topic := u.LocalTopic()
	loc := u.GetRouter().GetLocation()

	// Create Lobby
	return &Lobby{
		// General
		Type:  TopicType_LOCAL,
		Peers: make(map[string]*Peer),
		User:  u.GetPeer(),

		// Info
		Info: &Lobby_Info{
			Name:     topic[12:],
			Location: loc,
			Topic:    topic,
		},
	}
}

// Returns Lobby Peer Count
func (l *Lobby) Count() int {
	return len(l.Peers)
}

// Returns TOTAL Lobby Size with Peer
func (l *Lobby) Size() int {
	return len(l.Peers) + 1
}

// Returns Lobby Topic
func (l *Lobby) Topic() string {
	return l.GetInfo().GetTopic()
}

// Returns as Lobby Buffer
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Add/Update Peer in Lobby
func (l *Lobby) Add(peer *Peer) {
	// Update Peer with new data
	l.Peers[peer.PeerID()] = peer
}

// Remove Peer from Lobby
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
}

// ** ─── Local Event MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Exit Local Event
func NewJoinLocalEvent(peer *Peer) *LobbyEvent {
	return &LobbyEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: LobbyEvent_JOIN,
		Type:    TopicType_LOCAL,
	}
}

// Creates New Exit Local Event
func NewUpdateLocalEvent(peer *Peer) *LobbyEvent {
	return &LobbyEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: LobbyEvent_UPDATE,
		Type:    TopicType_LOCAL,
	}
}

// Creates New Exit Local Event
func NewExitLocalEvent(id string) *LobbyEvent {
	return &LobbyEvent{
		Id:      id,
		Subject: LobbyEvent_EXIT,
		Type:    TopicType_LOCAL,
	}
}

// ** ─── Peer Instance MANAGEMENT ────────────────────────────────────────────────────────
// Converts Peer to PeerInstance for Threads Storage
func (p *Peer) ToInstance() *PeerInstance {
	return &PeerInstance{
		SName:       p.GetSName(),
		PeerID:      p.PeerID(),
		MultiAddr:   p.Id.GetMultiAddr(),
		FirstName:   p.Profile.GetFirstName(),
		IsActive:    true,
		IsReachable: true,
	}
}

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────
// ^ Create New Peer from Connection Request and Host ID ^ //
func (u *User) NewPeer(id peer.ID, maddr multiaddr.Multiaddr) *SonrError {
	u.Peer = &Peer{
		SName: u.SName,
		Id: &Peer_ID{
			Peer:      id.String(),
			Device:    u.DeviceID(),
			MultiAddr: maddr.String(),
			PublicKey: u.KeyPair().GetPublic().GetBuffer(),
		},
		Profile:  u.Profile(),
		Platform: u.Device.Platform,
		Model:    u.Device.Model,
	}
	// Set Device Topic
	u.Router.DeviceTopic = fmt.Sprintf("/sonr/topic/%s", u.Peer.GetSName())
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

// ^ Returns Peer ID String Value
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// ^ Returns Peer Public Key ^ //
func (p *Peer) PublicKey() crypto.PubKey {
	buf := p.GetId().GetPublicKey()

	// Get Key from Buffer
	pubKey, err := crypto.UnmarshalPublicKey(buf)
	if err != nil {
		return nil
	}
	return pubKey
}

// ^ Checks if Two Peers are the Same by Device ID and Peer ID
func (p *Peer) IsSame(other *Peer) bool {
	return p.PeerID() == other.PeerID() && p.DeviceID() == other.DeviceID() && p.GetSName() == other.GetSName()
}

// ^ Checks if PeerDeviceIDID is the Same
func (p *Peer) IsSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// ^ Checks if PeerID is the Same
func (p *Peer) IsSamePeerID(pid peer.ID) bool {
	return p.PeerID() == pid.String()
}

// ^ Checks if Two Peers are NOT the Same by Device ID and Peer ID
func (p *Peer) IsNotSame(other *Peer) bool {
	return p.PeerID() != other.PeerID() && p.DeviceID() != other.DeviceID() && p.GetSName() != other.GetSName()
}

// ^ Checks if DeviceID is NOT the Same
func (p *Peer) IsNotSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// ^ Checks if PeerID is NOT the Same
func (p *Peer) IsNotSamePeerID(pid peer.ID) bool {
	return p.PeerID() != pid.String()
}
