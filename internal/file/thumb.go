package file

import (
	"bytes"
	"log"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Method to generate thumbnail for ProcessRequest^ //
func RequestThumbnail(req *md.ProcessRequest, sm *ProcessedFile) {
	// Initialize
	thumbBuffer := new(bytes.Buffer)

	// @ 1. Check for External File Request
	if req.ThumbnailPath != "" {
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
	preview := sm.GetPreview()

	// Convert to bytes
	data, err := proto.Marshal(preview)
	if err != nil {
		log.Println("Error Marshaling Metadata ", err)
	}

	// @ 3. Callback with Preview
	sm.OnQueued(data)
}
