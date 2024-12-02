package attns

import (
	"fmt"

	"github.com/onsonr/sonr/x/dwn/types/attns/capability"
	"github.com/onsonr/sonr/x/dwn/types/attns/policytype"
	"github.com/onsonr/sonr/x/dwn/types/attns/resourcetype"
	"github.com/ucan-wg/go-ucan"
)

// CreateVaultAttenuations creates default attenuations for a smart account
func CreateVaultAttenuations(
	accountAddr string,
) ucan.Attenuations {
	caps := AttentuationVault.GetCapabilities()
	return ucan.Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(capability.CAPOWNER.String()), Rsc: NewResource(resourcetype.RESACCOUNT, accountAddr)},

		// Operation capabilities
		{Cap: caps.Cap(capability.CAPEXECUTE.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPPROPOSE.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSIGN.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{Cap: caps.Cap(capability.CAPSETPOLICY.String()), Rsc: NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSETTHRESHOLD.String()), Rsc: NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// CreateVaultPolicyAttenuation creates attenuations for policy management
func CreateVaultPolicyAttenuation(
	accountAddr string,
	policyType policytype.PolicyType,
) ucan.Attenuations {
	caps := AttentuationVault.GetCapabilities()
	return ucan.Attenuations{
		{
			Cap: caps.Cap(capability.CAPSETPOLICY.String()),
			Rsc: NewResource(
				resourcetype.RESPOLICY,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}
