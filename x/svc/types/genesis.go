package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// DefaultAttenuations returns the default Attenuation
func DefaultAttenuations() []*Attenuation {
	return []*Attenuation{
		{
			Resource: &Resource{
				Kind:     "dwn",
				Template: "https://dwn.sonr.io/dwn",
			},
			Capabilities: []*Capability{
				{
					Name:    "store",
					Parent:  "dwn",
					Command: "store",
				},
				{
					Name:    "store-reply",
					Parent:  "dwn",
					Command: "store-reply",
				},
			},
		},
	}
}
