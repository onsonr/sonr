package highway

import (
	context "context"
	"errors"

	otv1 "github.com/sonr-io/blockchain/x/object/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Verify that object fields are not nil
	if req.GetInitialFields() == nil {
		return nil, errors.New("object to register must have fields")
	}

	// Translate ot fields to otv1
	fields := s.bufToBlockchain(req.GetInitialFields())

	// Build Transaction
	tx := &otv1.MsgCreateObject{
		Creator:       req.GetCreator(),
		Label:         req.GetLabel(),
		Description:   req.GetDescription(),
		InitialFields: fields,
	}

	// Broadcast the message
	res, err := s.cosmos.BroadcastCreateObject(tx)
	if err != nil {
		return nil, err
	}
	//v.GetValue()
	// decode otv1 fields to ot
	outputFields := make(map[string]*ot.ObjectField, len(res.WhatIs.ObjectDoc.Fields))
	for k, v := range res.WhatIs.ObjectDoc.Fields {
		outputFields[k] = &ot.ObjectField{
			Label: v.GetLabel(),
			Type:  ot.ObjectFieldType(v.GetType()),
			Did:   v.GetDid(),
			//Value:    v.Value,
			Metadata: v.GetMetadata(),
		}
	}

	return &ot.MsgCreateObjectResponse{
		Code:    res.GetCode(),
		Message: res.GetMessage(),
		WhatIs: &ot.WhatIs{
			Did: res.WhatIs.GetDid(),
			ObjectDoc: &ot.ObjectDoc{
				Label:       res.WhatIs.ObjectDoc.GetLabel(),
				Description: res.WhatIs.ObjectDoc.GetDescription(),
				Did:         res.WhatIs.ObjectDoc.GetDid(),
				BucketDid:   res.WhatIs.ObjectDoc.GetBucketDid(),
				Fields:      outputFields,
			},
			Creator:   res.WhatIs.GetCreator(),
			Timestamp: res.WhatIs.GetTimestamp(),
			IsActive:  res.WhatIs.GetIsActive(),
		},
	}, nil
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// // DeleteObject deletes an object.
// func (s *HighwayServer) DeleteObject(ctx context.Context, req *ot.MsgDeleteObject) (*ot.MsgDeleteObjectResponse, error) {
// 	return nil, ErrMethodUnimplemented
// }

// -----------------
// Helper Functions
// -----------------

// Translate ot (buf) fields to otv1 (blockchain) //TODO rename this
func (s *HighwayServer) bufToBlockchain(bufFields []*ot.ObjectField) []*otv1.ObjectField {
	var blockchainFields []*otv1.ObjectField
	for _, v := range bufFields {
		item := &otv1.ObjectField{
			Label: v.Label,
			Type:  otv1.ObjectFieldType(v.Type),
			Did:   v.Did,
			// Value: v.Value,
		}
		blockchainFields = append(blockchainFields, item)
	}

	return blockchainFields
}

// Translate otv1 (blockchain) fields to ot (buf) //TODO rename this
func (s *HighwayServer) blockchainToBuf(blockchainFields []*otv1.ObjectField) []*ot.ObjectField {
	var bufFields []*ot.ObjectField
	for _, v := range blockchainFields {
		item := &ot.ObjectField{
			Label: v.Label,
			Type:  ot.ObjectFieldType(v.Type),
			Did:   v.Did,
			// Value:    v.GetValue(),
			Metadata: v.GetMetadata(),
		}
		bufFields = append(bufFields, item)
	}

	return bufFields
}
