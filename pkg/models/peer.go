package models

import (
	"github.com/libp2p/go-libp2p-core/peer"
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
		Type:  Lobby_LOCAL,
		Peers: make(map[string]*Peer),
		User:  u.GetPeer(),

		// Info
		Info: &Lobby_LocalInfo{
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
