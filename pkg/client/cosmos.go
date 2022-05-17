package client

import (
	"context"

	"github.com/sonr-io/sonr/pkg/config"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	ct "github.com/sonr-io/sonr/x/channel/types"
	ot "github.com/sonr-io/sonr/x/object/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
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
	cosmos, err := cosmosclient.New(ctx, toCosmosOptions(config)...)
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
// BroadcastBuyAppAlias broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastBuyAlias(msg *rt.MsgBuyAlias) (*rt.MsgBuyAliasResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {

		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgBuyAliasResponse{}
	err = resp.Decode(respMsg)
	if err != nil {

		return nil, err
	}
	return respMsg, nil
}

// BroadcastBuyNameAlias broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastBuyNameAlias(msg *rt.MsgBuyAlias) (*rt.MsgBuyAliasResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {

		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgBuyAliasResponse{}
	err = resp.Decode(respMsg)
	if err != nil {

		return nil, err
	}
	return respMsg, nil
}

// BroadcastSellAlias broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastSellAlias(msg *rt.MsgSellAlias) (*rt.MsgSellAliasResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {

		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgSellAliasResponse{}
	err = resp.Decode(respMsg)
	if err != nil {

		return nil, err
	}
	return respMsg, nil
}

// BroadcastTransferNameAlias broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastTransferAlias(msg *rt.MsgTransferAlias) (*rt.MsgTransferAliasResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {

		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgTransferAliasResponse{}
	err = resp.Decode(respMsg)
	if err != nil {

		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateApplication broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateWhoIs(msg *rt.MsgCreateWhoIs) (*rt.MsgCreateWhoIsResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgCreateWhoIsResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateApplication broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateWhoIs(msg *rt.MsgUpdateWhoIs) (*rt.MsgUpdateWhoIsResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgUpdateWhoIsResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateName broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateWhoIs(msg *rt.MsgDeactivateWhoIs) (*rt.MsgDeactivateWhoIsResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &rt.MsgDeactivateWhoIsResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// NameExists checks if a name exists on the blockchain
func (cc *Cosmos) NameExists(name string) bool {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIs(context.Background(), &rt.QueryWhoIsRequest{Did: name})
	if err != nil {
		return false
	}

	// check if the name exists
	return queryResp.GetWhoIs().Alias[0].GetName() == name
}

// QueryAllWhoIs returns all DIDDocuments registered on the blockchain
func (cc *Cosmos) QueryAllWhoIs() ([]rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIsAll(context.Background(), &rt.QueryAllWhoIsRequest{})
	if err != nil {
		return nil, err
	}
	return queryResp.GetWhoIs(), nil
}

// QueryWhoIs returns a DIDDocument for the given name registered on the blockchain
func (cc *Cosmos) QueryWhoIs(did string) (*rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIs(context.Background(), &rt.QueryWhoIsRequest{
		Did: did,
	})
	if err != nil {
		return nil, err
	}
	whois := queryResp.GetWhoIs()
	return whois, nil
}

// QueryWhoIsAlias returns a DIDDocument for the given alias registered on the blockchain
func (cc *Cosmos) QueryWhoIsAlias(alias string) (*rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIsAlias(context.Background(), &rt.QueryWhoIsAliasRequest{
		Alias: alias,
	})
	if err != nil {
		return nil, err
	}
	whois := queryResp.GetWhoIs()
	return whois, nil
}

// QueryWhoIsController returns the controller for the given name registered on the blockchain
func (cc *Cosmos) QueryWhoIsController(controller string) (*rt.WhoIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.registryQuery.WhoIsController(context.Background(), &rt.QueryWhoIsControllerRequest{
		Controller: controller,
	})
	if err != nil {
		return nil, err
	}
	whois := queryResp.GetWhoIs()
	return whois, nil
}

// -------
// Buckets
// -------
// BroadcastCreateBucket broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateBucket(msg *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &bt.MsgCreateBucketResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateBucket broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateBucket(msg *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &bt.MsgUpdateBucketResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastDeactivateBucket broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateBucket(msg *bt.MsgDeactivateBucket) (*bt.MsgDeactivateBucketResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &bt.MsgDeactivateBucketResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// QueryAllBuckets returns all names registered on the blockchain
func (cc *Cosmos) QueryAllBuckets() ([]bt.WhichIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.bucketQuery.WhichIsAll(context.Background(), &bt.QueryAllWhichIsRequest{})
	if err != nil {
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
		return nil, err
	}
	whichIs := queryResp.GetWhichIs()
	return &whichIs, nil
}

// -------
// Channels
// -------
// BroadcastDeactivateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateChannel(msg *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &ct.MsgCreateChannelResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateChannel(msg *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &ct.MsgUpdateChannelResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateChannel(msg *ct.MsgDeactivateChannel) (*ct.MsgDeactivateChannelResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &ct.MsgDeactivateChannelResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// QueryAllChannels returns all names registered on the blockchain
func (cc *Cosmos) QueryAllChannels() ([]ct.HowIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.channelQuery.HowIsAll(context.Background(), &ct.QueryAllHowIsRequest{})
	if err != nil {
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
		return nil, err
	}
	howIs := queryResp.GetHowIs()
	return &howIs, nil
}

// -------
// Objects
// -------
// BroadcastCreateObject broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastCreateObject(msg *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &ot.MsgCreateObjectResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateObject broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastUpdateObject(msg *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &ot.MsgUpdateObjectResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// BroadcastUpdateChannel broadcasts a transaction to the blockchain
func (cc *Cosmos) BroadcastDeactivateObject(msg *ot.MsgDeactivateObject) (*ot.MsgDeactivateObjectResponse, error) {
	// broadcast the transaction to the blockchain
	resp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := &ot.MsgDeactivateObjectResponse{}
	err = resp.Decode(respMsg)
	if err != nil {
		return nil, err
	}
	return respMsg, nil
}

// QueryAllObjects returns all names registered on the blockchain
func (cc *Cosmos) QueryAllObjects() ([]ot.WhatIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.objectQuery.WhatIsAll(context.Background(), &ot.QueryAllWhatIsRequest{})
	if err != nil {
		return nil, err
	}
	return queryResp.GetWhatIs(), nil
}

// QueryObject returns all names registered on the blockchain
func (cc *Cosmos) QueryObject(did string) (*ot.WhatIs, error) {
	// query the blockchain using the client's `WhoIsAll` method to get all names
	queryResp, err := cc.objectQuery.WhatIs(context.Background(), &ot.QueryWhatIsRequest{
		Did: did,
	})
	if err != nil {
		return nil, err
	}
	whatIs := queryResp.GetWhatIs()
	return &whatIs, nil
}
