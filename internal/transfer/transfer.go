package transfer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type OnProgress func(p float32)

type Transfer struct {
	// Inherited Properties
	mutex      sync.Mutex
	metadata   *md.Metadata
	owner      *md.Peer
	onProgress OnProgress
	onComplete OnProtobuf
	path       string

	// Builders
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer

	// Tracking
	count int
	size  int
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewTransfer(docDir string, meta *md.Metadata, own *md.Peer, op OnProgress, oc OnProtobuf) *Transfer {
	return &Transfer{
		// Inherited Properties
		metadata:   meta,
		path:       docDir + "/" + meta.Name + "." + meta.Mime.Subtype,
		owner:      own,
		onProgress: op,
		onComplete: oc,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),

		// Tracking
		count: 0,
		size:  0,
	}
}

// ^ Check file type and use corresponding method ^ //
func (t *Transfer) AddBuffer(buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		return true, err
	}

	// @ Increment received count
	if t.count == 0 {
		t.size = int(chunk.Total)
	}

	// Set Tracking
	var n int
	t.count++

	// @ Check File Type
	if chunk.GetB64() != "" {
		// Add Base64 Chunk to Buffer
		n, err = t.stringsBuilder.WriteString(chunk.B64)
		if err != nil {
			return true, err
		}
	} else {
		// Add ByteChunk to Buffer
		n, err = t.bytesBuilder.Write(chunk.Buffer)
		if err != nil {
			return true, err
		}
	}

	// @ Update Tracking
	currW := t.count*BufferChunkSize + n
	currP := float32(currW) / float32(t.size)
	t.onProgress(currP)

	if t.stringsBuilder.Len() == t.size || t.bytesBuilder.Len() == t.size {
		t.Save()
		return true, nil
	}
	return false, nil
}

// ^ Check file type and use corresponding method to Save to Disk ^ //
func (t *Transfer) Save() {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// @ Create File at Path
	f, err := os.Create(t.path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// @ Set File Bytes by Type
	if t.metadata.Mime.Type == md.MIME_image {
		// Get Base64 Data
		data := t.stringsBuilder.String()

		// Get Bytes from base64
		b64Bytes, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Fatal("error:", err)
		}

		// Save Bytes from Base64
		if _, err := f.Write(b64Bytes); err != nil {
			log.Fatalln(err)
		}

		// Sync file
		if err := f.Sync(); err != nil {
			log.Fatalln(err)
		}
	} else {
		// Save Bytes from Buffer
		if _, err := f.Write(t.bytesBuilder.Bytes()); err != nil {
			log.Fatalln(err)
		}

		// Sync file
		if err := f.Sync(); err != nil {
			log.Fatalln(err)
		}
	}

	// Get Info
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// @ 3. Callback saved Metadata
	// Create Metadata
	saved := &md.Metadata{
		Name:       t.metadata.Name,
		Path:       t.path,
		Size:       int32(info.Size()),
		Mime:       t.metadata.Mime,
		Owner:      t.owner,
		LastOpened: int32(time.Now().Unix()),
	}

	// Convert to bytes
	bytes, err := proto.Marshal(saved)
	if err != nil {
		onError(err, "read")
		log.Fatalln(err)
	}

	// Send Complete Callback
	t.onComplete(bytes)
}
