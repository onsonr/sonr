package motor

import (
	"fmt"
	"net/http"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) QueryWhoIs(req mt.QueryWhoIsRequest) (*mt.QueryWhoIsResponse, error) {
	resp, err := mtr.GetClient().QueryWhoIs(req.Did)
	if err != nil {
		return nil, err
	}

	return &mt.QueryWhoIsResponse{
		Code:  http.StatusAccepted,
		WhoIs: resp,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhatIs(req mt.QueryWhatIsRequest) (*mt.QueryWhatIsResponse, error) {
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
	schema, err := mtr.Resources.StoreWhatIs(resp)
	if err != nil {
		return nil, fmt.Errorf("store WhatIs: %s", err)
	}

	return &mt.QueryWhatIsResponse{
		Code:   http.StatusAccepted,
		WhatIs: mtr.Resources.whatIsStore[req.Did],
		Schema: schema,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhatIsByCreator(req mt.QueryWhatIsByCreatorRequest) (*mt.QueryWhatIsByCreatorResponse, error) {
	resp, err := mtr.GetClient().QueryWhatIsByCreator(req.Creator, req.Pagination)
	if err != nil {
		return nil, err
	}

	// store reference to schema
	schemas := make(map[string]*st.SchemaDefinition)
	for _, w := range resp {
		def := w.Schema
		if err != nil {
			return nil, fmt.Errorf("store WhatIs: %s", err)
		}
		schemas[w.Schema.Did] = def
	}

	return &mt.QueryWhatIsByCreatorResponse{
		Code:    http.StatusAccepted,
		WhatIs:  resp,
		Schemas: schemas,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhatIsByDid(did string) (*mt.QueryWhatIsResponse, error) {
	resp, err := mtr.GetClient().QueryWhatIsByDid(did)
	if err != nil {
		return nil, err
	}

	s, err := mtr.Resources.StoreWhatIs(resp)
	if err != nil {
		return nil, fmt.Errorf("store WhatIs: %s", err)
	}

	return &mt.QueryWhatIsResponse{
		Code:   http.StatusOK,
		WhatIs: resp,
		Schema: s,
	}, nil
}

func (mtr *motorNodeImpl) QueryWhereIs(req mt.QueryWhereIsRequest) (*mt.QueryWhereIsResponse, error) {
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

func (mtr *motorNodeImpl) QueryWhereIsByCreator(req mt.QueryWhereIsByCreatorRequest) (*mt.QueryWhereIsByCreatorResponse, error) {
	resp, err := mtr.GetClient().QueryWhereIsByCreator(req.Creator, req.Pagination)
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

func (mtr *motorNodeImpl) QueryObject(cid string) (map[string]interface{}, error) {
	var dag map[string]interface{}
	err := mtr.sh.DagGet(cid, &dag)
	return dag, err
}
