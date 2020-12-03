package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"fmt"
	_ "image/gif"
	_ "image/jpeg"

	"github.com/h2non/filetype"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// Define Function Types
type OnProtobuf func(data []byte)
type OnProgress func(data float32)
type OnError func(err error, method string)

// Define Block Size
const BlockSize = 16000

// Struct to Implement Node Callback Methods
type FileCallback struct {
}

// ^ File that safely sets metadata and thumbnail in routine ^ //
type SafeFile struct {
	// Callbacks
	CallQueued   OnProtobuf
	CallProgress OnProgress
	CallError    OnError

	// Public Properties
	Mime *pb.MIME

	// Private Properties
	mutex sync.Mutex
	meta  pb.Metadata
	path  string
}

// ^ Create generates file metadata ^ //
func NewMetadata(filePath string, queueCall OnProtobuf, progCall OnProgress, errCall OnError) *SafeFile {
	// Create new SafeFile
	sf := &SafeFile{
		CallQueued:   queueCall,
		CallProgress: progCall,
		CallError:    errCall,
		path:         filePath,
	}

	// ** Lock ** //
	sf.mutex.Lock()

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(sf.path)
	if err != nil {
		log.Fatalln(err)
		sf.CallError(err, "AddFile")
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
		sf.CallError(err, "AddFile")
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		log.Fatalln(err)
		sf.CallError(err, "AddFile")
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		log.Fatalln(err)
		sf.CallError(err, "AddFile")
	}

	// @ 2. Set Metadata Protobuf Values
	// Set Metadata
	sf.meta = pb.Metadata{
		Name: getFileName(sf.path),
		Path: sf.path,
		Size: int32(info.Size()),
		Mime: &pb.MIME{
			Type:    pb.MIME_Type(pb.MIME_Type_value[kind.MIME.Type]),
			Subtype: kind.MIME.Subtype,
			Value:   kind.MIME.Value,
		},
	}

	// Set Mime
	sf.Mime = sf.meta.Mime

	// @ 3. Create Thumbnail in Goroutine
	go func(sf *SafeFile) {
		if filetype.IsImage(head) {
			// New File for ThumbNail
			thumbBytes, err := NewThumbnail(sf.path)
			if err != nil {
				fmt.Println(err)
				sf.CallError(err, "AddFile")
			}
			// Update Metadata Value
			sf.meta.Thumbnail = thumbBytes
		}

		// ** Unlock ** //
		sf.mutex.Unlock()

		// Get Metadata
		meta := sf.Metadata()

		// Convert to bytes
		data, err := proto.Marshal(meta)
		if err != nil {
			fmt.Println("Error Marshaling Metadata ", err)
		}

		// Callback with Metadata
		sf.CallQueued(data)
	}(sf)
	return sf
}

// ^ Safely returns metadata depending on lock ^ //
func (sf *SafeFile) Metadata() *pb.Metadata {
	// ** Lock File wait for access ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// @ 2. Return Value
	return &sf.meta
}

// ^ Helper: Get File name from Path ^ //
func getFileName(path string) string {
	// Get File Name
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(path))
}
