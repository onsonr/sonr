package types

import (
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

func NewWhichIsFromBuf(cd *bt.WhichIs) *WhichIs {
	return &WhichIs{
		Did:       cd.GetDid(),
		Creator:   cd.GetCreator(),
		Timestamp: cd.GetTimestamp(),
		IsActive:  cd.GetIsActive(),
		Bucket:    NewBucketDocFromBuf(cd.GetBucket()),
	}
}

func NewWhichIsToBuf(cd *WhichIs) *bt.WhichIs {
	return &bt.WhichIs{
		Did:       cd.GetDid(),
		Creator:   cd.GetCreator(),
		Timestamp: cd.GetTimestamp(),
		IsActive:  cd.GetIsActive(),
		Bucket:    NewBucketDocToBuf(cd.GetBucket()),
	}
}
