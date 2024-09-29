package keeper_test

import (
	"testing"

	apiv1 "github.com/onsonr/sonr/api/oracle/v1"
	"github.com/stretchr/testify/require"
)

func TestORM(t *testing.T) {
	f := SetupTest(t)

	dt := f.k.OrmDB.BalanceTable()
	amt := uint64(7)

	err := dt.Insert(f.ctx, &apiv1.Balance{
		Amount: amt,
	})
	require.NoError(t, err)

	require.NoError(t, err)
}
