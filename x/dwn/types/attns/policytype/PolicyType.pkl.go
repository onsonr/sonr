// Code generated from Pkl module `sonr.motr.ATN`. DO NOT EDIT.
package policytype

import (
	"encoding"
	"fmt"
)

type PolicyType string

const (
	POLICYTHRESHOLD PolicyType = "POLICY_THRESHOLD"
	POLICYTIMELOCK  PolicyType = "POLICY_TIMELOCK"
	POLICYWHITELIST PolicyType = "POLICY_WHITELIST"
)

// String returns the string representation of PolicyType
func (rcv PolicyType) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(PolicyType)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for PolicyType.
func (rcv *PolicyType) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "POLICY_THRESHOLD":
		*rcv = POLICYTHRESHOLD
	case "POLICY_TIMELOCK":
		*rcv = POLICYTIMELOCK
	case "POLICY_WHITELIST":
		*rcv = POLICYWHITELIST
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid PolicyType`, str)
	}
	return nil
}
