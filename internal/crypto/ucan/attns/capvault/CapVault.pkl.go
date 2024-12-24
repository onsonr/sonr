// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package capvault

import (
	"encoding"
	"fmt"
)

type CapVault string

const (
	CrudAsset      CapVault = "crud/asset"
	CrudAuthzgrant CapVault = "crud/authzgrant"
	CrudProfile    CapVault = "crud/profile"
	CrudRecord     CapVault = "crud/record"
	UseRecovery    CapVault = "use/recovery"
	UseSync        CapVault = "use/sync"
	UseSigner      CapVault = "use/signer"
)

// String returns the string representation of CapVault
func (rcv CapVault) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(CapVault)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for CapVault.
func (rcv *CapVault) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "crud/asset":
		*rcv = CrudAsset
	case "crud/authzgrant":
		*rcv = CrudAuthzgrant
	case "crud/profile":
		*rcv = CrudProfile
	case "crud/record":
		*rcv = CrudRecord
	case "use/recovery":
		*rcv = UseRecovery
	case "use/sync":
		*rcv = UseSync
	case "use/signer":
		*rcv = UseSigner
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid CapVault`, str)
	}
	return nil
}
