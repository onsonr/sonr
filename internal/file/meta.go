package file

import (
	"os"
	"path/filepath"
	"sync"

	"fmt"
	_ "image/gif"
	_ "image/jpeg"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// Define Function Types
type OnQueued func(data []byte)
type OnProgress func(data float32)
type OnError func(err error, method string)

// Define Block Size
const BlockSize = 16000

// Struct to Implement Node Callback Methods
type FileCallback struct {
	Queued OnQueued
	Error  OnError
}

// ^ File that safely sets metadata and thumbnail in routine ^ //
type SafeMeta struct {
	Path  string
	Call  FileCallback
	mutex sync.Mutex
	meta  pb.Metadata
}

// ^ Create generates file metadata ^ //
func (sm *SafeMeta) NewMetadata() {
	// ** Lock ** //
	sm.mutex.Lock()

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(sm.Path)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening File", err)
		sm.Call.Error(err, "AddFile")
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		sm.Call.Error(err, "AddFile")
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
		sm.Call.Error(err, "AddFile")
	}

	// Get Mime Type
	mime := &pb.MIME{
		Type:    kind.MIME.Type,
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 2. Set Metadata Protobuf Values
	sm.meta = pb.Metadata{
		Uuid: uuid.New().String(),
		Name: filepath.Base(sm.Path),
		Path: sm.Path,
		Size: int32(info.Size()),
		Mime: mime,
	}

	// @ 3. Create Thumbnail
	if filetype.IsImage(head) {
		// New File for ThumbNail
		thumbBytes, err := NewThumbnail(sm.Path)
		if err != nil {
			fmt.Println(err)
			sm.Call.Error(err, "AddFile")
		}
		// Update Metadata Value
		sm.meta.Thumbnail = thumbBytes
	}

	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Metadata
	meta := sm.Metadata()

	// Convert to bytes
	data, err := proto.Marshal(meta)
	if err != nil {
		fmt.Println("Error Marshaling Metadata ", err)
	}

	// Callback with Metadata
	sm.Call.Queued(data)
}

// ^ Safely returns metadata depending on lock ^ //
func (sf *SafeMeta) Metadata() *pb.Metadata {
	// ** Lock File wait for access ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// @ 2. Return Value
	return &sf.meta
}
