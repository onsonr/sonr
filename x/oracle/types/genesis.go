package types

import (
	"github.com/onsonr/sonr/pkg/orm/assettype"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// DefaultAssets returns the default asset infos: BTC, ETH, SNR, and USDC
func DefaultAssets() []*AssetInfo {
	return []*AssetInfo{
		{
			Name:      "Bitcoin",
			Symbol:    "BTC",
			Hrp:       "bc",
			Index:     0,
			AssetType: assettype.Native.String(),
			IconUrl:   "https://cdn.sonr.land/BTC.svg",
		},
		{
			Name:      "Ethereum",
			Symbol:    "ETH",
			Hrp:       "eth",
			Index:     64,
			AssetType: assettype.Native.String(),
			IconUrl:   "https://cdn.sonr.land/ETH.svg",
		},
		{
			Name:      "Sonr",
			Symbol:    "SNR",
			Hrp:       "idx",
			Index:     703,
			AssetType: assettype.Native.String(),
			IconUrl:   "https://cdn.sonr.land/SNR.svg",
		},
	}
}
