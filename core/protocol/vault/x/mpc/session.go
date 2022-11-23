package mpc

import (
	"context"
	"fmt"

	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/internal/node"
	common "github.com/sonr-io/sonr/pkg/common"
	ct "github.com/sonr-io/sonr/pkg/common/v1"
	st "github.com/sonr-io/sonr/core/protocol/discovery/types/v1"
)

// ErrFunc is a function that returns an error
type ErrFunc func() error

// sessionImpl is the protocol for managing local peers.
type sessionImpl struct {
	callback     common.MotorCallback
	node         node.Node
	ctx          context.Context
	eventHandler *ps.TopicEventHandler
	messages     chan *LobbyEvent
	subscription *ps.Subscription
	topic        *ps.Topic
	olc          string
	peers        []party.ID
	selfID       peer.ID
	updateFunc   ErrFunc
}

// Initializing the local struct.
// func (e *DiscoverProtocol) initLocal(topic *ps.Topic, cb common.MotorCallback) error {

// 	// Subscribe to Room
// 	sub, err := topic.Subscribe()
// 	if err != nil {
// 		fmt.Errorf("%s - Failed to Subscribe to OLC Topic", err)
// 		return err
// 	}

// 	// Create Room Handler
// 	handler, err := topic.EventHandler()
// 	if err != nil {
// 		fmt.Errorf("%s - Failed to Get Event Handler", err)
// 		return err
// 	}

// 	// Create Local Struct
// 	e.local = &Local{
// 		ctx:          e.ctx,
// 		selfID:       e.node.HostID(),
// 		node:         e.node,
// 		updateFunc:   e.Update,
// 		topic:        topic,
// 		subscription: sub,
// 		eventHandler: handler,
// 		olc:          sub.Topic(),
// 		messages:     make(chan *LobbyEvent),
// 		peers:        make([]*ct.Peer, 0),
// 	}

// 	// Handle Events
// 	go e.local.handleSub()
// 	go e.local.handleTopic()
// 	go e.local.handleEvents()
// 	go e.local.autoPushUpdates()
// 	return nil
// }

// Publish publishes a LobbyMessage to the Local Topic
func (p *sessionImpl) Publish(data *ct.Peer) error {
	// Create Message Buffer
	buf := createLobbyMsgBuf(data)
	err := p.topic.Publish(p.ctx, buf)
	if err != nil {
		return err
	}
	return nil
}

// handleSub method listens to Pubsub Events for Local Topic
func (p *sessionImpl) handleSub() {
	// Loop Events
	for {
		// Get next event
		event, err := p.eventHandler.NextPeerEvent(p.ctx)
		if err != nil {
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
func (p *sessionImpl) handleTopic() {
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
			data := &st.NearbyChannelMessage{}
			err = data.Unmarshal(msg.Data)
			if err != nil {
				continue
			}
			p.messages <- newLobbyEvent(msg.ReceivedFrom, data.GetFrom())
		}
	}
}

// callUpdate publishes a LobbyMessage to the Local Topic
func (lp *sessionImpl) callUpdate() error {
	// Create Event
	fmt.Println("Sending Update to Lobby")
	err := lp.updateFunc()
	if err != nil {
		return err
	}
	return nil
}

// createLobbyMsgBuf Creates a new Message Buffer for Local Topic
func createLobbyMsgBuf(p *ct.Peer) []byte {
	// Marshal Event
	event := &st.NearbyChannelMessage{From: p}
	buf, err := event.Marshal()
	if err != nil {
		return nil
	}
	return buf
}

// hasPeerID Checks if Peer ID is in Local Topic
func (lp *sessionImpl) hasPeerID(id peer.ID) bool {
	for _, p := range lp.topic.ListPeers() {
		if p == id {
			return true
		}
	}
	return false
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
