package stream

import (
	"encoding/base64"
)

const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 32000

// ^ Safely returns metadata depending on lock ^ //
func Base64(buffer []byte) (string, int) {
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
