package highway

import (
	context "context"
	"log"

	bt_v1 "github.com/sonr-io/blockchain/x/bucket/types"
	bt "go.buf.build/grpc/go/sonr-io/blockchain/bucket"
)

// CreateBucket creates a new bucket.
func (s *HighwayServer) CreateBucket(ctx context.Context, req *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	tx := &bt_v1.MsgCreateBucket{
		Creator:           req.GetCreator(),
		Label:             req.GetLabel(),
		Description:       req.GetDescription(),
		Kind:              req.GetKind(),
		InitialObjectDids: req.GetInitialObjectDids(),
		Session:           s.regSessToTypeSess(*req.GetSession()),
	}
	resp, err := s.cosmos.BroadcastCreateBucket(tx)
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &bt.MsgCreateBucketResponse{
		Code:    resp.GetCode(),
		Message: resp.GetMessage(),
		WhichIs: &bt.WhichIs{
			Did:     resp.WhichIs.GetDid(),
			Creator: resp.WhichIs.GetCreator(),
			Bucket: &bt.BucketDoc{
				Label:       resp.WhichIs.Bucket.GetLabel(),
				Description: resp.WhichIs.Bucket.GetDescription(),
				Type:        bt.BucketType(resp.WhichIs.Bucket.GetType()),
				Did:         resp.WhichIs.GetDid(),
				ObjectDids:  resp.WhichIs.Bucket.GetObjectDids(),
			},
			Timestamp: resp.WhichIs.Timestamp,
			IsActive:  resp.WhichIs.IsActive,
		},
	}, nil
}

// UpdateBucket updates a bucket.
func (s *HighwayServer) UpdateBucket(ctx context.Context, req *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	return nil, ErrMethodUnimplemented
}
