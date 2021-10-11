package lobby

import (
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/state"
)

var (
	logger = golog.Child("protocols/lobby")
)

func checkParams(host *host.SNRHost, olc string, em *state.Emitter) error {
	if host == nil {
		logger.Error("Host provided is nil", ErrParameters)
		return ErrParameters
	}
	if olc == "" {
		logger.Error("Location provided is nil", ErrParameters)
		return ErrParameters
	}
	if em == nil {
		logger.Error("Emitter provided is nil", ErrParameters)
		return ErrParameters
	}
	return host.HasRouting()
}

func createOlc(l *common.Location) string {
	code := olc.Encode(l.GetLatitude(), l.GetLongitude(), 8)
	if code == "" {
		logger.Error("Failed to Determine OLC Code, set to Global")
		return "global"
	}
	logger.Info("Calculated OLC for Location", golog.Fields{"olc": code, "latitude": l.GetLatitude(), "longitude": l.GetLongitude()})
	return code
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
