package types

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:  DefaultParams(),
		Records: DefaultRecords(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// DefaultRecords returns default records.
func DefaultRecords() []Record {
	return []Record{
		{
			Name:        "Localhost Dev",
			Origin:      "localhost",
			Description: "Sonr Localhost Validator Dashboard and Auth management.",
			Authority:   "gov",
		},
	}
}
