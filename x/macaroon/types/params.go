package types

import (
	"encoding/json"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Methods: DefaultMethods(),
		Scopes:  DefaultScopes(),
		Caveats: DefaultCaveats(),
	}
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
	return nil
}

func DefaultMethods() *Methods {
	return &Methods{
		Default:   "did:sonr",
		Supported: []string{"did:key", "did:web", "did:sonr", "did:ipfs", "did:btcr", "did:ethr"},
	}
}

func DefaultScopes() *Scopes {
	return &Scopes{
		Base:      "openid profile sonr.address",
		Supported: []string{"create", "read", "update", "delete", "sign", "verify", "simulate", "execute", "broadcast", "admin"},
	}
}

func DefaultCaveats() *Caveats {
	return &Caveats{
		// First party - JWT Format
		SupportedFirstParty: []string{"aud", "exp", "iat", "iss", "nbf", "nonce", "sub"},
		// Third party - UCAN Format
		SupportedThirdParty: []string{"cap", "nbf", "exp", "att", "prf", "rmt", "sig", "ucv"},
	}
}
