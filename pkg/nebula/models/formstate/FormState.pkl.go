// Code generated from Pkl module `models`. DO NOT EDIT.
package formstate

import (
	"encoding"
	"fmt"
)

type FormState string

const (
	Initial FormState = "initial"
	Error   FormState = "error"
	Success FormState = "success"
	Warning FormState = "warning"
)

// String returns the string representation of FormState
func (rcv FormState) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(FormState)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for FormState.
func (rcv *FormState) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "initial":
		*rcv = Initial
	case "error":
		*rcv = Error
	case "success":
		*rcv = Success
	case "warning":
		*rcv = Warning
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid FormState`, str)
	}
	return nil
}
