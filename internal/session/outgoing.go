package session

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	md "github.com/sonr-io/core/pkg/models"
)

type outgoingFile struct {
	// References
	Payload md.Payload
	Path    string
	mime    *md.MIME

	// Private Properties
	metadata *md.Metadata
	request  *md.InviteRequest
	preview  []byte
	owner    *md.Peer
	receiver *md.Peer
	size     int32
}

// ^ newOutgoingFile Processes Outgoing File ^ //
func newOutgoingFile(req *md.InviteRequest, p *md.Peer) *outgoingFile {
	var mime *md.MIME
	var payload md.Payload
	var size int32
	var err error

	// Check Values
	if req == nil || p == nil {
		return nil
	}

	// Get File Information
	file := req.Files[len(req.Files)-1]

	// Check if Mime Provided
	if file.Mime != nil {
		mime = file.GetMime()
	} else {
		// Get Mime
		mime, err = md.GetFileMime(file)
		if err != nil {
			return nil
		}
	}

	// Check if Size Provided
	if file.Size != 0 {
		size = file.GetSize()
	} else {
		// Get Size
		size, err = md.GetFileSize(file)
		if err != nil {
			return nil
		}
	}

	// Get Payload
	if req.Payload != md.Payload_UNDEFINED {
		payload = req.GetPayload()
	} else {
		payload = md.GetFilePayload(file)
	}

	// @ 1. Create new SafeFile
	sm := &outgoingFile{
		Path:     file.Path,
		Payload:  payload,
		request:  req,
		mime:     mime,
		receiver: req.To,
		owner:    p,
		metadata: file,
		size:     size,
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
			return nil
		}

		// @ Encode as Jpeg into buffer w/o scaling
		err = jpeg.Encode(thumbWriter, img, nil)
		if err != nil {
			log.Panicln(err)
			return nil
		}

		sm.preview = thumbWriter.Bytes()
	}
	return sm
}

// ^ Safely returns Preview depending on lock ^ //
func (sm *outgoingFile) Card() *md.TransferCard {
	// Create Card
	card := md.TransferCard{
		// SQL Properties
		Payload: sm.Payload,
		Preview: sm.preview,

		// Owner Properties
		Receiver: sm.receiver.GetProfile(),
		Owner:    sm.owner.GetProfile(),
		Metadata: sm.metadata,
	}

	if len(sm.preview) > 0 {
		card.Preview = sm.preview
	}
	return &card
}

// ^ Safely returns Preview depending on lock ^ //
func (s *Session) OutgoingCard() *md.TransferCard {
	return s.outgoing.Card()
}

// ^ Method adjusts extension for JPEG ^ //
func (pf *outgoingFile) Ext() string {
	if pf.mime.Subtype == "jpg" || pf.mime.Subtype == "jpeg" {
		return "jpeg"
	}
	return pf.mime.Subtype
}

// ^ Method Processes File at Path^ //
func (pf *outgoingFile) encodeFile(buf *bytes.Buffer) error {
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
