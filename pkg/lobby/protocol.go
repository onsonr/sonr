package lobby

import (
	"context"
	"errors"

	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_LIST_REFRESH = "lobby-list-refresh"
)

var (
	ErrParameters      = errors.New("Failed to create new LobbyProtocol, invalid parameters")
	ErrInvalidPeer     = errors.New("Peer object provided to LobbyProtocol is Nil")
	ErrTopicNotCreated = errors.New("Lobby Topic has not been Created")
)

// LobbyProtocol is the protocol for managing local peers.
type LobbyProtocol struct {
	node         api.NodeImpl
	ctx          context.Context
	host         *host.SNRHost // host
	eventHandler *ps.TopicEventHandler
	messages     chan *LobbyMessage
	subscription *ps.Subscription
	topic        *ps.Topic
	olc          string
	peers        []*common.Peer
}

// NewProtocol creates a new lobby protocol instance.
func NewProtocol(ctx context.Context, host *host.SNRHost, nu api.NodeImpl, options ...LobbyOption) (*LobbyProtocol, error) {
	opts := defaultLobbyOptions()
	for _, option := range options {
		option(opts)
	}

	olc := createOlc(opts.location)

	// Check parameters
	if err := checkParams(host); err != nil {
		logger.Errorf("%s - Failed to create LobbyProtocol", err)
		return nil, err
	}

	// Create Exchange Topic
	topic, err := host.Join(olc)
	if err != nil {
		logger.Errorf("%s - Failed to Join Local Pubsub Topic", err)
		return nil, err
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		logger.Errorf("%s - Failed to Subscribe to OLC Topic", err)
		return nil, err
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		logger.Errorf("%s - Failed to Get Event Handler", err)
		return nil, err
	}

	// Create Exchange Protocol
	lobProtocol := &LobbyProtocol{
		node:         nu,
		ctx:          ctx,
		host:         host,
		topic:        topic,
		subscription: sub,
		eventHandler: handler,
		olc:          olc,
		messages:     make(chan *LobbyMessage),
		peers:        make([]*common.Peer, 0),
	}

	// Handle Events
	go lobProtocol.HandleEvents()
	go lobProtocol.HandleMessages()

	// Return Protocol
	logger.Debug("✅  LobbyProtocol is Activated \n")
	return lobProtocol, nil
}

// Close closes the LobbyProtocol
func (p *LobbyProtocol) Close() error {
	p.eventHandler.Cancel()
	p.subscription.Cancel()
	err := p.topic.Close()
	if err != nil {
		// ignore
	}
	return nil
}

// Update method publishes peer data to the topic
func (p *LobbyProtocol) Update() error {
	// Verify Topic has been created
	if p.topic == nil {
		return ErrTopicNotCreated
	}

	// Verify Peer is not nil
	peer, err := p.node.Peer()
	if err != nil {
		return err
	}

	// Publish Event
	if buf := createLobbyMsgBuf(peer); buf != nil {
		err = p.topic.Publish(p.ctx, buf)
		if err != nil {
			logger.Errorf("%s - Failed to Publish Event", err)
			return err
		}
	}
	return nil
}

// HandleEvents method listens to Pubsub Events for room
func (p *LobbyProtocol) HandleEvents() {
	// Loop Events
	for {
		// Get next event
		event, err := p.eventHandler.NextPeerEvent(p.ctx)
		if err != nil {
			logger.Errorf("%s - Failed to Get Next Peer Event", err)
			return
		}

		// Check Event and Validate not User
		if event.Type == ps.PeerLeave && event.Peer != p.host.ID() {
			// Remove Peer, Emit Event
			if ok := p.removePeer(event.Peer); ok {
				p.callRefresh()
			}
			continue
		}

		// Check Event and Validate not User
		if event.Type == ps.PeerJoin && event.Peer != p.host.ID() {
			// Update Peer Data in Topic
			err := p.callUpdate()
			if err != nil {
				logger.Errorf("%s - Failed to send peer update to lobby topic", err)
				continue
			}
		}
	}
}

// HandleMessages method listens to Pubsub Messages for room
func (p *LobbyProtocol) HandleMessages() {
	// Loop Messages
	for {
		// Get next message
		msg, err := p.subscription.Next(p.ctx)
		if err != nil {
			return
		}

		// Check Message and Validate not User
		if msg.ReceivedFrom != p.host.ID() {
			// Unmarshal Message
			data := &LobbyMessage{}
			err = proto.Unmarshal(msg.Data, data)
			if err != nil {
				logger.Errorf("%s - Failed to Unmarshal Message", err)
				continue
			}

			// Update Peer, Emit Event
			if ok := p.updatePeer(msg.ReceivedFrom, data.GetPeer()); ok {
				p.callRefresh()
			}
		}
	}
}
