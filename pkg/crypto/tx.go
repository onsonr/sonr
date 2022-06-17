package crypto

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// BuildTx builds a transaction from the given inputs.
func BuildTx(w *MPCWallet, msgs ...sdk.Msg) (*types.AuthInfo, *types.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, len(msgs))
	for i, m := range msgs {
		msg, err := codectypes.NewAnyWithValue(m)
		if err != nil {
			return nil, nil, err
		}
		anyMsgs[i] = msg
	}
	// pubKey, err := w.PublicKeyProto()
	// if err != nil {
	// 	return nil, nil, err
	// }

	// Build signerInfo parameters
	//anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
	modeInfo := &types.ModeInfo_Single_{
		Single: &types.ModeInfo_Single{
			Mode: signing.SignMode_SIGN_MODE_DIRECT,
		},
	}

	// Create SignerInfos
	signerInfo := &types.SignerInfo{
		// PublicKey: anyPubKey,
		ModeInfo: &types.ModeInfo{
			Sum: modeInfo,
		},
	}

	// Create TXRaw and Marshal
	txBody := types.TxBody{
		Messages: anyMsgs,
	}

	// Create AuthInfo
	authInfo := types.AuthInfo{
		SignerInfos: []*types.SignerInfo{signerInfo},
		Fee: &types.Fee{
			Amount: sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(10))),
		},
	}
	return &authInfo, &txBody, nil
}
