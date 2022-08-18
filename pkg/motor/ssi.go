package motor

import (
	"errors"
)

func (mtr *motorNodeImpl) AddBucketServiceEndpoint(id string) error {
	if mtr.DIDDocument == nil {
		return errors.New("Document is not defined")
	}
	if _, ok := mtr.Resources.bucketStore[id]; !ok {
		return errors.New("Cannot resolve content for bucket, not found")
	}

	bucket := mtr.Resources.bucketStore[id]
	se := bucket.CreateBucketServiceEndpoint()

	mtr.DIDDocument.AddService(se)

	return nil
}
