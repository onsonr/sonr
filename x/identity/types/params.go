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
		DidBaseContext:   "https://www.w3.org/ns/did/v1",
		DidMethodContext: "https://docs.sonr.io/identity/1.0/",
		DidMethodName:    "snr",
		DidMethodVersion: "1.0",
		DidNetwork:       "devnet",
		IpfsGateway:      "https://ipfs.sonr.network/ipfs/",
		IpfsApi:          "https://ipfs.sonr.network",
		HnsTlds: []string{
			".snr",
			".sonr",
		},
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
