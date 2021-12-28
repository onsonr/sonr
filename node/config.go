package node

import (
	"errors"
	"fmt"

	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/device"
)

// Error Definitions
var (
	ErrEmptyQueue      = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery    = errors.New("No SName or PeerID provided.")
	ErrMissingParam    = errors.New("Paramater is missing.")
	ErrProtocolsNotSet = errors.New("Node Protocol has not been initialized.")
)

// Role is the type of the node (Client, Highway)
type Role int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	Role_UNSPECIFIED Role = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	Role_TEST

	// Role_MOTOR is for a Motor Node
	Role_MOTOR

	// Role_HIGHWAY is for a Highway Node
	Role_HIGHWAY
)

// Motor returns true if the node has a client.
func (m Role) IsMotor() bool {
	return m == Role_MOTOR
}

// Highway returns true if the node has a highway stub.
func (m Role) IsHighway() bool {
	return m == Role_HIGHWAY
}

// Prefix returns golog prefix for the node.
func (m Role) Prefix() string {
	var name string
	switch m {
	case Role_HIGHWAY:
		name = "highway"
	case Role_MOTOR:
		name = "motor"
	case Role_TEST:
		name = "test"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

type Configuration struct {
	deviceId         string
	location         *common.Location
	profile          *common.Profile
	connection       common.Connection
	homeDirectory    string
	supportDirectory string
	tempDirectory    string
}

func (c *Configuration) Apply(n *node) {
	// Initialize device
	device.Init(
		device.WithHomePath(c.homeDirectory),
		device.WithSupportPath(c.supportDirectory),
		device.SetDeviceID(c.deviceId),
	)
}
