package lib

import (
	"context"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
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
var started bool

// Start starts the host, node, and rpc service.
func Start(reqBytes []byte) *Client {
	// Check if already started
	if !started {
		// Unmarshal request
		req, fsOpts, err := parseInitializeRequest(reqBytes)
		if err != nil {
			logger.Fatal("Failed to Parse InitializeRequest", zap.Error(err))
		}

		// Initialize Device
		kc, err := device.Init(req.GetEnvironment().IsDev(), fsOpts...)
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
		host, err := host.NewHost(ctx, kc, req.GetConnection())
		if err != nil {
			logger.Panic("Failed to create Host", zap.Error(err))
		}

		// Create Node and Start Service
		n := node.NewNode(ctx, host, req.GetLocation(), req.GetProfile())
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
	}
	return client
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

// parseInitializeRequest parses the given buffer and returns the proto and fsOptions.
func parseInitializeRequest(buf []byte) (*node.InitializeRequest, []device.FSOption, error) {
	// Unmarshal request
	req := &node.InitializeRequest{}
	err := proto.Unmarshal(buf, req)
	if err != nil {
		return nil, nil, err
	}

	// Check FSOptions and Get Device Paths
	fsOpts := make([]device.FSOption, 0)
	if req.GetDeviceOptions() != nil {
		// Set Temporary Path
		fsOpts = append(fsOpts, device.FSOption{
			Path: req.GetDeviceOptions().GetCacheDir(),
			Type: device.Temporary,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetDocumentsDir(),
			Type: device.Documents,
		}, device.FSOption{
			Path: req.GetDeviceOptions().GetSupportDir(),
			Type: device.Support,
		})
	}
	return req, fsOpts, nil
}
