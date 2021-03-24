package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"sync"

	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
)

// @ File that safely sets metadata and thumbnail in routine
type ProcessedFile struct {
	// References
	Payload md.Payload
	call    dt.NodeCallback
	mime    *md.MIME
	Path    string

	// Private Properties
	mutex   sync.Mutex
	card    md.TransferCard
	request *md.InviteRequest
}

// @ ProcessedFileBuilder creates a new item and returns a pointer to it.
func ProcessedFileBuilder() interface{} {
	return &ProcessedFile{}
}

// ^ NewProcessedFile Processes Outgoing File ^ //
func NewProcessedFile(req *md.InviteRequest, p *md.Profile, callback dt.NodeCallback) *ProcessedFile {
	// Check Values
	if req == nil || p == nil {
		return nil
	}

	// Get File Information
	file := req.Files[len(req.Files)-1]
	info, err := dt.GetFileInfo(file.Path)
	if err != nil {
		callback.Error(err, "NewProcessedFile:GetFileInfo")
	}

	// @ 1. Create new SafeFile
	sm := &ProcessedFile{
		call:    callback,
		Path:    file.Path,
		Payload: info.Payload,
		request: req,
		mime:    info.Mime,
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
	if len(file.Thumbnail) > 0 {
		go HandleThumbnail(file, sm)
	} else {
		go RequestThumbnail(file, sm)
	}
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *ProcessedFile) Card() *md.TransferCard {
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
	// @ Handle Created File Request
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

	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Transfer Card
	preview := sm.Card()

	// @ 3. Callback with Preview
	sm.call.Queued(preview, sm.request)
}

// ^ Method to Handle Provided Thumbnail ^ //
func HandleThumbnail(reqFi *md.InviteRequest_FileInfo, sm *ProcessedFile) {
	// Initialize
	thumbWriter := new(bytes.Buffer)
	thumbReader := bytes.NewReader(reqFi.Thumbnail)

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
	preview := sm.Card()

	// @ 3. Callback with Preview
	sm.call.Queued(preview, sm.request)
}
