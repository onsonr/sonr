package highway

import (
	context "context"
	"errors"
	"net"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/channel"
	"github.com/sonr-io/core/config"
	hn "github.com/sonr-io/core/host"
	"github.com/sonr-io/core/host/discover"
	"github.com/sonr-io/core/host/exchange"
	v1 "go.buf.build/grpc/go/sonr-io/core/highway/v1"
	"google.golang.org/grpc"

	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

// Error Definitions
var (
	logger                 = golog.Default.Child("node/highway")
	ErrEmptyQueue          = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery        = errors.New("No SName or PeerID provided.")
	ErrMissingParam        = errors.New("Paramater is missing.")
	ErrProtocolsNotSet     = errors.New("Node Protocol has not been initialized.")
	ErrMethodUnimplemented = errors.New("Method is not implemented.")
)

// HighwayServer is the RPC Service for the Custodian Node.
type HighwayServer struct {
	v1.HighwayServer
	config.CallbackImpl
	node   hn.HostImpl
	cosmos cosmosclient.Client

	// Properties
	ctx      context.Context
	listener net.Listener
	grpc     *grpc.Server
	*discover.DiscoverProtocol
	*exchange.ExchangeProtocol

	// Configuration
	// ipfs *storage.IPFSService

	// List of Entries
	channels map[string]channel.Channel
}

// NewHighwayServer creates a new Highway service stub for the node.
func NewHighway(ctx context.Context, opts ...hn.Option) (*HighwayServer, error) {
	// Create a new HostImpl
	node, err := hn.NewHost(ctx, config.Role_HIGHWAY, opts...)
	if err != nil {
		return nil, err
	}

	// // Set IPFS Service
	// stub.ipfs, err = storage.Init()
	// if err != nil {
	// 	return nil, err
	// }

	lst, err := node.Listener()
	if err != nil {
		return nil, err
	}

	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("snr"))
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	stub := &HighwayServer{
		node:     node,
		ctx:      ctx,
		grpc:     grpc.NewServer(),
		cosmos:   cosmos,
		listener: lst,
	}

	// Set Discovery Protocol
	stub.DiscoverProtocol, err = discover.New(ctx, node, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	stub.ExchangeProtocol, err = exchange.New(ctx, node, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Register RPC Service
	v1.RegisterHighwayServer(stub.grpc, stub)
	return stub, nil
}

// Serve starts the RPC Service.
func (s *HighwayServer) Serve() {
	logger.Infof("Starting RPC Server on %s", s.listener.Addr().String())
	go s.serveCtxListener(s.ctx, s.listener)
}

// Serve serves the RPC Service on the given port.
func (s *HighwayServer) serveCtxListener(ctx context.Context, listener net.Listener) {
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := s.grpc.Serve(listener); err != nil {
		logger.Errorf("%s - Failed to start HTTP server", err)
	}
	s.node.Persist()
}
