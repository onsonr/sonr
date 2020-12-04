package transfer

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	pb "github.com/sonr-io/core/internal/models"
)

// ^ Helper: Encodes to Jpeg Image ^ //
func EncodeJpegBuffer(buf *bytes.Buffer, meta *pb.Metadata) error {
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
func EncodePngBuffer(buf *bytes.Buffer, meta *pb.Metadata) error {
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
