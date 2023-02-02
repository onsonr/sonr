package internal

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/sonrhq/core/pkg/client/rosetta"
	"github.com/sonrhq/core/pkg/common"
	"github.com/sonrhq/core/pkg/crypto/wallet"
)

// cosmosAccountImpl is an implementation of the Account interface for Cosmos
type cosmosAccountImpl struct {
	wallet.Account
	rosetta.Client
	rootAcc wallet.Account
}

// LoadCosmosAccount loads a Cosmos account from a wallet account
func LoadCosmosAccount(root wallet.Account, cosmos wallet.Account, client rosetta.Client) wallet.CosmosAccount {
	return &cosmosAccountImpl{
		Account: cosmos,
		rootAcc: root,
		Client:  client,
	}
}

// SignTx signs a transaction
func (a *cosmosAccountImpl) SendTx(note string, msgs ...sdk.Msg) (*tx.BroadcastTxResponse, error) {
	txBody, err := makeTxBody(note, msgs...)
	if err != nil {
		return nil, err
	}
	// Serialize the tx body
	txBytes, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}
	_, sig, err := signTxDocDirectAux(a, txBytes)
	if err != nil {
		return nil, err
	}

	// Create a signature list and append the signature
	sigList := make([][]byte, 1)
	sigList[0] = sig

	return &tx.BroadcastTxResponse{
		//TxHash: signDocDir.TxBodyBytes,
	}, nil
}

// GetSignerData returns the signer data for the account
func (a *cosmosAccountImpl) GetSignerData() authsigning.SignerData {
	return authsigning.SignerData{
		ChainID: common.CURRENT_CHAIN_ID,
		PubKey:  a.PubKey(),
	}
}

// Verify verifies a signature for a message
func (a *cosmosAccountImpl) VerifySignature(msg []byte, sig []byte) bool {
	ok, err := a.rootAcc.Verify(msg, sig)
	if err != nil {
		return false
	}
	return ok
}

//
// Cosmos TX Signing helper functions
//

// makeTxBody builds a transaction from the given inputs.
func makeTxBody(note string, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// func BuildTx(w *crypto.MPCWallet, msgs ...sdk.Msg) (*txtypes.TxBody, error) {
	// Create Any for each message
	anyMsgs := make([]*codectypes.Any, len(msgs))
	for i, m := range msgs {
		msg, err := codectypes.NewAnyWithValue(m)
		if err != nil {
			return nil, err
		}
		anyMsgs[i] = msg
	}

	// Create TXRaw and Marshal
	txBody := txtypes.TxBody{
		Messages: anyMsgs,
		Memo:     note,
	}
	return &txBody, nil
}

// createRawTxBytes is a helper function to create a raw raw transaction and Marshal it to bytes
func createRawTxBytes(txBody *txtypes.TxBody, sig []byte, authInfo *txtypes.AuthInfo) ([]byte, error) {
	// Serialize the tx body
	txBytes, err := txBody.Marshal()
	if err != nil {
		return nil, err
	}

	// Serialize the authInfo
	authInfoBytes, err := authInfo.Marshal()
	if err != nil {
		return nil, err
	}

	// Create a signature list and append the signature
	sigList := make([][]byte, 1)
	sigList[0] = sig

	// Create Raw TX
	txRaw := &txtypes.TxRaw{
		BodyBytes:     txBytes,
		AuthInfoBytes: authInfoBytes,
		Signatures:    sigList,
	}

	// Marshal the txRaw
	return txRaw.Marshal()
}

// getAuthInfoSingle returns the authentication information for the given message.
func getAuthInfoSingle(w wallet.Account, gas int) (*txtypes.AuthInfo, error) {
	// Build signerInfo parameters
	anyPubKey, err := codectypes.NewAnyWithValue(w.PubKey())
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
							Mode: 3,
						},
					},
				},
			},
		},
		Fee: &txtypes.Fee{
			Amount:   sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(int64(gas)))),
			GasLimit: uint64(300000),
		},
	}
	return &authInfo, nil
}

// It takes a transaction body and auth info, serializes them, and then creates a SignDoc object that
// contains the serialized transaction body and auth info, and the chain ID
func getSignDocBytes(authInfo *txtypes.AuthInfo, txBody *txtypes.TxBody) ([]byte, error) {
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
	}
	return signDoc.Marshal()
}

// Signing a transaction.
func signTxDocDirectAux(w wallet.Account, txBody []byte) (*txtypes.SignDocDirectAux, []byte, error) {
	// Build the public key.
	pk, err := codectypes.NewAnyWithValue(w.PubKey())
	if err != nil {
		return nil, nil, err
	}

	// Build the sign doc.
	doc := &txtypes.SignDocDirectAux{
		ChainId:   "sonr",
		PublicKey: pk,
		BodyBytes: txBody,
	}

	// Marshal the document.
	bz, err := doc.Marshal()
	if err != nil {
		return nil, nil, err
	}

	// Sign the document.
	sig, err := w.Sign(bz)
	if err != nil {
		return nil, nil, err
	}

	// Return the document and the signature.
	return doc, sig, nil
}
