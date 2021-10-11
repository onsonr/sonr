package lobby

import (
	peer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
)

// HasPeer Method Checks if Peer ID is Subscribed to Room
func (p *LobbyProtocol) HasPeerID(q peer.ID) bool {
	// Iterate through PubSub in room
	for _, id := range p.topic.ListPeers() {
		// If Found Match
		if id == q {
			return true
		}
	}
	return false
}

// isEventJoin Checks if PeerEvent is Join and NOT User
func (p *LobbyProtocol) isEventJoin(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerJoin && ev.Peer != p.host.ID()
}

// isEventExit Checks if PeerEvent is Exit and NOT User
func (p *LobbyProtocol) isEventExit(ev ps.PeerEvent) bool {
	return ev.Type == ps.PeerLeave && ev.Peer != p.host.ID()
}

// isValidMessage Checks if Message is NOT from User
func (p *LobbyProtocol) isValidMessage(msg *ps.Message) bool {
	return p.host.ID() != msg.ReceivedFrom && p.HasPeerID(msg.ReceivedFrom)
}




