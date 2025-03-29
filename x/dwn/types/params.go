package types

import (
	"encoding/json"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Attenuations: DefaultAttenuations(),
		AllowedOperators: []string{ // TODO:
			"localhost",
			"didao.xyz",
			"sonr.id",
		},
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
				Kind:     "dwn",
				Template: "{id}",
			},
			Capabilities: []*Capability{
				{
					Name:        "sign",
					Parent:      "/vault",
					Command:     "/vault/sign",
					Description: "Sign an arbitrary payload",
				},
				{
					Name:        "verify",
					Parent:      "/vault",
					Command:     "/vault/verify",
					Description: "Verify a signature",
				},
				{
					Name:        "refresh",
					Parent:      "/vault",
					Command:     "/vault/refresh",
					Description: "Refresh the Session KeyShares",
				},
			},
		},
	}
}
