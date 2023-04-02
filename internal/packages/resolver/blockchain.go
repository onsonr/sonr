package resolver

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	// "github.com/sonrhq/core/app"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/sonrhq/core/internal/local"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	"google.golang.org/grpc"
)

type BroadcastTxResponse = txtypes.BroadcastTxResponse

// GetDID returns the DID document with the given id
func GetDID(ctx context.Context, id string) (*identitytypes.ResolvedDidDocument, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {

		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Did(ctx, &identitytypes.QueryGetDidRequest{Did: id})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetAllDIDs returns all DID documents
func GetAllDIDs(ctx context.Context) ([]*identitytypes.DidDocument, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {

		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).DidAll(ctx, &identitytypes.QueryAllDidRequest{})
	if err != nil {

		return nil, err
	}
	list := make([]*identitytypes.DidDocument, len(resp.DidDocument))
	for i, d := range resp.DidDocument {
		list[i] = &d
	}
	return list, nil
}

// GetService returns the service with the given id
func GetService(ctx context.Context, origin string) (*identitytypes.Service, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {

		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Service(ctx, &identitytypes.QueryGetServiceRequest{Origin: origin})
	if err != nil {

		return nil, err
	}
	return &resp.Service, nil
}

// GetAllServices returns all services
func GetAllServices(ctx context.Context) ([]*identitytypes.Service, error) {
	snrctx := local.NewContext()
	conn, err := grpc.Dial(snrctx.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {

		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).ServiceAll(ctx, &identitytypes.QueryAllServiceRequest{})
	if err != nil {

		return nil, err
	}
	list := make([]*identitytypes.Service, len(resp.Services))
	for i, s := range resp.Services {
		list[i] = &s
	}
	return list, nil
}


// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func  BroadcastTx(txRawBytes []byte) (*BroadcastTxResponse, error) {
		snrctx := local.NewContext()
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		snrctx.GrpcEndpoint(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.

	txClient := txtypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txRawBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return nil, err
	}
	if grpcRes == nil {
		return nil, fmt.Errorf("no response from broadcast tx")
	}
	if grpcRes.GetTxResponse() != nil && grpcRes.GetTxResponse().Code != 0 {
		return nil, fmt.Errorf("failed to broadcast transaction: %s", grpcRes.GetTxResponse().RawLog)
	}
	return grpcRes, nil
}

// SimulateTx simulates a transaction on the Sonr blockchain network
func SimulateTx(txRawBytes []byte) (*txtypes.SimulateResponse, error) {
			snrctx := local.NewContext()
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		snrctx.RpcEndpoint(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := txtypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.Simulate(
		context.Background(),
		&txtypes.SimulateRequest{
			TxBytes: txRawBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return nil, err
	}
	return grpcRes, nil
}

func DecodeTxResponseData(d string, v proto.Unmarshaler) error {
	data, err := hex.DecodeString(d)
	if err != nil {
		return err
	}

	anyWrapper := new(types.Any)
	if err := proto.Unmarshal(data, anyWrapper); err != nil {
		return err
	}

	// TODO: figure out if there's a better 'cosmos' way of doing this
	// you have to unwrap the Any twice, and the first time the bytes get decoded
	// in the 'TypeUrl' field instead of 'Value' field
	any := new(types.Any)
	if err := proto.Unmarshal([]byte(anyWrapper.TypeUrl), any); err != nil {
		return err
	}

	return v.Unmarshal(any.Value)
}
