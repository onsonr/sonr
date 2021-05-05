package models

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// ***************************** //
// ** Sonr File Outgoing Info ** //
// ***************************** //
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

// Writes Buffer to File at Index with channel
func (f *SonrFile) AddItemAtIndex(index int, path string, first *Chunk, cCh chan *Chunk, pCh chan *Progress) {
	progress := NewItemProgress(first, int(f.Count), int(f.Size))
	stringsBuilder := new(strings.Builder)

	for {
		select {
		case chunk := <-cCh:
			// Add To Builder
			n, err := stringsBuilder.WriteString(chunk.Base)
			if err != nil {
				log.Println(err)
				break
			}
			// Send Progress if Met
			if met := progress.Add(n); met {
				// Get Progress
				p := progress.Progress()

				// Check Item Complete
				if p.ItemComplete {
					// Get Bytes from base64
					data, err := base64.StdEncoding.DecodeString(stringsBuilder.String())
					if err != nil {
						log.Println(err)
						break
					}

					if err := f.SaveItem(path, data, index); err != nil {
						log.Println(err)
						break
					}
				}
				pCh <- progress.Progress()
			}
		}
	}
}

// Returns SonrFile as TransferCard given Receiver and Owner
func (f *SonrFile) Card(receiver *Peer, owner *Peer) *TransferCard {
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
		return &card
	} else {
		// Update Direction
		f.Direction = SonrFile_Default

		// Create Card
		return &TransferCard{
			// SQL Properties
			Payload:  f.Payload,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    owner.GetProfile(),
			Receiver: receiver.GetProfile(),

			// Data Properties
			File: f,
		}
	}
}

// Method Encodes Single File into Buffer
func (f *SonrFile) Encode(index int, buf *bytes.Buffer) error {
	// Retreive File Metadata at Index
	pf, err := f.ItemAtIndex(index)
	if err != nil {
		return err
	}

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

// Method Returns SingleFile if Applicable
func (f *SonrFile) SingleFile() *SonrFile_Metadata {
	if f.IsSingle() {
		return f.Files[0]
	} else {
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

// Returns Item Transfer Chunk Count
func (f *SonrFile) TotalItemChunks(chunk *Chunk) int {
	return int(chunk.Total) / K_B64_CHUNK
}

// Returns Total Number of Transfer Chunks
func (f *SonrFile) TotalTranferChunks() int {
	return int(f.Size) / K_B64_CHUNK
}

// Returns Interval for Chunk
func (f *SonrFile) TransferInterval(chunk *Chunk) int {
	return f.TotalItemChunks(chunk) / 100
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
