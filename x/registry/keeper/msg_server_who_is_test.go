package keeper_test

import (
	"crypto/ed25519"
	cryptrand "crypto/rand"
	"fmt"
	"strings"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/x/registry/types"
)

func TestWhoIsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	owner := "A"
	doc, _ := CreateMockDidDocument(owner)
	encoded_doc, _ := doc.MarshalJSON()
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateWhoIs(ctx, &types.MsgCreateWhoIs{Creator: owner, DidDocument: encoded_doc})
		require.NoError(t, err)
		whoIs := resp.GetWhoIs()
		require.NotNil(t, whoIs)
	}
}

func TestWhoIsMsgServerUpdate(t *testing.T) {
	owner := "A"
	doc, _ := CreateMockDidDocument(owner)
	encoded_doc, _ := doc.MarshalJSON()
	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWhoIs
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateWhoIs{
				Creator:     owner,
				DidDocument: encoded_doc,
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateWhoIs(ctx, &types.MsgCreateWhoIs{
				Creator:     owner,
				DidDocument: encoded_doc,
			})
			require.NoError(t, err)

			_, err = srv.UpdateWhoIs(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWhoIsMsgServerDelete(t *testing.T) {
	owner := "A"
	doc, _ := CreateMockDidDocument(owner)
	encoded_doc, _ := doc.MarshalJSON()
	for _, tc := range []struct {
		desc    string
		request *types.MsgDeactivateWhoIs
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeactivateWhoIs{Creator: owner},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeactivateWhoIs{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateWhoIs(ctx, &types.MsgCreateWhoIs{Creator: owner, DidDocument: encoded_doc})
			require.NoError(t, err)
			_, err = srv.DeactivateWhoIs(ctx, tc.request)
			if tc.err != nil {
				require.NotNil(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// CreateMockDidDocument creates a mock did document for testing
func CreateMockDidDocument(acct string) (did.Document, error) {
	rawCreator := acct

	// Trim snr account prefix
	if strings.HasPrefix(rawCreator, "snr") {
		rawCreator = strings.TrimLeft(rawCreator, "snr")
	}

	// Trim cosmos account prefix
	if strings.HasPrefix(rawCreator, "cosmos") {
		rawCreator = strings.TrimLeft(rawCreator, "cosmos")
	}

	// UnmarshalJSON from DID document
	doc, err := did.NewDocument(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return nil, err
	}

	//webauthncred := CreateMockCredential()
	pubKey, _, err := ed25519.GenerateKey(cryptrand.Reader)
	if err != nil {
		return nil, err
	}

	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return nil, err
	}
	didController, err := did.ParseDID(fmt.Sprintf("did:snr:%s#test", rawCreator))
	if err != nil {
		return nil, err
	}

	vm, err := did.NewVerificationMethod(*didUrl, ssi.JsonWebKey2020, *didController, pubKey)
	if err != nil {
		return nil, err
	}
	doc.AddAuthenticationMethod(vm)
	return doc, nil
}
