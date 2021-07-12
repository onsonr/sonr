package models

import (
	"fmt"
	"log"

	"net/http"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	util "github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// ** ─── VerifyRequest MANAGEMENT ────────────────────────────────────────────────────────
// Checks if VerifyRequest is for String Value
func (vr *VerifyRequest) IsString() bool {
	switch vr.Data.(type) {
	case *VerifyRequest_TextValue:
		return true
	default:
		return false
	}
}

// Checks if VerifyRequest is for Buffer Value
func (vr *VerifyRequest) IsBuffer() bool {
	switch vr.Data.(type) {
	case *VerifyRequest_BufferValue:
		return true
	default:
		return false
	}
}

// ** ─── VerifyResponse MANAGEMENT ────────────────────────────────────────────────────────
// Create Marshalled VerifyResponse as GIVEN VALUE
func NewVerifyResponseBuf(result bool) []byte {
	if buf, err := proto.Marshal(&VerifyResponse{IsVerified: result}); err != nil {
		return nil
	} else {
		return buf
	}
}

// Create Marshalled VerifyResponse as TRUE
func NewValidVerifyResponseBuf() []byte {
	if buf, err := proto.Marshal(&VerifyResponse{IsVerified: true}); err != nil {
		return nil
	} else {
		return buf
	}
}

// Create Marshalled VerifyResponse as FALSE
func NewInvalidVerifyResponseBuf() []byte {
	if buf, err := proto.Marshal(&VerifyResponse{IsVerified: false}); err != nil {
		return nil
	} else {
		return buf
	}
}

// ** ─── URLLink MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Link
func NewURLLink(url string) *URLLink {
	link := &URLLink{
		Url:         url,
		Initialized: false,
	}
	link.SetData()
	return link
}

// Sets URLLink Data
func (u *URLLink) SetData() {
	if !u.Initialized {
		// Create Request
		resp, err := http.Get(u.Url)
		if err != nil {
			return
		}

		// Get Info
		info, err := util.GetPageData(resp)
		if err != nil {
			return
		}

		// Set Link
		u.Initialized = true
		u.Title = info.Title
		u.Type = info.Type
		u.Site = info.Site
		u.SiteName = info.SiteName
		u.Description = info.Description
		u.Locale = info.Locale

		// Get Images
		if info.Images != nil {
			for _, v := range info.Images {
				u.Images = append(u.Images, &URLLink_OpenGraphImage{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Width:     int32(v.Width),
					Height:    int32(v.Height),
					Type:      v.Type,
				})
			}
		}

		// Get Videos
		if info.Videos != nil {
			for _, v := range info.Videos {
				u.Videos = append(u.Videos, &URLLink_OpenGraphVideo{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Width:     int32(v.Width),
					Height:    int32(v.Height),
					Type:      v.Type,
				})
			}
		}

		// Get Audios
		if info.Audios != nil {
			for _, v := range info.Videos {
				u.Audios = append(u.Audios, &URLLink_OpenGraphAudio{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Type:      v.Type,
				})
			}
		}

		// Get Twitter
		if info.Twitter != nil {
			u.Twitter = &URLLink_TwitterCard{
				Card:        info.Twitter.Card,
				Site:        info.Twitter.Site,
				SiteId:      info.Twitter.SiteId,
				Creator:     info.Twitter.Creator,
				CreatorId:   info.Twitter.CreatorId,
				Description: info.Twitter.Description,
				Title:       info.Twitter.Title,
				Image:       info.Twitter.Image,
				ImageAlt:    info.Twitter.ImageAlt,
				Url:         info.Twitter.Url,
				Player: &URLLink_TwitterCard_Player{
					Url:    info.Twitter.Player.Url,
					Width:  int32(info.Twitter.Player.Width),
					Height: int32(info.Twitter.Player.Height),
					Stream: info.Twitter.Player.Stream,
				},
				Iphone: &URLLink_TwitterCard_IPhone{
					Name: info.Twitter.IPhone.Name,
					Id:   info.Twitter.IPhone.Id,
					Url:  info.Twitter.IPhone.Url,
				},
				Ipad: &URLLink_TwitterCard_IPad{
					Name: info.Twitter.IPad.Name,
					Id:   info.Twitter.IPad.Id,
					Url:  info.Twitter.IPad.Url,
				},
				GooglePlay: &URLLink_TwitterCard_GooglePlay{
					Name: info.Twitter.Googleplay.Name,
					Id:   info.Twitter.Googleplay.Id,
					Url:  info.Twitter.Googleplay.Url,
				},
			}
		}
	}
}

// ** ─── InviteResponse MANAGEMENT ────────────────────────────────────────────────────────
// Checks if Peer Accepted Transfer
func (r *InviteResponse) HasAcceptedTransfer() bool {
	return r.GetDecision() && r.GetType() == InviteResponse_Default
}

// Returns Protocol ID Set by Peer
func (r *InviteResponse) ProtocolID() protocol.ID {
	return protocol.ID(r.GetProtocol())
}

// ** ─── InviteRequest MANAGEMENT ────────────────────────────────────────────────────────
// Returns Invite Contact
func (i *InviteRequest) GetContact() *Contact {
	return i.GetTransfer().GetContact()
}

// Returns Invite File
func (i *InviteRequest) GetFile() *SFile {
	return i.GetTransfer().GetFile()
}

// Returns Invite URL
func (i *InviteRequest) GetUrl() *URLLink {
	return i.GetTransfer().GetUrl()
}

// Checks if Payload is Contact
func (i *InviteRequest) IsPayloadContact() bool {
	return i.Payload == Payload_CONTACT
}

// Checks if Payload is File Transfer
func (i *InviteRequest) IsPayloadTransfer() bool {
	return i.Payload == Payload_FILE || i.Payload == Payload_FILES || i.Payload == Payload_MEDIA || i.Payload == Payload_ALBUM
}

// Checks if Payload is Url
func (i *InviteRequest) IsPayloadUrl() bool {
	return i.Payload == Payload_URL
}

// Checks for Flat Invite
func (i *InviteRequest) IsFlatInvite() bool {
	return i.GetType() == InviteRequest_Flat
}

// Returns Protocol ID Set by Peer
func (r *InviteRequest) ProtocolID() protocol.ID {
	return protocol.ID(r.GetProtocol())
}

// Set Protocol for Invite and Return ID
func (i *InviteRequest) SetProtocol(p SonrProtocol, id peer.ID) protocol.ID {
	protocolName := fmt.Sprintf("/sonr/%s/%s", p.Method(), id.String())
	i.Protocol = protocolName
	return protocol.ID(protocolName)
}

// Validates InviteRequest has From Parameter
func (u *User) ValidateInvite(i *InviteRequest) *InviteRequest {
	// Set From
	if i.From == nil {
		i.From = u.GetPeer()
	}
	return i
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
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), 8)
}

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// Construct New Protocol ID given Method Name String and id Peer.ID
func (p SonrProtocol) NewIDProtocol(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", p.Method(), id.String()))
}

// Construct New Protocol ID given Method Name String and Value String
func (p SonrProtocol) NewValueProtocol(value string) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", p.Method(), value))
}

// Returns Method Name for this Protocol
func (p SonrProtocol) Method() string {
	switch p {
	case SonrProtocol_Authorize:
		return "auth-service"
	case SonrProtocol_Devices:
		return "user-devices"
	case SonrProtocol_Linker:
		return "user-linker"
	case SonrProtocol_LocalTransfer:
		return "local-transfer"
	case SonrProtocol_RemoteTransfer:
		return "remote-transfer"
	default:
		return ""
	}
}

// ** ─── Status MANAGEMENT ────────────────────────────────────────────────────────
// Update Connected Connection Status
func (u *User) SetConnected(value bool) *StatusEvent {
	// Update Status
	if value {
		u.Status = Status_CONNECTED
	} else {
		u.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusEvent{Value: u.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetAvailable(value bool) *StatusEvent {
	// Update Status
	if value {
		u.Status = Status_AVAILABLE
	} else {
		u.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusEvent{Value: u.GetStatus()}
}

// Update Node Status
func (u *User) SetStatus(ns Status) *StatusEvent {
	// Set Value
	u.Status = ns

	// Returns Status Update
	return &StatusEvent{Value: u.GetStatus()}
}

// Checks if Status is Given Value
func (u *User) IsStatus(gs Status) bool {
	return u.GetStatus() == gs
}

// Checks if Status is Not Given Value
func (u *User) IsNotStatus(gs Status) bool {
	return u.GetStatus() != gs
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

// ^ Return New Peer Not Found Error with Peer ID as Data ^ //
func NewPeerFoundError(err error, peer string) *SonrError {
	// Initialize
	message, severity := generateError(ErrorMessage_PEER_NOT_FOUND_INVITE)

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
			Data:     peer,
		},
		Capture:  capture,
		HasError: true,
	}
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

// @ Method Prints Error
func (errWrap *SonrError) Print() {
	log.Printf("ERROR: %s", errWrap.String())
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
	case ErrorMessage_HOST_TEXTILE:
		return "Failed to Start Textile Client", ErrorMessage_CRITICAL
	default:
		return "Unknown", ErrorMessage_LOG
	}
}
