package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) BuyAlias(msg rt.MsgBuyAlias) (rt.MsgBuyAliasResponse, error) {
	if msg.Creator == "" {
		return rt.MsgBuyAliasResponse{}, fmt.Errorf("creator cannot be empty")
	}
	if msg.Name == "" {
		return rt.MsgBuyAliasResponse{}, fmt.Errorf("name cannot be empty")
	}

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.registry.MsgBuyAlias", &msg)
	if err != nil {
		return rt.MsgBuyAliasResponse{}, fmt.Errorf("sign tx with wallet: %s", err)
	}

	txresp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return rt.MsgBuyAliasResponse{}, fmt.Errorf("broadcast tx: %s", err)
	}

	resp := &rt.MsgBuyAliasResponse{}
	if err := client.DecodeTxResponseData(txresp.TxResponse.Data, resp); err != nil {
		return rt.MsgBuyAliasResponse{}, fmt.Errorf("decode MsgBuyAliasResponse: %s", err)
	}

	return *resp, nil
}

func (mtr *motorNodeImpl) SellAlias(msg rt.MsgSellAlias) (rt.MsgSellAliasResponse, error) {
	if msg.Creator == "" {
		return rt.MsgSellAliasResponse{}, fmt.Errorf("creator cannot be empty")
	}
	if msg.Alias == "" {
		return rt.MsgSellAliasResponse{}, fmt.Errorf("alias cannot be empty")
	}

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.registry.MsgSellAlias", &msg)
	if err != nil {
		return rt.MsgSellAliasResponse{}, fmt.Errorf("sign tx with wallet: %s", err)
	}

	txresp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return rt.MsgSellAliasResponse{}, fmt.Errorf("broadcast tx: %s", err)
	}

	resp := &rt.MsgSellAliasResponse{}
	if err := client.DecodeTxResponseData(txresp.TxResponse.Data, resp); err != nil {
		return rt.MsgSellAliasResponse{}, fmt.Errorf("decode MsgSellAliasResponse: %s", err)
	}

	return *resp, nil
}

func (mtr *motorNodeImpl) TransferAlias(msg rt.MsgTransferAlias) (rt.MsgTransferAliasResponse, error) {
	if msg.Creator == "" {
		return rt.MsgTransferAliasResponse{}, fmt.Errorf("creator cannot be empty")
	}
	if msg.Alias == "" {
		return rt.MsgTransferAliasResponse{}, fmt.Errorf("alias cannot be empty")
	}
	if msg.Recipient == "" {
		return rt.MsgTransferAliasResponse{}, fmt.Errorf("recipient cannot be empty")
	}

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.registry.MsgTransferAlias", &msg)
	if err != nil {
		return rt.MsgTransferAliasResponse{}, fmt.Errorf("sign tx with wallet: %s", err)
	}

	txresp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return rt.MsgTransferAliasResponse{}, fmt.Errorf("broadcast tx: %s", err)
	}

	resp := &rt.MsgTransferAliasResponse{}
	if err := client.DecodeTxResponseData(txresp.TxResponse.Data, resp); err != nil {
		return rt.MsgTransferAliasResponse{}, fmt.Errorf("decode MsgTransferAliasResponse: %s", err)
	}

	return *resp, nil
}
