package bind

import (
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

func (mn *Node) isReady() bool {
	return mn.user.GetConnection().HasBootstrapped && mn.user.GetConnection().HasConnected
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

func (mn *Node) setBootstrapped(val bool) {
	// Update Status
	su := mn.user.SetBootstrapped(val)

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
