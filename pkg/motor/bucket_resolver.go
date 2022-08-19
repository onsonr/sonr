package motor

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/internal/bucket"
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
		err := mtr.QueryWhereIs(context, did)
		if err != nil {
			return nil, err
		}
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
	if _, ok := mtr.Resources.whereIsStore[did]; !ok {
		err := mtr.QueryWhereIs(context, did)
		wi := mtr.Resources.whereIsStore[did]

		if err != nil {
			return nil, err
		}
		b := bucket.New(addr, wi, mtr.Resources.shell, mtr.bucketQueryClient)

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

/*
	Takes the whereIs store and checks for a matching bucket in the cache, if its not present it will create it and get its sub buckets
	Does not query for new buckets, only respects what is currently present in the store
*/
func (mtr *motorNodeImpl) GetBuckets(context context.Context) ([]bucket.Bucket, error) {
	addr := mtr.GetAddress()

	var buckets []bucket.Bucket = make([]bucket.Bucket, len(mtr.Resources.whereIsStore))
	for _, wi := range mtr.Resources.whereIsStore {
		did := wi.Did
		if _, ok := mtr.Resources.bucketStore[did]; !ok {
			b := bucket.New(addr, wi, mtr.Resources.shell, mtr.bucketQueryClient)

			err := b.ResolveBuckets(addr)
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
