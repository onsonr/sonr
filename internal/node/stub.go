package node

import (
	context "context"
	"net"

	"github.com/sonr-io/core/pkg/discover"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/registry"
	"github.com/sonr-io/core/pkg/transmit"

	"google.golang.org/grpc"
)

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
	*discover.DiscoverProtocol
	*registry.RegistryProtocol
	*exchange.ExchangeProtocol
}

// startHighwayStub creates a new Highway service stub for the node.
func (n *Node) startHighwayStub(ctx context.Context, opts *options) (*NodeHighwayStub, error) {
	// Set Discovery Protocol
	discProtocol, err := discover.New(ctx, n.host, n, discover.WithLocation(opts.location))
	if err != nil {
		logger.Errorf("%s - Failed to start DiscoveryProtocol", err)
		return nil, err
	}

	// Set Transmit Protocol
	exchangeProtocol, err := exchange.New(ctx, n.host, n)
	if err != nil {
		logger.Errorf("%s - Failed to start TransmitProtocol", err)
		return nil, err
	}

	// Set Exchange Protocol
	registeryProtocol, err := registry.New(ctx, n.host, n)
	if err != nil {
		logger.Errorf("%s - Failed to start ExchangeProtocol", err)
		return nil, err
	}

	// Create the RPC Service
	grpcServer := grpc.NewServer()
	stub := &NodeHighwayStub{
		Node:             n,
		ctx:              ctx,
		grpcServer:       grpcServer,
		DiscoverProtocol: discProtocol,
		ExchangeProtocol: exchangeProtocol,
		RegistryProtocol: registeryProtocol,
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
