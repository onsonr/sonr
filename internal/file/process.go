package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"sync"

	md "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnProtobuf func([]byte)
type OnQueued func(card *md.TransferCard, req *md.InviteRequest)
type OnProgress func(data float32)
type OnError func(err error, method string)

// Package Error Callback
var onError OnError

// ******************* //
// ******************* //
// ** OUTGOING FILE ** //
// ******************* //
// ******************* //

// ^ File that safely sets metadata and thumbnail in routine ^ //
type ProcessedFile struct {
	// References
	Payload  md.Payload
	OnQueued OnQueued
	mime     *md.MIME
	Path     string

	// Private Properties
	mutex   sync.Mutex
	card    md.TransferCard
	request *md.InviteRequest
}

// ^ Method adjusts extension for JPEG ^ //
func (pf *ProcessedFile) Ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpeg"
	}
	return pf.mime.Subtype
}

// ^ NewProcessedFile Processes Outgoing File ^ //
func NewProcessedFile(req *md.InviteRequest, p *md.Profile, queueCall OnQueued, errCall OnError) *ProcessedFile {
	// Set Package Level Callbacks
	onError = errCall

	// Get File Information
	file := req.Files[len(req.Files)-1]
	info := GetFileInfo(file.Path)

	// @ 1. Create new SafeFile
	sm := &ProcessedFile{
		OnQueued: queueCall,
		Path:     file.Path,
		Payload:  info.Payload,
		request:  req,
		mime:     info.Mime,
	}

	// ** Lock ** //
	sm.mutex.Lock()

	// @ 2. Set Metadata Protobuf Values

	// Create Card
	sm.card = md.TransferCard{
		// SQL Properties
		Payload:  info.Payload,
		Platform: p.Platform,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		Properties: &md.TransferCard_Properties{
			Name: info.Name,
			Size: info.Size,
			Mime: info.Mime,
		},
	}

	// @ 3. Create Thumbnail in Goroutine
	if len(file.Thumbdata) > 0 {
		go HandleThumbdata(file, sm)
	} else {
		go RequestThumbnail(file, sm)
	}
	return sm
}

// ^ NewBatchProcessFiles Processes Multiple Outgoing Files ^ //
func NewBatchProcessFiles(req *md.InviteRequest, p *md.Profile, queueCall OnQueued, errCall OnError) []*ProcessedFile {
	// Set Package Level Callbacks
	onError = errCall
	files := make([]*ProcessedFile, 64)

	// Iterate Through Attached Files
	for _, file := range req.Files {
		// Get Info
		info := GetFileInfo(file.Path)

		// @ 1. Create new SafeFile
		sm := &ProcessedFile{
			OnQueued: queueCall,
			Path:     file.Path,
			request:  req,
			mime:     info.Mime,
		}

		// ** Lock ** //
		sm.mutex.Lock()

		// @ 2. Set Metadata Protobuf Values
		// Create Card
		sm.card = md.TransferCard{
			// SQL Properties
			Payload:  info.Payload,
			Platform: p.Platform,

			// Owner Properties
			Username:  p.Username,
			FirstName: p.FirstName,
			LastName:  p.LastName,

			Properties: &md.TransferCard_Properties{
				Name: info.Name,
				Size: info.Size,
				Mime: info.Mime,
			},
		}
	}
	return files
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *ProcessedFile) TransferCard() *md.TransferCard {
	// ** Lock File wait for access ** //
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// @ 2. Return Value
	return &sm.card
}

// ^ Method to generate thumbnail for ProcessRequest^ //
func RequestThumbnail(reqFi *md.InviteRequest_FileInfo, sm *ProcessedFile) {
	// Initialize
	thumbBuffer := new(bytes.Buffer)

	// @ 1. Check for External File Request
	if reqFi.HasThumbnail {
		// Encode Thumbnail
		err := EncodeThumb(thumbBuffer, reqFi.Thumbpath)
		if err != nil {
			log.Panicln(err)
		}

		// Update Thumbnail Value
		sm.card.Preview = thumbBuffer.Bytes()

		// @ 2. Handle Created File Request
	} else {
		// Validate Image
		if sm.mime.Type == md.MIME_image {
			// Encode Thumbnail
			err := GenerateThumb(thumbBuffer, reqFi.Path)
			if err != nil {
				log.Panicln(err)
			}

			// Update Thumbnail Value
			sm.card.Preview = thumbBuffer.Bytes()
		}
	}

	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Transfer Card
	preview := sm.TransferCard()

	// @ 3. Callback with Preview
	sm.OnQueued(preview, sm.request)
}

// ^ Method to generate thumbnail for ProcessRequest^ //
func HandleThumbdata(reqFi *md.InviteRequest_FileInfo, sm *ProcessedFile) {
	// Initialize
	thumbWriter := new(bytes.Buffer)
	thumbReader := bytes.NewReader(reqFi.Thumbdata)

	// Convert to Image Object
	img, _, err := image.Decode(thumbReader)
	if err != nil {
		log.Println(err)
	}

	// @ Encode as Jpeg into buffer w/o scaling
	err = jpeg.Encode(thumbWriter, img, nil)
	if err != nil {
		log.Panicln(err)
	}

	sm.card.Preview = thumbWriter.Bytes()

	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Transfer Card
	preview := sm.TransferCard()

	// @ 3. Callback with Preview
	sm.OnQueued(preview, sm.request)
}
