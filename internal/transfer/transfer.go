package transfer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	msgio "github.com/libp2p/go-msgio"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type OnProgress func(p float32)

type Transfer struct {
	// Inherited Properties
	mutex      sync.Mutex
	metadata   *md.Metadata
	onProgress OnProgress
	onComplete OnProtobuf
	path       string

	// Builders
	stringsBuilder *strings.Builder
	bytesBuilder   *bytes.Buffer

	// Tracking
	count int
	size  int
}

// ^ read Data from Msgio ^ //
func (dsc *PeerConnection) ReadStream(reader msgio.ReadCloser) {
	for i := 0; ; i++ {
		// @ Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			onError(err, "ReadStream")
			log.Fatalln(err)
			break
		}

		// @ Unmarshal Bytes into Proto
		chunk := md.Chunk{}
		err = proto.Unmarshal(buffer, &chunk)
		if err != nil {
			onError(err, "ReadStream")
			log.Fatalln(err)
		}
	}
}

// ^ Create new SonrFile struct with meta and documents directory ^ //
func NewTransfer(docDir string, meta *md.Metadata, op OnProgress, oc OnProtobuf) Transfer {
	return Transfer{
		// Inherited Properties
		metadata:   meta,
		path:       docDir + "/" + meta.Name + "." + meta.Mime.Subtype,
		onProgress: op,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),

		// Tracking
		count: 0,
		size:  0,
	}
}

// ^ Check file type and use corresponding method ^ //
func (t *Transfer) AddBuffer(buffer []byte) (bool, error) {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		return true, err
	}

	// @ Increment received count
	if t.count == 0 {
		t.size = int(chunk.Total)
	}

	// Set Tracking
	var n int
	t.count++

	// @ Check File Type
	if chunk.GetB64() != "" {
		// Add Base64 Chunk to Buffer
		n, err = t.stringsBuilder.WriteString(chunk.B64)
		if err != nil {
			return true, err
		}
	} else {
		// Add ByteChunk to Buffer
		n, err = t.bytesBuilder.Write(chunk.Buffer)
		if err != nil {
			return true, err
		}
	}

	// @ Update Tracking
	currW := t.count*BufferChunkSize + n
	currP := float32(currW) / float32(t.size)
	t.onProgress(currP)

	if t.stringsBuilder.Len() == t.size || t.bytesBuilder.Len() == t.size {
		return true, nil
	}
	return false, nil
}

// ^ Check file type and use corresponding method to Save to Disk ^ //
func (t *Transfer) Save(owner *md.Peer) (*md.Metadata, error) {
	// ** Lock/Unlock ** //
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// Initialize
	var fileBytes []byte

	// @ Set File Bytes by Type
	if t.metadata.Mime.Type == md.MIME_image {
		// Get Base64 Data
		data := t.stringsBuilder.String()

		// Get Bytes from base64
		b64Bytes, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Fatal("error:", err)
		}

		// Set Bytes from Base64
		fileBytes = b64Bytes
	} else {
		// Set Bytes from Buffer
		fileBytes = t.bytesBuilder.Bytes()
	}

	// @ Create File at Path
	f, err := os.Create(t.path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// @ Write Bytes to to file
	if _, err := f.Write(fileBytes); err != nil {
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
		Name:       t.metadata.Name,
		Path:       t.path,
		Size:       int32(info.Size()),
		Mime:       t.metadata.Mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}
