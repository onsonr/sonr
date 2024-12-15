package capvault

import "github.com/onsonr/sonr/crypto/ucan"

func NewCap(ty CapVault) ucan.Capability {
	return ucan.Capability(ty)
}

func (c CapVault) Contains(b ucan.Capability) bool {
	return c.String() == b.String()
}
