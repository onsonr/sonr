package types

import (
	"encoding/json"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	// TODO:
	return Params{}
}

// Stringer method for Params.
func (p Params) String() string {
	bz, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(bz)
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	// TODO:
	return nil
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
	if c.Command != that.Command {
		return false
	}
	if c.Description != that.Description {
		return false
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
