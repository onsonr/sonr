package motor

import (
	"context"

	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/host"
	_ "golang.org/x/mobile/bind"
)

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) error {

	_, err := host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR))
	if err != nil {
		return err
	}
	return nil
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
