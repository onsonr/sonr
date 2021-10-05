package lobby

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_LIST_REFRESH = "lobby-list-refresh"
)

// LobbyProtocol is the protocol for managing local peers.
type LobbyProtocol struct {
	ctx          context.Context
	host         *host.SNRHost  // host
	emitter      *state.Emitter // Handle to signal when done
	eventHandler *ps.TopicEventHandler
	lobbyEvents  chan *LobbyMessage
	location     *common.Location
	subscription *ps.Subscription
	topic        *ps.Topic
	peers        map[peer.ID]*common.Peer
}

// NewProtocol creates a new lobby protocol instance.
func NewProtocol(ctx context.Context, host *host.SNRHost, loc *common.Location, em *state.Emitter) (*LobbyProtocol, error) {
	// Create Exchange Topic
	topic, err := host.Pubsub().Join(loc.OLC())
	if err != nil {
		return nil, logger.Error("Failed to Join Local Pubsub Topic", err)
	}

	// Subscribe to Room
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, logger.Error("Failed to Subscribe to OLC Topic", err)
	}

	// Create Room Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, logger.Error("Failed to Get Event Handler", err)
	}

	// Create Exchange Protocol
	lobProtocol := &LobbyProtocol{
		ctx:          ctx,
		host:         host,
		emitter:      em,
		topic:        topic,
		subscription: sub,
		eventHandler: handler,
		lobbyEvents:  make(chan *LobbyMessage),
		location:     loc,
		peers:        make(map[peer.ID]*common.Peer),
	}

	// Handle Events and Return Protocol
	go lobProtocol.HandleEvents()
	go lobProtocol.HandleMessages()
	return lobProtocol, nil
}

// Update method publishes peer data to the topic
func (p *LobbyProtocol) Update(peer *common.Peer) error {
	// Create Event
	event := &LobbyMessage{Peer: peer}

	// Marshal Event
	eventBuf, err := proto.Marshal(event)
	if err != nil {
		return logger.Error("Failed to Marshal Event", err)
	}

	// Publish Event
	err = p.topic.Publish(p.ctx, eventBuf)
	if err != nil {
		return logger.Error("Failed to Publish Event", err)
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
			if p.isEventExit(event) {
				p.pushRefresh(event.Peer, nil)
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
	}()
}

// pushRefresh sends a refresh event to the emitter
func (p *LobbyProtocol) pushRefresh(id peer.ID, peer *common.Peer) {
	// Check if Peer was provided
	if peer == nil {
		// Remove Peer from map
		delete(p.peers, id)
	} else {
		// Add Peer to map
		p.peers[id] = peer
	}

	// Create Peer List from map
	peers := make([]*common.Peer, 0, len(p.peers))
	for _, peer := range p.peers {
		peers = append(peers, peer)
	}

	// Create RefreshEvent
	event := &common.RefreshEvent{
		Olc:      p.location.OLC(),
		Peers:    peers,
		Received: int64(time.Now().Unix()),
	}

	// Emit Event
	p.emitter.Emit(Event_LIST_REFRESH, event)
}
