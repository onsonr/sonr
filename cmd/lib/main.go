package lib

import (
	"context"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
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
		req, fsOpts, envVars, err := parseInitializeRequest(reqBytes)
		if err != nil {
			panic(logger.Error("Failed to Parse Initialize Request", err))
		}

		// Initialize Device
		err = device.Init(req.GetEnvOptions().GetEnvironment().IsDev(), fsOpts...)
		if err != nil {
			panic(logger.Error("Failed to initialize Device", err))
		}

		// Initialize environment
		ctx := context.Background()
		err = device.InitEnv(envVars)
		if err != nil {
			logger.Error("Failed to initialize environment variables", err)
		}

		// Initialize Host
		host, err := host.NewHost(ctx, req.GetConnection())
		if err != nil {
			panic(logger.Error("Failed to create Host", err))
		}

		// Create Node
		n, err := node.NewNode(ctx, host, req.GetLocation())
		if err != nil {
			panic(logger.Error("Failed to update Profile for Node", err))
		}

		// Create RPC Service
		service, err := node.NewRPCService(ctx, n)
		if err != nil {
			panic(logger.Error("Failed to start RPC Service", err))
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
			logger.Error("Failed to supply initial Profile", err)
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
