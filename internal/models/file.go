package models

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

// ^ Struct returned for Info about Outgoing File
type OutFileInfo struct {
	Mime    *MIME
	Payload Payload
	Name    string
	Path    string
	Size    int32
	IsImage bool
	Preview []byte
}

// ** Method Returns File Info at Path **
func GetOutFileInfo(path string) (*OutFileInfo, error) {
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
	return &OutFileInfo{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
	}, nil
}

// ** Method Returns File Info at Path **
func GetOutFileInfoWithPreview(path string, preview []byte) (*OutFileInfo, error) {
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
	return &OutFileInfo{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
		Preview: preview,
	}, nil
}

// ** Checks for Preview **
func (fi *OutFileInfo) HasPreview() bool {
	return len(fi.Preview) > 0
}

// ** Method adjusts extension for JPEG **
func (fi *OutFileInfo) Ext() string {
	if fi.Mime.Subtype == "jpg" || fi.Mime.Subtype == "jpeg" {
		return "jpeg"
	}
	return fi.Mime.Subtype
}

// ^ Struct returned for Info about Incoming File
type InFileInfo struct {
	// Inherited Properties
	Properties *TransferCard_Properties
	Preview    []byte

	// Tracking
	CurrentSize int
	Interval    int
	TotalChunks int
	TotalSize   int
}

// ** Method Creates New Transfer File **
func GetInFileInfo(inv *AuthInvite) *InFileInfo {
	// Return File
	return &InFileInfo{
		// Inherited Properties
		Properties: inv.Card.Properties,
		Preview:    inv.Card.Preview,
	}
}
