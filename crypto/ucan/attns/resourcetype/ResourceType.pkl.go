// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package resourcetype

import (
	"encoding"
	"fmt"
)

type ResourceType string

const (
	RESACCOUNT     ResourceType = "RES_ACCOUNT"
	RESTRANSACTION ResourceType = "RES_TRANSACTION"
	RESPOLICY      ResourceType = "RES_POLICY"
	RESRECOVERY    ResourceType = "RES_RECOVERY"
	RESVAULT       ResourceType = "RES_VAULT"
	RESIPFS        ResourceType = "RES_IPFS"
	RESIPNS        ResourceType = "RES_IPNS"
	RESKEYSHARE    ResourceType = "RES_KEYSHARE"
)

// String returns the string representation of ResourceType
func (rcv ResourceType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ResourceType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ResourceType.
func (rcv *ResourceType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "RES_ACCOUNT":
		*rcv = RESACCOUNT
	case "RES_TRANSACTION":
		*rcv = RESTRANSACTION
	case "RES_POLICY":
		*rcv = RESPOLICY
	case "RES_RECOVERY":
		*rcv = RESRECOVERY
	case "RES_VAULT":
		*rcv = RESVAULT
	case "RES_IPFS":
		*rcv = RESIPFS
	case "RES_IPNS":
		*rcv = RESIPNS
	case "RES_KEYSHARE":
		*rcv = RESKEYSHARE
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ResourceType`, str)
	}
	return nil
}
