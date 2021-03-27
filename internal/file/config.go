package file

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	md "github.com/sonr-io/core/internal/models"
)

// ^ Struct returned on GetInfo() Generate Preview/Metadata
type FileInfo struct {
	Mime    *md.MIME
	Payload md.Payload
	Name    string
	Path    string
	Size    int32
	IsImage bool
}

// ^ Method Returns File Info at Path ^ //
func GetFileInfo(path string) (*FileInfo, error) {
	// Initialize
	var mime *md.MIME
	var payload md.Payload

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		return nil, err
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		return nil, err
	}

	// @ 2. Create Mime Protobuf
	mime = &md.MIME{
		Type:    md.MIME_Type(md.MIME_Type_value[kind.MIME.Type]),
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Find Payload
	if mime.Type == md.MIME_image || mime.Type == md.MIME_video || mime.Type == md.MIME_audio {
		payload = md.Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(path)

		// Cross Check Extension
		if ext == ".pdf" {
			payload = md.Payload_PDF
		} else if ext == ".ppt" || ext == ".pptx" {
			payload = md.Payload_PRESENTATION
		} else if ext == ".xls" || ext == ".xlsm" || ext == ".xlsx" || ext == ".csv" {
			payload = md.Payload_SPREADSHEET
		} else if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".ttf" {
			payload = md.Payload_TEXT
		} else {
			payload = md.Payload_UNDEFINED
		}
	}

	// Return Object
	return &FileInfo{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
	}, nil
}
