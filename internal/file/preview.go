package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/h2non/filetype"
	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnProtobuf func(data []byte)
type OnError func(err error, method string)

// Package Error Callback
var onError OnError

// ^ File that safely sets metadata and thumbnail in routine ^ //
type SafePreview struct {
	// References
	OnQueued OnProtobuf
	Path     string
	Type     md.MIME_Type

	// Private Properties
	mutex   sync.Mutex
	preview md.Preview
	request *md.ProcessRequest
}

// ^ Create generates file metadata ^ //
func NewPreview(req *md.ProcessRequest, queueCall OnProtobuf, errCall OnError) *SafePreview {
	// Set Package Level Callbacks
	onError = errCall

	// Create new SafeFile
	sm := &SafePreview{
		OnQueued: queueCall,
		Path:     req.FilePath,
		request:  req,
	}

	// ** Lock ** //
	sm.mutex.Lock()

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(sm.Path)
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		log.Fatalln(err)
		onError(err, "AddFile")
	}

	// @ 2. Set Metadata Protobuf Values
	// Set Metadata
	base := filepath.Base(sm.Path)
	sm.Type = md.MIME_Type(md.MIME_Type_value[kind.MIME.Type])

	// Create Preview
	sm.preview = md.Preview{
		Name: strings.TrimSuffix(base, filepath.Ext(sm.Path)),
		Size: int32(info.Size()),
		Mime: &md.MIME{
			Type:      md.MIME_Type(md.MIME_Type_value[kind.MIME.Type]),
			Extension: filepath.Ext(sm.Path),
			Subtype:   kind.MIME.Subtype,
			Value:     kind.MIME.Value,
		},
	}

	// @ 3. Create Thumbnail in Goroutine
	go NewThumbnail(req, sm)
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *SafePreview) GetPreview() *md.Preview {
	// ** Lock File wait for access ** //
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// @ 2. Return Value
	return &sm.preview
}
