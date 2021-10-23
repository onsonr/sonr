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
	err := proto.Unmarshal(reqBuf, req)
	if err != nil {
		golog.Fatalf("%s - Failed to unmarshal InitializeRequest", err)
		return
	}
	app.Start(req)
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
	app.Exit(0)
}
