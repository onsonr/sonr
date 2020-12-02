package file

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type SonrFile struct {
	Metadata       *pb.Metadata
	Path           string
	progress       float32
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer
	mutex          sync.Mutex
}

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
)

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *pb.Metadata) SonrFile {
	return SonrFile{
		Metadata:       meta,
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
		Path:           docDir + "/" + meta.Name,
		progress:       0,
	}
}

// ^ Check file type and use corresponding method ^ //
func (sf *SonrFile) AddBuffer(buffer []byte) (bool, float32, error) {
	// @ Unmarshal Bytes into Proto
	chunk := pb.Chunk{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		fmt.Println("Unmarshal Error ", err)
		return true, 0, err
	}

	// Check File Type for Base64 Media
	if sf.Metadata.Mime.Type == "image" {
		hc, n, err := sf.addBase64(&chunk)
		return hc, n, err
	}

	// Add bytes buffer
	hc, n, err := sf.addBytes(&chunk)
	return hc, n, err
}

// ^ Add Bytes Buffer to SonrFile Buffer ^ //
func (sf *SonrFile) addBytes(chunk *pb.Chunk) (bool, float32, error) {
	// ** Lock ** //
	sf.mutex.Lock()
	// Add Block to Buffer
	n, err := sf.bytesBuilder.Write(chunk.Buffer)
	if err != nil {
		fmt.Println(err)
		return true, 0, err
	}
	sf.mutex.Unlock()
	// ** Unlock ** //

	// Update Progress
	sf.progress = float32(n)/float32(chunk.Total) + sf.progress

	// @ Check if Completed
	if sf.stringsBuilder.Len() == int(chunk.Total) {
		return true, 0, nil
	}
	return false, sf.progress, nil
}

// ^ Add Base64 Buffer to SonrFile Buffer ^ //
func (sf *SonrFile) addBase64(chunk *pb.Chunk) (bool, float32, error) {
	// ** Lock ** //
	sf.mutex.Lock()
	// Add Block to Buffer
	n, err := sf.stringsBuilder.WriteString(chunk.B64)
	if err != nil {
		fmt.Println(err)
		return true, 0, err
	}
	sf.mutex.Unlock()
	// ** Unlock ** //

	// Update Progress
	sf.progress = float32(n)/float32(chunk.Total) + sf.progress

	// @ Check if Completed
	if sf.stringsBuilder.Len() == int(chunk.Total) {
		return true, 0, nil
	}
	return false, sf.progress, nil
}

// ^ Check file type and use corresponding method ^ //
func (sf *SonrFile) Save(owner *pb.Peer) (*pb.Metadata, error) {
	// Check File Type for Base64 Media
	if sf.Metadata.Mime.Type == "image" {
		m, err := sf.saveBase64(owner)
		return m, err
	}

	// Add bytes buffer
	m, err := sf.saveBytes(owner)
	return m, err
}

// ^ Save file of type Base64 to Documents Directory and Return Path ^ //
func (sf *SonrFile) saveBytes(owner *pb.Peer) (*pb.Metadata, error) {
	// ** Lock/Unlock ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Get Base64 Data
	data := sf.bytesBuilder.Bytes()

	// Create File at Path
	f, err := os.Create(sf.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Write Bytes to to file
	if _, err := f.Write(data); err != nil {
		log.Fatalln(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(err)
	}

	// Get Info
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// Get File Type
	head := make([]byte, 261)
	f.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
	}

	// Get Mime Type
	mime := &pb.MIME{
		Type:    kind.MIME.Type,
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Set Metadata Protobuf Values
	return &pb.Metadata{
		Uuid:       uuid.New().String(),
		Name:       fileNameWithoutExtension(sf.Path),
		Path:       sf.Path,
		Size:       int32(info.Size()),
		Chunks:     int32(info.Size()) / BlockSize,
		Mime:       mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}

// ^ Save file of type Base64 to Documents Directory and Return Path ^ //
func (sf *SonrFile) saveBase64(owner *pb.Peer) (*pb.Metadata, error) {
	// ** Lock/Unlock ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Get Base64 Data
	data := sf.stringsBuilder.String()

	// Get Bytes from base64
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal("error:", err)
	}

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

	// Get Info
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// Get File Type
	head := make([]byte, 261)
	f.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
	}

	// Get Mime Type
	mime := pb.MIME{
		Type:    kind.MIME.Type,
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Set Metadata Protobuf Values
	return &pb.Metadata{
		Uuid:       uuid.New().String(),
		Name:       fileNameWithoutExtension(sf.Path),
		Path:       sf.Path,
		Size:       int32(info.Size()),
		Chunks:     int32(info.Size()) / BlockSize,
		Mime:       &mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}

func fileNameWithoutExtension(fileName string) string {
	base := filepath.Base(fileName)
	return strings.TrimSuffix(base, filepath.Ext(fileName))
}
