package types

// DefaultGenesisState returns the default middleware GenesisState.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Records: DefaultRecords(),
	}
}

// NewGenesisState initializes and returns a new GenesisState.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Records: []Record{},
	}
}

// Validate performs basic validation of the GenesisState.
func (gs *GenesisState) Validate() error {
	return nil
}

// DefaultRecords returns default records.
func DefaultRecords() []Record {
	return []Record{
		{
			Name:        "Sonr Localhost",
			Origin:      "localhost",
			Description: "Sonr Localhost Validator Dashboard and Auth management.",
			Authority:   "gov",
		},
	}
}
