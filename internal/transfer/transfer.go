package transfer

import (
	"bytes"
	"encoding/base64"
	"log"
	"os"
	"strings"
	"sync"

	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type OnProgress func(data float32)
type Transfer struct {
	// Inherited Properties
	mutex      sync.Mutex
	preview    *md.Preview
	owner      *md.Peer
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

// ^ Check file type and use corresponding method ^ //
func (t *Transfer) addBuffer(curr int, buffer []byte) (bool, error) {
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
func (t *Transfer) save() error {
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

	// Create Metadata
	meta := sf.GetMetadata(t.path, t.owner)

	// Convert Message to bytes
	bytes, err := proto.Marshal(meta)
	if err != nil {
		return err
	}

	// Send Complete Callback
	t.onComplete(bytes)
	return nil
}
