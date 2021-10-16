package sonr_lite

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/node"
	"google.golang.org/protobuf/proto"
)

type sonrLite struct {
	// Properties
	ctx  context.Context
	node common.NodeImpl
}

var (
	instance *sonrLite
)

func init() {
	golog.SetPrefix("[Sonr-Core.lib] ")
	golog.SetStacktraceLimit(2)
	golog.SetFormat("json", "    ")
}

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {
	// Prevent duplicate start
	if instance != nil {
		return
	}

	// Parse Initialize Request
	ctx := context.Background()

	// Unmarshal request
	req := &api.InitializeRequest{}
	err := proto.Unmarshal(reqBuf, req)
	if err != nil {
		golog.Errorf("Failed to unmarshal request: %v", err)
		return
	}

	// Initialize Device
	err = device.Init(req.ParseOpts()...)
	if err != nil {
		golog.Fatal("Failed to initialize Device", err)
		return
	}

	// Create Node
	n, _, err := node.NewNode(ctx, node.WithRequest(req), node.WithStubMode(node.StubMode_CLIENT))
	if err != nil {
		golog.Fatal("Failed to Create new node", err)
		return
	}

	// Set Lib
	instance = &sonrLite{
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
