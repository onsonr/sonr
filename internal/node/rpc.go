package node

import (
	context "context"
	"fmt"
	"net"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/emitter"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

const RPC_SERVER_PORT = 60214

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

	nrc.Node.On(Event_STATUS, func(e *emitter.Event) {
		isOk := e.Args[0].(bool)
		message := e.Args[1].(string)

		event := &common.StatusEvent{
			IsOk:    isOk,
			Message: message,
		}
		nrc.statusEvents <- event
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
	return nil, nil
}

// Search Method to find a Peer by SName
func (n *NodeRPCService) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	// Call Internal Search
	peer, err := n.Node.Find(req.GetSName())
	if err != nil {
		return &SearchResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &SearchResponse{
		Success: true,
		Peer:    peer,
	}, nil
}

// Respond method responds to a received InviteRequest.
func (n *NodeRPCService) Respond(ctx context.Context, req *RespondRequest) (*RespondResponse, error) {
	// // Unmarshal Data to Request
	// resp := &data.DecisionRequest{}
	// if err := proto.Unmarshal(buf, resp); err != nil {
	// 	n.handleError(data.NewError(err, data.ErrorEvent_UNMARSHAL))
	// 	return
	// }

	// // Send Response
	// n.client.Respond(resp.ToResponse())
	return nil, nil
}

// OnDecision method sends a decision event to the client.
func (n *NodeRPCService) OnStatus(e *Empty, stream NodeService_OnStatusServer) error {
	for {
		select {
		case m := <-n.statusEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
	}
}

// OnDecision method sends a decision event to the client.
func (n *NodeRPCService) OnDecision(e *Empty, stream NodeService_OnDecisionServer) error {
	for {
		select {
		case m := <-n.decisionEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
	}
}

// OnInvite method sends an invite event to the client.
func (n *NodeRPCService) OnInvite(e *Empty, stream NodeService_OnInviteServer) error {
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

// OnProgress method sends a progress event to the client.
func (n *NodeRPCService) OnProgress(e *Empty, stream NodeService_OnProgressServer) error {
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

// OnComplete method sends a complete event to the client.
func (n *NodeRPCService) OnComplete(e *Empty, stream NodeService_OnCompleteServer) error {
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
