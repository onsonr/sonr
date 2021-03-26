package file

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/libp2p/go-msgio"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

// @ File that safely sets metadata and thumbnail in routine
type FileItem struct {
	// References
	Payload md.Payload
	call    dt.NodeCallback
	mime    *md.MIME
	Path    string

	// Private Properties
	mutex   sync.Mutex
	card    md.TransferCard
	request *md.InviteRequest
}

// ^ NewFileItem Processes Outgoing File ^ //
func NewFileItem(req *md.InviteRequest, p *md.Profile, callback dt.NodeCallback) *FileItem {
	// Check Values
	if req == nil || p == nil {
		return nil
	}

	// Get File Information
	file := req.Files[len(req.Files)-1]
	info, err := GetFileInfo(file.Path)
	if err != nil {
		callback.Error(err, "NewProcessedFile:GetFileInfo")
	}

	// @ 1. Create new SafeFile
	sm := &FileItem{
		call:    callback,
		Path:    file.Path,
		Payload: info.Payload,
		request: req,
		mime:    info.Mime,
	}

	// ** Lock ** //
	sm.mutex.Lock()

	// @ 2. Set Metadata Protobuf Values
	// Create Card
	sm.card = md.TransferCard{
		// SQL Properties
		Payload:  info.Payload,
		Platform: p.Platform,

		// Owner Properties
		Username:  p.Username,
		FirstName: p.FirstName,
		LastName:  p.LastName,

		Properties: &md.TransferCard_Properties{
			Name: info.Name,
			Size: info.Size,
			Mime: info.Mime,
		},
	}

	// @ 3. Create Thumbnail in Goroutine
	if len(file.Thumbnail) > 0 {
		// Initialize
		thumbWriter := new(bytes.Buffer)
		thumbReader := bytes.NewReader(file.Thumbnail)

		// Convert to Image Object
		img, _, err := image.Decode(thumbReader)
		if err != nil {
			log.Println(err)
		}

		// @ Encode as Jpeg into buffer w/o scaling
		err = jpeg.Encode(thumbWriter, img, nil)
		if err != nil {
			log.Panicln(err)
		}

		sm.card.Preview = thumbWriter.Bytes()
	}
	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Transfer Card
	preview := sm.Card()

	// @ 3. Callback with Preview
	sm.call.Queued(preview, sm.request)
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *FileItem) Card() *md.TransferCard {
	// ** Lock File wait for access ** //
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// @ 2. Return Value
	return &sm.card
}

// ^ Method adjusts extension for JPEG ^ //
func (pf *FileItem) Ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpeg"
	}
	return pf.mime.Subtype
}

// ^ Method Processes File at Path^ //
func (pf *FileItem) EncodeMedia(buf *bytes.Buffer) error {
	// @ Jpeg Image
	if ext := pf.Ext(); ext == "jpg" {
		// Open File at Meta Path
		file, err := os.Open(pf.Path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			return err
		}

		// Encode as Jpeg into buffer
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
		return nil

		// @ PNG Image
	} else if ext == "png" {
		// Open File at Meta Path
		file, err := os.Open(pf.Path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			return err
		}

		// Encode as Jpeg into buffer
		err = png.Encode(buf, img)
		if err != nil {
			return err
		}
		return nil

		// @ Other - Open File at Path
	} else {
		dat, err := ioutil.ReadFile(pf.Path)
		if err != nil {
			return err
		}

		// Write Bytes to buffer
		_, err = buf.Write(dat)
		if err != nil {
			return err
		}
		return nil
	}
}

// ^ write fileItem as Base64 in Msgio to Stream ^ //
func (pf *FileItem) WriteBase64(writer msgio.WriteCloser, peer *md.Peer) {
	// Initialize Buffer and Encode File
	var base string
	if pf.Payload == md.Payload_MEDIA {
		buffer := new(bytes.Buffer)

		if err := pf.EncodeMedia(buffer); err != nil {
			log.Fatalln(err)
		}

		// Encode Buffer to base 64
		data := buffer.Bytes()
		base = base64.StdEncoding.EncodeToString(data)
	} else {
		data, err := ioutil.ReadFile(pf.Path)
		if err != nil {
			log.Fatalln(err)
		}
		base = base64.StdEncoding.EncodeToString(data)
	}

	// Set Total
	total := int32(len(base))

	// Iterate for Entire file as String
	for _, dat := range ChunkBase64(base) {
		// Create Block Protobuf from Chunk
		chunk := md.Chunk64{
			Size:  int32(len(dat)),
			Data:  dat,
			Total: total,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(&chunk)
		if err != nil {
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
		dt.GetState().NeedsWait()
	}

	// Call Completed Sending
	pf.call.Transmitted(peer)
}
