package types

// DefaultGenesisState returns the default middleware GenesisState.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// NewGenesisState initializes and returns a new GenesisState.
func NewGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate performs basic validation of the GenesisState.
func (gs *GenesisState) Validate() error {
	return nil
}
