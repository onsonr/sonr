package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
)



func (c *Client) CheckBalance(address string) (types.Coins, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		c.GetRPCAddress(), // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
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

// RequestFaucet requests a faucet from the Sonr network
func (c *Client) RequestFaucet(address string) error {
	values := faucetRequest{
		Address: address,
		Coins:   []string{"12snr"},
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
