package models

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// ***************************** //
// ** Sonr File Incoming Info ** //
// ***************************** //

// ***************************** //
// ** Sonr File Outgoing Info ** //
// ***************************** //
// Returns first item in File
func (f *SonrFile) FirstFile() *SonrFile_Metadata {
	return f.Files[0]
}

// Checks if File contains single item
func (f *SonrFile) IsSingle() bool {
	return len(f.Files) == 1
}

// Checks if Single File is Media
func (f *SonrFile) IsMedia() bool {
	return f.Payload == Payload_MEDIA && f.IsSingle()
}

// Checks if File contains multiple items
func (f *SonrFile) IsMultiple() bool {
	return len(f.Files) > 1
}

// Method Returns SingleFile if Applicable
func (f *SonrFile) SingleFile() *SonrFile_Metadata {
	if f.IsSingle() {
		return f.Files[0]
	} else {
		return nil
	}
}

// Method Returns Metadata Item at Given Index
func (f *SonrFile) ItemAtIndex(index int) (*SonrFile_Metadata, error) {
	if index < int(f.GetCount()) {
		return f.Files[index], nil
	}
	return nil, errors.New("Item does not exist")
}

// Method Returns Preview from Thumbnail if Single File
func (f *SonrFile) Preview() []byte {
	// Validate Single
	if f.IsSingle() {
		// Retrieve Meta
		meta := f.Files[0]
		props := meta.GetProperties()

		// Check if Thumbnail Provided
		if props.HasThumbnail {
			// Initialize
			thumbWriter := new(bytes.Buffer)
			thumbReader := bytes.NewReader(meta.Thumbnail)

			// Convert to Image Object
			img, _, err := image.Decode(thumbReader)
			if err != nil {
				log.Println(err)
				return nil
			}

			// Encode as Jpeg into buffer w/o scaling
			err = jpeg.Encode(thumbWriter, img, nil)
			if err != nil {
				log.Panicln(err)
				return nil
			}

			return thumbWriter.Bytes()
		}
	}
	return nil
}

// Method Encodes Single File into Buffer
func (f *SonrFile) EncodeSingle(buf *bytes.Buffer) error {
	// Retreive First File
	pf := f.FirstFile()

	// @ Jpeg Image
	if ext := pf.Mime.Ext(); ext == "jpg" {
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

// Saves Item Data to Disk and Sets Update Item Path at Index
func (f *SonrFile) SaveItem(path string, data []byte, index int) error {
	if f.Files[index] != nil {
		// Write File to Disk
		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}

		// Set Path for Item
		f.Files[index].Path = path
		f.Files[index].Size = int32(len(data))
		return nil
	}
	return errors.New("Invalid Item index")
}

// Converts SonrFile to become Incoming
func (f *SonrFile) SetIncoming() {
	f.Direction = SonrFile_Incoming
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (f *SonrFile) ToCard(receiver *Peer, owner *Peer, preview []byte) *TransferCard {
	if f.Direction == SonrFile_Outgoing {
		// Create Card
		card := TransferCard{
			// SQL Properties
			Payload: f.Payload,

			// Owner Properties
			Receiver: receiver.GetProfile(),
			Owner:    owner.GetProfile(),
			File:     f,
		}

		// Set Preview
		if preview != nil {
			card.Preview = preview
		}
		return &card
	} else {
		// Update Direction
		f.Direction = SonrFile_Default

		// Create Card
		return &TransferCard{
			// SQL Properties
			Payload:  f.Payload,
			Received: int32(time.Now().Unix()),
			Preview:  preview,

			// Transfer Properties
			Status: TransferCard_COMPLETED,

			// Owner Properties
			Owner:    owner.GetProfile(),
			Receiver: receiver.GetProfile(),

			// Data Properties
			File: f,
		}
	}
}

// ************************** //
// ** MIME Info Management ** //
// ************************** //
// Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// *********************************** //
// ** Incoming File Info Management ** //
// *********************************** //
type InFile struct {
	Payload       Payload
	Metadata      *SonrFile
	ChunkBaseChan chan Chunk64
	ChunkBufChan  chan ChunkBuffer
}
