package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidDocumentList:  []DidDocument{},
		DomainRecordList: []DomainRecord{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in didDocument
	didDocumentIndexMap := make(map[string]struct{})

	for _, elem := range gs.DidDocumentList {
		index := string(DidDocumentKey(elem.Id))
		if _, ok := didDocumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for didDocument")
		}
		didDocumentIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in DomainRecord
	DomainRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.DomainRecordList {
		index := string(DomainRecordKey(elem.Domain, elem.Index))
		if _, ok := DomainRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for DomainRecord")
		}
		DomainRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
