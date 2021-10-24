package node

import (
	context "context"
	"fmt"
	"net"

	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/mailbox"
	"github.com/sonr-io/core/pkg/transmit"

	"google.golang.org/grpc"
)

// StubMode is the type of the node (Client, Highway)
type StubMode int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	StubMode_LIB StubMode = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	StubMode_CLI

	// StubMode_BIN is the Node utilized for Desktop background process
	StubMode_BIN

	// StubMode_HIGHWAY is the Custodian Node that manages Network
	StubMode_HIGHWAY
)

// IsLib returns true if the node is a client node.
func (m StubMode) IsLib() bool {
	return m == StubMode_LIB
}

// IsBin returns true if the node is a bin node.
func (m StubMode) IsBin() bool {
	return m == StubMode_BIN
}

// IsCLI returns true if the node is a CLI node.
func (m StubMode) IsCLI() bool {
	return m == StubMode_CLI
}

// IsHighway returns true if the node is a highway node.
func (m StubMode) IsHighway() bool {
	return m == StubMode_HIGHWAY
}

// HasClient returns true if the node has a client.
func (m StubMode) HasClient() bool {
	return m.IsLib() || m.IsBin() || m.IsCLI()
}

// HasHighway returns true if the node has a highway stub.
func (m StubMode) HasHighway() bool {
	return m.IsHighway()
}

// Prefix returns golog prefix for the node.
func (m StubMode) Prefix() string {
	var name string
	switch m {
	case StubMode_LIB:
		name = "lib"
	case StubMode_CLI:
		name = "cli"
	case StubMode_BIN:
		name = "bin"
	case StubMode_HIGHWAY:
		name = "highway"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

// ClientNodeStub is the RPC Service for the Default Node.
type ClientNodeStub struct {
	// Interfaces
	ClientServiceServer

	// Properties
	ctx        context.Context
	node       *Node
	grpcServer *grpc.Server

	// Protocols
	*transmit.TransmitProtocol
	*exchange.ExchangeProtocol
	*mailbox.MailboxProtocol
}

// startClientService creates a new Client service stub for the node.
func (n *Node) startClientService(ctx context.Context, opts *options) (*ClientNodeStub, error) {

	// Set Exchange Protocol
	exchProtocol, err := exchange.NewProtocol(ctx, n.host, n, exchange.WithLocation(opts.location))
	if err != nil {
		logger.Errorf("%s - Failed to start ExchangeProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	transmitProtocol, err := transmit.NewProtocol(ctx, n.host, n)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	stub := &ClientNodeStub{
		ctx:              ctx,
		TransmitProtocol: transmitProtocol,
		ExchangeProtocol: exchProtocol,
		node:             n,
		grpcServer:       grpcServer,
	}

	// Start Routines
	RegisterClientServiceServer(grpcServer, stub)
	go stub.Serve(ctx, n.listener)
	return stub, nil
}

// HasProtocols returns true if the node has the protocols.
func (s *ClientNodeStub) HasProtocols() bool {
	return s.TransmitProtocol != nil && s.ExchangeProtocol != nil
}

// Serve serves the RPC Service on the given port.
func (s *ClientNodeStub) Serve(ctx context.Context, listener net.Listener) {
	// Handle Node Events
	if err := s.grpcServer.Serve(listener); err != nil {
		logger.Error("Failed to serve gRPC", err)
	}
	for {
		// Stop Serving if context is done
		select {
		case <-ctx.Done():
			s.grpcServer.Stop()
			s.ExchangeProtocol.Close()
			return
		}
	}
}

// Update method updates the node's properties in the Key/Value Store and Lobby
func (s *ClientNodeStub) Update() error {
	// Call Internal Edit
	peer, err := s.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to get Peer Ref", err)
		return err
	}

	// Check for Valid Protocols
	if s.HasProtocols() {
		// Update LobbyProtocol
		err = s.ExchangeProtocol.Update()
		if err != nil {
			logger.Errorf("%s - Failed to Update Lobby", err)
		} else {
			logger.Debug("🌎 Succesfully updated Lobby.")
		}

		// Update ExchangeProtocol
		err := s.ExchangeProtocol.Put(peer)
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

// HighwayNodeStub is the RPC Service for the Custodian Node.
type HighwayNodeStub struct {
	HighwayServiceServer
	ClientServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
}

// startHighwayService creates a new Highway service stub for the node.
func (n *Node) startHighwayService(ctx context.Context, opts *options) (*HighwayNodeStub, error) {
	// Create the RPC Service
	grpcServer := grpc.NewServer()
	stub := &HighwayNodeStub{
		Node:       n,
		ctx:        ctx,
		grpcServer: grpcServer,
	}
	// Register the RPC Service
	RegisterHighwayServiceServer(grpcServer, stub)
	go stub.Serve(ctx, n.listener)
	return stub, nil
}

// Serve serves the RPC Service on the given port.
func (s *HighwayNodeStub) Serve(ctx context.Context, listener net.Listener) {
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
