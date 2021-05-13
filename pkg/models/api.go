package models

import (
	"errors"

	olc "github.com/google/open-location-code/go"
)

// ** ─── ConnectionRequest MANAGEMENT ────────────────────────────────────────────────────────
func (req *ConnectionRequest) NewRouter() *Router {
	return &Router{
		Location:     req.GetLocation(),
		Connectivity: req.GetConnectivity(),
	}
}

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
