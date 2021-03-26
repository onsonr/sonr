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
	MinorOLC  string
	MajorOLC  string
	Latitude  float64
	Longitude float64
}

// ^ Creates New Protocol Router: Grouped, Local, Global ^ //
func NewProtocolRouter(req *md.ConnectionRequest) *ProtocolRouter {
	return &ProtocolRouter{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		MinorOLC:  olc.Encode(float64(req.Latitude), float64(req.Longitude), 8),
		MajorOLC:  olc.Encode(float64(req.Latitude), float64(req.Longitude), 4),
	}
}

// @ Local Rendevouz Advertising Point
func (pr *ProtocolRouter) LocalPoint() string {
	return fmt.Sprintf("/sonr/%s", pr.MinorOLC)
}

// @ Major Rendevouz Advertising Point
func (pr *ProtocolRouter) MajorPoint() string {
	return fmt.Sprintf("/sonr/%s", pr.MajorOLC)
}

// ^ Transfer Protocols ^ //
// @ Transfer Controller Data Protocol IDs
func (pr *ProtocolRouter) Transfer() protocol.ID {
	return protocol.ID("/sonr/transfer/data")
}

// ^ Lobby Protocols ^ //
// @ Local Lobby Topic Protocol IDs
func (pr *ProtocolRouter) LocalTopic() string {
	return fmt.Sprintf("/sonr/topic/%s", pr.MinorOLC)
}

// @ Lobby Topic Protocol IDs
func (pr *ProtocolRouter) Topic(name string) string {
	return fmt.Sprintf("/sonr/topic/%s", name)
}

// @ Lobby Exchange Protocol IDs
func (pr *ProtocolRouter) TopicService() protocol.ID {
	return protocol.ID("/sonr/topic/service")
}
