package room

import (
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/data"
)

// Check if PeerEvent is Join and NOT User
func (tm *RoomManager) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// Check if PeerEvent is Exit and NOT User
func (tm *RoomManager) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// Check if Message is NOT from User
func (tm *RoomManager) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}

// Returns RoomData Data instance
func (tm *RoomManager) Room() *data.Room {
	return tm.room
}
