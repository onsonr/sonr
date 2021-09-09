package exchange

import (
	"context"
	"fmt"

	ps "github.com/libp2p/go-libp2p-pubsub"
	psr "github.com/libp2p/go-libp2p-pubsub-router"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_PEER_EXIT   = "exit"
	Event_PEER_UPDATE = "update"
)

// TransferProtocol type
type ExchangeProtocol struct {
	*psr.PubsubValueStore
	ctx            context.Context
	host           *host.SHost      // host
	emitter        *emitter.Emitter // Handle to signal when done
	eventHandler   *ps.TopicEventHandler
	exchangeEvents chan *ExchangeEvent
	olc            string
	subscription   *ps.Subscription
	topic          *ps.Topic
}

// NewProtocol creates new ExchangeProtocol
func NewProtocol(host *host.SHost, loc *common.Location, em *emitter.Emitter) (*ExchangeProtocol, error) {
	// Create PubSub Value Store
	olc := loc.OLC(6)
	r, err := psr.NewPubsubValueStore(context.Background(), host.Host, host.Pubsub(), ExchangeValidator{})
	if err != nil {
		return nil, err
	}

	// Create Exchange Topic
	topic, err := host.Pubsub().Join(olc)
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

	exchProtocol := &ExchangeProtocol{
		host:             host,
		emitter:          em,
		PubsubValueStore: r,
		topic:            topic,
		subscription:     sub,
		eventHandler:     handler,
		exchangeEvents:   make(chan *ExchangeEvent),
		olc:              olc,
	}

	go exchProtocol.handleExchangeEvents(context.Background())
	go exchProtocol.handleExchangeMessages(context.Background())
	return exchProtocol, nil
}

// Find peer by name
func (p *ExchangeProtocol) Find(sName string) (*common.Peer, error) {
	// Find peer from sName in the store
	buf, err := p.PubsubValueStore.GetValue(context.Background(), fmt.Sprintf("store/%s", sName))
	if err != nil {
		return nil, err
	}

	// Unmarshal Peer from buffer
	peer := &common.Peer{}
	err = proto.Unmarshal(buf, peer)
	if err != nil {
		return nil, err
	}
	return peer, nil
}

func (p *ExchangeProtocol) Update(sName string, buf []byte) error {
	// Determine Key and Add Value to Store
	err := p.PubsubValueStore.PutValue(context.Background(), fmt.Sprintf("store/%s", sName), buf)
	if err != nil {
		return err
	}
	return nil
}

// handleExchangeEvents method listens to Pubsub Events for room
func (p *ExchangeProtocol) handleExchangeEvents(ctx context.Context) {
	// Loop Events
	for {
		// Get next event
		event, err := p.eventHandler.NextPeerEvent(ctx)
		if err != nil {

			p.eventHandler.Cancel()
			return
		}

		// Check Event and Validate not User
		if p.isEventJoin(event) {
			// pbuf, err := proto.Marshal()
			// if err != nil {

			// 	continue
			// }
			// err = rm.Exchange(event.Peer, pbuf)
			// if err != nil {

			// 	continue
			// }
		} else if p.isEventExit(event) {

		}
		state.GetState().NeedsWait()
	}
}

// handleExchangeMessages method listens for messages on pubsub room subscription
func (p *ExchangeProtocol) handleExchangeMessages(ctx context.Context) {
	for {
		// Get next msg from pub/sub
		msg, err := p.subscription.Next(ctx)
		if err != nil {
			logger.Error("Failed to get next subcription message", zap.Error(err))
			return
		}

		// Only forward messages delivered by others
		if p.isValidMessage(msg) {
			// Unmarshal RoomEvent
			m := &ExchangeEvent{}
			err = proto.Unmarshal(msg.Data, m)
			if err != nil {
				logger.Error("Failed to Unmarshal Message", zap.Error(err))
				continue
			}

			// Check Peer is Online, if not ignore
			if m.GetPeer().GetStatus() == common.Peer_ONLINE {
				p.emitter.Emit(emitter.EMIT_ROOM_EVENT, m)
			}
		}
		state.GetState().NeedsWait()
	}
}
