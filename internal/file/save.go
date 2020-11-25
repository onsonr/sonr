package file

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	pb "github.com/sonr-io/core/internal/models"
)

type SonrFile struct {
	Metadata *pb.Metadata
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
	}
}

// ^ Add Block to SonrFile Buffer ^ //
func (sf *SonrFile) AddBlock(block []byte) {
	// ** Lock ** //
	sf.mutex.Lock()
	// Add Block to Buffer
	written, err := sf.builder.Write(block)
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

	// Verify Image Type
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}

	// Open Base64 Decode
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	// Check Image Configuration, Retreive Format
	imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	if imgCfg.Width != 750 || imgCfg.Height != 685 {
		return "", ErrSize
	}

	// Write Data as FileName to Path
	fileName := sf.Metadata.Name + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	// Create Delay
	time.After(time.Second)

	// Return Block
	return fileName, nil
}
