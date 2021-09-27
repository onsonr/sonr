package lobby

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
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
	Event_LIST_REFRESH = "lobby-list-refresh"
)

type LobbyProtocol struct {
	ctx          context.Context
	host         *host.SNRHost    // host
	emitter      *emitter.Emitter // Handle to signal when done
	eventHandler *ps.TopicEventHandler
	lobbyEvents  chan *LobbyMessage
	location     *common.Location
	subscription *ps.Subscription
	topic        *ps.Topic
	peers        map[peer.ID]*common.Peer
}

// NewProtocol creates a new lobby protocol instance.
func NewProtocol(host *host.SNRHost, loc *common.Location, em *emitter.Emitter) (*LobbyProtocol, error) {
	// Create Exchange Topic
	topic, err := host.Pubsub().Join(loc.OLC())
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
		lobbyEvents:  make(chan *LobbyMessage),
		location:     loc,
		peers:        make(map[peer.ID]*common.Peer),
	}

	// Handle Events and Return Protocol
	go lobProtocol.HandleEvents()
	go lobProtocol.HandleMessages()
	return lobProtocol, nil
}

func (p *LobbyProtocol) Update(peer *common.Peer) error {
	// Create Event
	event := &LobbyMessage{
		Peer: peer,
		Olc:  p.location.OLC(),
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
			if p.isEventExit(event) {
				// Remove Peer from map
				delete(p.peers, event.Peer)
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
					logger.Error("Failed to Unmarshal Message", zap.Error(err))
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
	// Add Peer to map
	p.peers[id] = peer

	// Create Peer List from map
	peers := make([]*common.Peer, 0, len(p.peers))
	for _, peer := range p.peers {
		peers = append(peers, peer)
	}

	// Create RefreshEvent
	event := &common.RefreshEvent{
		Olc:   p.location.OLC(),
		Peers: peers,
	}

	// Emit Event
	p.emitter.Emit(Event_LIST_REFRESH, event)
}
