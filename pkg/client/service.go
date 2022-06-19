package client

import (
	"context"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonr-io/sonr/pkg/crypto"

	// "github.com/sonr-io/sonr/pkg/crypto"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"google.golang.org/grpc"
)

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func (c *Client) BroadcastTx(txBody *txtypes.TxBody, sig *ecdsa.Signature, authInfo *txtypes.AuthInfo) (*txtypes.BroadcastTxResponse, error) {
	// Create TXRaw and Marshal
	txRawBytes, err := createRawTxBytes(txBody, sig, authInfo)
	if err != nil {
		return nil, err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
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
	return grpcRes, nil
}

// SimulateTx simulates a transaction on the Sonr blockchain network
func (c *Client) SimulateTx(txBody *txtypes.TxBody, sig *ecdsa.Signature, authInfo *txtypes.AuthInfo) (*txtypes.SimulateResponse, error) {
	// Create TXRaw and Marshal
	txRawBytes, err := createRawTxBytes(txBody, sig, authInfo)
	if err != nil {
		return nil, err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
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

// createRawTxBytes is a helper function to create a raw raw transaction and Marshal it to bytes
func createRawTxBytes(txBody *txtypes.TxBody, sig *ecdsa.Signature, authInfo *txtypes.AuthInfo) ([]byte, error) {
	// Serialize the tx body
	txBytes, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the authInfo
	authInfoBytes, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create a signature list and append the signature
	sigList := make([][]byte, 1)
	sigbz, err := crypto.SerializeSignature(sig)
	if err != nil {
		return nil, err
	}
	sigList[0] = sigbz

	// Create Raw TX
	txRaw := &txtypes.TxRaw{
		BodyBytes:     txBytes,
		AuthInfoBytes: authInfoBytes,
		Signatures:    sigList,
	}

	// Marshal the txRaw
	return txRaw.Marshal()
}
