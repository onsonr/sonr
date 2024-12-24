// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package resinterchain

import (
	"encoding"
	"fmt"
)

type ResInterchain string

const (
	ChannnelPort  ResInterchain = "channnel/port"
	ChainId       ResInterchain = "chain/id"
	ChainName     ResInterchain = "chain/name"
	AccHost       ResInterchain = "acc/host"
	AccController ResInterchain = "acc/controller"
)

// String returns the string representation of ResInterchain
func (rcv ResInterchain) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(ResInterchain)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for ResInterchain.
func (rcv *ResInterchain) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "channnel/port":
		*rcv = ChannnelPort
	case "chain/id":
		*rcv = ChainId
	case "chain/name":
		*rcv = ChainName
	case "acc/host":
		*rcv = AccHost
	case "acc/controller":
		*rcv = AccController
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid ResInterchain`, str)
	}
	return nil
}
