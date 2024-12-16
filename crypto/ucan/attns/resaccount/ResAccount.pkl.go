// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package resaccount

import (
	"encoding"
	"fmt"
)

type ResAccount string

const (
	AccSequence ResAccount = "acc/sequence"
	AccNumber   ResAccount = "acc/number"
	ChainId     ResAccount = "chain/id"
	AssetCode   ResAccount = "asset/code"
	AuthzGrant  ResAccount = "authz/grant"
)

// String returns the string representation of ResAccount
func (rcv ResAccount) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ResAccount)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ResAccount.
func (rcv *ResAccount) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "acc/sequence":
		*rcv = AccSequence
	case "acc/number":
		*rcv = AccNumber
	case "chain/id":
		*rcv = ChainId
	case "asset/code":
		*rcv = AssetCode
	case "authz/grant":
		*rcv = AuthzGrant
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ResAccount`, str)
	}
	return nil
}
