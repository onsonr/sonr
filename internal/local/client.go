package local

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"google.golang.org/grpc"
)

type BroadcastTxResponse = txtypes.BroadcastTxResponse

// ! ||--------------------------------------------------------------------------------||
// ! ||                              x/identity RPC client                             ||
// ! ||--------------------------------------------------------------------------------||

// CheckAlias checks if the alias is available and returns the existing DID if it's not
func (c LocalContext) CheckAlias(ctx context.Context, alias string) (bool, *identitytypes.Identity, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return false, nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).AliasAvailable(ctx, &identitytypes.QueryAliasAvailableRequest{Alias: alias})
	if err != nil {
		return false, nil, err
	}
	return resp.Available, resp.ExistingDocument, nil
}

// GetDID returns the DID document with the given id
func (c LocalContext) GetDID(ctx context.Context, id string) (*identitytypes.Identity, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).Did(ctx, &identitytypes.QueryGetDidRequest{Did: id})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetDIDByAlias returns the DID document with the given alias
func (c LocalContext) GetDIDByAlias(ctx context.Context, alias string) (*identitytypes.Identity, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).DidByAlsoKnownAs(ctx, &identitytypes.QueryDidByAlsoKnownAsRequest{AkaId: alias})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetDIDByAlias returns the DID document with the given alias
func (c LocalContext) GetDIDByOwner(ctx context.Context, owner string) (*identitytypes.Identity, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).DidByOwner(ctx, &identitytypes.QueryDidByOwnerRequest{Owner: owner})
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetAllDIDs returns all DID documents
func (c LocalContext) GetAllDIDs(ctx context.Context) ([]*identitytypes.Identity, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {

		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).DidAll(ctx, &identitytypes.QueryAllDidRequest{})
	if err != nil {

		return nil, err
	}
	list := make([]*identitytypes.Identity, len(resp.DidDocument))
	for i, d := range resp.DidDocument {
		list[i] = &d
	}
	return list, nil
}

// OldestUnclaimedWallet returns the oldest unclaimed wallet
func (c LocalContext) GetUnclaimedWallet(ctx context.Context, id uint64) (*identitytypes.ClaimableWallet, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := identitytypes.NewQueryClient(conn).ClaimableWallet(ctx, &identitytypes.QueryGetClaimableWalletRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &resp.ClaimableWallet, nil
}

// GetService returns the service with the given id
func (c LocalContext) GetService(ctx context.Context, origin string) (*servicetypes.ServiceRecord, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := servicetypes.NewQueryClient(conn).ServiceRecord(ctx, &servicetypes.QueryServiceRecordRequest{Origin: origin})
	if err != nil {

		return nil, err
	}
	return &resp.ServiceRecord, nil
}

// GetAllServices returns all services
func (c LocalContext) GetAllServices(ctx context.Context) ([]*servicetypes.ServiceRecord, error) {
	conn, err := grpc.Dial(c.GrpcEndpoint(), grpc.WithInsecure())
	if err != nil {

		return nil, errors.New("failed to connect to grpc server: " + err.Error())
	}
	resp, err := servicetypes.NewQueryClient(conn).ListServiceRecords(ctx, &servicetypes.ListServiceRecordsRequest{})
	if err != nil {

		return nil, err
	}
	list := make([]*servicetypes.ServiceRecord, len(resp.ServiceRecord))
	for i, s := range resp.ServiceRecord {
		list[i] = &s
	}
	return list, nil
}

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
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
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
