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
type OnProtobuf func([]byte)
type OnProgress func(data float32)
type OnError func(err error, method string)

// Package Error Callback
var onError OnError

// ******************* //
// ******************* //
// ** OUTGOING FILE ** //
// ******************* //
// ******************* //

// ^ Outgoing Type for File Media ^
type OutgoingType int

const (
	Unknown OutgoingType = iota
	ExternalMedia
	Image
	Video
)

// ^ File that safely sets metadata and thumbnail in routine ^ //
type ProcessedFile struct {
	// References
	OnQueued OnProtobuf
	outType  OutgoingType
	mime     *md.MIME
	path     string

	// Private Properties
	mutex   sync.Mutex
	preview md.Preview
	request *md.ProcessRequest
}

func (pf *ProcessedFile) Ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpg"
	}
	return pf.mime.Subtype
}

// ^ NewProcessedFile Processes Outgoing File ^ //
func NewProcessedFile(req *md.ProcessRequest, queueCall OnProtobuf, errCall OnError) *ProcessedFile {
	// Set Package Level Callbacks
	onError = errCall

	// Get File Information
	info := GetInfo(req.FilePath)

	// @ 1. Create new SafeFile
	sm := &ProcessedFile{
		OnQueued: queueCall,
		path:     req.FilePath,
		request:  req,
		mime:     info.Mime,
		outType:  GetOutgoingType(req, info.Mime),
	}

	// ** Lock ** //
	sm.mutex.Lock()

	// @ 2. Set Metadata Protobuf Values
	// Create Preview
	sm.preview = md.Preview{
		Name: info.Name,
		Size: info.Size,
		Mime: info.Mime,
	}

	// @ 3. Create Thumbnail in Goroutine
	go RequestThumbnail(req, sm)
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *ProcessedFile) GetPreview() *md.Preview {
	// ** Lock File wait for access ** //
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// @ 2. Return Value
	return &sm.preview
}

// ********************** //
// ********************** //
// ** FILE INFORMATION ** //
// ********************** //
// ********************** //

// ^ Struct returned on GetInfo() Generate Preview/Metadata
type Info struct {
	Mime    *md.MIME
	Name    string
	Path    string
	Size    int32
	IsImage bool
}

// ^ Method Returns File Info at Path ^ //
func GetInfo(path string) Info {
	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
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

	// Check for Image
	var isImage bool
	if filetype.IsImage(head) {
		isImage = true
	}

	return Info{
		Mime: &md.MIME{
			Type:    md.MIME_Type(md.MIME_Type_value[kind.MIME.Type]),
			Subtype: kind.MIME.Subtype,
			Value:   kind.MIME.Value,
		},
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: isImage,
	}
}

// ^ Method Returns Outgoing Type ^ //
func GetOutgoingType(req *md.ProcessRequest, mime *md.MIME) OutgoingType {
	// @ External File
	if req.IsExternal {
		if mime.Type == md.MIME_image || mime.Type == md.MIME_video {
			if req.ThumbnailPath != "" {
				return ExternalMedia
			}
		}

		// @ Passed Internal File
	} else {
		if mime.Type == md.MIME_image {
			return Image
		} else if mime.Type == md.MIME_video {
			return Video
		}
	}
	return Unknown
}
