package sonr

import (
	//"github.com/kataras/golog"
	//"github.com/sonr-io/core/node"
	// motor "github.com/sonr-io/core/node/motor/v1"
	// "google.golang.org/protobuf/proto"
	_ "golang.org/x/mobile/bind"	
)

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {

	//	node.Start(reqBuf)
	// Unmarshal request
	// req := &motor.InitializeRequest{}
	// if err := proto.Unmarshal(reqBuf, req); err != nil {
	// 	golog.Warn("%s - Failed to unmarshal InitializeRequest. Using defaults...", err)
	// }

	// Start the app
	//node.NewHighway(context.Background(), node.WithLogLevel(node.DebugLevel))
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	//	node.Pause()
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	//	node.Resume()
}

// Stop closes the host, node, and rpc service.
func Stop() {
	//	node.Exit(0)
}
