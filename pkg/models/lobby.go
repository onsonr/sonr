package models

import (
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ** ─── Global MANAGEMENT ────────────────────────────────────────────────────────
// Returns as Lobby Buffer
func (g *Global) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(g)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

// Remove Peer from Lobby
func (g *Global) Delete(id peer.ID) {
	// Find Username
	u, err := g.FindSName(id)
	if err != nil {
		return
	}

	// Remove Peer
	delete(g.Peers, u)
}

// Method Finds PeerID for Username
func (g *Global) FindPeerID(u string) (string, error) {
	for k, v := range g.Peers {
		if k == u {
			return v, nil
		}
	}
	return "", errors.New("PeerID not Found")
}

// Method Finds Username for PeerID
func (g *Global) FindSName(id peer.ID) (string, error) {
	for _, v := range g.Peers {
		if v == id.String() {
			return v, nil
		}
	}
	return "", errors.New("Username not Found")
}

// Method Checks if Global has Username
func (g *Global) HasSName(u string) bool {
	for k := range g.Peers {
		if k == u {
			return true
		}
	}
	return false
}

// Sync Between Remote Peers Lobby
func (g *Global) Sync(rg *Global) {
	// Iterate Over Remote Map
	for otherSName, id := range rg.Peers {
		if g.SName != otherSName {
			g.Peers[otherSName] = id
		}
	}

	// Check Self Map
	if !g.HasSName(rg.SName) {
		g.Peers[rg.SName] = rg.UserPeerID
	}
}

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
		Info: &Lobby_Local{
			Local: &Lobby_LocalInfo{
				Name:     topic[12:],
				Location: loc,
				Topic:    topic,
			},
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
	topic := ""
	switch l.Info.(type) {
	// @ Create Remote
	case *Lobby_Local:
		topic = l.GetLocal().GetTopic()
	// @ Join Remote
	case *Lobby_Remote:
		topic = l.GetRemote().GetTopic()
	}
	return topic
}

// Returns as Lobby Buffer
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println(err)
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
