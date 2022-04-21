package highway

import (
	context "context"
	"fmt"

	otv1 "github.com/sonr-io/blockchain/x/object/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Create ctv1 message to broadcast
	fields := make([]*otv1.TypeField, len(req.GetInitialFields()))
	for i, f := range req.GetInitialFields() {
		fields[i] = &otv1.TypeField{
			Name: f.GetName(),
			Kind: otv1.TypeKind(f.GetKind()),
		}
	}

	// Build Transaction
	tx := &otv1.MsgCreateObject{
		Creator:       req.GetCreator(),
		Label:         req.GetLabel(),
		InitialFields: fields,
	}
	resp, err := s.cosmos.BroadcastCreateObject(tx)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	return &ot.MsgCreateObjectResponse{}, nil
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	return nil, ErrMethodUnimplemented
}

// // DeleteObject deletes an object.
// func (s *HighwayServer) DeleteObject(ctx context.Context, req *ot.MsgDeleteObject) (*ot.MsgDeleteObjectResponse, error) {
// 	return nil, ErrMethodUnimplemented
// }
