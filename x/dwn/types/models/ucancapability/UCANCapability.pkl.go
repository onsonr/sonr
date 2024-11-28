// Code generated from Pkl module `sonr.motr.ORM`. DO NOT EDIT.
package ucancapability

import (
	"encoding"
	"fmt"
)

type UCANCapability string

const (
	CAPOWNER        UCANCapability = "CAP_OWNER"
	CAPOPERATOR     UCANCapability = "CAP_OPERATOR"
	CAPOBSERVER     UCANCapability = "CAP_OBSERVER"
	CAPEXECUTE      UCANCapability = "CAP_EXECUTE"
	CAPPROPOSE      UCANCapability = "CAP_PROPOSE"
	CAPSIGN         UCANCapability = "CAP_SIGN"
	CAPSETPOLICY    UCANCapability = "CAP_SET_POLICY"
	CAPSETTHRESHOLD UCANCapability = "CAP_SET_THRESHOLD"
	CAPRECOVER      UCANCapability = "CAP_RECOVER"
	CAPSOCIAL       UCANCapability = "CAP_SOCIAL"
	CAPVAULT        UCANCapability = "CAP_VAULT"
)

// String returns the string representation of UCANCapability
func (rcv UCANCapability) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(UCANCapability)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for UCANCapability.
func (rcv *UCANCapability) UnmarshalBinary(data []byte) error {
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
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid UCANCapability`, str)
	}
	return nil
}
