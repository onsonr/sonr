package models

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

// ^ Struct returned on GetInfo() Generate Preview/Metadata
type FileInfo struct {
	Mime      *MIME
	Payload   Payload
	Name      string
	Path      string
	Size      int32
	IsImage   bool
	Thumbnail []byte
}

// ^ Method Returns File Info at Path ^ //
func GetFileInfo(path string, preview []byte) (*FileInfo, error) {
	// Initialize
	var mime *MIME
	var payload Payload

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
	mime = &MIME{
		Type:    MIME_Type(MIME_Type_value[kind.MIME.Type]),
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Find Payload
	if mime.Type == MIME_image || mime.Type == MIME_video || mime.Type == MIME_audio {
		payload = Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(path)

		// Cross Check Extension
		if ext == ".pdf" {
			payload = Payload_PDF
		} else if ext == ".ppt" || ext == ".pptx" {
			payload = Payload_PRESENTATION
		} else if ext == ".xls" || ext == ".xlsm" || ext == ".xlsx" || ext == ".csv" {
			payload = Payload_SPREADSHEET
		} else if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".ttf" {
			payload = Payload_TEXT
		} else {
			payload = Payload_UNDEFINED
		}
	}

	// Return Object
	return &FileInfo{
		Mime:      mime,
		Payload:   payload,
		Name:      strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:      path,
		Size:      int32(info.Size()),
		IsImage:   filetype.IsImage(head),
		Thumbnail: preview,
	}, nil
}
