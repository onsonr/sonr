package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestORM(t *testing.T) {
	f := SetupTest(t)

	dt := f.k.OrmDB.DWNTable()
	require.NotNil(t, dt)
}
