package types

// Capability hierarchy for smart account operations
// ----------------------------------------------
// OWNER
//
//	└─ OPERATOR
//	     ├─ EXECUTE
//	     ├─ PROPOSE
//	     └─ SIGN
//	└─ SET_POLICY
//	     └─ SET_THRESHOLD
//	└─ RECOVER
//	     └─ SOCIAL
//
// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// Equal checks if two Attenuation are equal
func (a *Attenuation) Equal(that *Attenuation) bool {
	if that == nil {
		return false
	}
	if a.Resource != nil {
		if that.Resource == nil {
			return false
		}
		if !a.Resource.Equal(that.Resource) {
			return false
		}
	}
	if len(a.Capabilities) != len(that.Capabilities) {
		return false
	}
	for i := range a.Capabilities {
		if !a.Capabilities[i].Equal(that.Capabilities[i]) {
			return false
		}
	}
	return true
}

// Equal checks if two Capability are equal
func (c *Capability) Equal(that *Capability) bool {
	if that == nil {
		return false
	}
	if c.Name != that.Name {
		return false
	}
	if c.Parent != that.Parent {
		return false
	}
	// TODO: check description
	if len(c.Resources) != len(that.Resources) {
		return false
	}
	for i := range c.Resources {
		if c.Resources[i] != that.Resources[i] {
			return false
		}
	}
	return true
}

// Equal checks if two Resource are equal
func (r *Resource) Equal(that *Resource) bool {
	if that == nil {
		return false
	}
	if r.Kind != that.Kind {
		return false
	}
	if r.Template != that.Template {
		return false
	}
	return true
}
