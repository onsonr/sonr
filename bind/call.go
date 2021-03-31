package bind

import (
	"log"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnStatus(data []byte)      // Node Status Updates
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ Passes binded Methods to Node ^
func (mn *Node) callbackNode() md.NodeCallback {
	return md.NodeCallback{
		// Direct
		Refreshed:  mn.call.OnRefreshed,
		Event:      mn.call.OnEvent,
		Responded:  mn.call.OnResponded,
		Progressed: mn.call.OnProgress,

		// Middleware
		Invited:     mn.invited,
		Received:    mn.received,
		Transmitted: mn.transmitted,
		Error:       mn.error,
	}
}

// ^ invite Callback with data for Lifecycle ^ //
func (mn *Node) invited(data []byte) {
	// Update Status
	mn.setStatus(md.Status_INVITED)

	// Check Invite for FlatContact
	invite := &md.AuthInvite{}
	err := proto.Unmarshal(data, invite)
	if err != nil {
		log.Println(err)
	}

	if invite.IsFlat {
		mn.node.Respond(true, mn.local, mn.user.FS, mn.user.Contact(), true)
	}

	// Callback with Data
	mn.call.OnInvited(data)
}

// ^ transmitted Callback middleware post transfer ^ //
func (mn *Node) transmitted(peer *md.Peer) {
	// Update Status
	mn.setStatus(md.Status_AVAILABLE)

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(peer)
	if err != nil {
		// sentry.CaptureException(err)
		log.Println(err)
		return
	}

	// Callback with Data
	mn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (mn *Node) received(card *md.TransferCard) {
	// Update Status
	mn.setStatus(md.Status_AVAILABLE)

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		// sentry.CaptureException(err)
		log.Println(err)
		return
	}

	// Callback with Data
	mn.call.OnReceived(msgBytes)
}

// ^ error Callback with error instance, and method ^
func (mn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println(err)
		return
	}
	// Send Callback
	mn.call.OnError(bytes)
}
