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

// ClientNodeStub is the RPC Service for the Node.
type ClientNodeStub struct {
	ClientServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	listener   net.Listener

	// Channels
	decisionEvents chan *common.DecisionEvent
	refreshEvents  chan *common.RefreshEvent
	inviteEvents   chan *common.InviteEvent
	progressEvents chan *common.ProgressEvent
	completeEvents chan *common.CompleteEvent
}

// startClientService creates a new Client service stub for the node.
func (nd *Node) startClientService(ctx context.Context, loc *common.Location, profile *common.Profile) (*ClientNodeStub, error) {
	go func(node *Node) {

		// Set Transfer Protocol
		node.TransferProtocol = transfer.NewProtocol(ctx, node.host, node.Emitter)

		// Set Exchange Protocol
		exch, err := exchange.NewProtocol(ctx, node.host, node.Emitter)
		if err != nil {
			logger.Error("Failed to start ExchangeProtocol", err)
		} else {
			node.ExchangeProtocol = exch
		}

		// Set Local Lobby Protocol if Location is provided
		if loc != nil {
			lobby, err := lobby.NewProtocol(node.host, loc, node.Emitter)
			if err != nil {
				logger.Error("Failed to start LobbyProtocol", err)
			} else {
				node.LobbyProtocol = lobby
			}
		}

		// // Initialize Store
		// store, err := store.NewStore(ctx, node.host, node.Emitter)
		// if err != nil {
		// 	logger.Error("Failed to initialize store", err)
		// } else {
		// 	node.store = store
		// }
	}(nd)

	// Bind RPC Service
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		return nil, logger.Error("Failed to bind to port", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	nrc := &ClientNodeStub{
		grpcServer:     grpcServer,
		listener:       listener,
		ctx:            ctx,
		Node:           nd,
		decisionEvents: make(chan *common.DecisionEvent),
		refreshEvents:  make(chan *common.RefreshEvent),
		inviteEvents:   make(chan *common.InviteEvent),
		progressEvents: make(chan *common.ProgressEvent),
		completeEvents: make(chan *common.CompleteEvent),
	}

	// Start Routines
	RegisterClientServiceServer(grpcServer, nrc)
	go nrc.serveRPC()
	go nrc.handleEmitter()
	go nrc.Node.pushAutomaticPings()

	// Return RPC Service
	return nrc, nil
}

// Supply supplies the node with the given amount of resources.
func (n *ClientNodeStub) Supply(ctx context.Context, req *SupplyRequest) (*SupplyResponse, error) {
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

// Edit method edits the node's properties in the Key/Value Store
func (n *ClientNodeStub) Edit(ctx context.Context, req *EditRequest) (*EditResponse, error) {
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

// Fetch method retreives Node properties from Key/Value Store
func (n *ClientNodeStub) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	// Call Internal Fetch
	profile, err := n.Node.store.GetProfile()
	if err != nil {
		return &FetchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &FetchResponse{
		Success: true,
		Profile: profile,
	}, nil
}

// Share method sends supplied files/urls with a peer
func (n *ClientNodeStub) Share(ctx context.Context, req *ShareRequest) (*ShareResponse, error) {
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
func (n *ClientNodeStub) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	// Call Internal Ping
	entry, err := n.Node.Query(exchange.NewQueryRequestFromSName(req.GetSName()))
	if err != nil {
		return &SearchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &SearchResponse{
		Success: true,
		Peer:    entry.Peer,
	}, nil
}

// Respond method responds to a received InviteRequest.
func (n *ClientNodeStub) Respond(ctx context.Context, req *RespondRequest) (*RespondResponse, error) {
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
func (n *ClientNodeStub) Stat(ctx context.Context, req *StatRequest) (*StatResponse, error) {
	resp, _ := n.Node.Stat()
	return resp, nil
}

// HandleEmitter handles the emitter events.
func (nrc *ClientNodeStub) handleEmitter() {
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
			// Check Direction
			if event.Direction == common.CompleteEvent_INCOMING {
				// Add Sender to Recents
				err := nrc.Node.store.AddRecent(event.GetFrom().GetProfile())
				if err != nil {
					logger.Error("Failed to add sender's profile to store.", err)
				}
			} else {
				// Add Receiver to Recents
				err := nrc.Node.store.AddRecent(event.GetTo().GetProfile())
				if err != nil {
					logger.Error("Failed to add receiver's profile to store.", err)
				}
			}
			nrc.completeEvents <- event
		})

		// Handle Lobby Join Events
		nrc.Node.On(lobby.Event_LIST_REFRESH, func(e *state.Event) {
			refreshEvent := e.Args[0].(*common.RefreshEvent)
			nrc.refreshEvents <- refreshEvent
		})
	}
}

// serveRPC Serves the RPC Service on the given port.
func (nrc *ClientNodeStub) serveRPC() {
	for {
		// Handle Node Events
		if err := nrc.grpcServer.Serve(nrc.listener); err != nil {
			logger.Error("Failed to serve gRPC", err)
			return
		}
	}
}

// OnLobbyRefresh method sends a lobby refresh event to the client.
func (n *ClientNodeStub) OnLobbyRefresh(e *Empty, stream ClientService_OnLobbyRefreshServer) error {
	for {
		select {
		case m := <-n.refreshEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferAccepted method sends an accepted event to the client.
func (n *ClientNodeStub) OnTransferAccepted(e *Empty, stream ClientService_OnTransferAcceptedServer) error {
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

	}
}

// OnTransferDeclinedmethod sends a decline event to the client.
func (n *ClientNodeStub) OnTransferDeclined(e *Empty, stream ClientService_OnTransferDeclinedServer) error {
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

	}
}

// OnTransferInvite method sends an invite event to the client.
func (n *ClientNodeStub) OnTransferInvite(e *Empty, stream ClientService_OnTransferInviteServer) error {
	for {
		select {
		case m := <-n.inviteEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}

// OnTransferProgress method sends a progress event to the client.
func (n *ClientNodeStub) OnTransferProgress(e *Empty, stream ClientService_OnTransferProgressServer) error {
	for {
		select {
		case m := <-n.progressEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
	}
}

// OnTransferComplete method sends a complete event to the client.
func (n *ClientNodeStub) OnTransferComplete(e *Empty, stream ClientService_OnTransferCompleteServer) error {
	for {
		select {
		case m := <-n.completeEvents:
			if m != nil {
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}

	}
}
