package ucan

import (
	"encoding/json"
	"fmt"
)

// Attenuations is a list of attenuations
type Attenuations []Attenuation

func (att Attenuations) String() string {
	str := ""
	for _, a := range att {
		str += fmt.Sprintf("%s\n", a)
	}
	return str
}

// Contains is true if all attenuations in b are contained
func (att Attenuations) Contains(b Attenuations) bool {
	// fmt.Printf("%scontains\n%s?\n\n", att, b)
LOOP:
	for _, bb := range b {
		for _, aa := range att {
			if aa.Contains(bb) {
				// fmt.Printf("%s contains %s\n", aa, bb)
				continue LOOP
			} else if aa.Rsc.Contains(bb.Rsc) {
				// fmt.Printf("%s < %s\n", aa, bb)
				// fmt.Printf("rscEq:%t rscContains: %t capContains:%t\n", aa.Rsc.Type() == bb.Rsc.Type(), aa.Rsc.Contains(bb.Rsc), aa.Cap.Contains(bb.Cap))
				return false
			}
		}
		return false
	}
	return true
}

// AttenuationConstructorFunc is a function that creates an attenuation from a map
// Users of this package provide an Attenuation Constructor to the parser to
// bind attenuation logic to a UCAN
type AttenuationConstructorFunc func(v map[string]interface{}) (Attenuation, error)

// Attenuation is a capability on a resource
type Attenuation struct {
	Cap Capability
	Rsc Resource
}

// String formats an attenuation as a string
func (a Attenuation) String() string {
	return fmt.Sprintf("cap:%s %s:%s", a.Cap, a.Rsc.Type(), a.Rsc.Value())
}

// Contains returns true if both
func (a Attenuation) Contains(b Attenuation) bool {
	return a.Rsc.Contains(b.Rsc) && a.Cap.Contains(b.Cap)
}

// MarshalJSON implements the json.Marshaller interface
func (a Attenuation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		a.Rsc.Type(): a.Rsc.Value(),
		CapKey:       a.Cap.String(),
	})
}

// Resource is a unique identifier for a thing, usually stored state. Resources
// are organized by string types
type Resource interface {
	Type() string
	Value() string
	Contains(b Resource) bool
}

type stringLengthRsc struct {
	t string
	v string
}

// NewStringLengthResource is a silly implementation of resource to use while
// I figure out what an OR filter on strings is. Don't use this.
func NewStringLengthResource(typ, val string) Resource {
	return stringLengthRsc{
		t: typ,
		v: val,
	}
}

func (r stringLengthRsc) Type() string {
	return r.t
}

func (r stringLengthRsc) Value() string {
	return r.v
}

func (r stringLengthRsc) Contains(b Resource) bool {
	return r.Type() == b.Type() && len(r.Value()) <= len(b.Value())
}

// Capability is an action users can perform
type Capability interface {
	// A Capability must be expressable as a string
	String() string
	// Capabilities must be comparable to other same-type capabilities
	Contains(b Capability) bool
}

// NestedCapabilities is a basic implementation of the Capabilities interface
// based on a hierarchal list of strings ordered from most to least capable
// It is both a capability and a capability factory with the .Cap method
type NestedCapabilities struct {
	cap       string
	idx       int
	hierarchy *[]string
}

// assert at compile-time NestedCapabilities implements Capability
var _ Capability = (*NestedCapabilities)(nil)

// NewNestedCapabilities creates a set of NestedCapabilities
func NewNestedCapabilities(strs ...string) NestedCapabilities {
	return NestedCapabilities{
		cap:       strs[0],
		idx:       0,
		hierarchy: &strs,
	}
}

// Cap creates a new capability from the hierarchy
func (nc NestedCapabilities) Cap(str string) Capability {
	idx := -1
	for i, c := range *nc.hierarchy {
		if c == str {
			idx = i
			break
		}
	}
	if idx == -1 {
		panic(fmt.Sprintf("%s is not a nested capability. must be one of: %v", str, *nc.hierarchy))
	}

	return NestedCapabilities{
		cap:       str,
		idx:       idx,
		hierarchy: nc.hierarchy,
	}
}

// String returns the Capability value as a string
func (nc NestedCapabilities) String() string {
	return nc.cap
}

// Contains returns true if cap is equal or less than the NestedCapability value
func (nc NestedCapabilities) Contains(cap Capability) bool {
	str := cap.String()
	for i, c := range *nc.hierarchy {
		if c == str {
			if i >= nc.idx {
				return true
			}
			return false
		}
	}
	return false
}
