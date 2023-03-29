package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PrimaryIdentities: []DidDocument{},
		BlockchainIdentities: []DidDocument{},
		ServiceList: []Service{
			{
				Id:         "sonr.io",
				Controller: "did:web:sonr.io",
				Type:       "LinkedDomains",
				Origin:     "sonr.io",
				Name:       "Sonr",
			},
			{
				Id:         "localhost",
				Controller: "did:web:localhost",
				Type:       "LinkedDomains",
				Origin:     "localhost",
				Name:       "Localhost",
			},
			{
				Id:         "sonr.id",
				Controller: "did:web:sonr.id",
				Type:       "LinkedDomains",
				Origin:     "sonr.id",
				Name:       "Sonr.ID",
			},
			{
				Id:         "sonr.ws",
				Controller: "did:web:sonr.ws",
				Type:       "LinkedDomains",
				Origin:     "sonr.ws",
				Name:       "Sonr WS",
			},
			{
				Id:         "sonr.wtf",
				Controller: "did:web:sonr.wtf",
				Type:       "LinkedDomains",
				Origin:     "sonr.wtf",
				Name:       "Sonr.wtf",
			},
			{
				Id:         "sonrhq.com",
				Controller: "did:web:sonrhq.com",
				Type:       "LinkedDomains",
				Origin:     "sonrhq.com",
				Name:       "Sonr HQ",
			},
			{
				Id:         "sonr.network",
				Controller: "did:web:sonr.network",
				Type:       "LinkedDomains",
				Origin:     "sonr.network",
				Name:       "Sonr Network",
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
	for _, elem := range gs.PrimaryIdentities {
		index := string(DidDocumentKey(elem.Id))
		if _, ok := didDocumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for did")
		}
		didDocumentIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in services
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
