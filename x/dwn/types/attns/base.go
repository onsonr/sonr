package attns

import (
	"github.com/onsonr/sonr/x/dwn/types/attns/capability"
	"github.com/onsonr/sonr/x/dwn/types/attns/policytype"
	"github.com/onsonr/sonr/x/dwn/types/attns/resourcetype"
	"github.com/ucan-wg/go-ucan"
)

const (
	CapOwner        = capability.CAPOWNER
	CapOperator     = capability.CAPOPERATOR
	CapObserver     = capability.CAPOBSERVER
	CapExecute      = capability.CAPEXECUTE
	CapPropose      = capability.CAPPROPOSE
	CapSign         = capability.CAPSIGN
	CapSetPolicy    = capability.CAPSETPOLICY
	CapSetThreshold = capability.CAPSETTHRESHOLD
	CapRecover      = capability.CAPRECOVER
	CapSocial       = capability.CAPSOCIAL

	ResAccount     = resourcetype.RESACCOUNT
	ResTransaction = resourcetype.RESTRANSACTION
	ResPolicy      = resourcetype.RESPOLICY
	ResRecovery    = resourcetype.RESRECOVERY
	ResVault       = resourcetype.RESVAULT

	PolicyThreshold = policytype.POLICYTHRESHOLD
	PolicyTimelock  = policytype.POLICYTIMELOCK
	PolicyWhitelist = policytype.POLICYWHITELIST
)

// NewVaultResource creates a new resource identifier
func NewResource(resType resourcetype.ResourceType, path string) ucan.Resource {
	return ucan.NewStringLengthResource(string(resType), path)
}

// Attenuation represents the type of attenuation
type Attenuation string

const (
	// AttentuationSmartAccount represents the smart account attenuation
	AttentuationSmartAccount = Attenuation("smart_account")

	// AttentuationVault represents the vault attenuation
	AttentuationVault = Attenuation("vault")
)

// NestedCapabilities returns the nested capabilities for the given Attenuation
func (a Attenuation) GetCapabilities() ucan.NestedCapabilities {
	var caps []string
	switch a {
	case AttentuationSmartAccount:
		caps = baseSmartAccountCapabilities()
	case AttentuationVault:
		caps = baseVaultCapabilities()
	}
	return ucan.NewNestedCapabilities(caps...)
}

// Equals returns true if the given Attenuation is equal to the receiver
func (a Attenuation) Equals(b Attenuation) bool {
	return a == b
}

// String returns the string representation of the Attenuation
func (a Attenuation) String() string {
	return string(a)
}

// SmartAccountCapabilities defines the capability hierarchy
func baseSmartAccountCapabilities() []string {
	return []string{
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
	}
}

// VaultCapabilities defines the capability hierarchy
func baseVaultCapabilities() []string {
	return []string{
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
	}
}
