package cosmos

import (
	// Import necessary packages

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/sonrhq/core/internal/protocol/packages/controller"
)

// SignTransaction signs a Cosmos transaction for Token Transfer
func SignTransaction(wa controller.Account, to string, amount sdk.Int, denom string) ([]byte, error) {
	// Build the transaction body
	txBody, err := buildTxBody("/cosmos.bank.v1beta1.MsgSend", &banktypes.MsgSend{
		FromAddress: wa.Address(),
		ToAddress:   to,
		Amount:      sdk.NewCoins(sdk.NewCoin(denom, amount)),
	})
	if err != nil {
		return nil, err
	}

	// Sign the transaction body
	bodyBz, sig, err := signTxBodyBytes(wa, txBody)
	if err != nil {
		return nil, err
	}

	// Create the raw transaction bytes
	rawTxBytes, err := createRawTxBytes(bodyBz, sig, wa)
	if err != nil {
		return nil, err
	}

	return rawTxBytes, nil
}

// SignAnyTransactions signs a Cosmos transaction for a list of arbitrary messages
func SignAnyTransactions(wa controller.Account, typeUrl string, msgs ...sdk.Msg) ([]byte, error) {
	// Build the transaction body
	txBody, err := buildTxBody(typeUrl, msgs...)
	if err != nil {
		return nil, err
	}

	// Sign the transaction body
	bodyBz, sig, err := signTxBodyBytes(wa, txBody)
	if err != nil {
		return nil, err
	}

	// Create the raw transaction bytes
	rawTxBytes, err := createRawTxBytes(bodyBz, sig, wa)
	if err != nil {
		return nil, err
	}

	return rawTxBytes, nil
}

//
// ! ||--------------------------------------------------------------------------------||
// ! ||                              /// Helper functions                              ||
// ! ||--------------------------------------------------------------------------------||
//

// buildTxBody builds a transaction from the given inputs.
func buildTxBody(typeUrl string, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, 0)
	for _, m := range msgs {
		anyMsgs = append(anyMsgs, codectypes.UnsafePackAny(m))
	}
	// Create TXRaw and Marshal
	txBody := txtypes.TxBody{
		Messages: anyMsgs,
	}
	return &txBody, nil
}

// createRawTxBytes is a helper function to create a raw raw transaction and Marshal it to bytes
func createRawTxBytes(body []byte, sig []byte, wa controller.Account) ([]byte, error) {
	// Get AuthInfo
	authInfo, err := wa.GetAuthInfo(sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(2))))
	if err != nil {
		return nil, err
	}

	// Serialize the authInfo
	authInfoBytes, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create Raw TX
	txRaw := &txtypes.SignDoc{
		BodyBytes:     body,
		AuthInfoBytes: authInfoBytes,
	}

	// Marshal the txRaw
	return txRaw.Marshal()
}

func signTxBodyBytes(wa controller.Account, txBody *txtypes.TxBody) ([]byte, []byte, error) {
	// Serialize the transaction body.
	txBodyBz, err := txBody.Marshal()
	if err != nil {
		return nil, nil, err
	}
	authInf, err := wa.GetAuthInfo(sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(0))))
	if err != nil {
		return nil, nil, err
	}

	// Build signerInfo parameters
	aiBz, err := authInf.Marshal()
	if err != nil {
		return nil, nil, err
	}
	// Create SignDoc
	signDoc := &txtypes.SignDoc{
		BodyBytes: txBodyBz,
		AuthInfoBytes: aiBz,
	}
	bodyBz, err := codec.ProtoMarshalJSON(signDoc, nil)
	if err != nil {
		return nil, nil, err
	}

	// Sign the transaction body.
	sig, err := wa.Sign(bodyBz)
	if err != nil {
		return nil, nil, err
	}
	return bodyBz, sig, nil
}
