package protocol_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ipfs/go-datastore/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/x/schema/types"

	"github.com/ipfs/go-datastore"
)

const IPFSShellUrl = "localhost:5001"

type mockCache struct {
	mock.Mock
}

func (m *mockCache) Get(ctx context.Context, key datastore.Key) (value []byte, err error) {
	args := m.Called(ctx, key)
	return []byte(args.String(0)), args.Error(1)
}

func (m *mockCache) Has(ctx context.Context, key datastore.Key) (exists bool, err error) {
	panic("implement me")
}

func (m *mockCache) GetSize(ctx context.Context, key datastore.Key) (size int, err error) {
	panic("implement me")
}

func (m *mockCache) Query(ctx context.Context, q query.Query) (query.Results, error) {
	panic("implement me")
}

func (m *mockCache) Put(ctx context.Context, key datastore.Key, value []byte) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *mockCache) Delete(ctx context.Context, key datastore.Key) error {
	panic("implement me")
}

func (m *mockCache) Sync(ctx context.Context, prefix datastore.Key) error {
	panic("implement me")
}

func (m *mockCache) Close() error {
	panic("implement me")
}

func TestIPFSShell_PutData(t *testing.T) {
	ctx := context.Background()
	cacheStore := new(mockCache)

	i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)

	data, err := json.Marshal(&types.SchemaDefinition{
		Creator: "snr1h48jyesl50ahruft5p350nmnycaegdej2pzkdx",
		Label:   "test-label",
	})
	assert.NoError(t, err)

	cid := "QmcHujytrGJ7LqiG38pr83WhZqgM2vLWGqsERVVVyqHLmS"
	cacheStore.
		On(
			"Put",
			ctx,
			datastore.NewKey(cid),
			data,
		).
		Return(nil)

	got, err := i.PutData(ctx, data)
	assert.NoError(t, err)
	assert.Equal(t, cid, got)
}
