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

func Start(reqBytes []byte) {
	ctx := context.Background()
	logger.Init(true)

	// Unmarshal request
	initReq := &node.InitializeRequest{}
	err := proto.Unmarshal(reqBytes, initReq)
	if err != nil {
		panic(err)
	}

	logger.Info("Initialize Request", zap.String("JSON:", initReq.String()))

	// Initialize Device
	kc, err := device.Init()
	if err != nil {
		logger.Panic("Failed to initialize Device", zap.Error(err))
	}

	host, err := host.NewHost(ctx, kc)
	if err != nil {
		logger.Panic("Failed to create Host", zap.Error(err))
	}

	// Create Node
	n := node.NewNode(context.Background(), host)

	// Create RPC Service
	service, err := node.NewRPCService(n)
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
	client.node.Wait()
}
