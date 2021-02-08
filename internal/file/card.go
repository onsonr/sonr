package file

import (
	"time"

	md "github.com/sonr-io/core/internal/models"
)

// ^ Method Generates new Transfer Card from Contact^ //
func NewCardFromContact(p *md.Profile, c *md.Contact) *md.TransferCard {
	return &md.TransferCard{
		// SQL Properties
		Payload:    md.Payload_CONTACT,
		LastOpened: int32(time.Now().Unix()),
		Preview:    p.Picture,
		Platform:   p.Platform,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		// Data Properties
		Contact: c,
	}
}

// ^ Method Generates new Transfer Card from Metadata^ //
func NewCardFromMetadata(p *md.Profile, m *md.Metadata) *md.TransferCard {
	// Get Payload Type
	var load md.Payload

	// Check Metadata MIME
	switch m.Mime.Type {
	case md.MIME_application:
		load = md.Payload_UNDEFINED
	case md.MIME_audio:
		load = md.Payload_AUDIO
	case md.MIME_image:
		load = md.Payload_IMAGE
	case md.MIME_text:
		load = md.Payload_TEXT
	case md.MIME_video:
		load = md.Payload_VIDEO
	default:
		load = md.Payload_FILE
	}

	// Return Card
	return &md.TransferCard{
		// SQL Properties
		Payload:    load,
		LastOpened: int32(time.Now().Unix()),
		Platform:   p.Platform,
		Preview:    m.Thumbnail,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		// Data Properties
		Metadata: m,
	}
}
