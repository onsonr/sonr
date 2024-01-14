package identity

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	uniq := make(map[string]bool)
	for _, counter := range gs.Counters {
		if _, ok := uniq[counter.Address]; ok {
			return ErrDuplicateAddress
		}

		uniq[counter.Address] = true
	}

	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return nil
}
