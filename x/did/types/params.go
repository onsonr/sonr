package types

import (
	"encoding/json"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Attenuations: DefaultAttenuations(),
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
	// TODO:
	return nil
}

// DefaultAttenuations returns the default Attenuation
func DefaultAttenuations() []*Attenuation {
	return []*Attenuation{
		{
			Resource: &Resource{
				Kind:     "did",
				Template: "{id}",
			},
			Capabilities: []*Capability{
				{
					Name:        "execute",
					Parent:      "/account",
					Command:     "/account/execute",
					Description: "Execute a transaction on behalf of the account",
				},
				{
					Name:        "link",
					Parent:      "/account",
					Command:     "/account/link",
					Description: "Link a Verification Method to the account",
				},
				{
					Name:        "unlink",
					Parent:      "/account",
					Command:     "/account/unlink",
					Description: "Unlink a Verification Method from the account",
				},
				{
					Name:        "authenticate",
					Parent:      "/account",
					Command:     "/account/authenticate",
					Description: "Authenticate the account using a Verification Method",
				},
			},
		},
	}
}
