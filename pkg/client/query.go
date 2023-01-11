package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CheckBalance returns the balance of the given address
func (c *ClientStub) CheckBalance(addr string) (types.Coins, error) {
	res, err := c.bankClient.AllBalances(context.Background(), &banktypes.QueryAllBalancesRequest{
		Address: addr, // "snr155huqeqlm4lh5vvgychum0sa2xw70654m3kucq7vufjkx89hzvuqx0jmqc",
	})
	if err != nil {
		return nil, err
	}
	return res.GetBalances(), nil
}

// RequestFaucet requests a faucet from the Sonr network
func (c *ClientStub) RequestFaucet(address string) error {
	values := faucetRequest{
		Address: address,
		Coins:   []string{"12000snr"},
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.faucetUrl, "application/json", bytes.NewBuffer(json_data))
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
