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

// Initializes Pretty Logger
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

// Method Logs a Info Message for Event
func (t GenericEvent_Type) Log(message string) {
	if loggerEnabled && loggerInfoEnabled && t != GenericEvent_ROOM {
		log.Info().Msgf("⚡️  %s", t.String())
		defaultLogger.Println("\t" + message + "\n")
	}
}

// Method Logs a Info Message for Response
func (t GenericResponse_Type) Log(message string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⚡️  %s", t.String())
		defaultLogger.Println("\t" + message + "\n")
	}
}

// Method Logs a Info Message for Request
func (t GenericRequest_Type) Log(message string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⚡️  %s", t.String())
		defaultLogger.Println("\t" + message + "\n")
	}
}

// Method Logs an Error Message
func LogError(err error) {
	if loggerEnabled && loggerWarningEnabled {
		log.Error().Msgf("💣  %s", err.Error())
	}
}

// Method Logs a Info Message
func LogFatal(err error) {
	if !loggerEnabled {
		InitLogger(nil)
	}
	log.Fatal().Msgf("💀 %s", err.Error())
}

// Method Logs a Info Message
func LogInfo(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("💡  %s", msg)
	}
}

// Method Logs a Activate Message
func LogActivate(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("⛷  Activating %s...", msg)
	}
}

// Method Logs a RPC Server Message
func LogRPC(event string, value interface{}) {
	ev := strings.ToUpper(event)
	val := fmt.Sprint(value)
	defaultLogger.Println(fmt.Sprintf("(SONR_RPC)-%s=%s", ev, val))
}

// Method Logs a Success Message
func LogSuccess(msg string) {
	if loggerEnabled && loggerInfoEnabled {
		log.Info().Msgf("✅  %s Successful", msg)
	}
}

// Checks for Error With Type ^ //
func NewError(err error, errType ErrorEvent_Type) *SonrError {
	if err != nil {
		// Initialize
		severity := errType.Severity()

		// Set Capture
		capture := false
		if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
			capture = true
		}

		// Create Error
		serr := &SonrError{
			data: &ErrorEvent{
				Message:  errType.Message(),
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

// Checks for Error With Type ^ //
func NewErrorGroup(errors ...SonrErrorOpt) *SonrError {
	if len(errors) > 0 {
		// Create Slice
		joined := []*ErrorEvent{}
		capture := false

		// Loop Errors
		for _, err := range errors {
			// Generate Message
			severity := err.Type.Severity()
			if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
				capture = true
			}

			// Add Joined Message
			joined = append(joined, &ErrorEvent{
				Message:  err.Type.Message(),
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

// Return New Peer Not Found Error with Peer ID as Data ^ //
func NewPeerFoundError(err error, peer string) *SonrError {
	// Initialize
	severity := ErrorEvent_PEER_NOT_FOUND_INVITE.Severity()

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  ErrorEvent_PEER_NOT_FOUND_INVITE.Message(),
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

// Returns Proto Marshal Error
func NewMarshalError(err error) *SonrError {
	// Initialize
	severity := ErrorEvent_MARSHAL.Severity()

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  ErrorEvent_MARSHAL.Message(),
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

// Returns Proto Unmarshal Error
func NewUnmarshalError(err error) *SonrError {
	// Return Error
	// Initialize
	severity := ErrorEvent_UNMARSHAL.Severity()

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Create Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  ErrorEvent_UNMARSHAL.Message(),
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

// Returns New Error based on Type Only
func NewErrorWithType(errType ErrorEvent_Type) *SonrError {
	// Initialize
	severity := errType.Severity()

	// Set Capture
	capture := false
	if severity == ErrorEvent_CRITICAL || severity == ErrorEvent_FATAL {
		capture = true
	}

	// Return Error
	serr := &SonrError{
		data: &ErrorEvent{
			Message:  errType.Message(),
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

// Return Message as Marshalled Bytes ^ //
func (errWrap *SonrError) Marshal() []byte {
	bytes, err := proto.Marshal(errWrap.data)
	if err != nil {
		return nil
	}
	return bytes
}

// Method Prints Error
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

// Return Protobuf Message for Error
func (errWrap *SonrError) Message() *ErrorEvent {
	return errWrap.data
}

// Return Message as String ^ //
func (errWrap *SonrError) String() string {
	return errWrap.data.String()
}

func (et ErrorEvent_Type) Severity() ErrorEvent_Severity {
	switch et {
	case ErrorEvent_HOST_PUBSUB:
		return ErrorEvent_FATAL
	case ErrorEvent_HOST_START:
		return ErrorEvent_FATAL
	case ErrorEvent_BOOTSTRAP:
		return ErrorEvent_FATAL
	case ErrorEvent_CRYPTO_GEN:
		return ErrorEvent_CRITICAL
	case ErrorEvent_HOST_DHT:
		return ErrorEvent_FATAL
	case ErrorEvent_HOST_KEY:
		return ErrorEvent_CRITICAL
	case ErrorEvent_HOST_STREAM:
		return ErrorEvent_CRITICAL
	case ErrorEvent_INCOMING:
		return ErrorEvent_CRITICAL
	case ErrorEvent_IP_LOCATE:
		return ErrorEvent_CRITICAL
	case ErrorEvent_IP_RESOLVE:
		return ErrorEvent_FATAL
	case ErrorEvent_MARSHAL:
		return ErrorEvent_WARNING
	case ErrorEvent_OUTGOING:
		return ErrorEvent_CRITICAL
	case ErrorEvent_SESSION:
		return ErrorEvent_CRITICAL
	case ErrorEvent_ROOM_HANDLER:
		return ErrorEvent_WARNING
	case ErrorEvent_ROOM_INVALID:
		return ErrorEvent_WARNING
	case ErrorEvent_ROOM_JOIN:
		return ErrorEvent_WARNING
	case ErrorEvent_ROOM_CREATE:
		return ErrorEvent_LOG
	case ErrorEvent_ROOM_LEAVE:
		return ErrorEvent_LOG
	case ErrorEvent_ROOM_MESSAGE:
		return ErrorEvent_WARNING
	case ErrorEvent_ROOM_UPDATE:
		return ErrorEvent_LOG
	case ErrorEvent_ROOM_RPC:
		return ErrorEvent_CRITICAL
	case ErrorEvent_ROOM_SUB:
		return ErrorEvent_CRITICAL
	case ErrorEvent_TRANSFER_CHUNK:
		return ErrorEvent_CRITICAL
	case ErrorEvent_TRANSFER_END:
		return ErrorEvent_CRITICAL
	case ErrorEvent_TRANSFER_START:
		return ErrorEvent_CRITICAL
	case ErrorEvent_UNMARSHAL:
		return ErrorEvent_WARNING
	case ErrorEvent_USER_CREATE:
		return ErrorEvent_FATAL
	case ErrorEvent_USER_FS:
		return ErrorEvent_FATAL
	case ErrorEvent_USER_SAVE:
		return ErrorEvent_CRITICAL
	case ErrorEvent_USER_LOAD:
		return ErrorEvent_CRITICAL
	case ErrorEvent_USER_UPDATE:
		return ErrorEvent_WARNING
	case ErrorEvent_PEER_NOT_FOUND_INVITE:
		return ErrorEvent_LOG
	case ErrorEvent_PEER_NOT_FOUND_REPLY:
		return ErrorEvent_LOG
	case ErrorEvent_PEER_NOT_FOUND_TRANSFER:
		return ErrorEvent_LOG
	case ErrorEvent_URL_HTTP_GET:
		return ErrorEvent_WARNING
	case ErrorEvent_URL_INFO_RESP:
		return ErrorEvent_WARNING
	case ErrorEvent_FAILED_CONNECTION:
		return ErrorEvent_WARNING
	case ErrorEvent_HOST_INFO:
		return ErrorEvent_CRITICAL
	case ErrorEvent_KEY_ID:
		return ErrorEvent_CRITICAL
	case ErrorEvent_KEY_SET:
		return ErrorEvent_WARNING
	case ErrorEvent_KEY_INVALID:
		return ErrorEvent_FATAL
	case ErrorEvent_STORE_FIND:
		return ErrorEvent_LOG
	case ErrorEvent_STORE_GET:
		return ErrorEvent_WARNING
	case ErrorEvent_STORE_PUT:
		return ErrorEvent_WARNING
	case ErrorEvent_STORE_INIT:
		return ErrorEvent_CRITICAL
	case ErrorEvent_TEXTILE_START_CLIENT:
		return ErrorEvent_FATAL
	case ErrorEvent_TEXTILE_TOKEN_CTX:
		return ErrorEvent_FATAL
	case ErrorEvent_TEXTILE_USER_CTX:
		return ErrorEvent_FATAL
	case ErrorEvent_THREADS_START_NEW:
		return ErrorEvent_WARNING
	case ErrorEvent_THREADS_START_EXISTING:
		return ErrorEvent_WARNING
	case ErrorEvent_THREADS_LIST_ALL:
		return ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_START_NEW:
		return ErrorEvent_FATAL
	case ErrorEvent_MAILBOX_START_EXISTING:
		return ErrorEvent_FATAL
	case ErrorEvent_MAILBOX_LIST_ALL:
		return ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_MESSAGE_OPEN:
		return ErrorEvent_CRITICAL
	case ErrorEvent_MAILBOX_MESSAGE_SEND:
		return ErrorEvent_CRITICAL
	case ErrorEvent_MAILBOX_MESSAGE_PEER_PUBKEY:
		return ErrorEvent_CRITICAL
	case ErrorEvent_HOST_MDNS:
		return ErrorEvent_WARNING
	case ErrorEvent_PEER_PUBKEY_DECODE:
		return ErrorEvent_WARNING
	case ErrorEvent_PEER_PUBKEY_UNMARSHAL:
		return ErrorEvent_WARNING
	case ErrorEvent_DEVICE_ID:
		return ErrorEvent_CRITICAL
	case ErrorEvent_PUSH_SINGLE:
		return ErrorEvent_WARNING
	case ErrorEvent_PUSH_MULTIPLE:
		return ErrorEvent_WARNING
	case ErrorEvent_PUSH_START_APP:
		return ErrorEvent_CRITICAL
	case ErrorEvent_PUSH_START_MESSAGING:
		return ErrorEvent_CRITICAL
	case ErrorEvent_MAILBOX_MESSAGE_READ:
		return ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_MESSAGE_UNMARSHAL:
		return ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_MESSAGE_DELETE:
		return ErrorEvent_WARNING
	case ErrorEvent_MAILBOX_ACTION_INVALID:
		return ErrorEvent_LOG
	case ErrorEvent_PEER_PUSH_TOKEN_EMPTY:
		return ErrorEvent_WARNING
	case ErrorEvent_LINK_GENERATE:
		return ErrorEvent_WARNING
	case ErrorEvent_LINK_SHARED_KEY:
		return ErrorEvent_CRITICAL
	case ErrorEvent_ACCOUNT_CREATE:
		return ErrorEvent_CRITICAL
	case ErrorEvent_ACCOUNT_SAVE:
		return ErrorEvent_CRITICAL
	case ErrorEvent_ACCOUNT_LOAD:
		return ErrorEvent_CRITICAL
	default:
		return ErrorEvent_LOG
	}
}

func (et ErrorEvent_Type) Message() string {
	switch et {
	case ErrorEvent_HOST_PUBSUB:
		return "Failed to start communication with peers"
	case ErrorEvent_HOST_START:
		return "Failed to start networking host"
	case ErrorEvent_BOOTSTRAP:
		return "Failed to bootstrap to peers"
	case ErrorEvent_CRYPTO_GEN:
		return "Failed to generate secret words"
	case ErrorEvent_HOST_DHT:
		return "Error occurred handling DHT"
	case ErrorEvent_HOST_KEY:
		return "Error occured managing Private Key"
	case ErrorEvent_HOST_STREAM:
		return "Error occurred handling Network Stream"
	case ErrorEvent_INCOMING:
		return "Error occurred handling Incoming File"
	case ErrorEvent_IP_LOCATE:
		return "Error occurred locating User"
	case ErrorEvent_IP_RESOLVE:
		return "Error occurred managing IP Address"
	case ErrorEvent_MARSHAL:
		return "Failed to Marshal Data"
	case ErrorEvent_OUTGOING:
		return "Error occurred handling Outgoing File"
	case ErrorEvent_SESSION:
		return "Error occurred managing Session"
	case ErrorEvent_ROOM_HANDLER:
		return "Error occurred handling Lobby Peers"
	case ErrorEvent_ROOM_INVALID:
		return "This Code does not exist"
	case ErrorEvent_ROOM_JOIN:
		return "Failed to join Lobby"
	case ErrorEvent_ROOM_CREATE:
		return "Failed to join Lobby"
	case ErrorEvent_ROOM_LEAVE:
		return "Failed to leave Lobby"
	case ErrorEvent_ROOM_MESSAGE:
		return "Failed to Send Message"
	case ErrorEvent_ROOM_UPDATE:
		return "Failed to Send Update"
	case ErrorEvent_ROOM_RPC:
		return "Error occurred exchanging data"
	case ErrorEvent_ROOM_SUB:
		return "Error occurred subscribing to Room"
	case ErrorEvent_TRANSFER_CHUNK:
		return "Error occurred during Transfer"
	case ErrorEvent_TRANSFER_END:
		return "Error occurred finishing Transfer"
	case ErrorEvent_TRANSFER_START:
		return "Error occurred starting Transfer"
	case ErrorEvent_UNMARSHAL:
		return "Error occured Unmarshalling data"
	case ErrorEvent_USER_CREATE:
		return "Error occurred Creating User"
	case ErrorEvent_USER_FS:
		return "Error occurred Accessing File System"
	case ErrorEvent_USER_SAVE:
		return "Error occurred Saving User"
	case ErrorEvent_USER_LOAD:
		return "Error occurred Loading User"
	case ErrorEvent_USER_UPDATE:
		return "Error occurred Sending Update"
	case ErrorEvent_PEER_NOT_FOUND_INVITE:
		return "Invited Peer was not Found"
	case ErrorEvent_PEER_NOT_FOUND_REPLY:
		return "Could not send Reply, Peer Not Found"
	case ErrorEvent_PEER_NOT_FOUND_TRANSFER:
		return "Could not start Transfer, Peer not Found"
	case ErrorEvent_URL_HTTP_GET:
		return "Invalid URL"
	case ErrorEvent_URL_INFO_RESP:
		return "Failed to parse URL Response"
	case ErrorEvent_FAILED_CONNECTION:
		return "Failed to connect to Nearby Peer"
	case ErrorEvent_HOST_INFO:
		return "Failed to generate User Peer Info"
	case ErrorEvent_KEY_ID:
		return "Cannot get PeerID from Public Key"
	case ErrorEvent_KEY_SET:
		return "Cannot overwrite existing key"
	case ErrorEvent_KEY_INVALID:
		return "Key is Invalid, May not Exist"
	case ErrorEvent_STORE_FIND:
		return "Failed to Find Key"
	case ErrorEvent_STORE_GET:
		return "Failed to Get Value for Key"
	case ErrorEvent_STORE_PUT:
		return "Failed to Get Value for Key"
	case ErrorEvent_STORE_INIT:
		return "Failed to Get Value for Key"
	case ErrorEvent_TEXTILE_START_CLIENT:
		return "Failed to Start Textile Client"
	case ErrorEvent_TEXTILE_TOKEN_CTX:
		return "Failed to Retreive Textile Token"
	case ErrorEvent_TEXTILE_USER_CTX:
		return "Failed to Retreive Textile User"
	case ErrorEvent_THREADS_START_NEW:
		return "Failed to Start New Textile Threads"
	case ErrorEvent_THREADS_START_EXISTING:
		return "Failed to Start Existing Textile Threads"
	case ErrorEvent_THREADS_LIST_ALL:
		return "Failed to List All Threads"
	case ErrorEvent_MAILBOX_START_NEW:
		return "Failed to Start New Mailbox"
	case ErrorEvent_MAILBOX_START_EXISTING:
		return "Failed to Start Existing Mailbox"
	case ErrorEvent_MAILBOX_LIST_ALL:
		return "Failed to List All Mailbox Messages"
	case ErrorEvent_MAILBOX_MESSAGE_OPEN:
		return "Failed to Open Mailbox Message"
	case ErrorEvent_MAILBOX_MESSAGE_SEND:
		return "Failed to Send Mailbox Message"
	case ErrorEvent_MAILBOX_MESSAGE_PEER_PUBKEY:
		return "Failed to Find Peers Public Key"
	case ErrorEvent_HOST_MDNS:
		return "Failed to Start Host MDNS Discovery"
	case ErrorEvent_PEER_PUBKEY_DECODE:
		return "Failed to Decode Peer Public Key from String"
	case ErrorEvent_PEER_PUBKEY_UNMARSHAL:
		return "Failed to Unmarshal Public Key from Peers String representation"
	case ErrorEvent_DEVICE_ID:
		return "Failed to retreive Device's machine ID."
	case ErrorEvent_PUSH_SINGLE:
		return "Failed to send Push Notification to peer."
	case ErrorEvent_PUSH_MULTIPLE:
		return "Failed to send any Push Notifications to peers."
	case ErrorEvent_PUSH_START_APP:
		return "Failed to start Firebase application."
	case ErrorEvent_PUSH_START_MESSAGING:
		return "Failed to start Firebase push notification messaging."
	case ErrorEvent_MAILBOX_MESSAGE_READ:
		return "Failed to read Mailbox message."
	case ErrorEvent_MAILBOX_MESSAGE_UNMARSHAL:
		return "Failed to Unmarshal Mailbox Message body."
	case ErrorEvent_MAILBOX_MESSAGE_DELETE:
		return "Failed to delete Mailbox message."
	case ErrorEvent_MAILBOX_ACTION_INVALID:
		return "Invalid Mailbox Action."
	case ErrorEvent_PEER_PUSH_TOKEN_EMPTY:
		return "Peer's Push Token is Empty."
	case ErrorEvent_LINK_GENERATE:
		return "Error occurred Generating Shared Key Func."
	case ErrorEvent_LINK_SHARED_KEY:
		return "Error occurred retreiving Shared Key"
	case ErrorEvent_ACCOUNT_CREATE:
		return "Error occurred Creating Account"
	case ErrorEvent_ACCOUNT_SAVE:
		return "Error occurred Saving Account"
	case ErrorEvent_ACCOUNT_LOAD:
		return "Error occurred Loading Account"
	default:
		return "Unknown"
	}
}
