package main

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
)

type SonrHighway struct {
	// Properties
	ctx  context.Context
	node *node.Node
}

var (
	logger      *golog.Logger
	sonrHighway *SonrHighway
)

func init() {
	logger = golog.New()
	logger.SetPrefix("[SonrHighway] ")
}

func main() {
	// Initialize Device
	ctx := context.Background()
	err := device.Init(false)
	if err != nil {
		logger.Fatal("Failed to initialize Device", err)
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, node.WithHighway())
	if err != nil {
		logger.Fatal("Failed to update Profile for Node", err)
	}
	logger.Info("Node Started: ", golog.Fields{"Response": resp.String()})

	// Set Lib
	sonrHighway = &SonrHighway{
		ctx:  ctx,
		node: n,
	}
}
