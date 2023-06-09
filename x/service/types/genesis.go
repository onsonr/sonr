package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ServiceRecordList: []ServiceRecord{
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
		ServiceRelationshipsList: []ServiceRelationship{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in serviceRecord
	serviceRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.ServiceRecordList {
		index := string(ServiceRecordKey(elem.Id))
		if _, ok := serviceRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for serviceRecord")
		}
		serviceRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in serviceRelationships
	serviceRelationshipsIdMap := make(map[string]bool)
	for _, elem := range gs.ServiceRelationshipsList {
		if _, ok := serviceRelationshipsIdMap[elem.Did]; ok {
			return fmt.Errorf("duplicated id for serviceRelationships")
		}
		serviceRelationshipsIdMap[elem.Did] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
