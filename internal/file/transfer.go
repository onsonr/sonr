package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

const B64ChunkSize = 63996 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 64000

// ^ Safely returns metadata depending on lock ^ //
func Base64(sf *SafeMeta) string {
	// Retreive Metadata
	meta := sf.Metadata()
	imgBuffer := new(bytes.Buffer)

	// New File for ThumbNail
	file, err := os.Open(meta.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Encode as Jpeg into buffer
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

	// Return B64 Encoded string
	b64 := base64.StdEncoding.EncodeToString(imgBuffer.Bytes())
	return b64
}

// ^ Chunks string based on B64ChunkSize ^ //
func ChunkBase64(sf *SafeMeta) []string {
	// Get String
	s := Base64(sf)
	// scanner := bufio.NewScanner(bytes.NewBufferString(Base64(sf)))
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
	// scanner := bufio.NewScanner(file)

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

func GetSize(sf *SafeMeta) int32 {
	// Get Metadata
	meta := sf.Metadata()

	// Check Type for image
	if meta.Mime.Type == "image" {
		// Return Adjusted Size
		return int32(len(Base64(sf)))
	} else {
		// Return Given Size
		return meta.Size
	}
}
