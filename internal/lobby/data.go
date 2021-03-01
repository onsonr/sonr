package lobby

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
)

// ** removePeer removes Peer from Map **
func (lob *Lobby) removePeer(id peer.ID) {
	// Update Peer with new data
	delete(lob.data.Peers, id.String())
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// ** standbyPeer puts a Peer in Standby Mode **
func (lob *Lobby) resumePeer(peer *md.Peer) {
	// Update Peer with new data
	delete(lob.data.Standby, peer.Id)
	lob.data.Peers[peer.Id] = peer
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// ** standbyPeer puts a Peer in Standby Mode **
func (lob *Lobby) standbyPeer(peer *md.Peer) {
	// Update Peer with new data
	delete(lob.data.Peers, peer.Id)
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Add to Standby
	lob.data.Standby[peer.Id] = peer

	// Callback with Updated Data
	lob.Refresh()
}

// ** updatePeer changes Peer values in Lobby **
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Update Peer with new data
	lob.data.Peers[peer.Id] = peer
	lob.data.Count = int32(len(lob.data.Peers))
	lob.data.Size = int32(len(lob.data.Peers)) + 1 // Account for User

	// Callback with Updated Data
	lob.Refresh()
}

// @ Helper: ID returns ONE Peer.ID in PubSub
func (lob *Lobby) ID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.pubSub.ListPeers(lob.data.Olc) {
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
	for _, peer := range lob.data.Peers {
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
