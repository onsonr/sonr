package topic

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	sh "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/data"
)

type OnPeerJoined func(ps.PeerEvent)
type OnPeerLeft func(ps.PeerEvent)
type OnMessage func(*ps.Message)

type TopicListener struct {
	OnPeerJoined OnPeerJoined
	OnPeerLeft   OnPeerLeft
	OnMessage    OnMessage
}

type Topic interface {
	Publish(msg []byte) error
	Close()
	HasPeer(q string) bool
	FindPeer(q string) (peer.ID, error)
}

type topic struct {
	Topic
	ctx          context.Context
	hostNode     sh.HostNode
	subscription *ps.Subscription
	handler      *ps.TopicEventHandler
	room         string
	topic        *ps.Topic
	listener     TopicListener
}

func NewTopic(ctx context.Context, host sh.HostNode, room string, listener TopicListener) (Topic, error) {
	// Join Room
	t, err := host.Pubsub().Join(room)
	if err != nil {
		return nil, err
	}

	// Subscribe to Room
	subscription, err := t.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Room Handler
	handler, err := t.EventHandler()
	if err != nil {
		return nil, err
	}

	topicObject := &topic{
		ctx:          ctx,
		hostNode:     host,
		room:         room,
		subscription: subscription,
		handler:      handler,
		topic:        t,
		listener:     listener,
	}

	go topicObject.handleEvents(topicObject.ctx)
	go topicObject.handleMessages(topicObject.ctx)
	return topicObject, nil
}

func (t *topic) FindPeer(q string) (peer.ID, error) {
	// Iterate through Room Peers
	for _, id := range t.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return id, nil
		}
	}
	return "", fmt.Errorf("Given ID (%s) was not found in Topic: \n\t- Name: %s \n\t- Active Peers: %v", q, t.room, len(t.topic.ListPeers()))
}

func (t *topic) HasPeer(q string) bool {
	// Iterate through PubSub in room
	for _, id := range t.topic.ListPeers() {
		// If Found Match
		if id.String() == q {
			return true
		}
	}
	return false
}

// Publish method publishes message to pubsub room
func (t *topic) Publish(msg []byte) error {
	return t.topic.Publish(t.ctx, msg)
}

// Close method closes pubsub room
func (t *topic) Close() {
	t.handler.Cancel()
	t.subscription.Cancel()
	t.topic.Close()
	t.ctx.Done()
}

// handleExchangeEvents method listens to Pubsub Events for room
func (t *topic) handleEvents(ctx context.Context) {
	// Loop Events
	for {
		// Get next event
		event, err := t.handler.NextPeerEvent(ctx)
		if err != nil {
			data.LogError(err)
			t.Close()
			return
		}

		// Check Event Type
		if event.Type == ps.PeerJoin {
			// Call listener for Peer Joined
			if t.listener.OnPeerJoined != nil {
				t.listener.OnPeerJoined(event)
			}
		} else if event.Type == ps.PeerLeave {
			// Call listener for Peer Left
			if t.listener.OnPeerLeft != nil {
				t.listener.OnPeerLeft(event)
			}
		}

		// Check State
		data.GetState().NeedsWait()
	}
}

// handleExchangeMessages method listens for messages on pubsub room subscription
func (t *topic) handleMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := t.subscription.Next(ctx)
		if err != nil {
			data.LogError(err)
			t.Close()
			return
		}

		// Check Message not from self
		if msg.GetFrom() != t.hostNode.ID() {
			// Call listener for Message
			if t.listener.OnMessage != nil {
				t.listener.OnMessage(msg)
			}
		}
		data.GetState().NeedsWait()
	}
}
