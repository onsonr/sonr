package node

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
)

// Channel is a wrapper around a pubsub topic
type Channel struct {
	Name              string
	Topic             *ps.Topic
	TopicEventHandler *ps.TopicEventHandler
	Subscription      *ps.Subscription
	DataStream        chan []byte
	messages          chan *ps.Message
	events            chan *ps.PeerEvent
	ctx               context.Context
	selfId            peer.ID
}

// NewChannel joins a Channel interface with an underlying pubsub topic and event handler
func (n *hostImpl) NewChannel(ctx context.Context, name string, opts ...ps.TopicOpt) (*Channel, error) {
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil, errors.New("NewTopic: Pubsub has not been set on SNRHost")
	}

	// Call Underlying Pubsub to Connect
	t, err := n.Join(name, opts...)
	if err != nil {
		return nil, err
	}

	// Create Event Handler
	h, err := t.EventHandler()
	if err != nil {
		return nil, err
	}

	// Create Subscriber
	s, err := t.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Channel
	c := &Channel{
		Name:              name,
		Topic:             t,
		TopicEventHandler: h,
		Subscription:      s,
		messages:          make(chan *ps.Message),
		events:            make(chan *ps.PeerEvent),
		ctx:               ctx,
		selfId:            n.host.ID(),
	}

	// Handle Messages
	go c.handleMessages()
	go c.handleEvents()
	return c, nil
}

// Close closes the channel
func (c *Channel) Close() error {
	// Close Topic
	err := c.Topic.Close()
	if err != nil {
		return err
	}

	// Close Channel
	close(c.messages)
	close(c.events)
	return nil
}

// Send sends a message to the channel
func (c *Channel) Send(data []byte) error {
	return c.Topic.Publish(c.ctx, data)
}

// NextEvent returns the next event from the channel
func (c *Channel) NextEvent() <-chan *ps.PeerEvent {
	return c.events
}

// NextMessage returns the next message from the channel
func (c *Channel) NextMessage() <-chan *ps.Message {
	return c.messages
}

// ListPeers returns a list of peers connected to the channel
func (c *Channel) ListPeers() []peer.ID {
	return c.Topic.ListPeers()
}

func (c *Channel) handleMessages() {
	for {
		msg, err := c.Subscription.Next(c.ctx)
		if err != nil {
			return
		}
		d := msg.Data
		c.DataStream <- d
		c.messages <- msg
	}
}

func (c *Channel) handleEvents() {
	for {
		event, err := c.TopicEventHandler.NextPeerEvent(c.ctx)
		if err != nil {
			return
		}
		c.events <- &event
	}
}
