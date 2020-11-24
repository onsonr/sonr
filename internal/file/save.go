package file

import (
	"bufio"
	"bytes"
	"fmt"
	"sync"

	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
	path     string
	writer   *bufio.Writer
	buffer   bytes.Buffer
	mutex    sync.Mutex
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *pb.Metadata) *SonrFile {
	docPath := fmt.Sprintf(docDir + "/" + meta.Name)

	sf := SonrFile{
		Metadata: meta,
		path:     docPath,
	}

	sf.writer = bufio.NewWriter(&sf.buffer)
	return &sf
}

// ^ Add Block to SonrFile Buffer ^ //
func (sf *SonrFile) AddBlock(block []byte) {
	sf.mutex.Lock()
	// Add Block to Buffer
	written, err := sf.writer.Write(block)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bytes Written: ", written)
	sf.mutex.Unlock()
}

// ^ Save file of Documents Directory and Return Path ^ //
func (sf *SonrFile) Save() string {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Wait for block to be added
	err := sf.writer.Flush()
	if err != nil {
		fmt.Println("error flushing sonr file")
	}

	if sf.Metadata.Kind == "image" {
		EncodeImage(sf.buffer, sf.path)
	}
	// Return Block
	return sf.path
}
