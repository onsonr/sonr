package stream

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

const B64ChunkSize = 63996 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 64000

// ^ Safely returns metadata depending on lock ^ //
func Base64(buffer []byte) (string, int) {
	// imgBuffer := new(bytes.Buffer)

	// // New File for ThumbNail
	// file, err := os.Open(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // Convert to Image Object
	// img, _, err := image.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Encode as Jpeg into buffer
	// err = png.Encode(file, img)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Return B64 Encoded string
	b64 := base64.StdEncoding.EncodeToString(buffer)
	return b64, len(b64)
}

// ^ Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
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
func ChunkBytes(path string, total int) [][]byte {
	// Open File
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	// scanner := bufio.NewScanner(file)

	// Set Chunk Variables
	size := BufferChunkSize
	pss := make([][]byte, 0, total/size+1)
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
