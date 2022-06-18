package crypto

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	at "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// BuildTx builds a transaction from the given inputs.
func BuildTx(w *MPCWallet, msgs ...sdk.Msg) (*txtypes.AuthInfo, *txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, len(msgs))
	for i, m := range msgs {
		msg, err := codectypes.NewAnyWithValue(m)
		if err != nil {
			return nil, nil, err
		}
		msg.TypeUrl = "/sonrio.sonr.registry.MsgCreateWhoIs"
		anyMsgs[i] = msg
	}
	// Get PublicKey
	pubKey, err := w.PublicKeyProto()
	if err != nil {
		return nil, nil, err
	}

	// Build signerInfo parameters
	anyPubKey, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		return nil, nil, err
	}
	modeInfo := &txtypes.ModeInfo_Single_{
		Single: &txtypes.ModeInfo_Single{
			Mode: signing.SignMode_SIGN_MODE_DIRECT,
		},
	}

	addr, err := w.Address()
	if err != nil {
		return nil, nil, err
	}

	// Create SignerInfos
	signerInfo := &txtypes.SignerInfo{
		PublicKey: anyPubKey,
		ModeInfo: &txtypes.ModeInfo{
			Sum: modeInfo,
		},
		Sequence: 0,
	}

	// Create TXRaw and Marshal
	txBody := txtypes.TxBody{
		Messages: anyMsgs,
	}

	// Create AuthInfo
	authInfo := txtypes.AuthInfo{
		SignerInfos: []*txtypes.SignerInfo{signerInfo},
		Fee: &txtypes.Fee{
			Amount:   sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(2))),
			GasLimit: uint64(100000),
			Payer:    addr,
		},
	}
	return &authInfo, &txBody, nil
}

func GetSignDocBytes(account *at.BaseAccount, authInfo *txtypes.AuthInfo, txBody *txtypes.TxBody) ([]byte, error) {
	// Serialize the transaction body.
	txBodyBz, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the auth info.
	authInfoBz, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create SignDoc
	signDoc := &txtypes.SignDoc{
		BodyBytes:     txBodyBz,
		AuthInfoBytes: authInfoBz,
		ChainId:       "sonr",
		AccountNumber: account.GetAccountNumber(),
	}
	return signDoc.Marshal()
}
