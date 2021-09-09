package exchange

import (
	"context"
	"fmt"
	"time"

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
func NewProtocol(ctx context.Context, host *host.SHost, loc *common.Location, em *emitter.Emitter) (*ExchangeProtocol, error) {
	// Create PubSub Value Store
	olc := loc.OLC(6)
	r, err := psr.NewPubsubValueStore(ctx, host.Host, host.Pubsub(), ExchangeValidator{}, psr.WithRebroadcastInterval(10*time.Second))
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
		ctx:              ctx,
		host:             host,
		emitter:          em,
		PubsubValueStore: r,
		topic:            topic,
		subscription:     sub,
		eventHandler:     handler,
		exchangeEvents:   make(chan *ExchangeEvent),
		olc:              olc,
	}

	//go exchProtocol.handleExchangeEvents(exchProtocol.ctx)
	//go exchProtocol.handleExchangeMessages(exchProtocol.ctx)
	return exchProtocol, nil
}

// Search peer Profile by name
func (p *ExchangeProtocol) Search(sName string) (*common.Peer, error) {
	// Find peer from sName in the store
	buf, err := p.PubsubValueStore.GetValue(p.ctx, fmt.Sprintf("store/%s", sName))
	if err != nil {
		logger.Error("Failed to GET peer from store", zap.Error(err))
		return nil, err
	}

	// Unmarshal Peer from buffer
	profile := &common.Peer{}
	err = proto.Unmarshal(buf, profile)
	if err != nil {
		logger.Error("Failed to Unmarshal Peer", zap.Error(err))
		return nil, err
	}
	return profile, nil
}

func (p *ExchangeProtocol) Update(sName string, buf []byte) error {
	// Determine Key and Add Value to Store
	err := p.PubsubValueStore.PutValue(p.ctx, fmt.Sprintf("store/%s", sName), buf)
	if err != nil {
		logger.Error("Failed to PUT peer from store", zap.Error(err))
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
