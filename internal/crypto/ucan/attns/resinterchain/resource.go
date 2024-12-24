package resinterchain

import "github.com/onsonr/sonr/internal/crypto/ucan"

func Build(ty ResInterchain, value string) ucan.Resource {
	return newStringLengthResource(ty.String(), value)
}

type stringLengthRsc struct {
	t string
	v string
}

// NewStringLengthResource is a silly implementation of resource to use while
// I figure out what an OR filter on strings is. Don't use this.
func newStringLengthResource(typ, val string) ucan.Resource {
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

func (r stringLengthRsc) Contains(b ucan.Resource) bool {
	return r.Type() == b.Type() && len(r.Value()) <= len(b.Value())
}
