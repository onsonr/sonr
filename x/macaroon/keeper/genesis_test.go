package keeper_test

import (
	"testing"

	"github.com/onsonr/sonr/x/macaroon/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	f := SetupTest(t)

	genesisState := &types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	f.k.InitGenesis(f.ctx, genesisState)

	got := f.k.ExportGenesis(f.ctx)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
