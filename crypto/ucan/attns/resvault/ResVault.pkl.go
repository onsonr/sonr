// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package resvault

import (
	"encoding"
	"fmt"
)

type ResVault string

const (
	KsUser    ResVault = "ks/user"
	KsVal     ResVault = "ks/val"
	LocCid    ResVault = "loc/cid"
	LocEntity ResVault = "loc/entity"
	LocIpns   ResVault = "loc/ipns"
	LocSonr   ResVault = "loc/sonr"
)

// String returns the string representation of ResVault
func (rcv ResVault) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ResVault)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ResVault.
func (rcv *ResVault) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "ks/user":
		*rcv = KsUser
	case "ks/val":
		*rcv = KsVal
	case "loc/cid":
		*rcv = LocCid
	case "loc/entity":
		*rcv = LocEntity
	case "loc/ipns":
		*rcv = LocIpns
	case "loc/sonr":
		*rcv = LocSonr
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ResVault`, str)
	}
	return nil
}
