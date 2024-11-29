// Code generated from Pkl module `sonr.motr.ORM`. DO NOT EDIT.
package ucanresourcetype

import (
	"encoding"
	"fmt"
)

type UCANResourceType string

const (
	RESACCOUNT     UCANResourceType = "RES_ACCOUNT"
	RESTRANSACTION UCANResourceType = "RES_TRANSACTION"
	RESPOLICY      UCANResourceType = "RES_POLICY"
	RESRECOVERY    UCANResourceType = "RES_RECOVERY"
)

// String returns the string representation of UCANResourceType
func (rcv UCANResourceType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(UCANResourceType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for UCANResourceType.
func (rcv *UCANResourceType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "RES_ACCOUNT":
		*rcv = RESACCOUNT
	case "RES_TRANSACTION":
		*rcv = RESTRANSACTION
	case "RES_POLICY":
		*rcv = RESPOLICY
	case "RES_RECOVERY":
		*rcv = RESRECOVERY
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid UCANResourceType`, str)
	}
	return nil
}
