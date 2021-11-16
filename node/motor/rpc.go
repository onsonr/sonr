package motor

import (
	context "context"
	"errors"
	"net"
	"strings"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/internal/host"
	api "github.com/sonr-io/core/node/api"
	"github.com/sonr-io/core/pkg/discover"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/transmit"
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

// MotorStub is the RPC Service for the Default Node.
type MotorStub struct {
	// Interfaces
	MotorStubServer
	api.CallbackImpl
	node api.NodeImpl

	// Properties
	ctx context.Context

	grpcServer *grpc.Server

	// Protocols
	*transmit.TransmitProtocol
	*discover.DiscoverProtocol
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

// startMotorStub creates a new Client service stub for the node.
func NewMotorStub(ctx context.Context, h *host.SNRHost, n api.NodeImpl, loc *common.Location, lst net.Listener) (*MotorStub, error) {
	// Create a new gRPC server
	var err error
	grpcServer := grpc.NewServer()
	stub := &MotorStub{
		ctx:            ctx,
		node:           n,
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
	stub.TransmitProtocol, err = transmit.New(ctx, h, n, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	stub.ExchangeProtocol, err = exchange.New(ctx, h, n, stub)
	if err != nil {
		logger.Errorf("%s - Failed to start ExchangeProtocol", err)
		return nil, err
	}

	// Start Routines
	RegisterMotorStubServer(grpcServer, stub)
	go stub.Serve(ctx, lst)
	return stub, nil
}

// HasProtocols returns true if the node has the protocols.
func (s *MotorStub) HasProtocols() bool {
	return s.TransmitProtocol != nil && s.DiscoverProtocol != nil
}

// Serve serves the RPC Service on the given port.
func (s *MotorStub) Serve(ctx context.Context, listener net.Listener) {
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

// Update method updates the node's properties in the Key/Value Store and Lobby
func (s *MotorStub) Update() error {
	// Call Internal Edit
	peer, err := s.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to get Peer Ref", err)
		return err
	}

	// Check for Valid Protocols
	if s.HasProtocols() {
		// Update LobbyProtocol
		err = s.DiscoverProtocol.Update()
		if err != nil {
			logger.Errorf("%s - Failed to Update Lobby", err)
		} else {
			logger.Debug("🌎 Succesfully updated Lobby.")
		}

		// Update ExchangeProtocol
		err := s.DiscoverProtocol.Put(peer)
		if err != nil {
			logger.Errorf("%s - Failed to Update Exchange", err)
		} else {
			logger.Debug("🌎 Succesfully updated Exchange.")
		}
		return err
	} else {
		return ErrProtocolsNotSet
	}
}

// Edit method edits the node's properties in the Key/Value Store
func (s *MotorStub) Edit(ctx context.Context, req *api.EditRequest) (*api.EditResponse, error) {
	// Call Internal Update
	if err := s.Update(); err != nil {
		return &api.EditResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &api.EditResponse{
		Success: true,
	}, nil
}

// Fetch method retreives Node properties from Key/Value Store
func (s *MotorStub) Fetch(ctx context.Context, req *api.FetchRequest) (*api.FetchResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// Call Internal Fetch4
	profile, err := s.node.Profile()
	if err != nil {
		return &api.FetchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &api.FetchResponse{
		Success: true,
		Profile: profile,
	}, nil
}

// Share method sends supplied files/urls with a peer
func (s *MotorStub) Share(ctx context.Context, req *api.ShareRequest) (*api.ShareResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// Request Peer to Transmit File
	if s.TransmitProtocol != nil {
		err := s.ExchangeProtocol.Request(req)
		if err != nil {
			return &api.ShareResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}
	} else {
		return &api.ShareResponse{
			Success: false,
			Error:   ErrProtocolsNotSet.Error(),
		}, nil
	}

	// Send Response
	return &api.ShareResponse{
		Success: true,
	}, nil
}

// Search Method to find a Peer by SName
func (s *MotorStub) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// Call Internal Ping
	if s.DiscoverProtocol != nil {
		// Call Internal Search
		entry, err := s.Get(strings.ToLower(req.GetSName()))
		if err != nil {
			return &api.SearchResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}

		// Send Response
		return &api.SearchResponse{
			Success: true,
			Peer:    entry,
		}, nil
	} else {
		return &api.SearchResponse{
			Success: false,
			Error:   ErrProtocolsNotSet.Error(),
		}, nil
	}

}

// Respond method responds to a received InviteRequest.
func (s *MotorStub) Respond(ctx context.Context, req *api.RespondRequest) (*api.RespondResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// buildRespFunc is a function that builds a response to a received InviteRequest.
	buildRespFunc := func(err error) *api.RespondResponse {
		if err == nil {
			return &api.RespondResponse{
				Success: true,
			}
		}
		return &api.RespondResponse{
			Success: false,
			Error:   err.Error(),
		}
	}

	// Get Request Parameters
	decs := req.GetDecision()
	peer := req.GetPeer()

	// Respond on ExchangeProtocol
	payload, err := s.ExchangeProtocol.Respond(decs, peer)
	if err != nil {
		return buildRespFunc(err), nil
	}

	// Check decision
	if decs {
		// Prepare on TransmitProtocol
		if err := s.TransmitProtocol.Incoming(payload, peer); err != nil {
			return buildRespFunc(err), nil
		}
	}

	// Send Response
	return buildRespFunc(nil), nil

}
