package controller

import (

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"context"
	"fmt"

	"github.com/highlight/highlight/sdk/highlight-go"
	"github.com/sonrhq/core/pkg/did/method/authr"
	"github.com/sonrhq/core/pkg/did/method/sonr"
	"github.com/sonrhq/core/pkg/did/types"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	"go.opentelemetry.io/otel/attribute"
)

var kDefaultMethod = types.DIDMethod("idxr")

// SonrController is a controller for the Sonr DID method
type SonrController struct {
	ID     types.DIDIdentifier
	Method types.DIDMethod

	primary       *sonr.Account
	Authenticator *authr.Authenticator
	account       *types.ControllerAccount
	email         string
	emails        types.DIDAccumulator
	credentials   types.DIDAccumulator
}

// New creates a new identifier for the IDX controller
func New(email string, cred *types.Credential, origin string) (*SonrController, error) {
	ctx := context.Background()
	auth, err := authr.NewAuthenticator(email, cred, origin)
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email), attribute.String("origin", origin))
		return nil, err
	}
	sec, err := auth.DIDSecretKey(email)
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email), attribute.String("origin", origin))
		return nil, err
	}
	primary, err := sonr.NewSonrAccount(sec)
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email), attribute.String("origin", origin))
		return nil, err
	}
	pubKey, err := primary.PublicKey()
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email), attribute.String("origin", origin))
		return nil, err
	}
	emailAcc := types.NewAccumulator(primary.ID, "emails")
	err = emailAcc.Add(email, sec)
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email), attribute.String("origin", origin))
		return nil, err
	}
	idx := &SonrController{
		ID:            primary.ID,
		Method:        kDefaultMethod,
		primary:       primary,
		Authenticator: auth,
		email:         email,
		emails:        types.NewAccumulator(primary.ID, "emails"),
		account: &types.ControllerAccount{
			Address:        primary.Address(),
			Authenticators: []string{auth.DID()},
			PublicKey:      pubKey.String(),
			Wallets:        []string{},
		},
	}
	return idx, nil
}

// AuthorizeIdentity resolves a DID to an IDX controller
func AuthorizeIdentity(email string, controllerAcc *identitytypes.ControllerAccount) (*SonrController, error) {
	m := kDefaultMethod
	id := types.DIDIdentifier(controllerAcc.Address)
	ctx := context.Background()
	if !m.Equals(kDefaultMethod) {
		return nil, fmt.Errorf("invalid method: %s, expected 'idxr'", m)
	}
	authnr, err := authr.ResolveAuthenticator(controllerAcc.Authenticators[0])
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email))
		return nil, err
	}
	secKey, err := authnr.DIDSecretKey(email)
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email))
		return nil, err
	}
	primary, err := sonr.ResolveAccount(controllerAcc.Address, secKey)
	if err != nil {
		highlight.RecordError(ctx, err, attribute.String("email", email))
		return nil, err
	}
	idx := &SonrController{
		ID:            id,
		Method:        m,
		Authenticator: authnr,
		primary:       primary,
		account:       controllerAcc,
		email:         email,
	}
	return idx, nil
}
