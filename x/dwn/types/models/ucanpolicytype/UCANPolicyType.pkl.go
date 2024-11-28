// Code generated from Pkl module `sonr.motr.ORM`. DO NOT EDIT.
package ucanpolicytype

import (
	"encoding"
	"fmt"
)

type UCANPolicyType string

const (
	POLICYTHRESHOLD UCANPolicyType = "POLICY_THRESHOLD"
	POLICYTIMELOCK  UCANPolicyType = "POLICY_TIMELOCK"
	POLICYWHITELIST UCANPolicyType = "POLICY_WHITELIST"
)

// String returns the string representation of UCANPolicyType
func (rcv UCANPolicyType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(UCANPolicyType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for UCANPolicyType.
func (rcv *UCANPolicyType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "POLICY_THRESHOLD":
		*rcv = POLICYTHRESHOLD
	case "POLICY_TIMELOCK":
		*rcv = POLICYTIMELOCK
	case "POLICY_WHITELIST":
		*rcv = POLICYWHITELIST
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid UCANPolicyType`, str)
	}
	return nil
}
