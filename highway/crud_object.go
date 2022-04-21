package highway

import (
	context "context"
	"fmt"

	otv1 "github.com/sonr-io/blockchain/x/object/types"
	ot "go.buf.build/sonr-io/grpc-gateway/sonr-io/blockchain/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.cosmos.BroadcastCreateObject(otv1.NewMsgCreateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())

	// Upload Object Schema to IPFS
	cid, err := s.ipfsProtocol.PutObjectSchema(resp.GetWhatIs().GetObjectDoc())
	if err != nil {
		return nil, err
	}
	fmt.Println(cid)
	return &ot.MsgCreateObjectResponse{}, nil
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.cosmos.BroadcastUpdateObject(otv1.NewMsgUpdateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	return &ot.MsgUpdateObjectResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhatIs:  otv1.NewWhatIsToBuf(resp.WhatIs),
	}, nil
}
