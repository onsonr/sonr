package keeper_test

import (
	"testing"

	"github.com/sonrhq/sonr/x/identity"
	"github.com/stretchr/testify/require"
)

func TestInitGenesis(t *testing.T) {
	fixture := initFixture(t)

	data := &identity.GenesisState{
		Counters: []identity.Counter{
			{
				Address: fixture.addrs[0].String(),
				Count:   5,
			},
		},
		Params: identity.DefaultParams(),
	}
	err := fixture.k.InitGenesis(fixture.ctx, data)
	require.NoError(t, err)

	params, err := fixture.k.Params.Get(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, identity.DefaultParams(), params)

	count, err := fixture.k.Counter.Get(fixture.ctx, fixture.addrs[0].String())
	require.NoError(t, err)
	require.Equal(t, uint64(5), count)
}

func TestExportGenesis(t *testing.T) {
	fixture := initFixture(t)

	_, err := fixture.msgServer.IncrementCounter(fixture.ctx, &identity.MsgIncrementCounter{
		Sender: fixture.addrs[0].String(),
	})
	require.NoError(t, err)

	out, err := fixture.k.ExportGenesis(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, identity.DefaultParams(), out.Params)
	require.Equal(t, uint64(1), out.Counters[0].Count)
}
