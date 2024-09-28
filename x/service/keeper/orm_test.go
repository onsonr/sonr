package keeper_test

import (
	"testing"
)

func TestORM(t *testing.T) {
	f := SetupTest(t)
	if f == nil {
		return
	}

	// dt := f.k.OrmDB.ExampleDataTable()
	// acc := []byte("test_acc")
	// amt := uint64(7)
	//
	// err := dt.Insert(f.ctx, &apiv1.ExampleData{
	// 	Account: acc,
	// 	Amount:  amt,
	// })
	// require.NoError(t, err)
	//
	// d, err := dt.Has(f.ctx, []byte("test_acc"))
	// require.NoError(t, err)
	// require.True(t, d)
	//
	// res, err := dt.Get(f.ctx, []byte("test_acc"))
	// require.NoError(t, err)
	// require.NotNil(t, res)
	// require.EqualValues(t, amt, res.Amount)
}
