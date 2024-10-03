// Code generated from Pkl module `models`. DO NOT EDIT.
package inputtype

import (
	"encoding"
	"fmt"
)

type InputType string

const (
	Text       InputType = "text"
	Password   InputType = "password"
	Email      InputType = "email"
	Credential InputType = "credential"
	File       InputType = "file"
)

// String returns the string representation of InputType
func (rcv InputType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(InputType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for InputType.
func (rcv *InputType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "text":
		*rcv = Text
	case "password":
		*rcv = Password
	case "email":
		*rcv = Email
	case "credential":
		*rcv = Credential
	case "file":
		*rcv = File
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid InputType`, str)
	}
	return nil
}
