package mpc

import (
	"fmt"

	"github.com/ucan-wg/go-ucan"
)

// Capability hierarchy for smart account operations
// ----------------------------------------------
// OWNER
//   └─ OPERATOR
//        ├─ EXECUTE
//        ├─ PROPOSE
//        └─ SIGN
//   └─ SET_POLICY
//        └─ SET_THRESHOLD
//   └─ RECOVER
//        └─ SOCIAL

// Define capability hierarchy for smart account operations
const (
	// Root capabilities
	CAP_OWNER    = "OWNER"    // Full account control
	CAP_OPERATOR = "OPERATOR" // Can perform operations
	CAP_OBSERVER = "OBSERVER" // Can view account state

	// Operation capabilities
	CAP_EXECUTE = "EXECUTE" // Can execute transactions
	CAP_PROPOSE = "PROPOSE" // Can propose transactions
	CAP_SIGN    = "SIGN"    // Can sign transactions

	// Policy capabilities
	CAP_SET_POLICY    = "SET_POLICY"    // Can modify account policies
	CAP_SET_THRESHOLD = "SET_THRESHOLD" // Can modify signing threshold

	// Recovery capabilities
	CAP_RECOVER = "RECOVER" // Can initiate recovery
	CAP_SOCIAL  = "SOCIAL"  // Can act as social recovery
)

// SmartAccountCapabilities defines the capability hierarchy
func NewSmartAccountCapabilities() ucan.NestedCapabilities {
	return ucan.NewNestedCapabilities(
		CAP_OWNER,
		CAP_OPERATOR,
		CAP_OBSERVER,
		CAP_EXECUTE,
		CAP_PROPOSE,
		CAP_SIGN,
		CAP_SET_POLICY,
		CAP_SET_THRESHOLD,
		CAP_RECOVER,
		CAP_SOCIAL,
	)
}

// Resource types for smart account operations
type ResourceType string

const (
	RES_ACCOUNT     = "account"
	RES_TRANSACTION = "tx"
	RES_POLICY      = "policy"
	RES_RECOVERY    = "recovery"
)

// NewSmartAccountResource creates a new resource identifier
func NewSmartAccountResource(resType ResourceType, path string) ucan.Resource {
	return ucan.NewStringLengthResource(string(resType), path)
}

// CreateSmartAccountAttenuations creates default attenuations for a smart account
func CreateSmartAccountAttenuations(
	caps ucan.NestedCapabilities,
	accountAddr string,
) ucan.Attenuations {
	return ucan.Attenuations{
		// Owner capabilities
		{caps.Cap(CAP_OWNER), NewSmartAccountResource(RES_ACCOUNT, accountAddr)},

		// Operation capabilities
		{caps.Cap(CAP_EXECUTE), NewSmartAccountResource(RES_TRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(CAP_PROPOSE), NewSmartAccountResource(RES_TRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(CAP_SIGN), NewSmartAccountResource(RES_TRANSACTION, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{caps.Cap(CAP_SET_POLICY), NewSmartAccountResource(RES_POLICY, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(CAP_SET_THRESHOLD), NewSmartAccountResource(RES_POLICY, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// Policy represents smart account execution policies
type PolicyType string

const (
	POLICY_THRESHOLD = "threshold"
	POLICY_TIMELOCK  = "timelock"
	POLICY_WHITELIST = "whitelist"
)

// CreatePolicyAttenuation creates attenuations for policy management
func CreatePolicyAttenuation(
	caps ucan.NestedCapabilities,
	accountAddr string,
	policyType PolicyType,
) ucan.Attenuations {
	return ucan.Attenuations{
		{
			caps.Cap(CAP_SET_POLICY),
			NewSmartAccountResource(
				RES_POLICY,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}
