package bind

import (
	md "github.com/sonr-io/core/pkg/models"
)

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnStatus(data []byte) // Node Status Updates
	OnEvent(data []byte) // Local Lobby Event
	OnResponse(data []byte) // Generic Response Callback
	OnRequest(data []byte)  // Generic Request Callback
	OnError(data []byte)    // Internal Error
}

// # Passes binded Methods to Node
func (mn *Node) callback() md.Callback {
	return md.Callback{
		// Direct
		OnEvent:    mn.call.OnEvent,
		OnResponse: mn.call.OnResponse,
		OnRequest:  mn.call.OnRequest,

		// Middleware
		OnError:   mn.handleError,
		SetStatus: mn.setStatus,
	}
}

// # handleError Callback with handleError instance, and method
func (mn *Node) handleError(errMsg *md.SonrError) {
	// Check for Error
	if errMsg.HasError {
		// Send Callback
		mn.call.OnError(errMsg.Marshal())
	}
}
