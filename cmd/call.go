package bind

import (
	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnStatus(data []byte)      // Node Status Updates
	OnAPIRequest(data []byte)  // Rest Request
	OnRefreshed(data []byte)   // Lobby Updates
	OnLocalEvent(data []byte)  // Local Lobby Event
	OnRemoteEvent(data []byte) // Local Lobby Event
	OnLink(data []byte)        // Link event
	OnInvited(data []byte)     // User Invited
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// ^ invite Callback with data for Lifecycle ^ //
func (mn *Node) invited(data []byte) {
	// Update Status
	mn.setStatus(md.Status_INVITED)

	// Callback with Data
	mn.call.OnInvited(data)
}

// ^ Passes binded Methods to Node ^
func (mn *Node) callbackNode() md.NodeCallback {
	return md.NodeCallback{
		// Direct
		APIRequest:  mn.call.OnAPIRequest,
		Refreshed:   mn.call.OnRefreshed,
		LocalEvent:  mn.call.OnLocalEvent,
		RemoteEvent: mn.call.OnRemoteEvent,
		Responded:   mn.call.OnResponded,
		Progressed:  mn.call.OnProgress,

		// Middleware
		Invited:     mn.invited,
		Received:    mn.received,
		Status:      mn.setStatus,
		Transmitted: mn.transmitted,
		Error:       mn.handleError,
	}
}

// ^ transmitted Callback middleware post transfer ^ //
func (mn *Node) transmitted(card *md.Transfer) {
	// Update Status
	mn.setStatus(md.Status_AVAILABLE)

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}

	// Callback with Data
	mn.call.OnTransmitted(msgBytes)
}

// ^ received Callback middleware post transfer ^ //
func (mn *Node) received(card *md.Transfer) {
	// Update Status
	mn.setStatus(md.Status_AVAILABLE)

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(card)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		return
	}

	// Callback with Data
	mn.call.OnReceived(msgBytes)
}

// ^ handleError Callback with handleError instance, and method ^
func (mn *Node) handleError(errMsg *md.SonrError) {
	// Check for Error
	if errMsg.HasError {
		// Capture Error
		if errMsg.Capture {
			sentry.CaptureMessage(errMsg.String())
		}

		// Send Callback
		bytes := errMsg.Bytes()
		mn.call.OnError(bytes)
	}
}
