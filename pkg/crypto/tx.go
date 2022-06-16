package crypto

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stdtx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/pkg/errors"
	"github.com/sonr-io/sonr/pkg/client"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"google.golang.org/grpc"
)

// Balances returns the balances of the given party.
func (w *MPCWallet) Balances() sdk.Coins {
	addr, err := w.Bech32Address()
	if err != nil {
		return nil
	}

	resp, err := client.CheckBalance(addr)
	if err != nil {
		return nil
	}
	fmt.Println("-- Check Balance --\n", resp)
	return resp
}

func (w *MPCWallet) BroadcastCreateWhoIsRaw() error {
	// addr, err := w.Bech32Address()
	// if err != nil {
	// 	return err
	// }
	// doc, err := w.DIDDocument()
	// if err != nil {
	// 	return err
	// }

	// docJSON, err := doc.MarshalJSON()
	// if err != nil {
	// 	return err
	// }

	// msg := rt.NewMsgCreateWhoIs(addr, docJSON, rt.WhoIsType_USER)
	// msgBytes, err := msg.Marshal()
	// if err != nil {
	// 	return err
	// }

	// // Sign the transaction.
	// tx, err := w.SignTx(msgBytes, "sonrio.sonr.registry/MsgCreateWhoIs")
	// if err != nil {
	// 	return err
	// }

	// resp, err := client.BroadcastTx(tx)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("-- TX Response --\n", resp)
	return nil

}

func (w *MPCWallet) BroadcastCreateWhoIsWithEncoding() error {
	// Create a new TxBuilder.
	txConfig := tx.NewTxConfig(rt.ModuleCdc, tx.DefaultSignModes)
	txBuilder := txConfig.NewTxBuilder()
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin("snr", sdk.NewInt(2))))

	// Create a new WhoIs document.
	addr, err := w.Bech32Address()
	if err != nil {
		return errors.Wrap(err, "failed to get address")
	}
	doc, err := w.DIDDocument()
	if err != nil {
		return errors.Wrap(err, "failed to create DID document")
	}
	docJSON, err := doc.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "failed to marshal DID document")
	}

	// Create a new WhoIs message.
	msg := rt.NewMsgCreateWhoIs(addr, docJSON, rt.WhoIsType_USER)
	err = txBuilder.SetMsgs(msg)
	if err != nil {
		return errors.WithMessage(err, "failed to set message")
	}

	// Get the transaction sign bytes.
	signBz, err := txConfig.SignModeHandler().GetSignBytes(signing.SignMode_SIGN_MODE_DIRECT, xauthsigning.SignerData{
		ChainID:       "sonr",
		AccountNumber: 0,
		Sequence:      0,
	}, txBuilder.GetTx())

	// Sign the transaction.
	sig, err := w.Sign(signBz)
	if err != nil {
		return errors.Wrap(err, "failed to sign transaction")
	}

	// Add the signature data to the transaction.
	txBuilder.SetSignatures(signing.SignatureV2{
		//	PubKey: &pubKey,
		Data: &signing.SingleSignatureData{
			Signature: ECDSASignatureToBytes(sig),
			SignMode:  txConfig.SignModeHandler().DefaultMode(),
		},
		Sequence: 0,
	})

	// Generate a JSON string.
	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return errors.WithMessage(err, "failed to marshal signed transaction")
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // Or your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	)
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.
	txClient := stdtx.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.Simulate(
		context.TODO(),
		&stdtx.SimulateRequest{
			// Mode:    btx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Broadcasted transaction:", grpcRes)
	return nil

}
