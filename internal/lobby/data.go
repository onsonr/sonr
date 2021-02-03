package lobby

import (
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) removePeer(id peer.ID) {
	// Update Peer with new data
	delete(lob.Data.Peers, id.String())
	lob.Data.Count = int32(len(lob.Data.Peers))
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.onRefresh(bytes)
}

// ^ sendPeerLeft notifies when a peer exited lobby ^
func (lob *Lobby) sendPeerLeft(id peer.ID) {
	// Create Event Message
	lobEvent := &md.LobbyEvent{
		Id:    id.String(),
		Event: md.LobbyEvent_EXIT,
	}

	// Marshal data to bytes
	bytes, err := proto.Marshal(lobEvent)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.onEvent(bytes)
}

// ^ sendPeerEvent notifies when a peer exited lobby ^
func (lob *Lobby) sendPeerEvent(event *md.LobbyEvent) {
	// Marshal data to bytes
	bytes, err := proto.Marshal(event)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.onEvent(bytes)
}

// ** updatePeer changes peer values in Lobby **
func (lob *Lobby) updatePeer(peer *md.Peer) {
	// Update Peer with new data
	id := peer.Id
	lob.Data.Peers[id] = peer
	lob.Data.Count = int32(len(lob.Data.Peers))
	lob.Data.Size = int32(len(lob.Data.Peers)) + 1 // Account for User

	// Marshal data to bytes
	bytes, err := proto.Marshal(lob.Data)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.onRefresh(bytes)
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
