package attns

import (
	"fmt"

	"github.com/onsonr/sonr/x/dwn/types/attns/capability"
	"github.com/onsonr/sonr/x/dwn/types/attns/resourcetype"
	"github.com/ucan-wg/go-ucan"
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
	)
}

// CreateSmartAccountAttenuations creates default attenuations for a smart account
func CreateSmartAccountAttenuations(
	caps ucan.NestedCapabilities,
	accountAddr string,
) ucan.Attenuations {
	return ucan.Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(capability.CAPOWNER.String()), Rsc: NewResource(resourcetype.RESACCOUNT, accountAddr)},

		// Operation capabilities
		{caps.Cap(capability.CAPEXECUTE.String()), NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(capability.CAPPROPOSE.String()), NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(capability.CAPSIGN.String()), NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{caps.Cap(capability.CAPSETPOLICY.String()), NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:*", accountAddr))},
		{caps.Cap(capability.CAPSETTHRESHOLD.String()), NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// CreateSmartAccountPolicyAttenuation creates attenuations for policy management
func CreateSmartAccountPolicyAttenuation(
	caps ucan.NestedCapabilities,
	accountAddr string,
	policyType string,
) ucan.Attenuations {
	return ucan.Attenuations{
		{
			caps.Cap(capability.CAPSETPOLICY.String()),
			NewResource(
				resourcetype.RESPOLICY,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}
