package models

import (
	"fmt"

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

// ** ─── InitializeRequest MANAGEMENT ────────────────────────────────────────────────────────
// Check for Logging Enabled
func (i *InitializeRequest) IsLoggingEnabled() bool {
	if i.GetLogLevel() != InitializeRequest_NONE {
		return true
	}
	return false
}

// Check if Info Log Enabled
func (i *InitializeRequest) HasInfoLog() bool {
	if i.GetLogLevel() >= InitializeRequest_INFO {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasDebugLog() bool {
	if i.GetLogLevel() >= InitializeRequest_DEBUG {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasWarningLog() bool {
	if i.GetLogLevel() >= InitializeRequest_WARNING {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasCriticalLog() bool {
	if i.GetLogLevel() >= InitializeRequest_CRITICAL {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasFatalLog() bool {
	if i.GetLogLevel() >= InitializeRequest_FATAL {
		return true
	}
	return false
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
	return r.GetDecision() && r.GetType() == InviteResponse_Transfer
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
	// Initialize
	protocolName := fmt.Sprintf("/sonr/%s/%s", p.Method(), id.String())

	// // Get Nano ID
	// nanoid, err := gonanoid.Generate(id.Pretty(), 24)
	// if err == nil {
	// 	protocolName = fmt.Sprintf("/sonr/%s/%s", p.Method(), nanoid)
	// }

	// Set Name and Return ID
	i.Protocol = protocolName
	return protocol.ID(protocolName)
}

// Validates InviteRequest has From Parameter
func (u *User) SignInvite(i *InviteRequest) *InviteRequest {
	// Set From
	if i.From == nil {
		i.From = u.GetPeer()
	}

	// Set Type
	if i.Type == InviteRequest_None {
		i.Type = InviteRequest_Local
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

// ** ─── Generic Callback MANAGEMENT ───────────────────────────────────────────

// Create New Generic CONNECTION Response Message
func (o *ConnectionResponse) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Response{
		Type: Response_CONNECTION,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic INVITE Request Message
func (o *InviteRequest) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Request{
		Type: Request_INVITE,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic INVITE Response Message
func (o *InviteResponse) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Response{
		Type: Response_INVITE,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic REST Request Message
func (o *RestRequest) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Request{
		Type: Request_REST,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic REST Response Message
func (o *RestResponse) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Response{
		Type: Response_REST,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic VERIFY Request Message
func (o *VerifyRequest) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Request{
		Type: Request_VERIFY,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic COMPLETE Event Message
func (o *CompleteEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Event{
		Type: Event_COMPLETE,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic MAIL Event Message
func (o *MailEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Event{
		Type: Event_MAIL,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic PROGRESS Event Message
func (o *ProgressEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Event{
		Type: Event_PROGRESS,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic TOPIC Event Message
func (o *TopicEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &Event{
		Type: Event_TOPIC,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}
