package simulation

import (
	"crypto/ed25519"
	cryptrand "crypto/rand"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/x/registry/keeper"
	"github.com/sonr-io/sonr/x/registry/types"
)

func SimulateMsgCreateWhoIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// Creates a mock did document for the provided simulated account
		doc, err := CreateMockDidDocument(simAccount)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "createWhoIs", "failed to create mock did document"), nil, err
		}

		// Marshal Json document
		docBytes, err := doc.MarshalJSON()
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "createWhoIs", "failed to marshal json document"), nil, err
		}

		msg := &types.MsgCreateWhoIs{
			Creator:     simAccount.Address.String(),
			DidDocument: docBytes,
			WhoisType:   types.WhoIsType_USER,
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateWhoIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			whoIs      = types.WhoIs{}
			msg        = &types.MsgUpdateWhoIs{}
			allWhoIs   = k.GetAllWhoIs(ctx)
			found      = false
		)
		for _, obj := range allWhoIs {
			simAccount, found = FindAccount(accs, obj.Owner)
			if found {
				whoIs = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "whoIs owner not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.DidDocument = whoIs.DidDocument
		msg.Did = whoIs.Owner

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteWhoIs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			whoIs      = types.WhoIs{}
			msg        = &types.MsgUpdateWhoIs{}
			allWhoIs   = k.GetAllWhoIs(ctx)
			found      = false
		)
		for _, obj := range allWhoIs {
			simAccount, found = FindAccount(accs, obj.Owner)
			if found {
				whoIs = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "whoIs owner not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Did = whoIs.Owner
		msg.DidDocument = whoIs.DidDocument

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// Mock Credential object from webauthn test bench https://github.com/psteniusubi/webauthn-tester
func CreateMockCredential() (*did.Credential, error) {

	return &did.Credential{
		ID:              []byte("ktIQAlFosR9OMGnyJnGthmKcIodPb323F3UqPVe9kvB-eOYrE-pNchsSuiN4ZE0ICyAaRiCb6vfF-7Y5nrvcoD-D42KQsXzhJd14ciqzibA"),
		AttestationType: "platform",
		PublicKey:       []byte("o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YVjULNeTz6C0GMu_DqhSIoYH2el7Mz1NsKQQF3Zq9ruMdVFFAAAAAK3OAAI1vMYKZIsLJfHwVQMAUJLSEAJRaLEfTjBp8iZxrYZinCKHT299txd1Kj1XvZLwfnjmKxPqTXIbErojeGRNCAsgGkYgm"),
		Authenticator: did.Authenticator{
			AAGUID:    []byte("eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoiOHhBM2t3dUVCM0xtc2UxMkJyT2FrSDlZUWlrIiwib3JpZ2luIjoiaHR0cHM6Ly9wc3Rlbml1c3ViaS5naXRodWIuaW8iLCJjcm9zc09yaWdpbiI6ZmFsc2V9"),
			SignCount: 1,
		},
	}, nil
}

// CreateMockDidDocument creates a mock did document for testing
func CreateMockDidDocument(simAccount simtypes.Account) (*did.Document, error) {
	//webauthncred := CreateMockCredential()
	// idStr := "ktIQAlFosR9OMGnyJnGthmKcIodPb323F3UqPVe9kvB-eOYrE-pNchsSuiN4ZE0ICyAaRiCb6vfF-7Y5nrvcoD-D42KQsXzhJd14ciqzibA"
	pubKey, _, err := ed25519.GenerateKey(cryptrand.Reader)
	if err != nil {
		return nil, err
	}

	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", simAccount.Address.String()))
	if err != nil {
		return nil, err
	}
	didController, err := did.ParseDID(fmt.Sprintf("did:snr:%s#test", simAccount.Address.String()))
	if err != nil {
		return nil, err
	}

	vm, err := did.NewVerificationMethod(*didUrl, ssi.JsonWebKey2020, *didController, pubKey)
	if err != nil {
		return nil, err
	}
	ctxUri, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
	if err != nil {
		return nil, err
	}

	doc := did.Document{
		ID:      *didUrl,
		Context: []ssi.URI{*ctxUri},
	}
	doc.AddAuthenticationMethod(vm)
	return &doc, nil
}
