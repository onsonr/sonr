package highway

import (
	context "context"
	"errors"
	"net"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/node/api"
	"github.com/sonr-io/core/pkg/discover"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/registry"
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
	HighwayStubServer
	api.CallbackImpl
	node api.NodeImpl

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	*discover.DiscoverProtocol
	*registry.RegistryProtocol
	*exchange.ExchangeProtocol

	// Channels
	// TransferProtocol - decisionEvents
	decisionEvents chan *api.DecisionEvent

	// LobbyProtocol - refreshEvents
	refreshEvents chan *api.RefreshEvent

	// MailboxProtocol - mailEvents
	mailEvents chan *api.MailboxEvent

	// TransferProtocol - inviteEvents
	inviteEvents chan *api.InviteEvent

	// TransferProtocol - progressEvents
	progressEvents chan *api.ProgressEvent

	// TransferProtocol - completeEvents
	completeEvents chan *api.CompleteEvent
}

// startHighwayStub creates a new Highway service stub for the node.
func NewHighwayStub(ctx context.Context, h *host.SNRHost, n api.NodeImpl, loc *common.Location, lst net.Listener) (*HighwayStub, error) {
	// Create the RPC Service
	var err error
	grpcServer := grpc.NewServer()
	stub := &HighwayStub{
		node:           n,
		ctx:            ctx,
		grpcServer:     grpcServer,
		decisionEvents: make(chan *api.DecisionEvent),
		refreshEvents:  make(chan *api.RefreshEvent),
		inviteEvents:   make(chan *api.InviteEvent),
		mailEvents:     make(chan *api.MailboxEvent),
		progressEvents: make(chan *api.ProgressEvent),
		completeEvents: make(chan *api.CompleteEvent),
	}

	// Set Discovery Protocol
	stub.DiscoverProtocol, err = discover.New(ctx, h, n, stub, discover.WithLocation(loc))
	if err != nil {
		logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	stub.ExchangeProtocol, err = exchange.New(ctx, h, n, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	stub.RegistryProtocol, err = registry.New(ctx, h, n, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start ExchangeProtocol", err)
		return nil, err
	}
	// Register the RPC Service
	RegisterHighwayStubServer(grpcServer, stub)
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
			// s.LobbyProtocol.Close()
			// s.ExchangeProtocol.Close()
			return
		}
	}
}

// Authorize Signing Method Request for Data
func (s *HighwayStub) Authorize(ctx context.Context, req *api.AuthenticateRequest) (*api.AuthenticateResponse, error) {
	logger.Debug("HighwayStub.Authorize() is Unimplemented")
	return nil, nil
}

// Link a new Device to the Node
func (s *HighwayStub) Link(ctx context.Context, req *api.LinkRequest) (*api.LinkResponse, error) {
	logger.Debug("HighwayStub.Link() is Unimplemented")
	return nil, nil
}

// Register a new domain with the Node on the highway
func (s *HighwayStub) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	// Get Values
	pfix := req.GetPrefix()
	name := req.GetSName()
	fprint := req.GetFingerprint()

	// Check Values
	if pfix == "" || name == "" || fprint == "" {
		return &api.RegisterResponse{
			Success: false,
			Error:   "Invalid request. One or more of the required fields are empty.",
		}, nil
	}

	// Create Record
	resp, err := s.RegistryProtocol.Register(req)
	if err != nil {
		return &api.RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Return Response
	return &api.RegisterResponse{
		Success: true,
		Records: resp,
	}, nil
}
