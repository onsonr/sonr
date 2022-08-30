package motor

import (
	"errors"
)

/*
	Adds a reference to the bucket as a service endpoint on the registered did document
	Such functionality might be better on chain as to keep transactions in a single block.
*/
func (mtr *motorNodeImpl) AddBucketServiceEndpoint(baseURI, id string) error {

	if mtr.DIDDocument == nil {
		return errors.New("Document is not defined")
	}
	if _, ok := mtr.Resources.bucketStore[id]; !ok {
		return errors.New("Cannot resolve content for bucket, not found")
	}

	bucket := mtr.Resources.bucketStore[id]
	se := bucket.CreateBucketServiceEndpoint()

	if mtr.DIDDocument.GetServices().FindByID(se.ID) == nil {
		mtr.DIDDocument.AddService(se)
	}

	_, err := updateWhoIs(mtr)
	if err != nil {
		return err
	}
	return nil
}
