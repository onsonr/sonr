package capaccount

import "github.com/onsonr/sonr/crypto/ucan"

func NewCap(ty CapAccount) ucan.Capability {
	return ucan.Capability(ty)
}

func (c CapAccount) Contains(b ucan.Capability) bool {
	return c.String() == b.String()
}
