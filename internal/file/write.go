package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type TransferFile struct {
	Metadata       *md.Metadata
	SavePath       string
	progress       float32
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer
	mutex          sync.Mutex
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewFile(docDir string, meta *md.Metadata) TransferFile {
	return TransferFile{
		Metadata:       meta,
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
		SavePath:       docDir + "/" + meta.Name + "." + meta.Mime.Subtype,
		progress:       0,
	}
}

// ^ Check file type and use corresponding method ^ //
func (sf *TransferFile) AddBuffer(buffer []byte) (bool, float32, error) {
	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		fmt.Println("Unmarshal Error ", err)
		return true, 0, err
	}

	// Check File Type for Base64 Media
	if sf.Metadata.Mime.Type == md.MIME_image {
		hc, n, err := sf.addBase64(&chunk)
		return hc, n, err
	}

	// Add bytes buffer
	hc, n, err := sf.addBytes(&chunk)
	return hc, n, err
}

// ^ Add Bytes Buffer to SonrFile Buffer ^ //
func (sf *TransferFile) addBytes(chunk *md.Chunk) (bool, float32, error) {
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
func (sf *TransferFile) addBase64(chunk *md.Chunk) (bool, float32, error) {
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
func (sf *TransferFile) Save(owner *md.Peer) (*md.Metadata, error) {
	// Check File Type for Base64 Media
	if sf.Metadata.Mime.Type == md.MIME_image {
		m, err := sf.saveBase64(owner)
		return m, err
	}

	// Add bytes buffer
	m, err := sf.saveBytes(owner)
	return m, err
}

// ^ Save file of type Base64 to Documents Directory and Return Path ^ //
func (sf *TransferFile) saveBytes(owner *md.Peer) (*md.Metadata, error) {
	// ** Lock/Unlock ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Get Base64 Data
	data := sf.bytesBuilder.Bytes()

	// Create File at Path
	f, err := os.Create(sf.SavePath)
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

	// @ 3. Set Metadata Protobuf Values
	return &md.Metadata{
		Name:       sf.Metadata.Name,
		Path:       sf.SavePath,
		Size:       int32(info.Size()),
		Mime:       sf.Metadata.Mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}

// ^ Save file of type Base64 to Documents Directory and Return Path ^ //
func (sf *TransferFile) saveBase64(owner *md.Peer) (*md.Metadata, error) {
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
	f, err := os.Create(sf.SavePath)
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

	// @ 3. Set Metadata Protobuf Values
	return &md.Metadata{
		Name:       sf.Metadata.Name,
		Path:       sf.SavePath,
		Size:       int32(info.Size()),
		Mime:       sf.Metadata.Mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}
