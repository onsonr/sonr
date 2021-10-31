package node

import (
	context "context"
	"fmt"
	"net"

	"github.com/sonr-io/core/pkg/discover"
	"github.com/sonr-io/core/pkg/exchange"
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

// HasMotor returns true if the node has a client.
func (m StubMode) HasMotor() bool {
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

// NodeMotorStub is the RPC Service for the Default Node.
type NodeMotorStub struct {
	// Interfaces
	MotorStubServer

	// Properties
	ctx        context.Context
	node       *Node
	grpcServer *grpc.Server

	// Protocols
	*transmit.TransmitProtocol
	*discover.DiscoverProtocol
	*exchange.ExchangeProtocol
}

// startMotorStub creates a new Client service stub for the node.
func (n *Node) startMotorStub(ctx context.Context, opts *options) (*NodeMotorStub, error) {

	// Set Discovery Protocol
	discProtocol, err := discover.New(ctx, n.host, n, discover.WithLocation(opts.location))
	if err != nil {
		logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	transmitProtocol, err := transmit.New(ctx, n.host, n)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	exchangeProtocol, err := exchange.New(ctx, n.host, n)
	if err != nil {
		logger.Errorf("%s - Failed to start ExchangeProtocol", err)
		return nil, err
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	stub := &NodeMotorStub{
		ctx:              ctx,
		TransmitProtocol: transmitProtocol,
		DiscoverProtocol: discProtocol,
		ExchangeProtocol: exchangeProtocol,
		node:             n,
		grpcServer:       grpcServer,
	}

	// Start Routines
	RegisterMotorStubServer(grpcServer, stub)
	go stub.Serve(ctx, n.listener)
	return stub, nil
}

// HasProtocols returns true if the node has the protocols.
func (s *NodeMotorStub) HasProtocols() bool {
	return s.TransmitProtocol != nil && s.DiscoverProtocol != nil
}

// Serve serves the RPC Service on the given port.
func (s *NodeMotorStub) Serve(ctx context.Context, listener net.Listener) {
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
func (s *NodeMotorStub) Update() error {
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

// NodeHighwayStub is the RPC Service for the Custodian Node.
type NodeHighwayStub struct {
	HighwayStubServer
	MotorStubServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
}

// startHighwayStub creates a new Highway service stub for the node.
func (n *Node) startHighwayStub(ctx context.Context, opts *options) (*NodeHighwayStub, error) {
	// Create the RPC Service
	grpcServer := grpc.NewServer()
	stub := &NodeHighwayStub{
		Node:       n,
		ctx:        ctx,
		grpcServer: grpcServer,
	}
	// Register the RPC Service
	RegisterHighwayStubServer(grpcServer, stub)
	go stub.Serve(ctx, n.listener)
	return stub, nil
}

// Serve serves the RPC Service on the given port.
func (s *NodeHighwayStub) Serve(ctx context.Context, listener net.Listener) {
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
