package motor

import (
	"context"
	"net/http"

	mt "github.com/sonr-io/sonr/pkg/motor/types"
)

func (mtr *motorNodeImpl) QueryWhereIs(ctx context.Context, did string) (mt.QueryWhereIsResponse, error) {
	_, err := mtr.Resources.GetWhereIs(ctx, did, mtr.Address)
	if err != nil {
		return mt.QueryWhereIsResponse{}, err
	}

	// use the item within the cache from GetWhereIs
	wi := mtr.Resources.whereIsStore[did]

	return mt.QueryWhereIsResponse{
		WhereIs: wi,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhereIsByCreator(ctx context.Context) (mt.QueryWhereIsByCreatorResponse, error) {
	resp, err := mtr.Resources.GetWhereIsByCreator(ctx, mtr.Address)

	if err != nil {
		return mt.QueryWhereIsByCreatorResponse{}, err
	}

	// using the returned value from GetWhereIsByCreator as to not loop through the cache
	return mt.QueryWhereIsByCreatorResponse{
		Code:    http.StatusAccepted,
		WhereIs: resp,
	}, nil
}
