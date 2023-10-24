package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DIDDocumentList:       []DIDDocument{},
		ControllerAccountList: []ControllerAccount{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	// Check for duplicated index in dIDDocument
	dIDDocumentIndexMap := make(map[string]struct{})

	for _, elem := range gs.DIDDocumentList {
		index := string(DIDDocumentKey(elem.Id))
		if _, ok := dIDDocumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for dIDDocument")
		}
		dIDDocumentIndexMap[index] = struct{}{}
	}

	// Check for duplicated ID in controllerAccount
	controllerAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.ControllerAccountList {
		index := string(ControllerAccountKey(elem.Address))
		if _, ok := controllerAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for dIDDocument")
		}
		controllerAccountIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
