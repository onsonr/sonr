package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
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

// CreateAccount creates a new account on the Sonr network
func (c *Client) CreateAccount(did string, options ...QueryWhoIsOption) error {
	// Create a new request.
	req := &rt.QueryWhoIsRequest{Did: did}
	for _, option := range options {
		option(req)
	}

	endpoint := fmt.Sprintf("/sonr-io/sonr/registry/who_is/%s&bech32=%s&pubkey.key=%s", req.Did, req.Bech32, req.Pubkey.Key)
	resp, err := http.Get(c.GetAPIAddress() + endpoint)
	if err != nil {
		return err
	}
	if resp == nil {
		return nil
	}
	return nil
}

// RequestFaucet requests a faucet from the Sonr network
func (c *Client) RequestFaucet(address string) error {
	values := faucetRequest{
		Address: address,
		Coins:   []string{"200snr"},
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
