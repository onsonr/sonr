package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const (
	WEBAUTHN_DID_METHOD = "webauthn"
	SONR_DID_METHOD     = "sonr"
	BTC_DID_METHOD      = "btcr"
	ETH_DID_METHOD      = "ethr"
	WEB_DID_METHOD      = "web"
	COMM_DID_METHOD     = "commx"
	DNS_DID_METHOD      = "dns"
)

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
			WEBAUTHN_DID_METHOD,
			SONR_DID_METHOD,
			BTC_DID_METHOD,
			ETH_DID_METHOD,
			WEB_DID_METHOD,
			COMM_DID_METHOD,
			DNS_DID_METHOD,
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

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
