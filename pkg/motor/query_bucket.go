package motor

import (
	"context"

	mt "github.com/sonr-io/sonr/pkg/motor/types"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) QueryWhereIs(ctx context.Context, req mt.QueryWhereIsRequest) (mt.QueryWhereIsResponse, error) {
	resp, err := mtr.Resources.bucketQueryClient.WhereIs(ctx, &bt.QueryGetWhereIsRequest{
		Creator: req.Creator,
		Did:     req.Did,
	})

	if err != nil {
		return mt.QueryWhereIsResponse{}, err
	}

	res := mt.QueryWhereIsResponse{
		WhereIs: &resp.WhereIs,
	}

	mtr.Resources.whereIsStore[res.WhereIs.Did] = res.WhereIs

	return res, nil
}
