package lobby

import (
	peer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
)

// HasPeer Method Checks if Peer ID is Subscribed to Room
func (tm *LobbyProtocol) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range tm.topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// isEventJoin Checks if PeerEvent is Join and NOT User
func (tm *LobbyProtocol) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != tm.host.ID()
}

// isEventExit Checks if PeerEvent is Exit and NOT User
func (tm *LobbyProtocol) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != tm.host.ID()
}

// isValidMessage Checks if Message is NOT from User
func (tm *LobbyProtocol) isValidMessage(msg *ps.Message) bool {
	return tm.host.ID() != msg.ReceivedFrom && tm.HasPeerID(msg.ReceivedFrom)
}
