package didcommon

import (
	"fmt"
	"strings"
)

// DIDProperty is a string that is stored under a DID identifier in its internal store.
type DIDProperty string

// NewProperty creates a new property for the DID
func NewProperty(id string, key string) DIDProperty {
	return DIDProperty(fmt.Sprintf("%s#%s", id, key))
}

// Key returns the key of the resource
func (d DIDProperty) Key() string {
	ptrs := strings.Split(string(d), "#")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[1]
}

// Identifier returns the identifier of the resource
func (d DIDProperty) Identifier() string {
	ptrs := strings.Split(string(d), "#")
	if len(ptrs) < 2 {
		return ""
	}
	return ptrs[0]
}

// String returns the string representation of the property
func (d DIDProperty) String() string {
	return string(d)
}
