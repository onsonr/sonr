package didcontroller

import (
	"fmt"

	"github.com/sonrhq/sonr/pkg/didcommon"
	"github.com/sonrhq/sonr/pkg/didwallets"
	"github.com/sonrhq/sonr/pkg/zkaccumulator"
	identitytypes "github.com/sonrhq/sonr/x/identity/types"
)

var defaultDidMethod = didcommon.Method("idxr")

// SonrController is a controller for the Sonr DID method
type SonrController struct {
	ID     didcommon.Identifier
	Method didcommon.Method

	primary       *didwallets.SonrAccount
	Authenticator *Authenticator
	account       *didcommon.ControllerAccount
	email         string
	emails        zkaccumulator.DIDAccumulator
	credentials   zkaccumulator.DIDAccumulator
}

// New creates a new identifier for the IDX controller
func New(email string, cred *didcommon.Credential, origin string) (*SonrController, error) {
	auth, err := NewAuthenticator(email, cred, origin)
	if err != nil {
		return nil, err
	}
	sec, err := auth.DIDSecretKey(email)
	if err != nil {
		return nil, err
	}
	primary, err := didwallets.NewSonrAccount(sec)
	if err != nil {
		return nil, err
	}
	pubKey, err := primary.PublicKey()
	if err != nil {
		return nil, err
	}
	emailAcc := zkaccumulator.NewAccumulator(primary.ID, "emails")
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
		emails:        zkaccumulator.NewAccumulator(primary.ID, "emails"),
		account: &didcommon.ControllerAccount{
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
	id := didcommon.Identifier(controllerAcc.Address)
	if !m.Equals(defaultDidMethod) {
		return nil, fmt.Errorf("invalid method: %s, expected 'idxr'", m)
	}
	authnr, err := ResolveAuthenticator(controllerAcc.Authenticators[0])
	if err != nil {
		return nil, err
	}
	secKey, err := authnr.DIDSecretKey(email)
	if err != nil {
		return nil, err
	}
	primary, err := didwallets.ResolveSonrAccount(controllerAcc.Address, secKey)
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
