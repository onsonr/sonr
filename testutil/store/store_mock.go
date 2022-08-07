package store

import (
	"context"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Get(ctx context.Context, key datastore.Key) (value []byte, err error) {
	args := m.Called(ctx, key)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCache) Has(ctx context.Context, key datastore.Key) (exists bool, err error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
}

func (m *MockCache) GetSize(ctx context.Context, key datastore.Key) (size int, err error) {
	panic("implement me")
}

func (m *MockCache) Query(ctx context.Context, q query.Query) (query.Results, error) {
	panic("implement me")
}

func (m *MockCache) Put(ctx context.Context, key datastore.Key, value []byte) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockCache) Delete(ctx context.Context, key datastore.Key) error {
	panic("implement me")
}

func (m *MockCache) Sync(ctx context.Context, prefix datastore.Key) error {
	panic("implement me")
}

func (m *MockCache) Close() error {
	panic("implement me")
}
