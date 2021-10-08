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
	sonrHighway *SonrHighway
)

func init() {
	golog.SetPrefix("[Sonr-Core.highway] ")
	golog.SetStacktraceLimit(2)
    golog.SetFormat("json", "    ")
}

func main() {
	// Initialize Device
	ctx := context.Background()
	err := device.Init(false)
	if err != nil {
		golog.Fatal("Failed to initialize Device", err)
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, node.WithHighway())
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", err)
	}
	golog.Info("Node Started: ", golog.Fields{"Response": resp.String()})

	// Set Lib
	sonrHighway = &SonrHighway{
		ctx:  ctx,
		node: n,
	}
}
