package transfer

import (
	"bytes"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
type IncomingFile struct {
	// Inherited Properties
	mutex  sync.Mutex
	invite *md.AuthInvite
	call   md.TransferCallback
	path   string
	name   string

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
func NewIncomingFile(inv *md.AuthInvite, dirs *md.Directories, tc md.TransferCallback) *IncomingFile {
	// Create File Name
	path, name := getPath(inv.Payload, dirs, inv.Card.Properties)

	// Return File
	return &IncomingFile{
		// Inherited Properties
		invite: inv,
		path:   path,
		call:   tc,
		name:   name,

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
	b64Bytes, err := base64.StdEncoding.DecodeString(t.stringsBuilder.String())
	if err != nil {
		log.Fatal("error:", err)
	}

	// Create File at Path
	f, err := os.Create(t.path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Save Bytes from Base64
	if _, err := f.Write(b64Bytes); err != nil {
		return err
	}

	// Sync file
	if err := f.Sync(); err != nil {
		return err
	}

	// @ 1. Get File Information

	// Get File Information
	p := t.invite.From.Profile

	// Create Card
	card := &md.TransferCard{
		// SQL Properties
		Payload:  t.invite.Payload,
		Received: int32(time.Now().Unix()),
		Platform: p.Platform,
		Preview:  t.invite.Card.Preview,

		// Transfer Properties
		Status: md.TransferCard_COMPLETED,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		// Data Properties
		Metadata: &md.Metadata{
			Name:      t.name,
			Path:      t.path,
			Size:      t.invite.Card.Properties.Size,
			Mime:      t.invite.Card.Properties.Mime,
			Thumbnail: t.invite.Card.Preview,
		},
	}

	// Send Complete Callback
	t.call.Received(card)
	return nil
}

// @ Helper Method to Set Path and FileName - (Path, File Name)
func getPath(load md.Payload, dirs *md.Directories, props *md.TransferCard_Properties) (string, string) {
	// Create File Name
	fileName := props.Name + "." + props.Mime.Subtype

	// Check Load
	if load == md.Payload_MEDIA {
		return filepath.Join(dirs.Temporary, fileName), fileName
	} else {
		return filepath.Join(dirs.Documents, fileName), fileName
	}
}
