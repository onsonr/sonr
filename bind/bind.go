package motor

import (

	// mtr "github.com/sonr-hq/sonr/pkg/motor"
	// "github.com/sonr-hq/sonr/pkg/motor/x/document"
	// ct "github.com/sonr-hq/sonr/pkg/types/common"
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/sonr-hq/sonr/core/motor"
	mt "github.com/sonr-hq/sonr/core/motor/types/bind/v1"

	// rt "github.com/sonr-hq/sonr/x/registry/types"
	_ "golang.org/x/mobile/bind"
)

var (
	ctx context.Context
	mtr *motor.MotorNode
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(data []byte)
	OnLinking(data []byte)
}

func Init(buf []byte, cb MotorCallback) ([]byte, error) {
	ctx = context.Background()
	// Unmarshal the request
	var req mt.InitializeRequest
	if err := proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	mtdr, err := motor.NewMotorInstance(ctx, &req)
	if err != nil {
		return nil, err
	}
	mtr = mtdr
	return nil, nil
}

func Connect(buf []byte) ([]byte, error) {
	// Unmarshal the request
	var req mt.ConnectRequest
	if err := proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	// // Connect to the network
	resp, err := mtr.Connect(&req)
	if err != nil {
		return nil, err
	}

	// Marshal the response
	return proto.Marshal(resp)
}
