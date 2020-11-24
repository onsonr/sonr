package file

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"sync"

	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
	path     string
	writer   *bufio.Writer
	buffer   bytes.Buffer
	//file     *os.File
	mutex sync.Mutex
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *pb.Metadata) SonrFile {
	docPath := fmt.Sprintf(docDir + "/" + meta.Name)

	buffer := new(bytes.Buffer)
	wr := bufio.NewWriter(buffer)

	return SonrFile{
		Metadata: meta,
		path:     docPath,
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

	// Flush to Buffer
	err := sf.writer.Flush()
	if err != nil {
		fmt.Println("error flushing sonr file")
	}

	// Decode Buffer
	imgByte := sf.buffer.Bytes()

	// Decode Buffer into Img
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		fmt.Println(err)
	}

	// Set Options
	var opts jpeg.Options
	opts.Quality = 1

	// Create File at Path
	f, err := os.Create(sf.path)
	defer f.Close()
	jpeg.Encode(f, img, &opts)

	// Return Block
	return sf.path
}
