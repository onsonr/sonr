package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onsonr/sonr/x/did/types"
)

func TestGenesis(t *testing.T) {
	f := SetupTest(t)

	genesisState := &types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	err := f.k.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)

	got := f.k.ExportGenesis(f.ctx)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
