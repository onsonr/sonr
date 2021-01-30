package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

// ^ Method Processes File at Path^ //
func (pf *ProcessedFile) EncodeFile(buf *bytes.Buffer) error {
	// @ Jpeg Image
	if ext := pf.Ext(); ext == "jpg" {
		// Open File at Meta Path
		file, err := os.Open(pf.path)
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
		file, err := os.Open(pf.path)
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
		dat, err := ioutil.ReadFile(pf.path)
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

// ^ Encodes Image to Thumbnail Image: (buf) is reference to buffer ^ //
func EncodeThumb(buf *bytes.Buffer, path string) error {
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

	// @ Encode as Jpeg into buffer w/o scaling
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}

// ^ Generates Scaled Thumbnail for Image: (buf) is reference to buffer, (isScaled) is to scale image or not ^ //
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
