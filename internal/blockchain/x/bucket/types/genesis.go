package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WhichIsList: []WhichIs{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in whichIs
	whichIsIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhichIsList {
		index := string(WhichIsKey(elem.Did))
		if _, ok := whichIsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for whichIs")
		}
		whichIsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
