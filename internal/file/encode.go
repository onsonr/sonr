package file

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	pb "github.com/sonr-io/core/internal/models"
)

// ^ Safely returns Base64, and Size as 32 ^ //
func (sf *SafeFile) Base64() (string, int32) {
	// ** Lock ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// Initialize
	var result string
	var size int
	meta := sf.Metadata()

	// @ Check Type for image
	if meta.Mime.Type == pb.MIME_image {
		// Initialize Buffer
		imgBuffer := new(bytes.Buffer)

		// @ Check Image type
		if meta.Mime.Subtype == "jpeg" {
			// Get JPEG Encoded Buffer
			err := getJpegBuffer(imgBuffer, meta)
			if err != nil {
				sf.CallError(err, "Base64")
				log.Fatalln(err)
			}
		} else if meta.Mime.Subtype == "png" {
			// Get PNG Encoded Buffer
			err := getPngBuffer(imgBuffer, meta)
			if err != nil {
				sf.CallError(err, "Base64")
				log.Fatalln(err)
			}
		}

		// Encode Buffer to base 64
		imgBytes := imgBuffer.Bytes()
		result = base64.StdEncoding.EncodeToString(imgBytes)
		size = len(result)
	}

	// ** Unlock
	sf.mutex.Unlock()

	// Return B64 Encoded string
	return result, int32(size)
}

// ^ Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string, B64ChunkSize int) []string {
	chunkSize := B64ChunkSize
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

// ^ Helper: Encodes to Jpeg Image ^ //
func getJpegBuffer(buf *bytes.Buffer, meta *pb.Metadata) error {
	// Open File at Meta Path
	file, err := os.Open(meta.Path)
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
}

// ^ Helper: Encodes to PNG Image ^ //
func getPngBuffer(buf *bytes.Buffer, meta *pb.Metadata) error {
	// Open File at Meta Path
	file, err := os.Open(meta.Path)
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
}
