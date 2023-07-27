package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                PortID,
		DIDDocumentList:       []DIDDocument{},
		ControllerAccountList: []ControllerAccount{},
		EscrowAccountList:     []EscrowAccount{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
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
	controllerAccountIdMap := make(map[uint64]bool)
	controllerAccountCount := gs.GetControllerAccountCount()
	for _, elem := range gs.ControllerAccountList {
		if _, ok := controllerAccountIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for controllerAccount")
		}
		if elem.Id >= controllerAccountCount {
			return fmt.Errorf("controllerAccount id should be lower or equal than the last id")
		}
		controllerAccountIdMap[elem.Id] = true
	}
	// Check for duplicated ID in escrowAccount
	escrowAccountIdMap := make(map[uint64]bool)
	escrowAccountCount := gs.GetEscrowAccountCount()
	for _, elem := range gs.EscrowAccountList {
		if _, ok := escrowAccountIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for escrowAccount")
		}
		if elem.Id >= escrowAccountCount {
			return fmt.Errorf("escrowAccount id should be lower or equal than the last id")
		}
		escrowAccountIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
