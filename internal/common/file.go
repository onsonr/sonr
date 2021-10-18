package common

import (
	"os"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sonr-io/core/internal/device"
)

// NewFileItem creates a new transfer file item
func NewFileItem(path string, tbuf []byte) (*Payload_Item, error) {
	// Extracts File Infrom from path
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Create MIME
	mime, err := NewMime(path)
	if err != nil {
		return nil, err
	}

	// Create File Item
	fileItem := &FileItem{
		Mime:         mime,
		Path:         path,
		Size:         fi.Size(),
		Name:         fi.Name(),
		LastModified: fi.ModTime().Unix(),
		Thumbnail: &Thumbnail{
			Buffer: tbuf,
			Mime:   mime,
		},
	}

	// Returns transfer item
	return &Payload_Item{
		Size: fi.Size(),
		Mime: mime,
		Data: &Payload_Item_File{
			File: fileItem,
		},
	}, nil
}

// ResetDir replaces the directory of the item path with DownloadsPath
func (fi *FileItem) ResetDir() string {
	// Set Path
	fi.Path, _ = device.NewDownloadsPath(fi.GetPath())
	return fi.GetPath()
}

// ToTransferItem Returns Transfer for FileItem
func (f *FileItem) ToTransferItem() *Payload_Item {
	return &Payload_Item{
		Mime: f.GetMime(),
		Data: &Payload_Item_File{
			File: f,
		},
	}
}

// ** ───────────────────────────────────────────────────────
// ** ─── MIME Management ───────────────────────────────────
// ** ───────────────────────────────────────────────────────
// DefaultUrlMIME is the standard MIME type for URLs
func DefaultUrlMIME() *MIME {
	return &MIME{
		Type:    MIME_URL,
		Subtype: ".html",
		Value:   "url/html",
	}
}

// NewMime creates a new MIME type from Path
func NewMime(path string) (*MIME, error) {
	// Check if path is URL
	if IsUrl(path) {
		return DefaultUrlMIME(), nil
	}

	// Check if path to file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	// Get MIME Type and Set Proto Enum
	mtype, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}

	// Return MIME Type
	return &MIME{
		Type:    MIME_Type(MIME_Type_value[strings.ToUpper(mtype.Parent().String())]),
		Value:   mtype.String(),
		Subtype: mtype.Extension(),
	}, nil
}

// Ext Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// IsAudio Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// IsImage Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// IsMedia Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.IsAudio() || m.IsImage() || m.IsVideo()
}

// IsPDF Checks if Mime is PDF
func (m *MIME) IsPDF() bool {
	return strings.Contains(m.GetSubtype(), "pdf")
}

// IsVideo Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// IsUrl Checks if Path is a URL
func (m *MIME) IsUrl() bool {
	return m.Type == MIME_URL
}

// PermitsThumbnail Checks if Mime Type Allows Thumbnail Generation.
// Image, Video, Audio, and PDF are allowed.
func (m *MIME) PermitsThumbnail() bool {
	return m.IsImage() || m.IsVideo() || m.IsAudio() || m.IsPDF()
}

// ** ───────────────────────────────────────────────────────
// ** ─── Profile Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// Add adds a new Profile to the List and
// updates LastModified time.
func (p *PayloadList) Add(load *Payload) {
	p.Payloads = append(p.Payloads, load)
	p.LastModified = time.Now().Unix()
}

// Count returns the number of Profiles in the List
func (p *PayloadList) Count() int {
	return len(p.Payloads)
}

// IndexAt returns profile at index
func (p *PayloadList) IndexAt(i int) *Payload {
	return p.Payloads[i]
}
