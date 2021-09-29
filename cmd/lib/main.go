package lib

import (
	"context"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

type Client struct {
	// Properties
	ctx     context.Context
	host    *host.SNRHost
	node    *node.Node
	service *node.NodeRPCService
}

var client *Client
var started bool

// Start starts the host, node, and rpc service.
func Start(reqBytes []byte) {
	// Check if already started
	if !started {
		// Unmarshal request
		req, fsOpts, err := parseInitializeRequest(reqBytes)
		if err != nil {
			logger.Fatal("Failed to Parse InitializeRequest", zap.Error(err))
		}

		// Initialize Device
		err = device.Init(req.GetEnvOptions().GetEnvironment().IsDev(), fsOpts...)
		if err != nil {
			logger.Panic("Failed to initialize Device", zap.Error(err))
		}

		// Initialize environment
		ctx := context.Background()
		err = device.InitEnv()
		if err != nil {
			logger.Error("Failed to initialize environment variables", zap.Error(err))
		}

		// Initialize Host
		host, err := host.NewHost(ctx, req.GetConnection())
		if err != nil {
			logger.Panic("Failed to create Host", zap.Error(err))
		}

		// Create Node
		n, err := node.NewNode(ctx, host, req.GetLocation())
		if err != nil {
			logger.Panic("Failed to update Profile for Node", zap.Error(err))
		}

		// Create RPC Service
		service, err := node.NewRPCService(ctx, n)
		if err != nil {
			logger.Panic("Failed to start RPC Service", zap.Error(err))
		}

		// Create Client
		client = &Client{
			ctx:     ctx,
			host:    host,
			node:    n,
			service: service,
		}

		// Set Started
		started = true

		// Supply Initial Profile
		editReq := &node.EditRequest{
			Profile: req.GetProfile(),
		}

		// Push EditRequest
		_, err = client.service.Edit(client.ctx, editReq)
		if err != nil {
			logger.Error("Failed to supply initial Profile", zap.Error(err))
		}
	}
	return
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	if started {
		// state.GetState().Pause()
	}
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	if started {
		// state.GetState().Resume()
	}
}

// Stop closes the host, node, and rpc service.
func Stop() {
	if started {
		client.host.Close()
		client.ctx.Done()
	}
}
