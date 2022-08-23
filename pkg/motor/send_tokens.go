package motor

import (
	"github.com/cosmos/cosmos-sdk/types"
	bt "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
)

func (m *motorNodeImpl) SendTokens(req mt.PaymentRequest) (*mt.PaymentResponse, error) {
	// Build Message
	fromAddr, err := types.AccAddressFromBech32(req.GetFrom())
	if err != nil {
		return nil, err
	}

	toAddr, err := types.AccAddressFromBech32(req.GetTo())
	if err != nil {
		return nil, err
	}

	amount := types.NewCoins(types.NewCoin("snr", types.NewInt(req.GetAmount())))
	msg1 := bt.NewMsgSend(fromAddr, toAddr, amount)
	txRaw, err := tx.SignTxWithWallet(m.Wallet, "/cosmos.bank.v1beta1.MsgSend", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	cwir := &bt.MsgSendResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}

	bal := m.GetBalance()

	return &mt.PaymentResponse{
		Code:           int32(resp.TxResponse.Code),
		Message:        resp.TxResponse.Info,
		TxHash:         resp.TxResponse.TxHash,
		UpdatedBalance: int32(bal),
	}, nil
}
