package main

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/tools/state"
)

type Sonr struct {
	// Properties
	ctx     context.Context
	node    common.NodeImpl
	emitter *state.Emitter
}

var (
	sonrHighway *Sonr
)

func init() {
	golog.SetPrefix("[Sonr-Core.highway] ")
	golog.SetStacktraceLimit(2)
	golog.SetFormat("json", "    ")
}

func main() {
	// Initialize Device
	ctx := context.Background()
	emitter := state.NewEmitter(2048)

	err := device.Init()
	if err != nil {
		golog.Fatal("Failed to initialize Device", err)
	}

	// Create Node
	n, resp, err := node.NewNode(ctx, node.WithStubMode(node.StubMode_HIGHWAY))
	if err != nil {
		golog.Fatal("Failed to update Profile for Node", err)
	}
	golog.Info("Node Started: ", golog.Fields{"Response": resp.String()})

	// Set Lib
	sonrHighway = &Sonr{
		ctx:     ctx,
		node:    n,
		emitter: emitter,
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
