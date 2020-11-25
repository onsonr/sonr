package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
	path     string
	buffer   *bytes.Buffer
	mutex    sync.Mutex
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *pb.Metadata) SonrFile {
	docPath := fmt.Sprintf(docDir + "/" + meta.Name)

	return SonrFile{
		Metadata: meta,
		path:     docPath,
		buffer:   new(bytes.Buffer),
	}
}

// ^ Add Block to SonrFile Buffer ^ //
func (sf *SonrFile) AddBlock(block string) {
	sf.mutex.Lock()
	// Add Block to Buffer
	written, err := sf.buffer.WriteString(block)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bytes Written: ", written)
	sf.mutex.Unlock()
}

// ^ Save file of Documents Directory and Return Path ^ //
func (sf *SonrFile) Save() (string, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Open output file
	output, err := os.Create(sf.path)
	if err != nil {
		return "", err
	}
	// Close output file
	defer output.Close()

	// Create base64 stream decoder from input file. *io.File implements the
	// io.Reader interface. In other words we can pass it to NewDecoder.
	decoder := base64.NewDecoder(base64.StdEncoding, sf.buffer)

	// Magic! Copy from base64 decoder to output file
	_, err = io.Copy(output, decoder)

	// Check for Error
	if err != nil {
		return "", err
	}

	// Create Delay
	time.After(time.Second)

	// Return Block
	return sf.path, nil
}
