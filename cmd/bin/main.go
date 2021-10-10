package main

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
)

type SonrBin struct {
	// Properties
	ctx  context.Context
	node *node.Node
}

var (
	sonrBin *SonrBin
)

func init() {
	golog.SetPrefix("[Sonr-Core.bin] ")
	golog.SetStacktraceLimit(2)
	golog.SetFormat("json", "    ")
}

// Start starts the host, node, and rpc service.
func main() {
	// Read Flag Values from Environment for Initialize Request

	// Initialize Device
	ctx := context.Background()
	err := device.Init(true)
	if err != nil {
		golog.Fatal("Failed to initialize Device", err)
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, node.WithMode(node.Mode_CLIENT))
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", err)
	}
	golog.Info("Node Started: ", golog.Fields{
		"Response": resp.String(),
	})

	// Set Lib
	sonrBin = &SonrBin{
		ctx:  ctx,
		node: n,
	}
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
