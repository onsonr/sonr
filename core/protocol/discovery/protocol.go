package discovery

import (
	"context"
	"log"

	"github.com/kataras/golog"

	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/internal/node"
	"github.com/sonr-io/sonr/pkg/common"
)

var (
	logger             = golog.Child("protocols/discover")
	ErrParameters      = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer     = errors.New("Peer object provided to ExchangeProtocol is Nil")
	ErrTopicNotCreated = errors.New("Lobby Topic has not been Created")
)

// DiscoverProtocol handles Global and Local Sonr Peer Exchange Protocol
type DiscoverProtocol struct {
	node     node.Node
	ctx      context.Context
	local    *Local
	callback common.MotorCallback
}

// New creates new DiscoverProtocol
func New(ctx context.Context, host node.Node, cb common.MotorCallback, options ...Option) (*DiscoverProtocol, error) {
	// Create Exchange Protocol
	protocol := &DiscoverProtocol{
		ctx:      ctx,
		callback: cb,
		node:     host,
	}

	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	opts.Apply(protocol)
	log.Println("✅  DiscoverProtocol is Activated")
	return protocol, nil
}

// Update method publishes peer data to the topic
func (p *DiscoverProtocol) Update() error {

	// // Verify Peer is not nil
	// peer, err := p.node.Peer()
	// if err != nil {
	// 	return err
	// }

	// // Publish Event
	// err = p.local.Publish(peer)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// Close method closes the ExchangeProtocol
func (p *DiscoverProtocol) Close() error {
	p.local.eventHandler.Cancel()
	p.local.subscription.Cancel()
	err := p.local.topic.Close()
	if err != nil {
		logger.Errorf("%s - Failed to Close Local Lobby Topic for Exchange", err)
	}
	return nil
}
