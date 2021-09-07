package bind

// * Interface: Callback is implemented from Plugin to receive updates * //
type Callback interface {
	OnStatus(buf []byte)   // Node Status Updates
	OnEvent(buf []byte)    // Local Lobby Event
	OnResponse(buf []byte) // Generic Response Callback
	OnRequest(buf []byte)  // Generic Request Callback
	OnError(buf []byte)    // Internal Error
}
