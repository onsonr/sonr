package bind

import (
	md "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

// ** ─── Node Status Checks ────────────────────────────────────────────────────────
// # Checks if Node is Ready for Actions
func (n *Node) isReady() bool {
	return n.device.IsNotStatus(md.Status_STANDBY) || n.device.IsNotStatus(md.Status_FAILED)
}

// # Sets Node to be Connected Status
func (n *Node) setConnected(val bool) {
	// Update Status
	su := n.device.SetConnected(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}

// # Sets Node to be Available Status
func (n *Node) setAvailable(val bool) {
	// Update Status
	su := n.device.SetAvailable(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}

// # Sets Node to be (Provided) Status
func (n *Node) setStatus(newStatus md.Status) {
	// Set Status
	su := n.device.SetStatus(newStatus)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		n.handleError(md.NewError(err, md.ErrorEvent_MARSHAL))
		return
	}
	n.call.OnStatus(data)
}
