package motor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sonr-io/sonr/pkg/did"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/x/bucket/types"
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

func (mtr *motorNodeImpl) QueryBucketBySchema(ctx context.Context, req mt.QueryBucketContentBySchemaRequest) (mt.QueryBucketContentBySchemaResponse, error) {
	_, err := did.ParseDID(req.BucketDid)
	if err != nil {
		return mt.QueryBucketContentBySchemaResponse{}, fmt.Errorf("cannot parse did: %s", req.BucketDid)
	}

	_, err = did.ParseDID(req.SchemaDid)

	if err != nil {
		return mt.QueryBucketContentBySchemaResponse{}, fmt.Errorf("cannot parse did: %s", req.BucketDid)
	}

	if _, ok := mtr.Resources.bucketStore[req.BucketDid]; !ok {
		b, err := mtr.QueryWhereIs(ctx, req.BucketDid)
		if b.WhereIs.Visibility == types.BucketVisibility_PRIVATE && b.WhereIs.Creator != mtr.Address {
			return mt.QueryBucketContentBySchemaResponse{}, fmt.Errorf("creator address does not match session creator: %s", mtr.Address)
		}
		if err != nil {
			return mt.QueryBucketContentBySchemaResponse{}, fmt.Errorf("error while querying WhereIs for bucket: %s err: %s", req.BucketDid, err)
		}
	}

	b := mtr.Resources.bucketStore[req.BucketDid]

	res, err := b.ResolveContentBySchema(req.SchemaDid)

	if err != nil {
		return mt.QueryBucketContentBySchemaResponse{}, nil
	}

	var contentBytes [][]byte = make([][]byte, 0)

	for _, item := range res {
		b, err := json.Marshal(item)
		if err != nil {
			return mt.QueryBucketContentBySchemaResponse{}, err
		}
		contentBytes = append(contentBytes, b)
	}

	return mt.QueryBucketContentBySchemaResponse{
		Status:    http.StatusAccepted,
		BucketDid: req.BucketDid,
		SchemaDid: req.SchemaDid,
		Content:   contentBytes,
	}, nil
}
