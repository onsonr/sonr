package controller

import (

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"fmt"

	"sonr.io/core/services/did/method/authr"
	"sonr.io/core/services/did/method/sonr"
	"sonr.io/core/services/did/types"
	identitytypes "sonr.io/core/x/identity/types"
)

var defaultDidMethod = types.DIDMethod("idxr")

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
	auth, err := authr.NewAuthenticator(email, cred, origin)
	if err != nil {
		return nil, err
	}
	sec, err := auth.DIDSecretKey(email)
	if err != nil {
		return nil, err
	}
	primary, err := sonr.NewSonrAccount(sec)
	if err != nil {
		return nil, err
	}
	pubKey, err := primary.PublicKey()
	if err != nil {
		return nil, err
	}
	emailAcc := types.NewAccumulator(primary.ID, "emails")
	err = emailAcc.Add(email, sec)
	if err != nil {
		return nil, err
	}
	idx := &SonrController{
		ID:            primary.ID,
		Method:        defaultDidMethod,
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
	m := defaultDidMethod
	id := types.DIDIdentifier(controllerAcc.Address)
	if !m.Equals(defaultDidMethod) {
		return nil, fmt.Errorf("invalid method: %s, expected 'idxr'", m)
	}
	authnr, err := authr.ResolveAuthenticator(controllerAcc.Authenticators[0])
	if err != nil {
		return nil, err
	}
	secKey, err := authnr.DIDSecretKey(email)
	if err != nil {
		return nil, err
	}
	primary, err := sonr.ResolveAccount(controllerAcc.Address, secKey)
	if err != nil {
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
