package file

import (
	"sync"

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

// ^ File that safely sets metadata and thumbnail in routine ^ //
type ProcessedFile struct {
	// References
	OnQueued OnProtobuf
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
func NewProcessedFile(req *md.ProcessRequest, p *md.Profile, queueCall OnProtobuf, errCall OnError) *ProcessedFile {
	// Set Package Level Callbacks
	onError = errCall

	// Get File Information
	info := GetFileInfo(req.FilePath)

	// @ 1. Create new SafeFile
	sm := &ProcessedFile{
		OnQueued: queueCall,
		path:     req.FilePath,
		request:  req,
		mime:     info.Mime,
	}

	// ** Lock ** //
	sm.mutex.Lock()

	// @ 2. Set Metadata Protobuf Values
	// Create Card
	sm.card = NewCardFromProcessRequest(p, req.FilePath)

	// @ 3. Create Thumbnail in Goroutine
	go RequestThumbnail(req, sm)
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *ProcessedFile) GetTransferCard() *md.TransferCard {
	// ** Lock File wait for access ** //
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// @ 2. Return Value
	return &sm.card
}
