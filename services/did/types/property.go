package types

import (
	"fmt"
	"strings"
)

// DIDProperty is a string that is stored under a DID identifier in its internal store.
type DIDProperty string

// NewProperty creates a new property for the DID
func NewProperty(id DIDIdentifier, key string) DIDProperty {
	return DIDProperty(fmt.Sprintf("%s#%s", id.String(), key))
}

// Key returns the key of the resource
func (d DIDProperty) Key() string {
	ptrs := strings.Split(string(d), "#")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[1]
}

// Value Returns the value for the Properties key
func (d DIDProperty) Value() string {
	return d.Identifier().GetKey(d.Key())
}

// Identifier returns the identifier of the resource
func (d DIDProperty) Identifier() DIDIdentifier {
	ptrs := strings.Split(string(d), "#")
	if len(ptrs) < 2 {
		return ""
	}
	return DIDIdentifier(ptrs[0])
}

// String returns the string representation of the property
func (d DIDProperty) String() string {
	return string(d)
}
