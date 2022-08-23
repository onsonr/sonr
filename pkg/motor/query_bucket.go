package motor

import (
	"context"
	"net/http"

	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) QueryWhereIs(ctx context.Context, did string) (mt.QueryWhereIsResponse, error) {
	_, err := mtr.GetClient().GetWhereIs(did, mtr.Address)
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
	resp, err := mtr.GetClient().GetWhereIsByCreator(mtr.Address)
	var ptrArr []*bt.WhereIs = make([]*bt.WhereIs, len(resp.WhereIs))
	for _, wi := range resp.WhereIs {
		mtr.Resources.whereIsStore[wi.Did] = &wi
		ptrArr = append(ptrArr, &wi)
	}

	if err != nil {
		return mt.QueryWhereIsByCreatorResponse{}, err
	}

	// using the returned value from GetWhereIsByCreator as to not loop through the cache
	return mt.QueryWhereIsByCreatorResponse{
		Code:    http.StatusAccepted,
		WhereIs: ptrArr,
	}, nil
}
