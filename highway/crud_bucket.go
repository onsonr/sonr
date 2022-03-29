package highway

import (
	context "context"

	bt "go.buf.build/grpc/go/sonr-io/sonr/bucket"
)

// CreateBucket creates a new bucket.
func (s *HighwayServer) CreateBucket(ctx context.Context, req *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// ReadBucket reads a bucket.
func (s *HighwayServer) ReadBucket(ctx context.Context, req *bt.MsgReadBucket) (*bt.MsgReadBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// UpdateBucket updates a bucket.
func (s *HighwayServer) UpdateBucket(ctx context.Context, req *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteBucket deletes a bucket.
func (s *HighwayServer) DeleteBucket(ctx context.Context, req *bt.MsgDeleteBucket) (*bt.MsgDeleteBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}
