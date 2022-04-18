package highway

import (
	context "context"
	"log"

	btt "github.com/sonr-io/blockchain/x/bucket/types"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

// CreateBucket creates a new bucket.
func (s *HighwayServer) CreateBucket(ctx context.Context, req *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	tx := &btt.MsgCreateBucket{
		Creator:           req.GetCreator(),
		Label:             req.GetLabel(),
		Description:       req.GetDescription(),
		Kind:              req.GetKind(),
		InitialObjectDids: req.GetInitialObjectDids(),
	}
	resp, err := s.cosmos.BroadcastCreateBucket(tx)
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return nil, nil
}

// UpdateBucket updates a bucket.
func (s *HighwayServer) UpdateBucket(ctx context.Context, req *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}
