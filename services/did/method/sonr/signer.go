package sonr

import (
	// Import necessary packages

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/services/did/types"
	identitytypes "github.com/sonrhq/core/x/identity/types"
)

// SignCosmosTx signs a transaction with the given inputs.
func SignCosmosTx(anySigner types.AnySignerEntity, msgs ...sdk.Msg) ([]byte, error) {
	// Build TxBody
	txBody, err := BuildTxBody(msgs...)
	if err != nil {
		return nil, err
	}
	// Build AuthInfo
	authInfo, err := GetAuthInfo(anySigner, crypto.NewSNRCoins(0))
	if err != nil {
		return nil, err
	}
	// Build SignDoc
	signDoc, err := GetSignDoc(txBody, authInfo)
	if err != nil {
		return nil, err
	}
	// Sign the SignDoc
	sig, err := SignDocWithKSS(anySigner, signDoc)
	if err != nil {
		return nil, err
	}
	// Create the raw transaction
	rawTx, err := SerializeRawTx(txBody, authInfo, sig)
	if err != nil {
		return nil, err
	}
	return rawTx, nil
}

// BuildTxBody builds a transaction from the given inputs.
func BuildTxBody(msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, 0)
	for _, m := range msgs {
		anyMsg, err := codectypes.NewAnyWithValue(m)
		if err != nil {
			return nil, err
		}
		// anyMsg.TypeUrl = typeUrl
		anyMsgs = append(anyMsgs, anyMsg)
	}
	// Create TXRaw and Marshal
	txBody := txtypes.TxBody{
		Messages: anyMsgs,
	}
	return &txBody, nil
}

// GetAuthInfo creates an AuthInfo instance for this account with the specified gas amount.
func GetAuthInfo(anySigner types.AnySignerEntity, gas sdk.Coins) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	pk, err := anySigner.PublicKey()
	if err != nil {
		return nil, err
	}
	anyPubKey, err := codectypes.NewAnyWithValue(pk)
	if err != nil {
		return nil, err
	}

	// Create AuthInfo
	authInfo := txtypes.AuthInfo{
		SignerInfos: []*txtypes.SignerInfo{
			{
				PublicKey: anyPubKey,
				ModeInfo: &txtypes.ModeInfo{
					Sum: &txtypes.ModeInfo_Single_{
						Single: &txtypes.ModeInfo_Single{
							Mode: 1,
						},
					},
				},
				Sequence: 0,
			},
		},
		Fee: &txtypes.Fee{
			Amount:   gas,
			GasLimit: uint64(300000),
		},
	}
	return &authInfo, nil
}

// GetSignDoc builds a SignDoc from the given inputs.
func GetSignDoc(txBody *txtypes.TxBody, authInfo *txtypes.AuthInfo) (*txtypes.SignDoc, error) {
	txBodyBz, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}
	aiBz, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}
	signDoc := &txtypes.SignDoc{
		BodyBytes:     txBodyBz,
		AuthInfoBytes: aiBz,
		// ChainId:       chainID,
		// AccountNumber: accountNumber,
	}
	return signDoc, nil
}

// SignDocWithKSS signs a SignDoc with the given KeyshareSet.
func SignDocWithKSS(anySigner types.AnySignerEntity, signDoc *txtypes.SignDoc) ([]byte, error) {
	bodyBz, err := codec.ProtoMarshalJSON(signDoc, identitytypes.ModuleCdc.InterfaceRegistry())
	if err != nil {
		return nil, err
	}
	sig, err := anySigner.Sign(bodyBz)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// SerializeRawTx is a helper function to create a raw raw transaction and Marshal it to bytes
func SerializeRawTx(body *txtypes.TxBody, authInfo *txtypes.AuthInfo, sig []byte) ([]byte, error) {
	// Create Tx
	tx := &txtypes.Tx{
		Body:       body,
		AuthInfo:   authInfo,
		Signatures: [][]byte{sig},
	}

	// Marshal the txRaw
	return tx.Marshal()
}
