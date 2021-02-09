package lobby

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
)

// ** removePeer removes Peer from Map **
func (lob *Lobby) removePeer(id peer.ID) {
	// Update Peer with new data
	delete(lob.Data.Peers, id.String())
	lob.Data.Count = int32(len(lob.Data.Peers))
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// ** standbyPeer puts a Peer in Standby Mode **
func (lob *Lobby) resumePeer(peer *md.Peer) {
	// Update Peer with new data
	delete(lob.Data.Standby, peer.Id)
	lob.Data.Peers[peer.Id] = peer
	lob.Data.Count = int32(len(lob.Data.Peers))
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// ** standbyPeer puts a Peer in Standby Mode **
func (lob *Lobby) standbyPeer(peer *md.Peer) {
	// Update Peer with new data
	delete(lob.Data.Peers, peer.Id)
	lob.Data.Count = int32(len(lob.Data.Peers))
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Add to Standby
	lob.Data.Standby[peer.Id] = peer

	// Callback with Updated Data
	lob.Refresh()
}

// ** updatePeer changes Peer values in Lobby **
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Update Peer with new data
	lob.Data.Peers[peer.Id] = peer
	lob.Data.Count = int32(len(lob.Data.Peers))
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// @ Helper: ID returns ONE Peer.ID in PubSub
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.ps.ListPeers(lob.Data.Olc) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	return ""
}

// @ Helper: Peer returns ONE Peer in Lobby
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

// @ Helper: Find returns Pointer to Peer.ID and Peer
func (lob *Lobby) Find(q string) (peer.ID, *md.Peer, error) {
	// Retreive Data
	peer := lob.Peer(q)
	id := lob.ID(q)

	if peer == nil || id == "" {
		return "", nil, errors.New("Search Error, peer was not found in map.")
	}

	return id, peer, nil
}
