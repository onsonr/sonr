package capinterchain

import "github.com/onsonr/sonr/internal/crypto/ucan"

func NewCap(ty CapInterchain) ucan.Capability {
	return ucan.Capability(ty)
}

func (c CapInterchain) Contains(b ucan.Capability) bool {
	return c.String() == b.String()
}
