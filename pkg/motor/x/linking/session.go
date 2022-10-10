package linking

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"

	"github.com/sonr-io/sonr/pkg/host"
	// motor "go.buf.build/grpc/go/sonr-io/motor/common/v1"
	// v1 "go.buf.build/grpc/go/sonr-io/motor/service/v1"
	"github.com/sonr-io/sonr/third_party/types/common"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	st "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
	"google.golang.org/protobuf/proto"
)

// ErrFunc is a function that returns an error
type ErrFunc func() error

// Session is the protocol for managing local peers.
type Session struct {
	callback     common.MotorCallback
	node         host.SonrHost
	ctx          context.Context
	eventHandler *ps.TopicEventHandler
	subscription *ps.Subscription
	topic        *ps.Topic
	olc          string
	peers        []*ct.Peer
	selfID       peer.ID
	updateFunc   ErrFunc
}

// Initializing the local struct.
func (e *LinkingProtocol) buildNewSession(topic *ps.Topic, cb common.MotorCallback) error {
	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		logger.Errorf("%s - Failed to Subscribe to OLC Topic", err)
		return err
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		logger.Errorf("%s - Failed to Get Event Handler", err)
		return err
	}

	// Create Local Struct
	e.session = &Session{
		ctx:          e.ctx,
		selfID:       e.node.HostID(),
		node:         e.node,
		topic:        topic,
		subscription: sub,
		eventHandler: handler,
		olc:          sub.Topic(),
		peers:        make([]*ct.Peer, 0),
	}

	// Handle Events
	go e.session.handleSub()
	go e.session.handleTopic()
	go e.session.handleEvents()
	go e.session.handleTimeout()
	return nil
}

// handleTimeout method closes the session if timeout is reached
func (p *Session) handleTimeout() {
	// Loop Messages
	for {
		err := p.callUpdate()
		if err != nil {
			logger.Error("Failed to send peer update to lobby topic", err)
			continue
		}
		time.Sleep(time.Second * 8)
	}
}

// handleSub method listens to Pubsub Events for Local Topic
func (p *Session) handleSub() {
	// Loop Events
	for {
		// Get next event
		event, err := p.eventHandler.NextPeerEvent(p.ctx)
		if err != nil {
			logger.Errorf("%s - Failed to Get Next Peer Event", err)
			return
		}

		// Check Event and Validate not User
		if event.Type == ps.PeerLeave && event.Peer != p.selfID {
			// Remove Peer, Emit Event
			p.messages <- newLobbyEvent(event.Peer, nil)
			continue
		}
	}
}

// handleTopic method listens to Pubsub Messages for Local Topic
func (p *Session) handleTopic() {
	// Loop Messages
	for {
		// Get next message
		msg, err := p.subscription.Next(p.ctx)
		if err != nil {
			return
		}

		// Check Message and Validate not User
		if msg.ReceivedFrom != p.selfID {
			// Unmarshal Message
			data := &st.LobbyMessage{}
			err = data.Unmarshal(msg.Data)
			if err != nil {
				logger.Errorf("%s - Failed to Unmarshal Message", err)
				continue
			}
			p.messages <- newLobbyEvent(msg.ReceivedFrom, data.GetFrom())
		}
	}
}

// handleEvents method listens to Lobby Events passed from the Local Topic
func (p *Session) handleEvents() {
	// Loop Messages
	for {
		// Get next message
		msg := <-p.messages

		// Update Peer, Emit Event
		if msg.isExit {
			p.removePeer(msg.ID)
		} else {
			p.updatePeer(msg.ID, msg.Peer)
		}
	}
}

// callRefresh calls back RefreshEvent to Node
func (lp *Session) callRefresh() {
	// Create Event
	logger.Debug("Calling Refresh Event")
	ev := &st.RefreshEvent{
		Peers:      lp.peers,
		TopicName:  lp.olc,
		ReceivedAt: int64(time.Now().Unix()),
	}

	// Emit Refresh Event
	buf, err := ev.Marshal()
	if err != nil {
		logger.Errorf("%s - Failed to Marshal Refresh Event", err)
		return
	}
	lp.callback.OnDiscover(buf)
}

// callUpdate publishes a LobbyMessage to the Local Topic
func (lp *Session) callUpdate() error {
	// Create Event
	logger.Debug("Sending Update to Lobby")
	err := lp.updateFunc()
	if err != nil {
		logger.Errorf("%s - Failed to update peer", err)
		return err
	}
	return nil
}

// createLobbyMsgBuf Creates a new Message Buffer for Local Topic
func createLobbyMsgBuf(p *ct.Peer) []byte {
	// Marshal Event
	//event := &v1.LobbyMessage{Peer: p}
	_, err := proto.Marshal(nil)
	if err != nil {
		logger.Errorf("%s - Failed to Marshal Event", err)
		return nil
	}
	// return eventBuf
	return nil
}

// hasPeer Checks if Peer is in Peer List
func (lp *Session) hasPeer(data *ct.Peer) bool {
	hasInList := false
	hasInTopic := false
	// Check if Peer is in Data List
	for _, p := range lp.peers {
		if p.GetAccountId() == data.GetAccountId() {
			hasInList = true
		}
	}

	// Check if Peer is in Topic
	for _, p := range lp.topic.ListPeers() {
		if p.String() == data.GetAccountId() {
			hasInTopic = true
		}
	}
	if !hasInList {
		logger.Warn("Peer is subscribed to Topic but does not have Peer Data")
	}
	return hasInList && hasInTopic
}

// hasPeerData Checks if Peer Data is in Local Peer-Data List
func (lp *Session) hasPeerData(data *ct.Peer) bool {
	for _, p := range lp.peers {
		if p.GetAccountId() == data.GetAccountId() {
			return true
		}
	}
	return false
}

// hasPeerID Checks if Peer ID is in Local Topic
func (lp *Session) hasPeerID(id peer.ID) bool {
	for _, p := range lp.topic.ListPeers() {
		if p == id {
			return true
		}
	}
	return false
}

// indexOfPeer Returns Peer Index in Local Peer-Data List
func (lp *Session) indexOfPeer(peer *ct.Peer) int {
	for i, p := range lp.peers {
		if p.GetAccountId() == peer.GetAccountId() {
			return i
		}
	}
	return -1
}

// removePeer Removes Peer from Local Peer-Data List
func (lp *Session) removePeer(peerID peer.ID) bool {
	for i, p := range lp.peers {
		if p.GetAccountId() == peerID.String() {
			lp.peers = append(lp.peers[:i], lp.peers[i+1:]...)
			lp.callRefresh()
			return true
		}
	}
	return false
}

// updatePeer Adds Peer to Local Peer List
func (lp *Session) updatePeer(peerID peer.ID, data *ct.Peer) bool {
	// Check if Peer is in Peer List and Topic already
	if ok := lp.hasPeerID(peerID); !ok {
		lp.removePeer(peerID)
		return false
	}

	// Add Peer to List and Check if Peer is in Local List
	idx := lp.indexOfPeer(data)
	if idx == -1 {
		lp.peers = append(lp.peers, data)
		lp.callUpdate()
	} else {
		lp.peers[idx] = data
	}
	lp.callRefresh()
	return true
}

// LobbyEvent is either Peer Update or Exit in Topic
type LobbyEvent struct {
	ID     peer.ID
	Peer   *ct.Peer
	isExit bool
}

// newLobbyEvent Creates a new LobbyEvent
func newLobbyEvent(i peer.ID, p *ct.Peer) *LobbyEvent {
	if p == nil {
		return &LobbyEvent{
			ID:     i,
			isExit: true,
		}
	}
	return &LobbyEvent{
		ID:     i,
		Peer:   p,
		isExit: false,
	}
}
