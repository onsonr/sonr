package motor

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/motor/types"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) NewBucketResolver(context context.Context, did string) (bucket.Bucket, error) {
	if did == "" {
		return nil, errors.New("creator address and did must be defined within the request")
	}
	addr, err := mtr.Wallet.Address()
	if err != nil {
		return nil, err
	}

	if _, ok := mtr.Resources.whereIsStore[did]; !ok {
		wiReq, err := mtr.QueryWhereIs(context, types.QueryWhereIsRequest{
			Creator: addr,
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

func (mtr *motorNodeImpl) GetBucket(context context.Context, did string) (bucket.Bucket, error) {
	addr := mtr.GetAddress()
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		wi, err := mtr.QueryWhereIs(context, types.QueryWhereIsRequest{
			Creator: addr,
			Did:     did,
		})

		if err != nil {
			return nil, err
		}
		b := bucket.New(addr, wi.WhereIs, mtr.Resources.shell, mtr.bucketQueryClient)

		err = b.ResolveBuckets(addr)
		if err != nil {
			return nil, err
		}
		err = b.ResolveContent()
		if err != nil {
			return nil, err
		}

		mtr.Resources.bucketStore[did] = b
		for _, sb := range b.GetBuckets() {
			mtr.Resources.bucketStore[sb.Id] = sb.Item.(bucket.Bucket)
		}
	}

	return mtr.Resources.bucketStore[did], nil
}

func (mtr *motorNodeImpl) GetBuckets(context context.Context) ([]bucket.Bucket, error) {

	addr := mtr.GetAddress()
	res, err := mtr.Resources.bucketQueryClient.WhereIsByCreator(context, &bt.QueryGetWhereIsByCreatorRequest{
		Creator:    addr,
		Pagination: nil,
	})

	if err != nil {
		return nil, err
	}
	var buckets []bucket.Bucket = make([]bucket.Bucket, len(res.WhereIs))
	for _, wi := range res.WhereIs {
		did := wi.Did
		if _, ok := mtr.Resources.bucketStore[did]; !ok {
			wi, err := mtr.QueryWhereIs(context, types.QueryWhereIsRequest{
				Creator: addr,
				Did:     did,
			})

			if err != nil {
				return nil, err
			}
			b := bucket.New(addr, wi.WhereIs, mtr.Resources.shell, mtr.bucketQueryClient)

			err = b.ResolveBuckets(addr)
			if err != nil {
				return nil, err
			}
			err = b.ResolveContent()
			if err != nil {
				return nil, err
			}

			mtr.Resources.bucketStore[did] = b
			for _, sb := range b.GetBuckets() {
				mtr.Resources.bucketStore[sb.Id] = sb.Item.(bucket.Bucket)
			}
		}
		buckets = append(buckets, mtr.Resources.bucketStore[did])
	}

	return buckets, nil
}
