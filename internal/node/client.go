package node

import (
	context "context"
	"fmt"
	"net"
	"strings"
	"time"

	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"

	grpc "google.golang.org/grpc"
)

// RPC_SERVER_PORT is the port the RPC service listens on.
const RPC_SERVER_PORT = 52006

// ClientNodeStub is the RPC Service for the Node.
type ClientNodeStub struct {
	ClientServiceServer
	*Node

	// Properties
	ctx context.Context

	// grpcServer is the gRPC server.
	grpcServer *grpc.Server

	// TCPListener for RPC Service
	listener net.Listener

	// TransferProtocol - the transfer protocol
	*transfer.TransferProtocol

	// ExchangeProtocol - the exchange protocol
	*exchange.ExchangeProtocol

	// LobbyProtocol - The lobby protocol
	*lobby.LobbyProtocol

	// MailboxProtocol - Offline Mailbox Protocol
	// *mailbox.MailboxProtocol
}

// startClientService creates a new Client service stub for the node.
func (n *Node) startClientService(ctx context.Context, olc string) (*ClientNodeStub, error) {
	// Set Transfer Protocol
	transferProtocol, err := transfer.NewProtocol(ctx, n.host, n.Emitter)
	if err != nil {
		logger.Child("Client").Error("Failed to start TransferProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	exchProtocol, err := exchange.NewProtocol(ctx, n.host, n.Emitter)
	if err != nil {
		logger.Child("Client").Error("Failed to start ExchangeProtocol", err)
		return nil, err
	}

	// Set Local Lobby Protocol if Location is provided
	lobbyProtocol, err := lobby.NewProtocol(ctx, n.host, n.Emitter, olc)
	if err != nil {
		logger.Child("Client").Error("Failed to start LobbyProtocol", err)
		return nil, err
	}

	// Bind RPC Service
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		logger.Child("Client").Error("Failed to bind to port", err)
		return nil, err
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	nrc := &ClientNodeStub{
		grpcServer:       grpcServer,
		listener:         listener,
		ctx:              ctx,
		Node:             n,
		TransferProtocol: transferProtocol,
		ExchangeProtocol: exchProtocol,
		LobbyProtocol:    lobbyProtocol,
	}

	// Start Routines
	RegisterClientServiceServer(grpcServer, nrc)

	// Handle Node Events
	if err := nrc.grpcServer.Serve(nrc.listener); err != nil {
		logger.Child("Client").Error("Failed to serve gRPC", err)
		return nil, err
	}
	go nrc.pushAutomaticPings(ctx, time.NewTicker(5*time.Second))
	return nrc, nil
}

// Update method updates the node's properties in the Key/Value Store and Lobby
func (n *ClientNodeStub) Update() error {
	// Call Internal Edit
	peer, err := n.Peer()
	if err != nil {
		logger.Child("Client").Error("Failed to push Auto Ping", err)
		return err
	}

	// Push Update to Exchange
	if n.ExchangeProtocol != nil {
		if err := n.ExchangeProtocol.Update(peer); err != nil {
			logger.Child("Client").Error("Failed to Update Exchange", err)
			return err
		}
	}

	// Push Update to Lobby
	if n.LobbyProtocol != nil {
		if err := n.LobbyProtocol.Update(peer); err != nil {
			logger.Child("Client").Error("Failed to Update Lobby", err)
			return err
		}
	}
	return nil
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
		// Call Internal Respond
		toId, inv, err := n.Node.NewRequest(req.GetPeer())
		if err != nil {
			return &SupplyResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}

		// Request Peer to Transfer File
		if n.TransferProtocol != nil {
			err = n.TransferProtocol.Request(toId, inv)
			if err != nil {
				return &SupplyResponse{
					Success: false,
					Error:   err.Error(),
				}, nil
			}
		} else {
			return &SupplyResponse{
				Success: false,
				Error:   ErrTransferNotCreated.Error(),
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
	// Call Internal Update
	if err := n.Update(); err != nil {
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
	// Call Internal Fetch4
	profile, err := n.Node.Profile()
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
	// Call Internal Respond
	toId, inv, err := n.Node.NewRequest(req.GetPeer())
	if err != nil {
		return &ShareResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Request Peer to Transfer File
	if n.TransferProtocol != nil {
		err = n.TransferProtocol.Request(toId, inv)
		if err != nil {
			return &ShareResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}
	} else {
		return &ShareResponse{
			Success: false,
			Error:   ErrTransferNotCreated.Error(),
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
	if n.ExchangeProtocol != nil {
		// Call Internal Search
		entry, err := n.Query(strings.ToLower(req.GetSName()))
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
	} else {
		return &SearchResponse{
			Success: false,
			Error:   ErrExchangeNotCreated.Error(),
		}, nil
	}

}

// Respond method responds to a received InviteRequest.
func (n *ClientNodeStub) Respond(ctx context.Context, req *RespondRequest) (*RespondResponse, error) {
	// Call Internal Respond
	if n.TransferProtocol != nil {
		toId, resp, err := n.Node.NewResponse(req.GetDecision(), req.GetPeer())
		if err != nil {
			return &RespondResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}

		// Respond on TransferProtocol
		err = n.TransferProtocol.Respond(toId, resp)
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
	} else {
		return &RespondResponse{
			Success: false,
			Error:   ErrTransferNotCreated.Error(),
		}, nil
	}

}

// Stat method returns the node's stats
func (n *ClientNodeStub) Stat(ctx context.Context, req *StatRequest) (*StatResponse, error) {
	resp, _ := n.Node.Stat()
	return resp, nil
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

// OnMailboxMessage method sends an accepted event to the client.
func (n *ClientNodeStub) OnMailboxMessage(e *Empty, stream ClientService_OnMailboxMessageServer) error {
	for {
		select {
		case m := <-n.mailEvents:
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
				// Check Direction
				if m.Direction == common.CompleteEvent_INCOMING {
					// Add Sender to Recents
					err := n.AddRecent(m.GetFrom().GetProfile())
					if err != nil {
						logger.Child("Client").Error("Failed to add sender's profile to store.", err)
						return err
					}
				} else {
					// Add Receiver to Recents
					err := n.AddRecent(m.GetTo().GetProfile())
					if err != nil {
						logger.Child("Client").Error("Failed to add receiver's profile to store.", err)
						return err
					}
				}
				stream.Send(m)
			}
		case <-n.ctx.Done():
			return nil
		}
	}
}

// pushAutomaticPings sends automatic pings to the network of Profile
func (n *ClientNodeStub) pushAutomaticPings(ctx context.Context, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			// Call Internal Update
			if err := n.Update(); err != nil {
				logger.Child("Client").Error("Failed to push Auto Ping", err)
				ticker.Stop()
				return
			}
		case <-ctx.Done():
			ticker.Stop()
			n.grpcServer.Stop()
			return
		}
	}
}
