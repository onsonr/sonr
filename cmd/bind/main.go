package sonr

import (
	"github.com/kataras/golog"
	"github.com/sonr-io/core/app"
	"github.com/sonr-io/core/node/api"
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

	// Check Enviornment
	var logLevel app.LogLevel
	if req.Environment.IsDev() {
		logLevel = app.DebugLevel
	} else {
		logLevel = app.InfoLevel
	}

	// Start the app
	app.Start(req, app.WithLogLevel(logLevel))
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	app.Pause()
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	app.Resume()
}

// Stop closes the host, node, and rpc service.
func Stop() {
	app.Exit(0)
}
