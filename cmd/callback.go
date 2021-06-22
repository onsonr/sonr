package bind

import (
	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnStatus(data []byte)      // Node Status Updates
	OnConnected(data []byte)   // Connection Response
	OnEvent(data []byte)       // Local Lobby Event
	OnInvited(data []byte)     // User Invited
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data []byte)    // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// # Passes binded Methods to Node
func (mn *Node) callback() md.Callback {
	return md.Callback{
		// Direct
		OnConnected:   mn.call.OnConnected,
		OnEvent:       mn.call.OnEvent,
		OnInvite:      mn.call.OnInvited,
		OnReply:       mn.call.OnResponded,
		OnProgress:    mn.call.OnProgress,
		OnReceived:    mn.call.OnReceived,
		OnTransmitted: mn.call.OnTransmitted,

		// Middleware
		OnError:   mn.handleError,
		SetStatus: mn.setStatus,
	}
}

// # handleError Callback with handleError instance, and method
func (mn *Node) handleError(errMsg *md.SonrError) {
	// Check for Error
	if errMsg.HasError {
		// Capture Error
		if errMsg.Capture {
			sentry.CaptureMessage(errMsg.String())
		}

		// Send Callback
		mn.call.OnError(errMsg.Bytes())
	}
}
