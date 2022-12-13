package node

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
)

// OnPeerEvent is called when a peer joins or exits the channel
type OnPeerEvent func(e *ps.PeerEvent)

// OnMessage is called when a message is received
type OnMessage func(msg *ps.Message)

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
	onMessage         OnMessage
	onPeerEvent       OnPeerEvent
}

// Join joins a Channel interface with an underlying pubsub topic and event handler
func (n *hostImpl) Join(name string, opts ...ChannelOption) (*Channel, error) {
	// Check if PubSub is Set
	if n.ps == nil {
		return nil, errors.New("NewTopic: Pubsub has not been set on SNRHost")
	}

	config := defaultChannelConfig
	for _, opt := range opts {
		opt(&config)
	}

	// Call Underlying Pubsub to Connect
	t, err := n.JoinTopic(name, config.psOptions...)
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
		ctx:               n.ctx,
		selfId:            n.host.ID(),
		onMessage:         config.onMessageFunc,
		onPeerEvent:       config.onPeerEventFunc,
	}

	// Handle Messages
	go c.handleMessages()
	go c.handleSubscription()
	go c.handleEvents()
	return c, nil
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

// Returning a list of peers connected to the channel
func (c *Channel) ListPeers() []ID {
	pids := c.Topic.ListPeers()
	ids := make([]ID, len(pids))
	for i, pid := range pids {
		ids[i], _ = ParseID(pid)
	}
	return ids
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

// handleMessages handles messages received from the channel
func (c *Channel) handleMessages() {
	for {
		select {
		case msg := <-c.messages:
			c.DataStream <- msg.Data
		case <-c.ctx.Done():
			return
		}
	}
}

// handleSubscription handles messages received from the Topic
func (c *Channel) handleSubscription() {
	for {
		msg, err := c.Subscription.Next(c.ctx)
		if err != nil {
			return
		}
		c.onMessage(msg)
		c.messages <- msg
	}
}

// Handling the events from the channel
func (c *Channel) handleEvents() {
	for {
		event, err := c.TopicEventHandler.NextPeerEvent(c.ctx)
		if err != nil {
			return
		}
		c.onPeerEvent(&event)
		c.events <- &event
	}
}

// channelConfig is a struct that contains the configuration for a channel
type channelConfig struct {
	onMessageFunc   OnMessage
	onPeerEventFunc OnPeerEvent
	psOptions       []ps.TopicOpt
}

// defaultChannelConfig is the default configuration for a channel
var defaultChannelConfig = channelConfig{
	onMessageFunc:   func(msg *ps.Message) {},
	onPeerEventFunc: func(e *ps.PeerEvent) {},
	psOptions:       []ps.TopicOpt{},
}

// ChannelOption is a function that configures a channel
type ChannelOption func(*channelConfig)

// WithOnMessage sets the OnMessage callback for a channel
func WithOnMessage(callback OnMessage) ChannelOption {
	return func(c *channelConfig) {
		c.onMessageFunc = callback
	}
}

// WithOnPeerEvent sets the OnPeerEvent callback for a channel
func WithOnPeerEvent(callback OnPeerEvent) ChannelOption {
	return func(c *channelConfig) {
		c.onPeerEventFunc = callback
	}
}

// WithPSOptions sets the pubsub options for a channel
func WithPSOptions(opts ...ps.TopicOpt) ChannelOption {
	return func(c *channelConfig) {
		c.psOptions = opts
	}
}
