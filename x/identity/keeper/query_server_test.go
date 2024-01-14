package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/x/identity"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Params(f.ctx, &identity.QueryParamsRequest{})
	require.NoError(err)
	require.Equal(identity.Params{}, resp.Params)
}
