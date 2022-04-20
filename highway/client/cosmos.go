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
	accName string
	address string
	cosmosclient.Client
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
		Client:        cosmos,
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

// Address returns the account name as string
func (cc *Cosmos) Address() string {
	return cc.address
}

// -------
// Registry
// -------
// BroadcastRegisterApplication broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastRegisterApplication(msg *rt.MsgRegisterApplication) (*rt.MsgRegisterApplicationResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgRegisterApplicationResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastRegisterName broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastRegisterName(msg *rt.MsgRegisterName) (*rt.MsgRegisterNameResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
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

// QueryAllNames returns all DIDDocuments registered on the blockchain
func (cc *Cosmos) QueryAllNames() ([]rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIsAll(context.Background(), &rt.QueryAllWhoIsRequest{})
	if err != nil {
		golog.Errorf("Error querying all names: %s", err)
		return nil, err
	}
	return queryResp.GetWhoIs(), nil
}

// QueryName returns a DIDDocument for the given name registered on the blockchain
func (cc *Cosmos) QueryName(name string) (*rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIs(context.Background(), &rt.QueryWhoIsRequest{
		Did: name,
	})
	if err != nil {
		golog.Errorf("Error querying name: %s", err)
		return nil, err
	}
	whois := queryResp.GetWhoIs()
	return &whois, nil
}

// -------
// Buckets
// -------
// BroadcastCreateBucket broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateBucket(msg *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &bt.MsgCreateBucketResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateBucket broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateBucket(msg *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &bt.MsgUpdateBucketResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastDeactivateBucket broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateBucket(msg *bt.MsgDeactivateBucket) (*bt.MsgDeactivateBucketResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &bt.MsgDeactivateBucketResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// QueryAllBuckets returns all names registered on the blockchain
func (cc *Cosmos) QueryAllBuckets() ([]bt.WhichIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.bucketQuery.WhichIsAll(context.Background(), &bt.QueryAllWhichIsRequest{})
	if err != nil {
		golog.Errorf("Error querying all names: %s", err)
		return nil, err
	}
	return queryResp.GetWhichIs(), nil
}

// QueryBucket returns all names registered on the blockchain
func (cc *Cosmos) QueryBucket(name string) (*bt.WhichIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.bucketQuery.WhichIs(context.Background(), &bt.QueryWhichIsRequest{
		Did: name,
	})
	if err != nil {
		golog.Errorf("Error querying name: %s", err)
		return nil, err
	}
	whois := queryResp.GetWhichIs()
	return &whois, nil
}

// -------
// Channels
// -------
// BroadcastDeactivateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateChannel(msg *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &ct.MsgCreateChannelResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateChannel(msg *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &ct.MsgUpdateChannelResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateChannel(msg *ct.MsgDeactivateChannel) (*ct.MsgDeactivateChannelResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &ct.MsgDeactivateChannelResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// QueryAllChannels returns all names registered on the blockchain
func (cc *Cosmos) QueryAllChannels() ([]ct.HowIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.channelQuery.HowIsAll(context.Background(), &ct.QueryAllHowIsRequest{})
	if err != nil {
		golog.Errorf("Error querying all names: %s", err)
		return nil, err
	}
	return queryResp.GetHowIs(), nil
}

// QueryChannel returns all names registered on the blockchain
func (cc *Cosmos) QueryChannel(name string) (*ct.HowIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.channelQuery.HowIs(context.Background(), &ct.QueryHowIsRequest{
		Did: name,
	})
	if err != nil {
		golog.Errorf("Error querying name: %s", err)
		return nil, err
	}
	whois := queryResp.GetHowIs()
	return &whois, nil
}

// -------
// Objects
// -------
// BroadcastCreateObject broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateObject(msg *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &ot.MsgCreateObjectResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateObject broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateObject(msg *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &ot.MsgUpdateObjectResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateObject(msg *ot.MsgDeactivateObject) (*ot.MsgDeactivateObjectResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		golog.Errorf("Error broadcasting transaction: %s", err)
		return nil, err
	}

	// Decode the response
	respMsg := &ot.MsgDeactivateObjectResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		golog.Errorf("Error decoding response: %v", err)
		return nil, err
	}
	return respMsg, nil
}

// QueryAllObjects returns all names registered on the blockchain
func (cc *Cosmos) QueryAllObjects() ([]ot.WhatIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.objectQuery.WhatIsAll(context.Background(), &ot.QueryAllWhatIsRequest{})
	if err != nil {
		golog.Errorf("Error querying all names: %s", err)
		return nil, err
	}
	return queryResp.GetWhatIs(), nil
}

// QueryObject returns all names registered on the blockchain
func (cc *Cosmos) QueryObject(name string) (*ot.WhatIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.objectQuery.WhatIs(context.Background(), &ot.QueryWhatIsRequest{
		Did: name,
	})
	if err != nil {
		golog.Errorf("Error querying name: %s", err)
		return nil, err
	}
	whois := queryResp.GetWhatIs()
	return &whois, nil
}
