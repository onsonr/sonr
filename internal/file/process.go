package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/h2non/filetype"
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
	info, err := GetFileInfo(file.Path)
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

// ^ Struct returned on GetInfo() Generate Preview/Metadata
type FileInfo struct {
	Mime    *md.MIME
	Payload md.Payload
	Name    string
	Path    string
	Size    int32
	IsImage bool
}

// ^ Method Returns File Info at Path ^ //
func GetFileInfo(path string) (*FileInfo, error) {
	// Initialize
	var mime *md.MIME
	var payload md.Payload

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
	mime = &md.MIME{
		Type:    md.MIME_Type(md.MIME_Type_value[kind.MIME.Type]),
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Find Payload
	if mime.Type == md.MIME_image || mime.Type == md.MIME_video || mime.Type == md.MIME_audio {
		payload = md.Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(path)

		// Cross Check Extension
		if ext == ".pdf" {
			payload = md.Payload_PDF
		} else if ext == ".ppt" || ext == ".pptx" {
			payload = md.Payload_PRESENTATION
		} else if ext == ".xls" || ext == ".xlsm" || ext == ".xlsx" || ext == ".csv" {
			payload = md.Payload_SPREADSHEET
		} else if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".ttf" {
			payload = md.Payload_TEXT
		} else {
			payload = md.Payload_UNDEFINED
		}
	}

	// Return Object
	return &FileInfo{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
	}, nil
}
