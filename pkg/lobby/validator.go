package lobby

import (
	"time"

	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	common "github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

// callRefresh calls back RefreshEvent to Node
func (lp *LobbyProtocol) callRefresh() {
	// Check Peer Data List Length
	if len(lp.peers) == 0 {
		return
	}

	// Create Event
	logger.Debug("Calling Refresh Event")
	lp.node.OnRefresh(&api.RefreshEvent{
		Olc:      lp.olc,
		Peers:    lp.peers,
		Received: int64(time.Now().Unix()),
	})
}

// callUpdate publishes a LobbyMessage to the Topic
func (lp *LobbyProtocol) callUpdate() error {
	// Check Lobby Topic Peers Length
	if len(lp.topic.ListPeers()) == 0 {
		return nil
	}

	// Create Event
	logger.Debug("Sending Update to Lobby")
	err := lp.Update()
	if err != nil {
		logger.Error("Failed to update peer", err)
		return err
	}
	return nil
}

// createLobbyMsgBuf Creates a new Message Buffer for Lobby Topic
func createLobbyMsgBuf(p *common.Peer) []byte {
	// Create Event
	if p == nil {
		logger.Errorf("%s - Peer provided is nil", ErrParameters)
		return nil
	}
	event := &LobbyMessage{Peer: p}

	// Marshal Event
	eventBuf, err := proto.Marshal(event)
	if err != nil {
		logger.Errorf("%s - Failed to Marshal Event", err)
		return nil
	}
	return eventBuf
}

// hasPeer Checks if Peer is in Peer List
func (lp *LobbyProtocol) hasPeer(data *common.Peer) bool {
	hasInList := false
	hasInTopic := false
	for _, p := range lp.peers {
		if p.GetPeerID() == data.GetPeerID() {
			hasInList = true
		}
	}

	for _, p := range lp.topic.ListPeers() {
		if p.String() == data.GetPeerID() {
			hasInTopic = true
		}
	}
	return hasInList && hasInTopic
}

// hasPeerData Checks if Peer Data is in Lobby Peer-Data List
func (lp *LobbyProtocol) hasPeerData(data *common.Peer) bool {
	for _, p := range lp.peers {
		if p.GetPeerID() == data.GetPeerID() {
			return true
		}
	}
	return false
}

// hasPeerID Checks if Peer ID is in Lobby Topic
func (lp *LobbyProtocol) hasPeerID(id peer.ID) bool {
	for _, p := range lp.topic.ListPeers() {
		if p == id {
			return true
		}
	}
	return false
}

// indexOfPeer Returns Peer Index in Peer-Data List
func (lp *LobbyProtocol) indexOfPeer(peerID peer.ID) int {
	for i, p := range lp.peers {
		if p.GetPeerID() == peerID.String() {
			return i
		}
	}
	return -1
}

// removePeer Removes Peer from Peer-Data List
func (lp *LobbyProtocol) removePeer(peerID peer.ID) bool {
	for i, p := range lp.peers {
		if p.GetPeerID() == peerID.String() {
			lp.peers = append(lp.peers[:i], lp.peers[i+1:]...)
			return true
		}
	}
	return false
}

// updatePeer Adds Peer to Peer List
func (lp *LobbyProtocol) updatePeer(peerID peer.ID, data *common.Peer) bool {
	// Check if Peer is in Peer List and Topic already
	if !lp.hasPeerID(peerID) {
		lp.removePeer(peerID)
		return false
	}

	// Check if Peer is in Peer-Data List
	if lp.hasPeer(data) {
		return false
	}

	// Add Peer to List and Check if Peer is List
	idx := lp.indexOfPeer(peerID)
	if idx == -1 {
		lp.peers = append(lp.peers, data)
	} else {
		lp.peers[idx] = data
	}
	return true
}
