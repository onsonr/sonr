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
				Id:          "localhost",
				Controller:  "did:web:localhost",
				Origins:     []string{"localhost"},
				Name:        "Sandbox",
				Description: "A sandbox environment for testing authenticated services.",
			},
			{
				Id:          "sonr.io",
				Controller:  "did:web:sonr.io",
				Origins:     []string{"sonr.io"},
				Name:        "Sonr",
				Description: "Sonr is a decentralized identity platform. This website is a preview of the utilities that will be available to users.",
			},
			{
				Id:          "sonr.dev",
				Controller:  "did:web:sonr.dev",
				Origins:     []string{"sonr.dev"},
				Name:        "Sonr Preview",
				Description: "Sonr is a decentralized identity platform. This website is a preview of the utilities that will be available to users.",
			},
			{
				Id:          "idx.is",
				Controller:  "did:web:idx.is",
				Origins:     []string{"idx.is"},
				Name:        "IDX",
				Description: "A fully peer-to-peer 'magic link' system for Web3 Identities.",
			},
		},
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
		index := string(KeyPrefix(elem.Id))
		if _, ok := serviceRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for serviceRecord")
		}
		serviceRecordIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
