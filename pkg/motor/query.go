package motor

import (
	"fmt"
	"net/http"

	mt "github.com/sonr-io/sonr/third_party/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) QueryBucket(req mt.QueryWhereIsRequest) (*mt.QueryWhereIsResponse, error) {
	// use the item within the cache from GetWhereIs
	if wi := mtr.Resources.whereIsStore[req.Did]; wi != nil {
		return &mt.QueryWhereIsResponse{
			WhereIs: wi,
		}, nil
	}

	// Query from chain
	resp, err := mtr.GetClient().QueryWhereIs(req.Did, mtr.Address)
	if err != nil {
		return nil, err
	}
	mtr.Resources.StoreWhereIs(resp)
	return &mt.QueryWhereIsResponse{
		WhereIs: resp,
	}, nil
}

func (mtr *motorNodeImpl) QueryBucketGroup(req mt.QueryWhereIsByCreatorRequest) (*mt.QueryWhereIsByCreatorResponse, error) {
	resp, err := mtr.GetClient().QueryWhereIsByCreator(req.Creator)
	var ptrArr []*bt.WhereIs = make([]*bt.WhereIs, 0)
	for _, wi := range resp.WhereIs {
		mtr.Resources.whereIsStore[wi.Did] = &wi
		ptrArr = append(ptrArr, &wi)
	}
	if err != nil {
		return nil, err
	}

	return &mt.QueryWhereIsByCreatorResponse{
		Code:    http.StatusAccepted,
		WhereIs: ptrArr,
	}, nil
}

func (mtr *motorNodeImpl) QueryRegistry(req mt.QueryWhoIsRequest) (*mt.QueryWhoIsResponse, error) {
	resp, err := mtr.GetClient().QueryWhoIs(req.Did)
	if err != nil {
		return nil, err
	}

	return &mt.QueryWhoIsResponse{
		Code:  http.StatusAccepted,
		WhoIs: resp,
	}, nil
}

func (mtr *motorNodeImpl) QuerySchema(req mt.QueryWhatIsRequest) (*mt.QueryWhatIsResponse, error) {
	if wi, _, ok := mtr.Resources.GetSchema(req.Did); ok {
		return &mt.QueryWhatIsResponse{
			Code:   http.StatusAccepted,
			WhatIs: wi,
		}, nil
	}

	resp, err := mtr.GetClient().QueryWhatIs(mtr.GetDID().String(), req.Did)
	if err != nil {
		return nil, err
	}

	// store reference to schema
	_, err = mtr.Resources.StoreWhatIs(resp)
	if err != nil {
		return nil, fmt.Errorf("store WhatIs: %s", err)
	}

	return &mt.QueryWhatIsResponse{
		Code:   http.StatusAccepted,
		WhatIs: mtr.Resources.whatIsStore[req.Did],
	}, nil
}

func (mtr *motorNodeImpl) QuerySchemaGroup(req mt.QueryWhatIsByCreatorRequest) (*mt.QueryWhatIsByCreatorResponse, error) {
	whatIss, err := mtr.GetClient().QueryWhatIsByCreator(req.Creator)

	if err != nil {
		return nil, err
	}
	resp := mt.QueryWhatIsByCreatorResponse{
		Code:   http.StatusAccepted,
		WhatIs: whatIss,
	}

	return &resp, nil
}
