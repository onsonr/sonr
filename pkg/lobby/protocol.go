package lobby

import (
	"context"
	"errors"
	"time"

	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
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
	api.NodeImpl
	ctx          context.Context
	host         *host.SNRHost  // host
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
		logger.Error("Failed to create LobbyProtocol", err)
		return nil, err
	}

	// Create Exchange Topic
	topic, err := host.Join(olc)
	if err != nil {
		logger.Error("Failed to Join Local Pubsub Topic", err)
		return nil, err
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		logger.Error("Failed to Subscribe to OLC Topic", err)
		return nil, err
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		logger.Error("Failed to Get Event Handler", err)
		return nil, err
	}

	// Create Exchange Protocol
	lobProtocol := &LobbyProtocol{
		NodeImpl:     nu,
		ctx:          ctx,
		host:         host,
		topic:        topic,
		subscription: sub,
		eventHandler: handler,
		olc:          olc,
		messages:     make(chan *LobbyMessage),
		peers:        make([]*common.Peer, 0),
	}

	// Handle Events and Return Protocol
	logger.Info("✅  LobbyProtocol is Activated \n")
	go lobProtocol.HandleEvents()
	go lobProtocol.HandleMessages()
	go lobProtocol.autoPushUpdates()
	return lobProtocol, nil
}

// Update method publishes peer data to the topic
func (p *LobbyProtocol) Update(peer *common.Peer) error {
	// Verify Topic has been created
	if p.topic == nil {
		return ErrTopicNotCreated
	}

	// Verify Peer is not nil
	if peer == nil {
		return ErrInvalidPeer
	}

	// Create Event
	event := &LobbyMessage{Peer: peer}

	// Marshal Event
	eventBuf, err := proto.Marshal(event)
	if err != nil {
		logger.Error("Failed to Marshal Event", err)
		return err
	}

	// Publish Event
	err = p.topic.Publish(p.ctx, eventBuf)
	if err != nil {
		logger.Error("Failed to Publish Event", err)
		return err
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
			p.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if p.isEventExit(event) {
			p.pushRefresh(event.Peer, nil)
			continue
		} else {
			// Update Peer Data in Topic
			err := p.sendUpdate()
			if err != nil {
				logger.Error("Failed to send peer update to lobby topic", err)
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
			p.subscription.Cancel()
			return
		}

		// Check Message and Validate not User
		if msg.ReceivedFrom != p.host.ID() {
			// Unmarshal Message
			data := &LobbyMessage{}
			err = proto.Unmarshal(msg.Data, data)
			if err != nil {
				logger.Error("Failed to Unmarshal Message", err)
				continue
			}

			// Update Peer Data in map
			p.pushRefresh(msg.ReceivedFrom, data.Peer)
		}
	}
}

// HandleMessages method listens to Pubsub Messages for room
func (p *LobbyProtocol) autoPushUpdates() {
	// Loop Messages
	for {
		err := p.sendUpdate()
		if err != nil {
			logger.Error("Failed to send peer update to lobby topic", err)
			continue
		}
		time.Sleep(time.Second * 5)
	}
}
