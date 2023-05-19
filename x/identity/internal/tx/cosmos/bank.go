package cosmos

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BuildMsgSend builds a message to send coins from one account to another.
func BuildMsgSend(from, to string, amount int) *banktypes.MsgSend {
	return &banktypes.MsgSend{FromAddress: from, ToAddress: to, Amount: types.NewCoins(types.NewCoin("snr", math.NewInt(int64(amount))))}
}

// BuildMsgMultiSend builds a message to send coins from one account to another.
