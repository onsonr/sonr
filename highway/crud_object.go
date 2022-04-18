package highway

import (
	context "context"
	"errors"

	ot_v1 "github.com/sonr-io/blockchain/x/object/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Verify that object fields are not nil
	if req.GetInitialFields() == nil {
		return nil, errors.New("object to register must have fields")
	}

	// Translate ot fields to ot_v1
	fields := s.bufToBlockchain(req.GetInitialFields())

	// Build Transaction
	tx := &ot_v1.MsgCreateObject{
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

	// decode ot_v1 fields to ot
	var outputFields map[string]*ot.ObjectField
	if res.WhatIs.ObjectDoc.Fields != nil {
		outputFields = s.decodeObjectDocFields(res.WhatIs.ObjectDoc.Fields)
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
	// Verify that there are requested changes
	if req.GetAddedFields() == nil && req.GetRemovedFields() == nil {
		return nil, errors.New("object to register must have fields")
	}

	// Translate ot fields to ot_v1
	addedFields := s.bufToBlockchain(req.GetAddedFields())
	removedFields := s.bufToBlockchain(req.GetRemovedFields())

	// Build Transaction
	tx := &ot_v1.MsgUpdateObject{
		Creator:       req.GetCreator(),
		Label:         req.GetLabel(),
		AddedFields:   addedFields,
		RemovedFields: removedFields,
	}

	// Broadcast the message
	res, err := s.cosmos.BroadcastUpdateObject(tx)
	if err != nil {
		return nil, err
	}

	// decode ot_v1 fields to ot
	var outputFields map[string]*ot.ObjectField
	if res.WhatIs.ObjectDoc.Fields != nil {
		outputFields = s.decodeObjectDocFields(res.WhatIs.ObjectDoc.Fields)
	}

	return &ot.MsgUpdateObjectResponse{
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

// DeleteObject deletes an object.
func (s *HighwayServer) DeactivateObject(ctx context.Context, req *ot.MsgDeactivateObject) (*ot.MsgDeactivateObjectResponse, error) {
	// Build Transaction
	tx := &ot_v1.MsgDeactivateObject{
		Creator: req.GetCreator(),
		Did:     req.GetDid(),
	}

	// Broadcast the message
	res, err := s.cosmos.BroadcastDeactivateObject(tx)
	if err != nil {
		return nil, err
	}

	return &ot.MsgDeactivateObjectResponse{
		Code:    res.GetCode(),
		Message: res.GetMessage(),
	}, nil
}

// -----------------
// Helper Functions
// -----------------

// Translate ot (buf) fields to ot_v1 (blockchain) //TODO rename this
func (s *HighwayServer) bufToBlockchain(bufFields []*ot.ObjectField) []*ot_v1.ObjectField {
	var blockchainFields []*ot_v1.ObjectField
	for _, v := range bufFields {
		item := &ot_v1.ObjectField{
			Label: v.Label,
			Type:  ot_v1.ObjectFieldType(v.Type),
			Did:   v.Did,
			// Value: v.Value,
		}
		blockchainFields = append(blockchainFields, item)
	}

	return blockchainFields
}

// Translate the objectDoc ot_v1 (blockchain) fields to ot (buf) //TODO rename this
func (s *HighwayServer) decodeObjectDocFields(blockchainFields map[string]*ot_v1.ObjectField) map[string]*ot.ObjectField {
	// decode ot_v1 fields to ot

	if blockchainFields == nil || len(blockchainFields) < 1 {
		return nil
	}

	outputFields := make(map[string]*ot.ObjectField, len(blockchainFields))
	for k, v := range blockchainFields {
		outputFields[k] = &ot.ObjectField{
			Label: v.GetLabel(),
			Type:  ot.ObjectFieldType(v.GetType()),
			Did:   v.GetDid(),
			//Value:    v.Value,
			Metadata: v.GetMetadata(),
		}
	}

	return outputFields
}
