package models

import (
	"errors"
	"os"

	"github.com/h2non/filetype"
)

// ^ Finds Mime for Single File ^ //
func (f *SonrFile) FindMime() (*MIME, error) {
	if !f.IsMultiple {
		// Get Single File
		m := f.GetSingleFile()

		// @ 1. Get File Information
		// Open File at Path
		file, err := os.Open(m.Path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

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

		// @ 1. Set Mime for Metadata
		m.Mime = &MIME{
			Type:    MIME_Type(MIME_Type_value[kind.MIME.Type]),
			Subtype: kind.MIME.Subtype,
			Value:   kind.MIME.Value,
		}

		// @ 2. Create Mime Protobuf
		return m.Mime, nil
	} else {
		return nil, errors.New("Too many files to Find Mime")
	}
}

// *********************************** //
// ** Outgoing File Info Management ** //
// *********************************** //
// ^ Method Returns File Size at Path ^ //
func (f *SonrFile) FindSize() (int32, error) {
	if !f.IsMultiple {
		// Get Single File
		m := f.GetSingleFile()

		// Open File at Path
		file, err := os.Open(m.Path)
		if err != nil {
			return 0, err
		}

		// Find Info
		defer file.Close()
		info, err := file.Stat()
		if err != nil {
			return 0, err
		}

		// Set Size
		m.Size = int32(info.Size())
		return m.Size, nil
	} else {
		return 0, errors.New("Too many files to Find Mime")
	}
}

func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// *********************************** //
// ** Incoming File Info Management ** //
// *********************************** //
type InFile struct {
	Payload       Payload
	Metadata      *SonrFile
	ChunkBaseChan chan Chunk64
	ChunkBufChan  chan ChunkBuffer
}
