package attns

import (
	"fmt"

	"github.com/onsonr/sonr/x/dwn/types/attns/capability"
	"github.com/onsonr/sonr/x/dwn/types/attns/resourcetype"
	"github.com/ucan-wg/go-ucan"
)

const (
	CapOwner        = capability.CAPOWNER
	CapOperator     = capability.CAPOPERATOR
	CapServer       = capability.CAPOBSERVER
	CapExecute      = capability.CAPEXECUTE
	CapPropose      = capability.CAPPROPOSE
	CapSign         = capability.CAPSIGN
	CapSetPolicy    = capability.CAPSETPOLICY
	CapSetThreshold = capability.CAPSETTHRESHOLD
	CapRecover      = capability.CAPRECOVER
	CapSocial       = capability.CAPSOCIAL
	CapVault        = capability.CAPVAULT

	ResAccount     = resourcetype.RESACCOUNT
	ResTransaction = resourcetype.RESTRANSACTION
	ResPolicy      = resourcetype.RESPOLICY
	ResRecovery    = resourcetype.RESRECOVERY

	PolicyThreshold = resourcetype.RESPOLICY
	PolicySet       = resourcetype.RESPOLICY
	PolicyGet       = resourcetype.RESPOLICY
	PolicyDelete    = resourcetype.RESPOLICY
)

// SmartAccountCapabilities defines the capability hierarchy
func NewSmartAccountCapabilities() ucan.NestedCapabilities {
	return ucan.NewNestedCapabilities(
		capability.CAPOWNER.String(),
		capability.CAPOPERATOR.String(),
		capability.CAPOBSERVER.String(),
		capability.CAPEXECUTE.String(),
		capability.CAPPROPOSE.String(),
		capability.CAPSIGN.String(),
		capability.CAPSETPOLICY.String(),
		capability.CAPSETTHRESHOLD.String(),
		capability.CAPRECOVER.String(),
		capability.CAPSOCIAL.String(),
		capability.CAPVAULT.String(),
	)
}

// NewSmartAccountResource creates a new resource identifier
func NewSmartAccountResource(resType resourcetype.ResourceType, path string) ucan.Resource {
	return ucan.NewStringLengthResource(string(resType), path)
}

// CreateSmartAccountAttenuations creates default attenuations for a smart account
func CreateSmartAccountAttenuations(
	caps ucan.NestedCapabilities,
	accountAddr string,
) ucan.Attenuations {
	return ucan.Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(capability.CAPOWNER.String()), Rsc: NewSmartAccountResource(resourcetype.RESACCOUNT, accountAddr)},

		// Operation capabilities
		{caps.Cap(capability.CAPEXECUTE.String()), NewSmartAccountResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(capability.CAPPROPOSE.String()), NewSmartAccountResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(capability.CAPSIGN.String()), NewSmartAccountResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{caps.Cap(capability.CAPSETPOLICY.String()), NewSmartAccountResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(capability.CAPSETTHRESHOLD.String()), NewSmartAccountResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// CreatePolicyAttenuation creates attenuations for policy management
func CreatePolicyAttenuation(
	caps ucan.NestedCapabilities,
	accountAddr string,
	policyType string,
) ucan.Attenuations {
	return ucan.Attenuations{
		{
			caps.Cap(capability.CAPSETPOLICY.String()),
			NewSmartAccountResource(
				resourcetype.RESPOLICY,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}
