package object_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/nullify"
	"github.com/sonr-io/sonr/internal/blockchain/x/object"
	"github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WhatIsList: []types.WhatIs{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ObjectKeeper(t)
	object.InitGenesis(ctx, *k, genesisState)
	got := object.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WhatIsList, got.WhatIsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
