package node

import (
	context "context"
	"fmt"
	"net"
	"time"

	api "github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/mailbox"
	"github.com/sonr-io/core/pkg/transfer"
	grpc "google.golang.org/grpc"
)

var DefaultAutoPingTicker = 5 * time.Second

// ClientNodeStub is the RPC Service for the Node.
type ClientNodeStub struct {
	// Interfaces
	api.NodeStubImpl
	ClientServiceServer

	// Properties
	ctx        context.Context
	isTerminal bool
	listener   net.Listener
	grpcServer *grpc.Server
	node       *Node

	// Protocols
	*transfer.TransferProtocol
	*exchange.ExchangeProtocol
	*lobby.LobbyProtocol
	*mailbox.MailboxProtocol
}

// startClientService creates a new Client service stub for the node.
func (n *Node) startClientService(ctx context.Context, opts *nodeOptions) (api.NodeStubImpl, error) {
	// Set Transfer Protocol
	transferProtocol, err := transfer.NewProtocol(ctx, n.host, n)
	if err != nil {
		logger.Error("Failed to start TransferProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	exchProtocol, err := exchange.NewProtocol(ctx, n.host, n)
	if err != nil {
		logger.Error("Failed to start ExchangeProtocol", err)
		return nil, err
	}

	// Set Local Lobby Protocol if Location is provided
	lobbyProtocol, err := lobby.NewProtocol(ctx, n.host, n, lobby.WithLocation(opts.location))
	if err != nil {
		logger.Error("Failed to start LobbyProtocol", err)
		return nil, err
	}

	// // Set Mailbox Protocol
	// mailboxProtocol, err := mailbox.NewProtocol(ctx, n.host, n.Emitter)
	// if err != nil {
	// 	logger.Error("Failed to start MailboxProtocol", err)
	// 	return nil, err
	// }

	// Open Listener on Port
	listener, err := net.Listen(opts.network, opts.address)
	if err != nil {
		logger.Fatal("Failed to bind listener to port ", err)
		return nil, err
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	stub := &ClientNodeStub{
		ctx:              ctx,
		isTerminal:       opts.isTerminal,
		TransferProtocol: transferProtocol,
		ExchangeProtocol: exchProtocol,
		LobbyProtocol:    lobbyProtocol,
		//MailboxProtocol:  mailboxProtocol,
		grpcServer: grpcServer,
		node:       n,
		listener:   listener,
	}

	// Start Routines
	RegisterClientServiceServer(grpcServer, stub)
	go stub.Serve(ctx, listener, DefaultAutoPingTicker)
	return stub, nil
}

// HasProtocols returns true if the node has the protocols.
func (s *ClientNodeStub) HasProtocols() bool {
	return s.TransferProtocol != nil && s.ExchangeProtocol != nil && s.LobbyProtocol != nil
}

// Close closes the RPC Service.
func (s *ClientNodeStub) Close() error {
	s.listener.Close()
	s.grpcServer.Stop()
	s.LobbyProtocol.Close()
	return nil
}

// Serve serves the RPC Service on the given port.
func (s *ClientNodeStub) Serve(ctx context.Context, listener net.Listener, ticker time.Duration) {
	// Handle Node Events
	if err := s.grpcServer.Serve(listener); err != nil {
		logger.Error("Failed to serve gRPC", err)
	}
	logger.Info("🍦  Serving Client Stub: " + listener.Addr().String())
	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			return
		}
	}
}

// Update method updates the node's properties in the Key/Value Store and Lobby
func (s *ClientNodeStub) Update() error {
	// Call Internal Edit
	peer, err := s.node.Peer()
	if err != nil {
		logger.Error("Failed to get Peer Ref", err)
		return err
	}

	if s.HasProtocols() {
		// Update LobbyProtocol
		err = s.LobbyProtocol.Update(peer)
		if err != nil {
			logger.Error("Failed to Update Lobby", err)
		} else {
			logger.Info("🌎 Succesfully updated Lobby.")
		}

		// Update ExchangeProtocol
		err := s.ExchangeProtocol.Put(peer)
		if err != nil {
			logger.Error("Failed to Update Exchange", err)
		} else {
			logger.Info("🌎 Succesfully updated Exchange.")
		}
		return err
	} else {
		return ErrProtocolsNotSet
	}
}

// HighwayNodeStub is the RPC Service for the Full Node.
type HighwayNodeStub struct {
	api.NodeStubImpl
	HighwayServiceServer
	ClientServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	listener   net.Listener
}

// startHighwayService creates a new Highway service stub for the node.
func (n *Node) startHighwayService(ctx context.Context, opts *nodeOptions) (api.NodeStubImpl, error) {

	// Create the listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
	if err != nil {
		return nil, err
	}

	// Create the RPC Service
	grpcServer := grpc.NewServer()
	stub := &HighwayNodeStub{
		Node:       n,
		ctx:        ctx,
		grpcServer: grpcServer,
		listener:   listener,
	}
	// Register the RPC Service
	RegisterHighwayServiceServer(stub.grpcServer, stub)
	go stub.Serve(ctx, listener, DefaultAutoPingTicker)
	return stub, nil
}

func (s *HighwayNodeStub) Serve(ctx context.Context, listener net.Listener, ticker time.Duration) {
	// Handle Node Events
	if err := s.grpcServer.Serve(s.listener); err != nil {
		logger.Error("Failed to serve gRPC", err)

	}
	logger.Info("🍦  Serving Highway Stub...")
	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (s *HighwayNodeStub) Close() error {
	s.listener.Close()
	s.grpcServer.Stop()
	return nil
}
