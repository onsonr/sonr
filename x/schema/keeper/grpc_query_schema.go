package keeper

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/x/schema/types"
)

func (k Keeper) QuerySchema(ctx context.Context, req *types.QuerySchemaRequest) (*types.QuerySchemaResponse, error) {
	// TODO: implement
	return nil, errors.New("unimplemented")
}
