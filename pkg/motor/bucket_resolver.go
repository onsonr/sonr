package motor

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/motor/types"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) NewBucketResolver(context context.Context, creator string, did string) (bucket.Bucket, error) {
	if creator == "" || did == "" {
		return nil, errors.New("creator address and did must be defined within the request")
	}
	addr, err := mtr.Wallet.Address()
	if err != nil {
		return nil, err
	}

	if _, ok := mtr.Resources.whereIsStore[did]; !ok {
		wiReq, err := mtr.QueryWhereIs(context, types.QueryWhereIsRequest{
			Creator: creator,
			Did:     did,
		})

		if err != nil {
			return nil, err
		}

		mtr.Resources.whereIsStore[did] = wiReq.WhereIs
	}
	wi := mtr.Resources.whereIsStore[did]
	s := mtr.Resources.shell
	bq := mtr.Resources.bucketQueryClient

	b := bucket.New(addr, wi, s, bq)

	mtr.Resources.bucketStore[did] = b

	return b, nil
}

func (mtr *motorNodeImpl) ResolveBucketsForBucket(did string) error {
	addr, err := mtr.Wallet.Address()

	if err != nil {
		return err
	}

	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return errors.New("Cannot resolve content for bucket, not found")
	}

	bucket := mtr.Resources.bucketStore[did]
	err = bucket.ResolveBuckets(addr)

	return err
}

func (mtr *motorNodeImpl) ResolveContentForBucket(did string) error {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return errors.New("Cannot resolve content for bucket, not found")
	}

	bucket := mtr.Resources.bucketStore[did]
	err := bucket.ResolveContent()

	return err
}

func (mtr *motorNodeImpl) GetBucketItems(did string) ([]*bt.BucketItem, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("Cannot resolve content for bucket, not found")
	}

	bucket := mtr.Resources.bucketStore[did]
	items := bucket.GetBucketItems()

	return items, nil
}

func (mtr *motorNodeImpl) GetBucketContent(did string, item *bt.BucketItem) (*bucket.BucketContent, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("Cannot resolve content for bucket, not found")
	}
	bucket := mtr.Resources.bucketStore[did]
	content, err := bucket.GetContentById(item.Uri)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (mtr *motorNodeImpl) GetAllBucketContent(did string) ([]*bucket.BucketContent, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("Cannot resolve content for bucket, not found")
	}
	b := mtr.Resources.bucketStore[did]
	var bc []*bucket.BucketContent
	items := b.GetBucketItems()
	for _, item := range items {
		content, err := b.GetContentById(item.Uri)
		if err != nil {
			return nil, err
		}

		bc = append(bc, content)
	}

	return bc, nil
}

func (mtr *motorNodeImpl) UpdateBucketItems(context context.Context, did string, items []*bt.BucketItem) (bucket.Bucket, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("Cannot resolve content for bucket, not found")
	}

	wi := mtr.Resources.whereIsStore[did]

	bi := make([]*bt.BucketItem, len(wi.Content))

	copy(wi.Content, bi)

	bi = append(bi, items...)

	req := types.UpdateBucketRequest{
		Creator:    wi.Creator,
		Did:        wi.Did,
		Label:      wi.Label,
		Role:       wi.Role,
		Visibility: wi.Visibility,
		Content:    bi,
	}
	b, err := mtr.UpdateBucket(req)

	if err != nil {
		return nil, err
	}

	return b, nil
}
