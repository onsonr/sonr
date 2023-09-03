package handler

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	storagepb "github.com/sonrhq/core/types/highway/storage/v1"
)

// StorageHandler is the handler for the authentication service
type StorageHandler struct {
	cctx client.Context
}

// GetCID returns the CID for the given data.
func (a *StorageHandler) GetCID(ctx context.Context, req *storagepb.GetCIDRequest) (*storagepb.GetCIDResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// PutData puts the data into the storage.
func (a *StorageHandler) PutData(ctx context.Context, req *storagepb.PutDataRequest) (*storagepb.PutDataResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
