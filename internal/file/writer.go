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
func (pf *FileItem) EncodeFile(buf *bytes.Buffer) error {
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
	buffer := new(bytes.Buffer)

	if err := pf.EncodeFile(buffer); err != nil {
		log.Fatalln(err)
	}

	// Encode Buffer to base 64
	data := buffer.Bytes()
	base := base64.StdEncoding.EncodeToString(data)

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

// ^ Helper: Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
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
