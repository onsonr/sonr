package linking

import (
	"context"
	"log"
	"time"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/third_party/types/common"

	host "github.com/sonr-io/sonr/pkg/host"
)

var (
	logger             = golog.Child("protocols/discover")
	ErrParameters      = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer     = errors.New("Peer object provided to ExchangeProtocol is Nil")
	ErrTopicNotCreated = errors.New("Lobby Topic has not been Created")
)

// LinkingProtocol handles Global and Local Sonr Peer Exchange Protocol
type LinkingProtocol struct {
	node     host.SonrHost
	ctx      context.Context
	session  *Session
	mode     config.Role
	callback common.MotorCallback
	timeout  time.Duration
}

// New creates new DiscoveryProtocol
func New(ctx context.Context, host host.SonrHost, cb common.MotorCallback, options ...Option) (*LinkingProtocol, error) {
	// Create Exchange Protocol
	protocol := &LinkingProtocol{
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
	log.Println("✅  ExchangeProtocol is Activated")
	return protocol, nil
}

// Update method publishes peer data to the topic
func (p *LinkingProtocol) Update() error {
	if p.mode.IsMotor() {
		// Verify Peer is not nil
		peer, err := p.node.Peer()
		if err != nil {
			return err
		}

		// Publish Event
		err = p.session.Publish(peer)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close method closes the ExchangeProtocol
func (p *LinkingProtocol) Close() error {
	p.session.eventHandler.Cancel()
	p.session.subscription.Cancel()
	err := p.session.topic.Close()
	if err != nil {
		logger.Errorf("%s - Failed to Close Local Lobby Topic for Exchange", err)
	}
	return nil
}
