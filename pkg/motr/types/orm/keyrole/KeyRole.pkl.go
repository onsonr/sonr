// Code generated from Pkl module `orm`. DO NOT EDIT.
package keyrole

import (
	"encoding"
	"fmt"
)

type KeyRole string

const (
	Authentication KeyRole = "authentication"
	Assertion      KeyRole = "assertion"
	Delegation     KeyRole = "delegation"
	Invocation     KeyRole = "invocation"
)

// String returns the string representation of KeyRole
func (rcv KeyRole) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(KeyRole)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for KeyRole.
func (rcv *KeyRole) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "authentication":
		*rcv = Authentication
	case "assertion":
		*rcv = Assertion
	case "delegation":
		*rcv = Delegation
	case "invocation":
		*rcv = Invocation
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid KeyRole`, str)
	}
	return nil
}
