package file

import (
	"bytes"
	"fmt"
	"sync"

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
func (sf *SonrFile) AddBlock(block []byte) {
	sf.mutex.Lock()
	// Add Block to Buffer
	written, err := sf.buffer.Write(block)
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

	// Encode as Jpeg
	bytes := sf.buffer.Bytes()
	_, err := EncodeImage(bytes, sf.path)
	if err != nil {
		return "", err
	}

	// Return Block
	return sf.path, nil
}
