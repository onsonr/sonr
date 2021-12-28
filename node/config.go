package node

import (
	"errors"

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

type Configuration struct {
	deviceId         string
	location         *common.Location
	profile          *common.Profile
	connection       common.Connection
	homeDirectory    string
	supportDirectory string
	tempDirectory    string
}

func (c *Configuration) Apply(n *Node) {
	// Initialize device
	device.Init(
		device.WithHomePath(c.homeDirectory),
		device.WithSupportPath(c.supportDirectory),
		device.SetDeviceID(c.deviceId),
	)
}
