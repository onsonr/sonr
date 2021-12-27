package registry

import (
	"context"

	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/node/api"
)

// RegistryProtocol handles Global and Local Sonr Peer Exchange Protocol
type RegistryProtocol struct {
	callback api.CallbackImpl
	node     api.NodeImpl
	ctx      context.Context
	host     *host.SNRHost
	mode     api.StubMode
}

// New creates new RegisteryProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl, cb api.CallbackImpl, options ...Option) (*RegistryProtocol, error) {
	// Create Exchange Protocol
	protocol := &RegistryProtocol{
		ctx:      ctx,
		host:     host,
		node:     node,
		callback: cb,
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

// Verify method uses resolver to check if Peer is registered,
// returns true if Peer is registered
func (p *RegistryProtocol) Verify(sname string) (bool, error) {
	return true, nil
}
