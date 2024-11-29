// Code generated from Pkl module `sonr.motr.ATN`. DO NOT EDIT.
package capability

import (
	"encoding"
	"fmt"
)

type Capability string

const (
	CAPOWNER        Capability = "CAP_OWNER"
	CAPOPERATOR     Capability = "CAP_OPERATOR"
	CAPOBSERVER     Capability = "CAP_OBSERVER"
	CAPEXECUTE      Capability = "CAP_EXECUTE"
	CAPPROPOSE      Capability = "CAP_PROPOSE"
	CAPSIGN         Capability = "CAP_SIGN"
	CAPSETPOLICY    Capability = "CAP_SET_POLICY"
	CAPSETTHRESHOLD Capability = "CAP_SET_THRESHOLD"
	CAPRECOVER      Capability = "CAP_RECOVER"
	CAPSOCIAL       Capability = "CAP_SOCIAL"
	CAPVAULT        Capability = "CAP_VAULT"
	CAPVOTE         Capability = "CAP_VOTE"
)

// String returns the string representation of Capability
func (rcv Capability) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Capability)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Capability.
func (rcv *Capability) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "CAP_OWNER":
		*rcv = CAPOWNER
	case "CAP_OPERATOR":
		*rcv = CAPOPERATOR
	case "CAP_OBSERVER":
		*rcv = CAPOBSERVER
	case "CAP_EXECUTE":
		*rcv = CAPEXECUTE
	case "CAP_PROPOSE":
		*rcv = CAPPROPOSE
	case "CAP_SIGN":
		*rcv = CAPSIGN
	case "CAP_SET_POLICY":
		*rcv = CAPSETPOLICY
	case "CAP_SET_THRESHOLD":
		*rcv = CAPSETTHRESHOLD
	case "CAP_RECOVER":
		*rcv = CAPRECOVER
	case "CAP_SOCIAL":
		*rcv = CAPSOCIAL
	case "CAP_VAULT":
		*rcv = CAPVAULT
	case "CAP_VOTE":
		*rcv = CAPVOTE
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Capability`, str)
	}
	return nil
}
