package bind

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
	host    *host.SHost
	node    *node.Node
	service *node.NodeRPCService
}

var client *Client

// Start starts the host, node, and rpc service.
func Start(reqBytes []byte) {
	ctx := context.Background()
	logger.Init(true)

	// Unmarshal request
	initReq := &node.InitializeRequest{}
	err := proto.Unmarshal(reqBytes, initReq)
	if err != nil {
		panic(err)
	}

	// Get Device Paths
	opts := make([]device.FSOption, 0)
	// Check FSOptions
	if initReq.GetFsoptions() != nil {
		// Set Temporary Path
		opts = append(opts, device.FSOption{
			Path: initReq.GetFsoptions().GetCacheDir(),
			Type: device.Temporary,
		})

		// Set Documents Path
		opts = append(opts, device.FSOption{
			Path: initReq.GetFsoptions().GetDocumentsDir(),
			Type: device.Documents,
		})

		// Set Support Path
		opts = append(opts, device.FSOption{
			Path: initReq.GetFsoptions().GetSupportDir(),
			Type: device.Support,
		})
	}

	// Initialize Device
	kc, err := device.Init(opts...)
	if err != nil {
		logger.Panic("Failed to initialize Device", zap.Error(err))
	}

	host, err := host.NewHost(ctx, kc)
	if err != nil {
		logger.Panic("Failed to create Host", zap.Error(err))
	}

	// Create Node
	n := node.NewNode(ctx, host)

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
	return
}

// Stop closes the host, node, and rpc service.
func Stop() {
	client.host.Close()
	client.node.Close()
}
