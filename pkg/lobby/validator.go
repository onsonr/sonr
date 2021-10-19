package lobby

import (
	"fmt"
	"time"

	peer "github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/api"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
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

// checkParams Checks if Non-nil Parameters were passed
func checkParams(host *host.SNRHost) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}

// createOlc Creates a new Olc from Location
func createOlc(l *common.Location) string {
	code := l.OLC()
	if code == "" {
		logger.Error("Failed to Determine OLC Code, set to Global")
		code = "global"
	}
	logger.Debug("Calculated OLC for Location: " + code)
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
		p.node.OnRefresh(buildEvent(p.removePeer(id)))
	} else {
		// Update Peer, Emit Event
		ok, list := p.updatePeer(id, peer)
		p.node.OnRefresh(buildEvent(list))
		if !ok {
			p.sendUpdate()
		}
	}
}

// sendUpdate sends a refresh event to the Lobby topic
func (lp *LobbyProtocol) sendUpdate() error {
	peer, err := lp.node.Peer()
	if err != nil {
		logger.Error("Failed to get peer", err)
		return err
	}
	err = lp.Update(peer)
	if err != nil {
		logger.Error("Failed to update peer", err)
		return err
	}
	return nil
}

// hasPeer Checks if Peer is in Peer List
func (lp *LobbyProtocol) hasPeer(data *common.Peer) bool {
	for _, p := range lp.peers {
		if p.GetPeerID() == data.GetPeerID() {
			return true
		}
	}
	return false
}

// removePeer Removes Peer from Peer List
func (lp *LobbyProtocol) removePeer(peerID peer.ID) []*common.Peer {
	for i, p := range lp.peers {
		if p.GetPeerID() == peerID.String() {
			return append(lp.peers[:i], lp.peers[i+1:]...)
		}
	}
	return lp.peers
}

// updatePeer Adds Peer to Peer List
func (lp *LobbyProtocol) updatePeer(peerID peer.ID, data *common.Peer) (bool, []*common.Peer) {
	if lp.hasPeer(data) {
		return false, lp.peers
	}
	lp.peers = append(lp.peers, data)
	return true, lp.peers
}
