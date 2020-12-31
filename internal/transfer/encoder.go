package transfer

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	md "github.com/sonr-io/core/internal/models"
)

// ^ Helper: Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
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
func EncodeJpegBuffer(buf *bytes.Buffer, preview *md.Preview) error {
	// Open File at Meta Path
	file, err := os.Open(preview.Path)
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
func EncodePngBuffer(buf *bytes.Buffer, preview *md.Preview) error {
	// Open File at Meta Path
	file, err := os.Open(preview.Path)
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
