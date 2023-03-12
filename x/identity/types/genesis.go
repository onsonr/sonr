package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultServices returns the default services
func DefaultServices() []Service {
	return []Service{
		{
			Id:     "did:web:sonr.io",
			Type:   "LinkedDomains",
			Origin: "https://sonr.io",
			Name:   "Sonr Home",
		},
		{
			Id:     "did:web:localhost",
			Type:   "LinkedDomains",
			Origin: "localhost",
			Name:   "Localhost",
		},
		{
			Id:     "did:web:mind.sonr.io",
			Type:   "LinkedDomains",
			Origin: "https://mind.sonr.io",
			Name:   "Sonr Mind",
		},
		{
			Id:     "did:web:auth.sonr.io",
			Type:   "LinkedDomains",
			Origin: "https://auth.sonr.io",
			Name:   "Sonr Auth",
		},
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidDocumentList: []DidDocument{},
		ServiceList:     DefaultServices(),
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
	ServiceIndexMap := make(map[string]struct{})

	for _, elem := range gs.ServiceList {
		index := string(ServiceKey(elem.Id))
		if _, ok := ServiceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for DomainRecord")
		}
		ServiceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
