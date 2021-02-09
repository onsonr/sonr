package file

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

const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
type TransferFile struct {
	// Inherited Properties
	mutex      sync.Mutex
	invite     *md.AuthInvite
	onProgress OnProgress
	onComplete OnProtobuf
	path       string

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
func NewTransfer(inv *md.AuthInvite, dirs *md.Directories, op func(data float32), oc func([]byte)) *TransferFile {
	// Create File Name
	fileName := inv.Card.Properties.Name + "." + inv.Card.Properties.Mime.Subtype

	// Return File
	return &TransferFile{
		// Inherited Properties
		invite:     inv,
		path:       filepath.Join(dirs.Temporary, fileName),
		onProgress: op,
		onComplete: oc,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
	}
}

// ^ Check file type and use corresponding method ^ //
func (t *TransferFile) AddBuffer(curr int, buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk{}
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
	n, err := t.stringsBuilder.WriteString(chunk.B64)
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
				t.onProgress(float32(t.currentSize) / float32(t.totalSize))
			}
		}
		return false, nil
	} else {
		return true, nil
	}
}

// ^ Check file type and use corresponding method to save to Disk ^ //
func (t *TransferFile) Save() error {
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
	// Create Card
	card := t.GetCard()

	// Convert Message to bytes
	bytes, err := proto.Marshal(card)
	if err != nil {
		return err
	}

	// Send Complete Callback
	t.onComplete(bytes)
	return nil
}

// ^ Method Generates new Transfer Card from TransferFile^ //
func (t *TransferFile) GetCard() *md.TransferCard {
	// Get File Information
	i := GetFileInfo(t.path)
	p := t.invite.From.Profile

	// Return Card
	return &md.TransferCard{
		// SQL Properties
		Payload:  i.Payload,
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
			Name:      i.Name,
			Path:      t.path,
			Size:      i.Size,
			Mime:      i.Mime,
			Thumbnail: t.invite.Card.Preview,
		},
	}
}
