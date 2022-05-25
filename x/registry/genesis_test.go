package registry_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/registry"
	"github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		WhoIsList: []types.WhoIs{
			{
				Owner: "0",
			},
			{
				Owner: "1",
			},
		},
		WhoIsCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RegistryKeeper(t)
	registry.InitGenesis(ctx, *k, genesisState)
	got := registry.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.WhoIsList, got.WhoIsList)
	require.Equal(t, genesisState.WhoIsCount, got.WhoIsCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
