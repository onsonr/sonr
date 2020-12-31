package lobby

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
)

// ^ Find returns Pointer to Peer.ID and Peer ^
func (lob *Lobby) Find(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	if peer == nil || id == "" {
		return "", nil, errors.New("Search Error, peer was not found in map.")
	}

	return id, peer, nil
}

// ^ ID returns ONE Peer.ID in PubSub ^
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.ps.ListPeers(lob.Data.Code) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// ^ Peer returns ONE Peer in Lobby ^
func (lob *Lobby) Peer(q string) *md.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range lob.Data.Peers {
		// If Found Match
		if peer.Id == q {
			return peer
		}
	}
	return nil
}

// ^ setPeer changes peer values in Lobby ^
func (lob *Lobby) setPeer(msg *md.LobbyMessage) {
	// Update Peer with new data
	lob.Data.Peers[msg.Id] = msg.Peer

	// Send Event
	lob.sendRefresh()
}

// ^ removePeer deletes peer from all maps ^
func (lob *Lobby) removePeer(id peer.ID) {
	// Remove Peer from Peers
	delete(lob.Data.Peers, id.String())

	// Send Event
	lob.sendRefresh()
}
