package local

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc"
)

type BroadcastTxResponse = txtypes.BroadcastTxResponse

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Tendermint Node RPC                              ||
// ! ||--------------------------------------------------------------------------------||

var kBaseAPIUrl = fmt.Sprintf("http://%s", NodeAPIHost())

// The `EncodeTx` function is a method of the `LocalContext` struct. It takes a transaction (`tx`) as input and returns the encoded transaction bytes as output.
func (c LocalContext) EncodeTx(tx *txtypes.Tx) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/cosmos/tx/v1beta1/encode", kBaseAPIUrl)
	req := &txtypes.TxEncodeRequest{
		Tx: tx,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	bz, err := PostJSON(endpoint, body)
	if err != nil {
		return nil, err
	}
	resp := new(txtypes.TxEncodeResponse)
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return resp.TxBytes, nil
}

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func (c LocalContext) BroadcastTx(txMsg *txtypes.Tx) (*BroadcastTxResponse, error) {
	// Encode the transaction to Protobuf binary.
	txRawBytes, err := txMsg.Marshal()
	if err != nil {
		return nil, err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		NodeGrpcHost(),      // Or your gRPC server address.
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
		NodeGrpcHost(),      // Or your gRPC server address.
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

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func (c LocalContext) BroadcastTxAPI(tx *txtypes.Tx) (*BroadcastTxResponse, error) {
	endpoint := fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", kBaseAPIUrl)
	txBz, err := c.EncodeTx(tx)
	if err != nil {
		return nil, err
	}
	// Encode the transaction to Protobuf binary.
	req := &txtypes.BroadcastTxRequest{
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_UNSPECIFIED,
		TxBytes: txBz, // Proto-binary of the signed transaction, see previous step.
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// Create a connection to the gRPC server.
	bz, err := PostJSON(endpoint, body)
	if err != nil {
		return nil, err
	}
	resp := new(txtypes.BroadcastTxResponse)
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SimulateTx simulates a transaction on the Sonr blockchain network
func (c LocalContext) SimulateTxAPI(txRawBytes []byte) (*txtypes.SimulateResponse, error) {
	endpoint := fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", kBaseAPIUrl)
	req := &txtypes.SimulateRequest{
		TxBytes: txRawBytes, // Proto-binary of the signed transaction, see previous step.
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Create a connection to the gRPC server.
	bz, err := PostJSON(endpoint, body)
	if err != nil {
		return nil, err
	}
	resp := new(txtypes.SimulateResponse)
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetJSON makes a GET request to the given URL and returns the response body as bytes
func GetJSON(url string) ([]byte, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// PostJSON makes a POST request to the given URL and returns the response body as bytes
func PostJSON(url string, body []byte) ([]byte, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
