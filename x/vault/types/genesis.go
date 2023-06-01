package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ClaimableWalletList: []ClaimableWallet{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
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
