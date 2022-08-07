package client

import (
	"context"

	rt "github.com/sonr-io/sonr/x/registry/types"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

type Cosmos struct {
	accName string
	address string
	cosmosclient.Client
	schemaQuery   st.QueryClient
	registryQuery rt.QueryClient
}

// NewCosmos creates a Sonr Blockchain client with the given account and provides helper functions
// func NewCosmos(ctx context.Context, config *config.Config) (*Cosmos, error) {
// 	// Create a new cosmos client
// 	cosmos, err := cosmosclient.New(ctx, toCosmosOptions(config)...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// get account from the keyring by account name and return a bech32 address
// 	account, err := cosmos.Account(config.CosmosAccountName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// create a new client instance
// 	return &Cosmos{
// 		accName:       config.CosmosAccountName,
// 		address:       account.Address("snr"),
// 		Client:        cosmos,
// 		bucketQuery:   bt.NewQueryClient(cosmos.Context),
// 		channelQuery:  ct.NewQueryClient(cosmos.Context),
// 		schemaQuery:   st.NewQueryClient(cosmos.Context),
// 		registryQuery: rt.NewQueryClient(cosmos.Context),
// 	}, nil
// }

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
// Schemas
// -------
func (cc *Cosmos) BroadcastCreateSchema(msg *st.MsgCreateSchema) (*st.MsgCreateSchemaResponse, error) {
	broadCastResp, err := cc.Client.BroadcastTx(cc.accName, msg)
	if err != nil {
		return nil, err
	}

	// Decode the response
	respMsg := st.MsgCreateSchemaResponse{}
	err = broadCastResp.Decode(&respMsg)
	if err != nil {
		return nil, err
	}
	return &respMsg, nil
}

func (cc *Cosmos) QuerySchema(creator string, schemaDID string) (*st.QuerySchemaResponse, error) {
	queryReq := st.QuerySchemaRequest{
		Creator: creator,
		Did:     schemaDID,
	}
	queryResp, err := cc.schemaQuery.Schema(context.Background(), &queryReq)

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}

func (cc *Cosmos) QueryWhatIsAll(queryReq *st.QueryAllWhatIsRequest) (*st.QueryAllWhatIsResponse, error) {
	queryResp, err := cc.schemaQuery.WhatIsAll(context.Background(), queryReq)

	if err != nil {
		return nil, err
	}

	return queryResp, nil
}
