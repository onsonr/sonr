package network

import (
	"fmt"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/protocol"
	md "github.com/sonr-io/core/internal/models"
)

// ! Protocol Router for routing Sonr Endpoints by Module !
type ProtocolRouter struct {
	// Open Location Code for Local Peers
	OLC string

	// User Joined Groups
	AuthGroups     []protocol.ID
	ExchangeGroups []protocol.ID
	TopicGroups    []protocol.ID
	TransferGroups []protocol.ID
}

// ^ Creates New Protocol Router: Grouped, Local, Global ^ //
func NewProtocolRouter(req *md.ConnectionRequest) *ProtocolRouter {
	// Get Open Location Code
	olcValue := olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)

	// Return Protocol ID
	return &ProtocolRouter{
		OLC: olcValue,
	}
}

// @ Host Protocol IDs
// Main Application Prefix
func (pr *ProtocolRouter) Prefix() protocol.ID {
	return protocol.ID("/sonr")
}

// Main Local Rendevouz Advertising Point
func (pr *ProtocolRouter) Point() string {
	return fmt.Sprintf("/sonr/%s", pr.OLC)
}

// @ Transfer Controller Auth Protocol IDs: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Auth(opts ...*protocolRouterOption) protocol.ID {
	// Check Options
	if len(opts) > 0 {
		// First Value
		opt := opts[0]

		// Local Authentication Point
		if opt.local {
			return protocol.ID(fmt.Sprintf("/sonr/transfer/%s/auth", pr.OLC))
		}
	}

	// Return Default
	return protocol.ID("/sonr/transfer/auth")
}

// @ Remote Point Topic: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Direct(n string) protocol.ID {
	// Return Default
	return protocol.ID(fmt.Sprintf("/sonr/remote/%s/direct", n))
}

// @ Remote Point Topic: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Remote(n string) string {
	// Return Default
	return fmt.Sprintf("/sonr/remote/%s/topic", n)
}

// @ Transfer Controller Data Protocol IDs: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Transfer(opts ...*protocolRouterOption) protocol.ID {
	// Check Options
	if len(opts) > 0 {
		// First Value
		opt := opts[0]

		// Local Authentication Point
		if opt.local {
			return protocol.ID(fmt.Sprintf("/sonr/transfer/%s/data", pr.OLC))
		}
	}

	// Return Default
	return protocol.ID("/sonr/transfer/data")
}

// @ Lobby Topic Protocol IDs: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Topic(opts ...*protocolRouterOption) string {
	// Check Options
	if len(opts) > 0 {
		// First Value
		opt := opts[0]

		// Local Authentication Point
		if opt.local {
			return fmt.Sprintf("/sonr/lobby/%s/topic", pr.OLC)
		}

		// Group Auth Point
		if opt.group {
			return fmt.Sprintf("/sonr/lobby/group/%s/topic", opt.groupName)
		}
	}
	return "/sonr/lobby/topic"
}

// @ Lobby Exchange Protocol IDs: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Exchange(opts ...*protocolRouterOption) protocol.ID {
	// Check Options
	if len(opts) > 0 {
		// First Value
		opt := opts[0]

		// Local Exchange Point
		if opt.local {
			return protocol.ID(fmt.Sprintf("/sonr/lobby/%s/exchange", pr.OLC))
		}

		// Group Exchange Point
		if opt.group {
			return protocol.ID(fmt.Sprintf("/sonr/lobby/group/%s/exchange", opt.groupName))
		}
	}
	return protocol.ID("/sonr/lobby/exchange")
}
