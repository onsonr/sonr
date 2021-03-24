package data

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sync"

	"github.com/nfnt/resize"
	md "github.com/sonr-io/core/internal/models"
)

// @ File that safely sets metadata and thumbnail in routine
type ProcessedFile struct {
	// References
	Payload md.Payload
	call    NodeCallback
	mime    *md.MIME
	Path    string

	// Private Properties
	mutex   sync.Mutex
	card    md.TransferCard
	request *md.InviteRequest
}

// @ ProcessedFileBuilder creates a new item and returns a pointer to it.
func ProcessedFileBuilder() interface{} {
	return &ProcessedFile{}
}

// ^ NewProcessedFile Processes Outgoing File ^ //
func NewProcessedFile(req *md.InviteRequest, p *md.Peer, callback NodeCallback) *ProcessedFile {
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
	sm := &ProcessedFile{
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
		Username:  p.Profile.Username,
		FirstName: p.Profile.FirstName,
		LastName:  p.Profile.LastName,

		Properties: &md.TransferCard_Properties{
			Name: info.Name,
			Size: info.Size,
			Mime: info.Mime,
		},
	}

	// @ 3. Create Thumbnail in Goroutine
	if len(file.Thumbnail) > 0 {
		go HandleThumbnail(file, sm, p)
	} else {
		go RequestThumbnail(file, sm, p)
	}
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *ProcessedFile) Card() *md.TransferCard {
	// ** Lock File wait for access ** //
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// @ 2. Return Value
	return &sm.card
}

// ^ Method to generate thumbnail for ProcessRequest^ //
func RequestThumbnail(reqFi *md.InviteRequest_FileInfo, sm *ProcessedFile, p *md.Peer) {
	// Initialize
	thumbBuffer := new(bytes.Buffer)
	// @ Handle Created File Request
	// Validate Image
	if sm.mime.Type == md.MIME_image {
		// Encode Thumbnail
		err := GenerateThumb(thumbBuffer, reqFi.Path)
		if err != nil {
			log.Panicln(err)
		}

		// Update Thumbnail Value
		sm.card.Preview = thumbBuffer.Bytes()

	}

	// ** Unlock ** //
	sm.mutex.Unlock()

	// @ 3. Callback with Preview
	sm.call.HandleInvite(sm.Card(), sm.request, p, sm)
}

// ^ Method to Handle Provided Thumbnail ^ //
func HandleThumbnail(reqFi *md.InviteRequest_FileInfo, sm *ProcessedFile, p *md.Peer) {
	// Initialize
	thumbWriter := new(bytes.Buffer)
	thumbReader := bytes.NewReader(reqFi.Thumbnail)

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

	// ** Unlock ** //
	sm.mutex.Unlock()

	// @ 3. Callback with Preview
	sm.call.HandleInvite(sm.Card(), sm.request, p, sm)
}

// ^ Method adjusts extension for JPEG ^ //
func (pf *ProcessedFile) ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpeg"
	}
	return pf.mime.Subtype
}

// ^ Method Processes File at Path^ //
func (pf *ProcessedFile) EncodeFile(buf *bytes.Buffer) error {
	// @ Jpeg Image
	if ext := pf.ext(); ext == "jpg" {
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

// ^ Generates Scaled Thumbnail for Image: (buf) is reference to buffer ^ //
func GenerateThumb(buf *bytes.Buffer, path string) error {
	// @ Open File at Meta Path
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Retreive Bounds
	b := img.Bounds()
	w, h := b.Max.X, b.Max.Y

	// ** Resize Constants ** //
	const MAX_WIDTH float64 = 320
	const MAX_HEIGHT float64 = 240

	// Get Ratio
	ratio := math.Min(MAX_WIDTH/float64(w), MAX_HEIGHT/float64(h))

	// Calculate Fit and Scale Image
	newW, newH := int(math.Ceil(float64(w)*ratio)), int(math.Ceil(float64(h)*ratio))
	scaledImage := resize.Resize(uint(newW), uint(newH), img, resize.Lanczos3)

	// @ Encode as Jpeg into buffer
	err = jpeg.Encode(buf, scaledImage, nil)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}
