package models

import (
	"fmt"
	"io"
	defaultLogger "log"
	"os"
	"strings"

	"github.com/phuslu/log"
	"google.golang.org/protobuf/proto"
)

// ** ─── Error MANAGEMENT ────────────────────────────────────────────────────────
type SonrError struct {
	data     *ErrorEvent
	Capture  bool
	HasError bool
	IsJoined bool
	Error    error
	Joined   []*ErrorEvent
}

type SonrErrorOpt struct {
	Error error
	Type  ErrorEvent_Type
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
					return fmt.Fprintf(w, "(sonr_core - %s) [%s %s] \n%s", strings.ToUpper(a.Level),
						a.Time, a.Caller, a.Message)
				},
			},
		}
	}
}

// ^ Method Logs a Info Message for Event
func (t GenericEvent_Type) Log(message string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⚡️  %s", t.String())
		defaultLogger.Println("\t" + message + "\n")
	}
}

// ^ Method Logs a Info Message for Response
func (t GenericResponse_Type) Log(message string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⚡️  %s", t.String())
		defaultLogger.Println("\t" + message + "\n")
	}
}

// ^ Method Logs a Info Message for Request
func (t GenericRequest_Type) Log(message string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⚡️  %s", t.String())
		defaultLogger.Println("\t" + message + "\n")
	}
}

// ^ Method Logs an Error Message
func LogError(err error) {
	if loggerEnabled && loggerWarningEnabled {
		log.Error().Msgf("💣  %s", err.Error())
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

// ^ Method Logs a RPC Server Message
func LogRPC(event string, value interface{}) {
	ev := strings.ToUpper(event)
	val := fmt.Sprint(value)
	defaultLogger.Println(fmt.Sprintf("(SONR_RPC)-%s=%s", ev, val))
}

// ^ Method Logs a Success Message
func LogSuccess(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("✅  %s Successful", msg)
	}
}

// ^ Checks for Error With Type ^ //
func NewError(err error, errType ErrorEvent_Type) *SonrError {
	if err != nil {
		// Initialize
		message, severity := generateError(errType)

		// Set Capture
		capture := false
		if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
			capture = true
		}

		// Create Error
		serr := &SonrError{
			data: &ErrorEvent{
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
		joined := []*ErrorEvent{}
		capture := false

		// Loop Errors
		for _, err := range errors {
			// Generate Message
			message, severity := generateError(err.Type)
			if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
				capture = true
			}

			// Add Joined Message
			joined = append(joined, &ErrorEvent{
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
	message, severity := generateError(ErrorEvent_PEER_NOT_FOUND_INVITE)

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorEvent_MARSHAL,
			Severity: severity,
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
	message, severity := generateError(ErrorEvent_MARSHAL)

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorEvent_MARSHAL,
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
	message, severity := generateError(ErrorEvent_UNMARSHAL)

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorEvent_UNMARSHAL,
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
func NewErrorWithType(errType ErrorEvent_Type) *SonrError {
	// Initialize
	message, severity := generateError(errType)

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Return Error
	serr := &SonrError{
		data: &ErrorEvent{
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

// @ Return Message as Marshalled Bytes ^ //
func (errWrap *SonrError) Marshal() []byte {
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
		case ErrorEvent_LOG:
			log.Info().Msgf("😬 (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
		case ErrorEvent_WARNING:
			if loggerWarningEnabled {
				log.Warn().Msgf("⚠️ (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
				log.Info().Msg("\n")
			}
		case ErrorEvent_CRITICAL:
			if loggerCriticalEnabled {
				log.Info().Msg("\n")
				log.Error().Msgf("🚨 (%s, %s) \n Message: %s", errType, errSeverity.String(), errMsg)
				log.Info().Msg("\n")
			}
		case ErrorEvent_FATAL:
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
func (errWrap *SonrError) Message() *ErrorEvent {
	return errWrap.data
}

// @ Return Message as String ^ //
func (errWrap *SonrError) String() string {
	return errWrap.data.String()
}

// # Helper Method to Generate Client Message, Severity with Type
func generateError(errType ErrorEvent_Type) (string, ErrorEvent_Severity) {
	switch errType {
	case ErrorEvent_HOST_PUBSUB:
		return "Failed to start communication with peers", ErrorEvent_FATAL
	case ErrorEvent_HOST_START:
		return "Failed to start networking host", ErrorEvent_FATAL
	case ErrorEvent_BOOTSTRAP:
		return "Failed to bootstrap to peers", ErrorEvent_FATAL
	case ErrorEvent_CRYPTO_GEN:
		return "Failed to generate secret words", ErrorEvent_CRITICAL
	case ErrorEvent_HOST_DHT:
		return "Error occurred handling DHT", ErrorEvent_FATAL
	case ErrorEvent_HOST_KEY:
		return "Error occured managing Private Key", ErrorEvent_CRITICAL
	case ErrorEvent_HOST_STREAM:
		return "Error occurred handling Network Stream", ErrorEvent_CRITICAL
	case ErrorEvent_INCOMING:
		return "Error occurred handling Incoming File", ErrorEvent_CRITICAL
	case ErrorEvent_IP_LOCATE:
		return "Error occurred locating User", ErrorEvent_CRITICAL
	case ErrorEvent_IP_RESOLVE:
		return "Error occurred managing IP Address", ErrorEvent_FATAL
	case ErrorEvent_MARSHAL:
		return "Failed to Marshal Data", ErrorEvent_WARNING
	case ErrorEvent_OUTGOING:
		return "Error occurred handling Outgoing File", ErrorEvent_CRITICAL
	case ErrorEvent_SESSION:
		return "Error occurred managing Session", ErrorEvent_CRITICAL
	case ErrorEvent_TOPIC_HANDLER:
		return "Error occurred handling Lobby Peers", ErrorEvent_WARNING
	case ErrorEvent_TOPIC_INVALID:
		return "This Code does not exist", ErrorEvent_WARNING
	case ErrorEvent_TOPIC_JOIN:
		return "Failed to join Lobby", ErrorEvent_WARNING
	case ErrorEvent_TOPIC_CREATE:
		return "Failed to join Lobby", ErrorEvent_LOG
	case ErrorEvent_TOPIC_LEAVE:
		return "Failed to leave Lobby", ErrorEvent_LOG
	case ErrorEvent_TOPIC_MESSAGE:
		return "Failed to Send Message", ErrorEvent_WARNING
	case ErrorEvent_TOPIC_UPDATE:
		return "Failed to Send Update", ErrorEvent_LOG
	case ErrorEvent_TOPIC_RPC:
		return "Error occurred exchanging data", ErrorEvent_CRITICAL
	case ErrorEvent_TOPIC_SUB:
		return "Error occurred subscribing to Topic", ErrorEvent_CRITICAL
	case ErrorEvent_TRANSFER_CHUNK:
		return "Error occurred during Transfer", ErrorEvent_CRITICAL
	case ErrorEvent_TRANSFER_END:
		return "Error occurred finishing Transfer", ErrorEvent_CRITICAL
	case ErrorEvent_TRANSFER_START:
		return "Error occurred starting Transfer", ErrorEvent_CRITICAL
	case ErrorEvent_UNMARSHAL:
		return "Error occured Unmarshalling data", ErrorEvent_WARNING
	case ErrorEvent_USER_CREATE:
		return "Error occurred Creating User", ErrorEvent_FATAL
	case ErrorEvent_USER_FS:
		return "Error occurred Accessing File System", ErrorEvent_FATAL
	case ErrorEvent_USER_SAVE:
		return "Error occurred Saving User", ErrorEvent_CRITICAL
	case ErrorEvent_USER_LOAD:
		return "Error occurred Loading User", ErrorEvent_CRITICAL
	case ErrorEvent_USER_UPDATE:
		return "Error occurred Sending Update", ErrorEvent_WARNING
	case ErrorEvent_PEER_NOT_FOUND_INVITE:
		return "Invited Peer was not Found", ErrorEvent_LOG
	case ErrorEvent_PEER_NOT_FOUND_REPLY:
		return "Could not send Reply, Peer Not Found", ErrorEvent_LOG
	case ErrorEvent_PEER_NOT_FOUND_TRANSFER:
		return "Could not start Transfer, Peer not Found", ErrorEvent_LOG
	case ErrorEvent_URL_HTTP_GET:
		return "Invalid URL", ErrorEvent_WARNING
	case ErrorEvent_URL_INFO_RESP:
		return "Failed to parse URL Response", ErrorEvent_WARNING
	case ErrorEvent_FAILED_CONNECTION:
		return "Failed to connect to Nearby Peer", ErrorEvent_WARNING
	case ErrorEvent_HOST_INFO:
		return "Failed to generate User Peer Info", ErrorEvent_CRITICAL
	case ErrorEvent_KEY_ID:
		return "Cannot get PeerID from Public Key", ErrorEvent_CRITICAL
	case ErrorEvent_KEY_SET:
		return "Cannot overwrite existing key", ErrorEvent_WARNING
	case ErrorEvent_KEY_INVALID:
		return "Key is Invalid, May not Exist", ErrorEvent_FATAL
	case ErrorEvent_STORE_FIND:
		return "Failed to Find Key", ErrorEvent_LOG
	case ErrorEvent_STORE_GET:
		return "Failed to Get Value for Key", ErrorEvent_WARNING
	case ErrorEvent_STORE_PUT:
		return "Failed to Get Value for Key", ErrorEvent_WARNING
	case ErrorEvent_STORE_INIT:
		return "Failed to Get Value for Key", ErrorEvent_CRITICAL
	case ErrorEvent_TEXTILE_START_CLIENT:
		return "Failed to Start Textile Client", ErrorEvent_FATAL
	case ErrorEvent_TEXTILE_TOKEN_CTX:
		return "Failed to Retreive Textile Token", ErrorEvent_FATAL
	case ErrorEvent_TEXTILE_USER_CTX:
		return "Failed to Retreive Textile User", ErrorEvent_FATAL
	case ErrorEvent_THREADS_START_NEW:
		return "Failed to Start New Textile Threads", ErrorEvent_WARNING
	case ErrorEvent_THREADS_START_EXISTING:
		return "Failed to Start Existing Textile Threads", ErrorEvent_WARNING
	case ErrorEvent_THREADS_LIST_ALL:
		return "Failed to List All Threads", ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_START_NEW:
		return "Failed to Start New Mailbox", ErrorEvent_FATAL
	case ErrorEvent_MAILBOX_START_EXISTING:
		return "Failed to Start Existing Mailbox", ErrorEvent_FATAL
	case ErrorEvent_MAILBOX_LIST_ALL:
		return "Failed to List All Mailbox Messages", ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_MESSAGE_OPEN:
		return "Failed to Open Mailbox Message", ErrorEvent_CRITICAL
	case ErrorEvent_MAILBOX_MESSAGE_SEND:
		return "Failed to Send Mailbox Message", ErrorEvent_CRITICAL
	case ErrorEvent_MAILBOX_MESSAGE_PEER_PUBKEY:
		return "Failed to Find Peers Public Key", ErrorEvent_CRITICAL
	case ErrorEvent_HOST_MDNS:
		return "Failed to Start Host MDNS Discovery", ErrorEvent_WARNING
	case ErrorEvent_PEER_PUBKEY_DECODE:
		return "Failed to Decode Peer Public Key from String", ErrorEvent_WARNING
	case ErrorEvent_PEER_PUBKEY_UNMARSHAL:
		return "Failed to Unmarshal Public Key from Peers String representation", ErrorEvent_WARNING
	case ErrorEvent_DEVICE_ID:
		return "Failed to retreive Device's machine ID.", ErrorEvent_CRITICAL
	case ErrorEvent_PUSH_SINGLE:
		return "Failed to send Push Notification to peer.", ErrorEvent_WARNING
	case ErrorEvent_PUSH_MULTIPLE:
		return "Failed to send any Push Notifications to peers.", ErrorEvent_WARNING
	case ErrorEvent_PUSH_START_APP:
		return "Failed to start Firebase application.", ErrorEvent_CRITICAL
	case ErrorEvent_PUSH_START_MESSAGING:
		return "Failed to start Firebase push notification messaging.", ErrorEvent_CRITICAL
	case ErrorEvent_MAILBOX_MESSAGE_READ:
		return "Failed to read Mailbox message.", ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_MESSAGE_UNMARSHAL:
		return "Failed to Unmarshal Mailbox Message body.", ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_MESSAGE_DELETE:
		return "Failed to delete Mailbox message.", ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_ACTION_INVALID:
		return "Invalid Mailbox Action.", ErrorEvent_LOG
	case ErrorEvent_PEER_PUSH_TOKEN_EMPTY:
		return "Peer's Push Token is Empty.", ErrorEvent_WARNING
	default:
		return "Unknown", ErrorEvent_LOG
	}
}
