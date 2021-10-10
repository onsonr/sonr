package lib

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"google.golang.org/protobuf/proto"
)

type sonrLib struct {
	// Properties
	ctx  context.Context
	node *node.Node
}

var (
	instance *sonrLib
)

func init() {
	golog.SetPrefix("[Sonr-Core.lib] ")
	golog.SetStacktraceLimit(2)
	golog.SetFormat("json", "    ")
}

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {
	ctx := context.Background()
	// Unmarshal request
	req := &api.InitializeRequest{}
	err := proto.Unmarshal(reqBuf, req)
	if err != nil {
		golog.Fatal("Failed to Unmarshal InitializeRequest", err)
	}

	// Initialize Device
	err = device.Init(device.WithDirectoryPaths(req.GetDeviceOptions().GetFolders()))
	if err != nil {
		golog.Fatal("Failed to initialize Device", err)
	}

	// Create Node
	n, _, err := node.NewNode(ctx, node.WithRequest(req), node.WithMode(node.Mode_CLIENT))
	if err != nil {
		golog.Fatal("Failed to Create new node", err)
	}

	// Set Lib
	instance = &sonrLib{
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
	instance.ctx.Done()
}
