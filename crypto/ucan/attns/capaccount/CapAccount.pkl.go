// Code generated from Pkl module `sonr.orm.UCAN`. DO NOT EDIT.
package capaccount

import (
	"encoding"
	"fmt"
)

type CapAccount string

const (
	ExecBroadcast CapAccount = "exec/broadcast"
	ExecQuery     CapAccount = "exec/query"
	ExecSimulate  CapAccount = "exec/simulate"
	ExecVote      CapAccount = "exec/vote"
	ExecDelegate  CapAccount = "exec/delegate"
	ExecInvoke    CapAccount = "exec/invoke"
	ExecSend      CapAccount = "exec/send"
)

// String returns the string representation of CapAccount
func (rcv CapAccount) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(CapAccount)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for CapAccount.
func (rcv *CapAccount) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "exec/broadcast":
		*rcv = ExecBroadcast
	case "exec/query":
		*rcv = ExecQuery
	case "exec/simulate":
		*rcv = ExecSimulate
	case "exec/vote":
		*rcv = ExecVote
	case "exec/delegate":
		*rcv = ExecDelegate
	case "exec/invoke":
		*rcv = ExecInvoke
	case "exec/send":
		*rcv = ExecSend
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid CapAccount`, str)
	}
	return nil
}
