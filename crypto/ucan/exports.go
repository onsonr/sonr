package ucan

import (
	"github.com/onsonr/sonr/crypto/ucan/attns/capability"
	"github.com/onsonr/sonr/crypto/ucan/attns/policytype"
	"github.com/onsonr/sonr/crypto/ucan/attns/resourcetype"
)

const (
	// Owner
	CapOwner    = capability.CAPOWNER
	CapOperator = capability.CAPOPERATOR
	CapObserver = capability.CAPOBSERVER

	// Auth
	CapAuthenticate = capability.CAPAUTHENTICATE
	CapAuthorize    = capability.CAPAUTHORIZE
	CapDelegate     = capability.CAPDELEGATE
	CapInvoke       = capability.CAPINVOKE
	CapExecute      = capability.CAPEXECUTE
	CapPropose      = capability.CAPPROPOSE
	CapSign         = capability.CAPSIGN
	CapSetPolicy    = capability.CAPSETPOLICY
	CapSetThreshold = capability.CAPSETTHRESHOLD
	CapRecover      = capability.CAPRECOVER
	CapSocial       = capability.CAPSOCIAL
	CapResolver     = capability.CAPRESOLVER
	CapProducer     = capability.CAPPRODUCER

	// Resources
	ResAccount     = resourcetype.RESACCOUNT
	ResTransaction = resourcetype.RESTRANSACTION
	ResPolicy      = resourcetype.RESPOLICY
	ResRecovery    = resourcetype.RESRECOVERY
	ResVault       = resourcetype.RESVAULT
	ResIPFS        = resourcetype.RESIPFS
	ResIPNS        = resourcetype.RESIPNS

	// PolicyTypes
	PolicyThreshold = policytype.POLICYTHRESHOLD
	PolicyTimelock  = policytype.POLICYTIMELOCK
	PolicyWhitelist = policytype.POLICYWHITELIST
)

// NewVaultResource creates a new resource identifier
func NewResource(resType resourcetype.ResourceType, path string) Resource {
	return NewStringLengthResource(string(resType), path)
}

// AttenuationPreset represents the type of attenuation
type AttenuationPreset string

const (
	// AttentuationSmartAccount represents the smart account attenuation
	AttentuationSmartAccount = AttenuationPreset("smart_account")

	// AttentuationService represents the service attenuation
	AttentuationService = AttenuationPreset("service")

	// AttentuationVault represents the vault attenuation
	AttentuationVault = AttenuationPreset("vault")
)

// Cap returns the capability for the given AttenuationPreset
func (a AttenuationPreset) NewCap(c capability.Capability) Capability {
	return a.GetCapabilities().Cap(c.String())
}

// NestedCapabilities returns the nested capabilities for the given AttenuationPreset
func (a AttenuationPreset) GetCapabilities() NestedCapabilities {
	var caps []string
	switch a {
	case AttentuationSmartAccount:
		caps = SmartAccountCapabilities()
	case AttentuationVault:
		caps = VaultCapabilities()
	}
	return NewNestedCapabilities(caps...)
}

// Equals returns true if the given AttenuationPreset is equal to the receiver
func (a AttenuationPreset) Equals(b AttenuationPreset) bool {
	return a == b
}

// String returns the string representation of the AttenuationPreset
func (a AttenuationPreset) String() string {
	return string(a)
}
