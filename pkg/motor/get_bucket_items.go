package motor

import (
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bi "github.com/sonr-io/sonr/x/bucket/internal"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) GetBucketItems(request mt.GetBucketItemsRequest) (*mt.GetBucketItemsResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	config, err := mtr.fetchBucketConfig(request.Bucket, request.Uuid, request.Creator, request.Name)
	if err != nil {
		return nil, err
	}

	var items []*common.BucketItem
	itemsW, err := bi.DownloadBucket(mtr.sh, config, mtr.GetAddress())
	if err != nil {
		return nil, err
	}

	for _, item := range itemsW {
		i, err := itemWrapperToBucketItem(item)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	doc, err := mtr.getRegistryDIDDocument()
	if err != nil {
		return nil, err
	}

	return &mt.GetBucketItemsResponse{
		Bucket:      config,
		DidDocument: doc,
		Items:       items,
		Uri:         config.GetDidService(mtr.GetAddress()).ServiceEndpoint,
	}, nil
}

func itemWrapperToBucketItem(item bt.ItemWrapper) (*common.BucketItem, error) {
	var i common.BucketItem
	err := i.Unmarshal(item.Content())
	if err != nil {
		return nil, err
	}
	return &i, nil
}
