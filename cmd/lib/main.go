package sonr

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/core/app"
	"github.com/sonr-io/core/internal/api"
	"google.golang.org/protobuf/proto"
)

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {
	// Unmarshal request
	req := &api.InitializeRequest{}
	if err := proto.Unmarshal(reqBuf, req); err != nil {
		golog.Warn("%s - Failed to unmarshal InitializeRequest. Using defaults...", err)
		req = api.DefaultInitializeRequest()
	}
	app.Start(req)
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	app.GetState().Pause()
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	app.GetState().Resume()
}

// Stop closes the host, node, and rpc service.
func Stop() {
	app.Exit(0)
}
