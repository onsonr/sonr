package client

import (
	"context"
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	btx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"google.golang.org/grpc"
)

type signature []byte

type buildTxOpts struct {
	gas        sdk.Coins
	protoCodec codec.ProtoCodecMarshaler
	signModes  []signing.SignMode
	signatures []signature
}

func defaultBuildTxOpts() *buildTxOpts {
	return &buildTxOpts{
		gas:        sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(2))),
		signModes:  tx.DefaultSignModes,
		signatures: nil,
	}
}

type TxOption func(*buildTxOpts)

// WithGas sets the gas limit for the transaction. default is 2 snr
func WithGas(gas int64) TxOption {
	return func(opts *buildTxOpts) {
		opts.gas = sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(gas)))
	}
}

// WithSignMode sets the sign mode for the transaction. default is tx.DefaultSignModes
func WithSignMode(signMode signing.SignMode) TxOption {
	return func(opts *buildTxOpts) {
		opts.signModes = append(opts.signModes, signMode)
	}
}

// WithSignatures sets the signatures for the transaction.
func WithSignatures(signatures ...signature) TxOption {
	return func(opts *buildTxOpts) {
		opts.signatures = signatures
	}
}

// BuildTx builds a transaction on the Sonr blockchain network and returns the bytes of the signed transaction
func BuildTx(msg sdk.Msg, opts ...TxOption) ([]byte, error) {
	// Configure tx build options.
	buildOpts := defaultBuildTxOpts()
	for _, opt := range opts {
		opt(buildOpts)
	}

	if buildOpts.signatures == nil {
		return nil, errors.New("no signatures provided")
	}

	// Create a new transaction builder.
	txConfig := tx.NewTxConfig(buildOpts.protoCodec, buildOpts.signModes)
	txBuilder := txConfig.NewTxBuilder()

	// Add the message to the transaction and set the gas limit.
	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return nil, err
	}
	txBuilder.SetFeeAmount(buildOpts.gas)

	// Set the signatures.
	signs := make([]signing.SignatureV2, len(buildOpts.signatures))
	for i, sig := range buildOpts.signatures {
		signs[i] = signing.SignatureV2{
			Data: &signing.SingleSignatureData{
				Signature: sig,
				SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			},
		}
	}
	txBuilder.SetSignatures(signs...)

	// Generate a JSON string.
	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func BroadcastTx(tx []byte) (*sdk.TxResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := btx.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&btx.BroadcastTxRequest{
			Mode:    btx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: tx, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return nil, err
	}
	return grpcRes.TxResponse, nil
}

// SimulateTx simulates a transaction on the Sonr blockchain network
func SimulateTx(tx []byte) (*sdk.Result, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := btx.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.Simulate(
		context.Background(),
		&btx.SimulateRequest{

			TxBytes: tx, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return nil, err
	}
	return grpcRes.Result, nil
}
