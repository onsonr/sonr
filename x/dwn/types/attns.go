package types

import (
	"fmt"

	"github.com/onsonr/sonr/x/dwn/types/models/ucancapability"
	"github.com/onsonr/sonr/x/dwn/types/models/ucanresourcetype"
	"github.com/ucan-wg/go-ucan"
)

// SmartAccountCapabilities defines the capability hierarchy
func NewSmartAccountCapabilities() ucan.NestedCapabilities {
	return ucan.NewNestedCapabilities(
		ucancapability.CAPOWNER.String(),
		ucancapability.CAPOPERATOR.String(),
		ucancapability.CAPOBSERVER.String(),
		ucancapability.CAPEXECUTE.String(),
		ucancapability.CAPPROPOSE.String(),
		ucancapability.CAPSIGN.String(),
		ucancapability.CAPSETPOLICY.String(),
		ucancapability.CAPSETTHRESHOLD.String(),
		ucancapability.CAPRECOVER.String(),
		ucancapability.CAPSOCIAL.String(),
		ucancapability.CAPVAULT.String(),
	)
}

// NewSmartAccountResource creates a new resource identifier
func NewSmartAccountResource(resType ucanresourcetype.UCANResourceType, path string) ucan.Resource {
	return ucan.NewStringLengthResource(string(resType), path)
}

// CreateSmartAccountAttenuations creates default attenuations for a smart account
func CreateSmartAccountAttenuations(
	caps ucan.NestedCapabilities,
	accountAddr string,
) ucan.Attenuations {
	return ucan.Attenuations{
		// Owner capabilities
		{caps.Cap(ucancapability.CAPOWNER.String()), NewSmartAccountResource(ResourceType(ucanresourcetype.RESACCOUNT.String()), accountAddr)},

		// Operation capabilities
		{caps.Cap(ucancapability.CAPEXECUTE.String()), NewSmartAccountResource(ucanresourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(ucancapability.CAPPROPOSE.String()), NewSmartAccountResource(ucanresourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(ucancapability.CAPSIGN.String()), NewSmartAccountResource(ucanresourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{caps.Cap(ucancapability.CAPSETPOLICY.String()), NewSmartAccountResource(ucanresourcetype.RESPOLICY, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(ucancapability.CAPSETTHRESHOLD.String()), NewSmartAccountResource(ucanresourcetype.RESPOLICY, fmt.Sprintf("%s:threshold", accountAddr))},
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
			caps.Cap(ucancapability.CAPSETPOLICY.String()),
			NewSmartAccountResource(
				ucanresourcetype.RESPOLICY,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}
