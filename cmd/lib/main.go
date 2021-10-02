package lib

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	// Properties
	ctx     context.Context
	host    *host.SNRHost
	node    *node.Node
	service *node.NodeRPCService
}

var client *Client

// Start starts the host, node, and rpc service.
func Start(reqBytes []byte) []byte {
	// Unmarshal request
	isDev, req, fsOpts, err := parseInitializeRequest(reqBytes)
	if err != nil {
		panic(logger.Error("Failed to Parse Initialize Request", err))
	}

	// Initialize Device
	ctx := context.Background()
	err = device.Init(isDev, fsOpts...)
	if err != nil {
		panic(logger.Error("Failed to initialize Device", err))
	}

	// Initialize Host
	host, err := host.NewHost(ctx, req.GetConnection())
	if err != nil {
		panic(logger.Error("Failed to create Host", err))
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, host, req.GetLocation())
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

	// Marshal Response
	buf, err := proto.Marshal(resp)
	if err != nil {
		logger.Error("Failed to Marshal InitializeResponse", err)
	}
	return buf
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	// if started {
	// 	// state.GetState().Pause()
	// }
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	// if started {
	// 	// state.GetState().Resume()
	// }
}

// Stop closes the host, node, and rpc service.
func Stop() {
	// if started {
	// 	client.host.Close()
	// 	client.ctx.Done()
	// }
}
