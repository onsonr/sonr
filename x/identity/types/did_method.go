package types

import (
	fmt "fmt"
	"strings"
)

// Format returns a string representation of the DIDMethod that is on the DID spec
func (m DIDMethod) Format(val string, options ...FormatOption) string {
	if m == DIDMethod_DIDMethod_BLOCKCHAIN {
		ct := findCoinTypeFromAddress(val)
		return fmt.Sprintf("did:%s:%s", ct.DidMethod(), val)
	}
	r := fmt.Sprintf("did:%s:%s", m.PrettyString(), val)
	for _, opt := range options {
		r = opt(r)
	}
	return r
}

// PrettyString returns a string representation of the DIDMethod that is on the DID spec
func (m DIDMethod) PrettyString() string {
	prts := strings.Split(m.String(), "_")
	return strings.ToLower(prts[len(prts)-1])
}

// FormatOption is a function that can be used to format a DIDMethod
type FormatOption func(string) string

// WithFragment returns a FormatOption that will append a fragment to the DID
func WithFragment(frag string) FormatOption {
	return func(did string) string {
		return fmt.Sprintf("%s#%s", did, frag)
	}
}

// WithPath returns a FormatOption that will append a path to the DID
func WithPath(path string) FormatOption {
	return func(did string) string {
		return fmt.Sprintf("%s/%s", did, path)
	}
}

// WithQuery returns a FormatOption that will append a query to the DID
func WithQuery(query string) FormatOption {
	return func(did string) string {
		return fmt.Sprintf("%s?%s", did, query)
	}
}
