package motor

import (
	"github.com/sonr-io/sonr/internal/motor"
	_ "golang.org/x/mobile/bind"
)

// Start starts the host, node, and rpc service.
func Start() error {
	_, err := motor.CreateAccount()
	if err != nil {
		return err
	}
	return nil
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
