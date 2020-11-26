package file

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
	Path     string
	builder  *strings.Builder
	mutex    sync.Mutex
}

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
)

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *pb.Metadata) SonrFile {

	return SonrFile{
		Metadata: meta,
		builder:  new(strings.Builder),
		Path:     docDir + "/" + meta.Name,
	}
}

// ^ Add Block to SonrFile Buffer ^ //
func (sf *SonrFile) AddBlock(block string) {
	// ** Lock ** //
	sf.mutex.Lock()
	// Add Block to Buffer
	written, err := sf.builder.WriteString(block)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bytes Written: ", written)
	fmt.Println("Bytes Total: ", sf.builder.Len())
	sf.mutex.Unlock()
	// ** Unlock ** //
}

// ^ Save file of Documents Directory and Return Path ^ //
func (sf *SonrFile) Save() (string, error) {
	// ** Lock/Unlock ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Get Base64 Data
	data := sf.builder.String()

	// Get Bytes from base64
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal("error:", err)
	}

	// Create Delay
	time.After(time.Millisecond * 500)

	// Create File at Path
	f, err := os.Create(sf.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Write Bytes to to file
	if _, err := f.Write(bytes); err != nil {
		log.Fatalln(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(err)
	}

	// Return Block
	return sf.Path, nil
}

// ^ Creates Metadata for File at Path ^ //
func GetMetadata(path string, owner *pb.Peer) (*pb.Metadata, error) {
	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening File", err)
		return nil, err
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

	// Get Mime Type
	mime := &pb.Metadata_MIME{
		Type:    kind.MIME.Type,
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Set Metadata Protobuf Values
	return &pb.Metadata{
		Uuid:       uuid.New().String(),
		Name:       fileNameWithoutExtension(path),
		Path:       path,
		Size:       int32(info.Size()),
		Chunks:     int32(info.Size()) / BlockSize,
		Mime:       mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}
