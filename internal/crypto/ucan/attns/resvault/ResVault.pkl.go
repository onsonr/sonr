// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package resvault

import (
	"encoding"
	"fmt"
)

type ResVault string

const (
	KsEnclave ResVault = "ks/enclave"
	LocCid    ResVault = "loc/cid"
	LocEntity ResVault = "loc/entity"
	LocIpns   ResVault = "loc/ipns"
	AddrSonr  ResVault = "addr/sonr"
	ChainCode ResVault = "chain/code"
)

// String returns the string representation of ResVault
func (rcv ResVault) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ResVault)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ResVault.
func (rcv *ResVault) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "ks/enclave":
		*rcv = KsEnclave
	case "loc/cid":
		*rcv = LocCid
	case "loc/entity":
		*rcv = LocEntity
	case "loc/ipns":
		*rcv = LocIpns
	case "addr/sonr":
		*rcv = AddrSonr
	case "chain/code":
		*rcv = ChainCode
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ResVault`, str)
	}
	return nil
}
