package ucan

import (
	"fmt"
)

var EmptyAttenuation = Attenuation{
	Cap: Capability(nil),
	Rsc: Resource(nil),
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
func (a Permissions) NewCap(c string) Capability {
	return a.GetCapabilities().Cap(c)
}

// NestedCapabilities returns the nested capabilities for the given AttenuationPreset
func (a Permissions) GetCapabilities() NestedCapabilities {
	var caps []string
	switch a {
	case AccountPermissions:
		// caps = SmartAccountCapabilities()
	case VaultPermissions:
		// caps = VaultCapabilities()
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
