package stream

import (
	"encoding/base64"
)

// ^ Safely returns metadata depending on lock ^ //
func Base64(buffer []byte) (string, int) {
	// Return B64 Encoded string
	b64 := base64.StdEncoding.EncodeToString(buffer)
	return b64, len(b64)
}
