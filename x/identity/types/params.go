package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		DidBaseContext:          "https://www.w3.org/ns/did/v1",
		AcccountDiscoveryReward: 1,
		AccountDidMethodName:    "sonr",
		AccountDidMethodContext: "https://docs.sonr.io/identity/1.0",
		SupportedDidMethods: []string{
			"sonr",
			"btcr",
			"ethr",
			"web",
			"key",
			"sov",
			"webauthn",
		},
		MaximumIdentityAliases: 2,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

func (p Params) IsSupportedDidMethod(method string) bool {
	for _, m := range p.SupportedDidMethods {
		if m == method {
			return true
		}
	}
	return false
}
