package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TLDRecordList: []TLDRecord{},
		SLDRecordList: []SLDRecord{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tLDRecord
	tLDRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.TLDRecordList {
		index := string(TLDRecordKey(elem.Index))
		if _, ok := tLDRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tLDRecord")
		}
		tLDRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in sLDRecord
	sLDRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.SLDRecordList {
		index := string(SLDRecordKey(elem.Index))
		if _, ok := sLDRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sLDRecord")
		}
		sLDRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
