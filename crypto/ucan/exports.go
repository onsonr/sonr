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

// NewResource returns a new resource identifier
func NewResource(resType string, path string) Resource {
	return NewStringLengthResource(resType, path)
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
		cap := preset.NewCap(capStr)
		if cap == nil {
			return EmptyAttenuation, fmt.Errorf("invalid capability %s for preset %s", capStr, preset)
		}

		// Create resource
		resource := NewResource(resType, path)

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
