// Code generated from Pkl module `orm`. DO NOT EDIT.
package keyencoding

import (
	"encoding"
	"fmt"
)

type KeyEncoding string

const (
	Raw       KeyEncoding = "raw"
	Hex       KeyEncoding = "hex"
	Multibase KeyEncoding = "multibase"
)

// String returns the string representation of KeyEncoding
func (rcv KeyEncoding) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(KeyEncoding)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for KeyEncoding.
func (rcv *KeyEncoding) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "raw":
		*rcv = Raw
	case "hex":
		*rcv = Hex
	case "multibase":
		*rcv = Multibase
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid KeyEncoding`, str)
	}
	return nil
}
