package motor

import (
	"context"

	"github.com/sonr-io/sonr/internal/bucket"
)

func (mtr *motorNodeImpl) GetBucket(did string) (bucket.Bucket, error) {
	addr := mtr.GetAddress()
	if _, ok := mtr.Resources.whereIsStore[did]; !ok {
		_, err := mtr.QueryWhereIs(did)
		wi := mtr.Resources.whereIsStore[did]

		if err != nil {
			return nil, err
		}
		b := bucket.New(addr, wi, mtr.Resources.shell, mtr.GetClient().GetRPCAddress())

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
			b := bucket.New(addr, wi, mtr.Resources.shell, mtr.GetClient().GetRPCAddress())

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
