package models

import "google.golang.org/protobuf/proto"

type SonrError struct {
	data     *ErrorMessage
	Capture  bool
	HasError bool
	Error    error
}

// ^ Checks for Error With Type ^ //
func NewError(err error, errType ErrorMessage_Type) *SonrError {
	if err != nil {
		// Initialize
		message, severity := generateError(errType)

		// Set Capture
		capture := false
		if severity == ErrorMessage_CRITICAL || severity == ErrorMessage_FATAL {
			capture = true
		}

		// Return Error
		return &SonrError{
			data: &ErrorMessage{
				Message:  message,
				Error:    err.Error(),
				Type:     errType,
				Severity: severity,
			},
			Capture:  capture,
			HasError: true,
		}
	}
	// Return Error
	return &SonrError{
		HasError: false,
	}
}

// ^ Returns New Error based on Type Only
func NewErrorWithType(errType ErrorMessage_Type) *SonrError {
	// Initialize
	message, severity := generateError(errType)

	// Set Capture
	capture := false
	if severity == ErrorMessage_CRITICAL || severity == ErrorMessage_FATAL {
		capture = true
	}

	// Return Error
	return &SonrError{
		data: &ErrorMessage{
			Message:  message,
			Type:     errType,
			Severity: severity,
		},
		Capture:  capture,
		HasError: true,
	}
}

// @ Return Message as Bytes ^ //
func (errWrap *SonrError) Bytes() []byte {
	bytes, err := proto.Marshal(errWrap.data)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Return Message as String ^ //
func (errWrap *SonrError) String() string {
	return errWrap.String()
}

// # Helper Method to Generate Client Message, Severity
func generateError(errType ErrorMessage_Type) (string, ErrorMessage_Severity) {
	// Check Type
	switch errType {
	case ErrorMessage_HOST_PUBSUB:
		return "Failed to start communication with peers", ErrorMessage_FATAL
	case ErrorMessage_HOST_START:
		return "Failed to start networking host", ErrorMessage_FATAL
	case ErrorMessage_BOOTSTRAP:
		return "Failed to bootstrap to peers", ErrorMessage_FATAL
	case ErrorMessage_CRYPTO_GEN:
		return "Failed to generate secret words", ErrorMessage_CRITICAL
	case ErrorMessage_HOST_DHT:
		return "Error occurred handling DHT", ErrorMessage_FATAL
	case ErrorMessage_HOST_KEY:
		return "Error occured managing Private Key", ErrorMessage_CRITICAL
	case ErrorMessage_HOST_STREAM:
		return "Error occurred handling Network Stream", ErrorMessage_CRITICAL
	case ErrorMessage_INCOMING:
		return "Error occurred handling Incoming File", ErrorMessage_CRITICAL
	case ErrorMessage_IP_LOCATE:
		return "Error occurred locating User", ErrorMessage_CRITICAL
	case ErrorMessage_IP_RESOLVE:
		return "Error occurred managing IP Address", ErrorMessage_FATAL
	case ErrorMessage_MARSHAL:
		return "Failed to Marshal Data", ErrorMessage_WARNING
	case ErrorMessage_OUTGOING:
		return "Error occurred handling Outgoing File", ErrorMessage_CRITICAL
	case ErrorMessage_SESSION:
		return "Error occurred managing Session", ErrorMessage_CRITICAL
	case ErrorMessage_TOPIC_HANDLER:
		return "Error occurred handling Lobby Peers", ErrorMessage_WARNING
	case ErrorMessage_TOPIC_INVALID:
		return "This Code does not exist", ErrorMessage_WARNING
	case ErrorMessage_TOPIC_JOIN:
		return "Failed to join Lobby", ErrorMessage_WARNING
	case ErrorMessage_TOPIC_CREATE:
		return "Failed to join Lobby", ErrorMessage_LOG
	case ErrorMessage_TOPIC_LEAVE:
		return "Failed to leave Lobby", ErrorMessage_LOG
	case ErrorMessage_TOPIC_MESSAGE:
		return "Failed to Send Message", ErrorMessage_WARNING
	case ErrorMessage_TOPIC_UPDATE:
		return "Failed to Send Update", ErrorMessage_LOG
	case ErrorMessage_TOPIC_RPC:
		return "Error occurred exchanging data", ErrorMessage_CRITICAL
	case ErrorMessage_TOPIC_SUB:
		return "Error occurred subscribing to Topic", ErrorMessage_CRITICAL
	case ErrorMessage_TRANSFER_CHUNK:
		return "Error occurred during Transfer", ErrorMessage_CRITICAL
	case ErrorMessage_TRANSFER_END:
		return "Error occurred finishing Transfer", ErrorMessage_CRITICAL
	case ErrorMessage_TRANSFER_START:
		return "Error occurred starting Transfer", ErrorMessage_CRITICAL
	case ErrorMessage_UNMARSHAL:
		return "Error occured Unmarshalling data", ErrorMessage_WARNING
	case ErrorMessage_USER_CREATE:
		return "Error occurred Creating User", ErrorMessage_FATAL
	case ErrorMessage_USER_FS:
		return "Error occurred Accessing File System", ErrorMessage_FATAL
	case ErrorMessage_USER_SAVE:
		return "Error occurred Saving User", ErrorMessage_CRITICAL
	case ErrorMessage_USER_LOAD:
		return "Error occurred Loading User", ErrorMessage_CRITICAL
	case ErrorMessage_USER_UPDATE:
		return "Error occurred Sending Update", ErrorMessage_WARNING
	case ErrorMessage_PEER_NOT_FOUND_INVITE:
		return "Invited Peer was not Found", ErrorMessage_LOG
	case ErrorMessage_PEER_NOT_FOUND_REPLY:
		return "Could not send Reply, Peer Not Found", ErrorMessage_LOG
	case ErrorMessage_PEER_NOT_FOUND_TRANSFER:
		return "Could not start Transfer, Peer not Found", ErrorMessage_LOG
	}
	return "Unknown", ErrorMessage_LOG
}
