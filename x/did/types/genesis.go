package types

import "encoding/json"

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		WhitelistedAssets: DefaultAssets(),
		WhitelistedChains: DefaultChains(),
		AllowedPublicKeys: DefaultKeyInfos(),
		OpenidConfig:      DefaultOpenIDConfig(),
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
