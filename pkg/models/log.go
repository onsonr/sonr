package models

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/phuslu/log"
	"google.golang.org/protobuf/proto"
)

// ** ─── Error MANAGEMENT ────────────────────────────────────────────────────────
type SonrError struct {
	data     *ErrorMessage
	Capture  bool
	HasError bool
	IsJoined bool
	Error    error
	Joined   []*ErrorMessage
}

type SonrErrorOpt struct {
	Error error
	Type  ErrorMessage_Type
}

// Logger Based Settings
var loggerEnabled = true
var loggerInfoEnabled = true
var loggerDebugEnabled = true
var loggerWarningEnabled = true
var loggerCriticalEnabled = true
var loggerFatalEnabled = true

// ^ Initializes Pretty Logger
func InitLogger(req *InitializeRequest) {
	// Check Terminal
	if log.IsTerminal(os.Stderr.Fd()) && !loggerEnabled {
		// Check Request
		if req != nil {
			// Set Logging by Preferences
			loggerEnabled = req.IsLoggingEnabled()
			loggerInfoEnabled = req.HasInfoLog()
			loggerDebugEnabled = req.HasDebugLog()
			loggerWarningEnabled = req.HasWarningLog()
			loggerCriticalEnabled = req.HasCriticalLog()
			loggerFatalEnabled = req.HasFatalLog()
		} else {
			loggerEnabled = true
		}
	}

	// Configure Logger from Enabled
	if loggerEnabled {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
				Formatter: func(w io.Writer, a *log.FormatterArgs) (int, error) {
					return fmt.Fprintf(w, "(sonr_core - %s) [%s %s] %s\n Stack: \n%s", strings.ToUpper(a.Level),
						a.Time, a.Caller, a.Message, a.Stack)
				},
			},
		}
	}
}

// ^ Method Logs a Info Message
func LogFatal(err error) {
	if !loggerEnabled {
		InitLogger(nil)
	}
	log.Fatal().Msgf("💀 %s", err.Error())
}

// ^ Method Logs a Info Message
func LogInfo(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("💡  %s", msg)
	}
}

// ^ Method Logs a Activate Message
func LogActivate(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⛷  Activating %s...", msg)
	}
}

// ^ Method Logs a Success Message
func LogSuccess(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("✅  %s Successful", msg)
	}
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

		// Create Error
		serr := &SonrError{
			data: &ErrorMessage{
				Message:  message,
				Error:    err.Error(),
				Type:     errType,
				Severity: severity,
			},
			Capture:  capture,
			HasError: true,
		}

		// Handle Logging
		serr.Log()
		return serr
	}
	// Return Error
	return &SonrError{
		HasError: false,
	}
}

// ^ Checks for Error With Type ^ //
func NewErrorGroup(errors ...SonrErrorOpt) *SonrError {
	if len(errors) > 0 {
		// Create Slice
		joined := []*ErrorMessage{}
		capture := false

		// Loop Errors
		for _, err := range errors {
			// Generate Message
			message, severity := generateError(err.Type)
			if severity == ErrorMessage_CRITICAL || severity == ErrorMessage_FATAL {
				capture = true
			}

			// Add Joined Message
			joined = append(joined, &ErrorMessage{
				Message:  message,
				Error:    err.Error.Error(),
				Type:     err.Type,
				Severity: severity,
			})
		}

		// Create Joined Error
		serr := &SonrError{
			IsJoined: true,
			HasError: true,
			Capture:  capture,
			Joined:   joined,
		}

		// Handle Logging
		serr.Log()
		return serr
	} else {
		// Return Error
		return &SonrError{
			HasError: false,
		}
	}
}

// ^ Return New Peer Not Found Error with Peer ID as Data ^ //
func NewPeerFoundError(err error, peer string) *SonrError {
	// Initialize
	message, severity := generateError(ErrorMessage_PEER_NOT_FOUND_INVITE)

	// Set Capture
	capture := false
	if severity == ErrorMessage_CRITICAL || severity == ErrorMessage_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorMessage{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorMessage_MARSHAL,
			Severity: severity,
			Data:     peer,
		},
		Capture:  capture,
		HasError: true,
	}

	// Handle Logging
	serr.Log()
	return serr
}

// ^ Returns Proto Marshal Error
func NewMarshalError(err error) *SonrError {
	// Initialize
	message, severity := generateError(ErrorMessage_MARSHAL)

	// Set Capture
	capture := false
	if severity == ErrorMessage_CRITICAL || severity == ErrorMessage_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorMessage{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorMessage_MARSHAL,
			Severity: severity,
		},
		Capture:  capture,
		HasError: true,
	}

	// Handle Logging
	serr.Log()
	return serr
}

// ^ Returns Proto Unmarshal Error
func NewUnmarshalError(err error) *SonrError {
	// Return Error
	// Initialize
	message, severity := generateError(ErrorMessage_UNMARSHAL)

	// Set Capture
	capture := false
	if severity == ErrorMessage_CRITICAL || severity == ErrorMessage_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorMessage{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorMessage_UNMARSHAL,
			Severity: severity,
		},
		Capture:  capture,
		HasError: true,
	}

	// Handle Logging
	serr.Log()
	return serr
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
	serr := &SonrError{
		data: &ErrorMessage{
			Message:  message,
			Type:     errType,
			Severity: severity,
		},
		Capture:  capture,
		HasError: true,
	}

	// Handle Logging
	serr.Log()
	return serr
}

// @ Return Message as Bytes ^ //
func (errWrap *SonrError) Bytes() []byte {
	bytes, err := proto.Marshal(errWrap.data)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Method Prints Error
func (err *SonrError) Log() {
	if loggerEnabled && loggerInfoEnabled {
		// Fetch Data
		errSeverity := err.Message().GetSeverity()
		errType := err.Message().GetType().String()
		errMsg := err.Message().GetError()

		// Start Line Break
		log.Info().Msg("\n")

		// Check Severity
		switch errSeverity {
		case ErrorMessage_LOG:
			log.Info().Msgf("😬 (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
		case ErrorMessage_WARNING:
			if loggerWarningEnabled {
				log.Warn().Msgf("⚠️ (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
				log.Info().Msg("\n")
			}
		case ErrorMessage_CRITICAL:
			if loggerCriticalEnabled {
				log.Info().Msg("\n")
				log.Error().Msgf("🚨 (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
				log.Info().Msg("\n")
			}
		case ErrorMessage_FATAL:
			if loggerFatalEnabled {
				log.Info().Msg("\n")
				log.Fatal().Msgf("💀 (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
				log.Info().Msg("\n")
			}
		}

		// End Line Break
		log.Info().Msg("\n")
	}
}

// @ Return Protobuf Message for Error
func (errWrap *SonrError) Message() *ErrorMessage {
	return errWrap.data
}

// @ Return Message as String ^ //
func (errWrap *SonrError) String() string {
	return errWrap.data.String()
}

// # Helper Method to Generate Client Message, Severity with Type
func generateError(errType ErrorMessage_Type) (string, ErrorMessage_Severity) {
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
	case ErrorMessage_URL_HTTP_GET:
		return "Invalid URL", ErrorMessage_WARNING
	case ErrorMessage_URL_INFO_RESP:
		return "Failed to parse URL Response", ErrorMessage_WARNING
	case ErrorMessage_FAILED_CONNECTION:
		return "Failed to connect to Nearby Peer", ErrorMessage_WARNING
	case ErrorMessage_HOST_INFO:
		return "Failed to generate User Peer Info", ErrorMessage_CRITICAL
	case ErrorMessage_KEY_ID:
		return "Cannot get PeerID from Public Key", ErrorMessage_CRITICAL
	case ErrorMessage_KEY_SET:
		return "Cannot overwrite existing key", ErrorMessage_WARNING
	case ErrorMessage_KEY_INVALID:
		return "Key is Invalid, May not Exist", ErrorMessage_FATAL
	case ErrorMessage_STORE_FIND:
		return "Failed to Find Key", ErrorMessage_LOG
	case ErrorMessage_STORE_GET:
		return "Failed to Get Value for Key", ErrorMessage_WARNING
	case ErrorMessage_STORE_PUT:
		return "Failed to Get Value for Key", ErrorMessage_WARNING
	case ErrorMessage_STORE_INIT:
		return "Failed to Get Value for Key", ErrorMessage_CRITICAL
	case ErrorMessage_TEXTILE_START_CLIENT:
		return "Failed to Start Textile Client", ErrorMessage_FATAL
	case ErrorMessage_TEXTILE_TOKEN_CTX:
		return "Failed to Retreive Textile Token", ErrorMessage_FATAL
	case ErrorMessage_TEXTILE_USER_CTX:
		return "Failed to Retreive Textile User", ErrorMessage_FATAL
	case ErrorMessage_THREADS_START_NEW:
		return "Failed to Start New Textile Threads", ErrorMessage_WARNING
	case ErrorMessage_THREADS_START_EXISTING:
		return "Failed to Start Existing Textile Threads", ErrorMessage_WARNING
	case ErrorMessage_THREADS_LIST_ALL:
		return "Failed to List All Threads", ErrorMessage_WARNING
	case ErrorMessage_MAILBOX_START_NEW:
		return "Failed to Start New Mailbox", ErrorMessage_FATAL
	case ErrorMessage_MAILBOX_START_EXISTING:
		return "Failed to Start Existing Mailbox", ErrorMessage_FATAL
	case ErrorMessage_MAILBOX_LIST_ALL:
		return "Failed to List All Mailbox Messages", ErrorMessage_WARNING
	case ErrorMessage_MAILBOX_MESSAGE_OPEN:
		return "Failed to Open Mailbox Message", ErrorMessage_CRITICAL
	case ErrorMessage_MAILBOX_MESSAGE_SEND:
		return "Failed to Send Mailbox Message", ErrorMessage_CRITICAL
	case ErrorMessage_MAILBOX_MESSAGE_PEER_PUBKEY:
		return "Failed to Find Peers Public Key", ErrorMessage_CRITICAL
	case ErrorMessage_HOST_MDNS:
		return "Failed to Start Host MDNS Discovery", ErrorMessage_WARNING
	case ErrorMessage_PEER_PUBKEY_DECODE:
		return "Failed to Decode Peer Public Key from String", ErrorMessage_WARNING
	case ErrorMessage_PEER_PUBKEY_UNMARSHAL:
		return "Failed to Unmarshal Public Key from Peers String representation", ErrorMessage_WARNING
	case ErrorMessage_DEVICE_ID:
		return "Failed to retreive Device's machine ID.", ErrorMessage_CRITICAL
	case ErrorMessage_PUSH_SINGLE:
		return "Failed to send Push Notification to peer.", ErrorMessage_WARNING
	case ErrorMessage_PUSH_MULTIPLE:
		return "Failed to send any Push Notifications to peers.", ErrorMessage_WARNING
	case ErrorMessage_PUSH_START_APP:
		return "Failed to start Firebase application.", ErrorMessage_CRITICAL
	case ErrorMessage_PUSH_START_MESSAGING:
		return "Failed to start Firebase push notification messaging.", ErrorMessage_CRITICAL
	default:
		return "Unknown", ErrorMessage_LOG
	}
}
