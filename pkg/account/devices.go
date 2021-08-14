package account

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

type GetRoomFunc func() *md.Room

// FindPeer @ Helper: Find returns Pointer to Peer.ID and Peer
func (tm *accountLinker) FindPeer(q string) (peer.ID, error) {
	// Iterate through Room Peers
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", errors.New("Peer ID was not found in room")
}

// Publish @ Publish message to specific peer in room
func (tm *accountLinker) Publish(msg *md.RoomEvent) error {
	if tm.room.IsLocal() || tm.room.IsGroup() {
		// Convert Event to Proto Binary
		bytes, err := proto.Marshal(msg)
		if err != nil {
			md.LogError(err)
			return err
		}

		// Publish to Room
		err = tm.Topic.Publish(tm.ctx, bytes)
		if err != nil {
			md.LogError(err)
			return err
		}
	}
	return nil
}

// Publish @ Publish message to specific peer in room
func (tm *accountLinker) Sync(msg *md.SyncEvent) error {
	if tm.room.IsDevices() {
		// Convert Event to Proto Binary
		bytes, err := proto.Marshal(msg)
		if err != nil {
			md.LogError(err)
			return err
		}

		// Publish to Room
		err = tm.Topic.Publish(tm.ctx, bytes)
		if err != nil {
			md.LogError(err)
			return err
		}
	}
	return nil
}

// HasPeer Method Checks if Peer ID String is Subscribed to Room
func (tm *accountLinker) HasPeer(q string) bool {
	// Iterate through PubSub in room
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// HasPeer Method Checks if Peer ID is Subscribed to Room
func (tm *accountLinker) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range tm.Topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// # Check if PeerEvent is Join and NOT User
func (tm *accountLinker) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// # Check if PeerEvent is Exit and NOT User
func (tm *accountLinker) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// # Check if Message is NOT from User
func (tm *accountLinker) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}

// Returns RoomData Data instance
func (tm *accountLinker) Room() *md.Room {
	return tm.room
}
