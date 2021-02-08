package file

import (
	"path/filepath"
	"time"

	md "github.com/sonr-io/core/internal/models"
)

// ^ Method Returns Payload by Extension ^ //
func GetPayloadFromPath(p string) md.Payload {
	// Get Extension
	ext := filepath.Ext(p)

	// Cross Check Extension
	if ext == "pdf" {
		return md.Payload_PDF
	} else if ext == "ppt" || ext == "pptx" {
		return md.Payload_PRESENTATION
	} else if ext == "xls" || ext == "xlsm" || ext == "xlsx" || ext == "csv" {
		return md.Payload_SPREADSHEET
	} else if ext == "txt" || ext == "doc" || ext == "docx" || ext == "ttf" {
		return md.Payload_TEXT
	}
	return md.Payload_UNDEFINED
}

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
	// Return Card
	return &md.TransferCard{
		// SQL Properties
		Payload:    GetPayloadFromPath(m.Path),
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

// ^ Method Generates new Transfer Card from URL ^ //
func NewCardFromUrl(p *md.Profile, s string) *md.TransferCard {
	// Return Card
	return &md.TransferCard{
		// SQL Properties
		Payload:  md.Payload_URL,
		Platform: p.Platform,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		// Data Properties
		Url: s,
	}
}
