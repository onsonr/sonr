package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	MIN_BUY_ALIAS = sdk.Coins{sdk.NewInt64Coin("snr", 10)}
)

func UnmarshalTxMsg(typeUrl string, data []byte) (sdk.Msg, error) {
	switch typeUrl {
	case "/sonrio.sonr.registry.MsgCreateWhoIs":
		req := &MsgCreateWhoIs{}
		err := req.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		return req, nil
	case "/sonrio.sonr.registry.MsgUpdateWhoIs":
		req := &MsgUpdateWhoIs{}
		err := req.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		return req, nil
	case "/sonrio.sonr.registry.MsgDeactivateWhoIs":
		req := &MsgDeactivateWhoIs{}
		err := req.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		return req, nil
	case "/sonrio.sonr.registry.MsgBuyAlias":
		req := &MsgBuyAlias{}
		err := req.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		return req, nil
	case "/sonrio.sonr.registry.MsgTransferAlias":
		req := &MsgTransferAlias{}
		err := req.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		return req, nil
	case "/sonrio.sonr.registry.MsgSellAlias":
		req := &MsgSellAlias{}
		err := req.Unmarshal(data)
		if err != nil {
			return nil, err
		}
		return req, nil

	}
	return nil, ErrUnknownTxMsgType
}
