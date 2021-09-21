package lobby

import (
	"context"

	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_PEER_EXIT   = "exchange-peer-exit"
	Event_PEER_UPDATE = "exchange-peer-update"
	Event_PEER_JOIN   = "exchange-peer-join"
)

type LobbyProtocol struct {
	ctx          context.Context
	host         *host.SNRHost    // host
	emitter      *emitter.Emitter // Handle to signal when done
	eventHandler *ps.TopicEventHandler
	lobbyEvents  chan *common.LobbyEvent
	location     *common.Location
	subscription *ps.Subscription
	topic        *ps.Topic
}

// NewProtocol creates a new lobby protocol instance.
func NewProtocol(host *host.SNRHost, loc *common.Location, em *emitter.Emitter) (*LobbyProtocol, error) {
	// Create Exchange Topic
	topic, err := host.Pubsub().Join(loc.OLC(6))
	if err != nil {
		return nil, err
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, err
	}

	// Create Exchange Protocol
	lobProtocol := &LobbyProtocol{
		ctx:          context.Background(),
		host:         host,
		emitter:      em,
		topic:        topic,
		subscription: sub,
		eventHandler: handler,
		lobbyEvents:  make(chan *common.LobbyEvent),
		location:     loc,
	}

	// Handle Events and Return Protocol
	go lobProtocol.HandleEvents()
	go lobProtocol.HandleMessages()
	return lobProtocol, nil
}

func (p *LobbyProtocol) Update(peer *common.Peer) error {
	// Create Event
	event := &PublishEvent{
		Peer: peer,
		Olc:  p.location.OLC(6),
	}

	// Marshal Event
	eventBuf, err := proto.Marshal(event)
	if err != nil {
		logger.Error("Failed to Marshal Event", zap.Error(err))
		return err
	}

	// Publish Event
	err = p.topic.Publish(p.ctx, eventBuf)
	if err != nil {
		logger.Error("Failed to Publish Event", zap.Error(err))
		return err
	}
	return nil
}

// HandleEvents method listens to Pubsub Events for room
func (p *LobbyProtocol) HandleEvents() {
	go func() {
		// Loop Events
		for {
			// Get next event
			event, err := p.eventHandler.NextPeerEvent(p.ctx)
			if err != nil {
				p.eventHandler.Cancel()
				return
			}

			// Check Event and Validate not User
			if p.isEventJoin(event) {
				// Handle Join Event
				continue
			} else if p.isEventExit(event) {
				continue
			}
		}
	}()
}

// HandleMessages method listens to Pubsub Messages for room
func (p *LobbyProtocol) HandleMessages() {
	go func() {
		// Loop Messages
		for {
			// Get next message
			msg, err := p.subscription.Next(p.ctx)
			if err != nil {
				p.subscription.Cancel()
				return
			}

			// Check Message and Validate not User
			if msg.ReceivedFrom == p.host.ID() {
				continue
			} else {
				// Unmarshal Message
				event := &PublishEvent{}
				err = proto.Unmarshal(msg.Data, event)
				if err != nil {
					logger.Error("Failed to Unmarshal Message", zap.Error(err))
					continue
				}

				// Emit Event
				p.emitter.Emit(Event_PEER_UPDATE, event)
			}
		}
	}()
}
