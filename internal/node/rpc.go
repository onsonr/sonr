package node

import (
	context "context"
	"fmt"
	"net"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	grpc "google.golang.org/grpc"
)

// RPC_SERVER_PORT is the port the RPC service listens on.
const RPC_SERVER_PORT = 52006

// NodeRPCService is the RPC Service for the Node.
type NodeRPCService struct {
	NodeServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	listener   net.Listener

	// Channels
	statusEvents   chan *common.StatusEvent
	decisionEvents chan *common.DecisionEvent
	exchangeEvents chan *common.RefreshEvent
	inviteEvents   chan *common.InviteEvent
	progressEvents chan *common.ProgressEvent
	completeEvents chan *common.CompleteEvent
}

// NewRPCService creates a new RPC service for the node.
func NewRPCService(ctx context.Context, n *Node) (*NodeRPCService, error) {
	// Bind RPC Service
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		return nil, logger.Error("Failed to bind to port", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	nrc := &NodeRPCService{
		grpcServer:     grpcServer,
		listener:       listener,
		ctx:            ctx,
		Node:           n,
		statusEvents:   make(chan *common.StatusEvent),
		decisionEvents: make(chan *common.DecisionEvent),
		exchangeEvents: make(chan *common.RefreshEvent),
		inviteEvents:   make(chan *common.InviteEvent),
		progressEvents: make(chan *common.ProgressEvent),
		completeEvents: make(chan *common.CompleteEvent),
	}

	// Start Routines
	RegisterNodeServiceServer(grpcServer, nrc)
	go nrc.serveRPC()
	go nrc.handleEmitter()

	// Return RPC Service
	return nrc, nil
}

// Supply supplies the node with the given amount of resources.
func (n *NodeRPCService) Supply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
	// Call Internal Supply
	err := n.Node.Supply(req.GetPaths())
	if err != nil {
		return &SupplyResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Check if Peer is provided
	if req.GetPeer() != nil {
		// Call Internal Share
		err = n.Node.Share(req.GetPeer())
		if err != nil {
			return &SupplyResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}
	}

	// Send Response
	return &SupplyResponse{
		Success: true,
	}, nil
}

// Edit method edits the node's user profile.
func (n *NodeRPCService) Edit(ctx context.Context, req *EditRequest) (*EditResponse, error) {
	// Call Internal Edit
	err := n.Node.Edit(req.GetProfile())
	if err != nil {
		return &EditResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &EditResponse{
		Success: true,
	}, nil
}

// Share method sends supplied files/urls with a peer
func (n *NodeRPCService) Share(ctx context.Context, req *ShareRequest) (*ShareResponse, error) {
	// Call Internal Share
	err := n.Node.Share(req.GetPeer())
	if err != nil {
		return &ShareResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &ShareResponse{
		Success: true,
	}, nil
}

// Search Method to find a Peer by SName
func (n *NodeRPCService) Find(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	// Call Internal Ping
	entry, err := n.Node.Query(exchange.NewQueryRequestFromSName(req.GetSName()))
	if err != nil {
		return &FindResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &FindResponse{
		Success: true,
		Peer:    entry.Peer,
	}, nil
}

// Respond method responds to a received InviteRequest.
func (n *NodeRPCService) Respond(ctx context.Context, req *RespondRequest) (*RespondResponse, error) {
	// Call Internal Respond
	err := n.Node.Respond(req.GetDecision(), req.GetPeer())
	if err != nil {
		return &RespondResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &RespondResponse{
		Success: true,
	}, nil
}

// Stat method returns the node's stats
func (n *NodeRPCService) Stat(ctx context.Context, req *StatRequest) (*StatResponse, error) {
	resp, _ := n.Node.Stat()
	return resp, nil
}

// HandleEmitter handles the emitter events.
func (nrc *NodeRPCService) handleEmitter() {
	for {
		// Handle Transfer Invite
		nrc.Node.On(transfer.Event_INVITED, func(e *state.Event) {
			event := e.Args[0].(*common.InviteEvent)
			nrc.inviteEvents <- event
		})

		// Handle Transfer Decision
		nrc.Node.On(transfer.Event_RESPONDED, func(e *state.Event) {
			event := e.Args[0].(*common.DecisionEvent)
			nrc.decisionEvents <- event
		})

		// Handle Transfer Progress
		nrc.Node.On(transfer.Event_PROGRESS, func(e *state.Event) {
			event := e.Args[0].(*common.ProgressEvent)
			nrc.progressEvents <- event
		})

		// Handle Transfer Completed
		nrc.Node.On(transfer.Event_COMPLETED, func(e *state.Event) {
			event := e.Args[0].(*common.CompleteEvent)
			nrc.completeEvents <- event
		})

		// Handle Lobby Join Events
		nrc.Node.On(lobby.Event_LIST_REFRESH, func(e *state.Event) {
			refreshEvent := e.Args[0].(*common.RefreshEvent)
			nrc.exchangeEvents <- refreshEvent
		})

		// Stop Emitter if context is done
		select {
		case <-nrc.ctx.Done():
			nrc.host.Close()
			return
		}
	}
}

// serveRPC Serves the RPC Service on the given port.
func (nrc *NodeRPCService) serveRPC() {
	for {
		// Handle Node Events
		if err := nrc.grpcServer.Serve(nrc.listener); err != nil {
			logger.Error("Failed to serve gRPC", err)
			return
		}

		// Stop Serving if context is done
		select {
		case <-nrc.ctx.Done():
			nrc.host.Close()
			return
		}
	}
}
