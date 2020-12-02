package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
)

const B64ChunkSize = 63996 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 64000

// ^ Safely returns metadata depending on lock ^ //
func Base64(sf *SafeMeta) (string, error) {
	// Retreive Metadata
	meta := sf.Metadata()
	imgBuffer := new(bytes.Buffer)

	// New File for ThumbNail
	file, err := os.Open(meta.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Encode as Jpeg into buffer
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return "", err
	}

	// Return B64 Encoded string
	b64 := base64.StdEncoding.EncodeToString(imgBuffer.Bytes())
	return b64, nil
}

// ^ Chunks string based on B64ChunkSize ^ //
func ChunkBase64(sf *SafeMeta) []string {
	// Get String
	s, _ := Base64(sf)
	size := B64ChunkSize
	ss := make([]string, 0, len(s)/size+1)
	for len(s) > 0 {
		if len(s) < size {
			size = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:size]), s[size:]
	}
	return ss
}

// ^ Chunk bytes from FilePath based on Buffer Chunk Size ^ //
func ChunkBytes(sf *SafeMeta) [][]byte {
	// Open File
	meta := sf.Metadata()
	file, err := os.Open(meta.Path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// Set Chunk Variables
	size := BufferChunkSize
	pss := make([][]byte, 0, int(meta.Size)/size+1)
	ps := make([]byte, BufferChunkSize)

	// Iterate file
	for {
		// Read Bytes
		bytesread, err := file.Read(ps)

		// Check for Error
		if err != nil {
			// Non EOF Error
			if err != io.EOF {
				fmt.Println(err)
			}
			// File Complete
			break
		}

		// Set Current Chunk to how many bytes read
		pss, ps = append(pss, ps[:bytesread]), ps[bytesread:]
	}
	return pss
}

func GetSize(sf *SafeMeta) (int32, error) {
	// Get Metadata
	meta := sf.Metadata()

	// Check Type for image
	if meta.Mime.Type == "image" {
		s, err := Base64(sf)

		// Handle Error
		if err != nil {
			return 0, err
		}

		// Return Adjusted Size
		return int32(len(s)), nil
	} else {
		// Return Given Size
		return meta.Size, nil
	}
}
