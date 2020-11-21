package lobby

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/sonr-io/core/pkg/models"
)

// ** getID returns ONE Peer.ID in PubSub **
func (lob *Lobby) getID(q string) peer.ID {
	// Iterate through PubSub in topic
	for _, id := range lob.ps.ListPeers(lob.Data.Code) {
		// If Found Match
		if id.String() == q {
			return id
		}
	}
	// Log Error
	fmt.Println("Error QueryId was not found in PubSub topic")
	return ""
}

// ** getPeer returns ONE Peer in Lobby **
func (lob *Lobby) getPeer(q string) *pb.Peer {
	// Iterate Through Peers, Return Matched Peer
	for _, peer := range lob.Data.Peers {
		// If Found Match
		if peer.Id == q {
			return peer
		}
	}
	return nil
}

// ** removePeer deletes a peer from Lobby ** //
func (lob *Lobby) removePeer(id string) {
	// Delete peer from Lobby Map
	delete(lob.Data.Peers, id)

	// Send Callback with updated peers
	lob.call.Refreshed(lob.Peers())
}

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) updatePeer(id string, data *pb.Peer) {
	// Update Peer with new data
	lob.Data.Peers[id] = data

	// Send Callback with updated peers
	lob.call.Refreshed(lob.Peers())
}
