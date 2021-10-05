package main

import (
	"context"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

type SonrHighway struct {
	// Properties
	ctx  context.Context
	node *node.Node
}

var sonrHighway *SonrHighway

func main() {
	// Initialize Device
	ctx := context.Background()
	err := device.Init(false)
	if err != nil {
		panic(logger.Error("Failed to initialize Device", err))
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, node.WithHighway())
	if err != nil {
		panic(logger.Error("Failed to update Profile for Node", err))
	}
	logger.Info("Node Started: ", zap.Any("Response", resp))

	// Set Lib
	sonrHighway = &SonrHighway{
		ctx:  ctx,
		node: n,
	}
}
