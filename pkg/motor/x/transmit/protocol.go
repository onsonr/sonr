package transmit

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/host"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	st "github.com/sonr-io/sonr/third_party/types/motor/api/v1/service/v1"
)

// TransmitProtocol type
type TransmitProtocol struct {
	node    host.SonrHost
	ctx     context.Context // Context
	current *st.Session     // current session
	mode    config.Role
}

// New creates a new TransferProtocol
func New(ctx context.Context, host host.SonrHost, options ...Option) (*TransmitProtocol, error) {
	// create a new transfer protocol
	protocol := &TransmitProtocol{
		ctx:  ctx,
		node: host,
	}
	// Set options
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	opts.Apply(protocol)

	// Setup Stream Handlers
	host.SetStreamHandler(FilePID, protocol.onIncomingTransfer)
	logger.Debug("✅  TransmitProtocol is Activated \n")
	return protocol, nil
}

// CurrentSession returns the current session
func (p *TransmitProtocol) CurrentSession() (*st.Session, error) {
	if p.current != nil {
		return p.current, nil
	}
	return nil, ErrNoSession
}

// Incoming is called by the node to accept an incoming transfer
func (p *TransmitProtocol) Incoming(payload *st.Payload, from *ct.Peer) error {
	// // Get User Peer
	// to, err := p.node.Peer()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Get User Peer", err)
	// 	return err
	// }

	// Create New TransferEntry
	p.current = NewInSession(payload, from, nil)
	return nil
}

// Outgoing is called by the node to initiate a transfer
func (p *TransmitProtocol) Outgoing(payload *st.Payload, to *ct.Peer) error {
	// // Get User Peer
	// from, err := p.node.Peer()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Get Peer", err)
	// 	return err
	// }

	// Get Id
	toId, err := Libp2pID(to)
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer ID", err)
		return err
	}

	// Create New TransferEntry
	p.current = NewOutSession(payload, nil, to)

	// Send Files
	if p.current.Payload.GetItems()[0].GetMime().Type != st.MIME_TYPE_URL {
		// Create New Stream
		stream, err := p.node.NewStream(p.ctx, toId, FilePID)
		if err != nil {
			logger.Errorf("%s - Failed to Create New Stream", err)
			return err
		}

		// Start Transfer
		p.onOutgoingTransfer(stream)
	}

	return nil
}

// // Reset resets the current session
// func (p *TransmitProtocol) Reset(event *v1.OnTransmitCompleteResponse) {
// 	logger.Debug("Resetting TransmitProtocol")
// 	p.node.Events().Emit(t.ON_COMPLETE, event)
// 	p.current = nil
// }

// onIncomingTransfer incoming transfer handler
func (p *TransmitProtocol) onIncomingTransfer(stream network.Stream) {
	logger.Debug("Received Incoming Transfer")
	// Find Entry in Queue
	// entry, err := p.CurrentSession()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to find transfer request", err)
	// 	stream.Close()
	// 	return
	// }

	// // Create New Reader
	// event, err := RouteSessionStream(entry, stream, p.node)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Read From Stream", err)
	// 	stream.Close()
	// 	return
	// }
	// p.Reset(event)
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransmitProtocol) onOutgoingTransfer(stream network.Stream) {
	logger.Debug("Received Accept Decision, Starting Outgoing Transfer")

	// Find Entry in Queues
	// entry, err := p.CurrentSession()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to find transfer request", err)
	// 	stream.Close()
	// 	return
	// }

	// Create New Writer
	// event, err := RouteSessionStream(entry, stream, p.node)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Write To Stream", err)
	// 	stream.Close()
	// 	return
	// }
	// p.Reset(event)
}

// Libp2pID returns the PeerID based on PublicKey from Profile
func Libp2pID(p *ct.Peer) (peer.ID, error) {
	// Return Peer ID
	id, err := peer.IDFromString(p.GetPeerId())
	if err != nil {
		return "", err
	}
	return id, nil
}
