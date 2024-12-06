// Code generated from Pkl module `sonr.motr.ORM`. DO NOT EDIT.
package keysharerole

import (
	"encoding"
	"fmt"
)

type KeyShareRole string

const (
	User      KeyShareRole = "user"
	Validator KeyShareRole = "validator"
)

// String returns the string representation of KeyShareRole
func (rcv KeyShareRole) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(KeyShareRole)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for KeyShareRole.
func (rcv *KeyShareRole) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "user":
		*rcv = User
	case "validator":
		*rcv = Validator
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid KeyShareRole`, str)
	}
	return nil
}
