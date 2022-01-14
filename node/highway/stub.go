package highway

import (
	context "context"
	"errors"
	"net"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/node"
	"github.com/sonr-io/core/node/highway/v1"
	"github.com/sonr-io/core/protocols/discover"
	"github.com/sonr-io/core/protocols/exchange"

	"google.golang.org/grpc"
)

// Error Definitions
var (
	logger             = golog.Default.Child("node/highway")
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrMissingParam    = errors.New("Paramater is missing.")
	ErrProtocolsNotSet = errors.New("Node Protocol has not been initialized.")
)

// HighwayStub is the RPC Service for the Custodian Node.
type HighwayStub struct {
	highway.HighwayServiceServer
	node.CallbackImpl
	node node.NodeImpl

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	*discover.DiscoverProtocol
	*exchange.ExchangeProtocol
}

// startHighwayStub creates a new Highway service stub for the node.
func NewHighwayStub(ctx context.Context, n node.NodeImpl, loc *common.Location, lst net.Listener) (*HighwayStub, error) {
	// Create the RPC Service
	var err error
	grpcServer := grpc.NewServer()
	stub := &HighwayStub{
		node:       n,
		ctx:        ctx,
		grpcServer: grpcServer,
	}

	// Set Discovery Protocol
	stub.DiscoverProtocol, err = discover.New(ctx, n, stub, discover.WithLocation(loc))
	if err != nil {
		logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	stub.ExchangeProtocol, err = exchange.New(ctx, n, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Register the RPC Service
	highway.RegisterHighwayServiceServer(grpcServer, stub)
	go stub.Serve(ctx, lst)
	return stub, nil
}

// Serve serves the RPC Service on the given port.
func (s *HighwayStub) Serve(ctx context.Context, listener net.Listener) {
	// Handle Node Events
	if err := s.grpcServer.Serve(listener); err != nil {
		logger.Error("Failed to serve gRPC", err)
	}
	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			s.grpcServer.Stop()
			s.DiscoverProtocol.Close()
			return
		}
	}
}
