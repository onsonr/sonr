package local

import (
	"context"
	"encoding/hex"
	"fmt"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

type BroadcastTxResponse = txtypes.BroadcastTxResponse

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Tendermint Node RPC                              ||
// ! ||--------------------------------------------------------------------------------||

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func (c LocalContext) BroadcastTx(txRawBytes []byte) (*BroadcastTxResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GrpcEndpoint(),    // Or your gRPC server address.
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
func (c LocalContext) SimulateTx(txRawBytes []byte) (*txtypes.SimulateResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.RpcEndpoint(),     // Or your gRPC server address.
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

// ! ||--------------------------------------------------------------------------------||
// ! ||                            Utility Helper Functions                            ||
// ! ||--------------------------------------------------------------------------------||

// DecodeTxResponseData decodes the data from a transaction response
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

// // RequestFaucet funds an account with the given address
// func (c LocalContext) RequestFaucet(ctx context.Context, address string) error {
// 	type FaucetRequest struct {
// 		Address string   `json:"address"`
// 		Coins   []string `json:"coins"`
// 	}
// 	type FaucetResponse struct {
// 		Error string `json:"error"`
// 	}

// 	faucetRequest := &FaucetRequest{
// 		Address: address,
// 		Coins:   []string{"200snr"},
// 	}

// 	// Marshal the request data into JSON format
// 	reqData, err := json.Marshal(faucetRequest)
// 	if err != nil {
// 		return errors.New("failed to marshal faucet request: " + err.Error())
// 	}

// 	// Create an HTTP request with JSON data
// 	req, err := http.NewRequestWithContext(ctx, "POST", c.FaucetEndpoint(), bytes.NewBuffer(reqData))
// 	if err != nil {
// 		return errors.New("failed to create a new HTTP request: " + err.Error())
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	// Send the HTTP request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return errors.New("failed to send faucet request: " + err.Error())
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return errors.New("faucet request returned a non-200 status code")
// 	}

// 	return nil
// }
