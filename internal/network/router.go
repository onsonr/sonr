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
	MinorOLC string
	MajorOLC string

	// User Joined Groups
	AuthGroups     []protocol.ID
	ExchangeGroups []protocol.ID
	TopicGroups    []protocol.ID
	TransferGroups []protocol.ID
}

// ^ Creates New Protocol Router: Grouped, Local, Global ^ //
func NewProtocolRouter(req *md.ConnectionRequest) *ProtocolRouter {
	// Get Open Location Code
	olcMinor := olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)
	olcMajor := olc.Encode(float64(req.Latitude), float64(req.Longitude), 4)

	// Return Protocol ID
	return &ProtocolRouter{
		MinorOLC: olcMinor,
		MajorOLC: olcMajor,
	}
}

// @ Host Protocol IDs
// Main Application Prefix
func (pr *ProtocolRouter) Prefix() protocol.ID {
	return protocol.ID("/sonr")
}

// Main Local Rendevouz Advertising Point
func (pr *ProtocolRouter) LocalPoint() string {
	return fmt.Sprintf("/sonr/%s", pr.MinorOLC)
}

// Main Local Rendevouz Advertising Point
func (pr *ProtocolRouter) MajorPoint() string {
	return fmt.Sprintf("/sonr/%s", pr.MajorOLC)
}

// @ Transfer Controller Auth Protocol IDs: ONLY ONE OPTION ALLOWED
func (pr *ProtocolRouter) Auth(opts ...*protocolRouterOption) protocol.ID {
	// Check Options
	if len(opts) > 0 {
		// First Value
		opt := opts[0]

		// Local Authentication Point
		if opt.local {
			return protocol.ID(fmt.Sprintf("/sonr/transfer/%s/auth", pr.MinorOLC))
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
			return protocol.ID(fmt.Sprintf("/sonr/transfer/%s/data", pr.MinorOLC))
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
			return fmt.Sprintf("/sonr/lobby/%s/topic", pr.MinorOLC)
		}
	}
	return "/sonr/lobby/topic"
}

// @ Lobby Exchange Protocol IDs
func (pr *ProtocolRouter) Exchange(pointName string) protocol.ID {
	return protocol.ID("/sonr/lobby/%/exchange")
}
