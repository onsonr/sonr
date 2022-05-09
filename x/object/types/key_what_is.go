package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// WhatIsKeyPrefix is the prefix to retrieve all WhatIs
	WhatIsKeyPrefix = "WhatIs/value/"
)

func WhatIsKey(did string) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
