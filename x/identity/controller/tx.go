package controller

import (
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/tx/cosmos"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
)

// CreatePrimaryIdentity sends a transaction to create a new DID document with the provided account
func (c *didController) CreatePrimaryIdentity(doc *types.DidDocument, acc models.Account, alias string) (*local.BroadcastTxResponse, error) {
	msg := types.NewMsgCreateDidDocument(acc.Address(), alias, doc)
	bz, err := cosmos.SignAnyTransactions(acc, msg)
	if err != nil {
		return nil, err
	}
	return local.Context().BroadcastTx(bz)
}

// UpdatePrimaryIdentity sends a transaction to update an existing DID document with the provided account
func (c *didController) UpdatePrimaryIdentity(docs ...*types.DidDocument) (*local.BroadcastTxResponse, error) {
	msg := types.NewMsgUpdateDidDocument(c.primary.Address(), c.primaryDoc, docs...)
	bz, err := cosmos.SignAnyTransactions(c.primary, msg)
	if err != nil {
		return nil, err
	}
	return local.Context().BroadcastTx(bz)
}
