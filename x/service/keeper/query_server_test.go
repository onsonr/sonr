package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/x/service"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)
	_, err := f.queryServer.Params(f.ctx, &service.QueryParamsRequest{})
	require.NoError(err)
}
