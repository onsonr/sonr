package highway

import (
	context "context"

	bt_v1 "github.com/sonr-io/blockchain/x/bucket/types"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

// CreateBucket creates a new bucket.
func (s *HighwayServer) CreateBucket(ctx context.Context, req *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	resp, err := s.cosmos.BroadcastCreateBucket(btt.NewMsgCreateBucketFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &bt.MsgCreateBucketResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhichIs: btt.NewWhichIsToBuf(resp.WhichIs),
	}, nil
}

// UpdateBucket updates a bucket.
func (s *HighwayServer) UpdateBucket(ctx context.Context, req *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	resp, err := s.cosmos.BroadcastUpdateBucket(btt.NewMsgUpdateBucketFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &bt.MsgUpdateBucketResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhichIs: btt.NewWhichIsToBuf(resp.WhichIs),
	}, nil
}
