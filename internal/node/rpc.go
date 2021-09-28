package node

import (
	context "context"
	"fmt"
	"net"
	"time"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

const RPC_SERVER_PORT = 52006

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
		logger.Error("Failed to bind to port", zap.Error(err))
		return nil, err
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
	// Call Internal Stat
	return &StatResponse{
		SName:   n.profile.SName,
		Profile: n.profile,
		Device: &StatResponse_Device{
			Id:        device.Stat().Id,
			Name:      device.Stat().Name,
			Os:        device.Stat().Os,
			Arch:      device.Stat().Arch,
			IsDesktop: device.Stat().IsDesktop,
			IsMobile:  device.Stat().IsMobile,
		},
		Network: &StatResponse_Network{
			PublicKey: n.host.Stat().PublicKey,
			PeerID:    n.host.Stat().PeerID,
			Multiaddr: n.host.Stat().MultAddr,
		},
	}, nil
}

// HandleEmitter handles the emitter events.
func (nrc *NodeRPCService) handleEmitter() {
	for {
		// Handle Transfer Invite
		nrc.Node.On(transfer.Event_INVITED, func(e *state.Event) {
			inv := e.Args[0].(*transfer.InviteRequest)
			invEvent := &common.InviteEvent{
				Received: int64(time.Now().Unix()),
				From:     inv.GetFrom(),
				Payload:  inv.GetPayload(),
			}
			nrc.inviteEvents <- invEvent
		})

		// Handle Transfer Decision
		nrc.Node.On(transfer.Event_RESPONDED, func(e *state.Event) {
			inv := e.Args[0].(*transfer.InviteResponse)
			decsEvent := &common.DecisionEvent{
				From:     inv.GetFrom(),
				Received: int64(time.Now().Unix()),
				Decision: inv.GetDecision(),
			}
			nrc.decisionEvents <- decsEvent
		})

		// Handle Transfer Progress
		nrc.Node.On(transfer.Event_PROGRESS, func(e *state.Event) {
			writtenBytes := e.Args[0].(int64)
			progEvent := &common.ProgressEvent{
				Received: int64(time.Now().Unix()),
				Current:  int32(writtenBytes),
				Total:    0,
				Progress: 0,
			}
			nrc.progressEvents <- progEvent
		})

		// Handle Transfer Completed
		nrc.Node.On(transfer.Event_COMPLETED, func(e *state.Event) {
			result := e.Args[0].(*common.Payload)
			compEvent := &common.CompleteEvent{
				Received: int64(time.Now().Unix()),
				Payload:  result,
			}
			nrc.completeEvents <- compEvent
		})

		// Handle Lobby Join Events
		nrc.Node.On(lobby.Event_LIST_REFRESH, func(e *state.Event) {
			refreshEvent := e.Args[0].(*common.RefreshEvent)
			refreshEvent.Received = int64(time.Now().Unix())
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
			logger.Error("Failed to serve gRPC", zap.Error(err))
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
