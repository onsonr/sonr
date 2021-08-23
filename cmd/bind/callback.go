package bind

import "github.com/sonr-io/core/pkg/data"

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnStatus(buf []byte)   // Node Status Updates
	OnEvent(buf []byte)    // Local Lobby Event
	OnResponse(buf []byte) // Generic Response Callback
	OnRequest(buf []byte)  // Generic Request Callback
	OnError(buf []byte)    // Internal Error
}

// Passes binded Methods to Node
func (mn *Node) callback() data.Callback {
	return data.Callback{
		// Direct
		OnEvent:    mn.call.OnEvent,
		OnResponse: mn.call.OnResponse,
		OnRequest:  mn.call.OnRequest,

		// Middleware
		OnError:   mn.handleError,
		SetStatus: mn.setStatus,
	}
}

// handleError Callback with handleError instance, and method
func (mn *Node) handleError(errMsg *data.SonrError) {
	// Check for Error
	if errMsg.HasError {
		// Send Callback
		mn.call.OnError(errMsg.Marshal())
	}
}
