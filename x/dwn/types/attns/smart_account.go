package attns

import (
	"fmt"

	"github.com/onsonr/sonr/x/dwn/types/attns/capability"
	"github.com/onsonr/sonr/x/dwn/types/attns/policytype"
	"github.com/ucan-wg/go-ucan"
)

// CreateSmartAccountAttenuations creates default attenuations for a smart account
func CreateSmartAccountAttenuations(
	accountAddr string,
) ucan.Attenuations {
	caps := AttentuationSmartAccount.GetCapabilities()
	return ucan.Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(CapOwner.String()), Rsc: NewResource(ResAccount, accountAddr)},

		// Operation capabilities
		{Cap: caps.Cap(capability.CAPEXECUTE.String()), Rsc: NewResource(ResTransaction, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPPROPOSE.String()), Rsc: NewResource(ResTransaction, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSIGN.String()), Rsc: NewResource(ResTransaction, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{Cap: caps.Cap(capability.CAPSETPOLICY.String()), Rsc: NewResource(ResPolicy, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSETTHRESHOLD.String()), Rsc: NewResource(ResPolicy, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// CreateSmartAccountPolicyAttenuation creates attenuations for policy management
func CreateSmartAccountPolicyAttenuation(
	accountAddr string,
	policyType policytype.PolicyType,
) ucan.Attenuations {
	caps := AttentuationSmartAccount.GetCapabilities()
	return ucan.Attenuations{
		{
			Cap: caps.Cap(capability.CAPSETPOLICY.String()),
			Rsc: NewResource(
				ResPolicy,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}
