package file

import (
	"bytes"
	"log"
	"sync"

	lf "github.com/sonr-io/core/internal/lifecycle"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// Package Error Callback
var onError lf.OnError

// ******************* //
// ******************* //
// ** OUTGOING FILE ** //
// ******************* //
// ******************* //

// ^ File that safely sets metadata and thumbnail in routine ^ //
type ProcessedFile struct {
	// References
	OnQueued lf.OnProtobuf
	mime     *md.MIME
	path     string

	// Private Properties
	mutex   sync.Mutex
	card    md.TransferCard
	request *md.ProcessRequest
}

func (pf *ProcessedFile) Ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpg"
	}
	return pf.mime.Subtype
}

// ^ NewProcessedFile Processes Outgoing File ^ //
func NewProcessedFile(req *md.ProcessRequest, p *md.Profile, calls lf.ProcessCallbacks) *ProcessedFile {
	// Set Package Level Callbacks
	onError = calls.CallError

	// Get File Information
	info := GetFileInfo(req.FilePath)

	// @ 1. Create new SafeFile
	sm := &ProcessedFile{
		OnQueued: calls.CallQueued,
		path:     req.FilePath,
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
	go RequestThumbnail(req, sm)
	return sm
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
func RequestThumbnail(req *md.ProcessRequest, sm *ProcessedFile) {
	// Initialize
	thumbBuffer := new(bytes.Buffer)

	// @ 1. Check for External File Request
	if req.HasThumbnail {
		// Encode Thumbnail
		err := EncodeThumb(thumbBuffer, req.ThumbnailPath)
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
			err := GenerateThumb(thumbBuffer, req.FilePath)
			if err != nil {
				log.Panicln(err)
			}

			// Update Thumbnail Value
			sm.card.Preview = thumbBuffer.Bytes()
		}
	}

	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Metadata
	preview := sm.TransferCard()
	preview.Status = md.TransferCard_PROCESSED

	// Convert to bytes
	data, err := proto.Marshal(preview)
	if err != nil {
		log.Println("Error Marshaling Metadata ", err)
	}

	// @ 3. Callback with Preview
	sm.OnQueued(data)
}
