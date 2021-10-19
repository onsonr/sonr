package sonr

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"google.golang.org/protobuf/proto"
)

type sonrLite struct {
	// Properties
	ctx  context.Context
	node api.NodeImpl
}

var (
	instance *sonrLite
)

func init() {
	golog.SetPrefix("[Sonr-Core.lite] ")
	golog.SetStacktraceLimit(2)
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
	err = req.Parse()
	if err != nil {
		golog.Errorf("Failed to parse and handle request: %v", err)
		return
	}

	// Create Node
	n, _, err := node.NewNode(ctx, node.WithRequest(req))
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
