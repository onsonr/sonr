// Package resolver provides the methods to resolve DIDs and UCANs
package resolver

import (
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/crypto/ucan/didkey"
)

type AttenuationConstructorFunc func(v map[string]interface{}) (ucan.Attenuation, error)

type UCANResolver interface {
	ResolveUCAN(did string) (string, error)
}

type DIDResolver interface {
	ResolveDID(did string) (didkey.ID, error)
}
