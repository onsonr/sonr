package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/x/service"
)

func TestInitGenesis(t *testing.T) {
	fixture := initFixture(t)

	data := &service.GenesisState{
		Params: service.DefaultParams(),
	}
	err := fixture.k.InitGenesis(fixture.ctx, data)
	require.NoError(t, err)

	params, err := fixture.k.Params.Get(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, service.DefaultParams(), params)
}

func TestExportGenesis(t *testing.T) {
	fixture := initFixture(t)

	out, err := fixture.k.ExportGenesis(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, service.DefaultParams(), out.Params)
}
