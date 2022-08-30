package motor

import (
	"context"

	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func ResolveBucket(did string) error {
	if instance == nil {
		return errWalletNotExists
	}

	bucket, err := instance.GetBucket(did)
	if err != nil {
		return err
	}

	if err = bucket.ResolveContent(); err != nil {
		return err
	}
	if err = bucket.ResolveBuckets(); err != nil {
		return err
	}

	return nil
}

func GetBucketObject(bucketDid, cid string) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	bucket, err := instance.GetBucket(bucketDid)
	if err != nil {
		return nil, err
	}

	c, err := bucket.GetContentById(cid)
	if err != nil {
		return nil, err
	}

	return c.Marshal()
}

func GetBucketObjects(bucketDid string) ([][]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	bucket, err := instance.GetBucket(bucketDid)
	if err != nil {
		return nil, err
	}

	result := make([][]byte, 0)
	items := bucket.GetBucketItems()
	for _, item := range items {
		c, err := bucket.GetContentById(item.Uri)
		if err != nil {
			return nil, err
		}

		if c.ContentType != bt.ResourceIdentifier_CID {
			continue
		}

		if b, err := c.Marshal(); err == nil {
			result = append(result, b)
		} else {
			return nil, err
		}
	}

	return result, nil
}

func ResolveSubBucket(bucketDid, subBucketDid string) error {
	if instance == nil {
		return errWalletNotExists
	}

	bucket, err := instance.GetBucket(bucketDid)
	if err != nil {
		return err
	}

	for _, b := range bucket.GetBuckets() {
		if err = b.ResolveContent(); err != nil {
			return err
		}
		if err = b.ResolveBuckets(); err != nil {
			return err
		}
	}

	return nil
}

func UpdateBucketLabel(bucketDid, label string) error {
	if instance == nil {
		return errWalletNotExists
	}

	_, err := instance.UpdateBucketLabel(context.Background(), bucketDid, label)

	return err
}

func UpdateBucketVisibility(bucketDid string, vis bt.BucketVisibility) error {
	if instance == nil {
		return errWalletNotExists
	}

	_, err := instance.UpdateBucketVisibility(context.Background(), bucketDid, vis)

	return err
}

func AddBucketObject(bucketDid string, obj []byte) error {
	if instance == nil {
		return errWalletNotExists
	}

	newObject := &bt.BucketItem{}
	if err := newObject.Unmarshal(obj); err != nil {
		return err
	}

	b, err := instance.GetBucket(bucketDid)
	if err != nil {
		return err
	}

	newObjects := append(b.GetBucketItems(), newObject)
	_, err = instance.UpdateBucketItems(context.Background(), bucketDid, newObjects)

	return err
}

func RemoveBucketObject(bucketDid, cid string) error {
	if instance == nil {
		return errWalletNotExists
	}

	b, err := instance.GetBucket(bucketDid)
	if err != nil {
		return err
	}

	items := b.GetBucketItems()
	newObjects := make([]*bt.BucketItem, len(items)-1)
	var skip int
	for i, item := range items {
		if item.Uri == cid {
			skip += 1
			continue
		}

		newObjects[i-skip] = item
	}

	_, err = instance.UpdateBucketItems(context.Background(), bucketDid, newObjects)
	return err
}
