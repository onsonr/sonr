package client

import (
	"context"

	"github.com/kataras/golog"
	bt "github.com/sonr-io/blockchain/x/bucket/types"
	ct "github.com/sonr-io/blockchain/x/channel/types"
	ot "github.com/sonr-io/blockchain/x/object/types"
	rt "github.com/sonr-io/blockchain/x/registry/types"
	"github.com/sonr-io/core/highway/config"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

type Cosmos struct {
	accName       string
	address       string
	instance      cosmosclient.Client
	bucketQuery   bt.QueryClient
	channelQuery  ct.QueryClient
	objectQuery   ot.QueryClient
	registryQuery rt.QueryClient
}

// NewCosmos creates a Sonr Blockchain client with the given account and provides helper functions
func NewCosmos(ctx context.Context, config *config.Config) (*Cosmos, error) {
	// Create a new cosmos client
	cosmos, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(config.CosmosAddressPrefix), cosmosclient.WithKeyringBackend(config.CosmosKeyringBackend))
	if err != nil {
		return nil, err
	}

	// get account from the keyring by account name and return a bech32 address
	account, err := cosmos.Account(config.CosmosAccountName)
	if err != nil {
		return nil, err
	}

	// create a new client instance
	return &Cosmos{
		accName:       config.CosmosAccountName,
		address:       account.Address("snr"),
		instance:      cosmos,
		bucketQuery:   bt.NewQueryClient(cosmos.Context),
		channelQuery:  ct.NewQueryClient(cosmos.Context),
		objectQuery:   ot.NewQueryClient(cosmos.Context),
		registryQuery: rt.NewQueryClient(cosmos.Context),
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
func (cc *Cosmos) BroadcastRegisterName(msg *rt.MsgRegisterName) (*rt.MsgRegisterNameResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.instance.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgRegisterNameResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// QueryAllNames returns all names registered on the blockchain
func (cc *Cosmos) QueryAllNames() ([]rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIsAll(context.Background(), &rt.QueryAllWhoIsRequest{})
	if err != nil {
		golog.Errorf("Error querying all names: %s", err)
		return nil, err
	}
	return queryResp.GetWhoIs(), nil
}

// QueryAllNames returns all names registered on the blockchain
func (cc *Cosmos) QueryName(name string) (*rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIs(context.Background(), &rt.QueryGetWhoIsRequest{
		Index: name,
	})
	if err != nil {
		golog.Errorf("Error querying name: %s", err)
		return nil, err
	}
	whois := queryResp.GetWhoIs()
	return &whois, nil
}
