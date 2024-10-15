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
		SupportedFirstParty: DefaultFirstPartyCaveats(),
		// Third party - UCAN Format
		SupportedThirdParty: DefaultThirdPartyCaveats(),
	}
}

func DefaultFirstPartyCaveats() []*Caveat {
	return []*Caveat{
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "aud",
			Description: "Audience must be a valid DID",
		},
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "exp",
			Description: "Expiration time must be a valid timestamp",
		},
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "iat",
			Description: "Issued at time must be a valid timestamp",
		},
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "iss",
			Description: "Issuer must be a valid DID",
		},
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "nbf",
			Description: "Not before time must be a valid timestamp",
		},
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "nonce",
			Description: "Nonce must be a valid string",
		},
		{
			Scopes:      []string{"openid", "profile", "sonr.address"},
			Caveat:      "sub",
			Description: "Subject must be a valid DID",
		},
	}
}

func DefaultThirdPartyCaveats() []*Caveat {
	return []*Caveat{
		{
			Scopes:      []string{"create", "read", "update", "delete", "sign", "verify", "simulate", "execute", "broadcast", "admin"},
			Caveat:      "cap",
			Description: "Capability must be a valid capability",
		},
		{
			Scopes:      []string{"create", "read", "update", "delete", "sign", "verify", "simulate", "execute", "broadcast", "admin"},
			Caveat:      "exp",
			Description: "Expiration time must be a valid timestamp",
		},
	}
}

func (c *Caveat) Equal(other *Caveat) bool {
	return c.Caveat == other.Caveat
}
