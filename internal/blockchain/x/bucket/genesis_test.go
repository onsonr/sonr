package bucket_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/internal/blockchain/testutil/keeper"
	"github.com/sonr-io/sonr/internal/blockchain/testutil/nullify"
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket"
	"github.com/sonr-io/sonr/internal/blockchain/x/bucket/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WhichIsList: []types.WhichIs{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BucketKeeper(t)
	bucket.InitGenesis(ctx, *k, genesisState)
	got := bucket.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WhichIsList, got.WhichIsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
