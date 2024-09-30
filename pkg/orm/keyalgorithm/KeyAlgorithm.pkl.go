// Code generated from Pkl module `orm`. DO NOT EDIT.
package keyalgorithm

import (
	"encoding"
	"fmt"
)

type KeyAlgorithm string

const (
	Es256  KeyAlgorithm = "es256"
	Es384  KeyAlgorithm = "es384"
	Es512  KeyAlgorithm = "es512"
	Eddsa  KeyAlgorithm = "eddsa"
	Es256k KeyAlgorithm = "es256k"
	Ecdsa  KeyAlgorithm = "ecdsa"
)

// String returns the string representation of KeyAlgorithm
func (rcv KeyAlgorithm) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(KeyAlgorithm)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for KeyAlgorithm.
func (rcv *KeyAlgorithm) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "es256":
		*rcv = Es256
	case "es384":
		*rcv = Es384
	case "es512":
		*rcv = Es512
	case "eddsa":
		*rcv = Eddsa
	case "es256k":
		*rcv = Es256k
	case "ecdsa":
		*rcv = Ecdsa
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid KeyAlgorithm`, str)
	}
	return nil
}
