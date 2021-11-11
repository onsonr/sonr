package discover

import (
	"context"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/beam"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/api"
	"github.com/sonr-io/core/x/core/common"
	"google.golang.org/protobuf/proto"
)

var (
	logger             = golog.Child("protocols/discover")
	ErrParameters      = errors.New("Failed to create new ExchangeProtocol, invalid parameters")
	ErrInvalidPeer     = errors.New("Peer object provided to ExchangeProtocol is Nil")
	ErrTopicNotCreated = errors.New("Lobby Topic has not been Created")
)

// DiscoverProtocol handles Global and Local Sonr Peer Exchange Protocol
type DiscoverProtocol struct {
	node   api.NodeImpl
	ctx    context.Context
	global beam.Beam
	host   *host.SNRHost
	local  *Local
	mode   api.StubMode
}

// New creates new DiscoveryProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl, options ...Option) (*DiscoverProtocol, error) {

	// Create BeamStore
	b, err := beam.New(ctx, host, beam.ID("global"))
	if err != nil {
		return nil, err
	}

	// Create Exchange Protocol
	protocol := &DiscoverProtocol{
		ctx:    ctx,
		global: b,
		host:   host,
		node:   node,
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

// FindPeerId method returns PeerID by SName
func (p *DiscoverProtocol) Get(sname string) (*common.Peer, error) {
	peer := &common.Peer{}
	// Get Peer from KadDHT store
	if buf, err := p.global.Get(sname); err == nil {
		// Unmarshal Peer
		err := proto.Unmarshal(buf, peer)
		if err != nil {
			logger.Errorf("%s - Failed to unmarshal Peer", err)
			return nil, err
		}
		return peer, nil
	} else {
		logger.Warn("Failed to get Peer from BeamStore: %s", err)
		return nil, err
	}
}

// Put method updates peer instance in the store
func (p *DiscoverProtocol) Put(peer *common.Peer) error {
	logger.Debug("Updating Peer in BeamStore")
	// Marshal Peer
	buf, err := peer.Buffer()
	if err != nil {
		logger.Errorf("Failed to Marshal Peer: %s", err)
		return err
	}

	// Add Peer to KadDHT store
	err = p.global.Put(peer.GetSName(), buf)
	if err != nil {
		logger.Errorf("Failed to put item in BeamStore: %s", err)
		return err
	}
	return nil
}

// Update method publishes peer data to the topic
func (p *DiscoverProtocol) Update() error {
	if p.mode.Motor() {
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
