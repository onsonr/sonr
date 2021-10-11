package lobby

import (
	"fmt"
	"time"
	olc "github.com/google/open-location-code/go"
	peer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/api"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/state"
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

func checkParams(host *host.SNRHost, em *state.Emitter) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	if em == nil {
		logger.Error("Emitter provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}

func createOlc(l *common.Location) string {
	code := olc.Encode(l.GetLatitude(), l.GetLongitude(), 6)
	if code == "" {
		logger.Error("Failed to Determine OLC Code, set to Global")
		code = "global"
	}
	logger.Info("Calculated OLC for Location: " + code)
	return fmt.Sprintf("sonr/topic/%s", code)
}

// pushRefresh sends a refresh event to the emitter
func (p *LobbyProtocol) pushRefresh(id peer.ID, peer *common.Peer) {
	// Function to build a refreshEvent
	var buildEvent = func(peers []*common.Peer) *api.RefreshEvent {
		return &api.RefreshEvent{
			Olc:      p.olc,
			Peers:    peers,
			Received: int64(time.Now().Unix()),
		}
	}

	// Check if Peer was provided
	if peer == nil {
		// Remove Peer, Emit Event
		p.emitter.Emit(Event_LIST_REFRESH, buildEvent(p.removePeer(id)))

	} else {
		// Update Peer, Emit Event
		p.emitter.Emit(Event_LIST_REFRESH, buildEvent(p.updatePeer(id, peer)))
	}

}

// removePeer Removes Peer from Peer List
func (lp *LobbyProtocol) removePeer(peerID peer.ID) []*common.Peer {
	for i, p := range lp.peers {
		if peer.ID(p.GetPeerID()) == peerID {
			return append(lp.peers[:i], lp.peers[i+1:]...)
		}
	}
	return lp.peers
}

// updatePeer Adds Peer to Peer List
func (lp *LobbyProtocol) updatePeer(peerID peer.ID, peer *common.Peer) []*common.Peer {
	lp.peers = append(lp.peers, peer)
	return lp.peers
}
