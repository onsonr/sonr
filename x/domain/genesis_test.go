package domain_test

import (
	"testing"

	keepertest "github.com/sonr-io/core/testutil/keeper"
	"github.com/sonr-io/core/testutil/nullify"
	"github.com/sonr-io/core/x/domain"
	"github.com/sonr-io/core/x/domain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UsernameRecordsList: []types.UsernameRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DomainKeeper(t)
	domain.InitGenesis(ctx, *k, genesisState)
	got := domain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UsernameRecordsList, got.UsernameRecordsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
