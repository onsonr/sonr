package motor

import (
	"github.com/sonr-io/sonr/internal/motor"
	_ "golang.org/x/mobile/bind"
)

var instance *motor.MotorNode

// Start starts the host, node, and rpc service.
func New() error {
	m, err := motor.CreateAccount()
	if err != nil {
		return err
	}
	instance = m
	return nil
}
