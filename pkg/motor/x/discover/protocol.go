package discover

import (
	"context"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/pkg/config"

	host "github.com/sonr-io/sonr/pkg/host"
)

var (
	logger             = golog.Child("protocols/discover")
	ErrParameters      = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer     = errors.New("Peer object provided to ExchangeProtocol is Nil")
	ErrTopicNotCreated = errors.New("Lobby Topic has not been Created")
)

// DiscoverProtocol handles Global and Local Sonr Peer Exchange Protocol
type DiscoverProtocol struct {
	node  host.SonrHost
	ctx   context.Context
	local *Local
	mode  config.Role
}

// New creates new DiscoveryProtocol
func New(ctx context.Context, host host.SonrHost, options ...Option) (*DiscoverProtocol, error) {
	// Create Exchange Protocol
	protocol := &DiscoverProtocol{
		ctx: ctx,

		node: host,
	}

	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	opts.Apply(protocol)
	logger.Debug("✅  ExchangeProtocol is Activated \n")
	return protocol, nil
}

// Update method publishes peer data to the topic
func (p *DiscoverProtocol) Update() error {
	if p.mode.IsMotor() {
		// Verify Peer is not nil
		peer, err := p.node.Peer()
		if err != nil {
			return err
		}

		// Publish Event
		err = p.local.Publish(peer)
		if err != nil {
			return err
		}
	}
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
