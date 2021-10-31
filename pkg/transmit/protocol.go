package transmit

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
)

// TransmitProtocol type
type TransmitProtocol struct {
	node    api.NodeImpl
	ctx     context.Context // Context
	host    *host.SNRHost   // local host
	current *Session        // current session
	mode    api.StubMode
}

// New creates a new TransferProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl, options ...Option) (*TransmitProtocol, error) {
	// create a new transfer protocol
	protocol := &TransmitProtocol{
		ctx:  ctx,
		host: host,
		node: node,
	}
	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	opts.Apply(protocol)

	// Setup Stream Handlers
	host.SetStreamHandler(IncomingPID, protocol.onIncomingTransfer)
	logger.Debug("✅  TransmitProtocol is Activated \n")
	return protocol, nil
}

// CurrentSession returns the current session
func (p *TransmitProtocol) CurrentSession() (*Session, error) {
	if p.current != nil {
		return p.current, nil
	}
	return nil, ErrNoSession
}

// Incoming is called by the node to accept an incoming transfer
func (p *TransmitProtocol) Incoming(payload *common.Payload, from *common.Peer) error {
	// Get User Peer
	to, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get User Peer", err)
		return err
	}

	// Create New TransferEntry
	p.current = NewInSession(payload, from, to)
	return nil
}

// Outgoing is called by the node to initiate a transfer
func (p *TransmitProtocol) Outgoing(payload *common.Payload, to *common.Peer) error {
	// Get User Peer
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer", err)
		return err
	}

	// Get Id
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer ID", err)
		return err
	}

	// Create New TransferEntry
	p.current = NewOutSession(payload, from, to)

	// Create New Stream
	stream, err := p.host.NewStream(p.ctx, toId, IncomingPID)
	if err != nil {
		logger.Errorf("%s - Failed to Create New Stream", err)
		return err
	}

	// Start Transfer
	p.onOutgoingTransfer(stream)
	return nil
}

// Reset resets the current session
func (p *TransmitProtocol) Reset(event *api.CompleteEvent) {
	logger.Debug("Resetting TransmitProtocol")
	p.node.OnComplete(event)
	p.current = nil
}

// onIncomingTransfer incoming transfer handler
func (p *TransmitProtocol) onIncomingTransfer(stream network.Stream) {
	logger.Debug("Received Incoming Transfer")
	// Find Entry in Queue
	entry, err := p.CurrentSession()
	if err != nil {
		logger.Errorf("%s - Failed to find transfer request", err)
		stream.Close()
		return
	}

	// Create New Reader
	event, err := entry.ReadFrom(stream, p.node)
	if err != nil {
		logger.Errorf("%s - Failed to Read From Stream", err)
		stream.Reset()
		return
	}
	p.Reset(event)
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransmitProtocol) onOutgoingTransfer(stream network.Stream) {
	logger.Debug("Received Accept Decision, Starting Outgoing Transfer")

	// Find Entry in Queues
	entry, err := p.CurrentSession()
	if err != nil {
		logger.Errorf("%s - Failed to find transfer request", err)
		stream.Close()
		return
	}

	// Create New Writer
	event, err := entry.WriteTo(stream, p.node)
	if err != nil {
		logger.Errorf("%s - Failed to Write To Stream", err)
		stream.Reset()
		return
	}
	p.Reset(event)
}
