package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidDocumentList: []DidDocument{},
		ServiceList: []Service{
			{
				Id:         "sonr.io",
				Controller: "did:web:sonr.io",
				Type:       "LinkedDomains",
				Origin:     "sonr.io",
				Name:       "Sonr Home",
			},
			{
				Id:         "localhost",
				Controller: "did:web:localhost",
				Type:       "LinkedDomains",
				Origin:     "localhost",
				Name:       "Localhost",
			},
			{
				Id:         "mind.sonr.io",
				Controller: "did:web:mind.sonr.io",
				Type:       "LinkedDomains",
				Origin:     "mind.sonr.io",
				Name:       "Sonr Mind",
			},
			{
				Id:         "auth.sonr.io",
				Controller: "did:web:auth.sonr.io",
				Type:       "LinkedDomains",
				Origin:     "auth.sonr.io",
				Name:       "Sonr Auth",
			},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params:        DefaultParams(),
		Relationships: []VerificationRelationship{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	relationshipMap := make(map[string]struct{})
	for _, elem := range gs.Relationships {
		index := string(RelationshipKey(elem.Reference))
		if _, ok := relationshipMap[index]; ok {
			return fmt.Errorf("duplicated id for relationship")
		}
		relationshipMap[elem.Reference] = struct{}{}
	}
	// Check for duplicated index in didDocument
	didDocumentIndexMap := make(map[string]struct{})

	for _, elem := range gs.DidDocumentList {
		index := string(DidDocumentKey(elem.Id))
		if _, ok := didDocumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for did")
		}
		didDocumentIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in DomainRecord
	ServiceIndexMap := make(map[string]struct{})

	for _, elem := range gs.ServiceList {
		index := string(ServiceKey(elem.Origin))
		if _, ok := ServiceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for Service")
		}
		ServiceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
