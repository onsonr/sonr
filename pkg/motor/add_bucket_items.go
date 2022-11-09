package motor

import (
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bi "github.com/sonr-io/sonr/x/bucket/service"
)

// TODO
func (mtr *motorNodeImpl) AddBucketItems(request mt.AddBucketItemsRequest) (*mt.AddBucketItemsResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	config, err := mtr.fetchBucketConfig(request.Bucket, request.Uuid, request.Creator, request.Name)
	if err != nil {
		return nil, err
	}

	paths, err := bi.WriteBucketItems(mtr.sh, config, mtr.GetAddress(), request.Items...)
	if err != nil {
		return nil, err
	}

	doc, err := mtr.getRegistryDIDDocument()
	if err != nil {
		return nil, err
	}

	return &mt.AddBucketItemsResponse{
		AddedItems:  paths,
		Bucket:      config,
		DidDocument: doc,
		Uri:         config.GetDidService(mtr.GetAddress(), "").ServiceEndpoint.Value[0],
	}, nil
}
