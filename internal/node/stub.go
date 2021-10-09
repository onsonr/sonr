package node

import (
	context "context"
	"fmt"
	"net"
	"time"

	"github.com/sonr-io/core/pkg/domain"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	grpc "google.golang.org/grpc"
)

// ClientNodeStub is the RPC Service for the Node.
type ClientNodeStub struct {
	NodeStub
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
func (n *Node) startClientService(ctx context.Context, listener net.Listener, olc string) (NodeStub, error) {
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

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	cns := &ClientNodeStub{
		grpcServer:       grpcServer,
		ctx:              ctx,
		Node:             n,
		TransferProtocol: transferProtocol,
		ExchangeProtocol: exchProtocol,
		LobbyProtocol:    lobbyProtocol,
		listener:         listener,
	}

	// Start Routines
	RegisterClientServiceServer(grpcServer, cns)
	return cns, nil
}

func (s *ClientNodeStub) Serve(ctx context.Context) error {
	// Handle Node Events
	if err := s.grpcServer.Serve(s.listener); err != nil {
		logger.Child("Client").Error("Failed to serve gRPC", err)
		return err
	}

	go s.pushAutomaticPings(ctx, time.NewTicker(5*time.Second))
	return nil
}

func (s *ClientNodeStub) Close() error {
	s.grpcServer.Stop()
	s.listener.Close()
	return nil
}

// Update method updates the node's properties in the Key/Value Store and Lobby
func (s *ClientNodeStub) Update() error {
	// Call Internal Edit
	peer, err := s.Peer()
	if err != nil {
		logger.Child("Client").Error("Failed to push Auto Ping", err)
		return err
	}

	// Push Update to Exchange
	if s.ExchangeProtocol != nil {
		if err := s.ExchangeProtocol.Update(peer); err != nil {
			logger.Child("Client").Error("Failed to Update Exchange", err)
			return err
		}
	}

	// Push Update to Lobby
	if s.LobbyProtocol != nil {
		if err := s.LobbyProtocol.Update(peer); err != nil {
			logger.Child("Client").Error("Failed to Update Lobby", err)
			return err
		}
	}
	return nil
}

// pushAutomaticPings sends automatic pings to the network of Profile
func (s *ClientNodeStub) pushAutomaticPings(ctx context.Context, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			// Call Internal Update
			if err := s.Update(); err != nil {
				logger.Child("Client").Error("Failed to push Auto Ping", err)
				ticker.Stop()
				return
			}
		case <-ctx.Done():
			ticker.Stop()
			s.grpcServer.Stop()
			return
		}
	}
}

// HighwayNodeStub is the RPC Service for the Full Node.
type HighwayNodeStub struct {
	NodeStub
	HighwayServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	listener   net.Listener
	*domain.DomainProtocol
}

// startHighwayService creates a new Highway service stub for the node.
func (n *Node) startHighwayService(ctx context.Context) (NodeStub, error) {
	// Create the listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		return nil, err
	}

	// Initialize Domain Protocol
	domainProtocol, err := domain.NewProtocol(ctx, n.host)
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	grpcServer := grpc.NewServer()
	stub := &HighwayNodeStub{
		Node:           n,
		ctx:            ctx,
		grpcServer:     grpcServer,
		listener:       listener,
		DomainProtocol: domainProtocol,
	}
	// Register the RPC Service
	RegisterHighwayServiceServer(stub.grpcServer, stub)
	return stub, nil
}

func (s *HighwayNodeStub) Serve(ctx context.Context) error {
	// Handle Node Events
	if err := s.grpcServer.Serve(s.listener); err != nil {
		logger.Child("Highway").Error("Failed to serve gRPC", err)
		return err
	}

	// Start the server
	go s.serveRPC(ctx)
	return nil
}

func (s *HighwayNodeStub) Close() error {
	s.grpcServer.Stop()
	return nil
}

// serveRPC Serves the RPC Service on the given port.
func (hrc *HighwayNodeStub) serveRPC(ctx context.Context) {
	for {

		// Stop Serving if context is done
		select {
		case <-hrc.ctx.Done():
			hrc.host.Close()
			return
		}
	}
}
