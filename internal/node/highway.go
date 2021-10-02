package node

import (
	context "context"
	"fmt"
	"net"

	"github.com/sonr-io/core/pkg/domain"
	"github.com/sonr-io/core/tools/logger"
	dnet "github.com/sonr-io/core/tools/net"
	grpc "google.golang.org/grpc"
)

// ClientNodeStub is the RPC Service for the Node.
type HighwayNodeStub struct {
	HighwayServiceServer
	*Node

	// Properties
	ctx        context.Context
	grpcServer *grpc.Server
	listener   net.Listener
	*domain.DomainProtocol
}

// startClientService creates a new Client service stub for the node.
func (n *Node) startHighwayService(ctx context.Context, clientKey string, clientSecret string) (*HighwayNodeStub, error) {
	// Initialize Domain Protocol
	domainProtocol, err := domain.NewProtocol(ctx, n.host, clientKey, clientSecret)
	if err != nil {
		return nil, err
	}

	// Create the listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", RPC_SERVER_PORT))
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

	// Start the server
	go stub.serveRPC()
	return stub, nil
}

// Authorize Signing Method Request for Data
func (hrc *HighwayNodeStub) Authorize(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error) {
	logger.Info("HighwayService.Authorize() is Unimplemented")
	return nil, nil
}

// Link a new Device to the Node
func (hrc *HighwayNodeStub) Link(ctx context.Context, req *LinkRequest) (*LinkResponse, error) {
	logger.Info("HighwayService.Link() is Unimplemented")
	return nil, nil
}

// Register a new domain with the Node on the highway
func (hrc *HighwayNodeStub) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	// Get Values
	pfix := req.GetPrefix()
	name := req.GetSName()
	fprint := req.GetFingerprint()

	// Check Values
	if pfix == "" || name == "" || fprint == "" {
		return &RegisterResponse{
			Success: false,
			Error:   "Invalid request. One or more of the required fields are empty.",
		}, nil
	}
	// Create Record
	resp, err := hrc.DomainProtocol.Register(name, dnet.NewNBAuthRecord(pfix, name, fprint))
	if err != nil {
		return &RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Return Response
	return &RegisterResponse{
		Success: true,
		Records: resp,
	}, nil
}

// serveRPC Serves the RPC Service on the given port.
func (hrc *HighwayNodeStub) serveRPC() {
	for {
		// Handle Node Events
		if err := hrc.grpcServer.Serve(hrc.listener); err != nil {
			logger.Error("Failed to serve gRPC", err)
			return
		}

		// Stop Serving if context is done
		select {
		case <-hrc.ctx.Done():
			hrc.host.Close()
			return
		}
	}
}
