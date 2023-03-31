package identity_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/identity"
	"github.com/sonr-io/sonr/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PrimaryIdentities: []types.DidDocument{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		ServiceList: []types.Service{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, *k, genesisState)
	got := identity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PrimaryIdentities, got.PrimaryIdentities)
	require.ElementsMatch(t, genesisState.ServiceList, got.ServiceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
