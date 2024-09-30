// Code generated from Pkl module `orm`. DO NOT EDIT.
package permissionscope

import (
	"encoding"
	"fmt"
)

type PermissionScope string

const (
	Profile      PermissionScope = "profile"
	Metadata     PermissionScope = "metadata"
	Permissions  PermissionScope = "permissions"
	Wallets      PermissionScope = "wallets"
	Transactions PermissionScope = "transactions"
	User         PermissionScope = "user"
	Validator    PermissionScope = "validator"
)

// String returns the string representation of PermissionScope
func (rcv PermissionScope) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(PermissionScope)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for PermissionScope.
func (rcv *PermissionScope) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "profile":
		*rcv = Profile
	case "metadata":
		*rcv = Metadata
	case "permissions":
		*rcv = Permissions
	case "wallets":
		*rcv = Wallets
	case "transactions":
		*rcv = Transactions
	case "user":
		*rcv = User
	case "validator":
		*rcv = Validator
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid PermissionScope`, str)
	}
	return nil
}
