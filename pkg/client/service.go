package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc"
)

const (
	// RPC Address for public node
	SONR_RPC_ADDR_PUBLIC = "143.198.29.209:9090"

	// HTTP Faucet Address
	SONR_HTTP_FAUCET = "http://143.198.29.209:8000"

	// RPC Address for local node
	SONR_RPC_ADDR_LOCAL = "127.0.0.1:9090"
)

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func BroadcastTx(tx []byte) (*sdk.TxResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		SONR_RPC_ADDR_PUBLIC, // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
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
		SONR_RPC_ADDR_PUBLIC, // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := txtypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.Simulate(
		context.Background(),
		&txtypes.SimulateRequest{
			TxBytes: tx, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return nil, err
	}
	return grpcRes.Result, nil
}

type FaucetRequest struct {
	Address string   `json:"address"`
	Coins   []string `json:"coins"`
}

// RequestFaucet requests a faucet from the Sonr network
func RequestFaucet(address string) error {
	values := FaucetRequest{
		Address: address,
		Coins:   []string{"12snr"},
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(SONR_HTTP_FAUCET, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
