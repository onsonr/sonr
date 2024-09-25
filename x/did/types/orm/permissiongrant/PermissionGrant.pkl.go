// Code generated from Pkl module `models`. DO NOT EDIT.
package permissiongrant

import (
	"encoding"
	"fmt"
)

type PermissionGrant string

const (
	None      PermissionGrant = "none"
	Read      PermissionGrant = "read"
	Write     PermissionGrant = "write"
	Verify    PermissionGrant = "verify"
	Broadcast PermissionGrant = "broadcast"
	Admin     PermissionGrant = "admin"
)

// String returns the string representation of PermissionGrant
func (rcv PermissionGrant) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(PermissionGrant)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for PermissionGrant.
func (rcv *PermissionGrant) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "none":
		*rcv = None
	case "read":
		*rcv = Read
	case "write":
		*rcv = Write
	case "verify":
		*rcv = Verify
	case "broadcast":
		*rcv = Broadcast
	case "admin":
		*rcv = Admin
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid PermissionGrant`, str)
	}
	return nil
}
