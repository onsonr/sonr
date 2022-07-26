package protocol_test

import (
	"context"
	"encoding/json"
	"github.com/ipfs/go-datastore/query"
	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/ipfs/go-datastore"
	shell "github.com/ipfs/go-ipfs-api"
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
	cacheStore := new(mockCache)

	type fields struct {
		Shell *shell.Shell
		cache datastore.Datastore
	}
	type args struct {
		ctx    context.Context
		schema *types.SchemaDefinition
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{name: "test#1", fields: fields{
			Shell: nil,
			cache: nil,
		}, args: args{
			ctx: nil,
			schema: &types.SchemaDefinition{
				Creator: "did:snr:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm",
				Label:   "test-label",
			},
		}, want: "QmW4Ghk82fyq4LsoBKwH5o66Zb1sEpZ735Tmn1yA7o1uGu", wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)

			data, err := json.Marshal(tt.args.schema)
			assert.NoError(t, err)

			cacheStore.
				On("Put", nil, datastore.NewKey("QmW4Ghk82fyq4LsoBKwH5o66Zb1sEpZ735Tmn1yA7o1uGu"), data).
				Return(tt.wantErr)

			got, err := i.PutData(tt.args.ctx, data)

			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
