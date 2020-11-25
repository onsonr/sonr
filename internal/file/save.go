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
	time.After(time.Second)

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
