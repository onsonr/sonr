package motor

import (
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/internal"
)

// TODO
func (mtr *motorNodeImpl) BurnBucket(request mt.BurnBucketRequest) (*mt.BurnBucketResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	config, err := mtr.fetchBucketConfig(request.Bucket, request.Uuid, request.Creator, request.Name)
	if err != nil {
		return nil, err
	}

	err = internal.PurgeBucketItems(mtr.sh, config, mtr.GetAddress())
	if err != nil {
		return nil, err
	}

	doc, err := mtr.getRegistryDIDDocument()
	if err != nil {
		return nil, err
	}

	return &mt.BurnBucketResponse{
		DidDocument: doc,
		Code:        200,
		Message:     "OK",
	}, nil
}
