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

	"github.com/libp2p/go-msgio"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

// ^ Method Processes File at Path^ //
func (pf *FileItem) EncodeMedia(buf *bytes.Buffer) error {
	// @ Jpeg Image
	if ext := pf.Ext; ext == "jpg" {
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

// ^ Returns Total Incoming Buffer ^ //
func (i *FileItem) GetIncomingBuffer() ([]byte, error) {
	// Get Bytes from base64
	return base64.StdEncoding.DecodeString(i.stringsBuilder.String())
}

// ^ Check file type and use corresponding method ^ //
func (t *FileItem) WriteFromStream(curr int, buffer []byte) (*md.InFileProgress, error) {
	// ** Lock/Unlock ** //

	// @ Unmarshal Bytes into Proto
	chunk := md.Chunk64{}
	err := proto.Unmarshal(buffer, &chunk)
	if err != nil {
		return nil, err
	}

	// @ Add Buffer by File Type
	// Add Base64 Chunk to Buffer
	n, err := t.stringsBuilder.WriteString(chunk.Data)
	if err != nil {
		return nil, err
	}

	// Update Tracking
	return t.inInfo.UpdateTracking(n, curr), nil
}

// ^ write fileItem as Base64 in Msgio to Stream ^ //
func (pf *FileItem) WriteToStream(writer msgio.WriteCloser, peer *md.Peer, hc chan bool) {
	// Initialize Buffer and Encode File
	var base string
	if pf.Payload == md.Payload_MEDIA {
		buffer := new(bytes.Buffer)

		if err := pf.EncodeMedia(buffer); err != nil {
			log.Fatalln(err)
			hc <- false
		}

		// Encode Buffer to base 64
		data := buffer.Bytes()
		base = base64.StdEncoding.EncodeToString(data)
	} else {
		data, err := ioutil.ReadFile(pf.Path)
		if err != nil {
			log.Fatalln(err)
			hc <- false
		}
		base = base64.StdEncoding.EncodeToString(data)
	}

	// Set Total
	total := int32(len(base))

	// Iterate for Entire file as String
	for _, dat := range chunkBase64(base) {
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
			hc <- false
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
			hc <- false
		}
		dt.GetState().NeedsWait()
	}

	// Call Completed Sending
	hc <- true
}

// @ Helper: Chunks string based on B64ChunkSize ^ //
func chunkBase64(s string) []string {
	chunkSize := K_B64_CHUNK
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
}
