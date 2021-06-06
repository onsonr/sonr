package bind

import (
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

func (mn *Node) isReady() bool {
	return mn.user.IsNotStatus(md.Status_STANDBY) || mn.user.IsNotStatus(md.Status_FAILED)
}

func (mn *Node) setConnected(val bool) {
	// Update Status
	su := mn.user.SetConnected(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setAvailable(val bool) {
	// Update Status
	su := mn.user.SetAvailable(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setStatus(newStatus md.Status) {
	// Set Status
	su := mn.user.SetStatus(newStatus)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}
