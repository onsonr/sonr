package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	//cdc "github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	acctypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sonr-io/sonr/app"
	"google.golang.org/grpc"
)

func (c *Client) CheckBalance(address string) (types.Coins, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	resp, err := banktypes.NewQueryClient(grpcConn).AllBalances(context.Background(), &banktypes.QueryAllBalancesRequest{
		Address: address, // "snr155huqeqlm4lh5vvgychum0sa2xw70654m3kucq7vufjkx89hzvuqx0jmqc",
	})
	if err != nil {
		return nil, err
	}
	return resp.GetBalances(), nil
}

//GetAccount fetches on-chain account details from the Sonr Network
func (c *Client) GetAccount(address string) (*acctypes.BaseAccount, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(),   // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()
	if err != nil {
		return nil, err
	}
	encodingConfig := app.MakeEncodingConfig()

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino)

	queryCli := acctypes.NewQueryClient(grpcConn)
	resp, err := queryCli.Account(context.Background(), &acctypes.QueryAccountRequest{Address: address})
	if err != nil {
		return nil, err
	}
	accBz, err := initClientCtx.Codec.MarshalJSON(resp.Account)
	if err != nil {
		return nil, err
	}
	var acc acctypes.BaseAccount
	initClientCtx.Codec.UnmarshalJSON(accBz, &acc)
	return &acc, nil
}

// RequestFaucet requests a faucet from the Sonr network
func (c *Client) RequestFaucet(address string) error {
	values := faucetRequest{
		Address: address,
		Coins:   []string{"120snr"},
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.GetFaucetAddress(), "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

type faucetRequest struct {
	Address string   `json:"address"`
	Coins   []string `json:"coins"`
}
