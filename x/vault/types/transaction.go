package types

import (
	// Import necessary packages

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// SignTransaction signs a Cosmos transaction for Token Transfer
func SignTransaction(wa Account, to string, amount sdk.Int, denom string) ([]byte, error) {
	// Build the transaction body
	txBody, err := buildTxBody(&banktypes.MsgSend{
		FromAddress: wa.Address(),
		ToAddress:   to,
		Amount:      sdk.NewCoins(sdk.NewCoin(denom, amount)),
	})
	if err != nil {
		return nil, err
	}

	// Get AuthInfo
	authInfo, err := wa.GetAuthInfo(sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(0))))
	if err != nil {
		return nil, err
	}

	// Sign the transaction body
	_, sig, err := signTxBodyBytes(wa, txBody, authInfo)
	if err != nil {
		return nil, err
	}

	// Create the raw transaction bytes
	rawTxBytes, err := createRawTxBytes(txBody, authInfo, sig)
	if err != nil {
		return nil, err
	}

	return rawTxBytes, nil
}

// SignAnyTransactions signs a Cosmos transaction for a list of arbitrary messages
func SignAnyTransactions(wa Account, msgs ...sdk.Msg) ([]byte, error) {
	// Build the transaction body
	txBody, err := buildTxBody(msgs...)
	if err != nil {
		return nil, err
	}
	// Get AuthInfo
	authInfo, err := wa.GetAuthInfo(sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(0))))
	if err != nil {
		return nil, err
	}

	// Sign the transaction body
	_, sig, err := signTxBodyBytes(wa, txBody, authInfo)
	if err != nil {
		return nil, err
	}

	// Create the raw transaction bytes
	rawTxBytes, err := createRawTxBytes(txBody, authInfo, sig)
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
func buildTxBody(msgs ...sdk.Msg) (*txtypes.TxBody, error) {
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

// createRawTxBytes is a helper function to create a raw raw transaction and Marshal it to bytes
func createRawTxBytes(body *txtypes.TxBody, authInfo *txtypes.AuthInfo, sig []byte) ([]byte, error) {
	// Create Tx
	tx := &txtypes.Tx{
		Body:       body,
		AuthInfo:   authInfo,
		Signatures: [][]byte{sig},
	}

	// Marshal the txRaw
	return tx.Marshal()
}

func signTxBodyBytes(wa Account, txBody *txtypes.TxBody, authInf *txtypes.AuthInfo) ([]byte, []byte, error) {
	// Serialize the transaction body.
	txBodyBz, err := txBody.Marshal()
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
		BodyBytes:     txBodyBz,
		AuthInfoBytes: aiBz,
		ChainId:       "core",
		AccountNumber: 0,
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
