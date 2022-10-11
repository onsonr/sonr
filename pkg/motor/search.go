package motor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sonr-io/sonr/pkg/did"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
)

/*
Stop gap implementation for
*/
func (mtr *motorNodeImpl) SeachBucketBySchema(req mt.SeachBucketContentBySchemaRequest) (mt.SearchBucketContentBySchemaResponse, error) {
	_, err := did.ParseDID(req.BucketDid)
	if err != nil {
		return mt.SearchBucketContentBySchemaResponse{}, fmt.Errorf("cannot parse did: %s", req.BucketDid)
	}

	_, err = did.ParseDID(req.SchemaDid)

	if err != nil {
		return mt.SearchBucketContentBySchemaResponse{}, fmt.Errorf("cannot parse did: %s", req.BucketDid)
	}

	if _, ok := mtr.Resources.bucketStore[req.BucketDid]; !ok {
		b, err := mtr.GetBucket(req.BucketDid)
		if b.GetVisibility() == types.BucketVisibility_PRIVATE && b.GetCreator() != mtr.Address {
			return mt.SearchBucketContentBySchemaResponse{}, fmt.Errorf("creator address does not match session creator: %s", mtr.Address)
		}
		if err != nil {
			return mt.SearchBucketContentBySchemaResponse{}, fmt.Errorf("error while querying WhereIs for bucket: %s err: %s", req.BucketDid, err)
		}
	}

	b := mtr.Resources.bucketStore[req.BucketDid]

	res, err := b.ResolveContentBySchema(req.SchemaDid)

	if err != nil {
		return mt.SearchBucketContentBySchemaResponse{}, nil
	}

	var contentBytes [][]byte = make([][]byte, 0)

	for _, item := range res {
		b, err := json.Marshal(item)
		if err != nil {
			return mt.SearchBucketContentBySchemaResponse{}, err
		}
		contentBytes = append(contentBytes, b)
	}

	return mt.SearchBucketContentBySchemaResponse{
		Status:    http.StatusAccepted,
		BucketDid: req.BucketDid,
		SchemaDid: req.SchemaDid,
		Content:   contentBytes,
	}, nil
}
