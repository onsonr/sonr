package lobby

import (
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
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
	for _, peer := range lob.Data.Available {
		// If Found Match
		if peer.Id == q {
			return peer
		}
	}
	return nil
}

// ^ setPeer changes peer values in Lobby ^
func (lob *Lobby) setPeer(msg *md.LobbyMessage) {
	// Remove Peer from Unavailable
	delete(lob.Data.Unavailable, msg.Id)

	// Update Peer with new data
	id := msg.Id
	lob.Data.Available[id] = msg.Peer
	lob.Data.Size = int32(len(lob.Data.Available)) + 1 // Account for User

	// Create Event
	event := md.LobbyEvent{
		Event:     md.LobbyEvent_UPDATE,
		Peer:      msg.Peer,
		Direction: msg.Direction,
	}

	// Marshal data to bytes
	bytes, err := proto.Marshal(&event)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callEvent(bytes)
}

// ^ setBusy changes peer values in Lobby ^
func (lob *Lobby) setUnavailable(msg *md.LobbyMessage) {
	// Remove Peer from Available
	delete(lob.Data.Available, msg.Id)

	// Add Peer to Unavailable Map
	lob.Data.Unavailable[msg.Id] = msg.Peer
	lob.Data.Size = int32(len(lob.Data.Available)) + 1 // Account for User

	// Create Event
	event := md.LobbyEvent{
		Event: md.LobbyEvent_BUSY,
		Peer:  msg.Peer,
	}

	// Marshal data to bytes
	bytes, err := proto.Marshal(&event)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callEvent(bytes)
}

// ^ removePeer deletes peer from all maps ^
func (lob *Lobby) removePeer(id string) {
	// Remove Peer from Available
	delete(lob.Data.Available, id)

	// Remove Peer from Unavailable
	delete(lob.Data.Unavailable, id)
	lob.Data.Size = int32(len(lob.Data.Available)) + 1 // Account for User

	// Create Event
	event := md.LobbyEvent{
		Event:  md.LobbyEvent_BUSY,
		PeerId: id,
	}

	// Marshal data to bytes
	bytes, err := proto.Marshal(&event)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Send Callback with updated peers
	lob.callEvent(bytes)
}
