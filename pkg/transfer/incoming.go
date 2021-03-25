package transfer

import (
	"bytes"
	"encoding/base64"
	"log"
	"strings"
	"sync"
	"time"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"

	"google.golang.org/protobuf/proto"
)

const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3

type IncomingFile struct {
	// Inherited Properties
	mutex      sync.Mutex
	call       dt.NodeCallback
	fs         *sf.FileSystem
	owner      *md.Profile
	payload    md.Payload
	properties *md.TransferCard_Properties
	preview    []byte

	// Builders
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer

	// Tracking
	currentSize int
	interval    int
	totalChunks int
	totalSize   int
}

// ^ Method Creates New Transfer File ^ //
func CreateIncomingFile(inv *md.AuthInvite, fs *sf.FileSystem, tc dt.NodeCallback) *IncomingFile {
	// Return File
	return &IncomingFile{
		// Inherited Properties
		properties: inv.Card.Properties,
		payload:    inv.Payload,
		owner:      inv.From.Profile,
		preview:    inv.Card.Preview,
		fs:         fs,
		call:       tc,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
	}
}

// ^ Check file type and use corresponding method ^ //
func (t *IncomingFile) AddBuffer(curr int, buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk64{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		return true, err
	}

	// @ Initialize Vars if First Chunk
	if curr == 0 {
		// Calculate Tracking Data
		totalChunks := int(chunk.Total) / B64ChunkSize
		interval := totalChunks / 100

		// Set Data
		t.totalSize = int(chunk.Total)
		t.totalChunks = totalChunks
		t.interval = interval
	}

	// @ Add Buffer by File Type
	// Add Base64 Chunk to Buffer
	n, err := t.stringsBuilder.WriteString(chunk.Data)
	if err != nil {
		return true, err
	}

	// Update Tracking
	t.currentSize = t.currentSize + n

	// @ Check Completed
	if t.currentSize < t.totalSize {
		// Validate Interval
		if t.interval != 0 {
			// Check for Interval
			if curr%t.interval == 0 {
				// Send Callback
				t.call.Progressed(float32(t.currentSize) / float32(t.totalSize))
			}
		}
		return false, nil
	} else {
		return true, nil
	}
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (t *IncomingFile) Save() error {
	// Get Bytes from base64
	data, err := base64.StdEncoding.DecodeString(t.stringsBuilder.String())
	if err != nil {
		log.Fatal("error:", err)
	}

	// Write File to Disk
	name, path, err := t.fs.WriteIncomingFile(t.payload, t.properties, data)
	if err != nil {
		return err
	}

	// @ 1. Get File Information
	// Create Card
	card := &md.TransferCard{
		// SQL Properties
		Payload:  t.payload,
		Received: int32(time.Now().Unix()),
		Platform: t.owner.Platform,
		Preview:  t.preview,

		// Transfer Properties
		Status: md.TransferCard_COMPLETED,

		// Owner Properties
		Username:  t.owner.Username,
		FirstName: t.owner.FirstName,
		LastName:  t.owner.LastName,

		// Data Properties
		Metadata: &md.Metadata{
			Name:      name,
			Path:      path,
			Size:      t.properties.Size,
			Mime:      t.properties.Mime,
			Thumbnail: t.preview,
		},
	}

	// Send Complete Callback
	t.call.Received(card)
	return nil
}
