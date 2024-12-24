// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package capinterchain

import (
	"encoding"
	"fmt"
)

type CapInterchain string

const (
	TransferSwap   CapInterchain = "transfer/swap"
	TransferSend   CapInterchain = "transfer/send"
	TransferAtomic CapInterchain = "transfer/atomic"
	TransferBatch  CapInterchain = "transfer/batch"
	TransferP2p    CapInterchain = "transfer/p2p"
)

// String returns the string representation of CapInterchain
func (rcv CapInterchain) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(CapInterchain)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for CapInterchain.
func (rcv *CapInterchain) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "transfer/swap":
		*rcv = TransferSwap
	case "transfer/send":
		*rcv = TransferSend
	case "transfer/atomic":
		*rcv = TransferAtomic
	case "transfer/batch":
		*rcv = TransferBatch
	case "transfer/p2p":
		*rcv = TransferP2p
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid CapInterchain`, str)
	}
	return nil
}
