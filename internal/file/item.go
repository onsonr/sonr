package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"

	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

// @ File that safely sets metadata and thumbnail in routine
type FileItem struct {
	// References
	Payload md.Payload
	call    dt.NodeCallback
	mime    *md.MIME
	Path    string

	// Private Properties
	card    md.TransferCard
	request *md.InviteRequest
}

// ^ NewOutgoingFileItem Processes Outgoing File ^ //
func NewOutgoingFileItem(req *md.InviteRequest, p *md.Peer, callback dt.NodeCallback) *FileItem {
	// Check Values
	if req == nil || p == nil {
		return nil
	}

	// Get File Information
	file := req.Files[len(req.Files)-1]
	info, err := GetFileInfo(file.Path)
	if err != nil {
		callback.Error(err, "NewProcessedFile:GetFileInfo")
		return nil
	}

	// @ 1. Create new SafeFile
	sm := &FileItem{
		call:    callback,
		Path:    file.Path,
		Payload: info.Payload,
		request: req,
		mime:    info.Mime,
	}

	// @ 2. Set Metadata Protobuf Values
	// Create Card
	sm.card = md.TransferCard{
		// SQL Properties
		Payload:  info.Payload,
		Platform: p.Platform,

		// Owner Properties
		Username:  p.Profile.Username,
		FirstName: p.Profile.FirstName,
		LastName:  p.Profile.LastName,

		Properties: &md.TransferCard_Properties{
			Name: info.Name,
			Size: info.Size,
			Mime: info.Mime,
		},
	}

	// @ 3. Create Thumbnail in Goroutine
	if len(file.Thumbnail) > 0 {
		// Initialize
		thumbWriter := new(bytes.Buffer)
		thumbReader := bytes.NewReader(file.Thumbnail)

		// Convert to Image Object
		img, _, err := image.Decode(thumbReader)
		if err != nil {
			log.Println(err)
			return nil
		}

		// @ Encode as Jpeg into buffer w/o scaling
		err = jpeg.Encode(thumbWriter, img, nil)
		if err != nil {
			log.Panicln(err)
			return nil
		}

		sm.card.Preview = thumbWriter.Bytes()
	}

	// Get Transfer Card
	preview := sm.Card()

	// @ 3. Callback with Preview
	sm.call.Queued(preview, sm.request)
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *FileItem) Card() *md.TransferCard {
	// @ 2. Return Value
	return &sm.card
}

// ^ Method adjusts extension for JPEG ^ //
func (pf *FileItem) Ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpeg"
	}
	return pf.mime.Subtype
}
