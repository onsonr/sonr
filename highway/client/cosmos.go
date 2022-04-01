package client

import (
	"context"

	"github.com/sonr-io/blockchain/x/registry/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

type Cosmos struct {
	accName  string
	address  string
	instance cosmosclient.Client
	query    types.QueryClient
}

// NewCosmos creates a Sonr Blockchain client with the given account and provides helper functions
func NewCosmos(ctx context.Context, accName string, options ...cosmosclient.Option) (*Cosmos, error) {
	// Create a new cosmos client
	cosmos, err := cosmosclient.New(ctx, options...)
	if err != nil {
		return nil, err
	}

	// get account from the keyring by account name and return a bech32 address
	acc, err := cosmos.Address(accName)
	if err != nil {
		return nil, err
	}

	return &Cosmos{
		accName:  accName,
		address:  acc.String(),
		instance: cosmos,
		query:    types.NewQueryClient(cosmos.Context),
	}, nil
}

// AccountName returns the account name as string
func (cc *Cosmos) AccountName() string {
	return cc.accName
}

// AccountName returns the account name as string
func (cc *Cosmos) Address() string {
	return cc.address
}

// BroadcastTx broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastRegisterName(msg *types.MsgRegisterName) (cosmosclient.Response, error) {
	return cc.instance.BroadcastTx(cc.accName, msg)
}

// QueryAllNames returns all names registered on the blockchain
func (cc *Cosmos) QueryAllNames() ([]types.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.query.WhoIsAll(context.Background(), &types.QueryAllWhoIsRequest{})
	if err != nil {
		return nil, err
	}
	return queryResp.GetWhoIs(), nil
}

// QueryAllNames returns all names registered on the blockchain
func (cc *Cosmos) QueryName(name string) (*types.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.query.WhoIs(context.Background(), &types.QueryGetWhoIsRequest{
		Index: name,
	})
	if err != nil {
		return nil, err
	}
	whois := queryResp.GetWhoIs()
	return &whois, nil
}
