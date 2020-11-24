package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
	path     string
	writer   *bufio.Writer
	file     *os.File
	mutex    sync.Mutex
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *pb.Metadata) SonrFile {
	docPath := fmt.Sprintf(docDir + "/" + meta.Name)

	file, err := os.Create(docPath)
	if err != nil {
		log.Fatal(err)
	}

	wr := bufio.NewWriter(file)

	return SonrFile{
		Metadata: meta,
		path:     docPath,
		file:     file,
		writer:   wr,
	}
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

	// Close The File
	err = sf.file.Close()
	if err != nil {
		fmt.Println("error closing sonr file")
	}

	// Encode as Jpeg
	
	// Return Block
	return sf.path
}
