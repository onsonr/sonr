package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidDocuments:        []Identity{},
		Relationships:       []VerificationRelationship{},
		ClaimableWalletList: []ClaimableWallet{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in relationships
	relationshipMap := make(map[string]struct{})
	for _, elem := range gs.Relationships {
		index := string(RelationshipKey(elem.Reference))
		if _, ok := relationshipMap[index]; ok {
			return fmt.Errorf("duplicated id for relationship")
		}
		relationshipMap[elem.Reference] = struct{}{}
	}

	// Check for duplicated index in primary identities
	didDocumentIndexMap := make(map[string]struct{})
	for _, elem := range gs.DidDocuments {
		index := string(DidDocumentKey(elem.Id))
		if _, ok := didDocumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for did")
		}
		didDocumentIndexMap[index] = struct{}{}
	}

	// Check for duplicated ID in claimableWallet
	claimableWalletIdMap := make(map[uint64]bool)
	claimableWalletCount := gs.GetClaimableWalletCount()
	for _, elem := range gs.ClaimableWalletList {
		if _, ok := claimableWalletIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for claimableWallet")
		}
		if elem.Id >= claimableWalletCount {
			return fmt.Errorf("claimableWallet id should be lower or equal than the last id")
		}
		claimableWalletIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
