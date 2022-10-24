package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonr-io/sonr/x/bucket/types"
)

func TestWhereIsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.DefineBucket(ctx, &types.MsgDefineBucket{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, resp.WhereIs)
	}
}
