package motor

import (
	"net/http"

	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) QueryWhereIs(did string) (mt.QueryWhereIsResponse, error) {
	// use the item within the cache from GetWhereIs
	if wi := mtr.Resources.whereIsStore[did]; wi != nil {
		return mt.QueryWhereIsResponse{
			WhereIs: wi,
		}, nil
	}

	resp, err := mtr.GetClient().QueryWhereIs(did, mtr.Address)
	if err != nil {
		return mt.QueryWhereIsResponse{}, err
	}

	return mt.QueryWhereIsResponse{
		WhereIs: resp,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhereIsForCreator() (mt.QueryWhereIsByCreatorResponse, error) {
	resp, err := mtr.GetClient().QueryWhereIsByCreator(mtr.Address)
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
