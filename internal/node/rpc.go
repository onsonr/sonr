package node

import (
	context "context"
	"fmt"
	"net"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/emitter"
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
	exchangeEvents chan *common.ExchangeEvent
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
		exchangeEvents: make(chan *common.ExchangeEvent),
		inviteEvents:   make(chan *common.InviteEvent),
		progressEvents: make(chan *common.ProgressEvent),
		completeEvents: make(chan *common.CompleteEvent),
	}

	// Register RPC Service
	RegisterNodeServiceServer(grpcServer, nrc)
	go func(nodeRpcService *NodeRPCService, grpcServer *grpc.Server) {
		if err := grpcServer.Serve(nodeRpcService.listener); err != nil {
			logger.Error("Failed to serve gRPC", zap.Error(err))
			return
		}
	}(nrc, grpcServer)

	// Register RPC Events
	nrc.Node.On(Event_STATUS, func(e *emitter.Event) {
		event := &common.StatusEvent{
			IsOk:    e.Args[0].(bool),
			Message: e.Args[1].(string),
		}
		nrc.statusEvents <- event
	})

	// Handle Node Events
	nrc.Node.On(transfer.Event_INVITED, func(e *emitter.Event) {
		inv := e.Args[0].(*transfer.InviteEvent)
		invEvent := &common.InviteEvent{
			InviteId: inv.GetInviteId(),
			From:     inv.GetFrom(),
			Transfer: inv.GetTransfer(),
		}
		nrc.inviteEvents <- invEvent
	})

	// Handle Node Events
	nrc.Node.On(exchange.Event_PEER_UPDATE, func(e *emitter.Event) {
		updEvent := e.Args[0].(*exchange.UpdateEvent)
		exchEvent := &common.ExchangeEvent{
			Peer: updEvent.GetPeer(),
			Type: common.ExchangeEvent_UPDATE,
		}
		nrc.exchangeEvents <- exchEvent
	})

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
		}, err
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
		}, err
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
		}, err
	}

	// Send Response
	return &ShareResponse{
		Success: true,
	}, nil
}

// Search Method to find a Peer by SName
func (n *NodeRPCService) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	// Call Internal Ping
	peer, err := n.Node.Ping(req.GetSName())
	if err != nil {
		return &PingResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Send Response
	return &PingResponse{
		Success: true,
		Peer:    peer,
	}, nil
}

// Respond method responds to a received InviteRequest.
func (n *NodeRPCService) Respond(ctx context.Context, req *RespondRequest) (*RespondResponse, error) {
	// Call Internal Respond
	err := n.Node.Respond(req.GetDecision())
	if err != nil {
		return &RespondResponse{
			Success: false,
			Error:   err.Error(),
		}, err
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
		SName: n.profile.SName,
		Peer:  n.Peer(),
		Device: &StatResponse_Device{
			Id:      device.Stat().Id,
			Name:    device.Stat().Name,
			Os:      device.Stat().Os,
			Arch:    device.Stat().Arch,
			Version: device.Stat().Version,
		},
		Network: &StatResponse_Network{
			PublicKey: n.SHost.Stat().PublicKey,
			PeerID:    n.SHost.Stat().PeerID,
			Multiaddr: n.SHost.Stat().MultAddr,
		},
	}, nil
}

// OnDecision method sends a decision event to the client.
func (n *NodeRPCService) OnNodeStatus(e *Empty, stream NodeService_OnNodeStatusServer) error {
	for {
		select {
		case m := <-n.statusEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnLocalJoin method sends a join event to the client.
func (n *NodeRPCService) OnLocalJoin(e *Empty, stream NodeService_OnLocalJoinServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				if m.GetType() == common.ExchangeEvent_JOIN {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnLocalJoin method sends a join event to the client.
func (n *NodeRPCService) OnLocalUpdate(e *Empty, stream NodeService_OnLocalUpdateServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				if m.GetType() == common.ExchangeEvent_UPDATE {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnLocalExit method sends a join event to the client.
func (n *NodeRPCService) OnLocalExit(e *Empty, stream NodeService_OnLocalExitServer) error {
	for {
		select {
		case m := <-n.exchangeEvents:
			if m != nil {
				if m.GetType() == common.ExchangeEvent_EXIT {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnDecision-Accepted method sends a decision event to the client.
func (n *NodeRPCService) OnTransferAccepted(e *Empty, stream NodeService_OnTransferAcceptedServer) error {
	for {
		select {
		case m := <-n.decisionEvents:
			if m != nil {
				if m.Decision {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnDecision-Declined method sends a decision event to the client.
func (n *NodeRPCService) OnTransferDeclined(e *Empty, stream NodeService_OnTransferDeclinedServer) error {
	for {
		select {
		case m := <-n.decisionEvents:
			if m != nil {
				if !m.Decision {
					stream.Send(m)
				}
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnTransferInvite method sends an invite event to the client.
func (n *NodeRPCService) OnTransferInvite(e *Empty, stream NodeService_OnTransferInviteServer) error {
	for {
		select {
		case m := <-n.inviteEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnTransferProgress method sends a progress event to the client.
func (n *NodeRPCService) OnTransferProgress(e *Empty, stream NodeService_OnTransferProgressServer) error {
	for {
		select {
		case m := <-n.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}

// OnTransferComplete method sends a complete event to the client.
func (n *NodeRPCService) OnTransferComplete(e *Empty, stream NodeService_OnTransferCompleteServer) error {
	for {
		select {
		case m := <-n.completeEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
		state.GetState().NeedsWait()
	}
}
