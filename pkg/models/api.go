package models

import (
	"fmt"
	"net/http"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/protobuf/proto"
)

// ** ─── ConnectionRequest MANAGEMENT ────────────────────────────────────────────────────────
// Checks for Auth Type
func (cr *ConnectionRequest) IsLink() bool {
	return cr.GetType() == ConnectionRequest_LINK
}

// Checks for Auth Type
func (cr *ConnectionRequest) IsStorj() bool {
	return cr.GetType() == ConnectionRequest_STORJ
}

// Checks for Connect Type
func (cr *ConnectionRequest) IsConnect() bool {
	return cr.GetType() == ConnectionRequest_CONNECT
}

// ** ─── Store MANAGEMENT ────────────────────────────────────────────────────────
func (sk StoreKeys) Bytes() []byte {
	return []byte(sk.String())
}

// ** ─── InviteResponse MANAGEMENT ────────────────────────────────────────────────────────
func (r *InviteResponse) HasAcceptedTransfer() bool {
	return r.Decision && r.Type == InviteResponse_Transfer
}

// ** ─── InviteRequest MANAGEMENT ────────────────────────────────────────────────────────
// Returns Invite Contact
func (i *InviteRequest) GetContact() *Contact {
	return i.GetData().GetContact()
}

// Returns Invite File
func (i *InviteRequest) GetFile() *SonrFile {
	return i.GetData().GetFile()
}

// Returns Invite URL
func (i *InviteRequest) GetUrl() *URLLink {
	return i.GetData().GetUrl()
}

// Checks if Payload is Contact
func (i *InviteRequest) IsPayloadContact() bool {
	return i.Payload == Payload_CONTACT || i.Payload == Payload_FLAT_CONTACT
}

// Checks if Payload is File Transfer
func (i *InviteRequest) IsPayloadFile() bool {
	return i.Payload == Payload_FILE || i.Payload == Payload_FILES || i.Payload == Payload_MEDIA
}

// Checks if Payload is Url
func (i *InviteRequest) IsPayloadUrl() bool {
	return i.Payload == Payload_URL
}

// Checks for Flat Invite
func (i *InviteRequest) IsFlat() bool {
	return i.Data.Properties.IsFlat
}

// Checks for Remote Invite
func (i *InviteRequest) IsRemote() bool {
	return i.Data.Properties.IsRemote
}

// Validates InviteRequest has From Parameter
func (u *User) ValidateInvite(i *InviteRequest) *InviteRequest {
	if i.From == nil {
		i.From = u.GetPeer()
	}
	return i
}

// ** ─── REST API MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Proto Request from HTTP Request
func NewRestRequest(r *http.Request) *RestRequest {
	// Method Type
	methodType := RestMethodType_DEFAULT

	// Find Method
	for k := range RestMethodType_value {
		if k == r.Method {
			methodType = RestMethodType(RestMethodType_value[r.Method])
		}
	}

	return &RestRequest{
		Type:   methodType,
		Method: extractHTTPFunction(r.RequestURI),
	}
}

// ** ─── Location MANAGEMENT ────────────────────────────────────────────────────────
func (l *Location) MinorOLC() string {
	lat := l.GetLatitude()
	lon := l.GetLongitude()
	return olc.Encode(lat, lon, 6)
}

func (l *Location) MajorOLC() string {
	lat := l.GetLatitude()
	lon := l.GetLongitude()
	return olc.Encode(lat, lon, 2)
}

func (l *Location) OLC() string {
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), 5)
}

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// @ Local Lobby Topic Protocol ID
func (r *User) LocalTopic() string {
	return fmt.Sprintf("/sonr/topic/%s", r.Location.OLC())
}

// @ LocalTransferProtocol Controller Data Protocol ID
func (r *User_Router) LocalTransferProtocol(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/local-transfer/%s", id.Pretty()))
}

// @ Lobby Topic Protocol ID
func (r *User_Router) LinkProtocol(username string) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/user-linker/%s", username))
}

// @ Transfer Controller Data Protocol ID
func (r *User_Router) RemoteTransfer(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/remote-transfer/%s", id.Pretty()))
}

// @ Lobby Topic Protocol ID
func (r *User_Router) Topic(name string) string {
	return fmt.Sprintf("/sonr/topic/%s", name)
}

// @ Major Rendevouz Advertising Point
func (u *User) GetRouter() *User_Router {
	return u.GetConnection().GetRouter()
}

// ** ─── Status MANAGEMENT ────────────────────────────────────────────────────────
// Update Connected Connection Status
func (u *User) SetConnected(value bool) *StatusUpdate {
	// Update Status
	if value {
		u.Connection.Status = Status_CONNECTED
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetAvailable(value bool) *StatusUpdate {
	// Update Status
	if value {
		u.Connection.Status = Status_AVAILABLE
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Node Status
func (u *User) SetStatus(ns Status) *StatusUpdate {
	// Set Value
	u.Connection.Status = ns

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Checks if Status is Given Value
func (u *User) IsStatus(gs Status) bool {
	return u.GetConnection().GetStatus() == gs
}

// Checks if Status is Not Given Value
func (u *User) IsNotStatus(gs Status) bool {
	return u.GetConnection().GetStatus() != gs
}

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

// ^ Checks for Error With Type ^ //
func NewErrorJoined(errors ...SonrErrorOpt) *SonrError {
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

		// Return Joined Error
		return &SonrError{
			IsJoined: true,
			HasError: true,
			Capture:  capture,
			Joined:   joined,
		}
	} else {
		// Return Error
		return &SonrError{
			HasError: false,
		}
	}
}

// ^ Returns Proto Marshal Error
func NewMarshalError(err error) *SonrError {
	// Return Error
	// Initialize
	message, severity := generateError(ErrorMessage_MARSHAL)

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
			Type:     ErrorMessage_MARSHAL,
			Severity: severity,
		},
		Capture:  capture,
		HasError: true,
	}
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

	// Return Error
	return &SonrError{
		data: &ErrorMessage{
			Message:  message,
			Error:    err.Error(),
			Type:     ErrorMessage_UNMARSHAL,
			Severity: severity,
		},
		Capture:  capture,
		HasError: true,
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
	default:
		return "Unknown", ErrorMessage_LOG
	}
}
