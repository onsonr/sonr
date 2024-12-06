package ucan

import (
	"fmt"

	"github.com/onsonr/sonr/crypto/ucan/attns/capability"
	"github.com/onsonr/sonr/crypto/ucan/attns/policytype"
	"github.com/onsonr/sonr/crypto/ucan/attns/resourcetype"
)

var EmptyAttenuation = Attenuation{
	Cap: Capability(nil),
	Rsc: Resource(nil),
}

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
	ResKeyShare    = resourcetype.RESKEYSHARE

	// PolicyTypes
	PolicyThreshold = policytype.POLICYTHRESHOLD
	PolicyTimelock  = policytype.POLICYTIMELOCK
	PolicyWhitelist = policytype.POLICYWHITELIST
	PolicyKeyShare  = policytype.POLICYKEYGEN
)

// NewVaultResource creates a new resource identifier
func NewResource(resType resourcetype.ResourceType, path string) Resource {
	return NewStringLengthResource(string(resType), path)
}

// Permissions represents the type of attenuation
type Permissions string

const (
	// AccountPermissions represents the smart account attenuation
	AccountPermissions = Permissions("account")

	// ServicePermissions represents the service attenuation
	ServicePermissions = Permissions("service")

	// VaultPermissions represents the vault attenuation
	VaultPermissions = Permissions("vault")
)

// Cap returns the capability for the given AttenuationPreset
func (a Permissions) NewCap(c capability.Capability) Capability {
	return a.GetCapabilities().Cap(c.String())
}

// NestedCapabilities returns the nested capabilities for the given AttenuationPreset
func (a Permissions) GetCapabilities() NestedCapabilities {
	var caps []string
	switch a {
	case AccountPermissions:
		caps = SmartAccountCapabilities()
	case VaultPermissions:
		caps = VaultCapabilities()
	}
	return NewNestedCapabilities(caps...)
}

// Equals returns true if the given AttenuationPreset is equal to the receiver
func (a Permissions) Equals(b Permissions) bool {
	return a == b
}

// String returns the string representation of the AttenuationPreset
func (a Permissions) String() string {
	return string(a)
}

// GetConstructor returns the AttenuationConstructorFunc for a Permission
func (a Permissions) GetConstructor() AttenuationConstructorFunc {
	return NewAttenuationFromPreset(a)
}

// NewAttenuationFromPreset creates an AttenuationConstructorFunc for the given preset
func NewAttenuationFromPreset(preset Permissions) AttenuationConstructorFunc {
	return func(v map[string]interface{}) (Attenuation, error) {
		// Extract capability and resource from map
		capStr, ok := v["cap"].(string)
		if !ok {
			return EmptyAttenuation, fmt.Errorf("missing or invalid capability in attenuation data")
		}

		resType, ok := v["type"].(string)
		if !ok {
			return EmptyAttenuation, fmt.Errorf("missing or invalid resource type in attenuation data")
		}

		path, ok := v["path"].(string)
		if !ok {
			path = "/" // Default path if not specified
		}

		// Create capability from preset
		cap := preset.NewCap(capability.Capability(capStr))
		if cap == nil {
			return EmptyAttenuation, fmt.Errorf("invalid capability %s for preset %s", capStr, preset)
		}

		// Create resource
		resource := NewResource(resourcetype.ResourceType(resType), path)

		return Attenuation{
			Cap: cap,
			Rsc: resource,
		}, nil
	}
}

// GetPresetConstructor returns the appropriate AttenuationConstructorFunc for a given type
func GetPresetConstructor(attType string) (AttenuationConstructorFunc, error) {
	preset := Permissions(attType)
	switch preset {
	case AccountPermissions, ServicePermissions, VaultPermissions:
		return NewAttenuationFromPreset(preset), nil
	default:
		return nil, fmt.Errorf("unknown attenuation preset: %s", attType)
	}
}

// ParseAttenuationData parses raw attenuation data into a structured format
func ParseAttenuationData(data map[string]interface{}) (Permissions, map[string]interface{}, error) {
	typeRaw, ok := data["preset"]
	if !ok {
		return "", nil, fmt.Errorf("missing preset type in attenuation data")
	}

	presetType, ok := typeRaw.(string)
	if !ok {
		return "", nil, fmt.Errorf("invalid preset type format")
	}

	return Permissions(presetType), data, nil
}
