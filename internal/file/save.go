package file

import (
	"bufio"
	"fmt"
	"log"
	"os"

	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
	path     string
	writer   *bufio.Writer
	file     *os.File
}

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

func (sf *SonrFile) AddBlock(block []byte) {
	// Add Block to Buffer
	written, err := sf.writer.Write(block)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bytes Written: ", written)
}

func (sf *SonrFile) Save() {
	err := sf.writer.Flush()
	if err != nil {
		fmt.Println("error flushing sonr file")
	}
}
