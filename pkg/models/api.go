package models

import (
	"fmt"

	olc "github.com/google/open-location-code/go"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── AuthInvite MANAGEMENT ────────────────────────────────────────────────────────
func (r *AuthReply) HasAcceptedTransfer() bool {
	return r.Decision && r.Type == AuthReply_Transfer
}

// ** ─── AuthInvite MANAGEMENT ────────────────────────────────────────────────────────
func (i *AuthInvite) GetContact() *Contact {
	return i.GetData().GetContact()
}

func (i *AuthInvite) GetFile() *SonrFile {
	return i.GetData().GetFile()
}

func (i *AuthInvite) GetUrl() *URLLink {
	return i.GetData().GetUrl()
}

func (i *AuthInvite) IsFlat() bool {
	return i.Data.Properties.IsFlat
}

func (i *AuthInvite) IsRemote() bool {
	return i.Data.Properties.IsRemote
}

// ** ─── Remote MANAGEMENT ────────────────────────────────────────────────────────
// Get Remote Topic
func (r *RemoteCreateRequest) GetTopic() string {
	return fmt.Sprintf("%s.remote.%s.snr/", r.Fingerprint, r.SName)
}

// Get Remote Point Info
func (r *RemoteCreateRequest) NewCreatedRemote(u *User) *Lobby {
	// Create Lobby
	return &Lobby{
		// General
		Type:  Lobby_REMOTE,
		Peers: make(map[string]*Peer),
		User:  u.GetPeer(),

		// Info
		Info: &Lobby_Remote{
			Remote: &Lobby_RemoteInfo{
				IsJoin:      false,
				Fingerprint: r.GetFingerprint(),
				Words:       r.GetWords(),
				Topic:       r.GetTopic(),
				File:        r.GetFile(),
				Owner:       u.GetPeer(),
			},
		},
	}
}

func (r *RemoteJoinRequest) NewJoinedRemote(u *User) *Lobby {
	// Create Lobby
	return &Lobby{
		// General
		Type:  Lobby_REMOTE,
		Peers: make(map[string]*Peer),
		User:  u.GetPeer(),

		// Info
		Info: &Lobby_Remote{
			Remote: &Lobby_RemoteInfo{
				IsJoin: true,
				Topic:  r.GetTopic(),
			},
		},
	}
}

// ** ─── Linker MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Linker from Link Request
func NewLinker(lr *LinkRequest) *Linker {
	return &Linker{
		Device:   lr.Device,
		Username: lr.Username,
		Router: &Linker_Router{
			LocalIPTopic: lr.Location.OLC(),
			Rendevouz:    fmt.Sprintf("/sonr/%s", lr.GetLocation().MajorOLC()),
			Location:     lr.GetLocation(),
		},
	}
}

// Creates New Peer for Linker
func (l *Linker) NewPeer(id peer.ID, maddr multiaddr.Multiaddr) {
	// Initialize
	deviceID := l.Device.GetId()

	// Set Peer
	l.Peer = &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
		},
		Platform: l.Device.Platform,
		Model:    l.Device.Model,
	}

	// Set Device Topic
	l.Router.DeviceTopic = fmt.Sprintf("/sonr/user/%s", l.Username)
}

// Returns Linker Private Key
func (l *Linker) PrivateKey() crypto.PrivKey {
	// Get Key from Buffer
	key, err := crypto.UnmarshalPrivateKey(l.GetDevice().GetPrivateKey())
	if err != nil {
		return nil
	}
	return key
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
	// Set Value
	u.Connection.HasConnected = value

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
func (u *User) SetBootstrapped(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasBootstrapped = value

	// Update Status
	if value {
		u.Connection.Status = Status_BOOTSTRAPPED
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetJoinedLocal(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasJoinedLocal = value

	// Update Status
	if value {
		u.Connection.Status = Status_AVAILABLE
	} else {
		u.Connection.Status = Status_BOOTSTRAPPED
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
	default:
		return "Unknown", ErrorMessage_LOG
	}
}
