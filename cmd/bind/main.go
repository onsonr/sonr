package bind

import (
	"context"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	// Properties
	ctx     context.Context
	host    *host.SHost
	node    *node.Node
	service *node.NodeRPCService
}

var client *Client
var started bool

// Start starts the host, node, and rpc service.
func Start(reqBytes []byte) {
	// Check if already started
	if !started {
		ctx := context.Background()
		logger.Init(true)

		// Unmarshal request
		req := &node.InitializeRequest{}
		err := proto.Unmarshal(reqBytes, req)
		if err != nil {
			panic(err)
		}

		// Get Device Paths
		fsOpts := make([]device.FSOption, 0)
		// Check FSOptions
		if req.GetFsoptions() != nil {
			// Set Temporary Path
			fsOpts = append(fsOpts, device.FSOption{
				Path: req.GetFsoptions().GetCacheDir(),
				Type: device.Temporary,
			})

			// Set Documents Path
			fsOpts = append(fsOpts, device.FSOption{
				Path: req.GetFsoptions().GetDocumentsDir(),
				Type: device.Documents,
			})

			// Set Support Path
			fsOpts = append(fsOpts, device.FSOption{
				Path: req.GetFsoptions().GetSupportDir(),
				Type: device.Support,
			})
		}

		// Initialize Device
		kc, err := device.Init(fsOpts...)
		if err != nil {
			logger.Panic("Failed to initialize Device", zap.Error(err))
		}

		// Initialize Host
		host, err := host.NewHost(ctx, kc)
		if err != nil {
			logger.Panic("Failed to create Host", zap.Error(err))
		}

		// Create Node
		n := node.NewNode(ctx, host, req.GetLocation())

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
	}
	return
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	if started {
		state.GetState().Pause()
	}
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	if started {
		state.GetState().Resume()
	}
}

// Stop closes the host, node, and rpc service.
func Stop() {
	if started {
		client.host.Close()
		client.ctx.Done()
	}
}
