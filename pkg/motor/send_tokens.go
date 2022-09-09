package motor

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/types"
	bt "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

func (m *motorNodeImpl) SendTokens(req mt.PaymentRequest) (*mt.PaymentResponse, error) {
	// Build Message
	amt := types.NewInt(int64(req.Amount))
	fromAddr, err := types.AccAddressFromBech32(req.GetFrom())
	if err != nil {
		return nil, err
	}

	toAddr, err := types.AccAddressFromBech32(req.GetTo())
	if err != nil {
		return nil, err
	}

	// Check user balance
	bal := m.GetBalance()
	if bal < amt.Int64() {
		return nil, errors.New("Failed to issue payment to user: insufficient funds")
	}

	// Build transaction
	amount := types.NewCoins(types.NewCoin("snr", amt))
	msg1 := bt.NewMsgSend(fromAddr, toAddr, amount)
	txRaw, err := tx.SignTxWithWallet(m.Wallet, "/cosmos.bank.v1beta1.MsgSend", msg1)
	if err != nil {
		return nil, err
	}

	// Send transaction
	resp, err := m.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	// Get updated balance and return response
	updatedBal := bal - amt.Int64()
	cwir := &bt.MsgSendResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}
	return &mt.PaymentResponse{
		Code:           int32(resp.TxResponse.Code),
		Message:        resp.TxResponse.Info,
		TxHash:         resp.TxResponse.TxHash,
		UpdatedBalance: int32(updatedBal),
	}, nil
}
