package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/thirdparty/types/common"
	ct "github.com/sonr-io/sonr/thirdparty/types/common"
	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) Query(req mt.QueryRequest) (mt.QueryResponse, error) {
	results := make([]*mt.QueryResultItem, 0)
	switch req.GetModule() {
	case ct.BlockchainModule_SCHEMA:
		i, err := mtr.handleSchemaQuery(req.GetQuery())
		if err != nil {
			return mt.QueryResponse{}, err
		}
		results = append(results, i)
	case ct.BlockchainModule_REGISTRY:
		i, err := mtr.handleRegistryQuery(req.GetQuery())
		if err != nil {
			return mt.QueryResponse{}, err
		}
		results = append(results, i)
	case ct.BlockchainModule_BUCKET:
		if req.GetKind() == common.EntityKind_ADDRESS {

			i, err := mtr.handleBucketGroupQuery(req.GetQuery())
			if err != nil {
				return mt.QueryResponse{}, err
			}
			results = append(results, i)

		} else {
			i, err := mtr.handleBucketQuery(req.GetQuery())
			if err != nil {
				return mt.QueryResponse{}, err
			}
			results = append(results, i)

		}
	}

	return mt.QueryResponse{
		Results: results,
		Query:   req.GetQuery(),
		Module:  req.GetModule(),
	}, nil
}

func (mtr *motorNodeImpl) handleBucketQuery(did string) (*mt.QueryResultItem, error) {
	// use the item within the cache from GetWhereIs
	if wi := mtr.Resources.whereIsStore[did]; wi != nil {
		return &mt.QueryResultItem{
			Kind:    common.EntityKind_DID,
			WhereIs: wi,
		}, nil
	}

	// Query from chain
	resp, err := mtr.GetClient().QueryWhereIs(did, mtr.Address)
	if err != nil {
		return nil, err
	}
	mtr.Resources.whereIsStore[resp.Did] = resp

	return &mt.QueryResultItem{
		Kind:    common.EntityKind_DID,
		WhereIs: resp,
	}, nil
}

func (mtr *motorNodeImpl) handleBucketGroupQuery(addr string) (*mt.QueryResultItem, error) {
	resp, err := mtr.GetClient().QueryWhereIsByCreator(addr)
	var ptrArr []*bt.WhereIs = make([]*bt.WhereIs, len(resp.WhereIs))
	for _, wi := range resp.WhereIs {
		mtr.Resources.whereIsStore[wi.Did] = &wi
		ptrArr = append(ptrArr, &wi)
	}
	if err != nil {
		return nil, err
	}

	return &mt.QueryResultItem{
		Kind:        common.EntityKind_ADDRESS,
		WhereIsList: ptrArr,
	}, nil
}

func (mtr *motorNodeImpl) handleRegistryQuery(did string) (*mt.QueryResultItem, error) {
	resp, err := mtr.GetClient().QueryWhoIs(did)
	if err != nil {
		return nil, err
	}

	return &mt.QueryResultItem{
		Kind:  common.EntityKind_DID,
		WhoIs: resp,
	}, nil
}

func (mtr *motorNodeImpl) handleSchemaQuery(did string) (*mt.QueryResultItem, error) {
	resp, err := mtr.GetClient().QueryWhatIs(mtr.GetAddress(), did)
	if err != nil {
		return nil, err
	}

	// store reference to schema
	_, err = mtr.Resources.StoreWhatIs(resp)
	if err != nil {
		return nil, fmt.Errorf("store WhatIs: %s", err)
	}

	return &mt.QueryResultItem{
		Did:    resp.GetDid(),
		Kind:   common.EntityKind_DID,
		WhatIs: resp,
	}, nil
}
