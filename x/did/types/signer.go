package types

import (
	"context"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type directHandler struct{}

func (s directHandler) Mode() signingv1beta1.SignMode {
	return signingv1beta1.SignMode_SIGN_MODE_DIRECT_AUX
}

func (s directHandler) GetSignBytes(
	_ context.Context,
	signerData signing.SignerData,
	txData signing.TxData,
) ([]byte, error) {
	txDoc := tx.SignDoc{
		BodyBytes:     txData.BodyBytes,
		AuthInfoBytes: txData.AuthInfoBytes,
		ChainId:       signerData.ChainID,
		AccountNumber: signerData.AccountNumber,
	}
	return txDoc.Marshal()
}
