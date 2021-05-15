package models

import (
	"errors"
	"fmt"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

// ************************** //
// ** MIME Info Management ** //
// ************************** //
// Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// ** ─── InviteRequest MANAGEMENT ────────────────────────────────────────────────────────
func (r *InviteRequest) GetContact() *Contact {
	return r.GetData().GetContact()
}

func (r *InviteRequest) GetFile() *SonrFile {
	return r.GetData().GetFile()
}

func (r *InviteRequest) GetUrl() *URLLink {
	return r.GetData().GetUrl()
}

// Prepare for Outgoing Session
func (r *InviteRequest) NewSession(u *User, tc NodeCallback) *Session {
	return &Session{
		file: r.GetFile(),
		peer: r.GetTo(),
		user: u,
		call: tc,
	}
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

// Prepare for Incoming Session
func (i *AuthInvite) NewSession(u *User, c NodeCallback) *Session {
	return &Session{
		file: i.GetFile(),
		peer: i.GetFrom(),
		user: u,
		call: c,
	}
}

// ** ─── Location MANAGEMENT ────────────────────────────────────────────────────────
func (l *Location) MinorOLC() string {
	lat := l.Latitude()
	lon := l.Longitude()
	return olc.Encode(lat, lon, 6)
}

func (l *Location) MajorOLC() string {
	lat := l.Latitude()
	lon := l.Longitude()
	return olc.Encode(lat, lon, 4)
}

func (l *Location) Latitude() float64 {
	if l.Geo != nil {
		return l.Geo.GetLatitude()
	}
	return l.Ip.GetLatitude()
}

func (l *Location) Longitude() float64 {
	if l.Geo != nil {
		return l.Geo.GetLongitude()
	}
	return l.Ip.GetLongitude()
}

func (l *Location) GeoOLC() (string, error) {
	if l.Geo != nil {
		return "", errors.New("Geo Location doesnt exist")
	}
	return olc.Encode(float64(l.Geo.GetLatitude()), float64(l.Geo.GetLongitude()), 5), nil
}

func (l *Location) IPOLC() string {
	return olc.Encode(float64(l.Ip.GetLatitude()), float64(l.Ip.GetLongitude()), 5)
}

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// @ Local Lobby Topic Protocol ID
func (r *User) LocalIPTopic() string {
	return fmt.Sprintf("/sonr/topic/%s", r.Location.IPOLC())
}

func (r *User) LocalGeoTopic() (string, error) {
	geoOlc, err := r.Location.GeoOLC()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/sonr/topic/%s", geoOlc), nil
}

// @ Transfer Controller Data Protocol ID
func (r *User_Router) Transfer(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/transfer/%s", id.Pretty()))
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
